package models

import (
	"sort"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type TaskDetail struct {
	ID              int64              `json:"ID"`
	JobID           string             `json:"JobID"`
	JobName         string             `json:"JobName"`
	DisplayJobName  string             `json:"DisplayJobName"`
	Status          string             `json:"Status"`
	DetailedStatus  string             `json:"DetailedStatus"`
	JobType         string             `json:"JobType"`
	CreatedUnix     timeutil.TimeStamp `json:"CreatedUnix"`
	WaitTime        string             `json:"WaitTime"`
	RunTime         string             `json:"RunTime"`
	StartTime       timeutil.TimeStamp `json:"StartTime"`
	EndTime         timeutil.TimeStamp `json:"EndTime"`
	ComputeResource string             `json:"ComputeResource"`
	Type            int                `json:"Type"`
	UserName        string             `json:"UserName"`
	RepoName        string             `json:"RepoName"`
	RepoAlias       string             `json:"RepoAlias"`
	RepoID          int64              `json:"RepoID"`
	IsDelete        bool               `json:"IsDelete"`
	CardNum         int                `json:"CardNum"`
	CardType        string             `json:"CardType"`
	CardDuration    string             `json:"CardDuration"`
	AiCenter        string             `json:"AiCenter"`
	FlavorName      string             `json:"FlavorName"`
	WorkServerNum   int64              `json:"WorkServerNum"`
	Spec            *Specification     `json:"Spec"`
}

type CloudbrainDurationStatistic struct {
	ID              int64 `xorm:"pk autoincr"`
	Cluster         string
	AiCenterCode    string `xorm:"INDEX"`
	AiCenterName    string
	ComputeResource string
	AccCardType     string `xorm:"INDEX"`

	DateTime           timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DateTimeUnix       timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DayTime            string             `xorm:"INDEX"`
	HourTime           int                `xorm:"INDEX"`
	CardsUseDuration   int
	CardsTotalDuration int
	CardsTotalNum      int

	DeletedUnix timeutil.TimeStamp `xorm:"deleted"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type CloudbrainTaskNumStatistic struct {
	ID           int64 `xorm:"pk autoincr"`
	Cluster      string
	AiCenterCode string `xorm:"INDEX"`
	AiCenterName string
	AccCardType  string `xorm:"INDEX"`

	DateTime     timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DateTimeUnix timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DayTime      string             `xorm:"INDEX"`
	HourTime     int                `xorm:"INDEX"`
	TaskNum      int                //分中心任务数

	DeletedUnix timeutil.TimeStamp `xorm:"deleted"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type DateCloudbrainNumJson struct {
	Date              string
	CloudbrainNumJson string `xorm:"TEXT"`
}
type QueueTaskDurationStatistic struct {
	ID      int64 `xorm:"pk autoincr"`
	QueueID int64 `xorm:"INDEX"`

	DateTime           timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DateTimeUnix       timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DayTime            string             `xorm:"INDEX"`
	HourTime           int                `xorm:"INDEX"`
	CardsUseDuration   int
	CardsTotalDuration int
	CardsTotalNum      int

	DeletedUnix timeutil.TimeStamp `xorm:"deleted"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type QueueTaskNumStatistic struct {
	ID      int64 `xorm:"pk autoincr"`
	QueueID int64 `xorm:"INDEX"`

	DateTime     timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DateTimeUnix timeutil.TimeStamp `xorm:"INDEX DEFAULT 0"`
	DayTime      string             `xorm:"INDEX"`
	HourTime     int                `xorm:"INDEX"`
	TaskNum      int                //分中心任务数

	DeletedUnix timeutil.TimeStamp `xorm:"deleted"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}
type DurationStatisticOptions struct {
	BeginTime    time.Time
	EndTime      time.Time
	AiCenterCode string
	QueueID      int64
}

type DurationRateStatistic struct {
	AiCenterTotalDurationStat map[string]int `json:"aiCenterTotalDurationStat"`
	AiCenterUsageDurationStat map[string]int `json:"aiCenterUsageDurationStat"`
	C2NetUsageDurationTable   map[string]int `json:"c2NetUsageDurationTable"`

	UsageRate map[string]float64 `json:"UsageRate"`
}
type ResourceDetail struct {
	QueueCode       string
	Cluster         string `xorm:"notnull"`
	AiCenterCode    string
	AiCenterName    string
	ComputeResource string
	AccCardType     string
	CardsTotalNum   int
	IsAutomaticSync bool
}

type DateUsageStatistic struct {
	Date          string  `json:"date"`
	UsageDuration int     `json:"usageDuration"`
	TotalDuration int     `json:"totalDuration"`
	UsageRate     float64 `json:"usageRate"`
}

type DateCenterTaskStatistic struct {
	Date          string         `json:"date"`
	UsageDuration map[string]int `json:"usageDuration"`
	TaskNum       map[string]int `json:"taskNum"`
}

type DateQueueTaskStatistic struct {
	Date               string                   `json:"date"`
	QueueTotalDuration map[int64]int            `json:"queueTotalDuration"`
	QueueUsageDuration map[int64]int            `json:"usageDuration"`
	QueueTaskNum       map[int64]int            `json:"taskNum"`
	QueueInfo          map[int64]*ResourceQueue `json:"queueInfo"`
}

type HourTimeStatistic struct {
	HourTimeUsageDuration map[string]int     `json:"hourTimeUsageDuration"`
	HourTimeTotalDuration map[string]int     `json:"hourTimeTotalDuration"`
	HourTimeUsageRate     map[string]float64 `json:"hourTimeUsageRate"`
}
type CloudbrainTypeDuration []struct {
	Type            int `xorm:"type"`
	DurationSum     int `xorm:"durationSum"`
	CardDurationSum int `xorm:"cardDurationSum"`
	Count           int `xorm:"count"`
}
type CloudbrainAllDuration struct {
	DurationSum     int `xorm:"durationSum"`
	CardDurationSum int `xorm:"cardDurationSum"`
	Count           int `xorm:"count"`
}

type XPUInfoBase struct {
	ID           int64 `xorm:"pk autoincr"`
	CardType     string
	CardTypeShow string
	ResourceType string
	Company      string
	AccessTime   string
}

type XPUInfoStatistic struct {
	ID           int64 `xorm:"pk autoincr"`
	InfoID       int64 `xorm:"index"`
	Type         int
	UsedDuration int64
	UsedCardHour int64
	UserCount    int64
	TaskCount    int64
	UpdatedUnix  int64
}

type XPUInfoStatisticExtendBase struct {
	CardType     string
	CardTypeShow string
	ResourceType string
	Company      string
	AccessTime   string
	UpdatedUnix  int64
	UsedCardHour int64
	UserCount    int64
	TaskCount    int64
}

func (XPUInfoStatisticExtendBase) TableName() string {
	return "xpu_info_statistic"
}

type XPUInfoStatisticShow struct {
	CardType     string `json:"card_type"`
	ResourceType string `json:"resource_type"`
	Company      string `json:"company"`
	AccessTime   string `json:"access_time"`
	Count        int64  `json:"count"`
	UpdatedUnix  int64  `json:"updated_unix"`
}

const TypeAllDays = 0
const TypeSevenDays = 1
const TypeThirtyDays = 2

var XpuInfoCategories = map[string]string{"card": "used_card_hour", "user": "user_count", "task": "task_count"}
var XpuInfoType = map[string]int{"all": TypeAllDays, "7": TypeSevenDays, "30": TypeThirtyDays}

func GetXPUInfos() ([]*XPUInfoBase, error) {
	infos := make([]*XPUInfoBase, 0)
	err := xStatistic.Find(&infos)
	return infos, err
}

func GetXPUStatisticInfos(dataType int, category string) ([]*XPUInfoStatisticShow, error) {
	statistics := make([]XPUInfoStatisticExtendBase, 0)

	s := xStatistic.Where("type=?", dataType).
		Join("INNER", "xpu_info_base", "xpu_info_base.id = xpu_info_statistic.info_id")
	err := s.OrderBy("xpu_info_statistic." + XpuInfoCategories[category] + " desc").Find(&statistics)

	if err != nil {
		return nil, err
	}
	var result = make([]*XPUInfoStatisticShow, 0)
	for _, statistic := range statistics {
		item := &XPUInfoStatisticShow{
			CardType:     statistic.CardTypeShow,
			ResourceType: statistic.ResourceType,
			Company:      statistic.Company,
			AccessTime:   statistic.AccessTime,
			UpdatedUnix:  statistic.UpdatedUnix,
		}

		if category == "card" {
			item.Count = statistic.UsedCardHour
		} else if category == "user" {
			item.Count = statistic.UserCount
		} else {
			item.Count = statistic.TaskCount
		}
		result = append(result, item)
	}
	return result, nil
}

func UpdateOrInsertXPUInfoStatistic(bean *XPUInfoStatistic) error {

	xpuInfoIndb := new(XPUInfoStatistic)
	has, err := xStatistic.Where("info_id = ? and type=?", bean.InfoID, bean.Type).Get(xpuInfoIndb)

	if err != nil {
		return err
	}

	if has {
		_, err = xStatistic.ID(xpuInfoIndb.ID).Cols("used_duration", "user_count", "used_card_hour", "task_count", "updated_unix").Update(bean)

	} else {
		_, err = xStatistic.Insert(bean)
	}
	return err

}

func HasTotalTypeXPUInfoStatisticRecord() bool {
	xpuInfo := new(XPUInfoStatistic)
	total, _ := xStatistic.Where("type =?", TypeAllDays).Count(xpuInfo)
	return total > 0
}

func FindTotalTypeXPUInfoStatisticRecords() ([]XPUInfoStatistic, error) {

	xpuInfos := make([]XPUInfoStatistic, 0)

	err := xStatistic.Where("type =?", TypeAllDays).Find(&xpuInfos)
	return xpuInfos, err
}
func GetTotalTypeLatestUpdateUnix() int64 {
	var updated []int64
	xStatistic.Table("xpu_info_statistic").Cols("updated_unix").Find(&updated)
	if len(updated) > 0 {
		return updated[0]
	}
	return 0
}

func GetTodayCreatorCount(beginTime time.Time, endTime time.Time) (int64, error) {
	countSql := "SELECT count(distinct user_id) FROM  " +
		"public.cloudbrain where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<" + strconv.FormatInt(endTime.Unix(), 10)
	return x.SQL(countSql).Count()
}
func GetTodayCloudbrainCount(beginTime time.Time, endTime time.Time) (int64, error) {
	countSql := "SELECT count(*) FROM  " +
		"public.cloudbrain where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<=" + strconv.FormatInt(endTime.Unix(), 10)
	return x.SQL(countSql).Count()
}
func GetTodayRunningCount(beginTime time.Time, endTime time.Time) (int64, error) {
	countSql := "SELECT count(*) FROM  " +
		"public.cloudbrain where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<=" + strconv.FormatInt(endTime.Unix(), 10) + " and (status='" + string(JobRunning) + "'" +
		" or status='" + string(ModelArtsTrainJobInit) + "')"
	return x.SQL(countSql).Count()
}
func GetTodayWaitingCount(beginTime time.Time, endTime time.Time) (int64, error) {
	countSql := "SELECT count(*) FROM  " +
		"public.cloudbrain where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<=" + strconv.FormatInt(endTime.Unix(), 10) + " and status='" + string(JobWaiting) + "'"
	return x.SQL(countSql).Count()
}

func GetCreatorCount() (int64, error) {
	countSql := "SELECT count(distinct user_id) FROM  public.cloudbrain"
	return x.SQL(countSql).Count()
}

func GetCloudbrainTypeCount() ([]map[string]string, error) {
	countSql := "SELECT type,count(*) num FROM public.cloudbrain group by type order by num desc"
	return x.QueryString(countSql)
}

func GetCloudbrainStatusCount() ([]map[string]string, error) {
	countSql := "SELECT status,count(*) num FROM public.cloudbrain group by status order by num desc"
	return x.QueryString(countSql)
}

func GetCloudbrainCardTimeAndCountGroupByAICenter() ([]map[string]string, error) {
	countSql := `select ai_center,SUM(
            COALESCE(a.duration *
				CASE
				WHEN a.work_server_number = 0 THEN 1
				ELSE COALESCE(a.work_server_number, 1)
				END *
				COALESCE(cloudbrain_spec.acc_cards_num, 1), 0)
   ) as card_duration,count(*) num from

(select  id,duration,work_server_number,case when type=0 then 'OpenIOne' when type=1 then 'OpenITwo' when type=3 then 'OpenIChengdu' else split_part(ai_center, '+',1)
        end ai_center
FROM public.cloudbrain ) a Left JOIN cloudbrain_spec on a.id = cloudbrain_spec.cloudbrain_id
where ai_center!='' group by  a.ai_center order by card_duration desc`
	return x.QueryString(countSql)
}

func GetCloudbrainTpyeDurationSum() ([]map[string]string, error) {
	countSql := "SELECT type,sum(duration) FROM public.cloudbrain group by type order by sum(duration) desc"
	return x.QueryString(countSql)
}

func GetRecordBeginTime() ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	sess.OrderBy("cloudbrain.id ASC limit 1")
	cloudbrains := make([]*CloudbrainInfo, 0)
	if err := sess.Table(&Cloudbrain{}).Unscoped().
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}
	return cloudbrains, nil
}

