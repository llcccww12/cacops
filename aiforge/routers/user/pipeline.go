package user

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/modules/setting"
)

var restyClient *resty.Client

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

type pipelineResult struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func NotifyPipelineUserLogout(uid int64) {
	if setting.MLOPS && setting.MlopsHost != "" && setting.MlopsToken != "" {
		client := getRestyClient()
		var result pipelineResult
		for i := 0; i < 4; i++ {
			if i != 0 {
				time.Sleep(10 * time.Second)
			}
			res, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetResult(&result).
				Get(setting.MlopsHost + "/pch-bff-pipeline/api/v1/pengcheng/logout?userId=" + strconv.FormatInt(uid, 10) + "&token=" + setting.MlopsToken)

			if err != nil {
				log.Warn("log out notify pipeline failed:", err)
			} else if res.StatusCode() == http.StatusOK && result.Code == 0 {
				break
			}

		}
	}
}
