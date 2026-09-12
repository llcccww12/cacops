package task

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/reward"
	"code.gitea.io/gitea/services/reward/limiter"
	"encoding/json"
	"fmt"
)

func Accomplish(action models.Action) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	taskType := models.GetTaskTypeFromAction(action.OpType)
	if taskType == "" {
		log.Info("Accomplish finished.taskType is not exist.action.ID=%d", action.ID)
		return
	}
	actions := make([]models.Action, 0)
	actions = append(actions, action)
	switch taskType {
	//only creating public repo can be rewarded
	case models.TaskCreatePublicRepo:
		if action.Repo.IsPrivate {
			return
		}
	//only creating public image can be rewarded
	case models.TaskCreateImage:
		if action.IsPrivate {
			return
		}
	case models.TaskBindWechat:
		n, err := models.CountWechatBindLog(action.Content, models.WECHAT_BIND)
		if err != nil {
			log.Error("CountWechatBindLog error when accomplish task,err=%v", err)
			return
		}
		//if wechatOpenId has been bound before,the action can not get reward
		if n > 1 && models.IsWechatOpenIdRewarded(action.Content) {

			log.Debug("the wechat account has been bound before,wechatOpenId = %s", action.Content)
			return
		}
	case models.TaskDatasetRecommended:
		actions[0].UserID = action.ActUserID
	case models.TaskCreateDataset:
		if action.DatasetID == nil || *action.DatasetID == "" {
			return
		}
		dataset, _ := models.GetDatasetRegistryByID(*action.DatasetID)
		if dataset == nil {
			return
		}
		dataset.GetOwner()
		if dataset.Owner == nil {
			return
		}
	case models.TaskAimodelRecommended:
		actions[0].UserID = action.ActUserID
	case models.TaskCreateAimodel:
		if action.AimodelID == nil || *action.AimodelID == "" {
			return
		}
		aimodel, _ := models.GetAimodelByID(*action.AimodelID)
		if aimodel == nil {
			return
		}
		aimodel.GetOwner()
		if aimodel.Owner == nil {
			return
		}
	}
	batchAccomplish(taskType, actions...)
}

func batchAccomplish(taskType models.TaskType, actions ...models.Action) {
	for _, act := range actions {
		go accomplish(act, taskType)
	}
}

func accomplish(action models.Action, taskType models.TaskType) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	log.Info("accomplish start.   actionId=%d  userId= %d", action.ID, action.UserID)
	userId := action.UserID
	if !isUserAvailable(userId) {
		return nil
	}

	//get task config
	config, err := GetTaskConfig(string(taskType))
	if err != nil {
		log.Error("GetTaskConfig error,%v", err)
		return err
	}
	if config == nil {
		log.Info("task config not exist,userId=%d taskType=%s", userId, taskType)
		return nil
	}

	//is limited?
	if isLimited(userId, config, models.JustReject) {
		log.Info("task accomplish maximum times are reached,userId=%d taskType=%s", userId, taskType)
		return nil
	}

	//add log
	_, err = models.InsertTaskAccomplishLog(&models.TaskAccomplishLog{
		ConfigId: config.ID,
		TaskCode: config.TaskCode,
		UserId:   userId,
		ActionId: action.ID,
	})
	if err != nil {
		log.Error("InsertTaskAccomplishLog error,%v", err)
		return err
	}

	var sourceContent string
	if contentByte, err := json.Marshal(action.ToShow()); err == nil {
		sourceContent = string(contentByte)
	}
	//reward
	reward.Operate(&models.RewardOperateContext{
		SourceType:       models.SourceTypeAccomplishTask,
		SourceId:         fmt.Sprint(action.ID),
		SourceTemplateId: string(taskType),
		Title:            config.Title,
		Reward: models.Reward{
			Amount: config.AwardAmount,
			Type:   models.GetRewardTypeInstance(config.AwardType),
		},
		TargetUserId:  userId,
		RequestId:     fmt.Sprintf("%d_%d", action.ID, userId),
		OperateType:   models.OperateTypeIncrease,
		RejectPolicy:  models.FillUp,
		SourceContent: sourceContent,
	})
	log.Debug("accomplish success,action=%v", action)
	return nil
}

func isLimited(userId int64, config *models.TaskConfig, rejectPolicy models.LimiterRejectPolicy) bool {
	if _, err := limiter.CheckLimit(config.TaskCode, models.LimitTypeTask, userId, 1, rejectPolicy); err != nil {
		log.Error(" isLimited CheckLimit error. %v", err)
		return true
	}
	return false

}

func isUserAvailable(userId int64) bool {
	if userId < 1 {
		return false
	}
	user, err := models.GetUserByID(userId)
	if err != nil || user == nil {
		return false
	}
	if user.IsOrganization() {
		return false
	}
	return true
}
