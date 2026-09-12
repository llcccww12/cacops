package redis_key

import "fmt"

const CLOUDBRAIN_PREFIX = "cloudbrain"

func CloudbrainBindingJobNameKey(jobType string, jobName string) string {
	return KeyJoin(CLOUDBRAIN_PREFIX, jobType, jobName, "redis_key")
}

func CloudbrainUniquenessKey(userId int64, jobType string) string {
	return KeyJoin(CLOUDBRAIN_PREFIX, fmt.Sprint(userId), jobType, "uniqueness")
}
