package task

import (
	"fmt"
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
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
)

type GrampusComfyuiExperienceTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusComfyuiExperienceTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeComfyuiExperience,
			Config:      GetGrampusComfyuiExperienceConfig,
		},
	}
	RegisterTask(models.JobTypeComfyuiExperience, entity.C2Net, t)
}

func GetGrampusComfyuiExperienceConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/tmp/code"
	datasetPath := "/tmp/dataset"
	pretrainModelPath := "/tmp/pretrainmodel"
	outputPath := "/tmp/output"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())

	config := &entity.AITaskBaseConfig{
		ActionType:          models.ActionCreateGrampusComfyuiExperienceTask,
		IsActionUseJobId:    false,
		DatasetsLimitSizeGB: setting.DebugAttachSize,
		DatasetsMaxNum:      setting.MaxDatasetNum,
		ModelLimitSizeGB:    setting.COMFYUI_EXPERIENCE_MODEL_SIZE_LIMIT_GB,
		ModelMaxNum:         setting.COMFYUI_EXPERIENCE_MODEL_NUM_LIMIT,
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
				AcceptStorageType:   []entity.StorageType{storageType},
				MKDIR:               false,
			},
		},
	}
	return config
}

func (t GrampusComfyuiExperienceTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Info("GrampusComfyuiExperienceTaskTemplate create")
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
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
		log.Error("create CreateComfyuiExperience err.%v", err)
		return nil, err
	}

	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (g GrampusComfyuiExperienceTaskTemplate) CheckMultiRequest(ctx *context.CreationContext) *response.BizError {

	log.Info("Start to online infer CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	if ctx.Request.EndPoint == "" && ctx.Request.Port > 0 {
		return response.ENDPOINT_AND_PORT
	}
	if ctx.Request.EndPoint != "" && ctx.Request.Port == 0 {
		return response.ENDPOINT_AND_PORT
	}
	// if ctx.Request.EndPoint != "" && ctx.Request.Port > 0 {
	// 	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH) {
	// 		//if !role.UserHasRole(ctx.User.ID, models.Subscriber) {
	// 		if !ctx.User.IsAdmin {
	// 			return response.SYSTEM_ERROR
	// 		}
	// 	}
	// }

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
		regexStr := "^[A-Za-z0-9/]+$"
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

	jobType := string(ctx.Request.JobType)
	log.Info("Start to CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, jobType)
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err)
		return response.SYSTEM_ERROR
	}
	limitNum := GetUserMultiLimitNum(ctx.User.ID, ctx.Request.JobType)
	log.Info("the user task count=" + fmt.Sprint(count) + " limitNum=" + fmt.Sprint(limitNum))
	if count >= limitNum {
		log.Error("the user already has running or waiting task.count=" + fmt.Sprint(count) + " limitNum=" + fmt.Sprint(limitNum))
		return response.MULTI_TASK.WithParams(count)
	}

	log.Info("CheckMulti online infer success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g GrampusComfyuiExperienceTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("create GrampusComfyuiExperienceTaskTemplate CallCreationAPI")
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
	var endPoint *entity.SelfEndPoint
	if form.Port > 0 && form.EndPoint != "" {
		endPoint = &entity.SelfEndPoint{
			Port:     int64(form.Port),
			EndPoint: form.EndPoint,
		}
	}
	req := entity.CreateNoteBookTaskRequest{
		Name:                 form.JobName,
		PrimitiveDatasetName: ctx.Request.DatasetNames,
		RepoName:             ctx.Repository.Name,
		Tasks: []entity.NoteBookTask{
			{
				Name:             form.JobName,
				ResourceSpecId:   ctx.Spec.SourceSpecId,
				ImageId:          form.ImageID,
				ImageUrl:         imageUrl,
				Datasets:         append(ctx.GetContainerDataArray(entity.ContainerDataset)),
				PreTrainModel:    append(ctx.GetContainerDataArray(entity.ContainerPreTrainModel)),
				Code:             ctx.GetContainerDataArray(entity.ContainerCode),
				OutPut:           ctx.GetContainerDataArray(entity.ContainerOutPutPath),
				AutoStopDuration: -1,
				Capacity:         setting.Capacity,
				Queues:           ctx.Queues,
				Spec:             ctx.Spec,
				BootFile:         ctx.Request.BootFile,
				EndPoint:         endPoint,
			},
		},
	}

	log.Info("endPoint is:%s", endPoint)

	createTime := timeutil.TimeStampNow()
	res, err := c.CreateComfyuiExperience(req,otel.GetTraceInfo(ctx.NewCloudbrain, "CreateComfyuiExperience", ctx.TraceContext))
	if err != nil {
		log.Error("GrampusNoteBookTask ComfyuiExperienceTask err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	if res.JobID == "" {
		log.Error("GrampusNoteBookTask ComfyuiExperienceTask failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.CREATE_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}
