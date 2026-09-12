package models

import (
	"fmt"
	"time"

	"xorm.io/builder"
)

const (
	SortHot     = "hot"
	SortCreated = "created"

	StatusPending    = "pending"
	StatusExecuting  = "executing"
	StatusFinish     = "finish"
	StatusFailed     = "failed"
	StatusRolledBack = "rolled_back"
)

type ResourceUpdateBatch struct {
	ID             int64  `xorm:"pk autoincr"`
	Sort           string //排序，hot 按近一个月热度， created 按时间
	BatchSize      int    //批次大小
	MaxSize        int    //大小上限 0 为没上限
	RealBatchSize  int    //实际批次大小，可能小于BatchSize
	Status         string // pending → executing → success/failed → rolled_back
	SuccessCount   int
	FailedCount    int
	PendingCount   int
	ExecutingCount int
	RollbackCount  int
	ResourceType   string //dataset or model

	CreatedAt    time.Time  `xorm:"created"`
	ExecutedAt   *time.Time // 开始执行时间
	FinishedAt   *time.Time // 完成/失败时间
	RolledBackAt *time.Time // 回滚时间
}

type ResourceUpdateBatchItem struct {
	ID           int64  `xorm:"pk autoincr"`
	BatchId      int64  //批次ID
	ResourceType string //dataset or model
	ResourceID   string `xorm:"uuid"`
	OldConfig    string //
	NewConfig    string //
	Status       string // pending → executing → success/failed/rolled_back
	ExecutedAt   *time.Time
	ErrorMsg     string // 失败原因
}

type ResourceItemDetail struct {
	ResourceID  string
	ItemID      int64
	Name        string
	UseCount    int64
	Size        int64
	CreatedAt   time.Time
	StorageType string
	Path        string
}

type ResourceBatchDetail struct {
	ResourceUpdateBatch
	Items []ResourceItemDetail
}

type CreateBatchReq struct {
	ResourceType string `json:"resource_type" binding:"Required"`
	Sort         string `json:"sort" binding:"Required"`
	BatchSize    int    `json:"batch_size" binding:"Required"`
	MaxSize      int    `json:"max_size"`
	ResourceIDs  []string
}

func GetResourceUpdateBatchByID(id int64) (*ResourceUpdateBatch, error) {
	batch := &ResourceUpdateBatch{ID: id}
	has, err := x.Get(batch)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("batch not found: %d", id)
	}
	return batch, nil
}

func LockResourceUpdateBatch(batchId int64) error {
	batch := &ResourceUpdateBatch{ID: batchId}
	_, err := x.ID(batchId).ForUpdate().Get(batch)
	return err
}

func UpdateResourceUpdateBatchStatus(batchId int64, status string) error {
	batch := &ResourceUpdateBatch{
		Status: status,
	}
	_, err := x.ID(batchId).Cols("status").Update(batch)
	return err
}

func ListResourceUpdateBatchItemsByBatchId(batchId int64) ([]ResourceUpdateBatchItem, error) {
	var items []ResourceUpdateBatchItem
	err := x.Where("batch_id = ?", batchId).Find(&items)
	return items, err
}

func ListResourceUpdateBatchItemsByBatchIdAndStatus(batchId int64, status []string) ([]ResourceUpdateBatchItem, error) {
	var items []ResourceUpdateBatchItem
	err := x.Where("batch_id = ?", batchId).In("status", status).Find(&items)
	return items, err
}

func ListResourceUpdateBatchItemsByResourceIdsAndType(resourceType string, resourceIds []string) ([]ResourceUpdateBatchItem, error) {
	var items []ResourceUpdateBatchItem
	err := x.Where("resource_type = ?", resourceType).In("resource_id", resourceIds).Find(&items)
	return items, err
}

