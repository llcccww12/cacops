package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

type ModelartsDeployQueue struct {
	JobID          string `xorm:"pk 'job_id'"`
	ModelID        string
	ModelName      string
	ServiceID      string
	DisplayJobName string
	CreateUnix     timeutil.TimeStamp
}

func CreateModelartsDeployQueue(deploy *ModelartsDeployQueue) (err error) {

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(deploy); err != nil {
		return err
	}
	return sess.Commit()
}

// write a function to get the deployment by job id
func GetModelartsDeployQueueByJobID(jobID string) (deploy *ModelartsDeployQueue, err error) {
	deploy = new(ModelartsDeployQueue)
	has, err := x.Where("job_id = ?", jobID).Get(deploy)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrModelartsDeployNotExist{jobID}
	}
	return deploy, nil
}

func GetModelartsDeployQueue(n int) ([]*ModelartsDeployQueue, error) {
	return getModelartsDeployQueue(x, n)
}
func getModelartsDeployQueue(e Engine, n int) ([]*ModelartsDeployQueue, error) {
	deploys := make([]*ModelartsDeployQueue, 0)
	return deploys, e.Asc("create_unix").Limit(n).Find(&deploys)
}

// delete data by jobid
func DeleteModelartsDeployQueueByJobID(jobID string) error {
	return deleteModelartsDeployQueueByJobID(x, jobID)
}

func deleteModelartsDeployQueueByJobID(e Engine, jobID string) error {
	_, err := e.Where("job_id = ?", jobID).Delete(&ModelartsDeployQueue{})
	log.Info("panguService: delte modelarts deploy queue by job id %s", jobID)
	return err
}
