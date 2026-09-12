package llm_chat

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"code.gitea.io/gitea/manager/client/grampus"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/ai_task"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

const (
	CHAT_TYPE           = "experience"
	CHAT_TYPE_MODELARTS = "modelarts"
)

var restyClient *resty.Client

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

type ChatConfig struct {
	ChatId            string        `json:"chat_id" binding:"Required"`
	TaskId            int           `json:"task_id"`
	Message           string        `json:"message" binding:"Required"`
	SystemMessage     string        `json:"system_message,omitempty"`
	MaxTokens         int           `json:"max_tokens,omitempty"`
	Temperature       float64       `json:"temperature,omitempty"`
	TopP              float64       `json:"top_p,omitempty"`
	RepetitionPenalty float64       `json:"repetition_penalty,omitempty"`
	History           []ChatHistory `json:"history,omitempty" `
}

type ChatHistory struct {
	UserMessage string `json:"user_message"`
	BotResponse string `json:"bot_response"`
}

func scanByCharacter(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	return 1, data[:1], nil
}

func ChatCompletionAPI(ctx *context.APIContext, config ChatConfig) {

	var errMsg string
	cloudbrainId := int64(config.TaskId)
	// Get cloudbrain job for modelId and modelName
	task, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		errMsg = ctx.Tr("llm_chat.chat_expired")
		log.Error("[llm_chat] Error getting cloudbrain by cloudbrainId: %v", err)
		ctx.JSON(503, map[string]string{"error": errMsg})
		return
	}
	if task.IsTerminalOrStopping() {
		errMsg = ctx.Tr("cloudbrain.Already_stopped")
		ctx.JSON(503, map[string]string{"error": errMsg})
		return
	}

	if task.Type == models.TypeIFLYTek {
		IFLYTekChat(ctx, config, task)
		return
	}

	// Get url
	tokenUrl := ""
	if task.ComputeResource != models.NPUResource {
		tokenUrl, _ = ai_task.GetSelfEndPointUrlById(cloudbrainId)
	}
	index := strings.Index(tokenUrl, "?token=")
	if index == -1 {
		log.Info("not found token=")
		tokenUrl = tokenUrl + "/chat"
	} else {
		token := tokenUrl[index:]
		tokenUrl = tokenUrl[0:index] + "/chat" + token
	}
	log.Info("[llm_chat] Url=" + tokenUrl)

	// Create a chat object for save DB
	var chatParams string
	chatParamsJson, err := json.Marshal(models.LlmChatParams{
		SystemMessage:     config.SystemMessage,
		MaxTokens:         config.MaxTokens,
		Temperature:       config.Temperature,
		TopP:              config.TopP,
		RepetitionPenalty: config.RepetitionPenalty,
	})
	if err != nil {
		log.Info("[llm_chat] Error marshaling chatParamsDB: %v", err)
	} else {
		chatParams = string(chatParamsJson)
	}

	messageID := uuid.NewV4().String()
	llmChat := &models.LlmChat{
		ID:        messageID,
		Prompt:    config.Message,
		Params:    chatParams,
		TaskId:    cloudbrainId,
		ChatId:    config.ChatId,
		UserId:    ctx.User.ID,
		ModelId:   task.ModelId,
		ModelName: task.ModelName,
		Endpoint:  tokenUrl,
		ChatType:  CHAT_TYPE,
		Count:     1,
	}

	// Send request to chat service
	var resp *resty.Response
	if task.ComputeResource == models.NPUResource {
		value, _ := redis_client.Get(redis_key.ModelExperienceModelLoadedKey(task.JobID))
		if value == "" {
			errMsg = ctx.Tr("llm_chat.loading")
			ctx.JSON(503, map[string]string{"error": errMsg})
			return
		}

		chatParams := grampus.ChatParameter{
			MaxTokens:         config.MaxTokens,
			Messages:          getChatMessages(config, true),
			RepetitionPenalty: config.RepetitionPenalty,
			Temperature:       config.Temperature,
			TopP:              config.TopP,
		}
		resp, err = grampus.ModelAppInference(task.JobID, chatParams)

	} else {

		client := getRestyClient()
		resp, err = client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(config).
			SetDoNotParseResponse(true).
			Post(tokenUrl)
	}
	if err != nil {
		errMsg = ctx.Tr("llm_chat.chat_expired")
		log.Error("[llm_chat] Error sending request to llm_chat Url: %v", err)
		ctx.JSON(503, map[string]string{"error": errMsg})
		return
	}
	defer resp.RawBody().Close()

	// Check if the status code is 200
	log.Info("[llm_chat] Url resp status: %s\n", resp.Status())
	if resp.StatusCode() != 200 {
		switch resp.StatusCode() {
		case 502:
			errMsg = ctx.Tr("llm_chat.loading")
		case 404:
			errMsg = ctx.Tr("llm_chat.chat_expired")
		default:
			errMsg = ctx.Tr("llm_chat.server_error")
		}
		ctx.JSON(503, map[string]string{"error": errMsg})
		return
	}

	// prepare response headers
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")

	var answer string
	var invalidCharBuffer string

	// read stream and return valid character
	scanner := bufio.NewScanner(resp.RawBody())
	scanner.Split(scanByCharacter)
	for scanner.Scan() {
		char := scanner.Text()
		if len(invalidCharBuffer) > 0 {
			char = invalidCharBuffer + char
			invalidCharBuffer = ""
		}
		if utf8.ValidString(char) {
			//runes := []rune(char)
			//fmt.Printf("%s -> %U \n", char, runes[0])
			// send stream to front end
			_, writeErr := ctx.Resp.Write([]byte(char))
			if writeErr != nil {
				log.Error("[llm_chat] Error flushing to front end: %v", writeErr)
				ctx.ServerError("[llm_chat] Error flushing to front end: ", writeErr)
				return
			}
			ctx.Resp.Flush()
			answer += char
		} else {
			invalidCharBuffer += char
		}
	}

	// left over invalid characters
	if len(invalidCharBuffer) > 0 {
		log.Info("[llm_chat] Unprocessed invalid UTF-8 characters: %s\n", invalidCharBuffer)
	}

	// scanner error
	if scanner.Err() != nil {
		errMsg = ctx.Tr("llm_chat.server_error")
		for _, v := range errMsg {
			ctx.Resp.Write([]byte(string(v)))
			ctx.Resp.Flush()
			//time.Sleep(50 * time.Millisecond)
		}
		log.Error("[llm_chat] scanner .Err(): %s", err)
	}

	// Save chat to database
	llmChat.Answer = answer
	bizErr := models.SaveChat(llmChat)
	if bizErr != nil {
		log.Error("[llm_chat] Error saving chat: %v", bizErr)
		ctx.ServerError("[llm_chat] Error saving chat", bizErr)
		return
	}
}

