package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type SourceType string

const (
	SourceTypeAccomplishTask    SourceType = "ACCOMPLISH_TASK"
	SourceTypeAdminOperate      SourceType = "ADMIN_OPERATE"
	SourceTypeRunCloudbrainTask SourceType = "RUN_CLOUDBRAIN_TASK"
)

func (r SourceType) Name() string {
	switch r {
	case SourceTypeAccomplishTask:
		return "ACCOMPLISH_TASK"
	case SourceTypeAdminOperate:
		return "ADMIN_OPERATE"
	case SourceTypeRunCloudbrainTask:
		return "RUN_CLOUDBRAIN_TASK"
	default:
		return ""
	}
}
func (r SourceType) ChineseName() string {
	switch r {
	case SourceTypeAccomplishTask:
		return "积分任务"
	case SourceTypeAdminOperate:
		return "管理员操作"
	case SourceTypeRunCloudbrainTask:
		return "运行云脑任务"
	default:
		return ""
	}
}

type RewardType string

const (
	RewardTypePoint RewardType = "POINT"
)

func (r RewardType) Name() string {
	switch r {
	case RewardTypePoint:
		return "POINT"
	default:
		return ""
	}
}
func (r RewardType) Show() string {
	switch r {
	case RewardTypePoint:
		return "积分"
	default:
		return ""
	}
}
func GetRewardTypeInstance(s string) RewardType {
	switch s {
	case RewardTypePoint.Name():
		return RewardTypePoint
	default:
		return ""
	}
}

type RewardOperateType string

func (r RewardOperateType) Name() string {
	switch r {
	case OperateTypeIncrease:
		return "INCREASE"
	case OperateTypeDecrease:
		return "DECREASE"
	default:
		return ""
	}
}
func (r RewardOperateType) Show() string {
	switch r {
	case OperateTypeIncrease:
		return "奖励"
	case OperateTypeDecrease:
		return "扣减"
	default:
		return ""
	}
}

func GetRewardOperateTypeInstance(s string) RewardOperateType {
	switch s {
	case OperateTypeIncrease.Name():
		return OperateTypeIncrease
	case OperateTypeDecrease.Name():
		return OperateTypeDecrease
	default:
		return ""
	}
}

const (
	OperateTypeIncrease RewardOperateType = "INCREASE"
	OperateTypeDecrease RewardOperateType = "DECREASE"
	OperateTypeNull     RewardOperateType = "NIL"
)

const (
	OperateStatusOperating = "OPERATING"
	OperateStatusSucceeded = "SUCCEEDED"
	OperateStatusFailed    = "FAILED"
)

const Semicolon = ";"

type RewardOperateOrderBy string

const (
	RewardOrderByIDDesc RewardOperateOrderBy = "reward_operate_record.id desc"
)

type RewardRecordList []*RewardOperateRecord
type RewardRecordShowList []*RewardOperateRecordShow

func (l RewardRecordShowList) loadAttribute(isAdmin bool) {
	l.loadAction()
	l.loadCloudbrain()
	if isAdmin {
		l.loadAdminLog()
	}
}

func (l RewardRecordShowList) loadAction() error {
	if len(l) == 0 {
		return nil
	}
	actionIds := make([]int64, 0)
	for _, r := range l {
		if r.SourceType != SourceTypeAccomplishTask.Name() {
			continue
		}
		if !setting.UseDynamicData {
			if r.SourceContent != "" {
				actShow := &ActionShow{}
				err := json.Unmarshal([]byte(r.SourceContent), actShow)
				if err == nil {
					r.Action = actShow
				}
			}
			continue
		}
		i, _ := strconv.ParseInt(r.SourceId, 10, 64)
		actionIds = append(actionIds, i)
	}
	actions, err := GetActionByIds(actionIds)
	if err != nil {
		return err
	}
	actionIdMap := make(map[string]*Action, 0)
	for _, v := range actions {
		actionIdMap[fmt.Sprint(v.ID)] = v
	}

	for i, r := range l {
		act := actionIdMap[r.SourceId]
		if act != nil {
			l[i].Action = act.ToShow()
		}
	}
	return nil
}

