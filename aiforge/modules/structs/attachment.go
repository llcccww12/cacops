// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package structs // import "code.gitea.io/gitea/modules/structs"

import (
	"time"
)

// Attachment a generic attachment
// swagger:model
type Attachment struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	DownloadCount int64  `json:"download_count"`
	// swagger:strfmt date-time
	Created       time.Time `json:"created_at"`
	UUID          string    `json:"uuid"`
	DownloadURL   string    `json:"browser_download_url"`
	S3DownloadURL string
}

// EditAttachmentOptions options for editing attachments
// swagger:model
type EditAttachmentOptions struct {
	Name string `json:"name"`
}

type Dataset struct {
	ID            int64           `json:"id"`
	Title         string          `json:"title"`
	Status        int32           `json:"status"`
	Category      string          `json:"category"`
	Description   string          `json:"description"`
	DownloadTimes int64           `json:"downloadTimes"`
	UseCount      int64           `json:"useCount"`
	NumStars      int             `json:"numStars"`
	Recommend     bool            `json:"recommend"`
	License       string          `json:"license"`
	Task          string          `json:"task"`
	ReleaseID     int64           `json:"releaseId"`
	UserID        int64           `json:"userId"`
	RepoID        int64           `json:"repoId"`
	Repo          *RepositoryShow `json:"repo"`
	CreatedUnix   int64           `json:"createdUnix"`
	UpdatedUnix   int64           `json:"updatedUnix"`

	Attachments []*AttachmentShow `json:"attachments"`
}

type RepositoryShow struct {
	OwnerName string `json:"ownerName"`
	Name      string `json:"name"`
}

type AttachmentShow struct {
	ID              int64  `json:"id"`
	UUID            string `json:"uuid"`
	DatasetID       int64  `json:"datasetId"`
	ReleaseID       int64  `json:"releaseId"`
	UploaderID      int64  `json:"uploaderId"`
	CommentID       int64  `json:"commentId"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	DownloadCount   int64  `json:"downloadCount"`
	UseNumber       int64  `json:"useNumber"`
	Size            int64  `json:"size"`
	IsPrivate       bool   `json:"isPrivate"`
	DecompressState int32  `json:"decompressState"`
	Type            int    `json:"type"`
	CreatedUnix     int64  `json:"createdUnix"`
}