func GetAllStatusCloudBrain() map[string]int {
	sess := x.NewSession()
	defer sess.Close()
	cloudbrains := make([]*CloudbrainInfo, 0)
	if err := sess.Table(&Cloudbrain{}).Unscoped().
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}
	cloudBrainStatusResult := make(map[string]int)
	for _, cloudbrain := range cloudbrains {
		if _, ok := cloudBrainStatusResult[cloudbrain.Status]; !ok {
			cloudBrainStatusResult[cloudbrain.Status] = 1
		} else {
			cloudBrainStatusResult[cloudbrain.Status] += 1
		}
	}
	return cloudBrainStatusResult
}

func GetWaittingTop() ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"cloudbrain.status": string(JobWaiting)},
	)
	sess.OrderBy("cloudbrain.created_unix ASC limit 10")
	cloudbrains := make([]*CloudbrainInfo, 0, 10)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}

	var ids []int64
	for _, task := range cloudbrains {
		ids = append(ids, task.RepoID)
	}
	repositoryMap, err := GetRepositoriesMapByIDs(ids)
	if err == nil {
		for _, task := range cloudbrains {
			task.Repo = repositoryMap[task.RepoID]
		}
	}
	return cloudbrains, nil
}

func GetRunningTop() ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"cloudbrain.status": string(JobRunning)},
	)
	sess.OrderBy("cloudbrain.duration DESC limit 10")
	cloudbrains := make([]*CloudbrainInfo, 0, 10)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}
	var ids []int64
	for _, task := range cloudbrains {
		ids = append(ids, task.RepoID)
	}
	repositoryMap, err := GetRepositoriesMapByIDs(ids)
	if err == nil {
		for _, task := range cloudbrains {
			task.Repo = repositoryMap[task.RepoID]
		}
	}
	return cloudbrains, nil
}

