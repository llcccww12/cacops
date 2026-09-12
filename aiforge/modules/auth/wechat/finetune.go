package wechat

import (
	"fmt"
	"time"

	"code.gitea.io/gitea/modules/setting"
)

type FinetuneStartMsg struct {
}

var FinetuneMsg = &FinetuneStartMsg{}

func (FinetuneStartMsg) Data(ctx *TemplateContext) interface{} {
	return &DefaultWechatTemplate{
		First:    TemplateValue{Value: setting.FineTune.Pangu.Wechat.Title},
		Keyword1: TemplateValue{Value: ctx.ModelartsDeploy.DisplayJobName},
		Keyword2: TemplateValue{Value: setting.FineTune.Pangu.Wechat.JobType},
		Keyword3: TemplateValue{Value: time.Unix(int64(ctx.ModelartsDeploy.UpdateUnix), 0).Format("2006-01-02 15:04:05")},
		Remark:   TemplateValue{Value: setting.CloudbrainStartedRemark},
	}
}

func (FinetuneStartMsg) ShouldSend(ctx *TemplateContext) bool {
	return setting.FineTune.Pangu.Wechat.Flag
}

func (FinetuneStartMsg) MsgId(ctx *TemplateContext) string {
	return "finetune_start" + "_" + fmt.Sprint(ctx.ModelartsDeploy.JobID)
}

func (FinetuneStartMsg) Url(ctx *TemplateContext) string {
	// http://192.168.207.34:8094/extension/modelbase/pangufinetune/inference?jobid=158204&type=0&jobcategory=1
	url := setting.AppURL + "extension/modelbase/pangufinetune/inference?"
	jobID := "jobid=" + fmt.Sprint(ctx.ModelartsDeploy.JobID)
	jobType := "type=" + fmt.Sprint(ctx.ModelartsDeploy.FinetuneModelType)
	jobCategory := "jobcategory=" + fmt.Sprint(ctx.ModelartsDeploy.FinetuneCategory)

	return url + jobID + "&" + jobType + "&" + jobCategory
}

func (FinetuneStartMsg) TemplateId(ctx *TemplateContext) string {
	return setting.FineTune.Pangu.Wechat.TemplateID
}
