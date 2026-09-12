package redis_key

import "fmt"

const OVERVIEW_REDIS_PREFIX = "overview"

// GetOverviewKeyByUserId 获取概览页面的key
func GetOverviewKeyByUserId(userId int64) string {
	return KeyJoin(OVERVIEW_REDIS_PREFIX, "overview", "userId", fmt.Sprint(userId))
}
