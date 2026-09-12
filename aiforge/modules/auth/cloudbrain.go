package auth

import (
	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
)

type CreateCloudBrainForm struct {
	JobName              string `form:"job_name" binding:"Required"`
	DisplayJobName       string `form:"display_job_name" binding:"Required"`
	Image                string `form:"image" binding:"Required"`
	Command              string `form:"command" binding:"Required"`
	Attachment           string `form:"attachment" binding:"Required"`
	JobType              string `form:"job_type" binding:"Required"`
	BenchmarkCategory    string `form:"get_benchmark_category"`
	GpuType              string `form:"gpu_type"`
	TrainUrl             string `form:"train_url"`
	TestUrl              string `form:"test_url"`
	Description          string `form:"description"`
	ResourceSpecId       int    `form:"resource_spec_id" binding:"Required"`
	BenchmarkTypeID      int    `form:"benchmark_types_id"`
	BenchmarkChildTypeID int    `form:"benchmark_child_types_id"`
	BootFile             string `form:"boot_file"`
	Params               string `form:"run_para_list"`
	BranchName           string `form:"branch_name"`
	ModelName            string `form:"model_name"`
	ModelVersion         string `form:"model_version"`
	CkptName             string `form:"ckpt_name"`
	ModelId              string `form:"model_id"`
	LabelName            string `form:"label_names"`
	PreTrainModelUrl     string `form:"pre_train_model_url"`
	DatasetName          string `form:"dataset_name"`
	SpecId               int64  `form:"spec_id"`
	PreJobName           string `form:"pre_job_name"`
	IsContinue           bool   `form:"is_continue"`
}

type CommitImageCloudBrainForm struct {
	Description            string `form:"description" binding:"Required"`
	Type                   int    `form:"type"  binding:"Required"`
	Tag                    string `form:"tag" binding:"Required;MaxSize(100)" `
	IsPrivate              bool   `form:"isPrivate" binding:"Required"`
	Topics                 string `form:"topics"`
	Framework              string `form:"framework" binding:"Required"`
	FrameworkVersion       string `form:"frameworkVersion"`
	CudaVersion            string `form:"cudaVersion" binding:"Required"`
	PythonVersion          string `form:"pythonVersion" binding:"Required"`
	OperationSystem        string `form:"operationSystem"`
	OperationSystemVersion string `form:"operationSystemVersion"`
	ThirdPackages          string `form:"thirdPackages"`
	ComputeResource        string `form:"computeResource"`
}
type CommitImageGrampusForm struct {
	Description            string `form:"description" binding:"Required"`
	Type                   int    `form:"type"  binding:"Required"`
	Tag                    string `form:"tag" binding:"Required;MaxSize(50)" `
	IsPrivate              bool   `form:"isPrivate" binding:"Required"`
	Topics                 string `form:"topics"`
	Framework              string `form:"framework" binding:"Required"`
	FrameworkVersion       string `form:"frameworkVersion"`
	CudaVersion            string `form:"cudaVersion" binding:"Required"`
	CannVersion            string `form:"cannVersion" binding:"Required"`
	DTKVersion             string `form:"dtkVersion" binding:"Required"`
	PythonVersion          string `form:"pythonVersion" binding:"Required"`
	OperationSystem        string `form:"operationSystem"`
	OperationSystemVersion string `form:"operationSystemVersion"`
	ThirdPackages          string `form:"thirdPackages"`
	ComputeResource        string `form:"computeResource"`
	ImageId                string `form:"imageId"`
	ImageTag               string `form:"imageTag"`
}

type CommitAdminImageCloudBrainForm struct {
	Description            string `form:"description" binding:"Required"`
	Type                   int    `form:"type"  binding:"Required"`
	Tag                    string `form:"tag" binding:"Required;MaxSize(100)" `
	IsPrivate              bool   `form:"isPrivate" binding:"Required"`
	Topics                 string `form:"topics"`
	Place                  string `form:"place"  binding:"Required"`
	IsRecommend            bool   `form:"isRecommend" binding:"Required"`
	Framework              string `form:"framework" binding:"Required"`
	FrameworkVersion       string `form:"frameworkVersion" binding:"Required"`
	CudaVersion            string `form:"cudaVersion" binding:"Required"`
	DTKVersion             string `form:"dtkVersion" binding:"Required"`
	PythonVersion          string `form:"pythonVersion" binding:"Required"`
	TrainType              string `form:"trainType" binding:"Required"`
	OperationSystem        string `form:"operationSystem"`
	OperationSystemVersion string `form:"operationSystemVersion"`
	ThirdPackages          string `form:"thirdPackages"`
	ComputeResource        string `form:"computeResource"`
}

type EditImageCloudBrainForm struct {
	ID                     int64  `form:"id" binding:"Required"`
	Description            string `form:"description" binding:"Required"`
	Topics                 string `form:"topics"`
	Framework              string `form:"framework" binding:"Required"`
	FrameworkVersion       string `form:"frameworkVersion" binding:"Required"`
	CudaVersion            string `form:"cudaVersion" binding:""`
	CannVersion            string `form:"cannVersion" binding:""`
	DTKVersion             string `form:"dtkVersion" binding:""`
	PythonVersion          string `form:"pythonVersion" binding:"Required"`
	OperationSystem        string `form:"operationSystem"`
	OperationSystemVersion string `form:"operationSystemVersion"`
	ThirdPackages          string `form:"thirdPackages"`
}

type ReviewImageForm struct {
	Message string `form:"message" binding:"Required"`
}

type CreateCloudBrainInferencForm struct {
	JobName           string `form:"job_name" binding:"Required"`
	DisplayJobName    string `form:"display_job_name" binding:"Required"`
	Image             string `form:"image" binding:"Required"`
	Command           string `form:"command" binding:"Required"`
	Attachment        string `form:"attachment" binding:"Required"`
	JobType           string `form:"job_type" binding:"Required"`
	BenchmarkCategory string `form:"get_benchmark_category"`
	GpuType           string `form:"gpu_type"`
	PreTrainModelUrl  string `form:"pre_train_model_url"`
	TestUrl           string `form:"test_url"`
	Description       string `form:"description"`
	ResourceSpecId    int    `form:"resource_spec_id" binding:"Required"`
	BootFile          string `form:"boot_file"`
	Params            string `form:"run_para_list"`
	BranchName        string `form:"branch_name"`
	ModelName         string `form:"model_name" binding:"Required"`
	ModelVersion      string `form:"model_version" binding:"Required"`
	CkptName          string `form:"ckpt_name" binding:"Required"`
	LabelName         string `form:"label_names" binding:"Required"`
	ModelId           string `form:"model_id" binding:"Required"`
	DatasetName       string `form:"dataset_name"`
	SpecId            int64  `form:"spec_id"`
}

func (f *CreateCloudBrainForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

func (f *CommitImageCloudBrainForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

func (f *EditImageCloudBrainForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

func (f *CreateCloudBrainInferencForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
