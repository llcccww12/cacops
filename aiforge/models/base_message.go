package models

type BaseMessage struct {
	Code    int
	Message string
}

var BaseOKMessage = BaseMessage{
	0, "",
}

func BaseErrorMessage(message string) BaseMessage {
	return BaseMessage{
		1, message,
	}
}

type BaseMessageApi struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var BaseOKMessageApi = BaseMessageApi{
	0, "",
}

func BaseErrorMessageApi(message string) BaseMessageApi {
	return BaseMessageApi{
		1, message,
	}
}

type BaseMessageWithDataApi struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
