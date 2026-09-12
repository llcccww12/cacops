package point

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/reward/point/account"
	"net/http"
)

func SearchPointAccount(ctx *context.Context) {
	q := ctx.Query("q")
	page := ctx.QueryInt("page")
	resopnse, err := account.SearchPointAccount(models.SearchPointAccountOpts{ListOptions: models.ListOptions{Page: page, PageSize: 20}, Keyword: q})
	if err != nil {
		log.Error("SearchPointAccount error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(resopnse))
	return
}
