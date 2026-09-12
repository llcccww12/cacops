package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/repository"
	"code.gitea.io/gitea/services/storage_limit"
	uuid "github.com/satori/go.uuid"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

const (
	tplIndex         base.TplName = "repo/datasets/index"
	tplDatasetCreate base.TplName = "repo/datasets/create"
	tplDatasetEdit   base.TplName = "repo/datasets/edit"
	taskstplIndex    base.TplName = "repo/datasets/tasks/index"
	tplReference     base.TplName = "repo/datasets/reference"
)

// MustEnableDataset check if repository enable internal dataset
func MustEnableDataset(ctx *context.Context) {
	if !ctx.Repo.CanRead(models.UnitTypeDatasets) {
		ctx.NotFound("MustEnableDataset", nil)
		return
	}
}

func NewFilterPrivateAttachments(ctx *context.Context, list []*models.Attachment, repo *models.Repository) []*models.Attachment {

	if ctx.Repo.CanWrite(models.UnitTypeDatasets) {
		log.Info("can write.")
		return list
	} else {
		if repo.Owner == nil {
			repo.GetOwner()
		}
		permission := false
		if !permission && ctx.User != nil {
			isCollaborator, _ := repo.IsCollaborator(ctx.User.ID)
			isInRepoTeam, _ := repo.IsInRepoTeam(ctx.User.ID)
			if isCollaborator || isInRepoTeam {
				log.Info("Collaborator user may visit the attach.")
				permission = true
			}
		}

		var publicList []*models.Attachment
		for _, attach := range list {
			if !attach.IsPrivate {
				publicList = append(publicList, attach)
			} else {
				if permission {
					publicList = append(publicList, attach)
				}
			}
		}
		return publicList
	}
}

func QueryDataSet(ctx *context.Context) []*models.Attachment {
	repo := ctx.Repo.Repository

	dataset, err := models.GetDatasetByRepo(repo)
	if err != nil {
		log.Error("zou  not found dataset 1")
		ctx.NotFound("GetDatasetByRepo", err)
		return nil
	}

	if ctx.Query("type") == "" {
		log.Error("zou  not found  type 2")
		ctx.NotFound("type error", nil)
		return nil
	}
	err = models.GetDatasetAttachments(ctx.QueryInt("type"), ctx.IsSigned, ctx.User, dataset)
	if err != nil {
		ctx.ServerError("GetDatasetAttachments", err)
		return nil
	}
	attachments := NewFilterPrivateAttachments(ctx, dataset.Attachments, repo)

	ctx.Data["SortType"] = ctx.Query("sort")

	sort.Slice(attachments, func(i, j int) bool {
		return attachments[i].CreatedUnix > attachments[j].CreatedUnix
	})

	return attachments
}

