package structs

type CreateGrampusTrainJobOption struct {
	DisplayJobName   string `json:"display_job_name" binding:"Required"`
	JobName          string `json:"job_name" binding:"Required" `
	Attachment       string `json:"attachment" binding:"Required"`
	BootFile         string `json:"boot_file" binding:"Required"`
	ImageID          string `json:"image_id" binding:"Required"`
	Params           string `json:"run_para_list" binding:"Required"`
	Description      string `json:"description"`
	BranchName       string `json:"branch_name" binding:"Required"`
	EngineName       string `json:"engine_name" binding:"Required"`
	WorkServerNumber int    `json:"work_server_number" binding:"Required"`
	Image            string `json:"image" binding:"Required"`
	DatasetName      string `json:"dataset_name" binding:"Required"`
	ModelName        string `json:"model_name"`
	ModelVersion     string `json:"model_version"`
	CkptName         string `json:"ckpt_name"`
	ModelId          string `json:"model_id"`
	LabelName        string `json:"label_names"`
	PreTrainModelUrl string `json:"pre_train_model_url"`
	SpecId           int64  `json:"spec_id" binding:"Required"`
}

type CreateFineTuneJobOption struct {
	Type           int    `json:"type"`
	DisplayJobName string `json:"display_job_name" binding:"Required"`
	//上传数据集和选择平台数据集需要输入
	Attachment  string `json:"attachment"`
	DatasetName string `json:"dataset_name"`
	//选择示例数据集必填
	SampleDatasetType int    `json:"sample_dataset_type"`
	Description       string `json:"description" `
	Params            string `json:"run_para_list" binding:"Required"`
	SpecId            int64  `json:"spec_id" binding:"Required"`
}

type CreateFineTuneDeployOption struct {
	JobID string `json:"job_id" binding:"Required"`
}

type UpdateFineTuneDeployOption struct {
	JobID  string `json:"job_id" binding:"Required"`
	Status string `json:"status" binding:"Required"`
}

type PanguInferenceOption struct {
	JobID string `json:"job_id" binding:"Required"`
	Text  string `json:"text" binding:"Required"`
}

type CreateTrainJobOption struct {
	Type              int         `json:"type"`
	DisplayJobName    string      `json:"display_job_name" binding:"Required"`
	ImageID           string      `json:"image_id"`
	Image             string      `json:"image" binding:"Required"`
	Attachment        string      `json:"attachment" binding:"Required"`
	DatasetName       string      `json:"dataset_name" binding:"Required"`
	Description       string      `json:"description" `
	BootFile          string      `json:"boot_file" binding:"Required"`
	BranchName        string      `json:"branch_name" binding:"Required"`
	Params            string      `json:"run_para_list" binding:"Required"`
	WorkServerNumber  int         `json:"work_server_number"`
	ModelName         string      `json:"model_name"`
	ModelVersion      string      `json:"model_version"`
	CkptName          string      `json:"ckpt_name"`
	ModelId           string      `json:"model_id"`
	LabelName         string      `json:"label_names"`
	PreTrainModelUrl  string      `json:"pre_train_model_url"`
	SpecId            int64       `json:"spec_id" binding:"Required"`
	FineTune          bool        `json:"-"`
	FineTuneModelType int         `json:"-"`
	FineTuneCategory  int         `json:"-"`
	CodeRepo          interface{} `json:"-"`
}
type CreateNotebookOption struct {
	Type             int    `json:"type"`
	DisplayJobName   string `json:"display_job_name" binding:"Required"`
	Attachment       string `json:"attachment"`
	ImageID          string `json:"image_id"`
	Description      string `json:"description"`
	BranchName       string `json:"branch_name" binding:"Required"`
	Image            string `json:"image" binding:"Required"`
	DatasetName      string `json:"dataset_name"`
	ModelName        string `json:"model_name"`
	ModelVersion     string `json:"model_version"`
	CkptName         string `json:"ckpt_name"`
	ModelId          string `json:"model_id"`
	LabelName        string `json:"label_names"`
	PreTrainModelUrl string `json:"pre_train_model_url"`
	SpecId           int64  `json:"spec_id"  binding:"Required"`
}

