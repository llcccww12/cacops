package cluster

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/manager/client/iflytek"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type IFLYTEKTrainingClusterAdapter struct {
}

func init() {
	//注册到一个Map
	AddCluster(entity.IFLYTEKTraining, new(IFLYTEKTrainingClusterAdapter))
}

func (c IFLYTEKTrainingClusterAdapter) CreateNoteBook(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) CreateOnlineInfer(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) CreateModelExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	t := req.Tasks[0]
	if len(t.PreTrainModel) == 0 {
		log.Error("IFLYTEKTrainingClusterAdapter CreateModelExperience error. pretrain model is empty.req=%+v", req)
		return nil, errors.New("pretrain model is empty")
	}
	modelPath := t.PreTrainModel[0].ObjectKey
	modelPath = strings.TrimSuffix(strings.TrimSuffix(modelPath, "/checkpoints/"), "/") + "/checkpoints/"
	modelName := t.PreTrainModel[0].Name
	serviceId := req.Name
	var versionId int64
	var source string
	if req.SourceCloudbrain != nil {
		versionId = req.SourceCloudbrain.VersionID
		source = req.SourceCloudbrain.JobName
	}
	//先发布模型
	publishRes, pulishErr := iflytek.PublishModel(iflytek.PublishModelReq{
		AcceleratorCard: setting.IFLYTekConfig.ComputeSource,
		BaseModel:       modelName,
		ModelPath:       modelPath,
		ModelType:       "SFT/LoRA",
		Name:            "model" + t.Name,
		Source:          source + "-V1",
		SystemStructure: setting.IFLYTekConfig.SystemStructure,
		TaskVersionId:   versionId,
		TrainType:       2,
		TrainingMethod:  "LoRA",
	})
	var modelId string
	if pulishErr != nil {
		if models.IsErrErrModelPublished(pulishErr) {
			modelId = req.SourceCloudbrain.Uuid
		} else {
			log.Error("publish model err.%v ", pulishErr)
			return nil, pulishErr
		}
	} else if publishRes.ID == "" {
		log.Error("publish model empty. ")
		return nil, errors.New("publish model empty")
	} else {
		modelId = publishRes.ID
		req.SourceCloudbrain.Uuid = publishRes.ID
		models.UpdateJob(req.SourceCloudbrain)
	}

	//查询modelPath
	queryModelRes, err := iflytek.QueryModelDetail(modelId)
	if err != nil {
		log.Error("QueryModelDetail err.%v ", err)
		return nil, err
	}
	if len(queryModelRes.FlintItemVoList.Content) == 0 {
		log.Error("FlintItemVoList empty. ")
		return nil, errors.New("FlintItemVoList empty")
	}
	for _, file := range queryModelRes.FlintItemVoList.Content {
		if strings.HasSuffix(file.Name, ".bin") && !file.IsDir {
			modelPath = strings.Replace(file.DataPath, "s3://", "/myjfs/", 1)
			break
		}
	}

	res, err := iflytek.CreateOnlineService(iflytek.CreateOnlineServiceReq{
		ServiceId:       serviceId,
		PretrainedModel: "spark",
		LoraPath:        modelPath,
	})
	if err != nil {
		log.Error("CreateModelExperience CreateOnlineService err.req=%+v,err=%v", req, err)
		return nil, err
	}
	if res.Code != 0 {
		log.Error("CreateModelExperience CreateOnlineService err.req=%+v,err=%+v", req, res)
		return nil, errors.New(fmt.Sprintf("code:%d message:%s", res.Code, res.Message))
	}
	return &entity.CreateNoteBookTaskResponse{
		StartedAt:   0,
		RunSec:      0,
		CompletedAt: 0,
		CreatedAt:   time.Now().Unix(),
		JobID:       serviceId,
		Name:        serviceId,
		Status:      models.GrampusStatusWaiting,
	}, nil
}

