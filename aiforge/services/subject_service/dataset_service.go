package subject_service

import (
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/storage_limit"

	gouuid "github.com/satori/go.uuid"
)

var NamePattern = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9\-_.]{0,98}[A-Za-z0-9])?$`)

var AlphaDashDotChinese = regexp.MustCompile(`^[\p{Han}A-Za-z0-9\-_.]{1,100}$`)

func CreateDataset(doer, owner *models.User, req entity.CreateDatasetReq) (*models.DatasetRegistry, *response.BizError) {
	if !NamePattern.MatchString(req.Name) {
		return nil, response.DATASET_NAME_INVALID
	}
	if req.Alias == "" {
		req.Alias = req.Name
	}
	if !AlphaDashDotChinese.MatchString(req.Alias) {
		return nil, response.DATASET_ALIAS_INVALID
	}
	redisKey := redis_key.DatasetNameLock(req.OwnerId, strings.ToLower(req.Name))
	lock := redis_lock.NewDistributeLock(redisKey)
	isOk, err := lock.LockWithWait(5*time.Second, 5*time.Second)
	if err != nil {
		log.Error("CreateDataset LockWithWait failed, redisKey=%s, err=%v", redisKey, err)
		return nil, response.NewBizError(err)
	}
	if !isOk {
		log.Error("CreateDataset LockWithWait failed, redisKey=%s, err=%v", redisKey, response.DATASET_NAME_EXIST.ToError())
		return nil, response.DATASET_NAME_EXIST
	}
	defer func() {
		err := lock.UnLock()
		if err != nil {
			log.Error("CreateDataset UnLock failed, redisKey=%s, err=%v", redisKey, err)
		}
	}()

	record, err := models.GetDatasetRegistryByOwnerAndName(req.OwnerId, req.Name)
	if err != nil {
		if !models.IsErrRecordNotExist(err) {
			log.Error("CreateDataset GetDatasetRegistryByOwnerAndName failed, ownerId=%d, name=%s, err=%v", req.OwnerId, req.Name, err)
			return nil, response.NewBizError(err)
		}
	}
	if record != nil {
		return nil, response.DATASET_NAME_EXIST
	}
	record, err = models.GetDatasetRegistryByOwnerAndAlias(req.OwnerId, req.Alias)
	if err != nil {
		if !models.IsErrRecordNotExist(err) {
			log.Error("CreateDataset GetDatasetRegistryByOwnerAndAlias failed, ownerId=%d, alias=%s, err=%v", req.OwnerId, req.Alias, err)
			return nil, response.NewBizError(err)
		}
	}
	if record != nil {
		return nil, response.DATASET_ALIAS_EXIST
	}

	uploaderHelper := GetUploadHelper(models.DatasetSubject)
	storageType := uploaderHelper.GetStorageType()
	if storageType == "" {
		log.Info("dataset storage type setting is empty")
		return nil, response.SYSTEM_ERROR
	}
	stroageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if stroageHelper == nil {
		log.Info("dataset storage type %s not support", storageType)
		return nil, response.SYSTEM_ERROR
	}

	dataId := gouuid.NewV4().String()

	prefix := path.Join(stroageHelper.GetBucket(), uploaderHelper.GetStoragePathPrefix(dataId)) + "/"
	path, err := stroageHelper.AllocateDatasetNamespace(req.Name, prefix)
	if err != nil {
		log.Error("AllocateDatasetNamespace failed, prefix=%s, err=%v", req.Name, err)
		return nil, response.NewBizError(err)
	}

	datasetDTO := &models.DatasetRegistry{
		ID:          dataId,
		Name:        req.Name,
		Alias:       req.Alias,
		LowerAlias:  strings.ToLower(req.Alias),
		LowerName:   strings.ToLower(req.Name),
		Tags:        req.Tags,
		License:     req.License,
		Tasks:       req.Tasks,
		IsPrivate:   req.IsPrivate,
		CreatorID:   req.CreatorId,
		OwnerID:     req.OwnerId,
		StorageType: string(storageType),
		Path:        path,
	}
	err = models.CreateDatasetRegistry(datasetDTO, doer)
	if err != nil {
		log.Error("CreateDatasetRegistry failed, name=%s, err=%v", req.Name, err)
		return nil, response.NewBizError(err)
	}

	go notification.NotifyCreateDataset(doer, owner, datasetDTO)

	return datasetDTO, nil
}

func EditDataset(req entity.CreateDatasetReq, dataset *models.DatasetRegistry) *response.BizError {
	if !NamePattern.MatchString(req.Name) {
		return response.DATASET_NAME_INVALID
	}
	if req.Alias == "" {
		req.Alias = req.Name
	}
	if !AlphaDashDotChinese.MatchString(req.Alias) {
		return response.DATASET_ALIAS_INVALID
	}
	if dataset.Name != req.Name {
		redisKey := redis_key.DatasetNameLock(dataset.OwnerID, strings.ToLower(req.Name))
		lock := redis_lock.NewDistributeLock(redisKey)
		isOk, err := lock.LockWithWait(5*time.Second, 5*time.Second)
		if err != nil {
			log.Error("DatasetNameLock LockWithWait failed, redisKey=%s, err=%v", redisKey, err)
			return response.NewBizError(err)
		}
		if !isOk {
			log.Error("CreateDataset LockWithWait failed, redisKey=%s, err=%v", redisKey, response.DATASET_NAME_EXIST.ToError())
			return response.DATASET_NAME_EXIST
		}
		defer func() {
			err := lock.UnLock()
			if err != nil {
				log.Error("EditDataset UnLock failed, redisKey=%s, err=%v", redisKey, err)
			}
		}()

		record, err := models.GetDatasetRegistryByOwnerAndName(dataset.OwnerID, req.Name)
		if err != nil {
			if !models.IsErrRecordNotExist(err) {
				log.Error("CreateDataset GetDatasetRegistryByOwnerAndName failed, ownerId=%d, name=%s, err=%v", dataset.OwnerID, req.Name, err)
				return response.NewBizError(err)
			}
		}
		if record != nil {
			if record.ID != dataset.ID {
				return response.DATASET_NAME_EXIST
			}
		}

	}
	if dataset.Alias != req.Alias {
		record, err := models.GetDatasetRegistryByOwnerAndAlias(dataset.OwnerID, req.Alias)
		if err != nil {
			if !models.IsErrRecordNotExist(err) {
				log.Error("CreateDataset GetDatasetRegistryByOwnerAndAlias failed, ownerId=%d, alias=%s, err=%v", dataset.OwnerID, req.Alias, err)
				return response.NewBizError(err)
			}
		}
		if record != nil {
			if record.ID != dataset.ID {
				return response.DATASET_ALIAS_EXIST
			}
		}
	}
	err := models.UpdateDatasetRegistry(models.DatasetRegistry{
		ID:         dataset.ID,
		Name:       req.Name,
		Alias:      req.Alias,
		Tags:       req.Tags,
		License:    req.License,
		Tasks:      req.Tasks,
		IsPrivate:  req.IsPrivate,
		LowerName:  strings.ToLower(req.Name),
		LowerAlias: strings.ToLower(req.Alias),
	})
	if err != nil {
		log.Error("UpdateDatasetRegistry failed, datasetId=%s, err=%v", dataset.ID, err)
		return response.NewBizError(err)
	}

	return nil
}

func GetDataset(id string) (*entity.DatasetInfo, error) {
	dataset, err := models.GetDatasetRegistryByID(id)
	if err != nil {
		log.Error("GetDatasetV2ByID failed, datasetId=%s", id)
		return nil, err
	}
	owner, err := models.GetUserByID(dataset.OwnerID)
	if err != nil {
		log.Error("GetUserByID failed, ownerId=%d", dataset.OwnerID)
		return nil, err
	}

	datasetInfo := &entity.DatasetInfo{
		ID:            dataset.ID,
		Name:          dataset.Name,
		Tags:          dataset.Tags,
		License:       dataset.License,
		Tasks:         dataset.Tasks,
		IsPrivate:     dataset.IsPrivate,
		UseCount:      dataset.UseCount,
		DownloadCount: dataset.DownloadCount,
		NumStars:      dataset.NumCollections,
		CreatedUnix:   dataset.CreatedUnix,
		UpdatedUnix:   dataset.UpdatedUnix,
		OwnerName:     owner.Name,
	}
	return datasetInfo, nil
}

const README_FILE_NAME = "README.md"

func DeleteDatasetFile(dataset *models.DatasetRegistry, parentDir, fileName string) error {
	objectKey := path.Join(dataset.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("StorageHelper failed, StorageType=%d", dataset.StorageType)
		return errors.New("Storage type error")
	}
	err := helper.DeleteFile(objectKey)
	if err != nil {
		log.Error("DeleteFile failed, objectKey=%s err=%v", objectKey, err)
		return err
	}
	DoAfterDatasetFileChanged(dataset.ID)
	return nil
}

func DeleteDataset(doer *models.User, dataset *models.DatasetRegistry) error {
	owner, err := models.GetUserByID(dataset.OwnerID)
	if err != nil {
		log.Error("GetUserByID failed, OwnerID=%s err=%v", dataset.OwnerID, err)
		return err
	}
	err = models.WithTx(func(ctx models.DBContext) error {
		err := models.DeleteDatasetRegistry(ctx, doer, dataset.OwnerID, dataset.ID)
		if err != nil {
			log.Error("DeleteDataset err,dataset=%+v,err=%v", dataset, err)
			return err
		}
		err = DeleteDatasetStorage(dataset)
		if err != nil {
			log.Error("DeleteDataset DeleteDatasetStorage err,dataset=%+v,err=%v", dataset, err)
			return err
		}
		notification.NotifyDeleteDataset(doer, owner, dataset)
		return nil
	})

	return err
}

func DeleteDatasetStorage(dataset *models.DatasetRegistry) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", models.DatasetSubject)
		return errors.New("Storage type error")
	}
	remainingFiles, err := helper.DeleteCollection(dataset.Path)
	if err == nil && remainingFiles {
		tmpErr := models.InsertStorageDeleteFailedDataset(dataset)
		if tmpErr != nil {
			log.Error("InsertStorageDeleteFailedDataset failed, datasetId=%s, datasetName=%s err=%v", dataset.ID, dataset.Name, tmpErr)
		}
	}
	return err
}

func BatchDeleteDataset(doer *models.User, datasetIds []string) []string {
	successIds := make([]string, 0)
	for _, datasetId := range datasetIds {
		if datasetId == "" {
			continue
		}
		dataset, err := models.GetDatasetRegistryByID(datasetId)
		if err != nil {
			log.Error("BatchDeleteDataset GetDatasetRegistryByID error.id=%s err=%v", datasetId, err)
			continue
		}
		if !doer.CanDeleteDatasetRegistry(*dataset) {
			log.Error("BatchDeleteDataset check permission failed.id=%s userId=%d", datasetId, doer.ID)
			continue
		}
		err = DeleteDataset(doer, dataset)
		if err != nil {
			log.Error("BatchDeleteDataset DeleteDataset error.id=%s err=%v", datasetId, err)
			continue
		}
		successIds = append(successIds, datasetId)
	}
	return successIds
}

func CollectDataset(user *models.User, dataset *models.DatasetRegistry) error {
	return models.CollectionDataset(user.ID, dataset.ID, true)
}

func UnCollectDataset(user *models.User, dataset *models.DatasetRegistry) error {
	return models.CollectionDataset(user.ID, dataset.ID, false)
}

func RecommendDataset(user *models.User, dataset *models.DatasetRegistry) error {
	var err error
	err = models.RecommendDatasetRegistry(dataset.ID, true)
	if err != nil {
		return err
	}

	notification.NotifyDatasetRegistryRecommend(user, dataset)
	return nil
}

func UnRecommendDataset(user *models.User, dataset *models.DatasetRegistry) error {
	var err error
	err = models.RecommendDatasetRegistry(dataset.ID, false)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDatasetRegistrySize(datasetId string) bool {
	dataset, err := models.GetDatasetRegistryByID(datasetId)
	if err != nil {
		log.Info("not found dataset, id=%s", datasetId)
		return false
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", dataset.StorageType)
		return false
	}
	size, err := helper.CountDirSize(dataset.Path)
	if err != nil {
		log.Error("Failed to query size . id=%s err=%v", datasetId, err)
		return false
	}
	models.UpdateDatasetRegistrySize(datasetId, size)
	return true
}

func EnqueueDatasetSizeChangedTask(datasetId string) {
	score := float64(time.Now().Add(sizeChangedDelayWindow).Unix())
	if err := redis_client.ZAdd(redis_key.DatasetSizeChangedQueue(), datasetId, score); err != nil {
		log.Error("EnqueueDatasetSizeChangedTask ZAdd failed, datasetId=%s, err=%v", datasetId, err)
	}
}

func getDatasetSizeChangedQueueBatchSize() int {
	if setting.DatasetSizeChangedQueueBatchSize <= 0 {
		return 10
	}
	return setting.DatasetSizeChangedQueueBatchSize
}

func ProcessDatasetSizeChangedQueue() {
	now := float64(time.Now().Unix())
	ids, err := redis_client.ZRangeByScoreLimit(redis_key.DatasetSizeChangedQueue(), 0, now, 0, getDatasetSizeChangedQueueBatchSize())
	if err != nil {
		log.Error("ProcessDatasetSizeChangedQueue ZRangeByScoreLimit failed, err=%v", err)
		return
	}
	for _, id := range ids {
		ReportDatasetSizeChanged(id)
		score, exists, err := redis_client.ZScore(redis_key.DatasetSizeChangedQueue(), id)
		if err != nil {
			log.Error("ProcessDatasetSizeChangedQueue ZScore failed, id=%s, err=%v", id, err)
			continue
		}
		if !exists || score > now {
			continue
		}
		if err := redis_client.ZRem(redis_key.DatasetSizeChangedQueue(), id); err != nil {
			log.Error("ProcessDatasetSizeChangedQueue ZRem failed, id=%s, err=%v", id, err)
		}
	}
}

func ReportDatasetSizeChanged(datasetId string) {
	dataset, err := models.GetDatasetRegistryByID(datasetId)
	if err != nil {
		log.Info("not found dataset, id=%s", datasetId)
		return
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", dataset.StorageType)
		return
	}
	if err := helper.ReportSizeChanged(dataset.Path); err != nil {
		log.Error("ReportDatasetSizeChanged failed, datasetId=%s err=%v", datasetId, err)
	}
}

func DoAfterDatasetFileChanged(datasetId string) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	UpdateDatasetVersion(datasetId)
	if UpdateDatasetRegistrySize(datasetId) {
		EnqueueDatasetSizeChangedTask(datasetId)
	}
}

func SearchDatasets(req models.SearchDatasetReq) ([]*entity.DatasetInfo, int64, *response.BizError) {
	datasets, total, err := models.SearchDatasetRegistry(req)
	if err != nil {
		log.Error("SearchDatasets failed, req=%+v, err=%v", req, err)
		return nil, 0, response.NewBizError(err)
	}
	if datasets == nil || len(datasets) == 0 {
		log.Info("SearchDatasets no data found, req=%+v", req)
		return nil, 0, nil
	}

	if len(datasets) > 0 {
		creatorIds := make([]int64, 0)
		for _, dataset := range datasets {
			if dataset.CreatorID > 0 {
				creatorIds = append(creatorIds, dataset.CreatorID)
			}
		}
		if len(creatorIds) > 0 {
			creators, err := models.GetUsersByIDs(creatorIds)
			if err != nil {
				log.Error("GetUsersByIDs failed, req=%+v, err=%v", req, err)
			} else {
				creatorMap := make(map[int64]*models.User)
				for _, creator := range creators {
					creatorMap[creator.ID] = creator
				}
				for _, dataset := range datasets {
					if dataset.CreatorID > 0 {
						if creator, ok := creatorMap[dataset.CreatorID]; ok {
							dataset.CreatorName = creator.Name
						}
					}
				}
			}
		}
	}

	datasetInfos := make([]*entity.DatasetInfo, len(datasets))
	for i, dataset := range datasets {
		datasetInfos[i] = entity.BuildDatasetInfo(dataset)
	}

	return datasetInfos, total, nil
}

func ExportTaskResult2Dataset(req entity.ExportTaskResultReq) (string, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(req.TaskId)
	if err != nil {
		log.Error("ExportTaskResult GetCloudbrainByCloudbrainID failed, taskId=%d, err=%v", req.TaskId, err)
		return "", err
	}
	if req.FileListStr == "" {
		return "", errors.New("file list is empty")
	}
	fileList := strings.Split(req.FileListStr, ",")
	if fileList == nil || len(fileList) == 0 {
		return "", errors.New("file list is empty")
	}
	processId := redis_key.ExportTaskResultProcessIdKey(req.Dataset.ID, req.TaskId)
	msgMap := make(map[string]int, 0)
	for i := 0; i < len(fileList); i++ {
		msgMap[fileList[i]] = 0
	}
	go asyncToExportDataset(req.Dataset, fileList, cloudbrain, req.Doer, msgMap, processId)
	return processId, nil
}

func setProgress(msgKey string, msgMap map[string]int) {
	msgMapJson, _ := json.Marshal(msgMap)
	redisValue := string(msgMapJson)
	log.Info("set redis key=%s value=%s", msgKey, redisValue)
	_, err := redis_client.Setex(msgKey, redisValue, 3600*24*time.Second)
	if err != nil {
		log.Info("set redis error:%v", err)
	}
}

func asyncToExportDataset(dataset *models.DatasetRegistry, fileList []string, aiTask *models.Cloudbrain, doer *models.User, msgMap map[string]int, processId string) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	aiConfig := aiTask.GetCloudbrainConfig()
	if aiConfig == nil {
		log.Error("GetCloudbrainConfig empty, cloudbrain=%s", aiTask.JobName)
		return
	}
	sourceStorage := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aiConfig.OutputStorageType), aiConfig.OutputBucket)
	targetStorage := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if sourceStorage == nil || targetStorage == nil {
		log.Error("asyncToExportDataset GetUploadHelper failed, sourceStorage=%s, targetStorage=%s", aiConfig.OutputStorageType, dataset.StorageType)
		return
	}
	for _, shortFile := range fileList {
		//获取文件大小判断是否超出限额
		souceObjectKey := path.Join(aiConfig.OutputObjectPrefix, shortFile)
		meta, err := sourceStorage.GetObjectMeta(souceObjectKey)
		if err != nil {
			log.Error("GetObjectMeta failed, objectKey=%s, err=%v", souceObjectKey, err)
			msgMap[shortFile] = -1
			setProgress(processId, msgMap)
			continue
		}
		size := meta.ContentLength
		if storage_limit.IsSubjectFileUploadOverLimit(size, dataset.ConvertSubjectAccessContext()) {
			msgMap[shortFile] = -3
			setProgress(processId, msgMap)
			continue
		}

		//判断是否已存在
		targetObjectKey := path.Join(dataset.Path, shortFile)

		has, err := targetStorage.HasObject(targetObjectKey)
		if err != nil {
			log.Error("HasObject failed, objectKey=%s, err=%v", targetObjectKey, err)
			msgMap[shortFile] = -1
			setProgress(processId, msgMap)
			continue
		}
		if has {
			msgMap[shortFile] = -2
			setProgress(processId, msgMap)
			continue
		}
		//copy到数据集
		err = storage_helper.CopyFileBetweenStorage(sourceStorage, targetStorage, souceObjectKey, targetObjectKey)
		if err != nil {
			log.Error("CopyFileBetweenStorage failed, sourceObjectKey=%s, targetObjectKey=%s, err=%v", souceObjectKey, targetObjectKey, err)
			msgMap[shortFile] = -1
			setProgress(processId, msgMap)
			continue
		}
		DoAfterDatasetFileChanged(dataset.ID)
		msgMap[shortFile] = 100
		setProgress(processId, msgMap)
	}
}

func PutDatasetReadme(req entity.DatasetReadmeReq) error {
	dataset := req.Dataset
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if storageHelper == nil {
		return fmt.Errorf("GetUploadHelper failed, StorageType=%s", dataset.StorageType)
	}
	path := path.Join(dataset.Path, README_FILE_NAME)
	err := storageHelper.UploadFile(path, strings.NewReader(req.Content))
	return err
}

func GetOwnedPublicDatasetTags(userId int64) (*models.DatasetTags, error) {
	return models.GetOwnedPublicDatasetTags(userId)
}
