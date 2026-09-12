package llm_chat

import (
	"code.gitea.io/gitea/models"
	baiduAPI "code.gitea.io/gitea/modules/baiduai"
	"code.gitea.io/gitea/modules/context"
	llmChatAPI "code.gitea.io/gitea/modules/llm_chat"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	VectorStoreType   = "faiss"
	EmbeddingModel    = "m3e-base"
	TopK              = 5
	ScoreThreshold    = 0.5
	DeleteContent     = true
	NotRefreshVsCache = false
	FlagTextInvalid   = "<illegal>"
	FlagAccountBanned = "<banned>"
	FlagTextDoc       = "<docs>"
	ValidationTool    = "baidu_api"
)

func getKnowledgeBaseName(ctx *context.Context) string {
	kbName := ctx.Query("knowledge_base_name")
	userID := strconv.FormatInt(ctx.User.ID, 10)
	if kbName != setting.LLM_CHAT_API.COMMON_KB {
		return userID + "_" + kbName
	}
	return kbName
}

func LLMChatService(ctx *context.Context, data api.LLMChatMessage, chat *models.LlmChatVisit) {
	log.Info("received by api %+v", data)
	res, err := llmChatAPI.SendLLMChat(data)
	if err != nil {
		log.Error("LLMChatService failed: %s", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
	}
	ctx.JSON(http.StatusOK, res)
}

func KBChatService(ctx *context.Context, data api.KBChatMessage, chat *models.LlmChatVisit) {
	if data.TopK == 0 || data.ScoreThreshold == 0 {
		data.TopK = TopK
		data.ScoreThreshold = ScoreThreshold
	}
	log.Info("received by api %+v", data)
	res, err := llmChatAPI.SendKBChat(data)
	if err != nil {
		log.Error("KnowledgeBaseChatService failed: %s", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
	}
	log.Info("received by resty %+v", res)
	ctx.JSON(http.StatusOK, res)
}

func isInvalidQuery(ctx *context.Context, chat *models.LlmChat, queryType string) bool {
	var query string
	if queryType == "prompt" {
		query = chat.Prompt
	} else {
		query = chat.Answer
	}

	if query == "" {
		return false
	}

	chat.InvalidCount = 0
	chat.InvalidTool = ValidationTool
	res, err := baiduAPI.CheckLegalText(query)
	if err != nil {
		log.Error("isInvalidQuery() failed: %s", err)
		return false
	}
	if res.ConclusionType != 1 {
		chat.InvalidCount = 1
		chat.InvalidType = queryType
		jsonRes, _ := json.Marshal(res)
		chat.InvalidDetail = string(jsonRes)
		err := models.SaveChat(chat)
		if err != nil {
			log.Error("isInvalidQuery() SaveChat failed: %s", err)
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
			return true
		}

		invalidTotal := models.QueryInvalidPromptCount(ctx.User.ID)
		if invalidTotal >= setting.LLM_CHAT_API.LEGAL_MAX_COUNT {
			log.Info("isInvalidQuery() invalid total reach max: %d\n", invalidTotal)
			ctx.User.ProhibitLogin = true
			models.UpdateUserCols(ctx.User, "prohibit_login")
			ctx.Resp.Write([]byte(FlagAccountBanned))
			ctx.Resp.Flush()
		} else {
			ctx.Resp.Write([]byte(FlagTextInvalid))
			ctx.Resp.Flush()
		}
		return true
	}
	return false
}

func StreamLLMChatService(ctx *context.Context, data api.LLMChatMessage, chat *models.LlmChatVisit) {
	uuid := uuid.NewV4()
	id := uuid.String()
	llmChat := &models.LlmChat{
		ID:         id,
		UserId:     ctx.User.ID,
		ChatId:     chat.ChatId,
		Prompt:     data.Query,
		ModelName:  data.ModelName,
		Endpoint:   llmChatAPI.GetEndpoint(data.ModelName),
		ChatType:   "llm",
		ChatStatus: 1,
		Count:      1,
	}

	var answer string
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")

	//call baiduai api to check legality of query
	if setting.LLM_CHAT_API.LEGAL_CHECK {
		invalidPrompt := isInvalidQuery(ctx, llmChat, "prompt")
		if invalidPrompt {
			log.Info("StreamLLMChatService() invalid prompt: %s\n", llmChat.Prompt)
			return
		}
	}

	resultChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})
	go llmChatAPI.StreamLLMChat(data, resultChan, errChan, done)

	for {
		select {
		case data := <-resultChan:
			answer += data
			ctx.Resp.Write([]byte(data))
			ctx.Resp.Flush()
		case err := <-errChan:
			response := ctx.Tr("llm_chat.server_error")
			for _, v := range response {
				ctx.Resp.Write([]byte(string(v)))
				ctx.Resp.Flush()
				time.Sleep(50 * time.Millisecond)
			}
			log.Error("StreamLLMChatService() failed: %s", err)
			log.Info("StreamLLMChatService() chat server api error, save to db")
			llmChat.ChatStatus = 0
			err = models.SaveChat(llmChat)
			if err != nil {
				log.Error("StreamLLMChatService() SaveChat failed: %s", err)
				ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
			}
			close(resultChan)
			close(errChan)
			close(done)
			return
		case <-done:
			llmChat.Answer = answer
			if llmChat.Answer == "" {
				llmChat.ChatStatus = 0
			}
			if setting.LLM_CHAT_API.LEGAL_CHECK {
				invalidAnswer := isInvalidQuery(ctx, llmChat, "answer")
				if invalidAnswer {
					log.Info("StreamLLMChatService() invalid answer: %s\n", llmChat.Answer)
					close(resultChan)
					close(errChan)
					return
				}
			}
			log.Info("StreamLLMChatService() nothing invalid, save to db")
			err := models.SaveChat(llmChat)
			if err != nil {
				log.Error("StreamLLMChatService() SaveChat failed: %s", err)
				ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
			}
			close(resultChan)
			close(errChan)
			return
		}
	}
}

