package repo

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

func GetResourceSpec(ctx *context.APIContext) {
	jobType := ctx.Query("jobType")
	computeResource := ctx.Query("compute")
	cluster := ctx.Query("cluster")
	aiCenterCode := ctx.Query("center")
	if jobType == "" || computeResource == "" || cluster == "" {
		log.Info("GetResourceSpec api.param error")
		ctx.JSON(200, response.OuterBizError(response.PARAM_ERROR))
		return
	}
	specs, err := resource.FindAvailableSpecs4Show(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobType(jobType),
		ComputeResource: computeResource,
		Cluster:         cluster,
		AiCenterCode:    aiCenterCode,
	})
	if err != nil {
		log.Error("GetResourceSpec api error. %v", err)
		ctx.JSON(200, response.OuterServerError(err.Error()))
		return
	}

	specMap := make(map[string]interface{}, 0)
	specMap["specs"] = specs
	ctx.JSON(200, response.OuterSuccessWithData(specMap))
}
