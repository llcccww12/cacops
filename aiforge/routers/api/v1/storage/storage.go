package storage

import (
	"net/http"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/role"
	"code.gitea.io/gitea/services/storage_limit"
)

func GetUserStorageSummary(ctx *context.APIContext) {
	userId := ctx.User.ID
	subjectID := ctx.Query("subject_id")
	subjectType := ctx.QueryInt("subject_type")
	if subjectID != "" && subjectType >= 0 {
		subjectContext, err := models.GetSubjectContext(subjectID, models.SubjectType(subjectType))
		if err != nil {
			ctx.JSON(http.StatusOK, entity.StorageSummary{
				StorageLimit:       0,
				DatasetUsedStorage: 0,
				ModelUsedStorage:   0,
				UsedStorage:        0,
				RemainingStorage:   0,
			})
			return
		}
		err = subjectContext.GetOwner()
		if err != nil {
			ctx.JSON(http.StatusOK, entity.StorageSummary{
				StorageLimit:       0,
				DatasetUsedStorage: 0,
				ModelUsedStorage:   0,
				UsedStorage:        0,
				RemainingStorage:   0,
			})
			return
		}
		userId = subjectContext.OwnerID
	}

	storageLimit := role.UserStorageNum(userId)

	usedDatasetSize, err := models.GetUsedDatasetRegistrySizeByUser(userId)
	if err != nil {
		log.Warn("GetUsedDatasetSizeByUser", err)
		ctx.JSON(http.StatusOK, entity.StorageSummary{
			StorageLimit:       storageLimit * 1024 * 1024 * 1024,
			DatasetUsedStorage: usedDatasetSize,
			ModelUsedStorage:   0,
			UsedStorage:        usedDatasetSize,
			RemainingStorage:   0,
		})
		return
	}
	usedModelSize, err := models.GetUsedModelSizeByOwner(userId)
	if err != nil {
		log.Warn("GetUsedModelSizeByOwner", err)
		ctx.JSON(http.StatusOK, entity.StorageSummary{
			StorageLimit:       storageLimit * 1024 * 1024 * 1024,
			DatasetUsedStorage: usedDatasetSize,
			ModelUsedStorage:   usedModelSize,
			UsedStorage:        usedDatasetSize + usedModelSize,
			RemainingStorage:   0,
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.StorageSummary{
		StorageLimit:       storageLimit * 1024 * 1024 * 1024,
		DatasetUsedStorage: usedDatasetSize,
		ModelUsedStorage:   usedModelSize,
		UsedStorage:        usedDatasetSize + usedModelSize,
		RemainingStorage:   storageLimit*1024*1024*1024 - usedDatasetSize - usedModelSize,
	})

}

func GetOrgStorageSummary(ctx *context.APIContext) {
	org := ctx.Org.Organization
	isOwner, err := org.IsOwnedBy(ctx.User.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}
	if !isOwner && !ctx.User.IsAdmin {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}
	summary, err := storage_limit.GetUserStorageSummary(org)
	if err != nil {
		log.Error("GetOrgStorageSummary err.org.name=%s org.ID=%d  err=%v", org.Name, org.ID, err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(summary))
}

func SearchDataSetAttachments(ctx *context.APIContext) {
	opts := &models.SearchDatasetAttachmentOptions{
		ListOptions: models.ListOptions{
			Page:     ctx.QueryInt("page"),
			PageSize: ctx.QueryInt("pageSize"),
		},

		UploaderID: ctx.User.ID,
		OrderBy:    ctx.Query("orderBy"),
	}

	attachments, total, err := storage_limit.SearchDatasetAttachment(opts)
	if err != nil {
		ctx.ServerError("SearchDatasetAttachment", err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"total": total,
		"data":  attachments,
	})

}
