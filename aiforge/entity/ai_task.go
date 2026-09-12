package entity

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"reflect"
	"strings"
	"sync"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/models"

	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"
)

// todo 暂时保留之前各种云脑属性的定义
type CreateReq struct {
	JobType               models.JobType           `json:"job_type" binding:"Required"`
	DisplayJobName        string                   `json:"display_job_name" binding:"Required"`
	JobName               string                   `json:"job_name"`
	SpecId                int64                    `json:"spec_id" binding:"Required"`
	ComputeSourceStr      string                   `json:"compute_source" binding:"Required"`
	Cluster               ClusterType              `json:"cluster" binding:"Required"`
	WorkServerNumber      int                      `json:"work_server_number"`
	BranchName            string                   `json:"branch_name"`
	ImageUrl              string                   `json:"image_url"`
	ImageID               string                   `json:"image_id"`
	ImageName             string                   `json:"image_name"`
	PretrainModelName     string                   `json:"pretrain_model_name"`
	PretrainModelId       string                   `json:"pretrain_model_id_str"`
	Description           string                   `json:"description"`
	LabelName             string                   `json:"label_names"`
	DatasetUUIDStr        string                   `json:"dataset_uuid_str"`
	Params                string                   `json:"params"`
	BootFile              string                   `json:"boot_file"`
	PoolId                string                   `json:"pool_id"`
	IsContinueRequest     bool                     `json:"is_continue"`
	SourceCloudbrainId    int64                    `json:"source_cloudbrain_id"`
	AppName               string                   `json:"app_name"`
	HasInternet           models.SpecInternetQuery `json:"has_internet"` //0 all;1 no internet;2 has internet
	TimeLimit             int                      `json:"time_limit" binding:"Range(-1,24)"`
	VisualizeRequired     bool                     `json:"visualize_required"`
	ParamArray            models.Parameters
	ComputeSource         *models.ComputeSource
	ReqCommitID           string
	IsFileNoteBookRequest bool
	FileRepository        *models.Repository
	FileBranchName        string
	IsRestartRequest      bool
	DatasetNames          string
	DatasetAlias          string
	ModelNames            string
	Port                  int    `json:"port"`
	EndPoint              string `json:"endPoint"`
	TemplateID            string `json:"template_id"`
	AutoSave              bool   `json:"auto_save"` //自动保存镜像
}

type CreationResponse struct {
	Error       error
	JobID       string
	Status      string //todo 考虑统一状态
	CreateTime  timeutil.TimeStamp
	VersionID   int64
	VersionName string
}
type DeleteIDs struct {
	ID []int64 `json:"id"`
}
type JobIdAndVersionId struct {
	JobID           string
	VersionID       int64
	JobType         string
	ComputeResource string
	TaskID          int64
}

type QueryAITaskRes struct {
	Task                 *AITaskDetailInfo   `json:"task"`
	EarlyVersionList     []*AITaskDetailInfo `json:"early_version_list"`
	CanCreateVersion     bool                `json:"can_create_version"`
	CanDownload          bool                `json:"can_download"`
	CanModify            bool                `json:"can_modify"`
	CanDelete            bool                `json:"can_delete"`
	CanCreateTemplate    bool                `json:"can_create_template"`
	CanExperience        bool                `json:"can_experience"`
	CanFintuneExperience bool                `json:"can_fintune_experience"`
	IsSubscriber         bool                `json:"is_subscriber"`
}

func (r *QueryAITaskRes) TryToRemoveNoNeedInfo(currentUser *models.User) {
	if r.Task != nil {
		r.Task.TryToRemoveDatasets(currentUser)
		r.Task.TryToRemovePretrainModelList(currentUser)
		r.Task.TryToRemoveSDKCode(currentUser)
		r.Task.TryToRemoveEndPoint(currentUser)
	}
	if r.EarlyVersionList != nil {
		for _, t := range r.EarlyVersionList {
			t.TryToRemoveDatasets(currentUser)
			t.TryToRemovePretrainModelList(currentUser)
			t.TryToRemoveSDKCode(currentUser)
		}
	}
}
func (r *QueryAITaskRes) Tr(language string) {
	if r.Task != nil {
		r.Task.Tr(language)
	}
	if r.EarlyVersionList != nil {
		for _, t := range r.EarlyVersionList {
			t.Tr(language)
		}
	}
}

