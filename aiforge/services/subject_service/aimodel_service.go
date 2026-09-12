package subject_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/storage_limit"
	uuid "github.com/satori/go.uuid"
)

func CreateAimodel(doer, owner *models.User, req entity.CreateAimodelReq) (*models.AiModelManage, *response.BizError) {
	if !NamePattern.MatchString(req.Name) {
		return nil, response.AIMODEL_NAME_INVALID
	}
	if req.Alias == "" {
		req.Alias = req.Name
	}
	if !AlphaDashDotChinese.MatchString(req.Alias) {
		return nil, response.AIMODEL_ALIAS_INVALID
	}

	redisKey := redis_key.AimodelNameLock(req.OwnerId, strings.ToLower(req.Name))
	lock := redis_lock.NewDistributeLock(redisKey)
	isOk, err := lock.LockWithWait(5*time.Second, 5*time.Second)
	if err != nil {
		log.Error("CreateAimodel LockWithWait failed, redisKey=%s, err=%v", redisKey, err)
		return nil, response.NewBizError(err)
	}
	if !isOk {
		log.Error("CreateAimodel LockWithWait failed, redisKey=%s, err=%v", redisKey, response.AIMODEL_NAME_EXIST.ToError())
		return nil, response.AIMODEL_NAME_EXIST
	}
	defer func() {
		err := lock.UnLock()
		if err != nil {
			log.Error("CreateAimodel UnLock failed, redisKey=%s, err=%v", redisKey, err)
		}
	}()

	record, err := models.GetAiModelByByOwnerAndName(req.OwnerId, req.Name)
	if err != nil && !models.IsErrRecordNotExist(err) {
		log.Error("CreateAimodel GetAiModelByByOwnerAndName failed, ownerId=%d, name=%s, err=%v", req.OwnerId, req.Name, err)
		return nil, response.NewBizError(err)
	}
	if record != nil {
		return nil, response.AIMODEL_NAME_EXIST
	}
	record, err = models.GetAiModelByOwnerAndAlias(req.OwnerId, req.Alias)
	if err != nil && !models.IsErrRecordNotExist(err) {
		log.Error("CreateAimodel GetAiModelByOwnerAndAlias failed, ownerId=%d, alias=%s, err=%v", req.OwnerId, req.Alias, err)
		return nil, response.NewBizError(err)
	}
	if record != nil {
		return nil, response.AIMODEL_ALIAS_EXIST
	}

	uploaderHelper := GetUploadHelper(models.AimodelSubject)
	storageType := uploaderHelper.GetStorageType()
	if storageType == "" {
		log.Info("aimodel storage type setting is empty")
		return nil, response.SYSTEM_ERROR
	}
	stroageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if stroageHelper == nil {
		log.Info("aimodel storage type %s not support", storageType)
		return nil, response.SYSTEM_ERROR
	}

	// 模型状态，线上与外部模型使用
	AimodelStatus := models.HfTransferStatusSuccess
	if req.AimodelType == models.MODEL_HF_TYPE {
		AimodelStatus = models.HfTransferStatusWaiting
	}

	// 训练任务信息,线上模型使用
	trainTaskInfo := ""
	if req.AimodelType == 0 && req.TaskId != 0 {
		aiTask, err := models.GetCloudbrainByCloudbrainID(req.TaskId)
		if err != nil {
			log.Error("ExportTaskResult GetCloudbrainByCloudbrainID failed, taskId=%d, err=%v", req.TaskId, err)
			return nil, response.AI_TASK_NOT_EXISTS
		}
		spec, err := resource.GetCloudbrainSpec(aiTask.ID)
		if err == nil {
			specJson, _ := json.Marshal(spec)
			aiTask.FlavorName = string(specJson)
		}
		aiTask.ContainerIp = ""
		aiTaskJson, _ := json.Marshal(aiTask)
		trainTaskInfo = string(aiTaskJson)
		AimodelStatus = models.HfTransferStatusOngoing
		parentModelIds := aiTask.GetModelIdArray()
		for _, modelId := range parentModelIds {
			err := models.ModifyModelDerivativeCount(modelId)
			if err != nil {
				log.Error("ModifyModelDerivativeCount faild." + err.Error())
			}
		}
	}

	aimodelId := uuid.NewV4().String()
	prefix := path.Join(stroageHelper.GetBucket(), uploaderHelper.GetStoragePathPrefix(aimodelId)) + "/"
	path, err := stroageHelper.AllocateDatasetNamespace(req.Name, prefix)
	if err != nil {
		log.Error("AllocateDatasetNamespace failed for aimodel, prefix=%s, err=%v", req.Name, err)
		return nil, response.NewBizError(err)
	}

	aimodelDTO := &models.AiModelManage{
		ID:            aimodelId,
		Name:          req.Name,
		Alias:         req.Alias,
		LowerAlias:    strings.ToLower(req.Alias),
		LowerName:     strings.ToLower(req.Name),
		UserId:        req.CreatorId,
		OwnerID:       req.OwnerId,
		ModelType:     req.AimodelType,
		License:       req.License,
		Label:         req.Label,
		Engine:        req.Engine,
		IsPrivate:     req.IsPrivate,
		StorageType:   string(storageType),
		Path:          path,
		ExternalName:  req.ExternalName,
		TrainTaskInfo: trainTaskInfo,
		Status:        AimodelStatus,
	}
	err = models.CreateAimodel(aimodelDTO, doer)
	if err != nil {
		log.Error("CreateAimodel failed, name=%s, err=%v", req.Name, err)
		return nil, response.NewBizError(err)
	}
	go notification.NotifyCreateAimodel(doer, owner, aimodelDTO)
	return aimodelDTO, nil
}

