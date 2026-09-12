package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/util"

	"xorm.io/builder"
	"xorm.io/xorm"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

type CloudbrainStatus string
type JobType string
type ModelArtsJobStatus string

const (
	TypeCloudBrainOne int = iota
	TypeCloudBrainTwo
	TypeC2Net    //智算网络
	TypeCDCenter //成都智算中心
	TypeIFLYTek  //讯飞火石

	TypeCloudBrainAll = -1
	AccCardsNumAll    = -1
	JobNoTeminal      = -1
)

const (
	PanguCustom int = iota
	PanguTextClassification
	PanguTranslation
	PanguOpenDialog
)
const (
	PanguModelFineTune int = iota
)

const (
	NPUResource = "NPU"
	GPUResource = "CPU/GPU"
	GCUResource = "GCU"
	CPUResource = "CPU"
	AllResource = "all"

	//notebook storage category
	EVSCategory = "EVS"
	EFSCategory = "EFS"

	ManagedOwnership  = "MANAGED"
	DetectedOwnership = "DEDICATED"

	NotebookFeature = "NOTEBOOK"
	DefaultFeature  = "DEFAULT"

	DefaultComfyuiTag = "comfyui"

	JobWaiting   CloudbrainStatus = "WAITING"
	JobStopped   CloudbrainStatus = "STOPPED"
	JobSucceeded CloudbrainStatus = "SUCCEEDED"
	JobFailed    CloudbrainStatus = "FAILED"
	JobRunning   CloudbrainStatus = "RUNNING"

	ModelSafetyTesting CloudbrainStatus = "TESTING"

	JobTypeDebug             JobType = "DEBUG"
	JobTypeBenchmark         JobType = "BENCHMARK"
	JobTypeModelSafety       JobType = "MODELSAFETY"
	JobTypeSnn4imagenet      JobType = "SNN4IMAGENET"
	JobTypeBrainScore        JobType = "BRAINSCORE"
	JobTypeSnn4Ecoset        JobType = "SNN4ECOSET"
	JobTypeSim2BrainSNN      JobType = "SIM2BRAIN_SNN"
	JobTypeTrain             JobType = "TRAIN"
	JobTypeInference         JobType = "INFERENCE"
	JobTypeOnlineInference   JobType = "ONLINEINFERENCE"
	JobTypeModelExperience   JobType = "MODELEXPERIENCE"
	JobTypeSuperCompute      JobType = "HPC"
	JobTypeGeneral           JobType = "GENERAL"
	JobTypeFINETUNE          JobType = "FINETUNE"
	JobTypeSdFinetune        JobType = "SDFINETUNE"
	JobTypeComfyuiExperience JobType = "ComfyuiExperience"
	JobTypeEval              JobType = "EVAL"

	//notebook
	ModelArtsCreateQueue  ModelArtsJobStatus = "CREATE_QUEUING" //免费资源创建排队中
	ModelArtsCreating     ModelArtsJobStatus = "CREATING"       //创建中
	ModelArtsCreateFailed ModelArtsJobStatus = "CREATE_FAILED"  //创建失败
	ModelArtsStartQueuing ModelArtsJobStatus = "START_QUEUING"  //免费资源启动排队中
	ModelArtsReadyToStart ModelArtsJobStatus = "READY_TO_START" //免费资源等待启动
	ModelArtsStarting     ModelArtsJobStatus = "STARTING"       //启动中
	ModelArtsRestarting   ModelArtsJobStatus = "RESTARTING"     //重启中
	ModelArtsStartFailed  ModelArtsJobStatus = "START_FAILED"   //启动失败
	ModelArtsRunning      ModelArtsJobStatus = "RUNNING"        //运行中
	ModelArtsStopping     ModelArtsJobStatus = "STOPPING"       //停止中
	ModelArtsStopped      ModelArtsJobStatus = "STOPPED"        //停止
	ModelArtsUnavailable  ModelArtsJobStatus = "UNAVAILABLE"    //故障
	ModelArtsDeleting     ModelArtsJobStatus = "DELETING"       //删除中
	ModelArtsDeleted      ModelArtsJobStatus = "DELETED"        //已删除
	ModelArtsResizing     ModelArtsJobStatus = "RESIZING"       //规格变更中
	ModelArtsResizFailed  ModelArtsJobStatus = "RESIZE_FAILED"  //规格变更失败

	//trainjob
	ModelArtsTrainJobUnknown               ModelArtsJobStatus = "UNKNOWN"                 //作业状态未知
	ModelArtsTrainJobInit                  ModelArtsJobStatus = "INIT"                    //作业初始化状态
	ModelArtsTrainJobImageCreating         ModelArtsJobStatus = "IMAGE_CREATING"          //作业镜像正在创建
	ModelArtsTrainJobImageFailed           ModelArtsJobStatus = "IMAGE_FAILED"            //作业镜像创建失败
	ModelArtsTrainJobSubmitTrying          ModelArtsJobStatus = "SUBMIT_TRYING"           //作业正在提交
	ModelArtsTrainJobSubmitFailed          ModelArtsJobStatus = "SUBMIT_FAILED"           //作业提交失败
	ModelArtsTrainJobDeleteFailed          ModelArtsJobStatus = "DELETE_FAILED"           //作业删除失败
	ModelArtsTrainJobWaiting               ModelArtsJobStatus = "WAITING"                 //作业正在排队中
	ModelArtsTrainJobRunning               ModelArtsJobStatus = "RUNNING"                 //作业正在运行中
	ModelArtsTrainJobKilling               ModelArtsJobStatus = "KILLING"                 //作业正在取消
	ModelArtsTrainJobCompleted             ModelArtsJobStatus = "COMPLETED"               //作业已经完成
	ModelArtsTrainJobFailed                ModelArtsJobStatus = "FAILED"                  //作业运行失败
	ModelArtsTrainJobKilled                ModelArtsJobStatus = "KILLED"                  //作业取消成功
	ModelArtsTrainJobCanceled              ModelArtsJobStatus = "CANCELED"                //作业取消
	ModelArtsTrainJobLost                  ModelArtsJobStatus = "LOST"                    //作业丢失
	ModelArtsTrainJobScaling               ModelArtsJobStatus = "SCALING"                 //作业正在扩容
	ModelArtsTrainJobSubmitModelFailed     ModelArtsJobStatus = "SUBMIT_MODEL_FAILED"     //提交模型失败
	ModelArtsTrainJobDeployServiceFailed   ModelArtsJobStatus = "DEPLOY_SERVICE_FAILED"   //部署服务失败
	ModelArtsTrainJobCheckInit             ModelArtsJobStatus = "CHECK_INIT"              //审核作业初始化
	ModelArtsTrainJobCheckRunning          ModelArtsJobStatus = "CHECK_RUNNING"           //审核作业正在运行中
	ModelArtsTrainJobCheckRunningCompleted ModelArtsJobStatus = "CHECK_RUNNING_COMPLETED" //审核作业已经完成
	ModelArtsTrainJobCheckFailed           ModelArtsJobStatus = "CHECK_FAILED"            //审核作业失败

	DURATION_STR_ZERO     = "00:00:00"
	CloudbrainKeyDuration = 24 * time.Hour

	//grampus
	GrampusStatusPending   = "pending"
	GrampusStatusRunning   = "RUNNING"
	GrampusStatusFailed    = "FAILED"
	GrampusStatusSucceeded = "SUCCEEDED"
	GrampusStatusStopped   = "STOPPED"
	GrampusStatusStopping  = "STOPPING"
	GrampusStatusUnknown   = "UNKNOWN"
	GrampusStatusWaiting   = "WAITING"

	ModelSuffix = "models.zip"

	//local status
	LocalStatusPreparing = "PREPARING"
	LocalStatusCreating  = "CONNECTING"
	LocalStatusFailed    = "CREATED_FAILED"
)

const (
	//cluster
	OpenICluster   = "OpenI"
	C2NetCluster   = "C2Net"
	IFLYTEKCluster = "IFLYTEKTraining"

	//cloudbrain two sdk PathValue
	LocalCodePath          = "/home/ma-user/work/code"
	LocalDatasetPath       = "/home/ma-user/work/dataset"
	LocalPretrainModelPath = "/home/ma-user/work/pretrainmodel"
	LocalOutputPath        = "/home/ma-user/work/output"

	DataDownloadMethodMount  = "MOUNT"
	DataDownloadMethodMoxing = "MOXING"

	CodeNeedUnzipTrue  = "true"
	CodeNeedUnzipFalse = "false"

	DatasetNeedUnzipTrue  = "true"
	DatasetNeedUnzipFalse = "false"

	PretrainModelNeedUnzipTrue  = "true"
	PretrainModelNeedUnzipFalse = "false"

	//AI center
	AICenterOfCloudBrainOne = "OpenIOne"
	AICenterOfCloudBrainTwo = "OpenITwo"
	AICenterOfChengdu       = "OpenIChengdu"
	AICenterOfHuoShi        = "OpenI-huoshi"

	//ComputeResource
	GPU      = "GPU"
	NPU      = "NPU"
	GCU      = "GCU"
	MLU      = "MLU"
	DCU      = "DCU"
	CPU      = "CPU"
	ILUVATAR = "ILUVATAR-GPGPU"
	METAX    = "METAX-GPGPU"
	GPGPU    = "GPGPU"
	BIREN    = "BIREN-GPU"

	ProcessorTypeNPU      = "npu.huawei.com/NPU"
	ProcessorTypeGPU      = "nvidia.com/gpu"
	ProcessorTypeGCU      = "enflame-tech.com/gcu"
	ProcessorTypeMLU      = "cambricon.com/mlu"
	ProcessorTypeDCU      = "ac.sugon.com/dcu"
	ProcessorTypeCPU      = "hpc/cpu"
	ProcessorTypeILUVATAR = "iluvatar.com/iluvatar-gpgpu"
	ProcessorTypeMETAX    = "metax-tech.com/metax-gpgpu"
	ProcessorTypeBIREN    = "biren.tech.com/biren-gpu"
)

const (
	ModelAppLoading = "loading"
	ModelAppLoaded  = "loaded"
)

const (
	OutputMigrateSystemUrchin1 = "urchin1.0"
	OutputMigrateSystemUrchin2 = "urchin2.0"
)

const CloudbrainTwoDefaultVersion = "/V0001"

type ComputeSource struct {
	Name             string
	CloudbrainFormat string
	FullName         string
}

func (c *ComputeSource) GetCloudbrainFormat() string {
	if c.CloudbrainFormat != "" {
		return c.CloudbrainFormat
	}
	return c.Name
}

func GetComputeSourceInstance(name string) *ComputeSource {
	for n, v := range ComputeSourceMap {
		if n == name || name == v.GetCloudbrainFormat() || name == v.FullName {
			return ComputeSourceMap[n]
		}
	}
	return nil
}

func GetComputeSourceCloudbrainFormat(name string) string {
	c := GetComputeSourceInstance(name)
	if c == nil {
		return ""
	}
	return c.CloudbrainFormat
}

func GetComputeSourceStandardFormat(name string) string {
	c := GetComputeSourceInstance(name)
	if c == nil {
		return ""
	}
	return c.Name
}

var ComputeSourceMap = map[string]*ComputeSource{
	GPU:      {Name: GPU, CloudbrainFormat: GPUResource, FullName: ProcessorTypeGPU},
	NPU:      {Name: NPU, FullName: ProcessorTypeNPU},
	GCU:      {Name: GCU, FullName: ProcessorTypeGCU},
	MLU:      {Name: MLU, FullName: ProcessorTypeMLU},
	DCU:      {Name: DCU, FullName: ProcessorTypeDCU},
	CPU:      {Name: CPU, FullName: ProcessorTypeCPU},
	ILUVATAR: {Name: ILUVATAR, FullName: ProcessorTypeILUVATAR},
	METAX:    {Name: METAX, FullName: ProcessorTypeMETAX},
	BIREN:    {Name: BIREN, FullName: ProcessorTypeBIREN},
}

const (
	AIModelPath = "aimodels/"
)

type Cloudbrain struct {
	ID                   int64  `xorm:"pk autoincr"`
	JobID                string `xorm:"INDEX NOT NULL"`
	JobType              string `xorm:"INDEX NOT NULL DEFAULT 'DEBUG'"`
	JobName              string
	DisplayJobName       string
	Status               string
	DetailedStatus       string `xorm:"DEFAULT '-'"`
	UserID               int64  `xorm:"INDEX NOT NULL"`
	RepoID               int64  `xorm:"INDEX NOT NULL"`
	SubTaskName          string
	ContainerID          string
	ContainerIp          string
	CreatedUnix          timeutil.TimeStamp `xorm:"INDEX"`
	UpdatedUnix          timeutil.TimeStamp `xorm:"INDEX updated"`
	Duration             int64              `xorm:"DEFAULT 0"` //运行时长 单位秒
	TrainJobDuration     string             `xorm:"DEFAULT '00:00:00'"`
	Image                string             //镜像名称
	GpuQueue             string             //GPU类型即GPU队列
	ResourceSpecId       int                //GPU规格id
	DeletedAt            time.Time          `xorm:"deleted"`
	CanDebug             bool               `xorm:"-"`
	CanDel               bool               `xorm:"-"`
	CanModify            bool               `xorm:"-"`
	Type                 int                `xorm:"INDEX"`
	BenchmarkTypeID      int
	BenchmarkChildTypeID int
	CardType             string
	Cluster              string

	VersionID       int64  //版本id
	VersionName     string `xorm:"INDEX"` //当前版本
	Uuid            string //数据集id
	DatasetName     string `xorm:"varchar(2000)"`
	DatasetAlias    string `xorm:"varchar(2000)"`
	VersionCount    int    //任务的当前版本数量，不包括删除的
	IsLatestVersion string //是否是最新版本，1是，0否
	CommitID        string //提交的仓库代码id
	PreVersionName  string //父版本名称
	ComputeResource string //计算资源，例如npu
	EngineID        int64  //引擎id
	ImageID         string //grampus image_id
	AiCenter        string //grampus ai center: center_id+center_name
	QueueCode       string
	FailedReason    string `xorm:"text"`

	TrainUrl              string      //输出模型的obs路径
	RemoteCodeUrl         string      //分中心下载代码地址
	BranchName            string      `xorm:"varchar(2550)"` //分支名称
	Parameters            string      `xorm:"varchar(2000)"` //传给modelarts的param参数
	BootFile              string      `xorm:"varchar(2550)"` //启动文件
	DataUrl               string      `xorm:"varchar(3500)"` //数据集的obs路径
	LogUrl                string      //日志输出的obs路径
	PreVersionId          int64       //父版本的版本id
	FlavorCode            string      //modelarts上的规格id
	Description           string      `xorm:"varchar(2550)"` //描述
	WorkServerNumber      int         //节点数
	FlavorName            string      //规格名称
	EngineName            string      //引擎名称
	TotalVersionCount     int         //任务的所有版本数量，包括删除的
	LabelName             string      //标签名称
	ModelName             string      //模型名称
	ModelVersion          string      //模型版本
	CkptName              string      `xorm:"varchar(2550)"` //权重文件名称
	ModelId               string      //模型ID
	ModelRepoName         string      `xorm:"-"`
	ModelRepoOwnerName    string      `xorm:"-"`
	PreTrainModelUrl      string      //预训练模型地址
	ResultUrl             string      //推理结果的obs路径
	ResultJson            string      `xorm:"varchar(4000)"`
	User                  *User       `xorm:"-"`
	Repo                  *Repository `xorm:"-"`
	BenchmarkType         string      `xorm:"-"` //算法评测，模型评测
	BenchmarkTypeName     string      `xorm:"-"`
	BenchmarkTypeRankLink string      `xorm:"-"`
	StartTime             timeutil.TimeStamp
	EndTime               timeutil.TimeStamp
	Cleared               bool `xorm:"DEFAULT false"`
	FineTune              bool `xorm:"DEFAULT false"`
	FineTuneModelType     int
	FineTuneCategory      int
	Spec                  *Specification    `xorm:"-"`
	Config                *CloudbrainConfig `xorm:"-"`
	AppName               string            //超算任务的应用类型
	HasInternet           int
	TimeLimit             int `xorm:"DEFAULT 0"` //just for subscriber
	Port                  int
	EndPoint              string
	VisualizeRequired     bool
	AimRequired           bool
	SourceID              int64
	TemplateID            string
}

type CloudbrainShow struct {
	ID               int64
	JobID            string
	RepoFullName     string
	RepoID           int64
	Type             int
	JobType          string
	DisplayJobName   string
	Duration         string
	ResourceSpec     *Specification
	ComputeResource  string
	AiCenter         string
	WorkServerNumber int
}

func (c *CloudbrainShow) GetChineseJobType() string {
	switch JobType(c.JobType) {
	case JobTypeDebug:
		return "调试任务"
	case JobTypeBenchmark, JobTypeModelSafety:
		return "评测任务"
	case JobTypeTrain:
		return "训练任务"
	case JobTypeInference:
		return "推理任务"
	case JobTypeOnlineInference:
		return "在线推理"
	case JobTypeModelExperience:
		return "在线体验"
	case JobTypeSuperCompute:
		return "超算任务"
	case JobTypeSdFinetune:
		return "CV微调任务"
	case JobTypeComfyuiExperience:
		return "Comfyui体验任务"
	}
	return "未知类型"
}

type CloudbrainShow4Action struct {
	ID              int64
	JobID           string
	Type            int
	JobType         string
	DisplayJobName  string
	ComputeResource string
}

func (task *Cloudbrain) ToShow() *CloudbrainShow {
	n := 1
	if task.WorkServerNumber > 1 {
		n = task.WorkServerNumber
	}
	task.LoadSpec()
	c := &CloudbrainShow{
		ID:               task.ID,
		JobID:            task.JobID,
		JobType:          task.JobType,
		Type:             task.Type,
		DisplayJobName:   task.DisplayJobName,
		Duration:         task.TrainJobDuration,
		ResourceSpec:     task.Spec,
		ComputeResource:  task.ComputeResource,
		WorkServerNumber: n,
		RepoID:           task.RepoID,
	}
	task.GetRepository()
	if task.Repo != nil {
		c.RepoFullName = task.Repo.FullName()
	}
	return c
}

func (task *Cloudbrain) GetRepository() *Repository {
	if task.Repo != nil {
		return task.Repo
	}
	if task.RepoID == 0 {
		return nil
	}
	repo, _ := GetRepositoryByID(task.RepoID)
	task.Repo = repo
	return repo
}

func (task *Cloudbrain) LoadSpec() error {
	if task.Spec != nil {
		return nil
	}
	spec := &CloudbrainSpec{}
	_, err := x.Where("cloudbrain_id = ?", task.ID).Get(spec)
	if err != nil {
		return err
	}
	task.Spec = spec.ConvertToSpecification()
	return nil
}

