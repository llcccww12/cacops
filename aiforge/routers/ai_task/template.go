package ai_task

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"net/http"
)

const (
	tplAiTaskTmplList base.TplName = "aitasktmpl/list"
	tplAiTaskTmplListMy base.TplName = "aitasktmpl/list_my"
	tplAiTaskTmplCreate base.TplName = "aitasktmpl/create"
	tplAiTaskTmplDetail base.TplName = "aitasktmpl/detail"
	tplAiTaskTmplEdit base.TplName = "aitasktmpl/edit"
)

func CanCreateAITask(ctx *context.Context) {
	var canCreate = ctx.IsUserRepoWriter([]models.UnitType{models.UnitTypeCloudBrain}) || ctx.IsUserRepoAdmin() || ctx.IsUserSiteAdmin()
	result := map[string]interface{}{"CanCreate": canCreate}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetTemplateEditAddress(ctx *context.Context) {
	if !ctx.Repo.CanEnableEditor() {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.REPO_CAN_NOT_CREATE_AI_TEMPLATE, ctx))
		return
	}
	_, err := ctx.Repo.Commit.GetTreeEntryByPath(setting.AI_TASK_TEMPLATE_FILE_NAME)
	isTemplateFileExists := true
	if err != nil {
		if git.IsErrNotExist(err) {
			isTemplateFileExists = false
			err = nil
		} else {
			ctx.ServerError("", err)
			return
		}
	}

	var url string
	if isTemplateFileExists {
		url = ctx.Repo.Repository.GetAITemplateEditURL(ctx.Repo.BranchName)
	} else {
		url = ctx.Repo.Repository.GetAITemplateNewURL(ctx.Repo.BranchName)
	}

	result := map[string]interface{}{
		"edit_url":  url,
		"file_name": setting.AI_TASK_TEMPLATE_FILE_NAME,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetAiTaskTmplListTpl(ctx *context.Context) {
	ctx.HTML(200, tplAiTaskTmplList)
}

func GetAiTaskTmplListMyTpl(ctx *context.Context) {
	ctx.HTML(200, tplAiTaskTmplListMy)
}

func GetAiTaskTmplCreateTpl(ctx *context.Context) {
	ctx.HTML(200, tplAiTaskTmplCreate)
}

func GetAiTaskTmplDetailTpl(ctx *context.Context) {
	ctx.HTML(200, tplAiTaskTmplDetail)
}

func GetAiTaskTmplEditTpl(ctx *context.Context) {
	ctx.HTML(200, tplAiTaskTmplEdit)
}