func (c IFLYTEKTrainingClusterAdapter) QueryModelExperience(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	res, err := iflytek.QueryOnlineServiceStatus(opts.JobID)
	if err != nil {
		log.Error("QueryModelExperience QueryOnlineServiceStatus err.req=%s,err=%v", opts.JobID, err)
		return nil, err
	}
	task, err := models.GetCloudbrainByJobID(opts.JobID)
	if err != nil {
		log.Error("QueryModelExperience GetCloudbrainByJobID err.req=%s,err=%v", opts.JobID, err)
		return nil, err
	}
	var status string
	var startTime, endTime timeutil.TimeStamp
	if task.Status == models.GrampusStatusRunning && res.Code != 0 {
		endTime = timeutil.TimeStampNow()
		status = models.GrampusStatusStopped
	} else {
		endTime = task.EndTime
	}
	if task.Status == models.GrampusStatusWaiting && res.Code != 0 {
		return &entity.QueryTaskResponse{
			Status:     models.GrampusStatusFailed,
			CenterName: "火石平台",
			JobId:      task.JobID,
		}, nil
	}

	if task.StartTime == 0 && res.Data {
		startTime = timeutil.TimeStampNow()
	} else {
		startTime = task.StartTime
	}

	if res.Code == 0 {
		status = convertIFLYTEKServiceStatus(res.Data)
	}
	return &entity.QueryTaskResponse{
		StartedAt:   startTime,
		CompletedAt: endTime,
		Status:      status,
		JobId:       task.JobID,
		VersionId:   task.VersionID,
		CenterId:    models.AICenterOfHuoShi,
		CenterName:  "火石平台",
	}, nil
}

func (c IFLYTEKTrainingClusterAdapter) StopModelExperience(opts entity.JobIdAndVersionId) error {
	res, err := iflytek.StopOnlineService(opts.JobID)
	if err != nil {
		log.Error("StopModelExperience StopModelExperience err.req=%s,err=%v", opts.JobID, err)
		return err
	}
	if res.Code != 0 {
		log.Error("StopModelExperience StopOnlineService err.req=%s,err=%+v", opts.JobID, res)
		return errors.New(fmt.Sprintf("code:%d message:%s", res.Code, res.Message))
	}
	return nil
}

func (c IFLYTEKTrainingClusterAdapter) DeleteModelExperience(opts entity.JobIdAndVersionId) error {
	return nil

}

func (c IFLYTEKTrainingClusterAdapter) GetSelfEndPointUrl(jobId string) (string, error) {
	return "", nil
}

func (c IFLYTEKTrainingClusterAdapter) CreateGeneralTask(req entity.CreateGeneralTaskRequest, trace *entity.TraceInfo) (*entity.CreateGeneralTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) CreateSdFinetune(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) CreateComfyuiExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetNotebookImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	return nil, true, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetTrainImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	return nil, false, nil
}

func (c IFLYTEKTrainingClusterAdapter) RestartNoteBook(string, int64, *entity.TraceInfo) (*entity.RestartNoteBookTaskResponse, error) {

	return nil, nil
}
func (c IFLYTEKTrainingClusterAdapter) DeleteNoteBook(entity.JobIdAndVersionId) error {
	return nil
}

func (c IFLYTEKTrainingClusterAdapter) StopNoteBook(opts entity.JobIdAndVersionId) error {
	return nil
}