func (task *Cloudbrain) IsRestartTask() bool {
	n, _ := x.Where("display_job_name = ?", task.DisplayJobName).Unscoped().Count(&Cloudbrain{})
	if n > 1 {
		return true
	}
	return false
}

func (task *Cloudbrain) HasUseModel(modelId string) bool {
	modelIDArray := task.GetModelIdArray()
	if modelIDArray == nil || len(modelIDArray) == 0 {
		return false
	}
	for _, id := range modelIDArray {
		if id == modelId {
			return true
		}
	}
	return false
}

func (task *Cloudbrain) GetModelIdArray() []string {
	if task.ModelId == "" {
		return []string{}
	}
	modelIdStr := strings.TrimSuffix(task.ModelId, ";")
	modelIDArray := strings.Split(modelIdStr, ";")
	return modelIDArray
}

func (task *Cloudbrain) GetModelNameArray() []string {
	if task.ModelName == "" {
		return []string{}
	}
	modelNameStr := strings.TrimSuffix(task.ModelName, ";")
	modelNameArray := strings.Split(modelNameStr, ";")
	return modelNameArray
}

func (task *Cloudbrain) GetStandardComputeSource() string {
	return GetComputeSourceStandardFormat(task.ComputeResource)
}
func (task *Cloudbrain) GetCloudbrainConfig() *CloudbrainConfig {
	if task.Config != nil {
		return task.Config
	}
	c, err := GetCloudbrainConfig(task.ID)
	if err != nil {
		return nil
	}
	task.Config = c
	return c
}

func (task *Cloudbrain) ComputeAndSetDuration() {
	var d int64
	if task.StartTime == 0 {
		d = 0
	} else if task.EndTime == 0 {
		if !task.IsTerminal() {
			d = time.Now().Unix() - task.StartTime.AsTime().Unix()
		}
	} else {
		d = task.EndTime.AsTime().Unix() - task.StartTime.AsTime().Unix()
	}

	if d < 0 {
		d = 0
	}
	task.Duration = d
	task.TrainJobDuration = ConvertDurationToStr(d)
}

func (task *Cloudbrain) CorrectCreateUnix() {
	if task.StartTime > 0 && task.CreatedUnix > task.StartTime {
		task.CreatedUnix = task.StartTime
	}
}
func (task *Cloudbrain) GetAiCenter() string {
	if task.Type == TypeCloudBrainOne {
		return AICenterOfCloudBrainOne
	} else if task.Type == TypeCloudBrainTwo {
		return AICenterOfCloudBrainTwo
	} else if task.Type == TypeCDCenter {
		return AICenterOfChengdu
	} else {
		return strings.Split(task.AiCenter, "+")[0]
	}

}

func (task *Cloudbrain) GetAiCenterName() string {
	if task.Type == TypeCloudBrainOne {
		return "云脑一"
	} else if task.Type == TypeCloudBrainTwo {
		return "云脑二"
	} else if task.Type == TypeCDCenter {
		return "启智成都智算"
	} else {
		tmpArray := strings.Split(task.AiCenter, "+")
		if len(tmpArray) == 2 {
			return tmpArray[1]
		}
	}
	return task.AiCenter
}

// 是否为在线notebook文件任务
func (task *Cloudbrain) IsFileNoteBookTask() bool {
	return task.JobType == string(JobTypeDebug) && task.BootFile != ""
}

func (task *Cloudbrain) CanUserModify(user *User) bool {
	if user == nil {
		return false
	}
	if user.IsAdmin {
		return true
	}

	return user.ID == task.UserID
}
func (task *Cloudbrain) CanUserDelete(user *User, isRepoOwner bool) bool {
	if user == nil {
		return false
	}
	return user.IsAdmin || user.ID == task.UserID
}

func (task *Cloudbrain) CanExperience(user *User) bool {
	if user == nil {
		return false
	}
	if user.IsAdmin {
		return true
	}
	log.Info("user.id=" + fmt.Sprint(user.ID) + " task_id=" + fmt.Sprint(task.UserID))
	return user.ID == task.UserID
}
func (task *Cloudbrain) CanFintuneExperience(user *User) bool {

	if task.JobType != string(JobTypeFINETUNE) {
		return false
	}
	if task.ComputeResource == ILUVATAR {
		return false
	}

	if (task.Status != string(ModelArtsTrainJobCompleted) && task.Status != string(JobSucceeded)) || task.Cleared {
		return false
	}

	migrateRecord, err := GetModelMigrateRecordByCloudbrainId(task.ID)
	if err != nil {
		log.Error("GetModelMigrateRecordByCloudbrainId error:%v", err)
		return false
	}
	if migrateRecord.Status != ModelMigrateSuccess {
		return false
	}

	if user == nil {
		return false
	}
	config := setting.GetModelFinetuneConfig()
	if config == nil {
		return false
	}
	matchResource := false
	if len(config.LLAMA.ModelInfos) > 0 {
		for _, modelInfo := range config.LLAMA.ModelInfos {
			if modelInfo.Id == task.ModelId {

				if modelInfo.Experience.ComputeResourceInfo.GPU.Parameters != "" && task.ComputeResource == GPUResource {
					matchResource = true
					break
				}
				if modelInfo.Experience.ComputeResourceInfo.NPU.Parameters != "" && task.ComputeResource == NPUResource {
					matchResource = true
					break
				}

			}
		}
	}

	if !matchResource {
		return false
	}

	if user.IsAdmin {
		return true
	}
	log.Info("user.id=" + fmt.Sprint(user.ID) + " task_id=" + fmt.Sprint(task.UserID))
	return user.ID == task.UserID
}

var CanCreateAITaskTemplateJobs = []JobType{JobTypeDebug, JobTypeTrain, JobTypeOnlineInference, JobTypeGeneral, JobTypeSuperCompute}

func (task *Cloudbrain) CanCreateTemplate(user *User) bool {
	if user == nil {
		return false
	}
	canCreate := false
	for _, v := range CanCreateAITaskTemplateJobs {
		if v == JobType(task.JobType) {
			canCreate = true
			break
		}
	}
	if !canCreate {
		return false
	}
	return user.IsAdmin || user.ID == task.UserID
}

func AllTerminalStatus() []string {
	return []string{string(ModelArtsTrainJobCompleted), string(ModelArtsTrainJobFailed),
		string(ModelArtsTrainJobKilled), string(ModelArtsStopped), string(ModelArtsCreateFailed),
		string(ModelArtsStartFailed), string(JobStopped), string(JobFailed),
		string(JobSucceeded), GrampusStatusFailed,
		GrampusStatusSucceeded, GrampusStatusStopped, LocalStatusFailed}
}

// AllFailedOrStoppedStatus 返回失败或主动停止的状态列表（不包含成功状态）
func AllFailedOrStoppedStatus() []string {
	return []string{string(ModelArtsTrainJobFailed), string(ModelArtsTrainJobKilled),
		string(ModelArtsStopped), string(ModelArtsCreateFailed), string(ModelArtsStartFailed),
		string(JobStopped), string(JobFailed), GrampusStatusFailed,
		GrampusStatusStopped, LocalStatusFailed, string(ModelArtsStopping),
		string(ModelArtsDeleting), string(ModelArtsTrainJobKilling), GrampusStatusStopping}
}

func IsCloudbrainTerminalStatus(status string) bool {
	for _, s := range AllTerminalStatus() {
		if strings.ToUpper(status) == strings.ToUpper(s) {
			return true
		}
	}
	return false
}

// IsCloudbrainFailedOrStopped 检查任务状态是否为失败或主动停止（不包含成功）
func IsCloudbrainFailedOrStopped(status string) bool {
	for _, s := range AllFailedOrStoppedStatus() {
		if strings.ToUpper(status) == strings.ToUpper(s) {
			return true
		}
	}
	return false
}

func AllStoppingStatus() []string {
	return []string{string(ModelArtsStopping), string(ModelArtsDeleting),
		string(ModelArtsTrainJobKilling), GrampusStatusStopping}
}
func AllRunningOrWaitingStatus() []string {
	return []string{string(JobRunning), string(JobWaiting),
		string(ModelArtsCreating), string(ModelArtsStartQueuing),
		string(ModelArtsReadyToStart), string(ModelArtsCreateQueue),
		string(ModelArtsStarting), string(ModelArtsRestarting),
		string(ModelArtsRunning), string(ModelArtsTrainJobSubmitTrying),
		string(ModelArtsTrainJobWaiting), string(ModelArtsTrainJobRunning),
		GrampusStatusPending, GrampusStatusRunning, GrampusStatusWaiting,
	}
}

func IsCloudbrainInstatus(status string, statusList []string) bool {
	for _, s := range statusList {
		if s == status {
			return true
		}
	}
	return false
}

func IsCloudbrainWaitingOrRunning(status string) bool {
	return IsCloudbrainInstatus(status, AllRunningOrWaitingStatus())
}

func AllStoppingAndTerminalStatus() []string {
	var status = AllTerminalStatus()
	return append(status, AllStoppingStatus()...)
}

func (task *Cloudbrain) IsTerminal() bool {
	status := task.Status
	return IsCloudbrainTerminalStatus(status)
}

// IsFailedOrStopped 检查任务是否为失败或主动停止状态（不包含成功）
func (task *Cloudbrain) IsFailedOrStopped() bool {
	return IsCloudbrainFailedOrStopped(task.Status)
}

func (task *Cloudbrain) IsPreparing() bool {
	return task.Status == LocalStatusPreparing
}
func (task *Cloudbrain) IsCreating() bool {
	return task.Status == LocalStatusCreating
}
func (task *Cloudbrain) NeedActiveStop() bool {
	return task.IsCreating() || (task.IsPreparing() && int64(task.CreatedUnix) < time.Now().Add(-1*setting.PREPARING_MAX_WAIT_DURATION).Unix())
}

// 是否允许创建多版本
// 目前只有启智NPU可以
func (task *Cloudbrain) IsAllowedToCreateMultipleVersions() bool {
	if task.Type == TypeCloudBrainTwo && task.ComputeResource == NPUResource && task.JobType != string(JobTypeDebug) {
		return true
	}
	return false
}

func (task *Cloudbrain) IsNewAITask() bool {
	for k, v := range setting.AI_TASK_RANGE {
		if k == task.JobType+"_"+fmt.Sprint(task.Type) {
			if len(v) == 0 {
				continue
			}
			for _, s := range v {
				if s == task.GetStandardComputeSource() {
					return true
				}
			}
		}
	}
	return false
}

func (task *Cloudbrain) IsTerminalOrStopping() bool {
	status := task.Status
	for _, s := range AllStoppingAndTerminalStatus() {
		if status == s {
			return true
		}
	}
	return false
}

func (task *Cloudbrain) IsRunning() bool {
	status := task.Status
	return status == string(ModelArtsTrainJobRunning) || status == string(ModelArtsRunning) ||
		status == string(JobRunning) || status == GrampusStatusRunning
}

func (task *Cloudbrain) IsUserHasRight(user *User) bool {
	if user == nil {
		return false
	}
	return user.IsAdmin || user.ID == task.UserID
}
func (task *Cloudbrain) IsGPUTask() bool {
	return task.ComputeResource == GPUResource
}
func (task *Cloudbrain) IsGCUTask() bool {
	return task.ComputeResource == GCUResource
}
func (task *Cloudbrain) IsNPUTask() bool {
	return task.ComputeResource == NPUResource
}

func (task *Cloudbrain) IsDCUTask() bool {
	return task.ComputeResource == DCU
}

func ConvertDurationToStr(duration int64) string {
	if duration <= 0 {
		return DURATION_STR_ZERO
	}
	return util.AddZero(duration/3600) + ":" + util.AddZero(duration%3600/60) + ":" + util.AddZero(duration%60)
}
func ConvertStrToDuration(trainJobDuration string) int64 {
	trainJobDurationList := strings.Split(trainJobDuration, ":")
	if len(trainJobDurationList) == 3 {
		i, _ := strconv.ParseInt(trainJobDurationList[0], 10, 64)
		j, _ := strconv.ParseInt(trainJobDurationList[1], 10, 64)
		k, _ := strconv.ParseInt(trainJobDurationList[2], 10, 64)
		return i*3600 + j*60 + k
	} else {
		return 0
	}
}

func IsTrainJobTerminal(status string) bool {
	return status == string(ModelArtsTrainJobCompleted) || status == string(ModelArtsTrainJobFailed) || status == string(ModelArtsTrainJobKilled) || status == GrampusStatusFailed || status == GrampusStatusStopped || status == GrampusStatusSucceeded
}

func IsModelArtsDebugJobTerminal(status string) bool {
	return status == string(ModelArtsStopped)
}

func IsCloudBrainOneDebugJobTerminal(status string) bool {
	return status == string(JobStopped) || status == string(JobFailed) || status == string(JobSucceeded)
}
func IsBenchMarkJobType(jobType string) bool {
	types := AllBenchMarkJobType()
	for _, t := range types {
		if jobType == t {
			return true
		}
	}
	return false
}

func AllJobType() []string {
	jobTypes := make([]string, 0)
	jobTypes = append(jobTypes, string(JobTypeDebug), string(JobTypeTrain), string(JobTypeInference))
	jobTypes = append(jobTypes, AllBenchMarkJobType()...)
	return jobTypes
}

func AllBenchMarkJobType() []string {
	jobTypes := make([]string, 0)
	jobTypes = append(jobTypes, string(JobTypeBenchmark), string(JobTypeModelSafety))
	jobTypes = append(jobTypes, AllModelMarkJobType()...)
	return jobTypes
}

func AllModelMarkJobType() []string {
	return []string{string(JobTypeSnn4imagenet), string(JobTypeBrainScore), string(JobTypeSnn4Ecoset), string(JobTypeSim2BrainSNN)}
}

func ParseAndSetDurationFromCloudBrainOne(result JobResultPayload, task *Cloudbrain) {
	isActivated := result.JobStatus.CreatedTime > 0
	if task.StartTime == 0 && isActivated {
		task.StartTime = timeutil.TimeStamp(result.JobStatus.CreatedTime / 1000)
	}
	if task.EndTime == 0 && IsCloudBrainOneDebugJobTerminal(result.JobStatus.State) && isActivated {
		if result.JobStatus.CompletedTime > 0 {
			task.EndTime = timeutil.TimeStamp(result.JobStatus.CompletedTime / 1000)
		}
	}
	task.CorrectCreateUnix()
	task.ComputeAndSetDuration()
}

func ParseAndSetDurationFromModelArtsNotebook(result *GetNotebook2Result, job *Cloudbrain) {
	if job.StartTime == 0 && result.Lease.UpdateTime > 0 {
		job.StartTime = timeutil.TimeStamp(result.Lease.UpdateTime / 1000)
	}
	job.Status = result.Status
	if job.EndTime == 0 && IsModelArtsDebugJobTerminal(job.Status) {
		job.EndTime = timeutil.TimeStampNow()
	}
	job.CorrectCreateUnix()
	job.ComputeAndSetDuration()
}

type CloudbrainInfo struct {
	Cloudbrain `xorm:"extends"`
	User       `xorm:"extends"`
}

type CloudBrainLoginResult struct {
	Code    string
	Msg     string
	Payload map[string]interface{}
}

type TaskRole struct {
	Name                  string `json:"name"`
	TaskNumber            int    `json:"taskNumber"`
	MinSucceededTaskCount int    `json:"minSucceededTaskCount"`
	MinFailedTaskCount    int    `json:"minFailedTaskCount"`
	CPUNumber             int    `json:"cpuNumber"`
	GPUNumber             int    `json:"gpuNumber"`
	MemoryMB              int    `json:"memoryMB"`
	ShmMB                 int    `json:"shmMB"`
	Command               string `json:"command"`
	NeedIBDevice          bool   `json:"needIBDevice"`
	IsMainRole            bool   `json:"isMainRole"`
	UseNNI                bool   `json:"useNNI"`
}

