package resources

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"net/http"
)

func GetAccCardList(ctx *context.Context) {
	specIds, err := models.GetAllSpecIds()
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return

	}
	list, err := resource.GetAccCardListV2(specIds)
	if err != nil {
		log.Error("GetAccCardList error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}

	m := map[string]interface{}{"list": list}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetAvailableAICenterList(ctx *context.Context) {
	specIds, err := models.GetAllSpecIds()
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return

	}
	list, err := resource.GetAvailableAICenterV2(specIds)
	if err != nil {
		log.Error("GetAvailableAICenterList error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	for i := 0; i < len(list); i++ {
		list[i].Tr(ctx.Language())
	}

	m := map[string]interface{}{"list": list}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}
