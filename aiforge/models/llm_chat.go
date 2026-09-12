package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"fmt"
)

type LlmChat struct {
	ID                string `xorm:"pk"`
	Prompt            string `xorm:"text"`
	Answer            string `xorm:"text"`
	Params            string `xorm:"text"`
	ChatId            string `xorm:"INDEX"`
	TaskId            int64  `xorm:"INDEX"`
	UserId            int64  `xorm:"INDEX"`
	ModelName         string
	ModelId           string
	Endpoint          string
	CreatedUnix       timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix       timeutil.TimeStamp `xorm:"updated"`
	ChatType          string
	Count             int
	InvalidCount      int
	InvalidType       string
	InvalidTool       string
	InvalidDetail     string `xorm:"text"`
	ChatStatus        int
	KnowledgeBaseName string
	VectorStoreType   string
	EmbeddingModel    string
}

type LlmChatParams struct {
	SystemMessage     string  `json:"system_message,omitempty"`
	MaxTokens         int     `json:"max_tokens,omitempty"`
	Temperature       float64 `json:"temperature,omitempty"`
	TopP              float64 `json:"top_p,omitempty"`
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty"`
}

type PayloadText struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type IFLYTekAPIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type IFLYTekAPIRequest struct {
	Model    string        `json:"model"`
	Messages []PayloadText `json:"messages"`
}

// IFLYTekAPIChoice represents the choice structure within the response.
type IFLYTekAPIChoice struct {
	PayloadText `json:"message"`
	Index       int `json:"index"`
}

// IFLYTekAPIUsage represents the usage statistics in the response.
type IFLYTekAPIUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Response represents the overall JSON response structure.
type IFLYTekAPIResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Sid     string             `json:"sid"`
	Choices []IFLYTekAPIChoice `json:"choices"`
	Usage   IFLYTekAPIUsage    `json:"usage"`
}

func SaveChat(llmChat *LlmChat) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	re, err := sess.Insert(llmChat)
	if err != nil {
		fmt.Printf("insert llmChat error %s\n", err.Error())
		return err
	}
	fmt.Printf("success to save llmChat db.re=%+v\n", fmt.Sprint(re))
	return nil
}

func QueryChatCount(userId int64, modelName string) int64 {
	sess := xStatistic.NewSession()
	defer sess.Close()
	query := "SELECT SUM(count) AS count FROM public.llm_chat WHERE chat_status = 1 AND user_id = ? AND model_name = ?"
	sumList, err := sess.QueryInterface(query, userId, modelName)
	if err == nil {
		if len(sumList) == 1 {
			return convertInterfaceToInt64(sumList[0]["count"])
		}
	}
	return 0
}

func QueryInvalidPromptCount(userId int64) int64 {
	sess := xStatistic.NewSession()
	defer sess.Close()
	query := "SELECT SUM(invalid_count) AS count FROM public.llm_chat WHERE invalid_type='prompt' and user_id = ?"
	sumList, err := sess.QueryInterface(query, userId)
	if err == nil {
		if len(sumList) == 1 {
			return convertInterfaceToInt64(sumList[0]["count"])
		}
	}
	return 0
}

func QueryChatStatistics() ([]map[string]interface{}, error) {
	sess := xStatistic.NewSession()
	query := `
		SELECT
			COALESCE(model_name, 'total') as model_name,
			COUNT(DISTINCT id) as chat,
			COUNT(DISTINCT user_id) as chat_user,
			COUNT(DISTINCT id) / COUNT(DISTINCT user_id) as chat_per_user_avg,
			COUNT(DISTINCT CASE WHEN chat_type = 'llm' THEN id END) as chat_llm,
			COUNT(DISTINCT CASE WHEN chat_type = 'llm' THEN user_id END) as chat_llm_user,
			COUNT(DISTINCT CASE WHEN chat_type = 'kb' THEN id END) as chat_kb,
			COUNT(DISTINCT CASE WHEN chat_type = 'kb' THEN user_id END) as chat_kb_user,
			COALESCE(SUM(invalid_count),0) as chat_illegal,
			COALESCE(SUM(CASE WHEN invalid_type = 'prompt' THEN 1 END),0) as chat_illegal_prompt,
			COALESCE(SUM(CASE WHEN invalid_type = 'answer' THEN 1 END),0) as chat_illegal_answer
		FROM
			llm_chat
		GROUP BY
			ROLLUP(model_name)`

	results, err := sess.SQL(query).QueryInterface()
	if err != nil {
		return nil, err
	}

	return results, nil
}
