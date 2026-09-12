package task

import (
	"strings"

	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/otel"

	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	clear_service "code.gitea.io/gitea/services/cloudbrain"
)

type GrampusNoteBookTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusNoteBookTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeDebug,
			Config:      GetGrampusNoteBookConfig,
		},
	}
	RegisterTask(models.JobTypeDebug, entity.C2Net, t)
}

func GetGrampusNoteBookConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/tmp/code"
	datasetPath := "/tmp/dataset"
	pretrainModelPath := "/tmp/pretrainmodel"
	outputPath := "/tmp/output"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())

	config := &entity.AITaskBaseConfig{
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:     codePath,
				ReadOnly:          false,
				AcceptStorageType: []entity.StorageType{storageType},
				VolumeFolder:      true,
				Uncompressed:      true,
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
			entity.ContainerOutPutPath: {
				ContainerPath:       outputPath,
				StorageRelativePath: cloudbrain.ModelMountPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{storageType},
				MKDIR:               false,
			},
		},
	}
	if opts.ComputeSource == models.CPU {
		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerCode: {
					ContainerPath:     codePath,
					ReadOnly:          false,
					AcceptStorageType: []entity.StorageType{storageType},
					VolumeFolder:      true,
					Uncompressed:      true,
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
				entity.ContainerOutPutPath: {
					ContainerPath:       outputPath,
					StorageRelativePath: cloudbrain.ModelMountPath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
					MKDIR:               false,
				},
			},
		}
	}

	if opts.ComputeSource == models.NPU {
		config = &entity.AITaskBaseConfig{
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
					VolumeFolder:      true,
				},
				entity.ContainerPreTrainModel: {
					ContainerPath:     pretrainModelPath,
					ReadOnly:          true,
					AcceptStorageType: []entity.StorageType{storageType},
				},
				entity.ContainerOutPutPath: {
					ContainerPath:       outputPath,
					StorageRelativePath: setting.OutPutPath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
					MKDIR:               false,
				},
			},
		}
	}

	if opts.ComputeSource == models.DCU {
		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerCode: {
					ContainerPath:     codePath,
					ReadOnly:          false,
					AcceptStorageType: []entity.StorageType{storageType},
					VolumeFolder:      true,
					Uncompressed:      true,
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
				entity.ContainerOutPutPath: {
					ContainerPath:       outputPath,
					StorageRelativePath: cloudbrain.ModelMountPath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
					MKDIR:               false,
				},
			},
		}
	}

	if opts.ComputeSource == models.GCU {
		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerCode: {
					ContainerPath:     codePath,
					ReadOnly:          false,
					AcceptStorageType: []entity.StorageType{storageType},
					VolumeFolder:      true,
					Uncompressed:      true,
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
				entity.ContainerOutPutPath: {
					ContainerPath:       outputPath,
					StorageRelativePath: cloudbrain.ModelMountPath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
					MKDIR:               false,
				},
			},
		}
	}
	//在线运行notebook配置
	if opts.IsFileNoteBookRequest {

		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerFileNoteBookCode: {},
				entity.ContainerCode: {
					ContainerPath:     codePath,
					ReadOnly:          false,
					AcceptStorageType: []entity.StorageType{storageType},
					Uncompressed:      true,
				},
				entity.ContainerDataset: {
					ContainerPath:     datasetPath,
					ReadOnly:          true,
					AcceptStorageType: []entity.StorageType{storageType},
				},
				entity.ContainerOutPutPath: {
					ContainerPath:       outputPath,
					StorageRelativePath: cloudbrain.ModelMountPath,
					ReadOnly:            false,
					AcceptStorageType:   []entity.StorageType{storageType},
					MKDIR:               false,
				},
			},
		}

	}

	switch opts.ComputeSource {
	case models.NPU:
		config.ActionType = models.ActionCreateGrampusNPUDebugTask
	case models.GPU:
		config.ActionType = models.ActionCreateGrampusGPUDebugTask
	case models.GCU:
		config.ActionType = models.ActionCreateGrampusGCUDebugTask
	case models.MLU:
		config.ActionType = models.ActionCreateGrampusMLUDebugTask
	case models.DCU:
		config.ActionType = models.ActionCreateGrampusDCUDebugTask
	case models.CPU:
		config.ActionType = models.ActionCreateSuperComputeTask
	case models.ILUVATAR:
		config.ActionType = models.ActionCreateGrampusILUVATARDebugTask
	case models.METAX:
		config.ActionType = models.ActionCreateGrampusMETAXDebugTask
	case models.BIREN:
		config.ActionType = models.ActionCreateGrampusBIRENGPUDebugTask
	}

	config.IsActionUseJobId = false
	config.DatasetsLimitSizeGB = setting.DebugAttachSize
	config.DatasetsMaxNum = setting.MaxDatasetNum
	config.ModelLimitSizeGB = setting.DEBUG_MODEL_SIZE_LIMIT_GB
	config.ModelMaxNum = setting.DEBUG_MODEL_NUM_LIMIT
	config.DebugAddressCheck = true
	return config
}

