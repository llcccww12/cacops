package repo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/manager/client/grampus"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	uuid "github.com/satori/go.uuid"
)

const (
	tplModelManageConvertIndex = "repo/modelmanage/convertIndex"
	tplModelConvertInfo        = "repo/modelmanage/convertshowinfo"
	PYTORCH_ENGINE             = 0
	TENSORFLOW_ENGINE          = 1
	MINDSPORE_ENGINE           = 2
	PADDLE_ENGINE              = 4
	MXNET_ENGINE               = 6
	ModelMountPath             = "/model"
	CodeMountPath              = "/code"
	DataSetMountPath           = "/tmp/dataset"
	LogFile                    = "log.txt"
	DefaultBranchName          = "master"
	SubTaskName                = "task1"
	//GpuQueue                   = "openidgx"
	Success = "S000"
	//GPU_PYTORCH_IMAGE      = "dockerhub.pcl.ac.cn:5000/user-images/openi:tensorRT_7_zouap"
	//GPU_TENSORFLOW_IMAGE   = "dockerhub.pcl.ac.cn:5000/user-images/openi:tf2onnx"
	//NPU_MINDSPORE_16_IMAGE = "swr.cn-south-222.ai.pcl.cn/openi/mindspore1.6.1_train_v1_openi:v3_ascend"
	//PytorchOnnxBootFile    = "convert_pytorch.py"
	//PytorchTrTBootFile     = "convert_pytorch_tensorrt.py"
	//MindsporeBootFile      = "convert_mindspore.py"
	//TensorFlowNpuBootFile  = "convert_tensorflow.py"
	//TensorFlowGpuBootFile  = "convert_tensorflow_gpu.py"

	//ConvertRepoPath = "https://openi.pcl.ac.cn/zouap/npu_test"

	CONVERT_FORMAT_ONNX = 0
	CONVERT_FORMAT_TRT  = 1

	NetOutputFormat_FP32 = 0
	NetOutputFormat_FP16 = 1

	//NPU_MINDSPORE_IMAGE_ID  = 37
	//NPU_TENSORFLOW_IMAGE_ID = 38

	//GPU_Resource_Specs_ID = 1 //cpu 1, gpu 1

	//NPU_FlavorCode = "modelarts.bm.910.arm.public.1"
	//NPU_PoolID     = "pool7908321a"
)

var (
	TrainResourceSpecs *models.ResourceSpecs
)

func SaveModelConvert(ctx *context.Context) {
	log.Info("save model convert start.")
	if !ctx.Repo.CanWrite(models.UnitTypeModelManage) {
		ctx.JSON(200, map[string]string{
			"code": "1",
			"msg":  ctx.Tr("repo.modelconvert.manage.no_operate_right"),
		})
		return
	}
	name := ctx.Query("name")
	desc := ctx.Query("desc")
	modelId := ctx.Query("modelId")
	modelPath := ctx.Query("modelFile")
	SrcEngine := ctx.QueryInt("srcEngine")
	InputShape := ctx.Query("inputshape")
	InputDataFormat := ctx.Query("inputdataformat")
	DestFormat := ctx.QueryInt("destFormat")
	NetOutputFormat := ctx.QueryInt("netOutputFormat")

	task, err := models.QueryModelById(modelId)
	if err != nil {
		log.Error("no such model!", err.Error())
		ctx.JSON(200, map[string]string{
			"code": "1",
			"msg":  ctx.Tr("repo.modelconvert.manage.model_not_exist"),
		})
		return
	}

	convertList, err := models.QueryModelConvertByRepoID(ctx.Repo.Repository.ID)
	if err == nil {
		for _, convert := range convertList {
			if convert.Name == name {
				log.Info("convert.Name=" + name + " convert.id=" + convert.ID)
				ctx.JSON(200, map[string]string{
					"code": "1",
					"msg":  ctx.Tr("repo.modelconvert.manage.create_error1"),
				})
				return
			}
		}
	}

	convertList, err = models.QueryModelConvertByUserID(ctx.User.ID)
	if err == nil {
		for _, convert := range convertList {
			if isRunningTask(convert.Status) {
				log.Info("convert.Status=" + convert.Status + " convert.id=" + convert.ID)
				ctx.JSON(200, map[string]string{
					"code": "1",
					"msg":  ctx.Tr("repo.modelconvert.manage.create_error2"),
				})
				return
			}
		}
	}

	uuid := uuid.NewV4()
	id := uuid.String()
	modelConvert := &models.AiModelConvert{
		ID:              id,
		Name:            name,
		Description:     desc,
		Status:          string(models.JobWaiting),
		SrcEngine:       SrcEngine,
		RepoId:          ctx.Repo.Repository.ID,
		ModelName:       task.Name,
		ModelVersion:    task.Version,
		ModelId:         modelId,
		ModelPath:       modelPath,
		DestFormat:      DestFormat,
		NetOutputFormat: NetOutputFormat,
		InputShape:      InputShape,
		InputDataFormat: InputDataFormat,
		UserId:          ctx.User.ID,
	}
	models.SaveModelConvert(modelConvert)
	go goCreateTask(modelConvert, ctx, task)

	ctx.JSON(200, map[string]string{
		"id":   id,
		"code": "0",
	})
}

func isRunningTask(status string) bool {
	stopStatus := []string{"COMPLETED", "STOPPED", "FAILED", "START_FAILED", "STOPPING", "SUCCEEDED"}
	for _, sta := range stopStatus {
		if sta == status {
			return false
		}
	}
	return true
}