func GetQueueTask(opts *QueueCloudbrainOptions) ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	if (opts.ComputeResource) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.compute_resource": opts.ComputeResource},
		)
	}

	if (opts.JobTypes) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.job_type": opts.JobTypes},
		)
	}

	if (opts.JobStatus) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.status": opts.JobStatus},
		)
	}

	cond = cond.And(
		builder.Or(
			builder.Eq{"cloudbrain.status": string(JobRunning)},
			builder.Eq{"cloudbrain.status": string(JobWaiting)},
			builder.Eq{"cloudbrain.status": string(GrampusStatusStopping)},
		),
	)

	cloudbrains := make([]*CloudbrainInfo, 0)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}
	var ids []int64
	for _, task := range cloudbrains {
		ids = append(ids, task.RepoID)
	}
	repositoryMap, err := GetRepositoriesMapByIDs(ids)
	if err == nil {
		for _, task := range cloudbrains {
			task.Repo = repositoryMap[task.RepoID]
		}
	}
	return cloudbrains, nil
}

func GetCreateHourPeriodCount(cloudbrains []*CloudbrainInfo, dateBeginTime time.Time, dateEndTime time.Time) (map[string]int64, error) {
	var createHourPeriodCount = make(map[string]int64)
	for _, cloudbrain := range cloudbrains {
		if cloudbrain.StartTime != 0 && cloudbrain.EndTime == 0 {
			cloudbrain.EndTime = timeutil.TimeStamp(time.Now().Unix())
		}
		hourBeginTime := dateBeginTime.Unix()
		hourEndTime := hourBeginTime + int64(3600)
		for hour := 0; hour < 24; hour++ {
			if int64(cloudbrain.Cloudbrain.CreatedUnix) >= hourBeginTime && int64(cloudbrain.Cloudbrain.CreatedUnix) < hourEndTime {
				createHourPeriodCount[strconv.Itoa(hour)] = createHourPeriodCount[strconv.Itoa(hour)] + 1
			}
			hourBeginTime = hourEndTime
			hourEndTime = hourEndTime + int64(3600)
		}
	}
	return createHourPeriodCount, nil
}

