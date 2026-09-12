package cloudbrainTask

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

func IsModelFileExists(model *models.AiModelManage, fileName string) bool {
	if model.Type == models.TypeCloudBrainTwo {
		key := models.AIModelPath + models.AttachmentRelativePath(model.ID) + "/" + fileName
		log.Info("IsModelFileExists TypeCloudBrainTwo key=%s", key)
		isExist, err := storage.IsObjectExist4Obs(setting.Bucket, key)
		if err != nil {
			return false
		}
		return isExist
	} else if model.Type == models.TypeCloudBrainOne {
		prefix := models.AIModelPath + models.AttachmentRelativePath(model.ID) + "/"
		objectName := prefix + fileName
		log.Info("IsModelFileExists TypeCloudBrainOne objectName=%s", objectName)
		isExist, err := storage.IsObjectExist4Minio(setting.Attachment.Minio.Bucket, objectName)
		if err != nil {
			return false
		}
		return isExist
	}
	return false
}

func CheckAndGetFileSize(model *models.AiModelManage, fileName string) (bool, int64) {
	if model.Type == models.TypeCloudBrainTwo {
		key := models.AIModelPath + models.AttachmentRelativePath(model.ID) + "/" + fileName
		log.Info("IsModelFileExists TypeCloudBrainTwo key=%s", key)
		return storage.ObsCheckAndGetFileSize(setting.Bucket, key)
	} else if model.Type == models.TypeCloudBrainOne {
		prefix := models.AIModelPath + models.AttachmentRelativePath(model.ID) + "/"
		objectName := prefix + fileName
		log.Info("IsModelFileExists TypeCloudBrainOne objectName=%s", objectName)
		return storage.MinioCheckAndGetFileSize(setting.Attachment.Minio.Bucket, objectName)

	}
	return false, 0
}
