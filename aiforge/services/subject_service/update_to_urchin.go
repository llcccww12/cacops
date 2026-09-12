package subject_service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/urchin_v2"
)

//现在补上必要的log，均用Resource to urchin2.0.作为开头，方便后续查询日志定位问题

func UpdateByBatchId(batchId int64, token string, node string, userID string) error {
	//1、查询batch信息
	batch, err := models.GetResourceUpdateBatchByID(batchId)
	if err != nil {
		return err
	}
	if batch.Status == models.StatusFinish {
		log.Info("Resource to urchin2.0.UpdateByBatchId batchId = %d has been successfully executed, skip, current status:%d", batchId, batch.Status)
		return nil
	}
	if batch.Status != models.StatusPending && batch.Status != models.StatusFailed && batch.Status != models.StatusRolledBack {
		log.Info("Resource to urchin2.0.UpdateByBatchId batchId = %d status is not waiting, current status:%d", batchId, batch.Status)
		return fmt.Errorf("batch can not excute,status = %s", batch.Status)
	}
	//加一个redis分布式锁
	lock := redis_lock.NewDistributeLock("update_to_urchin:update:lock")
	success, err := lock.Lock(5 * time.Minute)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("batch is being processed by other worker")
	}
	defer lock.UnLock()
	//2、更新状态为执行中
	err = models.UpdateResourceUpdateBatchStatus(batchId, models.StatusExecuting)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			models.UpdateResourceUpdateBatchStatus(batchId, models.StatusFailed)
		}
	}()
	//3、查询明细，顺序执行,修改明细状态
	items, err := models.ListResourceUpdateBatchItemsByBatchIdAndStatus(batchId, []string{models.StatusPending, models.StatusFailed, models.StatusRolledBack})
	if err != nil {
		return err
	}
	successCount := 0
	failedCount := 0
	for _, item := range items {
		log.Info("Resource to urchin2.0. batchID = %d successCount: %d, failedCount: %d", batchId, successCount, failedCount)
		log.Info("Resource to urchin2.0.Start to update batch item, batchID: %d, itemID: %d, resourceID: %s, resourceType: %s", batchId, item.ID, item.ResourceID, item.ResourceType)
		err = UpdateBatchItem(item, token, node, userID)
		if err != nil {
			log.Error("Resource to urchin2.0.UpdateBatchItem failed, itemID: %d, err: %v", item.ID, err)
			failedCount++
			//更新item状态为失败
			//截断消息，防止过长的错误消息导致数据库更新失败
			errorMsg := err.Error()
			if len(errorMsg) > 200 {
				errorMsg = errorMsg[:200]
			}
			models.UpdateResourceUpdateBatchItemStatus(item.ID, models.StatusFailed, errorMsg)
			continue
		}
		successCount++
		log.Info("Resource to urchin2.0.UpdateBatchItem success, batchID: %d, itemID: %d", batchId, item.ID)
	}
	//5、更新batch状态为成功/失败
	if failedCount > 0 {
		err = models.UpdateResourceUpdateBatchStatus(batchId, models.StatusFailed)
		if err != nil {
			return err
		}
		log.Info("Resource to urchin2.0.UpdateByBatchId batchID = %d finished with failed status, successCount: %d, failedCount: %d", batchId, successCount, failedCount)
		return fmt.Errorf("batch executed with some failed items, successCount: %d, failedCount: %d", successCount, failedCount)
	} else {
		err = models.UpdateResourceUpdateBatchStatus(batchId, models.StatusFinish)
		if err != nil {
			return err
		}
		log.Info("Resource to urchin2.0.UpdateByBatchId batchID = %d finished with success status, successCount: %d, failedCount: %d", batchId, successCount, failedCount)
	}

	return nil
}

