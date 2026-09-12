package cluster

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type CloudbrainOneClusterAdapter struct {
}

func init() {
	//注册到一个Map
	AddCluster(entity.OpenICloudbrainOne, new(CloudbrainOneClusterAdapter))
}

func (c CloudbrainOneClusterAdapter) CreateNoteBook(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	jobResult, err := cloudbrain.CreateJob(req.Name, convertNoteBookReq2CloudbrainOne(req))
	if err != nil {
		log.Error("CreateNoteBook failed: %v", err.Error())
		return nil, err
	}
	return convertCloudbrainOne2NoteBookRes(jobResult), nil
}

func (c CloudbrainOneClusterAdapter) CreateOnlineInfer(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) CreateModelExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) CreateSdFinetune(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) CreateComfyuiExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) QueryModelExperience(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return c.QueryNoteBook(opts)
}

func (c CloudbrainOneClusterAdapter) StopModelExperience(opts entity.JobIdAndVersionId) error {
	return nil
}
func (c CloudbrainOneClusterAdapter) DeleteModelExperience(opts entity.JobIdAndVersionId) error {
	return nil

}

func (c CloudbrainOneClusterAdapter) GetSelfEndPointUrl(jobId string) (string, error) {
	return "", nil
}

func (c CloudbrainOneClusterAdapter) CreateGeneralTask(req entity.CreateGeneralTaskRequest, trace *entity.TraceInfo) (*entity.CreateGeneralTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) GetNotebookImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	return nil, true, nil
}

func (c CloudbrainOneClusterAdapter) GetTrainImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error) {
	return c.GetNotebookImages(req, queues...)
}

var SubTaskName = "task1"

func convertNoteBookReq2CloudbrainOne(req entity.CreateNoteBookTaskRequest) models.CreateJobParams {
	var command = `pip3 install jupyterlab==3 -i https://pypi.tuna.tsinghua.edu.cn/simple;pip3 install -U "nbclassic>=0.2.8" -i https://pypi.tuna.tsinghua.edu.cn/simple;service ssh stop;jupyter lab --ServerApp.shutdown_no_activity_timeout=` + setting.CullIdleTimeout + ` --TerminalManager.cull_inactive_timeout=` + setting.CullIdleTimeout + ` --TerminalManager.cull_interval=` + setting.CullInterval + ` --MappingKernelManager.cull_idle_timeout=` + setting.CullIdleTimeout + ` --MappingKernelManager.cull_interval=` + setting.CullInterval + ` --MappingKernelManager.cull_connected=True --MappingKernelManager.cull_busy=True --no-browser --ip=0.0.0.0 --allow-root --notebook-dir="/code" --port=80 --ServerApp.token="" --LabApp.token="" --ServerApp.allow_origin="self https://cloudbrain.pcl.ac.cn" `
	t := req.Tasks[0]

	return models.CreateJobParams{
		JobName:    t.Name,
		RetryCount: 1,
		GpuType:    t.Spec.QueueCode,
		Image:      t.ImageUrl,
		TaskRoles: []models.TaskRole{
			{
				Name:                  SubTaskName,
				TaskNumber:            1,
				MinSucceededTaskCount: 1,
				MinFailedTaskCount:    1,
				CPUNumber:             t.Spec.CpuCores,
				GPUNumber:             t.Spec.AccCardsNum,
				MemoryMB:              int(t.Spec.MemGiB * 1024),
				ShmMB:                 int(t.Spec.ShareMemGiB * 1024),
				Command:               command,
				NeedIBDevice:          false,
				IsMainRole:            false,
				UseNNI:                false,
			},
		},
		Volumes: convertContainerDataArray2Volume(t.Code, t.Datasets, t.PreTrainModel, t.OutPut),
	}
}

func convertContainerDataArray2Volume(containerDataArray ...[]entity.ContainerData) []models.Volume {
	r := make([]models.Volume, 0)
	for _, array := range containerDataArray {
		for _, d := range array {
			r = append(r, convertContainerData2Volume(d))
		}
	}
	return r
}

func convertContainerData2Volume(d entity.ContainerData) models.Volume {
	return models.Volume{
		HostPath: models.StHostPath{
			Path:      d.RealPath,
			MountPath: d.ContainerPath,
			ReadOnly:  d.ReadOnly,
		},
	}
}

func convertCloudbrainOne2NoteBookRes(res *models.CreateJobResult) *entity.CreateNoteBookTaskResponse {
	playload := res.Payload
	return &entity.CreateNoteBookTaskResponse{
		JobID:  playload["jobId"].(string),
		Status: string(models.JobWaiting),
	}
}