func GetRunHourPeriodCount(cloudbrains []*CloudbrainInfo, dateBeginTime time.Time, dateEndTime time.Time) (map[string]int64, error) {
	var runHourPeriodCount = make(map[string]int64)
	for _, cloudbrain := range cloudbrains {
		if cloudbrain.StartTime != 0 && cloudbrain.EndTime == 0 {
			cloudbrain.EndTime = timeutil.TimeStamp(time.Now().Unix())
		}
		hourBeginTime := dateBeginTime.Unix()
		hourEndTime := hourBeginTime + int64(3600)
		for hour := 0; hour < 24; hour++ {
			if cloudbrain.StartTime.AsTime().Unix() < hourEndTime && cloudbrain.EndTime.AsTime().Unix() > hourBeginTime {
				runHourPeriodCount[strconv.Itoa(hour)] = runHourPeriodCount[strconv.Itoa(hour)] + 1
			}
			hourBeginTime = hourEndTime
			hourEndTime = hourEndTime + int64(3600)
		}
	}
	return runHourPeriodCount, nil
}

func GetCloudbrainRunning() ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"cloudbrain.status": string(JobRunning)},
	)
	sess.OrderBy("cloudbrain.created_unix ASC")
	cloudbrains := make([]*CloudbrainInfo, 0, 10)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}
	return cloudbrains, nil
}

