package context

import (
	"context"
	"encoding/json"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"

	"code.gitea.io/gitea/modules/git"
)

type CreationContext struct {
	Request          *entity.CreateReq
	ContainerData    map[entity.ContainerDataType][]entity.ContainerData
	GitRepo          *git.Repository
	Repository       *models.Repository
	Spec             *models.Specification
	User             *models.User
	CommitID         string
	Response         *entity.CreationResponse
	SourceCloudbrain *models.Cloudbrain
	NewCloudbrain    *models.Cloudbrain
	Config           *entity.AITaskBaseConfig
	Queues           []models.ResourceQueue
	TraceContext     *context.Context
}

func (ctx *CreationContext) AddContainerData(t entity.ContainerDataType, d []entity.ContainerData) {
	if ctx.ContainerData == nil {
		ctx.ContainerData = make(map[entity.ContainerDataType][]entity.ContainerData, 0)
	}
	ctx.ContainerData[t] = d
}
func (ctx *CreationContext) GetContainerDataArray(t entity.ContainerDataType) []entity.ContainerData {
	if ctx.ContainerData == nil {
		return nil
	}
	return ctx.ContainerData[t]
}
func (ctx *CreationContext) GetContainerData(t entity.ContainerDataType) entity.ContainerData {
	a := ctx.GetContainerDataArray(t)
	if a == nil || len(a) == 0 {
		return entity.ContainerData{}
	}
	return a[0]
}
func (ctx *CreationContext) WriteResponse(t entity.ContainerDataType) entity.ContainerData {
	a := ctx.GetContainerDataArray(t)
	if a == nil || len(a) == 0 {
		return entity.ContainerData{}
	}
	return a[0]
}

func (ctx *CreationContext) BuildCloudbrainConfig() *models.CloudbrainConfig {
	var aiConfigStr = ""
	s, err := json.Marshal(ctx.Config)
	if err == nil {
		aiConfigStr = string(s)
	}
	var containerDataStr = ""
	t, err := json.Marshal(ctx.ContainerData)
	if err == nil {
		containerDataStr = string(t)
	}
	output := ctx.GetContainerData(entity.ContainerOutPutPath)
	log := ctx.GetContainerData(entity.ContainerLogPath)
	c := &models.CloudbrainConfig{

		ConfigurationSnapshot: aiConfigStr,
		OutputBucket:          output.Bucket,
		OutputObjectPrefix:    output.ObjectKey,
		OutputStorageType:     string(output.StorageType),
		OutputEndpoint:        output.EndPoint,
		LogBucket:             log.Bucket,
		LogObjectPrefix:       log.ObjectKey,
		LogStorageType:        string(log.StorageType),
		LogEndpoint:           log.EndPoint,
		ContainerDataSnapshot: containerDataStr,
	}
	return c
}
