package task

import (
	"strings"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type GrampusGeneralTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusGeneralTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeGeneral,
			Config:      GetGeneralTaskConfig,
		},
	}
	RegisterTask(models.JobTypeGeneral, entity.C2Net, t)
}

func GetGeneralTaskConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/code"
	datasetPath := "/dataset"
	pretrainModelPath := "/pretrainmodel"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())

	config := &entity.AITaskBaseConfig{
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:     codePath,
				ReadOnly:          false,
				AcceptStorageType: []entity.StorageType{storageType},
				VolumeFolder:      true,
			},
			entity.ContainerDataset: {
				ContainerPath:     datasetPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{storageType},
			},
			entity.ContainerPreTrainModel: {
				ContainerPath:     pretrainModelPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{storageType},
			},
		},
		ActionType:          models.ActionCreateGeneralGPUTask,
		DatasetsLimitSizeGB: setting.DebugAttachSize,
		DatasetsMaxNum:      setting.MaxDatasetNum,
		ModelLimitSizeGB:    setting.DEBUG_MODEL_SIZE_LIMIT_GB,
		ModelMaxNum:         setting.DEBUG_MODEL_NUM_LIMIT,
		DebugAddressCheck:   true,
	}
	return config
}

func (t GrampusGeneralTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		//Next(t.CheckNotebookCount).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckBranchExists).
		Next(t.CheckModels).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create GrampusNoteBookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g GrampusGeneralTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return response.SYSTEM_ERROR
	}
	form := ctx.Request
	imageUrl := strings.TrimSpace(form.ImageUrl)
	if form.ImageID != "" {
		imageUrl = ""
	}
	req := entity.CreateGeneralTaskRequest{
		Name: form.JobName,
		Tasks: []entity.GeneralTask{
			{
				Name:           form.JobName,
				ResourceSpecId: ctx.Spec.SourceSpecId,
				ImageId:        form.ImageID,
				ImageUrl:       imageUrl,
				Datasets:       ctx.GetContainerDataArray(entity.ContainerDataset),
				PreTrainModel:  ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				Code:           ctx.GetContainerDataArray(entity.ContainerCode),
				EnvVariables:   map[string]interface{}{},
				Capacity:       setting.Capacity,
				Queues:         ctx.Queues,
				Spec:           ctx.Spec,
			},
		},
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateGeneralTask(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateGeneralTask",ctx.TraceContext))
	if err != nil {
		log.Error("GrampusGeneralTaskTemplate CreateGeneralTask err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	if res.JobID == "" {
		log.Error("GrampusGeneralTaskTemplate CreateGeneralTask failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.CREATE_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}

func (g GrampusGeneralTaskTemplate) GetImages(computeSource models.ComputeSource, accCardType string, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return nil, false, response.SYSTEM_ERROR
	}

	var images []entity.ClusterImage
	var customFlag bool
	var err error
	images, customFlag, err = c.GetNotebookImages(entity.GetImageReq{
		ComputeSource: computeSource,
		JobType:       models.JobTypeDebug,
		AccCardType:   accCardType,
	}, queues...)
	if err != nil {
		log.Error("GetImages err.computeSource=%s err =%v", computeSource.Name, err)
		return nil, false, response.NewBizError(err)
	}
	return images, customFlag, nil
}

func (g GrampusGeneralTaskTemplate) GetOperationProfile(cloudbrainId int64) (*entity.OperationProfile, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return nil, response.SYSTEM_ERROR
	}
	s, err := GetOperationProfile(cloudbrainId, c.GetNoteBookOperationProfile)
	if err != nil {
		log.Error("GetOperationProfile err.cloudbrainId=%d err =%v", cloudbrainId, err)
		return nil, nil
	}
	return s, nil
}
