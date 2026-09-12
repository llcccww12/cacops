package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	QueueTypePublic    = "public"
	QueueTypeExclusive = "exclusive"
)

type ResourceQueue struct {
	ID                  int64 `xorm:"pk autoincr"`
	QueueCode           string
	QueueName           string
	QueueType           string
	Cluster             string `xorm:"notnull"`
	AiCenterCode        string
	AiCenterName        string
	ComputeResource     string
	AccCardType         string
	CardsTotalNum       int
	HasInternet         int //0 unknown;1 no internet;2 has internet
	IsAutomaticSync     bool
	IsAvailable         bool
	EnableVisualization bool //support visualization or not
	Remark              string
	DeletedTime         timeutil.TimeStamp `xorm:"deleted"`
	CreatedTime         timeutil.TimeStamp `xorm:"created"`
	CreatedBy           int64
	UpdatedTime         timeutil.TimeStamp `xorm:"updated"`
	UpdatedBy           int64
}

func (r ResourceQueue) ConvertToRes() *ResourceQueueRes {
	return &ResourceQueueRes{
		ID:                  r.ID,
		QueueCode:           r.QueueCode,
		QueueName:           r.QueueName,
		QueueType:           r.QueueType,
		Cluster:             r.Cluster,
		AiCenterCode:        r.AiCenterCode,
		AiCenterName:        r.AiCenterName,
		ComputeResource:     r.ComputeResource,
		AccCardType:         r.AccCardType,
		CardsTotalNum:       r.CardsTotalNum,
		UpdatedTime:         r.UpdatedTime,
		Remark:              r.Remark,
		HasInternet:         AICenterInternetStatus(r.HasInternet),
		IsAvailable:         r.IsAvailable,
		EnableVisualization: r.EnableVisualization,
	}
}

type ResourceQueueReq struct {
	QueueCode       string
	Cluster         string `binding:"Required"`
	AiCenterCode    string
	ComputeResource string `binding:"Required"`
	AccCardType     string `binding:"Required"`
	CardsTotalNum   int
	HasInternet     int
	CreatorId       int64
	IsAutomaticSync bool
	Remark          string
	QueueName       string
	QueueType       string
	IsAvailable     int
}

func (r ResourceQueueReq) ToDTO() ResourceQueue {
	isAvailable := false
	if r.IsAvailable == 2 {
		isAvailable = true
	}
	q := ResourceQueue{
		QueueCode:       r.QueueCode,
		Cluster:         r.Cluster,
		AiCenterCode:    r.AiCenterCode,
		ComputeResource: strings.ToUpper(r.ComputeResource),
		AccCardType:     strings.ToUpper(r.AccCardType),
		CardsTotalNum:   r.CardsTotalNum,
		HasInternet:     r.HasInternet,
		IsAutomaticSync: r.IsAutomaticSync,
		Remark:          r.Remark,
		CreatedBy:       r.CreatorId,
		UpdatedBy:       r.CreatorId,
		QueueName:       r.QueueName,
		QueueType:       r.QueueType,
		IsAvailable:     isAvailable,
	}
	if r.Cluster == OpenICluster {
		if r.AiCenterCode == AICenterOfCloudBrainOne {
			q.AiCenterName = "云脑一"
		} else if r.AiCenterCode == AICenterOfCloudBrainTwo {
			q.AiCenterName = "云脑二"
		} else if r.AiCenterCode == AICenterOfChengdu {
			q.AiCenterName = "启智成都智算"
		}
	}
	return q
}

type SearchResourceQueueOptions struct {
	ListOptions
	Cluster             string
	AiCenterCode        string
	ComputeResource     string
	AccCardType         string
	HasInternet         SpecInternetQuery
	QueueType           string
	IsAvailable         int
	IsQueueExclusive    int
	EnableVisualization int
}

type ResourceQueueListRes struct {
	TotalSize int64
	List      []*ResourceQueueRes
}

type ResourceQueueCodesRes struct {
	ID              int64
	QueueCode       string
	QueueName       string
	QueueType       string
	Cluster         string
	AiCenterCode    string
	AiCenterName    string
	ComputeResource string
	AccCardType     string
}

func (ResourceQueueCodesRes) TableName() string {
	return "resource_queue"
}

type ResourceAiCenterRes struct {
	AiCenterCode string
	AiCenterName string
}

func (r *ResourceAiCenterRes) Tr(language string) {
	r.AiCenterName = GetAiCenterShow(r.AiCenterCode, r.AiCenterName, language)
}