func (c IFLYTEKTrainingClusterAdapter) QueryNoteBook(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) QueryOnlineInfer(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) QueryNoteBookByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetNoteBookLog(jobId string) (*entity.ClusterLog, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetNoteBookUrl(jobId string) (string, error) {
	return "", nil
}

func (c IFLYTEKTrainingClusterAdapter) GetNoteBookOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) CreateTrainJob(req entity.CreateTrainTaskRequest, trace *entity.TraceInfo) (*entity.CreateTrainTaskResponse, error) {
	//1、新增数据集+数据集版本
	task := req.Tasks[0]
	if len(task.PreTrainModel) != 1 {
		return nil, errors.New("params error")
	}
	pretrainModel := task.PreTrainModel[0]
	spec := task.Spec

	var datasets []iflytek.TrainDataSet
	for _, dataset := range task.Datasets {
		data, err := uploadDataset2IFLYTEK(dataset)
		if err != nil {
			log.Error("CreateTrainJob failed. uploadDataset2IFLYTEK error.%v", err)
			return nil, err
		}
		datasets = append(datasets, data...)
	}
	if len(datasets) == 0 {
		log.Error("CreateTrainJob failed.invalid dataset format.req=%+v", req)
		return nil, errors.New("invalid dataset format")
	}

	//构造请求参数
	metaData, err := iflytek.GetFlintTrainTaskMetaData(iflytek.GetMetaDataOpts{
		BaseModel:       pretrainModel.Name,
		TrainingMethod:  task.Method,
		SystemStructure: setting.IFLYTekConfig.SystemStructure,
		AcceleratorCard: setting.IFLYTekConfig.ComputeSource,
	})
	if err != nil {
		log.Error("CreateTrainJob failed. GetFlintTrainTaskMetaData error.%v", err)
		return nil, err
	}
	var extendInfo string
	if metaData != nil {
		paramArray := metaData.Params[setting.IFLYTekConfig.XNum+"-node-train"]
		if len(paramArray) > 0 {
			newArray := make([]iflytek.NodeParam, len(paramArray))
			for i := 0; i < len(paramArray); i++ {
				param := paramArray[i]
				tag := param.OriginParam.Tag
				if tag == "cpuNum" {
					param.Value = spec.CpuCores
				} else if tag == "memorySize" {
					param.Value = spec.MemGiB
				} else if tag == "dshmSize" {
					param.Value = spec.ShareMemGiB
				} else if tag == "numGpu" {
					param.Value = spec.AccCardsNum
				} else if tag == "numWorker" {
					param.Value = task.WorkServerNumber
				} else {
					for _, p := range task.Params.Parameter {
						if p.Label == tag {
							param.Value = p.Value
						}
					}
				}
				newArray[i] = param
			}
			metaData.Params[setting.IFLYTekConfig.XNum+"-node-train"] = newArray
		}
		categories, err := iflytek.GetFlintTrainTaskProfile()
		if err != nil {
			log.Error("CreateTrainJob failed.GetFlintTrainTaskProfile error.%v", err)
			return nil, err
		}
		for i := 0; i < len(categories); i++ {
			queryName := categories[i].QueryName
			if queryName == "baseModel" {
				categories[i].Value = pretrainModel.Name
			} else if queryName == "trainingMethod" {
				categories[i].Value = task.Method
			} else if queryName == "systemStructure" {
				categories[i].Value = setting.IFLYTekConfig.SystemStructure
			} else if queryName == "acceleratorCard" {
				categories[i].Value = setting.IFLYTekConfig.ComputeSource
			}
		}
		metaData.Profile = categories
		tmpByte, err := json.Marshal(metaData)
		if err != nil {
			log.Error("CreateTrainJob failed. json.Marshal error.%v", err)
			return nil, err
		}
		extendInfo = string(tmpByte)
	}

	//

	//提交创建任务请求
	taskRes, err := iflytek.CreateFlintTrainTask(iflytek.FlintTrainTaskReq{
		Description: req.Description,
		Name:        req.Name,
		TrainType:   1,
		InitVersion: iflytek.InitVersion{
			IncrementalTrain: false,
			AssetQueue:       spec.Remark,
			TrainDataSets:    datasets,
			BaseModelCode:    pretrainModel.Name,
			MaxRunTime:       0, //不限时
			TestDataSets:     []iflytek.TrainDataSet{},
			RunNow:           true,
			ExtendInfo:       extendInfo,
		},
	})
	if err != nil {
		log.Error("CreateTrainJob failed. CreateFlintTrainTask error.%v", err)
		return nil, err
	}
	var createData int64
	if taskRes.CreatedDate != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", taskRes.CreatedDate)
		if err != nil {
			return nil, err
		}
		createData = parsedTime.Unix()
	}

	return &entity.CreateTrainTaskResponse{
		StartedAt:   0,
		RunSec:      0,
		CompletedAt: 0,
		CreatedAt:   createData,
		JobID:       fmt.Sprint(taskRes.ID),
		Name:        taskRes.Name,
		Status:      models.GrampusStatusWaiting,
	}, nil
}

