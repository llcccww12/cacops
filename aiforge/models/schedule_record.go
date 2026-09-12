package models

import (
	"time"

	"code.gitea.io/gitea/modules/timeutil"
)

const (
	StorageUrchinScheduleSucceed int = iota
	StorageUrchinScheduleProcessing
	StorageUrchinScheduleFailed
	StorageUrchinNoFile
	StorageUrchinScheduleWaiting
)

const (
	MoveBucketSucceed int = iota
	MoveBucketOperating
	MoveBucketFailed
	MoveBucketWaiting
)

type ModelScheduleStatus int

const (
	ModelScheduleSucceed ModelScheduleStatus = iota
	ModelScheduleOperating
	ModelScheduleFailed
	ModelScheduleWaiting
)

const UrchinDefaultBucket = "urchincache"

type ScheduleRecord struct {
	ID                 int64  `xorm:"pk autoincr"`
	CloudbrainID       int64  `xorm:"INDEX NOT NULL unique"`
	EndPoint           string `xorm:"INDEX NOT NULL"`
	Bucket             string `xorm:"INDEX NOT NULL"`
	ObjectKey          string `xorm:"INDEX NOT NULL"`
	ProxyServer        string `xorm:"INDEX NOT NULL"`
	IsDir              bool   `xorm:"default false"`
	ComputeSource      string
	TargetObjectKey    string
	Status             int                `xorm:"INDEX NOT NULL DEFAULT 0"`
	CreatedUnix        timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix        timeutil.TimeStamp `xorm:"updated"`
	DeletedAt          time.Time          `xorm:"deleted"`
	Remark             string
	LocalOperateStatus int `xorm:"INDEX NOT NULL DEFAULT 0"`
}

func updateScheduleCols(e Engine, record *ScheduleRecord, cols ...string) error {
	_, err := e.ID(record.ID).Cols(cols...).Update(record)
	return err
}

func UpdateScheduleCols(record *ScheduleRecord, cols ...string) error {
	return updateScheduleCols(x, record, cols...)
}

func UpdateScheduleLocalOperateStatus(record *ScheduleRecord, newLocalOperateStatus int) error {
	record.LocalOperateStatus = newLocalOperateStatus
	return updateScheduleCols(x, record, "local_operate_status")
}

func GetSchedulingRecord() ([]*ScheduleRecord, error) {
	records := make([]*ScheduleRecord, 0, 10)
	return records, x.
		Where("status = ?", StorageUrchinScheduleProcessing).
		Limit(100).
		Find(&records)
}

func InsertScheduleRecord(record *ScheduleRecord) (_ *ScheduleRecord, err error) {

	if _, err := x.Insert(record); err != nil {
		return nil, err
	}

	return record, nil
}

func getScheduleRecordByPrID(e Engine, cloudbrainId int64) (*ScheduleRecord, error) {
	record := new(ScheduleRecord)
	has, err := e.Where("cloudbrain_id = ?", cloudbrainId).Get(record)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return record, nil
}

func GetScheduleRecordByCloudbrainID(cloudbrainId int64) (*ScheduleRecord, error) {
	return getScheduleRecordByPrID(x, cloudbrainId)
}
