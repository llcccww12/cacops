package entity

import (
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"
)

type CreateAimodelReq struct {
	Name         string `json:"name" binding:"Required"`
	Alias        string `json:"alias"`
	AimodelType  int    `json:"aimodel_type"` // 0: online, 1: local, 2: external
	Engine       int64  `json:"engine"`
	License      string `json:"license"`
	Label        string `json:"label"`
	IsPrivate    bool   `json:"is_private"`
	CreatorId    int64  `json:"creator_id"`
	OwnerId      int64  `json:"owner_id"`
	ExternalName string `json:"external_name"` // only for aimodel_type = 2
	TaskId       int64  `json:"task_id"`       // only for aimodel_type = 0
	FileListStr  string `json:"file_list"`     // only for aimodel_type = 0
}

type AimodelExportTaskResultReq struct {
	TaskId      int64                 `json:"task_id" `
	FileListStr string                `json:"file_list" `
	Cloudbrain  *models.Cloudbrain    `json:"-"`
	Aimodel     *models.AiModelManage `json:"-"`
	Doer        *models.User          `json:"-"`
}

type AimodelMigrationInfo struct {
	Status int    `json:"status"`
	Desc   string `json:"status_desc"`
}

type AimodelPermissionInfo struct {
	CanEditFile bool `json:"can_edit_file"`
	CanManage   bool `json:"can_manage"`
	CanDownload bool `json:"can_download"`
	CanDelete   bool `json:"can_delete"`
}

type AimodelInfo struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Alias           string             `json:"alias"`
	Label           string             `json:"tags"`
	License         string             `json:"licenses"`
	Engine          int64              `json:"engine"`
	IsPrivate       bool               `json:"is_private"`
	UseCount        int                `json:"use_count"`
	DownloadCount   int                `json:"download_count"`
	NumStars        int                `json:"num_stars"`
	DerivativeCount int                `json:"derivative_count"`
	CreatedUnix     timeutil.TimeStamp `json:"created_unix"`
	UpdatedUnix     timeutil.TimeStamp `json:"updated_unix"`
	Size            int64              `json:"size"`
	OwnerName       string             `json:"owner_name"`
	CreatorName     string             `json:"creator_name"`
	Recommend       bool               `json:"recommend"`
	AimodelType     int                `json:"aimodel_type"`
	ExternalName    string             `json:"external_name"`
	ExternalUrl     string             `json:"external_url"`
	IsCollected     bool               `json:"is_collected"`

	Owner      *models.User4Front     `json:"Owner,omitempty"`
	Permission *AimodelPermissionInfo `json:"permission,omitempty"`
	Migration  *AimodelMigrationInfo  `json:"migration,omitempty"`
}

func BuildAimodelInfo(aimodel *models.AiModelManage) *AimodelInfo {
	var ownerFront *models.User4Front
	var ownerName string
	if aimodel.Owner != nil {
		ownerFront = aimodel.Owner.ToFrontFormat()
		ownerName = aimodel.Owner.Name
	}

	isRecommended := false
	if aimodel.Recommend == 1 {
		isRecommended = true
	}

	migrationInfo := BuildAimodelMigrationInfo(aimodel)
	externalUrl := buildAimodelExternalUrl(aimodel)

	return &AimodelInfo{
		ID:              aimodel.ID,
		Name:            aimodel.Name,
		Alias:           aimodel.Alias,
		Label:           aimodel.Label,
		License:         aimodel.License,
		Engine:          aimodel.Engine,
		IsPrivate:       aimodel.IsPrivate,
		UseCount:        aimodel.ReferenceCount,
		DownloadCount:   aimodel.DownloadCount,
		NumStars:        aimodel.CollectedCount,
		DerivativeCount: aimodel.DerivativeCount,
		Size:            aimodel.Size,
		Recommend:       isRecommended,
		AimodelType:     aimodel.ModelType,
		ExternalName:    aimodel.ExternalName,
		ExternalUrl:     externalUrl,
		IsCollected:     aimodel.IsCollected,
		CreatorName:     aimodel.CreatorName,

		Migration:   migrationInfo,
		OwnerName:   ownerName,
		Owner:       ownerFront,
		CreatedUnix: aimodel.CreatedUnix,
		UpdatedUnix: aimodel.UpdatedUnix,
	}

}

