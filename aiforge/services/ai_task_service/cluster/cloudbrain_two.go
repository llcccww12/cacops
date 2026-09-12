package cluster

import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path"
	"strconv"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/manager/client/cloudbrain_two"
	"code.gitea.io/gitea/manager/client/cloudbrain_two_cd"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type CloudbrainTwoClusterAdapter struct {
}

func init() {
	AddCluster(entity.OpenICloudbrainTwo, new(CloudbrainTwoClusterAdapter))
}

func (c CloudbrainTwoClusterAdapter) CreateNoteBook(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	t := req.Tasks[0]

	appUrl := JointCloudbrainTwoReqUrl(t.Code)
	if appUrl != "" {
		appUrl = "s3:/" + appUrl
	}
	trainUrl := JointCloudbrainTwoReqUrl(t.OutPut)
	if trainUrl != "" {
		trainUrl = "s3:/" + trainUrl
	}
	datasetUrl := getCloudbrainTwoMultiDataUrl(t.Datasets)
	multiModelUrl := getCloudbrainTwoModelUrl(t.PreTrainModel)

	var jobResult *models.CreateNotebookResult
	var err error
	if setting.ModelartsCD.Enabled {
		jobResult, err = cloudbrain_two_cd.CreateNotebook(models.CreateNotebookWithoutPoolParams{
			JobName:     req.Name,
			Description: req.Description,
			Flavor:      t.Spec.SourceSpecId,
			Duration:    t.AutoStopDuration,
			ImageID:     t.ImageId,
			Feature:     models.NotebookFeature,
			Volume: models.VolumeReq{
				Capacity:  setting.Capacity,
				Category:  models.EVSCategory,
				Ownership: models.ManagedOwnership,
			},
			WorkspaceID: "0",
		})
	} else {
		jobResult, err = cloudbrain_two.CreateNotebook2(models.CreateNotebook2Params{
			JobName:     req.Name,
			Description: req.Description,
			Flavor:      t.Spec.SourceSpecId,
			Duration:    t.AutoStopDuration,
			ImageID:     t.ImageId,
			PoolID:      t.Spec.QueueCode,
			Feature:     models.NotebookFeature,
			Volume: models.VolumeReq{
				Capacity:  setting.Capacity,
				Category:  models.EVSCategory,
				Ownership: models.ManagedOwnership,
			},
			EnvVariables: models.CloudBrain2EnvVarReq{
				CodeObsUrl:             appUrl,
				DatasetObsUrl:          datasetUrl,
				PretrainedModelObsUrl:  multiModelUrl,
				OutputObsUrl:           trainUrl,
				LocalCodePath:          models.LocalCodePath,
				LocalDatasetPath:       models.LocalDatasetPath,
				LocalPretrainModelPath: models.LocalPretrainModelPath,
				LocalOutputPath:        models.LocalOutputPath,
				DataDownloadMethod:     models.DataDownloadMethodMoxing,
				CodeNeedUnzip:          models.CodeNeedUnzipTrue,
				DatasetNeedUnzip:       models.DatasetNeedUnzipTrue,
				PretrainModelNeedUnzip: models.PretrainModelNeedUnzipFalse,
			},
			WorkspaceID: "0",
		})
	}

	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	return convertCloudbrainTwo2NoteBookRes(jobResult), nil
}

func (c CloudbrainTwoClusterAdapter) GetSelfEndPointUrl(jobId string) (string, error) {
	return "", nil
}
func (c CloudbrainTwoClusterAdapter) CreateOnlineInfer(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}
func (c CloudbrainTwoClusterAdapter) CreateModelExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}
func (c CloudbrainTwoClusterAdapter) CreateComfyuiExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainTwoClusterAdapter) CreateSdFinetune(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainTwoClusterAdapter) QueryModelExperience(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return c.QueryNoteBook(opts)
}

func (c CloudbrainTwoClusterAdapter) StopModelExperience(opts entity.JobIdAndVersionId) error {
	return nil
}
func (c CloudbrainTwoClusterAdapter) DeleteModelExperience(opts entity.JobIdAndVersionId) error {
	return nil

}
func (c CloudbrainTwoClusterAdapter) CreateGeneralTask(req entity.CreateGeneralTaskRequest, trace *entity.TraceInfo) (*entity.CreateGeneralTaskResponse, error) {
	return nil, nil
}

var cloudbrainTwoNotebookImages []entity.ClusterImage

