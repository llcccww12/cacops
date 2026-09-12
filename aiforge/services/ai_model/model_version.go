package ai_model

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

const versionKey = "data_version"
const versionFileName = "openi_resource.version"

// UpdateModelMeta 更新模型的版本信息
func UpdateModelMeta(modelId string) error {
	lock := redis_lock.NewDistributeLock(redis_key.AIModelMetaUpdateLock(modelId))
	success, err := lock.LockWithWait(3*time.Second, 3*time.Second)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("InitModelMeta err")
	}
	defer lock.UnLock()
	m, err := models.QueryModelById(modelId)
	if err != nil {
		return err
	}
	info := map[string]interface{}{versionKey: time.Now().Unix()}
	metaInfo, _ := json.Marshal(info)
	storageHelper := storage_helper.SelectStorageHelperFromStorageIntType(m.Type)

	bucket := storageHelper.GetBucket()
	if len(m.Path) > len(bucket)+1 {
		err = storageHelper.UploadFile(m.Path[len(bucket)+1:]+versionFileName, strings.NewReader(string(metaInfo)))

	}

	return err
}

func QueryModelMeta(modelId string) (string, error) {
	m, err := models.QueryModelById(modelId)
	if err != nil {
		log.Error("QueryModelMeta QueryModelById err.%v")
		return "", err
	}

	helper := storage_helper.SelectStorageHelperFromStorageIntType(m.Type)
	bucket := helper.GetBucket()
	if len(m.Path) > len(bucket)+1 {

		ObjectKey := m.Path[len(bucket)+1:] + versionFileName
		reader, err := helper.OpenFile(ObjectKey)
		if err != nil {
			log.Error("QueryModelMeta GetObject err.%v")
			return "", err
		}
		if reader == nil {
			return "", nil
		}
		s, _ := io.ReadAll(reader)
		return string(s), nil
	}
	return "", nil
}

func InitModelMeta(modelId string) {
	data, _ := QueryModelMeta(modelId)
	if data != "" {
		return
	}
	UpdateModelMeta(modelId)
}
