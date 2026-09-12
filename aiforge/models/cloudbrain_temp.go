package models

import (
	"time"

	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	TempJobId     = "TEMP"
	TempVersionId = TempJobId
	TempJobStatus = TempJobId
)

type CloudbrainTemp struct {
	ID          int64              `xorm:"pk autoincr"`
	JobID       string             `xorm:"NOT NULL DEFAULT 'TEMP'"`
	VersionID   string             `xorm:"NOT NULL DEFAULT 'TEMP'"`
	JobName     string             `xorm:"NOT NULL "`
	Type        int                `xorm:"NOT NULL "`
	JobType     string             `xorm:"INDEX NOT NULL DEFAULT 'DEBUG'"`
	Status      string             `xorm:"INDEX NOT NULL DEFAULT 'TEMP'"`
	QueryTimes  int                `xorm:"INDEX NOT NULL DEFAULT 0"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated"`
	DeletedAt   time.Time          `xorm:"deleted"`
}

func InsertCloudbrainTemp(temp *CloudbrainTemp) (err error) {
	if _, err = x.Insert(temp); err != nil {
		return err
	}

	return nil
}

func getCloudBrainTemp(temp *CloudbrainTemp) (*CloudbrainTemp, error) {
	has, err := x.Get(temp)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrJobNotExist{}
	}
	return temp, nil
}

func GetCloudBrainTempJobs() ([]*CloudbrainTemp, error) {
	jobs := make([]*CloudbrainTemp, 0, 10)
	return jobs, x.In("status", TempJobStatus, string(ModelArtsStopping), string(ModelArtsTrainJobKilling)).
		And("query_times < ?", setting.MaxTempQueryTimes).
		Limit(100).
		Find(&jobs)
}

func DeleteCloudbrainTemp(temp *CloudbrainTemp) error {
	return deleteCloudbrainTemp(x, temp)
}

func deleteCloudbrainTemp(e Engine, temp *CloudbrainTemp) error {
	_, err := e.ID(temp.ID).Delete(temp)
	return err
}

func UpdateCloudbrainTemp(temp *CloudbrainTemp) error {
	_, err := x.ID(temp.ID).AllCols().Update(temp)
	return err
}
