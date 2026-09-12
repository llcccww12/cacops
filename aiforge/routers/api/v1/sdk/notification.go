package sdk

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
)

type Message struct {
	Msg   string `json:"msg"`
	Level string `json:"level"`
}

type Notify struct {
	Code     int       `json:"code"`
	Messages []Message `json:"messages"`
}

func parseRawJsonArray(jsonStr string) ([]Message, error) {
	var messages []Message
	if err := json.Unmarshal([]byte(jsonStr), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func SendNotification(ctx *context.APIContext) {
	code := 0

	notifyMessages, err := parseRawJsonArray(setting.SDK_NOTICE)
	if err != nil {
		log.Error("Failed to parse SDK notice JSON: %v", err)
		code = 99
	}

	ctx.JSON(200, Notify{
		Code:     code,
		Messages: notifyMessages,
	})
}
