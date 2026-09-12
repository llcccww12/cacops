package convert

import (
	"code.gitea.io/gitea/models"
	api "code.gitea.io/gitea/modules/structs"
)

func ToFineTuneJobShow(cloudbrain *models.Cloudbrain) *api.FinetuneJobShow {

	jobID := cloudbrain.JobID
	deployStatus, _ := models.GetModelartsDeployStatusByJobID(jobID)

	return &api.FinetuneJobShow{
		ID:             cloudbrain.ID,
		JobID:          cloudbrain.JobID,
		JobType:        cloudbrain.JobType,
		DisplayJobName: cloudbrain.DisplayJobName,
		Status:         cloudbrain.Status,
		CreatedUnix:    int64(cloudbrain.CreatedUnix),
		JobCategory:    cloudbrain.FineTuneCategory,
		DeployStatus:   deployStatus,
		Cleared:        cloudbrain.Cleared,
	}
}