func (t GrampusNoteBookTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		Next(t.CheckNotebookCount).
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

func (t GrampusNoteBookTaskTemplate) Restart(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.BuildRequest4Restart).
		Next(t.CheckSourceTaskIsCleared).
		Next(t.CheckModels).
		Next(t.CheckDatasets).
		Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CallRestartAPI).
		Next(t.CreateCloudbrainRecord4Restart).
		Next(t.NotifyCreation).
		Operate(ctx)
	if err != nil {
		log.Error("Restart GrampusNoteBookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID, Status: ctx.NewCloudbrain.Status}, nil
}

var autoStopDurationMs = int64(4 * 60 * 60 * 1000)

func getAutoStopDurationMs(taskTimeLimit int, uid int64) int64 {
	if taskTimeLimit == 0 || !role.UserHasOper(uid, role.ROLE_OPER_DEBUG_TIME) {
		return autoStopDurationMs
	} else {
		if taskTimeLimit > 0 {
			return int64(taskTimeLimit * 60 * 60 * 1000)
		}
		return -1
	}
}

func (g GrampusNoteBookTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CallCreationAPI.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

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
	req := entity.CreateNoteBookTaskRequest{
		Name: form.JobName,
		Tasks: []entity.NoteBookTask{
			{
				Name:             form.JobName,
				ResourceSpecId:   ctx.Spec.SourceSpecId,
				ImageId:          form.ImageID,
				ImageUrl:         imageUrl,
				Datasets:         ctx.GetContainerDataArray(entity.ContainerDataset),
				PreTrainModel:    ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				Code:             ctx.GetContainerDataArray(entity.ContainerCode),
				OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
				EnvVariables:     map[string]interface{}{},
				AutoStopDuration: getAutoStopDurationMs(ctx.NewCloudbrain.TimeLimit, ctx.User.ID),
				Capacity:         setting.Capacity,
				Queues:           ctx.Queues,
				Spec:             ctx.Spec,
				AutoSave:         ctx.Request.AutoSave,
			},
		},
		//IsSubscriber: role.UserHasRole(ctx.User.ID, models.Subscriber),
		IsSubscriber: role.UserHasOper(ctx.User.ID, role.ROLE_OPER_DEBUG_TIME),
	}
	createTime := timeutil.TimeStampNow()

	res, err := c.CreateNoteBook(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateNoteBook", ctx.TraceContext))
	if err != nil {
		log.Error("GrampusNoteBookTask CreateNoteBook err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	if res.JobID == "" {
		log.Error("GrampusNoteBookTask CreateNoteBook failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.CREATE_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	log.Info("CallCreationAPI success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (g GrampusNoteBookTaskTemplate) CallRestartAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return response.SYSTEM_ERROR
	}
	createTime := timeutil.TimeStampNow()
	autoStopDuration := getAutoStopDurationMs(ctx.SourceCloudbrain.TimeLimit, ctx.User.ID)
	res, err := c.RestartNoteBook(ctx.SourceCloudbrain.JobID, autoStopDuration, otel.GetTraceInfo(ctx.SourceCloudbrain, "RestartNoteBook",nil))
	if err != nil {
		if err.Error() == entity.GrampusJobNotExistInCenter.CodeTrCode {
			clear_service.ClearCloudbrain(ctx.SourceCloudbrain)
		}
		log.Error("GrampusNoteBookTask RestartNoteBook err.Cloudbrain.JobID=%s err=%v", ctx.SourceCloudbrain.JobID, err)
		return response.NewBizError(err)
	}
	if res.JobId == "" {
		log.Error("GrampusNoteBookTask RestartNoteBook failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.RESTART_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobId,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}
