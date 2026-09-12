package task

import (
	"strings"

	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

type GrampusModelSafetyTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusModelSafetyTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeModelSafety,
			Config:      GetGrampusModelSafetyTaskConfig,
		},
	}
	RegisterTask(models.JobTypeModelSafety, entity.C2Net, t)
}

func GetGrampusModelSafetyTaskConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/code"
	datasetPath := "/dataset"
	pretrainModelPath := "/pretrainmodel"
	outputPath := "/output"
	config := &entity.AITaskBaseConfig{
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:       "/cache" + codePath,
				StorageRelativePath: codePath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.OBS},
			},
			entity.ContainerDataset: {
				ContainerPath:     "/cache" + datasetPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
			entity.ContainerPreTrainModel: {
				ContainerPath:     "/cache" + pretrainModelPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.OBS},
			},
			entity.ContainerOutPutPath: {
				ContainerPath:       "/cache" + outputPath,
				StorageRelativePath: setting.OutPutPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.OBS},
			},
		},
	}

	config.ActionType = models.ActionCreateBenchMarkTask
	config.IsActionUseJobId = true
	return config
}

func (t GrampusModelSafetyTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.HandleReqParameters).
		Next(t.CheckBranchExists).
		Next(t.CheckBootFile).
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
		log.Error("create GrampusBenchmarkTaskTemplate err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g GrampusModelSafetyTaskTemplate) BriefQuery(cloudbrainId int64) (*entity.AITaskBriefInfo, *response.BizError) {

	t, err := QueryTaskBriefInfo(cloudbrainId)
	if err != nil {
		log.Error("QueryBrief err.err=%v", err)
		return nil, response.NewBizError(err)
	}

	// 镜像版本变更，补充为imageTag
	if t.ImageName != "" {

		imageInfo, _ := models.GetImageByPlace(t.ImageName)
		if imageInfo != nil && imageInfo.Tag != "" {
			t.ImageName = imageInfo.Tag
		}
	}

	return t, nil
}

func (g GrampusModelSafetyTaskTemplate) InsertCloudbrainRecord4Async(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to InsertCloudbrainRecord4Async.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	req := ctx.Request

	imageUrl := req.ImageUrl
	if req.ImageUrl == "" && req.ImageName != "" {
		imageUrl = req.ImageName
	}
	taskType := req.Cluster.GetCloudbrainType()

	branchName := req.BranchName

	//if a normal user set TimeLimit, it does not work
	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_DEBUG_TIME) {
		req.TimeLimit = 0
	}
	c := &models.Cloudbrain{
		Status:           models.LocalStatusPreparing,
		UserID:           ctx.User.ID,
		RepoID:           ctx.Repository.ID,
		JobName:          req.JobName,
		DisplayJobName:   req.DisplayJobName,
		JobType:          string(req.JobType),
		Type:             taskType,
		Uuid:             req.DatasetUUIDStr,
		DatasetName:      req.DatasetNames,
		CommitID:         ctx.CommitID,
		IsLatestVersion:  "1",
		VersionCount:     1,
		ComputeResource:  req.ComputeSource.GetCloudbrainFormat(),
		ImageID:          req.ImageID,
		Image:            imageUrl,
		BranchName:       branchName,
		Parameters:       req.Params,
		BootFile:         req.BootFile,
		Description:      req.Description,
		WorkServerNumber: req.WorkServerNumber,
		EngineName:       imageUrl,
		Spec:             ctx.Spec,
		ModelName:        req.ModelNames,
		ModelVersion:     models.GetParam(req.ParamArray.Parameter, "model_version"),
		CkptName:         models.GetParam(req.ParamArray.Parameter, "ckpt_name"),
		LabelName:        req.LabelName,
		ModelId:          req.PretrainModelId,
		SubTaskName:      models.SubTaskName,
		CreatedUnix:      timeutil.TimeStampNow(),
		UpdatedUnix:      timeutil.TimeStampNow(),
		GpuQueue:         ctx.Spec.QueueCode,
		AppName:          req.AppName,
		HasInternet:      int(req.HasInternet),
		TimeLimit:        req.TimeLimit,
	}

	err := models.CreateCloudbrain(c)

	if err != nil {
		log.Error("DefaultHandler InsertCloudbrainRecord(%s) failed:%v", req.DisplayJobName, err)
		return response.NewBizError(err)
	}
	ctx.NewCloudbrain = c
	log.Info("InsertCloudbrainRecord4Async success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (g GrampusModelSafetyTaskTemplate) LoadSpec(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to LoadSpec.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, ctx.Request.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: ctx.Request.ComputeSource.Name,
		Cluster:         ctx.Request.Cluster.GetParentCluster(),
		HasInternet:     ctx.Request.HasInternet,
	})
	if err != nil || spec == nil {
		return response.SPEC_NOT_AVAILABLE
	}
	ctx.Spec = spec
	log.Info("LoadSpec success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (GrampusModelSafetyTaskTemplate) GetAvailableQueues(ctx *context.CreationContext) *response.BizError {
	ctx.Queues = ctx.Spec.GetAvailableQueuesForNewRight(models.GetAvailableCenterIdOpts{
		UserId:      ctx.User.ID,
		JobType:     models.JobTypeTrain,
		HasInternet: ctx.Request.HasInternet,
	})
	return nil
}

func (g GrampusModelSafetyTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
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
	res, err := c.CreateTrainJob(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateBenchmarkTask", ctx.TraceContext))
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
