package operate_log

import (
	"code.gitea.io/gitea/models"
)

type LogBizType string

const (
	BadgeCategoryOperate LogBizType = "BadgeCategoryOperate"
	BadgeOperate         LogBizType = "BadgeOperate"
)

func Log(log models.AdminOperateLog) error {
	_, err := models.InsertAdminOperateLog(log)
	return err
}

func NewLogValues() *models.LogValues {
	return &models.LogValues{Params: make([]models.LogValue, 0)}
}

func Log4Add(bizType LogBizType, newValue interface{}, doerId int64, comment string) {
	Log(models.AdminOperateLog{
		BizType:     string(bizType),
		OperateType: "add",
		NewValue:    NewLogValues().Add("new", newValue).JsonString(),
		CreatedBy:   doerId,
		Comment:     comment,
	})
}

func Log4Edit(bizType LogBizType, oldValue interface{}, newValue interface{}, doerId int64, comment string) {
	Log(models.AdminOperateLog{
		BizType:     string(bizType),
		OperateType: "edit",
		NewValue:    NewLogValues().Add("new", newValue).JsonString(),
		OldValue:    NewLogValues().Add("old", oldValue).JsonString(),
		CreatedBy:   doerId,
		Comment:     comment,
	})
}

func Log4Del(bizType LogBizType, oldValue interface{}, doerId int64, comment string) {
	Log(models.AdminOperateLog{
		BizType:     string(bizType),
		OperateType: "del",
		OldValue:    NewLogValues().Add("old", oldValue).JsonString(),
		CreatedBy:   doerId,
		Comment:     comment,
	})
}
