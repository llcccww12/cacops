package llm_chat

import (
	"code.gitea.io/gitea/models"
	baiduAPI "code.gitea.io/gitea/modules/baiduai"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	llmService "code.gitea.io/gitea/services/llm_chat"
	"net/http"
	"strings"
	"time"
)

const (
	Chatglm     = "chatglm2-6b"
	Llama       = "llama2-7b-chat-hf"
	FlagExpired = "<expired>"
)

func chatPreCheck(ctx *context.APIContext, isPrompt bool) *models.LlmChatVisit {
	modelName := ctx.Query("model_name")
	if modelName == "" {
		ctx.Error(http.StatusBadRequest, "model_name can't be empty", "model_name can't be empty")
		return nil
	}
	userID := ctx.User.ID
	currentTime := time.Now()
	hasChat, _ := models.QueryRunningChat(userID, modelName, currentTime.Unix())

	if hasChat == nil {
		if isPrompt {
			ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
			ctx.Resp.Write([]byte(FlagExpired))
			ctx.Resp.Flush()
			return nil
		} else {
			ctx.JSON(http.StatusOK, map[string]string{
				"code": "-1",
				"msg":  ctx.Tr("llm_chat.chat_expired"),
			})
			log.Error("userID %d : no running chat session for model  %s.", ctx.User.ID, modelName)
			return nil
		}
	}

	counts := models.QueryChatCount(ctx.User.ID, modelName)
	maxTires := setting.LLM_CHAT_API.MAX_FREE_TRIES
	if counts >= maxTires {
		ctx.JSON(http.StatusOK, map[string]string{
			"code": "-1",
			"msg":  ctx.Tr("llm_chat.max_free_exceed"),
		})
		log.Error("userID %d : max free times exceed %d.", ctx.User.ID, maxTires)
		return nil
	}

	return hasChat
}

func promptPreCheck(ctx *context.APIContext, prompt string, modelName string) bool {
	queryLen := len(strings.TrimSpace(prompt))
	log.Info("query length check: %d tokens\n", queryLen)
	if queryLen == 0 {
		ctx.JSON(http.StatusOK, map[string]string{
			"code": "-1",
			"msg":  ctx.Tr("llm_chat.query_empty"),
		})
		log.Error("userID %d : query can't be empty.", ctx.User.ID)
		return false
	}

	lenFlag := false
	if modelName == Chatglm {
		lenFlag = queryLen > setting.LLM_CHAT_API.CHATGLM2_MAX_LENGTH
	} else if modelName == Llama {
		lenFlag = queryLen > setting.LLM_CHAT_API.LLAMA2_MAX_LENGTH
	}
	if lenFlag {
		ctx.JSON(http.StatusOK, map[string]string{
			"code": "-1",
			"msg":  ctx.Tr("llm_chat.query_too_long"),
		})
		log.Error("userID %d : query length too long.", ctx.User.ID)
		return false
	}

	return true
}

func LLMChat(ctx *context.APIContext, data api.LLMChatMessage) {
	log.Info("LLM chat by api.")
	hasChat := chatPreCheck(ctx, true)
	promptFlag := promptPreCheck(ctx, data.Query, data.ModelName)
	if !promptFlag || hasChat == nil {
		log.Error("userID %d : chat prompt pre-check failed.", ctx.User.ID)
		return
	}
	if data.Stream {
		llmService.StreamLLMChatService(ctx.Context, data, hasChat)
	} else {
		ctx.JSON(http.StatusInternalServerError, "currently not supported")
		//llmService.LLMChatService(ctx.Context, data, hasChat)
	}
}

func KBChat(ctx *context.APIContext, data api.KBChatMessage) {
	log.Info("LLM KnowledgeBase chat by api.")
	hasChat := chatPreCheck(ctx, true)
	promptFlag := promptPreCheck(ctx, data.Query, data.ModelName)
	if !promptFlag || hasChat == nil {
		log.Error("userID %d : chat prompt pre-check failed.", ctx.User.ID)
		return
	}
	if data.Stream {
		llmService.StreamKBChatService(ctx.Context, data, hasChat)
	} else {
		ctx.JSON(http.StatusInternalServerError, "currently not supported")
		//llmService.KBChatService(ctx.Context, data, hasChat)
	}
}

