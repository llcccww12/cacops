package models

import (
	"code.gitea.io/gitea/modules/timeutil"
)

type CloudbrainConfig struct {
	CloudbrainID          int64 `xorm:"pk"`
	OutputObjectPrefix    string
	OutputStorageType     string
	OutputBucket          string
	OutputEndpoint        string
	LogObjectPrefix       string
	LogStorageType        string
	LogBucket             string
	LogEndpoint           string
	ConfigurationSnapshot string             `xorm:"text"`
	ContainerDataSnapshot string             `xorm:"text"`
	CreatedTime           timeutil.TimeStamp `xorm:"created"`
	UpdatedTime           timeutil.TimeStamp `xorm:"updated"`
}

func GetCloudbrainConfig(cloudbrainId int64) (*CloudbrainConfig, error) {
	r := &CloudbrainConfig{}
	if has, err := x.Where("cloudbrain_id = ?", cloudbrainId).Get(r); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

func InsertCloudbrainConfig(config *CloudbrainConfig) (int64, error) {
	return x.Insert(config)
}

func UpdateCloudbrainConfigByTaskID(taskID int64, config *CloudbrainConfig) (int64, error) {
	return x.Where("cloudbrain_id = ? ", taskID).Update(config)
}
