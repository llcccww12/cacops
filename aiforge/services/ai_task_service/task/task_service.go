package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/cluster"
	creation_context "code.gitea.io/gitea/services/ai_task_service/context"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/lock"
	"code.gitea.io/gitea/services/role"
)

type QueryFunc func(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error)
type QueryListFunc func(string) ([]*entity.QueryTaskResponse, error)
type DeleteFunc func(opts entity.JobIdAndVersionId) error
type StopFunc func(opts entity.JobIdAndVersionId) error
type GetLogFunc func(entity.ClusterLogOpts) (*entity.ClusterLog, error)
type GetLogDownloadInfoFunc func(entity.ClusterLogDownloadInfoOpts) (*entity.FileDownloadInfo, error)
type GetNotebookUrlFunc func(string) (string, error)
type GetNodeInfoFunc func(entity.ClusterNodeInfoOpts) ([]entity.AITaskNodeInfo, error)
type GetOutputFunc func(entity.ClusterOutputOpts) (*entity.ClusterAITaskOutput, error)
type GetAllOutputFunc func(entity.ClusterOutputOpts) (*entity.AllAITaskOutput, error)
type GetSingleOutputDownloadInfoFunc func(req entity.ClusterSingleOutputDownloadInfoOpts) (*entity.FileDownloadInfo, error)
type DownloadAllOutputFunc func(req entity.DownloadOutputOpts) error
type GetOperationProfileFunc func(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error)
type GetResourceUsageFunc func(entity.ClusterResourceUsageOpts) (*entity.ResourceUsage, error)

func BuildAITaskInfo(cloudbrainId int64) (*entity.AITaskDetailInfo, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		log.Error("QueryTaskInfo GetCloudbrainByCloudbrainID err.id=%d err=%v", cloudbrain.ID, err)

		return nil, err
	}
	return BuildAITaskByCloudbrain(cloudbrain)
}

func BuildAITaskByCloudbrain(cloudbrain *models.Cloudbrain) (*entity.AITaskDetailInfo, error) {
	creator, err := models.GetUserByID(cloudbrain.UserID)
	if err != nil {
		return nil, err
	}
	c := GetDetailConfigInfoByCloudbrain(cloudbrain)
	return buildAITaskInfo(cloudbrain, creator, c)
}

