package repo

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	repo_service "code.gitea.io/gitea/services/repository"
	"net/http"
)

// PinnedProject 项目置顶
func PinnedProject(ctx *context.APIContext) {
	// 现在的系统时间
	nowUnix := timeutil.TimeStampNow()

	// Step 1 查看数据是否已经存在
	record, err := models.GetPinnedRecord(ctx.User.ID, ctx.Repo.Repository.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 2, Message: "internal error"})
		return
	}

	// Step 2 查看数据是否已经存在,存在直接更新
	if record != nil && record.ID > 0 {
		if err := models.UpdatePinnedRecord(&models.RepositoryPinnedRecord{
			UserID:       ctx.User.ID,
			RepositoryID: ctx.Repo.Repository.ID,
			PinnedUnix:   nowUnix,
		}); err != nil {
			ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 2, Message: "internal error"})
			return
		}
	} else {
		// Step 3 创建置顶
		if err = models.CreatePinnedRecord(&models.RepositoryPinnedRecord{
			UserID:       ctx.User.ID,
			RepositoryID: ctx.Repo.Repository.ID,
			PinnedUnix:   nowUnix,
			CreatedUnix:  nowUnix,
		}); err != nil {
			ctx.Error(http.StatusInternalServerError, "PinnedProject", "create failed")
			return
		}
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
	return
}

// CancelPinnedProject 取消项目置顶
func CancelPinnedProject(ctx *context.APIContext) {
	// Step 1 取消置顶
	if err := models.CancelPinnedRecord(&models.RepositoryPinnedRecord{
		UserID:       ctx.User.ID,
		RepositoryID: ctx.Repo.Repository.ID,
		CancelUnix:   timeutil.TimeStampNow(),
	}); err != nil {
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 2, Message: "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
	return
}

// DeleteProject 删除项目
func DeleteProject(ctx *context.APIContext) {

	// Step 1 删除用户项目的耦合关系
	// 参考代码/:username/:reponame/settings方法
	repo, err := models.GetRepositoryByID(ctx.Repo.Repository.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 2, Message: "internal error"})
		return
	}

	// 检查用户是否有权限删除
	canDelete, err := repo.CanUserDelete(ctx.User)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 2, Message: "internal error"})
		return
	}
	if !canDelete {
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 2, Message: "is not owner"})
		return
	}

	// 删除对应的项目依赖
	if err := repo_service.DeleteRepository(ctx.User, repo); err != nil {
		ctx.ServerError("DeleteRepository", err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
	return
}
