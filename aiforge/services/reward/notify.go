package reward

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/timeutil"
	"encoding/json"
	"fmt"
	"time"
)

func NotifyRewardOperation(userId int64, amount float64, sourceType models.SourceType, rewardType models.RewardType, operateType models.RewardOperateType) {
	switch sourceType {
	case models.SourceTypeRunCloudbrainTask:
		return
	}
	data := &models.UserRewardOperationRedis{
		UserId:      userId,
		Amount:      amount,
		RewardType:  rewardType,
		OperateType: operateType,
	}
	b, _ := json.Marshal(data)
	redis_client.ZAdd(redis_key.RewardOperateNotification(), string(b), float64(time.Now().Unix()))
}

func GetRewardOperation(since, until timeutil.TimeStamp) []models.UserRewardOperation {
	list, err := redis_client.ZRangeByScore(redis_key.RewardOperateNotification(), float64(since), float64(until))
	if err != nil {
		log.Error("GetRewardOperation ZRangeByScore error. %v", err)
		return nil
	}
	if len(list) == 0 {
		log.Debug("GetRewardOperation list length = 0")
		return nil
	}
	r := make([]models.UserRewardOperation, len(list))
	for _, v := range list {
		t := models.UserRewardOperationRedis{}
		json.Unmarshal([]byte(v), &t)
		r = append(r, models.UserRewardOperation{
			UserId: t.UserId,
			Msg:    v,
		})
	}
	redis_client.ZRemRangeByScore(redis_key.RewardOperateNotification(), float64(since), float64(until))
	return r
}

func GetRewardOperateMsg(u models.UserRewardOperationRedis) string {
	return u.OperateType.Show() + fmt.Sprint(u.Amount) + u.RewardType.Show()
}
