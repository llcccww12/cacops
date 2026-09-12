package setting

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"code.gitea.io/gitea/modules/log"
)

type PanguFineTune struct {
	Basic struct {
		OwnerName  string `json:"owner_name"`
		RepoName   string `json:"repo_name"`
		BootFile   string `json:"boot_file"`
		BranchName string `json:"branch_name"`
		ImageId    int64  `json:"image_id"`
		Image      string `json:"image"`
	}
	Dataset struct {
		Attachments  []string `json:"attachments"`
		DatasetNames []string `json:"dataset_names"`
	}

	Model struct {
		ModelNames       []string `json:"model_names"`
		ModelAttachments []string `json:"model_attachments"`
	}
	Deploy struct {
		CodeBucket       string `json:"code_bucket"`
		CodeObsPath      string `json:"code_obs_path"`
		ModelName        string `json:"deploy_model_name"`
		Runtime          string `json:"runtime"`
		MaxDeployNum     int    `json:"max_deploy_num"`
		MaxDeployPerUser int    `json:"max_deploy_per_user"`
		Duration         int    `json:"duration"`
		WarmupDuration   int    `json:"warmup_duration"`
		TimeUnit         string `json:"time_unit"`
		Specification    string `json:"specification"`
	}
	Service struct {
		ClusterID     string `json:"cluster_id"`
		CPU           int    `json:"cpu"`
		CPUArch       string `json:"cpu_arch"`
		Memory        int    `json:"memory"`
		MemoryInfo    int    `json:"memory_info"`
		MemoryUnit    string `json:"memory_unit"`
		NPU           int    `json:"npu"`
		NPUBrand      string `json:"npu_brand"`
		NPUVersion    string `json:"npu_version"`
		NPUMemory     int    `json:"npu_memory"`
		NPUMemoryUnit string `json:"npu_memory_unit"`
	}
	Wechat struct {
		Flag       bool   `json:"flag"`
		Title      string `json:"title"`
		JobType    string `json:"job_type"`
		TemplateID string `json:"template_id"`
	}
}

var FineTune = struct {
	MaxJobNum int
	Pangu     *PanguFineTune
}{
	MaxJobNum: 5,
	Pangu:     &PanguFineTune{},
}

func getFineTuneConfig() {
	sec := Cfg.Section("finetune")
	FineTune.MaxJobNum = sec.Key("MAX_JOB_NUM").MustInt(5)
	panguBasicConfig := sec.Key("pangu_basic").MustString("")

	if err := json.Unmarshal([]byte(panguBasicConfig), &FineTune.Pangu.Basic); err != nil {
		log.Error("Unmarshal(panguConfig) failed:%v", err)
	}
	panguDatasetConfig := sec.Key("pangu_dataset").MustString("")

	if err := json.Unmarshal([]byte(panguDatasetConfig), &FineTune.Pangu.Dataset); err != nil {
		log.Error("Unmarshal(panguConfig) failed:%v", err)
	}

	panguModelConfig := sec.Key("pangu_model").MustString("")

	if err := json.Unmarshal([]byte(panguModelConfig), &FineTune.Pangu.Model); err != nil {
		log.Error("Unmarshal(panguConfig) failed:%v", err)
	}

	panguDeployConfig := sec.Key("pangu_deploy").MustString("")

	if err := json.Unmarshal([]byte(panguDeployConfig), &FineTune.Pangu.Deploy); err != nil {
		log.Error("Unmarshal(panguConfig) failed:%v", err)
	}

	panguServiceConfig := sec.Key("pangu_service").MustString("")

	if err := json.Unmarshal([]byte(panguServiceConfig), &FineTune.Pangu.Service); err != nil {
		log.Error("Unmarshal(panguConfig) failed:%v", err)
	}

	panguWechatConfig := sec.Key("pangu_wechat").MustString("")

	if err := json.Unmarshal([]byte(panguWechatConfig), &FineTune.Pangu.Wechat); err != nil {
		log.Error("Unmarshal(panguConfig) failed:%v", err)
	}
}

type ModelFinetuneConfig struct {
	LLAMA   LLaMaConfig `json:"llm"`
	Updated int64       `json:"-"`
}
type LLaMaConfig struct {
	Repo_owner_name string      `json:"repo_owner_name"`
	Repo_name       string      `json:"repo_name"`
	ModelInfos      []ModelInfo `json:"model"`
}

type ComputeResourceInfo struct {
	GPU Params `json:"GPU"`
	NPU Params `json:"NPU"`
}

type Params struct {
	Parameters string `json:"parameters"`
}

type Experience struct {
	ComputeResourceInfo ComputeResourceInfo `json:"compute_sources"`
}
type ModelInfo struct {
	Id         string     `json:"id"`
	Experience Experience `json:"experience"`
}

var ModelFinetuneConfigInfo *ModelFinetuneConfig

const UPATE_TTL = 300

func GetModelFinetuneConfig() *ModelFinetuneConfig {
	url := RecommentRepoAddr + "model/modelfinetune.json"
	if ModelFinetuneConfigInfo != nil && ModelFinetuneConfigInfo.Updated+UPATE_TTL > time.Now().Unix() {
		return ModelFinetuneConfigInfo
	} else {

		content, err := ReadFromPromote(url)
		var modelFinetuneConfig ModelFinetuneConfig
		if err == nil {

			err := json.Unmarshal(content, &modelFinetuneConfig)
			if err == nil {
				ModelFinetuneConfigInfo = &modelFinetuneConfig
				ModelFinetuneConfigInfo.Updated = time.Now().Unix()
				return ModelFinetuneConfigInfo
			}
		}
	}
	return nil
}
func ReadFromPromote(url string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Info("not error.", err)
			return
		}
	}()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Info("Get organizations url error=" + err.Error())
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return nil, err
	}
	return bytes, nil
}