type GetQueueCodesOptions struct {
	Cluster []string
}

func NewResourceQueueListRes(totalSize int64, list []ResourceQueue) *ResourceQueueListRes {
	exclusiveMap := FindQueuesExclusiveMap()
	resList := make([]*ResourceQueueRes, len(list))
	for i, v := range list {
		resList[i] = v.ConvertToRes()
		if _, exists := exclusiveMap[v.ID]; exists {
			resList[i].IsQueueExclusive = true
		}
	}
	return &ResourceQueueListRes{
		TotalSize: totalSize,
		List:      resList,
	}
}

type ResourceQueueRes struct {
	ID                  int64
	QueueCode           string
	QueueType           string
	QueueName           string
	Cluster             string
	AiCenterCode        string
	AiCenterName        string
	ComputeResource     string
	AccCardType         string
	CardsTotalNum       int
	UpdatedTime         timeutil.TimeStamp
	Remark              string
	HasInternet         AICenterInternetStatus
	IsAvailable         bool
	IsQueueExclusive    bool
	EnableVisualization bool //support visualization or not
}

func InsertResourceQueue(queue ResourceQueue) (int64, error) {
	return x.Insert(&queue)
}

func UpdateResourceQueueById(queueId int64, queue ResourceQueue) (int64, error) {
	return x.ID(queueId).Update(&queue)
}

func UpdateResourceQueueCardNum(queue ResourceQueue) error {
	// 直接更新 cards_total_num 字段
	_, err := x.Cols("cards_total_num").Where("queue_code = ? and ai_center_code = ? and acc_card_type ILIKE ?", queue.QueueCode, queue.AiCenterCode, queue.AccCardType).Update(&ResourceQueue{
		CardsTotalNum: queue.CardsTotalNum,
	})
	if err != nil {
		return err
	} else {
		log.Info("QueueCode is %v, AiCenterCode is: %v, AccCardType is: %v, CardsTotalNum is: %v\n", queue.QueueCode, queue.AiCenterCode, queue.AccCardType, queue.CardsTotalNum)
	}

	return nil
}

func UpdateResourceCardsTotalNumAndInternetStatus(queueId int64, queue ResourceQueue, isAvailable int) (int64, error) {
	sess := x.NewSession()
	if err := sess.Begin(); err != nil {
		sess.Close()
		return 0, err
	}
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	cols := []string{"cards_total_num", "remark", "has_internet", "queue_type", "queue_name"}
	if isAvailable > 0 {
		if isAvailable == 1 {
			cols = append(cols, "is_available")
			queue.IsAvailable = false
		} else if isAvailable == 2 {
			cols = append(cols, "is_available")
			queue.IsAvailable = true
		}
	}

	n, err := sess.ID(queueId).Cols(cols...).Update(&queue)
	if err != nil {
		return 0, err
	}
	specIds := make([]int64, 0)
	if err = sess.Cols("resource_specification.id").Table("resource_specification").
		In("queue_id", queueId).Find(&specIds); err != nil {
		return 0, err
	}
	if len(specIds) == 0 {
		return n, nil
	}
	if isAvailable == 1 {
		if _, err = sess.Cols("status", "is_available").Table("resource_specification").In("id", specIds).Update(&ResourceSpecification{Status: SpecOffShelf, IsAvailable: false}); err != nil {
			return 0, err
		}
	} else if isAvailable == 2 {
		if _, err = sess.Cols("is_available").Table("resource_specification").In("id", specIds).Update(&ResourceSpecification{IsAvailable: true}); err != nil {
			return 0, err
		}
	}
	sess.Commit()
	return n, nil
}

