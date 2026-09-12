package attachment

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func ClearTemporaryAttachment() {
	attachmets, err := models.GetTemporaryAttachments()
	if err != nil {
		log.Error("GetTemporaryAttachments: %v", err)
		return
	}

	for _, attach := range attachmets {
		if err := models.DeleteAttachment(attach, true); err != nil {
			log.Error("DeleteAttachment: %v", err)
		}
	}

}