func (c CloudbrainTwoClusterAdapter) GetNotebookImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	if cloudbrainTwoNotebookImages == nil || len(cloudbrainTwoNotebookImages) == 0 {
		images := setting.StImageInfos.ImageInfo
		cloudbrainTwoNotebookImages = make([]entity.ClusterImage, len(images))
		for i := 0; i < len(images); i++ {
			cloudbrainTwoNotebookImages[i] = entity.ClusterImage{
				ImageId:   images[i].Id,
				ImageName: images[i].Value,
			}
		}
	}

	return cloudbrainTwoNotebookImages, false, nil
}

var cloudbrainTwoTrainImages []entity.ClusterImage

func (c CloudbrainTwoClusterAdapter) GetTrainImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	if cloudbrainTwoTrainImages == nil || len(cloudbrainTwoTrainImages) == 0 {
		var versionInfos modelarts.VersionInfo
		if err := json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
			log.Error("Unmarshal setting.EngineVersions err. req=%+v err=%v", req, err)
			return cloudbrainTwoTrainImages, false, nil
		}
		cloudbrainTwoTrainImages = make([]entity.ClusterImage, len(versionInfos.Version))
		for i := 0; i < len(versionInfos.Version); i++ {
			cloudbrainTwoTrainImages[i] = entity.ClusterImage{
				ImageId:   fmt.Sprint(versionInfos.Version[i].ID),
				ImageName: versionInfos.Version[i].Value,
				ImageUrl:  versionInfos.Version[i].Url,
			}
		}
	}

	return cloudbrainTwoTrainImages, false, nil
}

func (c CloudbrainTwoClusterAdapter) GetTrainImageByImageId(imageId string) (entity.ClusterImage, error) {
	if imageId == "" {
		return entity.ClusterImage{}, errors.New("imageId is empty")
	}
	images, _, _ := c.GetTrainImages(entity.GetImageReq{})
	if images == nil {
		return entity.ClusterImage{}, errors.New("image not setting correctly")
	}
	for _, image := range images {
		if image.ImageId == imageId {
			return image, nil
		}
	}
	return entity.ClusterImage{}, errors.New("image not exists")
}

var poolInfos *models.PoolInfos

func convertCloudbrainTwo2NoteBookRes(res *models.CreateNotebookResult) *entity.CreateNoteBookTaskResponse {
	return &entity.CreateNoteBookTaskResponse{
		JobID:  res.ID,
		Status: res.Status,
	}
}

func (c CloudbrainTwoClusterAdapter) RestartNoteBook(jobId string, autoStopDuration int64, traceInfo *entity.TraceInfo) (*entity.RestartNoteBookTaskResponse, error) {
	param := models.NotebookAction{
		Action: models.ActionStart,
	}
	task, err := models.GetNewestCloudbrainByJobId(jobId)
	if err != nil {
		return nil, err
	}

	var res *models.NotebookActionResult
	if task.Type == models.TypeCloudBrainTwo {
		res, err = cloudbrain_two.ManageNotebook2(task.JobID, param, int(autoStopDuration))
	} else if task.Type == models.TypeCDCenter {
		res, err = cloudbrain_two_cd.ManageNotebook(task.JobID, param, int(autoStopDuration))
	}
	if err != nil {
		log.Error("ManageNotebook err.jobID=%s err=%v", jobId, err)
		return nil, err
	}
	return convertCloudbrainTwo2NoteBookRestartRes(jobId, res), nil
}

func convertCloudbrainTwo2NoteBookRestartRes(jobId string, res *models.NotebookActionResult) *entity.RestartNoteBookTaskResponse {
	return &entity.RestartNoteBookTaskResponse{
		JobId:  jobId,
		Status: res.Status,
	}
}

func (c CloudbrainTwoClusterAdapter) DeleteNoteBook(opts entity.JobIdAndVersionId) error {
	task, err := models.GetNewestCloudbrainByJobId(opts.JobID)
	if err != nil {
		return err
	}

	if task.Type == models.TypeCloudBrainTwo {
		_, err = cloudbrain_two.DelNotebook2(task.JobID)
	} else if task.Type == models.TypeCDCenter {
		_, err = cloudbrain_two_cd.DelNotebook(task.JobID)
	}
	if err != nil {
		log.Error("DeleteNoteBook err.jobID=%s err=%v", opts, err)
		log.Info("error=" + err.Error())
		return nil
	}
	return nil
}

