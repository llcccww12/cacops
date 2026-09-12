package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/eval"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"

	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
	"code.gitea.io/gitea/services/role"

	"strconv"

	"code.gitea.io/gitea/entity"
	grampus_client "code.gitea.io/gitea/manager/client/grampus"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/cluster"
	"code.gitea.io/gitea/services/ai_task_service/container_builder"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

var taskMap = map[string]AITaskTemplate{}

func RegisterTask(jobType models.JobType, clusterType entity.ClusterType, task AITaskTemplate) {
	k := string(jobType) + "_" + string(clusterType)
	taskMap[k] = task
}

func GetTask(jobType models.JobType, clusterType entity.ClusterType) (AITaskTemplate, error) {
	if jobType == "" || clusterType == "" {
		return nil, errors.New("jobType or clusterType is empty")
	}
	k := string(jobType) + "_" + string(clusterType)
	if t, ok := taskMap[k]; !ok {
		return nil, errors.New("Task not exists")
	} else {
		return t, nil
	}
}

type AITaskTemplate interface {
	Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError)
	Delete(cloudbrainId int64) *response.BizError
	Stop(cloudbrainId int64) (*entity.AITaskBriefInfo, *response.BizError)
	Query(cloudbrainId int64) (*entity.AITaskDetailInfo, *response.BizError)
	BriefQuery(cloudbrainId int64) (*entity.AITaskBriefInfo, *response.BizError)
	Restart(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError)
	Update(cloudbrainId int64, res ...*entity.QueryTaskResponse) *response.BizError
	GetLog(opts entity.QueryLogOpts) (*entity.ClusterLog, *response.BizError)
	GetLogDownloadInfo(opts entity.GetLogDownloadInfoReq) (*entity.FileDownloadInfo, *response.BizError)
	GetSingleOutputDownloadInfo(opts entity.GetSingleDownloadInfoReq) (*entity.FileDownloadInfo, *response.BizError)
	DownloadAllOutput(opts entity.DownloadAllFileReq) *response.BizError
	GetOutput(cloudbrainId int64, parentDir string) (*entity.AITaskOutput, *response.BizError)
	GetAllOutput(opts entity.GetAllOutputReq) (*entity.AllAITaskOutput, *response.BizError)
	GetDebugUrl(cloudbrainId int64, fileName ...string) (string, *response.BizError)
	GetVisualizeUrl(cloudbrainId int64) (string, *response.BizError)
	GetSelfEndPointUrl(cloudbrainId int64, fileName ...string) (string, *response.BizError)
	GetOperationProfile(cloudbrainId int64) (*entity.OperationProfile, *response.BizError)
	GetResourceUsage(opts entity.GetResourceUsageOpts) (*entity.ResourceUsage, *response.BizError)
	GetImages(computeSource models.ComputeSource, accCardType string, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, *response.BizError)
	GetSpecs(opts entity.GetSpecOpts) ([]*api.SpecificationShow, *response.BizError)
	GetConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig
	GetNodeInfo(cloudbrainId int64) ([]entity.AITaskNodeInfo, *response.BizError)
	GetAllowedWorkerNum(userId int64, computeSource *models.ComputeSource) ([]int, *response.BizError)
	GetDisplayJobName(userName string) string
}

type GetConfigFunc func(entity.AITaskConfigKey) *entity.AITaskBaseConfig

type DefaultAITaskTemplate struct {
	DefaultCreationHandler
	ClusterType entity.ClusterType
	JobType     models.JobType
	Config      GetConfigFunc
}