func DatasetIndex(ctx *context.Context) {
	log.Info("dataset index 1")
	MustEnableDataset(ctx)
	ctx.Data["PageIsDataset"] = true
	ctx.Data["max_model_size"] = int64(setting.MaxModelSize * MODEL_MAX_SIZE)
	ctx.Data["SortType"] = ctx.Query("sort")

	repo := ctx.Repo.Repository

	dataset, err := models.GetDatasetByRepo(repo)
	ctx.Data["CanWrite"] = ctx.Repo.CanWrite(models.UnitTypeDatasets)
	if err != nil {
		log.Warn("query dataset, not found.")
		ctx.HTML(200, tplIndex)
		return
	}
	cloudbrainType := -1
	if ctx.Query("type") != "" {

		cloudbrainType = ctx.QueryInt("type")
	}
	err = models.GetDatasetAttachments(cloudbrainType, ctx.IsSigned, ctx.User, dataset)
	if err != nil {
		ctx.ServerError("GetDatasetAttachments", err)
		return
	}

	attachments := NewFilterPrivateAttachments(ctx, dataset.Attachments, repo)

	if ctx.Data["SortType"] == "nameAsc" {
		sort.Slice(attachments, func(i, j int) bool {
			return strings.ToLower(attachments[i].Name) < strings.ToLower(attachments[j].Name)
		})
	} else if ctx.Data["SortType"] == "nameDesc" {
		sort.Slice(attachments, func(i, j int) bool {
			return strings.ToLower(attachments[i].Name) > strings.ToLower(attachments[j].Name)
		})
	} else if ctx.Data["SortType"] == "sizeAsc" {
		sort.Slice(attachments, func(i, j int) bool {
			return attachments[i].Size < attachments[j].Size
		})
	} else if ctx.Data["SortType"] == "sizeDesc" {
		sort.Slice(attachments, func(i, j int) bool {
			return attachments[i].Size > attachments[j].Size
		})
	} else if ctx.Data["SortType"] == "timeAsc" {
		sort.Slice(attachments, func(i, j int) bool {
			return attachments[i].CreatedUnix < attachments[j].CreatedUnix
		})
	} else {
		sort.Slice(attachments, func(i, j int) bool {
			return attachments[i].CreatedUnix > attachments[j].CreatedUnix
		})
	}

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pagesize := ctx.QueryInt("pagesize")
	if pagesize <= 0 {
		pagesize = 10
	}
	pager := context.NewPagination(len(attachments), pagesize, page, 5)

	pageAttachments := getPageAttachments(attachments, page, pagesize)

	//load attachment creator
	for _, attachment := range pageAttachments {
		uploader, err1 := models.GetUserByID(attachment.UploaderID)
		if err1 != nil {
			log.Info("query dataset user error." + err1.Error())
		}
		attachment.Uploader = uploader
		if !strings.HasSuffix(attachment.Name, ".zip") && !strings.HasSuffix(attachment.Name, ".tar.gz") {
			attachment.DecompressState = -1 //非压缩文件
		}

	}
	uuidList := make([]string, 0, len(pageAttachments))
	uuidMap := make(map[string]int, 0)
	newDatasetMap := make(map[string]*models.DatasetRegistry)
	key := redis_key.ProcessingOldDataset()
	for _, attach := range pageAttachments {
		if !strings.HasSuffix(attach.Name, ".zip") && !strings.HasSuffix(attach.Name, ".tar.gz") {
			uuidMap[attach.UUID] = 0
			continue
		}
		if attach.DecompressState != models.DecompressStateDone {
			uuidMap[attach.UUID] = 0
			continue
		}

		flag, _ := redis_client.SISMember(key, attach.UUID)
		if flag > 0 {
			uuidMap[attach.UUID] = 9
		} else {
			uuidMap[attach.UUID] = 1
		}

		uuidList = append(uuidList, attach.UUID)
	}

	list, _ := models.GetHandledOldDatasetByIDs(uuidList)
	datasetIds := make([]string, 0, len(list))
	for _, r := range list {
		uuidMap[r.ID] = 2
		datasetIds = append(datasetIds, r.ID)
	}
	datasets, err := models.GetDatasetRegistryListByIDs(datasetIds)
	for _, d := range datasets {
		newDatasetMap[d.ID] = d
	}

	for i := 0; i < len(pageAttachments); i++ {
		pageAttachments[i].DatasetMigrateStatus = uuidMap[pageAttachments[i].UUID]
		if _, ok := newDatasetMap[pageAttachments[i].UUID]; ok {
			d := newDatasetMap[pageAttachments[i].UUID]
			d.GetOwner()
			pageAttachments[i].NewOwnerName = d.Owner.Name
			pageAttachments[i].NewDatasetName = d.Name
		}
	}

	ctx.Data["Page"] = pager

	ctx.Data["Title"] = ctx.Tr("dataset.show_dataset")
	ctx.Data["Link"] = ctx.Repo.RepoLink + "/datasets"
	ctx.Data["dataset"] = dataset
	ctx.Data["Attachments"] = pageAttachments
	ctx.Data["IsOwner"] = true
	ctx.Data["StoreType"] = setting.Attachment.StoreType
	ctx.Data["Type"] = cloudbrainType

	renderAttachmentSettings(ctx)
	log.Info("dataset index finished.")
	ctx.HTML(200, tplIndex)
}

func getPageAttachments(attachments []*models.Attachment, page int, pagesize int) []*models.Attachment {
	begin := (page - 1) * pagesize
	end := (page) * pagesize

	if begin > len(attachments)-1 {
		return nil
	}
	if end > len(attachments)-1 {
		return attachments[begin:]
	} else {
		return attachments[begin:end]
	}

}

func CreateDataset(ctx *context.Context) {

	MustEnableDataset(ctx)
	ctx.Data["PageIsDataset"] = true

	ctx.HTML(200, tplDatasetCreate)
}

