package task

import (
	"math/rand"
	"strconv"
	"time"

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

type IFLYTEKModelExperienceTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &IFLYTEKModelExperienceTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.IFLYTEKTraining,
			JobType:     models.JobTypeModelExperience,
			Config:      GetIFLYTEKModelExperienceConfig,
		},
	}
	RegisterTask(models.JobTypeModelExperience, entity.IFLYTEKTraining, t)
}

func GetIFLYTEKModelExperienceConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	config := &entity.AITaskBaseConfig{
		ActionType:     models.ActionCreateIFLYTEKModelExperienceTask,
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{},
	}
	return config
}

func (t IFLYTEKModelExperienceTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Info("IFLYTEKModelExperienceTaskTemplate create")
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.HandleReqParameters).
		Next(t.CheckMultiRequest).
		Next(t.CheckModels).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.GetAvailableQueues, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create CreateOnlineInfer err.%v", err)
		return nil, err
	}

	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g IFLYTEKModelExperienceTaskTemplate) CheckMultiRequest(ctx *context.CreationContext) *response.BizError {

	log.Info("Start to CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	task, err := models.GetModelExperienceNotFinalStatusTask(ctx.User.ID, []string{string(models.JobWaiting), string(models.JobRunning), string(models.LocalStatusPreparing), string(models.LocalStatusCreating)})
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err)
		return response.SYSTEM_ERROR
	}
	if len(task) >= setting.MODEL_EXPERIENCE.MAX_CREATES {
		log.Error("the user already has running or waiting task.")
		return response.MULTI_TASK.WithParams(len(task))
	}
	for _, v := range task {
		if v.AppName == ctx.Request.AppName {
			log.Error("the user already has running or waiting task.")
			return response.MULTI_TASK.WithParams(1)
		}
	}
	log.Info("CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g IFLYTEKModelExperienceTaskTemplate) Delete(cloudbrainId int64) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return response.SYSTEM_ERROR
	}

	err := DelTask(cloudbrainId, c.DeleteNoteBook)

	if err != nil {
		log.Error("DelTask error,cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}
	log.Info("DelTask model experience success.cloudbrainId=%d", cloudbrainId)
	return nil
}

func (g IFLYTEKModelExperienceTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("create IFLYTEKModelExperienceTaskTemplate CallCreationAPI")
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return response.SYSTEM_ERROR
	}
	form := ctx.Request

	req := entity.CreateNoteBookTaskRequest{
		SourceCloudbrain: ctx.SourceCloudbrain,
		Name:             form.JobName,
		RepoName:         ctx.Repository.Name,
		Tasks: []entity.NoteBookTask{
			{
				Name:           form.JobName,
				ResourceSpecId: ctx.Spec.SourceSpecId,
				PreTrainModel:  append(ctx.GetContainerDataArray(entity.ContainerPreTrainModel)),
				Queues:         ctx.Queues,
				Spec:           ctx.Spec,
			},
		},
	}

	createTime := timeutil.TimeStampNow()
	res, err := c.CreateModelExperience(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateIFLYTEKModelExperienceTask", ctx.TraceContext))
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

func (g IFLYTEKModelExperienceTaskTemplate) LoadSpec(ctx *context.CreationContext) *response.BizError {
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

func (IFLYTEKModelExperienceTaskTemplate) GetAvailableQueues(ctx *context.CreationContext) *response.BizError {
	ctx.Queues = ctx.Spec.GetAvailableQueuesForNewRight(models.GetAvailableCenterIdOpts{
		UserId:      ctx.User.ID,
		JobType:     models.JobTypeModelExperience,
		HasInternet: ctx.Request.HasInternet,
	})
	return nil
}

func (g IFLYTEKModelExperienceTaskTemplate) CheckModels(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckModels.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	if ctx.SourceCloudbrain == nil || ctx.SourceCloudbrain.DataUrl == "" {
		log.Error("CheckModels err.SourceCloudbrain is empty.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		return response.PARAM_ERROR
	}
	ctx.ContainerData = map[entity.ContainerDataType][]entity.ContainerData{
		entity.ContainerPreTrainModel: {
			entity.ContainerData{
				ObjectKey: ctx.SourceCloudbrain.DataUrl,
				Name:      ctx.SourceCloudbrain.AppName,
			}},
	}
	ctx.Request.AppName = ctx.SourceCloudbrain.AppName
	log.Info("CheckModels success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (IFLYTEKModelExperienceTaskTemplate) HandleReqParameters(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request
	req.JobName = "x" + time.Now().Format("20060102150405")[10:] + strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000)))
	return nil
}