func (c CloudbrainOneClusterAdapter) RestartNoteBook(string, int64, *entity.TraceInfo) (*entity.RestartNoteBookTaskResponse, error) {

	return nil, nil
}
func (c CloudbrainOneClusterAdapter) DeleteNoteBook(entity.JobIdAndVersionId) error {
	return nil
}

func (c CloudbrainOneClusterAdapter) StopNoteBook(opts entity.JobIdAndVersionId) error {
	err := cloudbrain.StopJob(opts.JobID)
	if err != nil {
		log.Error("StopNoteBook(%s) failed:%v", opts, err)
		return err
	}
	return nil
}

func (c CloudbrainOneClusterAdapter) QueryNoteBook(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	if opts.JobID == "" {
		log.Error("jobid is empty")
		return nil, errors.New("jobid is empty")
	}
	jobResult, err := cloudbrain.GetJob(opts.JobID)
	if err != nil {
		log.Error("QueryNoteBook failed:%v", err)
		return nil, err
	}
	if err != nil {
		log.Error("ConvertToJobResultPayload failed:%v", err)
		return nil, err
	}
	return entity.ConvertCloudbrainOneNotebookResponse(jobResult.Payload)
}

func (c CloudbrainOneClusterAdapter) QueryNoteBookByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	jobResult, err := cloudbrain.GetJobListByName(jobName)
	if err != nil {
		log.Error("GetJobListByName failed:%v", err)
		return nil, err
	}
	result, err := models.ConvertToJobListResultPayload(jobResult.Payload)
	if err != nil {
		log.Error("ConvertToJobListResultPayload failed:%v", err)
		return nil, err
	}
	r := make([]*entity.QueryTaskResponse, 0)
	for i := 0; i < len(result.Jobs); i++ {
		if result.Jobs[i].Name == jobName {
			r = append(r, entity.ConvertCloudbrainOneQueryNotebookByNameResponse(result.Jobs[i]))
		}
	}
	return r, nil
}

func (c CloudbrainOneClusterAdapter) QueryOnlineInfer(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) GetNoteBookLog(jobId string) (*entity.ClusterLog, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) GetNoteBookUrl(jobId string) (string, error) {
	return setting.DebugServerHost + "jpylab_" + jobId + "_" + models.SubTaskName, nil
}

func (c CloudbrainOneClusterAdapter) GetNoteBookOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	if opts.JobID == "" {
		log.Error("jobid is empty")
		return nil, errors.New("jobid is empty")
	}
	jobResult, err := cloudbrain.GetJob(opts.JobID)
	if err != nil {
		log.Error("QueryNoteBook failed:%v", err)
		return nil, err
	}
	result, err := models.ConvertToJobResultPayload(jobResult.Payload)
	if err != nil {
		log.Error("ConvertToJobResultPayload failed:%v", err)
		return nil, err
	}

	taskRoles := result.TaskRoles
	taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
	ExitDiagnostics := taskRes.TaskStatuses[0].ExitDiagnostics

	return parseDiagnosticsToOperationProfile(result.JobStatus.AppExitDiagnostics, ExitDiagnostics), nil
}

func parseDiagnosticsToOperationProfile(appExitDiagnostics string, exitDiagnostics string) *entity.OperationProfile {
	if appExitDiagnostics == "" {
		return nil
	}
	diagnostics := entity.CloudbrainOneAppExitDiagnostics{}
	err := json.Unmarshal([]byte(appExitDiagnostics), &diagnostics)
	if err != nil {
		log.Error("json.Unmarshal appExitDiagnostics err.%v", err)
		return nil
	}
	events := make([]entity.ProfileEvent, 0)
	podEvents := diagnostics.PodEvents.Task10
	for i := 0; i < len(podEvents); i++ {
		e := podEvents[i]
		if e.Message == "" {
			continue
		}
		events = append(events, entity.ProfileEvent{
			Message: e.Message,
			Reason:  e.Reason,
			Action:  e.Action,
		})
	}
	extras := diagnostics.Extras
	for i := 0; i < len(extras); i++ {
		e := extras[i]
		if e.Message == "" {
			continue
		}
		events = append(events, entity.ProfileEvent{
			Message: e.Message,
			Reason:  e.Reason,
			Action:  e.Action,
		})
	}
	if exitDiagnostics != "" {
		events = append(events, entity.ProfileEvent{
			Message: exitDiagnostics,
			Reason:  "Error",
		})
	}
	return &entity.OperationProfile{Events: events}
}

func (c CloudbrainOneClusterAdapter) CreateTrainJob(req entity.CreateTrainTaskRequest, trace *entity.TraceInfo) (*entity.CreateTrainTaskResponse, error) {
	jobResult, err := cloudbrain.CreateJob(req.Name, convertTrainJobReq2CloudbrainOne(req))
	if err != nil {
		log.Error("CreateTrainJob failed: %v", err.Error())
		return nil, err
	}
	return convertCloudbrainOne2TrainJobRes(jobResult), nil
}

