package modelarts

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
)

var (
	restyClient *resty.Client
	HOST        string
	TOKEN       string
)

const (
	methodPassword = "password"

	urlGetToken          = "/v3/auth/tokens"
	urlNotebook          = "/demanager/instances"
	urlTrainJob          = "/training-jobs"
	urlResourceSpecs     = "/job/resource-specs"
	urlTrainJobConfig    = "/training-job-configs"
	urlDeployModel       = "/models"
	urlDeployService     = "/services"
	errorCodeExceedLimit = "ModelArts.0118"

	//notebook 2.0
	urlNotebook2 = "/notebooks"

	//error code
	modelartsIllegalToken = "ModelArts.6401"
	NotebookNotFound      = "ModelArts.6404"
	NotebookNoPermission  = "ModelArts.6407"
	NotebookInvalid       = "ModelArts.6400"
	UnknownErrorPrefix    = "UNKNOWN:"

	ModelArtsJobInTargetState = "ModelArts.6357"
	ModelArtsJobNotExists     = "ModelArts.0102"
	ModelArtsJobInternalError = "ModelArts.0010"
)

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

func checkSetting() {
	if len(HOST) != 0 && len(TOKEN) != 0 && restyClient != nil {
		return
	}

	err := getToken()
	if err != nil {
		log.Error("getToken failed:%v", err)
	}
}

func getToken() error {
	HOST = setting.ModelArtsHost

	client := getRestyClient()
	params := models.GetTokenParams{
		Auth: models.Auth{
			Identity: models.Identity{
				Methods: []string{methodPassword},
				Password: models.Password{
					User: models.NotebookUser{
						Name:     setting.ModelArtsUsername,
						Password: setting.ModelArtsPassword,
						Domain: models.Domain{
							Name: setting.ModelArtsDomain,
						},
					},
				},
			},
			Scope: models.Scope{
				Project: models.Project{
					Name: setting.ProjectName,
				},
			},
		},
	}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		Post(setting.IamHost + urlGetToken)
	if err != nil {
		return fmt.Errorf("resty getToken: %v", err)
	}

	if res.StatusCode() != http.StatusCreated {
		return fmt.Errorf("getToken failed:%s", res.String())
	}

	TOKEN = res.Header().Get("X-Subject-Token")

	return nil
}