func DeleteAimodel(doer *models.User, aimodel *models.AiModelManage) error {
	owner, err := models.GetUserByID(aimodel.OwnerID)
	if err != nil {
		log.Error("GetUserByID failed, OwnerID=%s err=%v", aimodel.OwnerID, err)
		return err
	}

	err = models.WithTx(func(ctx models.DBContext) error {
		err := models.DeleteAimodel(ctx, doer, aimodel.OwnerID, aimodel.ID)
		if err != nil {
			log.Error("DeleteAimodel err,dataset=%+v,err=%v", aimodel, err)
			return err
		}
		err = DeleteAimodelStorage(aimodel)
		if err != nil {
			log.Error("DeleteAimodel DeleteAimodelStorage err,aimodel=%+v,err=%v", aimodel, err)
			return err
		}
		// delete hf transfer file record db
		if aimodel.ModelType == models.MODEL_HF_TYPE {
			err := models.DeleteHfFilesByModelIdWithContext(ctx, aimodel.ID)
			if err != nil {
				log.Info("DeleteHfFilesByModelId error." + err.Error())
				return err
			}
		}
		notification.NotifyDeleteAimodel(doer, owner, aimodel)
		return nil
	})

	return err
}

func DeleteAimodelStorage(aimodel *models.AiModelManage) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()

	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", models.AimodelSubject)
		return errors.New("Storage type error")
	}
	remainingFiles, err := helper.DeleteCollection(aimodel.Path)
	if err == nil && remainingFiles {
		tmpErr := models.InsertStorageDeleteFailedAimodel(aimodel)
		if tmpErr != nil {
			log.Error("InsertStorageDeleteFailedAimodel failed, aimodelId=%s, datasetName=%s err=%v", aimodel.ID, aimodel.Name, tmpErr)
		}
	}
	return err
}

func BatchDeleteAimodel(doer *models.User, aimodelIds []string) []string {
	successIds := make([]string, 0)
	for _, aimodelId := range aimodelIds {
		if aimodelId == "" {
			continue
		}
		aimodel, err := models.GetAimodelByID(aimodelId)
		if err != nil {
			log.Error("BatchDeleteDataset GetDatasetRegistryByID error.id=%s err=%v", aimodelId, err)
			continue
		}
		if !doer.CanDeleteAiModelManage(*aimodel) {
			log.Error("BatchDeleteDataset check permission failed.id=%s userId=%d", aimodelId, doer.ID)
			continue
		}
		err = DeleteAimodel(doer, aimodel)
		if err != nil {
			log.Error("BatchDeleteDataset DeleteDataset error.id=%s err=%v", aimodelId, err)
			continue
		}
		successIds = append(successIds, aimodelId)
	}
	return successIds
}

