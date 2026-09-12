package models

import (
	"fmt"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	SpecNotVerified int = iota + 1
	SpecOnShelf
	SpecOffShelf
)

type SearchSpecOrderBy int

const (
	SearchSpecOrderById SearchSpecOrderBy = iota
	SearchSpecOrder4Standard
)

type AICenterInternetStatus int

const (
	Unknown     AICenterInternetStatus = -1
	NoInternet  AICenterInternetStatus = 1
	HasInternet AICenterInternetStatus = 2
)

type SpecInternetQuery int

const (
	QueryAllSpecs         SpecInternetQuery = 0
	QueryNoInternetSpecs  SpecInternetQuery = 1
	QueryHasInternetSpecs SpecInternetQuery = 2
)

type ResourceSpecification struct {
	ID              int64  `xorm:"pk autoincr"`
	QueueId         int64  `xorm:"INDEX"`
	SourceSpecId    string `xorm:"INDEX"`
	AccCardsNum     int
	CpuCores        int
	MemGiB          float32
	GPUMemGiB       float32
	ShareMemGiB     float32
	UnitPrice       float64
	Status          int
	IsAvailable     bool
	IsAutomaticSync bool
	CreatedTime     timeutil.TimeStamp `xorm:"created"`
	CreatedBy       int64
	UpdatedTime     timeutil.TimeStamp `xorm:"updated"`
	UpdatedBy       int64
}

func (r ResourceSpecification) ConvertToRes() *ResourceSpecificationRes {
	return &ResourceSpecificationRes{
		ID:           r.ID,
		SourceSpecId: r.SourceSpecId,
		AccCardsNum:  r.AccCardsNum,
		CpuCores:     r.CpuCores,
		MemGiB:       r.MemGiB,
		ShareMemGiB:  r.ShareMemGiB,
		GPUMemGiB:    r.GPUMemGiB,
		UnitPrice:    r.UnitPrice,
		Status:       r.Status,
		IsAvailable:  r.IsAvailable,
		UpdatedTime:  r.UpdatedTime,
	}
}

type ResourceSpecificationReq struct {
	QueueId         int64 `binding:"Required"`
	SourceSpecId    string
	AccCardsNum     int
	CpuCores        int
	MemGiB          float32
	GPUMemGiB       float32
	ShareMemGiB     float32
	UnitPrice       float64
	Status          int
	IsAutomaticSync bool
	CreatorId       int64
}

func (r ResourceSpecificationReq) ToDTO() ResourceSpecification {
	return ResourceSpecification{
		QueueId:         r.QueueId,
		SourceSpecId:    r.SourceSpecId,
		AccCardsNum:     r.AccCardsNum,
		CpuCores:        r.CpuCores,
		MemGiB:          r.MemGiB,
		GPUMemGiB:       r.GPUMemGiB,
		ShareMemGiB:     r.ShareMemGiB,
		UnitPrice:       r.UnitPrice,
		Status:          r.Status,
		IsAutomaticSync: r.IsAutomaticSync,
		CreatedBy:       r.CreatorId,
		UpdatedBy:       r.CreatorId,
		IsAvailable:     true,
	}
}

type SearchResourceSpecificationOptions struct {
	ListOptions
	QueueId             int64
	Status              int
	Cluster             string
	AiCenterCode        string
	AvailableCode       int
	OrderBy             SearchSpecOrderBy
	AccCardsNum         int
	ComputeResource     string
	AccCardType         string
	HasInternet         SpecInternetQuery
	EnableVisualization int
}
type SearchSpecLogOptions struct {
	Page            int
	PageSize        int
	SpecId          int    // 对应 RelatedId
	OperateType     string // 操作类型: on-shelf, off-shelf
	StartTime       string
	EndTime         string
	ComputeResource string
	AccCardType     string
	AccCardNum      int
	AiCenterCode    string
	ResourceQueueID int
}
type SearchResourceBriefSpecificationOptions struct {
	QueueId int64
	Cluster string
}

type ResourceSpecAndQueueListRes struct {
	TotalSize int64
	List      []*ResourceSpecAndQueueRes
}

func NewResourceSpecAndQueueListRes(totalSize int64, list []ResourceSpecAndQueue) *ResourceSpecAndQueueListRes {
	resList := make([]*ResourceSpecAndQueueRes, len(list))
	for i, v := range list {
		resList[i] = v.ConvertToRes()
	}
	return &ResourceSpecAndQueueListRes{
		TotalSize: totalSize,
		List:      resList,
	}
}

type ResourceSpecificationRes struct {
	ID           int64
	SourceSpecId string
	AccCardsNum  int
	CpuCores     int
	MemGiB       float32
	GPUMemGiB    float32
	ShareMemGiB  float32
	UnitPrice    float64
	Status       int
	IsAvailable  bool
	UpdatedTime  timeutil.TimeStamp
}

func (ResourceSpecificationRes) TableName() string {
	return "resource_specification"
}

type ResourceSpecAndQueueRes struct {
	Spec  *ResourceSpecificationRes
	Queue *ResourceQueueRes
}

type ResourceSpecAndQueue struct {
	ResourceSpecification `xorm:"extends"`
	ResourceQueue         `xorm:"extends"`
}

func (*ResourceSpecAndQueue) TableName() string {
	return "resource_specification"
}

func (r ResourceSpecAndQueue) ConvertToRes() *ResourceSpecAndQueueRes {
	return &ResourceSpecAndQueueRes{
		Spec:  r.ResourceSpecification.ConvertToRes(),
		Queue: r.ResourceQueue.ConvertToRes(),
	}
}

func (r ResourceSpecAndQueue) ConvertToResourceSpecInfo() *ResourceSpecInfo {
	return &ResourceSpecInfo{
		ID:                  r.ResourceSpecification.ID,
		SourceSpecId:        r.SourceSpecId,
		AccCardsNum:         r.AccCardsNum,
		CpuCores:            r.CpuCores,
		MemGiB:              r.MemGiB,
		GPUMemGiB:           r.GPUMemGiB,
		ShareMemGiB:         r.ShareMemGiB,
		UnitPrice:           r.UnitPrice,
		Status:              r.ResourceSpecification.Status,
		UpdatedTime:         r.ResourceSpecification.UpdatedTime,
		Cluster:             r.Cluster,
		AiCenterCode:        r.AiCenterCode,
		AiCenterName:        r.AiCenterName,
		QueueCode:           r.QueueCode,
		QueueType:           r.QueueType,
		QueueName:           r.QueueName,
		QueueId:             r.QueueId,
		ComputeResource:     r.ComputeResource,
		AccCardType:         r.AccCardType,
		HasInternet:         r.HasInternet,
		EnableVisualization: r.EnableVisualization,
	}
}

type FindSpecsOptions struct {
	JobType         JobType
	ComputeResource string
	Cluster         string
	AiCenterCode    string
	SpecId          int64
	QueueCode       string
	SourceSpecId    string
	AccCardsNum     int
	UseAccCardsNum  bool
	AccCardType     string
	CpuCores        int
	UseCpuCores     bool
	MemGiB          float32
	UseMemGiB       bool
	GPUMemGiB       float32
	UseGPUMemGiB    bool
	ShareMemGiB     float32
	UseShareMemGiB  bool
	//if true,find specs no matter used or not used in scene. if false,only find specs used in scene
	RequestAll        bool
	SpecStatus        int
	HasInternet       SpecInternetQuery //0 all,1 no internet,2 has internet
	SceneType         string
	Ids               []int64
	VisualizeRequired bool
}

type QueueCapabilityAggregate struct {
	// 联网能力
	NetworkCapableQueuesExist   bool // true: 存在能联网的队列
	NetworkIncapableQueuesExist bool // true: 存在不能联网的队列

	// 可视化能力
	VisualizeCapableQueuesExist   bool // true: 存在能可视化的队列
	VisualizeIncapableQueuesExist bool // true: 存在不能可视化的队列
}

type AggregateSpecification struct {
	ID                  int64
	SourceSpecId        string
	AccCardsNum         int
	AccCardType         string
	CpuCores            int
	MemGiB              float32
	GPUMemGiB           float32
	ShareMemGiB         float32
	ComputeResource     string
	UnitPrice           float64
	QueueId             int64
	QueueCode           string
	QueueName           string
	QueueType           string
	HasInternet         int
	EnableVisualization bool
	Cluster             string
	AiCenterCode        string
	AiCenterName        string
	SpecExclusiveType   string `xorm:"is_spec_exclusive"`
	ExclusiveOrg        string
	Remark              string
	Capability          QueueCapabilityAggregate
}

