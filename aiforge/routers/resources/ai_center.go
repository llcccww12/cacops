package resources

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"net/http"
)

func GetActiveAICenterList(ctx *context.Context) {
	computeSource := ctx.Query("computeSource")
	accCardType := ctx.Query("accCardType")
	list, err := resource.GetActiveAICenterList(models.ActiveAICenterReq{
		ComputeSource: computeSource,
		AccCardType:   accCardType,
	})
	if err != nil {
		log.Error("GetAvailableAICenterList error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	for i := 0; i < len(list); i++ {
		list[i].TotalCardsUsageNum = 0
		(&list[i]).Tr(ctx.Language())
	}
	m := map[string]interface{}{"list": list}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetActiveCardInfoList(ctx *context.Context) {
	list, err := resource.GetActiveAccCardInfo()
	if err != nil {
		log.Error("GetAvailableAICenterList error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	m := map[string]interface{}{"list": list}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}