func (g DefaultAITaskTemplate) BuildContainerData(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to BuildContainerData.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	if ctx.Config == nil {
		return nil
	}
	err := container_builder.BuildContainerDataChain(ctx.Config.ContainerSteps).Run(ctx)
	if err != nil {
		return err
	}
	log.Info("BuildContainerData success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g DefaultAITaskTemplate) GetMyCluster() cluster.ClusterAdapter {
	c, err := cluster.GetCluster(g.ClusterType)
	if err != nil {
		log.Error("GetMyCluster err.%v", err)
		return nil
	}
	return c
}

func (g DefaultAITaskTemplate) NotifyCreation(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to NotifyCreation.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	req := ctx.Request
	jobID := ctx.Response.JobID
	config := ctx.Config
	user := ctx.User
	displayJobName := req.DisplayJobName

	if config.IsActionUseJobId {
		notification.NotifyOtherTask(user, nil, jobID, displayJobName, config.ActionType)
	} else {
		task, err := models.GetCloudbrainByJobID(jobID)
		if err != nil {
			log.Error("NotifyCreation GetCloudbrainByJobID failed: %v", err.Error())
			return nil
		}
		stringId := strconv.FormatInt(task.ID, 10)
		notification.NotifyOtherTask(user, nil, stringId, displayJobName, config.ActionType)
	}
	log.Info("NotifyCreation success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g DefaultAITaskTemplate) GetConfig(opts entity.AITaskConfigKey) *entity.AITaskBaseConfig {
	if g.Config == nil {
		return &entity.AITaskBaseConfig{}
	}
	c := g.Config(opts)
	if c == nil {
		return &entity.AITaskBaseConfig{}
	}
	return c
}

func (g DefaultAITaskTemplate) GetNodeInfo(cloudbrainId int64) ([]entity.AITaskNodeInfo, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return nil, response.SYSTEM_ERROR
	}
	res, err := GetAITaskNodeInfo(cloudbrainId, c.GetNodeInfo)
	if err != nil {
		log.Error("GetNodeInfo error,cloudbrainId=%d err=%v", cloudbrainId, err)
		return nil, response.NewBizError(err)
	}
	log.Info("GetNodeInfo success.cloudbrainId=%d", cloudbrainId)
	return res, nil
}

var GrampusNPUMultiNodeConfig *modelarts.MultiNodes
var CloudbrainTwoNPUMultiNodeConfig *modelarts.MultiNodes

func (g DefaultAITaskTemplate) GetAllowedWorkerNum(userId int64, computeSource *models.ComputeSource) ([]int, *response.BizError) {
	if g.JobType == models.JobTypeTrain && g.ClusterType == entity.C2Net && (computeSource.Name == models.NPU || computeSource.Name == models.GPU) {
		nums := role.AllUserOperNum(userId, role.ROLE_OPER_MULTI_NODE, computeSource.Name)

		if len(nums) > 0 {
			nums = append(nums, 1)
			sort.Ints(nums)
			return nums, nil
		}
		return []int{1}, nil
	}
	//未命中配置则只允许一个节点
	return []int{1}, nil
}

func (g DefaultAITaskTemplate) CheckDebugAddress(cloudbrain *models.Cloudbrain) bool {
	config := g.GetConfig(entity.AITaskConfigKey{
		ComputeSource:         cloudbrain.GetStandardComputeSource(),
		IsFileNoteBookRequest: cloudbrain.IsFileNoteBookTask(),
	})
	if !config.DebugAddressCheck || !setting.DEBUG_ADDRESS_CHECK_GLOBAL_TOGGLE {
		return true
	}
	debugUrl, err := g.GetDebugUrl(cloudbrain.ID)
	if err != nil || debugUrl == "" {
		log.Error("get debug url err.  cloudbrain.ID=%d, err=%v", cloudbrain.ID, err)
		return false
	}
	res, _ := http.Get(debugUrl)
	if res != nil && res.StatusCode == http.StatusBadGateway {
		log.Error("access debug url failed. cloudbrain.ID=%d, res.StatusCode=%d", cloudbrain.ID, res.StatusCode)
		return false
	}
	return true
}

func (g DefaultAITaskTemplate) Query(cloudbrainId int64) (*entity.AITaskDetailInfo, *response.BizError) {
	//查询时先更新，然后再查询本地数据
	g.Update(cloudbrainId)

	t, err := BuildAITaskInfo(cloudbrainId)
	if err != nil {
		log.Error("QueryTaskInfo err.cloudbrainId=%d err=%v", cloudbrainId, err)
		return nil, response.NewBizError(err)
	}
	return t, nil
}

func (g DefaultAITaskTemplate) BriefQuery(cloudbrainId int64) (*entity.AITaskBriefInfo, *response.BizError) {
	//查询时先更新，然后再查询本地数据
	g.Update(cloudbrainId)
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

func (g DefaultAITaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Error("Create func is not implements")
	return nil, response.SYSTEM_ERROR
}

func (g DefaultAITaskTemplate) Delete(cloudbrainId int64) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return response.SYSTEM_ERROR
	}
	var err error
	if g.JobType == models.JobTypeDebug {
		err = DelTask(cloudbrainId, c.DeleteNoteBook)
	} else {
		err = DelTask(cloudbrainId, c.DeleteTrainJob)
	}

	if err != nil {
		log.Error("DelTask error,cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}
	log.Info("DelTask success.cloudbrainId=%d", cloudbrainId)
	return nil
}

func (g DefaultAITaskTemplate) Stop(cloudbrainId int64) (*entity.AITaskBriefInfo, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return nil, response.SYSTEM_ERROR
	}
	var err error
	if g.JobType == models.JobTypeDebug {
		err = StopTask(cloudbrainId, c.StopNoteBook)
	} else if g.JobType == models.JobTypeModelExperience || g.JobType == models.JobTypeEval {
		err = StopTask(cloudbrainId, c.StopModelExperience)
	} else {
		err = StopTask(cloudbrainId, c.StopTrainJob)
	}

	if err != nil {
		log.Error("StopTask err.cloudbrainId=%d err=%v", cloudbrainId, err)
		if models.IsErrCannotStopCreatingGrampusJob(err) {
			return nil, response.CAN_NOT_STOP_CREATING_JOB
		}
		if models.IsErrCannotStopSavingImageJob(err) {
			return nil, response.CAN_NOT_STOP_SAVING_IMAGE_JOB
		}
		log.Error("StopTask err.cloudbrainId=%d err=%v", cloudbrainId, err)
		return nil, response.NewBizError(err)
	}
	log.Info("StopTask success.cloudbrainId=%d", cloudbrainId)
	return g.BriefQuery(cloudbrainId)
}

func (t DefaultAITaskTemplate) Restart(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	log.Error("Restart func is not implements")
	return nil, response.SYSTEM_ERROR
}

func (g DefaultAITaskTemplate) Update(cloudbrainId int64, reses ...*entity.QueryTaskResponse) *response.BizError {
	updateLock := redis_lock.NewDistributeLock(redis_key.AITaskUpdateLock(cloudbrainId))
	success, err := updateLock.LockWithWait(5*time.Second, 5*time.Second)
	if err != nil {
		log.Error("AI Task update lock err.cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}
	if !success {
		log.Error("AI Task update  lock failed.cloudbrainId=%d ", cloudbrainId)
		return nil
	}
	defer updateLock.UnLock()

	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return response.SYSTEM_ERROR
	}
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		log.Error("GetCloudbrainByCloudbrainID err when update ai task.cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}

	//如果任务已经是终态则无需更新
	if cloudbrain.IsTerminal() {
		log.Info("AI task is terminal.No need to update from remote.cloudbrainId=%d", cloudbrainId)
		return nil
	}

	//如果任务需要主动停止时执行主动停止逻辑
	//两种情况需要主动停止，一是处于CONNECTING状态，此时意味着异步创建任务出现了异常
	//二是处于PREPARING的时间超过了配置的等待时间，此时意味着异步创建任务时间过长或者出现了未知异常
	if cloudbrain.NeedActiveStop() {
		log.Info("AI task should active stop.cloudbrainId=%d", cloudbrainId)
		if g.JobType == models.JobTypeDebug || g.JobType == models.JobTypeOnlineInference {
			err = StopAITaskByJobNameFromRemote(cloudbrain, c.QueryNoteBookByJobName, c.StopNoteBook)
		} else if g.JobType == models.JobTypeModelExperience || g.JobType == models.JobTypeEval {
			err = StopAITaskByJobNameFromRemote(cloudbrain, c.QueryNoteBookByJobName, c.StopModelExperience)
		} else {
			err = StopAITaskByJobNameFromRemote(cloudbrain, c.QueryTrainJobByJobName, c.StopTrainJob)
		}

		if err != nil {
			log.Error("StopAITaskByJobNameFromRemote err.cloudbrainId=%d err=%v", cloudbrainId, err)
			return response.NewBizError(err)
		}
		return nil
	}

	//如果任务处于PREPARING状态，此时还处于等待异步创建结果中，无需更新
	if cloudbrain.IsPreparing() {
		log.Info("AI task is preparing.No need to update from remote.cloudbrainId=%d", cloudbrainId)
		return nil
	}

	//查询任务在集群上的状态
	log.Info("start to UpdateAITaskFromRemote.task.DisplayJobName = %s  task.Status = %s", cloudbrain.DisplayJobName, cloudbrain.Status)

	var res *entity.QueryTaskResponse
	if len(reses) > 0 {
		res = reses[0]
	} else {
		if g.JobType == models.JobTypeDebug || g.JobType == models.JobTypeOnlineInference {
			res, err = c.QueryNoteBook(entity.JobIdAndVersionId{JobID: cloudbrain.JobID, VersionID: cloudbrain.VersionID, TaskID: cloudbrain.ID})
		} else if g.JobType == models.JobTypeModelExperience || g.JobType == models.JobTypeEval {
			res, err = c.QueryModelExperience(entity.JobIdAndVersionId{JobID: cloudbrain.JobID, VersionID: cloudbrain.VersionID, ComputeResource: cloudbrain.ComputeResource, TaskID: cloudbrain.ID})
		} else {
			res, err = c.QueryTrainJob(entity.JobIdAndVersionId{JobID: cloudbrain.JobID, VersionID: cloudbrain.VersionID, TaskID: cloudbrain.ID})

		}
	}
	log.Info("remoteQueryFunc task.DisplayJobName = %s   res = %+v ", cloudbrain.DisplayJobName, res)
	if err != nil {
		log.Error("query from remote err.cloudbrainID = %d  err=%v", cloudbrain.ID, err)
		return response.NewBizError(err)
	}
	if res == nil || res.JobId == "" {
		log.Error("query from remote failed,response is empty,cloudbrainID = %d ", cloudbrain.ID)
		return response.NewBizError(errors.New("response is empty"))
	}

	if strings.ToUpper(res.Status) == string(models.JobRunning) && strings.ToUpper(cloudbrain.Status) == string(models.JobWaiting) {
		if !g.CheckDebugAddress(cloudbrain) {
			log.Error("check debug address failed.cloudbrain.id = %d", cloudbrain.ID)
			return nil
		}
	}

	err = UpdateByQueryResponse(res, cloudbrain)
	if err != nil {
		log.Error("UpdateByQueryResponse err.cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}

	err = modelExperienceLoad(cloudbrain, c, g)
	if err != nil {
		log.Error("UpdateByQueryResponse err.cloudbrainId=%d err=%v", cloudbrainId, err)
		return response.NewBizError(err)
	}

	log.Info("updateTask success.cloudbrainId=%d", cloudbrainId)
	return nil
}

func modelExperienceLoad(task *models.Cloudbrain, c cluster.ClusterAdapter, g DefaultAITaskTemplate) error {
	if task.JobType == string(models.JobTypeEval) && task.Status == string(models.JobRunning) {

		value, _ := redis_client.Get(redis_key.ModelEvalKey(task.JobID))
		if value == "" {
			redis_client.Setnx(redis_key.ModelEvalKey(task.JobID), "1", 30*24*3600*time.Second)

			var parameters models.Parameters

			if len(task.Parameters) != 0 {
				err := json.Unmarshal([]byte(task.Parameters), &parameters)
				if err != nil {
					log.Error("Failed to Unmarshal params: %s (%v)", task.Parameters, err)
					return err
				}
				datasetsStr := models.GetParam(parameters.Parameter, "datasets")
				limit, _ := strconv.Atoi(models.GetParam(parameters.Parameter, "limit"))
				log.Info("eval task start %v", task.ID)
				err = evalTask(task, datasetsStr, limit, g)
				if err != nil {
					log.Error("Failed to eval task:%v %v", task.ID, err)
					_, bizErr := g.Stop(task.ID)
					if bizErr == nil {
						task.Status = string(models.JobFailed)
						task.EndTime = timeutil.TimeStampNow()

						err := models.UpdateJob(task)
						if err != nil {
							log.Error("UpdateJob(%s) failed:%v", task.DisplayJobName, err)
							return err
						}
					}
				}

			}
		} else {
			evalResult, err := eval.GetEvalTaskResult(task.ID)
			if err != nil {
				log.Error("Failed to eval task:%v %v", task.ID, err)
				g.Stop(task.ID)
			}
			if evalResult.Code == 0 && evalResult.Status == "SUCCEED" {
				_, bizErr := g.Stop(task.ID)
				if bizErr == nil {

					task.Status = string(models.JobSucceeded)
					task.EndTime = timeutil.TimeStampNow()

					err := models.UpdateJob(task)
					if err != nil {
						log.Error("UpdateJob(%s) failed:%v", task.DisplayJobName, err)
						return err
					}
				}
			}

			if evalResult.Code != 0 && evalResult.Status == "FAILED" {

				_, bizErr := g.Stop(task.ID)
				if bizErr == nil {
					task.Status = string(models.JobFailed)
					task.EndTime = timeutil.TimeStampNow()

					err := models.UpdateJob(task)
					if err != nil {
						log.Error("UpdateJob(%s) failed:%v", task.DisplayJobName, err)
						return err
					}
				}

			}

		}

	}

	if (task.JobType == string(models.JobTypeModelExperience) || task.JobType == string(models.JobTypeEval)) && task.ComputeResource == models.NPUResource {
		if task.Status == string(models.JobRunning) {

			var parameters models.Parameters

			if len(task.Parameters) != 0 {
				err := json.Unmarshal([]byte(task.Parameters), &parameters)
				if err != nil {
					log.Error("Failed to Unmarshal params: %s (%v)", task.Parameters, err)
					return err
				}

				template := models.GetParam(parameters.Parameter, "template")
				sourceCloudBrainIdStr := models.GetParam(parameters.Parameter, "source_task_id")

				if template == "" {
					log.Error("do not have template parameter: %s (%v)", task.Parameters, err)
					return err

				} else {
					sourceCloudBrainId, _ := strconv.ParseInt(sourceCloudBrainIdStr, 10, 64)
					modelResultRelativePath := models.GetParam(parameters.Parameter, "sub_path")
					datasetsStr := ""
					limit := 0
					if task.JobType == string(models.JobTypeEval) {
						datasetsStr = models.GetParam(parameters.Parameter, "datasets")
						limit, _ = strconv.Atoi(models.GetParam(parameters.Parameter, "limit"))
					}
					go copyAndLoadModel(task, template, sourceCloudBrainId, modelResultRelativePath, c, datasetsStr, limit)

				}

			} else {
				log.Error("do not have template parameter: %s (%v)", task.Parameters)
				return errors.New("do not have template parameter")
			}

		}

	}
	return nil
}

func evalTask(task *models.Cloudbrain, datasetsStr string, limit int, g DefaultAITaskTemplate) error {
	evalResult, err := eval.EvalTask(&api.EvalRequest{
		TaskId:    task.ID,
		ModelName: task.AppName,
		Datasets:  datasetsStr,
		Limit:     limit,
	})
	if err != nil {
		return err

	}
	if evalResult.Code != 0 {
		return fmt.Errorf("failed to eval task %v", task.ID)
	}
	return nil
}

func copyAndLoadModel(task *models.Cloudbrain, template string, sourceCloudBrainId int64, modelResultRelativePath string, c cluster.ClusterAdapter, datasets string, limit int) {
	value, _ := redis_client.Get(redis_key.ModelExperienceModelLoadingKey(task.JobID))
	if value == "" {
		redis_client.Setnx(redis_key.ModelExperienceModelLoadingKey(task.JobID), "1", 30*24*3600*time.Second)
		err := grampus_client.ModelAppServiceCopyAndLoadModel(task.JobID, task.ModelId, sourceCloudBrainId, template, modelResultRelativePath)
		if err != nil {
			log.Error("copyAndLoadModel err.%v", err)
			StopAITaskByJobNameFromRemote(task, c.QueryNoteBookByJobName, c.StopModelExperience)
			return
		}

		redis_client.Setnx(redis_key.ModelExperienceModelLoadedKey(task.JobID), "1", 30*24*3600*time.Second)

	}

}

func (g DefaultAITaskTemplate) GetLog(opts entity.QueryLogOpts) (*entity.ClusterLog, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return nil, response.SYSTEM_ERROR
	}
	s, err := QueryTaskLog(opts, c.GetLog)
	if err != nil {
		log.Error("GetLog err.cloudbrainId=%d err =%v", opts, err)
		return &entity.ClusterLog{}, nil
	}

	return s, nil
}

func (g DefaultAITaskTemplate) GetLogDownloadInfo(opts entity.GetLogDownloadInfoReq) (*entity.FileDownloadInfo, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return nil, response.SYSTEM_ERROR
	}
	s, err := GetLogDownloadInfo(opts, c.GetLogDownloadInfo)
	if err != nil {
		log.Error("GetLog err.cloudbrainId=%d ", opts)
		return nil, nil
	}

	return s, nil
}

func (g DefaultAITaskTemplate) GetSingleOutputDownloadInfo(opts entity.GetSingleDownloadInfoReq) (*entity.FileDownloadInfo, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return nil, response.SYSTEM_ERROR
	}
	s, err := GetSingleOutputDownloadInfo(opts, c.GetSingleOutputDownloadInfo)
	if err != nil {
		log.Error("GetOutputDownloadInfo err.cloudbrainId=%d ", opts)
		return nil, nil
	}

	return s, nil
}

func (g DefaultAITaskTemplate) DownloadAllOutput(opts entity.DownloadAllFileReq) *response.BizError {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return response.SYSTEM_ERROR
	}
	err := DownloadAllOutput(opts, c.DownloadAllOutput)
	if err != nil {
		log.Error("GetOutputDownloadInfo err.cloudbrainId=%d ", opts)
		return nil
	}

	return nil
}

func (g DefaultAITaskTemplate) GetOutput(cloudbrainId int64, parentDir string) (*entity.AITaskOutput, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return nil, response.SYSTEM_ERROR
	}
	s, err := GetAITaskOutput(cloudbrainId, parentDir, c.GetOutput)
	if err != nil {
		log.Error("GetOutput err.cloudbrainId=%d err =%v", cloudbrainId, err)
		return nil, nil
	}
	return s, nil
}