func CollectAimodel(user *models.User, aimodel *models.AiModelManage, isCollect bool) error {
	record := models.QueryModelCollectByUserId(aimodel.ID, user.ID)
	if isCollect {
		if record == nil || len(record) == 0 {
			err := models.SaveModelCollect(&models.AiModelCollect{
				ModelID: aimodel.ID,
				UserId:  user.ID,
			})
			if err != nil {
				return err
			}
		}
	} else {
		if record != nil && len(record) > 0 {
			err := models.DeleteModelCollect(&models.AiModelCollect{
				ID: record[0].ID,
			})
			if err != nil {
				return err
			}
		}
	}
	num := models.QueryModelCollectNum(aimodel.ID)
	models.ModifyModelCollectedNum(aimodel.ID, num)

	return nil
}

func RecommendAimodel(user *models.User, aimodel *models.AiModelManage) error {
	var err error
	err = models.ModifyModelRecommend(aimodel.ID, 1)
	if err != nil {
		return err
	}

	notification.NotifyAimodelRecommend(user, aimodel)
	return nil
}

func UnRecommendAimodel(user *models.User, aimodel *models.AiModelManage) error {
	var err error
	err = models.ModifyModelRecommend(aimodel.ID, 0)
	if err != nil {
		return err
	}
	return nil
}

func SearchAimodel(req models.SearchAimodelReq) ([]*entity.AimodelInfo, int64, *response.BizError) {
	aimodels, total, err := models.SearchAimodel(req)
	if err != nil {
		log.Error("SearchDatasets failed, req=%+v, err=%v", req, err)
		return nil, 0, response.NewBizError(err)
	}
	if aimodels == nil || len(aimodels) == 0 {
		log.Info("SearchDatasets no data found, req=%+v", req)
		return nil, 0, nil
	}

	if len(aimodels) > 0 {
		creatorIds := make([]int64, 0, len(aimodels))
		for _, aimodel := range aimodels {
			if aimodel.UserId > 0 {
				creatorIds = append(creatorIds, aimodel.UserId)
			}
		}
		if len(creatorIds) > 0 {
			creators, err := models.GetUsersByIDs(creatorIds)
			if err != nil {
				log.Error("GetUsersByIDs failed, req=%+v, err=%v", req, err)
			} else {
				creatorMap := make(map[int64]*models.User, len(creators))
				for _, creator := range creators {
					creatorMap[creator.ID] = creator
				}
				for _, aimodel := range aimodels {
					if aimodel.UserId > 0 {
						if creator, ok := creatorMap[aimodel.UserId]; ok {
							aimodel.CreatorName = creator.Name
						}
					}
				}
			}
		}
	}

	aimodelInfos := make([]*entity.AimodelInfo, len(aimodels))
	for i, aimodel := range aimodels {
		aimodelInfos[i] = entity.BuildAimodelInfo(aimodel)
	}

	return aimodelInfos, total, nil
}

func PutAimodelReadme(req entity.AimodelReadmeReq) error {
	aimodel := req.Aimodel
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if storageHelper == nil {
		return fmt.Errorf("GetUploadHelper failed, StorageType=%s", aimodel.StorageType)
	}
	path := path.Join(aimodel.Path, README_FILE_NAME)
	err := storageHelper.UploadFile(path, strings.NewReader(req.Content))
	return err
}

func DeleteAimodelFile(aimodel *models.AiModelManage, parentDir, fileName string) error {
	objectKey := path.Join(aimodel.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("StorageHelper failed, StorageType=%d", aimodel.StorageType)
		return errors.New("Storage type error")
	}
	err := helper.DeleteFile(objectKey)
	if err != nil {
		log.Error("DeleteFile failed, objectKey=%s err=%v", objectKey, err)
		return err
	}
	DoAfterAimodelFileChanged(aimodel.ID)
	return nil
}