func SearchResourceQueue(opts SearchResourceQueueOptions) (int64, []ResourceQueue, error) {
	var cond = builder.NewCond()
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.Cluster != "" {
		cond = cond.And(builder.Eq{"cluster": opts.Cluster})
	}
	if opts.AiCenterCode != "" {
		cond = cond.And(builder.Eq{"ai_center_code": opts.AiCenterCode})
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"compute_resource": opts.ComputeResource})
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"acc_card_type": opts.AccCardType})
	}
	if opts.HasInternet == QueryNoInternetSpecs {
		cond = cond.And(builder.Eq{"has_internet": NoInternet})
	} else if opts.HasInternet == QueryHasInternetSpecs {
		cond = cond.And(builder.Eq{"has_internet": HasInternet})
	}
	if opts.QueueType != "" {
		cond = cond.And(builder.Eq{"queue_type": opts.QueueType})
	}
	if opts.IsAvailable > 0 {
		if opts.IsAvailable == 1 {
			cond = cond.And(builder.Eq{"is_available": false})
		} else if opts.IsAvailable == 2 {
			cond = cond.And(builder.Eq{"is_available": true})
		}
	}
	if opts.EnableVisualization > 0 {
		if opts.EnableVisualization == 1 {
			cond = cond.And(builder.Neq{"enable_visualization": true})
		} else if opts.EnableVisualization == 2 {
			cond = cond.And(builder.Eq{"enable_visualization": true})
		}
	}
	if opts.IsQueueExclusive > 0 {
		queueIds := FindExclusiveQueueIds()
		if opts.IsQueueExclusive == 1 {
			cond = cond.And(builder.NotIn("id", queueIds))
		} else if opts.IsQueueExclusive == 2 {
			cond = cond.And(builder.In("id", queueIds))
		}
	}
	n, err := x.Where(cond).Unscoped().Count(&ResourceQueue{})
	if err != nil {
		return 0, nil, err
	}

	r := make([]ResourceQueue, 0)
	err = x.Where(cond).Desc("id").Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Unscoped().Find(&r)
	if err != nil {
		return 0, nil, err
	}
	return n, r, nil
}

func GetResourceQueueCodes(opts GetQueueCodesOptions) ([]*ResourceQueueCodesRes, error) {
	cond := builder.NewCond()
	if len(opts.Cluster) > 0 {
		cond = cond.And(builder.In("cluster", opts.Cluster))
	}
	cond = cond.And(builder.Or(builder.IsNull{"deleted_time"}, builder.Eq{"deleted_time": 0}))
	r := make([]*ResourceQueueCodesRes, 0)
	err := x.Where(cond).OrderBy("cluster desc,ai_center_code asc").Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetResourceQueue(r *ResourceQueue) (*ResourceQueue, error) {
	has, err := x.Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}
func GetResourceQueueByID(id int64) (*ResourceQueue, error) {
	r := &ResourceQueue{ID: id}
	has, err := x.Unscoped().Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}

func ParseComputeResourceFormGrampus(grampusDeviceKind string) string {
	t := strings.Split(grampusDeviceKind, "/")
	if len(t) < 2 {
		return ""
	}
	return strings.ToUpper(t[1])
}

type MemSize struct {
	Sizes []string
	Hex   int
}

var memSize = MemSize{Sizes: []string{"K", "M", "G", "T", "P", "E"}, Hex: 1000}
var iMemSize = MemSize{Sizes: []string{"Ki", "Mi", "Gi", "Ti", "Pi", "Ei"}, Hex: 1024}

func MatchMemSize(memSize MemSize, val string) (int, float32, error) {
	for i, v := range memSize.Sizes {
		if strings.HasSuffix(val, v) {
			s := strings.TrimSuffix(val, v)
			f, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return 0, 0, err
			}
			return i, float32(f), nil
		}
	}
	return -1, 0, nil
}

// TransferMemSize transfer oldValue format from old index to new index
// eg:  memSize.Sizes = []string{"M", "G", "T", "P", "E"}, oldValue = 10 , oldIndex = 1 , newIndex = 0. it means transfer 10G to 10000M
// so it returns 10000
func TransferMemSize(memSize MemSize, oldValue float32, oldIndex int, newIndex int) float32 {
	diff := oldIndex - newIndex
	r := oldValue
	if diff > 0 {
		r = oldValue * float32(diff) * float32(memSize.Hex)
	} else if diff < 0 {
		r = oldValue / float32(-1*diff) / float32(memSize.Hex)
	}
	return r
}

// ParseMemSize find the memSize which matches value's format,and parse the number from value
func ParseMemSize(value string, memSize MemSize, newIndex int) (bool, float32, error) {
	index, r, err := MatchMemSize(memSize, value)
	if err != nil {
		return false, 0, err
	}
	if index < 0 {
		return false, 0, nil
	}
	return true, TransferMemSize(memSize, r, index, newIndex), nil
}

