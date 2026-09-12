package redis_key

import (
	"code.gitea.io/gitea/modules/setting"
	"fmt"
	"strings"
)

const REWARD_REDIS_PREFIX = "reward"

func RewardOperateLock(requestId string, sourceType string, operateType string) string {
	return KeyJoin(REWARD_REDIS_PREFIX, requestId, sourceType, operateType, "send")
}

func RewardOperateNotification() string {
	return KeyJoin(REWARD_REDIS_PREFIX, "operate", strings.ReplaceAll(setting.AppURL, "/", ""), "notification")
}

func RewardTaskRunningLock(taskId int64) string {
	return KeyJoin(REWARD_REDIS_PREFIX, "periodic_task", fmt.Sprint(taskId), "lock")
}

func RewardMsgNewestTime(userId int64, msgType string, code string) string {
	return KeyJoin(REWARD_REDIS_PREFIX, "msg", fmt.Sprint(userId), msgType, code, "newest_time")
}