type StHostPath struct {
	Path      string `json:"path"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}

type Volume struct {
	HostPath StHostPath `json:"hostPath"`
}

type CreateJobParams struct {
	JobName    string     `json:"jobName"`
	RetryCount int8       `json:"retryCount"`
	GpuType    string     `json:"gpuType"`
	Image      string     `json:"image"`
	TaskRoles  []TaskRole `json:"taskRoles"`
	Volumes    []Volume   `json:"volumes"`
}

type CreateJobResult struct {
	Code    string                 `json:"code"`
	Msg     string                 `json:"msg"`
	Payload map[string]interface{} `json:"payload"`
}

type QueueDetailResult struct {
	Code    string                 `json:"code"`
	Msg     string                 `json:"msg"`
	Payload map[string]QueueDetail `json:"payload"`
}

type QueueDetail struct {
	JobScheduleInfo JobScheduleInfo `json:"JobScheduleInfo"`
}

type JobScheduleInfo struct {
	Pending                     int `json:"Pending"`
	Running                     int `json:"Running"`
	MedianPendingJobDurationSec int `json:"MedianPendingJobDurationSec"`
}

type GetJobResult struct {
	Code    string                 `json:"code"`
	Msg     string                 `json:"msg"`
	Payload map[string]interface{} `json:"payload"`
}

type GetJobListResult struct {
	Code    string                 `json:"code"`
	Msg     string                 `json:"msg"`
	Payload map[string]interface{} `json:"payload"`
}

type GetImagesResult struct {
	Code    string           `json:"code"`
	Msg     string           `json:"msg"`
	Payload GetImagesPayload `json:"payload"`
}

type GetImagesPayload struct {
	Count      int          `json:"count"`
	TotalPages int          `json:"totalPages,omitempty"`
	ImageInfo  []*ImageInfo `json:"rows"`
}

type CloudbrainsOptions struct {
	ListOptions
	RepoID            int64 // include all repos if empty
	UserID            int64
	JobID             string
	SortType          string
	CloudbrainIDs     []int64
	JobStatus         []string
	JobStatusNot      bool
	Keyword           string
	Type              int
	JobTypes          []string
	VersionName       string
	IsLatestVersion   string
	JobTypeNot        bool
	NeedRepoInfo      bool
	RepoIDList        []int64
	BeginTime         time.Time
	EndTime           time.Time
	ComputeResource   string
	BeginTimeUnix     int64
	EndTimeUnix       int64
	DateBeginTimeUnix int64
	AiCenter          string
	NeedDeleteInfo    string
	Cluster           string
	AccCardType       string
	AccCardsNum       int
	WorkServerNumber  int
	QueueId           int64
	ExcludeJobTypes   []string
	IsAimRequired     bool
	AppName           string
	FilterCleared     bool
	OrderBy           string
}

type QueueCloudbrainOptions struct {
	ComputeResource string
	JobStatus       string
	JobTypes        string
}

type TaskPod struct {
	TaskRoleStatus struct {
		Name string `json:"name"`
	} `json:"taskRoleStatus"`
	//TaskStatuses []struct {
	//	TaskIndex       int       `json:"taskIndex"`
	//	PodUID          string    `json:"podUid"`
	//	PodIP           string    `json:"podIp"`
	//	PodName         string    `json:"podName"`
	//	ContainerID     string    `json:"containerId"`
	//	ContainerIP     string    `json:"containerIp"`
	//	ContainerGpus   string    `json:"containerGpus"`
	//	State           string    `json:"state"`
	//	StartAt         time.Time `json:"startAt"`
	//	FinishedAt      time.Time `json:"finishedAt"`
	//	ExitCode        int       `json:"exitCode"`
	//	ExitDiagnostics string    `json:"exitDiagnostics"`
	//	RetriedCount    int       `json:"retriedCount"`
	//	StartTime       string
	//	FinishedTime    string
	//} `json:"taskStatuses"`
	TaskStatuses []TaskStatuses `json:"taskStatuses"`
}

type TaskStatuses struct {
	TaskIndex       int       `json:"taskIndex"`
	PodUID          string    `json:"podUid"`
	PodIP           string    `json:"podIp"`
	PodName         string    `json:"podName"`
	ContainerID     string    `json:"containerId"`
	ContainerIP     string    `json:"containerIp"`
	ContainerGpus   string    `json:"containerGpus"`
	State           string    `json:"state"`
	StartAt         time.Time `json:"startAt"`
	FinishedAt      time.Time `json:"finishedAt"`
	ExitCode        int       `json:"exitCode"`
	ExitDiagnostics string    `json:"exitDiagnostics"`
	RetriedCount    int       `json:"retriedCount"`
	StartTime       string
	FinishedTime    string
}

type TaskInfo struct {
	Username          string   `json:"username"`
	TaskName          string   `json:"task_name"`
	CodeName          string   `json:"code_name"`
	BenchmarkCategory []string `json:"selected_category"`
	CodeLink          string   `json:"code_link"`
	GpuType           string   `json:"gpu_type"`
}

func ConvertToTaskPod(input map[string]interface{}) (TaskPod, error) {
	data, _ := json.Marshal(input)
	var taskPod TaskPod
	err := json.Unmarshal(data, &taskPod)
	taskPod.TaskStatuses[0].StartTime = time.Unix(taskPod.TaskStatuses[0].StartAt.Unix()+8*3600, 0).UTC().Format("2006-01-02 15:04:05")
	taskPod.TaskStatuses[0].FinishedTime = time.Unix(taskPod.TaskStatuses[0].FinishedAt.Unix()+8*3600, 0).UTC().Format("2006-01-02 15:04:05")
	//if the task is not finished or stopped,the cloudbrain renturns 0001-01-01 08:00:00, the finishedTime shows with -
	if strings.HasPrefix(taskPod.TaskStatuses[0].FinishedTime, "0001") {
		taskPod.TaskStatuses[0].FinishedTime = "-"
	}
	return taskPod, err
}

type JobListResultPayload struct {
	TotalSize int64                    `json:"totalSize"`
	Jobs      []JobResultInListPayload `json:"jobs"`
}

type JobResultInListPayload struct {
	AppExitCode        string `json:"appExitCode"`
	CompletedTime      int64  `json:"completedTime"`
	CreatedTime        int64  `json:"createdTime"`
	ExecutionType      string `json:"executionType"`
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Platform           string `json:"platform"`
	Retries            string `json:"retries"`
	State              string `json:"state"`
	SubState           string `json:"subState"`
	TaskRoleDetailInfo struct {
		Task1 struct {
			TaskRoleStatus struct {
				TaskRoleName string `json:"taskRoleName"`
			} `json:"taskRoleStatus"`
			TaskStatuses struct {
				TaskRoleName    string `json:"taskRoleName"`
				TaskStatusArray []struct {
					ContainerCompletedTimestamp time.Time `json:"containerCompletedTimestamp"`
					ContainerIp                 string    `json:"containerIp"`
					ContainerLaunchedTimestamp  time.Time `json:"containerLaunchedTimestamp"`
					TaskCreatedTimestamp        time.Time `json:"taskCreatedTimestamp"`
					TaskIndex                   int       `json:"taskIndex"`
					TaskRoleName                string    `json:"taskRoleName"`
					TaskState                   string    `json:"taskState"`
				} `json:"taskStatusArray"`
			} `json:"taskStatuses"`
		} `json:"task1"`
	} `json:"taskRoleDetailInfo"`
	Type           string `json:"type"`
	UserId         string `json:"userId"`
	VirtualCluster string `json:"virtualCluster"`
}

type JobResultPayload struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Platform  string `json:"platform"`
	JobStatus struct {
		Username           string      `json:"username"`
		State              string      `json:"state"`
		SubState           string      `json:"subState"`
		ExecutionType      string      `json:"executionType"`
		Retries            int         `json:"retries"`
		CreatedTime        int64       `json:"createdTime"`
		CompletedTime      int64       `json:"completedTime"`
		AppID              string      `json:"appId"`
		AppProgress        string      `json:"appProgress"`
		AppTrackingURL     string      `json:"appTrackingUrl"`
		AppLaunchedTime    int64       `json:"appLaunchedTime"`
		AppCompletedTime   interface{} `json:"appCompletedTime"`
		AppExitCode        int         `json:"appExitCode"`
		AppExitDiagnostics string      `json:"appExitDiagnostics"`
		AppExitType        interface{} `json:"appExitType"`
		VirtualCluster     string      `json:"virtualCluster"`
		StartTime          string
		EndTime            string
	} `json:"jobStatus"`
	TaskRoles map[string]interface{} `json:"taskRoles"`
	Resource  struct {
		CPU          int    `json:"cpu"`
		Memory       string `json:"memory"`
		NvidiaComGpu int    `json:"nvidia.com/gpu"`
	} `json:"resource"`
	Config struct {
		Image     string `json:"image"`
		JobID     string `json:"jobId"`
		GpuType   string `json:"gpuType"`
		JobName   string `json:"jobName"`
		JobType   string `json:"jobType"`
		TaskRoles []struct {
			Name                  string `json:"name"`
			ShmMB                 int    `json:"shmMB"`
			Command               string `json:"command"`
			MemoryMB              int    `json:"memoryMB"`
			CPUNumber             int    `json:"cpuNumber"`
			GpuNumber             int    `json:"gpuNumber"`
			IsMainRole            bool   `json:"isMainRole"`
			TaskNumber            int    `json:"taskNumber"`
			NeedIBDevice          bool   `json:"needIBDevice"`
			MinFailedTaskCount    int    `json:"minFailedTaskCount"`
			MinSucceededTaskCount int    `json:"minSucceededTaskCount"`
		} `json:"taskRoles"`
		RetryCount int `json:"retryCount"`
	} `json:"config"`
	Userinfo struct {
		User  string `json:"user"`
		OrgID string `json:"org_id"`
	} `json:"userinfo"`
}

func ConvertToJobResultPayload(input map[string]interface{}) (JobResultPayload, error) {
	data, _ := json.Marshal(input)
	var jobResultPayload JobResultPayload
	err := json.Unmarshal(data, &jobResultPayload)
	jobResultPayload.JobStatus.StartTime = time.Unix(jobResultPayload.JobStatus.CreatedTime/1000, 0).Format("2006-01-02 15:04:05")
	jobResultPayload.JobStatus.EndTime = time.Unix(jobResultPayload.JobStatus.CompletedTime/1000, 0).Format("2006-01-02 15:04:05")

	if jobResultPayload.JobStatus.State == string(JobWaiting) {
		jobResultPayload.JobStatus.StartTime = "-"
		jobResultPayload.JobStatus.EndTime = "-"
	}
	return jobResultPayload, err
}

func ConvertToJobListResultPayload(input map[string]interface{}) (JobListResultPayload, error) {
	data, _ := json.Marshal(input)
	var r JobListResultPayload
	err := json.Unmarshal(data, &r)
	if err != nil {
		log.Error("")
		return JobListResultPayload{}, err
	}
	for i := 0; i < len(r.Jobs); i++ {

		if r.Jobs[i].State == string(JobWaiting) {
			r.Jobs[i].CreatedTime = 0
			r.Jobs[i].CompletedTime = 0
		}
	}
	return r, nil
}

type ImagesResultPayload struct {
	Images []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Place       string `json:"place"`
		Description string `json:"description"`
		Provider    string `json:"provider"`
		Createtime  string `json:"createtime"`
		Remark      string `json:"remark"`
	} `json:"taskStatuses"`
}
type ImageInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Place       string `json:"place"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	Createtime  string `json:"createtime"`
	Remark      string `json:"remark"`
	IsPublic    int    `json:"isPublic"`
	PlaceView   string
}

type Categories struct {
	Category []*Category `json:"category"`
}

type Category struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type BenchmarkTypes struct {
	BenchmarkType []*BenchmarkType `json:"type"`
}

type BenchmarkType struct {
	Id       int                 `json:"id"`
	RankLink string              `json:"rank_link"`
	First    string              `json:"first"` //一级算法类型名称
	Second   []*BenchmarkDataset `json:"second"`
}

type BenchmarkDataset struct {
	Id         int    `json:"id"`
	Value      string `json:"value"`      //二级算法类型名称
	Attachment string `json:"attachment"` //数据集的uuid
	Owner      string `json:"owner"`      //评估脚本所在仓库的拥有者
	RepoName   string `json:"repo_name"`  //评估脚本所在仓库的名称
}

type GpuInfos struct {
	GpuInfo []*GpuInfo `json:"gpu_type"`
}

type GpuInfo struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
	Queue string `json:"queue"`
}

type ResourceSpecs struct {
	ResourceSpec []*ResourceSpec `json:"resorce_specs"`
}

type ResourceSpec struct {
	Id          int   `json:"id"`
	CpuNum      int   `json:"cpu"`
	GpuNum      int   `json:"gpu"`
	MemMiB      int   `json:"memMiB"`
	ShareMemMiB int   `json:"shareMemMiB"`
	UnitPrice   int64 `json:"unitPrice"`
}

type FlavorInfos struct {
	FlavorInfo []*FlavorInfo `json:"flavor_info"`
}

type FlavorInfo struct {
	Id        int    `json:"id"`
	Value     string `json:"value"`
	Desc      string `json:"desc"`
	UnitPrice int64  `json:"unitPrice"`
}

type SpecialPools struct {
	Pools []*SpecialPool `json:"pools"`
}
type SpecialPool struct {
	Org          string                `json:"org"`
	Type         string                `json:"type"`
	IsExclusive  bool                  `json:"isExclusive"`
	Pool         []*GpuInfo            `json:"pool"`
	JobType      []string              `json:"jobType"`
	ResourceSpec []*ResourceSpec       `json:"resourceSpecs"`
	Flavor       []*setting.FlavorInfo `json:"flavor"`
}

type PoolInfos struct {
	PoolInfo []*PoolInfo `json:"pool_info"`
}

type PoolInfo struct {
	PoolId   string `json:"pool_id"`
	PoolName string `json:"pool_name"`
	PoolType string `json:"pool_type"`
}

type CommitImageCloudBrainParams struct {
	Ip               string `json:"ip"`
	TaskContainerId  string `json:"taskContainerId"`
	ImageTag         string `json:"imageTag"`
	ImageDescription string `json:"imageDescription"`
}

type CommitImageParams struct {
	CommitImageCloudBrainParams
	IsPrivate              bool
	Topics                 []string
	CloudBrainType         int
	UID                    int64
	Place                  string
	Type                   int
	Framework              string
	FrameworkVersion       string
	CudaVersion            string
	DTKVersion             string
	PythonVersion          string
	OperationSystem        string
	OperationSystemVersion string
	ThirdPackages          string
	ComputeResource        string
	TrainType              string
}

type CommitImageResult struct {
	Code    string                 `json:"code"`
	Msg     string                 `json:"msg"`
	Payload map[string]interface{} `json:"payload"`
}

type GetJobLogParams struct {
	Size      string    `json:"size"`
	Sort      string    `json:"sort"`
	QueryInfo QueryInfo `json:"query"`
}

type QueryInfo struct {
	MatchInfo MatchInfo `json:"match"`
}

type MatchInfo struct {
	PodName string `json:"kubernetes.pod.name"`
}

type GetJobLogResult struct {
	ScrollID string `json:"_scroll_id"`
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Hits []Hits `json:"hits"`
	} `json:"hits"`
}

type Hits struct {
	Index  string `json:"_index"`
	Type   string `json:"_type"`
	ID     string `json:"_id"`
	Source struct {
		Message string `json:"message"`
	} `json:"_source"`
	Sort []int `json:"sort"`
}

type GetAllJobLogParams struct {
	Scroll   string `json:"scroll"`
	ScrollID string `json:"scroll_id"`
}

type DeleteJobLogTokenParams struct {
	ScrollID string `json:"scroll_id"`
}

type DeleteJobLogTokenResult struct {
	Succeeded bool `json:"succeeded"`
	NumFreed  int  `json:"num_freed"`
}

type CloudBrainResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type CreateNotebook2Params struct {
	JobName      string               `json:"name"`
	Description  string               `json:"description"`
	Duration     int64                `json:"duration"` //ms
	Feature      string               `json:"feature"`
	PoolID       string               `json:"pool_id"`
	Flavor       string               `json:"flavor"`
	ImageID      string               `json:"image_id"`
	WorkspaceID  string               `json:"workspace_id"`
	Volume       VolumeReq            `json:"volume"`
	EnvVariables CloudBrain2EnvVarReq `json:"env_variables"`
}

type CreateNotebookWithoutPoolParams struct {
	JobName     string    `json:"name"`
	Description string    `json:"description"`
	Duration    int64     `json:"duration"` //ms
	Feature     string    `json:"feature"`
	Flavor      string    `json:"flavor"`
	ImageID     string    `json:"image_id"`
	WorkspaceID string    `json:"workspace_id"`
	Volume      VolumeReq `json:"volume"`
}

type CloudBrain2EnvVarReq struct {
	CodeObsUrl             string `json:"CODE_URL"`
	DatasetObsUrl          string `json:"DATASET_URL"`
	PretrainedModelObsUrl  string `json:"PRETRAIN_MODEL_URL"`
	OutputObsUrl           string `json:"OUTPUT_URL"`
	LocalCodePath          string `json:"LOCAL_CODE_PATH"`
	LocalDatasetPath       string `json:"LOCAL_DATASET_PATH"`
	LocalPretrainModelPath string `json:"LOCAL_PRETRAIN_MODEL_PATH"`
	LocalOutputPath        string `json:"LOCAL_OUTPUT_PATH"`
	DataDownloadMethod     string `json:"DATA_DOWNLOAD_METHOD"`
	CodeNeedUnzip          string `json:"CODE_NEED_UNZIP"`
	DatasetNeedUnzip       string `json:"DATASET_NEED_UNZIP"`
	PretrainModelNeedUnzip string `json:"PRETRAIN_MODEL_NEED_UNZIP"`
}
type VolumeReq struct {
	Capacity  int    `json:"capacity"`
	Category  string `json:"category"`
	Ownership string `json:"ownership"`
	Uri       string `json:"uri"`
}

type CreateNotebookParams struct {
	JobName     string    `json:"name"`
	Description string    `json:"description"`
	ProfileID   string    `json:"profile_id"`
	Flavor      string    `json:"flavor"`
	Spec        Spec      `json:"spec"`
	Workspace   Workspace `json:"workspace"`
	Pool        Pool      `json:"pool"`
}

type Pool struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Workspace struct {
	ID string `json:"id"`
}

type Spec struct {
	Storage  Storage  `json:"storage"`
	AutoStop AutoStop `json:"auto_stop"`
}

type AutoStop struct {
	Enable   bool `json:"enable"`
	Duration int  `json:"duration"`
}

type Storage struct {
	Type     string   `json:"type"`
	Location Location `json:"location"`
}

type Location struct {
	Path string `json:"path"`
}

type NotebookResult struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type CreateNotebookResult struct {
	ErrorCode             string `json:"error_code"`
	ErrorMsg              string `json:"error_msg"`
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Status                string `json:"status"`
	CreationTimestamp     string `json:"creation_timestamp"`
	LatestUpdateTimestamp string `json:"latest_update_timestamp"`
	Profile               struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		DeType      string `json:"de_type"`
		FlavorType  string `json:"flavor_type"`
	} `json:"profile"`
	Flavor        string `json:"flavor"`
	FlavorDetails struct {
		Name          string `json:"name"`
		Status        string `json:"status"`
		QueuingNum    int    `json:"queuing_num"`
		QueueLeftTime int    `json:"queue_left_time"` //s
		Duration      int    `json:"duration"`        //auto_stop_time s
	} `json:"flavor_details"`
}

type GetNotebookResult struct {
	ErrorCode             string `json:"error_code"`
	ErrorMsg              string `json:"error_msg"`
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Status                string `json:"status"`
	CreationTimestamp     string `json:"creation_timestamp"`
	CreateTime            string
	LatestUpdateTimestamp string `json:"latest_update_timestamp"`
	LatestUpdateTime      string
	Profile               struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		DeType      string `json:"de_type"`
		FlavorType  string `json:"flavor_type"`
	} `json:"profile"`
	Flavor        string `json:"flavor"`
	FlavorDetails struct {
		Name          string `json:"name"`
		Status        string `json:"status"`
		QueuingNum    int    `json:"queuing_num"`
		QueueLeftTime int    `json:"queue_left_time"` //s
		Duration      int    `json:"duration"`        //auto_stop_time s
	} `json:"flavor_details"`
	QueuingInfo struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Flavor         string `json:"flavor"`
		DeType         string `json:"de_type"`
		Status         string `json:"status"`
		BeginTimestamp int    `json:"begin_timestamp"` //time of instance begin in queue
		BeginTime      string
		RemainTime     int `json:"remain_time"`   //remain time of instance
		EndTimestamp   int `json:"end_timestamp"` //
		EndTime        string
		Rank           int `json:"rank"` //rank of instance in queue
	} `json:"queuing_info"`
	Spec struct {
		Annotations struct {
			TargetDomain string `json:"target_domain"`
			Url          string `json:"url"`
		} `json:"annotations"`
	} `json:"spec"`
}