type Specification struct {
	ID                  int64
	SourceSpecId        string
	AccCardsNum         int
	AccCardType         string
	CpuCores            int
	MemGiB              float32
	GPUMemGiB           float32
	ShareMemGiB         float32
	ComputeResource     string
	UnitPrice           float64
	QueueId             int64
	QueueCode           string
	QueueName           string
	QueueType           string
	HasInternet         int
	EnableVisualization bool
	Cluster             string
	AiCenterCode        string
	AiCenterName        string
	SpecExclusiveType   string `xorm:"is_spec_exclusive"`
	ExclusiveOrg        string
	Remark              string
}

func (s Specification) IsSpecExclusive() bool {
	return s.SpecExclusiveType == SpecExclusive
}

func (Specification) TableName() string {
	return "resource_specification"
}

func (s *Specification) findRelatedSpecs(opts FindSpecsOptions, userId int64) []*Specification {
	defaultSpecs := make([]*Specification, 0)
	if s.SourceSpecId == "" {
		return defaultSpecs
	}

	r, err := FindSpecsForNewRight(opts)
	if err != nil {
		return defaultSpecs
	}
	newSpecs, err := filterSpecsByUser(r, opts.JobType, userId)
	return newSpecs
}

func filterSpecsByUser(sourceSpecs []*Specification, jobType JobType, userId int64) ([]*Specification, error) {
	newSpecs := make([]*Specification, 0)
	re, err := QueryAiforgeUserRoleByUserId(userId)
	if err != nil {
		log.Error("QueryAiforgeUserRoleByUserId err.userId=%d err=%v", userId, err)
		return newSpecs, err
	}
	ids := make([]int64, 0)
	minusRoleIdMap := make(map[int64]int64, 0)
	for _, tmp := range re {
		ids = append(ids, tmp.RoleId)
		if tmp.Type == MinusType {
			minusRoleIdMap[tmp.RoleId] = 1
		}
	}
	roles, err := QueryAiforgeRoleByIdsAndCommonRole(ids)
	specIdMap := make(map[int64]int64, 0)
	if err == nil {
		for _, tmp := range roles {
			if minusRoleIdMap[tmp.ID] == 1 {
				continue
			}
			setAvailSpecIdToMap(tmp, string(jobType), specIdMap)
		}
	}
	for specId, _ := range specIdMap {
		for _, spec := range sourceSpecs {
			if spec.ID == specId {
				newSpecs = append(newSpecs, spec)
			}
		}
	}
	return newSpecs, nil
}

func (s *Specification) ToShowString() string {
	var specName string
	specName += s.ComputeResource + ":" + fmt.Sprint(s.AccCardsNum) + "*" + s.AccCardType + ",CPU:" + fmt.Sprint(s.CpuCores)
	if s.GPUMemGiB > 0 {
		specName += ",显存:" + fmt.Sprint(s.GPUMemGiB) + "GB"
	}
	if s.MemGiB > 0 {
		specName += ",内存:" + fmt.Sprint(s.MemGiB) + "GB"
	}
	if s.ShareMemGiB > 0 {
		specName += ",共享内存:" + fmt.Sprint(s.ShareMemGiB) + "GB"
	}
	return specName
}

func (s *Specification) ParseResourceQueue() ResourceQueue {
	return ResourceQueue{
		ID:              s.QueueId,
		QueueCode:       s.QueueCode,
		QueueName:       s.QueueName,
		QueueType:       s.QueueType,
		Cluster:         s.Cluster,
		AiCenterCode:    s.AiCenterCode,
		AiCenterName:    s.AiCenterName,
		ComputeResource: s.ComputeResource,
		AccCardType:     s.AccCardType,
		CardsTotalNum:   s.AccCardsNum,
		HasInternet:     s.HasInternet,
	}
}

func (s *Specification) ParseQueueCapability() QueueCapabilityAggregate {
	return QueueCapabilityAggregate{
		NetworkCapableQueuesExist:     s.HasInternet == int(HasInternet),
		NetworkIncapableQueuesExist:   s.HasInternet != int(HasInternet),
		VisualizeCapableQueuesExist:   s.EnableVisualization,
		VisualizeIncapableQueuesExist: !s.EnableVisualization,
	}
}

func GetAvailableCenterIdsByASpec(ID int64) ([]string, error) {
	spec, err := GetResourceSpecification(&ResourceSpecification{
		ID: ID})
	if err != nil {
		return []string{}, err
	}
	var queueIds []int64
	err = x.Table("resource_specification").Cols("queue_id").Where("status=? and source_spec_id=?", SpecOnShelf, spec.SourceSpecId).Find(&queueIds)
	if err != nil || len(queueIds) == 0 {
		return []string{}, err
	}
	var centerIds []string
	err = x.Table("resource_queue").Cols("ai_center_code").In("id", queueIds).Find(&centerIds)
	return centerIds, err

}

type GetAvailableCenterIdOpts struct {
	UserId            int64
	JobType           JobType
	HasInternet       SpecInternetQuery
	VisualizeRequired bool
}

func (s *Specification) GetAvailableCenterIds(opts GetAvailableCenterIdOpts) []string {
	queues := s.GetAvailableQueuesForNewRight(opts)
	centerIds := make([]string, 0)
	for _, v := range queues {
		centerIds = append(centerIds, v.AiCenterCode)
	}
	return centerIds
}

func (s *Specification) GetAvailableQueuesForNewRight(opts GetAvailableCenterIdOpts) []ResourceQueue {
	//是否需要网络的调度策略如下：
	//需要联网时只能调度到有网的分中心；不需要联网时可以调度到所有的分中心
	hasInternet := opts.HasInternet
	if hasInternet == QueryNoInternetSpecs {
		hasInternet = QueryAllSpecs
	}

	log.Info("GetAvailableQueuesForNewRight opts.JobType=" + string(opts.JobType))
	relatedSpecs, err := QueryAvailResourceNotDistinct(FindAvailSpecsOptions{
		UserId:          opts.UserId,
		ComputeResource: s.ComputeResource,
		Cluster:         s.Cluster,
		JobType:         opts.JobType,
		SourceSpecId:    s.SourceSpecId,
		HasInternet:     hasInternet,
	})
	if err != nil {
		log.Error("GetAvailableQueuesForNewRight error.%v", err)
		return make([]ResourceQueue, 0)
	}

	if len(relatedSpecs) == 0 {
		log.Info("check centerIds opt%+v relatedSpecs is empty", opts)
		return make([]ResourceQueue, 0)
	}

	queueMap := make(map[int64]string, len(relatedSpecs))
	queues := make([]ResourceQueue, 0)
	for _, v := range relatedSpecs {
		if opts.VisualizeRequired {
			if !v.EnableVisualization {
				continue
			}
		}
		if _, ok := queueMap[v.QueueId]; ok {
			continue
		}
		queues = append(queues, v.ParseResourceQueue())
		log.Info("v.QueueId=" + fmt.Sprint(v.QueueId))
		queueMap[v.QueueId] = ""
	}
	return queues
}

func (s *Specification) GetAvailableQueues(opts GetAvailableCenterIdOpts) []ResourceQueue {
	//是否需要网络的调度策略如下：
	//需要联网时只能调度到有网的分中心；不需要联网时可以调度到所有的分中心
	hasInternet := opts.HasInternet
	if hasInternet == QueryNoInternetSpecs {
		hasInternet = QueryAllSpecs
	}

	specOpts := FindSpecsOptions{
		ComputeResource:   s.ComputeResource,
		Cluster:           s.Cluster,
		SourceSpecId:      s.SourceSpecId,
		RequestAll:        false,
		SpecStatus:        SpecOnShelf,
		JobType:           opts.JobType,
		HasInternet:       hasInternet,
		VisualizeRequired: opts.VisualizeRequired,
	}
	relatedSpecs := s.findRelatedSpecs(specOpts, opts.UserId)

	if len(relatedSpecs) == 0 {
		log.Info("check centerIds opt%+v relatedSpecs is empty", opts)
		return make([]ResourceQueue, 0)
	}

	queueMap := make(map[int64]string, len(relatedSpecs))
	queues := make([]ResourceQueue, 0)
	for _, v := range relatedSpecs {
		if _, ok := queueMap[v.QueueId]; ok {
			continue
		}
		queues = append(queues, v.ParseResourceQueue())
		queueMap[v.QueueId] = ""
	}
	return queues
}

