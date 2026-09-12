package entity

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/structs"
)

type CreationRequiredInfo struct {
	//排队信息、代码分支信息、查询是否有正在运行的任务、查询镜像列表、查询资源规格（积分余额，开关）
	Specs                 map[string][]*structs.SpecificationShow `json:"specs"`
	Images                []ClusterImage                          `json:"images"`
	CanUseAllImages       bool                                    `json:"can_use_all_images"`
	Branches              []string                                `json:"branches"`
	DefaultBranch         string                                  `json:"default_branch"`
	WaitCount             int64                                   `json:"wait_count"`
	NotStopTaskCount      int                                     `json:"not_stop_task_count"`
	CanCreateMore         bool                                    `json:"can_create_more"`
	NoteBookCanCreateMore bool                                    `json:"notebook_can_create_more"`
	DisplayJobName        string                                  `json:"display_job_name"`
	PointAccount          *PointAccountInfo                       `json:"point_account"`
	PaySwitch             bool                                    `json:"pay_switch"`
	Config                AITaskCreationConfig                    `json:"config"`
	AllowedWorkerNum      []int                                   `json:"allowed_worker_num"`
	IsSubscriber          bool                                    `json:"is_subscriber"`
	CodeSizeLimit         int                                     `json:"code_size_limit"`
	OutputSizeLimit       int                                     `json:"output_size_limit"`
	LimitCount            int                                     `json:"limit_count"`
	NotebookLimitCount    int                                     `json:"notebook_limit_count"`
}

type CreationRequiredRepoInfo struct {
	Branches      []string `json:"branches"`
	DefaultBranch string   `json:"default_branch"`
}

type ImageRequiredInfo struct {
	Images          []ClusterImage `json:"images"`
	CanUseAllImages bool           `json:"can_use_all_images"`
}

type AITaskCreationConfig struct {
	DatasetMaxSize int `json:"dataset_max_size"`
	DatasetsMaxNum int `json:"dataset_max_num"`
	ModelMaxSize   int `json:"model_max_size"`
	ModelMaxNum    int `json:"model_max_num"`
}

type SpecificationInfo struct {
	ID              int64   `json:"id"`
	SourceSpecId    string  `json:"source_spec_id"`
	AccCardsNum     int     `json:"acc_cards_num"`
	AccCardType     string  `json:"acc_card_type"`
	CpuCores        int     `json:"cpu_cores"`
	MemGiB          float32 `json:"mem_gi_b"`
	GPUMemGiB       float32 `json:"gpu_mem_gi_b"`
	ShareMemGiB     float32 `json:"share_mem_gi_b"`
	ComputeResource string  `json:"compute_resource"`
	UnitPrice       int     `json:"unit_price"`
	QueueId         int64   `json:"queue_id"`
	QueueCode       string  `json:"queue_code"`
	Cluster         string  `json:"cluster"`
	AiCenterCode    string  `json:"ai_center_code"`
	AiCenterName    string  `json:"ai_center_name"`
	IsExclusive     bool    `json:"is_exclusive"`
	ExclusiveOrg    string  `json:"exclusive_org"`
}

type PointAccountInfo struct {
	AccountCode   string  `json:"account_code"`
	Balance       float64 `json:"balance"`
	TotalEarned   float64 `json:"total_earned"`
	TotalConsumed float64 `json:"total_consumed"`
}

func ParsePointAccountInfo(p *models.PointAccount) *PointAccountInfo {
	return &PointAccountInfo{
		AccountCode:   p.AccountCode,
		Balance:       p.Balance,
		TotalEarned:   p.TotalEarned,
		TotalConsumed: p.TotalConsumed,
	}

}
