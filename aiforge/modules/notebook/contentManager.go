package notebook

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/modules/log"
)

var restyClient *resty.Client

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

type NotebookApiResponse struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type NotebookContent struct {
	Url      string
	Path     string
	Cookies  []*http.Cookie
	Xsrf     string
	PathType string //file  directory
	Token    string
}

func (c *NotebookContent) IsNotebookFileCanBrowser() bool {
	if c.Xsrf == "" {
		c.SetCookiesAndCsrf()
	}
	if c.Xsrf == "" {
		log.Warn("xsrf is empty, can not broswer url:" + c.Url)
		return false
	}
	return c.IsNoteBookContentsExist()

}

func (c *NotebookContent) SetCookiesAndCsrf() {
	log.Info("jupyter url:" + c.Url)
	if c.Xsrf != "" {
		return
	}
	var cookies []*http.Cookie
	const retryTimes = 10
	url := c.Url
	if c.Token != "" {
		url = c.Url + "?token=" + c.Token
	}
	for i := 0; i < retryTimes; i++ {
		if c.Xsrf == "" {
			c.setCookieAndCsrf(url, i, retryTimes, cookies)
			if c.Xsrf != "" {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}

}

func (c *NotebookContent) setCookieAndCsrf(url string, i int, retryTimes int, cookies []*http.Cookie) {
	client := http.Client{Transport: transport}
	res, err := client.Get(url)
	if err != nil {
		log.Error("browser jupyterUrl failed.", err)
		if i == retryTimes-1 {
			c.Cookies = cookies
		}
		return

	}
	defer res.Body.Close()
	cookies = res.Cookies()
	xsrf := ""
	for _, cookie := range cookies {
		if cookie.Name == "_xsrf" {
			xsrf = cookie.Value
			if len(cookies) > 1 {
				break
			}

		}

	}
	if xsrf != "" {
		c.Cookies = cookies
		c.Xsrf = xsrf
	}

}

func (c *NotebookContent) IsNoteBookContentsExist() bool {
	client := getRestyClient()
	uploadUrl := getJupyterBaseUrl(c.Url) + "api/contents/" + c.Path + "?type=" + c.PathType
	res, err := client.R().
		SetCookies(c.Cookies).
		SetHeader("X-XSRFToken", c.Xsrf).
		Get(uploadUrl)
	if err != nil {
		log.Warn("browser url error:"+uploadUrl, err)
		return false
	}
	return res.StatusCode() == http.StatusOK
}

func (c *NotebookContent) UploadNoteBookFile(task *models.Cloudbrain) error {

	err := c.MakeNoteBookDir()
	if err != nil {
		return err
	}

	codePath := setting.JobPath + task.JobName + cloudbrain.CodeMountPath
	fileContents, err := ioutil.ReadFile(codePath + "/" + c.Path)
	if err != nil {
		log.Error("read jupyter file failed:%v", task.DisplayJobName, err)
	}

	base64Content := base64.StdEncoding.EncodeToString(fileContents)
	client := getRestyClient()
	uploadUrl := getJupyterBaseUrl(c.Url) + "api/contents/" + c.Path
	res, err := client.R().
		SetCookies(c.Cookies).
		SetHeader("X-XSRFToken", c.Xsrf).
		SetBody(map[string]interface{}{
			"type":    "file",
			"format":  "base64",
			"name":    path.Base(c.Path),
			"path":    c.Path,
			"content": base64Content}).
		Put(uploadUrl)
	if err != nil {
		log.Error("upload jupyter file failed:%v", task.DisplayJobName, err)
		return err
	} else if res.StatusCode() != http.StatusCreated {
		log.Error("upload jupyter file failed:%v, status is %s", task.DisplayJobName, res.Status(), err)
		return fmt.Errorf("status:", res.StatusCode())
	}
	return nil
}

/*
*

	if c.Path is   a/b/c.txt
	makedir  a/b
	if c.Path is   a/b/c
	makedir  a/b
*/
func (c *NotebookContent) MakeNoteBookDir() error {
	filePaths := strings.Split(c.Path, "/")

	for i := 0; i < len(filePaths)-1; i++ {
		cTemp := &NotebookContent{
			Url:      c.Url,
			Cookies:  c.Cookies,
			Path:     path.Join(filePaths[0 : i+1]...),
			PathType: "directory",
			Xsrf:     c.Xsrf,
		}
		if !cTemp.IsNoteBookContentsExist() {

			createTempDirUrl := getJupyterBaseUrl(cTemp.Url) + "api/contents/" + cTemp.Path
			client := getRestyClient()
			var jobResult NotebookApiResponse
			res, err := client.R().
				SetCookies(c.Cookies).
				SetHeader("X-XSRFToken", c.Xsrf).
				SetBody(map[string]interface{}{
					"type": cTemp.PathType,
					"path": cTemp.Path,
				}).SetResult(&jobResult).
				Put(createTempDirUrl)
			if err != nil {
				return err
			}
			if res.StatusCode() != http.StatusCreated {
				return fmt.Errorf("status code:" + res.Status())
			}

		}

	}
	return nil
}

func getJupyterBaseUrl(url string) string {
	jupyterUrlLength := len(url)
	baseUrl := url
	if strings.HasSuffix(url, "lab") {
		baseUrl = url[0 : jupyterUrlLength-len(path.Base(url))]
	}
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl = baseUrl + "/"
	}

	return baseUrl
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}