func FilterExclusiveSpecs(r []*Specification, userId int64) []*Specification {
	if userId == 0 {
		return r
	}
	specs := make([]*Specification, 0, len(r))
	specMap := make(map[int64]string, 0)
	for i := 0; i < len(r); i++ {
		spec := r[i]
		if _, has := specMap[spec.ID]; has {
			continue
		}
		if !spec.IsSpecExclusive() {
			specs = append(specs, spec)
			specMap[spec.ID] = ""
			continue
		}
		orgs := strings.Split(spec.ExclusiveOrg, ";")
		for _, org := range orgs {
			isMember, _ := IsOrganizationMemberByOrgName(org, userId)
			if isMember {
				specs = append(specs, spec)
				specMap[spec.ID] = ""
				break
			}
		}
	}
	return specs
}

func GetSpecsInQueue(r []*Specification, queueIds []int64) []*Specification {
	specs := make([]*Specification, 0, len(r))
	for i := 0; i < len(r); i++ {
		spec := r[i]
		for _, queueId := range queueIds {
			if spec.QueueId == queueId {
				specs = append(specs, spec)
				break
			}
		}
	}
	return specs
}

func GetSpecsNotInQueue(r []*Specification, queueIds []int64) []*Specification {
	specs := make([]*Specification, 0, len(r))
	for i := 0; i < len(r); i++ {
		spec := r[i]
		flag := false
		for _, queueId := range queueIds {
			if spec.QueueId == queueId {
				flag = true
				break
			}
		}
		if !flag {
			specs = append(specs, spec)
		}
	}
	return specs
}

func HandleSpecialQueues(specs []*Specification, userId int64, opts FindSpecsOptions) []*Specification {
	if len(specs) == 0 {
		return specs
	}
	isUserInSpecialPool := IsUserInExclusivePool(userId)
	if isUserInSpecialPool {
		specs = handleExclusiveUserSpecs(specs, userId, opts)
	} else {
		specs = handleNormalUserSpecs(specs)
	}
	return specs
}

func handleExclusiveUserSpecs(specs []*Specification, userId int64, opts FindSpecsOptions) []*Specification {
	specialQueues := GetExclusiveQueueIds(opts)
	userOrgs, err := GetOrgsByUserID(userId, true)
	if err != nil {
		log.Error("handleSpecialUserSpecs GetOrgsByUserID error.%v", err)
		return []*Specification{}
	}
	userQueueIds := make([]int64, 0)
	for _, org := range userOrgs {
		for _, queue := range specialQueues {
			if strings.ToLower(org.Name) == strings.ToLower(queue.OrgName) {
				userQueueIds = append(userQueueIds, queue.QueueId)
			}
		}
	}
	specs = GetSpecsInQueue(specs, userQueueIds)
	return specs
}

func handleNormalUserSpecs(specs []*Specification) []*Specification {
	queues := GetAllExclusiveQueueIds()
	queueIds := make([]int64, 0)
	for _, queue := range queues {
		queueIds = append(queueIds, queue.QueueId)
	}
	specs = GetSpecsNotInQueue(specs, queueIds)
	return specs
}

func DistinctSpecs(r []*Specification) []*Specification {
	specs := make([]*Specification, 0, len(r))
	sourceSpecIdMap := make(map[string]string, 0)
	for i := 0; i < len(r); i++ {
		spec := r[i]
		if spec.SourceSpecId == "" {
			specs = append(specs, spec)
			continue
		}
		if _, has := sourceSpecIdMap[spec.SourceSpecId]; has {
			continue
		}
		specs = append(specs, spec)
		sourceSpecIdMap[spec.SourceSpecId] = ""
	}
	return specs
}

func InsertResourceSpecification(r ResourceSpecification) (int64, error) {
	return x.Insert(&r)
}

func UpdateResourceSpecificationById(queueId int64, spec ResourceSpecification) (int64, error) {
	return x.ID(queueId).Update(&spec)
}
func UpdateSpecUnitPriceById(id int64, unitPrice float64) error {
	_, err := x.Exec("update resource_specification set unit_price = ? ,updated_time = ? where id = ?", unitPrice, timeutil.TimeStampNow(), id)
	return err
}

func SearchResourceSpecification(opts SearchResourceSpecificationOptions) (int64, []ResourceSpecAndQueue, error) {
	var cond = builder.NewCond()
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.QueueId > 0 {
		cond = cond.And(builder.Eq{"resource_specification.queue_id": opts.QueueId})
	}
	if opts.Status > 0 {
		cond = cond.And(builder.Eq{"resource_specification.status": opts.Status})
	}
	if opts.Cluster != "" {
		cond = cond.And(builder.Eq{"resource_queue.cluster": opts.Cluster})
	}
	if opts.AiCenterCode != "" {
		cond = cond.And(builder.Eq{"resource_queue.ai_center_code": opts.AiCenterCode})
	}
	if opts.AccCardsNum >= 0 {
		cond = cond.And(builder.Eq{"resource_specification.acc_cards_num": opts.AccCardsNum})
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"resource_queue.compute_resource": opts.ComputeResource})
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"resource_queue.acc_card_type": opts.AccCardType})
	}
	if opts.AvailableCode == 1 {
		cond = cond.And(builder.Eq{"resource_specification.is_available": true})
	} else if opts.AvailableCode == 2 {
		cond = cond.And(builder.Eq{"resource_specification.is_available": false})
	}

	if opts.HasInternet == QueryNoInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": NoInternet})
	} else if opts.HasInternet == QueryHasInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": HasInternet})
	}
	if opts.EnableVisualization > 0 {
		if opts.EnableVisualization == 1 {
			cond = cond.And(builder.Neq{"resource_queue.enable_visualization": true})
		} else if opts.EnableVisualization == 2 {
			cond = cond.And(builder.Eq{"resource_queue.enable_visualization": true})
		}
	}
	//cond = cond.And(builder.Or(builder.Eq{"resource_queue.deleted_time": 0}).Or(builder.IsNull{"resource_queue.deleted_time"}))
	n, err := x.Where(cond).Join("INNER", "resource_queue", "resource_queue.ID = resource_specification.queue_id").
		Unscoped().Count(&ResourceSpecAndQueue{})
	if err != nil {
		return 0, nil, err
	}

	var orderby = ""
	switch opts.OrderBy {
	case SearchSpecOrder4Standard:
		orderby = "resource_queue.compute_resource asc,resource_queue.acc_card_type asc,resource_specification.acc_cards_num asc,resource_specification.cpu_cores asc,resource_specification.mem_gi_b asc,resource_specification.share_mem_gi_b asc"
	default:
		orderby = "resource_specification.id desc"
	}

	r := make([]ResourceSpecAndQueue, 0)
	err = x.Where(cond).
		Join("INNER", "resource_queue", "resource_queue.ID = resource_specification.queue_id").
		OrderBy(orderby).
		Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).
		Unscoped().Find(&r)
	if err != nil {
		return 0, nil, err
	}
	return n, r, nil
}

