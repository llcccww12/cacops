package response

import "errors"

type BizError struct {
	Code       int
	DefaultMsg string
	TrCode     string
	TrParams   []interface{}
}

// 当调用此方法时意味着错误信息中有占位符，需要传入参数
// 因此此时需要新建一个对象避免并发问题
func (e *BizError) WithParams(params ...interface{}) *BizError {
	newErr := &BizError{
		Code:       e.Code,
		DefaultMsg: e.DefaultMsg,
		TrCode:     e.TrCode,
	}
	if e.TrParams == nil {
		newErr.TrParams = params
	} else {
		newErr.TrParams = append(e.TrParams, params)
	}

	return newErr
}

func (e *BizError) ToError() error {
	msg := e.TrCode
	if msg == "" {
		msg = e.DefaultMsg
	}
	return errors.New(msg)
}

func NewBizError(err error) *BizError {
	return &BizError{Code: RESPONSE_CODE_ERROR_DEFAULT, DefaultMsg: err.Error(), TrCode: err.Error()}
}

func BuildBizError(code int, defaultMsg string, trCode ...string) *BizError {
	t := ""
	if len(t) == 0 && len(trCode) > 0 {
		t = trCode[0]
	}
	return &BizError{Code: code, DefaultMsg: defaultMsg, TrCode: t}
}
func BuildDefaultBizError(defaultMsg string, trCode ...string) *BizError {
	t := ""
	if len(trCode) > 0 {
		t = trCode[0]
	}
	return &BizError{Code: RESPONSE_CODE_ERROR_DEFAULT, DefaultMsg: defaultMsg, TrCode: t}
}
