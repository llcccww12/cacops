package redis_key

import "fmt"

const AI_TASK_PREFIX = "ai_task"

func AITaskUpdateLock(cloudbrainId int64) string {
	return KeyJoin(PREFIX, "update", fmt.Sprint(cloudbrainId), "lock")
}
