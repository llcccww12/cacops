package auth

import (
	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
)

type CreateModelArtsForm struct {
	JobName     string `form:"job_name" binding:"Required"`
	Attachment  string `form:"attachment"`
	Description string `form:"description"`
}

func (f *CreateModelArtsForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type CreateModelArtsNotebookForm struct {
	DisplayJobName   string `form:"display_job_name" binding:"Required"`
	JobName          string `form:"job_name" binding:"Required"`
	Attachment       string `form:"attachment"`
	Description      string `form:"description"`
	Flavor           string `form:"flavor" binding:"Required"`
	ImageId          string `form:"image_id" binding:"Required"`
	ModelName        string `form:"model_name"`
	ModelVersion     string `form:"model_version"`
	CkptName         string `form:"ckpt_name"`
	ModelId          string `form:"model_id"`
	LabelName        string `form:"label_names"`
	PreTrainModelUrl string `form:"pre_train_model_url"`
	SpecId           int64  `form:"spec_id" binding:"Required"`
	DatasetName      string `form:"dataset_name"`
}

func (f *CreateModelArtsNotebookForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type CreateModelArtsTrainJobForm struct {
	DisplayJobName        string `form:"display_job_name" binding:"Required"`
	JobName               string `form:"job_name" binding:"Required"`
	Attachment            string `form:"attachment" binding:"Required"`
	DatasetName           string `form:"dataset_name"`
	BootFile              string `form:"boot_file" binding:"Required"`
	WorkServerNumber      int    `form:"work_server_number" binding:"Required"`
	EngineID              int    `form:"engine_id" binding:"Required"`
	PoolID                string `form:"pool_id" binding:"Required"`
	Flavor                string `form:"flavor" binding:"Required"`
	Params                string `form:"run_para_list" binding:"Required"`
	Description           string `form:"description"`
	IsSaveParam           string `form:"is_save_para"`
	ParameterTemplateName string `form:"parameter_template_name"`
	PrameterDescription   string `form:"parameter_description"`
	BranchName            string `form:"branch_name" binding:"Required"`
	VersionName           string `form:"version_name" binding:"Required"`
	FlavorName            string `form:"flaver_names" binding:"Required"`
	EngineName            string `form:"engine_names" binding:"Required"`
	SpecId                int64  `form:"spec_id" binding:"Required"`
	ModelName             string `form:"model_name"`
	ModelId               string `form:"model_id"`
	ModelVersion          string `form:"model_version"`
	CkptName              string `form:"ckpt_name"`
	LabelName             string `form:"label_names"`
	PreTrainModelUrl      string `form:"pre_train_model_url"`
	IsContinue            bool   `form:"is_continue"`
}

type CreateModelArtsInferenceJobForm struct {
	DisplayJobName   string `form:"display_job_name" binding:"Required"`
	JobName          string `form:"job_name" binding:"Required"`
	Attachment       string `form:"attachment" binding:"Required"`
	BootFile         string `form:"boot_file" binding:"Required"`
	WorkServerNumber int    `form:"work_server_number" binding:"Required"`
	EngineID         int    `form:"engine_id" binding:"Required"`
	PoolID           string `form:"pool_id" binding:"Required"`
	Flavor           string `form:"flavor" binding:"Required"`
	Params           string `form:"run_para_list" binding:"Required"`
	Description      string `form:"description"`
	BranchName       string `form:"branch_name" binding:"Required"`
	VersionName      string `form:"version_name" binding:"Required"`
	FlavorName       string `form:"flaver_names" binding:"Required"`
	EngineName       string `form:"engine_names" binding:"Required"`
	LabelName        string `form:"label_names" binding:"Required"`
	PreTrainModelUrl string `form:"pre_train_model_url" binding:"Required"`
	ModelName        string `form:"model_name" binding:"Required"`
	ModelVersion     string `form:"model_version" binding:"Required"`
	CkptName         string `form:"ckpt_name" binding:"Required"`
	ModelId          string `form:"model_id"`
	SpecId           int64  `form:"spec_id" binding:"Required"`
}

func (f *CreateModelArtsTrainJobForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