func (l RewardRecordShowList) loadCloudbrain() error {
	if len(l) == 0 {
		return nil
	}
	cloudbrainIds := make([]int64, 0)
	cloudbrainMap := make(map[int64]*RewardOperateRecordShow, 0)
	for _, r := range l {
		if r.SourceType != SourceTypeRunCloudbrainTask.Name() {
			continue
		}
		if !setting.UseDynamicData && r.SourceContent != "" && r.Status == OperateStatusSucceeded {
			cloudbrainShow := &CloudbrainShow{}
			err := json.Unmarshal([]byte(r.SourceContent), cloudbrainShow)
			if err == nil {
				r.Cloudbrain = cloudbrainShow
				continue
			}
		}
		i, _ := strconv.ParseInt(r.SourceId, 10, 64)
		cloudbrainIds = append(cloudbrainIds, i)
		cloudbrainMap[i] = r
	}
	cloudbrains, err := GetCloudbrainByIds(cloudbrainIds)
	if err != nil {
		return err
	}
	var repoIds []int64
	var taskIds []int64
	for _, task := range cloudbrains {
		repoIds = append(repoIds, task.RepoID)
		taskIds = append(taskIds, task.ID)
	}
	repositoryMap, err := GetRepositoriesMapByIDs(repoIds)
	specMap, err := GetResourceSpecMapByCloudbrainIDs(taskIds)
	if err != nil {
		return err
	}
	for _, v := range cloudbrains {
		v.Repo = repositoryMap[v.RepoID]
		v.Spec = specMap[v.ID]
		cloudbrainMap[v.ID].Cloudbrain = v.ToShow()
	}

	return nil

}

func (l RewardRecordShowList) loadAdminLog() error {
	if len(l) == 0 {
		return nil
	}
	logIds := make([]string, 0)
	logMap := make(map[string]*RewardOperateRecordShow, 0)
	for _, r := range l {
		if r.SourceType != SourceTypeAdminOperate.Name() {
			continue
		}
		if !setting.UseDynamicData && r.SourceContent != "" {
			adminLogShow := &RewardAdminLogShow{}
			err := json.Unmarshal([]byte(r.SourceContent), adminLogShow)
			if err == nil {
				r.AdminLog = adminLogShow
				continue
			}
		}
		logIds = append(logIds, r.SourceId)
		logMap[r.SourceId] = r
	}
	adminLogs, err := GetRewardAdminLogByLogIds(logIds)
	if err != nil {
		return err
	}
	for _, v := range adminLogs {
		logMap[v.LogId].AdminLog = v.ToShow()
	}

	return nil

}

type RewardOperateRecord struct {
	ID               int64   `xorm:"pk autoincr"`
	SerialNo         string  `xorm:"INDEX NOT NULL"`
	UserId           int64   `xorm:"INDEX NOT NULL"`
	Amount           float64 `xorm:"NOT NULL"`
	LossAmount       float64
	Title            string
	RewardType       string `xorm:"NOT NULL"`
	SourceType       string `xorm:"NOT NULL"`
	SourceId         string `xorm:"INDEX NOT NULL"`
	SourceTemplateId string
	RequestId        string `xorm:"INDEX NOT NULL"`
	OperateType      string `xorm:"NOT NULL"`
	Status           string `xorm:"NOT NULL"`
	Remark           string
	CreatedUnix      timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix      timeutil.TimeStamp `xorm:"INDEX updated"`
	LastOperateUnix  timeutil.TimeStamp `xorm:"INDEX"`
	SourceContent    string             `xorm:"text"`
}

type AdminRewardOperateReq struct {
	TargetUserId int64             `binding:"Required"`
	OperateType  RewardOperateType `binding:"Required"`
	Amount       float64           `binding:"Required;Range(0,100000)"`
	Remark       string
	RewardType   RewardType
}

type RewardOperateRecordShow struct {
	SerialNo         string
	Status           string
	OperateType      string
	SourceId         string
	Amount           float64
	LossAmount       float64
	BalanceAfter     float64
	Remark           string
	SourceType       string
	SourceTemplateId string
	UserName         string
	UserId           int64
	LastOperateDate  timeutil.TimeStamp
	UnitPrice        float64
	SuccessCount     int
	IntervalSeconds  int64
	Action           *ActionShow
	Cloudbrain       *CloudbrainShow
	AdminLog         *RewardAdminLogShow
	SourceContent    string
}

type Column struct {
	Name  string
	Value string
}

