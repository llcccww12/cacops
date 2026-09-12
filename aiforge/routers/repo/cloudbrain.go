package repo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/services/lock"

	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"

	"code.gitea.io/gitea/modules/dataset"
	ai_task "code.gitea.io/gitea/services/ai_task_service/task"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward/point/account"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/modules/grampus"

	"code.gitea.io/gitea/modules/timeutil"
	"github.com/unknwon/i18n"

	grampus_client "code.gitea.io/gitea/manager/client/grampus"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/util"
)

const (
	tplCloudBrainNew        base.TplName = "repo/cloudbrain/new"
	tplCloudBrainShow       base.TplName = "repo/cloudbrain/show"
	tplCloudBrainShowModels base.TplName = "repo/cloudbrain/models/index"

	tplCloudBrainBenchmarkIndex base.TplName = "repo/cloudbrain/benchmark/index"
	tplCloudBrainBenchmarkNew   base.TplName = "repo/cloudbrain/benchmark/new"
	tplCloudBrainBenchmarkShow  base.TplName = "repo/cloudbrain/benchmark/show"

	tplCloudBrainModelSafetyNewGpu base.TplName = "repo/modelsafety/newgpu"
	tplCloudBrainModelSafetyNewNpu base.TplName = "repo/modelsafety/newnpu"

	tplCloudBrainImageSubmit base.TplName = "repo/cloudbrain/image/submit"
	tplCloudBrainImageEdit   base.TplName = "repo/cloudbrain/image/edit"
	tplCloudBrainImageApply  base.TplName = "repo/cloudbrain/image/apply"

	tplCloudBrainTrainJobNew  base.TplName = "repo/cloudbrain/trainjob/new"
	tplCloudBrainTrainJobShow base.TplName = "repo/cloudbrain/trainjob/show"

	tplCloudBrainInferenceJobNew  base.TplName = "repo/cloudbrain/inference/new"
	tplCloudBrainInferenceJobShow base.TplName = "repo/cloudbrain/inference/show"
)

var (
	gpuInfos               *models.GpuInfos
	categories             *models.Categories
	benchmarkTypes         *models.BenchmarkTypes
	benchmarkGpuInfos      *models.GpuInfos
	benchmarkResourceSpecs *models.ResourceSpecs
	trainGpuInfos          *models.GpuInfos
	inferenceGpuInfos      *models.GpuInfos
)

const BENCHMARK_TYPE_CODE = "repo.cloudbrain.benchmark.types"
const CLONE_FILE_PREFIX = "file:///"
const README = "README"

var benchmarkTypesMap = make(map[string]*models.BenchmarkTypes, 0)

var jobNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$`)

// MustEnableDataset check if repository enable internal cb
func MustEnableCloudbrain(ctx *context.Context) {
	if !ctx.Repo.CanRead(models.UnitTypeCloudBrain) {
		ctx.NotFound("MustEnableCloudbrain", nil)
		return
	}
}

func cloudBrainNewDataPrepare(ctx *context.Context, jobType string) error {
	ctx.Data["PageIsCloudBrain"] = true
	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName

	ctx.Data["command"] = cloudbrain.GetCloudbrainDebugCommand()
	ctx.Data["code_path"] = cloudbrain.CodeMountPath
	ctx.Data["dataset_path"] = cloudbrain.DataSetMountPath
	ctx.Data["model_path"] = cloudbrain.ModelMountPath
	ctx.Data["benchmark_path"] = cloudbrain.BenchMarkMountPath
	ctx.Data["is_benchmark_enabled"] = setting.IsBenchmarkEnabled

	if categories == nil {
		json.Unmarshal([]byte(setting.BenchmarkCategory), &categories)
	}
	ctx.Data["benchmark_categories"] = categories.Category

	ctx.Data["benchmark_types"] = GetBenchmarkTypes(ctx).BenchmarkType
	// queuesDetail, _ := cloudbrain.GetQueuesDetail()
	// if queuesDetail != nil {
	// 	ctx.Data["QueuesDetail"] = queuesDetail
	// }

	prepareCloudbrainOneSpecs(ctx)

	ctx.Data["params"] = ""
	ctx.Data["branchName"] = ctx.Repo.BranchName

	ctx.Data["snn4imagenet_path"] = cloudbrain.Snn4imagenetMountPath
	ctx.Data["is_snn4imagenet_enabled"] = setting.IsSnn4imagenetEnabled

	ctx.Data["brainscore_path"] = cloudbrain.BrainScoreMountPath
	ctx.Data["is_brainscore_enabled"] = setting.IsBrainScoreEnabled

	ctx.Data["datasetType"] = models.TypeCloudBrainOne
	defaultMode := ctx.Query("benchmarkMode")
	if defaultMode == "" {
		defaultMode = "alogrithm"
	}
	ctx.Data["benchmarkMode"] = defaultMode
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, jobType)
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount

	if ctx.Cloudbrain != nil {
		ctx.Data["branch_name"] = ctx.Cloudbrain.BranchName
		ctx.Data["image"] = ctx.Cloudbrain.Image
		ctx.Data["image_id"] = ctx.Cloudbrain.ImageID
		ctx.Data["boot_file"] = ctx.Cloudbrain.BootFile
		ctx.Data["description"] = ctx.Cloudbrain.Description
		spec, _ := resource.GetCloudbrainSpec(ctx.Cloudbrain.ID)
		if spec != nil {
			ctx.Data["spec_id"] = spec.ID
		}
		ctx.Data["run_para_list"] = ctx.Cloudbrain.Parameters
		ctx.Data["model_name"] = ctx.Cloudbrain.ModelName
		ctx.Data["label_name"] = ctx.Cloudbrain.LabelName
		ctx.Data["ckpt_name"] = ctx.Cloudbrain.CkptName
		ctx.Data["model_id"] = ctx.Cloudbrain.ModelId
		ctx.Data["model_version"] = ctx.Cloudbrain.ModelVersion
		ctx.Data["pre_train_model_url"] = ctx.Cloudbrain.PreTrainModelUrl
		ctx.Data["compute_resource"] = ctx.Cloudbrain.ComputeResource
		uuids, datasetNames := dataset.GetFilterDeletedAttachments(ctx.Cloudbrain.Uuid)
		ctx.Data["attachment"] = uuids
		ctx.Data["dataset_name"] = datasetNames
		ctx.Data["cluster_type"] = models.OpenICluster
	}

	return nil
}

func prepareCloudbrainOneSpecs(ctx *context.Context) {
	debugSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeDebug,
		ComputeResource: models.GPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne,
	})
	ctx.Data["debug_specs"] = debugSpecs

	trainSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.GPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne,
	})
	ctx.Data["train_specs"] = trainSpecs

	inferenceSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeInference,
		ComputeResource: models.GPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne,
	})
	ctx.Data["inference_specs"] = inferenceSpecs

	benchmarkSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeBenchmark,
		ComputeResource: models.GPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne,
	})
	ctx.Data["benchmark_specs"] = benchmarkSpecs
}

func CloudBrainNew(ctx *context.Context) {
	// err := cloudBrainNewDataPrepare(ctx, string(models.JobTypeDebug))
	// if err != nil {
	// 	ctx.ServerError("get new cloudbrain info failed", err)
	// 	return
	// }
	// ctx.Data["PageIsGPUDebug"] = true
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplCloudBrainNew)
}

func cloudBrainCreate(ctx *context.Context, form auth.CreateCloudBrainForm) {
	ctx.Data["PageIsCloudBrain"] = true
	displayJobName := form.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	image := strings.TrimSpace(form.Image)
	uuids := form.Attachment
	jobType := form.JobType
	codePath := setting.JobPath + jobName + cloudbrain.CodeMountPath
	branchName := form.BranchName
	bootFile := strings.TrimSpace(form.BootFile)
	repo := ctx.Repo.Repository
	tpl := tplCloudBrainNew

	if jobType == string(models.JobTypeTrain) {
		tpl = tplCloudBrainTrainJobNew
	}

	spec, err := resource.GetAndCheckSpec(ctx.User.ID, form.SpecId, models.FindSpecsOptions{
		JobType:         models.JobType(jobType),
		ComputeResource: models.GPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne})
	if err != nil || spec == nil {
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr("Resource specification not available", tpl, &form)
		return
	}

	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
		log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr(ctx.Tr("points.insufficient_points_balance"), tpl, &form)
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: displayJobName, JobType: form.JobType}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr(ctx.Tr(errMsg), tpl, &form)
		return
	}

	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, jobType, displayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist", ctx.Data["MsgID"])
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr("the job name did already exist", tpl, &form)
			return
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err, ctx.Data["MsgID"])
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr("system error", tpl, &form)
			return
		}
	}

	if !jobNamePattern.MatchString(displayJobName) {
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_jobname_err"), tpl, &form)
		return
	}

	if jobType != string(models.JobTypeBenchmark) && jobType != string(models.JobTypeDebug) && jobType != string(models.JobTypeTrain) {
		log.Error("jobtype error:", jobType, ctx.Data["MsgID"])
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr("jobtype error", tpl, &form)
		return
	}

	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, jobType)
	if err != nil {
		log.Error("GetCloudbrainCountByUserID failed:%v", err, ctx.Data["MsgID"])
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr("system error", tpl, &form)
		return
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr(ctx.Tr("repo.cloudbrain.morethanonejob"), tpl, &form)
			return
		}
	}
	var datasetInfos map[string]models.DatasetInfo
	var datasetNames string
	var attachSize int64
	if uuids != "" {
		datasetInfos, datasetNames, err = models.GetDatasetInfo(uuids)
		if err != nil {
			log.Error("GetDatasetInfo failed: %v", err, ctx.Data["MsgID"])
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr(ctx.Tr("cloudbrain.error.dataset_select"), tpl, &form)
			return
		}

		if jobType == string(models.JobTypeDebug) {
			for _, infos := range datasetInfos {
				attachSize += infos.Size
			}
			if attachSize > int64(setting.DebugAttachSize*1024*1024*1024) {
				log.Error("The DatasetSize exceeds the limit (%dGB)", setting.DebugAttachSize) // GB
				cloudBrainNewDataPrepare(ctx, jobType)
				ctx.RenderWithErr(ctx.Tr("cloudbrain.error.debug_datasetsize", setting.DebugAttachSize), tpl, &form)
				return
			}
		}
	}

	command := cloudbrain.GetCloudbrainDebugCommand()
	if jobType == string(models.JobTypeTrain) {
		bootFileExist, err := ctx.Repo.FileExists(bootFile, branchName)
		if err != nil || !bootFileExist {
			log.Error("Get bootfile error:", err, ctx.Data["MsgID"])
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_bootfile_err"), tpl, &form)
			return
		}
		tpl = tplCloudBrainTrainJobNew
		commandTrain, err := getTrainJobCommand(form)
		if err != nil {
			log.Error("getTrainJobCommand failed: %v", err)
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr(err.Error(), tpl, &form)
			return
		}

		command = commandTrain
	}

	if branchName == "" {
		branchName = cloudbrain.DefaultBranchName
	}
	errStr := loadCodeAndMakeModelPath(repo, codePath, branchName, jobName, cloudbrain.ModelMountPath)
	if errStr != "" {
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr(ctx.Tr(errStr), tpl, &form)
		return
	}

	commitID, _ := ctx.Repo.GitRepo.GetBranchCommitID(branchName)

	req := cloudbrain.GenerateCloudBrainTaskReq{
		Ctx:                  ctx,
		DisplayJobName:       displayJobName,
		JobName:              jobName,
		Image:                image,
		Command:              command,
		Uuids:                uuids,
		DatasetNames:         datasetNames,
		DatasetInfos:         datasetInfos,
		CodePath:             storage.GetMinioPath(jobName, cloudbrain.CodeMountPath+"/"),
		ModelPath:            storage.GetMinioPath(jobName, cloudbrain.ModelMountPath+"/"),
		BenchmarkPath:        storage.GetMinioPath(jobName, cloudbrain.BenchMarkMountPath+"/"),
		Snn4ImageNetPath:     storage.GetMinioPath(jobName, cloudbrain.Snn4imagenetMountPath+"/"),
		BrainScorePath:       storage.GetMinioPath(jobName, cloudbrain.BrainScoreMountPath+"/"),
		JobType:              jobType,
		Description:          form.Description,
		BranchName:           branchName,
		BootFile:             form.BootFile,
		Params:               form.Params,
		CommitID:             commitID,
		BenchmarkTypeID:      0,
		BenchmarkChildTypeID: 0,
		ResultPath:           storage.GetMinioPath(jobName, cloudbrain.ResultPath+"/"),
		Spec:                 spec,
	}

	if form.ModelName != "" { //使用预训练模型训练

		req.ModelName = form.ModelName
		req.LabelName = form.LabelName
		req.CkptName = form.CkptName
		req.ModelId = form.ModelId
		req.ModelVersion = form.ModelVersion
		minioPreModelURL, err := dealModelInfo(form.ModelId, jobName, form.CkptName)
		if err != nil {
			log.Error("Can not find model", err)
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr(ctx.Tr("repo.modelconvert.manage.model_not_exist"), tpl, &form)
			return
		}
		req.PreTrainModelPath = setting.Attachment.Minio.RealPath + minioPreModelURL
		req.PreTrainModelUrl = minioPreModelURL
	}

	if form.IsContinue { // qizhi GPU 继续训练，将旧任务输出文件拷贝至新任务输出路径
		srcPath := "jobs/" + form.PreJobName + "/model/"
		destPath := "jobs/" + jobName + "/model/"
		err := MinioCopyResults(srcPath, destPath)
		if err != nil {
			log.Error("Cloudbrain GPU: Copy Prev Task Result File failed:", err.Error())
			cloudBrainNewDataPrepare(ctx, jobType)
			ctx.RenderWithErr("Failed to copy output files from previous train job", tpl, &form)
			return
		}
	}

	_, err = cloudbrain.GenerateTask(req)
	if err != nil {
		cloudBrainNewDataPrepare(ctx, jobType)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}
	if jobType == string(models.JobTypeTrain) {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=all")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/debugjob?debugListType=all")
	}
}

func dealModelInfo(modelId string, jobName string, ckptName string) (string, error) {
	preModel, err := models.QueryModelById(modelId)
	if err != nil || preModel == nil || preModel.ID == "" {
		log.Error("Can not find model", err)
		return "", fmt.Errorf("Can not find model: %v", ckptName)
	}
	minioPreModelURL, err := downloadModelFromObs(preModel, jobName, cloudbrain.PretrainModelMountPath, ckptName)
	if err != nil {
		log.Error("Can not find model", err)

		return "", err
	}
	return minioPreModelURL, nil
}

func downloadModelFromObs(preModel *models.AiModelManage, jobName, suffixPath string, ckptFileName string) (string, error) {
	destPath := setting.CBCodePathPrefix + jobName + suffixPath + "/"
	destFile := destPath + ckptFileName
	returnStr := setting.Attachment.Minio.Bucket + "/" + destPath
	if len(preModel.Path) > len(setting.Bucket)+1 {
		srcUrl := preModel.Path[len(setting.Bucket)+1:] + ckptFileName
		log.Info("dest model Path=" + returnStr + " src path=" + preModel.Path + ckptFileName)
		body, err := storage.ObsDownloadAFile(setting.Bucket, srcUrl)
		if err == nil {
			defer body.Close()
			_, err = storage.Attachments.UploadContent(setting.Attachment.Minio.Bucket, destFile, body)
			if err != nil {
				log.Error("UploadObject(%s) failed: %s", preModel.Path+ckptFileName, err.Error())
				return "", err
			}
		} else {
			log.Info("download model failed. as " + err.Error())
			return "", err
		}
		log.Info("download model from obs succeed")
	}
	return returnStr, nil
}

func MinioCopyResults(srcPath string, destPath string) error {
	log.Info("prev task obs path:", setting.Attachment.Minio.Bucket+srcPath)
	log.Info("current task obs path:", setting.Attachment.Minio.Bucket+destPath)
	allfile, _ := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, srcPath)
	var fileNames []string
	for _, file := range allfile {
		if strings.Contains(file.FileName, "README") || strings.Contains(file.FileName, ".txt") {
			continue
		}
		fileNames = append(fileNames, file.FileName)
	}
	log.Info("Previous task all files", fileNames)

	fileSizeAll, err := storage.MinioCopyFiles(setting.Attachment.Minio.Bucket, srcPath, destPath, fileNames)
	log.Info("%v output files copied from previous task", fileSizeAll)

	return err
}

func loadCodeAndMakeModelPath(repo *models.Repository, codePath string, branchName string, jobName string, resultPath string) string {
	err := downloadCode(repo, codePath, branchName)
	if err != nil {
		return "cloudbrain.load_code_failed"
	}

	err = uploadCodeToMinio(codePath+"/", jobName, cloudbrain.CodeMountPath+"/")
	if err != nil {
		return "cloudbrain.load_code_failed"
	}

	return initModelPath(jobName, resultPath)

}

func initModelPath(jobName string, resultPath string) string {
	modelPath := setting.JobPath + jobName + resultPath + "/"
	err := mkModelPath(modelPath)
	if err != nil {
		return "cloudbrain.load_code_failed"
	}
	err = uploadCodeToMinio(modelPath, jobName, resultPath+"/")
	if err != nil {
		return "cloudbrain.load_code_failed"
	}

	return ""
}

func hasDatasetDeleted(task *models.Cloudbrain) bool {
	if task.Uuid == "" {
		return false
	}
	uuids := strings.Split(task.Uuid, ";")
	attachs, _ := models.GetAttachmentsByUUIDs(uuids)
	return len(attachs) < len(uuids)
}

func CloudBrainBenchMarkShow(ctx *context.Context) {
	cloudBrainShow(ctx, tplCloudBrainBenchmarkShow, models.JobTypeBenchmark)
}

func CloudBrainShow(ctx *context.Context) {
	// cloudBrainShow(ctx, tplCloudBrainShow, models.JobTypeDebug)
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplCloudBrainShow)
}

func CloudBrainTrainJobShow(ctx *context.Context) {
	// cloudBrainShow(ctx, tplCloudBrainTrainJobShow, models.JobTypeTrain)
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplCloudBrainTrainJobShow)
}

func cloudBrainShow(ctx *context.Context, tpName base.TplName, jobType models.JobType) {
	ctx.Data["PageIsCloudBrain"] = true
	debugListType := ctx.Query("debugListType")
	cloudbrain.InitSpecialPool()

	var task *models.Cloudbrain
	var err error
	if jobType == models.JobTypeTrain || jobType == models.JobTypeInference {
		task, err = models.GetCloudbrainByJobID(ctx.Params(":jobid"))
	} else {
		task, err = models.GetCloudbrainByIDWithDeleted(ctx.Params(":id"))
	}
	if task.JobType == string(models.JobTypeModelSafety) {
		GetAiSafetyTaskTmpl(ctx)
		return
	}
	if err != nil {
		log.Info("error:" + err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}
	prepareSpec4Show(ctx, task)
	if ctx.Written() {
		return
	}
	if task.Status == string(models.JobWaiting) || task.Status == string(models.JobRunning) {
		if task.IsNewAITask() {
			task, _ = ai_task.UpdateCloudbrain(task)
		} else {
			task, err = cloudbrainTask.SyncCloudBrainOneStatus(task)
			if err != nil {
				log.Info("error:" + err.Error())
				ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
				return
			}
		}

	}

	user, err := models.GetUserByID(task.UserID)
	if err == nil {
		task.User = user
	}

	if task.BenchmarkTypeID > 0 {
		for _, benchmarkType := range GetBenchmarkTypes(ctx).BenchmarkType {
			if task.BenchmarkTypeID == benchmarkType.Id {
				ctx.Data["BenchmarkTypeName"] = benchmarkType.First
				task.BenchmarkTypeName = benchmarkType.First
				for _, benchmarkChildType := range benchmarkType.Second {
					if task.BenchmarkChildTypeID == benchmarkChildType.Id {
						ctx.Data["BenchmarkChildTypeName"] = benchmarkChildType.Value
						break
					}
				}
				break
			}
		}
	}

	if task.JobType == string(models.JobTypeBenchmark) {
		task.BenchmarkType = ctx.Tr("repo.cloudbrain.benchmark.algorithm")
	} else if models.IsBenchMarkJobType(task.JobType) {
		task.BenchmarkType = ctx.Tr("repo.cloudbrain.benchmark.model")
		task.BenchmarkTypeName = task.JobType
		ctx.Data["BenchmarkTypeName"] = task.JobType
		if task.JobType == string(models.JobTypeBrainScore) {
			ctx.Data["BenchmarkChildTypeName"] = getBrainRegion(task.BenchmarkChildTypeID)
		}
		if task.JobType == string(models.JobTypeSim2BrainSNN) {
			ctx.Data["BenchmarkChildTypeName"] = getSim2BrainDatasetType(task.BenchmarkChildTypeID)
		}

	}

	if task.TrainJobDuration == "" {
		if task.Duration == 0 {
			var duration int64
			if task.Status == string(models.JobWaiting) {
				duration = 0
			} else if task.Status == string(models.JobRunning) {
				duration = time.Now().Unix() - int64(task.CreatedUnix)
			} else {
				duration = int64(task.UpdatedUnix) - int64(task.CreatedUnix)
			}
			task.Duration = duration
		}
		task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)
	}
	ctx.Data["duration"] = task.TrainJobDuration

	if len(task.Parameters) > 0 {
		var parameters models.Parameters

		err := json.Unmarshal([]byte(task.Parameters), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal Parameters: %s (%v)", task.Parameters, err)
			task.Parameters = ""
		} else {
			if len(parameters.Parameter) > 0 {
				paramTemp := ""
				for _, Parameter := range parameters.Parameter {
					param := Parameter.Label + " = " + Parameter.Value + "; "
					paramTemp = paramTemp + param
				}
				task.Parameters = paramTemp[:len(paramTemp)-2]
			} else {
				task.Parameters = ""
			}
		}

	}
	ctx.Data["datasetDownload"] = GetCloudBrainDataSetInfo(task.Uuid, task.DatasetName, false)
	ctx.Data["task"] = task
	labelName := strings.Fields(task.LabelName)
	ctx.Data["LabelName"] = labelName
	ctx.Data["jobName"] = task.JobName
	ctx.Data["displayJobName"] = task.DisplayJobName
	version_list_task := make([]*models.Cloudbrain, 0)
	version_list_task = append(version_list_task, task)
	ctx.Data["version_list_task"] = version_list_task
	ctx.Data["debugListType"] = debugListType
	ctx.Data["code_path"] = cloudbrain.CodeMountPath
	ctx.Data["dataset_path"] = cloudbrain.DataSetMountPath
	ctx.Data["model_path"] = cloudbrain.ModelMountPath
	ctx.Data["canDownload"] = cloudbrain.CanDownloadJob(ctx, task)
	ctx.Data["branchName"] = task.BranchName
	ctx.HTML(200, tpName)
}

func CloudBrainDebug(ctx *context.Context) {
	task := ctx.Cloudbrain
	debugUrl := setting.DebugServerHost + "jpylab_" + task.JobID + "_" + task.SubTaskName

	if ctx.QueryTrim("file") != "" {
		ctx.Redirect(getFileUrl(debugUrl, ctx.QueryTrim("file")))
	} else {
		if task.BootFile != "" {
			go cloudbrainTask.UploadNotebookFiles(task)
		}
		ctx.Redirect(debugUrl)
	}

}

func prepareSpec4Show(ctx *context.Context, task *models.Cloudbrain) {
	s, err := resource.GetCloudbrainSpec(task.ID)
	if err != nil {
		log.Info("error:" + err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}
	ctx.Data["Spec"] = s
}

func oldPrepareSpec4Show(ctx *context.Context, task *models.Cloudbrain) {
	hasSpec := false
	if task.JobType == string(models.JobTypeTrain) {
		if cloudbrain.TrainResourceSpecs == nil {
			json.Unmarshal([]byte(setting.TrainResourceSpecs), &cloudbrain.TrainResourceSpecs)
		}

		for _, tmp := range cloudbrain.TrainResourceSpecs.ResourceSpec {
			if tmp.Id == task.ResourceSpecId {
				hasSpec = true
				ctx.Data["GpuNum"] = tmp.GpuNum
				ctx.Data["CpuNum"] = tmp.CpuNum
				ctx.Data["MemMiB"] = tmp.MemMiB
				ctx.Data["ShareMemMiB"] = tmp.ShareMemMiB
				break
			}
		}

	} else if task.JobType == string(models.JobTypeInference) {
		if cloudbrain.InferenceResourceSpecs == nil {
			json.Unmarshal([]byte(setting.InferenceResourceSpecs), &cloudbrain.InferenceResourceSpecs)
		}
		for _, tmp := range cloudbrain.InferenceResourceSpecs.ResourceSpec {
			if tmp.Id == task.ResourceSpecId {
				hasSpec = true
				ctx.Data["GpuNum"] = tmp.GpuNum
				ctx.Data["CpuNum"] = tmp.CpuNum
				ctx.Data["MemMiB"] = tmp.MemMiB
				ctx.Data["ShareMemMiB"] = tmp.ShareMemMiB
				break
			}
		}
	} else {
		if cloudbrain.ResourceSpecs == nil {
			json.Unmarshal([]byte(setting.ResourceSpecs), &cloudbrain.ResourceSpecs)
		}
		for _, tmp := range cloudbrain.ResourceSpecs.ResourceSpec {
			if tmp.Id == task.ResourceSpecId {
				hasSpec = true
				ctx.Data["GpuNum"] = tmp.GpuNum
				ctx.Data["CpuNum"] = tmp.CpuNum
				ctx.Data["MemMiB"] = tmp.MemMiB
				ctx.Data["ShareMemMiB"] = tmp.ShareMemMiB
				break

			}
		}
	}

	if !hasSpec && cloudbrain.SpecialPools != nil {

		for _, specialPool := range cloudbrain.SpecialPools.Pools {

			if specialPool.ResourceSpec != nil {

				for _, spec := range specialPool.ResourceSpec {
					if task.ResourceSpecId == spec.Id {
						ctx.Data["GpuNum"] = spec.GpuNum
						ctx.Data["CpuNum"] = spec.CpuNum
						ctx.Data["MemMiB"] = spec.MemMiB
						ctx.Data["ShareMemMiB"] = spec.ShareMemMiB
						break
					}
				}
			}
		}
	}
}

func CloudBrainCommitImageShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["Type"] = ctx.Cloudbrain.Type
	ctx.HTML(200, tplCloudBrainImageSubmit)
}

func GetImage(ctx *context.Context) {

	var ID = ctx.Params(":id")
	id, _ := strconv.ParseInt(ID, 10, 64)

	image, err := models.GetImageByID(id)
	if err != nil {
		log.Error("GetImageByID failed:%v", err.Error())
		ctx.JSON(http.StatusNotFound, nil)
	}
	ctx.JSON(http.StatusOK, image)

}

func CloudBrainImageEdit(ctx *context.Context) {
	ctx.Data["PageIsImageEdit"] = true
	ctx.Data["PageFrom"] = ctx.Params(":from")
	var ID = ctx.Params(":id")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Error("GetImageByID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	image, err := models.GetImageByID(id)
	if err != nil {
		log.Error("GetImageByID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	ctx.Data["Image"] = image
	ctx.HTML(http.StatusOK, tplCloudBrainImageEdit)

}

func CloudBrainImageApplyRecommend(ctx *context.Context) {
	ctx.Data["PageIsImageEdit"] = true
	ctx.Data["PageFrom"] = "apply"
	var ID = ctx.Params(":id")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Error("GetImageByID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	image, err := models.GetImageByID(id)
	if err != nil {
		log.Error("GetImageByID failed:%v", err.Error())
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
	}
	ctx.Data["Image"] = image
	ctx.HTML(http.StatusOK, tplCloudBrainImageApply)

}

func CloudBrainImageRecommendApplyPost(ctx *context.Context, form auth.EditImageCloudBrainForm) {

	if utf8.RuneCountInString(form.Description) > 1000 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", 1000)))
		return
	}

	validTopics, errMessage := checkTopics(form.Topics)
	if errMessage != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr(errMessage)))
		return
	}
	image, err := models.GetImageByID(form.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_not_exist")))
		return
	}

	image.Description = form.Description
	image.Framework = form.Framework
	image.FrameworkVersion = form.FrameworkVersion
	image.CudaVersion = form.CudaVersion
	image.CannVersion = form.CannVersion
	image.PythonVersion = form.PythonVersion
	image.OperationSystem = form.OperationSystem
	image.OperationSystemVersion = form.OperationSystemVersion
	image.ThirdPackages = form.ThirdPackages
	image.DTKVersion = form.DTKVersion

	if image.ApplyStatus == models.Applying || image.ApplyStatus == models.OKApply {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_apply_duplicate")))
		return
	}

	image.ApplyStatus = models.Applying
	image.Message = ""

	err = models.WithTx(func(ctx models.DBContext) error {
		if err := models.UpdateLocalImage(image); err != nil {
			return err
		}
		if err := models.SaveImageTopics(image.ID, validTopics...); err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_not_exist")))

	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}

func CloudBrainImageEditPost(ctx *context.Context, form auth.EditImageCloudBrainForm) {

	if utf8.RuneCountInString(form.Description) > 1000 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", 1000)))
		return
	}

	validTopics, errMessage := checkTopics(form.Topics)
	if errMessage != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr(errMessage)))
		return
	}
	image, err := models.GetImageByID(form.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_not_exist")))
		return
	}

	image.Description = form.Description
	image.Framework = form.Framework
	image.FrameworkVersion = form.FrameworkVersion
	image.CudaVersion = form.CudaVersion
	image.CannVersion = form.CannVersion
	image.PythonVersion = form.PythonVersion
	image.OperationSystem = form.OperationSystem
	image.OperationSystemVersion = form.OperationSystemVersion
	image.ThirdPackages = form.ThirdPackages
	image.DTKVersion = form.DTKVersion

	err = models.WithTx(func(ctx models.DBContext) error {
		if err := models.UpdateLocalImage(image); err != nil {
			return err
		}
		if err := models.SaveImageTopics(image.ID, validTopics...); err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_not_exist")))

	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}

func CloudBrainImageDelete(ctx *context.Context) {
	var ID = ctx.Params(":id")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_not_exist")))
		return
	}
	image, err := models.GetImageByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_not_exist")))
		return
	}
	err = grampus.DeleteImage(image)
	if err != nil {
		log.Info("Delete remote delete failed.error=" + err.Error())
	}
	err = models.DeleteLocalImage(id)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_delete_fail")))
	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}

func CloudBrainCommitImageCheck(ctx *context.Context, form auth.CommitImageCloudBrainForm) {
	isExist, _ := models.IsImageExistByUser(form.Tag, ctx.User.ID)
	if isExist {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.image_overwrite")))
	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}

func CloudBrainAdminCommitImage(ctx *context.Context, form auth.CommitAdminImageCloudBrainForm) {

	if !NamePattern.MatchString(form.Tag) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.title_format_err")))
		return
	}

	if utf8.RuneCountInString(form.Description) > 1000 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", "1000")))
		return
	}

	validTopics, errMessage := checkTopics(form.Topics)
	if errMessage != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr(errMessage)))
		return
	}

	err := cloudbrain.CommitAdminImage(models.CommitImageParams{
		CommitImageCloudBrainParams: models.CommitImageCloudBrainParams{
			ImageDescription: form.Description,
			ImageTag:         form.Tag,
		},
		IsPrivate:              form.IsPrivate,
		CloudBrainType:         form.Type,
		Topics:                 validTopics,
		UID:                    ctx.User.ID,
		Type:                   models.GetRecommondType(form.IsRecommend),
		Place:                  form.Place,
		Framework:              form.Framework,
		FrameworkVersion:       form.FrameworkVersion,
		CudaVersion:            form.CudaVersion,
		PythonVersion:          form.PythonVersion,
		OperationSystem:        form.OperationSystem,
		OperationSystemVersion: form.OperationSystemVersion,
		ThirdPackages:          form.ThirdPackages,
		ComputeResource:        form.ComputeResource,
		DTKVersion:             form.DTKVersion,
		TrainType:              form.TrainType,
	}, ctx.User)
	if err != nil {
		log.Error("CommitImagefailed")
		if models.IsErrImageTagExist(err) {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_exist")))

		} else if models.IsErrorImageCommitting(err) {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_committing")))
		} else {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_commit_fail")))
		}

		return
	}

	ctx.JSON(200, models.BaseOKMessage)
}

func CloudBrainCommitImage(ctx *context.Context, form auth.CommitImageCloudBrainForm) {

	if !NamePattern.MatchString(form.Tag) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.images.name_format_err")))
		return
	}

	if utf8.RuneCountInString(form.Description) > 1000 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", "1000")))
		return
	}

	validTopics, errMessage := checkTopics(form.Topics)
	if errMessage != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr(errMessage)))
		return
	}

	err := cloudbrain.CommitImage(ctx.Cloudbrain.JobID, models.CommitImageParams{
		CommitImageCloudBrainParams: models.CommitImageCloudBrainParams{
			Ip:               ctx.Cloudbrain.ContainerIp,
			TaskContainerId:  ctx.Cloudbrain.ContainerID,
			ImageDescription: form.Description,
			ImageTag:         form.Tag,
		},
		IsPrivate:              form.IsPrivate,
		CloudBrainType:         form.Type,
		Topics:                 validTopics,
		UID:                    ctx.User.ID,
		Framework:              form.Framework,
		FrameworkVersion:       form.FrameworkVersion,
		CudaVersion:            form.CudaVersion,
		PythonVersion:          form.PythonVersion,
		OperationSystem:        form.OperationSystem,
		OperationSystemVersion: form.OperationSystemVersion,
		ThirdPackages:          form.ThirdPackages,
		ComputeResource:        form.ComputeResource,
	}, ctx.User)
	if err != nil {
		log.Error("CommitImage(%s) failed:%v", ctx.Cloudbrain.JobName, err.Error(), ctx.Data["msgID"])
		if models.IsErrImageTagExist(err) {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_exist")))

		} else if models.IsErrorImageCommitting(err) {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_committing")))
		} else if isOver20GError(err) {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_over_20g")))
		} else {
			ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_commit_fail")))
		}

		return
	}
	ctx.JSON(200, models.BaseOKMessage)
}

func isOver20GError(err error) bool {
	return strings.Contains(err.Error(), "over max image size 20GB")
}

func checkTopics(Topics string) ([]string, string) {
	var topics = make([]string, 0)
	var topicsStr = strings.TrimSpace(Topics)
	if len(topicsStr) > 0 {
		topics = strings.Split(topicsStr, ",")
	}

	validTopics, invalidTopics := models.SanitizeAndValidateImageTopics(topics)

	if len(validTopics) > 25 {
		return nil, "repo.topic.count_prompt"

	}

	if len(invalidTopics) > 0 {
		return nil, "repo.imagetopic.format_prompt"

	}
	return validTopics, ""
}

func CloudBrainStop(ctx *context.Context) {
	var ID = ctx.Params(":id")
	var resultCode = "0"
	var errorMsg = ""
	var status = ""

	task := ctx.Cloudbrain

	for {
		if task.IsNewAITask() {
			t, bizErr := ai_task.StopCloudbrain(task)
			if bizErr != nil {
				resultCode = "-1"
				errorMsg = bizErr.TrCode
				resultCode = task.Status
				break
			}
			status = t.Status
			break
		}

		if task.Status == string(models.JobStopped) || task.Status == string(models.JobFailed) || task.Status == string(models.JobSucceeded) {
			log.Error("the job(%s) has been stopped", task.JobName, ctx.Data["msgID"])
			resultCode = "-1"
			errorMsg = "cloudbrain.Already_stopped"
			resultCode = task.Status
			break
		}
		if res, isHandled, err := ai_task.HandleNewAITaskStop(task.ID); isHandled {
			if err != nil {
				log.Error("StopJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
				resultCode = "-1"
				errorMsg = "cloudbrain.Stopped_failed"
				break
			}
			status = res.Status
			break
		}

		err := cloudbrain.StopJob(task.JobID)
		if err != nil {
			log.Error("StopJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
			resultCode = "-1"
			errorMsg = "cloudbrain.Stopped_failed"
			break
		}
		oldStatus := task.Status
		task.Status = string(models.JobStopped)
		if task.EndTime == 0 {
			task.EndTime = timeutil.TimeStampNow()
		}
		task.ComputeAndSetDuration()
		if oldStatus != task.Status {
			notification.NotifyChangeCloudbrainStatus(task, oldStatus)
		}
		err = models.UpdateJob(task)
		if err != nil {
			log.Error("UpdateJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
			resultCode = "-1"
			errorMsg = "cloudbrain.Stopped_success_update_status_fail"
			break
		}
		status = task.Status
		break
	}

	ctx.JSON(200, map[string]interface{}{
		"result_code": resultCode,
		"error_msg":   ctx.Tr(errorMsg),
		"status":      status,
		"id":          ID,
		"StatusOK":    0,
	})
}

func StopJobsByUserID(userID int64) {
	cloudBrains, err := models.GetCloudbrainsNeededStopByUserID(userID)
	if err != nil {
		log.Warn("Failed to get cloudBrain info", err)
		return
	}
	StopJobs(cloudBrains)

}

func StopJobsByRepoID(repoID int64) {
	cloudBrains, err := models.GetCloudbrainsNeededStopByRepoID(repoID)
	if err != nil {
		log.Warn("Failed to get cloudBrain info", err)
		return
	}
	StopJobs(cloudBrains)
}

func DeleteJobsByRepoID(repoID int64) {
	cloudBrains, err := models.GetCloudbrainsNeededDeleteByRepoID(repoID)
	if err != nil {
		log.Warn("Failed to get cloudBrain info", err)
		return
	}
	ai_task.DelCloudbrains(cloudBrains)
}

/*
*
 */
func StopJobs(cloudBrains []*models.Cloudbrain) {

	newStatus := string(models.JobStopped)
	for _, taskInfo := range cloudBrains {

		if taskInfo.Type == models.TypeCloudBrainOne {
			err := retry(3, time.Second*30, func() error {
				return cloudbrain.StopJob(taskInfo.JobID)
			})

			logErrorAndUpdateJobStatus(err, taskInfo, newStatus)
		} else if taskInfo.Type == models.TypeCloudBrainTwo {
			if taskInfo.JobType == string(models.JobTypeTrain) || taskInfo.JobType == string(models.JobTypeInference) {
				err := retry(3, time.Second*30, func() error {
					_, err := modelarts.StopTrainJob(taskInfo.JobID, strconv.FormatInt(taskInfo.VersionID, 10))
					return err
				})
				logErrorAndUpdateJobStatus(err, taskInfo, string(models.ModelArtsTrainJobKilling))
			} else {
				param := models.NotebookAction{
					Action: models.ActionStop,
				}
				var res *models.NotebookActionResult
				err := retry(3, time.Second*30, func() error {
					var tmpErr error
					res, tmpErr = modelarts.ManageNotebook2(taskInfo.JobID, param)
					return tmpErr
				})
				if res != nil {
					newStatus = res.Status
				}
				logErrorAndUpdateJobStatus(err, taskInfo, newStatus)
			}
		} else if taskInfo.Type == models.TypeC2Net {
			var res *models.GrampusStopJobResponse
			err := retry(3, time.Second*30, func() error {
				var err error
				res, err = grampus.StopJob(taskInfo.JobID, taskInfo.JobType)
				return err
			})
			if res != nil {
				newStatus = cloudbrainTask.GetStopJobResponseStatus(res)
			}
			logErrorAndUpdateJobStatus(err, taskInfo, newStatus)
		}
	}
}

func DeleteJobs(cloudBrains []*models.Cloudbrain) {
	for _, taskInfo := range cloudBrains {
		if taskInfo.Type == models.TypeCloudBrainOne {
			cloudbrain.DelCloudBrainJob(taskInfo.JobName)
			cloudbrainTask.DeleteCloudbrainJobStorage(taskInfo.JobName, models.TypeCloudBrainOne)
		}
		if taskInfo.Type == models.TypeCloudBrainTwo {
			if taskInfo.JobType == string(models.JobTypeTrain) || taskInfo.JobType == string(models.JobTypeInference) {

				_, err := modelarts.DelTrainJob(taskInfo.JobID)
				if err != nil {
					log.Error("Failed to delete cloudbrain job on modelarts.")
					continue
				}
				DeleteJobStorage(taskInfo.JobName)
			}
			if taskInfo.JobType == string(models.JobTypeDebug) {
				modelarts.DelNotebook2(taskInfo.JobID)
			}
		}
		if taskInfo.Type == models.TypeC2Net {
			if taskInfo.JobType == string(models.JobTypeTrain) {
				cloudbrain.DelCloudBrainJob(taskInfo.JobName)
				cloudbrainTask.DeleteCloudbrainJobStorage(taskInfo.JobName, models.TypeCloudBrainOne)
			}
		}
		err := models.DeleteJob(taskInfo)
		if err != nil {
			log.Warn("Failed to DeleteJob:", err)
			return
		}
	}
}

func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Warn("retrying after error:", err)
			time.Sleep(sleep)
		}
		err = f()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func logErrorAndUpdateJobStatus(err error, taskInfo *models.Cloudbrain, newStatus string) {
	if err != nil {
		log.Warn("Failed to stop cloudBrain job:"+taskInfo.JobID, err)
	} else {
		oldStatus := taskInfo.Status
		taskInfo.Status = newStatus
		if taskInfo.EndTime == 0 {
			taskInfo.EndTime = timeutil.TimeStampNow()
		}
		taskInfo.ComputeAndSetDuration()
		if oldStatus != taskInfo.Status {
			notification.NotifyChangeCloudbrainStatus(taskInfo, oldStatus)
		}
		err = models.UpdateJob(taskInfo)
		if err != nil {
			log.Warn("UpdateJob failed", err)
		}
	}
}

func CloudBrainDel(ctx *context.Context) {
	var listType = ctx.Query("debugListType")
	if err := deleteCloudbrainJob(ctx); err != nil {
		log.Error("deleteCloudbrainJob failed: %v", err, ctx.Data["msgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/debugjob?debugListType=" + listType)
	}
}

func deleteCloudbrainJob(ctx *context.Context) error {
	task := ctx.Cloudbrain

	if isHandled, err := ai_task.HandleNewAITaskDelete(task.ID); isHandled {
		if err != nil {
			log.Error("DeleteJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
			return err
		}
		return nil
	}

	if task.Status != string(models.JobStopped) && task.Status != string(models.JobFailed) && task.Status != string(models.JobSucceeded) {
		log.Error("the job(%s) has not been stopped", task.JobName, ctx.Data["msgID"])
		return errors.New("the job has not been stopped")
	}

	err := models.DeleteJob(task)
	if err != nil {
		log.Error("DeleteJob failed: %v", err, ctx.Data["msgID"])
		return err
	}

	cloudbrainTask.DeleteCloudbrainJobStorage(task.JobName, models.TypeCloudBrainOne)

	return nil
}

func CloudBrainShowModels(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true

	ID := ctx.Params(":id")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")
	task, err := models.GetCloudbrainByID(ID)
	if err != nil {
		log.Error("no such job!", ctx.Data["msgID"])
		ctx.ServerError("no such job:", err)
		return
	}

	//get dirs
	dirs, err := GetModelDirs(task.JobName, parentDir)
	if err != nil {
		log.Error("GetModelDirs failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("GetModelDirs failed:", err)
		return
	}

	var fileInfos []storage.FileInfo
	err = json.Unmarshal([]byte(dirs), &fileInfos)
	if err != nil {
		log.Error("json.Unmarshal failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("json.Unmarshal failed:", err)
		return
	}

	for i, fileInfo := range fileInfos {
		temp, _ := time.Parse("2006-01-02 15:04:05", fileInfo.ModTime)
		fileInfos[i].ModTime = temp.Local().Format("2006-01-02 15:04:05")
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime > fileInfos[j].ModTime
	})

	ctx.Data["Path"] = dirArray
	ctx.Data["Dirs"] = fileInfos
	ctx.Data["task"] = task
	ctx.Data["ID"] = ID
	ctx.HTML(200, tplCloudBrainShowModels)
}

func GetImages(ctx *context.Context, opts *models.SearchImageOptions) {
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = 15
	}
	opts.ListOptions = models.ListOptions{
		Page:     page,
		PageSize: pageSize,
	}
	var orderBy models.SearchOrderBy
	switch ctx.Query("sort") {
	case "newest":
		orderBy = models.SearchOrderByIDReverse
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse
	case "mostused":
		orderBy = models.SearchOrderByUseCountReverse
	default:
		orderBy = "type desc, num_stars desc,id desc"
	}

	opts.SearchOrderBy = orderBy
	opts.Framework = ctx.Query("framework")
	opts.FrameworkVersion = ctx.Query("frameworkVersion")
	opts.CudaVersion = ctx.Query("cuda")
	opts.CannVersion = ctx.Query("cann")
	opts.DTKVersion = ctx.Query("dtk")
	opts.PythonVersion = ctx.Query("python")
	opts.OperationSystem = ctx.Query("os")
	opts.OperationSystemVersion = ctx.Query("osVersion")
	opts.ThirdPackages = ctx.Query("thirdParty")
	opts.OnlyOpenIImage = ctx.QueryBool("onlyOpenIImage")
	opts.TrainType = ctx.Query("trainType")
	opts.TrainTypeNum = models.ParseTrainType(opts.TrainType)
	jobType := ctx.Query("jobType")
	specId := ctx.QueryInt64("spec")
	hasInternet := ctx.QueryInt("has_internet")
	visualizeRequired := ctx.QueryBool("visualize_required")

	var (
		queues          = make([]models.ResourceQueue, 0)
		imageListResult = make([]*models.ImageResult, 0)
	)
	if specId > 0 {
		spec, _ := models.GetSpecificationById(specId)
		if spec != nil {
			queues = spec.GetAvailableQueues(models.GetAvailableCenterIdOpts{
				JobType:           models.JobType(jobType),
				UserId:            opts.UID,
				HasInternet:       models.SpecInternetQuery(hasInternet),
				VisualizeRequired: visualizeRequired,
			})
			if len(queues) > 0 {
				var aiCenterCodes []string
				for i := 0; i < len(queues); i++ {
					q := queues[i]
					aiCenterCodes = append(aiCenterCodes, q.AiCenterCode)
				}
				opts.AiCenterIds = aiCenterCodes
				opts.AccCardType = spec.AccCardType
			}
		}
	}
	computeResource := ctx.Query("computeResource")
	if computeResource != "" {
		opts.ComputeResource = computeResource
	}

	imageList, total, err := models.SearchImage(opts)
	if err != nil {
		log.Error("Can not get images:%v", err)
		ctx.JSON(http.StatusOK, models.ImagesPageResult{
			Count:  0,
			Images: []*models.ImageResult{},
		})
		return
	}
	if len(queues) > 0 {
		newImageList := models.ImageList{}
		for _, tmp := range imageList {
			if hasIntersection(tmp.AiCenterImages, opts.AccCardType, queues) {
				newImageList = append(newImageList, tmp)
			}
		}
		imageList = newImageList
	}

	for _, image := range imageList {
		result := &models.ImageResult{Image: image}

		// 填充分中心名称
		if len(result.AiCenterImages) > 0 {
			var names []string
			for _, img := range result.AiCenterImages {
				if name := GetAiCenterNameByCode(img.AiCenterId, ctx.Language()); name != "" {
					names = append(names, name)
				}
			}
			result.AiCenterName = strings.Join(names, "/")
		}

		imageListResult = append(imageListResult, result)
	}

	for _, tmp := range imageListResult {
		if tmp.Place == "" {
			aiCenterId := ""
			if len(opts.AiCenterIds) > 0 {
				aiCenterId = opts.AiCenterIds[0]
			}
			tmp.Place = getImageUrl(tmp.Image, opts.TrainType, aiCenterId)
		}

		// 如果为空且状态没有返回的成功话,为空则traintype不展示
		if tmp.Status != 1 {
			tmp.TrainType = models.TrainTypeToString(-1)
			continue
		}
		// 否则全部展示
		if tmp.TrainType == "" {
			tmp.TrainType = models.TrainTypeToString(0)
		}
	}
	ctx.JSON(http.StatusOK, models.ImagesPageResult{
		Count:  total,
		Images: imageListResult,
	})
}

func hasIntersection(imageCenterInfos []models.AiCenterImage, accCardType string, queues []models.ResourceQueue) bool {
	if len(queues) == 0 || len(imageCenterInfos) == 0 {
		//如果没传queues或者查询的镜像不含可用中心信息，不进行判断，直接返回true
		return true
	}
	for _, aicenterImage := range imageCenterInfos {
		for _, queue := range queues {
			if aicenterImage.AiCenterId == queue.AiCenterCode {
				if len(accCardType) == 0 || len(aicenterImage.AccDeviceModel) == 0 {
					return true
				}
				if len(accCardType) > 0 && len(aicenterImage.AccDeviceModel) > 0 {
					if strings.ToUpper(accCardType) != strings.ToUpper(aicenterImage.AccDeviceModel) {
						return false
					}
				}
				if len(aicenterImage.PoolIds) == 0 {
					return true
				}
				for _, poolId := range aicenterImage.PoolIds {
					if queue.QueueCode == poolId {
						return true
					}
				}
			}
		}
	}
	return false
}

func getImageUrl(image *models.Image, trainType, aiCenterID string) string {
	if image.AiCenterImages != nil {
		for _, tmp := range image.AiCenterImages {
			if aiCenterID == "" {
				return models.GetImageUrlByTrainType(trainType, tmp.NotebookUrl, tmp.TrainJobUrl)
			}
			if tmp.AiCenterId == aiCenterID {
				log.Info(fmt.Sprintf("trainType=[%s] tmp.NotebookUrl=[%s] tmp.TrainJobUrl=[%s]", trainType, tmp.NotebookUrl, tmp.TrainJobUrl))
				return models.GetImageUrlByTrainType(trainType, tmp.NotebookUrl, tmp.TrainJobUrl)
			}
		}
	}
	return ""
}

func getUID(ctx *context.Context) int64 {
	var uid int64 = -1
	if ctx.IsSigned {
		uid = ctx.User.ID
	}
	return uid
}

func GetAllImages(ctx *context.Context) {
	uid := getUID(ctx)

	// 解析智算中心参数，支持多个中心代码，使用 || 分隔（与 queue/list 接口保持一致）
	var aiCenterIds []string
	aiCenterCode := ctx.Query("center")
	if aiCenterCode != "" {
		aiCenterIds = strings.Split(aiCenterCode, "||")
	}

	opts := models.SearchImageOptions{
		UID:                    uid,
		Keyword:                ctx.Query("q"),
		ApplyStatus:            ctx.QueryInt("apply"),
		Topics:                 ctx.Query("topic"),
		IncludeOfficialOnly:    ctx.QueryBool("recommend"),
		CloudbrainType:         ctx.QueryInt("cloudbrainType"),
		Status:                 -1,
		Framework:              ctx.Query("framework"),
		FrameworkVersion:       ctx.Query("frameworkVesion"),
		CudaVersion:            ctx.Query("cuda"),
		PythonVersion:          ctx.Query("python"),
		OperationSystem:        ctx.Query("os"),
		OperationSystemVersion: ctx.Query("osVersion"),
		ThirdPackages:          ctx.Query("thirdParty"),
		TrainType:              ctx.Query("trainType"),
		ComputeResource:        ctx.Query("computeResource"),
		AccCardType:            ctx.Query("accCardType"),
		AiCenterIds:            aiCenterIds,
		OnlyAccCardType:        true, // 精确查询
	}

	if ctx.Query("private") != "" {
		if ctx.QueryBool("private") {
			opts.IncludePrivateOnly = true
		} else {
			opts.IncludePublicOnly = true
		}
	}
	GetImages(ctx, &opts)

}

func GetModelDirs(jobName string, parentDir string) (string, error) {
	var req string
	modelActualPath := storage.GetMinioPath(jobName, cloudbrain.ModelMountPath+"/")
	if parentDir == "" {
		req = "baseDir=" + modelActualPath
	} else {
		req = "baseDir=" + modelActualPath + "&parentDir=" + parentDir
	}

	return getDirs(req)
}

func GetResultDirs(jobName string, parentDir string) (string, error) {
	var req string
	modelActualPath := storage.GetMinioPath(jobName, cloudbrain.ResultPath+"/")
	if parentDir == "" {
		req = "baseDir=" + modelActualPath
	} else {
		req = "baseDir=" + modelActualPath + "&parentDir=" + parentDir
	}

	return getDirs(req)
}

func CloudBrainDownloadModel(ctx *context.Context) {
	parentDir := ctx.Query("parentDir")
	fileName := ctx.Query("fileName")
	jobName := ctx.Query("jobName")
	filePath := "jobs/" + jobName + "/model/" + parentDir
	url, err := storage.Attachments.PresignedGetURL(filePath, fileName)
	if err != nil {
		log.Error("PresignedGetURL failed: %v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("PresignedGetURL", err)
		return
	}
	ctx.Resp.Header().Set("Cache-Control", "max-age=0")
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
}

func CloudBrainDownloadMultiModel(ctx *context.Context) {
	parentDir := ctx.Query("parentDir")
	jobName := ctx.Query("jobName")
	filePath := "jobs/" + jobName + "/model/" + parentDir
	allFile, err := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, filePath)
	if err == nil {
		returnFileName := jobName + ".zip"
		MinioDownloadManyFile(filePath, ctx, returnFileName, allFile)
	} else {
		log.Info("error,msg=" + err.Error())
		ctx.ServerError("no file to download.", err)
	}
}

func CloudBrainDownloadInferenceResult(ctx *context.Context) {
	parentDir := ctx.Query("parentDir")
	fileName := ctx.Query("fileName")
	jobName := ctx.Query("jobName")
	filePath := "jobs/" + jobName + "/result/" + parentDir
	url, err := storage.Attachments.PresignedGetURL(filePath, fileName)
	if err != nil {
		log.Error("PresignedGetURL failed: %v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("PresignedGetURL", err)
		return
	}
	ctx.Resp.Header().Set("Cache-Control", "max-age=0")
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
}

func GetRate(ctx *context.Context) {
	isObjectDetcionAll := ctx.QueryBool("isObjectDetcionAll")
	if isObjectDetcionAll {
		ctx.Redirect(setting.BenchmarkServerHost + "?username=admin")
		return
	}

	var ID = ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		ctx.ServerError("GetCloudbrainByJobID failed", err)
		return
	}

	if job.JobType == string(models.JobTypeBenchmark) {
		log.Info("url=" + setting.BenchmarkServerHost + "?username=" + ctx.User.Name)
		ctx.Redirect(setting.BenchmarkServerHost + "?username=" + ctx.User.Name)
	} else if job.JobType == string(models.JobTypeSnn4imagenet) {
		ctx.Redirect(setting.Snn4imagenetServerHost)
	} else if job.JobType == string(models.JobTypeBrainScore) {
		ctx.Redirect(setting.BrainScoreServerHost)
	} else if job.JobType == string(models.JobTypeSnn4Ecoset) {
		ctx.Redirect(setting.Snn4EcosetServerHost)
	} else if job.JobType == string(models.JobTypeSim2BrainSNN) {
		ctx.Redirect(setting.Sim2BrainSnnServerHost)
	} else {
		log.Error("JobType error:%s", job.JobType, ctx.Data["msgID"])
	}
}

func downloadCode(repo *models.Repository, codePath, branchName string) error {
	//add "file:///" prefix to make the depth valid
	if err := git.Clone(CLONE_FILE_PREFIX+repo.RepoPath(), codePath, git.CloneRepoOptions{Branch: branchName, Depth: 1}); err != nil {
		log.Error("Failed to clone repository: %s (%v)", repo.FullName(), err)
		return err
	}

	configFile, err := os.OpenFile(codePath+"/.git/config", os.O_RDWR, 0666)
	if err != nil {
		log.Error("open file(%s) failed:%v", codePath+"/,git/config", err)
		return err
	}

	defer configFile.Close()

	pos := int64(0)
	reader := bufio.NewReader(configFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Error("not find the remote-url")
				return nil
			} else {
				log.Error("read error: %v", err)
				return err
			}
		}

		if strings.Contains(line, "url") && strings.Contains(line, ".git") {
			originUrl := "\turl = " + repo.CloneLink().HTTPS + "\n"
			if len(line) > len(originUrl) {
				originUrl += strings.Repeat(" ", len(line)-len(originUrl))
			}
			bytes := []byte(originUrl)
			_, err := configFile.WriteAt(bytes, pos)
			if err != nil {
				log.Error("WriteAt failed:%v", err)
				return err
			}
			break
		}

		pos += int64(len(line))
	}

	return nil
}

func downloadRateCode(repo *models.Repository, taskName, rateOwnerName, rateRepoName, codePath, benchmarkCategory, gpuType, userName string) error {
	err := os.MkdirAll(codePath, os.ModePerm)
	if err != nil {
		log.Error("mkdir codePath failed", err.Error())
		return err
	}

	repoExt, err := models.GetRepositoryByOwnerAndName(rateOwnerName, rateRepoName)
	if err != nil {
		log.Error("GetRepositoryByOwnerAndName(%s) failed", rateRepoName, err.Error())
		return err
	}

	if err := git.Clone(CLONE_FILE_PREFIX+repoExt.RepoPath(), codePath, git.CloneRepoOptions{Depth: 1}); err != nil {
		log.Error("Failed to clone repository: %s (%v)", repoExt.FullName(), err)
		return err
	}

	fileName := codePath + cloudbrain.TaskInfoName
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Error("OpenFile failed", err.Error())
		return err
	}

	defer f.Close()

	data, err := json.Marshal(models.TaskInfo{
		Username:          userName,
		TaskName:          taskName,
		CodeName:          repo.Name,
		BenchmarkCategory: strings.Split(benchmarkCategory, ","),
		CodeLink:          strings.TrimSuffix(repo.CloneLink().HTTPS, ".git"),
		GpuType:           gpuType,
	})
	if err != nil {
		log.Error("json.Marshal failed", err.Error())
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		log.Error("WriteString failed", err.Error())
		return err
	}

	return nil
}

func uploadCodeToMinio(codePath, jobName, parentDir string) error {
	files, err := readDir(codePath)
	if err != nil {
		log.Error("readDir(%s) failed: %s", codePath, err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err = uploadCodeToMinio(codePath+file.Name()+"/", jobName, parentDir+file.Name()+"/"); err != nil {
				log.Error("uploadCodeToMinio(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		} else {
			destObject := setting.CBCodePathPrefix + jobName + parentDir + file.Name()
			sourceFile := codePath + file.Name()
			err = storage.Attachments.UploadObject(destObject, sourceFile)
			if err != nil {
				log.Error("UploadObject(%s) failed: %s", file.Name(), err.Error())
				if strings.Contains(err.Error(), "no such file or directory") {
					continue
				}
				return err
			}
		}
	}

	return nil
}

func mkModelPath(modelPath string) error {
	return mkPathAndReadMeFile(modelPath, "You can put the files into this directory and download the files by the web page.")
}

func mkPathAndReadMeFile(path string, text string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Error("MkdirAll(%s) failed:%v", path, err)
		return err
	}

	fileName := path + README
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Error("OpenFile failed", err.Error())
		return err
	}

	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		log.Error("WriteString failed", err.Error())
		return err
	}

	return nil
}

func SyncCloudbrainStatus() {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()

	cloudBrains, err := models.GetCloudBrainUnStoppedJob()
	if err != nil {
		log.Error("GetCloudBrainUnStoppedJob failed:", err.Error())
		return
	}

	log.Info("cloudbrain length is %v ", len(cloudBrains))

	batchSyncTask := make(map[string]*models.Cloudbrain)
	batchSyncTaskIds := []string{}
	batch_size := 20

	for i, task := range cloudBrains {

		if task.JobType == string(models.JobTypeModelSafety) {
			continue
		}

		if task.JobType == string(models.JobTypeTrain) || task.JobType == string(models.JobTypeDebug) {
			batchSyncTask[task.JobID] = task
			batchSyncTaskIds = append(batchSyncTaskIds, task.JobID)

		} else {

			syncOnetaskStatus(task)

		}

		if len(batchSyncTaskIds) >= batch_size || (i == len(cloudBrains)-1 && len(batchSyncTaskIds) > 0) {

			res, err := grampus_client.GetJobs(batchSyncTaskIds)
			if err != nil {
				log.Error("GetJobs failed:", err.Error())
				batchSyncTask = make(map[string]*models.Cloudbrain)
				batchSyncTaskIds = []string{}
				continue
			}

			for _, jobInfo := range res.JobInfos {
				if tempTask, ok := batchSyncTask[jobInfo.JobID]; ok {

					syncOnetaskStatus(tempTask, entity.ConvertGrampusTrainResponse(jobInfo))

				}

			}
			batchSyncTask = make(map[string]*models.Cloudbrain)
			batchSyncTaskIds = []string{}

		}

	}

}

func syncOnetaskStatus(task *models.Cloudbrain, res ...*entity.QueryTaskResponse) {
	maxDuration := setting.MaxDuration
	if role.UserHasOper(task.UserID, role.ROLE_OPER_DEBUG_TIME) && task.TimeLimit != 0 {
		//if role.UserHasRole(task.UserID, models.Subscriber) && task.TimeLimit != 0 {
		if task.TimeLimit > 0 {
			maxDuration = int64(task.TimeLimit * 60 * 60)
		} else {
			maxDuration = math.MaxInt64
		}
	}

	task, _ = ai_task.UpdateCloudbrain(task, res...)
	if task.Duration >= maxDuration && task.JobType == string(models.JobTypeDebug) || task.Duration >= setting.Grampus.MMLSparkMaxTime && task.JobType == string(models.JobTypeSuperCompute) {
		if task.Status != string(models.ModelArtsStopping) {
			ai_task.StopCloudbrain(task)
		}
	}
	ai_task.TryToNotifyDebugAlmostEndingTask(task, maxDuration)
	go ai_task.TryToNotifyLongRunningTask(task)
	go ai_task.TryToStopFailedPullImage(task)

}

// SyncCommitImageStatus 镜像提交后，同步镜像的状态等信息
func SyncCommitImageStatus() {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()

	images, err := models.GetCommittingImages()
	if err != nil {
		log.Error("GetCommiting Image failed:", err.Error())
		return
	}

	for _, image := range images {

		if image.CloudbrainType == models.TypeC2Net {
			result, err := grampus.GetImageStatus(image)
			if err == nil {
				if result.ErrorCode != 0 {
					log.Warn("commit grampus image err:%d  %s", result.ErrorCode, result.ErrorMsg)
					updateFailedWhenTimeout(image)
					continue
				}

				if result.Image.ImageStatus == grampus.GrampusImageCommitMade {
					image.Status = models.IMAGE_STATUS_SUCCESS
					image.Place = result.Image.ImageFullAddr
					if image.AiCenterImages != nil {
						image.AiCenterImages[0].NotebookUrl = result.Image.NotebookUrl
						image.AiCenterImages[0].TrainJobUrl = result.Image.TrainJobUrl
						image.AiCenterImages[0].AccDeviceModel = result.Image.AccDeviceModel
					}

					trainType, trainNum := models.ConvertTrainType(result.Image.TrainType)
					image.TrainType = trainType
					image.TrainTypeNum = trainNum

					models.UpdateLocalImageStatusAndPlace(image)
				}

				if result.Image.ImageStatus == grampus.GrampusImageCommitFailed {
					image.Status = models.IMAGE_STATUS_Failed
					models.UpdateLocalImageStatusAndPlace(image)
				}

				if image.Status == models.IMAGE_STATUS_COMMIT && time.Now().Unix()-int64(image.CreatedUnix) > 3600 {
					image.Status = models.IMAGE_STATUS_Failed
					models.UpdateLocalImageStatusAndPlace(image)

				}

			} else {
				updateFailedWhenTimeout(image)
			}

		}

	}
}

func updateFailedWhenTimeout(image *models.Image) {
	if time.Now().Unix()-int64(image.CreatedUnix) > 3600 {
		image.Status = models.IMAGE_STATUS_Failed
		models.UpdateLocalImageStatusAndPlace(image)
	}
}

func HandleTaskWithNoDuration(ctx *context.Context) {
	mode := ctx.Query("mode")
	log.Info("HandleTaskWithNoDuration start")
	count := 0
	start := time.Now().Unix()
	for {
		var cloudBrains []*models.Cloudbrain
		var err error
		if mode == "1" {
			cloudBrains, err = models.GetStoppedJobWithNoStartTimeEndTime()
		} else {
			cloudBrains, err = models.GetStoppedJobWithNoDurationJob()
		}

		if err != nil {
			log.Error("HandleTaskWithNoTrainJobDuration failed:", err.Error())
			break
		}
		if len(cloudBrains) == 0 {
			log.Info("HandleTaskWithNoTrainJobDuration:no task need handle")
			break
		}
		handleNoDurationTask(cloudBrains)
		count += len(cloudBrains)
		if len(cloudBrains) < 100 {
			log.Info("HandleTaskWithNoTrainJobDuration:task less than 100")
			break
		}
		if time.Now().Unix()-start > 600 {
			log.Info("HandleTaskWithNoDuration : time out")
			ctx.JSON(200, fmt.Sprintf("task stop for time out,count=%d", count))
			return
		}
	}
	log.Info("HandleTaskWithNoTrainJobDuration:count=%d", count)
	ctx.JSON(200, fmt.Sprintf("success,count=%d", count))
}

func handleNoDurationTask(cloudBrains []*models.Cloudbrain) {
	for _, task := range cloudBrains {
		time.Sleep(time.Millisecond * 100)
		log.Info("Handle job ,%+v", task)
		if task.Type == models.TypeCloudBrainOne {
			result, err := cloudbrain.GetJob(task.JobID)
			if err != nil {
				log.Error("GetJob(%s) failed:%v", task.JobName, err)
				updateDefaultDuration(task)
				continue
			}

			if result != nil {
				if result.Msg != "success" {
					updateDefaultDuration(task)
					continue
				}
				jobRes, err := models.ConvertToJobResultPayload(result.Payload)
				if err != nil || len(jobRes.TaskRoles) == 0 {
					updateDefaultDuration(task)
					continue
				}
				taskRoles := jobRes.TaskRoles
				taskRes, err := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
				if err != nil || len(taskRes.TaskStatuses) == 0 {
					updateDefaultDuration(task)
					continue
				}
				task.Status = taskRes.TaskStatuses[0].State
				log.Info("task startTime = %v  endTime= %v ,jobId=%d", jobRes.JobStatus.StartTime, jobRes.JobStatus.EndTime, task.ID)
				if jobRes.JobStatus.CreatedTime > 0 {
					task.StartTime = timeutil.TimeStamp(jobRes.JobStatus.CreatedTime / 1000)
					if jobRes.JobStatus.CompletedTime > 0 {
						task.EndTime = timeutil.TimeStamp(jobRes.JobStatus.CompletedTime / 1000)
					} else {
						task.EndTime = task.UpdatedUnix
					}
				} else {
					task.StartTime = 0
					task.EndTime = 0
				}

				if task.EndTime < task.StartTime {
					log.Info("endTime[%v] is less than starTime[%v],jobId=%d", task.EndTime, task.StartTime, task.ID)
					st := task.StartTime
					task.StartTime = task.EndTime
					task.EndTime = st
				}
				task.Duration = task.EndTime.AsTime().Unix() - task.StartTime.AsTime().Unix()
				task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)
				err = models.UpdateJobDurationWithDeleted(task)
				if err != nil {
					log.Error("UpdateJob(%s) failed:%v", task.JobName, err)
				}
			}
		} else if task.Type == models.TypeCloudBrainTwo {
			if task.JobType == string(models.JobTypeDebug) {
				//result, err := modelarts.GetJob(task.JobID)
				result, err := modelarts.GetNotebook2(task.JobID)
				if err != nil {
					log.Error("GetJob(%s) failed:%v", task.JobName, err)
					updateDefaultDuration(task)
					continue
				}

				if result != nil {
					task.Status = result.Status
					startTime := result.Lease.CreateTime
					duration := result.Lease.Duration / 1000
					if startTime > 0 {
						task.StartTime = timeutil.TimeStamp(startTime / 1000)
						task.EndTime = task.StartTime.Add(duration)
					}
					task.CorrectCreateUnix()
					task.ComputeAndSetDuration()
					err = models.UpdateJobDurationWithDeleted(task)
					if err != nil {
						log.Error("UpdateJob(%s) failed:%v", task.JobName, err)
						continue
					}
				}
			} else if task.JobType == string(models.JobTypeTrain) || task.JobType == string(models.JobTypeInference) {
				result, err := modelarts.GetTrainJob(task.JobID, strconv.FormatInt(task.VersionID, 10))
				if err != nil {
					log.Error("GetTrainJob(%s) failed:%v", task.JobName, err)
					updateDefaultDuration(task)
					continue
				}

				if result != nil {
					startTime := result.StartTime / 1000
					if startTime > 0 {
						task.StartTime = timeutil.TimeStamp(startTime)
						task.EndTime = task.StartTime.Add(result.Duration / 1000)
					}
					task.ComputeAndSetDuration()
					err = models.UpdateJobDurationWithDeleted(task)
					if err != nil {
						log.Error("UpdateJob(%s) failed:%v", task.JobName, err)
						continue
					}
				}
			} else {
				log.Error("task.JobType(%s) is error:%s", task.JobName, task.JobType)
			}

		} else {
			log.Error("task.Type(%s) is error:%d", task.JobName, task.Type)
		}
	}
}

func updateDefaultDuration(task *models.Cloudbrain) {
	log.Info("updateDefaultDuration: taskId=%d", task.ID)
	task.StartTime = task.CreatedUnix
	task.EndTime = task.UpdatedUnix
	task.ComputeAndSetDuration()
	err := models.UpdateJobDurationWithDeleted(task)
	if err != nil {
		log.Error("UpdateJob(%s) failed:%v", task.JobName, err)
	}
}

func CloudBrainBenchmarkIndex(ctx *context.Context) {
	MustEnableCloudbrain(ctx)
	repo := ctx.Repo.Repository
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	jobTypes := models.AllBenchMarkJobType()
	ciTasks, count, err := models.Cloudbrains(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: setting.UI.IssuePagingNum,
		},
		RepoID:   repo.ID,
		Type:     -1,
		JobTypes: jobTypes,
	})
	if err != nil {
		ctx.ServerError("Get debugjob faild:", err)
		return
	}

	for i, task := range ciTasks {
		ciTasks[i].CanDel = cloudbrain.CanDeleteJob(ctx, &task.Cloudbrain)
		ciTasks[i].Cloudbrain.ComputeResource = task.ComputeResource
		if ciTasks[i].TrainJobDuration == "" {
			if ciTasks[i].Duration == 0 {
				var duration int64
				if task.Status == string(models.JobRunning) {
					duration = time.Now().Unix() - int64(task.Cloudbrain.CreatedUnix)
				} else {
					duration = int64(task.Cloudbrain.UpdatedUnix) - int64(task.Cloudbrain.CreatedUnix)
				}
				ciTasks[i].Duration = duration
			}
			ciTasks[i].TrainJobDuration = models.ConvertDurationToStr(ciTasks[i].Duration)
		}
		ciTasks[i].BenchmarkTypeName = ""
		if ciTasks[i].JobType == string(models.JobTypeBenchmark) {
			ciTasks[i].BenchmarkType = ctx.Tr("repo.cloudbrain.benchmark.algorithm")
		} else if models.IsBenchMarkJobType(ciTasks[i].JobType) {
			ciTasks[i].BenchmarkType = ctx.Tr("repo.cloudbrain.benchmark.model")
			ciTasks[i].BenchmarkTypeName = ciTasks[i].JobType

			if ciTasks[i].JobType == string(models.JobTypeSnn4imagenet) {
				ciTasks[i].BenchmarkTypeRankLink = setting.Snn4imagenetServerHost
			} else if ciTasks[i].JobType == string(models.JobTypeBrainScore) {
				ciTasks[i].BenchmarkTypeRankLink = setting.BrainScoreServerHost
			} else {
				ciTasks[i].BenchmarkTypeRankLink = setting.Snn4EcosetServerHost
			}

		}

		if task.BenchmarkTypeID > 0 {
			for _, benchmarkType := range GetBenchmarkTypes(ctx).BenchmarkType {
				if task.BenchmarkTypeID == benchmarkType.Id {
					ciTasks[i].BenchmarkTypeRankLink = benchmarkType.RankLink
					ciTasks[i].BenchmarkTypeName = benchmarkType.First
					break
				}
			}
		}
		if task.JobType == string(models.JobTypeModelSafety) {
			ciTasks[i].BenchmarkType = "安全评测"
			ciTasks[i].BenchmarkTypeName = "Image Classification"
		}
	}

	pager := context.NewPagination(int(count), setting.UI.IssuePagingNum, page, 5)

	ctx.Data["Page"] = pager
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["Tasks"] = ciTasks
	ctx.Data["CanCreate"] = cloudbrain.CanCreateOrDebugJob(ctx)
	ctx.Data["RepoIsEmpty"] = repo.IsEmpty
	ctx.HTML(200, tplCloudBrainBenchmarkIndex)
}

func GetChildTypes(ctx *context.Context) {
	benchmarkTypeID := ctx.QueryInt("benchmark_type_id")
	re := make(map[string]interface{})
	for {
		var isExist bool
		for _, benchmarkType := range GetBenchmarkTypes(ctx).BenchmarkType {
			if benchmarkTypeID == benchmarkType.Id {
				isExist = true
				re["child_types"] = benchmarkType.Second
				re["result_code"] = "0"
				break
			}
		}
		if !isExist {
			re["result_code"] = "1"
			log.Error("no such benchmark_type_id", ctx.Data["MsgID"])
			re["errMsg"] = "system error"
			break
		}
		break
	}
	ctx.JSON(200, re)
}

func getBrainRegion(benchmarkChildTypeID int) string {
	return setting.BrainScoreRegion[benchmarkChildTypeID]
}
func getSim2BrainDatasetType(benchmarkChildTypeID int) string {
	return setting.Sim2BrainSnnDatasetTypes[benchmarkChildTypeID]
}

func trimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(s, " ")

}

func BenchmarkDel(ctx *context.Context) {
	if err := deleteCloudbrainJob(ctx); err != nil {
		log.Error("deleteCloudbrainJob failed: %v", err, ctx.Data["msgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/cloudbrain/benchmark")
	}
}

func CloudBrainTrainJobNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true
	cloudBrainTrainJobCreate(ctx)
}
func CloudBrainTrainJobVersionNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = false
	cloudBrainTrainJobCreate(ctx)
}

func cloudBrainTrainJobCreate(ctx *context.Context) {
	// err := cloudBrainNewDataPrepare(ctx, string(models.JobTypeTrain))
	// if err != nil {
	// 	ctx.ServerError("get new train-job info failed", err)
	// 	return
	// }
	ctx.HTML(http.StatusOK, tplCloudBrainTrainJobNew)
}

func InferenceCloudBrainJobNew(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplCloudBrainInferenceJobNew)
}

func InferenceCloudBrainJobShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplCloudBrainInferenceJobShow)
}

func DownloadGPUInferenceResultFile(ctx *context.Context) {
	var jobID = ctx.Params(":jobid")
	task, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		return
	}
	parentDir := ctx.Query("parentDir")
	filePath := "jobs/" + task.JobName + "/result/" + parentDir
	log.Info("prefix=" + filePath)
	allFile, err := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, filePath)
	if err == nil {
		returnFileName := task.DisplayJobName + ".zip"
		MinioDownloadManyFile(filePath, ctx, returnFileName, allFile)
	} else {
		log.Info("error,msg=" + err.Error())
		ctx.ServerError("no file to download.", err)
	}
}

func getInferenceJobCommand(form auth.CreateCloudBrainInferencForm) (string, error) {
	var command string
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params

	if !strings.HasSuffix(bootFile, ".py") {
		log.Error("bootFile(%s) format error", bootFile)
		return command, errors.New("bootFile format error")
	}

	var parameters models.Parameters
	var param string
	if len(params) != 0 {
		err := json.Unmarshal([]byte(params), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", params, err)
			return command, err
		}

		for _, parameter := range parameters.Parameter {
			param += " --" + parameter.Label + "=" + parameter.Value
		}
	}

	param += " --modelname" + "='" + form.CkptName + "'"

	command += "python /code/" + bootFile + param + " > " + cloudbrain.ResultPath + "/" + form.DisplayJobName + "-" + cloudbrain.LogFile

	return command, nil
}

func getTrainJobCommand(form auth.CreateCloudBrainForm) (string, error) {
	var command string
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params

	if !strings.HasSuffix(bootFile, ".py") {
		log.Error("bootFile(%s) format error", bootFile)
		return command, errors.New("bootFile format error")
	}

	var parameters models.Parameters
	var param string
	if len(params) != 0 {
		err := json.Unmarshal([]byte(params), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", params, err)
			return command, err
		}

		for _, parameter := range parameters.Parameter {
			param += " --" + parameter.Label + "=" + parameter.Value
		}
	}
	if form.CkptName != "" {
		param += " --ckpt_url" + "=" + "'/pretrainmodel/" + form.CkptName + "'"
	}

	command += "python /code/" + bootFile + param + " > " + cloudbrain.ModelMountPath + "/" + form.DisplayJobName + "-" + cloudbrain.LogFile

	return command, nil
}

func CloudBrainTrainJobDel(ctx *context.Context) {
	var listType = ctx.Query("listType")
	if err := deleteCloudbrainJob(ctx); err != nil {
		log.Error("deleteCloudbrainJob failed: %v", err, ctx.Data["msgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=" + listType)
	}
}

func GetBenchmarkTypes(ctx *context.Context) *models.BenchmarkTypes {
	var lang = ctx.Locale.Language()
	if benchmarkTypesMap[lang] == nil {
		var val = i18n.Tr(lang, BENCHMARK_TYPE_CODE)
		//use config
		val = setting.BenchmarkTypes
		var tempType *models.BenchmarkTypes
		if err := json.Unmarshal([]byte(val), &tempType); err != nil {
			log.Error("json.Unmarshal BenchmarkTypes(%s) failed:%v", val, err, ctx.Data["MsgID"])
			return &models.BenchmarkTypes{}
		}
		benchmarkTypesMap[lang] = tempType
	}
	return benchmarkTypesMap[lang]
}

func GetCloudbrainAiCenter(task models.Cloudbrain, ctx *context.Context) string {
	if task.Spec != nil {
		return task.Spec.AiCenterName
	} else {
		return ""
	}
}
func getCutStringAiCenterByAiCenter(aiCenter string) string {
	if aiCenter == "" {
		return ""
	}
	index := strings.LastIndex(aiCenter, "+")
	return aiCenter[index+1:]

}
func GetCloudbrainCluster(task models.Cloudbrain, ctx *context.Context) string {
	if task.Type == models.TypeCloudBrainOne || task.Type == models.TypeCloudBrainTwo || task.Type == models.TypeCDCenter {
		return ctx.Tr("cloudbrain.resource_cluster_openi")
	} else if task.Type == models.TypeC2Net {
		return ctx.Tr("cloudbrain.resource_cluster_c2net")
	}
	return ""
}
func GetCloudbrainCardDuration(task models.Cloudbrain) string {
	cardNum := int(0)
	spec := task.Spec
	cardNum = 1
	if spec == nil {
		specdb, err := models.GetSpecAccCardsNumByID(task.ID)
		if err != nil {
			log.Info("error:" + err.Error())
			return ""
		}
		if specdb != nil {
			cardNum = specdb.AccCardsNum
		}
	} else {
		cardNum = spec.AccCardsNum
	}
	var workServerNumber int64
	if task.WorkServerNumber >= 1 {
		workServerNumber = int64(task.WorkServerNumber)
	} else {
		workServerNumber = 1
	}
	cardDuration := models.ConvertDurationToStr(workServerNumber * int64(cardNum) * task.Duration)
	return cardDuration
}
func GetCloudbrainWaitTime(task models.Cloudbrain) string {
	var waitTime string
	if task.Status == string(models.JobWaiting) {
		waitTimeInt := time.Now().Unix() - task.CreatedUnix.AsTime().Unix()
		waitTime = models.ConvertDurationToStr(waitTimeInt)
		if waitTimeInt < 0 {
			waitTime = "00:00:00"
		}
	} else if task.Status == string(models.JobStopped) && task.StartTime.AsTime().Unix() == 0 {
		waitTimeInt := task.EndTime.AsTime().Unix() - task.CreatedUnix.AsTime().Unix()
		waitTime = models.ConvertDurationToStr(waitTimeInt)
		if waitTimeInt < 0 {
			waitTime = "00:00:00"

		}
	} else {
		waitTimeInt := task.StartTime.AsTime().Unix() - task.CreatedUnix.AsTime().Unix()
		waitTime = models.ConvertDurationToStr(waitTimeInt)
		if waitTimeInt < 0 {
			waitTime = "00:00:00"
		}
	}
	return waitTime
}
func GetCloudbrainFlavorName(task models.Cloudbrain) string {
	if task.Spec != nil {
		flavorName := task.Spec.ComputeResource + ":" + fmt.Sprint(task.Spec.AccCardsNum) + "*" + task.Spec.AccCardType +
			",内存:" + strconv.FormatInt(int64(task.Spec.MemGiB), 10) + "GB,共享内存:" + strconv.FormatInt(int64(task.Spec.ShareMemGiB), 10) + "GB"
		return flavorName
	} else {
		return ""
	}
}
