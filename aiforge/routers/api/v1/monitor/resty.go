package monitor

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"github.com/go-resty/resty/v2"
)

var (
	restyClient *resty.Client
)

const (
	urlSummary      = "data_summary"
	urlSearchAlerts = "data_alerts"
)

type ComputeSummary struct {
	ComputeResource string `json:"compute_resource"`
	Value           int    `json:"value"`
}

type AlertSummary struct {
	Updated   int64             `json:"updated"`
	XPU       int               `json:"xpu"`
	XPUDetail []*ComputeSummary `json:"xpu_detail"`
	In        int               `json:"in"`
	InDetail  []*ComputeSummary `json:"in_detail"`
	Out       int               `json:"out"`
	OutDetail []*ComputeSummary `json:"out_detail"`
}

type AlertSummaryResponse struct {
	Summary AlertSummary `json:"alert_summary"`
}

type Alert struct {
	ID              int64                      `json:"id"`
	Discription     string                     `json:"discription"`
	AlertName       string                     `json:"alert_name"`
	TaskID          string                     `json:"task_id"`
	TaskName        string                     `json:"display_job_name"`
	JobType         string                     `json:"job_type"`
	OwnerName       string                     `json:"owner_name"`
	RepoName        string                     `json:"repo_name"`
	Duration        string                     `json:"duration"`
	Email           string                     `json:"email"`
	ComputeResource string                     `json:"compute_resource"`
	AlertPolicyID   int64                      `json:"alert_policy_id"`
	CreatedUnix     int64                      `json:"created"`
	Status          string                     `json:"status"`
	CardType        string                     `json:"card_type"`
	AICenter        string                     `json:"ai_center"`
	UserName        string                     `json:"user_name"`
	JobName         string                     `json:"job_name"`
	Spec            *structs.SpecificationShow `json:"spec"`
}

type DataAlert struct {
	Alerts []*Alert `json:"alerts"`
	Total  int      `json:"total"`
}

type AlertsResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data DataAlert `json:"data"`
}
type SearchAlertsOption struct {
	models.ListOptions
	AlertType string
	JobType   string
	Keyword   string
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

func GetCurrentMonitorSummary() (*AlertSummaryResponse, error) {

	client := getRestyClient()
	var result AlertSummaryResponse

	_, err := client.R().
		SetResult(&result).
		Get(setting.MonitorTaskHost + urlSummary)

	if err != nil {
		return nil, fmt.Errorf("resty GetCurrentMonitorSummary: %v", err)
	}

	return &result, nil
}

func GetTaskAlerts(opt SearchAlertsOption) (*AlertsResponse, error) {

	client := getRestyClient()
	var result AlertsResponse

	reqBaseUrl := setting.MonitorTaskHost + urlSearchAlerts + "?"
	
	if opt.AlertType != "" {
		reqBaseUrl += "&alert_type=" + opt.AlertType
	}
	if opt.JobType != "" {
		reqBaseUrl += "&job_type=" + opt.JobType
	}
	if opt.Keyword != "" {
		reqBaseUrl += "&q=" + opt.Keyword
	}
	if opt.Page > 0 {
		reqBaseUrl += "&page=" + strconv.Itoa(opt.Page)
	}
	if opt.PageSize > 0 {
		reqBaseUrl += "&page_size=" + strconv.Itoa(opt.PageSize)
	}

	_, err := client.R().
		SetResult(&result).
		Get(reqBaseUrl)

	if err != nil {
		return nil, fmt.Errorf("resty GetTaskAlerts: %v", err)
	}

	return &result, nil

}
