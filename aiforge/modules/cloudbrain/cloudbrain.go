package cloudbrain

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/timeutil"

	"code.gitea.io/gitea/modules/storage"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
)

const (
	//Command = `pip3 install jupyterlab==2.2.5 -i https://pypi.tuna.tsinghua.edu.cn/simple;service ssh stop;jupyter lab --no-browser --ip=0.0.0.0 --allow-root --notebook-dir="/code" --port=80 --LabApp.token="" --LabApp.allow_origin="self https://cloudbrain.pcl.ac.cn"`
	//CommandBenchmark      = `echo "start benchmark";python /code/test.py;echo "end benchmark"`
	CommandBenchmark       = `cd /benchmark && bash run_bk.sh >/model/benchmark-log.txt`
	CodeMountPath          = "/code"
	DataSetMountPath       = "/dataset"
	ModelMountPath         = "/model"
	PretrainModelMountPath = "/pretrainmodel"
	LogFile                = "log.txt"
	BenchMarkMountPath     = "/benchmark"
	BenchMarkResourceID    = 1
	Snn4imagenetMountPath  = "/snn4imagenet"
	BrainScoreMountPath    = "/brainscore"
	TaskInfoName           = "/taskInfo"
	Snn4imagenetCommand    = `/opt/conda/bin/python /benchmark/testSNN_script.py --modelname '%s' --modelpath '/pretrainmodel/%s' --modeldescription '%s' >/model/benchmark-log.txt`
	BrainScoreCommand      = `bash /benchmark/brainscore_test_par4shSrcipt.sh -b '%s' -n '%s' -p '/pretrainmodel/%s' -d '%s' >/model/benchmark-log.txt`
	Snn4EcosetCommand      = `/opt/conda/bin/python /benchmark/testSNN_script.py --datapath '/dataset' --modelname '%s' --modelpath '/pretrainmodel/%s'  --modeldescription '%s' >/model/benchmark-log.txt`
	Sim2BrianSnnCommand    = `cd /code && /opt/conda/bin/python /benchmark/score_a_model.py --benchmark '%s' --dataset-path '/dataset' --metric 'RSA' --model-name '%s' --model-checkpoint '/pretrainmodel/%s' --model-description '%s' >/model/benchmark-log.txt`
	SubTaskName            = "task1"

	Success = "S000"

	DefaultBranchName = "master"

	ResultPath = "/result"
)

var (
	ResourceSpecs          *models.ResourceSpecs
	TrainResourceSpecs     *models.ResourceSpecs
	InferenceResourceSpecs *models.ResourceSpecs
	SpecialPools           *models.SpecialPools
)

type GenerateCloudBrainTaskReq struct {
	Ctx                  *context.Context
	DisplayJobName       string
	JobName              string
	Image                string
	Command              string
	CodePath             string
	ModelPath            string
	BenchmarkPath        string
	Snn4ImageNetPath     string
	BrainScorePath       string
	JobType              string
	Description          string
	BranchName           string
	BootFile             string
	Params               string
	CommitID             string
	Uuids                string
	DatasetNames         string
	DatasetInfos         map[string]models.DatasetInfo
	BenchmarkTypeID      int
	BenchmarkChildTypeID int
	ResultPath           string
	TrainUrl             string
	ModelName            string
	ModelVersion         string
	CkptName             string
	ModelId              string
	LabelName            string
	PreTrainModelPath    string
	PreTrainModelUrl     string
	Spec                 *models.Specification
}

