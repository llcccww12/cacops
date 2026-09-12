package dataset

import (
	"strings"

	"code.gitea.io/gitea/models"
)

func GetResourceType(cloudbrainType int) string {
	if cloudbrainType == 0 {
		return "CPU/GPU"
	} else {
		return "NPU"
	}
}

func GetStatusText(isPrivate bool) string {
	if isPrivate {
		return "dataset.private"
	} else {
		return "dataset.public"
	}
}

func IsShowDataSetOfCurrentRepo(repoID int64) bool {
	repo := models.Repository{
		ID: repoID,
	}

	dataset, _ := models.GetDatasetByRepo(&repo)
	if dataset != nil {
		return true
	}
	if models.HasReferenceDataset(repoID) {
		return false
	}
	return true

}

func GetFilterDeletedAttachments(uuids string) (string, string) {
	attachments, err := models.GetAttachmentsByUUIDs(strings.Split(uuids, ";"))
	if err != nil {
		return "", ""
	}
	uuidR := ""
	filenames := ""
	for i, attachment := range attachments {
		if i == 0 {
			uuidR += attachment.UUID
			filenames += attachment.Name
		} else {
			uuidR += ";" + attachment.UUID
			filenames += ";" + attachment.Name
		}
	}
	return uuidR, filenames

}
