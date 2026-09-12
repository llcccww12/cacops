package models

import "code.gitea.io/gitea/modules/timeutil"

const (
	IncreaseAccountBalance = "increase"
	DecreaseAccountBalance = "decrease"
)

type PointAccountLog struct {
	ID             int64              `xorm:"pk autoincr"`
	AccountCode    string             `xorm:"INDEX NOT NULL"`
	UserId         int64              `xorm:"INDEX NOT NULL"`
	Type           string             `xorm:"NOT NULL"`
	SourceId       string             `xorm:"INDEX NOT NULL"`
	PointsAmount   float64            `xorm:"NOT NULL"`
	BalanceBefore  float64            `xorm:"NOT NULL"`
	BalanceAfter   float64            `xorm:"NOT NULL"`
	AccountVersion int64              `xorm:"NOT NULL"`
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
}
