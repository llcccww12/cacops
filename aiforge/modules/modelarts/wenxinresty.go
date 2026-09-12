package modelarts

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts_gateway/core"
	"code.gitea.io/gitea/modules/setting"
)

type CreateWenXinParams struct {
	Data       WenXinText        `json:"data"`
	Parameters map[string]string `json:"parameters"`
}

type WenXinText struct {
	Prompt string `json:"prompt"`
}

type WenXinResult struct {
	Result string `json:"result"`
}

var (
	cdHttpClient *http.Client
)

func getCDHttpClient() *http.Client {
	if cdHttpClient == nil {
		cdHttpClient = &http.Client{
			Timeout:   300 * time.Second,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}
	}
	return cdHttpClient
}

func CreateWenXinJobToCD(modelapp *models.ModelApp, url string) (*WenXinResult, error) {
	createJobParams := &CreateWenXinParams{
		Data: WenXinText{
			Prompt: modelapp.Desc,
		},
		Parameters: make(map[string]string),
	}
	var result WenXinResult

	client := getCDHttpClient()
	s := core.Signer{
		Key:    setting.ModelartsCD.AccessKey,
		Secret: setting.ModelartsCD.SecretKey,
	}

	req, _ := json.Marshal(createJobParams)
	r, _ := http.NewRequest(http.MethodPost, url, ioutil.NopCloser(bytes.NewBuffer(req)))
	log.Info("send to cd modelarts")
	r.Header.Add("content-type", "application/json")
	s.Sign(r)
	res, err := client.Do(r)
	if err == nil {
		if res.StatusCode == 200 {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Error("ioutil.ReadAll failed: %s", err.Error())
				return &result, fmt.Errorf("ioutil.ReadAll failed: %s", err.Error())
			}
			err = json.Unmarshal(body, &result)
			if err != nil {
				log.Error("json.Unmarshal failed: %s", err.Error())
				return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
			}
			return &result, nil
		} else {
			log.Info("res.status=" + fmt.Sprint(res.StatusCode))
			return nil, fmt.Errorf("Service unavailable")
		}
	} else {
		log.Info("error =" + err.Error())
		return nil, fmt.Errorf("Service unavailable")
	}

}

func CreateWenXinJob(modelapp *models.ModelApp, url string) (*WenXinResult, error) {
	createJobParams := &CreateWenXinParams{
		Data: WenXinText{
			Prompt: modelapp.Desc,
		},
		Parameters: make(map[string]string),
	}
	checkSetting()
	client := getRestyClient()
	var result WenXinResult

	retry := 0

sendjob:
	log.Info("token=" + TOKEN + " url=" + url)
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Auth-Token", TOKEN).
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(url)

	if err != nil {
		return nil, fmt.Errorf("resty CreateWenXinJob: %s", err)
	}
	log.Info("re status=" + res.Status())
	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}
	if res.StatusCode() == 200 {
		var response WenXinResult
		err = json.Unmarshal(res.Body(), &response)
		if err != nil {
			log.Error("json.Unmarshal failed: %s", err.Error())
			return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
		}
		return &result, nil
	} else {
		return nil, fmt.Errorf("Service unavailable")
	}
}

var expire_time int64
var baidu_token string

func SendPictureReivew(picture_base64 string) bool {
	weburl := setting.BaiduWenXin.PICTURE_BAIDU_REVIEW_URL + GetBaiDuAccessToken()
	picture_base64 = url.QueryEscape(picture_base64)
	// image 可以通过 GetFileContentAsBase64("C:\fakepath\框图.jpg") 方法获取
	payload := strings.NewReader("image=" + picture_base64)
	client := &http.Client{}
	req, err := http.NewRequest("POST", weburl, payload)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	bodystr := string(body)
	log.Info("img bodystr=" + bodystr)
	if strings.Index(bodystr, "\"conclusionType\": 1") > 0 || strings.Index(bodystr, "\"conclusionType\":1") > 0 {
		return true
	}
	if strings.Index(bodystr, "\"error_code\": 18") > 0 || strings.Index(bodystr, "\"error_code\":18") > 0 {
		return true
	}
	return false
}

/**
 * 获取文件base64编码
 * @param string  path 文件路径
 * @return string base64编码信息，不带文件头
 */
func GetFileContentAsBase64(path string) string {
	srcByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(srcByte)
}

var lock sync.Mutex
var textCache = make(map[string]bool, 2000)

func SendTextReview(text string) bool {
	if textCache[text] {
		return true
	}
	bodystr := sendTextReviewToBaidu(text)
	log.Info("text bodystr=" + bodystr)
	if strings.Index(bodystr, "\"conclusionType\": 1") > 0 || strings.Index(bodystr, "\"conclusionType\":1") > 0 {
		setCacheTrue(text, true)
		return true
	}
	if strings.Index(bodystr, "\"error_code\": 18") > 0 || strings.Index(bodystr, "\"error_code\":18") > 0 {
		for i := 1; i <= 3; i++ {
			log.Info("sleep seconds " + fmt.Sprint(i))
			time.Sleep(time.Duration(i) * time.Second)
			bodystr := sendTextReviewToBaidu(text)
			log.Info("text bodystr=" + bodystr)
			if strings.Index(bodystr, "\"conclusionType\": 1") > 0 || strings.Index(bodystr, "\"conclusionType\":1") > 0 {
				setCacheTrue(text, true)
				return true
			}
		}
	}
	setCacheTrue(text, false)
	return false
}

func setCacheTrue(text string, isRight bool) {
	lock.Lock()
	defer lock.Unlock()
	if len(textCache) >= 1500 {
		for k := range textCache {
			delete(textCache, k) //delete only one
			break
		}
	}
	textCache[text] = isRight
}

func sendTextReviewToBaidu(text string) string {
	url := setting.BaiduWenXin.TEXT_BAIDU_REVIEW_URL + GetBaiDuAccessToken()
	payload := strings.NewReader("text=" + text)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bodystr := string(body)
	return bodystr
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetBaiDuAccessToken() string {
	if baidu_token != "" {
		if time.Now().Unix()-expire_time < 20*24*3600 {
			return baidu_token
		}
	}

	url := setting.BaiduWenXin.BAIDU_TOKEN_URL
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", setting.BaiduWenXin.BAIDU_API_KEY, setting.BaiduWenXin.BAIDU_SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
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
	baidu_token = accessTokenObj["access_token"]
	expire_time = time.Now().Unix()
	return accessTokenObj["access_token"]
}
