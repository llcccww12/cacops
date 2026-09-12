package models

import (
	"errors"
	"fmt"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

// AgentType 智能体类型
type AgentType string

const (
	// AgentTypeOpenClaw OpenClaw 类型智能体
	AgentTypeOpenClaw AgentType = "OPENCLAW"
	// AgentTypeOpenCode OpenCode 类型智能体
	AgentTypeOpenCode AgentType = "OPENCODE"
)

// AgentStatus 智能体状态
type AgentStatus int8

const (
	// AgentStatusDraft 草稿状态（未发布）
	AgentStatusDraft AgentStatus = 0
	// AgentStatusPublished 已发布
	AgentStatusPublished AgentStatus = 1
	// AgentStatusUnpublished 已下架
	AgentStatusUnpublished AgentStatus = 2
	// AgentStatusPaused 暂停状态
	AgentStatusPaused AgentStatus = 3
)

// String 返回状态的字符串表示
func (s AgentStatus) String() string {
	switch s {
	case AgentStatusDraft:
		return "DRAFT"
	case AgentStatusPublished:
		return "PUBLISHED"
	case AgentStatusUnpublished:
		return "UNPUBLISHED"
	case AgentStatusPaused:
		return "PAUSED"
	default:
		return "UNKNOWN"
	}
}

// ChineseString 返回状态的中文表示
func (s AgentStatus) ChineseString() string {
	switch s {
	case AgentStatusDraft:
		return "草稿"
	case AgentStatusPublished:
		return "已发布"
	case AgentStatusUnpublished:
		return "已下架"
	case AgentStatusPaused:
		return "暂停"
	default:
		return "未知"
	}
}

// IsValid 检查状态是否有效
func (s AgentStatus) IsValid() bool {
	return s >= AgentStatusDraft && s <= AgentStatusPaused
}

// Agent 智能体
// 独立表结构，ID 自增
type Agent struct {
	ID            int64              `xorm:"pk autoincr"`              // 自增 id
	AgentType     AgentType          `xorm:"INDEX NOT NULL"`           // 智能体类型：OPENCLAW / OPENCODE
	TemplateID    string             `xorm:"INDEX NOT NULL"`           // 关联的模板 ID（ai_task_template.id）
	Status        AgentStatus        `xorm:"INDEX NOT NULL DEFAULT 0"` // 智能体状态（0 草稿 1 已发布 2 已下架）
	Name          string             `xorm:"varchar(255)"`             // 智能体昵称
	AvatarUrl     string             `xorm:"varchar(1024)"`            // 智能体头像框
	BackGroundUrl string             `xorm:"varchar(1024)"`            // 智能体背景图
	DetailURL     string             `xorm:"varchar(1024)"`            // 详情页跳转链接 URL
	Tag           []string           `xorm:"json"`                     // 打上 tag
	UseCount      int64              `xorm:"NOT NULL DEFAULT 0"`       // 运行次数
	CreatorId     int64              `xorm:"INDEX NOT NULL DEFAULT 0"` // 创建智能体的用户 id
	Description   string             `xorm:"text"`                     // 智能体描述
	CreatedUnix   timeutil.TimeStamp `xorm:"INDEX created"`            // 创建时间
	UpdatedUnix   timeutil.TimeStamp `xorm:"INDEX updated"`            // 创建时间
	RecommendUnix timeutil.TimeStamp `xorm:"INDEX"`                    // 推荐时间
}

// UserAgent 用户与智能体的关联表
// 用于实现用户与智能体的多对多关系（一个用户可以有多个智能体，一个智能体可以被多个用户使用）
type UserAgent struct {
	ID           int64              `xorm:"pk autoincr"`    // 自增主键
	UserID       int64              `xorm:"INDEX NOT NULL"` // 用户 ID
	AgentID      int64              `xorm:"INDEX NOT NULL"` // 智能体 ID（关联 agent.id）
	CloudBrainID int64              `xorm:"INDEX NOT NULL"` // 关联的 Cloudbrain ID
	AgentType    AgentType          `xorm:"-"`              // 智能体类型（通过关联 agent 表获取）
	CreatedUnix  timeutil.TimeStamp `xorm:"INDEX created"`  // 创建时间
	UpdatedUnix  timeutil.TimeStamp `xorm:"INDEX updated"`  // 更新时间
}

// UserAgentRecord 用户智能体使用记录（用于历史回溯）
// 当 UserAgent 关联被删除时，保存历史记录到此表
type UserAgentRecord struct {
	ID               int64              `xorm:"pk autoincr"`    // 自增主键
	UserID           int64              `xorm:"INDEX NOT NULL"` // 用户 ID
	AgentID          int64              `xorm:"INDEX NOT NULL"` // 智能体 ID
	CloudBrainID     int64              `xorm:"INDEX NOT NULL"` // 关联的 Cloudbrain ID
	CloudBrainName   string             `xorm:"VARCHAR(255)"`   // Cloudbrain 名称（冗余字段，方便查询）
	UserAgentID      int64              `xorm:"INDEX NOT NULL"` // 原 UserAgent 表的 ID
	DeleteReason     string             `xorm:"VARCHAR(50)"`    // 删除原因：STOPPED(主动停止), FAILED(失败)
	CloudBrainStatus string             `xorm:"VARCHAR(50)"`    // 删除时的 Cloudbrain 状态
	TemplateID       string             `xorm:"VARCHAR(255)"`   // 智能体模板 ID
	CreatedUnix      timeutil.TimeStamp `xorm:"INDEX created"`  // 原始创建时间
	DeletedUnix      timeutil.TimeStamp `xorm:"INDEX deleted"`  // 删除时间
}

// AgentExecutionRecord 智能体执行记录（RunAgent / StopAgent 时记录）
type AgentExecutionRecord struct {
	ID              int64              `xorm:"pk autoincr"`            // 自增主键
	AgentID         int64              `xorm:"INDEX NOT NULL"`         // 智能体 ID
	UserID          int64              `xorm:"INDEX NOT NULL"`         // 用户 ID
	TemplateID      string             `xorm:"VARCHAR(255) NOT NULL"`  // 模板 ID
	CloudBrainID    int64              `xorm:"INDEX NOT NULL"`         // Cloudbrain 任务 ID
	Action          string             `xorm:"VARCHAR(20) NOT NULL"`   // 操作类型：RUN / STOP
	Success         bool               `xorm:"NOT NULL DEFAULT true"`  // 是否成功
	Message         string             `xorm:"VARCHAR(500)"`           // 错误信息（成功时为空）
	ResourceSpecID  int64              `xorm:"NOT NULL DEFAULT 0"`     // 资源规格 ID
	AccCardType     string             `xorm:"VARCHAR(50)"`            // GPU/加速卡类型
	AccCardsNum     int                `xorm:"NOT NULL DEFAULT 0"`     // 加速卡数量
	CpuCores        int                `xorm:"NOT NULL DEFAULT 0"`     // CPU 核心数
	MemGiB          float32            `xorm:"NOT NULL DEFAULT 0"`     // 内存大小 (GiB)
	GPUMemGiB       float32            `xorm:"NOT NULL DEFAULT 0"`     // GPU 显存大小 (GiB)
	ShareMemGiB     float32            `xorm:"NOT NULL DEFAULT 0"`     // 共享内存大小 (GiB)
	ComputeResource string             `xorm:"VARCHAR(50)"`            // 算力资源（GPU/NPU）
	UnitPrice       float64            `xorm:"NOT NULL DEFAULT 0"`     // 单价（积分/小时）
	CreatedUnix     timeutil.TimeStamp `xorm:"INDEX NOT NULL created"` // 执行时间
}

// AgentListRes 智能体列表响应
type AgentListRes struct {
	Agents   []*AgentInfo `json:"agents"`
	Total    int64        `json:"total"`
	PageSize int          `json:"page_size"`
	Page     int          `json:"page"`
}

// UserAgentInfo 用户智能体信息（用于 API 返回，包含 cloudbrain 运行状态和用户信息）
type UserAgentInfo struct {
	ID           int64              `json:"id"`            // user_agent 表 ID
	UserID       int64              `json:"user_id"`       // 用户 ID
	AgentID      int64              `json:"agent_id"`      // 智能体 ID
	CloudBrainID int64              `json:"cloudbrain_id"` // 关联的 Cloudbrain ID
	AgentType    AgentType          `json:"agent_type"`    // 智能体类型
	AgentName    string             `json:"agent_name"`    // 智能体名称
	AgentStatus  AgentStatus        `json:"agent_status"`  // 智能体状态
	JobStatus    string             `json:"job_status"`    // Cloudbrain 运行状态
	CreatedUnix  timeutil.TimeStamp `json:"created_unix"`  // 创建时间
	UpdatedUnix  timeutil.TimeStamp `json:"updated_unix"`  // 更新时间
}

// UserAgentListRes 用户智能体关联列表响应
type UserAgentListRes struct {
	UserAgents          []*UserAgentInfo `json:"user_agents"`
	Total               int64            `json:"total"`
	PageSize            int              `json:"page_size"`
	Page                int              `json:"page"`
	RunningCount        int64            `json:"running_count"`         // 正在运行状态的智能体总个数
	TotalPointsConsumed float64          `json:"total_points_consumed"` // 总积分消耗
}

// AgentSpec 智能体规格信息（来自 AITaskTemplate 的 spec 参数）
type AgentSpec struct {
	AccCardsNum     int     `json:"acc_cards_num"`    // 加速卡数量
	AccCardType     string  `json:"acc_card_type"`    // 加速卡类型
	CpuCores        int     `json:"cpu_cores"`        // CPU 核心数
	MemGiB          float32 `json:"mem_gib"`          // 内存大小 (GiB)
	GPUMemGiB       float32 `json:"gpu_mem_gib"`      // GPU 显存大小 (GiB)
	ShareMemGiB     float32 `json:"share_mem_gib"`    // 共享内存大小 (GiB)
	ComputeResource string  `json:"compute_resource"` // 算力资源（GPU/NPU）
	UnitPrice       float64 `json:"unit_price"`       // 单价（积分/小时）
}

// AgentInfo 智能体信息（用于 API 返回，隐藏内部字段）
type AgentInfo struct {
	ID            int64              `json:"id"`
	AgentType     AgentType          `json:"agent_type"`
	TemplateID    string             `json:"template_id"`
	Status        AgentStatus        `json:"status"`
	Name          string             `json:"name"`
	AvatarUrl     string             `json:"avatar_url"`
	BackGroundUrl string             `json:"back_ground_url"`
	DetailURL     string             `json:"detail_url"`
	Tag           []string           `json:"tag"`
	UseCount      int64              `json:"use_count"`
	Description   string             `json:"description"`
	CreatedUnix   timeutil.TimeStamp `json:"created_unix"`
	UpdatedUnix   timeutil.TimeStamp `json:"updated_unix"`
	RecommendUnix timeutil.TimeStamp `json:"recommend_unix"`
	Spec          *AgentSpec         `json:"spec,omitempty"` // 智能体规格信息
	UserID        int64              `json:"user_id"`        // 创建者用户 ID
	UserName      string             `json:"user_name"`      // 创建者用户名
	UserAvatar    string             `json:"user_avatar"`    // 创建者用户头像
}

// ToAgentInfo 将 Agent 转换为 AgentInfo
func (a *Agent) ToAgentInfo() *AgentInfo {
	return &AgentInfo{
		ID:            a.ID,
		AgentType:     a.AgentType,
		TemplateID:    a.TemplateID,
		Status:        a.Status,
		Name:          a.Name,
		AvatarUrl:     a.AvatarUrl,
		BackGroundUrl: a.BackGroundUrl,
		DetailURL:     a.DetailURL,
		Tag:           a.Tag,
		UseCount:      a.UseCount,
		Description:   a.Description,
		CreatedUnix:   a.CreatedUnix,
		UpdatedUnix:   a.UpdatedUnix,
		RecommendUnix: a.RecommendUnix,
	}
}

// ToAgentInfoWithSpec 将 Agent 转换为 AgentInfo（带规格信息）
func (a *Agent) ToAgentInfoWithSpec(template *AITaskTemplate) *AgentInfo {
	info := a.ToAgentInfo()
	if template != nil {
		info.Spec = &AgentSpec{
			AccCardsNum:     template.AccCardsNum,
			AccCardType:     template.AccCardType,
			CpuCores:        template.CpuCores,
			MemGiB:          template.MemGiB,
			GPUMemGiB:       template.GPUMemGiB,
			ShareMemGiB:     template.ShareMemGiB,
			ComputeResource: template.ComputeSource,
		}
	}
	return info
}

// ToAgentInfoWithSpecAndPrice 将 Agent 转换为 AgentInfo（带规格信息和单价）
func (a *Agent) ToAgentInfoWithSpecAndPrice(template *AITaskTemplate, unitPrice float64) *AgentInfo {
	info := a.ToAgentInfo()
	if template != nil {
		info.Spec = &AgentSpec{
			AccCardsNum:     template.AccCardsNum,
			AccCardType:     template.AccCardType,
			CpuCores:        template.CpuCores,
			MemGiB:          template.MemGiB,
			GPUMemGiB:       template.GPUMemGiB,
			ShareMemGiB:     template.ShareMemGiB,
			ComputeResource: template.ComputeSource,
			UnitPrice:       unitPrice,
		}
	}
	return info
}

// GetAgentList 获取智能体列表（直接查询 agent 表，不涉及 user_agent 关联）
func GetAgentList(q string, page, pageSize int) (*AgentListRes, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	session := x.NewSession()
	defer session.Close()

	// 构建查询条件：status = 1（已发布）
	var cond builder.Cond = builder.Eq{"status": AgentStatusPublished}

	// 如果有搜索关键词，追加 name 或 description 的模糊查询（大小写不敏感）
	if q != "" {
		lowerQ := strings.ToLower(q)
		cond = cond.And(builder.Or(builder.Like{"LOWER(name)", lowerQ}, builder.Like{"LOWER(description)", lowerQ}))
	}

	total, err := session.Where(cond).Count(&Agent{})
	if err != nil {
		return nil, err
	}

	agents := make([]*Agent, 0)
	err = session.Where(cond).OrderBy("recommend_unix DESC").
		OrderBy("use_count DESC").
		OrderBy("updated_unix DESC").
		Limit(pageSize, (page-1)*pageSize).
		Find(&agents)
	if err != nil {
		return nil, err
	}

	// 收集所有 template_id（去重）
	templateIDSet := make(map[string]struct{})
	for _, agent := range agents {
		if agent.TemplateID != "" {
			templateIDSet[agent.TemplateID] = struct{}{}
		}
	}
	templateIDs := make([]string, 0, len(templateIDSet))
	for id := range templateIDSet {
		templateIDs = append(templateIDs, id)
	}

	// 批量查询模板数据
	templateMap, err := GetAITaskTemplatesByIDs(templateIDs)
	if err != nil {
		log.Error("GetAgentList query templates err. templateIDs=%v, err=%v", templateIDs, err)
	}

	// 收集所有规格参数用于批量查询单价
	type specParams struct {
		AccCardType     string
		AccCardsNum     int
		CpuCores        int
		MemGiB          float32
		GPUMemGiB       float32
		ShareMemGiB     float32
		ComputeResource string
		TemplateID      string
	}
	specParamsList := make([]specParams, 0)
	for _, agent := range agents {
		template := templateMap[agent.TemplateID]
		if template != nil && template.AccCardType != "" {
			specParamsList = append(specParamsList, specParams{
				AccCardType:     template.AccCardType,
				AccCardsNum:     template.AccCardsNum,
				CpuCores:        template.CpuCores,
				MemGiB:          template.MemGiB,
				GPUMemGiB:       template.GPUMemGiB,
				ShareMemGiB:     template.ShareMemGiB,
				ComputeResource: template.ComputeSource,
				TemplateID:      agent.TemplateID,
			})
		}
	}

	// 批量查询单价（按 specKey 分组）
	priceMap := make(map[string]float64)
	specKeyMap := make(map[string]string) // templateID -> specKey
	for _, params := range specParamsList {
		specKey := fmt.Sprintf("%s_%d_%d_%f_%f_%f_%s",
			params.AccCardType, params.AccCardsNum, params.CpuCores,
			params.MemGiB, params.GPUMemGiB, params.ShareMemGiB, params.ComputeResource)
		specKeyMap[params.TemplateID] = specKey

		if _, exists := priceMap[specKey]; !exists {
			price, err := GetUnitPriceBySpec(
				params.AccCardType, params.AccCardsNum, params.CpuCores,
				params.MemGiB, params.GPUMemGiB, params.ShareMemGiB, params.ComputeResource,
			)
			if err != nil {
				log.Warn("GetUnitPriceBySpec failed. params=%+v, err=%v", params, err)
				price = 0
			}
			priceMap[specKey] = price
		}
	}

	// 收集所有创建者 ID（去重）
	creatorIDs := make([]int64, 0)
	creatorIDSet := make(map[int64]struct{})
	for _, agent := range agents {
		if agent.CreatorId > 0 {
			if _, exists := creatorIDSet[agent.CreatorId]; !exists {
				creatorIDSet[agent.CreatorId] = struct{}{}
				creatorIDs = append(creatorIDs, agent.CreatorId)
			}
		}
	}

	// 批量查询创建者用户信息
	creatorMap := make(map[int64]*User)
	if len(creatorIDs) > 0 {
		users := make([]*User, 0)
		err := x.In("id", creatorIDs).Find(&users)
		if err != nil {
			log.Warn("GetAgentList query creators err. creatorIDs=%v, err=%v", creatorIDs, err)
		} else {
			for _, u := range users {
				creatorMap[u.ID] = u
			}
		}
	}

	// 将 Agent 转换为 AgentInfo（带规格信息和单价）
	agentInfos := make([]*AgentInfo, 0, len(agents))
	for _, agent := range agents {
		template := templateMap[agent.TemplateID]
		var unitPrice float64
		if specKey, ok := specKeyMap[agent.TemplateID]; ok {
			unitPrice = priceMap[specKey]
		}
		info := agent.ToAgentInfoWithSpecAndPrice(template, unitPrice)

		// 填充创建者信息
		if creator, ok := creatorMap[agent.CreatorId]; ok {
			info.UserID = creator.ID
			info.UserName = creator.Name
			info.UserAvatar = creator.AvatarLink()
		}

		agentInfos = append(agentInfos, info)
	}

	return &AgentListRes{
		Agents:   agentInfos,
		Total:    total,
		PageSize: pageSize,
		Page:     page,
	}, nil
}

// GetUserAgentList 根据用户 ID 获取用户智能体列表（带 cloudbrain job_status 状态，以及运行统计信息和用户信息）
func GetUserAgentList(userID int64, page, pageSize int) (*UserAgentListRes, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 创建基础 session 用于 COUNT 查询（只显示运行中、成功、停止中的关联）
	countSession := x.Join("INNER", "agent", "user_agent.agent_id = agent.id").
		Join("LEFT", "cloudbrain", "user_agent.cloud_brain_id = cloudbrain.id").
		Where("user_agent.user_id = ?", userID)

	total, err := countSession.Count(&UserAgent{})
	if err != nil {
		return nil, err
	}

	// 使用 map 接收查询结果
	type Result struct {
		ID           int64
		UserID       int64
		UserName     string
		UserAvatar   string
		AgentID      int64
		CloudBrainID int64
		AgentType    AgentType
		AgentName    string
		AgentStatus  AgentStatus
		JobStatus    string
		CreatedUnix  timeutil.TimeStamp
		UpdatedUnix  timeutil.TimeStamp
	}

	results := make([]*Result, 0)
	query := x.Table("user_agent").
		Select("user_agent.id, user_agent.user_id, public.user.name as user_name, public.user.avatar as user_avatar, user_agent.agent_id, user_agent.cloud_brain_id, agent.agent_type, agent.name as agent_name, agent.status as agent_status, cloudbrain.status as job_status, user_agent.created_unix, user_agent.updated_unix").
		Join("INNER", "agent", "user_agent.agent_id = agent.id").
		Join("LEFT", "cloudbrain", "user_agent.cloud_brain_id = cloudbrain.id").
		Join("LEFT", "user", "user_agent.user_id = public.user.id").
		Where("user_agent.user_id = ?", userID)
	err = query.
		OrderBy("user_agent.created_unix DESC").
		Limit(pageSize, (page-1)*pageSize).
		Find(&results)
	if err != nil {
		return nil, err
	}

	// 转换为 UserAgentInfo
	userAgentInfos := make([]*UserAgentInfo, 0, len(results))
	for _, r := range results {
		userAgentInfos = append(userAgentInfos, &UserAgentInfo{
			ID:           r.ID,
			UserID:       r.UserID,
			AgentID:      r.AgentID,
			CloudBrainID: r.CloudBrainID,
			AgentType:    r.AgentType,
			AgentName:    r.AgentName,
			AgentStatus:  r.AgentStatus,
			JobStatus:    r.JobStatus,
			CreatedUnix:  r.CreatedUnix,
			UpdatedUnix:  r.UpdatedUnix,
		})
	}

	// 统计正在运行状态的智能体总个数和总积分消耗
	runningCount, totalPointsConsumed, err := calculateUserAgentStats(userID)
	if err != nil {
		log.Error("GetUserAgentList calculateUserAgentStats failed. userID=%d, err=%v", userID, err)
		// 统计失败不影响主流程，返回 0 值
		runningCount = 0
		totalPointsConsumed = 0
	}

	return &UserAgentListRes{
		UserAgents:          userAgentInfos,
		Total:               total,
		PageSize:            pageSize,
		Page:                page,
		RunningCount:        runningCount,
		TotalPointsConsumed: totalPointsConsumed,
	}, nil
}

// calculateUserAgentStats 统计用户正在运行状态的智能体个数和总积分消耗（每小时）
func calculateUserAgentStats(userID int64) (int64, float64, error) {
	// 查询用户所有处于 RUNNING 状态的 cloudbrain 任务
	type RunningTask struct {
		CloudBrainID   int64
		ResourceSpecId int
	}

	runningTasks := make([]RunningTask, 0)
	err := x.Table("user_agent").
		Select("user_agent.cloud_brain_id, cloudbrain.resource_spec_id").
		Join("INNER", "cloudbrain", "user_agent.cloud_brain_id = cloudbrain.id").
		Where("user_agent.user_id = ? AND cloudbrain.status = ?", userID, JobRunning).
		Find(&runningTasks)
	if err != nil {
		return 0, 0, err
	}

	runningCount := int64(len(runningTasks))
	totalPointsConsumed := 0.0

	// 遍历每个运行中的任务，累加单价（每小时消耗的积分）
	for _, task := range runningTasks {
		if task.ResourceSpecId <= 0 {
			continue
		}

		// 查询规格单价
		spec := &ResourceSpecification{}
		has, err := x.ID(task.ResourceSpecId).Get(spec)
		if err != nil || !has {
			log.Warn("calculateUserAgentStats get spec failed. specId=%d, err=%v", task.ResourceSpecId, err)
			continue
		}

		// 累加单价（每小时消耗的积分）
		totalPointsConsumed += spec.UnitPrice
	}

	return runningCount, totalPointsConsumed, nil
}

// GetUserAgentRelation 获取用户与智能体的关联关系
func GetUserAgentRelation(userID int64, agentID int64) (*UserAgent, error) {
	userAgent := &UserAgent{}
	has, err := x.Where("user_id = ? AND agent_id = ?", userID, agentID).Get(userAgent)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userAgent, nil
}

// CheckUserAgentRelation 检查用户是否有权限访问指定智能体
func CheckUserAgentRelation(userID int64, agentID int64) bool {
	userAgent, _ := GetUserAgentRelation(userID, agentID)
	return userAgent != nil
}

// CreateUserAgent 创建用户智能体关联
func CreateUserAgent(userAgent *UserAgent) error {
	_, err := x.Insert(userAgent)
	return err
}

// DeleteUserAgentByCloudBrainID 根据用户 ID 和 CloudBrainID 删除关联
func DeleteUserAgentByCloudBrainID(userID int64, cloudBrainID int64) error {
	_, err := x.Where("user_id = ? AND cloud_brain_id = ?", userID, cloudBrainID).Delete(&UserAgent{})
	return err
}

// UpdateUserAgent 更新用户智能体关联
func UpdateUserAgent(userAgent *UserAgent) error {
	_, err := x.ID(userAgent.ID).Cols("cloudbrain_id", "updated_unix").Update(userAgent)
	return err
}

// CreateAgent 创建智能体
func CreateAgent(agent *Agent) error {
	_, err := x.Insert(agent)
	return err
}

// UpdateAgent 更新智能体
func UpdateAgent(agent *Agent) error {
	_, err := x.ID(agent.ID).Cols("name", "description", "template_id", "back_ground_url", "detail_url", "avatar_url", "tag", "updated_unix").Update(agent)
	return err
}

// DeleteAgent 删除智能体
func DeleteAgent(id int64) error {
	_, err := x.ID(id).Delete(&Agent{})
	return err
}

// GetUserAgentTotal 获取用户的智能体关联总数
func GetUserAgentTotal(userID int64) (int64, error) {
	return x.Where("user_id = ?", userID).Count(&UserAgent{})
}

// GetAgentByID 根据智能体 ID 获取 Agent
func GetAgentByID(id int64) (*Agent, error) {
	agent := &Agent{}
	has, err := x.ID(id).Get(agent)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return agent, nil
}

// GetAITaskTemplatesByIDs 根据模板 ID 列表批量获取模板
func GetAITaskTemplatesByIDs(templateIDs []string) (map[string]*AITaskTemplate, error) {
	if len(templateIDs) == 0 {
		return nil, nil
	}

	templates := make([]*AITaskTemplate, 0)
	err := x.In("id", templateIDs).Find(&templates)
	if err != nil {
		return nil, err
	}

	templateMap := make(map[string]*AITaskTemplate, len(templates))
	for _, template := range templates {
		templateMap[template.ID] = template
	}

	return templateMap, nil
}

// UpdateAgentStatus 更新智能体状态
func UpdateAgentStatus(id int64, status AgentStatus) error {
	if !status.IsValid() {
		return errors.New("invalid agent status")
	}
	_, err := x.ID(id).Cols("status", "updated_unix").Update(&Agent{
		Status:      status,
		UpdatedUnix: timeutil.TimeStamp(timeutil.TimeStampNow()),
	})
	return err
}

// IncrementAgentUseCount 智能体使用次数 +1
func IncrementAgentUseCount(id int64) error {
	_, err := x.ID(id).Incr("use_count").Update(&Agent{})
	return err
}

// IsDraft 检查智能体是否为草稿状态
func (a *Agent) IsDraft() bool {
	return a.Status == AgentStatusDraft
}

// IsPublished 检查智能体是否已发布
func (a *Agent) IsPublished() bool {
	return a.Status == AgentStatusPublished
}

// IsUnpublished 检查智能体是否已下架
func (a *Agent) IsUnpublished() bool {
	return a.Status == AgentStatusUnpublished
}

// IsPaused 检查智能体是否已暂停
func (a *Agent) IsPaused() bool {
	return a.Status == AgentStatusPaused
}

// GetUnitPriceBySpec 根据规格参数查询单价
// 从 resource_specification 表中查询匹配规格的单价
func GetUnitPriceBySpec(accCardType string, accCardsNum int, cpuCores int, memGiB float32, gpuMemGiB float32, shareMemGiB float32, computeResource string) (float64, error) {
	spec := &ResourceSpecification{}
	has, err := x.Table("resource_specification").
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Where("resource_specification.acc_cards_num = ?", accCardsNum).
		And("resource_specification.cpu_cores = ?", cpuCores).
		And("resource_specification.mem_gi_b = ?", memGiB).
		And("resource_specification.gpu_mem_gi_b = ?", gpuMemGiB).
		And("resource_specification.share_mem_gi_b = ?", shareMemGiB).
		And("resource_queue.acc_card_type = ?", accCardType).
		And("resource_queue.compute_resource = ?", computeResource).
		And("resource_specification.status = ?", 2). // 只查询上架状态的规格
		Get(spec)

	if err != nil {
		return 0, err
	}
	if !has {
		return 0, nil
	}
	return spec.UnitPrice, nil
}

// DeleteUserAgentByCloudBrainIDOnly 仅根据 CloudBrainID 删除关联并保存记录（用于状态变更监听）
// 删除原因：FAILED(失败), STOPPED(主动停止)
func DeleteUserAgentByCloudBrainIDOnly(cloudBrainID int64, deleteReason string, cloudBrainStatus string) (int64, error) {
	// 查询要删除的 UserAgent 记录
	userAgents := make([]*UserAgent, 0)
	err := x.Where("cloud_brain_id = ?", cloudBrainID).Find(&userAgents)
	if err != nil {
		return 0, err
	}
	if len(userAgents) == 0 {
		return 0, nil
	}

	// 获取 cloudbrain 名称和 agent 的 template_id
	cloudbrainName := ""
	templateID := ""
	cloudbrain, err := GetCloudbrainByCloudbrainID(cloudBrainID)
	if err == nil && cloudbrain != nil {
		cloudbrainName = cloudbrain.DisplayJobName
	}
	// 获取第一个 user_agent 关联的 agent 的 template_id
	if len(userAgents) > 0 {
		agent, err := GetAgentByID(userAgents[0].AgentID)
		if err == nil && agent != nil {
			templateID = agent.TemplateID
		}
	}

	// 创建会话，保证原子性操作
	session := x.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}

	// 保存历史记录并删除 UserAgent
	deletedCount := int64(0)
	for _, ua := range userAgents {
		// 创建 UserAgentRecord 记录
		record := &UserAgentRecord{
			UserID:           ua.UserID,
			AgentID:          ua.AgentID,
			CloudBrainID:     ua.CloudBrainID,
			CloudBrainName:   cloudbrainName,
			UserAgentID:      ua.ID,
			DeleteReason:     deleteReason,
			CloudBrainStatus: cloudBrainStatus,
			TemplateID:       templateID,
			CreatedUnix:      ua.CreatedUnix,
			DeletedUnix:      timeutil.TimeStamp(timeutil.TimeStampNow()),
		}
		if _, err := session.Insert(record); err != nil {
			session.Rollback()
			return 0, err
		}

		// 删除 UserAgent
		if _, err := session.ID(ua.ID).Delete(&UserAgent{}); err != nil {
			session.Rollback()
			return 0, err
		}
		deletedCount++
	}

	if err := session.Commit(); err != nil {
		return 0, err
	}

	log.Info("DeleteUserAgentByCloudBrainIDOnly deleted %d records. cloudbrainID=%d, reason=%s", deletedCount, cloudBrainID, deleteReason)
	return deletedCount, nil
}