type AITaskDetailInfo struct {
	ID                 int64                             `json:"id"`
	JobID              string                            `json:"job_id"`
	Status             string                            `json:"status"`
	DetailedStatus     string                            `json:"detailed_status"`
	JobType            string                            `json:"job_type"`
	Cluster            string                            `json:"cluster"`
	DisplayJobName     string                            `json:"display_job_name"`
	FormattedDuration  string                            `json:"formatted_duration"`
	ComputeSource      string                            `json:"compute_source"`
	AICenter           string                            `json:"ai_center"`
	BootFile           string                            `json:"boot_file"`
	PreVersionName     string                            `json:"pre_version_name"`
	CurrentVersionName string                            `json:"current_version_name"`
	WorkServerNumber   int                               `json:"work_server_number"`
	Spec               *structs.SpecificationShow        `json:"spec"`
	DatasetList        []*models.DatasetRegistryDownload `json:"dataset_list"`
	PretrainModelList  []*models.Model4Show              `json:"pretrain_model_list"`
	SDKCode            string                            `json:"sdk_code"`
	Parameters         *models.Parameters                `json:"parameters"`
	CreatedUnix        timeutil.TimeStamp                `json:"created_unix"`
	CodePath           string                            `json:"code_path"`
	DatasetPath        string                            `json:"dataset_path"`
	PretrainModelPath  string                            `json:"pretrain_model_path"`
	PretrainModelUrl   string                            `json:"pretrain_model_url"`
	OutputPath         string                            `json:"output_path"`
	CodeUrl            string                            `json:"code_url"`
	PretrainModelName  string                            `json:"pretrain_model_name"`
	PretrainModelId    string                            `json:"pretrain_model_id"`
	StartTime          timeutil.TimeStamp                `json:"start_time"`
	EndTime            timeutil.TimeStamp                `json:"end_time"`
	Description        string                            `json:"description"`
	CommitID           string                            `json:"commit_id"`
	BranchName         string                            `json:"branch_name"`
	ImageUrl           string                            `json:"image_url"`
	ImageID            string                            `json:"image_id"`
	ImageName          string                            `json:"image_name"`
	CreatorName        string                            `json:"creator_name"`
	EngineName         string                            `json:"engine_name"`
	FailedReason       string                            `json:"failed_reason"`
	UserId             int64                             `json:"-"`
	AppName            string                            `json:"app_name"`
	HasInternet        int                               `json:"has_internet"`
	TimeLimit          int                               `json:"time_limit"`
	DefaultTimeLimit   int                               `json:"default_time_limit"`
	Port               int                               `json:"port"`
	EndPoint           string                            `json:"end_point"`
	VisualizeRequired  bool                              `json:"visualize_required"`
	AimRequired        bool                              `json:"aim_required"`
	IsFileNotebook     bool                              `json:"is_file_notebook"`
	RepoID             int64                             `json:"repo_id"`
	RepoName           string                            `json:"repo_name"`
	RepoOwnerName      string                            `json:"repo_owner_name"`
	RepoAlias          string                            `json:"repo_alias"`
	SourceID           int64                             `json:"source_id"`
	CodeSizeLimit      int                               `json:"code_size_limit"`
	OutputSizeLimit    int                               `json:"output_size_limit"`
}

func (a *AITaskDetailInfo) Tr(language string) {
	aiCenterInfo := strings.Split(a.AICenter, "+")
	aiCenterCode := aiCenterInfo[0]
	aiCenterName := ""
	if len(aiCenterCode) >= 2 {
		aiCenterName = aiCenterInfo[1]
	}
	a.AICenter = models.GetAiCenterShow(aiCenterCode, aiCenterName, language)
}

func (a *AITaskDetailInfo) TryToRemoveEndPoint(currentUser *models.User) {
	//if currentUser == nil || a.UserId == 0 || (!currentUser.IsAdmin && currentUser.ID != a.UserId && !role.UserHasRole(currentUser.ID, models.MonitorAdmin)) {
	if currentUser == nil || a.UserId == 0 || (!currentUser.IsAdmin && currentUser.ID != a.UserId && !role.UserHasOper(currentUser.ID, role.ROLE_OPER_MonitorAdmin)) {
		a.EndPoint = ""
		a.Port = 0
	}
}