func buildAITaskInfo(task *models.Cloudbrain, creator *models.User, config *entity.AITaskDetailConfigInfo) (*entity.AITaskDetailInfo, error) {
	spec, err := resource.GetCloudbrainSpec(task.ID)
	if err != nil {
		log.Error("buildAITaskInfo GetCloudbrainSpec error,id =%d ,err =%v", task.ID, err)
		return nil, err
	}
	datasets := []*models.DatasetRegistryDownload{}
	pretrainModelList := []*models.Model4Show{}
	var datasetNames []string
	var repoOwnerName string
	var repoName string
	var repoAlias string
	if task.RepoID != 0 {
		task.GetRepository()
		if task.Repo != nil {
			repoOwnerName = task.Repo.OwnerName
			repoName = task.Repo.Name
			repoAlias = task.Repo.Alias
		}
	}
	if task.Uuid != "" {
		datasets = GetCloudBrainDataSetRegistryInfo(task)
		for i := 0; i < len(datasets); i++ {
			if datasets[i].DatasetName == "" {
				continue
			}
			datasetNames = append(datasetNames, datasets[i].DatasetName)
		}
	}
	var pretrainModelNames []string
	if task.ModelId != "" {
		pretrainModelList = GetModelDownload(task)
		for i := 0; i < len(pretrainModelList); i++ {
			if pretrainModelList[i].Name == "" {
				continue
			}
			pretrainModelNames = append(pretrainModelNames, pretrainModelList[i].Name)
		}
	}

	paramKeys := make([]string, 0)
	if task.Parameters != "" {
		params := parseAITaskParameters(task.Parameters)
		if params != nil {
			for i := 0; i < len(params.Parameter); i++ {
				paramKeys = append(paramKeys, params.Parameter[i].Label)
			}
		}

	}
	code := GenerateSDKCode(
		entity.SDKCodeOpts{
			DatasetNames:       datasetNames,
			PretrainModelNames: pretrainModelNames,
			ParameterKeys:      paramKeys,
			JobType:            models.JobType(task.JobType),
		})
	n := 1
	if task.WorkServerNumber > 1 {
		n = task.WorkServerNumber
	}
	imageName := task.Image

	// 取对应的image数据，name为tag
	if task.Image != "" {
		imageInfo, _ := models.GetImageByPlace(task.Image)
		imageName = imageInfo.Tag
	}

	if imageName == "" {
		imageName = task.Image
	}

	imageUrl := task.Image
	imageId := task.ImageID
	if imageName == "" {
		imageName = task.EngineName
	}
	if imageUrl == "" {
		imageUrl = task.EngineName
	}
	if imageId == "" && task.EngineID > 0 {
		imageId = fmt.Sprint(task.EngineID)
	}

	baseConfig := &entity.AITaskBaseConfig{}
	if config != nil && config.BaseConfig != nil {
		baseConfig = config.BaseConfig
	}
	// 获取容器存储配额
	codeSizeLimit, outputSizeLimit := role.GetUserContainerStorageLimits(task.UserID)
	return &entity.AITaskDetailInfo{
		ID:                 task.ID,
		JobID:              task.JobID,
		Status:             task.Status,
		DetailedStatus:     task.DetailedStatus,
		JobType:            task.JobType,
		DisplayJobName:     html.EscapeString(task.DisplayJobName),
		FormattedDuration:  task.TrainJobDuration,
		ComputeSource:      task.GetStandardComputeSource(),
		PreVersionName:     task.PreVersionName,
		CurrentVersionName: task.VersionName,
		WorkServerNumber:   n,
		Spec:               convert.ToSpecification(spec),
		DatasetList:        datasets,
		PretrainModelList:  pretrainModelList,
		SDKCode:            code,
		AICenter:           task.AiCenter,
		BootFile:           task.BootFile,
		Cluster:            string(entity.GetClusterTypeFromCloudbrainType(task.Type)),
		Parameters:         parseAITaskParameters(task.Parameters),
		CreatedUnix:        task.CreatedUnix,
		CodePath:           baseConfig.GetContainerPath(entity.ContainerCode),
		DatasetPath:        baseConfig.GetContainerPath(entity.ContainerDataset),
		PretrainModelPath:  baseConfig.GetContainerPath(entity.ContainerPreTrainModel),
		OutputPath:         baseConfig.GetContainerPath(entity.ContainerOutPutPath),
		CodeUrl:            task.RemoteCodeUrl,
		PretrainModelName:  html.EscapeString(task.ModelName),
		PretrainModelUrl:   task.PreTrainModelUrl,
		PretrainModelId:    task.ModelId,
		StartTime:          task.StartTime,
		EndTime:            task.EndTime,
		Description:        html.EscapeString(task.Description),
		FailedReason:       html.EscapeString(task.FailedReason),
		CommitID:           task.CommitID,
		BranchName:         task.BranchName,
		ImageName:          html.EscapeString(imageName),
		ImageID:            html.EscapeString(imageId),
		ImageUrl:           html.EscapeString(imageUrl),
		CreatorName:        creator.GetDisplayName(),
		EngineName:         task.EngineName,
		UserId:             task.UserID,
		AppName:            html.EscapeString(task.AppName),
		HasInternet:        task.HasInternet,
		TimeLimit:          task.TimeLimit,
		DefaultTimeLimit:   int(setting.MaxDuration / 3600),
		Port:               task.Port,
		EndPoint:           task.EndPoint,
		VisualizeRequired:  task.VisualizeRequired,
		AimRequired:        task.AimRequired,
		IsFileNotebook:     task.IsFileNoteBookTask(),
		SourceID:           task.SourceID,
		RepoID:             task.RepoID,
		RepoOwnerName:      repoOwnerName,
		RepoName:           repoName,
		RepoAlias:          repoAlias,
		CodeSizeLimit:      codeSizeLimit,
		OutputSizeLimit:    outputSizeLimit,
	}, nil
}

func parseAITaskParameters(paramStr string) *models.Parameters {
	if paramStr == "" {
		return nil
	}
	var parameters = &models.Parameters{}
	err := json.Unmarshal([]byte(paramStr), parameters)
	if err != nil {
		log.Error("parseAITaskParameters. Failed to Unmarshal Parameters: %s (%v)", paramStr, err)
		return &models.Parameters{}
	}
	if parameters != nil {
		for i := 0; i < len(parameters.Parameter); i++ {
			parameters.Parameter[i].Label = html.EscapeString(parameters.Parameter[i].Label)
			parameters.Parameter[i].Value = html.EscapeString(parameters.Parameter[i].Value)
		}
	}
	return parameters
}

