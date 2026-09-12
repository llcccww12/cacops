package redis_key

import "fmt"

const AI_MODEL_REDIS_PREFIX = "ai_model"

func AIModelMetaUpdateLock(modelId string) string {
	return KeyJoin(AI_MODEL_REDIS_PREFIX, modelId, "meta", "lock")
}

func AimodelNameLock(userId int64, aimodelName string) string {
	return KeyJoin(AI_MODEL_REDIS_PREFIX, fmt.Sprint(userId), aimodelName, "create_lock")
}

func AimodelVersionUpdateLock(aimodelId string) string {
	return KeyJoin(AI_MODEL_REDIS_PREFIX, aimodelId, "version", "lock")
}

func AimodelExportTaskResultProcessIdKey(aimodelId string, cloudbrainId int64) string {
	return KeyJoin(AI_MODEL_REDIS_PREFIX, aimodelId, "export_task", fmt.Sprint(cloudbrainId), "process_id")
}
