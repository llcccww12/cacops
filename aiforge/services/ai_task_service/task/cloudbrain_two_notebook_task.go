package task

import (
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

type CloudbrainTwoNotebookTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &CloudbrainTwoNotebookTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.OpenICloudbrainTwo,
			JobType:     models.JobTypeDebug,
			Config:      GetCloudbrainTwoNotebookConfig,
		},
	}
	RegisterTask(models.JobTypeDebug, entity.OpenICloudbrainTwo, t)
}

func GetCloudbrainTwoNotebookConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	//默认配置
	config := &entity.AITaskBaseConfig{
		ActionType:          models.ActionCreateDebugNPUTask,
		IsActionUseJobId:    false,
		DatasetsLimitSizeGB: setting.DebugAttachSize,
		DatasetsMaxNum:      setting.MaxDatasetNum,
		ModelLimitSizeGB:    setting.DEBUG_MODEL_SIZE_LIMIT_GB,
		ModelMaxNum:         setting.DEBUG_MODEL_NUM_LIMIT,
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				Disable:           false,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
			entity.ContainerDataset: {
				Disable:           false,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
			entity.ContainerPreTrainModel: {
				Disable:           false,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
		},
	}

	//在线运行notebook配置
	if opts.IsFileNoteBookRequest {
		config = &entity.AITaskBaseConfig{
			ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
				entity.ContainerFileNoteBookCode: {},
			},
		}
	}
	return config
}

func (t CloudbrainTwoNotebookTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		Next(t.CheckNotebookCount).
		Next(t.CheckModels).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckBranchExists).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create CloudbrainOneNotebookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil

}

func (t CloudbrainTwoNotebookTaskTemplate) Restart(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
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
	if err != nil {
		log.Error("Restart GrampusNoteBookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID, Status: ctx.NewCloudbrain.Status}, nil

}

func (g CloudbrainTwoNotebookTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("ctx is here:", ctx)
	c := g.GetMyCluster()
	if c == nil {
		return response.SYSTEM_ERROR
	}
	form := ctx.Request
	req := entity.CreateNoteBookTaskRequest{
		Name:        form.JobName,
		Description: form.Description,
		Tasks: []entity.NoteBookTask{
			{
				Name:             form.JobName,
				ResourceSpecId:   ctx.Spec.SourceSpecId,
				ImageId:          form.ImageID,
				ImageUrl:         strings.TrimSpace(form.ImageUrl),
				AutoStopDuration: getAutoStopDurationMs(ctx.NewCloudbrain.TimeLimit, ctx.User.ID),
				Spec:             ctx.Spec,
				Datasets:         ctx.GetContainerDataArray(entity.ContainerDataset),
				Code:             ctx.GetContainerDataArray(entity.ContainerCode),
				PreTrainModel:    ctx.GetContainerDataArray(entity.ContainerPreTrainModel),
				OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
			},
		},
	}
	createTime := timeutil.TimeStampNow()
	res, err := c.CreateNoteBook(req, nil)
	if err != nil {
		log.Error("CloudbrainTwoNotebookTaskTemplate CreateNoteBook err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}

	return nil
}

func (g CloudbrainTwoNotebookTaskTemplate) CallRestartAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return response.SYSTEM_ERROR
	}
	createTime := timeutil.TimeStampNow()
	autoStopDuration := getAutoStopDurationMs(ctx.SourceCloudbrain.TimeLimit, ctx.User.ID)
	res, err := c.RestartNoteBook(ctx.SourceCloudbrain.JobID, autoStopDuration, nil)
	if err != nil {
		log.Error("CloudbrainTwoNotebookTaskTemplate RestartNoteBook err.Cloudbrain.JobID=%s err=%v", ctx.SourceCloudbrain.JobID, err)
		return response.NewBizError(err)
	}
	if res.JobId == "" {
		log.Error("CloudbrainTwoNotebookTaskTemplate RestartNoteBook failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.RESTART_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobId,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}

func (g CloudbrainTwoNotebookTaskTemplate) GetSpecs(opts entity.GetSpecOpts) ([]*api.SpecificationShow, *response.BizError) {
	var aiCenterCode = models.AICenterOfCloudBrainTwo
	if setting.ModelartsCD.Enabled {
		aiCenterCode = models.AICenterOfChengdu
	}
	var specs []*models.AggregateSpecification
	var err error
	specs, err = resource.FindAvailableSpecsForNewRight(opts.UserId, models.FindSpecsOptions{
		JobType:         g.JobType,
		ComputeResource: opts.ComputeSource.Name,
		Cluster:         g.ClusterType.GetParentCluster(),
		AiCenterCode:    aiCenterCode,
		HasInternet:     opts.HasInternet,
	})

	if err != nil {
		log.Error("GetSpecs err.%v", err)
		return nil, response.SPEC_NOT_AVAILABLE
	}
	r := make([]*api.SpecificationShow, len(specs))
	for i, v := range specs {
		r[i] = convert.ToSpecificationFromAggregate(v)
	}
	return r, nil
}
