package sd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/ai_task"
	"code.gitea.io/gitea/services/ai_task_service/task"
	"github.com/go-resty/resty/v2"
)

const (
	StageAPI = "/stage"
	TrainAPI = "/train_lora"
	InferAPI = "/text2img"
)

var restyClient *resty.Client

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

func tokenUrlByTaskId(cloudbrainId int64) string {
	url, err := ai_task.GetSelfEndPointUrlById(cloudbrainId)
	if err != nil {
		log.Error("GetSelfEndPointUrlById error" + fmt.Sprint(err))
		return ""
	}
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	log.Info("sd-finetune url= " + url)
	return url
}

type TaskStage struct {
	Stage         int      `json:"stage"`
	Desc          string   `json:"desc"`
	Pretrainmodel string   `json:"pretrainmodel"`
	Total_Steps   int      `json:"total_steps"`
	Total_Epochs  int      `json:"total_epochs"`
	LoraModels    []string `json:"lora_models"`
	LoraSamples   []string `json:"lora_samples"`
}

func getTaskStageResty(tokenUrl string) (*TaskStage, int, error) {
	var result TaskStage

	url := tokenUrl + StageAPI
	client := getRestyClient()
	retry := 0

sendjob:
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&result).
		Get(url)

	if err != nil {
		log.Error("request failed: %s; url: %s", err.Error(), url)
		return nil, 0, err
	}

	if resp.StatusCode() != 200 {
		if retry < 1 {
			retry++
			goto sendjob
		} else {
			log.Error("resposne failed %d %s; resp-body: %s", resp.StatusCode(), url, resp.Body())
			return nil, resp.StatusCode(), fmt.Errorf(resp.String())
		}
	}

	return &result, resp.StatusCode(), nil
}

func getTaskInfo(id int64) (*entity.AITaskBriefInfo, error) {
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("failed GetAITaskTemplateByCloudbrainId: %s", err.DefaultMsg)
		return nil, err.ToError()
	}

	info, err := t.BriefQuery(id)
	if err != nil {
		log.Error("failed BriefQuery: %s", err.DefaultMsg)
		return nil, err.ToError()
	}

	return info, nil
}

func GetTaskStage(ctx *context.APIContext) {
	cloudbrainId := ctx.QueryInt64("task_id")
	tokenUrl := tokenUrlByTaskId(cloudbrainId)

	stage, statusCode, err := getTaskStageResty(tokenUrl)
	if err != nil {
		log.Error("GetTaskStage failed to get /stage; url: %s; err: %s", tokenUrl, err.Error())
		ctx.JSON(200, map[string]interface{}{
			"url":  "",
			"msg":  "failed to get /stage from service:" + err.Error(),
			"code": statusCode,
		})
		return
	}

	task_info, err := getTaskInfo(cloudbrainId)
	if err != nil {
		log.Error("GetTaskStage failed to get task info; id: %s", cloudbrainId)
		ctx.JSON(200, map[string]interface{}{
			"url":  "",
			"msg":  "failed to get task info:" + err.Error(),
			"code": statusCode,
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"url":           tokenUrl,
		"stage":         stage.Stage,
		"desc":          stage.Desc,
		"pretrianmodel": stage.Pretrainmodel,
		"total_steps":   stage.Total_Steps,
		"total_epochs":  stage.Total_Epochs,
		"lora_models":   stage.LoraModels,
		"lora_samples":  stage.LoraSamples,
		"task":          task_info,
	})
}

type SdServerResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Stage TaskStage   `json:"stage"`
}

type UserTrainConfig struct {
	// Goes in train config toml
	MaxTrainEpochs              int     `json:"max_train_epochs"`
	SaveEveryNEpochs            int     `json:"save_every_n_epochs"`
	SavePrecision               string  `json:"save_precision"`
	MixedPrecision              string  `json:"mixed_precision"`
	SampleSampler               string  `json:"sample_sampler"`
	LearningRate                float64 `json:"learning_rate"`
	UnetLR                      float64 `json:"unet_lr"`
	TextEncoderLR               float64 `json:"text_encoder_lr"`
	LRScheduler                 string  `json:"lr_scheduler"`
	OptimizerType               string  `json:"optimizer_type"`
	LRSchedulerNumCycles        int     `json:"lr_scheduler_num_cycles"`
	NetworkDim                  int     `json:"network_dim"`
	NetworkAlpha                float64 `json:"network_alpha"`
	KeepTokens                  int     `json:"keep_tokens"`
	MaxTokenLength              int     `json:"max_token_length"`
	Seed                        int     `json:"seed"`
	ClipSkip                    int     `json:"clip_skip"`
	NetworkTrainUnetOnly        bool    `json:"network_train_unet_only"`
	NetworkTrainTextEncoderOnly bool    `json:"network_train_text_encoder_only"`
	// optional params
	LrWarmupSteps           int     `json:"lr_warmup_steps,omitempty"`
	MinSnrGamma             int     `json:"min_snr_gamma,omitempty"`
	NoiseOffset             float64 `json:"noise_offset,omitempty"`
	MultiresNoiseIterations int     `json:"multires_noise_iterations,omitempty"`
	MultiresNoiseDiscount   float64 `json:"multires_noise_discount,omitempty"`
}

