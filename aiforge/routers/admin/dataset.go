package admin

import (
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

const (
	tplDatasets         base.TplName = "admin/dataset/list"
	tplAdminModelManage base.TplName = "admin/model/list"
)

func Datasets(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("admin.datasets")
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminDatasets"] = true

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	var (
		datasets []*models.Dataset
		count    int64
		err      error
		orderBy  models.SearchOrderBy
	)

	ctx.Data["SortType"] = ctx.Query("sort")
	switch ctx.Query("sort") {
	case "newest":
		orderBy = models.SearchOrderByNewest
	case "oldest":
		orderBy = models.SearchOrderByOldest
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated
	case "leastupdate":
		orderBy = models.SearchOrderByLeastUpdated
	case "reversealphabetically":
		orderBy = models.SearchOrderByAlphabeticallyReverse
	case "alphabetically":
		orderBy = models.SearchOrderByAlphabetically
	case "reversesize":
		orderBy = models.SearchOrderBySizeReverse
	case "size":
		orderBy = models.SearchOrderBySize
	case "downloadtimes":
		orderBy = models.SearchOrderByDownloadTimes
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse
	case "feweststars":
		orderBy = models.SearchOrderByStars
	case "mostforks":
		orderBy = models.SearchOrderByForksReverse
	case "fewestforks":
		orderBy = models.SearchOrderByForks
	case "mostusecount":
		orderBy = models.SearchOrderByUseCountReverse
	case "fewestusecount":
		orderBy = models.SearchOrderByUseCount
	default:
		ctx.Data["SortType"] = "recentupdate"
		orderBy = models.SearchOrderByRecentUpdated
	}

	keyword := strings.Trim(ctx.Query("q"), " ")

	datasets, count, err = models.SearchDataset(&models.SearchDatasetOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: setting.UI.ExplorePagingNum,
		},
		Keyword:        keyword,
		RecommendOnly:  ctx.QueryBool("recommend"),
		CloudBrainType: -1,
		SearchOrderBy:  orderBy,
	})
	if err != nil {
		ctx.ServerError("SearchDataset", err)
		return
	}

	ctx.Data["Keyword"] = keyword
	ctx.Data["Total"] = count
	ctx.Data["Datasets"] = datasets
	ctx.Data["Recommend"] = ctx.QueryBool("recommend")
	pager := context.NewPagination(int(count), setting.UI.ExplorePagingNum, page, 5)
	pager.SetDefaultParams(ctx)
	ctx.Data["Page"] = pager

	ctx.HTML(200, tplDatasets)
}

func DatasetAction(ctx *context.Context) {
	var err error
	datasetId, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)
	switch ctx.Params(":action") {

	case "recommend":
		err = models.RecommendDataset(datasetId, true)
	case "unrecommend":
		err = models.RecommendDataset(datasetId, false)
	}
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.star_fail", ctx.Params(":action"))))
	} else {
		d, _ := models.GetDatasetByID(datasetId)
		notification.NotifyDatasetRecommend(ctx.User, d, ctx.Params(":action"))
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}
}

func DeleteDataset(ctx *context.Context) {
	dataset, err := models.GetDatasetByID(ctx.QueryInt64("id"))
	if err != nil {
		ctx.ServerError("GetDatasetByID", err)
		return
	}

	if err := models.DeleteDataset(dataset.ID, ctx.User.ID); err != nil {
		ctx.ServerError("DeleteDataset", err)
		return
	}
	log.Trace("DeleteDataset deleted: %s", dataset.ID)

	ctx.Flash.Success(ctx.Tr("dataset.settings.deletion_success"))
	ctx.JSON(200, map[string]interface{}{
		"redirect": setting.AppSubURL + "/admin/datasets?page=" + ctx.Query("page") + "&sort=" + ctx.Query("sort"),
	})
}