func EditDataset(ctx *context.Context) {

	MustEnableDataset(ctx)
	ctx.Data["PageIsDataset"] = true
	datasetId, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)

	dataset, _ := models.GetDatasetByID(datasetId)
	if dataset == nil {
		ctx.Error(http.StatusNotFound, "")
		return
	}
	ctx.Data["Dataset"] = dataset

	ctx.HTML(200, tplDatasetEdit)
}

func CreateDatasetPost(ctx *context.Context, form auth.CreateDatasetForm) {

	dataset := &models.Dataset{}

	if !NamePattern.MatchString(form.Title) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.title_format_err")))
		return
	}
	if utf8.RuneCountInString(form.Description) > 1024 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", "1024")))
		return
	}

	dataset.RepoID = ctx.Repo.Repository.ID
	dataset.UserID = ctx.User.ID
	dataset.Category = form.Category
	dataset.Task = form.Task
	dataset.Title = form.Title
	dataset.License = form.License
	dataset.Description = form.Description
	dataset.DownloadTimes = 0
	if ctx.Repo.Repository.IsPrivate {
		dataset.Status = 0
	} else {
		dataset.Status = 1
	}
	err := models.CreateDataset(dataset)
	if err != nil {
		log.Error("fail to create dataset", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.create_dataset_fail")))
	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}
func ReferenceDatasetDelete(ctx *context.Context) {
	repoID := ctx.Repo.Repository.ID
	datasetId, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)

	oldDatasetIds := models.GetDatasetIdsByRepoID(repoID)

	var newDatasetIds []int64

	for _, tempDatasetId := range oldDatasetIds {
		if datasetId != tempDatasetId {
			newDatasetIds = append(newDatasetIds, tempDatasetId)
		}
	}
	err := models.NewDatasetIdsByRepoID(repoID, newDatasetIds)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage("dataset.cancel_reference_dataset_fail"))
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessage)

}

func ReferenceDatasetPost(ctx *context.Context, form auth.ReferenceDatasetForm) {
	repoID := ctx.Repo.Repository.ID
	err := models.NewDatasetIdsByRepoID(repoID, form.DatasetID)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage("dataset.reference_dataset_fail"))
	}

	ctx.JSON(http.StatusOK, models.BaseOKMessage)

}

func EditDatasetPost(ctx *context.Context, form auth.EditDatasetForm) {
	ctx.Data["PageIsDataset"] = true

	ctx.Data["Title"] = ctx.Tr("dataset.edit_dataset")

	if !NamePattern.MatchString(form.Title) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.title_format_err")))
		return
	}
	if utf8.RuneCountInString(form.Description) > 1024 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", 1024)))
		return
	}

	rel, err := models.GetDatasetByID(form.ID)
	ctx.Data["dataset"] = rel

	if err != nil {
		log.Error("failed to query dataset", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.query_dataset_fail")))
		return
	}

	rel.Title = form.Title
	rel.Description = form.Description
	rel.Category = form.Category
	rel.Task = form.Task
	rel.License = form.License
	if err = models.UpdateDataset(models.DefaultDBContext(), rel); err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.query_dataset_fail")))
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessage)
}

func DatasetAction(ctx *context.Context) {
	var err error
	datasetId, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)
	switch ctx.Params(":action") {
	case "star":
		err = models.StarDataset(ctx.User.ID, datasetId, true)
	case "unstar":
		err = models.StarDataset(ctx.User.ID, datasetId, false)

	}
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.star_fail", ctx.Params(":action"))))
	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}

func DatasetMultiple(ctx *context.Context, opts *models.SearchDatasetOptions) {
	page := ctx.QueryInt("page")
	keyword := strings.Trim(ctx.Query("q"), " ")
	opts.Keyword = keyword
	if opts.SearchOrderBy.String() == "" {
		opts.SearchOrderBy = models.SearchOrderByRecentUpdated
	}
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.DatasetPagingNum
	}
	opts.RecommendOnly = ctx.QueryBool("recommend")
	opts.ListOptions = models.ListOptions{
		Page:     page,
		PageSize: pageSize,
	}
	opts.User = ctx.User

	category := ctx.Query("category")
	task := ctx.Query("task")
	license := ctx.Query("license")
	opts.Category = category
	opts.Task = task
	opts.License = license
	var datasets models.DatasetList
	var count int64
	var err error
	if opts.ExploreMine {
		datasets, count, err = models.SearchMyDataset(opts)
	} else {
		datasets, count, err = models.SearchDataset(opts)
	}

	if err != nil {
		ctx.ServerError("datasets", err)
		return
	}

	data, err := json.Marshal(datasets)

	if err != nil {
		log.Error("json.Marshal failed:", err.Error())
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"error_msg":   err.Error(),
			"data":        "",
		})
		return
	}
	ctx.JSON(200, map[string]string{
		"result_code": "0",
		"data":        string(data),
		"count":       strconv.FormatInt(count, 10),
	})
}

