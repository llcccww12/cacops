package task

import (
	"errors"
	"fmt"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

// MaxUserAgentCount 用户最多可拥有的智能体数量
const MaxUserAgentCount = 10

// AgentResourceSpec 智能体资源规格（用于执行记录）
type AgentResourceSpec struct {
	ID              int64
	AccCardType     string
	AccCardsNum     int
	CpuCores        int
	MemGiB          float32
	GPUMemGiB       float32
	ShareMemGiB     float32
	ComputeResource string
	UnitPrice       float64
}

// AgentExecutionParam 智能体执行记录参数
type AgentExecutionParam struct {
	AgentID    int64
	UserID     int64
	TemplateID string
	CloudBrain int64
	Action     string // RUN / STOP
	Success    bool
	Message    string
	Spec       *AgentResourceSpec // 资源规格（可为 nil）
}

// CreateAgentReq 创建智能体请求
type CreateAgentReq struct {
	AgentType     string   `json:"agent_type"`      // 智能体类型：OPENCLAW / OPENCODE
	TemplateID    string   `json:"template_id"`     // 关联的模板 ID（ai_task_template.id）
	Status        int8     `json:"status"`          // 智能体状态：0 草稿 1 已发布 2 已下架
	Name          string   `json:"name"`            // 智能体名称
	AvatarUrl     string   `json:"avatar_url"`      // 智能体头像框
	BackGroundUrl string   `json:"back_ground_url"` // 智能体背景图 URL
	DetailURL     string   `json:"detail_url"`      // 详情页跳转链接 URL
	Tag           []string `json:"tag"`             // 标签
	CreatorId     int64    `json:"creator_id"`      // 创建智能体的用户 ID
	Description   string   `json:"description"`     // 智能体描述
}

// CreateAgent 创建智能体（不建立用户关联，用户关联需通过绑定接口单独建立）
// 一个模板只能构建一个智能体？
func CreateAgent(req *CreateAgentReq) *response.BizError {
	// 参数验证
	if req.TemplateID == "" {
		return response.NewBizError(errors.New("template_id is required"))
	}
	if req.AgentType == "" {
		return response.NewBizError(errors.New("agent_type is required"))
	}

	// 验证智能体类型
	agentType := models.AgentType(req.AgentType)

	// 创建智能体
	agent := &models.Agent{
		AgentType:     agentType,
		TemplateID:    req.TemplateID,
		Status:        models.AgentStatus(req.Status),
		Name:          req.Name,
		AvatarUrl:     req.AvatarUrl,
		BackGroundUrl: req.BackGroundUrl,
		DetailURL:     req.DetailURL,
		Tag:           req.Tag,
		CreatorId:     req.CreatorId,
		Description:   req.Description,
	}

	if err := models.CreateAgent(agent); err != nil {
		log.Error("CreateAgent create agent err. req=%+v, err=%v", req, err)
		return response.NewBizError(errors.New("failed to create agent"))
	}

	return nil
}

// UpdateAgentReq 更新智能体请求
type UpdateAgentReq struct {
	Name          string   `json:"name"`            // 智能体名称
	Description   string   `json:"description"`     // 智能体描述
	TemplateID    string   `json:"template_id"`     // 模板 ID
	BackGroundUrl string   `json:"back_ground_url"` // 智能体背景图
	AvatarUrl     string   `json:"avatar_url"`      // 智能体头像框
	DetailURL     string   `json:"detail_url"`      // 详情页跳转链接
	Tag           []string `json:"tag"`             // 标签
}

// UpdateAgentStatusReq 更新智能体状态请求
type UpdateAgentStatusReq struct {
	Status int8 `json:"status"` // 智能体状态：0 草稿 1 已发布 2 已下架
}

// GetAgentList 获取智能体列表（直接查询 agent 表，不涉及 user_agent 关联）
func GetAgentList(q string, page, pageSize int) (*models.AgentListRes, *response.BizError) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 查询 agent 表（model 层已处理模板 spec 数据和单价关联）
	listRes, err := models.GetAgentList(q, page, pageSize)
	if err != nil {
		log.Error("GetAgentList query agents err. q=%s, err=%v", q, err)
		return nil, response.NewBizError(err)
	}

	// 收集所有 template_id
	templateIDs := make([]string, 0, len(listRes.Agents))
	for _, agent := range listRes.Agents {
		if agent.TemplateID != "" {
			templateIDs = append(templateIDs, agent.TemplateID)
		}
	}

	// 批量查询模板数据（用于补充 Description 和 Tag）
	templateMap, err := models.GetAITaskTemplatesByIDs(templateIDs)
	if err != nil {
		log.Error("GetAgentList query templates err. templateIDs=%v, err=%v", templateIDs, err)
		// 模板查询失败不影响 agent 列表返回
		return listRes, nil
	}

	// 如果 agent 的 Description 或 Tag 为空，使用模板的数据；同时补充 ComputeResource
	for i, agent := range listRes.Agents {
		if template, ok := templateMap[agent.TemplateID]; ok {
			if agent.Description == "" && template.Description != "" {
				listRes.Agents[i].Description = template.Description
			}
			if len(agent.Tag) == 0 && len(template.Tags) > 0 {
				listRes.Agents[i].Tag = template.Tags
			}
			if agent.Spec != nil && agent.Spec.ComputeResource == "" {
				listRes.Agents[i].Spec.ComputeResource = template.ComputeSource
			}
		}
	}

	return listRes, nil
}

