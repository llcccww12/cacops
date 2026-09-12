package models

import (
	"encoding/json"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

type AdminOperateLog struct {
	ID              int64 `xorm:"pk autoincr"`
	BizType         string
	OperateType     string
	OldValue        string `xorm:"TEXT"`
	NewValue        string `xorm:"TEXT"`
	RelatedId       string `xorm:"INDEX"`
	Comment         string
	CreatedTime     timeutil.TimeStamp `xorm:"created"`
	CreatedBy       int64
	ComputeResource string
	AccCardType     string
	AccCardNum      int
	AiCenterCode    string
	AiCenterName    string
	ResourceQueueID int64
}

type LogValues struct {
	Params []LogValue
}

type LogValue struct {
	Key string
	Val interface{}
}

func (l *LogValues) Add(key string, val interface{}) *LogValues {
	l.Params = append(l.Params, LogValue{Key: key, Val: val})
	return l
}

func (l *LogValues) JsonString() string {
	if len(l.Params) == 0 {
		return ""
	}
	b, err := json.Marshal(l)
	if err != nil {
		log.Error("LogValues JsonString error . %v", err)
		return ""
	}
	return string(b)
}

func InsertAdminOperateLog(log AdminOperateLog) (int64, error) {
	return x.Insert(&log)
}
