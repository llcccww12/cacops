package modelmanage

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

func QueryModelFileByModel(model *models.AiModelManage) []storage.FileInfo {
	if model.Type == models.StorageTypeObs {
		if len(model.Path) > len(setting.Bucket)+1 {
			prefix := model.Path[len(setting.Bucket)+1:]
			fileinfos, _ := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, prefix)
			return fileinfos
		}

	} else if model.Type == models.StorageTypeMinio {
		if len(model.Path) > len(setting.Attachment.Minio.Bucket)+1 {
			prefix := model.Path[len(setting.Attachment.Minio.Bucket)+1:]
			fileinfos, _ := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, prefix)
			return fileinfos
		}
	}
	return nil
}

func HasModelFileByModelId(modelId string, ckptName string) bool {
	if modelId == "" {
		return false
	}
	model, err := models.QueryModelById(modelId)
	if err != nil || model == nil || model.ID == "" {
		log.Error("Can not find model", err)
		return false
	}
	fileInfos := QueryModelFileByModel(model)
	isFind := false
	if fileInfos != nil {
		for _, fileInfo := range fileInfos {
			if fileInfo.FileName == ckptName {
				isFind = true
				break
			}
		}

	}
	return isFind
}

func HasModelFile(task *models.Cloudbrain) bool {
	if task.ModelId == "" {
		return true
	}
	return HasModelFileByModelId(task.ModelId, task.CkptName)
}