func (g DefaultAITaskTemplate) GetAllOutput(opts entity.GetAllOutputReq) (*entity.AllAITaskOutput, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", opts)
		return nil, response.SYSTEM_ERROR
	}
	s, err := GetAllAITaskOutput(opts, c.GetAllOutput)
	if err != nil {
		log.Error("GetOutput err.cloudbrainId=%d err =%v", opts, err)
		return nil, nil
	}
	return s, nil
}

func (g DefaultAITaskTemplate) GetDebugUrl(cloudbrainId int64, fileName ...string) (string, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return "", response.SYSTEM_ERROR
	}
	f := ""
	if fileName != nil && len(fileName) > 0 {
		f = fileName[0]
	}
	s, err := QueryNoteBookUrl(cloudbrainId, c.GetNoteBookUrl, f)
	if err != nil {
		log.Error("QueryNoteBookUrl err.cloudbrainId=%d err =%v", cloudbrainId, err)
		return "", nil
	}
	return s, nil
}

func (g DefaultAITaskTemplate) GetVisualizeUrl(cloudbrainId int64) (string, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return "", response.SYSTEM_ERROR
	}
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		return "", response.NewBizError(err)
	}
	if cloudbrain.JobID == "" {
		return "", nil
	}
	url, err := c.GetVisualizeUrl(cloudbrain.JobID)
	if err != nil {
		log.Error("GetVisualizeUrl err.cloudbrainId=%d err =%v", cloudbrainId, err)
		return "", nil
	}
	return url, nil
}

