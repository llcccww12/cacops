package modelappservice

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"github.com/go-resty/resty/v2"
)

type CreateSDParams struct {
	Prompt                string `json:"prompt"`
	Height                int    `json:"height"`
	Width                 int    `json:"width"`
	Seed                  int    `json:"seed"`
	Steps                 int    `json:"steps"`
	Num_images_per_prompt int    `json:"num_images_per_prompt"`
	Negative_prompt       string `json:"negative_prompt"`
	Scheduler_name        string `json:"scheduler_name"`
	Guidance_scale        int    `json:"guidance_scale"`
}

type SDResult struct {
	Image_base64_list []string `json:"image_base64_list"`
}

var (
	sdRestyClient *resty.Client
)

func getRestyClient() *resty.Client {
	if sdRestyClient == nil {
		sdRestyClient = resty.New()
		sdRestyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return sdRestyClient
}

func CreateSDJob(modelapp *models.ModelApp, url string) (*SDResult, error, int) {
	createSDParams := &CreateSDParams{
		Prompt:          modelapp.Desc,
		Negative_prompt: modelapp.Negative_prompt,
	}

	if modelapp.Num_images_per_prompt > 0 {
		createSDParams.Num_images_per_prompt = modelapp.Num_images_per_prompt
	}

	if modelapp.Seed > 0 {
		createSDParams.Seed = modelapp.Seed
	}
	if modelapp.Steps > 0 {
		createSDParams.Steps = modelapp.Steps
	}
	if modelapp.Height > 0 {
		createSDParams.Height = modelapp.Height
	}
	if modelapp.Width > 0 {
		createSDParams.Width = modelapp.Width
	}
	if modelapp.Guidance_scale > 0 {
		createSDParams.Guidance_scale = modelapp.Guidance_scale
	}
	if modelapp.Scheduler_name != "" {
		createSDParams.Scheduler_name = modelapp.Scheduler_name
	}
	client := getRestyClient()
	var result SDResult
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(createSDParams).
		SetResult(&result).
		Post(url)

	if err != nil {
		return nil, fmt.Errorf("resty CreateSDJob: %s", err), 0
	}
	log.Info("re status=" + res.Status() + " status code =" + fmt.Sprint(res.StatusCode()))
	if res.StatusCode() == 200 {
		var response SDResult
		err = json.Unmarshal(res.Body(), &response)
		if err != nil {
			log.Error("json.Unmarshal failed: %s", err.Error())
			return &result, fmt.Errorf("son.Unmarshal failed: %s", err.Error()), 0
		}
		return &result, nil, 200
	} else {
		return nil, fmt.Errorf(res.Status()), res.StatusCode()
	}
}