func IFLYTekChat(ctx *context.APIContext, config ChatConfig, task *models.Cloudbrain) {
	content := getChatMessages(config, false)

	chatReq := models.IFLYTekAPIRequest{
		Model:    task.JobName,
		Messages: content,
	}

	urlStr := setting.IFLYTekConfig.OnlineServiceAPIUrl
	var errMsg string
	var result models.IFLYTekAPIResponse
	// Send request to chat service
	client := getRestyClient()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Consumer-Username", ctx.User.Name).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", setting.IFLYTekConfig.OnlineServiceAPIToken)).
		SetBody(chatReq).
		SetResult(&result).
		Post(urlStr)
	if err != nil {
		errMsg = ctx.Tr("llm_chat.chat_expired")
		log.Error("[llm_chat] Error sending request to llm_chat Url: %v", err)
		ctx.JSON(503, map[string]string{"error": errMsg})
		return
	}
	defer resp.RawBody().Close()

	// Check if the status code is 200
	log.Info("[llm_chat] Url resp status: %s\n", resp.Status())
	if resp.StatusCode() != 200 {
		switch resp.StatusCode() {
		case 502:
			errMsg = ctx.Tr("llm_chat.loading")
		case 404:
			errMsg = ctx.Tr("llm_chat.chat_expired")
		default:
			errMsg = ctx.Tr("llm_chat.server_error")
		}
		ctx.JSON(503, map[string]string{"error": errMsg})
		return
	}

	var resultContent string
	for i := 0; i < len(result.Choices); i++ {
		c := result.Choices[i]
		if c.PayloadText.Role == "assistant" {
			resultContent = resultContent + c.PayloadText.Content
		}
	}
	// prepare response headers
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")

	var answer string
	var invalidCharBuffer string

	// read stream and return valid character
	scanner := bufio.NewScanner(strings.NewReader(resultContent))
	scanner.Split(scanByCharacter)
	for scanner.Scan() {
		char := scanner.Text()
		if len(invalidCharBuffer) > 0 {
			char = invalidCharBuffer + char
			invalidCharBuffer = ""
		}
		if utf8.ValidString(char) {
			//runes := []rune(char)
			//fmt.Printf("%s -> %U \n", char, runes[0])
			// send stream to front end
			_, writeErr := ctx.Resp.Write([]byte(char))
			if writeErr != nil {
				log.Error("[llm_chat] Error flushing to front end: %v", writeErr)
				ctx.ServerError("[llm_chat] Error flushing to front end: ", writeErr)
				return
			}
			ctx.Resp.Flush()
			answer += char
		} else {
			invalidCharBuffer += char
		}
	}

	// left over invalid characters
	if len(invalidCharBuffer) > 0 {
		log.Info("[llm_chat] Unprocessed invalid UTF-8 characters: %s\n", invalidCharBuffer)
	}

	// scanner error
	if scanner.Err() != nil {
		errMsg = ctx.Tr("llm_chat.server_error")
		for _, v := range errMsg {
			ctx.Resp.Write([]byte(string(v)))
			ctx.Resp.Flush()
			//time.Sleep(50 * time.Millisecond)
		}
		log.Error("[llm_chat] scanner .Err(): %s", err)
	}
	// Create a chat object for save DB
	var chatParams string
	chatParamsJson, err := json.Marshal(models.LlmChatParams{
		SystemMessage:     config.SystemMessage,
		MaxTokens:         config.MaxTokens,
		Temperature:       config.Temperature,
		TopP:              config.TopP,
		RepetitionPenalty: config.RepetitionPenalty,
	})
	if err != nil {
		log.Info("[llm_chat] Error marshaling chatParamsDB: %v", err)
	} else {
		chatParams = string(chatParamsJson)
	}

	messageID := uuid.NewV4().String()
	llmChat := &models.LlmChat{
		ID:        messageID,
		Prompt:    config.Message,
		Params:    chatParams,
		TaskId:    task.ID,
		ChatId:    config.ChatId,
		UserId:    ctx.User.ID,
		ModelId:   task.ModelId,
		ModelName: task.ModelName,
		Endpoint:  urlStr,
		ChatType:  CHAT_TYPE,
		Count:     1,
	}

	// Save chat to database
	llmChat.Answer = answer
	bizErr := models.SaveChat(llmChat)
	if bizErr != nil {
		log.Error("[llm_chat] Error saving chat: %v", bizErr)
		ctx.ServerError("[llm_chat] Error saving chat", bizErr)
		return
	}
}