// GetUserAgentList 获取用户智能体列表（通过 user_agent 关联表，带 cloudbrain job_status）
func GetUserAgentList(userID int64, page, pageSize int) (*models.UserAgentListRes, *response.BizError) {
	if userID <= 0 {
		return nil, response.NewBizError(errors.New("user id is required"))
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 查询 user_agent 列表（带 cloudbrain job_status）
	listRes, err := models.GetUserAgentList(userID, page, pageSize)
	if err != nil {
		log.Error("GetUserAgentList query user_agents err. userID=%d, err=%v", userID, err)
		return nil, response.NewBizError(err)
	}

	return listRes, nil
}

// UpdateAgent 更新智能体
func UpdateAgent(id int64, req *UpdateAgentReq) (*models.Agent, *response.BizError) {
	if id <= 0 {
		return nil, response.NewBizError(errors.New("agent id is required"))
	}

	// 获取智能体
	agent, err := models.GetAgentByID(id)
	if err != nil || agent == nil {
		return nil, response.NewBizError(errors.New("agent not found"))
	}

	// 更新字段
	if req.Name != "" {
		agent.Name = req.Name
	}
	if req.Description != "" {
		agent.Description = req.Description
	}
	if req.TemplateID != "" {
		agent.TemplateID = req.TemplateID
	}
	if req.BackGroundUrl != "" {
		agent.BackGroundUrl = req.BackGroundUrl
	}
	if req.AvatarUrl != "" {
		agent.AvatarUrl = req.AvatarUrl
	}
	if req.DetailURL != "" {
		agent.DetailURL = req.DetailURL
	}
	if len(req.Tag) > 0 {
		agent.Tag = req.Tag
	}

	if err := models.UpdateAgent(agent); err != nil {
		log.Error("UpdateAgent update agent err. id=%d, err=%v", id, err)
		return nil, response.NewBizError(errors.New("failed to update agent"))
	}

	return agent, nil
}

// UpdateAgentStatus 更新智能体状态
func UpdateAgentStatus(id int64, req *UpdateAgentStatusReq) (*models.Agent, *response.BizError) {
	if id <= 0 {
		return nil, response.NewBizError(errors.New("agent id is required"))
	}

	status := models.AgentStatus(req.Status)
	if !status.IsValid() {
		return nil, response.NewBizError(errors.New("invalid status, must be 0 (draft), 1 (published), or 2 (unpublished)"))
	}

	// 获取智能体
	agent, err := models.GetAgentByID(id)
	if err != nil || agent == nil {
		return nil, response.NewBizError(errors.New("agent not found"))
	}

	// 更新状态
	if err := models.UpdateAgentStatus(id, status); err != nil {
		log.Error("UpdateAgentStatus update status err. id=%d, err=%v", id, err)
		return nil, response.NewBizError(errors.New("failed to update agent status"))
	}

	agent.Status = status
	return agent, nil
}

// GetAgentByID 根据 ID 获取智能体
func GetAgentByID(id int64) (*models.Agent, *response.BizError) {
	agent, err := models.GetAgentByID(id)
	if err != nil {
		log.Error("GetAgentByID err. id=%d, err=%v", id, err)
		return nil, response.NewBizError(errors.New("failed to get agent"))
	}
	if agent == nil {
		return nil, response.NewBizError(errors.New("agent not found"))
	}
	return agent, nil
}

// DeleteAgent 删除智能体
func DeleteAgent(agentId int64) *response.BizError {
	if agentId <= 0 {
		return response.NewBizError(errors.New("task id is required"))
	}

	if err := models.DeleteAgent(agentId); err != nil {
		log.Error("DeleteAgent delete agent err. agentId=%d, err=%v", agentId, err)
		return response.NewBizError(errors.New("failed to delete agent"))
	}

	return nil
}

// CheckAgentAccess 检查用户是否有权限访问智能体
func CheckAgentAccess(userID int64, agentID int64, isAdmin bool) bool {
	if isAdmin {
		return true // 管理员可以访问所有智能体
	}
	return models.CheckUserAgentRelation(userID, agentID)
}

// recordAgentExecutionAsync 异步记录智能体执行日志（含 panic 兜底）
func recordAgentExecutionAsync(p AgentExecutionParam) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("recordAgentExecution PANIC. agentId=%d, userId=%d, action=%s, err=%v\nStack: %s",
					p.AgentID, p.UserID, p.Action, err, log.Stack(2))
			}
		}()
		recordAgentExecution(p)
	}()
}

