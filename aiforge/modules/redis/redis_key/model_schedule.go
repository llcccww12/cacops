package redis_key

import "fmt"

const MODEL_SCHEDULE_PREFIX = "model_schedule"

func RecordHandleLock(jobId string) string {
	return KeyJoin(MODEL_SCHEDULE_PREFIX, fmt.Sprint(jobId), "handle")
}