func QueryTaskEarlyVersionList(id int64) ([]*entity.AITaskDetailInfo, error) {
	task, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return nil, err
	}
	if !task.IsAllowedToCreateMultipleVersions() {
		return []*entity.AITaskDetailInfo{}, nil
	}
	taskList, err := models.GetCloudbrainEarlyVersionList(task)
	if err != nil {
		log.Error("QueryTaskEarlyVersionList GetCloudbrainEarlyVersionList err.id=%d err=%v", id, err)
		return nil, err
	}
	resultList := make([]*entity.AITaskDetailInfo, len(taskList))
	for i := 0; i < len(taskList); i++ {
		t, err := BuildAITaskByCloudbrain(taskList[i])
		if err != nil {
			log.Error("QueryTaskEarlyVersionList convertCloudbrainToAITaskDetailInfo err.id=%d currentId=%d err=%v", id, taskList[i].ID, err)
			return nil, err
		}
		resultList[i] = t
	}
	return resultList, nil
}

func QueryTaskBriefInfo(id int64) (*entity.AITaskBriefInfo, error) {
	task, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		log.Error("QueryTaskBriefInfo GetCloudbrainByCloudbrainID err.id=%d err=%v", id, err)
		return nil, err
	}
	task.LoadSpec()
	return entity.ConvertCloudbrainToAITaskBriefInfo(task), nil
}

func StopAITaskByJobNameFromRemote(task *models.Cloudbrain, queryFunc QueryListFunc, stopFunc StopFunc) error {
	if task.IsTerminal() {
		return nil
	}
	res, err := queryFunc(task.JobName)
	if err != nil {
		log.Error("query from remote err.cloudbrainID = %d  err=%v", task.ID, err)
		return err
	}
	if res == nil {
		log.Error("query from remote failed,response is empty,cloudbrainID = %d ", task.ID)
		return errors.New("response is empty")
	}

	for _, v := range res {
		//如果匹配到的任务开始时间比当前云脑任务的创建时间还早五分钟以上，那认为非同一个任务
		if models.IsCloudbrainWaitingOrRunning(TransAITaskStatus(v.Status)) {
			if v.StartedAt < task.CreatedUnix-5*60 {
				continue
			}
			if err = stopFunc(entity.JobIdAndVersionId{JobID: v.JobId, VersionID: v.VersionId, ComputeResource: task.ComputeResource, TaskID: task.ID}); err != nil {
				log.Error("stop task err. name=%s jobId=%s err=%v", task.JobName, v.JobId)
				return err
			}

		}

	}
	models.UpdateJobStatus(task.ID, models.LocalStatusFailed, task.Status)
	return nil
}

const DEFAULT_DETAILED_STATUS = "-"

func UpdateByQueryResponse(res *entity.QueryTaskResponse, task *models.Cloudbrain) error {
	if res.JobId == "" {
		return nil
	}

	//云脑一的调试任务状态变为RUNNING后存在短时间内不能访问的情况，需要特殊处理
	if res.Status == string(models.JobRunning) && task.Status == string(models.JobWaiting) && (task.JobType == string(models.JobTypeDebug) || task.JobType == string(models.JobTypeOnlineInference)) && task.Type == models.TypeCloudBrainOne {
		if !isCloudbrainOneNotebookReady(res.JobId) {
			log.Info("task(%s) is not ready", task.DisplayJobName)
			return nil
		}
	}

	if res.CenterId != "" && res.CenterName != "" {
		task.AiCenter = res.CenterId + "+" + res.CenterName
	}
	task.QueueCode = res.QueueCode
	oldStatus := task.Status
	newStatus := TransAITaskStatus(res.Status)

	task.Status = newStatus
	task.DetailedStatus = res.DetailedStatus
	if res.DetailedStatus == "" || res.DetailedStatus == res.Status {
		task.DetailedStatus = DEFAULT_DETAILED_STATUS
	}
	if res.StartedAt > 0 && task.StartTime == 0 {
		task.StartTime = res.StartedAt
	}
	if res.StartedAt > 0 && res.CompletedAt > 0 {
		task.EndTime = res.CompletedAt
	}
	task.ComputeAndSetDuration()
	task.CorrectCreateUnix()

	task.JobID = res.JobId

	task.RemoteCodeUrl = res.CodeUrl
	task.DataUrl = res.DataUrl
	task.ContainerID = res.ContainerID
	task.ContainerIp = res.ContainerIP
	task.VersionID = res.VersionId

	err := models.UpdateJob(task)
	if err != nil {
		log.Error("UpdateJob(%s) failed:%v", task.DisplayJobName, err)
		return err
	}
	if oldStatus != task.Status {
		go notification.NotifyChangeCloudbrainStatus(task, oldStatus)
	}
	go correctAITaskSpec(task)
	return nil
}

var noteBookOKMap = make(map[string]int, 20)
var noteBookFailMap = make(map[string]int, 20)

// if a task notebook url can get successfulCount times,  the notebook can browser.
const successfulCount = 3
const maxSuccessfulCount = 10