func recordAgentExecution(p AgentExecutionParam) {
	record := &models.AgentExecutionRecord{
		AgentID:      p.AgentID,
		UserID:       p.UserID,
		TemplateID:   p.TemplateID,
		CloudBrainID: p.CloudBrain,
		Action:       p.Action,
		Success:      p.Success,
		Message:      p.Message,
	}
	if p.Spec != nil {
		record.ResourceSpecID = p.Spec.ID
		record.AccCardType = p.Spec.AccCardType
		record.AccCardsNum = p.Spec.AccCardsNum
		record.CpuCores = p.Spec.CpuCores
		record.MemGiB = p.Spec.MemGiB
		record.GPUMemGiB = p.Spec.GPUMemGiB
		record.ShareMemGiB = p.Spec.ShareMemGiB
		record.ComputeResource = p.Spec.ComputeResource
		record.UnitPrice = p.Spec.UnitPrice
	}
	if err := models.CreateAgentExecutionRecord(record); err != nil {
		log.Error("recordAgentExecution failed. agentId=%d, userId=%d, action=%s, err=%v", p.AgentID, p.UserID, p.Action, err)
	}
}

// RunAgentRes 启动智能体响应
type RunAgentRes struct {
	AgentID      int64 `json:"agent_id"`      // 智能体 ID
	CloudBrainID int64 `json:"cloudbrain_id"` // 关联的云脑任务 ID
	IsNewCreate  bool  `json:"is_new_create"` // 是否是新创建的任务
}

