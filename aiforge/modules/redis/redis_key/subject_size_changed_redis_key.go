package redis_key

const SUBJECT_SIZE_CHANGED_REDIS_PREFIX = "subject_size_changed"

func AimodelSizeChangedQueue() string {
	return KeyJoin(SUBJECT_SIZE_CHANGED_REDIS_PREFIX, "aimodel", "queue")
}

func DatasetSizeChangedQueue() string {
	return KeyJoin(SUBJECT_SIZE_CHANGED_REDIS_PREFIX, "dataset", "queue")
}