// 云脑一存在状态为RUNNING但实际不可用的情况，且存在访问调试链接成功后又失败的情况，因此需要一段时间内多次成功才认为调试可用
// 下列代码来源于services/cloudbrain/cloudbrainTask/sync_status.go:118
// func isNoteBookReady(task *models.Cloudbrain) bool
// 为了解决循环引用copy了一份到此类，稍有改编
func isCloudbrainOneNotebookReady(jobId string) bool {
	url, err := new(cluster.CloudbrainOneClusterAdapter).GetNoteBookUrl(jobId)
	if err != nil {
		return false
	}
	if url == "" {
		return false
	}
	res, err := http.Get(url)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		count := noteBookOKMap[jobId]
		if count == 0 { //如果是第一次成功，把失败数重置为0
			noteBookFailMap[jobId] = 0
		}

		if count < successfulCount-1 || (noteBookFailMap[jobId] == 0 && count < maxSuccessfulCount-1) {
			noteBookOKMap[jobId] = count + 1
			return false
		} else {
			log.Info("notebook success count:" + strconv.Itoa(count) + ",fail count:" + strconv.Itoa(noteBookFailMap[jobId]))
			delete(noteBookOKMap, jobId)
			delete(noteBookFailMap, jobId)
			return true
		}

	} else {
		noteBookFailMap[jobId] += 1
	}
	return false

}

func DelTask(id int64, deleteRemote DeleteFunc) error {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return err
	}

	if !cloudbrain.IsTerminal() {
		return errors.New("task not stopped")
	}

	//删除远端记录
	if cloudbrain.JobID != "" {
		err = deleteRemote(entity.JobIdAndVersionId{JobID: cloudbrain.JobID, VersionID: cloudbrain.VersionID, JobType: cloudbrain.JobType, ComputeResource: cloudbrain.ComputeResource, TaskID: cloudbrain.ID})
		if err != nil {
			log.Error("delete from remote err.%v", err)
			return err
		}
	}

	//删除本地记录
	err = models.DeleteJob(cloudbrain)
	//删除存储
	storageType := models.TypeCloudBrainOne
	if cloudbrain.ComputeResource == models.NPUResource {
		storageType = models.TypeCloudBrainTwo
	}
	go cloudbrainTask.DeleteCloudbrainJobStorage(cloudbrain.JobName, storageType)
	return nil
}

func StopTask(id int64, stopRemote StopFunc) error {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return errors.New(response.SYSTEM_ERROR.TrCode)
	}

	if cloudbrain.IsTerminal() {
		return nil
	}
	if cloudbrain.IsPreparing() || cloudbrain.IsCreating() {
		return nil
	}
	err = stopRemote(entity.JobIdAndVersionId{JobID: cloudbrain.JobID, VersionID: cloudbrain.VersionID, ComputeResource: cloudbrain.ComputeResource, TaskID: cloudbrain.ID})
	if err != nil {
		log.Error("stop from remote err.%v", err)
		if models.IsErrCannotStopSavingImageJob(err) {
			return err
		}
		return errors.New(response.STOP_FAILED.TrCode)
	}

	return nil
}

// jobId string, baseLine int64, lines int64, order int64
func QueryTaskLog(opts entity.QueryLogOpts, getLogRemote GetLogFunc) (*entity.ClusterLog, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		return nil, err
	}
	if cloudbrain.JobID == "" {
		return &entity.ClusterLog{
			Content: "",
		}, nil
	}
	aiConfig := GetDetailConfigInfoByCloudbrain(cloudbrain)
	return getLogRemote(entity.ClusterLogOpts{
		JobId:           cloudbrain.JobID,
		JobType:         cloudbrain.JobType,
		JobName:         cloudbrain.JobName,
		BaseLine:        opts.BaseLine,
		Lines:           opts.Lines,
		Direction:       opts.Order,
		ObjectKeyPrefix: aiConfig.OutputObjectPrefix,
		StorageType:     aiConfig.OutputStorageType,
		VersionID:       cloudbrain.VersionID,
		NodeId:          opts.NodeId,
		LogFileName:     opts.LogFileName,
		WorkServerNum:   cloudbrain.WorkServerNumber,
	})
}