type UserDatasetConfig struct {
	// Goes in dataset json
	NumRepeats int    `json:"num_repeats"`
	BatchSize  int    `json:"batch_size"`
	Resolution [2]int `json:"resolution"`
}

type UserPromptsConfig struct {
	// Goes in prompts json
	Prompts    string `json:"prompts"`
	Negative   string `json:"negative"`
	Seed       int    `json:"seed"`
	Resolution [2]int `json:"resolution"`
}

type UserConfig struct {
	TrainConfig   UserTrainConfig   `json:"TrainConfig"`
	DatasetConfig UserDatasetConfig `json:"DatasetConfig"`
	PromptsConfig UserPromptsConfig `json:"PromptsConfig"`
}

func sendTrainLoraResty(tokenUrl string, userConfig UserConfig) (*SdServerResponse, int, error) {
	var result SdServerResponse

	url := tokenUrl + TrainAPI
	client := getRestyClient()

	configJSON, err := json.Marshal(userConfig)
	if err != nil {
		log.Error("[sd-finetune] Failed to marshal userConfig: %s", err.Error())
		return nil, 0, err
	}
	log.Info("userConfig: %s", string(configJSON))

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(userConfig).
		SetResult(&result).
		Post(url)

	if err != nil {
		log.Error("[sd-finetune] %s; url: %s", err.Error(), url)
		return nil, 0, err
	}

	if resp.StatusCode() != 200 {
		log.Error("[sd-finetune] url %d %s; resp-body: %s", resp.StatusCode(), url, resp.Body())
		return nil, resp.StatusCode(), fmt.Errorf(resp.String())
	}

	return &result, resp.StatusCode(), nil
}

func SendTrainLora(ctx *context.APIContext, config UserConfig) {
	cloudbrainId := ctx.QueryInt64("task_id")
	tokenUrl := tokenUrlByTaskId(cloudbrainId)

	configJSON, err := json.Marshal(config)
	if err != nil {
		log.Error("[sd-finetune] Failed to marshal userConfig: %s", err.Error())
	}
	log.Info("userConfig: %s", string(configJSON))

	res, statusCode, err := sendTrainLoraResty(tokenUrl, config)
	if err != nil {
		log.Error("[sd-finetune] failed to get /train_lora from service; url: %s; err: %s", tokenUrl, err.Error())
		ctx.JSON(200, map[string]interface{}{
			"url":  tokenUrl,
			"msg":  "failed to start train from service" + err.Error(),
			"code": statusCode,
		})
		return
	}

	ctx.JSON(200, res)
}

type UserInferPrompt struct {
	Prompt         string  `json:"prompt"`
	NegativePrompt string  `json:"negative_prompt"`
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	Seed           int     `json:"seed"`
	Steps          int     `json:"steps"`
	CfgScale       float64 `json:"cfg_scale"`
	LoraWeights    string  `json:"lora_weights"`
	Guidance       float64 `json:"guidance,omitempty"`
}

func sendText2ImgResty(tokenUrl string, userConfig UserInferPrompt) (*SdServerResponse, int, error) {
	var result SdServerResponse

	url := tokenUrl + InferAPI
	client := getRestyClient()

	configJSON, err := json.Marshal(userConfig)
	if err != nil {
		log.Error("[sd-finetune] Failed to marshal userConfig: %s", err.Error())
		return nil, 0, err
	}
	log.Info("userConfig: %s", string(configJSON))

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(userConfig).
		SetResult(&result).
		Post(url)

	if err != nil {
		log.Error("[sd-finetune] %s; url: %s", err.Error(), url)
		return nil, 0, err
	}

	if resp.StatusCode() != 200 {
		log.Error("[sd-finetune] url %d %s; resp-body: %s", resp.StatusCode(), url, resp.Body())
		return nil, resp.StatusCode(), fmt.Errorf(resp.String())
	}

	return &result, resp.StatusCode(), nil
}

func SendText2Img(ctx *context.APIContext, config UserInferPrompt) {
	cloudbrainId := ctx.QueryInt64("task_id")
	tokenUrl := tokenUrlByTaskId(cloudbrainId)

	configJSON, err := json.Marshal(config)
	if err != nil {
		log.Error("[sd-finetune] Failed to marshal userConfig: %s", err.Error())
	}
	log.Info("userConfig: %s", string(configJSON))

	res, statusCode, err := sendText2ImgResty(tokenUrl, config)
	if err != nil {
		log.Error("[sd-finetune] failed to get /text2img from service; url: %s; err: %s", tokenUrl, err.Error())
		ctx.JSON(200, map[string]interface{}{
			"url":  tokenUrl,
			"msg":  "failed to start image generation from service" + err.Error(),
			"code": statusCode,
		})
		return
	}

	ctx.JSON(200, res)
}