func (g DefaultAITaskTemplate) GetSelfEndPointUrl(cloudbrainId int64, fileName ...string) (string, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return "", response.SYSTEM_ERROR
	}
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		return "", response.SYSTEM_ERROR
	}
	if cloudbrain.JobID == "" {
		return "", response.SYSTEM_ERROR
	}
	s, err := c.GetSelfEndPointUrl(cloudbrain.JobID)
	if err != nil {
		log.Error("QueryNoteBookUrl err.cloudbrainId=%d err =%v", cloudbrainId, err)
		return "", response.NewBizError(err)
	}
	return s, nil
}

func (g DefaultAITaskTemplate) GetOperationProfile(cloudbrainId int64) (*entity.OperationProfile, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,cloudbrainId=%d", cloudbrainId)
		return nil, response.SYSTEM_ERROR
	}
	var s *entity.OperationProfile
	var err error

	if g.JobType == models.JobTypeTrain || g.JobType == models.JobTypeFINETUNE {
		s, err = GetOperationProfile(cloudbrainId, c.GetTrainJobOperationProfile)
	} else {
		s, err = GetOperationProfile(cloudbrainId, c.GetNoteBookOperationProfile)
	}
	if err != nil {
		log.Error("GetOperationProfile err.cloudbrainId=%d err =%v", cloudbrainId, err)
		return nil, nil
	}
	return s, nil
}