type CreateFileNotebookJobOption struct {
	Type        int    `json:"type"` //0 CPU 1 GPU 2 NPU
	ImageId     string `json:"image_id"`
	Image       string `json:"image"`
	File        string `json:"file" binding:"Required"`
	BranchName  string `json:"branch_name" binding:"Required"`
	OwnerName   string `json:"owner_name" binding:"Required"`
	ProjectName string `json:"project_name" binding:"Required"`
	JobId       string `json:"job_id"`
	ID          int64  `json:"id"`
}

type Cloudbrain struct {
	ID               int64              `json:"id"`
	JobID            string             `json:"job_id"`
	JobType          string             `json:"job_type"`
	Type             int                `json:"type"`
	DisplayJobName   string             `json:"display_job_name"`
	Status           string             `json:"status"`
	CreatedUnix      int64              `json:"created_unix"`
	RepoID           int64              `json:"repo_id"`
	Duration         int64              `json:"duration"` //运行时长 单位秒
	TrainJobDuration string             `json:"train_job_duration"`
	ImageID          string             `json:"image_id"` //grampus image_id
	Image            string             `json:"image"`
	Uuid             string             `json:"uuid"` //数据集id
	DatasetName      string             `json:"dataset_name"`
	ComputeResource  string             `json:"compute_resource"` //计算资源，例如npu
	AiCenter         string             `json:"ai_center"`        //grampus ai center: center_id+center_name
	BranchName       string             `json:"branch_name"`      //分支名称
	Parameters       string             `json:"parameters"`       //传给modelarts的param参数
	BootFile         string             `json:"boot_file"`        //启动文件
	Description      string             `json:"description"`      //描述
	ModelName        string             `json:"model_name"`       //模型名称
	ModelVersion     string             `json:"model_version"`    //模型版本
	CkptName         string             `json:"ckpt_name"`        //权重文件名称
	ModelId          string             `json:"model_id"`         //权重文件名称
	StartTime        int64              `json:"start_time"`
	EndTime          int64              `json:"end_time"`
	VersionName      string             `json:"version_name"`
	Spec             *SpecificationShow `json:"spec"`
}

type SpecificationShow struct {
	ID                  int64   `json:"id"`
	AICenterCode        string  `json:"ai_center_code"`
	AccCardsNum         int     `json:"acc_cards_num"`
	AccCardType         string  `json:"acc_card_type"`
	CpuCores            int     `json:"cpu_cores"`
	MemGiB              float32 `json:"mem_gi_b"`
	GPUMemGiB           float32 `json:"gpu_mem_gi_b"`
	ShareMemGiB         float32 `json:"share_mem_gi_b"`
	ComputeResource     string  `json:"compute_resource"`
	UnitPrice           float64 `json:"unit_price"`
	SourceSpecId        string  `json:"source_spec_id"`
	HasInternet         int     `json:"has_internet"`         //信息不准，考虑到用这个结构体的地方较多，先保留
	EnableVisualization bool    `json:"enable_visualization"` //信息不准，考虑到用这个结构体的地方较多，先保留

	NetworkCapableQueuesExist   bool // true: 存在能联网的队列
	NetworkIncapableQueuesExist bool // true: 存在不能联网的队列

	VisualizeCapableQueuesExist   bool // true: 存在能可视化的队列
	VisualizeIncapableQueuesExist bool // true: 存在不能可视化的队列
}

type PointAccountShow struct {
	ID            int64   `json:"id"`
	AccountCode   string  `json:"account_code"`
	Balance       float64 `json:"balance"`
	TotalEarned   float64 `json:"total_earned"`
	TotalConsumed float64 `json:"total_consumed"`
	Status        int     `json:"status"`
	Version       int64   `json:"version"`
	CreatedUnix   int64   `json:"created_unix"`
	UpdatedUnix   int64   `json:"updated_unix"`
}

type CreateModelNoteBook struct {
	RepoId  int64  `json:"repoId" binding:"Required"`
	ModelId string `json:"modelId" binding:"Required"`
}