func (r *RewardOperateRecordShow) ConvertToExcelColumn() []Column {
	source := SourceType(r.SourceType).ChineseName()
	action := TaskType(r.SourceTemplateId).ChineseName()
	date := time.Unix(int64(r.LastOperateDate), 0).Format("2006-01-02 15:04:05")
	var operator = "--"
	var detail = r.Remark
	if r.AdminLog != nil {
		operator = r.AdminLog.CreatorName
	}
	result := make([]Column, 0)
	if r.OperateType == string(OperateTypeIncrease) {

		result = []Column{
			{Name: "流水号", Value: r.SerialNo},
			{Name: "用户名", Value: r.UserName},
			{Name: "用户ID", Value: fmt.Sprint(r.UserId)},
			{Name: "时间", Value: date},
			{Name: "场景", Value: source},
			{Name: "积分行为", Value: action},
			{Name: "说明", Value: detail},
			{Name: "操作者", Value: operator},
			{Name: "数量", Value: fmt.Sprint(r.Amount)},
			{Name: "积分余额", Value: fmt.Sprint(r.BalanceAfter)},
		}
	} else if r.OperateType == string(OperateTypeDecrease) {
		var status string
		switch r.Status {
		case OperateStatusOperating:
			status = "消耗中"
		case OperateStatusSucceeded:
			status = "已完成"
		case OperateStatusFailed:
			status = "失败"
		}

		var duration = "--"
		var displayJobName = "--"
		var aiCenter = "--"
		var spec = "--"
		if r.Cloudbrain != nil {
			duration = r.Cloudbrain.Duration
			displayJobName = r.Cloudbrain.DisplayJobName
			if r.Cloudbrain.ResourceSpec != nil {
				spec = "【" + r.Cloudbrain.GetChineseJobType() + "】" + "【" + r.Cloudbrain.ResourceSpec.ToShowString() + "】"
				clusterName := ""
				switch r.Cloudbrain.ResourceSpec.Cluster {
				case "OpenI":
					clusterName = "启智集群"
				case "C2Net":
					clusterName = "智算集群"
				}

				aiCenter = r.Cloudbrain.ResourceSpec.QueueCode + "(" + clusterName + "-" + r.Cloudbrain.ResourceSpec.AiCenterName + ")"
			}
		} else {
			spec = r.Remark
		}
		//每小时资源单价（积分/时）
		hourPrice := r.UnitPrice
		if r.IntervalSeconds > 0 {
			hourPrice = float64(3600/r.IntervalSeconds) * r.UnitPrice
		}
		result = []Column{
			{Name: "流水号", Value: r.SerialNo},
			{Name: "时间", Value: date},
			{Name: "状态", Value: status},
			{Name: "运行时长", Value: duration},
			{Name: "用户名", Value: r.UserName},
			{Name: "用户ID", Value: fmt.Sprint(r.UserId)},
			{Name: "场景", Value: source},
			{Name: "任务名称", Value: displayJobName},
			{Name: "智算中心", Value: aiCenter},
			{Name: "资源规格", Value: spec},
			{Name: "操作者", Value: operator},
			{Name: "每小时资源单价（积分/时）", Value: fmt.Sprintf("%.2f", hourPrice)},
			{Name: "每扣费周期资源单价（积分/每周期）", Value: fmt.Sprint(r.UnitPrice)},
			{Name: "扣费周期数量", Value: fmt.Sprint(r.SuccessCount)},
			{Name: "总额", Value: fmt.Sprint(r.Amount)},
			{Name: "积分余额", Value: fmt.Sprint(r.BalanceAfter)},
		}
	}
	return result
}

func getPointOperateRecord(tl *RewardOperateRecord) (*RewardOperateRecord, error) {
	has, err := x.Get(tl)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return tl, nil
}

func GetPointOperateRecordBySourceTypeAndRequestId(sourceType, requestId, operateType string) (*RewardOperateRecord, error) {
	t := &RewardOperateRecord{
		SourceType:  sourceType,
		RequestId:   requestId,
		OperateType: operateType,
	}
	return getPointOperateRecord(t)
}

func GetPointOperateRecordBySerialNo(serialNo string) (*RewardOperateRecord, error) {
	t := &RewardOperateRecord{
		SerialNo: serialNo,
	}
	return getPointOperateRecord(t)
}

func InsertRewardOperateRecord(tl *RewardOperateRecord) (int64, error) {
	return x.Insert(tl)
}

func UpdateRewardRecordToFinalStatus(sourceType, requestId, newStatus string) (int64, error) {
	r := &RewardOperateRecord{
		Status:          newStatus,
		LastOperateUnix: timeutil.TimeStampNow(),
	}
	return x.Cols("status", "last_operate_unix").Where("source_type=? and request_id=? and status=?", sourceType, requestId, OperateStatusOperating).Update(r)
}