func (g DefaultAITaskTemplate) GetResourceUsage(opts entity.GetResourceUsageOpts) (*entity.ResourceUsage, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed,opts=%+v", opts)
		return nil, response.SYSTEM_ERROR
	}
	res, err := GetResourceUsage(opts, c.GetResourceUsage)
	if err != nil {
		log.Error("GetOperationProfile err.opts=%+v err =%v", opts, err)
		if err.Error() == strconv.Itoa(response.RESOURCE_USAGE_NOT_SUPPORT.Code) {
			return nil, response.RESOURCE_USAGE_NOT_SUPPORT
		}
		return nil, nil
	}
	return res, nil
}

func (g DefaultAITaskTemplate) GetImages(computeSource models.ComputeSource, accCardType string, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return nil, false, response.SYSTEM_ERROR
	}

	var images []entity.ClusterImage
	var customFlag bool
	var err error
	if g.JobType == models.JobTypeDebug {
		images, customFlag, err = c.GetNotebookImages(entity.GetImageReq{
			ComputeSource: computeSource,
			JobType:       g.JobType,
			AccCardType:   accCardType,
		}, queues...)
	} else {
		images, customFlag, err = c.GetTrainImages(entity.GetImageReq{
			ComputeSource: computeSource,
			JobType:       g.JobType,
			AccCardType:   accCardType,
		}, queues...)
	}
	if err != nil {
		log.Error("GetImages err.computeSource=%s err =%v", computeSource.Name, err)
		return nil, false, response.NewBizError(err)
	}
	return images, customFlag, nil
}