func uploadDataset2IFLYTEK(dataset entity.ContainerData) ([]iflytek.TrainDataSet, error) {
	result := make([]iflytek.TrainDataSet, 0)
	//找到数据集中的jsonl
	helper := storage_helper.SelectStorageHelperFromStorageType(dataset.StorageType)
	if !dataset.IsDir {
		if !strings.HasSuffix(dataset.Name, ".jsonl") {
			return result, nil
		}
		res, err := uploadSingleFile2IFLYTEK(dataset)
		if err != nil {
			log.Error("uploadDataset2IFLYTEK uploadSingleFile2IFLYTEK err.dataset=%+v err=%v", dataset, err)
			return nil, err
		}
		return []iflytek.TrainDataSet{*res}, nil
	}
	fileList, err := helper.GetOneLevelObjectsUnderDir(dataset.ObjectKey)
	if err != nil {
		log.Error("uploadDataset2IFLYTEK GetOneLevelObjectsUnderDir err.dataset=%+v err=%v", dataset, err)
		return nil, err
	}
	for _, file := range fileList {
		if !strings.HasSuffix(file.FileName, ".jsonl") {
			continue
		}
		res, err := uploadSingleFile2IFLYTEK(entity.ContainerData{
			Name:        file.FileName,
			ObjectKey:   file.RelativePath,
			StorageType: dataset.StorageType,
		})
		if err != nil {
			log.Error("uploadDataset2IFLYTEK uploadSingleFile2IFLYTEK in loop err.fileName=%s sourceDataset=%+v err=%v", file.FileName, dataset, err)
			return nil, err
		}
		result = append(result, *res)
	}
	return result, nil
}

func uploadSingleFile2IFLYTEK(dataset entity.ContainerData) (*iflytek.TrainDataSet, error) {
	//新增数据集
	res, err := iflytek.AddFlintDatasetGroups(iflytek.FlintDatasetGroup{
		LabelId: setting.IFLYTekConfig.DefaultLabelId,
		Name:    dataset.Name,
	})
	if err != nil {
		log.Error("uploadDataset2IFLYTEK failed. AddFlintDatasetGroups error.%v", err)
		return nil, err
	}
	//上传文件
	helper := storage_helper.SelectStorageHelperFromStorageType(dataset.StorageType)
	fileStream, err := helper.OpenFile(dataset.ObjectKey)
	if err != nil {
		log.Error("uploadDataset2IFLYTEK failed.open dataset file error.%v", err)
		return nil, err
	}
	fullPath, err := iflytek.UploadFile(fileStream, dataset.Name)
	if err != nil {
		log.Error("uploadDataset2IFLYTEK failed.UploadFile error.%v", err)
		return nil, err
	}
	versionRes, err := iflytek.AddFlintDatasetVersion(iflytek.FlintDatasetVersionReq{
		Source:         0,
		DataType:       1,
		OriDataPath:    fullPath,
		MarkStatus:     1, //只允许已标注的数据集
		PremarkStatus:  0,
		MarkToolId:     16,
		DatasetGroupId: res.ID,
	})
	if err != nil {
		log.Error("uploadDataset2IFLYTEK failed.open dataset file error.%v", err)
		return nil, err
	}
	return versionRes, nil
}

