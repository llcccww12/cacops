package ai_task

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/task"
)

// CreateAgent 创建智能体
// @title CreateAgent
// @description 创建智能体，可利用在智能体广场
func CreateAgent(ctx *context.Context, form task.CreateAgentReq) {
	log.Info("CreateAgent begin. form=%+v", form)

	// 创建智能体
	err := task.CreateAgent(&form)
	if err != nil {
		log.Error("CreateAgent failed: %v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
}

// DeleteAgent 删除智能体
// @title DeleteAgent
// @description 删除智能体（仅管理员可操作）
func DeleteAgent(ctx *context.Context) {
	id := ctx.ParamsInt64(":id")
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	// 获取智能体进行权限校验
	agent, err := models.GetAgentByID(id)
	if err != nil || agent == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("agent not found"), ctx))
		return
	}

	bizErr := task.DeleteAgent(id)
	if bizErr != nil {
		log.Error("DeleteAgent failed: %v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
}

// UpdateAgent 更新智能体
// @title UpdateAgent
// @description 更新智能体配置
func UpdateAgent(ctx *context.Context, form task.UpdateAgentReq) {
	id := ctx.ParamsInt64(":id")
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	// 获取智能体进行权限校验
	agent, err := models.GetAgentByID(id)
	if err != nil || agent == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("agent not found"), ctx))
		return
	}

	// 权限校验：通过 user_agent 关联表校验
	if !task.CheckAgentAccess(ctx.User.ID, id, ctx.User.IsAdmin) {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx))
		return
	}

	_, bizErr := task.UpdateAgent(id, &form)
	if bizErr != nil {
		log.Error("UpdateAgent failed: %v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
}

// GetAgentList 获取智能体列表（直接查询 agent 表，不涉及 user_agent 关联）
func GetAgentList(ctx *context.Context) {
	q := ctx.Query("q")
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("page_size")

	result, err := task.GetAgentList(q, page, pageSize)
	if err != nil {
		log.Error("GetAgentList failed: %v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

// UpdateAgentStatus 更新智能体状态
// @title UpdateAgentStatus
// @description 更新智能体状态（发布/下架）
func UpdateAgentStatus(ctx *context.Context, form task.UpdateAgentStatusReq) {
	id := ctx.ParamsInt64(":id")
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	// 权限校验：仅管理员可更新
	if !ctx.User.IsAdmin {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx))
		return
	}

	// 获取智能体进行权限校验
	agent, err := models.GetAgentByID(id)
	if err != nil || agent == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("agent not found"), ctx))
		return
	}

	// 权限校验：通过 user_agent 关联表校验
	if !task.CheckAgentAccess(ctx.User.ID, id, ctx.User.IsAdmin) {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx))
		return
	}

	_, bizErr := task.UpdateAgentStatus(id, &form)
	if bizErr != nil {
		log.Error("UpdateAgentStatus failed: %v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(""))
}

// RunAgent 启用智能体
// @title RunAgent
// @description 启用智能体（在 user_agent 表创建关联记录，并创建 AI 任务）
func RunAgent(ctx *context.Context) {
	id := ctx.ParamsInt64(":id")
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	// 获取智能体进行权限校验
	agent, err := models.GetAgentByID(id)
	if err != nil || agent == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("agent not found"), ctx))
		return
	}

	code := 0
	result, bizErr := task.RunAgent(id, ctx.User.ID)
	if bizErr != nil {
		log.Error("RunAgent failed: %v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx))
		return
	}

	if bizErr != nil {
		switch bizErr.Code {
		case response.MULTI_TASK.Code:
			code = 2
		default:
			code = 1
		}
		msg := bizErr.DefaultMsg
		if bizErr.TrCode != "" {
			if bizErr.TrParams == nil || len(bizErr.TrParams) == 0 {
				msg = ctx.Tr(bizErr.TrCode)
			} else {
				msg = ctx.Tr(bizErr.TrCode, bizErr.TrParams...)
			}
		}
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: code, Message: msg})
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

// StopAgent 停止智能体（根据 cloudbrain_id 停止云脑任务）
// @title StopAgent
// @description 停止智能体（根据 cloudbrain_id 停止云脑任务）
func StopAgent(ctx *context.Context) {
	cloudbrainId := ctx.ParamsInt64(":id")
	if cloudbrainId <= 0 {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	// 获取 Cloudbrain 任务进行权限校验
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil || cloudbrain == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("cloudbrain not found"), ctx))
		return
	}

	// 权限校验：只允许任务创建者
	if cloudbrain.UserID != ctx.User.ID {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx))
		return
	}

	bizErr := task.StopAgent(cloudbrainId, ctx.User.ID)
	if bizErr != nil {
		log.Error("StopAgent failed: %v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

// GetUserAgentList 获取用户智能体列表（带 cloudbrain job_status 运行状态）
// @title GetUserAgentList
// @description 获取当前用户的智能体列表，包含 cloudbrain 运行状态
func GetUserAgentList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("page_size")

	result, err := task.GetUserAgentList(ctx.User.ID, page, pageSize)
	if err != nil {
		log.Error("GetUserAgentList failed: %v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

// DeleteUserAgentByCloudBrainID 删除用户智能体关联
// @title DeleteUserAgentByCloudBrainID
// @description 根据 cloudbrain_id 删除用户与智能体的关联记录
func DeleteUserAgentByCloudBrainID(ctx *context.Context) {
	id := ctx.ParamsInt64(":id")
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	// 根据 cloudbrain_id 和用户 id 删除关联记录
	err := models.DeleteUserAgentByCloudBrainID(ctx.User.ID, id)
	if err != nil {
		log.Error("DeleteUserAgentByCloudBrainID failed: userID=%d, id=%d, err=%v", ctx.User.ID, id, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("failed to delete user agent relation"), ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(""))
}

// GetUserAgentRecordList 获取用户智能体使用历史记录（管理员接口）
// @title GetUserAgentRecordList
// @description 获取用户智能体使用历史记录，支持按用户 ID 筛选
func GetUserAgentRecordList(ctx *context.Context) {
	userID := ctx.QueryInt64("user_id")
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("page_size")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	records, total, err := models.GetUserAgentRecords(userID, page, pageSize)
	if err != nil {
		log.Error("GetUserAgentRecordList failed: userID=%d, err=%v", userID, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.BuildDefaultBizError("failed to get user agent records"), ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{
		"records":   records,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}))
}
