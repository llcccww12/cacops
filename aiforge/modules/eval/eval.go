package eval

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"github.com/go-resty/resty/v2"
)

var restyClient *resty.Client

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		restyClient.SetTimeout(time.Duration(10 * time.Minute))
	}
	return restyClient
}

func EvalTask(req *structs.EvalRequest) (*structs.EvalResultResponse, error) {
	var result structs.EvalResultResponse
	restyClient := getRestyClient()

	_, err := restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post(setting.EvalService.ServerUrl + "/evaluation")
	return &result, err

}

func GetEvalTaskLog(taskId int64) (string, error) {

	restyClient := getRestyClient()
	req := &structs.TaskRequest{TaskId: taskId}
	var logContent string
	res, err := restyClient.R().
		SetBody(req).
		SetResult(&logContent).
		Post(setting.EvalService.ServerUrl + "/log")

	if err != nil {
		return logContent, fmt.Errorf("resty GetEvalJobLog: %v", err)
	}

	if res.StatusCode() != http.StatusOK {

		return logContent, fmt.Errorf("GetEvalJobLog failed(%d)", res.StatusCode())
	}

	logContent = res.String()

	return logContent, nil

}
func ClearEvalTaskResult(taskId int64) error {
	req := structs.TaskRequest{TaskId: taskId}
	res, err := getRestyClient().R().
		SetBody(req).
		Post(setting.EvalService.ServerUrl + "/delete")
	if err != nil {
		return err
	}

	if res.StatusCode() != http.StatusOK {

		return fmt.Errorf("clear EvalJob failed(%d)", res.StatusCode())
	}
	return nil

}

func GetEvalTaskResult(taskId int64) (*structs.EvalResultResponse, error) {
	req := structs.TaskRequest{TaskId: taskId}
	var result structs.EvalResultResponse
	res, err := getRestyClient().R().
		SetBody(req).
		SetResult(&result).
		Post(setting.EvalService.ServerUrl + "/result")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {

		return nil, fmt.Errorf("clear EvalJob failed(%d)", res.StatusCode())
	}
	return &result, nil

}

func GetEvalTaskDetailResult(taskRequest structs.TaskDetailRequest) (*structs.EvalDetailResultResponse, error) {
	
	var result structs.EvalDetailResultResponse
	res, err := getRestyClient().R().
		SetBody(taskRequest).
		SetResult(&result).
		Post(setting.EvalService.ServerUrl + "/detail")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {

		return nil, fmt.Errorf("get eval job detail result failed(%d)", res.StatusCode())
	}
	return &result, nil

}
