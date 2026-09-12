package task

import (
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type IFLYTEKFinetuneTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &IFLYTEKFinetuneTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.IFLYTEKTraining,
			JobType:     models.JobTypeFINETUNE,
			Config:      IFLYTEKFinetuneTaskConfig,
		},
	}
	RegisterTask(models.JobTypeFINETUNE, entity.IFLYTEKTraining, t)
}

func IFLYTEKFinetuneTaskConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())
	var config = &entity.AITaskBaseConfig{
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerDataset: {
				ContainerPath:     "/dataset",
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{storageType},
				Uncompressed:      true,
			}},
		ActionType: models.ActionCreateIFLYTEKFinetuneTask,
	}
	return config
}

func (t IFLYTEKFinetuneTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.HandleReqParameters).
		Next(t.CheckWorkerNum).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckModels).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create IFLYTEKFinetuneTaskTemplate err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g IFLYTEKFinetuneTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CallCreationAPI.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	c := g.GetMyCluster()
	if c == nil {
		return response.SYSTEM_ERROR
	}
	form := ctx.Request
	imageUrl := strings.TrimSpace(form.ImageUrl)
	if form.ImageID != "" {
		imageUrl = ""
	}
	req := entity.CreateTrainTaskRequest{
		Name:           form.JobName,
		DisplayJobName: form.DisplayJobName,
		Tasks: []entity.TrainTask{
			{
				Name:             form.JobName,
				ResourceSpecId:   ctx.Spec.SourceSpecId,
				ImageId:          form.ImageID,
				ImageUrl:         imageUrl,
				Datasets:         ctx.GetContainerDataArray(entity.ContainerDataset),
				Code:             ctx.GetContainerDataArray(entity.ContainerCode),
				Queues:           ctx.Queues,
				PreTrainModel:    ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				BootFile:         form.BootFile,
				OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
				Params:           form.ParamArray,
				Spec:             ctx.Spec,
				RepoName:         ctx.Repository.Name,
				WorkServerNumber: ctx.Request.WorkServerNumber,
				Method:           "LoRA",
			},
		},
		TaskConfig: ctx.Config,
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateTrainJob(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateIFLYTEKFinetuneTask",ctx.TraceContext))
	if err != nil {
		log.Error("IFLYTEKFinetuneTaskTemplate CreateTrainJob err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	log.Info("CallCreationAPI success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g IFLYTEKFinetuneTaskTemplate) CheckModels(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckModels.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	modelName := ctx.Request.AppName
	if modelName == "" {
		return response.PARAM_ERROR
	}
	ctx.Request.ModelNames = modelName
	ctx.ContainerData = map[entity.ContainerDataType][]entity.ContainerData{
		entity.ContainerPreTrainModel: {entity.ContainerData{Name: modelName}},
	}
	log.Info("CheckModels success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}
