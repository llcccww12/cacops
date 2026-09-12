package models

import (
	"errors"
	"time"

	"code.gitea.io/gitea/modules/log"
	"xorm.io/builder"

	"code.gitea.io/gitea/modules/timeutil"
)

type GrampusMigrateResponse int

const (
	GrampusMigrateResponseMigrateInit   GrampusMigrateResponse = 0
	GrampusMigrateResponseSuccess       GrampusMigrateResponse = 1
	GrampusMigrateResponseFailed        GrampusMigrateResponse = 2
	GrampusMigrateResponseMigrating     GrampusMigrateResponse = 3
	GrampusMigrateResponseNoNeedMigrate GrampusMigrateResponse = 4
)

func (r GrampusMigrateResponse) ConvertToModelMigrateStep() ModelMigrateStep {
	switch r {
	case GrampusMigrateResponseMigrateInit:
		return GrampusMigrateInit
	case GrampusMigrateResponseSuccess:
		return GrampusMigrateSuccess
	case GrampusMigrateResponseFailed:
		return GrampusMigrateFailed
	case GrampusMigrateResponseMigrating:
		return GrampusMigrating
	case GrampusMigrateResponseNoNeedMigrate:
		return GrampusMigrateNoNeed
	}
	return -1
}

type ModelMigrateStep int

const (
	GrampusMigrateInit    ModelMigrateStep = 0
	GrampusMigrating      ModelMigrateStep = 1
	GrampusMigrateSuccess ModelMigrateStep = 2
	GrampusMigrateFailed  ModelMigrateStep = 3
	GrampusMigrateNoNeed  ModelMigrateStep = 4
	BucketMoving          ModelMigrateStep = 10
	MigrateFinished       ModelMigrateStep = 11
	BucketMoveFailed      ModelMigrateStep = 12
)

func (m ModelMigrateStep) GetStatus() ModelMigrateStatus {
	switch m {
	case MigrateFinished, GrampusMigrateNoNeed:
		return ModelMigrateSuccess
	case GrampusMigrateFailed, BucketMoveFailed:
		return ModelMigrateFailed
	case GrampusMigrateInit:
		return ModelMigrateWaiting
	case GrampusMigrateSuccess, GrampusMigrating, BucketMoving:
		return ModelMigrating
	}
	return -1
}

type ModelMigrateStatus int

const (
	ModelMigrateSuccess ModelMigrateStatus = 0
	ModelMigrating      ModelMigrateStatus = 1
	ModelMigrateFailed  ModelMigrateStatus = 2
	ModelMigrateWaiting ModelMigrateStatus = 3
)

var UnFinishedMigrateSteps = []ModelMigrateStep{GrampusMigrateInit, GrampusMigrating, GrampusMigrateSuccess, BucketMoving}

type ModelMigrateRecord struct {
	ID                   int64 `xorm:"pk autoincr"`
	CloudbrainID         int64 `xorm:"INDEX NOT NULL unique"`
	DestBucket           string
	DestEndpoint         string
	DestObjectKey        string
	DestProxy            string
	SrcBucket            string
	SrcEndpoint          string
	SrcObjectKey         string
	Status               ModelMigrateStatus `xorm:"NOT NULL DEFAULT 3"`
	CurrentStep          ModelMigrateStep   `xorm:"NOT NULL DEFAULT 0"`
	RetryCount           int
	CreatedUnix          timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix          timeutil.TimeStamp `xorm:"updated"`
	DeletedAt            time.Time          `xorm:"deleted"`
	Remark               string             `xorm:"TEXT"`
	DataID               string
	GrampusMigrateTaskId int64
	MigrateService       string
}

func (r *ModelMigrateRecord) IsFinished() bool {
	for _, s := range UnFinishedMigrateSteps {
		if s == r.CurrentStep {
			return false
		}
	}
	return true
}

func updateModelMigrateRecordCols(e Engine, record *ModelMigrateRecord, cols ...string) error {
	_, err := e.ID(record.ID).Cols(cols...).Update(record)
	return err
}

func UpdateModelMigrateRecordCols(record *ModelMigrateRecord, cols ...string) error {
	return updateModelMigrateRecordCols(x, record, cols...)
}
func IncreaseModelMigrateRetryCount(recordId int64) error {
	_, err := x.ID(recordId).Incr("retry_count", 1).Update(&ModelMigrateRecord{})
	return err
}

