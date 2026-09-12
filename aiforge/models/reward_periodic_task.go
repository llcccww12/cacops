package models

import (
	"time"

	"code.gitea.io/gitea/modules/timeutil"
)

type PeriodicTaskStatus int

const (
	PeriodicTaskStatusRunning  = iota + 1 // 1
	PeriodicTaskStatusFinished            // 2
)

type PeriodType string

const (
	PeriodType30MinutesFree1HourCost PeriodType = "30MF1HC"
)

func (r PeriodType) Name() string {
	switch r {
	case PeriodType30MinutesFree1HourCost:
		return "30MF1HC"
	default:
		return ""
	}
}

type RewardPeriodicTask struct {
	ID              int64  `xorm:"pk autoincr"`
	OperateSerialNo string `xorm:"INDEX NOT NULL"`
	DelaySeconds    int64
	IntervalSeconds int64
	Amount          float64            `xorm:"NOT NULL"`
	NextExecuteTime timeutil.TimeStamp `xorm:"INDEX NOT NULL"`
	SuccessCount    int                `xorm:"NOT NULL default 0"`
	Status          int                `xorm:"NOT NULL"`
	CreatedUnix     timeutil.TimeStamp `xorm:"INDEX created"`
	FinishedUnix    timeutil.TimeStamp `xorm:"INDEX"`
	UpdatedUnix     timeutil.TimeStamp `xorm:"INDEX updated"`
}

type StartPeriodicTaskOpts struct {
	SourceType    SourceType
	SourceId      string
	SourceContent string
	Remark        string
	Title         string
	TargetUserId  int64
	RequestId     string
	OperateType   RewardOperateType
	Delay         time.Duration
	Interval      time.Duration
	UnitAmount    float64
	RewardType    RewardType
	StartTime     time.Time
}

func InsertPeriodicTask(tl *RewardPeriodicTask) (int64, error) {
	return x.Insert(tl)
}

func GetRunningRewardTask(now time.Time) ([]RewardPeriodicTask, error) {
	r := make([]RewardPeriodicTask, 0)
	err := x.Where("next_execute_time <= ? and status = ?", now.Unix(), PeriodicTaskStatusRunning).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, err
}

func IncrRewardTaskSuccessCount(t RewardPeriodicTask, count int64, nextTime timeutil.TimeStamp) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Exec("update reward_periodic_task set success_count = success_count + ? , next_execute_time = ?, updated_unix = ? where id = ?", count, nextTime, timeutil.TimeStampNow(), t.ID)
	if err != nil {
		sess.Rollback()
		return err
	}
	_, err = sess.Exec("update reward_operate_record set amount = amount + ? ,updated_unix = ? ,last_operate_unix = ? where serial_no = ?", t.Amount, timeutil.TimeStampNow(), timeutil.TimeStampNow(), t.OperateSerialNo)
	if err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

func GetPeriodicTaskBySourceIdAndType(sourceType SourceType, sourceId string, operateType RewardOperateType) (*RewardPeriodicTask, error) {
	r := RewardPeriodicTask{}
	_, err := x.SQL("select rpt.* from reward_periodic_task rpt "+
		"inner join reward_operate_record ror on rpt.operate_serial_no = ror.serial_no"+
		" where ror.source_type = ? and ror.source_id = ? and ror.operate_type = ? ", sourceType.Name(), sourceId, operateType.Name()).Get(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func StopPeriodicTask(taskId int64, operateSerialNo string, stopTime time.Time, sourceContent string) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Where("id = ? and status = ?", taskId, PeriodicTaskStatusRunning).Update(&RewardPeriodicTask{Status: PeriodicTaskStatusFinished, FinishedUnix: timeutil.TimeStamp(stopTime.Unix())})
	if err != nil {
		sess.Rollback()
		return err
	}
	_, err = sess.Where("serial_no = ? and status = ?", operateSerialNo, OperateStatusOperating).Update(&RewardOperateRecord{Status: OperateStatusSucceeded, SourceContent: sourceContent})
	if err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}
