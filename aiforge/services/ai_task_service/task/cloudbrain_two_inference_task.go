package task

import (
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
)

type CloudbrainTwoInferenceTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &CloudbrainTwoInferenceTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.OpenICloudbrainTwo,
			JobType:     models.JobTypeInference,
			Config:      GetCloudbrainTwoInferenceConfig,
		},
	}
	RegisterTask(models.JobTypeInference, entity.OpenICloudbrainTwo, t)
}

func GetCloudbrainTwoInferenceConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	var config = &entity.AITaskBaseConfig{
		ActionType:       models.ActionCreateInferenceTask,
		IsActionUseJobId: true,
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:     "/code",
				ReadOnly:          false,
				AcceptStorageType: []entity.StorageType{entity.OBS},
				Uncompressed:      true,
			},
			entity.ContainerDataset: {
				ContainerPath:     "/dataset",
				ReadOnly:          true,
				Uncompressed:      true,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
			entity.ContainerPreTrainModel: {
				ContainerPath:     "/pretrainmodel",
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
			entity.ContainerOutPutPath: {
				ContainerPath:       "/output",
				StorageRelativePath: "/output" + models.CloudbrainTwoDefaultVersion,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.OBS},
				MKDIR:               true,
			},
			entity.ContainerLogPath: {
				ContainerPath:       "/log",
				StorageRelativePath: "/log" + models.CloudbrainTwoDefaultVersion,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.OBS},
				MKDIR:               true,
			},
		},
	}
	return config
}

func (t CloudbrainTwoInferenceTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
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
		log.Error("create GrampusInferenceTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g CloudbrainTwoInferenceTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		return response.SYSTEM_ERROR
	}
	form := ctx.Request
	req := entity.CreateTrainTaskRequest{
		Name:           form.JobName,
		DisplayJobName: form.DisplayJobName,
		Description:    form.Description,
		Tasks: []entity.TrainTask{
			{
				Name:             form.JobName,
				ResourceSpecId:   ctx.Spec.SourceSpecId,
				ImageId:          form.ImageID,
				ImageUrl:         strings.TrimSpace(form.ImageUrl),
				Datasets:         ctx.GetContainerDataArray(entity.ContainerDataset),
				Code:             ctx.GetContainerDataArray(entity.ContainerCode),
				LogPath:          ctx.GetContainerDataArray(entity.ContainerLogPath),
				Queues:           ctx.Queues,
				PreTrainModel:    ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				BootFile:         form.BootFile,
				OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
				Params:           form.ParamArray,
				Spec:             ctx.Spec,
				PoolId:           ctx.Spec.QueueCode,
				WorkServerNumber: form.WorkServerNumber,
			},
		},
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateTrainJob(req, nil)
	if err != nil {
		log.Error("CloudbrainTwo InfereceTask Create err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	ctx.Response = &entity.CreationResponse{
		JobID:       res.JobID,
		Status:      res.Status,
		CreateTime:  createTime,
		VersionID:   res.VersionID,
		VersionName: res.VersionName,
	}
	return nil
}
