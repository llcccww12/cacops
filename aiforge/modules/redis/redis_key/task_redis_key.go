package redis_key

const TASK_REDIS_PREFIX = "task"

func TaskAccomplishLock(sourceId string, taskType string) string {
	return KeyJoin(TASK_REDIS_PREFIX, sourceId, taskType, "accomplish")
}

func TaskConfigList() string {
	return KeyJoin(TASK_REDIS_PREFIX, "config", "list")
}
func TaskConfigOperateLock(taskCode, rewardType string) string {
	return KeyJoin(TASK_REDIS_PREFIX, "config", "operate", "lock")
}