func getChatMessages(config ChatConfig, containsSystemMessage bool) []models.PayloadText {
	content := make([]models.PayloadText, 0)
	if containsSystemMessage {
		systemMsg := models.PayloadText{
			Role:    "system",
			Content: config.SystemMessage,
		}
		content = append(content, systemMsg)
	}

	for i := 0; i < len(config.History); i++ {
		h := config.History[i]
		userMsg := models.PayloadText{
			Role:    "user",
			Content: h.UserMessage,
		}
		content = append(content, userMsg)
		botMsg := models.PayloadText{
			Role:    "assistant",
			Content: h.BotResponse,
		}
		content = append(content, botMsg)
	}
	content = append(content, models.PayloadText{
		Role:    "user",
		Content: config.Message,
	})
	return content
}

func FeedBackDeepSeekAPI(ctx *context.APIContext, config modelarts.ChatFeedBack) {
	response, err := modelarts.SendDeepSeekFeedBack(&config)
	if err != nil {
		log.Error("Error in FeedBackDeepSeekAPI: %v", err)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.server_error"), "server_status_code": "0"})
		return
	}

	if response.StatusCode != 200 {
		log.Error("Error in FeedBackDeepSeekAPI response status code: %d", response.StatusCode)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.server_error"), "server_status_code": fmt.Sprintf("%d", response.StatusCode)})
		return
	}

	ctx.JSON(200, map[string]string{"msg": "success"})
}

func ChatDeepSeekAPINew(ctx *context.APIContext, config modelarts.ChatConfigDeepSeek) {
	response, err := modelarts.CreateDeepSeekJobNew(&config)
	if err != nil {
		log.Error("Error in ChatDeepSeekAPINew: %v", err)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.server_error"), "server_status_code": "0"})
		return
	}

	if response.StatusCode != 200 {
		log.Error("Error in ChatDeepSeekAPINew response status code: %d", response.StatusCode)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.server_error"), "server_status_code": fmt.Sprintf("%d", response.StatusCode)})
		return
	}

	defer response.Body.Close()
	FlushHttpResponse(ctx, response)
}

func FlushHttpResponse(ctx *context.APIContext, response *http.Response) {
	// Prepare response headers
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	ctx.Resp.Header().Set("X-Accel-Buffering", "no")

	var answer string
	var invalidCharBuffer string

	// read stream and return valid character
	scanner := bufio.NewScanner(response.Body)
	scanner.Split(scanByCharacter)
	for scanner.Scan() {
		char := scanner.Text()
		if len(invalidCharBuffer) > 0 {
			char = invalidCharBuffer + char
			invalidCharBuffer = ""
		}
		if utf8.ValidString(char) && char != "" {
			if strings.Contains(char, "None") {
				log.Info("[llm_chat] 'None' found, breaking loop.")
				break
			}
			runes := []rune(char)
			fmt.Printf("%s -> %U \n", char, runes[0])
			//send stream to front end
			_, writeErr := ctx.Resp.Write([]byte(char))
			if writeErr != nil {
				log.Error("[llm_chat] Error flushing to front end: %v", writeErr)
				ctx.ServerError("[llm_chat] Error flushing to front end: ", writeErr)
				return
			}
			ctx.Resp.Flush()
			answer += char
		} else {
			invalidCharBuffer += char
		}
	}

	// left over invalid characters
	if len(invalidCharBuffer) > 0 {
		log.Info("[llm_chat] Unprocessed invalid UTF-8 characters: %s\n", invalidCharBuffer)
	}

	// scanner error
	if scanner.Err() != nil {
		log.Error("[llm_chat] scanner .Err(): %s", scanner.Err())
	}

	defer response.Body.Close()
}