func SumRewardAmountInTaskPeriod(rewardType string, sourceType string, userId int64, period *PeriodResult) (float64, error) {
	var cond = builder.NewCond()
	if period != nil {
		cond = cond.And(builder.Gte{"created_unix": period.StartTime.Unix()})
		cond = cond.And(builder.Lt{"created_unix": period.EndTime.Unix()})
	}
	if sourceType != "" {
		cond = cond.And(builder.Eq{"source_type": sourceType})
	}
	cond = cond.And(builder.Eq{"reward_type": rewardType})
	cond = cond.And(builder.Eq{"user_id": userId})
	return x.Where(cond).Sum(&RewardOperateRecord{}, "amount")
}

type RewardOperateContext struct {
	SourceType        SourceType
	SourceId          string
	SourceTemplateId  string
	SourceContent     string
	Title             string
	Remark            string
	Reward            Reward
	TargetUserId      int64
	RequestId         string
	OperateType       RewardOperateType
	RejectPolicy      LimiterRejectPolicy
	PermittedNegative bool
	LossAmount        float64
}

type Reward struct {
	Amount float64
	Type   RewardType
}

type UserRewardOperationRedis struct {
	UserId      int64
	Amount      float64
	RewardType  RewardType
	OperateType RewardOperateType
}

type UserRewardOperation struct {
	UserId int64
	Msg    string
}

func AppendRemark(remark, appendStr string) string {
	return strings.TrimPrefix(remark+Semicolon+appendStr, Semicolon)
}

type RewardRecordListOpts struct {
	ListOptions
	UserId      int64
	UserName    string
	OperateType RewardOperateType
	RewardType  RewardType
	SourceType  string
	TaskType    string
	SerialNo    string
	OrderBy     RewardOperateOrderBy
	IsAdmin     bool
	Status      string
}

func (opts *RewardRecordListOpts) toCond() builder.Cond {
	if opts.Page <= 0 {
		opts.Page = 1
	}

	if len(opts.OrderBy) == 0 {
		opts.OrderBy = RewardOrderByIDDesc
	}

	cond := builder.NewCond()
	if opts.UserId > 0 {
		cond = cond.And(builder.Eq{"reward_operate_record.user_id": opts.UserId})
	}
	if opts.OperateType != OperateTypeNull {
		cond = cond.And(builder.Eq{"reward_operate_record.operate_type": opts.OperateType.Name()})
	}
	if opts.SourceType != "" {
		cond = cond.And(builder.Eq{"reward_operate_record.source_type": opts.SourceType})
	}
	if opts.TaskType != "" {
		cond = cond.And(builder.Eq{"reward_operate_record.source_template_id": opts.TaskType})
	}
	if opts.SerialNo != "" {
		cond = cond.And(builder.Like{"reward_operate_record.serial_no", opts.SerialNo})
	}
	if opts.Status != "" {
		cond = cond.And(builder.Like{"reward_operate_record.status", opts.Status})
	}

	cond = cond.And(builder.Eq{"reward_operate_record.reward_type": opts.RewardType.Name()})
	cond = cond.And(builder.Gt{"reward_operate_record.amount": 0})
	return cond
}

func GetRewardRecordShowList(opts *RewardRecordListOpts) (RewardRecordShowList, int64, error) {
	cond := opts.toCond()
	count, err := x.Where(cond).Count(&RewardOperateRecord{})
	if err != nil {
		return nil, 0, err
	}
	r := make([]*RewardOperateRecordShow, 0)
	err = x.Table("reward_operate_record").Cols("reward_operate_record.source_id", "reward_operate_record.serial_no",
		"reward_operate_record.status", "reward_operate_record.operate_type", "reward_operate_record.amount",
		"reward_operate_record.loss_amount", "reward_operate_record.remark", "reward_operate_record.source_type", "reward_operate_record.source_template_id",
		"reward_operate_record.last_operate_unix as last_operate_date", "reward_operate_record.source_content").
		Where(cond).Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).OrderBy(string(opts.OrderBy)).Find(&r)

	if err != nil {
		return nil, 0, err
	}
	RewardRecordShowList(r).loadAttribute(false)
	return r, count, nil
}

