package task

import (
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
)

type CloudbrainOneTrainTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &CloudbrainOneTrainTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.OpenICloudbrainOne,
			JobType:     models.JobTypeTrain,
			Config:      GetCloudbrainOneTrainConfig,
		},
	}
	RegisterTask(models.JobTypeTrain, entity.OpenICloudbrainOne, t)
}

func GetCloudbrainOneTrainConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	config := &entity.AITaskBaseConfig{
		ActionType:       models.ActionCreateGPUTrainTask,
		IsActionUseJobId: true,
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
			},
			entity.ContainerLogPath: {
				ContainerPath:       "/model",
				StorageRelativePath: cloudbrain.ModelMountPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.MINIO},
				MKDIR:               true,
			},
		}}

	return config
}

func (t CloudbrainOneTrainTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.HandleReqParameters).
		Next(t.CheckPrivilege4Continue).
		Next(t.CheckSourceTaskIsCleared).
		Next(t.CheckWorkerNum).
		Next(t.CheckMultiRequest).
		Next(t.CheckBranchExists).
		Next(t.CheckBootFile).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
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

func (g CloudbrainOneTrainTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		return response.SYSTEM_ERROR
	}
	form := ctx.Request
	req := entity.CreateTrainTaskRequest{
		Name:           form.JobName,
		DisplayJobName: form.DisplayJobName,
		Tasks: []entity.TrainTask{
			{
				Name:           form.JobName,
				ResourceSpecId: ctx.Spec.SourceSpecId,
				ImageId:        form.ImageID,
				ImageUrl:       strings.TrimSpace(form.ImageUrl),
				Datasets:       ctx.GetContainerDataArray(entity.ContainerDataset),
				Code:           ctx.GetContainerDataArray(entity.ContainerCode),
				Queues:         ctx.Queues,
				PreTrainModel:  ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				BootFile:       form.BootFile,
				OutPut:         ctx.GetContainerDataArray(entity.ContainerOutPutPath),
				LogPath:        ctx.GetContainerDataArray(entity.ContainerLogPath),
				Params:         form.ParamArray,
				Spec:           ctx.Spec,
			},
		},
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateTrainJob(req, nil)
	if err != nil {
		log.Error("GrampusNoteBookTask CreateNoteBook err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}