func (c IFLYTEKTrainingClusterAdapter) DeleteTrainJob(opts entity.JobIdAndVersionId) error {
	taskId, err := strconv.ParseInt(opts.JobID, 10, 64)
	if err != nil {
		log.Error("DeleteTrainJob failed. taskId format error.opts.JobID=%s  err=%v", opts.JobID, err)
		return errors.New("params error")
	}

	_, err = iflytek.DeleteFlintTrainTask(taskId)
	if err != nil {
		log.Error("DeleteTrainJob(%s) failed:%v", opts, err)
		return err
	}
	return nil
}
func (c IFLYTEKTrainingClusterAdapter) StopTrainJob(opts entity.JobIdAndVersionId) error {
	versionId := opts.VersionID
	if versionId == 0 {
		log.Error("StopTrainJob failed. taskId format error.opts.versionId=%s ", opts.VersionID)
		return errors.New("params error")
	}

	_, err := iflytek.StopFlintTrainTask(versionId)
	if err != nil {
		log.Error("StopTrainJob(%s) failed:%v", opts, err)
		return err
	}
	return nil
}
func (c IFLYTEKTrainingClusterAdapter) QueryTrainJob(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	taskId, err := strconv.ParseInt(opts.JobID, 10, 64)
	if err != nil {
		log.Error("QueryTrainJob failed. taskId format error.opts.JobID=%s  err=%v", opts.JobID, err)
		return nil, errors.New("params error")
	}
	res, err := iflytek.QueryFlintTrainTask(taskId)
	if err != nil {
		log.Error("QueryTrainJob failed.QueryFlintTrainTask error.opts. err=%v", err)
		return nil, err
	}

	var startTime, endTime timeutil.TimeStamp
	if res.TrainStartTime != nil {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprint(res.TrainStartTime), time.Local)
		if err != nil {
			log.Error("QueryTrainJob failed.  res.TrainStartTime(%s) format error. err=%v", res.TrainStartTime, err)
			return nil, err
		}
		startTime = timeutil.TimeStamp(t.Unix())
	}
	if res.TrainEndTime != nil {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprint(res.TrainEndTime), time.Local)
		if err != nil {
			log.Error("QueryTrainJob failed.  res.TrainEndTime(%s) format error. err=%v", res.TrainStartTime, err)
			return nil, err
		}
		endTime = timeutil.TimeStamp(t.Unix())
	}
	return &entity.QueryTaskResponse{
		StartedAt:   startTime,
		CompletedAt: endTime,
		Status:      convertIFLYTEKStatus(res.TrainStatus),
		JobId:       fmt.Sprint(res.TaskId),
		VersionId:   res.ID,
		CenterId:    models.AICenterOfHuoShi,
		CenterName:  "火石平台",
		DataUrl:     res.ModelPath,
	}, nil
}

func convertIFLYTEKStatus(iflytekStatus int) string {
	//（1未运行、2待运行、3运行中、4运行成功、5运行失败、6已停止）
	switch iflytekStatus {
	case 1, 2:
		return models.GrampusStatusWaiting
	case 3:
		return models.GrampusStatusRunning
	case 4:
		return models.GrampusStatusSucceeded
	case 5:
		return models.GrampusStatusFailed
	case 6:
		return models.GrampusStatusStopped
	}
	return models.GrampusStatusUnknown
}

func convertIFLYTEKServiceStatus(isRunning bool) string {
	//true 发布成功 false 正在发布
	var status = models.GrampusStatusWaiting
	if isRunning {
		status = models.GrampusStatusRunning
	}
	return status
}

func (c IFLYTEKTrainingClusterAdapter) QueryTrainJobByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetVisualizeUrl(jobId string) (string, error) {
	return "", nil
}