func CurrentRepoDatasetMultiple(ctx *context.Context) {
	datasetIds := models.GetDatasetIdsByRepoID(ctx.Repo.Repository.ID)
	searchOrderBy := getSearchOrderByInValues(datasetIds)
	opts := &models.SearchDatasetOptions{
		RepoID:          ctx.Repo.Repository.ID,
		NeedAttachment:  true,
		CloudBrainType:  ctx.QueryInt("type"),
		DatasetIDs:      datasetIds,
		SearchOrderBy:   searchOrderBy,
		JustNeedZipFile: true,
	}

	DatasetMultiple(ctx, opts)

}

func getSearchOrderByInValues(datasetIds []int64) models.SearchOrderBy {
	if len(datasetIds) == 0 {
		return ""
	}
	searchOrderBy := "CASE id "
	for i, id := range datasetIds {
		searchOrderBy += fmt.Sprintf(" WHEN %d THEN %d", id, i+1)
	}
	searchOrderBy += " ELSE 0 END"
	return models.SearchOrderBy(searchOrderBy)
}

func MyDatasetsMultiple(ctx *context.Context) {

	opts := &models.SearchDatasetOptions{
		UploadAttachmentByMe: true,
		NeedAttachment:       true,
		CloudBrainType:       ctx.QueryInt("type"),
		JustNeedZipFile:      true,
	}
	DatasetMultiple(ctx, opts)

}

func ReferenceDatasetAvailable(ctx *context.Context) {

	opts := &models.SearchDatasetOptions{
		PublicOnly:      true,
		NeedAttachment:  false,
		CloudBrainType:  models.TypeCloudBrainAll,
		SearchOrderBy:   models.SearchOrderByDefault,
		JustNeedZipFile: false,
	}
	dataset, _ := models.GetDatasetByRepo(&models.Repository{ID: ctx.Repo.Repository.ID})
	if dataset != nil {
		opts.ExcludeDatasetId = dataset.ID
	}
	DatasetMultiple(ctx, opts)

}

func PublicDatasetMultiple(ctx *context.Context) {

	opts := &models.SearchDatasetOptions{
		PublicOnly:      true,
		NeedAttachment:  true,
		CloudBrainType:  ctx.QueryInt("type"),
		SearchOrderBy:   models.SearchOrderByDefault,
		JustNeedZipFile: true,
	}
	DatasetMultiple(ctx, opts)

}

func MyFavoriteDatasetMultiple(ctx *context.Context) {

	opts := &models.SearchDatasetOptions{
		StarByMe:        true,
		DatasetIDs:      models.GetDatasetIdsStarByUser(ctx.User.ID),
		NeedAttachment:  true,
		CloudBrainType:  ctx.QueryInt("type"),
		JustNeedZipFile: true,
	}
	DatasetMultiple(ctx, opts)
}
func ReferenceDataset(ctx *context.Context) {
	MustEnableDataset(ctx)
	ctx.Data["PageIsDataset"] = true
	ctx.Data["MaxReferenceDatasetNum"] = setting.RepoMaxReferenceDatasetNum
	ctx.Data["CanWrite"] = ctx.Repo.CanWrite(models.UnitTypeDatasets)
	ctx.HTML(200, tplReference)

}

func ReferenceDatasetData(ctx *context.Context) {
	MustEnableDataset(ctx)
	datasetIds := models.GetDatasetIdsByRepoID(ctx.Repo.Repository.ID)
	var datasets models.DatasetList
	var err error
	if len(datasetIds) > 0 {

		opts := &models.SearchDatasetOptions{
			DatasetIDs:     datasetIds,
			NeedAttachment: false,
			CloudBrainType: models.TypeCloudBrainAll,
			ListOptions: models.ListOptions{
				Page:     1,
				PageSize: setting.RepoMaxReferenceDatasetNum,
			},
			SearchOrderBy:  getSearchOrderByInValues(datasetIds),
			QueryReference: true,
		}
		datasets, _, err = models.SearchDataset(opts)
		if err != nil {
			ctx.ServerError("SearchDatasets", err)
			return
		}
	}

	ctx.JSON(http.StatusOK, repository.ConvertToDatasetWithStar(ctx, datasets))

}

