package task

import (
	"fmt"
	"io/ioutil"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"

	"code.gitea.io/gitea/modules/eval"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"

	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

type GrampusEvalTaskTemplate struct {
	GrampusModelExperienceTaskTemplate
}

func init() {
	t := &GrampusEvalTaskTemplate{
		GrampusModelExperienceTaskTemplate: GrampusModelExperienceTaskTemplate{
			DefaultAITaskTemplate: DefaultAITaskTemplate{
				ClusterType: entity.C2Net,
				JobType:     models.JobTypeEval,
				Config:      GetGrampusModelEvalConfig,
			},
		},
	}
	RegisterTask(models.JobTypeEval, entity.C2Net, t)
}

func GetGrampusModelEvalConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/tmp/code"
	datasetPath := "/tmp/dataset"
	pretrainModelPath := "/tmp/pretrainmodel"
	outputPath := "/tmp/output"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())

	config := &entity.AITaskBaseConfig{
		ActionType:          models.ActionCreateGrampusEvalTask,
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
			ActionType:          models.ActionCreateGrampusEvalTask,
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

func (t GrampusEvalTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Info("GrampusModelExperienceTaskTemplate create")
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckCanEval).
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

func (g GrampusEvalTaskTemplate) CheckCanEval(ctx *context.CreationContext) *response.BizError {
	if ctx.SourceCloudbrain == nil {
		return nil
	}

	if !ctx.SourceCloudbrain.CanFintuneExperience(ctx.User) {
		return response.CAN_NOT_Eval
	}

	return nil

}

func (g GrampusEvalTaskTemplate) GetLog(opts entity.QueryLogOpts) (*entity.ClusterLog, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return nil, response.SYSTEM_ERROR
	}
	content, err := eval.GetEvalTaskLog(opts.CloudbrainId)
	if err != nil {
		log.Error("GetLog err.cloudbrainId=%d err =%v", opts, err)
		return &entity.ClusterLog{}, nil
	}

	return &entity.ClusterLog{
		Content:        content,
		CanLogDownload: false,
	}, nil
}

func (g GrampusEvalTaskTemplate) GetLogDownloadInfo(opts entity.GetLogDownloadInfoReq) (*entity.FileDownloadInfo, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return nil, response.SYSTEM_ERROR
	}
	content, err := eval.GetEvalTaskLog(opts.CloudbrainId)
	if err != nil {
		log.Error("GetLog err.cloudbrainId=%d err =%v", opts.CloudbrainId, err)
		return nil, nil
	}
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		log.Error("GetLog err.cloudbrainId=%d err =%v", opts.CloudbrainId, err)
		return nil, nil
	}

	fileName := cloudbrain.JobName + "-log.txt"
	if cloudbrain.WorkServerNumber > 1 {
		fileName = cloudbrain.JobName + "-" + fmt.Sprint(opts.NodeId) + "-log.txt"
	}
	return &entity.FileDownloadInfo{
		Readers:        []entity.FileReader{{Reader: ioutil.NopCloser(strings.NewReader(content))}},
		ResultType:     entity.FileTypeTXT,
		ResultFileName: fileName,
	}, nil

}
func (g GrampusEvalTaskTemplate) CheckMultiRequest(ctx *context.CreationContext) *response.BizError {

	jobType := string(ctx.Request.JobType)
	log.Info("Start to CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, jobType)
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err)
		return response.SYSTEM_ERROR
	}

	limitNum := GetUserMultiLimitNum(ctx.User.ID, ctx.Request.JobType)
	if count >= limitNum {
		log.Error("the user already has running or waiting task.")
		return response.MULTI_TASK.WithParams(count)
	}
	log.Info("CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (g GrampusEvalTaskTemplate) Delete(cloudbrainId int64) *response.BizError {
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
	err = eval.ClearEvalTaskResult(cloudbrainId)
	if err != nil {
		log.Error("ClearEvalTaskResult error,cloudbrainId=%d err=%v", cloudbrainId, err)
	}
	log.Info("DelTask model experience success.cloudbrainId=%d", cloudbrainId)
	return nil
}

func (g GrampusEvalTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
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
				},
			},
		}
	}

	createTime := timeutil.TimeStampNow()
	res, err := c.CreateModelExperience(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateEvalTask", ctx.TraceContext))
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

func (g GrampusEvalTaskTemplate) LoadSpec(ctx *context.CreationContext) *response.BizError {
	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, ctx.Request.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeEval,
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

func (GrampusEvalTaskTemplate) GetAvailableQueues(ctx *context.CreationContext) *response.BizError {
	ctx.Queues = ctx.Spec.GetAvailableQueuesForNewRight(models.GetAvailableCenterIdOpts{
		UserId:      ctx.User.ID,
		JobType:     models.JobTypeEval,
		HasInternet: ctx.Request.HasInternet,
	})
	return nil
}
