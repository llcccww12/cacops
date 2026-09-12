package baiduai

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func getAccessToken() string {
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", setting.BAIDU_AI.API_KEY, setting.BAIDU_AI.SECRET_KEY)
	resp, err := http.Post(setting.BAIDU_AI.URL, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}

/**
 * 百度api文本内容合规检测
 */
func CheckLegalText(query string) (*LegalTextResponse, error) {
	var result LegalTextResponse

	url := setting.BAIDU_AI.LEGAL_TEXT_URL + getAccessToken()
	payload := strings.NewReader("text=" + query)
	log.Info("resty CheckLegalText() payload %+v", payload)
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Error("resty CheckLegalText() Request error: %s", err.Error())
		return &result, fmt.Errorf("resty CheckLegalText(): %s", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Error("resty CheckLegalText() Response error: %s", err.Error())
		return &result, fmt.Errorf("resty CheckLegalText(): %s", err)
	}
	defer res.Body.Close()
	log.Error("resty CheckLegalText() Response status: %s\n", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("resty CheckLegalText() Response body error: %s", err.Error())
		return &result, fmt.Errorf("resty CheckLegalText(): %s", err)
	}

	response := string(body)
	json.Unmarshal([]byte(response), &result)
	log.Info("resty CheckLegalText() results: %+v\n", result)
	return &result, nil
}
