package ai_task_template

import (
	"net/http"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/template_service"

	"code.gitea.io/gitea/entity"
)

func CreateAITaskTemplate(ctx *context.Context, req entity.CreateAITaskTemplateReq) {
	req.OwnerID = ctx.User.ID
	id, err := template_service.CreateAITaskTemplate(req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"id": id,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func EditAITaskTemplate(ctx *context.Context, req entity.CreateAITaskTemplateReq) {
	id := ctx.Query("id")
	if id == "" {
		log.Error("EditAITaskTemplate failed, templateId is empty")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}
	template, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s err=%v", id, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	if !template.CanEdit(ctx.User) {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
		return
	}

	bizErr := template_service.EditAITaskTemplate(req, id)
	if bizErr != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx.Locale))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func GetAITaskTemplate(ctx *context.Context) {
	id := ctx.Query("id")
	template, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s", id)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	var userId int64
	if ctx.User != nil {
		userId = ctx.User.ID
	}

	if !template.CanRead(ctx.User) {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
		return
	}

	template.LoadAtrribute(userId)

	templateInfo := entity.BuildAITaskTemplateDetaiInfo(template, ctx.User)

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(templateInfo))
}

func parseSearchTemplateReq(ctx *context.APIContext) models.SearchAITaskTemplateReq {
	jobType := ctx.Query("job_type")
	computeSource := ctx.Query("compute_source")
	datasetID := ctx.Query("dataset_id")
	modelID := ctx.Query("model_id")
	repoID := ctx.Query("repo_id")
	ownerName := ctx.Query("owner_name")
	keyword := ctx.Query("q")
	recommend := ctx.Query("recommend")
	tagsStr := ctx.Query("tags")
	tags := make([]string, 0)
	if tagsStr != "" {
		tags = strings.Split(tagsStr, "|")
	}
	orderBy := ctx.Query("order_by")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("page_size")
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 12
	}

	req := models.SearchAITaskTemplateReq{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		JobType:       jobType,
		ComputeSource: computeSource,
		DatasetID:     datasetID,
		Recommend:     recommend,
		ModelID:       modelID,
		RepoID:        repoID,
		OwnerName:     ownerName,
		Keyword:       keyword,
		Tags:          tags,
		OrderBy:       orderBy,
		Operator:      ctx.User,
	}
	return req
}

func DelAITaskTemplate(ctx *context.Context) {
	id := ctx.Query("id")
	template, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s err=%v", id, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	if !template.CanDelete(ctx.User) {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
		return
	}

	err = template_service.DelAITaskTemplate(id)
	if err != nil {
		log.Error("DelAITaskTemplate failed, templateId=%s err=%v", id, err)
		ctx.JSON(http.StatusOK, response.NewBizError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

// 所有公开
func SearchAllPublicTemplates(ctx *context.APIContext) {
	req := parseSearchTemplateReq(ctx)
	req.Visibility = models.VisibilityPublic
	res, total, err := template_service.SearchAITaskTemplates(req)
	if err != nil {
		log.Error("SearchAllPublicTemplates failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"Templates": res,
		"Total":     total,
		"Page":      req.Page,
		"PageSize":  req.PageSize,
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 所有模板
func SearchAllTemplates(ctx *context.APIContext) {
	req := parseSearchTemplateReq(ctx)
	req.Visibility = models.VisibilityAll
	res, total, err := template_service.SearchAITaskTemplates(req)
	if err != nil {
		log.Error("SearchAllPublicTemplates failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"Templates": res,
		"Total":     total,
		"Page":      req.Page,
		"PageSize":  req.PageSize,
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 我收藏的
func SearchCollectedTemplates(ctx *context.APIContext) {
	req := parseSearchTemplateReq(ctx)
	req.Scope = models.ScopeCollected
	res, total, err := template_service.SearchAITaskTemplates(req)
	if err != nil {
		log.Error("SearchAllPublicTemplates failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"Templates": res,
		"Total":     total,
		"Page":      req.Page,
		"PageSize":  req.PageSize,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 我创建的
func SearchCreatedTemplates(ctx *context.APIContext) {
	req := parseSearchTemplateReq(ctx)
	req.Scope = models.ScopeCreated
	res, total, err := template_service.SearchAITaskTemplates(req)
	if err != nil {
		log.Error("SearchAllPublicTemplates failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"Templates": res,
		"Total":     total,
		"Page":      req.Page,
		"PageSize":  req.PageSize,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 我可访问的
func SearchAccessibleTemplates(ctx *context.APIContext) {
	req := parseSearchTemplateReq(ctx)
	req.Visibility = models.VisibilityPublic
	ownerUser, _ := models.GetUserByName(req.OwnerName)
	if ctx.User.ID == ownerUser.ID || ctx.IsUserSiteAdmin() {
		req.Visibility = models.VisibilityAll
	}
	res, total, err := template_service.SearchAITaskTemplates(req)
	if err != nil {
		log.Error("SearchAllPublicTemplates failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"Templates": res,
		"Total":     total,
		"Page":      req.Page,
		"PageSize":  req.PageSize,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

// 管理后台
func SearchAllTemplates4Admin(ctx *context.APIContext) {
	req := parseSearchTemplateReq(ctx)
	req.Scope = models.ScopeAll
	res, total, err := template_service.SearchAITaskTemplates(req)
	if err != nil {
		log.Error("SearchAllPublicTemplates failed, req=%+v err=%v", req, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx.Locale))
		return
	}
	m := map[string]interface{}{
		"Templates": res,
		"Total":     total,
		"Page":      req.Page,
		"PageSize":  req.PageSize,
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func CollectAITaskTemplate(ctx *context.Context) {
	id := ctx.Query("id")
	template, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s", id)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	if !template.CanRead(ctx.User) {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
		return
	}
	err = template_service.CollectAITaskTemplate(ctx.User.ID, id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func UnCollectAITaskTemplate(ctx *context.Context) {
	id := ctx.Query("id")
	_, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s", id)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	if ctx.User == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
		return
	}

	err = template_service.UnCollectAITaskTemplate(ctx.User.ID, id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func RecommendAITaskTemplate(ctx *context.Context) {
	id := ctx.Query("id")
	t, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s", id)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	if t.IsPrivate {
		ctx.JSON(http.StatusOK, response.OuterServerError("Private template can not be recommended"))
		return
	}
	err = template_service.RecommendAITaskTemplate(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func UnRecommendAITaskTemplate(ctx *context.Context) {
	id := ctx.Query("id")
	_, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("GetAITaskTemplate failed, templateId=%s", id)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	err = template_service.UnRecommendAITaskTemplate(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func HandleHistoricalAITaskTemplate(ctx *context.Context) {
	forceFull := ctx.QueryBool("force_full")
	count, err := template_service.HandleHistoricalAITaskTemplate(forceFull)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{"count": count}))
}
