package cloudbrain

import (
	"bytes"
	"code.gitea.io/gitea/modules/httplib"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	UrlToken       = "/rest-server/api/v1/token/"
	UrlGetUserInfo = "/rest-server/api/v1/user/"

	TokenTypeBear = "Bearer "

	SuccessCode = "S000"
)

type RespAuth struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type RespToken struct {
	Code    string       `json:"code"`
	Message string       `json:"msg"`
	Payload PayloadToken `json:"payload"`
}

type PayloadToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	IsAdmin  bool   `json:"admin"`
}

type RespUserInfo struct {
	Code    string          `json:"code"`
	Message string          `json:"msg"`
	Payload PayloadUserInfo `json:"payload"`
}

type PayloadUserInfo struct {
	UserInfo StUserInfo `json:"userInfo"`
}

type StUserInfo struct {
	Email string `json:"email"`
}

type CloudBrainUser struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func UserValidate(username string, password string) (string, error) {
	values := map[string]string{"username": username, "password": password}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(setting.RestServerHost+UrlToken,
		"application/json",
		bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Error("req user center failed:" + err.Error())
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("read resp body failed:" + err.Error())
		return "", err
	}

	var res RespToken
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Error("unmarshal res failed:" + err.Error())
		return "", err
	}

	if res.Code != SuccessCode {
		log.Error("req rest-server for token failed:", res.Message)
		return "", errors.New(res.Message)
	}

	return res.Payload.Token, nil
}

func GetUserInfo(username string, token string) (*CloudBrainUser, error) {
	user := &CloudBrainUser{}

	client := httplib.NewClientTimeOut(&httplib.Config{})
	reqHttp, err := http.NewRequest("GET", setting.RestServerHost+UrlGetUserInfo+username, strings.NewReader(""))
	if err != nil {
		log.Error("new req failed:", err.Error())
		return nil, err
	}

	reqHttp.Header.Set("Authorization", TokenTypeBear+token)
	resp, err := client.Do(reqHttp)
	if err != nil {
		log.Error("req rest-server failed:", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("read resp body failed:", err.Error())
		return nil, err
	}

	var res RespUserInfo
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Error("unmarshal resp failed:", err.Error())
		return nil, err
	}

	if res.Code != SuccessCode {
		log.Error("get userInfo failed:", err.Error())
		return nil, err
	}

	user.Email = res.Payload.UserInfo.Email
	return user, nil
}
