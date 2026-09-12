package entity

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/timeutil"
)

type CreateDatasetReq struct {
	Name      string `json:"name" binding:"Required"`
	Alias     string
	Tags      []string `json:"tags"`
	License   string   `json:"licenses"`
	Tasks     []string `json:"tasks"`
	IsPrivate bool     `json:"is_private"`
	CreatorId int64    `json:"creator_id"`
	OwnerId   int64    `json:"owner_id"`
	UUID      string   `json:"-"`
}

type DatasetInfo struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Alias         string             `json:"alias"`
	Tags          []string           `json:"tags"`
	License       string             `json:"licenses"`
	Tasks         []string           `json:"tasks"`
	IsPrivate     bool               `json:"is_private"`
	UseCount      int64              `json:"use_count"`
	DownloadCount int64              `json:"download_count"`
	NumStars      int                `json:"num_stars"`
	CreatedUnix   timeutil.TimeStamp `json:"created_unix"`
	UpdatedUnix   timeutil.TimeStamp `json:"updated_unix"`
	Size          int64              `json:"size"`
	OwnerName     string             `json:"owner_name"`
	CreatorName   string             `json:"creator_name"`
	Recommend     bool               `json:"recommend"`
	// isCollected, canEditFile, canManage
	IsCollected bool `json:"is_collected"`
	CanEditFile bool `json:"can_edit_file"`
	CanManage   bool `json:"can_manage"`
	CanDownload bool `json:"can_download"`
	CanDelete   bool `json:"can_delete"`
	Owner       *models.User4Front
	ExternalUrl    string `json:"external_url,omitempty"`
	ExternalSource string `json:"external_source,omitempty"`
}

func BuildDatasetInfo(dataset *models.DatasetRegistry) *DatasetInfo {
	var owner *models.User4Front
	var ownerName string
	if dataset.Owner != nil {
		owner = dataset.Owner.ToFrontFormat()
		ownerName = dataset.Owner.Name
	}
	return &DatasetInfo{
		ID:            dataset.ID,
		Name:          dataset.Name,
		Alias:         dataset.Alias,
		Tags:          dataset.Tags,
		License:       dataset.License,
		Tasks:         dataset.Tasks,
		IsPrivate:     dataset.IsPrivate,
		UseCount:      dataset.UseCount,
		DownloadCount: dataset.DownloadCount,
		NumStars:      dataset.NumCollections,
		CreatedUnix:   dataset.CreatedUnix,
		UpdatedUnix:   dataset.UpdatedUnix,
		Size:          dataset.Size,
		OwnerName:     ownerName,
		CreatorName:   dataset.CreatorName,
		Owner:         owner,
		IsCollected:   dataset.IsCollected,
		Recommend:     dataset.Recommend,
	}
}

type DatasetBriefInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	OwnerName string `json:"owner_name"`
}

func BuildDatasetBriefInfo(dataset *models.DatasetRegistry) *DatasetBriefInfo {
	var ownerName string
	if dataset.Owner != nil {
		ownerName = dataset.Owner.Name
	}
	return &DatasetBriefInfo{
		ID:        dataset.ID,
		Name:      dataset.Name,
		Alias:     dataset.Alias,
		OwnerName: ownerName,
	}
}

type SearchDatasetRes struct {
	Datasets []*DatasetInfo `json:"datasets"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

type DatasetReadmeReq struct {
	Content string                  `json:"content" binding:"Required"`
	Dataset *models.DatasetRegistry `json:"-"`
}

type DatasetReadmeResponse struct {
	Content       string `json:"content"`
	FileName      string `json:"fileName"`
	HtmlContent   string `json:"htmlcontent"`
	IsExistMDFile bool   `json:"isExistMDFile"`
}

type ExportTaskResultReq struct {
	TaskId      int64                   `json:"task_id" binding:"Required"`
	FileListStr string                  `json:"file_list" binding:"Required"`
	Cloudbrain  *models.Cloudbrain      `json:"-"`
	Dataset     *models.DatasetRegistry `json:"-"`
	Doer        *models.User            `json:"-"`
}
