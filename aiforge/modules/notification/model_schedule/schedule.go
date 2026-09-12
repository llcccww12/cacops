package model_schedule

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification/base"
)

type scheduleNotifier struct {
	base.NullNotifier
}

var (
	_ base.Notifier = &scheduleNotifier{}
)

// NewNotifier create a new wechatNotifier notifier
func NewNotifier() base.Notifier {
	return &scheduleNotifier{}
}

func (*scheduleNotifier) NotifyChangeCloudbrainStatus(cloudbrain *models.Cloudbrain, oldStatus string) {
	if !cloudbrain.IsTerminal() {
		return
	}
	log.Info("try to InsertModelMigrateRecord.cloudbrainId=%d oldStatus=%s newStatus=%d", cloudbrain.ID, oldStatus, cloudbrain.Status)
	switch cloudbrain.Type {
	case models.TypeC2Net:
		_, err := models.InsertModelMigrateRecord(&models.ModelMigrateRecord{
			CloudbrainID: cloudbrain.ID,
			Status:       models.ModelMigrating,
			CurrentStep:  models.GrampusMigrating,
		})
		if err != nil {
			log.Error("InsertModelMigrateRecord err.cloudbrain.id=%d  err=%v", cloudbrain.ID, err)
		}
	}
}
