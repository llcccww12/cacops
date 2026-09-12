package setting

import (
	"code.gitea.io/gitea/modules/log"
)

type Phone struct {
	Enabled          bool
	VerifyCodeLength int
	AccessKeyId      string
	AccessKeySecret  string
	SignName         string
	TemplateCode     string
	CodeTimeout      int
	RetryInterval    int
	MaxRetryTimes    int
}

var (
	// Phone  verify info
	PhoneService *Phone
)

func newPhoneService() {
	sec := Cfg.Section("phone")
	// Check phone setting.
	if !sec.Key("ENABLED").MustBool() {
		PhoneService = &Phone{
			Enabled: sec.Key("ENABLED").MustBool(),
		}
		return
	}

	PhoneService = &Phone{
		Enabled:          sec.Key("ENABLED").MustBool(),
		VerifyCodeLength: sec.Key("VERIFY_CODE_LEN").MustInt(6),
		AccessKeyId:      sec.Key("AccessKeyId").String(),
		AccessKeySecret:  sec.Key("AccessKeySecret").String(),
		SignName:         sec.Key("SignName").String(),
		TemplateCode:     sec.Key("TemplateCode").String(),
		CodeTimeout:      sec.Key("CODE_TIMEOUT").MustInt(60 * 5),
		RetryInterval:    sec.Key("RETRY_INTERVAL").MustInt(60 * 2),
		MaxRetryTimes:    sec.Key("MAX_RETRY").MustInt(5),
	}

	log.Info("Phone Service Enabled")
}