// RunAgent 启动智能体（在 user_agent 表创建关联，并创建 AI 任务）
func RunAgent(agentId int64, userId int64) (*RunAgentRes, *response.BizError) {
	if agentId <= 0 {
		return nil, response.NewBizError(errors.New("agent id is required"))
	}
	if userId <= 0 {
		return nil, response.NewBizError(errors.New("user id is required"))
	}

	// 获取智能体
	agent, err := models.GetAgentByID(agentId)
	if err != nil || agent == nil {
		return nil, response.NewBizError(errors.New("agent not found"))
	}

	// 获取关联的模板，并加载关联的数据集和模型列表（存储在 template_datasets / template_models 表）
	template, err := models.GetAITaskTemplateByID(agent.TemplateID)
	if err != nil || template == nil {
		return nil, response.NewBizError(errors.New("template not found"))
	}

	if err := template.LoadAtrribute(0); err != nil {
		log.Error("RunAgent LoadAtrribute failed. templateId=%s, err=%v", agent.TemplateID, err)
		return nil, response.NewBizError(errors.New("failed to load template attributes"))
	}

	// 打印模板详细数据，用于调试参数错误问题
	log.Info("RunAgent template detail data. templateId=%s, ComputeSource=%s, Cluster=%s, BootFile=%s, BranchName=%s, ImageUrl=%s, ImageID=%s,  AccCardsNum=%d, AccCardType=%s, CpuCores=%d, MemGiB=%d, GPUMemGiB=%d, ShareMemGiB=%d, HasInternet=%d, JobType=%s, RepoId=%d, Description=%s",
		template.ID, template.ComputeSource, template.Cluster, template.BootFile, template.BranchName,
		template.ImageUrl, template.ImageID, template.AccCardsNum, template.AccCardType,
		template.CpuCores, template.MemGiB, template.GPUMemGiB, template.ShareMemGiB,
		template.HasInternet, template.JobType, template.RepoId, template.Description)

	// 检查是否已存在关联记录
	userAgent, err := models.GetUserAgentRelation(userId, agentId)
	if err != nil {
		log.Error("RunAgent GetUserAgentRelation failed. userId=%d, agentId=%d, err=%v", userId, agentId, err)
		return nil, response.NewBizError(errors.New("failed to get user agent relation"))
	}

	// 如果已存在关联记录且已有 CloudBrainID，检查云脑任务状态
	if userAgent != nil && userAgent.CloudBrainID > 0 {
		// 查询云脑任务状态
		cloudbrain, err := models.GetCloudbrainByCloudbrainID(userAgent.CloudBrainID)
		if err != nil {
			log.Error("RunAgent GetCloudbrainByCloudbrainID failed. cloudbrainId=%d, err=%v", userAgent.CloudBrainID, err)
			return nil, response.NewBizError(errors.New("CloudBrainID not found"))
		}

		if cloudbrain != nil && cloudbrain.Status == string(models.JobRunning) {
			// 任务正在运行，直接返回
			return &RunAgentRes{
				AgentID:      agentId,
				CloudBrainID: userAgent.CloudBrainID,
				IsNewCreate:  false,
			}, nil
		}

		// 任务不在运行中（已停止/失败/完成/停止中），清除旧关联，重新创建
		log.Info("RunAgent old task is not running, clearing old userAgent association. userId=%d, agentId=%d, oldCloudBrainId=%d",
			userId, agentId, userAgent.CloudBrainID)
		userAgent = nil
	}

	// 如果不存在关联记录，需要创建新的，检查用户智能体数量上限
	if userAgent == nil {
		total, err := models.GetUserAgentTotal(userId)
		if err != nil {
			log.Error("RunAgent GetUserAgentTotal failed. userId=%d, err=%v", userId, err)
			return nil, response.NewBizError(errors.New("failed to get user agent count"))
		}
		if total >= MaxUserAgentCount {
			return nil, response.NewBizError(fmt.Errorf("user agent limit exceeded, maximum %d allowed", MaxUserAgentCount))
		}
	}

	// 如果没有关联记录或关联记录已删除，继续创建新任务
	// 根据模板参数查询匹配的规格
	spec, err := getSpecByTemplate(template, userId)
	if err != nil || spec == nil {
		log.Error("RunAgent getSpecByTemplate failed. templateId=%s, err=%v", template.ID, err)
		return nil, response.NewBizError(errors.New("failed to get spec id"))
	}

	log.Info("RunAgent template info. templateId=%s, ComputeSource=%s, Cluster=%s, BootFile=%s, BranchName=%s, specId=%d",
		template.ID, template.ComputeSource, template.Cluster, template.BootFile, template.BranchName, spec.ID)

	// 获取用户
	user, err := models.GetUserByID(userId)
	if err != nil {
		log.Error("RunAgent GetUserByID failed. userId=%d, err=%v", userId, err)
		return nil, response.NewBizError(errors.New("user not found"))
	}

	// 获取 repo（如果 template 有关联）
	var gitRepo *git.Repository
	var repo *models.Repository
	if template.RepoId > 0 {
		repo, err = models.GetRepositoryByID(template.RepoId)
		if err != nil {
			log.Error("RunAgent GetRepositoryByID failed. repoId=%d, err=%v", template.RepoId, err)
		} else if repo != nil {
			gitRepo, err = git.OpenRepository(repo.RepoPath())
			if err != nil {
				log.Error("RunAgent OpenRepository failed. repoId=%d, err=%v", template.RepoId, err)
			}
		}
	}

	// 构建创建任务的参数
	// 注意：ImageUrl 和 ImageID 互斥，ImageUrl 优先级更高
	imageUrl := strings.TrimSpace(template.ImageUrl)
	imageID := strings.TrimSpace(template.ImageID)
	if imageUrl != "" {
		// 如果有 ImageUrl，忽略 ImageID
		imageID = ""
	}
	// 生成 JobName（基于 DisplayJobName）
	displayJobName := cloudbrain.GetDisplayJobName(user.Name)
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)

	// 提取模板绑定的模型 ID 和数据集 ID（分号拼接）
	modelIds := make([]string, len(template.ModelList))
	for i, m := range template.ModelList {
		modelIds[i] = m.ModelID
	}
	datasetIds := make([]string, len(template.DatasetList))
	for i, d := range template.DatasetList {
		datasetIds[i] = d.DatasetID
	}

	createReq := entity.CreateReq{
		JobType:           models.JobTypeOnlineInference,
		DisplayJobName:    displayJobName,
		JobName:           jobName,
		SpecId:            spec.ID,
		ComputeSourceStr:  template.ComputeSource,
		Cluster:           entity.ClusterType(template.Cluster),
		BranchName:        template.BranchName,
		BootFile:          template.BootFile,
		ImageUrl:          imageUrl,
		ImageID:           imageID,
		Description:       template.Description,
		HasInternet:       models.SpecInternetQuery(template.HasInternet),
		VisualizeRequired: template.VisualizeRequired,
		TemplateID:        template.ID,
		PretrainModelId:   strings.Join(modelIds, ";"),
		DatasetUUIDStr:    strings.Join(datasetIds, ";"),
		WorkServerNumber:  1,
	}

	// 调用 CreateAITask 创建任务
	traceContext, _ := otel.StartTraceParent(&entity.TraceInfo{TaskName: displayJobName, JobType: string(models.JobTypeOnlineInference), SpanName: "CreateAITask"}, "POST")
	var traceErr error
	defer func() {
		otel.FinalizeSpan(traceContext, traceErr)
	}()

	taskRes, bizErr := CreateAITask(createReq, gitRepo, repo, user, &traceContext)
	if bizErr != nil {
		traceErr = bizErr.ToError()
		log.Error("RunAgent CreateAITask failed. req=%+v, err=%v", createReq, bizErr)
		return nil, bizErr
	}

	// 解析返回的 CloudBrainID
	cloudbrainId := taskRes.ID
	if cloudbrainId <= 0 {
		log.Error("RunAgent CreateAITask returned invalid cloudbrainId=%d", cloudbrainId)
		return nil, response.NewBizError(errors.New("failed to get cloudbrain id"))
	}

	// 创建 user_agent 关联记录
	userAgent = &models.UserAgent{
		UserID:       userId,
		AgentID:      agentId,
		CloudBrainID: cloudbrainId,
	}
	err = models.CreateUserAgent(userAgent)
	if err != nil {
		log.Error("RunAgent CreateUserAgent failed. userId=%d, agentId=%d, err=%v", userId, agentId, err)
		return nil, response.NewBizError(errors.New("failed to create user agent relation"))
	}

	log.Info("RunAgent success. agentId=%d, userId=%d, cloudbrainId=%d", agentId, userId, cloudbrainId)

	recordAgentExecutionAsync(AgentExecutionParam{
		AgentID: agentId, UserID: userId, TemplateID: agent.TemplateID,
		CloudBrain: cloudbrainId, Action: "RUN", Success: true,
		Spec: &AgentResourceSpec{ID: spec.ID, AccCardType: spec.AccCardType, AccCardsNum: spec.AccCardsNum, CpuCores: spec.CpuCores, MemGiB: spec.MemGiB, GPUMemGiB: spec.GPUMemGiB, ShareMemGiB: spec.ShareMemGiB, ComputeResource: spec.ComputeResource, UnitPrice: spec.UnitPrice},
	})

	// 使用次数 +1
	_ = models.IncrementAgentUseCount(agentId)

	return &RunAgentRes{
		AgentID:      agentId,
		CloudBrainID: cloudbrainId,
		IsNewCreate:  true,
	}, nil
}

