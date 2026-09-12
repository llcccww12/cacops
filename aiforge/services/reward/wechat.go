package reward

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/reward/point/account"
	"fmt"
	"strconv"
	"time"
)

func TryToNotifyInsufficientBalance(cloudbrain *models.Cloudbrain, unitPrice float64, nextExecuteTime timeutil.TimeStamp) {
	if cloudbrain.Spec == nil {
		models.LoadSpecs([]*models.Cloudbrain{cloudbrain})
		if cloudbrain.Spec == nil {
			log.Error("TryToNotifyInsufficientBalance cloudbrain.Spec == nil .cloudbrain.id=%d", cloudbrain.ID)
			models.LoadSpecs([]*models.Cloudbrain{cloudbrain})
		}
	}
	if unitPrice == 0 {
		log.Error("TryToNotifyInsufficientBalance unitPrice == 0 .cloudbrain.id=%d", cloudbrain.ID)
		return
	}
	account, err := account.GetAccount(cloudbrain.UserID)
	if err != nil {
		log.Error("TryToNotifyInsufficientBalance GetAccount error.cloudbrain.id=%d", cloudbrain.ID)
		return
	}
	if account.Balance >= unitPrice {
		return
	}
	if !cloudbrain.IsRunning() {
		return
	}
	if !isMsgNotSendBefore(account, "InsufficientBalance", cloudbrain.ID) {
		return
	}
	log.Info("start to notify CloudbrainTaskComingToFinished.id=%d", cloudbrain.ID)
	notification.NotifyCloudbrainTaskComingToFinished(cloudbrain, nextExecuteTime, account)
}

func isMsgNotSendBefore(account *models.PointAccount, msgType string, taskId int64) bool {
	updateUnix := account.UpdatedUnix
	k := redis_key.RewardMsgNewestTime(account.UserId, msgType, fmt.Sprint(taskId))
	s, err := redis_client.Get(k)
	if err != nil {
		log.Error("isMsgSendBefore get redis err.%v", err)
		return false
	}
	var t int64
	if s != "" {
		t, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Error("isMsgSendBefore ParseInt err.%v", err)
			return false
		}
	}

	if t >= int64(updateUnix) {
		return false
	}
	redis_client.Setex(k, fmt.Sprint(updateUnix), 1*time.Hour)
	return true
}