func GetCloudbrainByTime(beginTime int64, endTime int64) ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.Or(
		builder.And(builder.Gte{"cloudbrain.end_time": beginTime}, builder.Lte{"cloudbrain.start_time": beginTime}, builder.Gt{"cloudbrain.start_time": 0}),
	)
	cond = cond.Or(
		builder.And(builder.Gte{"cloudbrain.start_time": beginTime}, builder.Lte{"cloudbrain.start_time": endTime}, builder.Gt{"cloudbrain.start_time": 0}),
	)
	cond = cond.Or(
		builder.And(builder.Eq{"cloudbrain.status": string(JobRunning)}, builder.Lte{"cloudbrain.start_time": beginTime}),
	)
	sess.OrderBy("cloudbrain.id ASC")
	cloudbrains := make([]*CloudbrainInfo, 0, 10)
	if err := sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
		Find(&cloudbrains); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return cloudbrains, nil
}

func GetCloudbrainByCreateTime(beginTime int64, endTime int64) ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.Or(
		builder.And(builder.Gte{"cloudbrain.created_unix": beginTime}, builder.Lte{"cloudbrain.created_unix": endTime}),
	)
	sess.OrderBy("cloudbrain.id ASC")
	cloudbrains := make([]*CloudbrainInfo, 0, 10)
	if err := sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
		Find(&cloudbrains); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return cloudbrains, nil
}

func GetSpecByAiCenterCodeAndType(aiCenterCode string, accCardType string) ([]*CloudbrainSpec, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.And(builder.Eq{"cloudbrain_spec.ai_center_code": aiCenterCode}, builder.Eq{"cloudbrain_spec.acc_card_type": accCardType}),
	)
	cloudbrainSpecs := make([]*CloudbrainSpec, 0, 10)
	if err := sess.Table(&CloudbrainSpec{}).Where(cond).
		Find(&cloudbrainSpecs); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return cloudbrainSpecs, nil
}

func InsertCloudbrainDurationStatistic(cloudbrainDurationStatistic *CloudbrainDurationStatistic) (int64, error) {
	return xStatistic.Insert(cloudbrainDurationStatistic)
}

func InsertCloudbrainTaskNumStatistic(cloudbrainTaskNumStatistic *CloudbrainTaskNumStatistic) (int64, error) {
	return xStatistic.Insert(cloudbrainTaskNumStatistic)
}

func InsertCloudbrainTaskNumJson(dateCloudbrainTaskNumJson *DateCloudbrainNumJson) (int64, error) {
	return xStatistic.Insert(dateCloudbrainTaskNumJson)
}

func InsertQueueTaskDurationStatistic(QueueTaskDurationStatistic *QueueTaskDurationStatistic) (int64, error) {
	return xStatistic.Insert(QueueTaskDurationStatistic)
}
func InsertQueueTaskNumStatistic(QueueTaskNumStatistic *QueueTaskNumStatistic) (int64, error) {
	return xStatistic.Insert(QueueTaskNumStatistic)
}

func getDurationStatistic(cb *CloudbrainDurationStatistic) (*CloudbrainDurationStatistic, error) {
	has, err := x.Get(cb)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrJobNotExist{}
	}
	return cb, nil
}

func GetCanUseCardInfo() ([]*ResourceQueue, error) {
	sess := x.NewSession()
	defer sess.Close()
	sess.OrderBy("resource_queue.cluster DESC, resource_queue.ai_center_code ASC")
	ResourceQueues := make([]*ResourceQueue, 0, 10)
	if err := sess.Table(&ResourceQueue{}).Unscoped().Find(&ResourceQueues); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return ResourceQueues, nil
}

func CloudbrainDurationStatistics(opts *DurationStatisticOptions) ([]*CloudbrainDurationStatistic, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	if opts.BeginTime.Unix() > 0 && opts.EndTime.Unix() > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain_duration_statistic.date_time_unix": opts.BeginTime.Unix()}, builder.Lt{"cloudbrain_duration_statistic.date_time_unix": opts.EndTime.Unix()}),
		)
	}
	if opts.AiCenterCode != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain_duration_statistic.ai_center_code": opts.AiCenterCode},
		)
	}
	CloudbrainDurationStatistics := make([]*CloudbrainDurationStatistic, 0, 10)
	if err := sess.Table(&CloudbrainDurationStatistic{}).Where(cond).
		Find(&CloudbrainDurationStatistics); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return CloudbrainDurationStatistics, nil
}