func GetAdminRewardRecordShowList(opts *RewardRecordListOpts) (RewardRecordShowList, int64, error) {
	cond := opts.toCond()
	count, err := x.Where(cond).Count(&RewardOperateRecord{})
	if err != nil {
		return nil, 0, err
	}
	r := make([]*RewardOperateRecordShow, 0)
	switch opts.OperateType {
	case OperateTypeIncrease:
		err = x.Table("reward_operate_record").Cols("reward_operate_record.source_id", "reward_operate_record.serial_no",
			"reward_operate_record.status", "reward_operate_record.operate_type", "reward_operate_record.amount",
			"reward_operate_record.loss_amount", "reward_operate_record.remark", "reward_operate_record.source_type", "reward_operate_record.source_template_id",
			"reward_operate_record.last_operate_unix as last_operate_date", "reward_operate_record.source_content", "public.user.name as user_name", "public.user.id as user_id",
			"point_account_log.balance_after").
			Join("LEFT", "public.user", "reward_operate_record.user_id = public.user.id").
			Join("LEFT", "point_account_log", " reward_operate_record.serial_no = point_account_log.source_id").
			Where(cond).Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).OrderBy(string(opts.OrderBy)).Find(&r)
	case OperateTypeDecrease:
		err = x.Table("reward_operate_record").Cols("reward_operate_record.source_id", "reward_operate_record.serial_no",
			"reward_operate_record.status", "reward_operate_record.operate_type", "reward_operate_record.amount",
			"reward_operate_record.loss_amount", "reward_operate_record.remark", "reward_operate_record.source_type", "reward_operate_record.source_template_id",
			"reward_operate_record.last_operate_unix as last_operate_date", "reward_operate_record.source_content", "public.user.name as user_name", "public.user.id as user_id",
			"reward_periodic_task.amount as unit_price", "reward_periodic_task.success_count", "reward_periodic_task.interval_seconds").
			Join("LEFT", "public.user", "reward_operate_record.user_id = public.user.id").
			Join("LEFT", "reward_periodic_task", "reward_operate_record.serial_no = reward_periodic_task.operate_serial_no").
			Where(cond).Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).OrderBy(string(opts.OrderBy)).Find(&r)
	}

	if err != nil {
		return nil, 0, err
	}

	//查询每条消耗明细记录的实时积分余额
	if opts.OperateType == OperateTypeDecrease && len(r) > 0 {
		serialNos := make([]string, len(r))
		for i := 0; i < len(r); i++ {
			serialNos[i] = r[i].SerialNo
		}
		logList1 := make([]PointAccountLog, 0)
		x.Table("point_account_log").Select("source_id,max(account_version) as account_version").
			In("source_id", serialNos).GroupBy("source_id").Find(&logList1)
		newCond := builder.NewCond()
		for _, l := range logList1 {
			newCond = newCond.Or(builder.And(builder.Eq{"source_id": l.SourceId}, builder.Eq{"account_version": l.AccountVersion}))
		}
		logList2 := make([]PointAccountLog, 0)
		x.Table("point_account_log").Cols("source_id", "balance_after").
			Where(newCond).Find(&logList2)
		tmpMap := make(map[string]float64, 0)
		for _, l := range logList2 {
			tmpMap[l.SourceId] = l.BalanceAfter
		}
		for i, l := range r {
			r[i].BalanceAfter = tmpMap[l.SerialNo]
		}
	}
	RewardRecordShowList(r).loadAttribute(true)
	return r, count, nil
}

func IsWechatOpenIdRewarded(wechatOpenId string) bool {
	actions := make([]Action, 0)
	err := x.Where(" op_type = ? and content = ?", ActionBindWechat, wechatOpenId).Find(&actions)

	if err != nil {
		log.Error("IsWechatOpenIdRewarded find actions err.%v", err)
		return true
	}
	if len(actions) == 0 {
		return false
	}
	actionIds := make([]int64, len(actions))
	for i, v := range actions {
		actionIds[i] = v.ID
	}
	n, _ := x.Where(builder.Eq{"source_type": SourceTypeAccomplishTask}.And(builder.In("source_id", actionIds))).Count(&RewardOperateRecord{})
	return n > 0
}

func UpdateSourceContent(sourceId string, sourceType SourceType, sourceContent string) error {
	_, err := x.Exec("Update reward_operate_record set source_content = ? where source_type = ? and source_id = ?", sourceContent, sourceType, sourceId)
	return err
}

func GetNoSourceContentRewardRecord(lastId int64) ([]*RewardOperateRecord, error) {
	r := make([]*RewardOperateRecord, 0)
	err := x.Table("reward_operate_record").
		Where("(source_content = '' OR source_content IS NULL) and id > ?", lastId).Limit(100).OrderBy("id").Find(&r)

	if err != nil {
		return nil, err
	}
	return r, nil
}
