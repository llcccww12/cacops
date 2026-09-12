package task

import (
	"strings"

	"encoding/json"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type GrampusFinetuneTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusFinetuneTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeFINETUNE,
			Config:      GetGrampusFinetuneTaskConfig,
		},
	}
	RegisterTask(models.JobTypeFINETUNE, entity.C2Net, t)
}

func GetGrampusFinetuneTaskConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/code"
	datasetPath := "/dataset"
	pretrainModelPath := "/pretrainmodel"
	outputPath := "/output"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())
	var config = &entity.AITaskBaseConfig{
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:       "/tmp" + codePath,
				StorageRelativePath: codePath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{storageType},
			},
			entity.ContainerDataset: {
				ContainerPath:     "/tmp" + datasetPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{storageType},
			},
			entity.ContainerPreTrainModel: {
				ContainerPath:     "/tmp" + pretrainModelPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{storageType},
			},
			entity.ContainerOutPutPath: {
				ContainerPath:       "/tmp" + outputPath,
				StorageRelativePath: cloudbrain.ModelMountPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{storageType},
				MKDIR:               false,
			},
		},
	}

	if opts.ComputeSource == models.NPU {
		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerCode: {
					ContainerPath:       "/cache" + codePath,
					StorageRelativePath: codePath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
				},
				entity.ContainerDataset: {
					ContainerPath:     "/cache" + datasetPath,
					ReadOnly:          true,
					AcceptStorageType: []entity.StorageType{storageType},
				},
				entity.ContainerPreTrainModel: {
					ContainerPath:     "/cache" + pretrainModelPath,
					ReadOnly:          true,
					AcceptStorageType: []entity.StorageType{storageType},
				},
				entity.ContainerOutPutPath: {
					ContainerPath:       "/cache" + outputPath,
					StorageRelativePath: setting.OutPutPath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
				},
			},
		}
	}
	config.ActionType = models.ActionCreateGrampusFinetuneTask
	config.IsActionUseJobId = false
	config.DatasetsMaxNum = 2 * setting.MaxDatasetNum
	return config
}

func (t GrampusFinetuneTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.HandleReqParameters).
		Next(t.CheckSourceTaskIsCleared).
		Next(t.CheckImageAvailable).
		Next(t.CheckBranchExists).
		Next(t.CheckBootFile).
		Next(t.CheckWorkerNum).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckModels).
		Next(t.HandleModelDatasetReqParamArray).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create GrampusTrainTaskTemplate err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (GrampusFinetuneTaskTemplate) HandleModelDatasetReqParamArray(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request

	if req.Params != "" {
		parameters := req.ParamArray
		err := json.Unmarshal([]byte(req.Params), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", req.Params, err)
			return response.PARAM_ERROR
		}

		index := findParameter(parameters, "model_name")
		if index != -1 {
			remove(parameters.Parameter, index)

		}
		index = findParameter(parameters, "dataset")
		if index != -1 {
			remove(parameters.Parameter, index)

		}
		parameters.Parameter = append(parameters.Parameter, models.Parameter{
			Label: "model_name",
			Value: req.ModelNames,
		})

		parameters.Parameter = append(parameters.Parameter, models.Parameter{
			Label: "dataset",
			Value: strings.TrimSuffix(ctx.Request.DatasetNames, ","),
		})
		ctx.Request.ParamArray = parameters
	}
	return nil
}

func findParameter(parameters models.Parameters, name string) int {
	index := -1
	for i, p := range parameters.Parameter {
		if p.Label == name {
			index = i
		}
	}
	return index
}

func remove(slice []models.Parameter, s int) []models.Parameter {
	return append(slice[:s], slice[s+1:]...)
}

func (g GrampusFinetuneTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
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
			},
		},
		TaskConfig: ctx.Config,
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateTrainJob(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateFinetuneTask",ctx.TraceContext))
	if err != nil {
		log.Error("GrampusTrainTaskTemplate CreateTrainJob err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}
