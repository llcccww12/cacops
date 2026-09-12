package wechat

import (
	"errors"
	"fmt"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

type Template interface {
	ShouldSend(ctx *TemplateContext) bool
	Data(ctx *TemplateContext) interface{}
	MsgId(ctx *TemplateContext) string
	Url(ctx *TemplateContext) string
	TemplateId(ctx *TemplateContext) string
}

type TemplateContext struct {
	Cloudbrain       *models.Cloudbrain
	EstimatedEndTime timeutil.TimeStamp
	PointAccount     *models.PointAccount
	ModelartsDeploy  *models.ModelartsDeploy
	Duration         time.Duration
}

func SendTemplateMsg(template Template, ctx *TemplateContext, userId int64) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:", combinedErr)
		}
	}()
	if !template.ShouldSend(ctx) {
		log.Info("SendTemplateMsg should not Send.jobId=%d  jobType=%s", ctx.Cloudbrain.ID, ctx.Cloudbrain.JobType)
		return nil
	}

	openId := models.GetUserWechatOpenId(userId)
	if openId == "" {
		log.Error("Wechat openId not exist,userId=%d", userId)
		return errors.New("Wechat openId not exist")
	}
	d := template.Data(ctx)
	req := TemplateMsgRequest{
		ToUser:      openId,
		TemplateId:  template.TemplateId(ctx),
		Url:         template.Url(ctx),
		ClientMsgId: template.MsgId(ctx),
		Data:        d,
	}
	log.Info("SendTemplateMsg before calling.userId = %d,req=%v data=%v", userId, req, d)
	err, retryFlag := sendTemplateMsg(req)
	if retryFlag {
		log.Info("SendTemplateMsg calling.userId = %d", userId)
		refreshAccessToken()
		err, _ = sendTemplateMsg(req)
		if err != nil {
			log.Error("SendTemplateMsg userId = %d err. %v", userId, err)
			return err
		}
		return nil
	}
	if err != nil {
		log.Error("SendTemplateMsg userId = %d err. %v", userId, err)
		return err
	}
	log.Info("SendTemplateMsg success userId = %d", userId)
	return nil
}
