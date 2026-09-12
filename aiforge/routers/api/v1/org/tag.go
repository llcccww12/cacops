package org

import (
	"net/http"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/routers/response"
)

func GetAvailableOrgOfficialDatasetRegistry(ctx *context.APIContext) {
	isOwner, err := ctx.Org.Organization.IsOwnedBy(ctx.User.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}
	if !isOwner && !ctx.User.IsAdmin {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}

	r, err := models.GetAvailableOrgOfficialDatasetRegistry(ctx.Org.Organization.ID)
	if err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	ctx.JSON(200, response.OuterSuccessWithData(map[string]interface{}{
		"datasets": r,
	}))

}

func SubmitTagsToDatasetRegistry(ctx *context.APIContext, form auth.SubmitDatasetOfTagForm) {
	isOwner, err := ctx.Org.Organization.IsOwnedBy(ctx.User.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}
	if !isOwner && !ctx.User.IsAdmin {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}

	datasetid := form.Datasetids
	datasetids := make([]string, 0)
	if datasetid != "" {
		datasetids = strings.Split(datasetid, ",")
	}
	if len(datasetids) > 9 {
		ctx.JSON(200, response.OuterTrBizError(response.BuildDefaultBizError("Submit tags cannot over 9"), ctx.Locale))
		return
	}

	err = models.UpdateTagDatasetRegistryByID(ctx.Org.Organization.ID, datasetids)
	if err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	ctx.JSON(200, response.OuterSuccess())
}

func GetOrgOfficialDatasetCardList(ctx *context.APIContext) {
	datasetCards, err := models.GetOrgOfficialDatasetRegistryForCard(ctx.Org.Organization.ID)
	if err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	datasetInfos := make([]*entity.DatasetInfo, len(datasetCards))
	for i, dataset := range datasetCards {
		datasetInfos[i] = entity.BuildDatasetInfo(dataset)
	}
	ctx.JSON(200, response.OuterSuccessWithData(map[string]interface{}{
		"datasets": datasetInfos,
	}))
}

func GetAvailableOrgOfficialAimodel(ctx *context.APIContext) {
	isOwner, err := ctx.Org.Organization.IsOwnedBy(ctx.User.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}
	if !isOwner && !ctx.User.IsAdmin {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}

	r, err := models.GetAvailableOrgOfficialAimodel(ctx.Org.Organization.ID)
	if err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	ctx.JSON(200, response.OuterSuccessWithData(map[string]interface{}{
		"aimodels": r,
	}))

}

func SubmitTagsToAimodel(ctx *context.APIContext, form auth.SubmitModelOfTagForm) {
	isOwner, err := ctx.Org.Organization.IsOwnedBy(ctx.User.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}
	if !isOwner && !ctx.User.IsAdmin {
		ctx.JSON(http.StatusForbidden, "no access")
		return
	}

	aimodelid := form.Modelids
	aimodelids := make([]string, 0)
	if aimodelid != "" {
		aimodelids = strings.Split(aimodelid, ",")
	}
	if len(aimodelids) > 9 {
		ctx.JSON(200, response.OuterTrBizError(response.BuildDefaultBizError("Submit tags cannot over 9"), ctx.Locale))
		return
	}

	err = models.UpdateTagModelByID(ctx.Org.Organization.ID, aimodelids)
	if err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	ctx.JSON(200, response.OuterSuccess())
}

func GetOrgOfficialAimodelCardList(ctx *context.APIContext) {
	aimodelCards, err := models.GetOrgOfficialAimodelForCard(ctx.Org.Organization.ID)
	if err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	aimodelInfos := make([]*entity.AimodelInfo, len(aimodelCards))
	for i, aimodel := range aimodelCards {
		aimodelInfos[i] = entity.BuildAimodelInfo(aimodel)
	}
	ctx.JSON(200, response.OuterSuccessWithData(map[string]interface{}{
		"aimodels": aimodelInfos,
	}))
}
