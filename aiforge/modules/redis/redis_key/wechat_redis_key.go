package redis_key

import (
	"code.gitea.io/gitea/models"
	"fmt"
	"time"
)

const PREFIX = "wechat"

func WechatBindingUserIdKey(sceneStr string) string {
	return KeyJoin(PREFIX, sceneStr, "scene_userId")
}

func WechatAccessTokenKey() string {
	return KeyJoin(PREFIX, "access_token")
}
func AccessTokenLockKey() string {
	return KeyJoin(PREFIX, "access_token_lock")
}

func LongRunningNotificationKey(cloudbrain *models.Cloudbrain, duration time.Duration) string {
	return KeyJoin(PREFIX, fmt.Sprint(cloudbrain.ID), duration.String(), "long_running_notification")
}

func DebugEndNotificationKey(cloudbrain *models.Cloudbrain) string {
	return KeyJoin(PREFIX, fmt.Sprint(cloudbrain.ID), "_running_notification")
}