func ListKnowledgeBase(ctx *context.APIContext) {
	log.Info("LLM list KnowledgeBase by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.ListKnowledgeBaseService(ctx.Context)
}

func CreateKnowledgeBase(ctx *context.APIContext, data api.CreateKnowledgeBaseParams) {
	log.Info("LLM create KnowledgeBase by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.CreateKnowledgeBaseService(ctx.Context, data)
}

func DeleteKnowledgeBase(ctx *context.APIContext) {
	log.Info("LLM delete KnowledgeBase by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.DeleteKnowledgeBaseService(ctx.Context)
}

func ListFiles(ctx *context.APIContext) {
	log.Info("LLM list files by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.ListFilesService(ctx.Context)
}

func SearchDoc(ctx *context.APIContext, data api.SearchDocParams) {
	log.Info("LLM search doc by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.SearchDocService(ctx.Context, data)
}

func DeleteDoc(ctx *context.APIContext, data api.DeleteDocParams) {
	log.Info("LLM delete doc by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.DeleteDocService(ctx.Context, data)

}

func UpdateDoc(ctx *context.APIContext) {
	log.Info("LLM update doc by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.UpdateDocService(ctx.Context)
}

func RecreateVectorStore(ctx *context.APIContext) {
	log.Info("LLM recreate vector store by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.RecreateVectorStoreService(ctx.Context)
}

func UploadDocUrl(ctx *context.APIContext) {
	log.Info("LLM upload doc by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.UploadDocUrlService(ctx.Context)
}

func UploadDoc(ctx *context.APIContext, form api.LLMChatUploadForm) {
	log.Info("LLM upload doc by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.UploadDocService(ctx.Context, form)
}

func DownloadDoc(ctx *context.APIContext) {
	log.Info("LLM download doc by api.")
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	llmService.DownloadDocService(ctx.Context)
}

func GetFreeTries(ctx *context.APIContext) {
	log.Info("LLM get free tries by api.")
	llmService.GetFreeTriesService(ctx.Context)
}

func LegalText(ctx *context.APIContext, data api.LegalTextParams) {
	log.Info("LLM get chat counts by api.")
	res, err := baiduAPI.CheckLegalText(data.Text)
	if err != nil {
		log.Error("CheckLegalText failed: %s", err)
		ctx.Error(http.StatusInternalServerError, "CheckLegalText failed", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func NewVisit(ctx *context.APIContext) {
	log.Info("LLM new visit by api.")
	currentTime := time.Now()
	modelName := ctx.Query("model_name")

	hasChat, _ := models.QueryRunningChat(ctx.User.ID, modelName, currentTime.Unix())
	if hasChat != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"chat_id":    hasChat.ChatId,
			"model_name": hasChat.ModelName,
			"expired_at": hasChat.ExpiredTime,
		})
		log.Info("hasChat %s, expired at %s\n", hasChat.ChatId, hasChat.ExpiredTime)
		return
	}

	chatID := ctx.User.Name + currentTime.Format("20060102150405")
	duration := time.Duration(setting.LLM_CHAT_API.CHAT_EXPIRED_MINUTES)
	endTime := currentTime.Add(time.Minute * duration)
	entTimeStr := endTime.Format("2006-01-02 15:04:05")
	llmChatVisit := &models.LlmChatVisit{
		UserId:      ctx.User.ID,
		ChatId:      chatID,
		ModelName:   modelName,
		ExpiredUnix: endTime.Unix(),
		ExpiredTime: entTimeStr,
		Agreement:   0,
	}
	models.SaveVisit(llmChatVisit)
	log.Info("new chat %s, expired at %s\n", chatID, entTimeStr)
	ctx.JSON(http.StatusOK, map[string]string{
		"chat_id":    llmChatVisit.ChatId,
		"model_name": llmChatVisit.ModelName,
		"expired_at": llmChatVisit.ExpiredTime,
	})
}

func SaveAgreement(ctx *context.APIContext) {
	hasChat := chatPreCheck(ctx, false)
	if hasChat == nil {
		return
	}
	if hasChat.Agreement == 1 {
		ctx.JSON(http.StatusOK, map[string]string{
			"code": "-1",
			"msg":  "already saved Agreement",
		})
		return
	}

	hasChat.Agreement = 1
	models.UpdateChat(hasChat)
	ctx.JSON(http.StatusOK, map[string]string{
		"code": "1",
		"msg":  "successfully saved Agreement status",
	})
}

func GetChatStats(ctx *context.APIContext) {
	log.Info("LLM chat stats by api.")

	resChat, err := models.QueryChatStatistics()
	if err != nil {
		log.Error("QueryChatStatistics failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, "QueryChatStatistics failed")
		return
	}

	resVisit, err := models.QueryChatVisitStatistics()
	if err != nil {
		log.Error("QueryChatVisitStatistics failed: %s", err)
		ctx.JSON(http.StatusInternalServerError, "QueryChatVisitStatistics failed")
		return
	}

	res := make(map[string]interface{})
	for _, chat := range resChat {
		for _, visit := range resVisit {
			if chat["model_name"] == visit["model_name"] {
				chat["visit"] = visit["visit"]
				chat["visit_user"] = visit["visit_user"]
			}
		}
		res[chat["model_name"].(string)] = chat
	}

	ctx.JSON(http.StatusOK, res)
}
