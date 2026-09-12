package modelscope

import (
	"net/http"
	"strconv"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	ms "code.gitea.io/gitea/services/modelscope"
)

func ListModels(ctx *context.APIContext) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	search := ctx.Query("q")
	if search == "" {
		search = ctx.Query("search")
	}
	sort := ctx.Query("order_by")
	if sort == "" {
		sort = ctx.Query("sort")
	}

	models, total, err := ms.ListModels(search, sort, page, pageSize)
	if err != nil {
		log.Error("ListModels from modelscope failed: %v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"models":    models,
		"total":     total,
		"page":      pageOrDefault(page),
		"page_size": pageSizeOrDefault(pageSize),
		"source":    "modelscope",
	}))
}

func ListDatasets(ctx *context.APIContext) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	search := ctx.Query("q")
	if search == "" {
		search = ctx.Query("search")
	}
	sort := ctx.Query("order_by")
	if sort == "" {
		sort = ctx.Query("sort")
	}

	datasets, total, err := ms.ListDatasets(search, sort, page, pageSize)
	if err != nil {
		log.Error("ListDatasets from modelscope failed: %v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(entity.SearchDatasetRes{
		Datasets: datasets,
		Total:    total,
		Page:     pageOrDefault(page),
		PageSize: pageSizeOrDefault(pageSize),
	}))
}

func pageOrDefault(page int) int {
	if page <= 0 {
		return 1
	}
	return page
}

func pageSizeOrDefault(pageSize int) int {
	if pageSize <= 0 {
		return 30
	}
	return pageSize
}