func GetLogDownloadInfo(opts entity.GetLogDownloadInfoReq, getLogDownloadInfo GetLogDownloadInfoFunc) (*entity.FileDownloadInfo, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		return nil, err
	}
	if cloudbrain.JobID == "" {
		return nil, nil
	}
	aiConfig := GetDetailConfigInfoByCloudbrain(cloudbrain)
	return getLogDownloadInfo(entity.ClusterLogDownloadInfoOpts{
		JobType:         cloudbrain.JobType,
		JobId:           cloudbrain.JobID,
		ObjectKeyPrefix: aiConfig.OutputObjectPrefix,
		StorageType:     aiConfig.OutputStorageType,
		NodeId:          opts.NodeId,
		LogFileName:     opts.LogFileName,
		WorkServerNum:   cloudbrain.WorkServerNumber,
		JobName:         cloudbrain.JobName,
		DisplayJobName:  cloudbrain.DisplayJobName,
	})
}

func GetSingleOutputDownloadInfo(opts entity.GetSingleDownloadInfoReq, f GetSingleOutputDownloadInfoFunc) (*entity.FileDownloadInfo, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		return nil, err
	}
	if cloudbrain.JobID == "" {
		return nil, nil
	}
	aiConfig := GetDetailConfigInfoByCloudbrain(cloudbrain)
	fileRelativePath := path.Join(aiConfig.OutputObjectPrefix, opts.ParentDir, opts.FileName)
	return f(entity.ClusterSingleOutputDownloadInfoOpts{
		JobId:       cloudbrain.JobID,
		Path:        fileRelativePath,
		StorageType: aiConfig.OutputStorageType,
	})
}

func DownloadAllOutput(opts entity.DownloadAllFileReq, downloadFunc DownloadAllOutputFunc) error {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		return err
	}
	if cloudbrain.JobID == "" {
		return nil
	}
	aiConfig := GetDetailConfigInfoByCloudbrain(cloudbrain)
	return downloadFunc(entity.DownloadOutputOpts{
		JobId:       cloudbrain.JobID,
		Path:        aiConfig.OutputObjectPrefix,
		StorageType: aiConfig.OutputStorageType,
		JobName:     cloudbrain.JobName,
		ZIPWriter:   opts.ZIPWriter,
	})
}

func QueryNoteBookUrl(id int64, getNoteBookUrl GetNotebookUrlFunc, fileName string) (string, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return "", err
	}
	if cloudbrain.JobID == "" {
		return "", errors.New("JobID is empty")
	}
	url, err := getNoteBookUrl(cloudbrain.JobID)
	if err != nil {
		return "", err
	}
	if url == "" {
		return "", nil
	}
	if fileName != "" {
		url = transferFileNotebookUrl(url, fileName)
	}
	//针对在线运行notebook的特殊处理，触发一下上传文件
	if cloudbrain.IsFileNoteBookTask() {
		go cloudbrainTask.UploadNotebookFiles(cloudbrain)
	}
	return url, nil
}

func GetAITaskNodeInfo(id int64, getNodeInfo GetNodeInfoFunc) ([]entity.AITaskNodeInfo, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return nil, err
	}

	res, err := getNodeInfo(entity.ClusterNodeInfoOpts{
		JobId:         cloudbrain.JobID,
		WorkServerNum: cloudbrain.WorkServerNumber,
		VersionId:     cloudbrain.VersionID,
	})
	if err != nil {
		log.Error("getNodeInfo error.id = %d err=%v", id, err)
		return nil, err
	}
	return res, nil
}
func GetAITaskOutput(id int64, parentDir string, getOutput GetOutputFunc) (*entity.AITaskOutput, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return nil, err
	}
	aiConfig := GetDetailConfigInfoByCloudbrain(cloudbrain)
	res, err := getOutput(entity.ClusterOutputOpts{
		JobId:           cloudbrain.JobID,
		ObjectKeyPrefix: aiConfig.OutputObjectPrefix,
		StorageType:     aiConfig.OutputStorageType,
		ParentDir:       parentDir,
	})
	if err != nil {
		log.Error("GetAITaskOutput getOutput from cluster error.id=%d parentDir=%s err=%v ", id, parentDir, err)
		return nil, err
	}
	return &entity.AITaskOutput{
		Status:         res.Status,
		Path:           res.Path,
		FileList:       res.FileList,
		IsTaskTerminal: cloudbrain.IsTerminal(),
	}, nil
}

func GetAllAITaskOutput(opts entity.GetAllOutputReq, getOutput GetAllOutputFunc) (*entity.AllAITaskOutput, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		return nil, err
	}
	aiConfig := GetDetailConfigInfoByCloudbrain(cloudbrain)
	res, err := getOutput(entity.ClusterOutputOpts{
		JobId:           cloudbrain.JobID,
		ObjectKeyPrefix: aiConfig.OutputObjectPrefix,
		StorageType:     aiConfig.OutputStorageType,
	})
	if err != nil {
		log.Error("GetAllAITaskOutput getOutput from cluster error.id=%d  err=%v ", opts.CloudbrainId, err)
		return nil, err
	}

	return filterOutputFile(res, opts.Suffix), nil
}