func GetSpecScenes(specId int64) ([]ResourceSceneBriefRes, error) {
	r := make([]ResourceSceneBriefRes, 0)
	err := x.Where("resource_scene_spec.spec_id = ?", specId).
		Join("INNER", "resource_scene_spec", "resource_scene_spec.scene_id = resource_scene.id").
		Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func ResourceSpecOnShelf(id int64, unitPrice float64) error {
	// 创建更新字段的map
	updateFields := map[string]interface{}{
		"updated_time": timeutil.TimeStampNow(),
		"status":       SpecOnShelf,
	}

	// 只有当unitPrice不为0时才添加到更新字段中
	if unitPrice != 0 {
		updateFields["unit_price"] = unitPrice
	}

	_, err := x.Table("resource_specification").Where("id = ?", id).Update(updateFields)

	if err != nil {
		return err
	}

	return nil
}

func ResourceSpecOffShelf(id int64) (int64, error) {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	param := ResourceSpecification{
		Status: SpecOffShelf,
	}
	n, err := sess.Where("id = ? and status = ?", id, SpecOnShelf).Update(&param)
	if err != nil {
		return 0, err
	}
	sess.Commit()
	return n, err
}

func GetResourceSpecificationByIds(ids []int64) ([]*Specification, error) {
	r := make([]*Specification, 0)
	err := x.In("resource_specification.id", ids).
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Find(&r)
	return r, err

}

func GetResourceSpecification(r *ResourceSpecification) (*ResourceSpecification, error) {
	has, err := x.Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}

func GetSpecificationById(id int64) (*Specification, error) {
	r := &Specification{}
	has, err := x.Where("resource_specification.id = ?", id).
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, err

}

func SyncGrampusSpecs(updateList []ResourceSpecification, insertList []ResourceSpecification, existIds []int64) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()
	//delete specs and scene that no longer exists
	deleteIds := make([]int64, 0)
	cond := builder.NewCond()
	cond = cond.And(builder.NotIn("resource_specification.id", existIds)).And(builder.Eq{"resource_queue.cluster": C2NetCluster})
	if err := sess.Cols("resource_specification.id").Table("resource_specification").
		Where(cond).Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Find(&deleteIds); err != nil {
		return err
	}
	if len(deleteIds) > 0 {
		if _, err = sess.Cols("status", "is_available").In("id", deleteIds).Update(&ResourceSpecification{Status: SpecOffShelf, IsAvailable: false}); err != nil {
			return err
		}
	}

	//update exists specs
	// Update existing specs
	if len(updateList) > 0 {

		for _, v := range updateList {
			// 1. 查询数据库中旧的记录
			var oldSpec ResourceSpecification
			has, err := sess.ID(v.ID).Get(&oldSpec)
			if err != nil {
				log.Error("查询旧规格数据失败, ID: %d, Error: %v", v.ID, err)
				return err
			}
			if !has {
				log.Warn("规格记录不存在，跳过更新, ID: %d", v.ID)
				continue
			}

			//执行更新操作
			if _, err = sess.ID(v.ID).UseBool("is_available").Update(&v); err != nil {
				log.Error("更新资源规格失败, ID: %d, Error: %v", v.ID, err)
				return err
			}

			// 比较 is_available 字段是否发生变化
			// 只有当旧值 != 新值时，才记录日志
			if oldSpec.IsAvailable != v.IsAvailable {
				// 记录操作日志 (更新)
				spec, err := GetResourceSpecification(&ResourceSpecification{ID: v.ID})
				if err != nil {
					log.Error("GetResourceSpecification failed: %v", err)
				}

				ResourceQueue, err := GetResourceQueueByID(spec.QueueId)
				if err != nil {
					log.Error("GetResourceQueue error.%v", err)
				}

				InsertAdminOperateLog(AdminOperateLog{
					BizType:         "SpecOperate",
					OperateType:     "auto-update",
					RelatedId:       fmt.Sprint(v.ID),
					CreatedBy:       -1, // 表示系统自动同步,
					Comment:         fmt.Sprintf("自动同步更新规格配置"),
					ComputeResource: ResourceQueue.ComputeResource,
					AccCardType:     ResourceQueue.AccCardType,
					AccCardNum:      spec.AccCardsNum,
					AiCenterCode:    ResourceQueue.AiCenterCode,
					AiCenterName:    ResourceQueue.AiCenterName,
					ResourceQueueID: ResourceQueue.ID,
				})
			}
		}
	}

	//insert new specs
	if len(insertList) > 0 {
		if _, err = sess.Insert(insertList); err != nil {
			return err
		}
		for _, v := range insertList {
			// 记录操作日志 (更新)
			spec, err := GetResourceSpecification(&ResourceSpecification{ID: v.ID})
			if err != nil {
				log.Error("GetResourceSpecification failed: %v", err)
			}

			ResourceQueue, err := GetResourceQueueByID(spec.QueueId)
			if err != nil {
				log.Error("GetResourceQueue error.%v", err)
			}

			InsertAdminOperateLog(AdminOperateLog{
				BizType:         "SpecOperate",
				OperateType:     "auto-update",
				RelatedId:       fmt.Sprint(v.ID),
				CreatedBy:       -1, // 表示系统自动同步,
				Comment:         fmt.Sprintf("自动同步更新规格配置"),
				ComputeResource: ResourceQueue.ComputeResource,
				AccCardType:     ResourceQueue.AccCardType,
				AccCardNum:      spec.AccCardsNum,
				AiCenterCode:    ResourceQueue.AiCenterCode,
				AiCenterName:    ResourceQueue.AiCenterName,
				ResourceQueueID: ResourceQueue.ID,
			})
		}
	}

	return sess.Commit()
}

// FindSpecs
func FindSpecs(opts FindSpecsOptions) ([]*Specification, error) {
	var cond = builder.NewCond()
	if !opts.RequestAll && opts.JobType != "" {
		cond = cond.And(builder.Eq{"resource_scene.job_type": opts.JobType})
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"resource_queue.compute_resource": opts.ComputeResource})
	}
	if opts.Cluster != "" {
		cond = cond.And(builder.Eq{"resource_queue.cluster": opts.Cluster})
	}
	if opts.AiCenterCode != "" {
		cond = cond.And(builder.Eq{"resource_queue.ai_center_code": opts.AiCenterCode})
	}
	if opts.SpecId > 0 {
		cond = cond.And(builder.Eq{"resource_specification.id": opts.SpecId})
	}
	if opts.QueueCode != "" {
		cond = cond.And(builder.Eq{"resource_queue.queue_code": opts.QueueCode})
	}
	if opts.SourceSpecId != "" {
		cond = cond.And(builder.Eq{"resource_specification.source_spec_id": opts.SourceSpecId})
	}
	if opts.UseAccCardsNum {
		cond = cond.And(builder.Eq{"resource_specification.acc_cards_num": opts.AccCardsNum})
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"resource_queue.acc_card_type": opts.AccCardType})
	}
	if opts.UseCpuCores {
		cond = cond.And(builder.Eq{"resource_specification.cpu_cores": opts.CpuCores})
	}
	if opts.UseMemGiB {
		cond = cond.And(builder.Eq{"resource_specification.mem_gi_b": opts.MemGiB})
	}
	if opts.UseGPUMemGiB {
		cond = cond.And(builder.Eq{"resource_specification.gpu_mem_gi_b": opts.GPUMemGiB})
	}
	if opts.UseShareMemGiB {
		cond = cond.And(builder.Eq{"resource_specification.share_mem_gi_b": opts.ShareMemGiB})
	}
	if opts.SpecStatus > 0 {
		cond = cond.And(builder.Eq{"resource_specification.status": opts.SpecStatus})
	}
	if opts.HasInternet == QueryNoInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": NoInternet})
	} else if opts.HasInternet == QueryHasInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": HasInternet})
	}
	if opts.SceneType != "" {
		cond = cond.And(builder.Eq{"resource_scene.scene_type": opts.SceneType})
	}

	r := make([]*Specification, 0)
	s := x.Where(cond).
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id")

	if !opts.RequestAll {
		s = s.Join("INNER", "resource_scene_spec", "resource_scene_spec.spec_id = resource_specification.id").
			Join("INNER", "resource_scene", "resource_scene_spec.scene_id = resource_scene.id")
	}
	err := s.OrderBy("resource_queue.compute_resource asc,resource_queue.acc_card_type asc,resource_specification.acc_cards_num asc,resource_specification.cpu_cores asc,resource_specification.mem_gi_b asc,resource_specification.share_mem_gi_b asc").
		Unscoped().Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func FindSpecsForNewRight(opts FindSpecsOptions) ([]*Specification, error) {
	var cond = builder.NewCond()
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"resource_queue.compute_resource": opts.ComputeResource})
	}
	if opts.Cluster != "" {
		cond = cond.And(builder.Eq{"resource_queue.cluster": opts.Cluster})
	}
	if opts.Ids != nil && len(opts.Ids) > 0 {
		cond = cond.And(builder.In("resource_specification.id", opts.Ids))
	}
	if opts.SpecStatus > 0 {
		cond = cond.And(builder.Eq{"resource_specification.status": opts.SpecStatus})
	}
	if opts.SourceSpecId != "" {
		cond = cond.And(builder.Eq{"resource_specification.source_spec_id": opts.SourceSpecId})
	}
	if opts.HasInternet == QueryNoInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": NoInternet})
	} else if opts.HasInternet == QueryHasInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": HasInternet})
	}
	if opts.VisualizeRequired {
		cond = cond.And(builder.Eq{"resource_queue.enable_visualization": true})
	}

	r := make([]*Specification, 0)
	s := x.Where(cond).
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id")

	err := s.OrderBy("resource_queue.compute_resource asc,resource_queue.acc_card_type asc,resource_specification.acc_cards_num asc,resource_specification.cpu_cores asc,resource_specification.mem_gi_b asc,resource_specification.share_mem_gi_b asc").
		Unscoped().Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func InitQueueAndSpec(queue ResourceQueue, spec ResourceSpecification) (*Specification, error) {
	sess := x.NewSession()
	defer sess.Close()

	sess.Begin()
	param := ResourceQueue{
		QueueCode:       queue.QueueCode,
		Cluster:         queue.Cluster,
		AiCenterCode:    queue.AiCenterCode,
		ComputeResource: queue.ComputeResource,
		AccCardType:     queue.AccCardType,
	}
	_, err := sess.Get(&param)
	if err != nil {
		sess.Rollback()
		return nil, err
	}
	if param.ID == 0 {
		_, err = sess.InsertOne(&queue)
		if err != nil {
			sess.Rollback()
			return nil, err
		}
	} else {
		queue = param
	}

	spec.QueueId = queue.ID
	_, err = sess.InsertOne(&spec)
	if err != nil {
		sess.Rollback()
		return nil, err
	}
	sess.Commit()
	return BuildSpecification(queue, spec), nil
}