func goCreateTask(modelConvert *models.AiModelConvert, ctx *context.Context, task *models.AiModelManage) error {
	if modelConvert.IsGpuTrainTask() {
		log.Info("create gpu train job.")
		return createGpuTrainJob(modelConvert, ctx, task)
	} else {
		//create npu job
		log.Info("create npu train job.")
		return createNpuTrainJob(modelConvert, ctx, task.Path)
	}
}

func createNpuTrainJob(modelConvert *models.AiModelConvert, ctx *context.Context, modelRelativePath string) error {
	VersionOutputPath := "V0001"
	codeLocalPath := setting.JobPath + modelConvert.ID + modelarts.CodePath
	codeObsPath := "/" + setting.Bucket + modelarts.JobPath + modelConvert.ID + modelarts.CodePath
	outputObsPath := "/" + setting.Bucket + modelarts.JobPath + modelConvert.ID + modelarts.OutputPath + VersionOutputPath + "/"
	logObsPath := "/" + setting.Bucket + modelarts.JobPath + modelConvert.ID + modelarts.LogPath + VersionOutputPath + "/"
	dataPath := "/" + modelRelativePath

	_, err := ioutil.ReadDir(codeLocalPath)
	if err == nil {
		deleteLocalDir(codeLocalPath)
	}
	if err := downloadConvertCode(setting.ModelConvert.ConvertRepoPath, codeLocalPath, DefaultBranchName); err != nil {
		log.Error("downloadCode failed, server timed out: %s (%v)", setting.ModelConvert.ConvertRepoPath, err)
		return err
	}
	if err := obsMkdir(setting.CodePathPrefix + modelConvert.ID + modelarts.OutputPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_output: %s (%v)", modelConvert.ID+modelarts.OutputPath, err)
		return err
	}
	if err := obsMkdir(setting.CodePathPrefix + modelConvert.ID + modelarts.LogPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_log: %s (%v)", modelConvert.ID+modelarts.LogPath, err)
		return err
	}
	if err := uploadCodeToObs(codeLocalPath, modelConvert.ID, ""); err != nil {
		log.Error("Failed to uploadCodeToObs: %s (%v)", modelConvert.ID, err)
		return err
	}
	deleteLocalDir(codeLocalPath)
	intputshape := strings.Split(modelConvert.InputShape, ",")
	n := "256"
	c := "1"
	h := "28"
	w := "28"
	if len(intputshape) == 4 {
		n = intputshape[0]
		c = intputshape[1]
		h = intputshape[2]
		w = intputshape[3]
	}

	var engineId int64
	engineId = int64(setting.ModelConvert.NPU_MINDSPORE_IMAGE_ID)
	bootfile := setting.ModelConvert.MindsporeBootFile
	if modelConvert.SrcEngine == TENSORFLOW_ENGINE {
		engineId = int64(setting.ModelConvert.NPU_TENSORFLOW_IMAGE_ID)
		bootfile = setting.ModelConvert.TensorFlowNpuBootFile
	}
	userCommand := "/bin/bash /home/work/run_train.sh 's3://" + codeObsPath + "' 'code/" + bootfile + "' '/tmp/log/train.log' --'data_url'='s3://" + dataPath + "' --'train_url'='s3://" + outputObsPath + "'"
	userCommand += " --'model'='" + modelConvert.ModelPath + "'"
	userCommand += " --'n'='" + fmt.Sprint(n) + "'"
	userCommand += " --'c'='" + fmt.Sprint(c) + "'"
	userCommand += " --'h'='" + fmt.Sprint(h) + "'"
	userCommand += " --'w'='" + fmt.Sprint(w) + "'"

	req := &modelarts.GenerateTrainJobReq{
		JobName:          modelConvert.ID,
		DisplayJobName:   modelConvert.Name,
		DataUrl:          dataPath,
		Description:      modelConvert.Description,
		CodeObsPath:      codeObsPath,
		BootFileUrl:      codeObsPath + bootfile,
		BootFile:         bootfile,
		TrainUrl:         outputObsPath,
		FlavorCode:       setting.ModelConvert.NPU_FlavorCode,
		WorkServerNumber: 1,
		IsLatestVersion:  modelarts.IsLatestVersion,
		EngineID:         engineId,
		LogUrl:           logObsPath,
		PoolID:           setting.ModelConvert.NPU_PoolID,
		//Parameters:       param,
		BranchName:   DefaultBranchName,
		UserImageUrl: setting.ModelConvert.NPU_MINDSPORE_16_IMAGE,
		UserCommand:  userCommand,
	}
	result, err := modelarts.GenerateModelConvertTrainJob(req)
	if err == nil {
		log.Info("jobId=" + fmt.Sprint(result.JobID) + " versionid=" + fmt.Sprint(result.VersionID))
		models.UpdateModelConvertModelArts(modelConvert.ID, fmt.Sprint(result.JobID), fmt.Sprint(result.VersionID))
	} else {
		log.Info("create modelarts taks failed.error=" + err.Error())
		models.UpdateModelConvertFailed(modelConvert.ID, "FAILED", err.Error())
	}
	return err
}

