package redis_key

import "fmt"

const ACCOUNT_REDIS_PREFIX = "account"

func PointAccountOperateLock(userId int64) string {
	return KeyJoin(ACCOUNT_REDIS_PREFIX, fmt.Sprint(userId), "point", "operate", "lock")
}

func PointAccountInfo(userId int64) string {
	return KeyJoin(ACCOUNT_REDIS_PREFIX, fmt.Sprint(userId), "info")
}

func PointAccountInitLock(userId int64) string {
	return KeyJoin(ACCOUNT_REDIS_PREFIX, fmt.Sprint(userId), "init", "lock")
}