//升级某个批次中的某个明细
func UpdateBatchItem(batchItemId models.ResourceUpdateBatchItem, token, node, userID string) error {
	//根据类型处理
	switch batchItemId.ResourceType {
	case "dataset":
		//处理dataset升级
		err := UpdateDatasetToUrchin(batchItemId, token, node, userID)
		if err != nil {
			log.Error("Resource to urchin2.0.UpdateDatasetToUrchin batchId = %d err:%v", batchItemId.ID, err)
			return err
		}
	case "model":
		//处理model升级
		err := UpdateModelToUrchin(batchItemId, token, node, userID)
		if err != nil {
			log.Error("Resource to urchin2.0.UpdateModelToUrchin batchId = %d err:%v", batchItemId.ID, err)
			return err
		}
	default:
		return nil
	}
	return nil
}

func UpdateDatasetToUrchin(batchItem models.ResourceUpdateBatchItem, token, node, userID string) error {
	//1、查询是否已升级
	dataset, err := models.GetDatasetRegistryByID(batchItem.ResourceID)
	if err != nil {
		return err
	}
	if dataset.StorageType == "URCHIN_V2" {
		log.Info("Resource to urchin2.0.dataset %s has been upgraded to urchin, skip, datasetID: %d", dataset.Name, dataset.ID)
		return nil
	}
	if dataset.StorageType != "OBS" {
		log.Error("Resource to urchin2.0.dataset %s has invalid storage type: %s, datasetID: %d", dataset.Name, dataset.StorageType, dataset.ID)
		return errors.New("invalid storage type")
	}
	//2、调用升级接口
	urchinClient := urchin_v2.NewUrchinClient(setting.UrchinAddress, setting.UrchinReqTimeoutSecond, setting.UrchinMaxConnection)

	urchinClient.UrchinClient.SetToken(token)
	path := strings.TrimPrefix(dataset.Path, setting.Bucket+"/")

	err, objUuid := urchinClient.RelateObject(dataset.Name, "", node, path, userID)
	if err != nil {
		log.Error("Resource to urchin2.0.RelateObject failed, datasetID: %d, err: %v", dataset.ID, err)
		return err
	}
	//3、更新数据本身状态+item的状态
	err = models.UpdateDatasetRegistryUrchinInfo(batchItem.ID, *dataset, objUuid)
	//4、返回结果
	return err
}

func UpdateModelToUrchin(batchItem models.ResourceUpdateBatchItem, token, node, userID string) error {
	//1、查询是否已升级
	model, err := models.QueryModelById(batchItem.ResourceID)
	if err != nil {
		return err
	}
	if model.StorageType == "URCHIN_V2" {
		log.Info("Resource to urchin2.0.model %s has been upgraded to urchin, skip, modelID: %d", model.Name, model.ID)
		return nil
	}
	if model.StorageType != "OBS" {
		log.Error("Resource to urchin2.0.model %s has invalid storage type: %s, modelID: %d", model.Name, model.StorageType, model.ID)
		return errors.New("invalid storage type")
	}
	//2、调用升级接口
	urchinClient := urchin_v2.NewUrchinClient(setting.UrchinAddress, setting.UrchinReqTimeoutSecond, setting.UrchinMaxConnection)

	urchinClient.UrchinClient.SetToken(token)

	path := strings.TrimPrefix(model.Path, setting.Bucket+"/")

	err, objUuid := urchinClient.RelateObject(model.Name, "", node, path, userID)
	if err != nil {
		return err
	}
	//3、更新数据本身状态+item的状态
	err = models.UpdateModelUrchinInfo(batchItem.ID, *model, objUuid)
	//4、返回结果
	return err
}

func RollbackByBatchID(batchId int64) error {
	//1、查询batch信息
	batch, err := models.GetResourceUpdateBatchByID(batchId)
	if err != nil {
		return err
	}
	if batch.Status != models.StatusFinish && batch.Status != models.StatusFailed {
		log.Info("Resource to urchin2.0.RollbackByBatchID batchId = %d status is not success, current status:%d", batchId, batch.Status)
		return fmt.Errorf("batch is not success,status = %s", batch.Status)
	}
	//2、查询明细，顺序执行,修改明细状态
	items, err := models.ListResourceUpdateBatchItemsByBatchId(batchId)
	if err != nil {
		return err
	}
	for _, item := range items {
		err = RollbackBatchItem(item)
		if err != nil {
			log.Error("RollbackBatchItem failed, itemID: %d, err: %v", item.ID, err)
			continue
		}
	}
	//3、更新batch状态为已回滚
	err = models.UpdateResourceUpdateBatchStatus(batchId, models.StatusRolledBack)
	if err != nil {
		return err
	}
	return nil
}