func downloadConvertCode(repopath string, codePath, branchName string) error {
	//add "file:///" prefix to make the depth valid
	if err := git.Clone(repopath, codePath, git.CloneRepoOptions{Branch: branchName, Depth: 1}); err != nil {
		log.Error("Failed to clone repository: %s (%v)", repopath, err)
		return err
	}
	log.Info("srcPath=" + repopath + " codePath=" + codePath)
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
			originUrl := "\turl = " + repopath + "\n"
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

func downloadFromObsToLocal(task *models.AiModelManage, localPath string) error {
	path := Model_prefix + models.AttachmentRelativePath(task.ID) + "/"
	allFile, err := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, path)
	if err == nil {
		_, errState := os.Stat(localPath)
		if errState != nil {
			if err = os.MkdirAll(localPath, os.ModePerm); err != nil {
				return err
			}
		}
		for _, oneFile := range allFile {
			if oneFile.IsDir {
				log.Info(" dir name:" + oneFile.FileName)
			} else {
				allFileName := localPath + "/" + oneFile.FileName
				index := strings.LastIndex(allFileName, "/")
				if index != -1 {
					parentDir := allFileName[0:index]
					if err = os.MkdirAll(parentDir, os.ModePerm); err != nil {
						log.Info("make dir may be error," + err.Error())
					}
				}
				fDest, err := os.Create(allFileName)
				if err != nil {
					log.Info("create file error, download file failed: %s\n", err.Error())
					return err
				}
				body, err := storage.ObsDownloadAFile(setting.Bucket, path+oneFile.FileName)
				if err != nil {
					log.Info("download file failed: %s\n", err.Error())
					return err
				} else {
					defer body.Close()
					p := make([]byte, 1024)
					var readErr error
					var readCount int
					// 读取对象内容
					for {
						readCount, readErr = body.Read(p)
						if readCount > 0 {
							fDest.Write(p[:readCount])
						}
						if readErr != nil {
							break
						}
					}
				}
			}
		}
	} else {
		log.Info("error,msg=" + err.Error())
		return err
	}
	return nil
}