func (c CloudbrainTwoClusterAdapter) StopNoteBook(opts entity.JobIdAndVersionId) error {
	task, err := models.GetNewestCloudbrainByJobId(opts.JobID)
	if err != nil {
		return err
	}
	param := models.NotebookAction{
		Action: models.ActionStop,
	}
	if task.Type == models.TypeCloudBrainTwo {
		_, err = cloudbrain_two.ManageNotebook2(task.JobID, param, 0)
	} else if task.Type == models.TypeCDCenter {
		_, err = cloudbrain_two_cd.ManageNotebook(task.JobID, param, 0)
	}
	if err != nil {
		log.Error("StopNoteBook err.jobID=%s err=%v", opts, err)
		return err
	}
	return nil
}

func (c CloudbrainTwoClusterAdapter) QueryNoteBook(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	task, err := models.GetNewestCloudbrainByJobId(opts.JobID)
	if err != nil {
		return nil, err
	}

	var result *models.GetNotebook2Result
	if task.Type == models.TypeCloudBrainTwo {
		result, err = cloudbrain_two.GetNotebook2(task.JobID)
	} else if task.Type == models.TypeCDCenter {
		result, err = cloudbrain_two_cd.GetNotebook(task.JobID)
	}
	if err != nil {
		log.Error("GetNotebook(%s) failed:%v", task.DisplayJobName, err)
		return nil, err
	}
	if result == nil {
		log.Error("GetNotebook(%s) from cloudbrain 2 failed:result is empty", task.DisplayJobName)
		return nil, errors.New("result is empty")
	}
	return convertCloudbrainTwo2QueryRes(result, task), nil
}

func (c CloudbrainTwoClusterAdapter) QueryOnlineInfer(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return nil, nil
}

func convertCloudbrainTwo2QueryRes(res *models.GetNotebook2Result, task *models.Cloudbrain) *entity.QueryTaskResponse {
	startedAt := timeutil.TimeStamp(0)

	if task.IsRestartTask() {
		if res.Lease.UpdateTime > 0 {
			startedAt = timeutil.TimeStamp(res.Lease.UpdateTime / 1000)
		}
	} else {
		if res.Lease.CreateTime > 0 {
			startedAt = timeutil.TimeStamp(res.Lease.CreateTime / 1000)
		}
	}
	completedAt := timeutil.TimeStamp(0)
	if models.IsCloudbrainTerminalStatus(res.Status) {
		completedAt = timeutil.TimeStampNow()
	}
	return &entity.QueryTaskResponse{
		StartedAt:   startedAt,
		CompletedAt: completedAt,
		JobId:       res.ID,
		Status:      res.Status,
		Url:         res.Url,
		Token:       res.Token,
	}
}

func (c CloudbrainTwoClusterAdapter) QueryNoteBookByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	result, err := cloudbrain_two.GetNotebookList(1000, 0, "createTime", "DESC", jobName)
	if err != nil {
		log.Error("QueryNoteBookByJobName failed:jobName=%s err=%v", jobName, err)
		return nil, err
	}
	r := make([]*entity.QueryTaskResponse, 0)
	for i := 0; i < len(result.NotebookList); i++ {
		if result.NotebookList[i].JobName == jobName {
			r = append(r, convertCloudbrainTwoQueryNotebookByNameResponse(result.NotebookList[i]))
		}
	}
	return r, nil
}

func convertCloudbrainTwoQueryNotebookByNameResponse(notebook models.NotebookList) *entity.QueryTaskResponse {
	return &entity.QueryTaskResponse{
		StartedAt: timeutil.TimeStamp(notebook.Lease.CreateTime / 1000),
		Status:    notebook.Status,
		JobId:     notebook.JobID,
	}
}

func (c CloudbrainTwoClusterAdapter) GetNoteBookLog(jobId string) (*entity.ClusterLog, error) {
	return nil, nil
}

func (c CloudbrainTwoClusterAdapter) GetNoteBookUrl(jobId string) (string, error) {
	res, err := c.QueryNoteBook(entity.JobIdAndVersionId{
		JobID: jobId,
	})
	if err != nil {
		return "", err
	}
	return res.Url + "?token=" + res.Token, nil
}

func (c CloudbrainTwoClusterAdapter) GetNoteBookOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	task, err := models.GetNewestCloudbrainByJobId(opts.JobID)
	if err != nil {
		return nil, err
	}

	var result *models.GetNotebook2Result
	if task.Type == models.TypeCloudBrainTwo {
		result, err = cloudbrain_two.GetNotebook2(task.JobID)
	} else if task.Type == models.TypeCDCenter {
		result, err = cloudbrain_two_cd.GetNotebook(task.JobID)
	}
	if err != nil {
		log.Error("GetNotebook(%s) failed:%v", task.DisplayJobName, err)
		return nil, err
	}
	return parseCloudbrainTwoEventsToOperationProfile(result), nil
}

