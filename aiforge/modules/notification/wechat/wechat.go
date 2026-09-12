// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wechat

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth/wechat"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification/base"
	"code.gitea.io/gitea/modules/timeutil"
	"time"
)

type wechatNotifier struct {
	base.NullNotifier
}

var (
	_ base.Notifier = &wechatNotifier{}
)

// NewNotifier create a new wechatNotifier notifier
func NewNotifier() base.Notifier {
	return &wechatNotifier{}
}

func (*wechatNotifier) NotifyChangeCloudbrainStatus(cloudbrain *models.Cloudbrain, oldStatus string) {
	log.Info("NotifyChangeCloudbrainStatus cloudbrain.id=%d cloudbrain.status=%s oldStatus=%s", cloudbrain.ID, cloudbrain.Status, oldStatus)
	operateType := wechat.GetJobOperateTypeFromCloudbrainStatus(cloudbrain)
	if operateType == "" {
		log.Info("NotifyChangeCloudbrainStatus operateType is incorrect")
		return
	}
	template := wechat.GetTemplateFromOperateType(operateType)
	go wechat.SendTemplateMsg(template, &wechat.TemplateContext{Cloudbrain: cloudbrain}, cloudbrain.UserID)
}

func (*wechatNotifier) NotifyCloudbrainTaskComingToFinished(cloudbrain *models.Cloudbrain, endTime timeutil.TimeStamp, account *models.PointAccount) {
	log.Info("NotifyCloudbrainTaskComingToFinished cloudbrain.id=%d", cloudbrain.ID)
	template := wechat.ComingStopMsg
	go wechat.SendTemplateMsg(template, &wechat.TemplateContext{Cloudbrain: cloudbrain, EstimatedEndTime: endTime, PointAccount: account}, cloudbrain.UserID)
}

func (*wechatNotifier) NotifyChangeFinetuneStatus(deployment *models.ModelartsDeploy) {
	log.Info("盘古微调部署: NotifyChangeFineTuneStatus deployment.jobid=%d deployment.status=%s", deployment.JobID, deployment.Status)
	template := wechat.FinetuneMsg
	go wechat.SendTemplateMsg(template, &wechat.TemplateContext{ModelartsDeploy: deployment}, deployment.UserID)
}

func (*wechatNotifier) NotifyLongRunningAITask(cloudbrain *models.Cloudbrain, duration time.Duration) {
	log.Info("NotifyAITaskDuration cloudbrain.id=%d", cloudbrain.ID)
	template := wechat.LongRunningMsg
	go wechat.SendTemplateMsg(template, &wechat.TemplateContext{Cloudbrain: cloudbrain, Duration: duration}, cloudbrain.UserID)
}

func (*wechatNotifier) NotifyAlmostEndingAIDebugTask(cloudbrain *models.Cloudbrain, duration time.Duration) {
	log.Info("NotifyAlmostEndingAIDebugTask cloudbrain.id=%d", cloudbrain.ID)
	template := wechat.AlmostEndingMsg
	go wechat.SendTemplateMsg(template, &wechat.TemplateContext{Cloudbrain: cloudbrain, Duration: duration}, cloudbrain.UserID)

}