func ParseMemSizeFromGrampus(grampusMemSize string) (float32, error) {
	if grampusMemSize == "" {
		return 0, nil
	}
	memflag, memResult, err := ParseMemSize(grampusMemSize, memSize, 2)
	if err != nil {
		return 0, err
	}
	if memflag {
		return memResult, nil
	}

	iMemFlag, imemResult, err := ParseMemSize(grampusMemSize, iMemSize, 2)
	if err != nil {
		return 0, err
	}
	if iMemFlag {
		return imemResult, nil
	}
	return 0, errors.New("grampus memSize format error")
}

func SyncGrampusQueues(updateList []ResourceQueue, insertList []ResourceQueue, existIds []int64) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	//delete queues that no longer exists
	deleteQueueIds := make([]int64, 0)
	queueCond := builder.NewCond()
	queueCond = queueCond.And(builder.NotIn("resource_queue.id", existIds)).And(builder.Eq{"resource_queue.cluster": C2NetCluster})
	if err := sess.Cols("resource_queue.id").Table("resource_queue").
		Where(queueCond).Find(&deleteQueueIds); err != nil {
		return err
	}

	if len(deleteQueueIds) > 0 {
		if _, err = sess.Cols("is_available").Table("resource_queue").In("id", deleteQueueIds).Update(&ResourceQueue{IsAvailable: false}); err != nil {
			return err
		}

		//delete specs and scene that no longer exists
		deleteSpcIds := make([]int64, 0)
		if err := sess.Cols("resource_specification.id").Table("resource_specification").
			In("queue_id", deleteQueueIds).Find(&deleteSpcIds); err != nil {
			return err
		}
		if len(deleteSpcIds) > 0 {
			if _, err = sess.Cols("status", "is_available").Table("resource_specification").In("id", deleteSpcIds).Update(&ResourceSpecification{Status: SpecOffShelf, IsAvailable: false}); err != nil {
				return err
			}
		}

	}

	//update exists specs
	if len(updateList) > 0 {
		for _, v := range updateList {
			if _, err = sess.ID(v.ID).Update(&v); err != nil {
				return err
			}
			if _, err = sess.ID(v.ID).Cols("is_available", "enable_visualization", "cards_total_num").Table("resource_queue").Update(&v); err != nil {
				return err
			}
		}

	}

	//insert new specs
	if len(insertList) > 0 {
		if _, err = sess.Insert(insertList); err != nil {
			return err
		}
	}

	return sess.Commit()
}

