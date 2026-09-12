package wechat

import (
	"fmt"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
)

type JobOperateType string

const (
	JobOperateTypeStart        JobOperateType = "start"
	JobOperateTypeStop         JobOperateType = "stop"
	JobOperateTypeLongRunning  JobOperateType = "long_running"
	JobOperateTypeAlmostEnding JobOperateType = "almost_ending"
)

type CloudbrainStartMsg struct {
}

func (CloudbrainStartMsg) Data(ctx *TemplateContext) interface{} {
	return &DefaultWechatTemplate{
		First:    TemplateValue{Value: setting.CloudbrainStartedTitle},
		Keyword1: TemplateValue{Value: ctx.Cloudbrain.DisplayJobName},
		Keyword2: TemplateValue{Value: getJobTypeDisplayName(ctx.Cloudbrain.JobType)},
		Keyword3: TemplateValue{Value: time.Unix(int64(ctx.Cloudbrain.CreatedUnix), 0).Format("2006-01-02 15:04:05")},
		Remark:   TemplateValue{Value: setting.CloudbrainStartedRemark},
	}
}

func (CloudbrainStartMsg) ShouldSend(ctx *TemplateContext) bool {
	if len(setting.CloudbrainStartedNotifyList) == 0 {
		return false
	}
	for _, v := range setting.CloudbrainStartedNotifyList {
		if v == ctx.Cloudbrain.JobType {
			return true
		}
	}
	return false
}

func (CloudbrainStartMsg) MsgId(ctx *TemplateContext) string {
	return string(JobOperateTypeStart) + "_" + fmt.Sprint(ctx.Cloudbrain.ID)
}

func (CloudbrainStartMsg) Url(ctx *TemplateContext) string {
	return getCloudbrainTemplateUrl(*ctx.Cloudbrain)
}

func (CloudbrainStartMsg) TemplateId(ctx *TemplateContext) string {
	return setting.CloudbrainStartedTemplateId
}

type CloudbrainStopMsg struct {
}

func (CloudbrainStopMsg) Data(ctx *TemplateContext) interface{} {
	return &DefaultWechatTemplate{
		First:    TemplateValue{Value: fmt.Sprintf(setting.CloudbrainStoppedTitle, ctx.Cloudbrain.Status)},
		Keyword1: TemplateValue{Value: ctx.Cloudbrain.DisplayJobName},
		Keyword2: TemplateValue{Value: getJobTypeDisplayName(ctx.Cloudbrain.JobType)},
		Keyword3: TemplateValue{Value: time.Unix(int64(ctx.Cloudbrain.CreatedUnix), 0).Format("2006-01-02 15:04:05")},
		Keyword4: TemplateValue{Value: time.Unix(int64(ctx.Cloudbrain.EndTime), 0).Format("2006-01-02 15:04:05")},
		Remark:   TemplateValue{Value: setting.CloudbrainStoppedRemark},
	}
}

func (CloudbrainStopMsg) ShouldSend(ctx *TemplateContext) bool {
	if len(setting.CloudbrainStoppedNotifyList) == 0 {
		return false
	}
	for _, v := range setting.CloudbrainStoppedNotifyList {
		if v == ctx.Cloudbrain.JobType {
			if ctx.Cloudbrain.Duration > 0 && ctx.Cloudbrain.EndTime > 0 {
				return true
			}
			break
		}
	}
	return false
}

func (CloudbrainStopMsg) MsgId(ctx *TemplateContext) string {
	return string(JobOperateTypeStop) + "_" + fmt.Sprint(ctx.Cloudbrain.ID)
}

func (CloudbrainStopMsg) Url(ctx *TemplateContext) string {
	return getCloudbrainTemplateUrl(*ctx.Cloudbrain)
}

func (CloudbrainStopMsg) TemplateId(ctx *TemplateContext) string {
	return setting.CloudbrainStoppedTemplateId
}

var startMsg = &CloudbrainStartMsg{}
var stopMsg = &CloudbrainStopMsg{}
var ComingStopMsg = &CloudbrainComingToStopMsg{}
var LongRunningMsg = &CloudbrainLongRunningMsg{}
var AlmostEndingMsg = &CloudbrainAlmostEndingMsg{}

func GetTemplateFromOperateType(operate JobOperateType) Template {
	switch operate {
	case JobOperateTypeStart:
		return startMsg
	case JobOperateTypeStop:
		return stopMsg
	}
	return nil
}

func GetJobOperateTypeFromCloudbrainStatus(cloudbrain *models.Cloudbrain) JobOperateType {
	if cloudbrain.IsTerminal() {
		return JobOperateTypeStop
	}
	if cloudbrain.IsRunning() {
		return JobOperateTypeStart
	}
	return ""
}

func getCloudbrainTemplateUrl(cloudbrain models.Cloudbrain) string {
	url := strings.TrimSuffix(setting.AppURL, "/") + "/cloudbrains/detail/" + fmt.Sprint(cloudbrain.ID)
	return url
}