func BuildSpecification(queue ResourceQueue, spec ResourceSpecification) *Specification {
	return &Specification{
		ID:              spec.ID,
		SourceSpecId:    spec.SourceSpecId,
		AccCardsNum:     spec.AccCardsNum,
		AccCardType:     queue.AccCardType,
		CpuCores:        spec.CpuCores,
		MemGiB:          spec.MemGiB,
		GPUMemGiB:       spec.GPUMemGiB,
		ShareMemGiB:     spec.ShareMemGiB,
		ComputeResource: queue.ComputeResource,
		UnitPrice:       spec.UnitPrice,
		QueueId:         queue.ID,
		QueueCode:       queue.QueueCode,
		Cluster:         queue.Cluster,
		AiCenterCode:    queue.AiCenterCode,
		AiCenterName:    queue.AiCenterName,
	}
}

func GetCloudbrainOneAccCardType(queueCode string) string {
	switch queueCode {
	case "a100":
		return "A100"
	case "openidebug":
		return "T4"
	case "openidgx":
		return "V100"

	}
	return ""
}

var cloudbrainTwoSpecsInitFlag = false
var cloudbrainTwoSpecs map[string]*Specification

func GetCloudbrainTwoSpecs() (map[string]*Specification, error) {
	if !cloudbrainTwoSpecsInitFlag {
		r, err := InitCloudbrainTwoSpecs()
		if err != nil {
			return nil, err
		}
		cloudbrainTwoSpecsInitFlag = true
		cloudbrainTwoSpecs = r
	}
	return cloudbrainTwoSpecs, nil
}

func InitCloudbrainTwoSpecs() (map[string]*Specification, error) {
	r := make(map[string]*Specification, 0)

	queue, err := GetResourceQueue(&ResourceQueue{QueueCode: "openisupport"})
	if err != nil {
		return nil, err
	}
	if queue == nil {
		queue = &ResourceQueue{
			QueueCode:       "openisupport",
			Cluster:         OpenICluster,
			AiCenterCode:    AICenterOfCloudBrainTwo,
			AiCenterName:    "云脑二",
			ComputeResource: NPU,
			AccCardType:     "ASCEND910",
			Remark:          "处理历史云脑任务时自动生成",
		}
		_, err = x.InsertOne(queue)
		if err != nil {
			return nil, err
		}
	}
	for i := 1; i <= 8; i = i * 2 {
		sourceSpecId := "modelarts.bm.910.arm.public." + fmt.Sprint(i)
		spec, err := GetResourceSpecification(&ResourceSpecification{
			SourceSpecId: sourceSpecId,
			QueueId:      queue.ID,
		})
		if err != nil {
			return nil, err
		}
		if spec == nil {
			spec = &ResourceSpecification{
				QueueId:      queue.ID,
				SourceSpecId: sourceSpecId,
				AccCardsNum:  i,
				CpuCores:     i * 24,
				MemGiB:       float32(i * 256),
				GPUMemGiB:    float32(32),
				Status:       SpecOffShelf,
				IsAvailable:  true,
			}
			_, err = x.Insert(spec)
			if err != nil {
				return nil, err
			}
		}
		r[sourceSpecId] = BuildSpecification(*queue, *spec)
	}
	return r, nil
}

var grampusSpecsInitFlag = false
var grampusSpecs map[string]*Specification

func GetGrampusSpecs() (map[string]*Specification, error) {
	if !grampusSpecsInitFlag {
		specMap := make(map[string]*Specification, 0)
		r, err := FindSpecs(FindSpecsOptions{
			Cluster:    C2NetCluster,
			RequestAll: true,
		})
		if err != nil {
			return nil, err
		}
		for _, spec := range r {
			specMap[spec.SourceSpecId] = spec
			specMap[spec.SourceSpecId+"_"+spec.AiCenterCode] = spec
		}
		grampusSpecsInitFlag = true
		grampusSpecs = specMap
	}
	return grampusSpecs, nil
}

type GetResourceListOpts struct {
	ListOptions
	Resource           []string
	AccCardType        string
	AccCardNum         int
	ExcludeAccCardNums []int
	AICenterCode       string
	MinPrice           float64
	MaxPrice           float64
}

type ResourceDetailInfo struct {
	Spec             ResourceSpecificationRes
	IsQueueExclusive bool
	AICenterList     []ResourceAiCenterRes
}

type ResourceInfo4CardRequest struct {
	ComputeResource string
	AccCardType     string
	AccCardsNum     int
	CpuCores        int
	MemGiB          float32
	GPUMemGiB       float32
	ShareMemGiB     float32
	UnitPrice       float64
	IsExclusive     bool
	IsSpecExclusive string
	AICenterList    []*ResourceAiCenterRes
}

func (r *ResourceInfo4CardRequest) Tr(language string) {
	if r.AICenterList == nil {
		return
	}
	for i := 0; i < len(r.AICenterList); i++ {
		r.AICenterList[i].Tr(language)
	}
}

type ResourceWithAICenter4CardRequest struct {
	Cluster         string
	AICenterCode    string
	AICenterName    string
	ComputeResource string
	AccCardType     string
	AccCardsNum     int
	CpuCores        int
	MemGiB          float32
	GPUMemGiB       float32
	ShareMemGiB     float32
	UnitPrice       float64
	IsExclusive     bool
	IsSpecExclusive string
}