func CloudbrainTaskNumStatistics(opts *DurationStatisticOptions) ([]*CloudbrainTaskNumStatistic, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	if opts.BeginTime.Unix() > 0 && opts.EndTime.Unix() > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain_task_num_statistic.date_time_unix": opts.BeginTime.Unix()}, builder.Lt{"cloudbrain_task_num_statistic.date_time_unix": opts.EndTime.Unix()}),
		)
	}
	if opts.AiCenterCode != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain_task_num_statistic.ai_center_code": opts.AiCenterCode},
		)
	}
	CloudbrainTaskNumStatistics := make([]*CloudbrainTaskNumStatistic, 0, 10)
	if err := sess.Table(&CloudbrainTaskNumStatistic{}).Where(cond).
		Find(&CloudbrainTaskNumStatistics); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return CloudbrainTaskNumStatistics, nil
}

func CheckDateExistsInCloudbrainNumJson(date string) int64 {
	dateCloudbrainNumJson := new(DateCloudbrainNumJson)
	total, _ := xStatistic.Where("date =?", date).Count(dateCloudbrainNumJson)
	return total
}

func GetAllCloudbrainNumJson() ([]*DateCloudbrainNumJson, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	DateCloudbrainNumJsons := make([]*DateCloudbrainNumJson, 0, 10)
	if err := sess.Table(&DateCloudbrainNumJson{}).
		Find(&DateCloudbrainNumJsons); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return DateCloudbrainNumJsons, nil
}

func GetQueueTaskDurationStatistics(opts *DurationStatisticOptions) ([]*QueueTaskDurationStatistic, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	if opts.BeginTime.Unix() > 0 && opts.EndTime.Unix() > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"queue_task_duration_statistic.date_time_unix": opts.BeginTime.Unix()}, builder.Lt{"queue_task_duration_statistic.date_time_unix": opts.EndTime.Unix()}),
		)
	}
	if opts.QueueID > 0 {
		cond = cond.And(
			builder.Eq{"queue_task_duration_statistic.queue_id": opts.QueueID},
		)
	}
	QueueTaskDurationStatistics := make([]*QueueTaskDurationStatistic, 0, 10)
	if err := sess.Table(&QueueTaskDurationStatistic{}).Where(cond).
		Find(&QueueTaskDurationStatistics); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return QueueTaskDurationStatistics, nil
}

func GetQueueTaskNumStatistics(opts *DurationStatisticOptions) ([]*QueueTaskNumStatistic, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	if opts.BeginTime.Unix() > 0 && opts.EndTime.Unix() > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"queue_task_num_statistic.date_time_unix": opts.BeginTime.Unix()}, builder.Lt{"queue_task_num_statistic.date_time_unix": opts.EndTime.Unix()}),
		)
	}
	QueueTaskNumStatistics := make([]*QueueTaskNumStatistic, 0, 10)
	if err := sess.Table(&QueueTaskNumStatistic{}).Where(cond).
		Find(&QueueTaskNumStatistics); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return QueueTaskNumStatistics, nil
}

func GetDurationRecordBeginTime() ([]*CloudbrainDurationStatistic, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()

	cond = cond.And(
		builder.Gt{"cloudbrain_duration_statistic.date_time_unix": 0},
	)

	sess.OrderBy("cloudbrain_duration_statistic.date_time_unix ASC limit 1")
	CloudbrainDurationStatistics := make([]*CloudbrainDurationStatistic, 0)
	if err := sess.Table(&CloudbrainDurationStatistic{}).Where(cond).Find(&CloudbrainDurationStatistics); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return CloudbrainDurationStatistics, nil
}

func GetDurationRecordUpdateTime() ([]*CloudbrainDurationStatistic, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()

	cond = cond.And(
		builder.Gt{"cloudbrain_duration_statistic.date_time_unix": 1577808000},
	)
	sess.OrderBy("cloudbrain_duration_statistic.date_time_unix DESC limit 1")
	CloudbrainDurationStatistics := make([]*CloudbrainDurationStatistic, 0)
	if err := sess.Table(&CloudbrainDurationStatistic{}).Where(cond).Find(&CloudbrainDurationStatistics); err != nil {
		log.Error("find error.")
		return nil, err
	}
	return CloudbrainDurationStatistics, nil
}

func DeleteCloudbrainDurationStatistic(beginTime timeutil.TimeStamp, endTime timeutil.TimeStamp) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("DELETE FROM cloudbrain_duration_statistic WHERE cloudbrain_duration_statistic.date_time_unix BETWEEN ? AND ?", beginTime, endTime); err != nil {
		log.Error("DELETE cloudbrain_duration_statistic data error.")
		return err
	}
	return nil
}

