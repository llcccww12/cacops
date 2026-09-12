package redis_key

import (
	"code.gitea.io/gitea/modules/timeutil"
	"time"
)

const RESOURCE_REDIS_PREFIX = "resource"

func EveryDayAICenterList(t time.Time) string {
	str := timeutil.GetDateString(t)
	return KeyJoin(RESOURCE_REDIS_PREFIX, str, "ai_center_list")
}

func UpdateEveryDayAICenterListLock(t time.Time) string {
	str := timeutil.GetDateString(t)
	return KeyJoin(RESOURCE_REDIS_PREFIX, str, "lock", "ai_center_list")
}

func EveryDayCardInfo(t time.Time) string {
	str := timeutil.GetDateString(t)
	return KeyJoin(RESOURCE_REDIS_PREFIX, str, "card_info")
}

func UpdateCardInfoLock(t time.Time) string {
	str := timeutil.GetDateString(t)
	return KeyJoin(RESOURCE_REDIS_PREFIX, str, "lock", "card_info")
}