func StreamKBChatService(ctx *context.Context, data api.KBChatMessage, chat *models.LlmChatVisit) {
	userID := strconv.FormatInt(ctx.User.ID, 10)
	if data.KnowledgeBaseName != setting.LLM_CHAT_API.COMMON_KB {
		data.KnowledgeBaseName = userID + "_" + data.KnowledgeBaseName
	}

	uuid := uuid.NewV4()
	id := uuid.String()
	llmChat := &models.LlmChat{
		ID:                id,
		UserId:            ctx.User.ID,
		ChatId:            chat.ChatId,
		Prompt:            data.Query,
		ModelName:         data.ModelName,
		Endpoint:          llmChatAPI.GetEndpoint(data.ModelName),
		KnowledgeBaseName: data.KnowledgeBaseName,
		VectorStoreType:   VectorStoreType,
		EmbeddingModel:    EmbeddingModel,
		ChatType:          "kb",
		ChatStatus:        1,
		Count:             1,
	}

	var answer string
	var docs string
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")

	//call baiduai api to check legality of query
	if setting.LLM_CHAT_API.LEGAL_CHECK {
		invalidPrompt := isInvalidQuery(ctx, llmChat, "prompt")
		if invalidPrompt {
			log.Info("StreamKBChatService() invalid prompt: %s\n", llmChat.Prompt)
			return
		}
	}

	resultChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})
	go llmChatAPI.StreamKBChat(data, resultChan, errChan, done)

	for {
		select {
		case data := <-resultChan:
			if strings.Contains(data, "answer") {
				var result api.KBChatAnswer
				json.Unmarshal([]byte(data), &result)
				//ctx.JSON(http.StatusOK, result)
				ctx.Resp.Write([]byte(result.Answer))
				ctx.Resp.Flush()
				answer += result.Answer
			}
			if strings.Contains(data, "docs") {
				docs += data
			}
		case err := <-errChan:
			response := ctx.Tr("llm_chat.server_error")
			for _, v := range response {
				ctx.Resp.Write([]byte(string(v)))
				ctx.Resp.Flush()
				time.Sleep(50 * time.Millisecond)
			}
			log.Error("StreamKBChatService() failed: %s", err)
			log.Info("StreamKBChatService() chat server api error, save to db")
			llmChat.ChatStatus = 0
			err = models.SaveChat(llmChat)
			if err != nil {
				log.Error("SaveChat failed: %s", err)
				ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
			}
			close(resultChan)
			close(errChan)
			close(done)
			return
		case <-done:
			if docs != "" {
				ctx.Resp.Write([]byte(FlagTextDoc + docs))
				ctx.Resp.Flush()
			}
			llmChat.Answer = answer
			if llmChat.Answer == "" {
				llmChat.ChatStatus = 0
			}
			//call baiduai api to check legality of query
			if setting.LLM_CHAT_API.LEGAL_CHECK {
				invalidAnswer := isInvalidQuery(ctx, llmChat, "answer")
				if invalidAnswer {
					log.Info("StreamKBChatService() invalid answer: %s\n", llmChat.Answer)
					close(resultChan)
					close(errChan)
					return
				}
			}
			log.Info("StreamKBChatService() nothing invalid, save to db")
			err := models.SaveChat(llmChat)
			if err != nil {
				log.Error("StreamKBChatService() SaveChat failed: %s", err)
				ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
			}
			close(resultChan)
			close(errChan)
			return
		}
	}
}