func createGpuTrainJob(modelConvert *models.AiModelConvert, ctx *context.Context, model *models.AiModelManage) error {
	modelRelativePath := model.Path
	command := ""
	IMAGE_URL := setting.ModelConvert.GPU_PYTORCH_IMAGE
	dataActualPath := setting.Attachment.Minio.RealPath + modelRelativePath
	if model.Type == models.TypeCloudBrainTwo {
		//如果模型在OBS上，需要下载到本地，并上传到minio中
		relatetiveModelPath := setting.JobPath + modelConvert.ID + "/dataset"
		log.Info("local dataset path:" + relatetiveModelPath)
		downloadFromObsToLocal(model, relatetiveModelPath)
		uploadCodeToMinio(relatetiveModelPath+"/", modelConvert.ID, "/dataset/")
		deleteLocalDir(relatetiveModelPath)
		dataActualPath = setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + "/" + setting.CBCodePathPrefix + modelConvert.ID + "/dataset"
	}
	log.Info("dataActualPath=" + dataActualPath)
	bootfile := ""
	runParms := make(map[string]interface{}, 0)
	if modelConvert.SrcEngine == PYTORCH_ENGINE {
		if modelConvert.DestFormat == CONVERT_FORMAT_ONNX {
			bootfile = setting.ModelConvert.PytorchOnnxBootFile
			runParms = getGpuModelConvertRunParams(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.PytorchOnnxBootFile)
			//command = getGpuModelConvertCommand(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.PytorchOnnxBootFile)
		} else if modelConvert.DestFormat == CONVERT_FORMAT_TRT {
			bootfile = setting.ModelConvert.PytorchTrTBootFile
			runParms = getGpuModelConvertRunParams(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.PytorchTrTBootFile)
			//command = getGpuModelConvertCommand(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.PytorchTrTBootFile)
		} else {
			return errors.New("Not support the format.")
		}
	} else if modelConvert.SrcEngine == TENSORFLOW_ENGINE {
		IMAGE_URL = setting.ModelConvert.GPU_TENSORFLOW_IMAGE
		if modelConvert.DestFormat == CONVERT_FORMAT_ONNX {
			bootfile = setting.ModelConvert.TensorFlowGpuBootFile
			runParms = getGpuModelConvertRunParams(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.TensorFlowGpuBootFile)
			//command = getGpuModelConvertCommand(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.TensorFlowGpuBootFile)
		} else {
			return errors.New("Not support the format.")
		}
	} else if modelConvert.SrcEngine == PADDLE_ENGINE {
		IMAGE_URL = setting.ModelConvert.GPU_PADDLE_IMAGE
		if modelConvert.DestFormat == CONVERT_FORMAT_ONNX {
			bootfile = setting.ModelConvert.PaddleOnnxBootFile
			runParms = getGpuModelConvertRunParams(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.PaddleOnnxBootFile)
			//command = getGpuModelConvertCommand(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.PaddleOnnxBootFile)
		} else {
			return errors.New("Not support the format.")
		}
	} else if modelConvert.SrcEngine == MXNET_ENGINE {
		IMAGE_URL = setting.ModelConvert.GPU_MXNET_IMAGE
		if modelConvert.DestFormat == CONVERT_FORMAT_ONNX {
			bootfile = setting.ModelConvert.MXnetOnnxBootFile
			runParms = getGpuModelConvertRunParams(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.MXnetOnnxBootFile)
			//command = getGpuModelConvertCommand(modelConvert.ID, modelConvert.ModelPath, modelConvert, setting.ModelConvert.MXnetOnnxBootFile)
		} else {
			return errors.New("Not support the format.")
		}
	}

	log.Info("command=" + command)
	codePath := setting.JobPath + modelConvert.ID + CodeMountPath
	codeTmpPath := setting.JobPath + modelConvert.ID + CodeMountPath + "tmp"
	uploader := storage_helper.SelectStorageHelperFromStorageType(entity.MINIO)
	codeRemoteDir := path.Join(uploader.GetJobDefaultObjectKeyPrefix(modelConvert.ID), "code")
	log.Info("codePath=" + codePath)
	log.Info("codeTmpPath=" + codeTmpPath)
	log.Info("codeRemoteDir=" + codeRemoteDir)

	downloadConvertCode(setting.ModelConvert.ConvertRepoPath, codeTmpPath, DefaultBranchName)

	Zip(codePath+"/master.zip", codeTmpPath)

	uploadCodeToMinio(codePath+"/", modelConvert.ID, CodeMountPath+"/")

	deleteLocalDir(codePath)
	deleteLocalDir(codeTmpPath)

	minioCodePath := setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + "/" + setting.CBCodePathPrefix + modelConvert.ID + "/code"
	log.Info("minio codePath=" + minioCodePath)

	modelPath := setting.JobPath + modelConvert.ID + ModelMountPath + "/"
	log.Info("local modelPath=" + modelPath)
	mkModelPath(modelPath)

	uploadCodeToMinio(modelPath, modelConvert.ID, ModelMountPath+"/")
	deleteLocalDir(modelPath)

	minioModelPath := setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + "/" + setting.CBCodePathPrefix + modelConvert.ID + "/model"
	log.Info("minio model path=" + minioModelPath)

	datasetRemoteDir := path.Join(uploader.GetJobDefaultObjectKeyPrefix(modelConvert.ID), "dataset")

	outputRemoteDir := path.Join(uploader.GetJobDefaultObjectKeyPrefix(modelConvert.ID), "model")

	datasetDirectoryObjectKey := datasetRemoteDir
	if !strings.HasSuffix(datasetRemoteDir, "/") {
		datasetDirectoryObjectKey = datasetRemoteDir + "/"
	}
	codeObjectKey := codeRemoteDir + "/master.zip"
	log.Info("codeObjectKey=" + codeObjectKey)
	log.Info("uploader.GetRealPath(codeObjectKey)=" + uploader.GetRealPath(codeObjectKey))

	req := entity.CreateTrainTaskRequest{
		Name:           modelConvert.ID,
		DisplayJobName: modelConvert.Name,
		Description:    "",
		TaskConfig:     getGrampusTrainTaskConfig(),
		Tasks: []entity.TrainTask{{
			Command:        command,
			Name:           modelConvert.ID,
			ResourceSpecId: setting.ModelConvert.GPU_Resource_Specs_ID, //toDO
			ImageId:        "",                                         //form.ImageID,
			ImageUrl:       IMAGE_URL,
			Datasets: []entity.ContainerData{
				entity.ContainerData{
					ContainerPath:   DataSetMountPath,
					Name:            "dataset",
					ReadOnly:        false,
					ObjectKey:       datasetDirectoryObjectKey,
					RealPath:        uploader.GetRealPath(datasetRemoteDir),
					Bucket:          uploader.GetBucket(),
					EndPoint:        uploader.GetEndpoint(),
					GetBackEndpoint: uploader.GetEndpoint(),
					IsDir:           true,
					StorageType:     entity.MINIO,
				},
			},
			Code: []entity.ContainerData{
				entity.ContainerData{
					Name:          strings.ToLower(ctx.Repo.Repository.Name),
					Bucket:        uploader.GetBucket(),
					EndPoint:      uploader.GetEndpoint(),
					ObjectKey:     codeObjectKey,
					ReadOnly:      true,
					ContainerPath: "/tmp/code/master.zip",
					RealPath:      uploader.GetRealPath(codeObjectKey),
					IsDir:         true,
					S3DownloadUrl: uploader.GetS3DownloadUrl(codeObjectKey),
					StorageType:   entity.MINIO,
				},
			},
			Queues: []models.ResourceQueue{
				models.ResourceQueue{
					AiCenterCode: setting.ModelConvert.GPU_AiCenter_Code,
				},
			},
			PreTrainModel: nil,
			BootFile:      bootfile,
			OutPut: []entity.ContainerData{{
				ContainerPath:   "/tmp/output",
				ReadOnly:        false,
				ObjectKey:       outputRemoteDir,
				RealPath:        uploader.GetRealPath(outputRemoteDir),
				Bucket:          uploader.GetBucket(),
				EndPoint:        uploader.GetEndpoint(),
				GetBackEndpoint: uploader.GetEndpoint(),
				IsDir:           true,
				StorageType:     entity.MINIO,
			}},
			Params: models.Parameters{},
			Spec: &models.Specification{
				ID:              setting.ModelConvert.GPU_Spec_ID,
				SourceSpecId:    setting.ModelConvert.GPU_Resource_Specs_ID,
				AccCardsNum:     1,
				AccCardType:     setting.ModelConvert.GPU_Spec_AccCardType,
				ComputeResource: "GPU",
				AiCenterCode:    setting.ModelConvert.GPU_AiCenter_Code,
			},
			RepoName:         ctx.Repo.Repository.Name,
			WorkServerNumber: 1,
		}},
	}

	reqJson, _ := json.Marshal(req)
	log.Info("reqJson=" + string(reqJson))

	jobResult, err := createGrampusTrainJob(req, command, runParms)

	if err != nil {
		log.Error("CreateJob failed:", err.Error(), ctx.Data["MsgID"])
		models.UpdateModelConvertFailed(modelConvert.ID, "FAILED", err.Error())
		return err
	}
	jobResultJson, _ := json.Marshal(jobResult)
	log.Info("jobResultJson=" + string(jobResultJson))

	if jobResult.ErrorCode != 0 {
		log.Error("CreateJob(%s) failed:%s", modelConvert.ID, jobResult.ErrorMsg, ctx.Data["MsgID"])
		models.UpdateModelConvertFailed(modelConvert.ID, "FAILED", jobResult.ErrorMsg)
		return errors.New(jobResult.ErrorMsg)
	}

	var jobID = jobResult.JobInfo.JobID
	log.Info("jobId=" + jobID)
	models.UpdateModelConvertCBTI(modelConvert.ID, jobID)

	return nil
}

