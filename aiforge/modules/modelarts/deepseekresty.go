package modelarts

import (
	"bytes"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	AUTO_MODEL_CHAT            = "/inference"
	AUTO_MODEL_CHAT_WITH_ID    = "/chat"
	AUTO_MODEL_CHAT_WITH_MODEL = "/model_chat"
	AUTO_MODEL_FEEDBACK        = "/save"
)

type ChatConfigDeepSeek struct {
	Messages []ChatMessages `json:"messages" binding:"Required"`
	Flag     bool           `json:"flag"`
	Id       int            `json:"id,omitempty"`
	Model    int            `json:"model,omitempty"`
	Service  string         `json:"service,omitempty"` // hub,auto

	MaxTokens         int     `json:"max_tokens,omitempty"`
	Temperature       float64 `json:"temperature,omitempty"`
	TopP              float64 `json:"top_p,omitempty"`
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty"`
	ChatId            string  `json:"chat_id,omitempty"`
}

func (r *ChatConfigDeepSeek) Init() {
	if r.Service == "" {
		r.Service = "hub"
	}
}

type ChatMessages struct {
	Role    string `json:"role" binding:"Required"`
	Content string `json:"content" binding:"Required"`
}

type ChatFeedBack struct {
	Data struct {
		Messages string      `json:"messages"`
		Info     string      `json:"info"`
		Feedback interface{} `json:"feedback"`
	} `json:"data"`
}

func CreateDeepSeekJobNew(chatConfig *ChatConfigDeepSeek) (*http.Response, error) {
	client := getCDHttpClient()

	chatConfig.Init()
	var url string
	if chatConfig.Service == "auto" {
		url = setting.DeepSeekModelarts.SERVICE_AUTO_MODEL_URL + AUTO_MODEL_CHAT
		if chatConfig.Model != 0 {
			url = setting.DeepSeekModelarts.SERVICE_AUTO_MODEL_URL + AUTO_MODEL_CHAT_WITH_MODEL
		}
		if chatConfig.Id != 0 {
			url = setting.DeepSeekModelarts.SERVICE_AUTO_MODEL_URL + AUTO_MODEL_CHAT_WITH_ID
		}
	} else {
		url = setting.DeepSeekModelarts.SERVICE_HUB_URL
	}
	log.Info("DeepSeek URL: %s\n", url)
	if url == "" {
		return nil, fmt.Errorf("service URL is empty")
	}

	reqBytes, reqErr := json.Marshal(chatConfig)
	if reqErr != nil {
		log.Error("Error marshaling request parameters: %v", reqErr)
		return nil, fmt.Errorf("error marshaling request parameters: %w", reqErr)
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Error("Error creating new HTTP request: %v", err)
		return nil, fmt.Errorf("error creating new HTTP request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")

	log.Info("Sending request to deepseek service hub")
	response, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	log.Info("return response from service hub")
	return response, nil
}

func SendDeepSeekFeedBack(body *ChatFeedBack) (*http.Response, error) {
	client := getCDHttpClient()
	url := setting.DeepSeekModelarts.SERVICE_AUTO_MODEL_URL + AUTO_MODEL_FEEDBACK

	reqBytes, reqErr := json.Marshal(body)
	if reqErr != nil {
		log.Error("Error marshaling request parameters: %v", reqErr)
		return nil, fmt.Errorf("error marshaling request parameters: %w", reqErr)
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Error("Error creating new HTTP request: %v", err)
		return nil, fmt.Errorf("error creating new HTTP request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")

	log.Info("Sending request to deepseek service hub")
	response, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	log.Info("return response from service hub")
	return response, nil
}
