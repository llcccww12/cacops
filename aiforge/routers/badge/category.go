package badge

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/badge"
	"errors"
	"net/http"
)

func GetBadgeCategoryList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := 50
	n, r, err := badge.GetBadgeCategoryList(models.ListOptions{Page: page, PageSize: pageSize})
	if err != nil {
		log.Error("GetCategoryList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	m := make(map[string]interface{})
	m["List"] = r
	m["Total"] = n
	m["PageSize"] = pageSize
	ctx.JSON(http.StatusOK, response.SuccessWithData(m))
}

func OperateBadgeCategory(ctx *context.Context, category models.BadgeCategory4Show) {
	action := ctx.Params(":action")

	var err *response.BizError
	switch action {
	case "edit":
		err = badge.EditBadgeCategory(category, ctx.User)
	case "new":
		err = badge.AddBadgeCategory(category, ctx.User)
	case "del":
		err = badge.DelBadgeCategory(category.ID, ctx.User)
	default:
		err = response.NewBizError(errors.New("action type error"))
	}

	if err != nil {
		log.Error("OperateBadgeCategory error ,%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}