func UpdateDatasetRegistryUrchinInfo(itemId int64, dataset DatasetRegistry, objUuid string) error {
	//更新item的old config和new config
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	_, err := sess.ID(itemId).Cols("old_config", "new_config", "status").Update(&ResourceUpdateBatchItem{
		OldConfig: dataset.Path,
		NewConfig: objUuid,
		Status:    StatusFinish,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	newDataset := &DatasetRegistry{
		Path:        objUuid,
		StorageType: "URCHIN_V2",
	}
	_, err = x.ID(dataset.ID).Cols("path", "storage_type").Update(newDataset)
	if err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func UpdateModelUrchinInfo(itemId int64, model AiModelManage, objUuid string) error {
	//更新item的old config和new config
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	_, err := sess.ID(itemId).Cols("old_config", "new_config", "status").Update(&ResourceUpdateBatchItem{
		OldConfig: model.Path,
		NewConfig: objUuid,
		Status:    StatusFinish,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	newModel := &AiModelManage{
		Path:        objUuid,
		StorageType: "URCHIN_V2",
	}
	_, err = x.ID(model.ID).Cols("path", "storage_type").Update(newModel)
	if err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

//CreateBatch 生成批次
func CreateBatch(req *CreateBatchReq) (*ResourceUpdateBatch, error) {
	batch := &ResourceUpdateBatch{
		ResourceType: req.ResourceType,
		Sort:         req.Sort,
		BatchSize:    req.BatchSize,
		MaxSize:      req.MaxSize,
		Status:       StatusPending,
	}
	_, err := x.Insert(batch)
	if err != nil {
		return nil, err
	}
	//根据规则查出对应的资源ID，生成batch item
	var resourceIds []string
	if req.ResourceType == "dataset" {
		resourceIds, err = GetDatasetItemsForUpdate(req.Sort, req.MaxSize, req.BatchSize)
	} else if req.ResourceType == "model" {
		resourceIds, err = GetModelItemsForUpdate(req.Sort, req.MaxSize, req.BatchSize)
	} else {
		return nil, fmt.Errorf("invalid resource type: %s", req.ResourceType)
	}
	if err != nil {
		return nil, err
	}
	var items []ResourceUpdateBatchItem
	for _, resourceId := range resourceIds {
		items = append(items, ResourceUpdateBatchItem{
			ResourceType: req.ResourceType,
			ResourceID:   resourceId,
			Status:       StatusPending,
			BatchId:      batch.ID,
		})
	}
	if len(items) > 0 {
		_, err = x.Insert(&items)
		if err != nil {
			return nil, err
		}
		batch.RealBatchSize = len(items)
		_, err = x.ID(batch.ID).Cols("real_batch_size").Update(batch)
		if err != nil {
			return nil, err
		}
	}
	return batch, err
}

//CreateBatch4Resource 为指定资源生成批次
func CreateBatch4Resource(req *CreateBatchReq) (*ResourceUpdateBatch, error) {
	batch := &ResourceUpdateBatch{
		ResourceType: req.ResourceType,
		Status:       StatusPending,
	}
	_, err := x.Insert(batch)
	if err != nil {
		return nil, err
	}
	//根据规则查出对应的资源ID，生成batch item
	var resourceIds = req.ResourceIDs
	var items []ResourceUpdateBatchItem
	for _, resourceId := range resourceIds {
		items = append(items, ResourceUpdateBatchItem{
			ResourceType: req.ResourceType,
			ResourceID:   resourceId,
			Status:       StatusPending,
			BatchId:      batch.ID,
		})
	}
	if len(items) > 0 {
		_, err = x.Insert(&items)
		if err != nil {
			return nil, err
		}
		batch.RealBatchSize = len(items)
		_, err = x.ID(batch.ID).Cols("real_batch_size").Update(batch)
		if err != nil {
			return nil, err
		}
	}
	return batch, err
}

func GetDatasetItemsForUpdate(sort string, maxSize int, batchSize int) ([]string, error) {
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"storage_type": "OBS"})
	if maxSize > 0 {
		cond = cond.And(builder.Lte{"size": maxSize})
	}
	if batchSize == 0 {
		batchSize = 50
	}
	var orderBy string
	if sort == SortHot {
		orderBy = "use_count DESC"
	} else {
		orderBy = "id DESC"
	}
	//还得不在任何batch里
	subQuery := builder.Select("resource_id").From("resource_update_batch_item").Where(builder.Eq{"resource_type": "dataset"})
	cond = cond.And(builder.NotIn("id", subQuery))
	var datasets []DatasetRegistry
	err := x.Where(cond).OrderBy(orderBy).Limit(batchSize).Find(&datasets)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, d := range datasets {
		ids = append(ids, d.ID)
	}
	return ids, nil
}

func GetModelItemsForUpdate(sort string, maxSize int, batchSize int) ([]string, error) {
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"storage_type": "OBS"})
	if maxSize > 0 {
		cond = cond.And(builder.Lte{"size": maxSize})
	}
	if batchSize == 0 {
		batchSize = 50
	}
	var orderBy string
	if sort == SortHot {
		orderBy = "reference_count DESC"
	} else {
		orderBy = "id DESC"
	}
	subQuery := builder.Select("resource_id::varchar").From("resource_update_batch_item").Where(builder.Eq{"resource_type": "model"})
	cond = cond.And(builder.NotIn("id", subQuery))
	var modelList []AiModelManage
	err := x.Where(cond).OrderBy(orderBy).Limit(batchSize).Find(&modelList)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, d := range modelList {
		ids = append(ids, d.ID)
	}
	return ids, nil
}

func QueryBatchList(resourceType, status, resourceId string, batchID int64) ([]ResourceUpdateBatch, error) {
	cond := builder.NewCond()
	if resourceType != "" {
		cond = cond.And(builder.Eq{"resource_type": resourceType})
	}
	if status != "" {
		cond = cond.And(builder.Eq{"status": status})
	}
	if resourceId != "" {
		subQuery := builder.Select("batch_id").From("resource_update_batch_item").Where(builder.Eq{"resource_id": resourceId})
		cond = cond.And(builder.In("id", subQuery))
	}
	if batchID != 0 {
		cond = cond.And(builder.Eq{"id": batchID})
	}
	var batches []ResourceUpdateBatch
	err := x.Where(cond).OrderBy("id DESC").Find(&batches)
	//查询每个batch的成功数和失败数
	for i := range batches {
		//invalid opration can not index counts
		type CountResult struct {
			Status string
			Count  int
		}
		var results []CountResult
		err = x.Table("resource_update_batch_item").Select("status, count(*) as count").Where("batch_id = ?", batches[i].ID).GroupBy("status").Find(&results)
		if err != nil {
			return nil, err
		}
		for _, r := range results {
			if r.Status == StatusFinish {
				batches[i].SuccessCount = r.Count
			} else if r.Status == StatusFailed {
				batches[i].FailedCount = r.Count
			} else if r.Status == StatusPending {
				batches[i].PendingCount = r.Count
			} else if r.Status == StatusExecuting {
				batches[i].ExecutingCount = r.Count
			} else if r.Status == StatusRolledBack {
				batches[i].RollbackCount = r.Count
			}
		}
	}
	return batches, err
}

func QueryBatchDetailByID(batchId int64) (*ResourceBatchDetail, error) {
	var batch ResourceUpdateBatch
	has, err := x.ID(batchId).Get(&batch)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("batch not found: %d", batchId)
	}
	var items []ResourceUpdateBatchItem
	err = x.Where("batch_id = ?", batchId).Find(&items)
	if err != nil {
		return nil, err
	}
	detail := &ResourceBatchDetail{
		ResourceUpdateBatch: batch,
		Items:               make([]ResourceItemDetail, 0),
	}
	//按类型批量聚合查询对应的资源具体情况，先组合资源id，再统一查询
	var datasetIds []string
	var modelIds []string
	for _, item := range items {
		if item.ResourceType == "dataset" {
			datasetIds = append(datasetIds, item.ResourceID)
		} else if item.ResourceType == "model" {
			modelIds = append(modelIds, item.ResourceID)
		}
	}
	datasetMap := make(map[string]DatasetRegistry)
	if len(datasetIds) > 0 {
		var datasets []DatasetRegistry
		err = x.In("id", datasetIds).Find(&datasets)
		if err != nil {
			return nil, err
		}
		for _, d := range datasets {
			datasetMap[d.ID] = d
		}
	}
	modelMap := make(map[string]AiModelManage)
	if len(modelIds) > 0 {
		var models []AiModelManage
		err = x.In("id", modelIds).Find(&models)
		if err != nil {
			return nil, err
		}
		for _, m := range models {
			modelMap[m.ID] = m
		}
	}
	//在user表里查询owner的名字，
	ownerIds := make(map[int64]struct{})
	for _, item := range items {
		if item.ResourceType == "dataset" {
			if dataset, ok := datasetMap[item.ResourceID]; ok {
				ownerIds[dataset.OwnerID] = struct{}{}
			}
		} else if item.ResourceType == "model" {
			if model, ok := modelMap[item.ResourceID]; ok {
				ownerIds[model.OwnerID] = struct{}{}
			}
		}
	}
	ownerMap := make(map[int64]User)
	if len(ownerIds) > 0 {
		var users []User
		ownerIDs := make([]int64, 0, len(ownerIds))
		for id := range ownerIds {
			ownerIDs = append(ownerIDs, id)
		}
		err = x.In("id", ownerIDs).Find(&users)
		if err != nil {
			return nil, err
		}
		for _, user := range users {
			ownerMap[user.ID] = user
		}
	}

	//组合结果
	for _, item := range items {
		if item.ResourceType == "dataset" {
			if dataset, ok := datasetMap[item.ResourceID]; ok {
				ownerName := ""
				if owner, ok := ownerMap[dataset.OwnerID]; ok {
					ownerName = owner.Name
				}
				detail.Items = append(detail.Items, ResourceItemDetail{
					ResourceID:  item.ResourceID,
					ItemID:      item.ID,
					Name:        ownerName + "/" + dataset.Name,
					UseCount:    dataset.UseCount,
					Size:        dataset.Size,
					CreatedAt:   time.Unix(int64(dataset.CreatedUnix), 0),
					StorageType: dataset.StorageType,
					Path:        dataset.Path,
				})
			}
		} else if item.ResourceType == "model" {

			if model, ok := modelMap[item.ResourceID]; ok {
				ownerName := ""
				if owner, ok := ownerMap[model.OwnerID]; ok {
					ownerName = owner.Name
				}
				detail.Items = append(detail.Items, ResourceItemDetail{
					ResourceID:  item.ResourceID,
					ItemID:      item.ID,
					Name:        ownerName + "/" + model.Name,
					UseCount:    int64(model.ReferenceCount),
					Size:        model.Size,
					CreatedAt:   time.Unix(int64(model.CreatedUnix), 0),
					StorageType: model.StorageType,
					Path:        model.Path,
				})
			}
		}
	}

	return detail, nil
}

func RollbackDataset(batchItem ResourceUpdateBatchItem) error {
	//更新item的old config和new config
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	_, err := sess.ID(batchItem.ID).Cols("status").Update(&ResourceUpdateBatchItem{
		Status: StatusRolledBack,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	newDataset := &DatasetRegistry{
		Path:        batchItem.OldConfig,
		StorageType: "OBS",
	}
	_, err = x.ID(batchItem.ResourceID).Cols("path", "storage_type").Update(newDataset)
	if err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func RollbackModel(batchItem ResourceUpdateBatchItem) error {
	//更新item的old config和new config
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	_, err := sess.ID(batchItem.ID).Cols("status").Update(&ResourceUpdateBatchItem{
		Status: StatusRolledBack,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	newModel := &AiModelManage{
		Path:        batchItem.OldConfig,
		StorageType: "OBS",
	}
	_, err = sess.ID(batchItem.ResourceID).Cols("path", "storage_type").Update(newModel)
	if err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

type ResourceUpdateStatisticDetail struct {
	ResourceItemTotalCount      int64
	ResourceItemSuccessCount    int64
	ResourceItemFailedCount     int64
	ResourceItemPendingCount    int64
	ResourceItemExecutingCount  int64
	ResourceItemRolledBackCount int64
	ResourceOBSCount            int64
	ResourceURCHINCount         int64
}

func CountResourceUpdatedData() (map[string]ResourceUpdateStatisticDetail, error) {
	//需要按资源类型统计每一种状态的数量，只需要数量即可
	type CountResult struct {
		ResourceType string
		Status       string
		Count        int64
	}
	var results []CountResult
	err := x.Table("resource_update_batch_item").Select("resource_type, status, count(*) as count").GroupBy("resource_type, status").Find(&results)
	if err != nil {
		return nil, err
	}
	data := make(map[string]ResourceUpdateStatisticDetail)
	for _, r := range results {
		stat, ok := data[r.ResourceType]
		if !ok {
			stat = ResourceUpdateStatisticDetail{}
		}
		stat.ResourceItemTotalCount += r.Count
		if r.Status == StatusFinish {
			stat.ResourceItemSuccessCount = r.Count
		} else if r.Status == StatusFailed {
			stat.ResourceItemFailedCount = r.Count
		} else if r.Status == StatusPending {
			stat.ResourceItemPendingCount = r.Count
		} else if r.Status == StatusExecuting {
			stat.ResourceItemExecutingCount = r.Count
		} else if r.Status == StatusRolledBack {
			stat.ResourceItemRolledBackCount = r.Count
		}
		data[r.ResourceType] = stat
	}
	//再统计每种资源类型下，不同存储类型的数量
	type StorageCountResult struct {
		ResourceType string
		StorageType  string
		Count        int64
	}
	var storageResults []StorageCountResult
	err = x.Table("dataset_registry").Select("storage_type, count(*) as count, 'dataset' as resource_type").GroupBy("storage_type").Find(&storageResults)

	if err != nil {
		return nil, err
	}
	for _, r := range storageResults {
		stat, ok := data[r.ResourceType]
		if !ok {
			stat = ResourceUpdateStatisticDetail{}
		}
		if r.StorageType == "URCHIN_V2" {
			stat.ResourceURCHINCount = r.Count
		} else if r.StorageType == "OBS" {
			stat.ResourceOBSCount = r.Count
		}
		data[r.ResourceType] = stat
	}
	storageResults = []StorageCountResult{}
	err = x.Table("ai_model_manage").Select("storage_type, count(*) as count , 'model' as resource_type").GroupBy("storage_type").Find(&storageResults)
	if err != nil {
		return nil, err
	}
	for _, r := range storageResults {
		stat, ok := data[r.ResourceType]
		if !ok {
			stat = ResourceUpdateStatisticDetail{}
		}
		if r.StorageType == "URCHIN_V2" {
			stat.ResourceURCHINCount = r.Count
		} else if r.StorageType == "OBS" {
			stat.ResourceOBSCount = r.Count
		}
		data[r.ResourceType] = stat
	}
	return data, nil
}

func UpdateResourceUpdateBatchItemStatus(itemId int64, status string, errorMsg string) error {
	item := &ResourceUpdateBatchItem{
		ID:       itemId,
		Status:   status,
		ErrorMsg: errorMsg,
	}
	_, err := x.ID(itemId).Cols("status", "error_msg").Update(item)
	return err
}
