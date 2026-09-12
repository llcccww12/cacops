package models

import (
	"code.gitea.io/gitea/modules/timeutil"
)

//用户活跃信息表
type UserLoginActionLog struct {
	ID          int64              `xorm:"pk autoincr"`
	UId         int64              `xorm:"NOT NULL"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

func SaveLoginActionToDb(uid int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	var dateRecord UserLoginActionLog
	dateRecord.UId = uid
	statictisSess.Insert(&dateRecord)
}
