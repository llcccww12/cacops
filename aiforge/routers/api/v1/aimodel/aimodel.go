package aimodel

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/markup/markdown"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/routers/api/v1/dataset"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
	"code.gitea.io/gitea/services/modelscope"
	"code.gitea.io/gitea/services/subject_service"
)

func CreateAimodel(ctx *context.Context, req entity.CreateAimodelReq) {
	owner := checkContextUser4Aimodel(ctx, req.OwnerId, req.AimodelType)
	if ctx.Written() {
		return
	}
	req.CreatorId = ctx.User.ID
	req.OwnerId = owner.ID
	res, err := subject_service.CreateAimodel(ctx.User, owner, req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}

	m := map[string]interface{}{
		"id":           res.ID,
		"name":         res.Name,
		"owner_name":   owner.Name,
		"aimodel_type": req.AimodelType,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// ExportTaskResult2NewAimodel 导出任务结果到新创建的线上模型
func ExportTaskResult2NewAimodel(ctx *context.Context, req entity.CreateAimodelReq) {
	// 固定为线上模型类型
	if req.AimodelType != models.MODEL_ONLINE_TYPE {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(
			response.NewBizError(errors.New("only online model type (0) is supported for export")), ctx.Locale))
		return
	}

	owner := checkContextUser4Aimodel(ctx, req.OwnerId, req.AimodelType)
	if ctx.Written() {
		return
	}
	req.CreatorId = ctx.User.ID
	req.OwnerId = owner.ID

	// 1. 创建新模型
	res, bizErr := subject_service.CreateAimodel(ctx.User, owner, req)
	if bizErr != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx.Locale))
		return
	}

	// 2. 触发导出任务
	processId, errExport := subject_service.ExportTaskResult2Aimodel(
		&entity.AimodelExportTaskResultReq{
			TaskId:      req.TaskId,
			FileListStr: req.FileListStr,
			Aimodel:     res,
			Doer:        ctx.User,
		})
	if errExport != nil {
		log.Error("ExportTaskResult2Aimodel failed, taskId=%d err=%v", req.TaskId, errExport)
		ctx.JSON(http.StatusOK, response.ResponseError(errExport))
		return
	}

	m := map[string]interface{}{
		"id":         res.ID,
		"name":       res.Name,
		"owner_name": owner.Name,
		"process_id": processId,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// GetExportAimodelByMsgId 获取模型导出进度
func GetExportAimodelByMsgId(ctx *context.Context) {
	progressId := ctx.Query("id")
	if progressId == "" {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}

	progress, err := redis_client.Get(progressId)
	if err != nil {
		log.Error("GetExportAimodelByMsgId failed, id=%s err=%v", progressId, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	msgMap := make(map[string]int, 0)
	err = json.Unmarshal([]byte(progress), &msgMap)
	if err != nil {
		log.Error("GetExportAimodelByMsgId Unmarshal failed, id=%s err=%v", progressId, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(msgMap))
}

func checkContextUser4Aimodel(ctx *context.Context, uid int64, aimodelType int) *models.User {
	// Not equal means current user is an organization.
	if uid == ctx.User.ID || uid == 0 {
		return ctx.User
	}

	org, err := models.GetUserByID(uid)
	if models.IsErrUserNotExist(err) {
		return ctx.User
	}

	if err != nil {
		ctx.ServerError("GetUserByID", fmt.Errorf("[%d]: %v", uid, err))
		return nil
	}

	// Check ownership of organization.
	if !org.IsOrganization() {
		ctx.Error(403)
		return nil
	}
	if org.Name == models.HF_MODEL_ORG_NAME && aimodelType == models.MODEL_HF_TYPE {
		return org
	}
	if !ctx.User.IsAdmin {
		canCreate, err := org.CanCreateOrgAimodel(ctx.User.ID)
		if err != nil {
			ctx.ServerError("CanCreateOrgRepo", err)
			return nil
		} else if !canCreate {
			ctx.Error(403)
			return nil
		}
	}
	return org
}

func GetAimodelCreationAvailableUsers(ctx *context.Context) {
	orgs, err := models.GetOrgsCanCreateAimodelByUserID(ctx.User.ID)
	if err != nil {
		log.Error("GetOrgsCanCreateAimodelByUserID failed, userID=%d err=%v", ctx.User.ID, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	users := make([]*models.User4Front, 0, len(orgs)+1)
	existsUser := make(map[int64]int, 0)
	users = append(users, ctx.User.ToFrontFormat())
	existsUser[ctx.User.ID] = 0
	for i := 0; i < len(orgs); i++ {
		org := orgs[i]
		if org == nil {
			continue
		}
		if _, ok := existsUser[org.ID]; !ok {
			users = append(users, org.ToFrontFormat())
			existsUser[org.ID] = 0
		}

	}

	//针对平台管理员在自身非组织成员时创建组织数据集的情况
	orgID := ctx.QueryInt64("org")
	if orgID > 0 {
		user := checkContextUser4Aimodel(ctx, orgID, 2)
		if user != nil {
			if _, ok := existsUser[user.ID]; !ok {
				users = append(users, user.ToFrontFormat())
			}
		}
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"users": users,
	}))
}

func GetAimodel(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel
	owner, err := models.GetUserByID(aimodel.OwnerID)
	if err != nil {
		log.Error("GetUserByID failed, ownerId=%d", aimodel.OwnerID)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	aimodel.Owner = owner

	var isCollected, canEditFile, canManage, canDownload, canDelete bool
	if ctx.User != nil {
		aimodelCollected := models.QueryModelCollectByUserId(aimodel.ID, ctx.User.ID)
		isCollected = false
		if len(aimodelCollected) > 0 {
			isCollected = true
		}
		canEditFile = ctx.AccessContext.IsWriter()
		canManage = ctx.AccessContext.IsAdmin()
		canDownload = ctx.AccessContext.IsReader()
		canDelete = ctx.AccessContext.IsOwner()
	}
	if !aimodel.IsPrivate {
		canDownload = true
	}
	aimodelPermission := &entity.AimodelPermissionInfo{
		CanDelete:   canDelete,
		CanEditFile: canEditFile,
		CanManage:   canManage,
		CanDownload: canDownload,
	}
	aimodel.IsCollected = isCollected
	aimodelInfo := entity.BuildAimodelInfo(aimodel)
	aimodelInfo.Permission = aimodelPermission

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(aimodelInfo))
}

func DelAimodel(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel
	err := subject_service.DeleteAimodel(ctx.User, aimodel)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func CollectAimodel(ctx *context.Context) {
	if ctx.User == nil {
		ctx.JSON(http.StatusOK, response.OuterBizError(response.BuildBizError(http.StatusUnauthorized, "user not login")))
		return
	}
	aimodel := ctx.AccessContext.Aimodel
	err := subject_service.CollectAimodel(ctx.User, aimodel, true)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func UnCollectAimodel(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel
	err := subject_service.CollectAimodel(ctx.User, aimodel, false)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func RecommendAimodel(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel
	err := subject_service.RecommendAimodel(ctx.User, aimodel)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func UnRecommendAimodel(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel
	err := subject_service.UnRecommendAimodel(ctx.User, aimodel)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

const README_FILE_NAME = "README.md"

func PutAimodelReadme(ctx *context.Context, req entity.AimodelReadmeReq) {
	aimodel := ctx.AccessContext.Aimodel
	req.Aimodel = aimodel
	err := subject_service.PutAimodelReadme(req)
	if err != nil {
		log.Error("PutAimodelReadme failed err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func GetAimodelReadme(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel
	if aimodel == nil {
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}

	// 存储 helper
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if storageHelper == nil {
		log.Error("GetAimodelReadme failed, StorageType=%d", aimodel.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}

	metas := map[string]string{"include_toc": "true"}
	path := path.Join(aimodel.Path, README_FILE_NAME)

	// 尝试打开 README 文件
	rc, err := storageHelper.OpenFile(path)
	if err != nil {
		log.Info("GetAimodelReadme open readme failed, path=%s err=%v", path, err)
		// fallback：从推荐 repo 拉取
		f := fetcher.NewGitLocalFetcher(setting.DyncConfigRepoOwner, setting.DyncConfigRepoName, setting.DyncConfigBranch)
		config, err := f.Fetch("model/" + README_FILE_NAME)
		if err != nil {
			log.Error("GetAimodelReadme Fetch failed, path=%s err=%v", path, err)
			ctx.JSON(http.StatusOK, response.ResponseError(err))
			return
		}

		ctx.JSON(http.StatusOK, response.OuterSuccessWithData(&entity.DatasetReadmeResponse{
			Content:       fmt.Sprint(config.Value),
			HtmlContent:   string(markdown.RenderRaw([]byte(fmt.Sprint(config.Value)), "", false, metas)),
			FileName:      README_FILE_NAME,
			IsExistMDFile: false,
		}))
		return
	}
	defer rc.Close()

	// 读内容
	content, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Error("ReadAll failed, path=%s err=%v", path, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(&entity.DatasetReadmeResponse{
		Content:       string(content),
		HtmlContent:   string(markdown.RenderRaw(content, "", false, metas)),
		FileName:      README_FILE_NAME,
		IsExistMDFile: true,
	}))
}

func GetAimodelFileList(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	marker := ctx.Query("marker")
	pageSize := ctx.QueryInt("page_size")
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 20
	}
	aimodel := ctx.AccessContext.Aimodel
	objectKey := path.Join(aimodel.Path, parentDir)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", aimodel.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}
	res, err := helper.GetOneLevelObjectsUnderDirWithMarker(objectKey, marker, pageSize)
	if err != nil {
		log.Error("GetOneLevelObjectsUnderDir failed, objectKey=%s err=%v", objectKey, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	fileList := res.Objects
	if fileList == nil {
		ctx.JSON(200, map[string]interface{}{
			"file_list": []storage.FileInfo{},
			"has_next":  false,
			"marker":    "",
		})
		return
	}

	for i := 0; i < len(res.Objects); i++ {
		fileList[i].ParenDir = parentDir
		fileList[i].RelativePath = ""
		if fileList[i].IsDir {
			continue
		}
		fileList[i].IsSupportPrivew = dataset.IsSupportedPreview(fileList[i].FileName)
	}
	m := map[string]interface{}{
		"file_list": fileList,
		"has_next":  res.IsTruncated,
		"marker":    res.NextMarker,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func DownloadAimodelFile(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	aimodel := ctx.AccessContext.Aimodel
	objectKey := path.Join(aimodel.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", aimodel.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}

	url, err := helper.GetSignedDownloadUrl(objectKey)
	if err != nil {
		log.Error("GetSignedDownloadUrl failed, objectKey=%s err=%v", objectKey, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	if ctx.Req.Method == "GET" {
		models.ModifyModelDownloadCount(aimodel.ID)
	}
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
}

func DownloadAimodel(ctx *context.Context) {
	aimodel := ctx.AccessContext.Aimodel

	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, dataType=%s", aimodel.StorageType)
		ctx.JSON(http.StatusOK, response.ResponseError(errors.New("Storage type error")))
		return
	}
	resultFileName := aimodel.Name + ".zip"
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(resultFileName))
	ctx.Resp.Header().Set("Content-Type", "application/zip")
	ctx.Resp.Header().Set("X-Content-Type-Options", "nosniff")

	zipWriter := zip.NewWriter(ctx.Resp)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			log.Error("failed to close zip writer: %v", err)
		}
	}()

	fileList := make([]storage.FileInfo, 0, 1000)
	marker := ""
	index := 0
	for {
		res, err := helper.GetAllObjectsUnderDirWithMarker(aimodel.Path, marker, 1000)
		if err != nil {
			log.Error("GetAllObjectsUnderDir err.objectKeyPrefix=%s,err=%v", aimodel.Path, err)
			ctx.JSON(http.StatusOK, response.ResponseError(err))
			return
		}
		fileList = res.Objects
		if len(fileList) == 0 {
			if index == 0 {
				ctx.JSON(http.StatusOK, response.ResponseError(errors.New("No files found")))
				return
			}
			break
		}

		marker = res.NextMarker

		buf := make([]byte, 32*1024) // 32KB buffer

		for _, file := range fileList {
			if file.IsDir {
				continue
			}

			if err := dataset.AddFileToZip(helper, file, zipWriter, buf); err != nil {
				// 一旦开始写入zip，就不能更改HTTP状态码
				log.Error("failed to add file to zip: %v", err)
				return
			}
		}

		// 显式刷新以确保所有数据写入
		if err := zipWriter.Flush(); err != nil {
			log.Error("failed to flush zip writer: %v", err)
		}
		if !res.IsTruncated {
			break
		}
	}

	models.ModifyModelDownloadCount(aimodel.ID)

}

func GetAimodelFileMeta(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	aimodel := ctx.AccessContext.Aimodel
	objectKey := path.Join(aimodel.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", aimodel.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}

	res, err := helper.GetObjectMeta(objectKey)
	if err != nil {
		log.Error("GetAimodelFileMeta failed, objectKey=%s err=%v", objectKey, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(200, response.OuterSuccessWithData(res))
}

func DeleteAimodelFile(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	aimodel := ctx.AccessContext.Aimodel

	err := subject_service.DeleteAimodelFile(aimodel, parentDir, fileName)
	if err != nil {
		log.Error("DeleteAimodelFile failed, aimodel.ID=%s parentDir=%s fileName=%s err=%v", aimodel.ID, parentDir, fileName, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func BatchDeleteAimodel(ctx *context.Context) {
	aimodelIdStr := ctx.Query("aimodel_ids")
	dataIds := strings.Split(aimodelIdStr, ",")
	successIds := subject_service.BatchDeleteAimodel(ctx.User, dataIds)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"success": successIds,
	}))
}

func EditAimodel(ctx *context.Context, req entity.CreateAimodelReq) {
	err := subject_service.EditAimodel(req, ctx.AccessContext.Aimodel)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func parseSearchAimodelReq(ctx *context.APIContext) models.SearchAimodelReq {
	creatorName := ctx.Query("creator_name")
	ownerName := ctx.Query("owner_name")
	keyword := ctx.Query("q")
	var aimodelName string
	if keyword != "" {
		//如果搜索关键词的格式是 a/b，则将a解析成ownerName，b解析成aimodelName
		parts := strings.Split(keyword, "/")
		if len(parts) == 2 {
			v1 := strings.TrimSpace(parts[0])
			v2 := strings.TrimSpace(parts[1])
			if v1 != "" && v2 != "" {
				keyword = ""
				if ownerName == "" {
					//只有当搜索条件中没有owner_name时才取关键词中的ownerName
					ownerName = v1
				}
				aimodelName = v2
			}
		}
	}
	orderBy := ctx.Query("order_by")
	recommend := ctx.Query("recommend")
	label := ctx.Query("label")
	engineStr := ctx.Query("engine")
	engine := -1
	if engineStr != "" {
		engine, _ = strconv.Atoi(engineStr)
	}
	aimodelTypeStr := ctx.Query("aimodel_type")
	aimodelType := models.MODEL_ALL_TYPE
	if aimodelTypeStr != "" {
		aimodelType, _ = strconv.Atoi(aimodelTypeStr)

	}
	ownerType := ctx.Query("owner_type")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("page_size")
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 12
	}
	req := models.SearchAimodelReq{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Keyword:      keyword,
		LabelFilter:  label,
		EngineFilter: engine,
		Recommend:    recommend,
		OrderBy:      orderBy,
		OwnerName:    ownerName,
		AimodelName:  aimodelName,
		AimodelType:  aimodelType,
		OwnerType:    ownerType,
		CreatorName:  creatorName,
	}
	req.User = ctx.User
	return req
}

func SearchAllPublicAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityPublic
	req.Scope = models.ScopeAll
	req.From = "profile"
	res, total, err := subject_service.SearchAimodel(req)
	if err != nil {
		log.Error("SearchAllPublicAimodel failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	if total == 0 && ctx.Query("source") != "local" && req.AimodelType != 2 {
		if msRes, msTotal, msErr := modelscope.ListModels(req.Keyword, req.OrderBy, req.Page, req.PageSize); msErr == nil && msTotal > 0 {
			res = msRes
			total = msTotal
		}
	}
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func SearchMyOwnedAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeOwned
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyOwnedAimodel failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func SearchMyCollaboratedAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeCollaborated
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyCollaboratedAimodel failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))

}

func SearchMyCollectedAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeCollected
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyCollectedAimodel failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func SearchMyAccessibleAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAccessible
	req.From = "profile"
	if ctx.User != nil && ctx.User.IsAdmin {
		req.UseAdminPermission = true
	}
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyAccessibleAimodel failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// for model_type=2: external model
func SearchInvolvedAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeInvolved
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyInvolvedAimodels failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func SearchOrgAimodels(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAccessible
	req.OwnerId = ctx.Org.Organization.ID
	if ctx.User != nil && ctx.User.IsAdmin {
		req.UseAdminPermission = true
	}
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchOrgAimodels failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 存储配额-数据集
func SearchStorage4Aimodels(ctx *context.APIContext) {
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
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAll
	req.OwnerId = ctx.Org.Organization.ID
	res, total, bizErr := subject_service.SearchAimodel(req)
	if bizErr != nil {
		log.Error("SearchOrgDatasets failed, req=%+v err=%v", req, bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx.Locale))
		return
	}
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	for _, aimodel := range m.Aimodels {
		aimodel.Permission = &entity.AimodelPermissionInfo{
			CanManage:   true,
			CanEditFile: true,
			CanDownload: true,
			CanDelete:   true,
		}
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 管理后台-数据集列表
func SearchAllAimodelsForAdmin(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	visibilityFilter := ctx.Query("visibility")
	switch visibilityFilter {
	case "public":
		req.Visibility = models.VisibilityPublic
	case "private":
		req.Visibility = models.VisibilityPrivate
	default:
		req.Visibility = models.VisibilityAll
	}
	req.Scope = models.ScopeAll
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchAllAimodelsForAdmin failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 存储管理-个人数据集
func SearchAimodelsForPersonalStorage(ctx *context.APIContext) {
	req := parseSearchAimodelReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAccessible
	req.OwnerId = ctx.User.ID
	res, total, err := subject_service.SearchAimodel(req)
	m := entity.SearchAimodelRes{
		Aimodels: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchAimodelsForPersonalStorage failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func AimodelEvolutionGraph(ctx *context.APIContext) {
	aimodel := ctx.AccessContext.Aimodel
	doer := ctx.User

	aimodelGraph, err := BuildAimodelEvolutionGraph(doer, aimodel)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(aimodelGraph))
}

func GetOrgAimodelLabels(ctx *context.APIContext) {
	GetOwnedPublicAimodelLabels(ctx, ctx.Org.Organization.ID)
}

func GetUserAimodelLabels(ctx *context.APIContext) {
	GetOwnedPublicAimodelLabels(ctx, ctx.User.ID)
}

func GetOwnedPublicAimodelLabels(ctx *context.APIContext, userId int64) {
	res, err := subject_service.GetOwnedPublicAimodelLabels(userId)
	if err != nil {
		log.Error("GetOwnedPublicAimodelLabels failed, userId=%d err=%v", ctx.Org.Organization.ID, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}

func GetAimodelRelatedData(ctx *context.APIContext) {
	aimodel := ctx.AccessContext.Aimodel
	doer := ctx.User
	res, err := subject_service.GetAimodelRelatedData(aimodel, doer)
	if err != nil {
		log.Error("GetAimodelExtraInfo failed, aimodelId=%d err=%v", aimodel.ID, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}
