package redis_key

const COPY_LOAD_MODEL_KEY_PREFIX = "npu_model_experience_load"
const COPY_LOAD_MODEL_FINISHED_KEY_PREFIX = "npu_model_experience_loaded"

const EVAL_MODEL_KEY_PREFIX = "model_eval"

func ModelExperienceModelLoadingKey(jobId string) string {
	return KeyJoin(COPY_LOAD_MODEL_KEY_PREFIX, jobId)
}

func ModelExperienceModelLoadedKey(jobId string) string {
	return KeyJoin(COPY_LOAD_MODEL_FINISHED_KEY_PREFIX, jobId)
}

func ModelEvalKey(jobId string) string {
	return KeyJoin(EVAL_MODEL_KEY_PREFIX, jobId)
}
