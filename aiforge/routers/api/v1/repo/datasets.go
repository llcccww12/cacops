package repo

import (
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/services/role"
	"fmt"
	"strings"

	"code.gitea.io/gitea/modules/convert"

	api "code.gitea.io/gitea/modules/structs"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	routerRepo "code.gitea.io/gitea/routers/repo"
)

func PublicDatasetMultiple(ctx *context.APIContext) {

	opts := &models.SearchDatasetOptions{
		PublicOnly:     true,
		NeedAttachment: true,
		CloudBrainType: ctx.QueryInt("type"),
	}
	datasetMultiple(ctx, opts)

}

func MyFavoriteDatasetMultiple(ctx *context.APIContext) {

	opts := &models.SearchDatasetOptions{
		StarByMe:       true,
		DatasetIDs:     models.GetDatasetIdsStarByUser(ctx.User.ID),
		NeedAttachment: true,
		CloudBrainType: ctx.QueryInt("type"),
	}
	datasetMultiple(ctx, opts)
}

func CurrentRepoDatasetMultiple(ctx *context.APIContext) {
	datasetIds := models.GetDatasetIdsByRepoID(ctx.Repo.Repository.ID)
	searchOrderBy := getSearchOrderByInValues(datasetIds)
	opts := &models.SearchDatasetOptions{
		RepoID:         ctx.Repo.Repository.ID,
		NeedAttachment: true,
		CloudBrainType: ctx.QueryInt("type"),
		DatasetIDs:     datasetIds,
		SearchOrderBy:  searchOrderBy,
	}

	datasetMultiple(ctx, opts)

}

func CurrentRepoDatasetInfoWithoutAttachment(ctx *context.APIContext) {
	dataset, err := models.GetDatasetByRepo(ctx.Repo.Repository)

	if err != nil {
		log.Warn("can not get dataset.", err)
		ctx.JSON(200, map[string]interface{}{
			"code":    0,
			"message": "",
			"data":    []*api.Dataset{},
		})
		return
	}

	dataset.Repo = ctx.Repo.Repository
	ctx.JSON(200, map[string]interface{}{
		"code":    0,
		"message": "",
		"data":    []*api.Dataset{convert.ToDataset(dataset)},
	})

}

func MyDatasetsMultiple(ctx *context.APIContext) {

	opts := &models.SearchDatasetOptions{
		UploadAttachmentByMe: true,
		NeedAttachment:       true,
		CloudBrainType:       ctx.QueryInt("type"),
	}
	datasetMultiple(ctx, opts)

}
func datasetMultiple(ctx *context.APIContext, opts *models.SearchDatasetOptions) {
	page := ctx.QueryInt("page")
	if page < 1 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize < 1 {
		pageSize = setting.UI.DatasetPagingNum
	}

	keyword := strings.Trim(ctx.Query("q"), " ")
	opts.Keyword = keyword
	if opts.SearchOrderBy.String() == "" {
		opts.SearchOrderBy = models.SearchOrderByRecentUpdated
	}

	opts.RecommendOnly = ctx.QueryBool("recommend")
	opts.ListOptions = models.ListOptions{
		Page:     page,
		PageSize: pageSize,
	}
	opts.JustNeedZipFile = true
	opts.User = ctx.User

	datasets, count, err := models.SearchDataset(opts)

	if err != nil {
		log.Error("json.Marshal failed:", err.Error())
		ctx.JSON(200, map[string]interface{}{
			"code":    1,
			"message": err.Error(),
			"data":    []*api.Dataset{},
			"count":   0,
		})
		return
	}
	var convertDatasets []*api.Dataset
	for _, dataset := range datasets {
		convertDatasets = append(convertDatasets, convert.ToDataset(dataset))
	}

	ctx.JSON(200, map[string]interface{}{
		"code":    0,
		"message": "",
		"data":    convertDatasets,
		"count":   count,
	})
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

func GetDataSetSelectItemByJobId(ctx *context.APIContext) {
	routerRepo.GetDataSetSelectItemByJobId(ctx.Context)
}

func GetExportDataSetByMsgId(ctx *context.APIContext) {
	routerRepo.GetExportDataSetByMsgId(ctx.Context)
}

func ExportModelToExistDataSet(ctx *context.APIContext) {
	routerRepo.ExportModelToExistDataSet(ctx.Context)
}

func ListDatasetFilesForSDK(ctx *context.APIContext) {
	log.Info("ListDatasetFilesForSDK by api.")

	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		ctx.JSON(200, map[string]interface{}{
			"code": -1,
			"msg":  "平台新版数据集已上线，请升级至 openi 3.0.0版本以使用最新功能。执行以下命令进行升级： pip install openi==3.0.0",
			"data": nil,
		})
		return
	}

	repo := ctx.Repo.Repository
	dataset, err := models.GetDatasetByRepo(repo)

	if err != nil {
		log.Warn("can not get dataset.", err)
		ctx.JSON(200, map[string]interface{}{
			"code": -1,
			"msg":  "error for getting dataset.",
			"data": nil,
		})
		return
	}
	ctx.Data["CanWrite"] = ctx.Repo.CanWrite(models.UnitTypeDatasets)

	cloudbrainType := ctx.QueryInt("type")
	err = models.GetDatasetAttachments(cloudbrainType, ctx.IsSigned, ctx.User, dataset)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"code": -2,
			"msg":  "error for getting attachments.",
			"data": nil,
		})
		return
	}
	attachments := routerRepo.NewFilterPrivateAttachments(ctx.Context, dataset.Attachments, repo)

	dataset.Attachments = attachments
	dataset.Repo = ctx.Repo.Repository
	ctx.JSON(200, map[string]interface{}{
		"code": 1,
		"msg":  "",
		"data": convert.ToDataset(dataset),
	})
}

func CreateDatasetAPI(ctx *context.APIContext, form auth.CreateDatasetForm) {
	log.Info("Create dataset by api.")
	routerRepo.CreateDatasetPost(ctx.Context, form)
}
