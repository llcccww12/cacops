package task

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
)

func GetAITaskTemplate(jobType models.JobType, clusterType entity.ClusterType) (AITaskTemplate, *response.BizError) {
	t, err := GetTask(jobType, clusterType)
	if err != nil || t == nil {
		return nil, response.PARAM_ERROR
	}
	return t, nil
}

func GetAITaskTemplateByCloudbrainId(id int64) (AITaskTemplate, *response.BizError) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		log.Error("GetAITaskByJobId GetCloudbrainByJobID err.%v", err)
		return nil, response.AI_TASK_NOT_EXISTS
	}
	return GetAITaskTemplateFromCloudbrain(cloudbrain)
}

func GetAITaskTemplateByJobId(jobId string) (AITaskTemplate, *response.BizError) {
	cloudbrain, err := models.GetCloudbrainByJobID(jobId)
	if err != nil {
		log.Error("GetAITaskByJobId GetCloudbrainByJobID err.%v", err)
		return nil, response.AI_TASK_NOT_EXISTS
	}
	return GetAITaskTemplateFromCloudbrain(cloudbrain)
}

func GetAITaskTemplateFromCloudbrain(cloudbrain *models.Cloudbrain) (AITaskTemplate, *response.BizError) {
	jobType := models.JobType(cloudbrain.JobType)
	cluster := entity.GetClusterTypeFromCloudbrainType(cloudbrain.Type)
	return GetAITaskTemplate(jobType, cluster)
}