func filterOutputFile(sourceOutput *entity.AllAITaskOutput, suffixArray []string) *entity.AllAITaskOutput {
	if len(suffixArray) == 0 || sourceOutput == nil || len(sourceOutput.FileList) == 0 {
		return sourceOutput
	}
	sourceFiles := sourceOutput.FileList
	var files = make([]storage.FileInfo, 0)
	for i := 0; i < len(sourceFiles); i++ {
		f := sourceFiles[i]
		for j := 0; j < len(suffixArray); j++ {
			if strings.HasSuffix(f.FileName, suffixArray[j]) {
				files = append(files, f)
				break
			}
		}

	}
	return &entity.AllAITaskOutput{FileList: files}
}

func transferFileNotebookUrl(oldUrl string, fileName string) string {
	u, err := url.Parse(oldUrl)
	if err != nil {
		return oldUrl
	}
	p := u.Path
	u.Path = GetFileNoteBookDebugUrl(p, fileName)
	return u.String()
}

func GetFileNoteBookDebugUrl(url string, filename string) string {
	middle := ""
	if url[len(url)-3:] == "lab" || url[len(url)-4:] == "lab/" {
		if url[len(url)-1] == '/' {
			middle = "tree/"
		} else {
			middle = "/tree/"
		}
	} else {
		if url[len(url)-1] == '/' {
			middle = "lab/tree/"
		} else {
			middle = "/lab/tree/"
		}
	}

	return url + middle + filename
}

func GetResourceUsage(opts entity.GetResourceUsageOpts, fun GetResourceUsageFunc) (*entity.ResourceUsage, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(opts.CloudbrainId)
	if err != nil {
		return nil, err
	}
	if cloudbrain.JobID == "" {
		return &entity.ResourceUsage{}, nil
	}
	return fun(entity.ClusterResourceUsageOpts{
		JobId:            cloudbrain.JobID,
		StartTime:        int64(cloudbrain.StartTime),
		EndTime:          int64(cloudbrain.EndTime),
		NodeId:           opts.NodeId,
		ComputeSource:    models.GetComputeSourceStandardFormat(cloudbrain.ComputeResource),
		WorkServerNumber: cloudbrain.WorkServerNumber,
		VersionID:        cloudbrain.VersionID,
		LogFileName:      opts.LogFileName,
	})
}

func CreateAITask(form entity.CreateReq, gitRepo *git.Repository, repo *models.Repository, user *models.User, traceContext *context.Context) (*entity.CreateTaskRes, *response.BizError) {
	t, err := GetAITaskTemplate(form.JobType, form.Cluster)
	if err != nil {
		log.Error("param error")
		return nil, err
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: form.DisplayJobName, JobType: string(form.JobType)}, User: user})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()
	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg)
		return nil, response.BuildDefaultBizError(errMsg, errMsg)
	}
	return t.Create(&creation_context.CreationContext{
		Request:    &form,
		GitRepo:    gitRepo,
		Repository: repo,
		User:       user,
		Config: t.GetConfig(entity.AITaskConfigKey{ComputeSource: form.ComputeSourceStr,
			IsFileNoteBookRequest: form.IsFileNoteBookRequest}),
		TraceContext: traceContext,
	})
}

func RestartAITask(cloudbrainId int64, gitRepo *git.Repository, repo *models.Repository, user *models.User) (*entity.CreateTaskRes, *response.BizError) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		log.Error("RestartAITask GetCloudbrainByJobID err.%v", err)
		return nil, response.AI_TASK_NOT_EXISTS
	}
	t, bizErr := GetAITaskTemplateFromCloudbrain(cloudbrain)
	if bizErr != nil {
		log.Error("param error")
		return nil, response.PARAM_ERROR
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainRestart(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: cloudbrain.DisplayJobName, JobType: cloudbrain.JobType}, User: user})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()
	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg)
		return nil, response.BuildDefaultBizError(errMsg, errMsg)
	}
	return t.Restart(&creation_context.CreationContext{
		User:             user,
		SourceCloudbrain: cloudbrain,
		Config:           t.GetConfig(entity.AITaskConfigKey{ComputeSource: models.GetComputeSourceStandardFormat(cloudbrain.ComputeResource)}),
	})
}

