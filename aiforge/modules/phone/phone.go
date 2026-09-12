package phone

import (
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/setting"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := "^1[3-9]\\d{9}$"
	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}

func GenerateVerifyCode(n int) string {
	min := int(math.Pow10(n - 1))
	max := int(math.Pow10(n))
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(max-min) + min)

}

func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendVerifyCode(phoneNumber string, verifyCode string) error {
	client, _err := createClient(&setting.PhoneService.AccessKeyId, &setting.PhoneService.AccessKeySecret)
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(setting.PhoneService.SignName),
		TemplateCode:  tea.String(setting.PhoneService.TemplateCode),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String("{\"code\":\"" + verifyCode + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
	return _err
}