func GetDatasetStatus(ctx *context.Context) {

	var (
		err error
	)

	UUID := ctx.Params(":uuid")
	attachment, err := models.GetAttachmentByUUID(UUID)
	if err != nil {
		log.Error("GetDatasetStarByUser failed:", err.Error())
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"error_msg":   err.Error(),
			"data":        "",
		})
		return
	}

	ctx.JSON(200, map[string]string{
		"result_code":      "0",
		"UUID":             UUID,
		"AttachmentStatus": fmt.Sprint(attachment.DecompressState),
	})
}

func GetDataSetSelectItemByJobId(ctx *context.Context) {
	JobID := ctx.Query("jobId")
	versionName := ctx.Query("versionName")
	re := map[string]interface{}{
		"code": "-1",
	}
	task, err := models.GetCloudbrainByJobIDAndVersionName(JobID, versionName)
	if err != nil {
		task, err = models.GetRepoCloudBrainByJobID(ctx.Repo.Repository.ID, JobID)
		if err != nil {
			log.Info("query task error." + err.Error())
			re["msg"] = "Query cloudbrain task error." + err.Error()
			ctx.JSON(200, re)
			return
		}
	}

	taskConfig := task.GetCloudbrainConfig()
	storageType := storage_helper.GetStorageIntTypeFromStorageType(entity.StorageType(taskConfig.OutputStorageType))

	result, err := getModelFromObjectSave(task.JobName, task.ComputeResource, storageType, task.VersionName)
	if err != nil {
		re["msg"] = "Query model file error, " + err.Error()
		ctx.JSON(200, re)
		return
	}
	reFile := make([]storage.FileInfo, 0)
	for _, file := range result {
		if strings.HasSuffix(strings.ToLower(file.FileName), ".zip") || strings.HasSuffix(strings.ToLower(file.FileName), ".tar.gz") {
			reFile = append(reFile, file)
		}
	}
	re["code"] = 0
	re["files"] = reFile
	ctx.JSON(200, re)
}

func GetExportDataSetByMsgId(ctx *context.Context) {
	progressId := ctx.Query("progressId")
	progress, err := redis_client.Get(progressId)
	if err == nil {
		ctx.JSON(200, progress)
	} else {
		ctx.JSON(200, -1)
	}
}

func GetCurrentDataSet(ctx *context.Context) {
	re := map[string]interface{}{
		"code": "-1",
	}
	dataset, err := models.GetDatasetByRepo(ctx.Repo.Repository)
	if err == nil {
		re["code"] = 0
		re["dataset"] = dataset
		ctx.JSON(200, re)
	} else {
		ctx.JSON(200, re)
	}
}

func ExportModelToExistDataSet(ctx *context.Context) {
	modelSelectedFile := ctx.Query("modelSelectedFile")
	log.Info("modelSelectedFile=" + modelSelectedFile)
	datasetId := ctx.QueryInt64("datasetId")
	log.Info("datasetId=" + fmt.Sprint(datasetId))

	re := map[string]string{
		"code": "-1",
	}
	description := ctx.Query("description")
	jobId := ctx.Query("jobId")
	cloudbrainId := ctx.QueryInt64("cloudbrain_id")
	storeType := models.GetDefaultStorageType()
	versionName := ctx.Query("versionName")
	dataset, err := models.GetDatasetByID(datasetId)
	if err != nil || dataset == nil {
		log.Info("Not found dataset.")
		re["msg"] = "Not found dataset."
		ctx.JSON(200, re)
		return
	}
	var aiTask *models.Cloudbrain
	if cloudbrainId > 0 {
		aiTask, err = models.GetCloudbrainByCloudbrainID(cloudbrainId)
	} else {
		aiTask, err = models.GetCloudbrainByJobIDAndVersionName(jobId, versionName)
		if err != nil {
			aiTask, err = models.GetRepoCloudBrainByJobID(ctx.Repo.Repository.ID, jobId)
		}
	}
	if err != nil {
		log.Info("query task error." + err.Error())
		re["msg"] = "Query cloudbrain task error." + err.Error()
		ctx.JSON(200, re)
		return
	}
	msgKey := fmt.Sprint(datasetId) + "_" + aiTask.JobID + "_" + aiTask.VersionName
	msgMap := make(map[string]int, 0)
	msgMap["##type##"] = storeType
	filterFiles := strings.Split(modelSelectedFile, ";")
	for _, shortFile := range filterFiles {
		msgMap[shortFile] = 0
	}
	setProgress(msgKey, msgMap)
	go asyncToExportDataset(dataset, storeType, modelSelectedFile, aiTask, ctx.User, msgKey, msgMap, aiTask.VersionName, description, ctx)
	ctx.JSON(200, map[string]string{
		"code":       "0",
		"progressId": msgKey,
	})
}