func (a *AITaskDetailInfo) TryToRemoveDatasets(currentUser *models.User) {
	//if currentUser == nil || a.UserId == 0 || (!currentUser.IsAdmin && currentUser.ID != a.UserId && !role.UserHasRole(currentUser.ID, models.MonitorAdmin)) {
	if currentUser == nil || a.UserId == 0 || (!currentUser.IsAdmin && currentUser.ID != a.UserId && !role.UserHasOper(currentUser.ID, role.ROLE_OPER_MonitorAdmin)) {
		a.DatasetList = []*models.DatasetRegistryDownload{}
	}
}
func (a *AITaskDetailInfo) TryToRemovePretrainModelList(currentUser *models.User) {
	if currentUser == nil || a.UserId == 0 || (!currentUser.IsAdmin && currentUser.ID != a.UserId) {
		a.PretrainModelList = []*models.Model4Show{}
	}
}
func (a *AITaskDetailInfo) TryToRemoveSDKCode(currentUser *models.User) {
	if currentUser == nil || a.UserId == 0 || (!currentUser.IsAdmin && currentUser.ID != a.UserId) {
		a.SDKCode = ""
	}
}

type CreateTaskRes struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type GetAITaskCreationInfoReq struct {
	User              *models.User
	JobType           models.JobType
	ClusterType       ClusterType
	ComputeSource     *models.ComputeSource
	Repo              *models.Repository
	GitRepo           *git.Repository
	IsOnlineType      bool
	VisualizeRequired bool
}

type GetAITaskCreationImageInfoReq struct {
	JobType           models.JobType
	ClusterType       ClusterType
	ComputeSource     *models.ComputeSource
	Spec              *models.Specification
	UserID            int64
	VisualizeRequired bool
	HasInternet       models.SpecInternetQuery
}

type AITaskBriefInfo struct {
	ID                int64              `json:"id"`
	JobType           string             `json:"job_type"`
	Status            string             `json:"status"`
	DetailedStatus    string             `json:"detailed_status"`
	DisplayJobName    string             `json:"display_job_name"`
	CreatedUnix       timeutil.TimeStamp `json:"created_unix"`
	StartTime         timeutil.TimeStamp `json:"start_time"`
	EndTime           timeutil.TimeStamp `json:"end_time"`
	FormattedDuration string             `json:"formatted_duration"`
	Cluster           string             `json:"cluster"`
	ComputeSource     string             `json:"compute_source"`
	AICenterCode      string             `json:"ai_center_code"`
	AICenter          string             `json:"ai_center"`
	IsFileNotebook    bool               `json:"is_file_notebook"`
	IsFineTuneTask    bool               `json:"is_fine_tune_task"`
	APPName           string             `json:"app_name"`
	AccCardType       string             `json:"acc_card_type"`
	JobName           string             `json:"job_name"`
	LabelName         string             `json:"label_name"`
	JobId             string             `json:"job_id"`
	EndPoint          string             `json:"end_point"`
	Port              int                `json:"port"`
	ImageId           string             `json:"image_id"`
	ImageName         string             `json:"image_name"`
}

func (a *AITaskBriefInfo) Tr(language string) {
	aiCenterInfo := strings.Split(a.AICenter, "+")
	aiCenterCode := aiCenterInfo[0]
	aiCenterName := ""
	if len(aiCenterInfo) >= 2 {
		aiCenterName = aiCenterInfo[1]
	}
	a.AICenter = models.GetAiCenterShow(aiCenterCode, aiCenterName, language)
}

func (a *AITaskBriefInfo) ClearNonPublicFields() *AITaskBriefInfo {
	a.JobName = ""
	return a
}

type AITaskListRes struct {
	Tasks         []*AITaskInfo4List `json:"tasks"`
	Total         int64              `json:"total"`
	PageSize      int                `json:"page_size"`
	Page          int                `json:"page"`
	CanCreateTask bool               `json:"can_create_task"`
	IsRepoEmpty   bool               `json:"is_repo_empty"`
	IsSubscriber  bool               `json:"is_subscriber"`
	IsAdmin       bool               `json:"is_admin"`
}
type AITaskInfo4List struct {
	Task                 *AITaskBriefInfo `json:"task"`
	Creator              UserBriefInfo    `json:"creator"`
	CanModify            bool             `json:"can_modify"`
	CanDelete            bool             `json:"can_delete"`
	CanCreateTemplate    bool             `json:"can_create_template"`
	CanExperience        bool             `json:"can_experience"`
	CanFintuneExperience bool             `json:"can_fintune_experience"`
	RepoID               int64            `json:"repo_id"`
	RepoName             string           `json:"repo_name"`
	RepoAlias            string           `json:"repo_alias"`
	OwnerName            string           `json:"owner_name"`
}

