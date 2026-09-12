package aim

import (
	"crypto/tls"
	"fmt"

	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
)

var (
	restyClient *resty.Client
)

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

type Overview struct {
	Url string `json:"run_overview_url"`
}

type Compare struct {
	Url string `json:"compare_url"`
}

func GetDetailUrl(job_name string) (string, error) {

	client := getRestyClient()
	var result Overview
	urlGet := setting.AimConfig.DetailUrl + "?query=" + job_name
	resp, err := client.R().
		SetResult(&result).
		Get(urlGet)
	if err != nil {
		return "", err
	}

	// 检查响应状态码是否为 200 OK
	if !resp.IsSuccess() {
		return "", fmt.Errorf("HTTP Error: %d", resp.StatusCode())
	}

	return result.Url, nil
}

func GetComparelUrl(job_names []string) (string, error) {

	client := getRestyClient()
	var result Compare

	queryStr := "["

	for i, job_name := range job_names {
		if i == 0 {
			queryStr += "'" + job_name + "'"
		} else {
			queryStr += ",'" + job_name + "'"
		}
	}
	queryStr = queryStr + "]"
	urlGet := setting.AimConfig.CompareUrl + "?query=" + queryStr
	resp, err := client.R().
		SetResult(&result).
		Get(urlGet)
	if err != nil {
		return "", err
	}

	// 检查响应状态码是否为 200 OK
	if !resp.IsSuccess() {
		return "", fmt.Errorf("HTTP Error: %d", resp.StatusCode())
	}

	return result.Url, nil
}