func getGrampusTrainTaskConfig() *entity.AITaskBaseConfig {
	codePath := "/tmp/code"
	datasetPath := "/tmp/dataset"
	pretrainModelPath := "/tmp/pretrainmodel"
	outputPath := "/tmp/output"
	var config = &entity.AITaskBaseConfig{
		ContainerSteps: map[entity.ContainerDataType]*entity.ContainerBuildOpts{
			entity.ContainerCode: {
				ContainerPath:       codePath,
				StorageRelativePath: cloudbrain.CodeMountPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.MINIO, entity.OBS},
			},
			entity.ContainerDataset: {
				ContainerPath:     datasetPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.MINIO, entity.OBS},
			},
			entity.ContainerPreTrainModel: {
				ContainerPath:     pretrainModelPath,
				ReadOnly:          true,
				AcceptStorageType: []entity.StorageType{entity.MINIO, entity.OBS},
			},
			entity.ContainerOutPutPath: {
				ContainerPath:       outputPath,
				StorageRelativePath: cloudbrain.ModelMountPath,
				ReadOnly:            false,
				AcceptStorageType:   []entity.StorageType{entity.MINIO},
				MKDIR:               false,
			},
		},
	}
	config.ActionType = models.ActionCreateGrampusGPUTrainTask
	config.IsActionUseJobId = true
	return config
}

func createGrampusTrainJob(req entity.CreateTrainTaskRequest, exeCommand string, runParam map[string]interface{}) (*models.CreateGrampusJobResponse, error) {
	jobResult, err := grampus.CreateJob(convertTrainReq2Grampus(req, exeCommand, runParam), nil)
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	return jobResult, nil
}

func convertTrainReq2Grampus(req entity.CreateTrainTaskRequest, exeCommand string, runParam map[string]interface{}) models.CreateGrampusJobRequest {
	//command := generateGrampusTrainCommand(req, exeCommand)
	command := ""
	tasks := make([]models.GrampusTasks, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		tasks[i] = convertTrainTask2Grampus(t, command, runParam)
	}

	return models.CreateGrampusJobRequest{Name: req.Name, Tasks: tasks}
}

func convertTrainTask2Grampus(t entity.TrainTask, command string, runParam map[string]interface{}) models.GrampusTasks {
	return models.GrampusTasks{
		Name:             t.Name,
		ResourceSpecId:   t.ResourceSpecId,
		ImageId:          t.ImageId,
		ImageUrl:         t.ImageUrl,
		Datasets:         convertContainerArray2GrampusArray(t.Datasets),
		Code:             convertContainerArray2Grampus(t.Code),
		Command:          command,
		CenterID:         []string{t.Queues[0].AiCenterCode},
		ReplicaNum:       1,
		Models:           convertContainerArray2GrampusArray(t.PreTrainModel),
		BootFile:         t.BootFile,
		OutPut:           convertContainerArray2Grampus(t.OutPut),
		WorkServerNumber: t.WorkServerNumber,
		RunParams:        runParam,
	}
}

func convertContainerArray2GrampusArray(containerDatas []entity.ContainerData) []models.GrampusDataset {
	res := make([]models.GrampusDataset, len(containerDatas))
	for i := 0; i < len(containerDatas); i++ {
		d := containerDatas[i]
		res[i] = convertContainer2Grampus(d)
	}
	return res
}

func convertContainerArray2Grampus(containerDatas []entity.ContainerData) models.GrampusDataset {
	res := models.GrampusDataset{}
	if containerDatas != nil && len(containerDatas) > 0 {
		res = convertContainer2Grampus(containerDatas[0])
	}
	return res
}

func convertContainer2Grampus(d entity.ContainerData) models.GrampusDataset {
	return models.GrampusDataset{
		Name:            d.Name,
		Bucket:          d.Bucket,
		EndPoint:        d.EndPoint,
		ObjectKey:       d.ObjectKey,
		ContainerPath:   d.ContainerPath,
		ReadOnly:        d.ReadOnly,
		GetBackEndpoint: d.GetBackEndpoint,
		Size:            d.Size,
		SizeLimit:       d.SizeLimit,
	}
}

func generateGrampusTrainCommand(req entity.CreateTrainTaskRequest, exeCommand string) string {
	t := req.Tasks[0]
	containerConfig := req.TaskConfig
	computeResource := t.Spec.ComputeResource
	var codePath = containerConfig.GetContainerPath(entity.ContainerCode)
	var modelPath = containerConfig.GetContainerPath(entity.ContainerPreTrainModel)
	var datasetPath = containerConfig.GetContainerPath(entity.ContainerDataset)
	var outputPath = containerConfig.GetContainerPath(entity.ContainerOutPutPath)

	builder := &entity.CommandBuilder{}
	builder.
		//mkdir dirs
		Add(buildMkdirCommand(codePath, modelPath, datasetPath, outputPath)).
		//unzip code
		Add(buildUnzipCodeCommand(codePath, t.Code[0].ContainerPath, computeResource)).
		//unzip dataset
		Add(buildUnzipDatasetCommand(t.Datasets, datasetPath, computeResource)).
		//export
		Add(buildExportCommand(req.Name, computeResource)).
		//exec code
		Add(buildExeCommand(exeCommand))

	return builder.ToString()
}
func buildExeCommand(exeCommand ...string) *entity.CommandBuilder {
	builder := &entity.CommandBuilder{}
	for _, dir := range exeCommand {
		builder.Next(entity.NewCommand(dir))
	}
	return builder
}