func parseCloudbrainTwoEventsToOperationProfile(result *models.GetNotebook2Result) *entity.OperationProfile {
	events := make([]entity.ProfileEvent, 0)
	if result.ActionProgress == nil || len(result.ActionProgress) == 0 {
		return nil
	}
	for i := 0; i < len(result.ActionProgress); i++ {
		e := result.ActionProgress[i]
		if e.Description == "" {
			continue
		}
		events = append(events, entity.ProfileEvent{
			Message: e.Description,
			Reason:  fmt.Sprint(e.Step),
			Name:    e.Status,
		})
	}
	return &entity.OperationProfile{Events: events}
}

func (c CloudbrainTwoClusterAdapter) CreateTrainJob(req entity.CreateTrainTaskRequest, trace *entity.TraceInfo) (*entity.CreateTrainTaskResponse, error) {
	t := req.Tasks[0]
	var jobResult *models.CreateTrainJobResult
	var createErr error
	//imageId 为0或者负数时代表自定义镜像
	imageId := int64(-1)
	i, err := strconv.ParseInt(t.ImageId, 10, 32)
	if err == nil {
		imageId = i
	}
	if imageId <= 0 {
		image, err := c.GetTrainImageByImageId(t.ImageId)
		if err != nil {
			log.Error("GetTrainImageByImageId error.req=%+v err=%v", req, err)
			return nil, err
		}
		jobResult, createErr = modelarts.CreateTrainJobUserImage(convertCloudbrainTwoTrainJobUserImageReq(req, image))
	} else {
		param, err := convertCloudbrainTwoTrainJobReq(req)
		if err != nil {
			return nil, err
		}
		jobResult, createErr = modelarts.CreateTrainJob(param)
	}
	if createErr != nil {
		log.Error("CloudbrainTwo createTrainJob failed: %v", createErr.Error())
		if strings.HasPrefix(createErr.Error(), modelarts.UnknownErrorPrefix) {
			return nil, models.NetworkError{}
		}
		return nil, createErr
	}
	return convertCloudbrainTwoRes2Standard(jobResult), nil
}

func (c CloudbrainTwoClusterAdapter) GetVisualizeUrl(jobId string) (string, error) {
	return "", nil
}

func convertCloudbrainTwoRes2Standard(res *models.CreateTrainJobResult) *entity.CreateTrainTaskResponse {
	var jobId string
	if res.JobID > 0 {
		jobId = fmt.Sprint(res.JobID)
	}
	return &entity.CreateTrainTaskResponse{
		CreatedAt:   res.CreateTime,
		JobID:       jobId,
		Name:        res.JobName,
		Status:      modelarts.TransTrainJobStatus(res.Status),
		VersionID:   res.VersionID,
		VersionName: res.VersionName,
	}
}

func convertCloudbrainTwoTrainJobReq(req entity.CreateTrainTaskRequest) (models.CreateTrainJobParams, error) {
	t := req.Tasks[0]
	imageId, err := strconv.ParseInt(t.ImageId, 10, 64)
	if err != nil {
		log.Error("Parse imageId err.imageIdStr=%s err=%v", t.ImageId, err)
		return models.CreateTrainJobParams{}, err
	}
	return models.CreateTrainJobParams{
		JobName:     req.Name,
		Description: req.Description,
		Config: models.Config{
			WorkServerNum: t.WorkServerNumber,
			AppUrl:        JointCloudbrainTwoReqUrl(t.Code),
			BootFileUrl:   path.Join(JointCloudbrainTwoReqUrl(t.Code), t.BootFile),
			DataUrl:       JointCloudbrainTwoReqUrl(t.Datasets),
			TrainUrl:      JointCloudbrainTwoReqUrl(t.OutPut),
			LogUrl:        JointCloudbrainTwoReqUrl(t.LogPath),
			PoolID:        t.PoolId,
			CreateVersion: true,
			Flavor: models.Flavor{
				Code: t.Spec.SourceSpecId,
			},
			EngineID:  imageId,
			Parameter: handleCloudbrainTwoParameter(req).Parameter,
			ShareAddr: setting.ModelArtsShareAddr,
			MountPath: setting.ModelArtsMountPath,
			NasType:   setting.ModelArtsNasType,
		},
	}, nil

}

