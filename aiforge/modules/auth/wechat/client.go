package wechat

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

var (
	client *resty.Client
)

const (
	GRANT_TYPE          = "client_credential"
	ACCESS_TOKEN_PATH   = "/cgi-bin/token"
	QR_CODE_PATH        = "/cgi-bin/qrcode/create"
	GET_MATERIAL_PATH   = "/cgi-bin/material/batchget_material"
	SEND_TEMPLATE_PATH  = "/cgi-bin/message/template/send"
	ACTION_QR_STR_SCENE = "QR_STR_SCENE"

	ERR_CODE_ACCESSTOKEN_EXPIRE  = 42001
	ERR_CODE_ACCESSTOKEN_INVALID = 40001
)

type AccessTokenResponse struct {
	Access_token string
	Expires_in   int
}

type QRCodeResponse struct {
	Ticket         string `json:"ticket"`
	Expire_Seconds int    `json:"expire_seconds"`
	Url            string `json:"url"`
}

type QRCodeRequest struct {
	Action_name    string     `json:"action_name"`
	Action_info    ActionInfo `json:"action_info"`
	Expire_seconds int        `json:"expire_seconds"`
}

type MaterialRequest struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

type TemplateMsgRequest struct {
	ToUser      string      `json:"touser"`
	TemplateId  string      `json:"template_id"`
	Url         string      `json:"url"`
	ClientMsgId string      `json:"client_msg_id"`
	Data        interface{} `json:"data"`
}
type TemplateValue struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type DefaultWechatTemplate struct {
	First    TemplateValue `json:"first"`
	Keyword1 TemplateValue `json:"keyword1"`
	Keyword2 TemplateValue `json:"keyword2"`
	Keyword3 TemplateValue `json:"keyword3"`
	Keyword4 TemplateValue `json:"keyword4"`
	Remark   TemplateValue `json:"remark"`
}

type WechatDurationWechatTemplate struct {
	Keyword1 TemplateValue `json:"thing11"`
	Keyword2 TemplateValue `json:"time9"`
	Keyword3 TemplateValue `json:"thing3"`
}

type ActionInfo struct {
	Scene Scene `json:"scene"`
}

type Scene struct {
	Scene_str string `json:"scene_str"`
}

type ErrorResponse struct {
	Errcode int
	Errmsg  string
}

func getWechatRestyClient() *resty.Client {
	if client == nil {
		client = resty.New()
		client.SetTimeout(time.Duration(setting.WechatApiTimeoutSeconds) * time.Second)
	}
	return client
}

// api doc:https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func callAccessToken() *AccessTokenResponse {
	client := getWechatRestyClient()

	log.Info("start to get wechat access token")
	var result AccessTokenResponse
	_, err := client.R().
		SetQueryParam("grant_type", GRANT_TYPE).
		SetQueryParam("appid", setting.WechatAppId).
		SetQueryParam("secret", setting.WechatAppSecret).
		SetResult(&result).
		Get(setting.WechatApiHost + ACCESS_TOKEN_PATH)
	if err != nil {
		log.Error("get wechat access token failed,e=%v", err)
		return nil
	}
	log.Info("get wechat access token result=%v", result)
	return &result
}

//callQRCodeCreate  call the wechat api to create qr-code,
// api doc: https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func callQRCodeCreate(sceneStr string) (*QRCodeResponse, bool) {
	client := getWechatRestyClient()

	body := &QRCodeRequest{
		Action_name:    ACTION_QR_STR_SCENE,
		Action_info:    ActionInfo{Scene: Scene{Scene_str: sceneStr}},
		Expire_seconds: setting.WechatQRCodeExpireSeconds,
	}
	bodyJson, _ := json.Marshal(body)
	var result QRCodeResponse
	r, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("access_token", GetWechatAccessToken()).
		SetBody(bodyJson).
		SetResult(&result).
		Post(setting.WechatApiHost + QR_CODE_PATH)
	if err != nil {
		log.Error("create QR code failed,e=%v", err)
		return nil, false
	}
	errCode := getErrorCodeFromResponse(r)
	if errCode == ERR_CODE_ACCESSTOKEN_EXPIRE || errCode == ERR_CODE_ACCESSTOKEN_INVALID {
		return nil, true
	}
	if result.Url == "" {
		return nil, false
	}
	log.Info("%v", r)
	return &result, false
}

//getMaterial
// api doc: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func getMaterial(mType string, offset, count int) (interface{}, bool) {
	client := getWechatRestyClient()

	body := &MaterialRequest{
		Type:   mType,
		Offset: offset,
		Count:  count,
	}
	bodyJson, _ := json.Marshal(body)
	r, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("access_token", GetWechatAccessToken()).
		SetBody(bodyJson).
		Post(setting.WechatApiHost + GET_MATERIAL_PATH)
	if err != nil {
		log.Error("create QR code failed,e=%v", err)
		return nil, false
	}
	a := r.Body()
	resultMap := make(map[string]interface{}, 0)
	json.Unmarshal(a, &resultMap)
	errcode := resultMap["errcode"]
	if errcode == fmt.Sprint(ERR_CODE_ACCESSTOKEN_EXPIRE) || errcode == fmt.Sprint(ERR_CODE_ACCESSTOKEN_INVALID) {
		return nil, true
	}
	log.Info("%v", r)
	return &resultMap, false
}

func getErrorCodeFromResponse(r *resty.Response) int {
	a := r.Body()
	resultMap := make(map[string]interface{}, 0)
	json.Unmarshal(a, &resultMap)
	code := resultMap["errcode"]
	if code == nil {
		return -1
	}
	c, _ := strconv.Atoi(fmt.Sprint(code))
	return c
}

func sendTemplateMsg(req TemplateMsgRequest) (error, bool) {
	client := getWechatRestyClient()

	bodyJson, _ := json.Marshal(req)
	r, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("access_token", GetWechatAccessToken()).
		SetBody(bodyJson).
		Post(setting.WechatApiHost + SEND_TEMPLATE_PATH)
	if err != nil {
		log.Error("sendTemplateMsg,e=%v", err)
		return nil, false
	}
	a := r.Body()
	resultMap := make(map[string]interface{}, 0)
	json.Unmarshal(a, &resultMap)
	errcode := resultMap["errcode"]
	log.Info("sendTemplateMsg,%v", r)
	if errcode == fmt.Sprint(ERR_CODE_ACCESSTOKEN_EXPIRE) || errcode == fmt.Sprint(ERR_CODE_ACCESSTOKEN_INVALID) {
		return nil, true
	}
	return nil, false
}
