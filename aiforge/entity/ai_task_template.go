package entity

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/timeutil"
)

type CreateAITaskTemplateReq struct {
	Name        string
	Description string
	Tags        []string
	IsPrivate   bool
	OwnerID     int64

	//task base
	JobType           string
	Cluster           string
	ComputeSource     string
	HasInternet       int
	VisualizeRequired bool
	Parameters        string

	//spec
	AccCardsNum int
	AccCardType string
	CpuCores    int
	MemGiB      float32
	GPUMemGiB   float32
	ShareMemGiB float32

	//image
	ImageID   string
	ImageName string
	ImageUrl  string

	//repo
	RepoName      string
	RepoOwnerName string
	BranchName    string
	BootFile      string

	DatasetIDs []string
	ModelIDs   []string
}

type AITaskTemplateDetailInfo struct {
	ID             string
	Name           string
	Description    string
	Tags           []string
	IsPrivate      bool
	UseCount       int64
	NumCollections int
	Recommend      bool
	CreatedUnix    timeutil.TimeStamp
	UpdatedUnix    timeutil.TimeStamp

	//task base
	JobType           string
	Cluster           string
	ComputeSource     string
	HasInternet       int
	VisualizeRequired bool
	Parameters        string

	//spec
	AccCardsNum int
	AccCardType string
	CpuCores    int
	MemGiB      float32
	GPUMemGiB   float32
	ShareMemGiB float32

	//image
	ImageID   string
	ImageName string
	ImageUrl  string

	//repo
	RepoID        int64
	RepoName      string
	RepoOwnerName string
	BranchName    string
	BootFile      string
	IsRepoDeleted bool

	OwnerId int64

	IsCollected bool

	CanEdit   bool
	CanDelete bool
	CanRead   bool

	Owner *models.User4Front

	DatasetLists []*models.TemplateDatasetInfo
	ModelLists   []*models.TemplateModelInfo
}

func BuildAITaskTemplateDetaiInfo(templateDTO *models.AITaskTemplate, operator *models.User) *AITaskTemplateDetailInfo {

	var ownerFront *models.User4Front
	if templateDTO.Owner != nil {
		ownerFront = templateDTO.Owner.ToFrontFormat()
	}
	return &AITaskTemplateDetailInfo{
		ID:             templateDTO.ID,
		Name:           templateDTO.Name,
		Description:    templateDTO.Description,
		Tags:           templateDTO.Tags,
		IsPrivate:      templateDTO.IsPrivate,
		UseCount:       templateDTO.UseCount,
		NumCollections: templateDTO.NumCollections,
		Recommend:      templateDTO.Recommend,
		CreatedUnix:    templateDTO.CreatedUnix,
		UpdatedUnix:    templateDTO.UpdatedUnix,
		//task base
		JobType:           templateDTO.JobType,
		Cluster:           templateDTO.Cluster,
		ComputeSource:     templateDTO.ComputeSource,
		HasInternet:       templateDTO.HasInternet,
		VisualizeRequired: templateDTO.VisualizeRequired,
		Parameters:        templateDTO.Parameters,

		//spec
		AccCardsNum: templateDTO.AccCardsNum,
		AccCardType: templateDTO.AccCardType,
		CpuCores:    templateDTO.CpuCores,
		MemGiB:      templateDTO.MemGiB,
		GPUMemGiB:   templateDTO.GPUMemGiB,
		ShareMemGiB: templateDTO.ShareMemGiB,

		//image
		ImageID:   templateDTO.ImageID,
		ImageName: templateDTO.ImageName,
		ImageUrl:  templateDTO.ImageUrl,

		//repo
		RepoID:        templateDTO.RepoId,
		RepoName:      templateDTO.RepoName,
		RepoOwnerName: templateDTO.RepoOwnerName,
		BranchName:    templateDTO.BranchName,
		BootFile:      templateDTO.BootFile,
		IsRepoDeleted: templateDTO.IsRepoDeleted,

		IsCollected: templateDTO.IsCollected,
		Owner:       ownerFront,
		OwnerId:     templateDTO.OwnerID,

		CanRead:   templateDTO.CanRead(operator),
		CanEdit:   templateDTO.CanEdit(operator),
		CanDelete: templateDTO.CanDelete(operator),

		DatasetLists: templateDTO.DatasetList,
		ModelLists:   templateDTO.ModelList,
	}

}
