package task

import (
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
)

type CloudbrainOneNotebookTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &CloudbrainOneNotebookTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.OpenICloudbrainOne,
			JobType:     models.JobTypeDebug,
			Config:      GetCloudbrainOneNotebookConfig,
		},
	}
	RegisterTask(models.JobTypeDebug, entity.OpenICloudbrainOne, t)
}

func GetCloudbrainOneNotebookConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	//默认配置
	config := &entity.AITaskBaseConfig{
		ActionType:          models.ActionCreateDebugGPUTask,
		IsActionUseJobId:    false,
		DatasetsLimitSizeGB: setting.DebugAttachSize,
		DatasetsMaxNum:      setting.MaxDatasetNum,
		ModelLimitSizeGB:    setting.DEBUG_MODEL_SIZE_LIMIT_GB,
		ModelMaxNum:         setting.DEBUG_MODEL_NUM_LIMIT,
		DebugAddressCheck:   true,
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:     "/code",
				ReadOnly:          false,
				AcceptStorageType: []entity.StorageType{entity.MINIO},
				Uncompressed:      true,
			},
			entity.ContainerDataset: {
				ContainerPath:     "/dataset",
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.MINIO},
				Uncompressed:      true,
			},
			entity.ContainerPreTrainModel: {
				ContainerPath:     "/pretrainmodel",
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.MINIO},
			},
			entity.ContainerOutPutPath: {
				ContainerPath:       "/model",
				StorageRelativePath: cloudbrain.ModelMountPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.MINIO},
				MKDIR:               true,
			},
		},
	}
	//在线运行notebook配置
	if opts.IsFileNoteBookRequest {
		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerFileNoteBookCode: {},
				entity.ContainerCode: {
					ContainerPath:     "/code",
					ReadOnly:          false,
					AcceptStorageType: []entity.StorageType{entity.MINIO},
					Uncompressed:      true,
				},
			},
			DebugAddressCheck: true,
		}

	}
	return config
}

func (t CloudbrainOneNotebookTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.CheckNotebookCount).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckModels).
		Next(t.CheckBranchExists).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create CloudbrainOneNotebookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil

}

func (t CloudbrainOneNotebookTaskTemplate) Restart(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.BuildRequest4Restart).
		Next(t.CheckSourceTaskIsCleared).
		Next(t.CheckModels).
		Next(t.CheckDatasets).
		Next(t.CheckModels).
		Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.BuildContainerData).
		Next(t.GetAvailableQueues).
		Next(t.CallRestartAPI).
		Next(t.CreateCloudbrainRecord4Restart).
		Next(t.NotifyCreation).
		Operate(ctx)
	if err != nil {
		log.Error("Restart GrampusNoteBookTask err.%v", err)
		return nil, err
	}
	if err != nil {
		log.Error("Restart GrampusNoteBookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID, Status: ctx.NewCloudbrain.Status}, nil

}

func (g CloudbrainOneNotebookTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		return response.SYSTEM_ERROR
	}
	form := ctx.Request

	req := entity.CreateNoteBookTaskRequest{
		Name: form.JobName,
		Tasks: []entity.NoteBookTask{
			{
				Name:             form.JobName,
				ResourceSpecId:   ctx.Spec.SourceSpecId,
				ImageId:          form.ImageID,
				ImageUrl:         strings.TrimSpace(form.ImageUrl),
				Datasets:         ctx.GetContainerDataArray(entity.ContainerDataset),
				Code:             ctx.GetContainerDataArray(entity.ContainerCode),
				PreTrainModel:    ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
				AutoStopDuration: getAutoStopDurationMs(ctx.NewCloudbrain.TimeLimit, ctx.User.ID),
				Capacity:         setting.Capacity,
				Queues:           ctx.Queues,
				Spec:             ctx.Spec,
			},
		},
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateNoteBook(req,nil)
	if err != nil {
		log.Error("CloudbrainOneNotebookTask CreateNoteBook err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}

	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}

	return nil
}

func (g CloudbrainOneNotebookTaskTemplate) CallRestartAPI(ctx *context.CreationContext) *response.BizError {
	//云脑一没有再次调试接口，通过使用同样的参数新建接口来模拟
	return g.CallCreationAPI(ctx)
}
