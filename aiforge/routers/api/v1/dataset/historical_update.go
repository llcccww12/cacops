package dataset

import (
	"fmt"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/subject_service"
)

func CreateBatch(ctx *context.APIContext) {
	sourt := ctx.Query("sort")
	batchSize := ctx.QueryInt("batch_size")
	maxSize := ctx.QueryInt("max_size")
	resourceType := ctx.Query("resource_type")
	//1、根据入参创建批次
	batch, err := models.CreateBatch(&models.CreateBatchReq{
		ResourceType: resourceType,
		Sort:         sourt,
		BatchSize:    batchSize,
		MaxSize:      maxSize,
	})
	if err != nil {
		log.Error("Resource to urchin2.0. CreateBatch failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("Resource to urchin2.0.CreateBatch success, batchID: %d", batch.ID)
	ctx.JSON(200, response.OuterSuccessWithData(batch))
}

func QueryBatchList(ctx *context.APIContext) {
	resourceType := ctx.Query("resource_type")
	status := ctx.Query("status")
	resourceId := ctx.Query("resource_id")
	batchID := ctx.QueryInt64("batch_id")

	batches, err := models.QueryBatchList(resourceType, status, resourceId, batchID)
	if err != nil {
		log.Error("Resource to urchin2.0.QueryBatchList failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("Resource to urchin2.0.QueryBatchList success, batch count: %d", len(batches))
	ctx.JSON(200, response.OuterSuccessWithData(batches))
}

func QueryBatchDetail(ctx *context.APIContext) {
	batchID := ctx.QueryInt64("batch_id")

	batch, err := models.QueryBatchDetailByID(batchID)
	if err != nil {
		log.Error("QueryBatchDetailByID failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("QueryBatchDetailByID success, batchID: %d", batchID)
	ctx.JSON(200, response.OuterSuccessWithData(batch))
}

func UpdateByBatch(ctx *context.APIContext) {
	batchID := ctx.QueryInt64("batch_id")
	userID := ctx.Query("user_id")
	token := ctx.Query("admin_token")
	node := ctx.Query("node")

	err := subject_service.UpdateByBatchId(batchID, token, node, userID)
	if err != nil {
		log.Error("Resource to urchin2.0.UpdateByBatch failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("Resource to urchin2.0.UpdateByBatch success, batchID: %d", batchID)
	ctx.JSON(200, response.OuterSuccess())
}

func UpdateToUrchinByResourceIds(ctx *context.APIContext) {
	userID := ctx.Query("user_id")
	token := ctx.Query("admin_token")
	node := ctx.Query("node")
	resourceType := ctx.Query("resource_type")
	resourceIdsTr := ctx.Query("resource_ids")
	resourceIds := strings.Split(resourceIdsTr, ",")
	if len(resourceIds) == 0 {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(fmt.Errorf("resourceIds can not be empty")), ctx.Locale))
		return
	}
	batchID, err := subject_service.UpdateUrchinByResourceIds(resourceType, token, node, userID, resourceIds)
	if err != nil {
		log.Error("Resource to urchin2.0.UpdateUrchinByResourceIds failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("Resource to urchin2.0.UpdateUrchinByResourceIds success, batchID: %d", batchID)
	ctx.JSON(200, response.OuterSuccessWithData(map[string]interface{}{"batch_id": batchID}))
}

func RollbackBatch(ctx *context.APIContext) {
	batchID := ctx.QueryInt64("batch_id")

	err := subject_service.RollbackByBatchID(batchID)
	if err != nil {
		log.Error("Resource to urchin2.0.RollbackByBatchId failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("Resource to urchin2.0.RollbackByBatchId success, batchID: %d", batchID)
	ctx.JSON(200, response.OuterSuccess())
}

func CountResourceUpdatedData(ctx *context.APIContext) {

	countData, err := models.CountResourceUpdatedData()
	if err != nil {
		log.Error("Resource to urchin2.0.CountResourceUpdatedData failed, err: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	log.Info("Resource to urchin2.0.CountResourceUpdatedData success, countData: %v", countData)
	ctx.JSON(200, response.OuterSuccessWithData(countData))
}