func setProgress(msgKey string, msgMap map[string]int) {
	msgMapJson, _ := json.Marshal(msgMap)
	redisValue := string(msgMapJson)
	log.Info("set redis key=" + msgKey + " value=" + redisValue)
	re, err := redis_client.Setex(msgKey, redisValue, 3600*24*time.Second)
	if err == nil {
		log.Info("re =" + fmt.Sprint(re))
	} else {
		log.Info("set redis error:" + err.Error())
	}
}

func asyncToExportDataset(dataset *models.Dataset, storeType int, modelSelectedFile string, aiTask *models.Cloudbrain, user *models.User, msgKey string, msgMap map[string]int, versionName string, description string, ctx *context.Context) {
	filterFiles := strings.Split(modelSelectedFile, ";")
	aiConfig := aiTask.GetCloudbrainConfig()
	if aiConfig == nil {
		log.Error("GetCloudbrainConfig empty, cloudbrain=%s", aiTask.JobName)
		return
	}
	for _, shortFile := range filterFiles {
		uuid := uuid.NewV4()
		id := uuid.String()
		fileName := getFileName(shortFile)
		log.Info("shortSrcFile=" + shortFile + "  fileName=" + fileName)

		if aiConfig.OutputStorageType == string(entity.MINIO) {
			size := getFileSizeFromMinio(aiConfig.OutputObjectPrefix, aiConfig.OutputBucket, shortFile)
			if storage_limit.IsFileUploadOverLimit(user.ID, size) {
				msgMap[shortFile] = -3
			} else if isExistInAttachment(fileName, size, dataset, storeType) {
				msgMap[shortFile] = -2
			} else {
				err := exportModelToDataFromMinio(id, aiConfig.OutputObjectPrefix, aiConfig.OutputBucket, shortFile, storeType)
				if err != nil {
					msgMap[shortFile] = -1
				} else {
					msgMap[shortFile] = 100
					finishedUploadAttachment(dataset.ID, size, id, fileName, description, storeType, user)
				}
			}

			setProgress(msgKey, msgMap)

		} else if aiConfig.OutputStorageType == string(entity.OBS) {
			size := getFileSizeFromObs(aiConfig.OutputObjectPrefix, aiConfig.OutputBucket, shortFile)
			if storage_limit.IsFileUploadOverLimit(user.ID, size) {
				msgMap[shortFile] = -3
			} else if isExistInAttachment(fileName, size, dataset, storeType) {
				msgMap[shortFile] = -2
			} else {
				err := exportModelToDataFromObs(id, aiConfig.OutputObjectPrefix, aiConfig.OutputBucket, shortFile, storeType)
				if err != nil {
					msgMap[shortFile] = -1
				} else {
					msgMap[shortFile] = 100
					finishedUploadAttachment(dataset.ID, size, id, fileName, description, storeType, user)
				}
			}
			setProgress(msgKey, msgMap)
		} else {
			log.Info("ExportModelToExistDataSet cannot suppport the ComputeResource" + aiTask.ComputeResource)
			msgMap[shortFile] = -1
			setProgress(msgKey, msgMap)
		}
	}
}

func isExistInAttachment(fileName string, size int64, dataset *models.Dataset, storeType int) bool {
	err := models.GetDatasetAttachments(storeType, false, nil, dataset)
	if err == nil && dataset.Attachments != nil && len(dataset.Attachments) > 0 {
		for _, attach := range dataset.Attachments {
			if attach.Name == fileName && attach.Size == size {
				log.Info("Has found same file. so not to create.")
				return true
			}
		}
	} else {
		log.Info("Not found attachments....")
		if err != nil {
			log.Info("error=" + err.Error())
		}
	}
	return false
}

