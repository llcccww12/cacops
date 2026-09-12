package auth

import (
	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
)

type CreateGrampusTrainJobForm struct {
	DisplayJobName   string `form:"display_job_name" binding:"Required"`
	JobName          string `form:"job_name" binding:"Required"`
	Attachment       string `form:"attachment" binding:"Required"`
	BootFile         string `form:"boot_file" binding:"Required"`
	ImageID          string `form:"image_id" binding:"Required"`
	Params           string `form:"run_para_list" binding:"Required"`
	Description      string `form:"description"`
	BranchName       string `form:"branch_name" binding:"Required"`
	EngineName       string `form:"engine_name" binding:"Required"`
	WorkServerNumber int    `form:"work_server_number" binding:"Required"`
	Image            string `form:"image"`
	DatasetName      string `form:"dataset_name"`
	ModelName        string `form:"model_name"`
	ModelVersion     string `form:"model_version"`
	CkptName         string `form:"ckpt_name"`
	ModelId          string `form:"model_id"`
	LabelName        string `form:"label_names"`
	PreTrainModelUrl string `form:"pre_train_model_url"`
	SpecId           int64  `form:"spec_id"`
	PreJobName       string `form:"pre_job_name"`
	IsContinue       bool   `form:"is_continue"`
}

func (f *CreateGrampusTrainJobForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type CreateGrampusNotebookForm struct {
	Type             int    `form:"type"`
	DisplayJobName   string `form:"display_job_name" binding:"Required"`
	Attachment       string `form:"attachment"`
	ImageID          string `form:"image_id" binding:"Required"`
	Description      string `form:"description"`
	BranchName       string `form:"branch_name" binding:"Required"`
	Image            string `form:"image" binding:"Required"`
	DatasetName      string `form:"dataset_name"`
	ModelName        string `form:"model_name"`
	ModelVersion     string `form:"model_version"`
	CkptName         string `form:"ckpt_name"`
	ModelId          string `form:"model_id"`
	LabelName        string `form:"label_names"`
	PreTrainModelUrl string `form:"pre_train_model_url"`
	SpecId           int64  `form:"spec_id"  binding:"Required"`
}

func (f *CreateGrampusNotebookForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