func JointCloudbrainTwoReqUrl(data []entity.ContainerData) string {
	if len(data) > 0 {
		d := data[0]
		s := path.Join("/", d.Bucket, d.ObjectKey)
		if d.IsDir {
			s = strings.TrimSuffix(path.Join("/", d.Bucket, d.ObjectKey, "/"), "/") + "/"
		}
		return s
	}
	return ""
}

func convertCloudbrainTwoTrainJobUserImageReq(req entity.CreateTrainTaskRequest, image entity.ClusterImage) models.CreateUserImageTrainJobParams {
	t := req.Tasks[0]
	appUrl := JointCloudbrainTwoReqUrl(t.Code)
	bootFileUrl := path.Join(JointCloudbrainTwoReqUrl(t.Code), t.BootFile)
	dataUrl := JointCloudbrainTwoReqUrl(t.Datasets)
	trainUrl := JointCloudbrainTwoReqUrl(t.OutPut)
	logUrl := JointCloudbrainTwoReqUrl(t.LogPath)
	params := handleCloudbrainTwoParameter(req)

	return models.CreateUserImageTrainJobParams{
		JobName:     req.Name,
		Description: req.Description,
		Config: models.UserImageConfig{
			WorkServerNum: t.WorkServerNumber,
			AppUrl:        appUrl,
			BootFileUrl:   bootFileUrl,
			DataUrl:       dataUrl,
			TrainUrl:      trainUrl,
			LogUrl:        logUrl,
			PoolID:        t.PoolId,
			CreateVersion: true,
			Flavor: models.Flavor{
				Code: t.Spec.SourceSpecId,
			},
			UserImageUrl: image.ImageUrl,
			UserCommand:  getCloudbrainTwoUserCommand(appUrl, t.BootFile, dataUrl, trainUrl, params),
			ShareAddr:    setting.ModelArtsShareAddr,
			MountPath:    setting.ModelArtsMountPath,
			NasType:      setting.ModelArtsNasType,
		},
	}
}

func getCloudbrainTwoDataUrl(data []entity.ContainerData) string {
	if len(data) == 0 {
		return ""
	}
	return data[0].ObjectKey
}

func handleCloudbrainTwoParameter(req entity.CreateTrainTaskRequest) models.Parameters {
	t := req.Tasks[0]

	var param = models.Parameters{}

	datasetUrl := getCloudbrainTwoMultiDataUrl(t.Datasets)
	if datasetUrl != "" {
		param.Parameter = append(param.Parameter, models.Parameter{
			Label: modelarts.MultiDataUrl,
			Value: datasetUrl,
		})
	}
	multiModelUrl := getCloudbrainTwoModelUrl(t.PreTrainModel)
	if multiModelUrl != "" {
		param.Parameter = append(param.Parameter, models.Parameter{
			Label: modelarts.PretrainUrl,
			Value: multiModelUrl,
		}, models.Parameter{
			Label: modelarts.CkptUrl,
			Value: t.PreTrainModel[0].S3DownloadUrl,
		})
	}

	param.Parameter = append(param.Parameter, models.Parameter{
		Label: modelarts.ResultUrl,
		Value: t.OutPut[0].S3DownloadUrl,
	})

	existDeviceTarget := false
	for _, parameter := range t.Params.Parameter {
		if parameter.Label == modelarts.DeviceTarget {
			existDeviceTarget = true
		}
		if parameter.Label != modelarts.TrainUrl && parameter.Label != modelarts.DataUrl {
			param.Parameter = append(param.Parameter, models.Parameter{
				Label: parameter.Label,
				Value: parameter.Value,
			})
		}
	}
	if !existDeviceTarget {
		param.Parameter = append(param.Parameter, models.Parameter{
			Label: modelarts.DeviceTarget,
			Value: modelarts.Ascend,
		})
	}
	return param
}