func CreateJob(createJobParams models.CreateNotebookParams) (*models.CreateNotebookResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateNotebookResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlNotebook)

	if err != nil {
		return nil, fmt.Errorf("resty create notebook: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("createNotebook failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		if response.ErrorCode == errorCodeExceedLimit {
			response.ErrorMsg = "所选规格使用数量已超过最大配额限制。"
		}
		return &result, fmt.Errorf("createNotebook failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func GetJob(jobID string) (*models.GetNotebookResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetNotebookResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlNotebook + "/" + jobID)

	if err != nil {
		return nil, fmt.Errorf("resty GetJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("GetJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		return &result, fmt.Errorf("GetJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func GetNotebook2(jobID string) (*models.GetNotebook2Result, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetNotebook2Result

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlNotebook2 + "/" + jobID)

	if err != nil {
		return nil, fmt.Errorf("resty GetJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("GetJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		if response.ErrorCode == modelartsIllegalToken && retry < 1 {
			retry++
			_ = getToken()
			goto sendjob
		}
		return &result, fmt.Errorf("GetJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func ManageNotebook(jobID string, param models.NotebookAction) (*models.NotebookActionResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NotebookActionResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlNotebook + "/" + jobID + "/action")

	if err != nil {
		return &result, fmt.Errorf("resty StopJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("ManageNotebook failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		return &result, fmt.Errorf("ManageNotebook failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func ManageNotebook2(jobID string, param models.NotebookAction) (*models.NotebookActionResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NotebookActionResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlNotebook2 + "/" + jobID + "/" + param.Action + "?duration=" + strconv.Itoa(AutoStopDurationMs))

	if err != nil {
		return &result, fmt.Errorf("resty ManageNotebook2: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if res.StatusCode() == http.StatusBadGateway {
		return &result, fmt.Errorf(UnknownErrorPrefix+"createNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	if len(response.ErrorCode) != 0 {
		log.Error("ManageNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		if response.ErrorCode == modelartsIllegalToken && retry < 1 {
			retry++
			_ = getToken()
			goto sendjob
		}
		return &result, fmt.Errorf("ManageNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func DelNotebook(jobID string) (*models.NotebookDelResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NotebookDelResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + "/v1/" + setting.ProjectID + urlNotebook + "/" + jobID)

	if err != nil {
		return &result, fmt.Errorf("resty DelJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("DelJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		return &result, fmt.Errorf("DelJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func DelNotebook2(jobID string) (*models.NotebookDelResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NotebookDelResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + "/v1/" + setting.ProjectID + urlNotebook2 + "/" + jobID)

	if err != nil {
		return &result, fmt.Errorf("resty DelJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("DelNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		if response.ErrorCode == modelartsIllegalToken && retry < 1 {
			retry++
			_ = getToken()
			goto sendjob
		}
		if response.ErrorCode == ModelArtsJobNotExists || response.ErrorCode == ModelArtsJobInTargetState {
			//任务不存在或者已经处于被删除的状态，此时认为删除成功
			return &models.NotebookDelResult{}, nil
		}
		if result.ErrorCode == ModelArtsJobInternalError {
			log.Error("ModelArt internal error when del job,jobId=%s", jobID)
			return &models.NotebookDelResult{}, nil
		}
		return &result, fmt.Errorf("DelNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func DelJob(jobID string) (*models.NotebookDelResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NotebookDelResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + "/v1/" + setting.ProjectID + urlNotebook + "/" + jobID)

	if err != nil {
		return &result, fmt.Errorf("resty DelJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("DelJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		return &result, fmt.Errorf("DelJob failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func GetJobToken(jobID string) (*models.NotebookGetJobTokenResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NotebookGetJobTokenResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlNotebook + "/" + jobID + "/token")

	if err != nil {
		return &result, fmt.Errorf("resty GetJobToken: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(response.ErrorCode) != 0 {
		log.Error("GetJobToken failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		return &result, fmt.Errorf("GetJobToken failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func CreateTrainJobUserImage(createJobParams models.CreateUserImageTrainJobParams) (*models.CreateTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob)

	if err != nil {
		return nil, fmt.Errorf("resty create train-job: %s", err)
	}

	req, _ := json.Marshal(createJobParams)
	log.Info("postapi json: %s", req)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("createTrainJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		bootFileErrorMsg := "Invalid OBS path '" + createJobParams.Config.BootFileUrl + "'."
		dataSetErrorMsg := "Invalid OBS path '" + createJobParams.Config.DataUrl + "'."
		if temp.ErrorMsg == bootFileErrorMsg {
			log.Error("启动文件错误！createTrainJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("启动文件错误！")
		}
		if temp.ErrorMsg == dataSetErrorMsg {
			log.Error("数据集错误！createTrainJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("数据集错误！")
		}
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createTrainJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createTrainJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	if !result.IsSuccess {
		log.Error("createTrainJobUserImage failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("createTrainJobUserImage failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func CreateTrainJob(createJobParams models.CreateTrainJobParams) (*models.CreateTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobResult

	retry := 0
	req, _ := json.Marshal(createJobParams)
	log.Info("postapi json: %s", req)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob)

	if err != nil {
		return nil, fmt.Errorf("resty create train-job: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("createTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		bootFileErrorMsg := "Invalid OBS path '" + createJobParams.Config.BootFileUrl + "'."
		dataSetErrorMsg := "Invalid OBS path '" + createJobParams.Config.DataUrl + "'."
		if temp.ErrorMsg == bootFileErrorMsg {
			log.Error("启动文件错误！createTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("启动文件错误！")
		}
		if temp.ErrorMsg == dataSetErrorMsg {
			log.Error("数据集错误！createTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("数据集错误！")
		}
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	if !result.IsSuccess {
		log.Error("createTrainJob failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("createTrainJob failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func createTrainJobVersion(createJobVersionParams models.CreateTrainJobVersionParams, jobID string) (*models.CreateTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobVersionParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions")

	if err != nil {
		return nil, fmt.Errorf("resty create train-job version: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}

		log.Error("createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		bootFileErrorMsg := "Invalid OBS path '" + createJobVersionParams.Config.BootFileUrl + "'."
		dataSetErrorMsg := "Invalid OBS path '" + createJobVersionParams.Config.DataUrl + "'."
		if temp.ErrorMsg == bootFileErrorMsg {
			log.Error("启动文件错误！createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("启动文件错误！")
		}
		if temp.ErrorMsg == dataSetErrorMsg {
			log.Error("数据集错误！createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("数据集错误！")
		}
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	if !result.IsSuccess {
		log.Error("createTrainJobVersion failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("createTrainJobVersion failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func createTrainJobVersionUserImage(createJobVersionParams models.CreateTrainJobVersionUserImageParams, jobID string) (*models.CreateTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobVersionParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions")

	if err != nil {
		return nil, fmt.Errorf("resty create train-job version: %s", err)
	}

	req, _ := json.Marshal(createJobVersionParams)
	log.Info("%s", req)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		BootFileErrorMsg := "Invalid OBS path '" + createJobVersionParams.Config.BootFileUrl + "'."
		DataSetErrorMsg := "Invalid OBS path '" + createJobVersionParams.Config.DataUrl + "'."
		if temp.ErrorMsg == BootFileErrorMsg {
			log.Error("启动文件错误！createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("启动文件错误！")
		}
		if temp.ErrorMsg == DataSetErrorMsg {
			log.Error("数据集错误！createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("数据集错误！")
		}
		return &result, fmt.Errorf("createTrainJobVersion failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("createTrainJobVersion failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("createTrainJobVersion failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetResourceSpecs() (*models.GetResourceSpecsResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetResourceSpecsResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlResourceSpecs)

	if err != nil {
		return nil, fmt.Errorf("resty GetResourceSpecs: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetResourceSpecs failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("GetResourceSpecs failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetResourceSpecs failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetResourceSpecs failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func CreateTrainJobConfig(req models.CreateConfigParams) (*models.CreateTrainJobConfigResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobConfigResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJobConfig)

	if err != nil {
		return nil, fmt.Errorf("resty CreateTrainJobConfig: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("CreateTrainJobConfig failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("CreateTrainJobConfig failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("CreateTrainJobConfig failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("CreateTrainJobConfig failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetConfigList(perPage, page int, sortBy, order, searchContent, configType string) (*models.GetConfigListResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetConfigListResult

	retry := 0

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"per_page":       strconv.Itoa(perPage),
			"page":           strconv.Itoa(page),
			"sortBy":         sortBy,
			"order":          order,
			"search_content": searchContent,
			"config_type":    configType,
		}).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJobConfig)

	if err != nil {
		return nil, fmt.Errorf("resty GetConfigList: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetConfigList failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("获取参数配置列表失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetConfigList failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("获取参数配置列表失败(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetParaConfig(configName, configType string) (models.GetConfigResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetConfigResult

	retry := 0

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"config_type": configType,
		}).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJobConfig + "/" + configName)

	if err != nil {
		return result, fmt.Errorf("resty GetParaConfig: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetParaConfig failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return result, fmt.Errorf("获取参数配置详情失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetParaConfig failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return result, fmt.Errorf("获取参数配置详情失败(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return result, nil
}

func GetTrainJob(jobID, versionID string) (*models.GetTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetTrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions/" + versionID)

	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("获取作业详情失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetTrainJob(%s) failed", jobID)
		return &result, fmt.Errorf("获取作业详情失败")
	}

	return &result, nil
}

func GetTrainJobLog(jobID, versionID, baseLine, logFile, order string, lines int) (*models.GetTrainJobLogResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetTrainJobLogResult

	retry := 0

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"base_line": baseLine,
			"lines":     strconv.Itoa(lines),
			"log_file":  logFile,
			"order":     order,
		}).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions/" + versionID + "/aom-log")

	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobLog: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJobLog failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("获取作业日志失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetTrainJobLog(%s) failed", jobID)
		return &result, fmt.Errorf("获取作业日志失败:%s", result.ErrorMsg)
	}

	return &result, nil
}

func GetTrainJobLogFileNames(jobID, versionID string) (*models.GetTrainJobLogFileNamesResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetTrainJobLogFileNamesResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions/" + versionID + "/log/file-names")

	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobLogFileNames: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJobLogFileNames failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("GetTrainJobLogFileNames failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetTrainJobLogFileNames(%s) failed", jobID)
		return &result, fmt.Errorf("获取作业日志文件失败:%s", result.ErrorMsg)
	}

	return &result, nil
}

func DelTrainJob(jobID string) (*models.TrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.TrainJobResult

	//get cloudbrain job by jobid
	if finetuneJob, err := models.GetCloudbrainByJobID(jobID); finetuneJob != nil && err == nil {
		if finetuneJob.FineTune {
			err := ServiceDelete(jobID)
			if err != nil {
				log.Error("panguService: Delete Deploy failed:%s %v", jobID, err.Error())
				return nil, err
			}
		}
	} else if err != nil {
		log.Warn("DelTrainJob GetCloudbrainByJobID from DB failed:%s %v", jobID, err.Error())
	}

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID)

	if err != nil {
		return &result, fmt.Errorf("resty DelTrainJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("DelTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		if temp.ErrorCode == ModelArtsJobNotExists || temp.ErrorCode == ModelArtsJobInTargetState {
			//任务不存在或者已经处于被删除的状态，此时认为删除成功
			return &models.TrainJobResult{IsSuccess: true}, nil
		}
		if result.ErrorCode == ModelArtsJobInternalError {
			log.Error("ModelArt internal error when del job,jobId=%s", jobID)
			return &models.TrainJobResult{IsSuccess: true}, nil
		}
		return &result, fmt.Errorf("删除训练作业失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("DelTrainJob(%s) failed", jobID)
		if result.ErrorCode == ModelArtsJobNotExists || result.ErrorCode == ModelArtsJobInTargetState {
			//任务不存在或者已经处于被删除的状态，此时认为删除成功
			return &models.TrainJobResult{IsSuccess: true}, nil
		}
		if result.ErrorCode == ModelArtsJobInternalError {
			log.Error("ModelArt internal error when del job,jobId=%s", jobID)
			return &models.TrainJobResult{IsSuccess: true}, nil
		}
		return &result, fmt.Errorf("删除训练作业失败:%s", result.ErrorMsg)
	}

	return &result, nil
}

func StopTrainJob(jobID, versionID string) (*models.TrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.TrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions/" + versionID + "/stop")

	if err != nil {
		return &result, fmt.Errorf("resty StopTrainJob: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("StopTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("停止训练作业失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("StopTrainJob(%s) failed", jobID)
		return &result, fmt.Errorf("停止训练作业失败:%s", result.ErrorMsg)
	}

	return &result, nil
}

func DelTrainJobVersion(jobID string, versionID string) (*models.TrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.TrainJobResult

	//get cloudbrain job by jobid
	if finetuneJob, err := models.GetCloudbrainByJobID(jobID); finetuneJob != nil && err == nil {
		if finetuneJob.FineTune {
			err := ServiceDelete(jobID)
			if err != nil {
				log.Error("panguService: Delete Deploy failed:%s %v", jobID, err.Error())
				return nil, err
			}
		}
	} else if err != nil {
		log.Warn("DelTrainJobVersion GetCloudbrainByJobID failed, cannnot get job from DB:%s %v", jobID, err.Error())
	}

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions/" + versionID)

	if err != nil {
		return &result, fmt.Errorf("resty DelTrainJobVersion: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}

		if temp.ErrorCode == ModelArtsJobNotExists || temp.ErrorCode == ModelArtsJobInTargetState {
			//任务不存在或者已经处于被删除的状态，此时认为删除成功
			return &models.TrainJobResult{IsSuccess: true}, nil
		}
		if result.ErrorCode == ModelArtsJobInternalError {
			log.Error("ModelArt internal error when del job,jobId=%s", jobID)
			return &models.TrainJobResult{IsSuccess: true}, nil
		}
		log.Error("DelTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("删除训练作业版本失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("DelTrainJob(%s) failed", jobID)
		return &result, fmt.Errorf("删除训练作业版本失败:%s", result.ErrorMsg)
	}

	return &result, nil
}

func createInferenceJob(createJobParams models.CreateInferenceJobParams) (*models.CreateTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob)

	if err != nil {
		return nil, fmt.Errorf("resty create inference-job: %s", err)
	}

	req, _ := json.Marshal(createJobParams)
	log.Info("%s", req)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("createInferenceJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		BootFileErrorMsg := "Invalid OBS path '" + createJobParams.InfConfig.BootFileUrl + "'."
		DataSetErrorMsg := "Invalid OBS path '" + createJobParams.InfConfig.DataUrl + "'."
		if temp.ErrorMsg == BootFileErrorMsg {
			log.Error("启动文件错误！createInferenceJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("启动文件错误！")
		}
		if temp.ErrorMsg == DataSetErrorMsg {
			log.Error("数据集错误！createInferenceJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("数据集错误！")
		}
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createInferenceJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createInferenceJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	if !result.IsSuccess {
		log.Error("createInferenceJob failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("createInferenceJob failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func createInferenceJobUserImage(createJobParams models.CreateInfUserImageParams) (*models.CreateTrainJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateTrainJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlTrainJob)

	if err != nil {
		return nil, fmt.Errorf("resty create train-job: %s", err)
	}

	req, _ := json.Marshal(createJobParams)
	log.Info("%s", req)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("createInferenceJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		bootFileErrorMsg := "Invalid OBS path '" + createJobParams.Config.BootFileUrl + "'."
		dataSetErrorMsg := "Invalid OBS path '" + createJobParams.Config.DataUrl + "'."
		if temp.ErrorMsg == bootFileErrorMsg {
			log.Error("启动文件错误！createInferenceJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("启动文件错误！")
		}
		if temp.ErrorMsg == dataSetErrorMsg {
			log.Error("数据集错误！createInferenceJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
			return &result, fmt.Errorf("数据集错误！")
		}
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createInferenceJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createInferenceJobUserImage failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	if !result.IsSuccess {
		log.Error("createInferenceJobUserImage failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("createInferenceJobUserImage failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func createNotebook2(createJobParams models.CreateNotebook2Params) (*models.CreateNotebookResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateNotebookResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlNotebook2)

	if err != nil {
		return nil, fmt.Errorf("resty create notebook2: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	var response models.NotebookResult
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("json.Unmarshal failed: %s", err.Error())
	}

	if res.StatusCode() == http.StatusBadGateway {
		return &result, fmt.Errorf(UnknownErrorPrefix+"createNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	if len(response.ErrorCode) != 0 {
		log.Error("createNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
		if response.ErrorCode == errorCodeExceedLimit {
			response.ErrorMsg = "所选规格使用数量已超过最大配额限制。"
		}
		if response.ErrorCode == modelartsIllegalToken && retry < 1 {
			retry++
			_ = getToken()
			goto sendjob
		}
		return &result, fmt.Errorf("createNotebook2 failed(%s): %s", response.ErrorCode, response.ErrorMsg)
	}

	return &result, nil
}

func GetTrainJobMetricStatistic(jobID, versionID, podName string) (*models.GetTrainJobMetricStatisticResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetTrainJobMetricStatisticResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions/" + versionID + "/pod/" + podName + "/metric-statistic?statistic_type=each")

	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobMetricStatistic: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJobMetricStatistic failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("GetTrainJobMetricStatistic failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetTrainJobMetricStatistic(%s) failed", jobID)
		return &result, fmt.Errorf("获取任务资源占用情况失败:%s", result.ErrorMsg)
	}

	return &result, nil
}

func GetTrainJobList(perPage, page int, sortBy, order, searchContent string) (*models.GetTrainJobListResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetTrainJobListResult

	retry := 0

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"per_page":       strconv.Itoa(perPage),
			"page":           strconv.Itoa(page),
			"sortBy":         sortBy,
			"order":          order,
			"search_content": searchContent,
		}).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJob)

	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobList: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJobList failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf(temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetTrainJobList failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf(result.ErrorMsg)
	}

	return &result, nil
}

func GetTrainJobVersionList(perPage, page int, jobID string) (*models.GetTrainJobVersionListResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetTrainJobVersionListResult

	retry := 0

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"per_page": strconv.Itoa(perPage),
			"page":     strconv.Itoa(page),
		}).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlTrainJob + "/" + jobID + "/versions")

	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobVersionList: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJobVersionList failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf(temp.ErrorMsg)
	}

	if !result.IsSuccess {
		log.Error("GetTrainJobVersionList failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf(result.ErrorMsg)
	}

	return &result, nil
}

func GetNotebookList(limit, offset int, sortBy, order, searchContent string) (*models.GetNotebookListResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetNotebookListResult

	retry := 0

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"limit":    strconv.Itoa(limit),
			"offset":   strconv.Itoa(offset),
			"name":     searchContent,
			"sort_key": sortBy,
			"sort_dir": order,
		}).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlNotebook2)

	if err != nil {
		return nil, fmt.Errorf("resty GetNotebookList: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetNotebookList failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf(temp.ErrorMsg)
	}

	return &result, nil
}

func createDeployModel(createDeployModelParams models.CreateDeployModelParams) (*models.CreateDeployModelResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateDeployModelResult

	retry := 0
	req, _ := json.Marshal(createDeployModelParams)
	log.Info("postapi json: %s", req)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createDeployModelParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlDeployModel)

	if err != nil {
		return nil, fmt.Errorf("resty create deploy model: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("createDeployModel failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createDeployModel failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createDeployModel failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	return &result, nil
}

func GetDeployModel(modelID string) (*models.GetDeployModelResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetDeployModelResult

	retry := 0

sendjob:
	//https://modelarts.cn-east-3.myhuaweicloud.com/v1/{project_id}/models/{model_id}
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlDeployModel + "/" + modelID)

	if err != nil {
		return nil, fmt.Errorf("resty GetDeployModel: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("获取部署模型失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	return &result, nil
}

func createDeployService(CreateDeployServiceParams models.CreateDeployServiceParams) (*models.CreateDeployServiceResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateDeployServiceResult

	retry := 0
	req, _ := json.Marshal(CreateDeployServiceParams)
	log.Info("postapi json: %s", req)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(CreateDeployServiceParams).
		SetResult(&result).
		Post(HOST + "/v1/" + setting.ProjectID + urlDeployService)

	if err != nil {
		return nil, fmt.Errorf("resty create deploy model: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("createDeployModel failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		if res.StatusCode() == http.StatusBadGateway {
			return &result, fmt.Errorf(UnknownErrorPrefix+"createDeployModel failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		} else {
			return &result, fmt.Errorf("createDeployModel failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		}
	}

	return &result, nil
}

func GetDeployService(serviceID string) (*models.GetDeployServiceResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetDeployServiceResult

	retry := 0

sendjob:
	//https://modelarts.cn-east-3.myhuaweicloud.com/v1/{project_id}/models/{model_id}
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + "/v1/" + setting.ProjectID + urlDeployService + "/" + serviceID)

	if err != nil {
		return nil, fmt.Errorf("resty GetDeployService: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return &result, fmt.Errorf("获取部署模型失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	return &result, nil
}

func DeleteDeployModel(modelID string) error {
	checkSetting()
	client := getRestyClient()
	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		Delete(HOST + "/v1/" + setting.ProjectID + urlDeployModel + "/" + modelID)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("DelTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return fmt.Errorf("删除部署模型失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	return nil
}

func DeleteDeployService(serviceID string) error {
	checkSetting()
	client := getRestyClient()

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		Delete(HOST + "/v1/" + setting.ProjectID + urlDeployService + "/" + serviceID)

	if err != nil {
		return fmt.Errorf("resty DelDeployService: %v", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("DelTrainJob failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return fmt.Errorf("删除部署任务失败(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	return nil
}

func UpdateDeployService(serviceID string, UpdateDeployServiceParams models.UpdateDeployServiceParams) error {
	checkSetting()
	client := getRestyClient()

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(UpdateDeployServiceParams).
		Put(HOST + "/v1/" + setting.ProjectID + urlDeployService + "/" + serviceID)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.ErrorResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			//return &result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
	}

	return nil
}

// https://modelarts-inference.cloudbrain2.pcl.ac.cn/v1/infers/cc1e11c5-16a0-4202-a031-819a8e178a0d
func SendInferenceDeploy(inferAddr string, text string) (*models.PanguInferResult, error) {
	panguInferParams := &models.PanguInferParams{
		Text: text,
	}
	checkSetting()
	client := getRestyClient()
	var result models.PanguInferResult

	retry := 0
	req, _ := json.Marshal(panguInferParams)
	log.Info("盘古微调部署：postapi json: ", req)
	log.Info("盘古微调部署 推理地址: ", inferAddr)
	log.Info("盘古微调部署 推理token: ", TOKEN)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Auth-Token", TOKEN).
		SetAuthToken(TOKEN).
		SetBody(panguInferParams).
		SetResult(&result).
		Post(inferAddr)
	//Post(HOST + "/v1/" + urlDeployInfer + serviceID)

	if err != nil {
		return nil, fmt.Errorf("resty post text inference failed: %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		var erroResult models.PanguInferError
		if err = json.Unmarshal([]byte(res.String()), &erroResult); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("盘古微调部署 inference failed(%d):%s", res.StatusCode(), res.String())
		log.Error("盘古微调部署 %s, %s", erroResult.ErrorCode, erroResult.ErrorMsg)
		return &result, fmt.Errorf("盘古微调部署 部署任务推理失败(%d): %s", res.StatusCode(), erroResult.ErrorMsg)
	}
	log.Info("盘古微调部署 推理返回状态码: %s", res.StatusCode())
	log.Info("盘古微调部署 推理结果: %s", result)

	return &result, nil
}

func ServiceDelete(jobID string) error {
	if deploy, _ := models.GetModelartsDeployByJobID(jobID); deploy != nil {
		if deploy.Status == "STOP" || deploy.Status == "FAILED" {

			if deploy.ServiceID != "" {
				if err := DeleteDeployService(deploy.ServiceID); err != nil {
					log.Error("panguService: Delete DeployService API failed:%s %v", jobID, err.Error())
				} else {
					log.Info("panguService: deploy service delete success %s", jobID)
				}
			}

			if deploy.ModelID != "" {
				if err := DeleteDeployModel(deploy.ModelID); err != nil {
					log.Error("panguService: Delete DeployModel API failed:%s %v", jobID, err.Error())
				} else {
					log.Info("panguService: deploy model delete success %s", jobID)
				}
			}

			if err := models.DeleteModelartsDeploy(jobID); err != nil {
				log.Error("panguService: Delete ModelartsDeploy from DB failed:%s %v", jobID, err.Error())
				return err
			} else {
				log.Info("panguService: deploy DB record delete success %s", jobID)
			}

		} else {
			log.Error("the job(%s) is a deploying finetune job, can be not deleted", jobID)
			return fmt.Errorf("1")
		}
	}
	return nil
}