func buildMkdirCommand(dirs ...string) *entity.CommandBuilder {
	builder := &entity.CommandBuilder{}
	for _, dir := range dirs {
		builder.Next(entity.NewCommand("mkdir", "-p", dir))
	}
	return builder
}

func buildUnzipCodeCommand(codeConfigPath, codeFilePath, computeSource string) *entity.CommandBuilder {
	builder := &entity.CommandBuilder{}
	if computeSource == models.NPU {
		return builder
	}
	builder.
		Next(entity.NewCommand("echo", "'start to unzip code'")).
		Next(entity.NewCommand("cd", codeConfigPath)).
		Next(entity.NewCommand("unzip", "-q", codeFilePath)).
		Next(entity.NewCommand("echo", "'unzip code finished'")).
		Next(entity.NewCommand("ls", "-l"))
	return builder
}
func buildUnzipDatasetCommand(datasets []entity.ContainerData, datasetPath, computeSource string) *entity.CommandBuilder {
	builder := &entity.CommandBuilder{}
	if computeSource == models.NPU {
		return builder
	}
	if len(datasets) == 0 {
		return nil
	}
	builder.Next(entity.NewCommand("cd", datasetPath)).
		Next(entity.NewCommand("echo", "'start to unzip datasets'"))

	fileDatasets := make([]entity.ContainerData, 0)
	for _, dataset := range datasets {
		if !dataset.IsDir {
			fileDatasets = append(fileDatasets, dataset)
		}
	}
	//单数据集
	if len(fileDatasets) == 1 {
		if strings.HasSuffix(fileDatasets[0].Name, ".tar.gz") {
			builder.Next(entity.NewCommand("tar", "--strip-components=1", "-zxvf", "'"+fileDatasets[0].Name+"'"))
		} else {
			builder.Next(entity.NewCommand("unzip", "-q", "'"+fileDatasets[0].Name+"'"))
		}
		builder.Next(entity.NewCommand("ls", "-l"))
		builder.Next(entity.NewCommand("echo", "'unzip datasets finished'"))
		return builder
	}
	//多数据集
	for i := 0; i < len(fileDatasets); i++ {
		name := fileDatasets[i].Name
		if strings.HasSuffix(name, ".tar.gz") {
			builder.Next(entity.NewCommand("tar", "-zxvf", name))
		} else {
			builder.Next(entity.NewCommand("unzip", "-q", "'"+name+"'", "-d", "'./"+strings.TrimSuffix(name, ".zip")+"'"))
		}
	}
	builder.Next(entity.NewCommand("ls", "-l"))
	builder.Next(entity.NewCommand("echo", "'unzip datasets finished'"))
	return builder
}

func buildExportCommand(jobName, computeResource string) *entity.CommandBuilder {
	builder := &entity.CommandBuilder{}

	if computeResource == models.NPU {
		outputRemotePath := setting.CodePathPrefix + jobName + modelarts.OutputPath
		builder.Next(entity.NewCommand("export", "bucket="+setting.Grampus.Env, "&&", "export", "remote_path="+outputRemotePath))
	} else {
		outputRemotePath := setting.CBCodePathPrefix + jobName + cloudbrain.ModelMountPath + "/"
		builder.Next(entity.NewCommand("export", "env="+setting.Grampus.Env, "&&", "export", "remote_path="+outputRemotePath))
	}
	return builder
}

func deleteLocalDir(dirpath string) {
	//TODO delete
	_err := os.RemoveAll(dirpath)
	if _err == nil {
		log.Info("Delete local file:" + dirpath)
	} else {
		log.Info("Delete local file error: path=" + dirpath)
	}
}

func getGpuModelConvertCommand(name string, modelFile string, modelConvert *models.AiModelConvert, bootfile string) string {
	var command string

	inputshape := strings.Split(modelConvert.InputShape, ",")
	n := "256"
	c := "1"
	h := "28"
	w := "28"
	if len(inputshape) == 4 {
		n = inputshape[0]
		c = inputshape[1]
		h = inputshape[2]
		w = inputshape[3]
	}
	command += "list -all /tmp/code;list -all /tmp/dataset;python3 /tmp/code/" + bootfile + " --model " + modelFile + " --n " + n + " --c " + c + " --h " + h + " --w " + w
	if modelConvert.DestFormat == CONVERT_FORMAT_TRT {
		if modelConvert.NetOutputFormat == NetOutputFormat_FP16 {
			command += " --fp16 True"
		}
	}
	command += " > /tmp/output/" + name + "-" + LogFile
	return command
}

func getGpuModelConvertRunParams(name string, modelFile string, modelConvert *models.AiModelConvert, bootfile string) map[string]interface{} {
	re := make(map[string]interface{}, 0)
	inputshape := strings.Split(modelConvert.InputShape, ",")
	n := "256"
	c := "1"
	h := "28"
	w := "28"
	if len(inputshape) == 4 {
		n = inputshape[0]
		c = inputshape[1]
		h = inputshape[2]
		w = inputshape[3]
	}
	re["model"] = modelFile
	re["n"] = n
	re["c"] = c
	re["h"] = h
	re["w"] = w
	if modelConvert.DestFormat == CONVERT_FORMAT_TRT {
		if modelConvert.NetOutputFormat == NetOutputFormat_FP16 {
			re["fp16"] = "True"

		}
	}
	return re
}