// GetUserAgentRecords 获取用户智能体历史记录
// userID 为 0 时获取所有记录，否则只获取指定用户的记录
func GetUserAgentRecords(userID int64, page, pageSize int) ([]*UserAgentRecord, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	session := x.NewSession()
	defer session.Close()

	// 如果 userID 不为 0，则按用户筛选
	if userID > 0 {
		session.Where("user_id = ?", userID)
	}

	total, err := session.Count(&UserAgentRecord{})
	if err != nil {
		return nil, 0, err
	}

	records := make([]*UserAgentRecord, 0)
	recordsSession := x.NewSession()
	defer recordsSession.Close()

	// 如果 userID 不为 0，则按用户筛选
	if userID > 0 {
		recordsSession.Where("user_id = ?", userID)
	}

	err = recordsSession.OrderBy("deleted_unix DESC").
		Limit(pageSize, (page-1)*pageSize).
		Find(&records)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// CreateAgentExecutionRecord 创建智能体执行记录
func CreateAgentExecutionRecord(record *AgentExecutionRecord) error {
	_, err := x.Insert(record)
	return err
}

// GetAgentExecutionRecordByCloudBrainID 根据 CloudBrainID 获取 UserAgent 关联
func GetAgentExecutionRecordByCloudBrainID(cloudBrainID int64) (*UserAgent, error) {
	ua := &UserAgent{}
	has, err := x.Where("cloud_brain_id = ?", cloudBrainID).Get(ua)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return ua, nil
}
