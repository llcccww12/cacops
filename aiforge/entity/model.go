package entity

import "code.gitea.io/gitea/models"

type ModelBriefInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	OwnerName string `json:"owner_name"`
}

func BuildModelBriefInfo(dataset *models.AiModelManage) *ModelBriefInfo {
	return nil
}