func (c IFLYTEKTrainingClusterAdapter) GetLog(opts entity.ClusterLogOpts) (*entity.ClusterLog, error) {
	versionId := opts.VersionID
	if versionId == 0 {
		log.Error("GetLog failed. taskId format error.opts.versionId=%s ", opts.VersionID)
		return nil, errors.New("params error")
	}
	res, err := iflytek.GetFlintTrainTaskLog(versionId)
	if err != nil {
		log.Error("GetLog failed. GetFlintTrainTaskLog error.opts.JobID=%s  err=%v", opts.JobId, err)
		return nil, err
	}

	var log string

	for i := 0; i < len(res.Data.Nodes); i++ {
		node := res.Data.Nodes[i]
		for j := 0; j < len(node.Log); j++ {
			l := node.Log[j]
			if l.Log != "" {
				log = log + "\n" + l.Log
			}
		}
	}

	return &entity.ClusterLog{
		Content: log,
	}, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetLogDownloadInfo(opts entity.ClusterLogDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
	//获取任务退出信息
	existStr := getCloudbrainOneExitDiagnostics(opts.JobId)

	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)

	//查找日志文件
	files := getLogFilesInStorage(helper, opts.ObjectKeyPrefix, "log.txt")
	if len(files) == 0 {
		//此时未找符合条件的文件
		if existStr != "" {
			return &entity.FileDownloadInfo{
				ResultType:     entity.FileTypeTXT,
				ResultFileName: "exit.log.txt",
				Readers: []entity.FileReader{{
					Reader: ioutil.NopCloser(strings.NewReader(existStr)),
				}},
			}, nil
		}

		return nil, nil
	}

	//默认选择第一个文件
	file := files[0]

	//获取日志reader
	reader, err := helper.OpenFile(file.RelativePath)
	if err != nil {
		log.Error("GetLogDownloadInfo OpenFile err.opts=%+v,err =%v", opts, err)
		return nil, err
	}

	return &entity.FileDownloadInfo{
		ResultType:     entity.FileTypeTXT,
		ResultFileName: file.FileName,
		Readers: []entity.FileReader{{
			Reader: reader,
		}},
	}, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetSingleOutputDownloadInfo(opts entity.ClusterSingleOutputDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
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

func (c IFLYTEKTrainingClusterAdapter) DownloadAllOutput(opts entity.DownloadOutputOpts) error {
	return DownloadAllOutput(opts)
}

func (c IFLYTEKTrainingClusterAdapter) GetNodeInfo(opts entity.ClusterNodeInfoOpts) ([]entity.AITaskNodeInfo, error) {
	return nil, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetResourceUsage(opts entity.ClusterResourceUsageOpts) (*entity.ResourceUsage, error) {
	return &entity.ResourceUsage{
		Interval:    0,
		MetricsInfo: []entity.MetricsInfo{},
	}, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetTrainJobOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	return c.GetNoteBookOperationProfile(opts)
}

func (c IFLYTEKTrainingClusterAdapter) GetOutput(opts entity.ClusterOutputOpts) (*entity.ClusterAITaskOutput, error) {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	fileList, err := helper.GetOneLevelObjectsUnderDir(path.Join(opts.ObjectKeyPrefix, opts.ParentDir))
	if err != nil {
		log.Error("GetOneLevelObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.ObjectKeyPrefix, err)
		return nil, err
	}
	return &entity.ClusterAITaskOutput{
		Status:   models.ModelMigrateSuccess,
		Path:     opts.ParentDir,
		FileList: fileList,
	}, nil
}

func (c IFLYTEKTrainingClusterAdapter) GetAllOutput(opts entity.ClusterOutputOpts) (*entity.AllAITaskOutput, error) {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	fileList, err := helper.GetAllObjectsUnderDir(path.Join(opts.ObjectKeyPrefix, opts.ParentDir))
	if err != nil {
		log.Error("GetAllObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.ObjectKeyPrefix, err)
		return nil, err
	}
	return &entity.AllAITaskOutput{FileList: fileList}, nil
}