func ConvertCloudbrainToAITaskBriefInfo(task *models.Cloudbrain) *AITaskBriefInfo {
	accCardType, aiCenterCode := "", ""
	if task.Spec != nil {
		accCardType = task.Spec.AccCardType
		aiCenterCode = task.Spec.AiCenterCode
	}
	aiCenter := task.AiCenter
	switch task.Type {
	case models.TypeCloudBrainOne:
		aiCenter = models.AICenterOfCloudBrainOne
	case models.TypeCloudBrainTwo:
		aiCenter = models.AICenterOfCloudBrainTwo
	case models.TypeCDCenter:
		aiCenter = models.AICenterOfChengdu
	}

	return &AITaskBriefInfo{
		ID:                task.ID,
		JobType:           task.JobType,
		Status:            task.Status,
		DetailedStatus:    task.DetailedStatus,
		DisplayJobName:    html.EscapeString(task.DisplayJobName),
		CreatedUnix:       task.CreatedUnix,
		FormattedDuration: task.TrainJobDuration,
		Cluster:           GetClusterTypeFromCloudbrainType(task.Type).GetParentCluster(),
		ComputeSource:     task.GetStandardComputeSource(),
		StartTime:         task.StartTime,
		EndTime:           task.EndTime,
		AICenter:          aiCenter,
		IsFileNotebook:    task.IsFileNoteBookTask(),
		IsFineTuneTask:    task.FineTune,
		APPName:           html.EscapeString(task.AppName),
		AccCardType:       accCardType,
		JobName:           task.JobName,
		LabelName:         html.EscapeString(task.LabelName),
		JobId:             task.JobID,
		EndPoint:          task.EndPoint,
		Port:              task.Port,
		AICenterCode:      aiCenterCode,
		ImageId:           task.ImageID,
		ImageName:         task.Image,
	}
}

type NotebookDataset struct {
	DatasetUrl    string `json:"dataset_url"`
	DatasetName   string `json:"dataset_name"`
	ContainerPath string `json:"containerPath"`
	ReadOnly      bool   `json:"readOnly"`
}

type QueryLogOpts struct {
	CloudbrainId int64
	BaseLine     int64
	Lines        int64
	Order        Direction
	UserId       int64
	NodeId       int
	LogFileName  string
}

type FinetuneLoss struct {
	Loss         float32 `json:"loss"`
	GradNorm     float32 `json:"grad_norm"`
	LearningRate float32 `json:"learning_rate"`
	Epoch        float32 `json:"epoch"`
}

type GetLogDownloadInfoReq struct {
	CloudbrainId int64
	NodeId       int
	LogFileName  string
}
type GetAllOutputReq struct {
	CloudbrainId int64
	Suffix       []string
}

type DownloadAllFileReq struct {
	CloudbrainId int64
	FileName     string
	ParentDir    string
	ZIPWriter    *zip.Writer
}

type GetSingleDownloadInfoReq struct {
	CloudbrainId int64
	FileName     string
	ParentDir    string
}

type Direction string

const (
	UP   Direction = "up"
	DOWN Direction = "down"
)

func (o Direction) Reverse() Direction {
	switch o {
	case DOWN:
		return UP
	case UP:
		return DOWN
	}
	return ""
}

type FileReader struct {
	Reader io.ReadCloser
	Name   string
}

type GetTaskListReq struct {
	models.ListOptions
	ComputeSource   *models.ComputeSource
	JobTypes        []string
	RepoID          int64
	AICenter        string
	JobStatus       string
	ExcludeStatus   []string
	Operator        *models.User
	IsRepoOwner     bool
	Cluster         string
	ExcludeJobTypes []string
	IsAimRequired   bool
}
type GetMyTaskListReq struct {
	models.ListOptions
	ComputeSource *models.ComputeSource
	JobType       string
	JobStatus     string
	User          *models.User
	IsRepoOwner   bool
	Keyword       string
	AICenter      string
	Cluster       string
	ExcludeStatus []string
	AppName       string
	BeginTimeUnix int64
	EndTimeUnix   int64
	OrderBy       string
}

type AITaskBaseConfig struct {
	ContainerSteps      map[ContainerDataType]*ContainerBuildOpts `json:"container_configs"`
	ActionType          models.ActionType                         `json:"action_type"`
	IsActionUseJobId    bool                                      `json:"is_action_use_job_id"`
	DatasetsLimitSizeGB int
	DatasetsMaxNum      int
	ModelMaxNum         int
	ModelLimitSizeGB    int
	AutoStopDuration    int64
	//主动检测调式地址是否可用
	DebugAddressCheck bool
}

func GetAITaskConfigFromCloudbrainConfig(config *models.CloudbrainConfig) *AITaskBaseConfig {
	if config == nil {
		return nil
	}
	s := config.ConfigurationSnapshot
	c := &AITaskBaseConfig{}
	err := json.Unmarshal([]byte(s), c)
	if err != nil {
		log.Error("GetAITaskConfigFromCloudbrain err,config=%+v err=&v", config, err)
		return nil
	}
	return c
}

