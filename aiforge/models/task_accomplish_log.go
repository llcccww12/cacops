package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"time"
)

type TaskAccomplishLog struct {
	ID          int64  `xorm:"pk autoincr"`
	ConfigId    int64  `xorm:"NOT NULL"`
	TaskCode    string `xorm:"NOT NULL"`
	UserId      int64  `xorm:"INDEX NOT NULL"`
	ActionId    int64
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

type PeriodResult struct {
	StartTime time.Time
	EndTime   time.Time
	LeftTime  time.Duration
}

func getTaskAccomplishLog(tl *TaskAccomplishLog) (*TaskAccomplishLog, error) {
	has, err := x.Get(tl)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return tl, nil
}

func CountTaskAccomplishLogInTaskPeriod(taskCode string, userId int64, period *PeriodResult) (float64, error) {
	var r int64
	var err error
	if period == nil {
		r, err = x.Where("task_code = ? and user_id = ?", taskCode, userId).Count(&TaskAccomplishLog{})
	} else {
		r, err = x.Where("task_code = ? and user_id = ? and created_unix >= ? and created_unix < ? ", taskCode, userId, period.StartTime.Unix(), period.EndTime.Unix()).Count(&TaskAccomplishLog{})
	}
	return float64(r), err
}

func InsertTaskAccomplishLog(tl *TaskAccomplishLog) (int64, error) {
	return x.Insert(tl)
}
