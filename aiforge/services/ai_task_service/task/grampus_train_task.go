package task

import (
	"strings"

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
	"code.gitea.io/gitea/services/role"
)

type GrampusTrainTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusTrainTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeTrain,
			Config:      GetGrampusTrainTaskConfig,
		},
	}
	RegisterTask(models.JobTypeTrain, entity.C2Net, t)
}

func GetGrampusTrainTaskConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
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
	switch opts.ComputeSource {
	case models.NPU:
		config.ActionType = models.ActionCreateGrampusNPUTrainTask
	case models.GPU:
		config.ActionType = models.ActionCreateGrampusGPUTrainTask
	case models.GCU:
		config.ActionType = models.ActionCreateGrampusGCUTrainTask
	case models.ILUVATAR:
		config.ActionType = models.ActionCreateGrampusILUVATARTrainTask
	case models.DCU:
		config.ActionType = models.ActionCreateGrampusDCUTrainTask
	case models.BIREN:
		config.ActionType = models.ActionCreateGrampusBIRENGPUTrainTask
	case models.METAX:
		config.ActionType = models.ActionCreateGrampusMETAXGPGPUTrainTask
	}
	config.IsActionUseJobId = true
	config.DatasetsMaxNum = setting.MaxDatasetNum
	config.ModelMaxNum = setting.DEBUG_MODEL_NUM_LIMIT
	return config
}

func (t GrampusTrainTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.HandleReqParameters).
		Next(t.CheckPrivilege4Continue).
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
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create GrampusTrainTaskTemplate err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g GrampusTrainTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
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

	var envVariables map[string]interface{} = map[string]interface{}{
		"TRACKSERVER": "",
		"RUNID":       "",
	}
	if setting.AimConfig.Enabled && role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		envVariables = map[string]interface{}{
			"TRACKSERVER": setting.AimConfig.EnvHost,
			"RUNID":       form.DisplayJobName,
		}

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
				EnvVariables:     envVariables,
			},
		},
		TaskConfig: ctx.Config,
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateTrainJob(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateTrainTask",ctx.TraceContext))
	if err != nil {
		log.Error("GrampusTrainTaskTemplate CreateTrainJob err.req=%+v err=%v", req, err)
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