type AITaskDetailConfigInfo struct {
	BaseConfig         *AITaskBaseConfig
	OutputObjectPrefix string
	OutputStorageType  StorageType
	LogObjectPrefix    string
	LogStorageType     StorageType
}

func BuildAITaskDetailConfigInfo(config *models.CloudbrainConfig) *AITaskDetailConfigInfo {
	c := &AITaskBaseConfig{}
	json.Unmarshal([]byte(config.ConfigurationSnapshot), c)
	return &AITaskDetailConfigInfo{
		BaseConfig:         c,
		OutputObjectPrefix: config.OutputObjectPrefix,
		OutputStorageType:  StorageType(config.OutputStorageType),
		LogObjectPrefix:    config.LogObjectPrefix,
		LogStorageType:     StorageType(config.LogStorageType),
	}
}

type AITaskConfigKey struct {
	ComputeSource         string
	IsFileNoteBookRequest bool
}

func (opts AITaskConfigKey) GetKey() string {
	v := reflect.ValueOf(opts)
	t := v.Type()
	b := strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		fieldValue := field.Interface()
		if !field.IsZero() {
			b.WriteString(fmt.Sprintf("%s:%v;", fieldName, fieldValue))
		}
	}
	return b.String()
}

func (c *AITaskBaseConfig) GetContainerConfig(containerDataType ContainerDataType) *ContainerBuildOpts {
	containerConfigs := c.ContainerSteps
	if containerConfigs != nil {
		return containerConfigs[containerDataType]
	}
	return nil

}
func (c *AITaskBaseConfig) GetContainerPath(containerDataType ContainerDataType) string {
	config := c.GetContainerConfig(containerDataType)
	if config == nil {
		return ""
	}
	return config.ContainerPath

}

type AITaskConfigMap struct {
	mu        sync.RWMutex
	ConfigMap map[string]*AITaskBaseConfig
}

func (h *AITaskConfigMap) Add(opts AITaskConfigKey, config *AITaskBaseConfig) *AITaskConfigMap {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.ConfigMap == nil {
		h.ConfigMap = make(map[string]*AITaskBaseConfig, 0)
	}
	h.ConfigMap[opts.GetKey()] = config
	return h
}
func (h *AITaskConfigMap) Default(config *AITaskBaseConfig) *AITaskConfigMap {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.ConfigMap == nil {
		h.ConfigMap = make(map[string]*AITaskBaseConfig, 0)
	}
	h.ConfigMap[AITaskConfigKey{}.GetKey()] = config
	return h
}

func (h AITaskConfigMap) Get(opts AITaskConfigKey) *AITaskBaseConfig {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.ConfigMap == nil {
		return nil
	}
	key := opts.GetKey()
	if _, isOk := h.ConfigMap[key]; isOk {
		return h.ConfigMap[key]
	}
	return nil
}

func (h AITaskConfigMap) IsEmpty() bool {
	return h.ConfigMap == nil || len(h.ConfigMap) == 0
}

type AITaskOutput struct {
	Status          models.ModelMigrateStatus `json:"status"`
	Path            string                    `json:"path"`
	FileList        []storage.FileInfo        `json:"file_list"`
	IsTaskTerminal  bool                      `json:"is_task_terminal"`
	CanReschedule   bool                      `json:"can_reschedule"`
	CanDownload     bool                      `json:"can_download"`
	OutputSizeLimit int                       `json:"output_size_limit"`
}
type AllAITaskOutput struct {
	FileList []storage.FileInfo `json:"file_list"`
}

type GetResourceUsageOpts struct {
	CloudbrainId int64
	NodeId       int
	LogFileName  string
}
type GetSpecOpts struct {
	UserId            int64
	ComputeSource     models.ComputeSource
	JobType           models.JobType
	HasInternet       models.SpecInternetQuery //0 all;1 no internet;2 has internet
	VisualizeRequired bool
}

type AITaskNodeInfo struct {
	ID          int    `json:"id"`
	LogFileName string `json:"log_file_name"`
}

type StorageObjectInfo struct {
	ObjectKey   string
	StorageType StorageType
	Bucket      string
}

type SDKCodeOpts struct {
	DatasetNames       []string
	PretrainModelNames []string
	ParameterKeys      []string
	JobType            models.JobType
	VisualizeRequired  bool
}

type TraceInfo struct {
	TaskId    int64
	TaskName  string
	JobType   string
	SpanName  string
	ParentCtx *context.Context
}
