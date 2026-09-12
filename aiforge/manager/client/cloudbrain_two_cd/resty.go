package cloudbrain_two_cd

import (
	"bytes"
	"code.gitea.io/gitea/modules/modelarts_gateway/core"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

var (
	httpClient         *http.Client
	HOST               string
	TOKEN              string
	autoStopDurationMs = 4 * 60 * 60 * 1000
)

const (
	errorCodeExceedLimit = "ModelArts.0118"

	//notebook 2.0
	urlNotebook2 = "/notebooks"

	//error code
	modelartsIllegalToken     = "ModelArts.6401"
	NotebookNotFound          = "ModelArts.6404"
	NotebookNoPermission      = "ModelArts.6407"
	NotebookInvalid           = "ModelArts.6400"
	UnknownErrorPrefix        = "UNKNOWN:"
	ModelArtsJobNotExists     = "ModelArts.0102"
	ModelArtsJobInTargetState = "ModelArts.6357"
	ModelArtsJobInternalError = "ModelArts.0010"
)

func getHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout:   30 * time.Second,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}
	}
	return httpClient
}

func GetNotebook(jobID string) (*models.GetNotebook2Result, error) {
	var result models.GetNotebook2Result

	client := getHttpClient()
	s := core.Signer{
		Key:    setting.ModelartsCD.AccessKey,
		Secret: setting.ModelartsCD.SecretKey,
	}
	r, _ := http.NewRequest(http.MethodGet,
		setting.ModelartsCD.EndPoint+"/v1/"+setting.ModelartsCD.ProjectID+urlNotebook2+"/"+jobID,
		nil)

	r.Header.Add("content-type", "application/json")
	s.Sign(r)

	resp, err := client.Do(r)
	if err != nil {
		log.Error("client.Do failed: %s", err.Error())
		return &result, fmt.Errorf("client.Do failed: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll failed: %s", err.Error())
		return &result, fmt.Errorf("ioutil.ReadAll failed: %s", err.Error())
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(result.ErrorCode) != 0 {
		log.Error("GetNotebook failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetNotebook failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func ManageNotebook(jobID string, param models.NotebookAction, autoStopDuration int) (*models.NotebookActionResult, error) {
	var result models.NotebookActionResult

	client := getHttpClient()
	s := core.Signer{
		Key:    setting.ModelartsCD.AccessKey,
		Secret: setting.ModelartsCD.SecretKey,
	}
	r, _ := http.NewRequest(http.MethodPost,
		setting.ModelartsCD.EndPoint+"/v1/"+setting.ModelartsCD.ProjectID+urlNotebook2+"/"+jobID+"/"+param.Action+"?duration="+strconv.Itoa(autoStopDuration),
		nil)

	r.Header.Add("content-type", "application/json")
	s.Sign(r)

	resp, err := client.Do(r)
	if err != nil {
		log.Error("client.Do failed: %s", err.Error())
		return &result, fmt.Errorf("client.Do failed: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll failed: %s", err.Error())
		return &result, fmt.Errorf("ioutil.ReadAll failed: %s", err.Error())
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(result.ErrorCode) != 0 {
		log.Error("ManageNotebook2 failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("ManageNotebook failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func DelNotebook(jobID string) (*models.NotebookDelResult, error) {
	var result models.NotebookDelResult

	client := getHttpClient()
	s := core.Signer{
		Key:    setting.ModelartsCD.AccessKey,
		Secret: setting.ModelartsCD.SecretKey,
	}

	r, _ := http.NewRequest(http.MethodDelete,
		setting.ModelartsCD.EndPoint+"/v1/"+setting.ModelartsCD.ProjectID+urlNotebook2+"/"+jobID,
		nil)

	r.Header.Add("content-type", "application/json")
	s.Sign(r)

	resp, err := client.Do(r)
	if err != nil {
		log.Error("client.Do failed: %s", err.Error())
		return &result, fmt.Errorf("client.Do failed: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll failed: %s", err.Error())
		return &result, fmt.Errorf("ioutil.ReadAll failed: %s", err.Error())
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error())
	}

	if len(result.ErrorCode) != 0 {
		log.Error("DelNotebook2 failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		if result.ErrorCode == ModelArtsJobNotExists || result.ErrorCode == ModelArtsJobInTargetState {
			//任务不存在或者已经处于被删除的状态，此时认为删除成功
			return &models.NotebookDelResult{}, nil
		}

		if result.ErrorCode == ModelArtsJobInternalError {
			log.Error("ModelArt internal error when del job,jobId=%s", jobID)
			return &models.NotebookDelResult{}, nil
		}
		return &result, fmt.Errorf("DelNotebook2 failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func CreateNotebook(createJobParams models.CreateNotebookWithoutPoolParams) (*models.CreateNotebookResult, error) {
	var result models.CreateNotebookResult
	client := getHttpClient()
	s := core.Signer{
		Key:    setting.ModelartsCD.AccessKey,
		Secret: setting.ModelartsCD.SecretKey,
	}

	req, _ := json.Marshal(createJobParams)
	r, _ := http.NewRequest(http.MethodPost,
		setting.ModelartsCD.EndPoint+"/v1/"+setting.ModelartsCD.ProjectID+urlNotebook2,
		ioutil.NopCloser(bytes.NewBuffer(req)))

	r.Header.Add("content-type", "application/json")
	s.Sign(r)

	resp, err := client.Do(r)
	if err != nil {
		log.Error("client.Do failed: %s", err.Error())
		return &result, fmt.Errorf("client.Do failed: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll failed: %s", err.Error())
		return &result, fmt.Errorf("ioutil.ReadAll failed: %s", err.Error())
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error("json.Unmarshal failed: %s", err.Error())
		return &result, fmt.Errorf("json.Unmarshal failed: %s", err.Error())
	}

	if len(result.ErrorCode) != 0 {
		log.Error("createNotebook failed(%s): %s", result.ErrorCode, result.ErrorMsg)
		if result.ErrorCode == errorCodeExceedLimit {
			result.ErrorMsg = "所选规格使用数量已超过最大配额限制。"
		}
		return &result, fmt.Errorf("createNotebook failed(%s): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}
