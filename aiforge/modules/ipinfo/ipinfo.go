package ipinfo

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"code.gitea.io/gitea/modules/setting"

	"github.com/go-resty/resty/v2"
)

var restyClient *resty.Client

type IpInfoResponse struct {
	Ip    string `json:"ip"`
	Loc   string `json:"loc"`
	Bogon bool   `json:"bogon"`
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

func GetLocationByIp(ip string) (*IpInfoResponse, error) {
	client := getRestyClient()
	var result IpInfoResponse
	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(setting.IPInfo.Token).
		SetResult(&result).
		Get(setting.IPInfo.Host + "/" + ip)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("http status is %d", res.StatusCode())
	}
	return &result, nil
}