type GetNotebook2Result struct {
	ErrorCode        string `json:"error_code"`
	ErrorMsg         string `json:"error_msg"`
	FailReason       string `json:"fail_reason"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	Url              string `json:"url"`   //实例访问的URL
	Token            string `json:"token"` //notebook鉴权使用的token信息
	Flavor           string `json:"flavor"`
	CreateTime       string
	LatestUpdateTime string
	CreateAt         int64 `json:"create_at"` //实例创建的时间，UTC毫秒
	UpdateAt         int64 `json:"update_at"` //实例最后更新（不包括保活心跳）的时间，UTC毫秒
	Image            struct {
		Name          string `json:"name"`
		Status        string `json:"status"`
		QueuingNum    int    `json:"queuing_num"`
		QueueLeftTime int    `json:"queue_left_time"` //s
		Duration      int    `json:"duration"`        //auto_stop_time s
	} `json:"image"`
	Lease struct {
		CreateTime int64 `json:"create_at"` //实例创建的时间，UTC毫秒
		Duration   int64 `json:"duration"`  //实例运行时长，以创建时间为起点计算，即“创建时间+duration > 当前时刻”时，系统会自动停止实例
		UpdateTime int64 `json:"update_at"` //实例最后更新（不包括保活心跳）的时间，UTC毫秒
	} `json:"lease"` //实例自动停止的倒计时信息
	VolumeRes struct {
		Capacity  int    `json:"capacity"`
		Category  string `json:"category"`
		MountPath string `json:"mount_path"`
		Ownership string `json:"ownership"`
		Status    string `json:"status"`
	} `json:"volume"`
	ActionProgress []struct {
		Step        int    `json:"step"`
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"action_progress"`
}

type GetTokenParams struct {
	Auth Auth `json:"auth"`
}

type Auth struct {
	Identity Identity `json:"identity"`
	Scope    Scope    `json:"scope"`
}

type Scope struct {
	Project Project `json:"project"`
}

type Project struct {
	Name string `json:"name"`
}

type Identity struct {
	Methods  []string `json:"methods"`
	Password Password `json:"password"`
}

type Password struct {
	User NotebookUser `json:"user"`
}

type NotebookUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   Domain `json:"domain"`
}

type Domain struct {
	Name string `json:"name"`
}

const (
	ActionStart   = "start"
	ActionStop    = "stop"
	ActionRestart = "restart"
	ActionQueue   = "queue"
	ActionDequeue = "dequeue"
)

type NotebookAction struct {
	Action string `json:"action"`
}

type NotebookActionResult struct {
	ErrorCode     string `json:"error_code"`
	ErrorMsg      string `json:"error_msg"`
	CurrentStatus string `json:"current_status"`
	PreviousState string `json:"previous_state"`
	Status        string `json:"status"`
}

type NotebookGetJobTokenResult struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Token     string `json:"token"`
}

type NotebookDelResult struct {
	NotebookResult
	InstanceID string `json:"instance_id"`
}

type CreateUserImageTrainJobParams struct {
	JobName     string          `json:"job_name"`
	Description string          `json:"job_desc"`
	Config      UserImageConfig `json:"config"`
	WorkspaceID string          `json:"workspace_id"`
}

type UserImageConfig struct {
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"`  //训练作业需要的数据集OBS路径URL
	TrainUrl      string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl        string      `json:"log_url"`
	UserImageUrl  string      `json:"user_image_url"`
	UserCommand   string      `json:"user_command"`
	CreateVersion bool        `json:"create_version"`
	Flavor        Flavor      `json:"flavor"`
	PoolID        string      `json:"pool_id"`
	ShareAddr     string      `json:"nas_share_addr"`
	MountPath     string      `json:"nas_mount_path"`
	NasType       string      `json:"nas_type"`
}

type CreateTrainJobParams struct {
	JobName     string `json:"job_name"`
	Description string `json:"job_desc"`
	Config      Config `json:"config"`
	WorkspaceID string `json:"workspace_id"`
}

type Config struct {
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID      int64       `json:"engine_id"`
	TrainUrl      string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl        string      `json:"log_url"`
	//UserImageUrl 		string       	`json:"user_image_url"`
	//UserCommand  		string 			`json:"user_command"`
	CreateVersion bool   `json:"create_version"`
	Flavor        Flavor `json:"flavor"`
	PoolID        string `json:"pool_id"`
	ShareAddr     string `json:"nas_share_addr"`
	MountPath     string `json:"nas_mount_path"`
	NasType       string `json:"nas_type"`
}

type CreateInferenceJobParams struct {
	JobName     string    `json:"job_name"`
	Description string    `json:"job_desc"`
	InfConfig   InfConfig `json:"config"`
	WorkspaceID string    `json:"workspace_id"`
}

type CreateInfUserImageParams struct {
	JobName     string             `json:"job_name"`
	Description string             `json:"job_desc"`
	Config      InfUserImageConfig `json:"config"`
	WorkspaceID string             `json:"workspace_id"`
}

type InfConfig struct {
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID      int64       `json:"engine_id"`
	LogUrl        string      `json:"log_url"`
	CreateVersion bool        `json:"create_version"`
	Flavor        Flavor      `json:"flavor"`
	PoolID        string      `json:"pool_id"`
}

type InfUserImageConfig struct {
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID      int64       `json:"engine_id"`
	LogUrl        string      `json:"log_url"`
	CreateVersion bool        `json:"create_version"`
	Flavor        Flavor      `json:"flavor"`
	PoolID        string      `json:"pool_id"`
	UserImageUrl  string      `json:"user_image_url"`
	UserCommand   string      `json:"user_command"`
}

type CreateTrainJobVersionParams struct {
	Description string                `json:"job_desc"`
	Config      TrainJobVersionConfig `json:"config"`
}

type CreateTrainJobVersionUserImageParams struct {
	Description string                         `json:"job_desc"`
	Config      TrainJobVersionUserImageConfig `json:"config"`
}

type TrainJobVersionConfig struct {
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID      int64       `json:"engine_id"`
	TrainUrl      string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl        string      `json:"log_url"`
	Flavor        Flavor      `json:"flavor"`
	PoolID        string      `json:"pool_id"`
	PreVersionId  int64       `json:"pre_version_id"`
	ShareAddr     string      `json:"nas_share_addr"`
	MountPath     string      `json:"nas_mount_path"`
	NasType       string      `json:"nas_type"`
}

type TrainJobVersionUserImageConfig struct {
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"`  //训练作业需要的数据集OBS路径URL
	TrainUrl      string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl        string      `json:"log_url"`
	Flavor        Flavor      `json:"flavor"`
	PoolID        string      `json:"pool_id"`
	PreVersionId  int64       `json:"pre_version_id"`
	UserImageUrl  string      `json:"user_image_url"`
	UserCommand   string      `json:"user_command"`
	ShareAddr     string      `json:"nas_share_addr"`
	MountPath     string      `json:"nas_mount_path"`
	NasType       string      `json:"nas_type"`
}

type CreateConfigParams struct {
	ConfigName    string      `json:"config_name"`
	Description   string      `json:"config_desc"`
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID      int64       `json:"engine_id"`
	TrainUrl      string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl        string      `json:"log_url"`
	Flavor        Flavor      `json:"flavor"`
	PoolID        string      `json:"pool_id"`
	Volumes       []Volumes   `json:"volumes"`
}

type Parameter struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Parameters struct {
	Parameter []Parameter `json:"parameter"`
}
type Datasurl struct {
	DatasetUrl  string `json:"dataset_url"`
	DatasetName string `json:"dataset_name"`
}

type ModelUrls struct {
	ModelUrl  string `json:"model_url"`
	ModelName string `json:"model_name"`
}

type DatasetDownload struct {
	UUID                string `json:"uuid"`
	DatasetName         string `json:"dataset_name"`
	DatasetDownloadLink string `json:"dataset_download_link"`
	RepositoryLink      string `json:"repository_link"`
	IsDelete            bool   `json:"is_delete"`
	Size                int64  `json:"size"`
}

type DatasetRegistryDownload struct {
	UUID         string `json:"uuid"`
	DatasetName  string `json:"dataset_name"`
	DatasetAlias string `json:"dataset_alias"`
	OwnerName    string `json:"owner_name"`
	IsDelete     bool   `json:"is_delete"`
	Size         int64  `json:"size"`
}

type ModelDownload struct {
	ModelName      string `json:"model_name"`
	Name           string `json:"name"`
	DownloadLink   string `json:"download_link"`
	RepositoryLink string `json:"repository_link"`
	IsDelete       bool   `json:"is_delete"`
}

type Model4Show struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RepositoryLink string `json:"repository_link"`
	IsDelete       bool   `json:"is_delete"`
	//DownloadLink   string `json:"download_link"`
	Size      int64  `json:"size"`
	Alias     string `json:"alias"`
	OwnerName string `json:"owner_name"`
}

type DataSource struct {
	DatasetID      string `json:"dataset_id"`
	DatasetVersion string `json:"dataset_version"`
	Type           string `json:"type"`
	DataUrl        string `json:"data_url"`
}

type Volumes struct {
	Nfs      Nfs      `json:"nfs"`
	HostPath HostPath `json:"host_path"`
}

type Nfs struct {
	ID         string `json:"id"`
	SourcePath string `json:"src_path"`
	DestPath   string `json:"dest_path"`
	ReadOnly   bool   `json:"read_only"`
}

type HostPath struct {
	SourcePath string `json:"src_path"`
	DestPath   string `json:"dest_path"`
	ReadOnly   bool   `json:"read_only"`
}

type Flavor struct {
	Code string `json:"code"`
}

type CreateTrainJobResult struct {
	ErrorCode   string `json:"error_code"`
	ErrorMsg    string `json:"error_msg"`
	IsSuccess   bool   `json:"is_success"`
	JobName     string `json:"job_name"`
	JobID       int64  `json:"job_id"`
	Status      int    `json:"status"`
	CreateTime  int64  `json:"create_time"`
	VersionID   int64  `json:"version_id"`
	ResourceID  string `json:"resource_id"`
	VersionName string `json:"version_name"`
}

type CreateTrainJobConfigResult struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	IsSuccess bool   `json:"is_success"`
}

type GetResourceSpecsResult struct {
	ErrorCode      string  `json:"error_code"`
	ErrorMsg       string  `json:"error_msg"`
	IsSuccess      bool    `json:"is_success"`
	SpecTotalCount int     `json:"spec_total_count"`
	Specs          []Specs `json:"specs"`
}

type Specs struct {
	Core          string `json:"core"`
	Cpu           string `json:"cpu"`
	IsNoResource  bool   `json:"no_resource"`
	GpuType       string `json:"gpu_type"`
	SpecID        int64  `json:"spec_id"`
	GpuNum        int    `json:"gpu_num"`
	SpecCode      string `json:"spec_code"`
	Storage       string `json:"storage"`
	MaxNum        int    `json:"max_num"`
	UnitNum       int    `json:"unit_num"`
	InterfaceType int    `json:"interface_type"`
}

type GetConfigListResult struct {
	ErrorCode        string       `json:"error_code"`
	ErrorMsg         string       `json:"error_msg"`
	IsSuccess        bool         `json:"is_success"`
	ConfigTotalCount int          `json:"config_total_count"`
	ParaConfigs      []ParaConfig `json:"configs"`
}

type ParaConfig struct {
	ConfigName    string `json:"config_name"`
	ConfigDesc    string `json:"config_desc"`
	CreateTime    int64  `json:"create_time"`
	EngineType    int    `json:"engine_type"`
	EngineName    string `json:"engine_name"`
	EngineId      int64  `json:"engine_id"`
	EngineVersion string `json:"engine_version"`
	UserImageUrl  string `json:"user_image_url"`
	UserCommand   string `json:"user_command"`
	Result        GetConfigResult
}

type GetConfigResult struct {
	ErrorCode     string      `json:"error_code"`
	ErrorMsg      string      `json:"error_msg"`
	IsSuccess     bool        `json:"is_success"`
	ConfigName    string      `json:"config_name"`
	Description   string      `json:"config_desc"`
	WorkServerNum int         `json:"worker_server_num"`
	AppUrl        string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl   string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter     []Parameter `json:"parameter"`
	DataUrl       string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID      int64       `json:"engine_id"`
	TrainUrl      string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl        string      `json:"log_url"`

	Flavor Flavor `json:"flavor"`
	PoolID string `json:"pool_id"`
}

type ErrorResult struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_message"`
	IsSuccess bool   `json:"is_success"`
}

type GetTrainJobResult struct {
	IsSuccess        bool   `json:"is_success"`
	JobName          string `json:"job_name"`
	JobID            int64  `json:"job_id"`
	Description      string `json:"job_desc"`
	IntStatus        int    `json:"status"`
	Status           string
	LongCreateTime   int64 `json:"create_time"`
	CreateTime       string
	Duration         int64       `json:"duration"` //训练作业的运行时间，单位为毫秒
	TrainJobDuration string      //训练作业的运行时间，格式为hh:mm:ss
	VersionID        int64       `json:"version_id"`
	ResourceID       string      `json:"resource_id"`
	VersionName      string      `json:"version_name"`
	PreVersionID     int64       `json:"pre_version_id"`
	WorkServerNum    int         `json:"worker_server_num"`
	AppUrl           string      `json:"app_url"`       //训练作业的代码目录
	BootFileUrl      string      `json:"boot_file_url"` //训练作业的代码启动文件，需要在代码目录下
	Parameter        []Parameter `json:"parameter"`
	DataUrl          string      `json:"data_url"` //训练作业需要的数据集OBS路径URL
	EngineID         int64       `json:"engine_id"`
	EngineName       string      `json:"engine_name"`
	EngineVersion    string      `json:"engine_version"`
	TrainUrl         string      `json:"train_url"` //训练作业的输出文件OBS路径URL
	LogUrl           string      `json:"log_url"`
	Flavor           Flavor      `json:"flavor"`
	PoolID           string      `json:"pool_id"`
	PoolName         string      `json:"pool_name"`
	NasMountPath     string      `json:"nas_mount_path"`
	NasShareAddr     string      `json:"nas_share_addr"`
	DatasetName      string
	ModelMetricList  string `json:"model_metric_list"` //列表里包含f1_score，recall，precision，accuracy，若有的话
	StartTime        int64  `json:"start_time"`        //训练作业开始时间。
}

type GetTrainJobLogResult struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	IsSuccess bool   `json:"is_success"`
	Content   string `json:"content"`
	Lines     int    `json:"lines"`
	StartLine string `json:"start_line"`
	EndLine   string `json:"end_line"`
}

type GetTrainJobLogFileNamesResult struct {
	ErrorCode   string   `json:"error_code"`
	ErrorMsg    string   `json:"error_msg"`
	IsSuccess   bool     `json:"is_success"`
	LogFileList []string `json:"log_file_list"`
}

type TrainJobResult struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	IsSuccess bool   `json:"is_success"`
}

type LogFile struct {
	Name string
}

type JobList struct {
	JobName      string `json:"job_name"`
	JobID        int64  `json:"job_id"`
	VersionID    int64  `json:"version_id"`
	VersionCount int64  `json:"version_count"`
	Description  string `json:"job_desc"`
	IntStatus    int    `json:"status"`
}

type GetTrainJobListResult struct {
	ErrorResult
	JobTotalCount int       `json:"job_total_count"` //查询到的用户创建作业总数
	JobCountLimit int       `json:"job_count_limit"` //用户还可以创建训练作业的数量
	Quotas        int       `json:"quotas"`          //训练作业的运行数量上限
	JobList       []JobList `json:"jobs"`
}

type JobVersionList struct {
	VersionName string `json:"version_name"`
	VersionID   int64  `json:"version_id"`
	IntStatus   int    `json:"status"`
}

type GetTrainJobVersionListResult struct {
	ErrorResult
	JobID          int64            `json:"job_id"`
	JobName        string           `json:"job_name"`
	JobDesc        string           `json:"job_desc"`
	VersionCount   int64            `json:"version_count"`
	JobVersionList []JobVersionList `json:"versions"`
}

type NotebookList struct {
	JobName string `json:"name"`
	JobID   string `json:"id"`
	Status  string `json:"status"`
	Lease   struct {
		CreateTime int64 `json:"create_at"` //实例创建的时间，UTC毫秒
		Duration   int64 `json:"duration"`  //实例运行时长，以创建时间为起点计算，即“创建时间+duration > 当前时刻”时，系统会自动停止实例
		UpdateTime int64 `json:"update_at"` //实例最后更新（不包括保活心跳）的时间，UTC毫秒
	} `json:"lease"` //实例自动停止的倒计时信息
}

type GetNotebookListResult struct {
	TotalCount   int64          `json:"total"`   //总的记录数量
	CurrentPage  int            `json:"current"` //当前页数
	TotalPages   int            `json:"pages"`   //总的页数
	Size         int            `json:"size"`    //每一页的数量
	NotebookList []NotebookList `json:"data"`
}

// Grampus
type GrampusResult struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

type GrampusJobInfo struct {
	StartedAt      int64          `json:"startedAt"`
	RunSec         int64          `json:"runSec"`
	CompletedAt    int64          `json:"completedAt"`
	CreatedAt      int64          `json:"createdAt"`
	UpdatedAt      int64          `json:"updatedAt"`
	Desc           string         `json:"desc"`
	JobID          string         `json:"id"`
	Name           string         `json:"name"`
	Status         string         `json:"status"`
	DetailedStatus string         `json:"detailedStatus"`
	UserID         string         `json:"userId"`
	Tasks          []GrampusTasks `json:"tasks"`
}

type GrampusNotebookInfo struct {
	StartedAt      int64                 `json:"startedAt"`
	RunSec         int64                 `json:"runSec"`
	CompletedAt    int64                 `json:"completedAt"`
	CreatedAt      int64                 `json:"createdAt"`
	UpdatedAt      int64                 `json:"updatedAt"`
	Desc           string                `json:"desc"`
	JobID          string                `json:"id"`
	Name           string                `json:"name"`
	Status         string                `json:"status"`
	DetailedStatus string                `json:"detailedStatus"`
	UserID         string                `json:"userId"`
	Tasks          []GrampusNotebookTask `json:"tasks"`
}

const (
	GrampusNetAccess    = "1"
	GrampusNotNetAccess = "2"

	GrampusPoolTypePublic    = "1"
	GrampusPoolTypeExclusive = "2"

	GrampusTrue  = "1"
	GrampusFalse = "2"
)

type Center struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ResourceSpec []struct {
		ID              string `json:"id"`
		PoolType        string `json:"poolType"`
		Name            string `json:"name"`
		IsNetAccess     string `json:"isNetAccess"`
		IsSupportVisual string `json:"isSupportVisual"`
	} `json:"resourceSpec"`
}
type GrampusSpec struct {
	CreatedAt     int64    `json:"createdAt"`
	UpdatedAt     int64    `json:"updatedAt"`
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	ProcessorType string   `json:"processorType"`
	Centers       []Center `json:"centers"`
	SpecInfo      SpecInfo `json:"specInfo"`
}

