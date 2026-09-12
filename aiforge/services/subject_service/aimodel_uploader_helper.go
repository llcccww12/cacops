package subject_service

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
	"path"
	"strings"
)

const AIMODEL_PREFIX = "aimodels"

type AimodelUploaderHelper struct {
	Aimodel *models.AiModelManage
}

func (h *AimodelUploaderHelper) GetFileStoragePath(fileName string, aimodelId string) string {
	return strings.TrimPrefix(path.Join(h.GetStoragePathPrefix(aimodelId), fileName), "/")
}

func (h *AimodelUploaderHelper) GetFileNameByStoragePath(path string, aimodelId string) string {
	prefix := h.GetStoragePathPrefix(aimodelId)
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

func (h *AimodelUploaderHelper) GetStoragePathPrefix(aimodelId string) string {
	if h.Aimodel != nil {
		return h.Aimodel.Path
	}
	return strings.TrimPrefix(path.Join(AIMODEL_PREFIX, path.Join(aimodelId[0:1], aimodelId[1:2], aimodelId)), "/") + "/"
}

func (h *AimodelUploaderHelper) CheckPermission(user *models.User, aimodelId string) error {
	return nil
}

func (h *AimodelUploaderHelper) GetStorageType() entity.StorageType {
	if h.Aimodel != nil {
		return entity.StorageType(h.Aimodel.StorageType)
	}
	var storageType string
	if setting.AIMODEL_STORAGE_TYPE != "" {
		storageType = setting.AIMODEL_STORAGE_TYPE
	} else {
		storageType = setting.StorageDefaultType
	}
	return entity.StorageType(strings.ToUpper(storageType))
}

func (h *AimodelUploaderHelper) DoAfterUploadedSuccess(dataID string) error {
	DoAfterAimodelFileChanged(dataID)
	return nil
}

func (h *AimodelUploaderHelper) getAimodel(id string) *models.AiModelManage {
	aimodel, err := models.GetAimodelByID(id)
	if err != nil {
		return nil
	}
	h.Aimodel = aimodel
	return nil
}
