package point

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/reward/limiter"
	"net/http"
)

func GetSingleDailyPointLimitConfig(ctx *context.Context) {
	r, err := limiter.GetSingleDailyPointLimitConfig()
	if err != nil {
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	resultMap := make(map[string]interface{}, 0)
	if r == nil {
		resultMap["LimitNum"] = ""
	} else {
		resultMap["LimitNum"] = r.LimitNum
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(resultMap))
}

func SetSingleDailyPointLimitConfig(ctx *context.Context, config models.LimitConfigVO) {
	err := limiter.SetSingleDailyPointLimitConfig(config.LimitNum, ctx.User)
	if err != nil {
		log.Error("Set single daily point limit config error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func DeletePointLimitConfig(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	err := limiter.DeleteLimitConfig(id, ctx.User)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}