func (g DefaultAITaskTemplate) GetSpecs(opts entity.GetSpecOpts) ([]*api.SpecificationShow, *response.BizError) {
	specs, err := resource.FindAvailableSpecsForNewRight(opts.UserId, models.FindSpecsOptions{
		JobType:           g.JobType,
		ComputeResource:   opts.ComputeSource.Name,
		Cluster:           g.ClusterType.GetParentCluster(),
		HasInternet:       opts.HasInternet,
		VisualizeRequired: opts.VisualizeRequired,
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

func GetSpecsNew(opts entity.GetSpecOpts) ([]*api.SpecificationShow, *response.BizError) {

	return nil, nil
}

func (g DefaultAITaskTemplate) CheckWorkerNum(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckMultiNode.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	serverNum := ctx.Request.WorkServerNumber
	if serverNum <= 1 {
		return nil
	}
	workerNums, _ := g.GetAllowedWorkerNum(ctx.User.ID, ctx.Request.ComputeSource)
	if !isInNodes(workerNums, serverNum) {
		return response.NO_NODE_RIGHR
	}
	return nil
}

func isInNodes(nodes []int, num int) bool {
	for _, node := range nodes {
		if node == num {
			return true
		}
	}
	return false

}
func (g DefaultAITaskTemplate) GetDisplayJobName(userName string) string {
	return cloudbrainService.GetDisplayJobName(userName)
}