func GetResourceListPaging(opts GetResourceListOpts) ([]*ResourceInfo4CardRequest, int64, error) {
	cond := builder.NewCond()
	resourceList := make([]string, 0)
	for i := 0; i < len(opts.Resource); i++ {
		if opts.Resource[i] != "" {
			resourceList = append(resourceList, opts.Resource[i])
		}
	}
	if len(resourceList) > 0 {
		cond = cond.And(builder.In("resource_queue.compute_resource", resourceList))
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"resource_queue.acc_card_type": opts.AccCardType})
	}
	if opts.AccCardNum >= 0 {
		if opts.AccCardNum > 999 {
			cond = cond.And(builder.NotIn("resource_specification.acc_cards_num", opts.ExcludeAccCardNums))
		} else {
			cond = cond.And(builder.Eq{"resource_specification.acc_cards_num": opts.AccCardNum})
		}
	}
	if opts.AICenterCode != "" {
		cond = cond.And(builder.Eq{"resource_queue.ai_center_code": opts.AICenterCode})
	}
	if opts.MaxPrice >= 0 && opts.MinPrice >= 0 && opts.MaxPrice < opts.MinPrice {
		opts.MaxPrice = -1
		opts.MinPrice = -1
	}
	if opts.MaxPrice >= 0 {
		cond = cond.And(builder.Lte{"resource_specification.unit_price": opts.MaxPrice})
	}
	if opts.MinPrice >= 0 {
		cond = cond.And(builder.Gte{"resource_specification.unit_price": opts.MinPrice})
	}
	cond = cond.And(builder.Or(builder.Eq{"resource_queue.deleted_time": 0}, builder.IsNull{"resource_queue.deleted_time"}))
	cond = cond.And(builder.Eq{"resource_specification.status": 2})
	//先按多字段去重分页查询资源规格
	//再基于结果查询智算中心信息
	resourceInfos := make([]*ResourceInfo4CardRequest, 0)
	err := x.Table("resource_specification").
		Join("LEFT", "resource_exclusive_pool", "resource_specification.queue_id = resource_exclusive_pool.queue_id").
		Join("INNER", "resource_queue", "resource_specification.queue_id = resource_queue.id").
		Join("INNER", "resource_scene_spec", "resource_specification.id = resource_scene_spec.spec_id").
		Join("INNER", "resource_scene", "resource_scene.id = resource_scene_spec.scene_id").
		Select("Distinct resource_queue.compute_resource, resource_queue.acc_card_type," +
			"resource_specification.acc_cards_num, resource_specification.cpu_cores, resource_specification.mem_gi_b, " +
			"resource_specification.gpu_mem_gi_b, resource_specification.share_mem_gi_b, resource_specification.unit_price," +
			"COALESCE(resource_exclusive_pool.queue_id IS NOT NULL, false) AS is_exclusive,resource_scene.is_spec_exclusive").
		Where(cond).
		OrderBy(" resource_queue.compute_resource DESC,resource_queue.acc_card_type ,is_exclusive," +
			"resource_specification.acc_cards_num DESC,resource_specification.gpu_mem_gi_b DESC," +
			"resource_specification.cpu_cores DESC,resource_specification.mem_gi_b DESC").
		Find(&resourceInfos)

	if err != nil {
		return nil, 0, err
	}
	tmpResourceInfos := make([]*ResourceInfo4CardRequest, 0)
	for i := 0; i < len(resourceInfos); i++ {
		//此处是为了过滤那些专属池中的规格又被配置到共享场景中的情况
		if !resourceInfos[i].IsExclusive || (resourceInfos[i].IsExclusive && resourceInfos[i].IsSpecExclusive == "") {
			tmpResourceInfos = append(tmpResourceInfos, resourceInfos[i])
		}
	}
	resourceInfos = tmpResourceInfos
	if len(resourceInfos) == 0 {
		return []*ResourceInfo4CardRequest{}, 0, nil
	}

	total := int64(len(resourceInfos))
	startIndex := int64((opts.Page - 1) * opts.PageSize)
	endIndex := int64(opts.Page * opts.PageSize)
	if startIndex >= total {
		return []*ResourceInfo4CardRequest{}, 0, nil
	}
	if endIndex > total {
		endIndex = total
	}
	resourceInfos = resourceInfos[startIndex:endIndex]

	newCond := builder.NewCond()
	for _, spec := range resourceInfos {
		if spec.IsExclusive {
			newCond = newCond.Or(builder.And(builder.Eq{"resource_queue.compute_resource": spec.ComputeResource},
				builder.Eq{"resource_queue.acc_card_type": spec.AccCardType},
				builder.Eq{"resource_specification.acc_cards_num": spec.AccCardsNum},
				builder.Eq{"resource_specification.cpu_cores": spec.CpuCores},
				builder.Eq{"resource_specification.mem_gi_b": spec.MemGiB},
				builder.Eq{"resource_specification.gpu_mem_gi_b": spec.GPUMemGiB},
				builder.Eq{"resource_specification.share_mem_gi_b": spec.ShareMemGiB},
				builder.Eq{"resource_specification.unit_price": spec.UnitPrice},
				builder.NotNull{"resource_exclusive_pool.queue_id"}))
		} else if spec.IsSpecExclusive == SpecExclusive {
			newCond = newCond.Or(builder.And(builder.Eq{"resource_queue.compute_resource": spec.ComputeResource},
				builder.Eq{"resource_queue.acc_card_type": spec.AccCardType},
				builder.Eq{"resource_specification.acc_cards_num": spec.AccCardsNum},
				builder.Eq{"resource_specification.cpu_cores": spec.CpuCores},
				builder.Eq{"resource_specification.mem_gi_b": spec.MemGiB},
				builder.Eq{"resource_specification.gpu_mem_gi_b": spec.GPUMemGiB},
				builder.Eq{"resource_specification.share_mem_gi_b": spec.ShareMemGiB},
				builder.Eq{"resource_specification.unit_price": spec.UnitPrice},
				builder.Eq{"resource_scene.is_spec_exclusive": SpecExclusive},
				builder.IsNull{"resource_exclusive_pool.queue_id"}))
		} else {
			newCond = newCond.Or(builder.And(builder.Eq{"resource_queue.compute_resource": spec.ComputeResource},
				builder.Eq{"resource_queue.acc_card_type": spec.AccCardType},
				builder.Eq{"resource_specification.acc_cards_num": spec.AccCardsNum},
				builder.Eq{"resource_specification.cpu_cores": spec.CpuCores},
				builder.Eq{"resource_specification.mem_gi_b": spec.MemGiB},
				builder.Eq{"resource_specification.gpu_mem_gi_b": spec.GPUMemGiB},
				builder.Eq{"resource_specification.share_mem_gi_b": spec.ShareMemGiB},
				builder.Eq{"resource_specification.unit_price": spec.UnitPrice},
				builder.Or(builder.Eq{"resource_scene.is_spec_exclusive": SpecPublic}, builder.IsNull{"resource_scene.is_spec_exclusive"}),
				builder.IsNull{"resource_exclusive_pool.queue_id"}))
		}

	}
	newCond = newCond.And(builder.Or(builder.Eq{"resource_queue.deleted_time": 0}, builder.IsNull{"resource_queue.deleted_time"}))
	newCond = newCond.And(builder.Eq{"resource_specification.status": 2})

	withCenterInfos := make([]ResourceWithAICenter4CardRequest, 0)
	err = x.Table("resource_specification").
		Join("LEFT", "resource_exclusive_pool", "resource_specification.queue_id = resource_exclusive_pool.queue_id").
		Join("INNER", "resource_queue", "resource_specification.queue_id = resource_queue.id").
		Join("INNER", "resource_scene_spec", "resource_specification.id = resource_scene_spec.spec_id").
		Join("INNER", "resource_scene", "resource_scene.id = resource_scene_spec.scene_id").
		Select("resource_queue.cluster,resource_queue.ai_center_code,resource_queue.ai_center_name,resource_queue.compute_resource, resource_queue.acc_card_type," +
			"resource_specification.acc_cards_num, resource_specification.cpu_cores, resource_specification.mem_gi_b, " +
			"resource_specification.gpu_mem_gi_b, resource_specification.share_mem_gi_b, resource_specification.unit_price," +
			"COALESCE(resource_exclusive_pool.queue_id IS NOT NULL, false) AS is_exclusive,resource_scene.is_spec_exclusive").
		Where(newCond).
		Find(&withCenterInfos)

	if err != nil {
		return nil, 0, err
	}
	tmpMap := make(map[string][]*ResourceAiCenterRes, 0)
	for i := 0; i < len(withCenterInfos); i++ {
		t := withCenterInfos[i]
		key := fmt.Sprintf("%s_%s_%d_%d_%f_%f_%f_%f_%t_%s", t.ComputeResource, t.AccCardType, t.AccCardsNum,
			t.CpuCores, t.MemGiB, t.GPUMemGiB, t.ShareMemGiB, t.UnitPrice, t.IsExclusive, t.IsSpecExclusive)
		if _, exists := tmpMap[key]; exists {
			centerExists := false
			for _, center := range tmpMap[key] {
				if center.AiCenterCode == t.AICenterCode {
					centerExists = true
				}
			}
			if centerExists {
				continue
			}
			tmpMap[key] = append(tmpMap[key], &ResourceAiCenterRes{
				AiCenterCode: t.AICenterCode,
				AiCenterName: t.AICenterName,
			})
		} else {
			tmpMap[key] = []*ResourceAiCenterRes{{
				AiCenterCode: t.AICenterCode,
				AiCenterName: t.AICenterName,
			}}
		}
	}

	for i := 0; i < len(resourceInfos); i++ {
		t := resourceInfos[i]
		key := fmt.Sprintf("%s_%s_%d_%d_%f_%f_%f_%f_%t_%s", t.ComputeResource, t.AccCardType, t.AccCardsNum,
			t.CpuCores, t.MemGiB, t.GPUMemGiB, t.ShareMemGiB, t.UnitPrice, t.IsExclusive, t.IsSpecExclusive)
		resourceInfos[i].AICenterList = tmpMap[key]
	}
	return resourceInfos, total, nil
}