type GrampusSpecPools struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CreatedAt       int64  `json:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt"`
	CardNum         int    `json:"cardNum"`
	CardType        string `json:"cardType"`
	CenterId        string `json:"centerId"`
	CenterName      string `json:"centerName"`
	ClusterType     string `json:"clusterType"`
	MarketState     int    `json:"marketState"`
	PoolType        string `json:"poolType"`
	ProcessorType   string `json:"processorType"`
	IsNetAccess     string `json:"isNetAccess"`
	IsSupportVisual string `json:"isSupportVisual"`
}

type GrampusAiCenter struct {
	AccDevices   []GrampusAccDevice      `json:"accDevices"`
	Id           string                  `json:"id"`
	Name         string                  `json:"name"`
	Resource     []GrampusCenterResource `json:"resource"`
	IsNetAccess  bool                    `json:"isNetAccess"`
	AccessTime   int64                   `json:"accessTime"`
	City         string                  `json:"city"`
	Province     string                  `json:"province"`
	ComputeScale int                     `json:"computeScale"`
}

type GrampusAccDevice struct {
	Kind  string `json:"kind"`  //加速卡类别, npu.huawei.com/NPU,nvidia.com/gpu,cambricon.com/mlu
	Model string `json:"model"` //加速卡型号
}

type GrampusCenterResource struct {
	Allocated string `json:"allocated"`
	Capacity  string `json:"capacity"`
	Name      string `json:"name"`
}

type SpecInfo struct {
	AccDeviceKind   string `json:"accDeviceKind"`
	AccDeviceMemory string `json:"accDeviceMemory"`
	AccDeviceModel  string `json:"accDeviceModel"`
	AccDeviceNum    int    `json:"accDeviceNum"`
	CpuCoreNum      int    `json:"cpuCoreNum"`
	MemorySize      string `json:"memorySize"`
}

type GrampusJobEvents struct {
	Message   string `json:"message"`
	Name      string `json:"name"`
	Reason    string `json:"reason"`
	Timestamp string `json:"timestamp"`
}

type GetGrampusResourceSpecsResult struct {
	GrampusResult
	Infos []GrampusSpec `json:"resourceSpecs"`
}

type GetGrampusResourceSpecPoolsResult struct {
	GrampusResult
	Infos []GrampusSpecPools `json:"resourceSpecPools"`
}

type GetGrampusAiCentersResult struct {
	GrampusResult
	Infos []GrampusAiCenter `json:"aiCenterInfos"`
}

type GrampusResourceQueue struct {
	QueueCode           string
	QueueName           string
	QueueType           string
	AiCenterCode        string
	AiCenterName        string
	ComputeResource     string
	AccCardType         string
	HasInternet         int //0 unknown;1 no internet;2 has internet
	EnableVisualization bool
	CardsTotalNum       int
}

type GrampusImage struct {
	CreatedAt     int64           `json:"createdAt"`
	UpdatedAt     int64           `json:"updatedAt"`
	ID            string          `json:"id"`
	CenterImageId string          `json:"centerImageId"` // 这个镜像是用来虎鲸去中心查得，目前测试出现对应不上的问题先兼容20251201
	Name          string          `json:"name"`
	ProcessorType string          `json:"processorType"`
	TrainType     string          `json:"trainType"`
	AICenterImage []AiCenterImage `json:"aiCenterImages"`
}

type GetGrampusImagesResult struct {
	GrampusResult
	TotalSize int            `json:"totalSize"`
	Infos     []GrampusImage `json:"images"`
}

type CreateGrampusJobResponse struct {
	GrampusResult
	JobInfo GrampusJobInfo `json:"otJob"`
}

type GetGrampusJobResponse struct {
	GrampusResult
	JobInfo         GrampusJobInfo `json:"otJob"`
	ExitDiagnostics string         `json:"exitDiagnostics"`
}

type GetGrampusJobListResponse struct {
	GrampusResult
	JobInfos  []GrampusJobInfo `json:"otJobs"`
	TotalSize int              `json:"totalSize"`
}

type GrampusNotebookResponse struct {
	GrampusResult
	JobInfo         GrampusNotebookInfo `json:"otJob"`
	ExitDiagnostics string              `json:"exitDiagnostics"`
}

type GrampusNotebookRestartResponse struct {
	GrampusResult
	NewId  string `json:"newId"`
	Status string `json:"status"`
}

type GrampusStopJobResponse struct {
	GrampusResult
	StoppedAt int64  `json:"stoppedAt"`
	ID        string `json:"id"`
	Status    string `json:"status"`
}

type GrampusModelMigrateInfoResponse struct {
	GrampusResult
	DestBucket     string `json:"destBucket"`
	DestEndpoint   string `json:"destEndpoint"`
	DestObjectKey  string `json:"destObjectKey"`
	DestProxy      string `json:"destProxy"`
	FailedReason   string `json:"failedReason"`
	SrcBucket      string `json:"srcBucket"`
	SrcEndpoint    string `json:"srcEndpoint"`
	SrcObjectKey   string `json:"srcObjectKey"`
	Status         int    `json:"status"` //0:初始化 1：成功 2：失败 3：调度中
	DataId         string `json:"dataId"`
	MigrateTaskId  int64  `json:"migrateTaskId"`
	MigrateService string `json:"migrateService"` //海胆1.0 = "urchin1.0" 海胆2.0= SystemUrchin2 = "urchin2.0"
}

type GetGrampusJobEventsResponse struct {
	GrampusResult
	JobEvents []GrampusJobEvents `json:"jobEvents"`
	TotalSize int                `json:"totalSize"`
}
type GetGrampusDebugJobEventsResponse struct {
	GrampusResult
	NotebookEvents []GrampusJobEvents `json:"notebookEvents"`
	TotalSize      int                `json:"totalSize"`
}

type GrampusTasks struct {
	Command             string                 `json:"command"`
	Name                string                 `json:"name"`
	ImageId             string                 `json:"imageId"`
	ResourceSpecId      string                 `json:"resourceSpecId"`
	ImageUrl            string                 `json:"imageUrl"`
	CenterID            []string               `json:"centerID"`
	CenterName          []string               `json:"centerName"`
	ReplicaNum          int                    `json:"replicaNum"`
	Datasets            []GrampusDataset       `json:"datasets"`
	Models              []GrampusDataset       `json:"models"`
	Code                GrampusDataset         `json:"code"`
	BootFile            string                 `json:"bootFile"`
	OutPut              GrampusDataset         `json:"output"`
	WorkServerNumber    int                    `json:"nodeCount"`
	RunParams           map[string]interface{} `json:"runParams"`
	TensorboardEndpoint string                 `json:"tensorboardEndpoint"`
	PoolId              string                 `json:"poolId"`
	EnvVariables        map[string]interface{} `json:"envVariables"`
}
type GrampusNotebookTask struct {
	AutoStopDuration         int64                  `json:"autoStopDuration"`
	Name                     string                 `json:"name"`
	Capacity                 int                    `json:"capacity"`
	CenterID                 []string               `json:"centerID"`
	CenterName               []string               `json:"centerName"`
	PoolId                   string                 `json:"poolId"`
	Code                     GrampusDataset         `json:"code"`
	Datasets                 []GrampusDataset       `json:"datasets"`
	PreTrainModel            []GrampusDataset       `json:"models"`
	OutPut                   GrampusDataset         `json:"output"`
	CodeUrl                  string                 `json:"codeUrl"`
	DataUrl                  string                 `json:"dataUrl"`
	ImageId                  string                 `json:"imageId"`
	ImageUrl                 string                 `json:"imageUrl"`
	ResourceSpecId           string                 `json:"resourceSpecId"`
	Token                    string                 `json:"token"`
	PToken                   string                 `json:"ptoken"`
	Url                      string                 `json:"url"`
	Status                   string                 `json:"status"`
	Command                  string                 `json:"command"`
	EnvVariables             map[string]interface{} `json:"envVariables"`
	NoActAutoShutDownTimeout int                    `json:"noActAutoShutDownTimeout"`
	DominUrl                 []string               `json:"selfDomains"`
	AutoSave                 bool                   `json:"autoSave"`
}

type GrampusInferenceTask struct {
	AutoStopDuration int64                  `json:"autoStopDuration"`
	Name             string                 `json:"name"`
	Capacity         int                    `json:"capacity"`
	CenterID         []string               `json:"centerID"`
	CenterName       []string               `json:"centerName"`
	PoolId           string                 `json:"poolId"`
	Code             GrampusDataset         `json:"code"`
	Datasets         []GrampusDataset       `json:"datasets"`
	PreTrainModel    []GrampusDataset       `json:"models"`
	OutPut           GrampusDataset         `json:"output"`
	CodeUrl          string                 `json:"codeUrl"`
	DataUrl          string                 `json:"dataUrl"`
	ImageId          string                 `json:"imageId"`
	ImageUrl         string                 `json:"imageUrl"`
	ResourceSpecId   string                 `json:"resourceSpecId"`
	Token            string                 `json:"token"`
	Url              string                 `json:"url"`
	Status           string                 `json:"status"`
	Command          string                 `json:"command"`
	EnvVariables     map[string]interface{} `json:"envVariables"`
	BootFile         string                 `json:"bootFile"`
	Endpoints        []*SelfEndPoint        `json:"endpoints"`
	AutoSave         bool                   `json:"autoSave"`
}

type SelfEndPoint struct {
	Port     int64  `json:"port"`
	EndPoint string `json:"endPoint"`
}

type GrampusDataset struct {
	Name              string `json:"name"`
	Bucket            string `json:"bucket"`
	EndPoint          string `json:"endPoint"`
	ObjectKey         string `json:"objectKey"`
	ContainerPath     string `json:"containerPath"`
	ReadOnly          bool   `json:"readOnly"`
	GetBackEndpoint   string `json:"getBackEndpoint"`
	Size              int64  `json:"size"`
	IsOverwrite       bool   `json:"isOverwrite"`
	IsNeedUnzip       bool   `json:"isNeedUnzip"`
	IsNeedTensorboard bool   `json:"isNeedTensorboard"`
	Id                string `json:"id"`
	SizeLimit         int    `json:"sizeLimit"` // 目录大小限制，单位GB，0表示不限制，仅启智混合集群生效
}

type CreateGrampusJobRequest struct {
	Name  string         `json:"name"`
	Tasks []GrampusTasks `json:"tasks"`
}

type CreateGrampusNotebookRequest struct {
	Name  string                `json:"name"`
	Tasks []GrampusNotebookTask `json:"tasks"`
}

type CreateGrampusInferenceRequest struct {
	Name  string                 `json:"name"`
	Tasks []GrampusInferenceTask `json:"tasks"`
}

type GrampusModelAppTask struct {
	CenterID       []string `json:"centerId"`
	ModelId        string   `json:"modelId"`
	Name           string   `json:"name"`
	ResourceSpecId string   `json:"resourceSpecId"`
}

type CreateGrampusModelAppRequest struct {
	Name  string                `json:"name"`
	Tasks []GrampusModelAppTask `json:"tasks"`
}

type GetTrainJobMetricStatisticResult struct {
	TrainJobResult
	Interval    int       `json:"interval"` //查询的时间间隔，单位为分钟
	MetricsInfo []Metrics `json:"metrics"`  //监控详情
}

type Metrics struct {
	Metric string   `json:"metric"` //监控指标项
	Value  []string `json:"value"`  //获取的监控值的序列，元素为String类型
}

type NewModelArtsMetricStatisticResult struct {
	MetricsInfo []NewModelArtsMetrics `json:"metrics"` //监控详情
	Step        int64                 `json:"step"`
}

type GrampusMetricStatisticResult struct {
	GrampusResult
	MetricsInfo []NewModelArtsMetrics `json:"metrics"` //监控详情
}

type NewModelArtsMetrics struct {
	Metric string    `json:"metric"` //监控指标项
	Value  []float32 `json:"value"`  //获取的监控值的序列，元素为float类型
}

type GrampusDeleteJobResponse struct {
	GrampusResult
	Info string `json:"info"`
}

type GrampusServiceModels struct {
	GrampusResult
	ServiceModels []ServiceModel `json:"serviceModels"`
}
type ServiceModel struct {
	CenterServiceModels []CenterInfo `json:"centerServiceModels"`
	CreatedAt           int64        `json:"createdAt"`
	ID                  string       `json:"id"`
	Name                string       `json:"name"`
	ExperienceType      string       `json:"type"`
	CenterIds           []string     `json:"centerId"`
	UpdatedAt           int64        `json:"updatedAt"`
}
type CenterInfo struct {
	AICenterId string `json:"aiCenterId"`
}

// CountCloudBrainsByUserId 统计云脑当前用户的任务数量
func CountCloudBrainsByUserId(userId int64) (int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	cond := builder.NewCond().And(builder.Eq{"cloudbrain.user_id": userId})
	// 训练版本问题兼容
	cond = cond.And(builder.Or(builder.And(builder.Eq{"cloudbrain.is_latest_version": "1"}, builder.Eq{"cloudbrain.job_type": "TRAIN"}), builder.Neq{"cloudbrain.job_type": "TRAIN"}))

	count, err := sess.Where(cond).Count(new(Cloudbrain))
	return count, err
}

// CountCloudBrainsByJobStatus 统计云脑当前用户的指定状态中的数量
func CountCloudBrainsByJobStatus(userId int64, jobStatus CloudbrainStatus) (int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	cond := builder.NewCond().And(builder.Eq{"cloudbrain.user_id": userId}).And(builder.Eq{"cloudbrain.status": jobStatus})
	// 训练版本问题兼容
	cond = cond.And(builder.Or(builder.And(builder.Eq{"cloudbrain.is_latest_version": "1"}, builder.Eq{"cloudbrain.job_type": "TRAIN"}), builder.Neq{"cloudbrain.job_type": "TRAIN"}))

	count, err := sess.Where(cond).Count(new(Cloudbrain))
	return count, err
}

// CountAccCardTotalDuration 计算用户累计卡时
func CountAccCardTotalDuration(userID int64) (int64, error) {
	var totalSum int64

	sql := `
           SELECT
        SUM(
            COALESCE(cloudbrain.duration *
				CASE
				WHEN cloudbrain.work_server_number = 0 THEN 1
				ELSE COALESCE(cloudbrain.work_server_number, 1)
				END *
				COALESCE(cloudbrain_spec.acc_cards_num, 1), 0)
   ) as cardDurationSum
    FROM cloudbrain
    LEFT JOIN cloudbrain_spec
        ON cloudbrain.id = cloudbrain_spec.cloudbrain_id where cloudbrain.user_id = ?`

	_, err := x.SQL(sql, userID).Unscoped().Get(&totalSum)
	return totalSum, err
}

func Cloudbrains(opts *CloudbrainsOptions) ([]*CloudbrainInfo, int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()
	if opts.RepoID > 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.repo_id": opts.RepoID},
		)
	}

	if opts.UserID > 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.user_id": opts.UserID},
		)
	}

	if (opts.JobID) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.job_id": opts.JobID},
		)
	}
	if (opts.ComputeResource) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.compute_resource": opts.ComputeResource},
		)
	}

	if (opts.Type) >= 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": opts.Type},
		)
	}
	if opts.FilterCleared {
		cond = cond.And(
			builder.Eq{"cloudbrain.cleared": false},
		)
	}

	if len(opts.JobTypes) > 0 {
		if opts.JobTypeNot {
			cond = cond.And(
				builder.NotIn("cloudbrain.job_type", opts.JobTypes),
			)
		} else {
			cond = cond.And(
				builder.In("cloudbrain.job_type", opts.JobTypes),
			)
		}
	}
	if len(opts.ExcludeJobTypes) > 0 {
		cond = cond.And(
			builder.NotIn("cloudbrain.job_type", opts.ExcludeJobTypes),
		)
	}
	if (opts.AiCenter) != "" {
		if opts.AiCenter == AICenterOfCloudBrainOne {
			cond = cond.And(
				builder.Eq{"cloudbrain.type": TypeCloudBrainOne},
			)
		} else if opts.AiCenter == AICenterOfCloudBrainTwo {
			cond = cond.And(
				builder.Eq{"cloudbrain.type": TypeCloudBrainTwo},
			)
		} else if opts.AiCenter == AICenterOfChengdu {
			cond = cond.And(
				builder.Eq{"cloudbrain.type": TypeCDCenter},
			)
		} else {
			cond = cond.And(
				builder.Like{"cloudbrain.ai_center", opts.AiCenter},
			)
		}
	}
	if opts.Cluster == "resource_cluster_openi" {
		cond = cond.And(
			builder.Or(builder.Eq{"cloudbrain.type": TypeCloudBrainOne}, builder.Eq{"cloudbrain.type": TypeCloudBrainTwo}, builder.Eq{"cloudbrain.type": TypeCDCenter}),
		)
	} else if opts.Cluster == "resource_cluster_c2net" {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": TypeC2Net},
		)
	} else if opts.Cluster == "resource_cluster_c2net" {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": TypeC2Net},
		)
	} else if opts.Cluster == IFLYTEKCluster {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": TypeIFLYTek},
		)
	}

	if (opts.IsLatestVersion) != "" {
		cond = cond.And(builder.Or(builder.And(builder.Eq{"cloudbrain.is_latest_version": opts.IsLatestVersion}, builder.Eq{"cloudbrain.job_type": "TRAIN"}), builder.Neq{"cloudbrain.job_type": "TRAIN"}))
	}

	if (opts.AppName) != "" {
		cond = cond.And(builder.Eq{"cloudbrain.app_name": opts.AppName})
	}

	if len(opts.CloudbrainIDs) > 0 {
		cond = cond.And(builder.In("cloudbrain.id", opts.CloudbrainIDs))
	}
	if opts.IsAimRequired {
		cond = cond.And(builder.Eq{"cloudbrain.aim_required": opts.IsAimRequired})
	}

	if len(opts.JobStatus) > 0 {
		if opts.JobStatusNot {
			cond = cond.And(
				builder.NotIn("cloudbrain.status", opts.JobStatus),
			)
		} else {
			cond = cond.And(
				builder.In("cloudbrain.status", opts.JobStatus),
			)
		}
	}
	if len(opts.RepoIDList) > 0 {
		cond = cond.And(
			builder.In("cloudbrain.repo_id", opts.RepoIDList),
		)

	}

	if opts.BeginTimeUnix > 0 && opts.EndTimeUnix > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain.created_unix": opts.BeginTimeUnix}, builder.Lte{"cloudbrain.created_unix": opts.EndTimeUnix}),
		)
	}

	var count int64
	var err error
	condition := "cloudbrain.user_id = `user`.id"
	if len(opts.Keyword) == 0 {
		count, err = sess.Where(cond).Count(new(Cloudbrain))
	} else {
		lowerKeyWord := strings.ToLower(opts.Keyword)

		cond = cond.And(builder.Or(builder.Like{"LOWER(cloudbrain.job_name)", lowerKeyWord}, builder.Like{"LOWER(cloudbrain.display_job_name)", lowerKeyWord}, builder.Like{"`user`.lower_name", lowerKeyWord}))
		count, err = sess.Table(&Cloudbrain{}).Where(cond).
			Join("left", "`user`", condition).Count(new(CloudbrainInfo))

	}

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}

	if opts.OrderBy == "repo" {
		sess.OrderBy("cloudbrain.repo_id DESC,cloudbrain.id DESC")
	} else {
		sess.OrderBy("cloudbrain.created_unix DESC,cloudbrain.id DESC")
	}

	cloudbrains := make([]*CloudbrainInfo, 0, setting.UI.IssuePagingNum)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Join("left", "`user`", condition).
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	if opts.NeedRepoInfo {
		var ids []int64
		for _, task := range cloudbrains {
			ids = append(ids, task.RepoID)
		}
		repositoryMap, err := GetRepositoriesMapByIDs(ids)
		if err == nil {
			for _, task := range cloudbrains {
				if task.RepoID != 0 {
					task.Repo = repositoryMap[task.RepoID]
				}
			}
		}

	}

	return cloudbrains, count, nil
}