func getCloudbrainTwoUserCommand(appUrl, bootFile, dataUrl, trainUrl string, params models.Parameters) string {
	userCommand := ""

	tmpCodeObsPaths := strings.Split(strings.Trim(appUrl, "/"), "/")
	lastCodeDir := "code"
	if len(tmpCodeObsPaths) > 0 {
		lastCodeDir = tmpCodeObsPaths[len(tmpCodeObsPaths)-1]
	}
	var multi_data_url string
	var pretrain_url string
	for _, param := range params.Parameter {
		if param.Label == "multi_data_url" {
			multi_data_url = string(param.Value)
		}
		if param.Label == "pretrain_url" {
			pretrain_url = string(param.Value)
		}
	}
	//配置环境变量，适配c2net的sdk用于训练脚本获取代码，数据集，模型等
	var envCodeCommand = "export CODE_URL=" + "s3://" + appUrl + ";" + "export LOCAL_CODE_PATH=/cache/code;"
	var envDataCommand = "export DATASET_URL=" + "'" + multi_data_url + "'" + ";" + "export LOCAL_DATASET_PATH=/cache/dataset;"
	var envPretrainCommand = "export PRETRAIN_MODEL_URL=" + "'" + pretrain_url + "'" + ";" + "export LOCAL_PRETRAIN_MODEL_PATH=/cache/pretrainmodel;"
	var envOutputCommand = "export OUTPUT_URL=" + "s3://" + trainUrl + ";" + "export LOCAL_OUTPUT_PATH=/cache/output;"
	var envMoxingCommand = "export DATA_DOWNLOAD_METHOD=MOXING;"
	var envNeedUnzipCommand = "export CODE_NEED_UNZIP=false;export DATASET_NEED_UNZIP=false;export PRETRAIN_MODEL_NEED_UNZIP=false;"
	var envCommand = envCodeCommand + envDataCommand + envPretrainCommand + envOutputCommand + envMoxingCommand + envNeedUnzipCommand
	userCommand = envCommand + "/bin/bash /home/work/run_train.sh 's3://" + appUrl + "' '" + lastCodeDir + "/" + bootFile + "' '/tmp/log/train.log' --'data_url'='s3://" + dataUrl + "' --'train_url'='s3://" + trainUrl + "'"
	for _, param := range params.Parameter {
		userCommand += " --'" + param.Label + "'='" + param.Value + "'"
	}
	return userCommand
}

func getCloudbrainTwoMultiDataUrl(datasets []entity.ContainerData) string {
	if len(datasets) == 0 {
		return ""
	}
	var datasUrlList []models.Datasurl
	for _, d := range datasets {
		datasUrlList = append(datasUrlList, models.Datasurl{
			DatasetUrl:  d.S3DownloadUrl,
			DatasetName: d.Name,
		})
	}
	jsondata, _ := json.Marshal(datasUrlList)
	return string(jsondata)
}

func getCloudbrainTwoModelUrl(pretrainModels []entity.ContainerData) string {
	if len(pretrainModels) == 0 {
		return ""
	}
	var modelUrlList []models.ModelUrls
	for _, d := range pretrainModels {
		modelUrlList = append(modelUrlList, models.ModelUrls{
			ModelUrl:  d.S3DownloadUrl,
			ModelName: d.Name,
		})
	}
	jsondata, _ := json.Marshal(modelUrlList)
	return string(jsondata)
}

func (c CloudbrainTwoClusterAdapter) DeleteTrainJob(opts entity.JobIdAndVersionId) error {
	_, err := modelarts.DelTrainJobVersion(opts.JobID, strconv.FormatInt(opts.VersionID, 10))
	if err != nil {
		log.Error("DeleteTrainJob err.jobID=%s err=%v", opts, err)
		log.Info("error=" + err.Error())
		return nil
	}
	return nil
}

func (c CloudbrainTwoClusterAdapter) StopTrainJob(opts entity.JobIdAndVersionId) error {
	_, err := modelarts.StopTrainJob(opts.JobID, strconv.FormatInt(opts.VersionID, 10))
	if err != nil {
		log.Error("StopTrainJob(%s) failed:%v", opts, err)
		return err
	}
	return nil
}

func (c CloudbrainTwoClusterAdapter) QueryTrainJobByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	res, err := modelarts.GetTrainJobList(20, 1, "create_time", "desc", jobName)
	if err != nil {
		log.Error("GetTrainJobList failed:%v", err)
		return nil, err
	}
	result := make([]*entity.QueryTaskResponse, 0)
	if res != nil {
		for i := 0; i < len(res.JobList); i++ {
			if res.JobList[i].JobName == jobName {
				result = append(result, convertJobList2QueryRes(res.JobList[i]))
			}

		}
	}
	return result, nil
}

func convertJobList2QueryRes(res models.JobList) *entity.QueryTaskResponse {
	return &entity.QueryTaskResponse{
		JobId:     strconv.FormatInt(res.JobID, 10),
		Status:    transCloudbrainTwoTrainJobStatus(res.IntStatus),
		VersionId: res.VersionID,
	}
}

