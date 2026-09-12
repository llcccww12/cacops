package subject_service

import (
	"path"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
)

const DATASET_PREFIX = "attachment"

type DatasetUploaderHelper struct {
	Dataset *models.DatasetRegistry
}

func (h *DatasetUploaderHelper) GetFileStoragePath(fileName string, dataId string) string {
	return strings.TrimPrefix(path.Join(h.GetStoragePathPrefix(dataId), fileName), "/")
}

func (h *DatasetUploaderHelper) GetFileNameByStoragePath(path string, dataId string) string {
	prefix := h.GetStoragePathPrefix(dataId)
	var fileName string
	if strings.HasPrefix(path, prefix) {
		fileName = strings.TrimPrefix(path, prefix)
	} else {
		parts := strings.Split(strings.TrimSuffix(path, "/"), "/")
		if len(parts) >= 1 {
			fileName = parts[len(parts)-1]
		}

	}
	return fileName
}

func (h *DatasetUploaderHelper) GetStoragePathPrefix(dataId string) string {
	if h.Dataset != nil {
		return h.Dataset.Path
	}
	return strings.TrimPrefix(path.Join(DATASET_PREFIX, path.Join(dataId[0:1], dataId[1:2], dataId+dataId)), "/") + "/"
}

func (h *DatasetUploaderHelper) CheckPermission(user *models.User, dataId string) error {
	return nil
}

func (h *DatasetUploaderHelper) GetStorageType() entity.StorageType {
	if h.Dataset != nil {
		return entity.StorageType(h.Dataset.StorageType)
	}
	var storageType string
	if setting.DATASET_STORAGE_TYPE != "" {
		storageType = setting.DATASET_STORAGE_TYPE
	} else {
		storageType = setting.StorageDefaultType
	}
	return entity.StorageType(strings.ToUpper(storageType))
}

func (h *DatasetUploaderHelper) DoAfterUploadedSuccess(dataID string) error {
	DoAfterDatasetFileChanged(dataID)
	return nil
}

func (h *DatasetUploaderHelper) getDataset(id string) *models.DatasetRegistry {
	dataset, err := models.GetDatasetRegistryByID(id)
	if err != nil {
		return nil
	}
	h.Dataset = dataset
	return nil
}
