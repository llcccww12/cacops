package repo

import (
	"net/http"
	"strings"
	"sync"

	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	routeRepo "code.gitea.io/gitea/routers/repo"
)

var mutex *sync.Mutex = new(sync.Mutex)
var modelMutex *sync.Mutex = new(sync.Mutex)

func GetSuccessChunks(ctx *context.APIContext) {
	if errStr := checkDatasetPermission(ctx); errStr != "" {
		ctx.JSON(http.StatusForbidden, ctx.Tr(errStr))
	}
	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		ctx.JSON(200, map[string]interface{}{
			"result_code": -1,
			"msg":         "平台新版数据集已上线，请升级至 openi 3.0.0版本以使用最新功能。执行以下命令进行升级： pip install openi==3.0.0",
			"data":        nil,
		})
		return
	}

	routeRepo.GetSuccessChunks(ctx.Context)
}

func checkDatasetPermission(ctx *context.APIContext) string {
	datasetId := ctx.QueryInt64("dataset_id")

	dataset, err := models.GetDatasetByID(datasetId)
	if err != nil {
		log.Warn("can not find dataset", err)

		return "dataset.query_dataset_fail"
	}
	repo, err := models.GetRepositoryByID(dataset.RepoID)
	if err != nil {
		log.Warn("can not find repo", err)
		return "dataset.query_dataset_fail"
	}

	permission, err := models.GetUserRepoPermission(repo, ctx.User)
	if err != nil {
		log.Warn("can not find repo permission for user", err)
		return "dataset.query_dataset_fail"
	} else {
		log.Info("permission.AccessMode=" + string(permission.AccessMode))
	}
	if permission.AccessMode >= models.AccessModeAdmin {
		return ""
	}
	if !permission.CanWrite(models.UnitTypeDatasets) {
		return "error.no_right"
	}
	return ""
}

func GetAttachmentDir(ctx *context.APIContext) {
	routeRepo.GetDirSomeFiles(ctx.Context)
}

func GetAttachmentImageContent(ctx *context.APIContext) {
	routeRepo.GetImageContent(ctx.Context)
}

func GetAttachmentTxtContent(ctx *context.APIContext) {
	routeRepo.GetTxtContent(ctx.Context)
}

func NewMultipart(ctx *context.APIContext) {
	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		ctx.JSON(200, map[string]interface{}{
			"result_code": -1,
			"msg":         "平台新版数据集已上线，请升级至 openi 3.0.0版本以使用最新功能。执行以下命令进行升级： pip install openi==3.0.0",
			"data":        nil,
		})
		return
	}

	if errStr := checkDatasetPermission(ctx); errStr != "" {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         ctx.Tr(errStr),
		})
		return
	}
	ignore := false
	if setting.FLOW_CONTROL.IGNORE_FLAG != "" {
		//if ctx.Query("IGNORE_FLAG") == setting.FLOW_CONTROL.IGNORE_FLAG {
		if role.UserHasOper(ctx.User.ID, role.ROLE_OPER_IGNORE_FLOW_CONTROL) {
			ignore = true
		}
	}
	if !ignore {
		if err := routeRepo.CheckFlowForDatasetSDK(); err != nil {
			ctx.JSON(200, map[string]string{
				"result_code": "-1",
				"msg":         err.Error(),
			})
			return
		}
	}
	mutex.Lock()
	defer mutex.Unlock()
	datasetId := ctx.QueryInt64("dataset_id")
	fileName := ctx.Query("file_name")

	re, err := routeRepo.NewMultipartForApi(ctx.Context, !ignore)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         err.Error(),
		})
	} else {
		if !ignore {
			routeRepo.AddFileNameToCache(datasetId, fileName, ctx.User.ID)
		}
		re["result_code"] = "0"
		ctx.JSON(200, re)
	}
}
func GetMultipartUploadUrl(ctx *context.APIContext) {
	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		ctx.JSON(200, map[string]interface{}{
			"result_code": -1,
			"msg":         "平台新版数据集已上线，请升级至 openi 3.0.0版本以使用最新功能。执行以下命令进行升级： pip install openi==3.0.0",
			"data":        nil,
		})
		return
	}

	if errStr := checkDatasetPermission(ctx); errStr != "" {
		ctx.JSON(http.StatusForbidden, ctx.Tr(errStr))
	}
	routeRepo.GetMultipartUploadUrl(ctx.Context)
}