func buildAimodelExternalUrl(aimodel *models.AiModelManage) string {
	if aimodel == nil || strings.TrimSpace(aimodel.ExternalName) == "" {
		return ""
	}
	domain := strings.TrimSpace(setting.ExternalTransfer.HfDomain)
	if domain == "" {
		return ""
	}
	return strings.TrimRight(domain, "/") + "/" + strings.TrimLeft(aimodel.ExternalName, "/")
}

func BuildAimodelMigrationInfo(aimodel *models.AiModelManage) *AimodelMigrationInfo {
	if aimodel.ModelType == models.MODEL_LOCAL_TYPE {
		return nil
	}
	return &AimodelMigrationInfo{
		Status: aimodel.Status,
		Desc:   aimodel.StatusDesc,
	}
}

type SearchAimodelRes struct {
	Aimodels []*AimodelInfo `json:"aimodels"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

type AimodelReadmeReq struct {
	Content string                `json:"content" binding:"Required"`
	Aimodel *models.AiModelManage `json:"-"`
}

type AimodelGraphNode struct {
	Type        int                    `json:"type"` //0:aitask; 1:model
	IsParent    bool                   `json:"is_parent"`
	IsCurrent   bool                   `json:"is_current"`
	IsPrivate   bool                   `json:"is_private"`
	Visible     bool                   `json:"visible"`
	AimodelInfo *AimodelInfo           `json:"aimodel_info"`
	TaskInfo    *TaskInfo4AimodelGraph `json:"task_info"`
	Parents     []*AimodelGraphNode    `json:"parents"`
	Next        []*AimodelGraphNode    `json:"next"`
}

type TaskInfo4AimodelGraph struct {
	ID                int64                             `json:"id"`
	JobID             string                            `json:"job_id"`
	DisplayJobName    string                            `json:"display_job_name"`
	FormattedDuration string                            `json:"formatted_duration"`
	Spec              *structs.SpecificationShow        `json:"spec"`
	DatasetList       []*models.DatasetRegistryDownload `json:"dataset_list"`
	PretrainModelList []*models.Model4Show              `json:"pretrain_model_list"`
	RepoName          string                            `json:"repo_name"`
	RepoOwnerName     string                            `json:"repo_owner_name"`
	BranchName        string                            `json:"branch_name"`
	ImageName         string                            `json:"image_name"`
	CreatorName       string                            `json:"creator_name"`
}

func BuildTaskInfo4AimodelGraph(taskDetailedInfo *AITaskDetailInfo) *TaskInfo4AimodelGraph {
	return &TaskInfo4AimodelGraph{
		ID:                taskDetailedInfo.ID,
		JobID:             taskDetailedInfo.JobID,
		DisplayJobName:    taskDetailedInfo.DisplayJobName,
		FormattedDuration: taskDetailedInfo.FormattedDuration,
		Spec:              taskDetailedInfo.Spec,
		DatasetList:       taskDetailedInfo.DatasetList,
		PretrainModelList: taskDetailedInfo.PretrainModelList,
		RepoName:          taskDetailedInfo.RepoName,
		RepoOwnerName:     taskDetailedInfo.RepoOwnerName,
		BranchName:        taskDetailedInfo.BranchName,
		ImageName:         taskDetailedInfo.ImageName,
		CreatorName:       taskDetailedInfo.CreatorName,
	}
}

type AimodelRelatedData struct {
	Datasets []*RelatedDataset `json:"datasets"`
	Repos    []*RelatedRepo    `json:"repos"`
}

type RelatedDataset struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	Alias     string `json:"alias"`
}

type RelatedRepo struct {
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	Alias     string `json:"alias"`
}
