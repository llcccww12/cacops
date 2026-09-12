package subject_service

import (
	"encoding/json"
	"errors"
	"io"
	"path"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

func UpdateAimodelVersion(aimodelId string) error {
	lock := redis_lock.NewDistributeLock(redis_key.AimodelVersionUpdateLock(aimodelId))
	success, err := lock.LockWithWait(3*time.Second, 3*time.Second)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("UpdateAimodelVersion err")
	}
	defer lock.UnLock()
	aimodel, err := models.GetAimodelByID(aimodelId)
	if err != nil {
		return err
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", aimodel.StorageType)
		return errors.New("Data type error")
	}

	info := map[string]interface{}{versionKey: time.Now().Unix()}

	metaInfo, _ := json.Marshal(info)
	objectKey := path.Join(helper.TrimBucketPrefix(aimodel.Path), versionFileName)
	err = helper.UploadFile(objectKey, strings.NewReader(string(metaInfo)))
	return err
}

func QueryAimodelVersion(aimodelId string) (string, error) {
	aimodel, err := models.GetAimodelByID(aimodelId)
	if err != nil {
		return "", err
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", aimodel.StorageType)
		return "", errors.New("Data type error")
	}
	objectKey := path.Join(helper.TrimBucketPrefix(aimodel.Path), versionFileName)
	rd, err := helper.OpenFile(objectKey)
	if err != nil {
		log.Error("QueryModelMeta GetObject err.%v")
		return "", err
	}
	if rd == nil {
		return "", nil
	}
	s, _ := io.ReadAll(rd)
	return string(s), nil
}

func InitAimodelVersion(aimodelId string) {
	data, _ := QueryAimodelVersion(aimodelId)
	if data != "" {
		return
	}
	UpdateAimodelVersion(aimodelId)
}
