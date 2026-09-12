package task

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"path"
	"strings"
)

func GetDetailConfigInfoByCloudbrain(cloudbrain *models.Cloudbrain) *entity.AITaskDetailConfigInfo {
	aiConfig := cloudbrain.GetCloudbrainConfig()
	if aiConfig != nil {
		return entity.BuildAITaskDetailConfigInfo(aiConfig)
	}
	return getHistoricalConfigInfo(cloudbrain)
}

//历史任务在cloudbrain_config表中没有对应的记录，因此根据实时配置模拟
func getHistoricalConfigInfo(cloudbrain *models.Cloudbrain) *entity.AITaskDetailConfigInfo {
	t, _ := GetAITaskTemplateFromCloudbrain(cloudbrain)
	if t == nil {
		return &entity.AITaskDetailConfigInfo{}
	}
	c := t.GetConfig(entity.AITaskConfigKey{ComputeSource: cloudbrain.GetStandardComputeSource()})
	return &entity.AITaskDetailConfigInfo{
		BaseConfig:         c,
		OutputObjectPrefix: GetContainerStorageObjectPrefix(c, cloudbrain.JobName, cloudbrain.VersionName, entity.ContainerOutPutPath),
		OutputStorageType:  GetContainerStorageType(c, entity.ContainerOutPutPath),
		LogObjectPrefix:    GetContainerStorageObjectPrefix(c, cloudbrain.JobName, cloudbrain.VersionName, entity.ContainerLogPath),
		LogStorageType:     GetContainerStorageType(c, entity.ContainerLogPath),
	}
}

func GetContainerStorageObjectPrefix(c *entity.AITaskBaseConfig, jobName string, versionName string, containerType entity.ContainerDataType) string {
	config := c.GetContainerConfig(containerType)
	if config == nil {
		return ""
	}
	st := config.AcceptStorageType
	if st == nil && len(st) == 0 {
		return ""
	}
	uploader := storage_helper.SelectStorageHelperFromStorageType(st[0])
	//兼容历史任务所以加上了versionName,另外云脑二训练任务为了适配modelarts接口加上了默认版本，此时要剔除
	localPath := config.GetLocalPath()
	localPath = strings.TrimSuffix(localPath, models.CloudbrainTwoDefaultVersion)
	objectKey := path.Join(uploader.GetJobDefaultObjectKeyPrefix(jobName), localPath, versionName)
	return objectKey
}

func GetContainerStorageType(c *entity.AITaskBaseConfig, containerType entity.ContainerDataType) entity.StorageType {
	outputConfig := c.GetContainerConfig(containerType)
	if outputConfig == nil {
		return ""
	}
	st := outputConfig.AcceptStorageType
	if st == nil && len(st) == 0 {
		return ""
	}
	return st[0]
}