func EditAimodel(req entity.CreateAimodelReq, aimodel *models.AiModelManage) *response.BizError {
	if !NamePattern.MatchString(req.Name) {
		return response.AIMODEL_NAME_INVALID
	}
	if req.Alias == "" {
		req.Alias = req.Name
	}
	if !AlphaDashDotChinese.MatchString(req.Alias) {
		return response.AIMODEL_ALIAS_INVALID
	}
	if aimodel.Name != req.Name {
		redisKey := redis_key.AimodelNameLock(aimodel.OwnerID, strings.ToLower(req.Name))
		lock := redis_lock.NewDistributeLock(redisKey)
		isOk, err := lock.LockWithWait(5*time.Second, 5*time.Second)
		if err != nil {
			log.Error("AimodelNameLock LockWithWait failed, redisKey=%s, err=%v", redisKey, err)
			return response.NewBizError(err)
		}
		if !isOk {
			log.Error("EditAimodel LockWithWait failed, redisKey=%s, err=%v", redisKey, response.AIMODEL_NAME_EXIST.ToError())
			return response.AIMODEL_NAME_EXIST
		}
		defer func() {
			err := lock.UnLock()
			if err != nil {
				log.Error("EditAimodel UnLock failed, redisKey=%s, err=%v", redisKey, err)
			}
		}()

		record, err := models.GetAiModelByByOwnerAndName(aimodel.OwnerID, req.Name)
		if err != nil {
			if !models.IsErrRecordNotExist(err) {
				log.Error("EditAimodel GetAiModelByByOwnerAndName failed, ownerId=%d, name=%s, err=%v", aimodel.OwnerID, req.Name, err)
				return response.NewBizError(err)
			}
		}
		if record != nil {
			return response.AIMODEL_NAME_EXIST
		}

	}
	if aimodel.Alias != req.Alias {
		record, err := models.GetAiModelByOwnerAndAlias(aimodel.OwnerID, req.Alias)
		if err != nil {
			if !models.IsErrRecordNotExist(err) {
				log.Error("EditAimodel GetAiModelByOwnerAndAlias failed, ownerId=%d, alias=%s, err=%v", aimodel.OwnerID, req.Alias, err)
				return response.NewBizError(err)
			}
		}
		if record != nil {
			return response.AIMODEL_ALIAS_EXIST
		}
	}
	err := models.UpdateAimodel(models.AiModelManage{
		ID:         aimodel.ID,
		Name:       req.Name,
		Alias:      req.Alias,
		Label:      req.Label,
		Engine:     req.Engine,
		License:    req.License,
		IsPrivate:  req.IsPrivate,
		LowerName:  strings.ToLower(req.Name),
		LowerAlias: strings.ToLower(req.Alias),
	})
	if err != nil {
		log.Error("UpdateAimodel failed, aimodelID=%s, err=%v", aimodel.ID, err)
		return response.NewBizError(err)
	}

	return nil
}

const sizeChangedDelayWindow = 5 * time.Minute

func getAimodelSizeChangedQueueBatchSize() int {
	if setting.AimodelSizeChangedQueueBatchSize <= 0 {
		return 10
	}
	return setting.AimodelSizeChangedQueueBatchSize
}

func EnqueueAimodelSizeChangedTask(aimodelId string) {
	score := float64(time.Now().Add(sizeChangedDelayWindow).Unix())
	if err := redis_client.ZAdd(redis_key.AimodelSizeChangedQueue(), aimodelId, score); err != nil {
		log.Error("EnqueueAimodelSizeChangedTask ZAdd failed, aimodelId=%s, err=%v", aimodelId, err)
	}
}

func ProcessAimodelSizeChangedQueue() {
	now := float64(time.Now().Unix())
	ids, err := redis_client.ZRangeByScoreLimit(redis_key.AimodelSizeChangedQueue(), 0, now, 0, getAimodelSizeChangedQueueBatchSize())
	if err != nil {
		log.Error("ProcessAimodelSizeChangedQueue ZRangeByScoreLimit failed, err=%v", err)
		return
	}
	for _, id := range ids {
		ReportAimodelSizeChanged(id)
		score, exists, err := redis_client.ZScore(redis_key.AimodelSizeChangedQueue(), id)
		if err != nil {
			log.Error("ProcessAimodelSizeChangedQueue ZScore failed, id=%s, err=%v", id, err)
			continue
		}
		if !exists || score > now {
			continue
		}
		if err := redis_client.ZRem(redis_key.AimodelSizeChangedQueue(), id); err != nil {
			log.Error("ProcessAimodelSizeChangedQueue ZRem failed, id=%s, err=%v", id, err)
		}
	}
}