func GetResourceAiCenters() ([]ResourceAiCenterRes, error) {
	r := make([]ResourceAiCenterRes, 0)

	err := x.SQL("SELECT  t.ai_center_code, t.ai_center_name FROM (SELECT DISTINCT ai_center_code, ai_center_name,cluster FROM resource_queue WHERE (deleted_time IS NULL OR deleted_time=0)) t ORDER BY cluster desc,ai_center_code asc").Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetAvailableResourceAiCenters() ([]*ResourceAiCenterRes, error) {
	r := make([]*ResourceAiCenterRes, 0)
	sql := "SELECT  t.ai_center_code, t.ai_center_name FROM (SELECT DISTINCT resource_queue.ai_center_code, resource_queue.ai_center_name,resource_queue.cluster FROM resource_queue inner join resource_specification  on resource_queue.id = resource_specification.queue_id inner join resource_scene_spec on resource_specification.id = resource_scene_spec.spec_id WHERE (resource_queue.deleted_time IS NULL OR resource_queue.deleted_time=0) and resource_queue.is_available = true and resource_specification.status = 2 ) t ORDER BY cluster desc,ai_center_code asc"

	err := x.SQL(sql).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func GetAvailableResourceAiCentersV2(specIds []int64) ([]*ResourceAiCenterRes, error) {
	if len(specIds) == 0 {
		// 如果没有 specIds，返回空结果或所有结果？
		// 根据业务需求决定，这里返回空
		return []*ResourceAiCenterRes{}, nil
	}

	r := make([]*ResourceAiCenterRes, 0)

	// 构建 SQL 查询
	query := `
        SELECT DISTINCT
            rq.ai_center_code,
            rq.ai_center_name,
            rq.cluster
        FROM resource_queue rq
        INNER JOIN resource_specification rs ON rq.id = rs.queue_id
        WHERE  rs.id IN (%s)
        ORDER BY rq.cluster DESC, rq.ai_center_code ASC
    `

	// 构建 IN 条件的占位符
	placeholders := make([]string, len(specIds))
	args := make([]interface{}, len(specIds))
	for i, id := range specIds {
		placeholders[i] = "?"
		args[i] = id
	}

	sql := fmt.Sprintf(query, strings.Join(placeholders, ","))

	// 执行查询
	err := x.SQL(sql, args...).Find(&r)
	if err != nil {

		return nil, err
	}

	return r, nil
}
func GetExclusiveQueueIds(opts FindSpecsOptions) []*ResourceExclusivePool {
	pools, err := FindExclusivePools()
	if err != nil {
		log.Error("GetSpecialQueueIds FindSpecialQueueConfig err.%v", err)
		return nil
	}

	queues := make([]*ResourceExclusivePool, 0)
	for _, queue := range pools {
		if queue.JobType != string(opts.JobType) {
			continue
		}

		if queue.Cluster != opts.Cluster {
			continue
		}
		if queue.ComputeResource != opts.ComputeResource {
			continue
		}
		queues = append(queues, queue)
	}
	return queues
}

func GetAllExclusiveQueueIds() []*ResourceExclusivePool {
	pools, err := FindExclusivePools()
	if err != nil {
		log.Error("GetSpecialQueueIds FindSpecialQueueConfig err.%v", err)
		return nil
	}
	return pools
}

func IsUserInExclusivePool(userId int64) bool {
	userOrgs, err := GetOrgsByUserID(userId, true)
	if err != nil {
		log.Error("GetSpecialQueueIds GetOrgsByUserID error.%v", err)
		return false
	}
	pools, err := FindExclusivePools()
	if err != nil {
		log.Error("IsUserInSpecialPool FindExclusivePools err.%v", err)
		return false
	}
	if len(pools) == 0 {
		return false
	}
	for _, org := range userOrgs {
		for _, queue := range pools {
			if strings.ToLower(org.Name) == strings.ToLower(queue.OrgName) {
				return true
			}
		}
	}
	return false
}

var defaultLanguage = "zh-CN"

func GetAiCenterShow(aiCenterCode, aiCenterName, language string) string {
	if aiCenterCode == "" {
		return aiCenterName
	}
	if aiCenterName == "" {
		aiCenterName = aiCenterCode
	}
	if aiCenterCode == AICenterOfCloudBrainOne {
		if language == defaultLanguage {
			return "云脑一"
		} else {
			return AICenterOfCloudBrainOne
		}
	} else if aiCenterCode == AICenterOfCloudBrainTwo {
		if language == defaultLanguage {
			return "云脑二"
		} else {
			return AICenterOfCloudBrainTwo
		}
	} else if aiCenterCode == AICenterOfChengdu {
		if language == defaultLanguage {
			return "启智成都智算"
		} else {
			return AICenterOfChengdu
		}
	}
	if setting.AiCenterCodeAndNameAndLocMapInfo != nil {
		if info, ok := setting.AiCenterCodeAndNameAndLocMapInfo[aiCenterCode]; ok {
			if language == defaultLanguage {
				return info.Content
			} else {
				return info.ContentEN
			}
		} else {
			return aiCenterName
		}
	} else {
		return aiCenterName
	}
	return ""
}

func GetAccCardInfoByAccCardTypeList(accCardTypeList []string) ([]AccCardInfo, error) {
	queryRes := make([]ResourceQueue, 0)
	err := x.Distinct("acc_card_type", "compute_resource").
		Where(builder.NewCond().And(builder.In("acc_card_type", accCardTypeList))).Find(&queryRes)
	if err != nil {
		return nil, err
	}
	result := make([]AccCardInfo, 0)
	tmpMap := make(map[string][]string, 0)
	for i := 0; i < len(queryRes); i++ {
		tmpQueue := queryRes[i]
		if _, ok := tmpMap[tmpQueue.ComputeResource]; ok {
			tmpMap[tmpQueue.ComputeResource] = append(tmpMap[tmpQueue.ComputeResource], tmpQueue.AccCardType)
		} else {
			tmpMap[tmpQueue.ComputeResource] = []string{tmpQueue.AccCardType}
		}
	}
	for k, v := range tmpMap {
		accCard := AccCardInfo{
			ComputeSource: k,
			CardList:      v,
		}
		result = append(result, accCard)
	}
	return result, nil
}

func GetAccCardInfoByAccCardTypeAndQueueTypeList() ([]ResourceQueue, error) {
	queryRes := make([]ResourceQueue, 0)
	err := x.Distinct("compute_resource", "acc_card_type", "queue_type").Find(&queryRes)
	if err != nil {
		return nil, err
	}
	return queryRes, nil
}
