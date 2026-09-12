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
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/routers/utils"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	clear_service "code.gitea.io/gitea/services/cloudbrain"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/role"
)

type GrampusOnlineInferTaskTemplate struct {
	DefaultAITaskTemplate
}

func init() {
	t := &GrampusOnlineInferTaskTemplate{
		DefaultAITaskTemplate: DefaultAITaskTemplate{
			ClusterType: entity.C2Net,
			JobType:     models.JobTypeOnlineInference,
			Config:      GetGrampusOnlineInferConfig,
		},
	}
	RegisterTask(models.JobTypeOnlineInference, entity.C2Net, t)
}

func GetGrampusOnlineInferConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	codePath := "/tmp/code"
	datasetPath := "/tmp/dataset"
	pretrainModelPath := "/tmp/pretrainmodel"
	outputPath := "/tmp/output"
	storageType := storage_helper.GetStorageTypeFromIntType(models.GetDefaultStorageType())

	config := &entity.AITaskBaseConfig{
		ActionType:          models.ActionCreateGrampusGPUOnlineInferTask,
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

func (t GrampusOnlineInferTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Info("GrampusOnlineInferTaskTemplate create")
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckEndPoint).
		Next(t.CheckMultiRequest).
		Next(t.CheckBootFile).
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

func (g GrampusOnlineInferTaskTemplate) CheckEndPoint(ctx *context.CreationContext) *response.BizError {

	if ctx.Request.EndPoint == "" && ctx.Request.Port > 0 {
		return response.ENDPOINT_AND_PORT
	}
	if ctx.Request.EndPoint != "" && ctx.Request.Port == 0 {
		return response.ENDPOINT_AND_PORT
	}
	if ctx.Request.EndPoint != "" && ctx.Request.Port > 0 {
		if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH) {

			if !ctx.User.IsAdmin {
				return response.SYSTEM_ERROR
			}
		}
	}

	if ctx.Request.EndPoint == "" {
		ctx.Request.EndPoint = utils.SubdomainFromUUID()
		ctx.Request.Port = setting.MODEL_EXPERIENCE.PORT_MIN
	}

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

	return nil
}

func (g GrampusOnlineInferTaskTemplate) CheckMultiRequest(ctx *context.CreationContext) *response.BizError {

	log.Info("Start to online infer CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	jobType := string(ctx.Request.JobType)
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

func (g GrampusOnlineInferTaskTemplate) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Info("create GrampusOnlineInferTaskTemplate CallCreationAPI")
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
		User:                 ctx.User,
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
				AutoSave:         ctx.Request.AutoSave,
			},
		},
	}

	createTime := timeutil.TimeStampNow()
	res, err := c.CreateOnlineInfer(req, otel.GetTraceInfo(ctx.NewCloudbrain, "CreateOnlineInferTask", ctx.TraceContext))
	if err != nil {
		log.Error("GrampusNoteBookTask CreateOnlineInfer err.req=%+v err=%v", req, err)
		return response.NewBizError(err)
	}
	if res.JobID == "" {
		log.Error("GrampusNoteBookTask CreateOnlineInfer failed.Cloudbrain.JobID=%s", ctx.SourceCloudbrain.JobID)
		return response.CREATE_FAILED
	}
	ctx.Response = &entity.CreationResponse{
		JobID:      res.JobID,
		Status:     res.Status,
		CreateTime: createTime,
	}
	return nil
}

func (g GrampusOnlineInferTaskTemplate) LoadSpec(ctx *context.CreationContext) *response.BizError {
	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, ctx.Request.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeOnlineInference,
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

func (GrampusOnlineInferTaskTemplate) GetAvailableQueues(ctx *context.CreationContext) *response.BizError {
	ctx.Queues = ctx.Spec.GetAvailableQueuesForNewRight(models.GetAvailableCenterIdOpts{
		UserId:      ctx.User.ID,
		JobType:     models.JobTypeOnlineInference,
		HasInternet: ctx.Request.HasInternet,
	})
	return nil
}

func (t GrampusOnlineInferTaskTemplate) Restart(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
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
	sourceCloudbrain := ctx.SourceCloudbrain
	if sourceCloudbrain != nil {
		aiConfig := sourceCloudbrain.GetCloudbrainConfig()
		if aiConfig != nil {
			//清理已经存在的日志
			helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aiConfig.OutputStorageType))
			//查找日志文件
			files := getLogFilesInStorage(helper, aiConfig.OutputObjectPrefix, ".txt")
			if len(files) == 0 {
				//此时未找符合条件的文件
				log.Info("Not found file.")
			} else {
				filePath := strings.TrimSuffix(aiConfig.OutputObjectPrefix, "/") + "/" + files[0].FileName
				log.Info("start to del exist online infer log file=" + filePath)
				err := helper.DeleteFile(filePath)
				if err != nil {
					log.Error("Deletefile(%s) failed:%v", filePath, err)
				}
			}
		}

	}

	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID, Status: ctx.NewCloudbrain.Status}, nil
}

func getLogFilesInStorage(helper storage_helper.StorageHelper, objectKeyPrefix string, logSuffix string) []storage.FileInfo {
	//获取日志输出目录下文件列表
	log.Info("filepath=" + objectKeyPrefix)
	fileList, err := helper.GetAllObjectsUnderDir(objectKeyPrefix)
	if err != nil {
		log.Error("GetTrainLog read dir err.objectKeyPrefix=%s,err=%v", objectKeyPrefix, err)
		return nil
	}
	if len(fileList) == 0 {
		log.Info("not found file.....")
		return nil
	}
	logFiles := make([]storage.FileInfo, 0)
	for _, f := range fileList {
		if f.IsDir {
			continue
		}
		log.Info("f.FileName=" + f.FileName)
		if strings.HasSuffix(f.FileName, logSuffix) {
			logFiles = append(logFiles, f)
		}
	}
	return logFiles
}

func (g GrampusOnlineInferTaskTemplate) Delete(cloudbrainId int64) *response.BizError {
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
	log.Info("DelTask online info success.cloudbrainId=%d", cloudbrainId)
	return nil
}

func (g GrampusOnlineInferTaskTemplate) CallRestartAPI(ctx *context.CreationContext) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return response.SYSTEM_ERROR
	}
	log.Info("online infer to restart.")
	createTime := timeutil.TimeStampNow()
	autoStopDuration := getAutoStopDurationMs(ctx.SourceCloudbrain.TimeLimit, ctx.User.ID)
	res, err := c.RestartNoteBook(ctx.SourceCloudbrain.JobID, autoStopDuration, otel.GetTraceInfo(ctx.SourceCloudbrain, "RestartOnlineInferTask", nil))
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
