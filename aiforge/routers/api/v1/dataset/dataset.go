package dataset

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/markup/markdown"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
	"code.gitea.io/gitea/services/modelscope"
	"code.gitea.io/gitea/services/subject_service"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func CreateDataset(ctx *context.Context, req entity.CreateDatasetReq) {
	owner := checkContextUser4Dataset(ctx, req.OwnerId)
	if ctx.Written() {
		return
	}
	req.CreatorId = ctx.User.ID
	res, err := subject_service.CreateDataset(ctx.User, owner, req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"id":         res.ID,
		"name":       res.Name,
		"owner_name": owner.Name,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func checkContextUser4Dataset(ctx *context.Context, uid int64) *models.User {
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
	if !ctx.User.IsAdmin {
		canCreate, err := org.CanCreateOrgDataset(ctx.User.ID)
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

func GetDataset(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	owner, err := models.GetUserByID(dataset.OwnerID)
	if err != nil {
		log.Error("GetUserByID failed, ownerId=%d", dataset.OwnerID)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	var isCollected, canEditFile, canManage, canDownload, canDelete bool
	if ctx.User != nil {
		isCollected = models.IsDatasetCollecting(ctx.User.ID, dataset.ID)
		canEditFile = ctx.AccessContext.IsWriter()
		canManage = ctx.AccessContext.IsAdmin()
		canDownload = ctx.AccessContext.IsReader()
		canDelete = ctx.AccessContext.IsOwner()
	}
	if !dataset.IsPrivate {
		canDownload = true
	}
	ownerFront := owner.ToFrontFormat()
	ownerFront.Email = ""

	datasetInfo := &entity.DatasetInfo{
		ID:            dataset.ID,
		Name:          dataset.Name,
		Alias:         dataset.Alias,
		Tags:          []string(dataset.Tags),
		License:       dataset.License,
		Tasks:         []string(dataset.Tasks),
		IsPrivate:     dataset.IsPrivate,
		UseCount:      dataset.UseCount,
		DownloadCount: dataset.DownloadCount,
		NumStars:      dataset.NumCollections,
		CreatedUnix:   dataset.CreatedUnix,
		UpdatedUnix:   dataset.UpdatedUnix,
		CanEditFile:   canEditFile,
		CanManage:     canManage,
		CanDownload:   canDownload,
		IsCollected:   isCollected,
		OwnerName:     owner.Name,
		Owner:         ownerFront,
		Size:          dataset.Size,
		CanDelete:     canDelete,
		Recommend:     dataset.Recommend,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(datasetInfo))
}

const README_FILE_NAME = "README.md"

func PutDatasetReadme(ctx *context.Context, req entity.DatasetReadmeReq) {
	dataset := ctx.AccessContext.Dataset
	req.Dataset = dataset
	err := subject_service.PutDatasetReadme(req)
	if err != nil {
		log.Error("PutDatasetReadme failed err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func GetDatasetReadme(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if storageHelper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", dataset.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}
	metas := map[string]string{"include_toc": "true"}
	path := path.Join(dataset.Path, README_FILE_NAME)
	rc, err := storageHelper.OpenFile(path)
	if err != nil {
		log.Info("GetDatasetReadme open readme failed, path=%s err=%v", path, err)
		f := fetcher.NewGitLocalFetcher(setting.DyncConfigRepoOwner, setting.DyncConfigRepoName, setting.DyncConfigBranch)
		config, err := f.Fetch("dataset/" + README_FILE_NAME)
		if err != nil {
			log.Error("GetDatasetReadme Fetch failed, path=%s err=%v", path, err)
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

func GetDatasetFileList(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	marker := ctx.Query("marker")
	pageSize := ctx.QueryInt("page_size")
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 20
	}
	dataset := ctx.AccessContext.Dataset
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", dataset.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}

	res, err := ListFiles(dataset.Path, parentDir, marker, helper, pageSize)
	if err != nil {
		log.Error("ListFiles failed, dataset.id=%d name=%s", dataset.ID, dataset.Name)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}

	if res.Objects == nil {
		ctx.JSON(200, map[string]interface{}{
			"file_list": []storage.FileInfo{},
			"has_next":  false,
			"marker":    "",
		})
		return
	}
	fileList := res.Objects

	for i := 0; i < len(res.Objects); i++ {
		fileList[i].ParenDir = parentDir
		fileList[i].RelativePath = ""
		if fileList[i].IsDir {
			continue
		}
		fileList[i].IsSupportPrivew = IsSupportedPreview(fileList[i].FileName)
	}
	m := map[string]interface{}{
		"file_list": fileList,
		"has_next":  res.IsTruncated,
		"marker":    res.NextMarker,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

var priorityFiles = []string{
	"README.md",
	"openi_resource.version",
}

func ListFiles(prefix, parentDir, marker string, helper storage_helper.StorageHelper, pageSize int) (*storage_helper.ObjectListResponse, error) {
	if parentDir == "" && marker == "" {
		return listFirstPage(prefix, helper, pageSize)
	}
	return listFollowPage(prefix, parentDir, marker, helper, pageSize)
}

func listFollowPage(prefix, parentDir, marker string, helper storage_helper.StorageHelper, pageSize int) (*storage_helper.ObjectListResponse, error) {
	ojectKey := path.Join(prefix, parentDir)

	if parentDir != "" {
		res, err := helper.GetOneLevelObjectsUnderDirWithMarker(ojectKey, marker, pageSize)
		if err != nil {
			log.Error("GetOneLevelObjectsUnderDir failed, objectKey=%s err=%v", prefix, err)
			return nil, err
		}
		return &storage_helper.ObjectListResponse{
			NextMarker:  res.NextMarker,
			Objects:     res.Objects,
			IsTruncated: res.IsTruncated,
		}, nil
	}

	//根目录需要过滤优先文件
	fileList := make([]storage.FileInfo, 0, pageSize)
	var nextMarker = marker
	var truncated = false

	remain := pageSize
	for {
		if remain <= 0 {
			break
		}
		res, err := helper.GetOneLevelObjectsUnderDirWithMarker(ojectKey, nextMarker, remain)
		if err != nil {
			log.Error("GetOneLevelObjectsUnderDir failed, objectKey=%s err=%v", prefix, err)
			return nil, err
		}
		for i := 0; i < len(res.Objects); i++ {
			obj := res.Objects[i]
			var isPriority = false
			for _, val := range priorityFiles {
				if val == obj.FileName {
					isPriority = true
					break
				}
			}
			if isPriority {
				continue
			}
			fileList = append(fileList, obj)
		}
		if len(fileList) == pageSize || !res.IsTruncated {
			nextMarker = res.NextMarker
			truncated = res.IsTruncated
			break
		}
		remain = pageSize - len(fileList)
		nextMarker = res.NextMarker
	}
	return &storage_helper.ObjectListResponse{
		NextMarker:  nextMarker,
		Objects:     fileList,
		IsTruncated: truncated,
	}, nil

}

func listFirstPage(prefix string, helper storage_helper.StorageHelper, pageSize int) (*storage_helper.ObjectListResponse, error) {
	fileList := make([]storage.FileInfo, 0, pageSize)
	//第一页优先显示README.md、openi_resource.version
	for i := 0; i < len(priorityFiles); i++ {
		fileName := priorityFiles[i]
		key := path.Join(prefix, fileName)
		obj, _ := helper.GetObjectMeta(key)
		if obj != nil {
			fileList = append(fileList, storage.FileInfo{
				FileName: fileName,
				ModTime:  obj.LastModified.Local().Format("2006-01-02 15:04:05"),
				IsDir:    false,
				Size:     obj.ContentLength,
				ParenDir: "",
			})
		}
	}
	var nextMarker string
	var truncated = false
	remain := pageSize - len(fileList)
	for {
		if remain <= 0 {
			break
		}
		res, err := helper.GetOneLevelObjectsUnderDirWithMarker(prefix, nextMarker, remain)
		if err != nil {
			log.Error("GetOneLevelObjectsUnderDir failed, objectKey=%s err=%v", prefix, err)
			return nil, err
		}
		for i := 0; i < len(res.Objects); i++ {
			obj := res.Objects[i]
			var isPriority = false
			for _, val := range priorityFiles {
				if val == obj.FileName {
					isPriority = true
					break
				}
			}
			if isPriority {
				continue
			}
			fileList = append(fileList, obj)
		}
		if len(fileList) == pageSize || !res.IsTruncated {
			nextMarker = res.NextMarker
			truncated = res.IsTruncated
			break
		}
		remain = pageSize - len(fileList)
		nextMarker = res.NextMarker
	}
	return &storage_helper.ObjectListResponse{
		NextMarker:  nextMarker,
		Objects:     fileList,
		IsTruncated: truncated,
	}, nil

}

var (
	supportPreviewExts = []string{".txt", ".xml", ".html", ".json", ".py", ".sh", ".md", ".csv", ".log", ".js", ".css", ".ipynb", ".jpg", ".jpeg", ".png", ".gif", ".bmp"}
)

func IsSupportedPreview(filename string) bool {
	lowerName := strings.ToLower(filename)
	for _, ext := range supportPreviewExts {
		if strings.HasSuffix(lowerName, ext) {
			return true
		}
	}
	return false
}

func DownloadDatasetFile(ctx *context.Context) {

	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	dataset := ctx.AccessContext.Dataset
	objectKey := path.Join(dataset.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", dataset.StorageType)
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
		models.IncreaseDatasetRegistryDownloadCount(dataset.ID)
	}
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
}

func GetDatasetFileMeta(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	dataset := ctx.AccessContext.Dataset
	objectKey := path.Join(dataset.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, StorageType=%d", dataset.StorageType)
		ctx.JSON(http.StatusOK, response.OuterBizError(response.SYSTEM_ERROR))
		return
	}

	res, err := helper.GetObjectMeta(objectKey)
	if err != nil {
		log.Error("GetDatasetFileMeta failed, objectKey=%s err=%v", objectKey, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(200, response.OuterSuccessWithData(res))
}

// 所有公开
func SearchAllPublicDataset(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityPublic
	req.Scope = models.ScopeAll
	req.From = "profile"
	res, total, err := subject_service.SearchDatasets(req)
	if err != nil {
		log.Error("SearchAllPublicDataset failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	if total == 0 && ctx.Query("source") != "local" {
		if msRes, msTotal, msErr := modelscope.ListDatasets(req.Keyword, req.OrderBy, req.Page, req.PageSize); msErr == nil && msTotal > 0 {
			res = msRes
			total = msTotal
		}
	}
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 我拥有的
func SearchMyOwnedDatasets(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeOwned
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyOwnedDatasets failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 我协作的
func SearchMyCollaboratedDatasets(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeCollaborated
	req.MinAccessMode = ctx.Query("min_access_mode")
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyCollaboratedDatasets failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 我收藏的
func SearchMyCollectedDatasets(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeCollected
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchMyCollectedDatasets failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 新建云脑任务-所有
func SearchMyAccessibleDatasets(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAccessible
	req.From = "profile"
	if ctx.User != nil && ctx.User.IsAdmin {
		req.UseAdminPermission = true
	}
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchAllAccessibleDatasets failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 组织主页-数据集
func SearchOrgDatasets(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAccessible
	req.OwnerId = ctx.Org.Organization.ID
	if ctx.User != nil && ctx.User.IsAdmin {
		req.UseAdminPermission = true
	}
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchOrgDatasets failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 存储配额-数据集
func SearchStorage4Datasets(ctx *context.APIContext) {
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
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAll
	req.OwnerId = ctx.Org.Organization.ID
	res, total, bizErr := subject_service.SearchDatasets(req)
	if bizErr != nil {
		log.Error("SearchOrgDatasets failed, req=%+v err=%v", req, bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx.Locale))
		return
	}
	//只有组织owner才能访问这个接口，所以所有权限都有
	for i := 0; i < len(res); i++ {
		res[i].CanEditFile = true
		res[i].CanManage = true
		res[i].CanDownload = true
		res[i].CanDelete = true
	}
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 管理后台-数据集列表
func SearchAllDatasetsForAdmin(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
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
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchAllDatasetsForAdmin failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 存储管理-个人数据集
func SearchDatasetsForPersonalStorage(ctx *context.APIContext) {
	req := parseSearchDatasetReq(ctx)
	req.Visibility = models.VisibilityAll
	req.Scope = models.ScopeAccessible
	req.OwnerId = ctx.User.ID
	res, total, err := subject_service.SearchDatasets(req)
	m := entity.SearchDatasetRes{
		Datasets: res,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if err != nil {
		log.Error("SearchOrgDatasets failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func parseSearchDatasetReq(ctx *context.APIContext) models.SearchDatasetReq {
	creatorName := ctx.Query("creator_name")
	ownerName := ctx.Query("owner_name")
	keyword := ctx.Query("q")
	var datasetName string
	if keyword != "" {
		//如果搜索关键词的格式是 a/b，则将a解析成ownerName，b解析成datasetName
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
				datasetName = v2
			}
		}
	}
	tasksStr := ctx.Query("tasks")
	tasks := make([]string, 0)
	if tasksStr != "" {
		tasks = strings.Split(tasksStr, "|")
	}
	tagsStr := ctx.Query("tags")
	tags := make([]string, 0)
	if tagsStr != "" {
		tags = strings.Split(tagsStr, "|")
	}
	liscense := ctx.Query("license")
	recommend := ctx.Query("recommend")
	orderBy := ctx.Query("order_by")
	page := ctx.QueryInt("page")
	ownerType := ctx.Query("owner_type")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("page_size")
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 12
	}
	req := models.SearchDatasetReq{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Keyword:     keyword,
		Tasks:       tasks,
		Tags:        tags,
		Liscense:    liscense,
		Recommend:   recommend,
		OrderBy:     orderBy,
		OwnerName:   ownerName,
		DatasetName: datasetName,
		OwnerType:   ownerType,
		CreatorName: creatorName,
	}
	req.User = ctx.User
	return req
}

func DeleteDatasetFile(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	dataset := ctx.AccessContext.Dataset

	err := subject_service.DeleteDatasetFile(dataset, parentDir, fileName)
	if err != nil {
		log.Error("DeleteFile failed, dataset.ID=%s parentDir=%s fileName=%s err=%v", dataset.ID, parentDir, fileName, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func PreviewDatasetFile(ctx *context.Context) {
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	if fileName == "" {
		log.Error("PreviewDatasetFile failed, fileName is empty")
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}
	if !IsSupportedPreview(fileName) {
		log.Error("PreviewDatasetFile failed, fileName=%s is not supported for preview", fileName)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.UNSUPPORTED_PREVIEW_FILE_TYPE, ctx.Locale))
		return
	}
	dataset := ctx.AccessContext.Dataset
	objectKey := path.Join(dataset.Path, parentDir, fileName)
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", dataset.StorageType)
		ctx.JSON(http.StatusOK, response.ServerError("Storage type error"))
		return
	}
	obj, err := helper.GetObject(objectKey)
	if err != nil {
		log.Error("GetObject failed, objectKey=%s err=%v", objectKey, err)
		ctx.JSON(http.StatusOK, response.ServerError("GetObject failed"))
		return
	}
	defer obj.Body.Close()

	if obj.ContentLength > setting.MAX_PREVIEW_FILE_SIZE {
		log.Error("PreviewDatasetFile failed, file too large, objectKey=%s, size=%d", objectKey, obj.ContentLength)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.OVER_PREVIEW_SIZE, ctx.Locale))
		return
	}

	contentType := determineContentType(objectKey, obj.ContentType)

	rc := WrapWithUTF8(contentType, "", obj.Body)
	defer rc.Close()

	ctx.Resp.Header().Set("Content-Type", contentType)
	_, err = io.Copy(ctx.Resp, rc)
	if err != nil {
		log.Error("io.Copy failed, objectKey=%s err=%v", objectKey, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
}

type readerWithCloser struct {
	io.Reader
	io.Closer
}

func WrapWithUTF8(contentType, forcedCharset string, r io.ReadCloser) io.ReadCloser {
	if !needEncodingCheck(contentType) {
		return r
	}

	br := bufio.NewReader(r)

	// 用户指定编码优先
	var charset string
	if forcedCharset != "" {
		charset = forcedCharset
	} else {
		head, _ := br.Peek(4096) // 大文件也只取前 4KB 检测
		charset = detectEncoding(head)
	}

	if charset == "UTF-8" {
		// UTF-8 直接返回原 reader（复用 br）
		return struct {
			io.Reader
			io.Closer
		}{br, r}
	}

	enc := getEncodingByName(charset)
	if enc == nil {
		// 不支持的编码，原样返回
		return struct {
			io.Reader
			io.Closer
		}{br, r}
	}

	// 转 UTF-8（流式）
	utf8Reader := transform.NewReader(br, enc.NewDecoder())
	return struct {
		io.Reader
		io.Closer
	}{utf8Reader, r}
}

// 判断是否需要关注编码
func needEncodingCheck(contentType string) bool {
	if strings.HasPrefix(contentType, "text/") ||
		contentType == "application/json" ||
		contentType == "application/xml" ||
		contentType == "application/csv" {
		return true
	}
	return false
}

// 基础编码检测：BOM > UTF-8 > GBK 识别 > 默认 GBK
func detectEncoding(buf []byte) string {
	// BOM 检测
	if bytes.HasPrefix(buf, []byte{0xEF, 0xBB, 0xBF}) {
		return "UTF-8"
	}
	if bytes.HasPrefix(buf, []byte{0xFF, 0xFE}) {
		return "UTF-16LE"
	}
	if bytes.HasPrefix(buf, []byte{0xFE, 0xFF}) {
		return "UTF-16BE"
	}

	// UTF-8 校验
	if utf8.Valid(buf) {
		return "UTF-8"
	}

	// 3. 尝试 GBK 解码
	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(buf), decoder)
	decoded, err := ioutil.ReadAll(reader)
	if err == nil && isMostlyReadable(string(decoded)) {
		return "GBK"
	}

	return "UTF-8"
}

func isMostlyReadable(s string) bool {
	if len(s) == 0 {
		return false
	}

	readable := 0
	chinese := 0
	replacement := 0
	total := 0

	for _, r := range s {
		total++

		if r == '\n' || r == '\r' || r == '\t' {
			readable++
			continue
		}

		if r == '\uFFFD' {
			replacement++
			continue
		}

		if unicode.IsPrint(r) {
			readable++
		}

		if r >= 0x4E00 && r <= 0x9FFF {
			chinese++
		}
	}

	// 可读率
	readableRatio := float64(readable) / float64(total)
	// 中文率
	chineseRatio := float64(chinese) / float64(total)
	// 替换符率
	replacementRatio := float64(replacement) / float64(total)

	return readableRatio > 0.99 && (chineseRatio > 0.05 || replacementRatio < 0.05)
}

// 根据编码名返回 encoding.Encoding
func getEncodingByName(name string) encoding.Encoding {
	name = strings.ToUpper(name)
	switch name {
	case "UTF-8":
		return encoding.Nop
	case "GBK", "GB2312":
		return simplifiedchinese.GBK
	case "BIG5":
		return traditionalchinese.Big5
	case "SHIFT_JIS", "SJIS":
		return japanese.ShiftJIS
	case "EUC-JP":
		return japanese.EUCJP
	case "ISO-8859-1":
		return charmap.ISO8859_1
	case "EUC-KR":
		return korean.EUCKR
	default:
		return nil
	}
}

func determineContentType(objectKey, s3ContentType string) string {
	// 优先使用S3返回的Content-Type
	if s3ContentType != "" && s3ContentType != "application/octet-stream" {
		if strings.HasPrefix(s3ContentType, "text/") ||
			s3ContentType == "application/json" ||
			s3ContentType == "application/xml" ||
			s3ContentType == "application/csv" {
			s3ContentType = s3ContentType + ""
		}
		return s3ContentType
	}

	// 根据文件扩展名猜测Content-Type
	ext := filepath.Ext(objectKey)
	if ext == "" {
		return "application/octet-stream"
	}

	// 常见图片和文本类型的映射
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".bmp":
		return "image/bmp"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	case ".txt":
		return "text/plain"
	case ".csv":
		return "text/csv"
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	default:
		// 使用mime包的标准类型检测
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			return mimeType
		}
		return "application/octet-stream"
	}
}

func DownloadDataset(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset

	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("GetUploadHelper failed, dataType=%s", dataset.StorageType)
		ctx.JSON(http.StatusOK, response.ResponseError(errors.New("Storage type error")))
		return
	}
	resultFileName := dataset.Name + ".zip"
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
		res, err := helper.GetAllObjectsUnderDirWithMarker(dataset.Path, marker, 1000)
		if err != nil {
			log.Error("GetAllObjectsUnderDir err.objectKeyPrefix=%s,err=%v", dataset.Path, err)
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

			if err := AddFileToZip(helper, file, zipWriter, buf); err != nil {
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

	models.IncreaseDatasetRegistryDownloadCount(dataset.ID)

}

func GetDownloadDatasetMeta(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	ctx.JSON(200, response.OuterSuccessWithData(&storage_helper.ObjectMeta{
		ContentLength: dataset.Size,
		ContentType:   "application/zip",
	}))
}

func AddFileToZip(helper storage_helper.StorageHelper, file storage.FileInfo, zipWriter *zip.Writer, buf []byte) error {
	reader, err := helper.OpenFile(file.FullPath)
	if err != nil {
		log.Error("OpenFile err.filePath=%+v,err=%v", file.FullPath, err)
		return fmt.Errorf("failed to open file %s: %w", file.FullPath, err)
	}
	defer reader.Close()

	fileName := path.Join(file.ParenDir, file.FileName)
	safeName := filepath.ToSlash(filepath.Clean(fileName))
	if safeName == "" || safeName == "." || safeName == ".." {
		return fmt.Errorf("invalid filename: %s", file.FileName)
	}

	fDest, err := zipWriter.Create(safeName)
	if err != nil {
		log.Error("zipWriter.Create error.%v", err)
		return fmt.Errorf("failed to create zip entry: %w", err)
	}

	if _, err := io.CopyBuffer(fDest, reader, buf); err != nil {
		log.Error("failed to copy file content: %v", err)
		return fmt.Errorf("failed to write file content: %w", err)
	}

	return nil
}
func EditDataset(ctx *context.Context, req entity.CreateDatasetReq) {
	err := subject_service.EditDataset(req, ctx.AccessContext.Dataset)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func CollectDataset(ctx *context.Context) {
	if ctx.User == nil {
		ctx.JSON(http.StatusOK, response.OuterBizError(response.BuildBizError(http.StatusUnauthorized, "user not login")))
		return
	}
	dataset := ctx.AccessContext.Dataset
	err := subject_service.CollectDataset(ctx.User, dataset)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func UnCollectDataset(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	err := subject_service.UnCollectDataset(ctx.User, dataset)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func RecommendDataset(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	err := subject_service.RecommendDataset(ctx.User, dataset)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func UnRecommendDataset(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	err := subject_service.UnRecommendDataset(ctx.User, dataset)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func DelDataset(ctx *context.Context) {
	dataset := ctx.AccessContext.Dataset
	err := subject_service.DeleteDataset(ctx.User, dataset)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func GetDatasetCreationAvailableUsers(ctx *context.Context) {
	orgs, err := models.GetOrgsCanCreateDatasetByUserID(ctx.User.ID)
	if err != nil {
		log.Error("GetOrgsCanCreateRepoByUserID failed, userID=%d err=%v", ctx.User.ID, err)
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
		user := checkContextUser4Dataset(ctx, orgID)
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

func BatchDeleteDatasets(ctx *context.Context) {
	datasetIdStr := ctx.Query("dataset_ids")
	dataIds := strings.Split(datasetIdStr, ",")
	successIds := subject_service.BatchDeleteDataset(ctx.User, dataIds)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"success": successIds,
	}))
}

func ExportTaskResult(ctx *context.APIContext, req entity.ExportTaskResultReq) {
	dataset := ctx.AccessContext.Dataset
	req.Dataset = dataset
	req.Doer = ctx.User
	id, err := subject_service.ExportTaskResult2Dataset(req)
	if err != nil {
		log.Error("ExportTaskResult failed, taskId=%d err=%v", req.TaskId, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"id": id,
	}))
}

func GetExportDatasetByMsgId(ctx *context.Context) {
	progressId := ctx.Query("id")
	progress, err := redis_client.Get(progressId)
	if err != nil {
		log.Error("GetExportDatasetByMsgId failed, id=%s err=%v", progressId, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
	}
	msgMap := make(map[string]int, 0)
	err = json.Unmarshal([]byte(progress), &msgMap)
	if err != nil {
		log.Error("GetExportDataSetByMsgId Unmarshal failed, id=%s err=%v", progressId, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(msgMap))
}

func StartDatasetMigrate(ctx *context.Context) {
	producerCount := ctx.QueryInt("producer")
	consumerCount := ctx.QueryInt("consumer")
	if producerCount == 0 {
		producerCount = 2
	}
	if consumerCount == 0 {
		consumerCount = 5
	}
	err := subject_service.StartDatasetMigrate(producerCount, consumerCount)
	if err != nil {
		log.Error("StartDatasetMigrate failed err=%v", err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func StopDatasetMigrate(ctx *context.Context) {
	err := subject_service.StopDatasetMigrate()
	if err != nil {
		log.Error("StartDatasetMigrate failed err=%v", err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func HandleAllOldDatasets(ctx *context.Context) {
	err := subject_service.HandleAllOldDatasets()
	if err != nil {
		log.Error("HandleAllOldDatasets failed err=%v", err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func GetDatasetMigrateStatus(ctx *context.Context) {
	res := subject_service.GetDatasetMigrateStatus()
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}

func HandleOneOldDataset(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	err := subject_service.HandleOneOldDataset(uuid)
	if err != nil {
		log.Error("HandleOneOldDataset failed err=%v", err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func CheckOld(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	sourceId := ctx.QueryInt64("source_id")
	isChanged := subject_service.IsOldDatasetNameChanged(uuid, sourceId)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"is_changed": isChanged,
	}))
}

// 组织主页-获取数据集标签合集
func GetOwnedPublicDatasetTags(ctx *context.APIContext) {
	res, err := subject_service.GetOwnedPublicDatasetTags(ctx.Org.Organization.ID)
	if err != nil {
		log.Error("GetOwnedPublicDatasetTags failed, userId=%d err=%v", ctx.Org.Organization.ID, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}