func DeleteModelConvert(ctx *context.Context) {
	log.Info("delete model convert start.")
	id := ctx.Params(":id")
	task, err := models.QueryModelConvertById(id)
	if err == nil {
		go deleteCloudBrainTask(task)
	}
	err = models.DeleteModelConvertById(id)
	//TODO delete OBS文件及云脑任务
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelmanage/convert_model")
	}
}

func deleteCloudBrainTask(task *models.AiModelConvert) {
	if task.IsGpuTrainTask() {
		log.Info("delete grampus model convert task.")
		_, err := grampus.DeleteJob(&entity.JobIdAndVersionId{JobID: task.CloudBrainTaskId})
		if err != nil {
			log.Error("Delete grampus job failed:%v", err)
		}
		dirPath := setting.CBCodePathPrefix + task.ID + "/"
		err = storage.Attachments.DeleteDir(dirPath)
		if err != nil {
			log.Error("DeleteDir(%s) failed:%v", dirPath, err)
		}
	} else {
		log.Info("delete cloudbrain two resource.")
		_, err := modelarts.DelTrainJob(task.CloudBrainTaskId)
		if err != nil {
			log.Error("DelTrainJob(%s) failed:%v", task.CloudBrainTaskId, err.Error())
		}
		DeleteJobStorage(task.ID)
	}
}

func stopModelConvert(id string) error {
	job, err := models.QueryModelConvertById(id)
	if err != nil {
		return err
	}
	if job.IsGpuTrainTask() {
		_, err = grampus.StopJob(&entity.JobIdAndVersionId{JobID: job.CloudBrainTaskId})
		if err != nil {
			log.Error("Stop cloudbrain Job(%s) failed:%v", job.CloudBrainTaskId, err)
		}
	} else {
		_, err = modelarts.StopTrainJob(job.CloudBrainTaskId, job.ModelArtsVersionId)
		if err != nil {
			log.Error("Stop modelarts Job(%s) failed:%v", job.CloudBrainTaskId, err)
		}
	}
	job.Status = string(models.JobStopped)
	if job.EndTime == 0 {
		job.EndTime = timeutil.TimeStampNow()
	}
	models.ModelConvertSetDuration(job)
	err = models.UpdateModelConvert(job)
	if err != nil {
		log.Error("UpdateModelConvert failed:", err)
		return err
	}
	return nil
}

func StopModelConvertApi(ctx *context.Context) {
	id := ctx.Query("id")
	log.Info("stop model convert start.id=" + id)
	err := stopModelConvert(id)
	if err == nil {
		ctx.JSON(200, map[string]string{
			"code": "0",
			"msg":  "succeed",
		})
	} else {
		ctx.JSON(200, map[string]string{
			"code": "1",
			"msg":  err.Error(),
		})
	}
}

func StopModelConvert(ctx *context.Context) {
	id := ctx.Params(":id")
	log.Info("stop model convert start.id=" + id)
	err := stopModelConvert(id)
	if err != nil {
		ctx.ServerError("Not found task.", err)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelmanage/convert_model")
}

func ShowModelConvertInfo(ctx *context.Context) {
	ctx.Data["ID"] = ctx.Query("id")
	ctx.Data["isModelManage"] = true
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)

	job, err := models.QueryModelConvertById(ctx.Query("id"))
	if err == nil {
		if job.TrainJobDuration == "" {
			job.TrainJobDuration = "00:00:00"
		}
		ctx.Data["task"] = job
	} else {
		ctx.ServerError("Not found task.", err)
		return
	}
	ctx.Data["Name"] = job.Name
	ctx.Data["canDownload"] = isOperModifyOrDelete(ctx, job.UserId)
	user, err := models.GetUserByID(job.UserId)
	if err == nil {
		job.UserName = user.Name
		job.UserRelAvatarLink = user.RelAvatarLink()
	}

	if job.IsGpuTrainTask() {
		ctx.Data["npu_display"] = "none"
		ctx.Data["gpu_display"] = "block"
		if job.CloudBrainTaskId == "" {
			ctx.Data["ExitDiagnostics"] = ""
			ctx.Data["AppExitDiagnostics"] = ""
			ctx.HTML(200, tplModelConvertInfo)
			return
		}
		jobResult, err := grampus.GetJob(&entity.JobIdAndVersionId{JobID: job.CloudBrainTaskId})
		if err != nil {
			log.Info("error:" + err.Error())
			ctx.Data["error"] = err.Error()
			ctx.HTML(200, tplModelConvertInfo)
			return
		}
		jobEvent, err := grampus.GetTrainJobEvents(&entity.JobIdAndVersionId{JobID:job.CloudBrainTaskId})
		jobEventJson, _ := json.Marshal(jobEvent)
		log.Info("jobEventJson=" + string(jobEventJson))
		if jobEvent != nil {
			ctx.Data["AppExitDiagnostics"] = string(jobEventJson)
		}
		if jobResult != nil {
			ctx.Data["ExitDiagnostics"] = jobResult.ExitDiagnostics
			if jobResult.JobInfo.Status == models.GrampusStatusPending {
				job.Status = models.GrampusStatusWaiting
			} else {
				job.Status = strings.ToUpper(jobResult.JobInfo.Status)
			}
			job.StartTime = timeutil.TimeStamp(jobResult.JobInfo.StartedAt)
			job.EndTime = timeutil.TimeStamp(jobResult.JobInfo.CompletedAt)
			if strings.ToUpper(jobResult.JobInfo.Status) != models.GrampusStatusWaiting && jobResult.JobInfo.Status != models.GrampusStatusPending {
				models.ModelConvertSetDuration(job)
				err = models.UpdateModelConvert(job)
				if err != nil {
					log.Error("UpdateJob failed:", err)
				}
			}
		}
	} else {
		if job.CloudBrainTaskId != "" {
			result, err := modelarts.GetTrainJob(job.CloudBrainTaskId, job.ModelArtsVersionId)
			if err != nil {
				log.Info("error:" + err.Error())
				ctx.Data["error"] = err.Error()
				return
			}
			job.Status = modelarts.TransTrainJobStatus(result.IntStatus)
			job.RunTime = result.Duration / 1000
			job.TrainJobDuration = models.ConvertDurationToStr(job.RunTime)
			err = models.UpdateModelConvert(job)
			if err != nil {
				log.Error("UpdateJob failed:", err)
			}
		}
		ctx.Data["npu_display"] = "block"
		ctx.Data["gpu_display"] = "none"
		ctx.Data["ExitDiagnostics"] = ""
		ctx.Data["AppExitDiagnostics"] = ""
	}

	ctx.HTML(200, tplModelConvertInfo)
}

