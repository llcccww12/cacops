package task

import (
	"regexp"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/routers/utils"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"github.com/google/uuid"
)

type GrampusModelExperienceTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusModelExperienceTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeModelExperience,
			Config:      GetGrampusModelExperienceConfig,
		},
	}
	RegisterTask(models.JobTypeModelExperience, entity.C2Net, t)
}

func GetGrampusModelExperienceConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/tmp/code"
	datasetPath := "/tmp/dataset"
	pretrainModelPath := "/tmp/pretrainmodel"
	outputPath := "/tmp/output"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())

	config := &entity.AITaskBaseConfig{
		ActionType:          models.ActionCreateGrampusModelExperienceTask,
		IsActionUseJobId:    false,
		DatasetsLimitSizeGB: setting.DebugAttachSize,
		DatasetsMaxNum:      setting.MaxDatasetNum,
		ModelLimitSizeGB:    setting.DEBUG_MODEL_SIZE_LIMIT_GB,
		ModelMaxNum:         setting.DEBUG_MODEL_NUM_LIMIT,
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:     codePath,
				ReadOnly:          false,
				AcceptStorageType: []entity.StorageType{storageType},
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
			entity.ContainerOutputAsModel: {
				ContainerPath:     pretrainModelPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{storageType},
			},
			entity.ContainerOutPutPath: {
				ContainerPath:       outputPath,
				StorageRelativePath: cloudbrain.ModelMountPath,
				AcceptStorageType:   []entity.StorageType{storageType},
				MKDIR:               false,
			},
		},
	}

	if opts.ComputeSource == models.NPUResource {
		config = &entity.AITaskBaseConfig{
			ActionType:          models.ActionCreateGrampusModelExperienceTask,
			IsActionUseJobId:    false,
			DatasetsLimitSizeGB: setting.DebugAttachSize,
			DatasetsMaxNum:      setting.MaxDatasetNum,
			ModelLimitSizeGB:    setting.DEBUG_MODEL_SIZE_LIMIT_GB,
			ModelMaxNum:         setting.DEBUG_MODEL_NUM_LIMIT,
			ContainerSteps:      map[entity.ContainerDataType]*entity.ContainerBuildOpts{},
		}

	}

	return config
}

