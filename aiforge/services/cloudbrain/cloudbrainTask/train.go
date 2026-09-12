package cloudbrainTask

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/timeutil"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/modules/obs"

	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/storage"
	"github.com/unknwon/com"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward/point/account"
)

var jobNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$`)

const TaskTypeCloudbrainOne = 0
const TaskTypeModelArts = 1
const TaskTypeGrampusGPU = 2
const TaskTypeGrampusNPU = 3

func CloudbrainOneTrainJobCreate(ctx *context.Context, option api.CreateTrainJobOption) {

	displayJobName := option.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	image := strings.TrimSpace(option.Image)
	uuids := option.Attachment
	jobType := string(models.JobTypeTrain)

	codePath := setting.JobPath + jobName + cloudbrain.CodeMountPath
	branchName := option.BranchName
	repo := ctx.Repo.Repository

	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainBindingJobNameKey(jobType, displayJobName))
	defer lock.UnLock()
	spec, datasetInfos, datasetNames, err := checkParameters(ctx, option, lock, repo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	command, err := getTrainJobCommand(option)
	if err != nil {
		log.Error("getTrainJobCommand failed: %v", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	errStr := loadCodeAndMakeModelPath(repo, codePath, branchName, jobName, cloudbrain.ModelMountPath)
	if errStr != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr(errStr)))
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
		Description:          option.Description,
		BranchName:           branchName,
		BootFile:             option.BootFile,
		Params:               option.Params,
		CommitID:             commitID,
		BenchmarkTypeID:      0,
		BenchmarkChildTypeID: 0,
		ResultPath:           storage.GetMinioPath(jobName, cloudbrain.ResultPath+"/"),
		Spec:                 spec,
	}

	if option.ModelName != "" { //使用预训练模型训练
		req.ModelName = option.ModelName
		req.LabelName = option.LabelName
		req.CkptName = option.CkptName
		req.ModelId = option.ModelId
		req.ModelVersion = option.ModelVersion
		minioPreModelURL, err := dealModelInfo(option.ModelId, jobName, option.CkptName)
		if err != nil {
			log.Error("Can not find model", err)
			//cloudBrainNewDataPrepare(ctx, jobType)
			//ctx.RenderWithErr(ctx.Tr("repo.modelconvert.manage.model_not_exist"), tpl, &form)
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.modelconvert.manage.model_not_exist")))
			return
		}
		req.PreTrainModelPath = setting.Attachment.Minio.RealPath + minioPreModelURL
		req.PreTrainModelUrl = minioPreModelURL

		//req.PreTrainModelPath = setting.Attachment.Minio.RealPath + option.PreTrainModelUrl
		//req.PreTrainModelUrl = option.PreTrainModelUrl

	}

	jobId, err := cloudbrain.GenerateTask(req)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{
		Code:    0,
		Message: jobId,
	})
}
func ModelArtsTrainJobNpuCreate(ctx *context.Context, option api.CreateTrainJobOption) {
	VersionOutputPath := modelarts.GetOutputPathByCount(modelarts.TotalVersionCount)
	displayJobName := option.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	uuid := option.Attachment
	description := option.Description
	workServerNumber := option.WorkServerNumber
	engineID, _ := strconv.Atoi(option.ImageID)
	bootFile := strings.TrimSpace(option.BootFile)
	params := option.Params
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + modelarts.CodePath
	codeObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.CodePath + VersionOutputPath + "/"
	outputObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.OutputPath + VersionOutputPath + "/"
	logObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.LogPath + VersionOutputPath + "/"
	branchName := option.BranchName
	isLatestVersion := modelarts.IsLatestVersion
	VersionCount := modelarts.VersionCountOne
	EngineName := option.Image

	errStr := checkMultiNode(ctx.User.ID, option.WorkServerNumber)
	if errStr != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr(errStr)))
		return
	}

	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainBindingJobNameKey(string(models.JobTypeTrain), displayJobName))
	defer lock.UnLock()

	spec, _, _, err := checkParameters(ctx, option, lock, repo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	//todo: del the codeLocalPath
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}
	var codeDownloadRepo = repo
	if option.CodeRepo != nil {
		codeDownloadRepo = option.CodeRepo.(*models.Repository)
	}

	gitRepo, _ := git.OpenRepository(codeDownloadRepo.RepoPath())
	commitID, _ := gitRepo.GetBranchCommitID(branchName)

	if err := downloadCode(codeDownloadRepo, codeLocalPath, branchName); err != nil {
		log.Error("downloadCode failed, server timed out: %s (%v)", repo.FullName(), err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	//todo: upload code (send to file_server todo this work?)
	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.OutputPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_output: %s (%v)", repo.FullName(), err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Failed to obsMkdir_output"))
		return
	}

	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.LogPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_log: %s (%v)", repo.FullName(), err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Failed to obsMkdir_log"))
		return
	}

	parentDir := VersionOutputPath + "/"
	if err := uploadCodeToObs(codeLocalPath, jobName, parentDir); err != nil {
		// if err := uploadCodeToObs(codeLocalPath, jobName, parentDir); err != nil {
		log.Error("Failed to uploadCodeToObs: %s (%v)", repo.FullName(), err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	var parameters models.Parameters
	param := make([]models.Parameter, 0)
	existDeviceTarget := false
	if len(params) != 0 {
		err := json.Unmarshal([]byte(params), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", params, err)
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("运行参数错误"))
			return
		}

		for _, parameter := range parameters.Parameter {
			if parameter.Label == modelarts.DeviceTarget {
				existDeviceTarget = true
			}
			if parameter.Label != modelarts.TrainUrl && parameter.Label != modelarts.DataUrl {
				param = append(param, models.Parameter{
					Label: parameter.Label,
					Value: parameter.Value,
				})
			}
		}
	}
	if !existDeviceTarget {
		param = append(param, models.Parameter{
			Label: modelarts.DeviceTarget,
			Value: modelarts.Ascend,
		})
	}
	datasUrlList, dataUrl, datasetNames, isMultiDataset, err := getDatasUrlListByUUIDS(uuid)
	if err != nil {
		log.Error("Failed to getDatasUrlListByUUIDS: %v", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Failed to getDatasUrlListByUUIDS:"+err.Error()))
		return
	}
	dataPath := dataUrl
	jsondatas, err := json.Marshal(datasUrlList)
	if err != nil {
		log.Error("Failed to Marshal: %v", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("json error:"+err.Error()))
		return
	}
	if isMultiDataset {
		param = append(param, models.Parameter{
			Label: modelarts.MultiDataUrl,
			Value: string(jsondatas),
		})
	}
	if option.ModelName != "" { //使用预训练模型训练
		ckptUrl := "/" + option.PreTrainModelUrl + option.CkptName
		param = append(param, models.Parameter{
			Label: modelarts.CkptUrl,
			Value: "s3:/" + ckptUrl,
		})
	}

	req := &modelarts.GenerateTrainJobReq{
		JobName:           jobName,
		DisplayJobName:    displayJobName,
		DataUrl:           dataPath,
		Description:       description,
		CodeObsPath:       codeObsPath,
		BootFileUrl:       codeObsPath + bootFile,
		BootFile:          bootFile,
		TrainUrl:          outputObsPath,
		WorkServerNumber:  workServerNumber,
		EngineID:          int64(engineID),
		LogUrl:            logObsPath,
		PoolID:            getPoolId(),
		Uuid:              uuid,
		Parameters:        param,
		CommitID:          commitID,
		IsLatestVersion:   isLatestVersion,
		BranchName:        branchName,
		Params:            option.Params,
		EngineName:        EngineName,
		VersionCount:      VersionCount,
		TotalVersionCount: modelarts.TotalVersionCount,
		DatasetName:       datasetNames,
		Spec:              spec,
		FineTune:          option.FineTune,
		FineTuneModelType: option.FineTuneModelType,
		FineTuneCategory:  option.FineTuneCategory,
	}
	if option.ModelName != "" { //使用预训练模型训练
		req.ModelName = option.ModelName
		req.LabelName = option.LabelName
		req.CkptName = option.CkptName
		req.ModelId = option.ModelId
		req.ModelVersion = option.ModelVersion
		req.PreTrainModelUrl = option.PreTrainModelUrl

	}

	userCommand, userImageUrl := getUserCommand(engineID, req)
	req.UserCommand = userCommand
	req.UserImageUrl = userImageUrl

	//将params转换Parameters.Parameter，出错时返回给前端
	var Parameters modelarts.Parameters
	if err := json.Unmarshal([]byte(params), &Parameters); err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("json.Unmarshal failed:"+err.Error()))
		return
	}

	jobId, err := modelarts.GenerateTrainJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error())
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{
		Code:    0,
		Message: jobId,
	})

}

func GrampusTrainJobGpuCreate(ctx *context.Context, option api.CreateTrainJobOption) {

	displayJobName := option.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	uuid := option.Attachment
	description := option.Description
	bootFile := strings.TrimSpace(option.BootFile)
	params := option.Params
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + cloudbrain.CodeMountPath + "/"
	branchName := option.BranchName
	image := strings.TrimSpace(option.Image)

	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainBindingJobNameKey(string(models.JobTypeTrain), displayJobName))
	defer lock.UnLock()
	spec, datasetInfos, datasetNames, err := checkParameters(ctx, option, lock, repo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	//prepare code and out path
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	if err := downloadZipCode(ctx, codeLocalPath, branchName); err != nil {
		log.Error("downloadZipCode failed, server timed out: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))

	}

	//todo: upload code (send to file_server todo this work?)
	//upload code
	if err := uploadCodeToMinio(codeLocalPath+"/", jobName, cloudbrain.CodeMountPath+"/"); err != nil {
		log.Error("Failed to uploadCodeToMinio: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	modelPath := setting.JobPath + jobName + cloudbrain.ModelMountPath + "/"
	if err := mkModelPath(modelPath); err != nil {
		log.Error("Failed to mkModelPath: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	//init model readme
	if err := uploadCodeToMinio(modelPath, jobName, cloudbrain.ModelMountPath+"/"); err != nil {
		log.Error("Failed to uploadCodeToMinio: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	var datasetRemotePath, allFileName string
	for _, datasetInfo := range datasetInfos {
		if datasetRemotePath == "" {
			datasetRemotePath = datasetInfo.DataLocalPath
			allFileName = datasetInfo.FullName
		} else {
			datasetRemotePath = datasetRemotePath + ";" + datasetInfo.DataLocalPath
			allFileName = allFileName + ";" + datasetInfo.FullName
		}

	}

	command, err := generateCommand(repo.Name, grampus.ProcessorTypeGPU, bootFile, params, setting.CBCodePathPrefix+jobName+cloudbrain.ModelMountPath+"/", datasetNames, option.CkptName, "")
	if err != nil {
		log.Error("Failed to generateCommand: %s (%v)", displayJobName, err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Create task failed, internal error"))
		return
	}

	commitID, _ := ctx.Repo.GitRepo.GetBranchCommitID(branchName)

	req := &grampus.GenerateTrainJobReq{
		JobName:         jobName,
		DisplayJobName:  displayJobName,
		ComputeResource: models.GPUResource,
		ProcessType:     grampus.ProcessorTypeGPU,
		Command:         command,
		ImageUrl:        image,
		Description:     description,
		BootFile:        bootFile,
		Uuid:            uuid,
		CommitID:        commitID,
		BranchName:      branchName,
		Params:          option.Params,
		EngineName:      image,
		DatasetNames:    datasetNames,
		DatasetInfos:    datasetInfos,

		IsLatestVersion:  modelarts.IsLatestVersion,
		VersionCount:     modelarts.VersionCountOne,
		WorkServerNumber: 1,
		Spec:             spec,
	}

	if option.ModelName != "" { //使用预训练模型训练
		req.ModelName = option.ModelName
		req.LabelName = option.LabelName
		req.CkptName = option.CkptName
		req.ModelId = option.ModelId
		req.ModelVersion = option.ModelVersion
		req.PreTrainModelUrl = option.PreTrainModelUrl
		preTrainModelPath := getPreTrainModelPath(option.PreTrainModelUrl, option.CkptName)
		req.PreTrainModelPath = preTrainModelPath
	}

	jobId, err := grampus.GenerateTrainJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 0, Message: jobId})
}

func checkParameters(ctx *context.Context, option api.CreateTrainJobOption, lock *redis_lock.DistributeLock, repo *models.Repository) (*models.Specification, map[string]models.DatasetInfo, string, error) {
	//check specification
	computeType := models.GPU

	if isNpuTask(option) {
		computeType = models.NPU
	}
	cluster := models.OpenICluster
	if isC2NetTask(option) {
		cluster = models.C2NetCluster
	}
	aiCenterCode := ""
	if option.Type == TaskTypeCloudbrainOne {
		aiCenterCode = models.AICenterOfCloudBrainOne
	} else if option.Type == TaskTypeModelArts {
		aiCenterCode = models.AICenterOfCloudBrainTwo
	}

	spec, err := resource.GetAndCheckSpec(ctx.User.ID, option.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: computeType,
		Cluster:         cluster,
		AiCenterCode:    aiCenterCode,
	})
	if err != nil || spec == nil {
		return nil, nil, "", fmt.Errorf("Resource specification is not available.")
	}

	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
		log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
		return nil, nil, "", fmt.Errorf(ctx.Tr("points.insufficient_points_balance"))
	}

	isOk, err := lock.Lock(models.CloudbrainKeyDuration)
	if !isOk {
		log.Error("lock processed failed:%v", err, ctx.Data["MsgID"])

		return nil, nil, "", fmt.Errorf(ctx.Tr("repo.cloudbrain_samejob_err"))
	}

	if !jobNamePattern.MatchString(option.DisplayJobName) {
		return nil, nil, "", fmt.Errorf(ctx.Tr("repo.cloudbrain_jobname_err"))
	}
	if option.CodeRepo == nil {
		bootFileExist, err := ctx.Repo.FileExists(option.BootFile, option.BranchName)
		if err != nil || !bootFileExist {
			log.Error("Get bootfile error:", err, ctx.Data["MsgID"])
			return nil, nil, "", fmt.Errorf(ctx.Tr("repo.cloudbrain_bootfile_err"))
		}
	}

	count, err := GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	if err != nil {
		log.Error("GetCountByUserID failed:%v", err, ctx.Data["MsgID"])
		return nil, nil, "", fmt.Errorf("system error")
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
			return nil, nil, "", fmt.Errorf("you have already a running or waiting task, can not create more.")
		}
	}

	//check param
	if err := paramCheckCreateTrainJob(option.BootFile, option.BranchName); err != nil {
		log.Error("paramCheckCreateTrainJob failed:(%v)", err, ctx.Data["MsgID"])
		return nil, nil, "", err
	}

	//check whether the task name in the project is duplicated
	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, string(models.JobTypeTrain), option.DisplayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist", ctx.Data["MsgID"])
			return nil, nil, "", fmt.Errorf("The job name did already exist.")
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err, ctx.Data["MsgID"])
			return nil, nil, "", fmt.Errorf("system error")
		}
	}

	//check dataset
	var datasetInfos map[string]models.DatasetInfo
	var datasetNames string
	if option.Type != TaskTypeModelArts {
		if isC2NetTask(option) {
			datasetInfos, datasetNames, err = models.GetDatasetInfo(option.Attachment, computeType)
		} else {
			datasetInfos, datasetNames, err = models.GetDatasetInfo(option.Attachment)
		}

		if err != nil {
			log.Error("GetDatasetInfo failed: %v", err, ctx.Data["MsgID"])
			return nil, nil, "", fmt.Errorf(ctx.Tr("cloudbrain.error.dataset_select"))
		}
	}
	return spec, datasetInfos, datasetNames, err
}

func isNpuTask(option api.CreateTrainJobOption) bool {
	return option.Type == TaskTypeModelArts || option.Type == TaskTypeGrampusNPU
}

func isC2NetTask(option api.CreateTrainJobOption) bool {
	return option.Type == TaskTypeGrampusGPU || option.Type == TaskTypeGrampusNPU
}

func GrampusTrainJobNpuCreate(ctx *context.Context, option api.CreateTrainJobOption) {

	displayJobName := option.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	uuid := option.Attachment
	description := option.Description
	bootFile := strings.TrimSpace(option.BootFile)
	params := option.Params
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + modelarts.CodePath
	codeObsPath := grampus.JobPath + jobName + modelarts.CodePath
	branchName := option.BranchName
	isLatestVersion := modelarts.IsLatestVersion
	versionCount := modelarts.VersionCountOne
	engineName := option.Image

	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainBindingJobNameKey(string(models.JobTypeTrain), displayJobName))
	defer lock.UnLock()
	spec, datasetInfos, datasetNames, err := checkParameters(ctx, option, lock, repo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	errMsg := CheckGrampusNPUMultiNode(ctx.User.ID, option.WorkServerNumber)
	if errMsg != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr(errMsg)))
		return
	}

	//prepare code and out path
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	if err := downloadZipCode(ctx, codeLocalPath, branchName); err != nil {
		log.Error("downloadZipCode failed, server timed out: %s (%v)", repo.FullName(), err)

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	//todo: upload code (send to file_server todo this work?)
	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.OutputPath); err != nil {
		log.Error("Failed to obsMkdir_output: %s (%v)", repo.FullName(), err)

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	if err := uploadCodeToObs(codeLocalPath, jobName, ""); err != nil {
		log.Error("Failed to uploadCodeToObs: %s (%v)", repo.FullName(), err)

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	var datasetRemotePath, allFileName string
	for _, datasetInfo := range datasetInfos {
		if datasetRemotePath == "" {
			datasetRemotePath = datasetInfo.DataLocalPath + "'" + datasetInfo.FullName + "'"
			allFileName = datasetInfo.FullName
		} else {
			datasetRemotePath = datasetRemotePath + ";" + datasetInfo.DataLocalPath + "'" + datasetInfo.FullName + "'"
			allFileName = allFileName + ";" + datasetInfo.FullName
		}

	}

	//prepare command
	preTrainModelPath := getPreTrainModelPath(option.PreTrainModelUrl, option.CkptName)
	command, err := generateCommand(repo.Name, grampus.ProcessorTypeNPU, bootFile, params, setting.CodePathPrefix+jobName+modelarts.OutputPath, datasetNames, option.CkptName, grampus.GetNpuModelRemoteObsUrl(jobName))

	if err != nil {
		log.Error("Failed to generateCommand: %s (%v)", displayJobName, err, ctx.Data["MsgID"])

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Create task failed, internal error"))
		return
	}

	commitID, _ := ctx.Repo.GitRepo.GetBranchCommitID(branchName)

	req := &grampus.GenerateTrainJobReq{
		JobName:           jobName,
		DisplayJobName:    displayJobName,
		ComputeResource:   models.NPUResource,
		ProcessType:       grampus.ProcessorTypeNPU,
		Command:           command,
		ImageId:           option.ImageID,
		Description:       description,
		CodeObsPath:       codeObsPath,
		BootFileUrl:       codeObsPath + bootFile,
		BootFile:          bootFile,
		WorkServerNumber:  option.WorkServerNumber,
		Uuid:              uuid,
		CommitID:          commitID,
		IsLatestVersion:   isLatestVersion,
		BranchName:        branchName,
		Params:            option.Params,
		EngineName:        engineName,
		VersionCount:      versionCount,
		TotalVersionCount: modelarts.TotalVersionCount,
		DatasetNames:      datasetNames,
		DatasetInfos:      datasetInfos,
		Spec:              spec,
		CodeName:          strings.ToLower(repo.Name),
	}
	if option.ModelName != "" { //使用预训练模型训练
		req.ModelName = option.ModelName
		req.LabelName = option.LabelName
		req.CkptName = option.CkptName
		req.ModelId = option.ModelId
		req.ModelVersion = option.ModelVersion
		req.PreTrainModelUrl = option.PreTrainModelUrl
		req.PreTrainModelPath = preTrainModelPath
	}

	jobId, err := grampus.GenerateTrainJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error())

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 0, Message: jobId})
}

func obsMkdir(dir string) error {
	input := &obs.PutObjectInput{}
	input.Bucket = setting.Bucket
	input.Key = dir
	_, err := storage.ObsCli.PutObject(input)
	if err != nil {
		log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
		return err
	}

	return nil
}
func uploadCodeToObs(codePath, jobName, parentDir string) error {
	files, err := readDir(codePath)
	if err != nil {
		log.Error("readDir(%s) failed: %s", codePath, err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			input := &obs.PutObjectInput{}
			input.Bucket = setting.Bucket
			input.Key = parentDir + file.Name() + "/"
			_, err = storage.ObsCli.PutObject(input)
			if err != nil {
				log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
				return err
			}

			if err = uploadCodeToObs(codePath+file.Name()+"/", jobName, parentDir+file.Name()+"/"); err != nil {
				log.Error("uploadCodeToObs(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		} else {
			input := &obs.PutFileInput{}
			input.Bucket = setting.Bucket
			input.Key = setting.CodePathPrefix + jobName + "/code/" + parentDir + file.Name()
			input.SourceFile = codePath + file.Name()
			_, err = storage.ObsCli.PutFile(input)
			if err != nil {
				log.Error("PutFile(%s) failed: %s", input.SourceFile, err.Error())
				return err
			}
		}
	}

	return nil
}

func paramCheckCreateTrainJob(bootFile string, branchName string) error {
	if !strings.HasSuffix(strings.TrimSpace(bootFile), ".py") {
		log.Error("the boot file(%s) must be a python file", bootFile)
		return errors.New("启动文件必须是python文件")
	}

	if branchName == "" {
		log.Error("the branch must not be null!", branchName)
		return errors.New("代码分支不能为空！")
	}

	return nil
}

func CheckGrampusNPUMultiNode(userId int64, serverNum int) string {
	if serverNum == 1 {
		return ""
	}
	grampus.InitMultiNode()
	var isServerNumValid = false
	if grampus.MultiNodeConfig != nil {
		for _, info := range grampus.MultiNodeConfig.Info {
			if isInOrg, _ := models.IsOrganizationMemberByOrgName(info.Org, userId); isInOrg {
				if isInNodes(info.Node, serverNum) {
					isServerNumValid = true
					break
				}

			}
		}
	}
	if isServerNumValid {
		return ""
	} else {
		return "repo.modelarts.no_node_right"
	}
}
func downloadZipCode(ctx *context.Context, codePath, branchName string) error {
	archiveType := git.ZIP
	archivePath := codePath

	if !com.IsDir(archivePath) {
		if err := os.MkdirAll(archivePath, os.ModePerm); err != nil {
			log.Error("MkdirAll failed:" + err.Error())
			return err
		}
	}

	// Get corresponding commit.
	var (
		commit *git.Commit
		err    error
	)

	gitRepo := ctx.Repo.GitRepo
	if err != nil {
		log.Error("OpenRepository failed:" + err.Error())
		return err
	}

	if gitRepo.IsBranchExist(branchName) {
		commit, err = gitRepo.GetBranchCommit(branchName)
		if err != nil {
			log.Error("GetBranchCommit failed:" + err.Error())
			return err
		}
	} else {
		log.Error("the branch is not exist: " + branchName)
		return fmt.Errorf("The branch does not exist.")
	}

	archivePath = path.Join(archivePath, grampus.CodeArchiveName)
	if !com.IsFile(archivePath) {
		if err := commit.CreateArchive(archivePath, git.CreateArchiveOpts{
			Format: archiveType,
			Prefix: setting.Repository.PrefixArchiveFiles,
		}); err != nil {
			log.Error("CreateArchive failed:" + err.Error())
			return err
		}
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

				if strings.Contains(err.Error(), "no such file or directory") {
					continue
				}
				log.Error("UploadObject(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		}
	}

	return nil
}

func uploadOneFileToMinio(codePath, filePath, jobName, parentDir string) error {
	destObject := setting.CBCodePathPrefix + jobName + parentDir + path.Base(filePath)
	sourceFile := codePath + "/" + filePath
	err := storage.Attachments.UploadObject(destObject, sourceFile)
	if err != nil {
		log.Error("UploadObject(%s) failed: %s", filePath, err.Error())
		return err
	}
	return nil

}

func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	list, err := f.Readdir(0)
	f.Close()
	if err != nil {
		//todo: can not upload empty folder
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	//sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
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

	fileName := path + "README"
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

func getPreTrainModelPath(pretrainModelDir string, fileName string) string {
	index := strings.Index(pretrainModelDir, "/")
	if index > 0 {
		filterBucket := pretrainModelDir[index+1:]
		return filterBucket + fileName
	} else {
		return ""
	}

}

func generateCommand(repoName, processorType, bootFile, paramSrc, outputRemotePath, datasetName, pretrainModelFileName, modelRemoteObsUrl string) (string, error) {
	var command string

	//prepare
	workDir := grampus.NpuWorkDir
	if processorType == grampus.ProcessorTypeNPU {
		command += "pwd;cd " + workDir + grampus.CommandPrepareScriptNpu
	} else if processorType == grampus.ProcessorTypeGPU {
		workDir = grampus.GpuWorkDir
		command += "pwd; cd " + workDir + fmt.Sprintf(grampus.CommandPrepareScriptGpu)
	}
	//unzip code & dataset
	if processorType == grampus.ProcessorTypeNPU {
		//no need to process
	} else if processorType == grampus.ProcessorTypeGPU {
		unZipDatasetCommand := GenerateDatasetUnzipCommand(datasetName)
		commandUnzip := "cd " + workDir + "code;echo \"start unzip code\";unzip -q master.zip;echo \"start to unzip dataset\";cd " + workDir + "dataset;" + unZipDatasetCommand
		command += commandUnzip
	}

	command += "echo \"unzip finished;start to exec code;\";"

	// set export
	var commandExport string
	if processorType == grampus.ProcessorTypeNPU {
		commandExport = "export bucket=" + setting.Bucket + " && export remote_path=" + outputRemotePath + ";"
	} else if processorType == grampus.ProcessorTypeGPU {
		commandExport = "export env=" + setting.Grampus.Env + " && export remote_path=" + outputRemotePath + ";"
	}

	command += commandExport

	//exec code
	var parameters models.Parameters
	var paramCode string

	if len(paramSrc) != 0 {
		err := json.Unmarshal([]byte(paramSrc), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", paramSrc, err)
			return command, err
		}

		for _, parameter := range parameters.Parameter {
			paramCode += " --" + parameter.Label + "=" + parameter.Value
		}
	}

	var commandCode string
	if processorType == grampus.ProcessorTypeNPU {
		paramCode += " --model_url=" + modelRemoteObsUrl
		commandCode = "source /home/ma-user/.bashrc;python /home/ma-user/davinci/train/davincirun.py python /home/ma-user/openi.py " + paramCode + ";"
	} else if processorType == grampus.ProcessorTypeGPU {
		if pretrainModelFileName != "" {
			paramCode += " --ckpt_url" + "=" + workDir + "pretrainmodel/" + pretrainModelFileName
		}
		commandCode = "cd " + workDir + "code/" + strings.ToLower(repoName) + ";python " + bootFile + paramCode + ";"
	}

	command += commandCode

	//get exec result
	commandGetRes := "result=$?;"
	command += commandGetRes

	////upload models
	//if processorType == grampus.ProcessorTypeNPU {
	//	// no need to upload
	//} else if processorType == grampus.ProcessorTypeGPU {
	//	commandUpload := "cd " + workDir + setting.Grampus.SyncScriptProject + "/;./uploader_for_gpu " + setting.Grampus.Env + " " + outputRemotePath + " " + workDir + "output/;"
	//	command += commandUpload
	//}

	//check exec result
	commandCheckRes := "bash -c \"[[ $result -eq 0 ]] && exit 0 || exit -1\""
	command += commandCheckRes

	return command, nil
}
func processPretrainModelParameter(pretrainModelPath string, pretrainModelFileName string, commandDownload string) string {
	commandDownloadTemp := commandDownload
	if pretrainModelPath != "" {
		commandDownloadTemp += " '" + pretrainModelPath + "' '" + pretrainModelFileName + "'"
	}
	commandDownloadTemp += ";"
	return commandDownloadTemp
}

func GenerateDatasetUnzipCommand(datasetName string) string {
	var unZipDatasetCommand string

	datasetNameArray := strings.Split(datasetName, ";")
	if len(datasetNameArray) == 1 { //单数据集
		unZipDatasetCommand = "unzip -q '" + datasetName + "';"
		if strings.HasSuffix(datasetNameArray[0], ".tar.gz") {
			unZipDatasetCommand = "tar --strip-components=1 -zxvf '" + datasetName + "';"
		}
		//unZipDatasetCommand += "rm -f '" + datasetName + "';"

	} else { //多数据集
		for _, datasetNameTemp := range datasetNameArray {
			if strings.HasSuffix(datasetNameTemp, ".tar.gz") {
				unZipDatasetCommand = unZipDatasetCommand + "tar -zxvf '" + datasetNameTemp + "';"
			} else {
				unZipDatasetCommand = unZipDatasetCommand + "unzip -q '" + datasetNameTemp + "' -d './" + strings.TrimSuffix(datasetNameTemp, ".zip") + "';"
			}
			//unZipDatasetCommand += "rm -f '" + datasetNameTemp + "';"
		}

	}
	return unZipDatasetCommand
}

func getPoolId() string {
	var resourcePools modelarts.ResourcePool
	json.Unmarshal([]byte(setting.ResourcePools), &resourcePools)

	return resourcePools.Info[0].ID
}

func PrepareSpec4Show(task *models.Cloudbrain) {
	s, err := resource.GetCloudbrainSpec(task.ID)
	if err != nil {
		log.Info("error:" + err.Error())
		return
	}
	task.Spec = s
}

func IsTaskNotStop(task *models.Cloudbrain) bool {
	statuses := CloudbrainOneNotFinalStatuses
	if task.Type == models.TypeCloudBrainTwo || task.Type == models.TypeCDCenter {
		statuses = CloudbrainTwoNotFinalStatuses
	} else {
		statuses = GrampusNotFinalStatuses
	}

	for _, status := range statuses {
		if task.Status == status {
			return true
		}
	}
	return false

}

func SyncTaskStatus(task *models.Cloudbrain) error {
	if task.Type == models.TypeCloudBrainOne {
		result, err := cloudbrain.GetJob(task.JobID)
		if err != nil {
			log.Info("error:" + err.Error())
			return fmt.Errorf("repo.cloudbrain_query_fail")
		}

		if result != nil {
			jobRes, _ := models.ConvertToJobResultPayload(result.Payload)
			taskRoles := jobRes.TaskRoles
			taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))

			oldStatus := task.Status
			task.Status = taskRes.TaskStatuses[0].State

			task.ContainerID = taskRes.TaskStatuses[0].ContainerID
			models.ParseAndSetDurationFromCloudBrainOne(jobRes, task)

			if task.DeletedAt.IsZero() { //normal record
				if oldStatus != task.Status {
					notification.NotifyChangeCloudbrainStatus(task, oldStatus)
				}
				err = models.UpdateJob(task)
				if err != nil {
					return fmt.Errorf("repo.cloudbrain_query_fail")

				}
			}

		} else {
			log.Info("error:" + err.Error())
			return fmt.Errorf("repo.cloudbrain_query_fail")
		}
	} else if task.Type == models.TypeCloudBrainTwo || task.Type == models.TypeCDCenter {
		err := modelarts.HandleTrainJobInfo(task)
		if err != nil {
			return fmt.Errorf("repo.cloudbrain_query_fail")
		}

	} else if task.Type == models.TypeC2Net {
		result, err := grampus.GetJob(task.JobID)
		if err != nil {
			log.Error("GetJob failed:" + err.Error())
			return fmt.Errorf("repo.cloudbrain_query_fail")
		}

		if result != nil {
			if len(result.JobInfo.Tasks[0].CenterID) == 1 && len(result.JobInfo.Tasks[0].CenterName) == 1 {
				task.AiCenter = result.JobInfo.Tasks[0].CenterID[0] + "+" + result.JobInfo.Tasks[0].CenterName[0]
			}
			oldStatus := task.Status
			task.Status = grampus.TransTrainJobStatus(result.JobInfo.Status)

			if task.Status != oldStatus || task.Status == models.GrampusStatusRunning {
				task.Duration = result.JobInfo.RunSec
				if task.Duration < 0 {
					task.Duration = 0
				}
				task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)

				if task.StartTime == 0 && result.JobInfo.StartedAt > 0 {
					task.StartTime = timeutil.TimeStamp(result.JobInfo.StartedAt)
				}
				if task.EndTime == 0 && models.IsTrainJobTerminal(task.Status) && task.StartTime > 0 {
					task.EndTime = task.StartTime.Add(task.Duration)
				}
				task.CorrectCreateUnix()
				if oldStatus != task.Status {
					notification.NotifyChangeCloudbrainStatus(task, oldStatus)
				}
				err = models.UpdateJob(task)
				if err != nil {
					log.Error("UpdateJob failed:" + err.Error())
					return fmt.Errorf("repo.cloudbrain_query_fail")
				}
			}
		}
	}
	return nil

}

func getTrainJobCommand(option api.CreateTrainJobOption) (string, error) {
	var command string
	bootFile := strings.TrimSpace(option.BootFile)
	params := option.Params

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
	if option.CkptName != "" {
		param += " --ckpt_url" + "=" + "/pretrainmodel/" + option.CkptName
	}

	command += "python /code/" + bootFile + param + " > " + cloudbrain.ModelMountPath + "/" + option.DisplayJobName + "-" + cloudbrain.LogFile

	return command, nil
}

func checkMultiNode(userId int64, serverNum int) string {
	if serverNum == 1 {
		return ""
	}
	modelarts.InitMultiNode()
	var isServerNumValid = false
	if modelarts.MultiNodeConfig != nil {
		for _, info := range modelarts.MultiNodeConfig.Info {
			if isInOrg, _ := models.IsOrganizationMemberByOrgName(info.Org, userId); isInOrg {
				if isInNodes(info.Node, serverNum) {
					isServerNumValid = true
					break
				}

			}
		}
	}
	if isServerNumValid {
		return ""
	} else {
		return "repo.modelarts.no_node_right"
	}
}

func isInNodes(nodes []int, num int) bool {
	for _, node := range nodes {
		if node == num {
			return true
		}
	}
	return false

}

func getUserCommand(engineId int, req *modelarts.GenerateTrainJobReq) (string, string) {
	userImageUrl := ""
	userCommand := ""
	if engineId < 0 {
		tmpCodeObsPath := strings.Trim(req.CodeObsPath, "/")
		tmpCodeObsPaths := strings.Split(tmpCodeObsPath, "/")
		lastCodeDir := "code"
		if len(tmpCodeObsPaths) > 0 {
			lastCodeDir = tmpCodeObsPaths[len(tmpCodeObsPaths)-1]
		}
		userCommand = "/bin/bash /home/work/run_train.sh 's3://" + req.CodeObsPath + "' '" + lastCodeDir + "/" + req.BootFile + "' '/tmp/log/train.log' --'data_url'='s3://" + req.DataUrl + "' --'train_url'='s3://" + req.TrainUrl + "'"
		var versionInfos modelarts.VersionInfo
		if err := json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
			log.Info("json parse err." + err.Error())
		} else {
			for _, engine := range versionInfos.Version {
				if engine.ID == engineId {
					userImageUrl = engine.Url
					break
				}
			}
		}
		for _, param := range req.Parameters {
			userCommand += " --'" + param.Label + "'='" + param.Value + "'"
		}
		return userCommand, userImageUrl
	}
	return userCommand, userImageUrl
}
