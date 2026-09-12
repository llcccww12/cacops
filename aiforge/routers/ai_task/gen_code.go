package ai_task

import (
	"code.gitea.io/gitea/models"
	"net/http"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/task"
)

func GenerateModelUploadCode(ctx *context.Context) {

	pretrainModelName := ctx.Query("model_name")

	code, cli := task.GenerateModelUploadCode(pretrainModelName, ctx.Repo.Repository.OwnerName, ctx.Repo.Repository.Name)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]string{"code": code, "cli": cli}))
}

func GenerateModelFileUploadCode(ctx *context.Context) {

	pretrainModelName := ctx.Query("model_name")

	code, cli := task.GenerateModelFileUploadCode(pretrainModelName, ctx.Repo.Repository.OwnerName, ctx.Repo.Repository.Name)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]string{"code": code, "cli": cli}))
}

func GenerateModelDownloadCode(ctx *context.Context) {
	file_name := ctx.Query("file_name")
	aimodel := ctx.AccessContext.Aimodel
	owner, err := models.GetUserByID(aimodel.OwnerID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	code, cli := task.GenerateSDKDownloadCode(file_name, owner.Name, aimodel.Name, "model")
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]string{"code": code, "cli": cli}))
}

func GenerateDatasetDownloadCode(ctx *context.Context) {
	file_name := ctx.Query("file_name")
	dataset := ctx.AccessContext.Dataset
	owner, err := models.GetUserByID(dataset.OwnerID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	code, cli := task.GenerateSDKDownloadCode(file_name, owner.Name, dataset.Name, "dataset")
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]string{"code": code, "cli": cli}))
}

func GenerateDatasetUploadCode(ctx *context.Context) {

	dataset_file_name := ctx.Query("dataset_file_name")

	code, cli := task.GenerateDatasetUploadCode(dataset_file_name, ctx.Repo.Repository.OwnerName, ctx.Repo.Repository.Name)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]string{"code": code, "cli": cli}))
}