func GetOperationProfile(id int64, getOperationProfile GetOperationProfileFunc) (*entity.OperationProfile, error) {
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		return nil, err
	}

	defaultRes := &entity.OperationProfile{Events: []entity.ProfileEvent{}}
	errMsg := cloudbrain.FailedReason
	if errMsg != "" {
		defaultRes = &entity.OperationProfile{Events: []entity.ProfileEvent{{Reason: "Error", Message: errMsg}}}
	}
	if cloudbrain.JobID == "" {
		return defaultRes, nil
	}
	s, err := getOperationProfile(entity.JobIdAndVersionId{JobID: cloudbrain.JobID, TaskID: cloudbrain.ID})
	if err != nil || s == nil {
		return defaultRes, nil
	}
	return s, err
}

func HandleNoJobIdAITasks() {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()

	cloudBrains, err := models.GetActiveStopCloudBrainJob()
	if err != nil {
		log.Error("GetLocalCreatingCloudBrainJob failed:", err.Error())
		return
	}

	for _, task := range cloudBrains {
		UpdateCloudbrain(task)
	}
}

func UpdateCloudbrain(task *models.Cloudbrain, res ...*entity.QueryTaskResponse) (*models.Cloudbrain, *response.BizError) {
	log.Info("use new ai task update method.task.DisplayJobName= %s", task.DisplayJobName)
	t, bizErr := GetAITaskTemplateFromCloudbrain(task)
	if bizErr != nil {
		log.Error("GetAITaskFromCloudbrain err.%v", bizErr)
		return task, bizErr
	}
	bizErr = t.Update(task.ID, res...)
	if bizErr != nil {
		return task, bizErr
	}
	newTask, err := models.GetCloudbrainByCloudbrainID(task.ID)
	if err != nil {
		return task, response.NewBizError(err)
	}
	return newTask, nil
}
func StopCloudbrain(task *models.Cloudbrain) (*entity.AITaskBriefInfo, *response.BizError) {
	log.Info("use new ai task stop method.task.DisplayJobName= %s", task.DisplayJobName)
	t, err := GetAITaskTemplateFromCloudbrain(task)
	if err != nil {
		log.Error("GetAITaskFromCloudbrain err.%v", err)
		return nil, err
	}
	return t.Stop(task.ID)
}
func DelCloudbrain(task *models.Cloudbrain) *response.BizError {
	log.Info("use new ai task delete method.task.DisplayJobName= %s", task.DisplayJobName)
	t, err := GetAITaskTemplateFromCloudbrain(task)
	if err != nil {
		log.Error("GetAITaskFromCloudbrain err.%v", err)
		return err
	}
	return t.Delete(task.ID)
}

func DelCloudbrains(tasks []*models.Cloudbrain) *response.BizError {
	for _, t := range tasks {
		err := DelCloudbrain(t)
		if err != nil {
			log.Error("delete cloudbrain err.id=%d err=%v", t.ID, err)
		}
	}
	return nil
}

func HandleNewAITaskStop(cloudbrainId int64) (result *entity.AITaskBriefInfo, isHandled bool, err error) {
	task, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		return nil, false, err
	}
	if !task.IsNewAITask() {
		return nil, false, nil
	}
	result, bizErr := StopCloudbrain(task)
	if bizErr != nil {
		return nil, true, errors.New(bizErr.DefaultMsg)
	}
	return result, true, nil
}
func HandleNewAITaskDelete(cloudbrainId int64) (isHandled bool, err error) {
	task, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err != nil {
		return false, err
	}
	if !task.IsNewAITask() {
		return false, nil
	}
	bizErr := DelCloudbrain(task)
	if bizErr != nil {
		return true, errors.New(bizErr.TrCode)
	}
	return true, nil
}

func ClearNotebook() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
		}
	}()

	if !setting.NotebookStrategy.ClearEnabled {
		return
	}

	userCountInfo, err := models.GetNotebookCountGreaterThanN(setting.NotebookStrategy.MaxNumberPerUser)
	if err != nil {
		log.Error("can not get Notebook user count info", err)
		return
	}
	deleteCount := 0
	for _, userCount := range userCountInfo {
		ids, err := models.GetNotebooksByUser(userCount.UserID, setting.NotebookStrategy.MaxNumberPerUser)
		if err != nil {
			log.Error("can not get Notebook by user id", err)
			continue
		}
		for _, id := range ids {
			t, _ := GetAITaskTemplateByCloudbrainId(id)
			if t == nil {
				log.Error("can not get task template")
				continue
			}
			err := t.Delete(id)
			if err != nil {
				log.Error("Delete error.%v", err)
				continue
			}
			log.Info("Clear Notebook id is " + strconv.FormatInt(id, 10))
			deleteCount += 1
			if deleteCount >= setting.NotebookStrategy.ClearBatchSize {
				return
			}
		}
	}

}

