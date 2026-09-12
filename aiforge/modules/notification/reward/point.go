package reward

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification/base"
)

type pointNotifier struct {
	base.NullNotifier
}

var (
	_ base.Notifier = &pointNotifier{}
)

// NewNotifier create a new wechatNotifier notifier
func NewNotifier() base.Notifier {
	return &pointNotifier{}
}

func (*pointNotifier) NotifyChangeCloudbrainStatus(cloudbrain *models.Cloudbrain, oldStatus string) {
	log.Info("pointNotifier NotifyChangeCloudbrainStatus cloudbrain.id=%d cloudbrain.status=%s oldStatus=%s", cloudbrain.ID, cloudbrain.Status, oldStatus)
	if cloudbrain.IsRunning() || cloudbrain.IsTerminal() {
		models.StatusChangeChan <- cloudbrain
	}
}