func DoAfterAimodelFileChanged(aimodelId string) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	UpdateAimodelVersion(aimodelId)
	if UpdateAimodelSize(aimodelId) {
		EnqueueAimodelSizeChangedTask(aimodelId)
	}
}

func UpdateAimodelSize(aimodelId string) bool {
	aimodel, err := models.GetAimodelByID(aimodelId)
	if err != nil {
		log.Info("not found aimodel, id=%s", aimodelId)
		return false
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", aimodel.StorageType)
		return false
	}
	files, err := helper.GetAllObjectsUnderDir(aimodel.Path)
	if err != nil {
		log.Error("Failed to query size . id=%s err=%v", aimodelId, err)
		return false
	}
	var size int64
	for _, file := range files {
		size += file.Size
	}
	models.UpdateAimodelBySize(aimodelId, size)
	return true
}

func ReportAimodelSizeChanged(aimodelId string) {
	aimodel, err := models.GetAimodelByID(aimodelId)
	if err != nil {
		log.Info("not found aimodel, id=%s", aimodelId)
		return
	}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", aimodel.StorageType)
		return
	}
	if err := helper.ReportSizeChanged(aimodel.Path); err != nil {
		log.Error("ReportAimodelSizeChanged failed, aimodelId=%s err=%v", aimodelId, err)
	}
}

func ExportTaskResult2Aimodel(req *entity.AimodelExportTaskResultReq) (string, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(req.TaskId)
	if err != nil {
		log.Error("ExportTaskResult GetCloudbrainByCloudbrainID failed, taskId=%d, err=%v", req.TaskId, err)
		return "", err
	}
	if cloudbrain.UserID != req.Doer.ID && !req.Doer.IsAdmin {
		return "", errors.New("permission denied")
	}
	if req.FileListStr == "" {
		return "", errors.New("file list is empty")
	}
	fileList := strings.Split(req.FileListStr, ",")
	if fileList == nil || len(fileList) == 0 {
		return "", errors.New("file list is empty")
	}
	processId := redis_key.AimodelExportTaskResultProcessIdKey(req.Aimodel.ID, req.TaskId)
	msgMap := make(map[string]int, 0)
	for i := 0; i < len(fileList); i++ {
		msgMap[fileList[i]] = 0
	}
	log.Info("ExportTaskResult2Aimodel processId=%s, fileList=%v", processId, fileList)
	go asyncToExportAimodel(req.Aimodel, fileList, cloudbrain, req.Doer, msgMap, processId)
	return processId, nil
}

func asyncToExportAimodel(aimodel *models.AiModelManage, fileList []string, aiTask *models.Cloudbrain, doer *models.User, msgMap map[string]int, processId string) {
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
	targetStorage := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if sourceStorage == nil || targetStorage == nil {
		log.Error("asyncToExportAimodel GetUploadHelper failed, sourceStorage=%s, targetStorage=%s", aiConfig.OutputStorageType, aimodel.StorageType)
		return
	}
	var errExportStr string
	for _, shortFile := range fileList {
		//获取文件大小判断是否超出限额
		souceObjectKey := path.Join(aiConfig.OutputObjectPrefix, shortFile)
		meta, err := sourceStorage.GetObjectMeta(souceObjectKey)
		if err != nil {
			log.Error("GetObjectMeta failed, objectKey=%s, err=%v", souceObjectKey, err)
			msgMap[shortFile] = -1
			setProgress(processId, msgMap)
			errExportStr += err.Error() + ";"
			continue
		}
		size := meta.ContentLength
		if storage_limit.IsSubjectFileUploadOverLimit(size, aimodel.ConvertSubjectAccessContext()) {
			msgMap[shortFile] = -3
			setProgress(processId, msgMap)
			continue
		}

		//判断是否已存在
		targetObjectKey := path.Join(aimodel.Path, shortFile)

		has, err := targetStorage.HasObject(targetObjectKey)
		if err != nil {
			log.Error("HasObject failed, objectKey=%s, err=%v", targetObjectKey, err)
			msgMap[shortFile] = -1
			setProgress(processId, msgMap)
			errExportStr += err.Error() + ";"
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
			errExportStr += err.Error() + ";"
			continue
		}
		DoAfterAimodelFileChanged(aimodel.ID)
		msgMap[shortFile] = 100
		setProgress(processId, msgMap)
	}

	// 修改线上模型迁移状态
	if errExportStr != "" {
		err := models.UpdateAimodelByStatus(aimodel.ID, models.HfTransferStatusFailure, errExportStr)
		if err != nil {
			log.Error("UpdateAimodelByStatus faild." + err.Error())
		}
		log.Info("asyncToExportAimodel faild." + errExportStr)
	} else {
		err := models.UpdateAimodelByStatus(aimodel.ID, models.HfTransferStatusSuccess, "")
		if err != nil {
			log.Error("UpdateAimodelByStatus faild." + err.Error())
		}
	}
	//if aiTask.ModelId != "" {
	//	modelids := strings.Split(aiTask.ModelId, ";")
	//	for _, v := range modelids {
	//		log.Info("model:" + v + "  DerivativeCount +1")
	//		err := models.ModifyModelDerivativeCount(v)
	//		if err != nil {
	//			log.Error("ModifyModelDerivativeCount faild." + err.Error())
	//		}
	//	}
	//}
}