func getFileSizeFromObs(objectKeyPrefix, bucketName string, modelSelectedFile string) int64 {
	objectkey := strings.TrimPrefix(objectKeyPrefix, "/")
	modelSrcPrefix := objectkey + "/"
	modelFile := modelSrcPrefix + modelSelectedFile
	log.Info("modelfile=" + modelFile)
	totalSize := storage.ObsGetFilesSize(bucketName, []string{modelFile})
	return totalSize
}

func exportModelToDataFromObs(id string, objectKeyPrefix, bucketName string, modelSelectedFile string, storeType int) error {
	objectkey := strings.TrimPrefix(objectKeyPrefix, "/")
	modelSrcPrefix := objectkey + "/"
	modelFile := modelSrcPrefix + modelSelectedFile
	log.Info("modelfile=" + modelFile)
	//totalSize := storage.ObsGetFilesSize(setting.Bucket, []string{modelFile})

	fileName := getFileName(modelSelectedFile)

	if storeType == models.StorageTypeMinio {
		reader, err := storage.ObsDownloadAFile(bucketName, modelFile)
		if err == nil {
			defer reader.Close()
			destPath := strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath, path.Join(id[0:1], id[1:2], id)), "/")
			log.Info("destPath=" + destPath)
			bucketName := setting.Attachment.Minio.Bucket
			_, minioErr := storage.Attachments.UploadContent(bucketName, destPath, reader)
			if minioErr != nil {
				log.Info("upload to minio failed 2.err=" + minioErr.Error())
				return minioErr
			}
		} else {
			log.Info("upload to minio failed 1.err=" + err.Error())
			return err
		}
	} else {
		destObjectKey := strings.TrimPrefix(path.Join(setting.BasePath, path.Join(id[0:1], id[1:2], id, fileName)), "/")
		log.Info("destPath=" + destObjectKey)
		log.Info("copy to bucket=" + setting.Bucket + " objectKey=" + destObjectKey)
		obsErr := storage.ObsCopyFile(bucketName, modelFile, setting.Bucket, destObjectKey)
		if obsErr != nil {
			log.Info("upload to obs failed.err=" + obsErr.Error())
			return obsErr
		}
	}
	return nil
}

func getFileSizeFromMinio(objectKeyPrefix, bucketName string, modelSelectedFile string) int64 {
	modelSrcPrefix := strings.TrimSuffix(objectKeyPrefix, "/") + "/"
	log.Info("   modelSrcPrefix=" + modelSrcPrefix + "  bucket=" + bucketName)
	modelFile := modelSrcPrefix + modelSelectedFile
	totalSize := storage.MinioGetFilesSize(bucketName, []string{modelFile})
	return totalSize
}

func exportModelToDataFromMinio(id string, objectKeyPrefix, bucketName string, modelSelectedFile string, storeType int) error {
	modelSrcPrefix := strings.TrimSuffix(objectKeyPrefix, "/") + "/"
	log.Info("   modelSrcPrefix=" + modelSrcPrefix + "  bucket=" + bucketName)
	modelFile := modelSrcPrefix + modelSelectedFile
	//totalSize := storage.MinioGetFilesSize(bucketName, []string{modelFile})

	fileName := getFileName(modelSelectedFile)
	if storeType == models.StorageTypeMinio {
		destPath := strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath, path.Join(id[0:1], id[1:2], id)), "/")
		log.Info("destPath=" + destPath)
		_, minioErr := storage.MinioCopyAFile(bucketName, modelFile, bucketName, destPath)
		if minioErr != nil {
			log.Info("upload to minio failed.err=" + minioErr.Error())
			return minioErr
		}
	} else {
		reader, err := storage.Attachments.DownloadAFile(bucketName, modelFile)
		if err == nil {
			defer reader.Close()
			destObjectKey := strings.TrimPrefix(path.Join(setting.BasePath, path.Join(id[0:1], id[1:2], id, fileName)), "/")
			log.Info("destPath=" + destObjectKey)
			log.Info("upload to bucket=" + setting.Bucket + " objectKey=" + destObjectKey)
			obsErr := storage.PutReaderToObs(setting.Bucket, destObjectKey, reader)
			if obsErr != nil {
				log.Info("upload to obs failed 1.err=" + obsErr.Error())
				return obsErr
			}
		} else {
			log.Info("upload to obs failed 2.err=" + err.Error())
			return err
		}
	}
	return nil
}

func getFileName(shortFile string) string {
	index := strings.LastIndex(shortFile, "/")
	if index > 0 {
		return shortFile[index+1:]
	}
	return shortFile
}