func RollbackBatchItem(batchItem models.ResourceUpdateBatchItem) error {
	//根据类型处理
	switch batchItem.ResourceType {
	case "dataset":
		//处理dataset回滚
		err := models.RollbackDataset(batchItem)
		if err != nil {
			log.Error("Resource to urchin2.0.RollbackDataset failed, itemID: %d, err: %v", batchItem.ID, err)
			return err
		}
	case "model":
		//处理model回滚
		err := models.RollbackModel(batchItem)
		if err != nil {
			log.Error("Resource to urchin2.0.RollbackModel failed, itemID: %d, err: %v", batchItem.ID, err)
			return err
		}
	default:
		return nil
	}
	return nil
}

func UpdateUrchinByResourceIds(resourceType, token, node, userID string, resourceIds []string) (int64, error) {
	//加一个redis分布式锁
	lock := redis_lock.NewDistributeLock("update_to_urchin:update:lock")
	success, err := lock.Lock(5 * time.Minute)
	if err != nil {
		return 0, err
	}
	if !success {
		return 0, errors.New("batch is being processed by other worker")
	}
	defer lock.UnLock()
	//判断资源是否存在
	switch resourceType {
	case "dataset":
		datasets, err := models.GetDatasetRegistryListByIDs(resourceIds)
		if err != nil {
			return 0, err
		}
		if len(datasets) != len(resourceIds) {
			return 0, errors.New("resource count mismatch")
		}
	case "model":
		aiModels, err := models.QueryModelByIds(resourceIds)
		if err != nil {
			return 0, err
		}
		if len(aiModels) != len(resourceIds) {
			return 0, errors.New("resource count mismatch")
		}
	default:
		return 0, errors.New("invalid resource type")
	}
	//排除已经创建的item
	existingItems, err := models.ListResourceUpdateBatchItemsByResourceIdsAndType(resourceType, resourceIds)
	if err != nil {
		return 0, err
	}
	if len(existingItems) > 0 {
		itemMap := make(map[string]int64)
		for _, item := range existingItems {
			itemMap[item.ResourceID] = item.BatchId
		}
		return 0, fmt.Errorf("some resources are already in batch, resourceType: %s,itemMap: %v", resourceType, itemMap)
	}
	//创建batch和item，维护好状态
	req := &models.CreateBatchReq{
		ResourceType: resourceType,
		ResourceIDs:  resourceIds,
	}
	batch, err := models.CreateBatch4Resource(req)
	if err != nil {
		return 0, err
	}
	log.Info("Resource to urchin2.0.Create batch for resourceIDs success, batchID: %d", batch.ID)
	err = models.UpdateResourceUpdateBatchStatus(batch.ID, models.StatusExecuting)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			models.UpdateResourceUpdateBatchStatus(batch.ID, models.StatusFailed)
		}
	}()
	items, err := models.ListResourceUpdateBatchItemsByBatchIdAndStatus(batch.ID, []string{models.StatusPending, models.StatusFailed, models.StatusRolledBack})
	if err != nil {
		return 0, err
	}
	for _, item := range items {
		err = UpdateBatchItem(item, token, node, userID)
		if err != nil {
			log.Error("Resource to urchin2.0.UpdateBatchItem failed, itemID: %d, err: %v", item.ID, err)
			continue
		}
	}
	//更新batch状态为完成
	err = models.UpdateResourceUpdateBatchStatus(batch.ID, models.StatusFinish)
	if err != nil {
		return 0, err
	}
	return batch.ID, nil
}