func ConvertModelTemplate(ctx *context.Context) {
	ctx.Data["isModelManage"] = true
	ctx.Data["TRAIN_COUNT"] = 0
	SetModelCount(ctx)
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)
	ShowModelConvertPageInfo(ctx)
	ctx.HTML(200, tplModelManageConvertIndex)
}

func ShowModelConvertPageInfo(ctx *context.Context) {
	log.Info("ShowModelConvertInfo start.")
	if !isQueryRight(ctx) {
		log.Info("no right.")
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	modelResult, count, err := GetModelConvertPageData(ctx)
	if err == nil {
		pager := context.NewPagination(int(count), page, pageSize, 5)
		ctx.Data["Page"] = pager
		ctx.Data["Tasks"] = modelResult
		ctx.Data["MODEL_CONVERT_COUNT"] = count
	} else {
		ctx.ServerError("Query data error.", err)
	}
}

func GetModelConvertById(ctx *context.Context) (*models.AiModelConvert, error) {
	id := ctx.Query("id")
	return models.QueryModelConvertById(id)
}

func GetModelConvertByName(ctx *context.Context) ([]*models.AiModelConvert, error) {
	name := ctx.Query("name")
	return models.QueryModelConvertByName(name, ctx.Repo.Repository.ID)
}

func GetModelConvertPageData(ctx *context.Context) ([]*models.AiModelConvert, int64, error) {
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	repoId := ctx.Repo.Repository.ID
	modelResult, count, err := models.QueryModelConvert(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		RepoID: repoId,
	})
	if err != nil {
		log.Info("query db error." + err.Error())
		return nil, 0, err
	}
	userIds := make([]int64, len(modelResult))
	for i, model := range modelResult {
		model.IsCanOper = isOperModifyOrDelete(ctx, model.UserId)
		model.IsCanDelete = isCanDelete(ctx, model.UserId)
		userIds[i] = model.UserId
	}
	userNameMap := models.QueryUserName(userIds)
	for _, model := range modelResult {
		value := userNameMap[model.UserId]
		if value != nil {
			model.UserName = value.Name
			model.UserRelAvatarLink = value.RelAvatarLink()
		}
	}
	return modelResult, count, nil
}

func ModelConvertDownloadModel(ctx *context.Context) {
	log.Info("enter here......")
	id := ctx.Params(":id")
	job, err := models.QueryModelConvertById(id)
	if err != nil {
		ctx.ServerError("Not found task.", err)
		return
	}
	AllDownload := ctx.QueryBool("allDownload")
	if AllDownload {
		if job.IsGpuTrainTask() {
			path := setting.CBCodePathPrefix + job.ID + "/model/"
			allFile, err := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, path)
			if err == nil {
				returnFileName := job.Name + ".zip"
				MinioDownloadManyFile(path, ctx, returnFileName, allFile)
			} else {
				log.Info("error,msg=" + err.Error())
				ctx.ServerError("no file to download.", err)
			}
		} else {
			Prefix := path.Join(setting.TrainJobModelPath, job.ID, "output/", "V0001", "") + "/"
			log.Info("bucket=" + setting.Bucket + "prefix=" + Prefix)
			allFile, err := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, Prefix)
			if err == nil {
				returnFileName := job.Name + ".zip"
				ObsDownloadManyFile(Prefix, ctx, returnFileName, allFile)
			} else {
				log.Info("error,msg=" + err.Error())
				ctx.ServerError("no file to download.", err)
			}
		}
	} else {
		parentDir := ctx.Query("parentDir")
		fileName := ctx.Query("fileName")
		jobName := ctx.Query("jobName")
		if job.IsGpuTrainTask() {
			filePath := "jobs/" + jobName + "/model/" + parentDir
			if !strings.HasSuffix(filePath, fileName) {
				filePath += fileName
			}
			url, err := storage.Attachments.PresignedGetURL(filePath, fileName)
			if err != nil {
				log.Error("PresignedGetURL failed: %v", err.Error(), ctx.Data["msgID"])
				ctx.ServerError("PresignedGetURL", err)
				return
			}
			//ctx.JSON(200, url)
			http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusTemporaryRedirect)
		} else {
			ObjectKey := path.Join(setting.TrainJobModelPath, job.ID, "output/", "V0001", parentDir, fileName)
			log.Info("ObjectKey=" + ObjectKey)
			url, err := storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, ObjectKey)
			if err != nil {
				log.Error("GetObsCreateSignedUrl failed: %v", err.Error(), ctx.Data["msgID"])
				ctx.ServerError("GetObsCreateSignedUrl", err)
				return
			}
			http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusTemporaryRedirect)
		}
	}
}
