package response

import "gitea.com/macaron/macaron"

type AiforgeOuterResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OuterSuccess() *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: RESPONSE_CODE_SUCCESS, Msg: RESPONSE_MSG_SUCCESS}
}

func OuterError(code int, msg string) *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: code, Msg: msg}
}

func OuterServerError(msg string) *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: RESPONSE_CODE_ERROR_DEFAULT, Msg: msg}
}

func OuterBizError(err *BizError) *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: err.Code, Msg: err.DefaultMsg}
}

type TrFunc func(string, ...interface{}) string

func OuterTrBizError(err *BizError, locale macaron.Locale) *AiforgeOuterResponse {
	msg := err.DefaultMsg
	if locale != nil && err.TrCode != "" {
		if err.TrParams == nil || len(err.TrParams) == 0 {
			msg = locale.Tr(err.TrCode)
		} else {
			msg = locale.Tr(err.TrCode, err.TrParams...)
		}
		if msg == "" {
			msg = err.DefaultMsg
		}
	}
	return &AiforgeOuterResponse{Code: err.Code, Msg: msg}
}

func OuterSuccessWithData(data interface{}) *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: RESPONSE_CODE_SUCCESS, Msg: RESPONSE_MSG_SUCCESS, Data: data}
}
func OuterErrorWithData(code int, msg string, data interface{}) *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: code, Msg: msg, Data: data}
}
func OuterResponseError(err error) *AiforgeOuterResponse {
	return &AiforgeOuterResponse{Code: RESPONSE_CODE_ERROR_DEFAULT, Msg: err.Error()}
}