func TryToNotifyDebugAlmostEndingTask(cloudbrain *models.Cloudbrain, maxDuration int64) {

	if cloudbrain.JobType == string(models.JobTypeDebug) && cloudbrain.IsRunning() && maxDuration-cloudbrain.Duration <= 1800 {
		val, err := redis_client.Get(redis_key.DebugEndNotificationKey(cloudbrain))
		if val == "" || err != nil {
			redis_client.Setex(redis_key.DebugEndNotificationKey(cloudbrain), "1", 2700*time.Second)
			notification.NotifyAlmostEndingAIDebugTask(cloudbrain, time.Duration(maxDuration-cloudbrain.Duration)*time.Second)
		}

	}

}

func TryToStopFailedPullImage(cloudbrain *models.Cloudbrain) {
	if cloudbrain.Status == string(models.JobWaiting) {
		created_unix := int64(cloudbrain.CreatedUnix)
		currentTimestamp := time.Now().Unix()
		if currentTimestamp-created_unix > 10*60 {
			log.Info("The task has waiting 10 minutes, so get it brief.task name=" + cloudbrain.JobName)
			t, bizErr := GetAITaskTemplateFromCloudbrain(cloudbrain)
			if bizErr != nil {
				log.Error("param error")
				return
			}
			events, bizErr := t.GetOperationProfile(cloudbrain.ID)
			if bizErr != nil {
				log.Error("param error")
				return
			}
			if events.Events != nil {
				for _, event := range events.Events {
					if strings.Contains(event.Message, "Failed to pull image") || strings.Contains(event.Reason, "Failed to pull image") {
						//to stop the task
						log.Info("The task has failed to pull image,so stop it. task name=" + cloudbrain.JobName)
						t.Stop(cloudbrain.ID)
					}
					if strings.Contains(event.Message, "ErrImagePull") || strings.Contains(event.Reason, "ErrImagePull") {
						//to stop the task
						log.Info("The task has failed to pull image,so stop it. task name=" + cloudbrain.JobName)
						t.Stop(cloudbrain.ID)
					}
					if strings.Contains(event.Message, "InvalidImageName") || strings.Contains(event.Reason, "InvalidImageName") {
						//to stop the task
						log.Info("The task has failed to pull image,so stop it. task name=" + cloudbrain.JobName)
						t.Stop(cloudbrain.ID)
					}
				}
			}
		}

	}
}

func TryToNotifyLongRunningTask(cloudbrain *models.Cloudbrain) {
	startTimestamp := int64(cloudbrain.StartTime)
	if setting.CloudbrainLongRunningNotifyInterval == 0 || startTimestamp == 0 {
		return
	}
	currentTimestamp := time.Now().Unix()
	matchIndex := 0
	i := 1
	for {
		duration := time.Duration(i) * setting.CloudbrainLongRunningNotifyInterval
		s := int64(duration.Seconds())
		if currentTimestamp-startTimestamp < s {
			matchIndex = i - 1
			break
		}
		i++
	}
	if matchIndex <= 0 {
		return
	}
	d := time.Duration(matchIndex) * setting.CloudbrainLongRunningNotifyInterval
	if hasLongRunningNotificationSendBefore(cloudbrain, d) {
		return
	}
	notification.NotifyLongRunningAITask(cloudbrain, d)
}

func hasLongRunningNotificationSendBefore(cloudbrain *models.Cloudbrain, duration time.Duration) bool {
	key := redis_key.LongRunningNotificationKey(cloudbrain, duration)
	lock := redis_lock.NewDistributeLock(key)
	success, err := lock.Lock(duration + 1*time.Hour)
	if err != nil {
		return true
	}
	return !success
}

func GetUserMultiLimitNum(userId int64, jobType models.JobType) int {
	defaultLimit := 1
	if jobType == models.JobTypeFINETUNE {
		defaultLimit = setting.MAX_MULTI_FINE_TUNE_TASK
	}
	if jobType == models.JobTypeModelExperience {
		defaultLimit = setting.MODEL_EXPERIENCE.MAX_CREATES
	}
	if jobType == models.JobTypeSdFinetune {
		defaultLimit = setting.MAX_MULTI_SD_FINETUNE_TASK

	}

	nums := role.AllUserOperNum(userId, role.ROLE_OPER_MULTI_TASK, string(jobType))

	if len(nums) > 0 {
		sort.Ints(nums)
		re := nums[len(nums)-1]
		if re > 0 {
			if re > defaultLimit {
				return re
			} else {
				return defaultLimit
			}
		}

	}
	return defaultLimit

}

func CopyModelAndLoadModel(cloudbrain *models.Cloudbrain) {

}