func GetAllSpecIds() ([]int64, error) {
	roleList, err := ListCommonAiforgeRole()
	if err != nil {
		return []int64{}, err
	}

	specIds := GetSpecIdsByList(roleList)

	return specIds, nil
}

func GetResourceListPagingV2(opts GetResourceListOpts, specIds []int64) ([]*ResourceInfo4CardRequest, int64, error) {
	// step 1 构建查询的基本条件
	var (
		cond         = builder.NewCond()
		resourceList = make([]string, 0)
	)

	for i := 0; i < len(opts.Resource); i++ {
		if opts.Resource[i] != "" {
			resourceList = append(resourceList, opts.Resource[i])
		}
	}
	if len(resourceList) > 0 {
		cond = cond.And(builder.In("resource_queue.compute_resource", resourceList))
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"resource_queue.acc_card_type": opts.AccCardType})
	}
	if opts.AccCardNum >= 0 {
		if opts.AccCardNum > 999 {
			cond = cond.And(builder.NotIn("resource_specification.acc_cards_num", opts.ExcludeAccCardNums))
		} else {
			cond = cond.And(builder.Eq{"resource_specification.acc_cards_num": opts.AccCardNum})
		}
	}
	if opts.AICenterCode != "" {
		cond = cond.And(builder.Eq{"resource_queue.ai_center_code": opts.AICenterCode})
	}
	if opts.MaxPrice >= 0 && opts.MinPrice >= 0 && opts.MaxPrice < opts.MinPrice {
		opts.MaxPrice = -1
		opts.MinPrice = -1
	}
	if opts.MaxPrice >= 0 {
		cond = cond.And(builder.Lte{"resource_specification.unit_price": opts.MaxPrice})
	}
	if opts.MinPrice >= 0 {
		cond = cond.And(builder.Gte{"resource_specification.unit_price": opts.MinPrice})
	}

	if len(specIds) > 0 {
		cond = cond.And(builder.In("resource_specification.id", specIds))
	}

	// 删除时间
	cond = cond.And(builder.Or(builder.Eq{"resource_queue.deleted_time": 0}, builder.IsNull{"resource_queue.deleted_time"}))
	cond = cond.And(builder.Eq{"resource_specification.status": 2})

	//先按多字段去重分页查询资源规格
	//再基于结果查询智算中心信息
	resourceInfos := make([]*ResourceInfo4CardRequest, 0)
	err := x.Table("resource_specification").
		Join("INNER", "resource_queue", "resource_specification.queue_id = resource_queue.id").
		Select(`DISTINCT
        resource_queue.compute_resource,
        resource_queue.acc_card_type,
        resource_specification.acc_cards_num,
        resource_specification.cpu_cores,
        resource_specification.mem_gi_b,
        resource_specification.gpu_mem_gi_b,
        resource_specification.share_mem_gi_b,
        resource_specification.unit_price`).
		Where(cond).
		OrderBy(`resource_queue.compute_resource,
        resource_queue.acc_card_type,
        resource_specification.acc_cards_num DESC,
        resource_specification.gpu_mem_gi_b DESC,
        resource_specification.cpu_cores DESC,
        resource_specification.mem_gi_b DESC`).
		Find(&resourceInfos)

	if err != nil {
		return nil, 0, err
	}
	tmpResourceInfos := make([]*ResourceInfo4CardRequest, 0)

	// step 3 过滤专属池
	for i := 0; i < len(resourceInfos); i++ {
		//此处是为了过滤那些专属池中的规格又被配置到共享场景中的情况
		if !resourceInfos[i].IsExclusive || (resourceInfos[i].IsExclusive && resourceInfos[i].IsSpecExclusive == "") {
			tmpResourceInfos = append(tmpResourceInfos, resourceInfos[i])
		}
	}
	resourceInfos = tmpResourceInfos
	if len(resourceInfos) == 0 {
		return []*ResourceInfo4CardRequest{}, 0, nil
	}

	total := int64(len(resourceInfos))
	startIndex := int64((opts.Page - 1) * opts.PageSize)
	endIndex := int64(opts.Page * opts.PageSize)
	if startIndex >= total {
		return []*ResourceInfo4CardRequest{}, 0, nil
	}
	if endIndex > total {
		endIndex = total
	}
	resourceInfos = resourceInfos[startIndex:endIndex]

	// 为了找到中心的资源？
	newCond := builder.NewCond()

	newCond = newCond.And(builder.Or(builder.Eq{"resource_queue.deleted_time": 0}, builder.IsNull{"resource_queue.deleted_time"}))
	newCond = newCond.And(builder.Eq{"resource_specification.status": 2})
	newCond = newCond.And(builder.In("resource_specification.id", specIds))

	withCenterInfos := make([]ResourceWithAICenter4CardRequest, 0)
	err = x.Table("resource_specification").
		Join("INNER", "resource_queue", "resource_specification.queue_id = resource_queue.id").
		Select(`resource_queue.cluster,
        resource_queue.ai_center_code,
        resource_queue.ai_center_name,
        resource_queue.compute_resource,
        resource_queue.acc_card_type,
        resource_specification.acc_cards_num,
        resource_specification.cpu_cores,
        resource_specification.mem_gi_b,
        resource_specification.gpu_mem_gi_b,
        resource_specification.share_mem_gi_b,
        resource_specification.unit_price`).
		Where(newCond).
		Find(&withCenterInfos)

	if err != nil {
		return nil, 0, err
	}
	tmpMap := make(map[string][]*ResourceAiCenterRes, 0)
	for i := 0; i < len(withCenterInfos); i++ {
		t := withCenterInfos[i]
		key := fmt.Sprintf("%s_%s_%d_%d_%f_%f_%f_%f_%t_%s", t.ComputeResource, t.AccCardType, t.AccCardsNum,
			t.CpuCores, t.MemGiB, t.GPUMemGiB, t.ShareMemGiB, t.UnitPrice, t.IsExclusive, t.IsSpecExclusive)
		if _, exists := tmpMap[key]; exists {
			centerExists := false
			for _, center := range tmpMap[key] {
				if center.AiCenterCode == t.AICenterCode {
					centerExists = true
				}
			}
			if centerExists {
				continue
			}
			tmpMap[key] = append(tmpMap[key], &ResourceAiCenterRes{
				AiCenterCode: t.AICenterCode,
				AiCenterName: t.AICenterName,
			})
		} else {
			tmpMap[key] = []*ResourceAiCenterRes{{
				AiCenterCode: t.AICenterCode,
				AiCenterName: t.AICenterName,
			}}
		}
	}

	for i := 0; i < len(resourceInfos); i++ {
		t := resourceInfos[i]
		key := fmt.Sprintf("%s_%s_%d_%d_%f_%f_%f_%f_%t_%s", t.ComputeResource, t.AccCardType, t.AccCardsNum,
			t.CpuCores, t.MemGiB, t.GPUMemGiB, t.ShareMemGiB, t.UnitPrice, t.IsExclusive, t.IsSpecExclusive)
		resourceInfos[i].AICenterList = tmpMap[key]
	}
	return resourceInfos, total, nil
}

type AccCardInfo struct {
	ComputeSource string
	CardList      []string
}

type AccCardInfoList []AccCardInfo

func (a AccCardInfoList) Len() int {
	return len(a)
}

func (a AccCardInfoList) Less(i, j int) bool {
	return strings.Compare(a[i].ComputeSource, a[j].ComputeSource) < 0
}