func getJobTypeDisplayName(jobType string) string {
	switch jobType {
	case string(models.JobTypeDebug):
		return "调试任务"
	case string(models.JobTypeOnlineInference):
		return "在线推理"
	case string(models.JobTypeBenchmark):
		return "评测任务"
	case string(models.JobTypeTrain):
		return "训练任务"
	case string(models.JobTypeInference),
		string(models.JobTypeModelSafety),
		string(models.JobTypeSnn4imagenet),
		string(models.JobTypeBrainScore),
		string(models.JobTypeSnn4Ecoset):
		return "推理任务"
	case string(models.JobTypeGeneral):
		return "通用任务"
	case string(models.JobTypeModelExperience):
		return "在线体验"
	case string(models.JobTypeFINETUNE):
		return "NLP微调任务"
	case string(models.JobTypeSdFinetune):
		return "CV微调任务"
	case string(models.JobTypeComfyuiExperience):
		return "Comfyui体验任务"
	}

	return ""
}

type CloudbrainLongRunningMsg struct {
}

func (CloudbrainLongRunningMsg) Data(ctx *TemplateContext) interface{} {
	formattedDuration := formatDuration(ctx.Duration)
	formattedRemark := ""
	if formattedDuration == "" {
		formattedRemark = "任务运行中"
	} else {
		formattedRemark = fmt.Sprintf("任务运行时间超过%s", formattedDuration)
	}
	d := int64(ctx.Duration.Seconds())
	t := int64(ctx.Cloudbrain.StartTime) + d
	return &WechatDurationWechatTemplate{
		Keyword1: TemplateValue{Value: ctx.Cloudbrain.DisplayJobName},
		Keyword2: TemplateValue{Value: time.Unix(t, 0).Format("2006-01-02 15:04:05")},
		Keyword3: TemplateValue{Value: formattedRemark},
	}
}

func formatDuration(duration time.Duration) string {
	switch {
	case duration >= time.Second && duration < time.Minute:
		return fmt.Sprintf("%d秒", int(duration.Seconds()))
	case duration >= time.Minute && duration < time.Hour:
		return fmt.Sprintf("%d分钟", int(duration.Minutes()))
	case duration >= time.Hour && duration < 24*time.Hour:
		return fmt.Sprintf("%d小时", int(duration.Hours()))
	case duration >= 24*time.Hour:
		days := duration / (24 * time.Hour)
		return fmt.Sprintf("%d天", days)
	default:
		return ""
	}
}

func (CloudbrainLongRunningMsg) ShouldSend(ctx *TemplateContext) bool {
	if len(setting.CloudbrainLongRunningNotifyList) == 0 {
		return false
	}
	for _, v := range setting.CloudbrainLongRunningNotifyList {
		if v == ctx.Cloudbrain.JobType {
			return true
		}
	}
	return false
}

func (CloudbrainLongRunningMsg) MsgId(ctx *TemplateContext) string {
	return string(JobOperateTypeLongRunning) + "_" + fmt.Sprint(ctx.Cloudbrain.ID) + "_" + ctx.Duration.String()
}

func (CloudbrainLongRunningMsg) Url(ctx *TemplateContext) string {
	return getCloudbrainTemplateUrl(*ctx.Cloudbrain)
}

func (CloudbrainLongRunningMsg) TemplateId(ctx *TemplateContext) string {
	return setting.CloudbrainLongRunningTemplateId
}

type CloudbrainAlmostEndingMsg struct {
}

func (CloudbrainAlmostEndingMsg) Data(ctx *TemplateContext) interface{} {
	formattedDuration := formatDuration(ctx.Duration)
	formattedRemark := ""
	if formattedDuration == "" {
		formattedRemark = "任务运行中"
	} else {
		formattedRemark = fmt.Sprintf("%s后调试任务将被关闭,请及时保存", formattedDuration)
	}
	d := int64(ctx.Duration.Seconds())
	t := int64(ctx.Cloudbrain.StartTime) + d
	return &WechatDurationWechatTemplate{
		Keyword1: TemplateValue{Value: ctx.Cloudbrain.DisplayJobName},
		Keyword2: TemplateValue{Value: time.Unix(t, 0).Format("2006-01-02 15:04:05")},
		Keyword3: TemplateValue{Value: formattedRemark},
	}
}

func (CloudbrainAlmostEndingMsg) ShouldSend(ctx *TemplateContext) bool {
	return true
}

func (CloudbrainAlmostEndingMsg) MsgId(ctx *TemplateContext) string {
	return string(JobOperateTypeAlmostEnding) + "_" + fmt.Sprint(ctx.Cloudbrain.ID) + "_" + ctx.Duration.String()
}

func (CloudbrainAlmostEndingMsg) Url(ctx *TemplateContext) string {
	return getCloudbrainTemplateUrl(*ctx.Cloudbrain)
}

func (CloudbrainAlmostEndingMsg) TemplateId(ctx *TemplateContext) string {
	return setting.CloudbrainLongRunningTemplateId
}