func DeleteCloudbrainTaskNumStatistic(beginTime timeutil.TimeStamp, endTime timeutil.TimeStamp) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("DELETE FROM cloudbrain_task_num_statistic WHERE cloudbrain_task_num_statistic.date_time_unix BETWEEN ? AND ?", beginTime, endTime); err != nil {
		log.Error("DELETE cloudbrain_task_num_statistic data error.")
		return err
	}
	return nil
}

func DeleteQueueTaskDurationStatistic(beginTime timeutil.TimeStamp, endTime timeutil.TimeStamp) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("DELETE FROM queue_task_duration_statistic WHERE queue_task_duration_statistic.date_time_unix BETWEEN ? AND ?", beginTime, endTime); err != nil {
		log.Error("DELETE queue_task_duration_statistic data error.")
		return err
	}
	return nil
}

func DeleteQueueTaskNumStatistic(beginTime timeutil.TimeStamp, endTime timeutil.TimeStamp) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("DELETE FROM queue_task_num_statistic WHERE queue_task_num_statistic.date_time_unix BETWEEN ? AND ?", beginTime, endTime); err != nil {
		log.Error("DELETE queue_task_num_statistic data error.")
		return err
	}
	return nil
}

func GetCloudbrainTypeCardDuration() (CloudbrainTypeDuration, error) {
	query := `
    SELECT
		cloudbrain.type,
		SUM(cloudbrain.duration) as durationSum,
        SUM(
            COALESCE(cloudbrain.duration *
				CASE
				WHEN cloudbrain.work_server_number = 0 THEN 1
				ELSE COALESCE(cloudbrain.work_server_number, 1)
				END *
				COALESCE(cloudbrain_spec.acc_cards_num, 1), 0)
   ) as cardDurationSum,
        COUNT(*) as count
    FROM cloudbrain
    LEFT JOIN cloudbrain_spec
        ON cloudbrain.id = cloudbrain_spec.cloudbrain_id
	GROUP BY cloudbrain.type
    `
	// 执行查询
	var results CloudbrainTypeDuration
	if err := x.SQL(query).Find(&results); err != nil {
		panic(err)
	}
	return results, nil
}

func GetCloudbrainAllCardDuration() (CloudbrainAllDuration, error) {
	query := `
    SELECT
		SUM(cloudbrain.duration) as durationSum,
        SUM(
            COALESCE(cloudbrain.duration *
				CASE
				WHEN cloudbrain.work_server_number = 0 THEN 1
				ELSE COALESCE(cloudbrain.work_server_number, 1)
				END *
				COALESCE(cloudbrain_spec.acc_cards_num, 1), 0)
   ) as cardDurationSum,
        COUNT(*) as count
    FROM cloudbrain
    LEFT JOIN cloudbrain_spec
        ON cloudbrain.id = cloudbrain_spec.cloudbrain_id
    `
	// 执行查询
	var result CloudbrainAllDuration
	if _, err := x.SQL(query).Get(&result); err != nil {
		panic(err)
	}
	return result, nil
}

func GetActiveAccCardType(startTime, endTime time.Time) ([]AccCardInfo, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	queryRes := make([]CloudbrainDurationStatistic, 0)
	err := sess.Table("cloudbrain_duration_statistic").
		Distinct("compute_resource", "acc_card_type").
		Where("cards_use_duration > 0 and date_time_unix >= ? and date_time_unix <= ? and compute_resource is not null and acc_card_type is not null and compute_resource != '' and acc_card_type != ''", startTime.Unix(), endTime.Unix()).
		OrderBy("compute_resource asc,acc_card_type asc").
		Find(&queryRes)
	if err != nil {
		return nil, err
	}

	result := make([]AccCardInfo, 0)
	tmpMap := make(map[string][]string, 0)
	for i := 0; i < len(queryRes); i++ {
		tmpQueue := queryRes[i]
		computeResource := tmpQueue.ComputeResource
		accCardType := tmpQueue.AccCardType
		if computeResource == "" || accCardType == "" {
			continue
		}
		if computeResource == ILUVATAR || computeResource == METAX {
			computeResource = GPGPU
		}
		if _, ok := tmpMap[computeResource]; ok {
			tmpMap[computeResource] = append(tmpMap[computeResource], accCardType)
		} else {
			tmpMap[computeResource] = []string{accCardType}
		}
	}
	for k, v := range tmpMap {
		sort.Strings(v)
		accCard := AccCardInfo{
			ComputeSource: k,
			CardList:      v,
		}
		result = append(result, accCard)
	}
	sort.Sort(AccCardInfoList(result))
	return result, nil
}

