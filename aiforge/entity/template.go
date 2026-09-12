package entity

import "code.gitea.io/gitea/models"

type SaveAITaskTemplateReq struct {
	AITaskId   int64                  `json:"AITaskId"`
	AITemplate *models.AITemplateText `json:"AITemplate"`
}
