package cluster

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/util"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/manager/client/grampus"
	"code.gitea.io/gitea/models"
	model_grampus "code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/schedule"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
)

type C2NetClusterAdapter struct {
}

func init() {
	//注册到一个Map
	AddCluster(entity.C2Net, new(C2NetClusterAdapter))
}

func (c C2NetClusterAdapter) CreateNoteBook(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	newReq, err := convertNoteBookReq2Grampus(req)
	if err != nil {
		log.Error("CreateNoteBook err.req=%+v err=%v", req, err)
		return nil, err
	}
	jobResult, err := grampus.CreateNotebookJob(newReq, trace)
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	if jobResult.ErrorCode > 0 {
		log.Error("CreateNotebookJob err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
		return nil, errors.New(fmt.Sprintf("CreateNotebookJob err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
	}
	return convertGrampus2NoteBookRes(jobResult), nil
}

func (c C2NetClusterAdapter) CreateOnlineInfer(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	log.Info("start to CreateOnlineInfer ")
	newReq, err := convertOnlineInfer2Grampus(req)
	if err != nil {
		log.Error("CreateOnlineInfer err.req=%+v err=%v", req, err)
		return nil, err
	}
	jobResult, err := grampus.CreateInferenceJob(newReq, trace)
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	if jobResult.ErrorCode > 0 {
		log.Error("CreateNotebookJob err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
		return nil, errors.New(fmt.Sprintf("CreateNotebookJob err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
	}
	return convertGrampus2NoteBookRes(jobResult), nil
}

func (c C2NetClusterAdapter) CreateModelExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	log.Info("start to CreateModelExperience ")

	if req.Tasks[0].ExperienceType == "" {

		newReq, err := convertModelExperience2Grampus(req)
		if err != nil {
			log.Error("CreateModelExperience err.req=%+v err=%v", req, err)
			return nil, err
		}
		jobResult, err := grampus.CreateInferenceJob(newReq, trace)
		if err != nil {
			log.Error("CreateNoteBook failed: %v", err.Error())
			return nil, err
		}
		if jobResult.ErrorCode > 0 {
			log.Error("CreateNotebookJob err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
			return nil, errors.New(fmt.Sprintf("CreateNotebookJob err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
		}
		return convertGrampus2NoteBookRes(jobResult), nil
	} else {
		modelSerives, err := grampus.GetAvailableModelServices()
		if err != nil {
			log.Error("Get available Model service err.req=%+v err=%v", req, err)
			return nil, err
		}

		var serviceModel *models.ServiceModel
		for _, v := range modelSerives.ServiceModels {
			if v.ExperienceType == req.Tasks[0].ExperienceType {
				serviceModel = &v
			}
		}
		if serviceModel == nil {
			log.Error("not found matched service model err.req=%+v", req)
			return nil, errors.New("Can not find matched service model.")
		}
		newReq := convertModelAPPExperience2Grampus(req, serviceModel)

		jobResult, err := grampus.CreateModelAppServiceJob(newReq, trace)
		if err != nil {
			log.Error("CreateModelAppService failed: %v", err.Error())
			return nil, err
		}
		if jobResult.ErrorCode > 0 {
			log.Error("CreateModelAppService err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
			return nil, errors.New(fmt.Sprintf("CreateModelAppService err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
		}
		return convertGrampus2NoteBookRes(jobResult), nil

	}

}

func (c C2NetClusterAdapter) CreateSdFinetune(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	log.Info("start to CreateSdFinetune ")
	newReq, err := convertSdFinetune2Grampus(req)
	if err != nil {
		log.Error("CreateModelExperience err.req=%+v err=%v", req, err)
		return nil, err
	}
	jobResult, err := grampus.CreateInferenceJob(newReq, trace)
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	if jobResult.ErrorCode > 0 {
		log.Error("CreateNotebookJob err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
		return nil, errors.New(fmt.Sprintf("CreateNotebookJob err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
	}
	return convertGrampus2NoteBookRes(jobResult), nil
}

func (c C2NetClusterAdapter) CreateComfyuiExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	log.Info("start to CreateComfyuiExperience ")
	newReq, err := convertComfyuiExperience2Grampus(req)
	if err != nil {
		log.Error("CreateComfyuiExperienceerr.req=%+v err=%v", req, err)
		return nil, err
	}
	jobResult, err := grampus.CreateInferenceJob(newReq, trace)
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	if jobResult.ErrorCode > 0 {
		log.Error("CreateNotebookJob err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
		return nil, errors.New(fmt.Sprintf("CreateNotebookJob err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
	}
	return convertGrampus2NoteBookRes(jobResult), nil
}

func (c C2NetClusterAdapter) QueryModelExperience(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	if opts.JobID == "" {
		return nil, errors.New("jobID is empty")
	}
	var result *models.GrampusNotebookResponse
	var err error
	if opts.ComputeResource == models.NPUResource {

		result, err = grampus.GetModelAppServiceJob(&opts)
	} else {
		result, err = grampus.GetNotebookJob(&opts)
	}

	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	return entity.ConvertGrampusNotebookResponse(result.JobInfo), nil
}

func (c C2NetClusterAdapter) StopModelExperience(opts entity.JobIdAndVersionId) error {
	if opts.ComputeResource == models.NPUResource {

		_, err := grampus.StopModelAppServiceJob(&opts)
		return err
	} else {
		return c.StopTrainJob(opts)
	}

}

func (c C2NetClusterAdapter) DeleteModelExperience(opts entity.JobIdAndVersionId) error {
	if opts.ComputeResource == models.NPUResource {

		_, err := grampus.DeleteModelAppServiceJob(&opts)
		return err
	} else {
		return c.DeleteNoteBook(opts)
	}

}

func (c C2NetClusterAdapter) CreateGeneralTask(req entity.CreateGeneralTaskRequest, trace *entity.TraceInfo) (*entity.CreateGeneralTaskResponse, error) {
	log.Info("start to CreateGeneralTask req=%+v", req)
	newReq, err := convertGeneralTaskReq2Grampus(req)
	if err != nil {
		log.Error("CreateGeneralTask err.req=%+v err=%v", req, err)
		return nil, err
	}
	jobResult, err := grampus.CreateNotebookJob(newReq, trace)
	if err != nil {
		log.Error("CreateGeneralTask failed: %v", err.Error())
		return nil, err
	}
	if jobResult.ErrorCode > 0 {
		log.Error("CreateGeneralTask call grampus err.req.Name = %s ErrorCode = %d ErrorMsg = %s", req.Name, jobResult.ErrorCode, jobResult.ErrorMsg)
		return nil, errors.New(fmt.Sprintf("CreateGeneralTask err[%d%s]", jobResult.ErrorCode, jobResult.ErrorMsg))
	}
	return convertGrampus2GeneralTaskRes(jobResult), nil
}

func (c C2NetClusterAdapter) GetNotebookImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	processType := req.ComputeSource.FullName
	images, err := grampus.GetImages(processType, string(req.JobType))
	if err != nil {
		log.Error("GetImages failed:%v", err)
		return nil, false, err
	}
	if images == nil || images.Infos == nil || len(images.Infos) == 0 {
		return nil, true, err
	}

	r := make([]entity.ClusterImage, 0)
	for _, v := range images.Infos {
		if hasIntersection(v.AICenterImage, req.AccCardType, queues) {
			r = append(r, ConvertGrampusImageToStandard(v))
		}
	}
	if len(r) == 0 {
		return nil, false, nil
	}

	return r, false, nil
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

func (c C2NetClusterAdapter) GetTrainImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	return c.GetNotebookImages(req, queues...)
}

func ConvertGrampusImageToStandard(image models.GrampusImage) entity.ClusterImage {
	return entity.ClusterImage{
		ImageId:   image.ID,
		ImageName: image.Name,
	}
}

func convertNoteBookReq2Grampus(req entity.CreateNoteBookTaskRequest) (models.CreateGrampusNotebookRequest, error) {
	codePath := "/code"
	if len(req.Tasks[0].Code) > 0 {
		codePath = req.Tasks[0].Code[0].ContainerPath
		if strings.HasSuffix(codePath, ".zip") {
			codePath = codePath[0:strings.LastIndex(codePath, "/")]
		}
	}

	tasks := make([]models.GrampusNotebookTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertNoteBookTask2Grampus(t, "")
		if !req.IsSubscriber {
			task.NoActAutoShutDownTimeout = 1800000
		}
		if err != nil {
			return models.CreateGrampusNotebookRequest{}, err
		}
		tasks[i] = task
	}

	return models.CreateGrampusNotebookRequest{Name: req.Name, Tasks: tasks}, nil
}

func convertModelExperience2Grampus(req entity.CreateNoteBookTaskRequest) (models.CreateGrampusInferenceRequest, error) {
	command := "echo gpu;"
	tasks := make([]models.GrampusInferenceTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertOnlineInference2Grampus(t, command)
		if err != nil {
			return models.CreateGrampusInferenceRequest{}, nil
		}
		tasks[i] = task
	}
	return models.CreateGrampusInferenceRequest{Name: req.Name, Tasks: tasks}, nil
}

func convertModelAPPExperience2Grampus(req entity.CreateNoteBookTaskRequest, serviceModel *models.ServiceModel) models.CreateGrampusModelAppRequest {

	tasks := make([]models.GrampusModelAppTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {

		task := models.GrampusModelAppTask{
			Name:           req.Name,
			ModelId:        serviceModel.ID,
			ResourceSpecId: req.Tasks[i].Spec.SourceSpecId,
		}

		tasks[i] = task
	}
	return models.CreateGrampusModelAppRequest{Name: req.Name, Tasks: tasks}
}

func convertOnlineInfer2Grampus(req entity.CreateNoteBookTaskRequest) (models.CreateGrampusInferenceRequest, error) {
	userInfo := req.User
	userName, userId := userInfo.Name, userInfo.ID

	command := "echo online"
	command += "; export OPENI_USERNAME=" + userName + " OPENI_USERID=" + strconv.FormatInt(userId, 10) + ";"
	// if req.Tasks[0].Spec.ComputeResource == "GCU" {
	// 	command = "echo gcu;"
	// }
	//:= generateCommand(req.RepoName, req.Tasks[0].BootFile, req.PrimitiveDatasetName)
	tasks := make([]models.GrampusInferenceTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertOnlineInference2Grampus(t, command)
		if err != nil {
			return models.CreateGrampusInferenceRequest{}, nil
		}
		tasks[i] = task
	}

	return models.CreateGrampusInferenceRequest{Name: req.Name, Tasks: tasks}, nil
}

func convertSdFinetune2Grampus(req entity.CreateNoteBookTaskRequest) (models.CreateGrampusInferenceRequest, error) {
	port := req.Tasks[0].EndPoint.Port
	endpoint := req.Tasks[0].EndPoint.EndPoint
	command := "export SD_PORT=" + strconv.FormatInt(port, 10) + " SD_ENDPOINT=" + endpoint + ";"

	tasks := make([]models.GrampusInferenceTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertOnlineInference2Grampus(t, command)
		if err != nil {
			return models.CreateGrampusInferenceRequest{}, nil
		}
		tasks[i] = task
	}
	return models.CreateGrampusInferenceRequest{Name: req.Name, Tasks: tasks}, nil
}

func convertComfyuiExperience2Grampus(req entity.CreateNoteBookTaskRequest) (models.CreateGrampusInferenceRequest, error) {
	port := req.Tasks[0].EndPoint.Port
	endpoint := req.Tasks[0].EndPoint.EndPoint
	command := "export COMFYUI_PORT=" + strconv.FormatInt(port, 10) + " COMFYUI_ENDPOINT=" + endpoint + ";"

	tasks := make([]models.GrampusInferenceTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertOnlineInference2Grampus(t, command)
		if err != nil {
			return models.CreateGrampusInferenceRequest{}, nil
		}
		tasks[i] = task
	}
	return models.CreateGrampusInferenceRequest{Name: req.Name, Tasks: tasks}, nil
}

func convertGeneralTaskReq2Grampus(req entity.CreateGeneralTaskRequest) (models.CreateGrampusNotebookRequest, error) {
	tasks := make([]models.GrampusNotebookTask, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertGeneralTask2Grampus(t)
		if err != nil {
			return models.CreateGrampusNotebookRequest{}, err
		}
		tasks[i] = task
	}

	return models.CreateGrampusNotebookRequest{Name: req.Name, Tasks: tasks}, nil
}

func generateCommand(repoName, bootFile, datasetName string) string {

	//prepare
	//command := "bash && cd /code && unzip master.zip  && cd test-export-data && uvicorn train:app --host 0.0.0.0 --port $OCTOPUS_NOTEBOOK_PORT"
	workDir := "/"
	command := "mkdir /output;pip install gradio fastapi -i https://pypi.tuna.tsinghua.edu.cn/simple;"
	command += "pwd; cd " + workDir + fmt.Sprintf(model_grampus.CommandPrepareScriptGpu)

	//unzip code & dataset
	unZipDatasetCommand := cloudbrainTask.GenerateDatasetUnzipCommand(datasetName)
	bootFile = strings.ReplaceAll(bootFile, "\\", "/")
	bootfilepath := ""
	bootonlyfile := bootFile
	if strings.Index(bootFile, "/") != -1 {
		bootfilepath = bootFile[0:strings.LastIndex(bootFile, "/")]
		if strings.HasPrefix(bootfilepath, "/") {
			bootfilepath = bootfilepath[1:]
		}
		bootonlyfile = bootFile[strings.LastIndex(bootFile, "/")+1:]
	}
	log.Info("bootfilepath=" + bootfilepath + " bootonlyfile=" + bootonlyfile)
	copyDatasetCmd := getCopyCmd(datasetName, repoName, bootfilepath)
	copyDatasetPath := "/code/" + strings.ToLower(repoName) + "/" + bootfilepath
	commandUnzip := "export OPENI_GRADIO_URL=$OCTOPUS_NOTEBOOK_BASE_URL;" + "cd " + workDir + "code;echo \"start unzip code\";unzip -q master.zip; " + copyDatasetCmd + " echo \"start to unzip dataset\";cd " + copyDatasetPath + "; " + unZipDatasetCommand
	//commandUnzip := "cd " + workDir + "code;echo \"start unzip code\";unzip -q master.zip;echo \"start to unzip dataset\";cd " + workDir + "dataset;" + unZipDatasetCommand

	command = command + commandUnzip
	command += "echo \"unzip finished;start to exec code;\";"
	if strings.HasSuffix(bootonlyfile, ".py") {
		bootonlyfile = bootonlyfile[0 : len(bootonlyfile)-3]
	}
	currentTimeNow := time.Now()
	timePostfix := currentTimeNow.Format("20060102_150405")
	command += "cd " + copyDatasetPath + ";echo " + timePostfix + " | tee /output/log_" + timePostfix + ".txt; uvicorn " + bootonlyfile + ":app --log-level debug --host 0.0.0.0 --port $OCTOPUS_NOTEBOOK_PORT 2>&1 | tee -a /output/log_" + timePostfix + ".txt;"

	log.Info("comand=" + command)
	return command
}
func getCopyCmd(datasetName, repoName, bootfilepath string) string {
	cmd := ""
	datasetNameArray := strings.Split(datasetName, ";")
	for _, datasetNameTemp := range datasetNameArray {
		cmd += "cp /dataset/" + datasetNameTemp + " /code/" + strings.ToLower(repoName) + "/" + bootfilepath + ";"
	}
	return cmd
}

func convertOnlineInference2Grampus(t entity.NoteBookTask, command string) (models.GrampusInferenceTask, error) {
	code := models.GrampusDataset{}
	codeArray := convertContainerArray2GrampusArray(t.Code)
	if codeArray != nil && len(codeArray) > 0 {
		code = codeArray[0]
	}
	output := models.GrampusDataset{}
	outputArray := convertContainerArray2GrampusArray(t.OutPut)
	if outputArray != nil && len(outputArray) > 0 {
		output = outputArray[0]
	}
	centerIds, err := getGrampusAvailableCenterIds(t.Queues, t.ImageId, *models.GetComputeSourceInstance(t.Spec.ComputeResource), models.JobTypeDebug)
	log.Info("check centerIds getGrampusAvailableCenterIds ImageId=%s queues=%v centerIds=%v", t.ImageId, t.Queues, centerIds)
	if err != nil {
		log.Error("check centerIds getGrampusAvailableCenterIds err.%v", err)
		return models.GrampusInferenceTask{}, err
	}
	var endPoint []*models.SelfEndPoint
	if t.EndPoint != nil {
		endPoint = append(endPoint, &models.SelfEndPoint{
			Port:     t.EndPoint.Port,
			EndPoint: t.EndPoint.EndPoint,
		})
	}
	return models.GrampusInferenceTask{
		Name:             t.Name,
		ResourceSpecId:   t.Spec.SourceSpecId,
		ImageId:          t.ImageId,
		ImageUrl:         t.ImageUrl,
		Datasets:         convertContainerArray2GrampusArray(t.Datasets),
		PreTrainModel:    convertContainerArray2GrampusArray(t.PreTrainModel),
		Code:             code,
		OutPut:           output,
		EnvVariables:     t.EnvVariables,
		AutoStopDuration: t.AutoStopDuration,
		Capacity:         t.Capacity,
		Command:          command,
		CenterID:         centerIds,
		BootFile:         t.BootFile,
		Endpoints:        endPoint,
		AutoSave:         t.AutoSave,
	}, nil
}

func convertNoteBookTask2Grampus(t entity.NoteBookTask, command string) (models.GrampusNotebookTask, error) {
	code := models.GrampusDataset{}
	codeArray := convertContainerArray2GrampusArray(t.Code)
	if codeArray != nil && len(codeArray) > 0 {
		code = codeArray[0]
	}
	output := models.GrampusDataset{}
	outputArray := convertContainerArray2GrampusArray(t.OutPut)
	if outputArray != nil && len(outputArray) > 0 {
		output = outputArray[0]
	}
	centerIds, err := getGrampusAvailableCenterIds(t.Queues, t.ImageId, *models.GetComputeSourceInstance(t.Spec.ComputeResource), models.JobTypeDebug)
	if err != nil {
		return models.GrampusNotebookTask{}, err
	}
	return models.GrampusNotebookTask{
		Name:             t.Name,
		ResourceSpecId:   t.Spec.SourceSpecId,
		ImageId:          t.ImageId,
		ImageUrl:         t.ImageUrl,
		Datasets:         convertContainerArray2GrampusArray(t.Datasets),
		PreTrainModel:    convertContainerArray2GrampusArray(t.PreTrainModel),
		Code:             code,
		OutPut:           output,
		EnvVariables:     t.EnvVariables,
		AutoStopDuration: t.AutoStopDuration,
		Capacity:         t.Capacity,
		Command:          command,
		CenterID:         centerIds,
		AutoSave:         t.AutoSave,
	}, nil
}

func convertGeneralTask2Grampus(t entity.GeneralTask) (models.GrampusNotebookTask, error) {
	code := models.GrampusDataset{}
	codeArray := convertContainerArray2GrampusArray(t.Code)
	if codeArray != nil && len(codeArray) > 0 {
		code = codeArray[0]
	}
	output := models.GrampusDataset{}
	outputArray := convertContainerArray2GrampusArray(t.OutPut)
	if outputArray != nil && len(outputArray) > 0 {
		output = outputArray[0]
	}
	centerIds, err := getGrampusAvailableCenterIds(t.Queues, t.ImageId, *models.GetComputeSourceInstance(t.Spec.ComputeResource), models.JobTypeDebug)
	if err != nil {
		return models.GrampusNotebookTask{}, err
	}
	return models.GrampusNotebookTask{
		Name:             t.Name,
		ResourceSpecId:   t.Spec.SourceSpecId,
		ImageId:          t.ImageId,
		ImageUrl:         t.ImageUrl,
		Datasets:         convertContainerArray2GrampusArray(t.Datasets),
		PreTrainModel:    convertContainerArray2GrampusArray(t.PreTrainModel),
		Code:             code,
		OutPut:           output,
		EnvVariables:     t.EnvVariables,
		AutoStopDuration: t.AutoStopDuration,
		Capacity:         t.Capacity,
		CenterID:         centerIds,
	}, nil
}

func getGrampusAvailableCenterIds(queues []models.ResourceQueue, imageId string, computeSource models.ComputeSource, jobType models.JobType) ([]string, error) {
	if len(queues) == 0 {
		return []string{}, nil
	}
	var intersectionCenterIds []string
	if imageId == "" {
		for _, queue := range queues {
			code := strings.TrimSuffix(queue.AiCenterCode+"+"+queue.QueueCode, "+")
			intersectionCenterIds = append(intersectionCenterIds, code)
		}
		return intersectionCenterIds, nil
	}

	processType := computeSource.FullName
	log.Info("processType=" + computeSource.FullName + " jobType=" + string(jobType))
	images, err := grampus.GetImages(processType, string(jobType))
	if err != nil {
		log.Warn("can not get image info from grampus", err)
		return []string{}, err
	}
	var imageCenterIds []string
	for _, image := range images.Infos {

		if image.ID == imageId {

			for _, centerInfo := range image.AICenterImage {
				for _, queue := range queues {
					if centerInfo.AiCenterId != queue.AiCenterCode {
						continue
					}
					if centerInfo.AccDeviceModel != "" {
						accCardType := strings.ToUpper(centerInfo.AccDeviceModel)
						if accCardType != queue.AccCardType {
							continue
						}
					}

					if len(centerInfo.PoolIds) == 0 {
						imageCenterIds = append(imageCenterIds, centerInfo.AiCenterId+"+"+queue.QueueCode)
						continue
					}
					for _, poolId := range centerInfo.PoolIds {
						if poolId == queue.QueueCode {
							imageCenterIds = append(imageCenterIds, centerInfo.AiCenterId+"+"+queue.QueueCode)
							continue
						}
					}
				}
			}
			break
		}
	}
	images, err = grampus.GetUserImages(processType, string(jobType))
	if err == nil {
		for _, image := range images.Infos {
			// 加多一个判断，用户镜像
			if image.ID == imageId {
				for _, centerInfo := range image.AICenterImage {
					for _, queue := range queues {
						if centerInfo.AiCenterId != queue.AiCenterCode {
							continue
						}
						if centerInfo.AccDeviceModel != "" {
							accCardType := strings.ToUpper(centerInfo.AccDeviceModel)
							if accCardType != queue.AccCardType {
								continue
							}
						}
						if len(centerInfo.PoolIds) == 0 {
							imageCenterIds = append(imageCenterIds, centerInfo.AiCenterId+"+"+queue.QueueCode)
							continue
						}
						for _, poolId := range centerInfo.PoolIds {
							if poolId == queue.QueueCode {
								imageCenterIds = append(imageCenterIds, centerInfo.AiCenterId+"+"+queue.QueueCode)
								continue
							}
						}
					}
				}
				break
			}
		}
	} else {
		log.Warn("can not get user image info from grampus", err)
	}
	if len(imageCenterIds) == 0 {
		return []string{}, errors.New("image not available")
	}

	imageCenterIds = util.UniqueStr(imageCenterIds)
	log.Info("get available center and queueId = %v", imageCenterIds)
	return imageCenterIds, nil
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

func convertParameters2Grampus(parameters models.Parameters) map[string]interface{} {
	req := make(map[string]interface{})
	for _, param := range parameters.Parameter {
		req[param.Label] = param.Value
	}

	return req
}

func convertContainer2Grampus(d entity.ContainerData) models.GrampusDataset {
	return models.GrampusDataset{
		Name:              d.Name,
		Bucket:            d.Bucket,
		EndPoint:          d.EndPoint,
		ObjectKey:         d.ObjectKey,
		ContainerPath:     d.ContainerPath,
		ReadOnly:          d.ReadOnly,
		GetBackEndpoint:   d.GetBackEndpoint,
		Size:              d.Size,
		IsOverwrite:       d.IsOverwrite,
		IsNeedUnzip:       d.IsNeedUnzip,
		IsNeedTensorboard: d.IsNeedTensorboard,
		Id:                d.Id,
		SizeLimit:         d.SizeLimit,
	}
}

func convertGrampus2NoteBookRes(res *models.GrampusNotebookResponse) *entity.CreateNoteBookTaskResponse {
	jobInfo := res.JobInfo
	return &entity.CreateNoteBookTaskResponse{
		StartedAt:   jobInfo.StartedAt,
		RunSec:      jobInfo.RunSec,
		CompletedAt: jobInfo.CompletedAt,
		CreatedAt:   jobInfo.CreatedAt,
		UpdatedAt:   jobInfo.UpdatedAt,
		Desc:        jobInfo.Desc,
		JobID:       jobInfo.JobID,
		Name:        jobInfo.Name,
		Status:      jobInfo.Status,
		UserID:      jobInfo.UserID,
	}
}

func convertGrampus2GeneralTaskRes(res *models.GrampusNotebookResponse) *entity.CreateGeneralTaskResponse {
	jobInfo := res.JobInfo
	return &entity.CreateGeneralTaskResponse{
		StartedAt:   jobInfo.StartedAt,
		RunSec:      jobInfo.RunSec,
		CompletedAt: jobInfo.CompletedAt,
		CreatedAt:   jobInfo.CreatedAt,
		UpdatedAt:   jobInfo.UpdatedAt,
		Desc:        jobInfo.Desc,
		JobID:       jobInfo.JobID,
		Name:        jobInfo.Name,
		Status:      jobInfo.Status,
		UserID:      jobInfo.UserID,
	}
}

func (c C2NetClusterAdapter) RestartNoteBook(jobId string, autoStopDuration int64, trace *entity.TraceInfo) (*entity.RestartNoteBookTaskResponse, error) {
	res, err := grampus.RestartNotebookJob(jobId, autoStopDuration, trace)
	if err != nil {
		log.Error("RestartNotebookJob err jobId=%s .%v", jobId, err)
		return nil, err
	}
	if res.ErrorCode > 0 {
		log.Error("RestartNotebookJob err.jobId = %s ErrorCode = %d ErrorMsg = %s", jobId, res.ErrorCode, res.ErrorMsg)
		if entity.GrampusJobCanNotRestart.IsMatch(res.ErrorCode) {
			return nil, errors.New(entity.GrampusJobCanNotRestart.CodeTrCode)
		}
		if entity.GrampusJobNotExistInCenter.IsMatch(res.ErrorCode) {
			return nil, errors.New(entity.GrampusJobNotExistInCenter.CodeTrCode)
		}
		return nil, errors.New(response.RESTART_FAILED.TrCode)
	}
	return convertToCreateNoteBookTaskResponse(res), nil
}

func convertToCreateNoteBookTaskResponse(res *models.GrampusNotebookRestartResponse) *entity.RestartNoteBookTaskResponse {
	return &entity.RestartNoteBookTaskResponse{
		JobId:  res.NewId,
		Status: res.Status,
	}
}

func (c C2NetClusterAdapter) DeleteNoteBook(opts entity.JobIdAndVersionId) error {
	_, err := grampus.DeleteJob(&opts)
	if err != nil {
		log.Error("DeleteNoteBook(%s) failed:%v", opts, err)
		log.Info("error=" + err.Error())
		return nil
	}
	return nil
}

func (c C2NetClusterAdapter) StopNoteBook(opts entity.JobIdAndVersionId) error {
	opts.JobType = string(models.JobTypeDebug)
	_, err := grampus.StopJob(&opts)
	if err != nil {
		log.Error("StopNoteBook(%s) failed:%v", opts, err)
		return err
	}
	return nil
}

func (c C2NetClusterAdapter) QueryNoteBook(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	if opts.JobID == "" {
		return nil, errors.New("jobID is empty")
	}
	result, err := grampus.GetNotebookJob(&opts)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	return entity.ConvertGrampusNotebookResponse(result.JobInfo), nil
}
func (c C2NetClusterAdapter) QueryOnlineInfer(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	if opts.JobID == "" {
		return nil, errors.New("jobID is empty")
	}
	result, err := grampus.GetInferenceJob(&opts)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	return entity.ConvertGrampusNotebookResponse(result.JobInfo), nil
}

func (c C2NetClusterAdapter) QueryNoteBookByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	res, err := grampus.GetJobListByJobName(jobName)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.QueryTaskResponse, 0)
	if res != nil {
		for i := 0; i < len(res.JobInfos); i++ {
			if res.JobInfos[i].Name == jobName {
				result = append(result, entity.ConvertGrampusTrainResponse(res.JobInfos[i]))
			}

		}
	}
	return result, nil
}

func (c C2NetClusterAdapter) GetNoteBookLog(jobId string) (*entity.ClusterLog, error) {
	return nil, nil
}

func (c C2NetClusterAdapter) GetNoteBookUrl(jobId string) (string, error) {
	res, err := c.QueryNoteBook(entity.JobIdAndVersionId{JobID: jobId})
	if err != nil {
		return "", err
	}
	if res.Url == "" {
		return "", errors.New("notebook task error")
	}
	ptokenStr := ""
	if res.PToken != "" {
		ptokenStr = "&ptoken=" + res.PToken
	}
	return res.Url + "?token=" + res.Token + ptokenStr, nil
}

func (c C2NetClusterAdapter) GetSelfEndPointUrl(jobId string) (string, error) {
	res, err := c.QueryOnlineInfer(entity.JobIdAndVersionId{JobID: jobId})
	log.Info("grampus res is: %s", res)
	if err != nil {
		return "", err
	}
	if res.Url == "" {
		return "", errors.New("notebook task error")
	}
	cloudbrainTask, bizErr := models.GetCloudbrainByJobID(jobId)
	if bizErr == nil {
		if cloudbrainTask.JobType != string(models.JobTypeOnlineInference) && cloudbrainTask.JobType != string(models.JobTypeModelExperience) && cloudbrainTask.JobType != string(models.JobTypeSdFinetune) && cloudbrainTask.JobType != string(models.JobTypeComfyuiExperience) {
			return "", nil
		}
	} else {
		return "", nil
	}
	url := res.Url
	if res.DomainUrl != "" {
		url = res.DomainUrl
	}
	if !checkUrlCanAccess(url) {
		return "", errors.New("aimodel.loading")
	}

	onlyForOnlineInference(url, jobId)
	return url, nil
}

func checkUrlCanAccess(url string) bool {
	res, _ := http.Get(url)
	if res != nil && res.StatusCode == http.StatusBadGateway {
		log.Error("access debug url failed. res.StatusCode=%d", res.StatusCode)
		return false
	}
	return true
}

func onlyForOnlineInference(url string, jobId string) string {
	cloudbrainTask, bizErr := models.GetCloudbrainByJobID(jobId)
	if bizErr == nil {
		if cloudbrainTask.JobType != string(models.JobTypeOnlineInference) && cloudbrainTask.JobType != string(models.JobTypeSdFinetune) && cloudbrainTask.JobType != string(models.JobTypeComfyuiExperience) {
			return ""
		}
	} else {
		return ""
	}
	tmpurl := url
	if strings.HasPrefix(url, setting.Grampus.NoteBookDomainURL) {

		if strings.HasSuffix(url, "/") {
			tmpurl += "queue/join"
		} else {
			tmpurl += "/queue/join"
		}
		var wsurl string
		if strings.HasPrefix(url, "https") {
			wsurl = strings.Replace(tmpurl, "https", "wss", 1)
		} else {
			wsurl = strings.Replace(tmpurl, "http", "ws", 1)
		}

		grampus.SendMsgToWebsocket(wsurl, "{\"msg\":\"send_hash\"}")
	}
	return ""
}

func (c C2NetClusterAdapter) GetNoteBookOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	if opts.JobID == "" {
		log.Error("jobid is empty")
		return nil, errors.New("jobid is empty")
	}
	jobResult, err := grampus.GetDebugJobEvents(&opts)
	if err != nil {
		log.Error("GetDebugJobEvents failed:%v", err)
		return nil, err
	}

	r := parseC2NetEventsToOperationProfile(jobResult.NotebookEvents)

	getJobResult, err := grampus.GetNotebookJob(&opts)
	if err == nil && getJobResult != nil && getJobResult.ExitDiagnostics != "" {
		r.Events = append(r.Events, entity.ProfileEvent{
			Message: getJobResult.ExitDiagnostics,
			Reason:  "Exit",
		})
	}
	return r, nil
}

func parseC2NetEventsToOperationProfile(notebookEvents []models.GrampusJobEvents) *entity.OperationProfile {
	events := make([]entity.ProfileEvent, 0)
	for i := 0; i < len(notebookEvents); i++ {
		e := notebookEvents[i]
		if e.Message == "" {
			continue
		}
		events = append(events, entity.ProfileEvent{
			Message:   e.Message,
			Reason:    e.Reason,
			Name:      e.Name,
			Timestamp: e.Timestamp,
		})
	}
	return &entity.OperationProfile{Events: events}
}

func (c C2NetClusterAdapter) CreateTrainJob(req entity.CreateTrainTaskRequest, trace *entity.TraceInfo) (*entity.CreateTrainTaskResponse, error) {
	newReq, err := convertTrainReq2Grampus(req)
	if err != nil {
		log.Error("CreateTrainJob err.req=%+v err=%v", req, err)
		return nil, err
	}
	log.Info("name = %s centers = %v", newReq.Name, newReq.Tasks[0].CenterID)
	jobResult, err := grampus.CreateJob(newReq, trace)
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	return convertGrampus2TrainRes(jobResult), nil
}

func convertTrainReq2Grampus(req entity.CreateTrainTaskRequest) (models.CreateGrampusJobRequest, error) {
	command := ""

	tasks := make([]models.GrampusTasks, len(req.Tasks))
	for i := 0; i < len(req.Tasks); i++ {
		t := req.Tasks[i]
		task, err := convertTrainTask2Grampus(t, command)
		if err != nil {
			return models.CreateGrampusJobRequest{}, err
		}
		tasks[i] = task
	}

	return models.CreateGrampusJobRequest{Name: req.Name, Tasks: tasks}, nil
}

func convertTrainTask2Grampus(t entity.TrainTask, command string) (models.GrampusTasks, error) {
	centerIds, err := getGrampusAvailableCenterIds(t.Queues, t.ImageId, *models.GetComputeSourceInstance(t.Spec.ComputeResource), models.JobTypeTrain)
	if err != nil {
		return models.GrampusTasks{}, err
	}

	return models.GrampusTasks{
		Name:             t.Name,
		ResourceSpecId:   t.ResourceSpecId,
		ImageId:          t.ImageId,
		ImageUrl:         t.ImageUrl,
		Datasets:         convertContainerArray2GrampusArray(t.Datasets),
		Code:             convertContainerArray2Grampus(t.Code),
		Command:          command,
		CenterID:         centerIds,
		ReplicaNum:       1,
		Models:           convertContainerArray2GrampusArray(t.PreTrainModel),
		BootFile:         t.BootFile,
		OutPut:           convertContainerArray2Grampus(t.OutPut),
		WorkServerNumber: t.WorkServerNumber,
		RunParams:        convertParameters2Grampus(t.Params),
		EnvVariables:     t.EnvVariables,
	}, nil
}

func convertGrampus2TrainRes(res *models.CreateGrampusJobResponse) *entity.CreateTrainTaskResponse {
	jobInfo := res.JobInfo
	return &entity.CreateTrainTaskResponse{
		StartedAt:   jobInfo.StartedAt,
		RunSec:      jobInfo.RunSec,
		CompletedAt: jobInfo.CompletedAt,
		CreatedAt:   jobInfo.CreatedAt,
		UpdatedAt:   jobInfo.UpdatedAt,
		Desc:        jobInfo.Desc,
		JobID:       jobInfo.JobID,
		Name:        jobInfo.Name,
		Status:      jobInfo.Status,
		UserID:      jobInfo.UserID,
	}
}

func (c C2NetClusterAdapter) DeleteTrainJob(opts entity.JobIdAndVersionId) error {
	_, err := grampus.DeleteJob(&opts)
	if err != nil {
		log.Error("Delete train job(%s) failed:%v", opts, err)
		log.Info("error=" + err.Error())
		return nil
	}
	return nil
}

func (c C2NetClusterAdapter) StopTrainJob(opts entity.JobIdAndVersionId) error {
	_, err := grampus.StopJob(&opts)
	if err != nil {
		log.Error("StopNoteBook(%s) failed:%v", opts, err)
		return err
	}
	return nil
}
func (c C2NetClusterAdapter) QueryTrainJob(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	if opts.JobID == "" {
		return nil, errors.New("jobID is empty")
	}
	result, err := grampus.GetJob(&opts)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return convertGrampusTrainJobResponse(result.JobInfo), nil
}

func (c C2NetClusterAdapter) QueryTrainJobByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	res, err := grampus.GetJobListByJobName(jobName)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.QueryTaskResponse, 0)
	if res != nil {
		for i := 0; i < len(res.JobInfos); i++ {
			if res.JobInfos[i].Name == jobName {
				result = append(result, entity.ConvertGrampusTrainResponse(res.JobInfos[i]))
			}

		}
	}
	return result, nil
}

func (c C2NetClusterAdapter) GetVisualizeUrl(jobId string) (string, error) {
	res, err := c.QueryTrainJob(entity.JobIdAndVersionId{JobID: jobId})
	if err != nil {
		return "", err
	}
	return res.TensorboardEndpoint, nil
}

func convertGrampusTrainJobResponse(job models.GrampusJobInfo) *entity.QueryTaskResponse {
	if len(job.Tasks) == 0 {
		return nil
	}
	task := job.Tasks[0]
	centerId := ""
	if len(task.CenterID) > 0 {
		centerId = task.CenterID[0]
	}
	centerName := ""
	if len(task.CenterName) > 0 {
		centerName = task.CenterName[0]
	}
	return &entity.QueryTaskResponse{
		StartedAt:           timeutil.TimeStamp(job.StartedAt),
		CompletedAt:         timeutil.TimeStamp(job.CompletedAt),
		Status:              job.Status,
		DetailedStatus:      job.DetailedStatus,
		CenterId:            centerId,
		CenterName:          centerName,
		JobId:               job.JobID,
		TensorboardEndpoint: task.TensorboardEndpoint,
		QueueCode:           task.PoolId,
	}
}

func (c C2NetClusterAdapter) GetLog(opts entity.ClusterLogOpts) (*entity.ClusterLog, error) {
	var content string
	var err error
	if opts.JobType == string(models.JobTypeSdFinetune) || opts.JobType == string(models.JobTypeOnlineInference) || opts.JobType == string(models.JobTypeModelExperience) || opts.JobType == string(models.JobTypeComfyuiExperience) {
		log.Info("return JobTypeOnlineInference log.")
		cloudbrain, _ := models.GetCloudbrainByJobID(opts.JobId)
		if cloudbrain != nil && !cloudbrain.IsTerminal() {
			return &entity.ClusterLog{
				//任务还没结束，请结束以后再来查看日志
				Content: "The AI task is still running. Please check the logs once it completes.",
			}, nil
		}
		status, _ := schedule.GetModelScheduleStatus(opts.JobId)
		if status == models.ModelMigrating {
			return &entity.ClusterLog{
				Content: "Hold on! We're moving your log files. This’ll just take a moment!",
			}, nil
		}
		content = getOnlineInferenceLog(opts)
		return &entity.ClusterLog{
			Content: content,
		}, nil

	}
	exitDiagnostics := getGrampusExitDiagnostics(opts.JobId)
	if opts.WorkServerNum > 1 {
		if opts.WorkServerNum < 1 || opts.NodeId > opts.WorkServerNum-1 {
			return nil, errors.New("query parameter is wrong")
		}
		content, err = grampus.GetTrainJobLog(opts.JobId, opts.NodeId)
	} else {
		content, err = grampus.GetTrainJobLog(opts.JobId)
	}
	if err != nil {
		log.Error("GetLog err.opts=%+v,err=%v", opts, err)
		content = ""
	}

	return &entity.ClusterLog{
		Content: content + "\n" + exitDiagnostics,
	}, nil
}

func getOnlineInferenceLog(opts entity.ClusterLogOpts) string {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	//查找日志文件
	files := getLogFilesInStorage(helper, opts.ObjectKeyPrefix, ".txt")
	if len(files) == 0 {
		//此时未找符合条件的文件
		log.Info("Not found file.")
		return ""
	}
	//默认选择第一个文件
	file := files[0]
	opts.Direction = entity.UP
	//计算开始行和结束行
	_, endLine := findStartAndEnd(opts, file.RelativePath, helper)
	//获取日志内容
	result, _, _ := getLogInStorage(1, endLine, helper, file.RelativePath)

	return result
}

func getGrampusExitDiagnostics(jobId string) string {
	result, err := grampus.GetJob(&entity.JobIdAndVersionId{JobID: jobId})
	if err != nil {
		log.Error("GetJob(%s) failed:%v", jobId, err)
		return ""
	}
	if result != nil {
		return result.ExitDiagnostics
	}
	return ""
}

func (c C2NetClusterAdapter) GetLogDownloadInfo(opts entity.ClusterLogDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
	res, err := c.GetLog(entity.ClusterLogOpts{
		JobType:         opts.JobType,
		JobName:         opts.JobName,
		JobId:           opts.JobId,
		NodeId:          opts.NodeId,
		WorkServerNum:   opts.WorkServerNum,
		StorageType:     opts.StorageType,
		ObjectKeyPrefix: opts.ObjectKeyPrefix,
	})
	if err != nil {
		log.Error("error occurs when attempting to get log content.opts=%+v err=%v", opts, err)
		return nil, err
	}
	fileName := opts.JobName + "-log.txt"
	if opts.WorkServerNum > 1 {
		fileName = opts.JobName + "-" + fmt.Sprint(opts.NodeId) + "-log.txt"
	}
	return &entity.FileDownloadInfo{
		Readers:        []entity.FileReader{{Reader: ioutil.NopCloser(strings.NewReader(res.Content))}},
		ResultType:     entity.FileTypeTXT,
		ResultFileName: fileName,
	}, nil
}

func (c C2NetClusterAdapter) GetSingleOutputDownloadInfo(opts entity.ClusterSingleOutputDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	url, err := helper.GetSignedDownloadUrl(opts.Path)
	if err != nil {
		log.Error("GetSignedDownloadUrl err.opts=%+v,err =%v", opts, err)
		return nil, err
	}
	return &entity.FileDownloadInfo{
		DownloadUrl: url,
	}, nil
}

func (c C2NetClusterAdapter) DownloadAllOutput(opts entity.DownloadOutputOpts) error {
	return DownloadAllOutput(opts)
}

func (c C2NetClusterAdapter) GetNodeInfo(opts entity.ClusterNodeInfoOpts) ([]entity.AITaskNodeInfo, error) {
	workServerNum := opts.WorkServerNum
	if opts.WorkServerNum < 1 {
		workServerNum = 1
	}

	res := make([]entity.AITaskNodeInfo, workServerNum)
	for i := 0; i < workServerNum; i++ {
		res[i] = entity.AITaskNodeInfo{
			ID: i,
		}
	}
	return res, nil
}

func (c C2NetClusterAdapter) GetResourceUsage(opts entity.ClusterResourceUsageOpts) (*entity.ResourceUsage, error) {
	var err error

	startTime := opts.StartTime
	endTime := opts.EndTime
	nodeId := opts.NodeId
	jobId := opts.JobId

	if startTime == 0 {
		startTime = time.Now().Unix() - 30*60
	}
	if endTime == 0 {
		endTime = time.Now().Unix()
	}
	var result models.NewModelArtsMetricStatisticResult
	if opts.WorkServerNumber <= 1 {
		result, err = grampus.GetGrampusMetrics(jobId, startTime, endTime)
	} else {
		if nodeId > opts.WorkServerNumber-1 {
			return nil, response.PARAM_ERROR.ToError()
		}
		result, err = grampus.GetGrampusMetrics(opts.JobId, startTime, endTime, nodeId)
	}

	if err != nil {
		log.Error("GetGrampusMetrics error. opts=%+v err= %v", opts, err)
		return nil, err
	}
	return transferGrampusMetrics2Standard(result), nil
}

func transferGrampusMetrics2Standard(result models.NewModelArtsMetricStatisticResult) *entity.ResourceUsage {
	m := make([]entity.MetricsInfo, 0)
	for i := 0; i < len(result.MetricsInfo); i++ {
		m = append(m, entity.MetricsInfo{
			Name:  result.MetricsInfo[i].Metric,
			Value: result.MetricsInfo[i].Value,
		})
	}
	return &entity.ResourceUsage{
		Interval:    int(result.Step / 60),
		MetricsInfo: m,
	}
}

func (c C2NetClusterAdapter) GetTrainJobOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	if opts.JobID == "" {
		log.Error("jobid is empty")
		return nil, errors.New("jobid is empty")
	}
	jobResult, err := grampus.GetTrainJobEvents(&opts)
	if err != nil {
		log.Error("GetTrainJobEvents failed:%v", err)
		return nil, err
	}

	r := parseC2NetEventsToOperationProfile(jobResult.JobEvents)
	getJobResult, err := grampus.GetJob(&opts)
	if err == nil && getJobResult != nil && getJobResult.ExitDiagnostics != "" {
		r.Events = append(r.Events, entity.ProfileEvent{
			Message: getJobResult.ExitDiagnostics,
			Reason:  "Exit",
		})
	}
	return r, nil
}

func (c C2NetClusterAdapter) GetOutput(opts entity.ClusterOutputOpts) (*entity.ClusterAITaskOutput, error) {
	status, err := schedule.GetModelScheduleStatus(opts.JobId)
	if err != nil {
		log.Error("GetModelScheduleStatus(%s) failed:%v", opts.JobId, err)
		return nil, err
	}
	if status != models.ModelMigrateSuccess {
		return &entity.ClusterAITaskOutput{
			Status:   status,
			Path:     opts.ParentDir,
			FileList: []storage.FileInfo{},
		}, nil
	}

	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	fileList, err := helper.GetOneLevelObjectsUnderDir(path.Join(opts.ObjectKeyPrefix, opts.ParentDir))
	if err != nil {
		log.Error("GetOneLevelObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.ObjectKeyPrefix, err)
		return nil, err
	}
	return &entity.ClusterAITaskOutput{
		Status:   status,
		Path:     opts.ParentDir,
		FileList: fileList,
	}, nil
}

func (c C2NetClusterAdapter) GetAllOutput(opts entity.ClusterOutputOpts) (*entity.AllAITaskOutput, error) {
	status, err := schedule.GetModelScheduleStatus(opts.JobId)
	if err != nil {
		log.Error("GetModelScheduleStatus(%s) failed:%v", opts.JobId, err)
		return nil, err
	}
	if status != models.ModelMigrateSuccess {
		return &entity.AllAITaskOutput{FileList: []storage.FileInfo{}}, nil
	}

	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	fileList, err := helper.GetAllObjectsUnderDir(path.Join(opts.ObjectKeyPrefix, opts.ParentDir))
	if err != nil {
		log.Error("GetOneLevelObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.ObjectKeyPrefix, err)
		return nil, err
	}
	return &entity.AllAITaskOutput{FileList: fileList}, nil
}