func GetActiveAICenterList(startTime, endTime time.Time) ([]string, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	centerCodeList := make([]string, 0)
	err := sess.Table("cloudbrain_duration_statistic").
		Distinct("ai_center_code").
		Where("cards_use_duration > 0 and cluster = ? and date_time_unix >= ? and date_time_unix <= ?", C2NetCluster, startTime.Unix(), endTime.Unix()).
		Find(&centerCodeList)
	return centerCodeList, err
}

func GetAICenterTotalCardsUsageMap(centerCodeList []string, starTime, endTime time.Time) (map[string]int, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	centerList := make([]AICenterWithCardsUsage, 0)
	err := sess.Table("cloudbrain_duration_statistic").
		Select("ai_center_code,sum(cards_use_duration) as total_cards_usage_num").
		Where(builder.And(
			builder.In("ai_center_code", centerCodeList),
			builder.Eq{"cluster": C2NetCluster},
			builder.Gte{"date_time_unix": starTime.Unix()},
			builder.Lte{"date_time_unix": endTime.Unix()},
		)).
		GroupBy("ai_center_code").
		Find(&centerList)

	if err != nil {
		return nil, err
	}
	tmpMap := make(map[string]int, 0)
	for _, v := range centerList {
		tmpMap[v.AICenterCode] = v.TotalCardsUsageNum
	}
	result := make(map[string]int, len(centerCodeList))
	for i := 0; i < len(centerCodeList); i++ {
		centerCode := centerCodeList[i]
		if _, ok := tmpMap[centerCode]; ok {
			result[centerCode] = tmpMap[centerCode]
		} else {
			result[centerCode] = 0
		}
	}
	return result, nil
}

func GetAICenterComputeResourceMap(centerCodeList []string, starTime, endTime time.Time) (map[string][]string, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	centerList := make([]CloudbrainDurationStatistic, 0)
	err := sess.Distinct("ai_center_code", "compute_resource").
		Where(builder.And(
			builder.In("ai_center_code", centerCodeList),
			builder.Eq{"cluster": C2NetCluster},
			builder.Gte{"date_time_unix": starTime.Unix()},
			builder.Lte{"date_time_unix": endTime.Unix()},
		)).
		Find(&centerList)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string][]string, 0)
	for _, v := range centerList {
		computeResource := v.ComputeResource
		if computeResource == "" {
			continue
		}
		if computeResource == ILUVATAR || computeResource == METAX {
			computeResource = GPGPU
		}
		if _, ok := resultMap[v.AiCenterCode]; ok {
			isExists := false
			for _, v := range resultMap[v.AiCenterCode] {
				if v == computeResource {
					isExists = true
					break
				}
			}
			if !isExists {
				resultMap[v.AiCenterCode] = append(resultMap[v.AiCenterCode], computeResource)
			}
		} else {
			resultMap[v.AiCenterCode] = []string{computeResource}
		}
	}
	return resultMap, nil
}

func GetAICenterCardTypeMap(centerCodeList []string, starTime, endTime time.Time) (map[string][]string, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	centerList := make([]CloudbrainDurationStatistic, 0)
	err := sess.Distinct("ai_center_code", "acc_card_type").
		Where(builder.And(
			builder.In("ai_center_code", centerCodeList),
			builder.Eq{"cluster": C2NetCluster},
			builder.Gte{"date_time_unix": starTime.Unix()},
			builder.Lte{"date_time_unix": endTime.Unix()},
		)).
		Find(&centerList)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string][]string, 0)
	for _, v := range centerList {
		if v.AccCardType == "" {
			continue
		}
		if _, ok := resultMap[v.AiCenterCode]; ok {
			resultMap[v.AiCenterCode] = append(resultMap[v.AiCenterCode], v.AccCardType)
		} else {
			resultMap[v.AiCenterCode] = []string{v.AccCardType}
		}
	}
	return resultMap, nil
}

func GetCloudbrainRunByUserId(uid int64) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	sess := x.Where("user_id=?", uid)
	sess.In("status", JobWaiting, JobRunning, ModelArtsCreateQueue, ModelArtsCreating, ModelArtsStarting,
		ModelArtsReadyToStart, ModelArtsResizing, ModelArtsStartQueuing, ModelArtsRunning, ModelArtsDeleting, ModelArtsRestarting, ModelArtsTrainJobInit,
		ModelArtsTrainJobImageCreating, ModelArtsTrainJobSubmitTrying, ModelArtsTrainJobWaiting, ModelArtsTrainJobRunning, ModelArtsStopping, ModelArtsResizing,
		ModelArtsTrainJobScaling, ModelArtsTrainJobCheckInit, ModelArtsTrainJobCheckRunning, ModelArtsTrainJobKilling, ModelArtsTrainJobCheckRunningCompleted)
	err := sess.Find(&cloudbrains)
	return cloudbrains, err
}
