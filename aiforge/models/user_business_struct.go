package models

import "code.gitea.io/gitea/modules/timeutil"

type UserBusinessAnalysisCurrentYear struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`
	UserLocation       string  `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysisLast30Day struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`
	UserLocation       string  `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysisLastMonth struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`
	UserLocation       string  `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount    int `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysisCurrentMonth struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`
	UserLocation       string  `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysisCurrentWeek struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`

	UserLocation string `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysisYesterday struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`

	UserLocation string `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysisLastWeek struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`
	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`
	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`
	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`
	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`
	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`
	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`
	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`

	UserLocation string `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserAnalysisPara struct {
	Key   string  `xorm:"NOT NULL"`
	Value float64 `xorm:"NOT NULL DEFAULT 0"`
}

type UserMetrics struct {
	CountDate                  int64   `xorm:"pk"`
	ActivateRegistUser         int     `xorm:"NOT NULL DEFAULT 0"` //当天激活用户
	NotActivateRegistUser      int     `xorm:"NOT NULL DEFAULT 0"` //当天未激活用户
	ActivateIndex              float64 `xorm:"NOT NULL DEFAULT 0"` //激活比率
	RegistActivityUser         int     `xorm:"NOT NULL DEFAULT 0"` //当天注册激活的人中，有贡献活动的人
	HasActivityUser            int     `xorm:"NOT NULL DEFAULT 0"` //当天有贡献活动的人
	TotalUser                  int     `xorm:"NOT NULL DEFAULT 0"`
	TotalRegistUser            int     `xorm:"-"`
	TotalActivateRegistUser    int     `xorm:"NOT NULL DEFAULT 0"`
	TotalNotActivateRegistUser int     `xorm:"-"`
	TotalHasActivityUser       int     `xorm:"NOT NULL DEFAULT 0"`
	DisplayDate                string  `xorm:"-"`
	DataDate                   string  `xorm:"NULL"`
	DaysForMonth               int     `xorm:"NOT NULL DEFAULT 0"`
	HasActivityUserJson        string  `xorm:"text NULL"`          //贡献活动用户列表
	ActivityUserJson           string  `xorm:"text NULL"`          //激活用户列表
	CurrentDayRegistUser       int     `xorm:"NOT NULL DEFAULT 0"` //当天注册用户
}

type UserBusinessAnalysisAll struct {
	ID int64 `xorm:"pk"`

	CountDate int64 `xorm:"pk"`

	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`

	//action :ActionCommitRepo                               // 5
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`

	//action :ActionCreateIssue                             // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`

	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`

	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`

	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`

	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`

	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`

	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`

	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`

	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`

	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`

	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`

	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`

	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员

	//user
	Name string `xorm:"NOT NULL"`

	DataDate string `xorm:"NULL"`

	//cloudbraintask
	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`

	UserLocation string `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}

type UserBusinessAnalysis struct {
	ID        int64  `xorm:"pk"`
	DataDate  string `xorm:"pk"`
	CountDate int64  `xorm:"NULL"`

	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`

	//action :ActionCommitRepo                               // 5
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`

	//action :ActionCreateIssue                             // 6
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`

	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`

	//watch table current date
	FocusRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//star table current date
	StarRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`

	// user table
	GiteaAgeMonth int `xorm:"NOT NULL DEFAULT 0"`

	//
	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`

	//attachement table
	CommitDatasetSize int `xorm:"NOT NULL DEFAULT 0"`

	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`

	//issue, issueassignees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`

	//baike
	EncyclopediasCount int `xorm:"NOT NULL DEFAULT 0"`

	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`

	//repo
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//login count, from elk
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`

	//openi index
	OpenIIndex float64 `xorm:"NOT NULL DEFAULT 0"`

	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员

	//user
	Name string `xorm:"NOT NULL"`

	CloudBrainTaskNum  int     `xorm:"NOT NULL DEFAULT 0"`
	GpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuDebugJob        int     `xorm:"NOT NULL DEFAULT 0"`
	GpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuTrainJob        int     `xorm:"NOT NULL DEFAULT 0"`
	NpuInferenceJob    int     `xorm:"NOT NULL DEFAULT 0"`
	GpuBenchMarkJob    int     `xorm:"NOT NULL DEFAULT 0"`
	CloudBrainRunTime  int     `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum   int     `xorm:"NOT NULL DEFAULT 0"`
	UserIndex          float64 `xorm:"NOT NULL DEFAULT 0"`
	UserIndexPrimitive float64 `xorm:"NOT NULL DEFAULT 0"`

	UserLocation string `xorm:"NULL"`

	FocusOtherUser   int `xorm:"NOT NULL DEFAULT 0"`
	CollectDataset   int `xorm:"NOT NULL DEFAULT 0"`
	CollectedDataset int `xorm:"NOT NULL DEFAULT 0"`
	RecommendDataset int `xorm:"NOT NULL DEFAULT 0"`
	CollectImage     int `xorm:"NOT NULL DEFAULT 0"`
	CollectedImage   int `xorm:"NOT NULL DEFAULT 0"`
	RecommendImage   int `xorm:"NOT NULL DEFAULT 0"`

	Phone             string `xorm:"NULL"`
	InvitationUserNum int    `xorm:"NOT NULL DEFAULT 0"`
	ModelConvertCount int    `xorm:"NOT NULL DEFAULT 0"`

	LoginActionCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitDatasetNumNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有数据集的数量
	CollectDatasetNew   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏数据集的数量
	CollectedDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的数据集数量
	RecommendDatasetNew int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的数据集数量

	OwnAimodel         int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录拥有模型的数量
	CollectAimodel     int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录收藏模型的数量
	CollectedAimodel   int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被收藏的模型数量
	RecommendedAimodel int `xorm:"NOT NULL DEFAULT 0"` // 新增字段，记录被推荐的模型数量
}
