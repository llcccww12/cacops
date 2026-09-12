package entity

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
)

type CreateNoteBookTaskRequest struct {
	Name                 string
	Description          string
	Tasks                []NoteBookTask
	PrimitiveDatasetName string
	RepoName             string
	IsSubscriber         bool
	User                 *models.User
	SourceCloudbrain     *models.Cloudbrain
}

type NoteBookTask struct {
	AutoStopDuration int64
	Name             string
	Capacity         int
	Queues           []models.ResourceQueue
	Code             []ContainerData
	Datasets         []ContainerData
	PreTrainModel    []ContainerData
	OutPut           []ContainerData
	ImageId          string
	ImageUrl         string
	ResourceSpecId   string
	BootFile         string
	ExperienceType   string
	Spec             *models.Specification
	EnvVariables     map[string]interface{}
	EndPoint         *SelfEndPoint
	AutoSave         bool
}

type SelfEndPoint struct {
	Port     int64
	EndPoint string
}

type CreateNoteBookTaskResponse struct {
	StartedAt   int64
	RunSec      int64
	CompletedAt int64
	CreatedAt   int64
	UpdatedAt   int64
	Desc        string
	JobID       string
	Name        string
	Status      string
	UserID      string
}
type CreateGeneralTaskRequest struct {
	Name                 string
	Description          string
	Tasks                []GeneralTask
	PrimitiveDatasetName string
	RepoName             string
}

type GeneralTask struct {
	AutoStopDuration int64
	Name             string
	Capacity         int
	Queues           []models.ResourceQueue
	Code             []ContainerData
	Datasets         []ContainerData
	PreTrainModel    []ContainerData
	OutPut           []ContainerData
	ImageId          string
	ImageUrl         string
	ResourceSpecId   string
	BootFile         string
	Spec             *models.Specification
	EnvVariables     map[string]interface{}
}

type CreateGeneralTaskResponse struct {
	StartedAt   int64
	RunSec      int64
	CompletedAt int64
	CreatedAt   int64
	UpdatedAt   int64
	Desc        string
	JobID       string
	Name        string
	Status      string
	UserID      string
}

type RestartNoteBookTaskResponse struct {
	JobId  string `json:"newId"`
	Status string `json:"status"`
}

type CreateTrainTaskRequest struct {
	Name           string
	DisplayJobName string
	Description    string
	TaskConfig     *AITaskBaseConfig
	Tasks          []TrainTask
}

type QueryTaskResponse struct {
	StartedAt           timeutil.TimeStamp `json:"started_at"`
	CompletedAt         timeutil.TimeStamp `json:"completed_at"`
	JobId               string             `json:"job_id"`
	Status              string             `json:"status"`
	DetailedStatus      string             `json:"detailed_status"`
	Url                 string             `json:"url"`
	Token               string             `json:"token"`
	PToken              string             `json:"ptoken"`
	CenterId            string             `json:"center_id"`
	CenterName          string             `json:"center_name"`
	QueueCode           string             `json:"queue_code"`
	CodeUrl             string             `json:"code_url"`
	DataUrl             string             `json:"data_url"`
	ContainerIP         string             `json:"container_ip"`
	ContainerID         string             `json:"container_id"`
	VersionId           int64              `json:"version_id"`
	TensorboardEndpoint string             `json:"tensorboardEndpoint"`
	DomainUrl           string             `json:"domain_url"`
}

