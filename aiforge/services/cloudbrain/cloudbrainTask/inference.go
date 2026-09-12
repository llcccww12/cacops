package cloudbrainTask

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
	"code.gitea.io/gitea/services/lock"

	"code.gitea.io/gitea/modules/modelarts"

	"code.gitea.io/gitea/modules/git"

	api "code.gitea.io/gitea/modules/structs"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward/point/account"
)

const CLONE_FILE_PREFIX = "file:///"

func dealModelInfo(modelId string, jobName string, ckptName string) (string, error) {
	preModel, err := models.QueryModelById(modelId)
	if err != nil {
		log.Error("Can not find model", err)
		return "", err
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
	return returnStr, nil
}

func ModelArtsInferenceJobCreate(ctx *context.Context, option api.CreateTrainJobOption) {
	ctx.Data["PageIsTrainJob"] = true
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
	codeObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.CodePath
	resultObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.ResultPath + VersionOutputPath + "/"
	logObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.LogPath + VersionOutputPath + "/"

	branchName := option.BranchName
	EngineName := option.Image
	LabelName := option.LabelName
	isLatestVersion := modelarts.IsLatestVersion
	VersionCount := modelarts.VersionCountOne
	trainUrl := option.PreTrainModelUrl
	modelName := option.ModelName
	modelVersion := option.ModelVersion
	ckptName := option.CkptName
	ckptUrl := "/" + option.PreTrainModelUrl + option.CkptName

	errStr := checkInferenceJobMultiNode(ctx.User.ID, option.WorkServerNumber)
	if errStr != "" {

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr(errStr)))
		return
	}

	spec, err := resource.GetAndCheckSpec(ctx.User.ID, option.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeInference,
		ComputeResource: models.NPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainTwo})
	if err != nil || spec == nil {

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Resource specification not available"))
		return
	}
	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
		log.Error("point balance is not enough,userId=%d specId=%d ", ctx.User.ID, spec.ID)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("points.insufficient_points_balance")))
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: displayJobName, JobType: string(models.JobTypeInference)}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr(errMsg)))
		return
	}

	count, err := GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeInference))
	if err != nil {
		log.Error("GetCloudbrainInferenceJobCountByUserID failed:%v", err, ctx.Data["MsgID"])

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("system error"))
		return
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting inference task", ctx.Data["MsgID"])

			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("you have already a running or waiting inference task, can not create more"))
			return
		}
	}

	if err := paramCheckCreateInferenceJob(option); err != nil {
		log.Error("paramCheckCreateInferenceJob failed:(%v)", err)

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	bootFileExist, err := ctx.Repo.FileExists(bootFile, branchName)
	if err != nil || !bootFileExist {
		log.Error("Get bootfile error:", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.cloudbrain_bootfile_err")))
		return
	}

	//Determine whether the task name of the task in the project is duplicated
	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, string(models.JobTypeInference), displayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist", ctx.Data["MsgID"])

			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("the job name did already exist"))
			return
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err, ctx.Data["MsgID"])

			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("system error"))
			return
		}
	}

	//todo: del the codeLocalPath
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	gitRepo, _ := git.OpenRepository(repo.RepoPath())
	commitID, _ := gitRepo.GetBranchCommitID(branchName)

	if err := downloadCode(repo, codeLocalPath, branchName); err != nil {
		log.Error("Create task failed, server timed out: %s (%v)", repo.FullName(), err)

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	//todo: upload code (send to file_server todo this work?)
	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.ResultPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_result: %s (%v)", repo.FullName(), err)

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Failed to obsMkdir_result"))
		return
	}

	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.LogPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_log: %s (%v)", repo.FullName(), err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Failed to obsMkdir_log"))
		return
	}

	if err := uploadCodeToObs(codeLocalPath, jobName, ""); err != nil {
		log.Error("Failed to uploadCodeToObs: %s (%v)", repo.FullName(), err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
		return
	}

	var parameters models.Parameters
	param := make([]models.Parameter, 0)
	param = append(param, models.Parameter{
		Label: modelarts.ResultUrl,
		Value: "s3:/" + resultObsPath,
	}, models.Parameter{
		Label: modelarts.CkptUrl,
		Value: "s3:/" + ckptUrl,
	})

	datasUrlList, dataUrl, datasetNames, isMultiDataset, err := getDatasUrlListByUUIDS(uuid)
	if err != nil {

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
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

	req := &modelarts.GenerateInferenceJobReq{
		JobName:           jobName,
		DisplayJobName:    displayJobName,
		DataUrl:           dataPath,
		Description:       description,
		CodeObsPath:       codeObsPath,
		BootFileUrl:       codeObsPath + bootFile,
		BootFile:          bootFile,
		TrainUrl:          trainUrl,
		WorkServerNumber:  workServerNumber,
		EngineID:          int64(engineID),
		LogUrl:            logObsPath,
		PoolID:            getPoolId(),
		Uuid:              uuid,
		Parameters:        param, //modelarts train parameters
		CommitID:          commitID,
		BranchName:        branchName,
		Params:            option.Params,
		EngineName:        EngineName,
		LabelName:         LabelName,
		IsLatestVersion:   isLatestVersion,
		VersionCount:      VersionCount,
		TotalVersionCount: modelarts.TotalVersionCount,
		ModelName:         modelName,
		ModelVersion:      modelVersion,
		CkptName:          ckptName,
		ModelId:           option.ModelId,
		ResultUrl:         resultObsPath,
		Spec:              spec,
		DatasetName:       datasetNames,
		JobType:           string(models.JobTypeInference),
	}

	jobId, err := modelarts.GenerateInferenceJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error())

		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: 0, Message: jobId})
}