func UpdateModelMigrateStatusByStep(record *ModelMigrateRecord, newStep ModelMigrateStep) error {
	status := newStep.GetStatus()
	if status < 0 {
		log.Error("Step format error.id = %d,newStep = %d", record.ID, newStep)
		return errors.New("Step format error")
	}
	record.Status = status
	record.CurrentStep = newStep
	//正常情况下状态只能向更大的状态更新
	n, err := x.Where(builder.NewCond().And(builder.Eq{"id": record.ID}).
		And(builder.Lt{"current_step": newStep})).
		Cols("status", "current_step").
		Update(record)
	if err != nil {
		log.Error("UpdateModelMigrateStatusByStep err.%v", err)
		return err
	}

	if n == 0 {
		log.Error("UpdateModelMigrateStatusByStep total num is 0.r.ID=%d", record.ID)
		return errors.New("current_step not valid")
	}

	return nil
}

func RollBackMigrateStatus(record *ModelMigrateRecord, newStep ModelMigrateStep) error {
	status := newStep.GetStatus()
	if status < 0 {
		log.Error("Step format error.id = %d,newStep = %d", record.ID, newStep)
		return errors.New("Step format error")
	}
	record.Status = status
	record.CurrentStep = newStep
	_, err := x.ID(record.ID).
		Cols("status", "current_step").
		Update(record)
	if err != nil {
		log.Error("RollBackMigrateStatus err.%v", err)
		return err
	}

	return nil
}

func UpdateModelMigrateRecordByStep(record *ModelMigrateRecord) error {
	n, err := x.
		Where(builder.NewCond().And(builder.Eq{"id": record.ID})).
		Update(record)
	if err != nil {
		log.Error("UpdateModelMigrateRecordByStep err. ID=%d  err=%v", record.ID, err)
		return err
	}
	if n == 0 {
		log.Error("UpdateModelMigrateRecordByStep total num is 0.r.ID=%d", record.ID)
		return errors.New("current_step not valid")
	}
	return nil
}

func GetUnfinishedModelMigrateRecordsPaging(pageSize, page int) ([]*ModelMigrateRecord, error) {
	records := make([]*ModelMigrateRecord, 0, 10)
	return records, x.Cols("model_migrate_record.id", "model_migrate_record.cloudbrain_id", "model_migrate_record.dest_bucket", "model_migrate_record.dest_endpoint", "model_migrate_record.dest_object_key", "model_migrate_record.dest_proxy", "model_migrate_record.src_bucket", "model_migrate_record.src_endpoint", "model_migrate_record.src_object_key", "model_migrate_record.status", "model_migrate_record.current_step", "model_migrate_record.retry_count", "model_migrate_record.created_unix", "model_migrate_record.updated_unix", "model_migrate_record.deleted_at", "model_migrate_record.remark").Table("model_migrate_record").
		Join("inner", "cloudbrain", "cloudbrain.id  = model_migrate_record.cloudbrain_id").
		Where(builder.NewCond().And(
			builder.In("model_migrate_record.current_step", UnFinishedMigrateSteps)).
			And(builder.Eq{"cloudbrain.deleted_at": "0001-01-01 00:00:00"}.Or(builder.IsNull{"cloudbrain.deleted_at"})).
			And(builder.Eq{"cloudbrain.cleared": false}.Or(builder.IsNull{"cloudbrain.cleared"})).
			And(builder.Gt{"cloudbrain.created_unix": time.Now().Unix() - int64(24*30*time.Hour.Seconds())})).
		OrderBy("cloudbrain.id").
		Limit(pageSize, (page-1)*pageSize).
		Find(&records)
}

func InsertModelMigrateRecord(record *ModelMigrateRecord) (_ *ModelMigrateRecord, err error) {

	if _, err := x.Insert(record); err != nil {
		return nil, err
	}

	return record, nil
}

func GetModelMigrateRecordByCloudbrainId(cloudbrainId int64) (*ModelMigrateRecord, error) {
	r := &ModelMigrateRecord{}
	if has, err := x.Where("cloudbrain_id = ?", cloudbrainId).Get(r); err != nil {
		log.Error("GetModelMigrateRecordByCloudbrainId err. %v", err)
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil

}
func GetModelMigrateRecordById(id int64) (*ModelMigrateRecord, error) {
	r := &ModelMigrateRecord{}
	if has, err := x.ID(id).Get(r); err != nil {
		log.Error("GetModelMigrateRecordByCloudbrainId err. %v", err)
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil

}
