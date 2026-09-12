package cluster

import (
	"errors"

	"code.gitea.io/gitea/models"

	"code.gitea.io/gitea/entity"
)

var clusterMap = map[entity.ClusterType]ClusterAdapter{}

func AddCluster(t entity.ClusterType, cluster ClusterAdapter) {
	clusterMap[t] = cluster
}

func GetCluster(t entity.ClusterType) (ClusterAdapter, error) {
	if t == "" {
		return nil, errors.New("ClusterType is empty")
	}
	if c, ok := clusterMap[t]; !ok {
		return nil, errors.New("Cluster not exists")
	} else {
		return c, nil
	}
}

type ClusterAdapter interface {
	CreateNoteBook(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error)
	RestartNoteBook(jobId string, autoStopDuration int64, trace *entity.TraceInfo) (*entity.RestartNoteBookTaskResponse, error)
	DeleteNoteBook(opts entity.JobIdAndVersionId) error
	StopNoteBook(opts entity.JobIdAndVersionId) error
	QueryNoteBook(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error)
	QueryNoteBookByJobName(jobName string) ([]*entity.QueryTaskResponse, error)
	GetNoteBookLog(jobId string) (*entity.ClusterLog, error)
	GetNoteBookUrl(jobId string) (string, error)
	GetSelfEndPointUrl(jobId string) (string, error)
	GetNoteBookOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error)
	CreateTrainJob(req entity.CreateTrainTaskRequest, trace *entity.TraceInfo) (*entity.CreateTrainTaskResponse, error)
	DeleteTrainJob(opts entity.JobIdAndVersionId) error
	StopTrainJob(opts entity.JobIdAndVersionId) error
	QueryTrainJob(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error)
	QueryTrainJobByJobName(jobName string) ([]*entity.QueryTaskResponse, error)
	GetVisualizeUrl(jobId string) (string, error)
	GetLog(opts entity.ClusterLogOpts) (*entity.ClusterLog, error)
	GetLogDownloadInfo(entity.ClusterLogDownloadInfoOpts) (*entity.FileDownloadInfo, error)
	GetTrainJobOperationProfile(opts entity.JobIdAndVersionId) (*entity.OperationProfile, error)
	GetOutput(opts entity.ClusterOutputOpts) (*entity.ClusterAITaskOutput, error)
	GetAllOutput(opts entity.ClusterOutputOpts) (*entity.AllAITaskOutput, error)
	GetSingleOutputDownloadInfo(opts entity.ClusterSingleOutputDownloadInfoOpts) (*entity.FileDownloadInfo, error)
	DownloadAllOutput(opts entity.DownloadOutputOpts) error
	GetNodeInfo(opts entity.ClusterNodeInfoOpts) ([]entity.AITaskNodeInfo, error)
	GetResourceUsage(zopts entity.ClusterResourceUsageOpts) (*entity.ResourceUsage, error)
	//GetImages return  available list of clusters
	//The second parameter will return true if image is no limit
	GetNotebookImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error)
	GetTrainImages(req entity.GetImageReq, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, error)
	CreateOnlineInfer(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error)
	CreateModelExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error)
	QueryOnlineInfer(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error)
	QueryModelExperience(opts entity.JobIdAndVersionId) (*entity.QueryTaskResponse, error)
	StopModelExperience(opts entity.JobIdAndVersionId) error
	DeleteModelExperience(opts entity.JobIdAndVersionId) error
	CreateGeneralTask(req entity.CreateGeneralTaskRequest, trace *entity.TraceInfo) (*entity.CreateGeneralTaskResponse, error)
	CreateSdFinetune(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error)
	CreateComfyuiExperience(req entity.CreateNoteBookTaskRequest, trace *entity.TraceInfo) (*entity.CreateNoteBookTaskResponse, error)
}
