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

const versionKey = "data_version"
const versionFileName = "openi_resource.version"

func UpdateDatasetVersion(datasetId string) error {
	lock := redis_lock.NewDistributeLock(redis_key.DatasetVersionUpdateLock(datasetId))
	success, err := lock.LockWithWait(3*time.Second, 3*time.Second)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("UpdateDatasetVersion err")
	}
	defer lock.UnLock()
	dataset, err := models.GetDatasetRegistryByID(datasetId)
	if err != nil {
		return err
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", dataset.StorageType)
		return errors.New("Data type error")
	}

	info := map[string]interface{}{versionKey: time.Now().Unix()}

	metaInfo, _ := json.Marshal(info)
	objectKey := path.Join(helper.TrimBucketPrefix(dataset.Path), versionFileName)
	err = helper.UploadFile(objectKey, strings.NewReader(string(metaInfo)))
	return err
}

func QueryDatasetVersion(datasetId string) (string, error) {
	dataset, err := models.GetDatasetRegistryByID(datasetId)
	if err != nil {
		return "", err
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", dataset.StorageType)
		return "", errors.New("Data type error")
	}
	objectKey := path.Join(helper.TrimBucketPrefix(dataset.Path), versionFileName)
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

func InitDatasetVersion(datasetId string) {
	data, _ := QueryDatasetVersion(datasetId)
	if data != "" {
		return
	}
	UpdateDatasetVersion(datasetId)
}