func (c CloudbrainTwoClusterAdapter) QueryTrainJob(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	result, err := modelarts.GetTrainJob(opts.JobID, strconv.FormatInt(opts.VersionID, 10))
	if err != nil {
		log.Error("GetTrainJob(%s) failed:%v", opts, err)
		return nil, err
	}
	return convertCloudbrainTwoTrainJob2QueryRes(result), nil
}

func convertCloudbrainTwoTrainJob2QueryRes(res *models.GetTrainJobResult) *entity.QueryTaskResponse {
	status := transCloudbrainTwoTrainJobStatus(res.IntStatus)
	startedAt := timeutil.TimeStamp(0)
	if res.StartTime > 0 {
		startedAt = timeutil.TimeStamp(res.StartTime / 1000)
	}

	duration := res.Duration / 1000
	completedAt := timeutil.TimeStamp(0)
	if startedAt > 0 && models.IsCloudbrainTerminalStatus(status) {
		completedAt = startedAt.Add(duration)
	}
	return &entity.QueryTaskResponse{
		StartedAt:   startedAt,
		CompletedAt: completedAt,
		JobId:       fmt.Sprint(res.JobID),
		Status:      status,
		VersionId:   res.VersionID,
	}
}

func transCloudbrainTwoTrainJobStatus(status int) string {
	switch status {
	case 0:
		return "UNKNOWN"
	case 1:
		return "INIT"
	case 2:
		return "IMAGE_CREATING"
	case 3:
		return "IMAGE_FAILED"
	case 4:
		return "SUBMIT_TRYING"
	case 5:
		return "SUBMIT_FAILED"
	case 6:
		return "DELETE_FAILED"
	case 7:
		return "WAITING"
	case 8:
		return "RUNNING"
	case 9:
		return "KILLING"
	case 10:
		return "COMPLETED"
	case 11:
		return "FAILED"
	case 12:
		return "KILLED"
	case 13:
		return "CANCELED"
	case 14:
		return "LOST"
	case 15:
		return "SCALING"
	case 16:
		return "SUBMIT_MODEL_FAILED"
	case 17:
		return "DEPLOY_SERVICE_FAILED"
	case 18:
		return "CHECK_INIT"
	case 19:
		return "CHECK_RUNNING"
	case 20:
		return "CHECK_RUNNING_COMPLETED"
	case 21:
		return "CHECK_FAILED"

	default:
		return strconv.Itoa(status)
	}
}

func transferCloudbrain2LogOrder(direction entity.Direction) string {
	if direction == entity.UP {
		return "asc"
	} else if direction == entity.DOWN {
		return "desc"
	}
	return ""
}

func (c CloudbrainTwoClusterAdapter) GetLog(opts entity.ClusterLogOpts) (*entity.ClusterLog, error) {
	baseLine := fmt.Sprint(opts.BaseLine)
	order := transferCloudbrain2LogOrder(opts.Direction)
	if opts.IsHeadRequest() {
		baseLine = ""
		order = "asc"
	} else if opts.IsBottomRequest() {
		baseLine = ""
		order = "desc"
	}
	result, err := getModelartsTrainJob(opts.JobId, opts.VersionID, baseLine, order, int(opts.Lines), opts.LogFileName)
	if err != nil {
		log.Error("getModelartsTrainJob(%s) failed:%v", opts.JobId, err)
		return nil, err
	}
	lines := int64(result.Lines)
	return &entity.ClusterLog{
		Content:   result.Content,
		StartLine: result.StartLine,
		EndLine:   result.EndLine,
		Lines:     lines,
	}, nil
}

func getModelartsTrainJob(jobID string, versionID int64, baseLine string, order string, lines int, logFileName string) (*models.GetTrainJobLogResult, error) {
	result, err := modelarts.GetTrainJobLog(jobID, strconv.FormatInt(versionID, 10), baseLine, logFileName, order, lines)
	if err != nil {
		log.Error("GetTrainJobLog(%s) failed:%v", jobID, err.Error())
		return nil, err
	}

	return result, err
}