func QueryModelTrainJobVersionList(jobId string) ([]*Cloudbrain, int, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()

	cond = cond.And(
		builder.Eq{"cloudbrain.job_id": jobId},
	)
	// cond = cond.And(
	// 	builder.In("cloudbrain.Status", "COMPLETED", "SUCCEEDED"),
	// 	//builder.Eq{"cloudbrain.Status": "COMPLETED"},
	// )

	sess.OrderBy("cloudbrain.created_unix DESC")
	cloudbrains := make([]*Cloudbrain, 0)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	return cloudbrains, int(len(cloudbrains)), nil
}

func QueryModelTrainJobList(repoId int64) ([]*Cloudbrain, int, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"repo_id": repoId},
	)
	// cond = cond.And(
	// 	builder.In("Status", "COMPLETED", "SUCCEEDED"),
	// )
	cond = cond.And(
		builder.Eq{"job_type": "TRAIN"},
	)
	// cond = cond.And(
	// 	builder.In("type", 0, 1),
	// )
	cond = cond.And(
		builder.In("compute_resource", NPUResource, GPUResource, GCUResource, ILUVATAR),
	)

	cloudbrains := make([]*Cloudbrain, 0)
	if err := sess.Select("*").Table(&Cloudbrain{}).Where(cond).OrderBy("created_unix DESC").
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	keys := make(map[string]string)
	uniqueElements := make([]*Cloudbrain, 0)
	for _, entry := range cloudbrains {
		if _, value := keys[entry.JobID]; !value {
			keys[entry.JobID] = entry.DisplayJobName
			uniqueElements = append(uniqueElements, entry)
		}
	}

	return uniqueElements, int(len(uniqueElements)), nil
}
func CountByRawSql(sql string) (int64, error) {
	return x.SQL(sql).Count()
}
func QueryByRawSql(sql string) ([]map[string]string, error) {
	return x.QueryString(sql)
}

func CloudbrainsVersionList(opts *CloudbrainsOptions) ([]*CloudbrainInfo, int, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()
	if opts.RepoID > 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.repo_id": opts.RepoID},
		)
	}

	if opts.UserID > 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.user_id": opts.UserID},
		)
	}

	if (opts.Type) >= 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": opts.Type},
		)
	}

	if (opts.JobID) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.job_id": opts.JobID},
		)
	}

	if len(opts.JobTypes) > 0 {
		cond = cond.And(
			builder.In("cloudbrain.job_type", opts.JobTypes),
		)
	}

	if len(opts.CloudbrainIDs) > 0 {
		cond = cond.And(builder.In("cloudbrain.id", opts.CloudbrainIDs))
	}

	count, err := sess.Where(cond).Count(new(Cloudbrain))
	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}

	sess.OrderBy("cloudbrain.created_unix DESC")
	cloudbrains := make([]*CloudbrainInfo, 0, setting.UI.IssuePagingNum)
	if err := sess.Table(&Cloudbrain{}).Where(cond).
		Join("left", "`user`", "cloudbrain.user_id = `user`.id").
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	return cloudbrains, int(count), nil
}

func GetCloudbrainEarlyVersionList(task *Cloudbrain) ([]*Cloudbrain, error) {
	if task.JobID == "" {
		return []*Cloudbrain{}, nil
	}
	cloudbrains := make([]*Cloudbrain, 0)
	if err := x.Where(builder.NewCond().
		And(builder.Eq{"cloudbrain.repo_id": task.RepoID}).
		And(builder.Eq{"cloudbrain.type": task.Type}).
		And(builder.Eq{"cloudbrain.job_id": task.JobID}).
		And(builder.Eq{"cloudbrain.job_type": task.JobType}).
		And(builder.Lt{"cloudbrain.created_unix": task.CreatedUnix})).
		OrderBy("cloudbrain.created_unix DESC").
		Find(&cloudbrains); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}

	return cloudbrains, nil
}

func CreateCloudbrain(cloudbrain *Cloudbrain) (err error) {
	session := x.NewSession()
	defer session.Close()

	err = session.Begin()
	cloudbrain.TrainJobDuration = DURATION_STR_ZERO
	if _, err = session.NoAutoTime().InsertOne(cloudbrain); err != nil {
		session.Rollback()
		return err
	}

	if cloudbrain.Spec != nil {
		if _, err = session.Insert(NewCloudBrainSpec(cloudbrain.ID, *cloudbrain.Spec)); err != nil {
			session.Rollback()
			return err
		}
	}
	if cloudbrain.Config != nil {
		cloudbrain.Config.CloudbrainID = cloudbrain.ID
		if _, err = session.Insert(cloudbrain.Config); err != nil {
			session.Rollback()
			return err
		}
	}
	if cloudbrain.TemplateID != "" {
		if _, err = session.Exec("update ai_task_template set use_count = use_count + 1 where id = ?", cloudbrain.TemplateID); err != nil {
			session.Rollback()
			return err
		}
	}
	session.Commit()
	go updateReferenceCount(cloudbrain)
	OperateRepoAITaskNum(cloudbrain.RepoID, 1)
	return nil
}

func updateReferenceCount(cloudbrain *Cloudbrain) {
	increaseDatasetRegistryUseCount(cloudbrain.Uuid)
	increaseImageUseCount(cloudbrain.Image, cloudbrain.ImageID)
	for _, id := range cloudbrain.GetModelIdArray() {
		increaseModelReference(id)
	}
}

func increaseImageUseCount(imageUrl, imageId string) {
	if imageUrl != "" {
		res, _ := x.Exec("UPDATE `image` SET use_count=use_count+1 WHERE place=?", imageUrl)
		affected, _ := res.RowsAffected()
		if affected > 0 {
			return
		}
	}
	if imageId != "" {
		x.Exec("UPDATE `image` SET use_count=use_count+1 WHERE image_id=?", imageId)
	}
}

func increaseModelReference(modelId string) {
	if modelId != "" {
		log.Info("increase model count.")
		if _, err := x.Exec("UPDATE `ai_model_manage` SET reference_count = reference_count + 1 WHERE id = ?", modelId); err != nil {
			log.Info("err=" + err.Error())
		}
	}
}

func getRepoCloudBrain(cb *Cloudbrain) (*Cloudbrain, error) {
	has, err := x.Get(cb)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrJobNotExist{}
	}
	return cb, nil
}

func getRepoCloudBrainWithDeleted(cb *Cloudbrain) (*Cloudbrain, error) {
	has, err := x.Unscoped().Get(cb)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrJobNotExist{}
	}
	return cb, nil
}

func GetRepoCloudBrainByJobID(repoID int64, jobID string) (*Cloudbrain, error) {
	cb := &Cloudbrain{JobID: jobID, RepoID: repoID}
	return getRepoCloudBrain(cb)
}

func GetCloudbrainByJobID(jobID string) (*Cloudbrain, error) {
	cb := &Cloudbrain{JobID: jobID}
	return getRepoCloudBrain(cb)
}

func GetCloudbrainListByJobID(jobID string) ([]*Cloudbrain, error) {
	r := make([]*Cloudbrain, 0)
	if err := x.Where("job_id = ?", jobID).OrderBy("id desc").Find(&r); err != nil {
		return nil, err
	}
	return r, nil

}

func GetNewestCloudbrainByJobId(jobID string) (*Cloudbrain, error) {
	r := &Cloudbrain{}
	if has, err := x.Where("job_id = ?", jobID).OrderBy("id desc").Limit(1).Get(r); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil

}

func GetCloudbrainByJobIDWithDeleted(jobID string) (*Cloudbrain, error) {
	cb := &Cloudbrain{JobID: jobID}
	return getRepoCloudBrainWithDeleted(cb)
}

func GetCloudbrainByID(id string) (*Cloudbrain, error) {
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	cb := &Cloudbrain{ID: idInt64}
	return getRepoCloudBrain(cb)
}

func GetCloudbrainByCloudbrainID(id int64) (*Cloudbrain, error) {
	cb := &Cloudbrain{ID: id}
	return getRepoCloudBrain(cb)
}

func IsCloudbrainExistByJobName(jobName string) (bool, error) {
	return x.Unscoped().Exist(&Cloudbrain{
		JobName: jobName,
	})
}

func GetCloudbrainByIDWithDeleted(id string) (*Cloudbrain, error) {
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	cb := &Cloudbrain{ID: idInt64}
	return getRepoCloudBrainWithDeleted(cb)
}

func GetCloudbrainByJobIDAndVersionName(jobID string, versionName string) (*Cloudbrain, error) {
	cb := &Cloudbrain{JobID: jobID, VersionName: versionName}
	return getRepoCloudBrain(cb)
}

func GetCloudbrainByJobIDAndIsLatestVersion(jobID string, isLatestVersion string) (*Cloudbrain, error) {
	cb := &Cloudbrain{JobID: jobID, IsLatestVersion: isLatestVersion}
	return getRepoCloudBrain(cb)
}

func GetCloudbrainsNeededStopByUserID(userID int64) ([]*Cloudbrain, error) {
	cloudBrains := make([]*Cloudbrain, 0)
	err := x.Cols("job_id", "status", "type", "job_type", "version_id", "start_time").Where("user_id=? AND status !=?", userID, string(JobStopped)).Find(&cloudBrains)
	return cloudBrains, err
}

func GetModelartsReDebugTaskByJobId(jobID string) ([]*Cloudbrain, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"cloudbrain.job_id": jobID},
	)
	sess.OrderBy("cloudbrain.created_unix ASC limit 1")
	cloudbrains := make([]*Cloudbrain, 0, 10)
	if err := sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
		Find(&cloudbrains); err != nil {
		log.Info("find error.")
	}
	return cloudbrains, nil
}

func GetCloudbrainsNeededStopByRepoID(repoID int64) ([]*Cloudbrain, error) {
	cloudBrains := make([]*Cloudbrain, 0)
	err := x.Cols("job_id", "status", "type", "job_type", "version_id", "start_time").Where("repo_id=? AND status !=?", repoID, string(JobStopped)).Find(&cloudBrains)
	return cloudBrains, err
}

func GetCloudbrainsNeededDeleteByRepoID(repoID int64) ([]*Cloudbrain, error) {
	cloudBrains := make([]*Cloudbrain, 0)
	err := x.Where("repo_id=?", repoID).Find(&cloudBrains)
	return cloudBrains, err
}

func GetCloudbrainsByDisplayJobName(repoID int64, jobType string, displayJobName string) ([]*Cloudbrain, error) {
	cloudBrains := make([]*Cloudbrain, 0)
	err := x.Cols("job_id", "job_name", "repo_id", "user_id", "job_type", "display_job_name").Where("repo_id=? AND job_type =? AND lower(display_job_name) = lower(?)", repoID, jobType, displayJobName).Find(&cloudBrains)
	return cloudBrains, err
}

func SetCloudbrainStatusByJobID(jobID string, status CloudbrainStatus) (err error) {
	cb := &Cloudbrain{JobID: jobID, Status: string(status)}
	_, err = x.Cols("status").Where("cloudbrain.job_id=?", jobID).Update(cb)
	return
}

func SetTrainJobStatusByJobID(jobID string, status string, duration int64, trainjobduration string) (err error) {
	cb := &Cloudbrain{JobID: jobID, Status: string(status), Duration: duration, TrainJobDuration: trainjobduration}
	_, err = x.Cols("status", "duration", "train_job_duration").Where("cloudbrain.job_id=?", jobID).Update(cb)
	return
}

func SetVersionCountAndLatestVersion(jobID string, versionName string, versionCount int, isLatestVersion string, totalVersionCount int) (err error) {
	cb := &Cloudbrain{JobID: jobID, VersionName: versionName, VersionCount: versionCount, IsLatestVersion: isLatestVersion, TotalVersionCount: totalVersionCount}
	_, err = x.Cols("version_Count", "is_latest_version", "total_version_count").Where("cloudbrain.job_id=? AND cloudbrain.version_name=?", jobID, versionName).Update(cb)
	return
}

func UpdateJob(job *Cloudbrain) error {
	return updateJob(x, job)
}

func UpdateJobStatus(id int64, newStatus, oldStatus string) (int64, error) {
	return x.Where("id = ? and status = ?", id, oldStatus).Cols("status").Update(&Cloudbrain{Status: newStatus})
}

func UpdateJobNewStatus(id int64, newStatus string) (int64, error) {
	return x.Where("id = ?", id).Cols("status").Update(&Cloudbrain{Status: newStatus})
}

func UpdateJobDescription(id int64, description string) (int64, error) {
	return x.Where("id = ?", id).Cols("description").Update(&Cloudbrain{Description: description})
}

func UpdateJobDurationWithDeleted(job *Cloudbrain) error {
	_, err := x.Exec("update cloudbrain set start_time=?, end_time=?,train_job_duration=?,duration=? where id=?", job.StartTime, job.EndTime, job.TrainJobDuration, job.Duration, job.ID)
	return err
}

func updateJob(e Engine, job *Cloudbrain) error {
	_, err := e.ID(job.ID).AllCols().Update(job)
	return err
}

func UpdateTrainJobVersion(job *Cloudbrain) error {
	return updateJobTrainVersion(x, job)
}

func updateJobTrainVersion(e Engine, job *Cloudbrain) error {
	var sess *xorm.Session
	sess = e.Where("job_id = ? AND version_name=?", job.JobID, job.VersionName)
	_, err := sess.Cols("status", "train_job_duration", "duration", "start_time", "end_time", "created_unix", "ai_center").Update(job)
	return err
}

func DeleteJob(job *Cloudbrain) error {
	return deleteJob(x, job)
}

func deleteJob(e Engine, job *Cloudbrain) error {
	_, err := e.ID(job.ID).Delete(job)
	if err == nil {
		go updateAITaskNumWhenDeleteJob(job)
	}
	return err
}

func updateAITaskNumWhenDeleteJob(job *Cloudbrain) {
	repoId := job.RepoID
	if repoId == 0 {
		t := &Cloudbrain{}
		_, tempErr := x.ID(job.ID).Unscoped().Get(t)
		if tempErr != nil {
			log.Error("updateAITaskNumWhenDeleteJob error.%v", tempErr)
			return
		}
		repoId = t.RepoID
	}

	if repoId > 0 {
		go OperateRepoAITaskNum(repoId, -1)
	}
}

func GetCloudbrainByName(jobName string) (*Cloudbrain, error) {
	cb := &Cloudbrain{JobName: jobName}
	return getRepoCloudBrain(cb)
}
func GetWaitOrRunFileNotebookByRepo(repoId int64, userID int64, computeResource string, image string) (*Cloudbrain, error) {
	cloudBrain := new(Cloudbrain)
	if computeResource != "" {
		if image != "" {
			has, err := x.In("status", JobWaiting, JobRunning, ModelArtsCreateQueue, ModelArtsCreating, ModelArtsStarting,
				ModelArtsReadyToStart, ModelArtsResizing, ModelArtsStartQueuing, ModelArtsRunning, ModelArtsDeleting, ModelArtsRestarting).Where("repo_id=? and user_id=? and type=? and boot_file!='' and job_type=? and compute_resource=? and image=? ", repoId, userID, TypeC2Net, string(JobTypeDebug), computeResource, image).Get(cloudBrain)
			if has {
				return cloudBrain, err
			}
			return nil, err
		} else {
			has, err := x.In("status", JobWaiting, JobRunning, ModelArtsCreateQueue, ModelArtsCreating, ModelArtsStarting,
				ModelArtsReadyToStart, ModelArtsResizing, ModelArtsStartQueuing, ModelArtsRunning, ModelArtsDeleting, ModelArtsRestarting).Where("repo_id=? and user_id=? and type=? and boot_file!='' and job_type=? and compute_resource=?", repoId, userID, TypeC2Net, string(JobTypeDebug), computeResource).Get(cloudBrain)
			if has {
				return cloudBrain, err
			}
			return nil, err
		}
	} else {
		has, err := x.In("status", JobWaiting, JobRunning, ModelArtsCreateQueue, ModelArtsCreating, ModelArtsStarting,
			ModelArtsReadyToStart, ModelArtsResizing, ModelArtsStartQueuing, ModelArtsRunning, ModelArtsDeleting, ModelArtsRestarting).Where("repo_id=? and user_id=? and type=? and boot_file!='' and job_type=?", repoId, userID, TypeC2Net, JobTypeDebug).Get(cloudBrain)

		if has {
			return cloudBrain, err
		}
		return nil, err
	}

}

func CanDelJob(isSigned bool, user *User, job *CloudbrainInfo) bool {
	if !isSigned || (job.Status != string(JobStopped) && job.Status != string(JobFailed) && job.Status != string(ModelArtsStartFailed) && job.Status != string(ModelArtsCreateFailed)) {
		return false
	}
	repo, err := GetRepositoryByID(job.RepoID)
	if err != nil {
		log.Error("GetRepositoryByID failed:%v", err.Error())
		return false
	}
	permission, _ := GetUserRepoPermission(repo, user)
	if err != nil {
		log.Error("GetUserRepoPermission failed:%v", err.Error())
		return false
	}

	if (user.ID == job.UserID && permission.AccessMode >= AccessModeWrite) || user.IsAdmin || permission.AccessMode >= AccessModeAdmin {
		return true
	}
	return false
}