func GetCloudbrainDebugCommand() string {
	var command = `pip3 install jupyterlab==3 -i https://pypi.tuna.tsinghua.edu.cn/simple;pip3 install -U "nbclassic>=0.2.8" -i https://pypi.tuna.tsinghua.edu.cn/simple;service ssh stop;jupyter lab --ServerApp.shutdown_no_activity_timeout=` + setting.CullIdleTimeout + ` --TerminalManager.cull_inactive_timeout=` + setting.CullIdleTimeout + ` --TerminalManager.cull_interval=` + setting.CullInterval + ` --MappingKernelManager.cull_idle_timeout=` + setting.CullIdleTimeout + ` --MappingKernelManager.cull_interval=` + setting.CullInterval + ` --MappingKernelManager.cull_connected=True --MappingKernelManager.cull_busy=True --no-browser --ip=0.0.0.0 --allow-root --notebook-dir="/code" --port=80 --ServerApp.token="" --LabApp.token="" --ServerApp.allow_origin="self https://cloudbrain.pcl.ac.cn" `
	return command
}

func IsAdminOrOwnerOrJobCreater(ctx *context.Context, job *models.Cloudbrain, err error) bool {
	if !ctx.IsSigned {
		return false
	}
	if err != nil {

		return ctx.IsUserRepoOwner() || ctx.IsUserSiteAdmin()
	} else {
		return ctx.IsUserRepoOwner() || ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	}

}

func isAdminOrOwnerOrJobCreater(ctx *context.Context, job *models.Cloudbrain, err error) bool {
	if !ctx.IsSigned {
		return false
	}
	if err != nil {

		return ctx.IsUserRepoOwner() || ctx.IsUserSiteAdmin()
	} else {
		return ctx.IsUserRepoOwner() || ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	}

}

func CanDeleteJob(ctx *context.Context, job *models.Cloudbrain) bool {

	return isAdminOrOwnerOrJobCreater(ctx, job, nil)
}

func CanCreateOrDebugJob(ctx *context.Context) bool {
	if !ctx.IsSigned {
		return false
	}
	return ctx.Repo.CanWrite(models.UnitTypeCloudBrain)
}

func CanModifyJob(ctx *context.Context, job *models.Cloudbrain) bool {
	return isAdminOrJobCreater(ctx, job, nil)
}
func CanDownloadJob(ctx *context.Context, job *models.Cloudbrain) bool {
	return isAdminOrOwnerOrJobCreater(ctx, job, nil)
}

func isAdminOrJobCreater(ctx *context.Context, job *models.Cloudbrain, err error) bool {
	if !ctx.IsSigned {
		return false
	}
	if err != nil {
		return ctx.IsUserSiteAdmin()
	} else {
		return ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	}

}

func isAdminOrImageCreater(ctx *context.Context, image *models.Image, err error) bool {
	if !ctx.IsSigned {
		return false
	}
	if err != nil {
		return ctx.IsUserSiteAdmin()
	} else {
		return ctx.IsUserSiteAdmin() || ctx.User.ID == image.UID
	}

}