func ListKnowledgeBaseService(ctx *context.Context) {
	userID := strconv.FormatInt(ctx.User.ID, 10)
	res, err := llmChatAPI.ListKnowledgeBase()
	if err != nil {
		log.Error("LLMChatService failed: %s", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	log.Info("received by resty %+v", res)
	log.Info("######## received by resty %+v\n", res)
	var realKB []string
	for i := len(res.Data) - 1; i >= 0; i-- {
		v := res.Data[i]
		if strings.Contains(v, userID) {
			substr := strings.Replace(v, userID+"_", "", -1)
			if strings.TrimSpace(substr) != "" {
				realKB = append(realKB, substr)
			}
		}
	}

	if setting.LLM_CHAT_API.COMMON_KB != "" {
		realKB = append(realKB, setting.LLM_CHAT_API.COMMON_KB)
	}

	realData := llmChatAPI.LLMBasicMsgWithData{
		Code: res.Code,
		Msg:  res.Msg,
		Data: realKB,
	}
	log.Info("######## sent %+v\n", realData)
	ctx.JSON(http.StatusOK, realData)
}

func CreateKnowledgeBaseService(ctx *context.Context, data api.CreateKnowledgeBaseParams) {
	userID := strconv.FormatInt(ctx.User.ID, 10)
	realKB := userID + "_" + data.KnowledgeBaseName
	params := api.CreateKnowledgeBaseParams{
		KnowledgeBaseName: realKB,
		VectorStoreType:   VectorStoreType,
		EmbedModel:        EmbeddingModel,
	}
	log.Info("received by api %+v\n", params)
	res, err := llmChatAPI.CreateKnowledgeBase(params)
	if err != nil {
		log.Error("KnowledgeBaseChatService failed: %s", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	log.Info("#####  received by resty %+v\n", res)
	if strings.Contains(res.Msg, userID) {
		substr := strings.Replace(res.Msg, userID+"_", "", -1)
		if strings.TrimSpace(substr) != "" {
			res.Msg = substr
		}
	}
	log.Info("#####  sent to client %+v\n", res)
	ctx.JSON(http.StatusOK, res)
}

func DeleteKnowledgeBaseService(ctx *context.Context) {
	userID := strconv.FormatInt(ctx.User.ID, 10)
	data := getKnowledgeBaseName(ctx)
	log.Info("### received by api %+v\n", data)
	if data == setting.LLM_CHAT_API.COMMON_KB {
		ctx.Error(http.StatusForbidden, "You can't operate %s", data)
		return
	}
	res, err := llmChatAPI.DeleteKnowledgeBase(data)
	if err != nil {
		log.Error("DeleteKnowledgeBaseService failed: %s", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	log.Info("######## received by resty %+v\n", res)
	if strings.Contains(res.Msg, userID) {
		substr := strings.Replace(res.Msg, userID+"_", "", -1)
		if strings.TrimSpace(substr) != "" {
			res.Msg = substr
		}
	}
	log.Info("######## sent to client %+v\n", res)
	ctx.JSON(http.StatusOK, res)
}

func ListFilesService(ctx *context.Context) {
	data := getKnowledgeBaseName(ctx)
	log.Info("received by api %+v", data)
	res, err := llmChatAPI.KBListFiles(data)
	if err != nil {
		log.Error("ListFiles failed: %s", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	log.Info("received by resty %+v", res)
	ctx.JSON(http.StatusOK, res)
}

func SearchDocService(ctx *context.Context, data api.SearchDocParams) {
	if data.TopK == 0 || data.ScoreThreshold == 0 {
		data.TopK = TopK
		data.ScoreThreshold = ScoreThreshold
	}
	realKB := getKnowledgeBaseName(ctx)
	data.KnowledgeBaseName = realKB
	log.Info("received by api %+v", data)
	res, err := llmChatAPI.KBSearchDoc(data)
	if err != nil {
		log.Error("SearchDocService failed: %s", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	log.Info("received by resty %+v", res)
	ctx.JSON(http.StatusOK, res)
}

func DeleteDocService(ctx *context.Context, data api.DeleteDocParams) {
	data.DeleteContent = DeleteContent
	data.NotRefreshVsCache = NotRefreshVsCache
	userID := strconv.FormatInt(ctx.User.ID, 10)
	realKB := userID + "_" + data.KnowledgeBaseName
	data.KnowledgeBaseName = realKB
	log.Info("received by api %+v", data)
	if data.KnowledgeBaseName == setting.LLM_CHAT_API.COMMON_KB {
		ctx.Error(http.StatusForbidden, "You can't operate %s", data.KnowledgeBaseName)
		return
	}
	res, err := llmChatAPI.KBDeleteDoc(data)
	if err != nil {
		log.Error("LLMChatService failed: %s", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func UpdateDocService(ctx *context.Context) {
	data := api.UpdateDocParams{
		KnowledgeBaseName: getKnowledgeBaseName(ctx),
		FileNames:         ctx.Query("file_name"),
		NotRefreshVsCache: NotRefreshVsCache,
	}
	log.Info("received by api %+v", data)
	if data.KnowledgeBaseName == setting.LLM_CHAT_API.COMMON_KB {
		ctx.Error(http.StatusForbidden, "You can't operate %s", data.KnowledgeBaseName)
		return
	}
	res, err := llmChatAPI.KBUpdateDoc(data)
	if err != nil {
		log.Error("LLMChatService failed: %s", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func RecreateVectorStoreService(ctx *context.Context) {
	data := api.RecreateVectorStoreParams{
		KnowledgeBaseName: getKnowledgeBaseName(ctx),
		AllowEmptyKb:      true,
		VsType:            VectorStoreType,
		EmbedModel:        EmbeddingModel,
	}
	log.Info("received by api %+v", data)
	if data.KnowledgeBaseName == setting.LLM_CHAT_API.COMMON_KB {
		ctx.Error(http.StatusForbidden, "You can't operate %s", data.KnowledgeBaseName)
		return
	}

	resultChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})

	go llmChatAPI.KBRecreateVectorStore(data, resultChan, errChan, done)
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")

	for {
		select {
		case data := <-resultChan:
			_, err := ctx.Resp.Write([]byte(data))
			if err != nil {
				log.Error("Error writing response: %s", err)
				ctx.JSON(http.StatusInternalServerError, models.BaseErrorMessageApi(err.Error()))
				return
			}
			log.Info("%s\n", []byte(data))
			ctx.Resp.Flush() // Flush the response to send it immediately
		case err := <-errChan:
			log.Error("Error writing response: %s", err)
			ctx.JSON(http.StatusInternalServerError, models.BaseErrorMessageApi(err.Error()))
			return
		case <-done:
			close(resultChan)
			close(errChan)
			return
		}
	}
}

func UploadDocUrlService(ctx *context.Context) {
	data := getKnowledgeBaseName(ctx)
	if data == setting.LLM_CHAT_API.COMMON_KB {
		ctx.Error(http.StatusForbidden, "You can't operate %s", data)
		return
	}
	url, _ := llmChatAPI.GetUploadDocUrl()
	log.Info("received by api %+v", url)

	ctx.JSON(http.StatusOK, url)
}

func UploadDocService(ctx *context.Context, form api.LLMChatUploadForm) {
	log.Info("######### received request %+v\n", ctx.Req.Request)
	log.Info("######### form api.LLMChatUploadForm %+v\n", form)
	modelName := ctx.Query("model_name")
	userID := strconv.FormatInt(ctx.User.ID, 10)
	if form.KnowledgeBaseName != setting.LLM_CHAT_API.COMMON_KB {
		form.KnowledgeBaseName = userID + "_" + form.KnowledgeBaseName
	}
	res, err := llmChatAPI.UploadDocs(modelName, form)
	log.Info("######### received by resty %+v\n", res)

	if err != nil {
		log.Error("UploadDocService failed: %s", err)
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func DownloadDocService(ctx *context.Context) {
	data := getKnowledgeBaseName(ctx)
	if data == setting.LLM_CHAT_API.COMMON_KB {
		ctx.Error(http.StatusForbidden, "You can't operate %s", data)
		return
	}
	fileName := ctx.Query("file_name")
	log.Info("received by api knowledgeBaseName:%s, fileName: %s", data, fileName)
	url, _ := llmChatAPI.GetDownloadDocUrl(data, fileName)
	log.Info("received by api %+v", url)
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
}

func GetFreeTriesService(ctx *context.Context) {
	modelName := ctx.Query("model_name")
	maxTries := setting.LLM_CHAT_API.MAX_FREE_TRIES
	counts := models.QueryChatCount(ctx.User.ID, modelName)
	//firstVisit := models.QueryFirstVisit(ctx.User.ID)

	data := api.LLMChatCountsResults{
		MaxTries: maxTries,
		Counts:   counts,
		CanChat:  counts < maxTries,
		//FirstVisit: firstVisit == 0,
	}
	log.Info("user %+v, GetFreeTriesService() data= %+v", ctx, data)
	ctx.JSON(http.StatusOK, data)
}
