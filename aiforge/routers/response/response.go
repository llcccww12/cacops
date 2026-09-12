package response

const (
	RESPONSE_CODE_SUCCESS       = 0
	RESPONSE_MSG_SUCCESS        = "ok"
	RESPONSE_CODE_ERROR_DEFAULT = 99
)

type AiforgeResponse struct {
	Code int
	Msg  string
	Data interface{}
}

func Success() *AiforgeResponse {
	return &AiforgeResponse{Code: RESPONSE_CODE_SUCCESS, Msg: RESPONSE_MSG_SUCCESS}
}

func Error(code int, msg string) *AiforgeResponse {
	return &AiforgeResponse{Code: code, Msg: msg}
}

func ServerError(msg string) *AiforgeResponse {
	return &AiforgeResponse{Code: RESPONSE_CODE_ERROR_DEFAULT, Msg: msg}
}

func ResponseBizError(err *BizError) *AiforgeResponse {
	return &AiforgeResponse{Code: err.Code, Msg: err.DefaultMsg}
}

func ResponseError(err error) *AiforgeResponse {
	return &AiforgeResponse{Code: RESPONSE_CODE_ERROR_DEFAULT, Msg: err.Error()}
}

func SuccessWithData(data interface{}) *AiforgeResponse {
	return &AiforgeResponse{Code: RESPONSE_CODE_SUCCESS, Msg: RESPONSE_MSG_SUCCESS, Data: data}
}
func ErrorWithData(code int, msg string, data interface{}) *AiforgeResponse {
	return &AiforgeResponse{Code: code, Msg: msg, Data: data}
}