func AdminOrOwnerOrJobCreaterRight(ctx *context.Context) {

	var id = ctx.Params(":id")
	job, err := GetCloudBrainByIdOrJobId(id, "id")
	if err != nil {
		log.Error("GetCloudbrainByID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	ctx.Cloudbrain = job
	if !isAdminOrOwnerOrJobCreater(ctx, job, err) {
		log.Error("!isAdminOrOwnerOrJobCreater error:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}

}

func AdminOrJobCreaterRight(ctx *context.Context) {

	var id = ctx.Params(":id")
	if id == "" {
		id = ctx.Query("id")
	}
	job, err := GetCloudBrainByIdOrJobId(id, "id")
	if err != nil {
		log.Error("GetCloudbrainByID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	ctx.Cloudbrain = job
	if !isAdminOrJobCreater(ctx, job, err) {
		log.Error("!isAdminOrJobCreater error:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}

}

func AdminOrOwnerOrJobCreaterRightForTrain(ctx *context.Context) {

	var jobID = ctx.Params(":jobid")
	job, err := GetCloudBrainByIdOrJobId(jobID, "jobid")
	if err != nil {
		log.Error("GetCloudbrainByJobID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	ctx.Cloudbrain = job
	if !isAdminOrOwnerOrJobCreater(ctx, job, err) {
		log.Error("!isAdminOrOwnerOrJobCreater failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}

}

func AdminOrJobCreaterRightForTrain(ctx *context.Context) {

	var jobID = ctx.Params(":jobid")
	job, err := GetCloudBrainByIdOrJobId(jobID, "jobid")
	if err != nil {
		log.Error("GetCloudbrainByJobID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	ctx.Cloudbrain = job
	if !isAdminOrJobCreater(ctx, job, err) {
		log.Error("!isAdminOrJobCreater errot:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}

}

func AdminOrImageCreaterRight(ctx *context.Context) {

	id, err := strconv.ParseInt(ctx.Params(":id"), 10, 64)
	var image *models.Image
	if err != nil {
		log.Error("Get Image by ID failed:%v", err.Error())

	} else {
		image, err = models.GetImageByID(id)
		if err != nil {
			log.Error("Get Image by ID failed:%v", err.Error())
			return
		}
	}

	if !isAdminOrImageCreater(ctx, image, err) {
		log.Error("!isAdminOrImageCreater error:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}

}

func GenerateTask(req GenerateCloudBrainTaskReq) (string, error) {
	var versionCount int
	if req.JobType == string(models.JobTypeTrain) {
		versionCount = 1
	}

	volumes := []models.Volume{
		{
			HostPath: models.StHostPath{
				Path:      req.CodePath,
				MountPath: CodeMountPath,
				ReadOnly:  false,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      req.ModelPath,
				MountPath: ModelMountPath,
				ReadOnly:  false,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      req.BenchmarkPath,
				MountPath: BenchMarkMountPath,
				ReadOnly:  true,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      req.ResultPath,
				MountPath: ResultPath,
				ReadOnly:  false,
			},
		},
	}

	if req.PreTrainModelUrl != "" { //预训练
		volumes = append(volumes, models.Volume{
			HostPath: models.StHostPath{
				Path:      req.PreTrainModelPath,
				MountPath: PretrainModelMountPath,
				ReadOnly:  true,
			},
		})
	}

	if len(req.DatasetInfos) == 1 {
		volumes = append(volumes, models.Volume{
			HostPath: models.StHostPath{
				Path:      req.DatasetInfos[req.Uuids].DataLocalPath,
				MountPath: DataSetMountPath,
				ReadOnly:  true,
			},
		})
	} else if len(req.DatasetInfos) > 1 {
		for _, dataset := range req.DatasetInfos {
			volumes = append(volumes, models.Volume{
				HostPath: models.StHostPath{
					Path:      dataset.DataLocalPath,
					MountPath: DataSetMountPath + "/" + dataset.Name,
					ReadOnly:  true,
				},
			})
		}
	}

	createTime := timeutil.TimeStampNow()
	jobResult, err := CreateJob(req.JobName, models.CreateJobParams{
		JobName:    req.JobName,
		RetryCount: 1,
		GpuType:    req.Spec.QueueCode,
		Image:      req.Image,
		TaskRoles: []models.TaskRole{
			{
				Name:                  SubTaskName,
				TaskNumber:            1,
				MinSucceededTaskCount: 1,
				MinFailedTaskCount:    1,
				CPUNumber:             req.Spec.CpuCores,
				GPUNumber:             req.Spec.AccCardsNum,
				MemoryMB:              int(req.Spec.MemGiB * 1024),
				ShmMB:                 int(req.Spec.ShareMemGiB * 1024),
				Command:               req.Command,
				NeedIBDevice:          false,
				IsMainRole:            false,
				UseNNI:                false,
			},
		},
		Volumes: volumes,
	})
	if err != nil {
		log.Error("CreateJob failed:", err.Error(), req.Ctx.Data["MsgID"])
		return "", err
	}
	if jobResult.Code != Success {
		log.Error("CreateJob(%s) failed:%s", req.JobName, jobResult.Msg, req.Ctx.Data["MsgID"])
		return "", errors.New(jobResult.Msg)
	}

	var jobID = jobResult.Payload["jobId"].(string)
	err = models.CreateCloudbrain(&models.Cloudbrain{
		Status:               string(models.JobWaiting),
		UserID:               req.Ctx.User.ID,
		RepoID:               req.Ctx.Repo.Repository.ID,
		JobID:                jobID,
		JobName:              req.JobName,
		DisplayJobName:       req.DisplayJobName,
		SubTaskName:          SubTaskName,
		JobType:              req.JobType,
		Type:                 models.TypeCloudBrainOne,
		Uuid:                 req.Uuids,
		Image:                req.Image,
		GpuQueue:             req.Spec.QueueCode,
		ComputeResource:      models.GPUResource,
		BenchmarkTypeID:      req.BenchmarkTypeID,
		BenchmarkChildTypeID: req.BenchmarkChildTypeID,
		Description:          req.Description,
		IsLatestVersion:      "1",
		VersionCount:         versionCount,
		BranchName:           req.BranchName,
		BootFile:             req.BootFile,
		DatasetName:          req.DatasetNames,
		Parameters:           req.Params,
		TrainUrl:             req.TrainUrl,
		ModelName:            req.ModelName,
		ModelVersion:         req.ModelVersion,
		CkptName:             req.CkptName,
		ModelId:              req.ModelId,
		ResultUrl:            req.ResultPath,
		LabelName:            req.LabelName,
		PreTrainModelUrl:     req.PreTrainModelUrl,
		CreatedUnix:          createTime,
		UpdatedUnix:          createTime,
		CommitID:             req.CommitID,
		Spec:                 req.Spec,
	})

	if err != nil {
		return "", err
	}

	task, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err.Error())
		return "", err
	}

	stringId := strconv.FormatInt(task.ID, 10)

	if models.IsBenchMarkJobType(req.JobType) {
		notification.NotifyOtherTask(req.Ctx.User, req.Ctx.Repo.Repository, stringId, req.DisplayJobName, models.ActionCreateBenchMarkTask)
	} else if string(models.JobTypeTrain) == req.JobType {
		notification.NotifyOtherTask(req.Ctx.User, req.Ctx.Repo.Repository, jobID, req.DisplayJobName, models.ActionCreateGPUTrainTask)
	} else if string(models.JobTypeInference) == req.JobType {
		notification.NotifyOtherTask(req.Ctx.User, req.Ctx.Repo.Repository, jobID, req.DisplayJobName, models.ActionCreateInferenceTask)
	} else {
		notification.NotifyOtherTask(req.Ctx.User, req.Ctx.Repo.Repository, stringId, req.DisplayJobName, models.ActionCreateDebugGPUTask)
	}

	return jobID, nil
}

func GetWaitingCloudbrainCount(cloudbrainType int, computeResource string, jobTypes ...models.JobType) int64 {
	num, err := models.GetWaitingCloudbrainCount(cloudbrainType, computeResource, jobTypes...)
	if err != nil {
		log.Warn("Get waiting count err", err)
		num = 0
	}
	return num
}

func RestartTask(ctx *context.Context, task *models.Cloudbrain, newID *string) error {
	jobName := task.JobName

	spec := task.Spec
	var datasetInfos map[string]models.DatasetInfo
	if task.Uuid != "" {
		var err error
		datasetInfos, _, err = models.GetDatasetInfo(task.Uuid)
		if err != nil {
			log.Error("GetDatasetInfo failed:%v", err, ctx.Data["MsgID"])
			return err
		}
	}

	volumes := []models.Volume{
		{
			HostPath: models.StHostPath{
				Path:      storage.GetMinioPath(jobName, CodeMountPath+"/"),
				MountPath: CodeMountPath,
				ReadOnly:  false,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      storage.GetMinioPath(jobName, ModelMountPath+"/"),
				MountPath: ModelMountPath,
				ReadOnly:  false,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      storage.GetMinioPath(jobName, BenchMarkMountPath+"/"),
				MountPath: BenchMarkMountPath,
				ReadOnly:  true,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      storage.GetMinioPath(jobName, Snn4imagenetMountPath+"/"),
				MountPath: Snn4imagenetMountPath,
				ReadOnly:  true,
			},
		},
		{
			HostPath: models.StHostPath{
				Path:      storage.GetMinioPath(jobName, BrainScoreMountPath+"/"),
				MountPath: BrainScoreMountPath,
				ReadOnly:  true,
			},
		},
	}
	if datasetInfos != nil {
		if len(datasetInfos) == 1 {
			volumes = append(volumes, models.Volume{
				HostPath: models.StHostPath{
					Path:      datasetInfos[task.Uuid].DataLocalPath,
					MountPath: DataSetMountPath,
					ReadOnly:  true,
				},
			})
		} else {
			for _, dataset := range datasetInfos {
				volumes = append(volumes, models.Volume{
					HostPath: models.StHostPath{
						Path:      dataset.DataLocalPath,
						MountPath: DataSetMountPath + "/" + dataset.Name,
						ReadOnly:  true,
					},
				})
			}
		}
	}

	if task.PreTrainModelUrl != "" { //预训练
		volumes = append(volumes, models.Volume{
			HostPath: models.StHostPath{
				Path:      setting.Attachment.Minio.RealPath + task.PreTrainModelUrl,
				MountPath: PretrainModelMountPath,
				ReadOnly:  true,
			},
		})
	}

	createTime := timeutil.TimeStampNow()
	jobResult, err := CreateJob(jobName, models.CreateJobParams{
		JobName:    jobName,
		RetryCount: 1,
		GpuType:    task.GpuQueue,
		Image:      task.Image,
		TaskRoles: []models.TaskRole{
			{
				Name:                  SubTaskName,
				TaskNumber:            1,
				MinSucceededTaskCount: 1,
				MinFailedTaskCount:    1,
				CPUNumber:             spec.CpuCores,
				GPUNumber:             spec.AccCardsNum,
				MemoryMB:              int(spec.MemGiB * 1024),
				ShmMB:                 int(spec.ShareMemGiB * 1024),
				Command:               GetCloudbrainDebugCommand(), //Command,
				NeedIBDevice:          false,
				IsMainRole:            false,
				UseNNI:                false,
			},
		},
		Volumes: volumes,
	})
	if err != nil {
		log.Error("CreateJob failed:%v", err.Error(), ctx.Data["MsgID"])
		return err
	}
	if jobResult.Code != Success {
		log.Error("CreateJob(%s) failed:%s", jobName, jobResult.Msg, ctx.Data["MsgID"])
		return errors.New(jobResult.Msg)
	}

	var jobID = jobResult.Payload["jobId"].(string)
	newTask := &models.Cloudbrain{
		Status:          string(models.JobWaiting),
		UserID:          task.UserID,
		RepoID:          task.RepoID,
		JobID:           jobID,
		JobName:         task.JobName,
		DisplayJobName:  task.DisplayJobName,
		SubTaskName:     task.SubTaskName,
		JobType:         task.JobType,
		Type:            task.Type,
		Uuid:            task.Uuid,
		DatasetName:     task.DatasetName,
		Image:           task.Image,
		GpuQueue:        task.GpuQueue,
		ResourceSpecId:  task.ResourceSpecId,
		ComputeResource: task.ComputeResource,

		CreatedUnix:      createTime,
		UpdatedUnix:      createTime,
		BranchName:       task.BranchName,
		Spec:             spec,
		ModelName:        task.ModelName,
		ModelVersion:     task.ModelVersion,
		LabelName:        task.LabelName,
		PreTrainModelUrl: task.PreTrainModelUrl,
		CkptName:         task.CkptName,
		ModelId:          task.ModelId,
	}

	err = models.RestartCloudbrain(task, newTask)
	if err != nil {
		log.Error("RestartCloudbrain(%s) failed:%v", jobName, err.Error(), ctx.Data["MsgID"])
		return err
	}

	stringId := strconv.FormatInt(newTask.ID, 10)
	*newID = stringId

	notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, stringId, task.DisplayJobName, models.ActionCreateDebugGPUTask)

	return nil
}

func geMatchResourceSpec(jobType string, gpuQueue string, resourceSpecId int) *models.ResourceSpec {

	for _, specialPool := range SpecialPools.Pools {

		if specialPool.ResourceSpec != nil {
			if IsElementExist(specialPool.JobType, jobType) && IsQueueInSpecialtPool(specialPool.Pool, gpuQueue) {
				for _, spec := range specialPool.ResourceSpec {
					if resourceSpecId == spec.Id {
						return spec
					}
				}
			}
		}
	}
	return nil
}

func DelCloudBrainJob(jobId string) string {
	task, err := models.GetCloudbrainByJobID(jobId)
	if err != nil {
		log.Error("get cloud brain err:", err)
		return "cloudbrain.Delete_failed"

	}
	if task.Status != string(models.JobStopped) && task.Status != string(models.JobFailed) && task.Status != string(models.JobSucceeded) {
		log.Error("the job(%s) has not been stopped", task.JobName)
		return "cloudbrain.Not_Stopped"
	}

	err = models.DeleteJob(task)
	if err != nil {
		log.Error("DeleteJob failed:", err)
		return "cloudbrain.Delete_failed"
	}

	deleteJobStorage(task.JobName)

	return ""
}

func deleteJobStorage(jobName string) error {
	//delete local
	localJobPath := setting.JobPath + jobName
	err := os.RemoveAll(localJobPath)
	if err != nil {
		log.Error("RemoveAll(%s) failed:%v", localJobPath, err)
	}

	dirPath := setting.CBCodePathPrefix + jobName + "/"
	err = storage.Attachments.DeleteDir(dirPath)
	if err != nil {
		log.Error("DeleteDir(%s) failed:%v", localJobPath, err)
	}

	return nil
}

func InitSpecialPool() {
	if SpecialPools == nil && setting.SpecialPools != "" {
		json.Unmarshal([]byte(setting.SpecialPools), &SpecialPools)
	}
}

func IsResourceSpecInSpecialPool(resourceSpecs []*models.ResourceSpec, resourceSpecId int) bool {
	if resourceSpecs == nil || len(resourceSpecs) == 0 {
		return true
	}
	for _, v := range resourceSpecs {
		if v.Id == resourceSpecId {
			return true
		}
	}
	return false
}

func IsQueueInSpecialtPool(pool []*models.GpuInfo, queue string) bool {
	for _, v := range pool {
		if v.Queue == queue {
			return true
		}
	}
	return false
}

func IsElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetCloudBrainByIdOrJobId(id string, initialQuery string) (*models.Cloudbrain, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	var job *models.Cloudbrain
	if err != nil {

		job, err = models.GetCloudbrainByJobID(id)
	} else {

		if strings.EqualFold(initialQuery, "id") {
			job, err = models.GetCloudbrainByID(id)
			if err != nil {
				job, err = models.GetCloudbrainByJobID(id)
			}
		} else {
			job, err = models.GetCloudbrainByJobID(id)
			if err != nil {
				job, err = models.GetCloudbrainByID(id)
			}
		}

	}
	return job, err
}

type GenerateModelArtsNotebookReq struct {
	JobName        string
	DisplayJobName string
	Uuid           string
	Description    string

	BootFile string

	ImageId            string
	AutoStopDurationMs int64
	BranchName         string

	Spec             *models.Specification
	ModelName        string
	LabelName        string
	CkptName         string
	ModelId          string
	ModelVersion     string
	PreTrainModelUrl string
}
