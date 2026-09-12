package redis_key

import (
	"code.gitea.io/gitea/models"
	"fmt"
)

const LIMIT_REDIS_PREFIX = "limit"

func LimitCount(userId int64, limitCode string, limitType string, scope string, period *models.PeriodResult) string {
	if scope == models.LimitScopeAllUsers.Name() {
		if period == nil {
			return KeyJoin(LIMIT_REDIS_PREFIX, limitCode, limitType, "count")
		}
		return KeyJoin(LIMIT_REDIS_PREFIX, limitCode, limitType, fmt.Sprint(period.StartTime.Unix()), fmt.Sprint(period.EndTime.Unix()), "count")
	}
	if period == nil {
		return KeyJoin(LIMIT_REDIS_PREFIX, "uid", fmt.Sprint(userId), limitCode, limitType, "count")
	}
	return KeyJoin(LIMIT_REDIS_PREFIX, "uid", fmt.Sprint(userId), limitCode, limitType, fmt.Sprint(period.StartTime.Unix()), fmt.Sprint(period.EndTime.Unix()), "count")

}

func LimitConfig(limitType string) string {
	return KeyJoin(LIMIT_REDIS_PREFIX, limitType, "config")
}
