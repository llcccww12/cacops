package redis_key

import (
	"fmt"
)

const DATASET_REDIS_PREFIX = "dataset"

func DatasetNameLock(userId int64, datasetName string) string {
	return KeyJoin(DATASET_REDIS_PREFIX, fmt.Sprint(userId), datasetName, "create_lock")
}

func DatasetVersionUpdateLock(datasetId string) string {
	return KeyJoin(DATASET_REDIS_PREFIX, datasetId, "version", "lock")
}

func ExportTaskResultProcessIdKey(datasetId string, cloudbrainId int64) string {
	return KeyJoin(DATASET_REDIS_PREFIX, datasetId, "export_task", fmt.Sprint(cloudbrainId), "process_id")
}

func HandleOldDatasetLock(datasetId string) string {
	return KeyJoin(DATASET_REDIS_PREFIX, datasetId, "handle_old")
}

func ProcessingOldDataset() string {
	return KeyJoin(DATASET_REDIS_PREFIX, "proccessing", "old")
}