func getDatasUrlListByUUIDS(uuidStr string) ([]models.Datasurl, string, string, bool, error) {
	var isMultiDataset bool
	var dataUrl string
	var datasetNames string
	var datasUrlList []models.Datasurl
	uuids := strings.Split(uuidStr, ";")
	if len(uuids) > setting.MaxDatasetNum {
		log.Error("the dataset count(%d) exceed the limit", len(uuids))
		return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("the dataset count exceed the limit")
	}

	datasetInfos := make(map[string]models.DatasetInfo)
	attachs, err := models.GetAttachmentsByUUIDs(uuids)
	if err != nil || len(attachs) != len(uuids) {
		log.Error("GetAttachmentsByUUIDs failed: %v", err)
		return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("GetAttachmentsByUUIDs failed")
	}

	for i, tmpUuid := range uuids {
		var attach *models.Attachment
		for _, tmpAttach := range attachs {
			if tmpAttach.UUID == tmpUuid {
				attach = tmpAttach
				break
			}
		}
		if attach == nil {
			log.Error("GetAttachmentsByUUIDs failed: %v", err)
			return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("GetAttachmentsByUUIDs failed")
		}
		fileName := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(attach.Name, ".zip"), ".tar.gz"), ".tgz")
		for _, datasetInfo := range datasetInfos {
			if fileName == datasetInfo.Name {
				log.Error("the dataset name is same: %v", attach.Name)
				return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("the dataset name is same")
			}
		}
		if len(attachs) <= 1 {
			//dataUrl = "/" + setting.Bucket + "/" + setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + attach.UUID + "/"
			dataUrl = "/" + setting.Bucket + "/" + attach.UnzipPath
			isMultiDataset = false
		} else {
			//dataUrl = "/" + setting.Bucket + "/" + setting.BasePath + path.Join(attachs[0].UUID[0:1], attachs[0].UUID[1:2]) + "/" + attachs[0].UUID + attachs[0].UUID + "/"
			dataUrl = "/" + setting.Bucket + "/" + attachs[0].UnzipPath
			//datasetUrl := "s3://" + setting.Bucket + "/" + setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + attach.UUID + "/"
			datasetUrl := "s3://" + setting.Bucket + "/" + attach.UnzipPath
			datasUrlList = append(datasUrlList, models.Datasurl{
				DatasetUrl:  datasetUrl,
				DatasetName: fileName,
			})
			isMultiDataset = true
		}

		if i == 0 {
			datasetNames = attach.Name
		} else {
			datasetNames += ";" + attach.Name
		}
	}

	return datasUrlList, dataUrl, datasetNames, isMultiDataset, nil
}
func checkInferenceJobMultiNode(userId int64, serverNum int) string {
	if serverNum == 1 {
		return ""
	}

	return "repo.modelarts.no_node_right"

}

func paramCheckCreateInferenceJob(option api.CreateTrainJobOption) error {
	if !strings.HasSuffix(strings.TrimSpace(option.BootFile), ".py") {
		log.Error("the boot file(%s) must be a python file", strings.TrimSpace(option.BootFile))
		return errors.New("启动文件必须是python文件")
	}

	if option.ModelName == "" {
		log.Error("the ModelName(%d) must not be nil", option.ModelName)
		return errors.New("模型名称不能为空")
	}
	if option.ModelVersion == "" {
		log.Error("the ModelVersion(%d) must not be nil", option.ModelVersion)
		return errors.New("模型版本不能为空")
	}
	if option.CkptName == "" {
		log.Error("the CkptName(%d) must not be nil", option.CkptName)
		return errors.New("权重文件不能为空")
	}
	if option.BranchName == "" {
		log.Error("the Branch(%d) must not be nil", option.BranchName)
		return errors.New("分支名不能为空")
	}

	if utf8.RuneCountInString(option.Description) > 255 {
		log.Error("the Description length(%d) must not more than 255", option.Description)
		return errors.New("描述字符不能超过255个字符")
	}

	return nil
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

	modelPath := setting.JobPath + jobName + resultPath + "/"
	err = mkModelPath(modelPath)
	if err != nil {
		return "cloudbrain.load_code_failed"
	}
	err = uploadCodeToMinio(modelPath, jobName, resultPath+"/")
	if err != nil {
		return "cloudbrain.load_code_failed"
	}

	return ""
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

func getInferenceJobCommand(option api.CreateTrainJobOption) (string, error) {
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

	param += " --modelname" + "=" + option.CkptName

	command += "python /code/" + bootFile + param + " > " + cloudbrain.ResultPath + "/" + option.DisplayJobName + "-" + cloudbrain.LogFile

	return command, nil
}