func GetCloudBrainUnStoppedJob() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0, 10)
	return cloudbrains, x.
		NotIn("status",
			JobStopped, JobSucceeded, JobFailed, ModelArtsCreateFailed, ModelArtsStartFailed, ModelArtsUnavailable, ModelArtsResizFailed, ModelArtsDeleted,
			ModelArtsStopped, ModelArtsTrainJobCanceled, ModelArtsTrainJobCheckFailed, ModelArtsTrainJobCompleted, ModelArtsTrainJobDeleteFailed, ModelArtsTrainJobDeployServiceFailed,
			ModelArtsTrainJobFailed, ModelArtsTrainJobImageFailed, ModelArtsTrainJobKilled, ModelArtsTrainJobLost, ModelArtsTrainJobSubmitFailed, ModelArtsTrainJobSubmitModelFailed, LocalStatusFailed).
		// Limit(1000).
		Find(&cloudbrains)
}

func GetLocalPreparingCreatingCloudBrainJob() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.Where(builder.And(builder.In("status", LocalStatusCreating, LocalStatusPreparing))).
		Find(&cloudbrains)
}
func GetActiveStopCloudBrainJob() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.Where(builder.And(builder.Eq{"status": LocalStatusCreating}).
		Or(builder.Eq{"status": LocalStatusPreparing}.
			And(builder.Lt{"created_unix": time.Now().Add(-1 * setting.PREPARING_MAX_WAIT_DURATION).Unix()}))).
		Find(&cloudbrains)
}

func GetGPUStoppedNotDebugJobDaysAgo(days int, limit int) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0, 10)
	endTimeBefore := time.Now().Unix() - int64(days)*24*3600
	missEndTimeBefore := endTimeBefore - 24*3600
	return cloudbrains, x.Unscoped().AllCols().
		In("status",
			JobStopped, JobSucceeded, JobFailed, ModelArtsCreateFailed, ModelArtsStartFailed, ModelArtsUnavailable, ModelArtsResizFailed, ModelArtsDeleted,
			ModelArtsStopped, ModelArtsTrainJobCanceled, ModelArtsTrainJobCheckFailed, ModelArtsTrainJobCompleted, ModelArtsTrainJobDeleteFailed, ModelArtsTrainJobDeployServiceFailed,
			ModelArtsTrainJobFailed, ModelArtsTrainJobImageFailed, ModelArtsTrainJobKilled, ModelArtsTrainJobLost, ModelArtsTrainJobSubmitFailed, ModelArtsTrainJobSubmitModelFailed).
		Where("(((end_time is null or end_time=0) and updated_unix<? and updated_unix != 0 ) or (end_time<? and end_time != 0)) and cleared=false  and job_type != 'DEBUG' and (type=0 or type=2)", missEndTimeBefore, endTimeBefore).
		Limit(limit).
		Find(&cloudbrains)
}

func GetNPUStoppedNotDebugJobDaysAgo(days int, limit int) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0, 10)
	endTimeBefore := time.Now().Unix() - int64(days)*24*3600
	//missEndTimeBefore := endTimeBefore - 24*3600
	return cloudbrains, x.Unscoped().AllCols().
		In("status",
			JobStopped, JobSucceeded, JobFailed, ModelArtsCreateFailed, ModelArtsStartFailed, ModelArtsUnavailable, ModelArtsResizFailed, ModelArtsDeleted,
			ModelArtsStopped, ModelArtsTrainJobCanceled, ModelArtsTrainJobCheckFailed, ModelArtsTrainJobCompleted, ModelArtsTrainJobDeleteFailed, ModelArtsTrainJobDeployServiceFailed,
			ModelArtsTrainJobFailed, ModelArtsTrainJobImageFailed, ModelArtsTrainJobKilled, ModelArtsTrainJobLost, ModelArtsTrainJobSubmitFailed, ModelArtsTrainJobSubmitModelFailed).
		Where("updated_unix<? and updated_unix != 0 and cleared=false and job_type != 'DEBUG' and type=1", endTimeBefore).
		Limit(limit).
		Find(&cloudbrains)
}

/*
*

	本方法考虑了再次调试的情况，多次调试取最后一次的任务的结束时间
*/
func GetGPUStoppedDebugJobDaysAgo(days int, limit int) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0, 10)
	endTimeBefore := time.Now().Unix() - int64(days)*24*3600
	missEndTimeBefore := endTimeBefore - 24*3600
	sql := `SELECT * from (SELECT DISTINCT ON (job_name)
        id, job_name, job_id,status,end_time,updated_unix,cleared,compute_resource,job_type,type
        FROM   cloudbrain
        where  job_type='DEBUG'
        ORDER  BY job_name, updated_unix DESC) a
        where status in ('STOPPED','SUCCEEDED','FAILED') and (((end_time is null or end_time=0) and updated_unix<? and updated_unix != 0 ) or (end_time<? and end_time != 0)) and (type=0 or type =2 ) and cleared=false`
	return cloudbrains, x.Unscoped().SQL(sql, missEndTimeBefore, endTimeBefore).Limit(limit).Find(&cloudbrains)

}

func GetGPUStoppedDebugJobNotDaysAgo(days int, limit int) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0, 10)
	endTimeBefore := time.Now().Unix() - int64(days)*24*3600
	missEndTimeBefore := endTimeBefore - 24*3600
	sql := `SELECT * from (SELECT DISTINCT ON (job_name)
        id, job_name, job_id,status,end_time,updated_unix,cleared,compute_resource,job_type,type,duration
        FROM   cloudbrain
        where  job_type='DEBUG'
        ORDER  BY job_name, updated_unix DESC) a
        where status in ('STOPPED','SUCCEEDED','FAILED') and (((end_time is null or end_time=0) and updated_unix>=? and updated_unix != 0 ) or (end_time>=? and end_time != 0)) and (type=0 or type =2 ) and cleared=false`

	return cloudbrains, x.Unscoped().SQL(sql, missEndTimeBefore, endTimeBefore).Limit(limit).Find(&cloudbrains)

}

func GetNPUStoppedDebugJobDaysAgo(days int, limit int) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0, 10)
	endTimeBefore := time.Now().Unix() - int64(days)*24*3600
	sql := `SELECT * from (SELECT DISTINCT ON (job_name)
        id, job_name, job_id,status,end_time,updated_unix,cleared,compute_resource,job_type,type
        FROM   cloudbrain
        where  job_type='DEBUG'
        ORDER  BY job_name, updated_unix DESC) a
        where status in ('STOPPED','SUCCEEDED','FAILED') and updated_unix<? and updated_unix != 0  and type=1 and cleared=false`
	//(type=0 or (type =2 and compute_resource='CPU/GPU')) and
	return cloudbrains, x.Unscoped().SQL(sql, endTimeBefore).Limit(limit).Find(&cloudbrains)

}

func UpdateCloudBrainRecordsCleared(ids []int64) error {
	pageSize := 150
	n := len(ids) / pageSize

	var err error

	for i := 1; i <= n+1; i++ {
		tempIds := getPageIds(ids, i, pageSize)
		if len(tempIds) > 0 {
			idsIn := ""
			for i, id := range tempIds {
				if i == 0 {
					idsIn += strconv.FormatInt(id, 10)
				} else {
					idsIn += "," + strconv.FormatInt(id, 10)
				}
			}

			_, errTemp := x.Unscoped().Exec("update cloudbrain set cleared=true where id in (" + idsIn + ")")
			if errTemp != nil {
				err = errTemp
			}

		}

	}
	return err

}

func getPageIds(ids []int64, page int, pagesize int) []int64 {
	begin := (page - 1) * pagesize
	end := (page) * pagesize

	if begin > len(ids)-1 {
		return []int64{}
	}
	if end > len(ids)-1 {
		return ids[begin:]
	} else {
		return ids[begin:end]
	}

}

func GetStoppedJobWithNoDurationJob() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.
		In("status", ModelArtsTrainJobCompleted, ModelArtsTrainJobFailed, ModelArtsTrainJobKilled, ModelArtsStopped, JobStopped, JobFailed, JobSucceeded).
		Where("train_job_duration is null or train_job_duration = '' ").
		Limit(100).
		Find(&cloudbrains)
}
func GetStoppedJobWithNoStartTimeEndTime() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.SQL("select * from cloudbrain where status in (?,?,?,?,?,?,?) and (start_time is null or end_time is null) limit 100", ModelArtsTrainJobCompleted, ModelArtsTrainJobFailed, ModelArtsTrainJobKilled, ModelArtsStopped, JobStopped, JobFailed, JobSucceeded).Find(&cloudbrains)
}
func GetC2NetWithAiCenterWrongJob() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.
		In("status", ModelArtsTrainJobCompleted, ModelArtsTrainJobFailed, ModelArtsTrainJobKilled, ModelArtsStopped, JobStopped, JobFailed, JobSucceeded).
		Where("type = ?", TypeC2Net).
		Find(&cloudbrains)
}

func GetModelSafetyTestTask() ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	sess := x.Where("job_type=?", string(JobTypeModelSafety))
	err := sess.Find(&cloudbrains)
	return cloudbrains, err
}

func GetModelExperienceRunningTask(userID int64) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	sess := x.Where("job_type=?", string(JobTypeModelExperience))
	sess.And("user_id = ? ", userID)
	sess.And("status = ? ", JobRunning)
	err := sess.Find(&cloudbrains)
	return cloudbrains, err
}

func GetModelExperienceNotFinalStatusTask(userID int64, notFinalStatus []string) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	sess := x.Where("job_type=?", string(JobTypeModelExperience))
	sess.And("user_id = ? ", userID)
	sess.In("status", notFinalStatus)
	err := sess.Find(&cloudbrains)
	return cloudbrains, err
}

func GetOnlineInfoTaskByEndPoint(endPoint string) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	sess := x.Where("end_point=?", endPoint)
	err := sess.Find(&cloudbrains)
	return cloudbrains, err
}

func GetCloudbrainRunCountByRepoID(repoID int64) (int, error) {
	count, err := x.In("status", JobWaiting, JobRunning, ModelArtsCreateQueue, ModelArtsCreating, ModelArtsStarting,
		ModelArtsReadyToStart, ModelArtsResizing, ModelArtsStartQueuing, ModelArtsRunning, ModelArtsDeleting, ModelArtsRestarting, ModelArtsTrainJobInit,
		ModelArtsTrainJobImageCreating, ModelArtsTrainJobSubmitTrying, ModelArtsTrainJobWaiting, ModelArtsTrainJobRunning, ModelArtsStopping, ModelArtsResizing,
		ModelArtsTrainJobScaling, ModelArtsTrainJobCheckInit, ModelArtsTrainJobCheckRunning, ModelArtsTrainJobKilling, ModelArtsTrainJobCheckRunningCompleted).And("repo_id = ?", repoID).Count(new(Cloudbrain))
	return int(count), err
}

func GetModelSafetyCountByUserID(userID int64) (int, error) {
	count, err := x.In("status", JobWaiting, JobRunning, ModelArtsTrainJobInit, ModelArtsTrainJobImageCreating, ModelArtsTrainJobSubmitTrying, ModelArtsTrainJobScaling, ModelArtsTrainJobCheckInit, ModelArtsTrainJobCheckRunning, ModelArtsTrainJobCheckRunningCompleted).And("job_type = ? and user_id = ?", string(JobTypeModelSafety), userID).Count(new(Cloudbrain))
	return int(count), err
}

func GetWaitingCloudbrainCount(cloudbrainType int, computeResource string, jobTypes ...JobType) (int64, error) {
	sess := x.Where(builder.NewCond().And(builder.In("status", JobWaiting, LocalStatusPreparing, LocalStatusCreating)).And(builder.Eq{"type": cloudbrainType}))
	if len(jobTypes) > 0 {
		sess.In("job_type", jobTypes)
	}
	if computeResource != "" {
		sess.And("compute_resource=?", computeResource)
	}
	return sess.Count(new(Cloudbrain))
}
func GetNotFinalStatusTaskCount(userID int64, notFinalStatus []string, jobTypes []string) (int, error) {
	count, err := x.In("status", notFinalStatus).
		In("job_type", jobTypes).
		And("user_id = ? ", userID).Count(new(Cloudbrain))
	return int(count), err
}

func RestartCloudbrain(old *Cloudbrain, new *Cloudbrain) (err error) {
	sess := x.NewSession()
	defer sess.Close()

	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Delete(old); err != nil {
		sess.Rollback()
		return err
	}

	if _, err = sess.NoAutoTime().InsertOne(new); err != nil {
		sess.Rollback()
		return err
	}

	if new.Spec != nil {
		if _, err = sess.Insert(NewCloudBrainSpec(new.ID, *new.Spec)); err != nil {
			sess.Rollback()
			return err
		}
	}
	config := old.GetCloudbrainConfig()
	config.CloudbrainID = new.ID
	sess.Insert(config)

	if err = sess.Commit(); err != nil {
		return err
	}

	go updateReferenceCount(new)

	return nil
}
func CloudbrainAll(opts *CloudbrainsOptions) ([]*CloudbrainInfo, int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()
	if opts.RepoID > 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.repo_id": opts.RepoID},
		)
	}

	if opts.UserID > 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.user_id": opts.UserID},
		)
	}

	if (opts.JobID) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.job_id": opts.JobID},
		)
	}
	if (opts.ComputeResource) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain.compute_resource": opts.ComputeResource},
		)
	}

	if (opts.Type) >= 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": opts.Type},
		)
	}

	if len(opts.JobTypes) > 0 {
		if opts.JobTypeNot {
			cond = cond.And(
				builder.NotIn("cloudbrain.job_type", opts.JobTypes),
			)
		} else {
			cond = cond.And(
				builder.In("cloudbrain.job_type", opts.JobTypes),
			)
		}
	}

	if (opts.AiCenter) != "" {
		cond = cond.And(
			builder.Eq{"cloudbrain_spec.ai_center_code": opts.AiCenter},
		)
	}
	if (opts.NeedDeleteInfo) != "" {
		if opts.NeedDeleteInfo == "yes" {
			cond = cond.And(
				builder.And(builder.NotNull{"cloudbrain.deleted_at"}),
			)
		}
		if opts.NeedDeleteInfo == "no" {
			cond = cond.And(
				builder.And(builder.IsNull{"cloudbrain.deleted_at"}),
			)
		}
	}

	if (opts.IsLatestVersion) != "" {
		cond = cond.And(builder.Or(builder.And(builder.Eq{"cloudbrain.is_latest_version": opts.IsLatestVersion},
			builder.Eq{"cloudbrain.job_type": "TRAIN"}), builder.Neq{"cloudbrain.job_type": "TRAIN"}))
	}

	if len(opts.CloudbrainIDs) > 0 {
		cond = cond.And(builder.In("cloudbrain.id", opts.CloudbrainIDs))
	}

	if len(opts.JobStatus) > 0 {
		if opts.JobStatusNot {
			cond = cond.And(
				builder.NotIn("cloudbrain.status", opts.JobStatus),
			)
		} else {
			cond = cond.And(
				builder.In("cloudbrain.status", opts.JobStatus),
			)
		}
	}
	if len(opts.RepoIDList) > 0 {
		cond = cond.And(
			builder.In("cloudbrain.repo_id", opts.RepoIDList),
		)

	}
	if opts.BeginTimeUnix > 0 && opts.EndTimeUnix > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain.created_unix": opts.BeginTimeUnix}, builder.Lte{"cloudbrain.created_unix": opts.EndTimeUnix}),
		)
	}

	if opts.WorkServerNumber > 0 {
		if opts.WorkServerNumber == 1 {
			cond = cond.And(builder.Or(
				builder.Eq{"cloudbrain.work_server_number": 0},
				builder.Eq{"cloudbrain.work_server_number": 1},
				builder.IsNull{"cloudbrain.work_server_number"},
			))
		} else {
			cond = cond.And(
				builder.Eq{"cloudbrain.work_server_number": opts.WorkServerNumber},
			)
		}
	}

	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"cloudbrain_spec.acc_card_type": opts.AccCardType})
	}
	if opts.AccCardsNum >= 0 {
		cond = cond.And(builder.Eq{"cloudbrain_spec.acc_cards_num": opts.AccCardsNum})
	}
	if opts.QueueId > 0 {
		cond = cond.And(builder.Eq{"cloudbrain_spec.queue_id": opts.QueueId})
	}

	var count int64
	var err error
	condition := "cloudbrain.user_id = `user`.id"
	if len(opts.Keyword) == 0 {
		count, err = sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
			Join("left", "`user`", condition).
			Join("left", "cloudbrain_spec", "cloudbrain.id = cloudbrain_spec.cloudbrain_id").
			Count(new(CloudbrainInfo))
	} else {
		lowerKeyWord := strings.ToLower(opts.Keyword)

		cond = cond.And(builder.Or(builder.Like{"LOWER(cloudbrain.job_name)", lowerKeyWord},
			builder.Like{"LOWER(cloudbrain.display_job_name)", lowerKeyWord}, builder.Like{"`user`.lower_name", lowerKeyWord}))
		count, err = sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
			Join("left", "`user`", condition).
			Join("left", "cloudbrain_spec", "cloudbrain.id = cloudbrain_spec.cloudbrain_id").
			Count(new(CloudbrainInfo))
	}

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}

	sess.OrderBy("cloudbrain.created_unix DESC")
	cloudbrains := make([]*CloudbrainInfo, 0, setting.UI.IssuePagingNum)
	if err := sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
		Join("left", "`user`", condition).
		Join("left", "cloudbrain_spec", "cloudbrain.id = cloudbrain_spec.cloudbrain_id").
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}
	if opts.NeedRepoInfo {
		var ids []int64
		for _, task := range cloudbrains {
			ids = append(ids, task.RepoID)
		}
		repositoryMap, err := GetRepositoriesMapByIDs(ids)
		if err == nil {
			for _, task := range cloudbrains {
				task.Repo = repositoryMap[task.RepoID]
			}
		}

	}

	return cloudbrains, count, nil
}

func CloudbrainTotalForDashBoard(opts *CloudbrainsOptions) (int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()

	if opts.BeginTimeUnix > 0 && opts.EndTimeUnix > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain.created_unix": opts.BeginTimeUnix}, builder.Lte{"cloudbrain.created_unix": opts.EndTimeUnix}),
		)
	}
	var count int64
	var err error
	count, err = sess.Unscoped().Table(&Cloudbrain{}).Where(cond).
		Count(new(CloudbrainInfo))

	if err != nil {
		return 0, fmt.Errorf("Count: %v", err)
	}
	return count, nil
}