func CompleteMultipart(ctx *context.APIContext) {
	if errStr := checkDatasetPermission(ctx); errStr != "" {
		ctx.JSON(http.StatusForbidden, ctx.Tr(errStr))
	}
	datasetId := ctx.QueryInt64("dataset_id")
	fileName := ctx.Query("file_name")
	routeRepo.RemoveFileFromCache(datasetId, fileName, ctx.User.ID)
	routeRepo.CompleteMultipart(ctx.Context)

}
func GetAttachment(ctx *context.APIContext) {
	routeRepo.GetAttachment(ctx.Context)
}

func GetModelChunks(ctx *context.APIContext) {
	log.Info("GetModelChunks by api.")
	modeluuid := ctx.Query("modeluuid")
	model, err := models.QueryModelById(modeluuid)
	if err == nil {
		if errStr := checkModelPermission(ctx, model); errStr != "" {
			ctx.JSON(200, map[string]string{
				"result_code": "-1",
				"msg":         errStr,
			})
			return
		}
	} else {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         "model not exist.",
		})
		return
	}
	routeRepo.GetModelChunks(ctx.Context)
}

func NewModelMultipart(ctx *context.APIContext) {
	log.Info("NewModelMultipart by api.")
	modeluuid := ctx.Query("modeluuid")
	model, err := models.QueryModelById(modeluuid)
	if err == nil {
		if errStr := checkModelPermission(ctx, model); errStr != "" {
			ctx.JSON(200, map[string]string{
				"result_code": "-1",
				"msg":         errStr,
			})
			return
		}
	} else {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         "model not exist.",
		})
		return
	}
	ignore := false
	if setting.FLOW_CONTROL.IGNORE_FLAG != "" {
		//if ctx.Query("IGNORE_FLAG") == setting.FLOW_CONTROL.IGNORE_FLAG {
		if role.UserHasOper(ctx.User.ID, role.ROLE_OPER_IGNORE_FLOW_CONTROL) {
			ignore = true
		}
	}
	if !ignore {
		if err := routeRepo.CheckFlowForModelSDK(); err != nil {
			ctx.JSON(200, map[string]string{
				"result_code": "-1",
				"msg":         err.Error(),
			})
			return
		}
	}
	modelMutex.Lock()
	defer modelMutex.Unlock()
	fileName := ctx.Query("file_name")
	re, err := routeRepo.NewModelMultipartForApi(ctx.Context, !ignore)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         err.Error(),
		})
	} else {
		if !ignore {
			routeRepo.AddModelFileNameToCache(modeluuid, fileName, ctx.User.ID)
		}
		re["result_code"] = "0"
		ctx.JSON(200, re)
	}
}

func checkModelPermission(ctx *context.APIContext, model *models.AiModelManage) string {
	if ctx.User == nil {
		return "User not login."
	}
	if ctx.Repo.Repository == nil {
		repo, err := models.GetRepositoryByID(model.RepoId)
		if err == nil {
			ctx.Repo.Repository = repo
			owner, err := models.GetUserByID(repo.OwnerID)
			if err == nil {
				ctx.Repo.Owner = owner
			}
		} else {
			return "Repo is not exist."
		}
	}

	if ctx.User.IsAdmin || ctx.User.ID == model.UserId {
		return ""
	} else {
		return "User has not right to operate."
	}

}

func GetModelMultipartUploadUrl(ctx *context.APIContext) {
	log.Info("GetModelMultipartUploadUrl by api.")
	routeRepo.GetModelMultipartUploadUrl(ctx.Context)
}

func CompleteModelMultipart(ctx *context.APIContext) {
	log.Info("CompleteModelMultipart by api.")
	modeluuid := ctx.Query("modeluuid")
	//fileName := ctx.Query("file_name")
	uuid := ctx.Query("uuid")
	fileChunk, err := models.GetModelFileChunkByUUID(uuid)
	if err == nil {
		log.Info("fileChunk.ObjectName=" + fileChunk.ObjectName)
		objectNames := strings.Split(fileChunk.ObjectName, "/")
		routeRepo.RemoveModelFileFromCache(modeluuid, objectNames[len(objectNames)-1], ctx.User.ID)
	}
	routeRepo.CompleteModelMultipart(ctx.Context)
}
