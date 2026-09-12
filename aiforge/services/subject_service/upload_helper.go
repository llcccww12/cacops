package subject_service

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
)

type UploadHelper interface {
	GetFileStoragePath(fileName string, dataId string) string
	GetFileNameByStoragePath(path string, dataId string) string
	CheckPermission(user *models.User, dataId string) error
	GetStoragePathPrefix(dataId string) string
	GetStorageType() entity.StorageType
	DoAfterUploadedSuccess(dataId string) error
}

func GetUploadHelper(subjectType models.SubjectType, dataIdArray ...string) UploadHelper {
	var dataId string
	if len(dataIdArray) > 0 {
		dataId = dataIdArray[0]
	}
	switch subjectType {
	case models.DatasetSubject:
		if dataId == "" {
			return &DatasetUploaderHelper{}
		}
		dataset, _ := models.GetDatasetRegistryByID(dataId)
		if dataset == nil {
			return &DatasetUploaderHelper{}
		}
		return &DatasetUploaderHelper{
			Dataset: dataset,
		}
	case models.AimodelSubject:
		if dataId == "" {
			return &AimodelUploaderHelper{}
		}
		aimodel, _ := models.GetAimodelByID(dataId)
		if aimodel == nil {
			return &AimodelUploaderHelper{}
		}
		return &AimodelUploaderHelper{
			Aimodel: aimodel,
		}
	default:
		return nil
	}
}

func GetUploadHelperBySubjectContext(ctx *models.SubjectAccessContext) UploadHelper {
	switch ctx.SubjectType {
	case models.DatasetSubject:
		return &DatasetUploaderHelper{
			Dataset: ctx.Dataset,
		}
	case models.AimodelSubject:
		return &AimodelUploaderHelper{
			Aimodel: ctx.Aimodel,
		}
	default:
		return nil
	}
}
