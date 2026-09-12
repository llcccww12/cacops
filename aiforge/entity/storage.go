package entity

import "code.gitea.io/gitea/models"

type StorageType string

const (
	MINIO     StorageType = "MINIO"
	OBS       StorageType = "OBS"
	URCHIN_V2 StorageType = "URCHIN_V2"
)

func GetStorageTypeFromCloudbrainType(cloudbrainType int) StorageType {
	switch cloudbrainType {
	case models.TypeCloudBrainOne:
		return MINIO
	case models.TypeCloudBrainTwo:
		return OBS

	}
	return ""
}

type StorageSummary struct {
	StorageLimit       int64 `json:"storage_limit"`
	DatasetUsedStorage int64 `json:"dataset_used_storage"`
	ModelUsedStorage   int64 `json:"model_used_storage"`
	UsedStorage        int64 `json:"used_storage"`
	RemainingStorage   int64 `json:"remaining_storage"`
}

type DatasetAttachmentShow struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	UUID            string `json:"uuid"`
	DecompressState int32  `json:"decompress_state"`
	Size            int64  `json:"size"`
	DatasetName     string `json:"dataset_name"`
	DatasetID       int64  `json:"dataset_id"`
	RepoID          int64  `json:"repo_id"`
	RepoOwner       string `json:"repo_owner"`
	RepoName        string `json:"repo_name"`
	IsPrivate       bool   `json:"is_private"`
	CreatedUnix     int64  `json:"created_unix"`
}
