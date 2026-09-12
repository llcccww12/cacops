package wechat

import (
	"fmt"
	"time"

	"code.gitea.io/gitea/modules/setting"
)

type CloudbrainComingToStopMsg struct {
}

func (CloudbrainComingToStopMsg) Data(ctx *TemplateContext) interface{} {
	var balance float64
	if ctx.PointAccount != nil {
		balance = ctx.PointAccount.Balance
	}
	return &DefaultWechatTemplate{
		First:    TemplateValue{Value: setting.CloudbrainComingStopTitle},
		Keyword1: TemplateValue{Value: ctx.Cloudbrain.DisplayJobName + "(任务名称)"},
		Keyword2: TemplateValue{Value: fmt.Sprintf("%.2f积分", balance)},
		Keyword3: TemplateValue{Value: setting.CloudbrainComingStopChargeLink},
		Keyword4: TemplateValue{Value: time.Unix(int64(ctx.EstimatedEndTime), 0).Format("2006-01-02 15:04:05") + "(预计停止时间)"},
		Remark:   TemplateValue{Value: setting.CloudbrainComingStopRemark},
	}
}

func (CloudbrainComingToStopMsg) ShouldSend(ctx *TemplateContext) bool {
	return setting.CloudbrainComingStopSendFlag
}

func (CloudbrainComingToStopMsg) MsgId(ctx *TemplateContext) string {
	return "coming_stop_" + fmt.Sprint(ctx.Cloudbrain.ID) + "_" + fmt.Sprint(time.Now().Unix())
}

func (CloudbrainComingToStopMsg) Url(ctx *TemplateContext) string {
	return getCloudbrainTemplateUrl(*ctx.Cloudbrain)
}

func (CloudbrainComingToStopMsg) TemplateId(ctx *TemplateContext) string {
	return setting.CloudbrainComingStopTemplateId
}