func (c CloudbrainTwoClusterAdapter) GetLogDownloadInfo(opts entity.ClusterLogDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
	var err error
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)

	//查找日志文件
	files := getLogFilesInStorage(helper, opts.ObjectKeyPrefix, ".log")
	if len(files) == 0 {
		//此时未找符合条件的文件
		return nil, nil
	}

	//只有一个日志文件时直接使用obs的下载链接
	if len(files) == 1 {
		var url string
		url, err = helper.GetSignedDownloadUrl(files[0].RelativePath)
		if err != nil {
			log.Error("GetObsCreateSignedUrlByBucketAndKey failed when GetLogDownloadInfo opts=%+v: err=%v", opts, err)
			return nil, err
		}
		return &entity.FileDownloadInfo{
			DownloadUrl: url,
		}, nil
	}

	readerList := make([]entity.FileReader, 0)
	defer func() {
		if err != nil {
			for _, r := range readerList {
				if r.Reader != nil {
					r.Reader.Close()
				}
			}
		}
	}()
	//多个文件时需要打包后下载
	for _, file := range files {
		//获取日志reader
		var reader io.ReadCloser
		reader, err = helper.OpenFile(file.RelativePath)
		if err != nil {
			log.Error("GetLogDownloadInfo OpenFile err.opts=%+v,err =%v", opts, err)
			return nil, err
		}
		readerList = append(readerList, entity.FileReader{
			Reader: reader,
			Name:   file.FileName,
		})
	}
	return &entity.FileDownloadInfo{
		Readers:        readerList,
		ResultType:     entity.FileTypeZIP,
		ResultFileName: opts.DisplayJobName + ".zip",
	}, nil
}

func (c CloudbrainTwoClusterAdapter) GetSingleOutputDownloadInfo(opts entity.ClusterSingleOutputDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
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

func (c CloudbrainTwoClusterAdapter) DownloadAllOutput(opts entity.DownloadOutputOpts) error {
	return DownloadAllOutput(opts)
}

func (c CloudbrainTwoClusterAdapter) GetTrainJobOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	return c.GetNoteBookOperationProfile(opts)
}

func (c CloudbrainTwoClusterAdapter) GetResourceUsage(opts entity.ClusterResourceUsageOpts) (*entity.ResourceUsage, error) {
	result, err := modelarts.GetTrainJobMetricStatistic(opts.JobId, strconv.FormatInt(opts.VersionID, 10), opts.LogFileName)
	if err != nil {
		log.Error("GetTrainJobMetricStatistic(%s) failed:%v", opts.JobId, err.Error())
		return nil, err
	}
	return transferModelartsMetricsToStandard(result), nil
}

func transferModelartsMetricsToStandard(result *models.GetTrainJobMetricStatisticResult) *entity.ResourceUsage {
	m := make([]entity.MetricsInfo, 0)
	for i := 0; i < len(result.MetricsInfo); i++ {
		valArray := result.MetricsInfo[i].Value
		temp := make([]float32, len(valArray))
		for j := 0; j < len(valArray); j++ {
			val, err := strconv.ParseFloat(valArray[j], 32)
			if err != nil {
				log.Error("parse metrics value error, val=%v err=%v result=%+v", valArray[j], err, result)
				return nil
			}
			temp[j] = float32(val)
		}

		m = append(m, entity.MetricsInfo{
			Name:  result.MetricsInfo[i].Metric,
			Value: temp,
		})
	}
	return &entity.ResourceUsage{
		Interval:    result.Interval,
		MetricsInfo: m,
	}
}

func (c CloudbrainTwoClusterAdapter) GetNodeInfo(opts entity.ClusterNodeInfoOpts) ([]entity.AITaskNodeInfo, error) {
	resultLogFile, err := modelarts.GetTrainJobLogFileNames(opts.JobId, strconv.FormatInt(opts.VersionId, 10))
	if err != nil {
		log.Error("GetTrainJobLogFileNames(%s) failed:%v", opts.JobId, err.Error())
		return nil, nil
	}
	if resultLogFile == nil {
		return nil, nil
	}
	res := make([]entity.AITaskNodeInfo, len(resultLogFile.LogFileList))
	for i := 0; i < len(resultLogFile.LogFileList); i++ {
		res[i] = entity.AITaskNodeInfo{LogFileName: resultLogFile.LogFileList[i]}
	}
	return res, nil
}

func (c CloudbrainTwoClusterAdapter) GetOutput(opts entity.ClusterOutputOpts) (*entity.ClusterAITaskOutput, error) {
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

func (c CloudbrainTwoClusterAdapter) GetAllOutput(opts entity.ClusterOutputOpts) (*entity.AllAITaskOutput, error) {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	fileList, err := helper.GetAllObjectsUnderDir(path.Join(opts.ObjectKeyPrefix, opts.ParentDir))
	if err != nil {
		log.Error("GetAllObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.ObjectKeyPrefix, err)
		return nil, err
	}
	return &entity.AllAITaskOutput{FileList: fileList}, nil
}
