package models

import "code.gitea.io/gitea/modules/timeutil"

// Overview 概览结构体
type Overview struct {
	BindingInfo BindingOverview `json:"binding_info"` // 用户绑定相关
	AiTaskInfo  AiTaskOverview  `json:"ai_task_info"` // ai计算任务相关
	StorageInfo StorageOverview `json:"storage_info"` // 存储相关
	PointInfo   PointOverview   `json:"point_info"`   // 积分相关

	RepositoryInfo RepositoryOverview `json:"repository_info"` // 项目相关
	DatasetInfo    DatasetOverview    `json:"dataset_info"`    // 数据集相关
	ModelInfo      ModelOverview      `json:"model_info"`      // 模型相关
}

// RepositoryOverview 项目相关概览
type RepositoryOverview struct {
	// 项目总数
	RepositoryCount int64 `json:"repository_count"`
	// 自建项目数量（用户/组织手动创建的项目）
	SelfBuiltCount int64 `json:"self_built_count"`
	// 派生项目数量（从其他项目Fork而来的项目）
	ForkedCount int64 `json:"forked_count"`
	// 镜像项目数量（从外部仓库镜像同步的项目）
	MirroredCount int64 `json:"mirrored_count"`
	// 协作项目数量（用户参与协作但非自己创建的项目）
	CollaborationCount int64 `json:"collaboration_count"`
}

// DatasetOverview 数据集相关概览
type DatasetOverview struct {
	// 数据集总数
	DatasetCount int64 `json:"dataset_count"`
	// 公开数据集数量（对所有用户可见）
	PublicCount int64 `json:"public_count"`
	// 私有数据集数量（仅对所有者或授权用户可见）
	PrivateCount int64 `json:"private_count"`
}

// ModelOverview 模型相关概览
type ModelOverview struct {
	// 模型总数
	ModelCount int64 `json:"model_count"`
	// 公开模型数量（对所有用户可见）
	PublicCount int64 `json:"public_count"`
	// 私有数据集数量（仅对所有者或授权用户可见）
	PrivateCount int64 `json:"private_count"`
}

type BindingOverview struct {
	// 是否绑定手机号 暂时去掉
	//IsBindPhone bool `json:"is_bind_phone,omitempty"`
	// 是否绑定微信号
	IsBindWechat bool `json:"is_bind_wechat,omitempty"`
}

// AiTaskOverview 计算任务相关
type AiTaskOverview struct {
	// 我的AI计算任务数量
	TotalAITasks int64 `json:"total_ai_tasks"`
	// 运行中的AI计算任务数量
	RunningAITasks int64 `json:"running_ai_tasks"`
	// AI计算累计运行的卡时（单位：卡时）
	AllCardDurationCount int64 `json:"all_card_duration_count"`

	// AI计算任务模板总数
	TemplateCount int64 `json:"template_count"`
	// AI计算任务模板运行总数
	TemplateUseCount int64 `json:"template_use_count"`
}

// StorageOverview 存储相关
type StorageOverview struct {
	// 存储总配额（计算得出，方便前端显示百分比）
	StorageLimit int64 `json:"storage_limit"`
	// 数据集占用存储
	DatasetUsedStorage int64 `json:"dataset_used_storage"`
	// 模型占用存储
	ModelUsedStorage int64 `json:"model_used_storage"`
	// 使用的容量
	UsedStorage int64 `json:"used_storage"`
	// 剩余容量
	RemainingStorage int64 `json:"remaining_storage"`
}

// PointOverview 资源配额相关
type PointOverview struct {
	// 获取的积分
	TotalEarned float64 `json:"total_earned"`
	// 剩余的积分
	Balance float64 `json:"balance"`
	// 消耗的积分
	TotalConsumed float64 `json:"total_consumed"`
}

// OverViewAction 概览页面的动态
type OverViewAction struct {
	OpType        ActionType               `json:"opType"`
	UserName      string                   `json:"userName"`
	CreatedUnix   timeutil.TimeStamp       `json:"createdUnix"`
	Comment       *OverViewComment         `json:"comment,omitempty"`
	Repo          *OverViewRepository      `json:"repo,omitempty"`
	Dataset       *OverViewDatasetRegistry `json:"dataset,omitempty"`
	CloudbrainId  int64                    `json:"cloudbrainId,omitempty"`
	Aimodel       *OverViewAiModel         `json:"aimodel,omitempty"`
	IsDeleted     bool                     `json:"isDeleted"`
	RefName       string                   `json:"refName,omitempty"`
	IsPrivate     bool                     `json:"isPrivate"`
	IsTransformed bool                     `json:"isTransformed"`
	Content       string                   `json:"content,omitempty"`
}

// OverViewComment 概览页面的评论
type OverViewComment struct {
	ID    int64  `json:"id"`
	Issue *Issue `json:"issue,omitempty"`
}

// OverViewRepository 概览页面的项目
type OverViewRepository struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias,omitempty"`
	OwnerName string `json:"ownerName,omitempty"`
}

// OverViewDatasetRegistry 概览页面的数据集
type OverViewDatasetRegistry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias,omitempty"`
	OwnerName string `json:"ownerName,omitempty"`
}

// OverViewAiModel 概览页面的模型
type OverViewAiModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias,omitempty"`
	OwnerName string `json:"ownerName,omitempty"`
}