func (t GrampusModelExperienceTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Info("GrampusModelExperienceTaskTemplate create")
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckCanFinetuneExperience).
		Next(t.HandleReqParameters).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckBranchExists).
		Next(t.CheckModels).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create CreateOnlineInfer err.%v", err)
		return nil, err
	}

	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g GrampusModelExperienceTaskTemplate) CheckCanFinetuneExperience(ctx *context.CreationContext) *response.BizError {
	if ctx.SourceCloudbrain == nil {
		return nil
	}

	if !ctx.SourceCloudbrain.CanFintuneExperience(ctx.User) {
		return response.CAN_NOT_FINETUNE_EXPERIENCE
	}

	return nil

}
func (g GrampusModelExperienceTaskTemplate) CheckMultiRequest(ctx *context.CreationContext) *response.BizError {

	log.Info("Start to CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	ctx.Request.EndPoint = utils.SubdomainFromUUID()
	ctx.Request.Port = setting.MODEL_EXPERIENCE.PORT_MIN

	if ctx.Request.EndPoint != "" {
		if len(ctx.Request.EndPoint) > 32 || strings.HasPrefix(ctx.Request.EndPoint, "/") || strings.HasSuffix(ctx.Request.EndPoint, "/") {
			return response.ENDPOINT_NOT_START_SLASH
		}
		if ctx.Request.Port < setting.MODEL_EXPERIENCE.PORT_MIN {
			return response.PORT_RANGE.WithParams(setting.MODEL_EXPERIENCE.PORT_MIN, setting.MODEL_EXPERIENCE.PORT_MAX)
		}
		if ctx.Request.Port > setting.MODEL_EXPERIENCE.PORT_MAX {
			return response.PORT_RANGE.WithParams(setting.MODEL_EXPERIENCE.PORT_MIN, setting.MODEL_EXPERIENCE.PORT_MAX)
		}
		regexStr := "^[a-z][a-z0-9]+$"
		urlRgex := regexp.MustCompile(regexStr)
		if !urlRgex.MatchString(ctx.Request.EndPoint) {
			return response.ENDPOINT_MUST_BE_VALID
		}
		task, err := models.GetOnlineInfoTaskByEndPoint(ctx.Request.EndPoint)
		if err != nil {
			log.Error("GetGrampusCountByUserID failed:%v", err)
			return response.SYSTEM_ERROR
		}
		if len(task) >= 1 {
			for _, tmp := range task {
				if tmp.JobName != ctx.Request.JobName {
					log.Error("the url has been used.path=" + ctx.Request.EndPoint + " jobName=" + tmp.JobName)
					return response.ENDPOINT_NOT_EQAUL_TASK
				}
			}
		}
	}

	task, err := models.GetModelExperienceNotFinalStatusTask(ctx.User.ID, []string{string(models.JobWaiting), string(models.JobRunning), string(models.LocalStatusPreparing), string(models.LocalStatusCreating)})
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err)
		return response.SYSTEM_ERROR
	}
	limitNum := GetUserMultiLimitNum(ctx.User.ID, ctx.Request.JobType)
	if len(task) >= limitNum {
		log.Error("the user already has running or waiting task.")
		return response.MULTI_TASK.WithParams(len(task))
	}
	for _, v := range task {
		if v.AppName == ctx.Request.AppName {
			log.Error("the user already has running or waiting task.")
			return response.MULTI_MODEL_TASK.WithParams(1)
		}
	}
	log.Info("CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g GrampusModelExperienceTaskTemplate) Delete(cloudbrainId int64) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return response.SYSTEM_ERROR
	}

	err := DelTask(cloudbrainId, c.DeleteModelExperience)

	if err != nil {
		log.Error("DelTask error,cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}
	log.Info("DelTask model experience success.cloudbrainId=%d", cloudbrainId)
	return nil
}

func (g GrampusModelExperienceTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("create GrampusModelExperienceTaskTemplate CallCreationAPI")
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
	var req entity.CreateNoteBookTaskRequest

	if form.ComputeSource.Name == models.NPUResource {

		parameters := form.ParamArray.Parameter
		if len(parameters) == 0 {
			log.Error("not have experience type parameter")
			return response.PARAM_ERROR
		}
		experirenceType := ""
		for _, param := range parameters {
			if param.Label == "experience_type" {
				experirenceType = param.Value
			}
		}
		req = entity.CreateNoteBookTaskRequest{
			Name: form.JobName,
			Tasks: []entity.NoteBookTask{
				{
					Name:           form.JobName,
					ResourceSpecId: ctx.Spec.SourceSpecId,
					ExperienceType: experirenceType,
					Spec:           ctx.Spec,
				},
			},
		}
	} else {

		var envVariables map[string]interface{} = map[string]interface{}{
			"TEMPLATE": models.GetParam(form.ParamArray.Parameter, "template"),
		}

		var endPoint *entity.SelfEndPoint

		if form.Port > 0 && form.EndPoint != "" {
			endPoint = &entity.SelfEndPoint{
				Port:     int64(form.Port),
				EndPoint: form.EndPoint,
			}
		} else {

			endPoint = &entity.SelfEndPoint{
				Port:     int64(setting.MODEL_EXPERIENCE.PORT_MIN),
				EndPoint: strings.ReplaceAll(uuid.New().String(), "-", ""),
			}

		}

		req = entity.CreateNoteBookTaskRequest{
			Name:     form.JobName,
			RepoName: "",
			Tasks: []entity.NoteBookTask{
				{
					Name:             form.JobName,
					ResourceSpecId:   ctx.Spec.SourceSpecId,
					ImageId:          form.ImageID,
					ImageUrl:         imageUrl,
					Datasets:         append(ctx.GetContainerDataArray(entity.ContainerDataset)),
					PreTrainModel:    append(ctx.GetContainerDataArray(entity.ContainerPreTrainModel), ctx.GetContainerDataArray(entity.ContainerOutputAsModel)...),
					Code:             ctx.GetContainerDataArray(entity.ContainerCode),
					OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
					AutoStopDuration: -1,
					Capacity:         setting.Capacity,
					Queues:           ctx.Queues,
					Spec:             ctx.Spec,
					BootFile:         ctx.Request.BootFile,
					EnvVariables:     envVariables,
					EndPoint:         endPoint,
				},
			},
		}
	}

	createTime := timeutil.TimeStampNow()
	res, err := c.CreateModelExperience(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateModelExperienceTask",ctx.TraceContext))
	if err != nil {
		log.Error("GrampusNoteBookTask ModelExperienceTask err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	if res.JobID == "" {
		log.Error("GrampusNoteBookTask ModelExperienceTask failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.CREATE_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}

func (g GrampusModelExperienceTaskTemplate) LoadSpec(ctx *context.CreationContext) *response.BizError {
	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, ctx.Request.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeModelExperience,
		ComputeResource: ctx.Request.ComputeSource.Name,
		Cluster:         ctx.Request.Cluster.GetParentCluster(),
		HasInternet:     ctx.Request.HasInternet,
	})
	if err != nil || spec == nil {
		return response.SPEC_NOT_AVAILABLE
	}
	ctx.Spec = spec
	return nil
}

func (GrampusModelExperienceTaskTemplate) GetAvailableQueues(ctx *context.CreationContext) *response.BizError {
	ctx.Queues = ctx.Spec.GetAvailableQueuesForNewRight(models.GetAvailableCenterIdOpts{
		UserId:      ctx.User.ID,
		JobType:     models.JobTypeModelExperience,
		HasInternet: ctx.Request.HasInternet,
	})
	return nil
}