func ConvertGrampusNotebookResponse(job models.GrampusNotebookInfo) *QueryTaskResponse {
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
	var url, token, codeUrl, dataUrl, ptoken, domainUrl string

	if len(job.Tasks) > 0 {
		t := job.Tasks[0]
		url = t.Url
		token = t.Token
		codeUrl = t.CodeUrl
		dataUrl = t.DataUrl
		ptoken = t.PToken
		if len(t.DominUrl) > 0 {
			domainUrl = t.DominUrl[0]
		}
	}
	return &QueryTaskResponse{
		StartedAt:      timeutil.TimeStamp(job.StartedAt),
		CompletedAt:    timeutil.TimeStamp(job.CompletedAt),
		Status:         job.Status,
		DetailedStatus: job.DetailedStatus,
		CenterId:       centerId,
		CenterName:     centerName,
		QueueCode:      task.PoolId,
		Url:            url,
		Token:          token,
		JobId:          job.JobID,
		CodeUrl:        codeUrl,
		DataUrl:        dataUrl,
		PToken:         ptoken,
		DomainUrl:      domainUrl,
	}
}
func ConvertGrampusTrainResponse(job models.GrampusJobInfo) *QueryTaskResponse {
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
	return &QueryTaskResponse{
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

func ConvertCloudbrainOneQueryNotebookByNameResponse(result models.JobResultInListPayload) *QueryTaskResponse {
	if result.State == "" {
		return nil
	}
	return &QueryTaskResponse{
		StartedAt:   timeutil.TimeStamp(result.CreatedTime / 1000),
		CompletedAt: timeutil.TimeStamp(result.CompletedTime / 1000),
		Status:      result.State,
		JobId:       result.Id,
	}
}

func ConvertCloudbrainOneNotebookResponse(input map[string]interface{}) (*QueryTaskResponse, error) {
	data, _ := json.Marshal(input)
	var jobResultPayload models.JobResultPayload
	err := json.Unmarshal(data, &jobResultPayload)
	if err != nil {
		log.Error("parse cloudbrain one result err,result=%+v err=%v", input, err)
		return nil, err
	}
	if jobResultPayload.JobStatus.State == "" {
		return nil, nil
	}

	startTime := jobResultPayload.JobStatus.AppLaunchedTime / 1000
	var endTime int64
	switch jobResultPayload.JobStatus.AppCompletedTime.(type) {
	case float64:
		f := jobResultPayload.JobStatus.AppCompletedTime.(float64)
		s := fmt.Sprintf("%.0f", f)
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			endTime = i / 1000
		}
	}

	if jobResultPayload.JobStatus.State == string(models.JobWaiting) {
		startTime = 0
		endTime = 0
	}
	var containerIP, containerID string
	taskRoles := jobResultPayload.TaskRoles
	if taskRoles != nil && len(taskRoles) > 0 {
		subTask := taskRoles[cloudbrain.SubTaskName]
		if subTask != nil {
			taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
			if taskRes.TaskStatuses != nil && len(taskRes.TaskStatuses) > 0 {
				containerIP = taskRes.TaskStatuses[0].ContainerIP
				containerID = taskRes.TaskStatuses[0].ContainerID
			}
		}
	}

	res := &QueryTaskResponse{
		StartedAt:   timeutil.TimeStamp(startTime),
		CompletedAt: timeutil.TimeStamp(endTime),
		Status:      jobResultPayload.JobStatus.State,
		JobId:       jobResultPayload.ID,
		ContainerIP: containerIP,
		ContainerID: containerID,
	}
	return res, nil
}

type ClusterLog struct {
	Content        string `json:"content"`
	CanLogDownload bool   `json:"can_log_download"`
	//云脑二返回的startline和baseline在前端会丢失精度。因此改为string类型
	StartLine string `json:"start_line"`
	EndLine   string `json:"end_line"`
	Lines     int64  `json:"lines"`
}

type TrainTask struct {
	Command          string                 `json:"command"`
	Name             string                 `json:"name"`
	ImageId          string                 `json:"imageId"`
	ImageUrl         string                 `json:"imageUrl"`
	ResourceSpecId   string                 `json:"resourceSpecId"`
	Queues           []models.ResourceQueue `json:"centerID"`
	ReplicaNum       int                    `json:"replicaNum"`
	Datasets         []ContainerData        `json:"datasets"`
	PreTrainModel    []ContainerData        `json:"models"`
	Code             []ContainerData        `json:"code"`
	BootFile         string                 `json:"bootFile"`
	OutPut           []ContainerData        `json:"output"`
	LogPath          []ContainerData        `json:"logPath"`
	PoolId           string                 `json:"poolId"`
	EnvVariables     map[string]interface{} `json:"envVariables"`
	Params           models.Parameters
	Spec             *models.Specification
	RepoName         string
	WorkServerNumber int
	Method           string
}

type CreateTrainTaskResponse struct {
	StartedAt   int64  `json:"startedAt"`
	RunSec      int64  `json:"runSec"`
	CompletedAt int64  `json:"completedAt"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	Desc        string `json:"desc"`
	JobID       string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	UserID      string `json:"userId"`
	VersionID   int64  `json:"versionID"`
	VersionName string `json:"versionName"` //当前版本
}

type ClusterType string

const (
	OpenICloudbrainOne ClusterType = "OpenICloudbrainOne"
	OpenICloudbrainTwo ClusterType = "OpenICloudbrainTwo"
	C2Net              ClusterType = "C2Net"
	IFLYTEKTraining    ClusterType = "IFLYTEKTraining"
)

func (t ClusterType) GetParentCluster() string {
	switch t {
	case OpenICloudbrainTwo, OpenICloudbrainOne:
		return models.OpenICluster
	case C2Net:
		return models.C2NetCluster
	case IFLYTEKTraining:
		return models.IFLYTEKCluster
	}
	return ""
}

func (t ClusterType) GetCloudbrainType() int {
	switch t {
	case OpenICloudbrainOne:
		return models.TypeCloudBrainOne
	case OpenICloudbrainTwo:
		return models.TypeCloudBrainTwo
	case C2Net:
		return models.TypeC2Net
	case IFLYTEKTraining:
		return models.TypeIFLYTek
	}
	return -1
}

func GetClusterTypeFromCloudbrainType(t int) ClusterType {
	switch t {
	case models.TypeCloudBrainOne:
		return OpenICloudbrainOne
	case models.TypeCloudBrainTwo:
		return OpenICloudbrainTwo
	case models.TypeC2Net:
		return C2Net
	case models.TypeCDCenter:
		return OpenICloudbrainTwo
	case models.TypeIFLYTek:
		return IFLYTEKTraining
	}
	return ""
}

type ClusterLogOpts struct {
	JobId           string
	JobType         string
	JobName         string
	BaseLine        int64
	Lines           int64
	Direction       Direction
	ObjectKeyPrefix string
	StorageType     StorageType
	VersionID       int64
	NodeId          int
	LogFileName     string
	WorkServerNum   int
}

func (opts ClusterLogOpts) IsBottomRequest() bool {
	return opts.BaseLine == 0 && opts.Direction == UP
}
func (opts ClusterLogOpts) IsHeadRequest() bool {
	return opts.BaseLine == 0 && opts.Direction == DOWN
}

type ClusterLogDownloadInfoOpts struct {
	JobId           string
	JobType         string
	ObjectKeyPrefix string
	StorageType     StorageType
	NodeId          int
	LogFileName     string
	WorkServerNum   int
	JobName         string
	DisplayJobName  string
}

type DownloadOutputOpts struct {
	JobId       string
	Path        string
	JobName     string
	StorageType StorageType
	ZIPWriter   *zip.Writer
}
type ClusterSingleOutputDownloadInfoOpts struct {
	JobId       string
	Path        string
	JobName     string
	StorageType StorageType
}
type ClusterNodeInfoOpts struct {
	JobId         string
	WorkServerNum int
	VersionId     int64
}
type ClusterResourceUsageOpts struct {
	JobId            string
	NodeId           int
	LogFileName      string
	VersionID        int64
	StartTime        int64
	EndTime          int64
	ComputeSource    string
	WorkServerNumber int
}

type ClusterOutputOpts struct {
	JobId           string
	ObjectKeyPrefix string
	StorageType     StorageType
	ParentDir       string
}

type ClusterAITaskOutput struct {
	Status          models.ModelMigrateStatus
	Path            string
	FileList        []storage.FileInfo
	OutputSizeLimit int `json:"output_size_limit"`
}