func convertTrainJobReq2CloudbrainOne(req entity.CreateTrainTaskRequest) models.CreateJobParams {
	var command = getTrainJobCommand(req)
	t := req.Tasks[0]

	return models.CreateJobParams{
		JobName:    t.Name,
		RetryCount: 1,
		GpuType:    t.Spec.QueueCode,
		Image:      t.ImageUrl,
		TaskRoles: []models.TaskRole{
			{
				Name:                  SubTaskName,
				TaskNumber:            1,
				MinSucceededTaskCount: 1,
				MinFailedTaskCount:    1,
				CPUNumber:             t.Spec.CpuCores,
				GPUNumber:             t.Spec.AccCardsNum,
				MemoryMB:              int(t.Spec.MemGiB * 1024),
				ShmMB:                 int(t.Spec.ShareMemGiB * 1024),
				Command:               command,
				NeedIBDevice:          false,
				IsMainRole:            false,
				UseNNI:                false,
			},
		},
		Volumes: convertContainerDataArray2Volume(t.Code, t.Datasets, t.PreTrainModel, t.OutPut),
	}
}

func getTrainJobCommand(req entity.CreateTrainTaskRequest) string {
	form := req.Tasks[0]
	var command string
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params

	var param string
	if params.Parameter != nil && len(params.Parameter) != 0 {
		for _, parameter := range params.Parameter {
			param += " --'" + parameter.Label + "'='" + parameter.Value + "'"
		}
	}

	//启智GPU训练暂未支持多模型，此处先视为只会有一个模型文件
	if form.PreTrainModel != nil && len(form.PreTrainModel) > 0 {
		param += " --ckpt_url" + "=" + "'/pretrainmodel/" + form.PreTrainModel[0].Name + "'"
	}

	logPath := cloudbrain.ModelMountPath
	if form.LogPath != nil && len(form.LogPath) > 0 {
		logPath = form.LogPath[0].ContainerPath
	}
	command += "python -u /code/" + bootFile + param + " > " + logPath + "/" + req.DisplayJobName + "-" + cloudbrain.LogFile

	return command
}

func convertCloudbrainOne2TrainJobRes(res *models.CreateJobResult) *entity.CreateTrainTaskResponse {
	playload := res.Payload
	return &entity.CreateTrainTaskResponse{
		JobID:  playload["jobId"].(string),
		Status: string(models.JobWaiting),
	}
}

func (c CloudbrainOneClusterAdapter) DeleteTrainJob(entity.JobIdAndVersionId) error {
	return nil
}
func (c CloudbrainOneClusterAdapter) StopTrainJob(opts entity.JobIdAndVersionId) error {
	err := cloudbrain.StopJob(opts.JobID)
	if err != nil {
		log.Error("StopNoteBook(%s) failed:%v", opts, err)
		return err
	}
	return nil
}
func (c CloudbrainOneClusterAdapter) QueryTrainJob(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error) {
	return c.QueryNoteBook(opts)
}

func (c CloudbrainOneClusterAdapter) QueryTrainJobByJobName(jobName string) ([]*entity.QueryTaskResponse, error) {
	jobResult, err := cloudbrain.GetJobListByName(jobName)
	if err != nil {
		log.Error("GetJobListByName failed:%v", err)
		return nil, err
	}
	result, err := models.ConvertToJobListResultPayload(jobResult.Payload)
	if err != nil {
		log.Error("ConvertToJobListResultPayload failed:%v", err)
		return nil, err
	}
	r := make([]*entity.QueryTaskResponse, 0)
	for i := 0; i < len(result.Jobs); i++ {
		if result.Jobs[i].Name == jobName {
			r = append(r, entity.ConvertCloudbrainOneQueryNotebookByNameResponse(result.Jobs[i]))
		}
	}
	return r, nil
}

func (c CloudbrainOneClusterAdapter) GetVisualizeUrl(jobId string) (string, error) {
	return "", nil
}