func IsAimodelVisibleByUser(doer *models.User, aimodel *models.AiModelManage) (bool, error) {
	if aimodel.IsPrivate {
		subjectContext := aimodel.ConvertSubjectAccessContext()
		permission, err := models.GetUserSubjectPermission(subjectContext, doer)
		if err != nil {
			return false, err
		}
		accessMode := permission.AccessMode
		if accessMode < models.AccessModeRead {
			return false, nil
		}
	}
	return true, nil
}

func QueryAimodelVisibility(doer *models.User, aimodel *models.AiModelManage) bool {
	visible, err := IsAimodelVisibleByUser(doer, aimodel)
	if err != nil {
		log.Error("QueryAimodelVisibility IsAimodelVisibleByUser failed, doerId=%d, aimodelId=%s, err=%v", doer.ID, aimodel.ID, err)
		return false
	}
	return visible
}

func GetOwnedPublicAimodelLabels(userId int64) (*models.AimodelLabels, error) {
	return models.GetOwnedPublicAimodelLabels(userId)
}

func GetAimodelRelatedData(aimodel *models.AiModelManage, user *models.User) (*entity.AimodelRelatedData, error) {
	cloudbrains, err := models.GetCloudBrainByModelId(aimodel.ID)
	if err != nil {
		log.Error("GetCloudBrainByModelId failed, modelId=%s, err=%v", aimodel.ID, err)
		return nil, err
	}
	if cloudbrains == nil || len(cloudbrains) <= 0 {
		log.Info("GetCloudBrainByModelId no data found, modelId=%s", aimodel.ID)
		return nil, err
	}

	var (
		datasets     []*entity.RelatedDataset
		repos        []*entity.RelatedRepo
		datasetIDSet = make(map[string]struct{})
		repoIDSet    = make(map[int64]struct{})
	)

	for _, task := range cloudbrains {
		if len(datasets) < 5 && task.Uuid != "" {
			for _, uuid := range strings.Split(task.Uuid, ";") {
				if len(datasets) >= 5 {
					break
				}
				if _, exists := datasetIDSet[uuid]; exists {
					continue
				}
				if dataset, err := models.GetDatasetRegistryByID(uuid); err == nil {
					if dataset.GetOwner() == nil {
						datasets = append(datasets, &entity.RelatedDataset{
							ID:        dataset.ID,
							Name:      dataset.Name,
							Alias:     dataset.Alias,
							OwnerName: dataset.Owner.Name,
						})
						datasetIDSet[uuid] = struct{}{}
					}
				}
			}
		}

		if task.RepoID != 0 {
			if _, exists := repoIDSet[task.RepoID]; exists {
				continue
			}
			if repo, err := models.GetRepositoryByID(task.RepoID); err == nil {
				if len(repos) < 5 {
					repos = append(repos, &entity.RelatedRepo{
						Name:      repo.Name,
						OwnerName: repo.OwnerName,
						Alias:     repo.Alias,
					})
				}
			}
		}
		repoIDSet[task.RepoID] = struct{}{}

		if len(datasets) >= 5 && len(repos) >= 5 {
			break
		}
	}

	res := &entity.AimodelRelatedData{
		Datasets: datasets,
		Repos:    repos,
	}
	return res, nil
}