func (a AccCardInfoList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func GetAccCardList() ([]AccCardInfo, error) {
	res := make([]AccCardInfo, 0)
	r := make([]*Specification, 0)
	err := x.Where("resource_specification.status = ? and (resource_queue.deleted_time = 0 or resource_queue.deleted_time is null)", SpecOnShelf).
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Join("INNER", "resource_scene_spec", "resource_scene_spec.spec_id = resource_specification.id").
		Join("INNER", "resource_scene", "resource_scene_spec.scene_id = resource_scene.id").
		OrderBy("resource_queue.compute_resource asc,resource_queue.acc_card_type asc").
		Unscoped().Distinct("resource_queue.compute_resource,resource_queue.acc_card_type").Find(&r)
	if err != nil {
		return nil, err
	}
	tmpMap := make(map[string][]string, 0)
	keys := make([]string, 0)
	for i := 0; i < len(r); i++ {
		spec := r[i]
		if _, exists := tmpMap[spec.ComputeResource]; exists {
			tmpMap[spec.ComputeResource] = append(tmpMap[spec.ComputeResource], spec.AccCardType)
		} else {
			keys = append(keys, spec.ComputeResource)
			tmpMap[spec.ComputeResource] = []string{spec.AccCardType}
		}
	}
	for i := 0; i < len(keys); i++ {
		res = append(res, AccCardInfo{
			ComputeSource: keys[i],
			CardList:      tmpMap[keys[i]],
		})
	}

	return res, nil
}

func GetAccCardListV2(specIds []int64) ([]AccCardInfo, error) {
	res := make([]AccCardInfo, 0)
	r := make([]*Specification, 0)

	err := x.Where("resource_specification.status = ? and (resource_queue.deleted_time = 0 or resource_queue.deleted_time is null)", SpecOnShelf).
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		In("resource_specification.id", specIds).
		OrderBy("resource_queue.compute_resource asc,resource_queue.acc_card_type asc").
		Unscoped().Distinct("resource_queue.compute_resource,resource_queue.acc_card_type").Find(&r)
	if err != nil {
		return nil, err
	}
	tmpMap := make(map[string][]string, 0)
	keys := make([]string, 0)
	for i := 0; i < len(r); i++ {
		spec := r[i]
		if _, exists := tmpMap[spec.ComputeResource]; exists {
			tmpMap[spec.ComputeResource] = append(tmpMap[spec.ComputeResource], spec.AccCardType)
		} else {
			keys = append(keys, spec.ComputeResource)
			tmpMap[spec.ComputeResource] = []string{spec.AccCardType}
		}
	}
	for i := 0; i < len(keys); i++ {
		res = append(res, AccCardInfo{
			ComputeSource: keys[i],
			CardList:      tmpMap[keys[i]],
		})
	}

	return res, nil
}

// FindSpecs
func FindAllSpecs() ([]*Specification, error) {
	return nil, nil
}

// resource_spec.go 或 vo/resource_spec_vo.go

type ResourceSpecLogVO struct {
	ID              int64              `json:"id"`
	BizType         string             `json:"biz_type"`
	OperateType     string             `json:"operate_type"`
	SpecId          string             `json:"spec_id"`
	CreatedBy       int64              `json:"created_by"`
	CreatedByName   string             `json:"created_by_name"`
	CreatedTime     timeutil.TimeStamp `json:"created_time"`
	Comment         string             `json:"comment"`
	ComputeResource string
	AccCardType     string
	AccCardNum      int
	AiCenterCode    string
	AiCenterName    string
	ResourceQueueID int64

	ResourceSpec  *ResourceSpecification
	ResourceQueue *ResourceQueue
}

func SearchResourceSpecLogs(opts SearchSpecLogOptions) (int64, []ResourceSpecLogVO, error) {
	var (
		logs  []AdminOperateLog
		total int64
	)
	var cond = builder.NewCond()

	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 {
		opts.PageSize = 20 // 默认每页20条
	}

	cond = cond.And(builder.Eq{"admin_operate_log.biz_type": "SpecOperate"})

	if opts.SpecId > 0 {
		cond = cond.And(builder.Eq{"admin_operate_log.related_id": opts.SpecId})
	}

	// 按操作类型查询 (上架/下架)
	if opts.OperateType != "" {
		cond = cond.And(builder.Eq{"admin_operate_log.operate_type": opts.OperateType})
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"admin_operate_log.compute_resource": opts.ComputeResource})
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"admin_operate_log.acc_card_type": opts.AccCardType})
	}
	if opts.AccCardNum >= 0 {
		cond = cond.And(builder.Eq{"admin_operate_log.acc_card_num": opts.AccCardNum})
	}

	if opts.AiCenterCode != "" {
		cond = cond.And(builder.Eq{"admin_operate_log.ai_center_code": opts.AiCenterCode})
	}
	if opts.ResourceQueueID > 0 {
		cond = cond.And(builder.Eq{"admin_operate_log.resource_queue_id": opts.ResourceQueueID})
	}
	//筛选掉历史任务，只选择新增的记录：有规格id
	cond = cond.And(builder.Gte{"admin_operate_log.resource_queue_id": 0})

	// 按时间范围查询
	if opts.StartTime != "" {
		cond = cond.And(builder.Gte{"admin_operate_log.created_time": opts.StartTime})
	}
	if opts.EndTime != "" {
		cond = cond.And(builder.Lte{"admin_operate_log.created_time": opts.EndTime})
	}

	n, err := x.Where(cond).Unscoped().Count(&AdminOperateLog{})
	if err != nil {
		return 0, nil, err
	}
	total = n
	session := x.Where(cond).
		OrderBy("id DESC").
		Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).
		Unscoped()
	err = session.Find(&logs)
	if err != nil {
		return 0, nil, err
	}

	// 转换数据：将 AdminOperateLog 转换为 ResourceSpecLogVO
	var result []ResourceSpecLogVO

	// 如果没有日志，直接返回
	if len(logs) == 0 {
		return total, result, nil
	}

	userIDs := make([]int64, 0)
	specIDs := make([]int64, 0)

	userIDSet := make(map[int64]bool)
	specIDSet := make(map[int64]bool)

	for _, log := range logs {
		// 收集用户ID (排除系统ID -1)
		if log.CreatedBy != -1 {
			if _, exists := userIDSet[log.CreatedBy]; !exists {
				userIDSet[log.CreatedBy] = true
				userIDs = append(userIDs, log.CreatedBy)
			}
		}

		// 收集规格ID (RelatedId 是 string 类型，需要转换)
		if log.RelatedId != "" {
			if id, err := strconv.ParseInt(log.RelatedId, 10, 64); err == nil {
				if _, exists := specIDSet[id]; !exists {
					specIDSet[id] = true
					specIDs = append(specIDs, id)
				}
			}
		}
	}

	userMap := make(map[int64]*User)
	if len(userIDs) > 0 {
		users, err := GetUsersByIDs(userIDs)
		if err != nil {
			log.Error("批量获取用户失败: %v", err)
		} else {
			for _, u := range users {
				userMap[u.ID] = u
			}
		}
	}

	for _, adminOperateLog := range logs {
		vo := ResourceSpecLogVO{
			ID:              adminOperateLog.ID,
			BizType:         adminOperateLog.BizType,
			OperateType:     adminOperateLog.OperateType,
			SpecId:          adminOperateLog.RelatedId, // 原样返回 string
			CreatedBy:       adminOperateLog.CreatedBy,
			CreatedTime:     adminOperateLog.CreatedTime,
			Comment:         adminOperateLog.Comment,
			ComputeResource: adminOperateLog.ComputeResource,
			AccCardType:     adminOperateLog.AccCardType,
			AccCardNum:      adminOperateLog.AccCardNum,
			AiCenterCode:    adminOperateLog.AiCenterCode,
			AiCenterName:    adminOperateLog.AiCenterName,
			ResourceQueueID: adminOperateLog.ResourceQueueID,
		}

		if adminOperateLog.CreatedBy == -1 {
			vo.CreatedByName = "系统自动同步"
		} else {
			if user, exists := userMap[adminOperateLog.CreatedBy]; exists {
				vo.CreatedByName = user.Name
			} else {
				vo.CreatedByName = fmt.Sprintf("用户(%d)", adminOperateLog.CreatedBy)
			}
		}

		if adminOperateLog.RelatedId != "" {
			if specID, err := strconv.ParseInt(adminOperateLog.RelatedId, 10, 64); err == nil {
				spec, err := GetResourceSpecification(&ResourceSpecification{ID: specID})
				if err != nil {
					log.Error("GetResourceSpecification failed: %v", err)
				}
				vo.ResourceSpec = spec
				vo.ResourceQueue, err = GetResourceQueueByID(spec.QueueId)
				if err != nil {
					log.Error("GetResourceQueue error.%v", err)
				}
			}

			result = append(result, vo)
		}
	}
	return total, result, err
}