func (c CloudbrainOneClusterAdapter) GetLog(opts entity.ClusterLogOpts) (*entity.ClusterLog, error) {
	if opts.Lines <= 0 || opts.ObjectKeyPrefix == "" {
		return nil, nil
	}
	//获取任务退出信息
	existStr := getCloudbrainOneExitDiagnostics(opts.JobId)

	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)

	//查找日志文件
	files := getLogFilesInStorage(helper, opts.ObjectKeyPrefix, "log.txt")
	if len(files) == 0 {
		//此时未找符合条件的文件
		startLine, endLine, lines := handleOverLines(opts)
		return &entity.ClusterLog{
			Content:   existStr,
			StartLine: fmt.Sprint(startLine),
			EndLine:   fmt.Sprint(endLine),
			Lines:     lines,
		}, nil
	}

	//默认选择第一个文件
	file := files[0]

	//计算开始行和结束行
	startLine, endLine := findStartAndEnd(opts, file.RelativePath, helper)

	//获取日志内容
	result, realEndLine, contentLines := getLogInStorage(startLine, endLine, helper, file.RelativePath)

	//处理到达顶部或者底部时的情况
	if contentLines == 0 {
		startLine, realEndLine, contentLines = handleOverLines(opts)
	}

	return &entity.ClusterLog{
		Content:   result,
		StartLine: fmt.Sprint(startLine),
		EndLine:   fmt.Sprint(realEndLine),
		Lines:     contentLines,
	}, nil
}

func (c CloudbrainOneClusterAdapter) GetLogDownloadInfo(opts entity.ClusterLogDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
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

func (c CloudbrainOneClusterAdapter) GetSingleOutputDownloadInfo(opts entity.ClusterSingleOutputDownloadInfoOpts) (*entity.FileDownloadInfo, error) {
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

func (c CloudbrainOneClusterAdapter) DownloadAllOutput(opts entity.DownloadOutputOpts) error {
	return DownloadAllOutput(opts)
}

func (c CloudbrainOneClusterAdapter) GetNodeInfo(opts entity.ClusterNodeInfoOpts) ([]entity.AITaskNodeInfo, error) {
	return nil, nil
}

func (c CloudbrainOneClusterAdapter) GetResourceUsage(opts entity.ClusterResourceUsageOpts) (*entity.ResourceUsage, error) {
	return &entity.ResourceUsage{
		Interval:    0,
		MetricsInfo: []entity.MetricsInfo{},
	}, nil
}

func getLogInStorage(startLine, endLine int64, helper storage_helper.StorageHelper, path string) (content string, realEndLine int64, total int64) {
	log.Info("getLogInStorage path=%s", path)
	reader, err := helper.OpenFile(path)
	if err != nil {
		log.Info("elper.OpenFile error,path=%s err=%v", path, err)
		return "", 0, 0
	}
	defer reader.Close()

	return GetLocalLog(reader, startLine, endLine)
}

func handleOverLines(opts entity.ClusterLogOpts) (int64, int64, int64) {
	var startLine, endLine int64
	if opts.Direction == entity.DOWN {
		endLine = opts.BaseLine
		startLine = endLine + 1 - opts.Lines
		if startLine < 1 {
			startLine = 1
		}
	} else {
		startLine = 1
		endLine = startLine + opts.Lines - 1
	}
	return startLine, endLine, 0
}

func getCloudbrainOneExitDiagnostics(jobId string) string {
	jobResult, _ := cloudbrain.GetJob(jobId)
	if jobResult != nil {
		jobRes, _ := models.ConvertToJobResultPayload(jobResult.Payload)
		taskRoles := jobRes.TaskRoles
		taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
		return taskRes.TaskStatuses[0].ExitDiagnostics
	}
	return ""
}

// findStartAndEnd 基于baseLine，根据方向向上或者向下计算
func findStartAndEnd(opts entity.ClusterLogOpts, filePath string, helper storage_helper.StorageHelper) (startLine int64, endLine int64) {
	baseLine := opts.BaseLine
	if opts.Direction == entity.UP {
		if baseLine == 0 {
			endLine = getAllLineFromFile(helper, filePath)
		} else {
			endLine = baseLine - 1
		}
		startLine = endLine - opts.Lines + 1
		if startLine <= 0 {
			startLine = 1
		}
	} else {
		startLine = baseLine + 1
		endLine = startLine + opts.Lines - 1
	}
	return startLine, endLine
}

func (c CloudbrainOneClusterAdapter) GetTrainJobOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error) {
	return c.GetNoteBookOperationProfile(opts)
}

func (c CloudbrainOneClusterAdapter) GetOutput(opts entity.ClusterOutputOpts) (*entity.ClusterAITaskOutput, error) {
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

func (c CloudbrainOneClusterAdapter) GetAllOutput(opts entity.ClusterOutputOpts) (*entity.AllAITaskOutput, error) {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	fileList, err := helper.GetAllObjectsUnderDir(path.Join(opts.ObjectKeyPrefix, opts.ParentDir))
	if err != nil {
		log.Error("GetAllObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.ObjectKeyPrefix, err)
		return nil, err
	}
	return &entity.AllAITaskOutput{FileList: fileList}, nil
}