// getSpecByTemplate 根据模板参数查询匹配的规格
func getSpecByTemplate(template *models.AITaskTemplate, userId int64) (*models.AggregateSpecification, error) {
	// 查询可用的规格
	specs, err := resource.FindAvailableSpecsForNewRight(userId, models.FindSpecsOptions{
		ComputeResource: template.ComputeSource,
		Cluster:         template.Cluster,
		JobType:         models.JobType(template.JobType),
		SourceSpecId:    "", // 不限 sourceSpecId
	})
	if err != nil {
		return nil, err
	}

	// 根据模板的 spec 参数匹配
	for _, spec := range specs {
		if spec.AccCardsNum == template.AccCardsNum &&
			spec.AccCardType == template.AccCardType &&
			spec.CpuCores == template.CpuCores &&
			spec.MemGiB == template.MemGiB &&
			spec.GPUMemGiB == template.GPUMemGiB &&
			spec.ShareMemGiB == template.ShareMemGiB {
			return spec, nil
		}
	}

	// 如果没有找到匹配的规格，返回第一个可用规格
	if len(specs) > 0 {
		return specs[0], nil
	}

	return nil, errors.New("no available spec found")
}

// StopAgent 停止智能体（根据 cloudbrain_id 停止云脑任务）
func StopAgent(cloudbrainId int64, userId int64) *response.BizError {
	if cloudbrainId <= 0 {
		return response.NewBizError(errors.New("cloudbrain id is required"))
	}
	if userId <= 0 {
		return response.NewBizError(errors.New("user id is required"))
	}

	// 获取 Cloudbrain 任务
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		log.Error("StopAgent GetCloudbrainByCloudbrainID failed. cloudbrainId=%d, err=%v", cloudbrainId, err)
		return response.NewBizError(errors.New("failed to get cloudbrain task"))
	}
	if cloudbrain == nil {
		log.Error("StopAgent cloudbrain not found. cloudbrainId=%d", cloudbrainId)
		return response.NewBizError(errors.New("cloudbrain task not found"))
	}

	// 调用 StopCloudbrain 停止任务
	// 注意：StopCloudbrain 会触发 NotifyChangeCloudbrainStatus 通知，
	// 进而通过 StatusChangeChan 异步删除 user_agent 关联（在 AcceptStatusChangeAction 中处理）
	_, bizErr := StopCloudbrain(cloudbrain)
	if bizErr != nil {
		log.Error("StopAgent StopCloudbrain failed. cloudbrainId=%d, err=%v", cloudbrainId, bizErr)
		return response.NewBizError(errors.New("failed to stop cloudbrain task"))
	}

	// 立即删除 user_agent 关联记录，避免异步清理延迟导致 run 不起
	if _, err := models.DeleteUserAgentByCloudBrainIDOnly(cloudbrainId, "STOPPED", cloudbrain.Status); err != nil {
		log.Error("StopAgent DeleteUserAgentByCloudBrainIDOnly failed. cloudbrainId=%d, err=%v", cloudbrainId, err)
	}

	log.Info("StopAgent success. cloudbrainId=%d, userId=%d", cloudbrainId, userId)

	// 获取智能体和模板信息（用于记录执行日志）
	userAgent, err := models.GetAgentExecutionRecordByCloudBrainID(cloudbrainId)
	var agentId int64
	var templateId string
	if userAgent != nil && err == nil {
		agentId = userAgent.AgentID
		agent, aerr := models.GetAgentByID(agentId)
		if agent != nil && aerr == nil {
			templateId = agent.TemplateID
		}
	}

	// 获取规格信息（用于记录执行日志）
	spec, _ := models.GetCloudbrainSpecByID(cloudbrainId)
	if agentId > 0 && spec != nil {
		recordAgentExecutionAsync(AgentExecutionParam{
			AgentID: agentId, UserID: userId, TemplateID: templateId,
			CloudBrain: cloudbrainId, Action: "STOP", Success: true,
			Spec: &AgentResourceSpec{ID: spec.SpecId, AccCardType: spec.AccCardType, AccCardsNum: spec.AccCardsNum, CpuCores: spec.CpuCores, MemGiB: spec.MemGiB, GPUMemGiB: spec.GPUMemGiB, ShareMemGiB: spec.ShareMemGiB, ComputeResource: spec.ComputeResource, UnitPrice: spec.UnitPrice},
		})
	}

	return nil
}
