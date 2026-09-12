package entity

import (
	"code.gitea.io/gitea/models"
)

type UploadChunkResponse struct {
	UUID     string `json:"uuid"`
	Uploaded int    `json:"uploaded"`
	Chunks   []int  `json:"chunks"`
	FileName string `json:"fileName"`
}
type UploadChunkRequest struct {
	FileMD5     string `json:"md5"`
	FileName    string `json:"file_name"`
	SubjectID   string `json:"subjectId"`
	SubjectType int    `json:"data_type"`
	UserId      int64
}
type NewMultipartRequest struct {
	//dataId, fileName, fileType string, dataType, storageType int, user *models.User
	SubjectId            string `json:"subject_id"`
	FileName             string `json:"file_name"`
	FileType             string `json:"file_type"`
	SubjectType          int    `json:"subject_type"`
	TotalChunkCounts     int    `json:"total_chunk_counts"`
	Size                 int64  `json:"size"`
	MD5                  string `json:"md5"`
	User                 *models.User
	SubjectAccessContext *models.SubjectAccessContext
}

type GetMultipartUrlRequest struct {
	UUID       string `json:"uuid"`
	PartNumber int    `json:"chunkNumber"`
	Size       int64  `json:"size"`
}

type NewMuiltipartResponse struct {
	UUID string `json:"uuid"`
}

type CompleteMultipartRequest struct {
	UUID string `json:"uuid"`
}

type GetUploadUrlRequest struct {
	SubjectContext *models.SubjectAccessContext
	FileName       string
	User           *models.User
	Size           int64
	FileType       string
}

type CompleteUploadRequest struct {
	FileNameList   []string                     `json:"file_name_list"`
	SubjectContext *models.SubjectAccessContext `json:"-"`
}