func CloudbrainAllForDashBoard(opts *CloudbrainsOptions) ([]*CloudbrainInfo, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()

	if opts.BeginTimeUnix > 0 && opts.EndTimeUnix > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain.created_unix": opts.BeginTimeUnix}, builder.Lte{"cloudbrain.created_unix": opts.EndTimeUnix}),
		)
	}
	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}
	sess.OrderBy("cloudbrain.created_unix DESC")
	cloudbrains := make([]*CloudbrainInfo, 0, setting.UI.IssuePagingNum)
	if err := sess.Unscoped().Table(&Cloudbrain{}).Where(cond).
		Find(&cloudbrains); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}
	if opts.NeedRepoInfo {
		var ids []int64
		for _, task := range cloudbrains {
			ids = append(ids, task.RepoID)
		}
		repositoryMap, err := GetRepositoriesMapByIDs(ids)
		if err == nil {
			for _, task := range cloudbrains {
				task.Repo = repositoryMap[task.RepoID]
			}
		}
	}

	return cloudbrains, nil
}

func CloudbrainAllStatic(opts *CloudbrainsOptions) ([]*CloudbrainInfo, int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()

	if (opts.Type) >= 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": opts.Type},
		)
	}
	if opts.BeginTimeUnix > 0 && opts.EndTimeUnix > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain.created_unix": opts.BeginTimeUnix}, builder.Lte{"cloudbrain.created_unix": opts.EndTimeUnix}),
		)
	}
	if opts.DateBeginTimeUnix > 0 || len(opts.JobStatus) > 0 {
		cond = cond.And(builder.Or(
			builder.Gte{"cloudbrain.end_time": opts.DateBeginTimeUnix},
			builder.In("cloudbrain.status", opts.JobStatus),
		))
	}
	var count int64
	var err error
	count, err = sess.Unscoped().Where(cond).Count(new(Cloudbrain))

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}
	// sess.OrderBy("cloudbrain.created_unix DESC")
	cloudbrains := make([]*CloudbrainInfo, 0, setting.UI.IssuePagingNum)
	if err := sess.Table(&Cloudbrain{}).Unscoped().Where(cond).
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}
	if opts.NeedRepoInfo {
		var ids []int64
		for _, task := range cloudbrains {
			ids = append(ids, task.RepoID)
		}
		repositoryMap, err := GetRepositoriesMapByIDs(ids)
		if err == nil {
			for _, task := range cloudbrains {
				task.Repo = repositoryMap[task.RepoID]
			}
		}

	}
	return cloudbrains, count, nil
}

func GetLastestNCloudbrain(n int) ([]*Cloudbrain, error) {
	r := make([]*Cloudbrain, 0)
	err := x.Where("ai_center!='' or type!=2").Desc("id").Limit(n).Unscoped().Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil

}
func CloudbrainAllKanBan(opts *CloudbrainsOptions) ([]*CloudbrainInfo, int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	var cond = builder.NewCond()

	if (opts.Type) >= 0 {
		cond = cond.And(
			builder.Eq{"cloudbrain.type": opts.Type},
		)
	}
	if opts.BeginTimeUnix > 0 && opts.EndTimeUnix > 0 {
		cond = cond.And(
			builder.And(builder.Gte{"cloudbrain.created_unix": opts.BeginTimeUnix}, builder.Lte{"cloudbrain.created_unix": opts.EndTimeUnix}),
		)
	}
	var count int64
	var err error
	count, err = sess.Unscoped().Where(cond).Count(new(Cloudbrain))

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}
	// sess.OrderBy("cloudbrain.created_unix DESC")
	cloudbrains := make([]*CloudbrainInfo, 0, setting.UI.IssuePagingNum)
	if err := sess.Cols("id", "type", "work_server_number", "duration", "train_job_duration", "ai_center", "cluster").Table(&Cloudbrain{}).Unscoped().Where(cond).
		Find(&cloudbrains); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}
	if opts.NeedRepoInfo {
		var ids []int64
		for _, task := range cloudbrains {
			ids = append(ids, task.RepoID)
		}
		repositoryMap, err := GetRepositoriesMapByIDs(ids)
		if err == nil {
			for _, task := range cloudbrains {
				task.Repo = repositoryMap[task.RepoID]
			}
		}

	}
	return cloudbrains, count, nil
}

func GetStartedCloudbrainTaskByUpdatedUnix(startTime, endTime time.Time) ([]Cloudbrain, error) {
	r := make([]Cloudbrain, 0)
	err := x.Where("updated_unix >= ? and updated_unix <= ? and start_time > 0", startTime.Unix(), endTime.Unix()).Unscoped().Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCloudbrainByIds(ids []int64) ([]*Cloudbrain, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	cloudbrains := make([]*Cloudbrain, 0)
	err := x.In("id", ids).Unscoped().Find(&cloudbrains)
	if err != nil {
		return nil, err
	}
	return cloudbrains, nil
}

type DatasetInfo struct {
	DataLocalPath string
	Name          string
	FullName      string
	Type          int
	Size          int64
	DownloadUrl   string
}

type DatasetInfo4AITask struct {
	Compressed   DatasetBaseInfo
	Uncompressed DatasetBaseInfo
	Type         int
	Size         int64
}

type DatasetBaseInfo struct {
	RealPath        string
	ObjectKey       string
	HttpDownloadUrl string
	S3DownloadUrl   string
	Name            string
}

func GetDatasetInfo(uuidStr string, grampusType ...string) (map[string]DatasetInfo, string, error) {
	var datasetNames string
	uuids := strings.Split(uuidStr, ";")
	if len(uuids) > setting.MaxDatasetNum {
		log.Error("the dataset count(%d) exceed the limit", len(uuids))
		return nil, datasetNames, errors.New("the dataset count exceed the limit")
	}

	datasetInfos := make(map[string]DatasetInfo)
	attachs, err := GetAttachmentsByUUIDs(uuids)
	if err != nil {
		log.Error("GetAttachmentsByUUIDs failed: %v", err)
		return nil, datasetNames, err
	}
	for i, tmpUuid := range uuids {
		var attach *Attachment
		for _, tmpAttach := range attachs {
			if tmpAttach.UUID == tmpUuid {
				attach = tmpAttach
				break
			}
		}
		if attach == nil {
			log.Error("GetAttachmentsByUUIDs failed: %v", err)
			return nil, datasetNames, err
		}
		fileName := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(attach.Name, ".zip"), ".tar.gz"), ".tgz")
		for _, datasetInfo := range datasetInfos {
			if fileName == datasetInfo.Name {
				log.Error("the dataset name is same: %v", attach.Name)
				return nil, datasetNames, errors.New("the dataset name is same")
			}
		}
		var dataLocalPath string
		var downloadUrl string
		if len(grampusType) > 0 {
			if grampusType[0] == GPU {
				dataLocalPath = setting.Attachment.Minio.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID
			} else if grampusType[0] == NPU {
				dataLocalPath = setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + "/"
			} else {
				if attach.Type == TypeCloudBrainOne {
					dataLocalPath = setting.Attachment.Minio.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID
				} else {
					dataLocalPath = setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + "/"
				}
			}
		} else {
			if attach.Type == TypeCloudBrainOne {
				dataLocalPath = setting.Attachment.Minio.RealPath +
					setting.Attachment.Minio.Bucket + "/" +
					attach.UnzipPath
				// setting.Attachment.Minio.BasePath +
				// AttachmentRelativePath(attach.UUID) +
				// attach.UUID
			} else {
				//downloadUrl = "s3://" + setting.Bucket + "/" + setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + attach.UUID + "/"
				downloadUrl = "s3://" + setting.Bucket + "/" + attach.UnzipPath
				//dataLocalPath = setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + "/"
				dataLocalPath = attach.UnzipPath
			}

		}

		datasetInfos[attach.UUID] = DatasetInfo{
			DataLocalPath: dataLocalPath,
			Name:          fileName,
			FullName:      attach.Name,
			Type:          attach.Type,
			Size:          attach.Size,
			DownloadUrl:   downloadUrl,
		}
		if i == 0 {
			datasetNames = attach.Name
		} else {
			datasetNames += ";" + attach.Name
		}
	}

	return datasetInfos, datasetNames, nil
}

func GetDatasetInfo4AITask(uuidStr string) (map[string]DatasetInfo4AITask, error) {
	uuids := strings.Split(uuidStr, ";")
	attachments, err := GetAttachmentsByUUIDs(uuids)
	if err != nil {
		log.Error("GetAttachmentsByUUIDs failed: %v", err)
		return nil, err
	}

	attachMap := make(map[string]*Attachment, 0)
	attachNameMap := make(map[string]string, 0)
	for _, attach := range attachments {
		fileName := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(attach.Name, ".zip"), ".tar.gz"), ".tgz")
		if _, exits := attachNameMap[fileName]; exits {
			return nil, errors.New("the dataset name is same")
		}
		attachNameMap[fileName] = ""
		attachMap[attach.UUID] = attach
	}

	datasetInfos := make(map[string]DatasetInfo4AITask)
	for _, tmpUuid := range uuids {
		attach := attachMap[tmpUuid]
		if attach == nil {
			log.Error("GetAttachmentsByUUIDs failed: %v", err)
			return nil, err
		}

		var compressedRealPath, compressedObjectKey, compressedS3DownloadUrl string
		var uncompressedRealPath, uncompressedObjectKey, uncompressedS3DownloadUrl string
		fileName := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(attach.Name, ".zip"), ".tar.gz"), ".tgz")
		if attach.Type == TypeCloudBrainOne {
			uncompressedRealPath = setting.Attachment.Minio.RealPath +
				setting.Attachment.Minio.Bucket + "/" +
				attach.UnzipPath
			uncompressedObjectKey = attach.UnzipPath
			compressedRealPath = setting.Attachment.Minio.RealPath +
				setting.Attachment.Minio.Bucket + "/" +
				attach.Path
			compressedObjectKey = attach.Path
		} else {
			compressedObjectKey = attach.Path
			compressedS3DownloadUrl = "s3://" + setting.Bucket + "/" + compressedObjectKey
			//uncompressedObjectKey = setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + attach.UUID + "/"
			uncompressedObjectKey = attach.UnzipPath
			uncompressedS3DownloadUrl = "s3://" + setting.Bucket + "/" + uncompressedObjectKey
		}

		datasetInfos[attach.UUID] = DatasetInfo4AITask{
			Compressed: DatasetBaseInfo{
				RealPath:      compressedRealPath,
				ObjectKey:     compressedObjectKey,
				S3DownloadUrl: compressedS3DownloadUrl,
				Name:          attach.Name,
			},
			Uncompressed: DatasetBaseInfo{
				RealPath:      uncompressedRealPath,
				ObjectKey:     uncompressedObjectKey,
				S3DownloadUrl: uncompressedS3DownloadUrl,
				Name:          fileName,
			},
			Type: attach.Type,
			Size: attach.Size,
		}
	}

	return datasetInfos, nil
}

var (
	SpecsMapInitFlag                    = false
	CloudbrainDebugResourceSpecsMap     map[int]*ResourceSpec
	CloudbrainTrainResourceSpecsMap     map[int]*ResourceSpec
	CloudbrainInferenceResourceSpecsMap map[int]*ResourceSpec
	CloudbrainBenchmarkResourceSpecsMap map[int]*ResourceSpec
	CloudbrainSpecialResourceSpecsMap   map[int]*ResourceSpec
	GpuInfosMapInitFlag                 = false
	CloudbrainDebugGpuInfosMap          map[string]*GpuInfo
	CloudbrainTrainGpuInfosMap          map[string]*GpuInfo
	CloudbrainInferenceGpuInfosMap      map[string]*GpuInfo
	CloudbrainBenchmarkGpuInfosMap      map[string]*GpuInfo
	CloudbrainSpecialGpuInfosMap        map[string]*GpuInfo
)

func GetNewestJobsByAiCenter() ([]int64, error) {
	ids := make([]int64, 0)
	return ids, x.
		Select("max(id) as id").
		Where("type=? and ai_center!='' and ai_center is not null", TypeC2Net).
		GroupBy("ai_center").
		Table(Cloudbrain{}).
		Find(&ids)
}

func GetNewestJobsByType() ([]int64, error) {
	ids := make([]int64, 0)
	return ids, x.
		Select("max(id) as id").
		In("type", TypeCloudBrainOne, TypeCloudBrainTwo).
		GroupBy("type").
		Table(Cloudbrain{}).
		Find(&ids)
}
func GetFinetuneCloudbrainsByUser(uid int64) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.
		Where("fine_tune=true and user_id=?", uid).Desc("id").
		Find(&cloudbrains)
}

func GetFinetuneCloudbrainsCountByUser(uid int64) (int64, error) {
	cloudbrain := new(Cloudbrain)
	return x.
		Where("fine_tune=true and user_id=?", uid).
		Count(cloudbrain)
}

func GetNewestCloudbrainByUser(uid int64) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	err := x.Unscoped().Where("user_id=? and job_type!=?", uid, JobTypeSuperCompute).Desc("id").Limit(1).Find(&cloudbrains)
	return cloudbrains, err

}

func GetCloudbrainByIDs(ids []int64) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.
		In("id", ids).
		Find(&cloudbrains)
}

type CountPerUserID struct {
	Count  int64
	UserID int64
}

func GetNotebookCountGreaterThanN(n int) ([]CountPerUserID, error) {
	cpuis := []CountPerUserID{}
	err := x.
		Table("cloudbrain").
		GroupBy("user_id").Having("count(*)>"+strconv.Itoa(n)).
		Select("user_id, count(*) AS count").
		Where("job_type=? and (deleted_at=? or deleted_at is NULL)", "DEBUG", "0001-01-01 00:00:00").OrderBy("count(*) desc").
		Find(&cpuis)
	return cpuis, err

}
func GetNotebooksByUser(uid int64, offset int) ([]int64, error) {
	var ints []int64
	err := x.Table("cloudbrain").Cols("id").Where("job_type=? and user_id=? and (deleted_at=? or deleted_at is NULL)", "DEBUG", uid, "0001-01-01 00:00:00").Desc("id").Limit(1000, offset).Find(&ints)
	return ints, err
}

func GetNotebooksCountByUser(uid int64) (int64, error) {
	cloudbrain := new(Cloudbrain)
	return x.Where("user_id=? and job_type=?", uid, "DEBUG").Count(cloudbrain)

}

func GetCloudbrainWithDeletedByIDs(ids []int64) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.
		In("id", ids).Unscoped().Find(&cloudbrains)
}

func GetCloudbrainByJobNameDesc(jobName string) ([]*Cloudbrain, error) {
	cloudbrains := make([]*Cloudbrain, 0)
	return cloudbrains, x.Where("job_name = ?", jobName).Desc("created_unix").Unscoped().Find(&cloudbrains)
}

func GetCloudbrainCountByJobName(jobName, jobType string, typeCloudbrain int) (int, error) {
	count, err := x.Where("job_name = ? and job_type= ? and type = ?", jobName, jobType, typeCloudbrain).Count(new(Cloudbrain))
	return int(count), err
}

func LoadSpecs(tasks []*Cloudbrain) error {
	cloudbrainIds := make([]int64, len(tasks))
	for i, v := range tasks {
		cloudbrainIds[i] = v.ID
	}
	specs := make([]*CloudbrainSpec, 0)
	err := x.In("cloudbrain_id", cloudbrainIds).Find(&specs)
	if err != nil {
		return err
	}
	specMap := make(map[int64]*CloudbrainSpec)
	for _, v := range specs {
		specMap[v.CloudbrainID] = v
	}
	for _, v := range tasks {
		if specMap[v.ID] != nil {
			v.Spec = specMap[v.ID].ConvertToSpecification()
		}
	}
	return nil
}

func LoadSpecs4CloudbrainInfo(tasks []*CloudbrainInfo) error {
	cloudbrainIds := make([]int64, len(tasks))
	for i, v := range tasks {
		cloudbrainIds[i] = v.Cloudbrain.ID
	}
	specs := make([]*CloudbrainSpec, 0)
	err := x.In("cloudbrain_id", cloudbrainIds).Find(&specs)
	if err != nil {
		return err
	}
	specMap := make(map[int64]*CloudbrainSpec)
	for _, v := range specs {
		specMap[v.CloudbrainID] = v
	}
	for _, v := range tasks {
		if specMap[v.Cloudbrain.ID] != nil {
			v.Cloudbrain.Spec = specMap[v.Cloudbrain.ID].ConvertToSpecification()
		}
	}
	return nil
}

func GetCloudBrainByModelId(modelId string) ([]*Cloudbrain, error) {
	cloudBrains := make([]*Cloudbrain, 0)
	err := x.AllCols().Where("model_id like ?", "%"+modelId+"%").OrderBy("created_unix desc").Find(&cloudBrains)
	return cloudBrains, err
}

func GetCloudBrainByRepoIdAndModelName(repoId int64, modelName string) ([]*Cloudbrain, error) {
	cloudBrains := make([]*Cloudbrain, 0)
	err := x.AllCols().Where("model_name like ? and repo_id=?", "%"+modelName+"%", repoId).OrderBy("created_unix asc").Find(&cloudBrains)
	return cloudBrains, err
}

func GetParam(parameterList []Parameter, name string) string {
	for _, v := range parameterList {
		if v.Label == name {
			return v.Value
		}
	}
	return ""
}

var SubTaskName = "task1"

type Dataset4Template struct {
	ID          string `yaml:"id"`
	DatasetName string `yaml:"datasetname"`
}

type PretrainModel4Template struct {
	ID        string `yaml:"id"`
	ModelName string `yaml:"modelname"`
}

type Image4Template struct {
	ImageUrl  string
	ImageID   string
	ImageName string
}

type AITemplateText struct {
	Name              string
	Description       string
	JobType           string                   `yaml:"jobtype"`
	Cluster           string                   `yaml:"cluster"`
	ComputeSource     string                   `yaml:"computesource"`
	BootFile          string                   `yaml:"bootfile"`
	WorkServerNumber  int                      `yaml:"workservernumber"`
	AccCardsNum       int                      `yaml:"acccardsnum"`
	AccCardType       string                   `yaml:"acccardtype"`
	CpuCores          int                      `yaml:"cpucores"`
	MemGiB            float32                  `yaml:"memgib"`
	GPUMemGiB         float32                  `yaml:"gpumemgib"`
	ShareMemGiB       float32                  `yaml:"sharememgib"`
	HasInternet       int                      `yaml:"hasinternet"`
	DatasetList       []Dataset4Template       `yaml:"datasetlist"`
	PretrainModelList []PretrainModel4Template `yaml:"pretrainmodellist"`
	Parameters        []Parameter              `yaml:"parameters"`
	BranchName        string                   `yaml:"branchname"`
	Image             Image4Template           `yaml:"image"`
}

type ParseAITemplateTextReq struct {
	TemplateList []*AITemplateText
}

type RunAITemplateReq struct {
	AITemplate  AITemplateText
	RedirectUrl string
}
