package aisafety

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"

	"github.com/go-resty/resty/v2"
)

var (
	restyClient *resty.Client
	HOST        string
	KEY         string
)

type TaskReq struct {
	UnionId     string   //评测任务ID，唯一标识，由后端生成UUID
	EvalName    string   //评测任务名称
	EvalContent string   //评测任务描述
	TLPath      string   //
	Indicators  []string //评测指标，由GetAlgorithmList接口返回的指标列表中的title属性值
	CDPath      string
	CDName      string //对抗数据集名称
	BDPath      string
	BDName      string //原数据集名称
}

type ReturnMsg struct {
	Code  string     `json:"code"`
	Msg   string     `json:"msg"`
	Data  ReturnData `json:"data"`
	Times int64      `json:"times"`
}

type ReturnData struct {
	ID           int    `json:"id"`
	No           string `json:"no"`
	StandardJson string `json:"standardJson"`
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	Status       int    `json:"status"`
}

const (
	APPID               = "1"
	LogPageSize         = 500
	LogPageTokenExpired = "5m"
	pageSize            = 15
	Success             = "S000"
)

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
	}
	return restyClient
}

func checkSetting() {
	if len(HOST) != 0 && len(KEY) != 0 {
		return
	}
	_ = loginCloudbrain()
}

func loginCloudbrain() error {
	HOST = setting.ModelSafetyTest.HOST
	KEY = setting.ModelSafetyTest.KEY
	return nil
}

func createSign(params map[string]interface{}, signKey string) string {
	var sceneList []string
	for k := range params {
		sceneList = append(sceneList, k)
	}
	sort.Strings(sceneList)
	re := ""
	for _, key := range sceneList {
		if params[key] != nil {
			re += key + "=" + fmt.Sprint(params[key]) + "&"
		}
	}
	re += "key=" + signKey
	log.Info("sign key:" + re)
	h := md5.New()
	h.Write([]byte(re))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func getParams(req TaskReq) (map[string]interface{}, string) {
	params := make(map[string]interface{})
	reStr := ""

	params["unionId"] = req.UnionId
	reStr += "unionId=" + req.UnionId
	params["evalName"] = req.EvalName
	reStr += "&evalName=" + req.EvalName
	params["evalContent"] = req.EvalContent
	reStr += "&evalContent=" + url.QueryEscape(req.EvalContent)
	params["TLPath"] = req.TLPath
	reStr += "&TLPath=" + url.QueryEscape(req.TLPath)

	params["CDName"] = req.CDName
	reStr += "&CDName=" + url.QueryEscape(req.CDName)

	params["BDName"] = req.BDName
	reStr += "&BDName=" + url.QueryEscape(req.BDName)

	if req.CDPath != "" {
		params["CDPath"] = req.CDPath
		reStr += "&CDPath=" + url.QueryEscape(req.CDPath)
	}
	if req.BDPath != "" {
		params["BDPath"] = req.BDPath
		reStr += "&BDPath=" + url.QueryEscape(req.BDPath)
	}
	indicators := ""
	if len(req.Indicators) > 0 {
		for _, tmp := range req.Indicators {
			indicators += tmp + "|"
		}
	}
	if len(indicators) > 0 {
		indicators = indicators[0 : len(indicators)-1]
	}

	params["Indicators"] = indicators
	log.Info("indicators=" + indicators)
	reStr += "&Indicators=" + url.QueryEscape(indicators)

	return params, reStr
}

func CreateSafetyTask(req TaskReq, jsonstr string) (string, error) {
	checkSetting()
	client := getRestyClient()

	//reqPara, _ := json.Marshal(body)
	//log.Warn("job req:", jsonstr)

	params, urlQuerys := getParams(req)

	bodyMap := make(map[string]interface{})
	//reJsonMap := make(map[string]interface{})
	bodyMap["resJson"] = jsonstr
	//bodyMap["externalEvalParam"] = reJsonMap

	//reqPara, _ := json.Marshal(bodyMap)
	//log.Warn("job req json:", string(reqPara))

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("appId", APPID).
		SetHeader("sign", createSign(params, KEY)).
		//SetAuthToken(TOKEN).
		SetBody(bodyMap).
		Post(HOST + "/v1/external/eval-standard/create?" + urlQuerys)
	log.Info("url=" + HOST + "/v1/external/eval-standard/create?" + urlQuerys)

	responseStr := string(res.Body())
	log.Info("CreateSafetyTask responseStr=" + responseStr + "  res code=" + fmt.Sprint(res.StatusCode()))

	if err != nil {
		return "", fmt.Errorf("resty create job: %s", err)
	} else {
		log.Info("result string=" + "  res code=" + fmt.Sprint(res.StatusCode()))
	}

	reMap := make(map[string]interface{})

	err = json.Unmarshal(res.Body(), &reMap)
	if err == nil && reMap["code"] == "0" {
		dataMap := (reMap["data"]).(map[string]interface{})
		return fmt.Sprint(dataMap["serialNo"]), nil
	}
	return "", nil
}

func GetAlgorithmList() (map[string]interface{}, error) {
	checkSetting()
	client := getRestyClient()
	params := make(map[string]interface{})

	jsonResult := make(map[string]interface{})
	sign := createSign(params, KEY)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("appId", APPID).
		SetHeader("sign", sign).
		Get(HOST + "/v1/external/eval-standard/algorithmList")

	log.Info("url=" + HOST + "/v1/external/eval-standard/algorithmList" + " sign=" + sign + "  appId=" + APPID)

	jsonerr := json.Unmarshal(res.Body(), &jsonResult)
	if jsonerr == nil {
		log.Info("jsonResult code=" + fmt.Sprint(jsonResult["msg"]))
	}
	responseStr := string(res.Body())
	log.Info("GetAlgorithmList responseStr=" + responseStr + "  res code=" + fmt.Sprint(res.StatusCode()))

	if err != nil {
		log.Info("error =" + err.Error())
		return nil, fmt.Errorf("resty GetJob: %v", err)
	} else {
		reMap := make(map[string]interface{})
		err = json.Unmarshal(res.Body(), &reMap)
		if err == nil && reMap["code"] == "0" {
			return reMap, nil
		} else {
			return nil, fmt.Errorf("get error,code not 0")
		}
	}

}

func GetTaskStatus(jobID string) (*ReturnMsg, error) {
	checkSetting()
	client := getRestyClient()
	var reMap ReturnMsg
	params := make(map[string]interface{})
	params["serialNo"] = jobID

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("appId", APPID).
		SetHeader("sign", createSign(params, KEY)).
		SetResult(&reMap).
		Get(HOST + "/v1/external/eval-standard/query?serialNo=" + jobID)

	log.Info("url=" + HOST + "/v1/external/eval-standard/query?serialNo=" + jobID)
	responseStr := string(res.Body())
	log.Info("GetTaskStatus responseStr=" + responseStr + "  res code=" + fmt.Sprint(res.StatusCode()))

	if err != nil {
		log.Info("error =" + err.Error())
		return nil, fmt.Errorf("Get task status error: %v", err)
	} else {
		return &reMap, nil
	}
}
