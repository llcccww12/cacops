package modelarts

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/cloudbrain"

	"code.gitea.io/gitea/modules/modelarts_cd"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	//notebook

	storageTypeOBS     = "obs"
	autoStopDuration   = 4 * 60 * 60
	AutoStopDurationMs = 4 * 60 * 60 * 1000

	CodePath         = "/code/"
	OutputPath       = "/output/"
	ResultPath       = "/result/"
	LogPath          = "/log/"
	JobPath          = "/job/"
	OrderDesc        = "desc" //向下查询
	OrderAsc         = "asc"  //向上查询
	Lines            = 500
	TrainUrl         = "train_url"
	DataUrl          = "data_url"
	MultiDataUrl     = "multi_data_url"
	ResultUrl        = "result_url"
	CkptUrl          = "ckpt_url"
	PretrainUrl      = "pretrain_url"
	DeviceTarget     = "device_target"
	Ascend           = "Ascend"
	PerPage          = 10
	IsLatestVersion  = "1"
	NotLatestVersion = "0"
	VersionCountOne  = 1

	SortByCreateTime  = "create_time"
	ConfigTypeCustom  = "custom"
	TotalVersionCount = 1
)

var (
	poolInfos        *models.PoolInfos
	TrainFlavorInfos *Flavor
	SpecialPools     *models.SpecialPools
	MultiNodeConfig  *MultiNodes
)

type GenerateTrainJobReq struct {
	JobName           string
	DisplayJobName    string
	Uuid              string
	Description       string
	CodeObsPath       string
	BootFile          string
	BootFileUrl       string
	DataUrl           string
	TrainUrl          string
	LogUrl            string
	PoolID            string
	WorkServerNumber  int
	EngineID          int64
	Parameters        []models.Parameter
	CommitID          string
	IsLatestVersion   string
	Params            string
	BranchName        string
	PreVersionId      int64
	PreVersionName    string
	FlavorCode        string
	FlavorName        string
	VersionCount      int
	EngineName        string
	TotalVersionCount int
	UserImageUrl      string
	UserCommand       string
	DatasetName       string
	Spec              *models.Specification
	ModelName         string
	LabelName         string
	CkptName          string
	ModelId           string
	ModelVersion      string
	PreTrainModelUrl  string
	FineTune          bool
	FineTuneModelType int
	FineTuneCategory  int
}

type GenerateInferenceJobReq struct {
	JobName           string
	DisplayJobName    string
	Uuid              string
	Description       string
	CodeObsPath       string
	BootFile          string
	BootFileUrl       string
	DataUrl           string
	TrainUrl          string
	LogUrl            string
	PoolID            string
	WorkServerNumber  int
	EngineID          int64
	Parameters        []models.Parameter
	CommitID          string
	Params            string
	BranchName        string
	FlavorName        string
	EngineName        string
	LabelName         string
	IsLatestVersion   string
	VersionCount      int
	TotalVersionCount int
	ModelName         string
	ModelVersion      string
	CkptName          string
	ModelId           string
	ResultUrl         string
	Spec              *models.Specification
	DatasetName       string
	JobType           string
	UserImageUrl      string
	UserCommand       string
}

type GenerateDeployModelReq struct {
	JobID          string
	ModelName      string
	ModelVersion   string
	ModelType      string
	SourceLocation string
	UserID         int64
	Runtime        string
	InstallType    []string
}

type GenerateDeployServiceReq struct {
	JobID         string
	InferType     string
	ServiceName   string
	Spec          string
	Duration      int
	TimeUnit      string
	ScheduleType  string
	ModelID       string
	InstanceCount int
}

type VersionInfo struct {
	Version []struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
		Url   string `json:"url"`
	} `json:"version"`
}

type Flavor struct {
	Info []struct {
		Code      string `json:"code"`
		Value     string `json:"value"`
		UnitPrice int64  `json:"unitPrice"`
	} `json:"flavor"`
}

type Engine struct {
	Info []struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"engine"`
}

type ResourcePool struct {
	Info []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"resource_pool"`
}

type MultiNodes struct {
	Info []OrgMultiNode `json:"multinode"`
}
type OrgMultiNode struct {
	Org  string `json:"org"`
	Node []int  `json:"node"`
}

type Parameters struct {
	Parameter []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"parameter"`
}

func GenerateNotebook2(ctx *context.Context, req cloudbrain.GenerateModelArtsNotebookReq) (string, error) {
	if poolInfos == nil {
		json.Unmarshal([]byte(setting.PoolInfos), &poolInfos)
	}

	imageName, err := GetNotebookImageName(req.ImageId)
	if err != nil {
		log.Error("GetNotebookImageName failed: %v", err.Error())
		return "", err
	}
	createTime := timeutil.TimeStampNow()
	jobResult, err := createNotebook2(models.CreateNotebook2Params{
		JobName:     req.JobName,
		Description: req.Description,
		Flavor:      req.Spec.SourceSpecId,
		Duration:    req.AutoStopDurationMs,
		ImageID:     req.ImageId,
		PoolID:      poolInfos.PoolInfo[0].PoolId,
		Feature:     models.NotebookFeature,
		Volume: models.VolumeReq{
			Capacity:  setting.Capacity,
			Category:  models.EVSCategory,
			Ownership: models.ManagedOwnership,
		},
		WorkspaceID: "0",
	})
	if err != nil {
		log.Error("createNotebook2 failed: %v", err.Error())
		if strings.HasPrefix(err.Error(), UnknownErrorPrefix) {
			log.Info("(%s)unknown error, set temp status", req.DisplayJobName)
			errTemp := models.InsertCloudbrainTemp(&models.CloudbrainTemp{
				JobID:     models.TempJobId,
				VersionID: models.TempVersionId,
				Status:    models.TempJobStatus,
				Type:      models.TypeCloudBrainTwo,
				JobName:   req.JobName,
				JobType:   string(models.JobTypeDebug),
			})
			if errTemp != nil {
				log.Error("InsertCloudbrainTemp failed: %v", errTemp.Error())
				return "", errTemp
			}
		}
		return "", err
	}
	task := &models.Cloudbrain{
		Status:           jobResult.Status,
		UserID:           ctx.User.ID,
		RepoID:           ctx.Repo.Repository.ID,
		JobID:            jobResult.ID,
		JobName:          req.JobName,
		FlavorCode:       req.Spec.SourceSpecId,
		DisplayJobName:   req.DisplayJobName,
		JobType:          string(models.JobTypeDebug),
		Type:             models.TypeCloudBrainTwo,
		Uuid:             req.Uuid,
		ComputeResource:  models.NPUResource,
		Image:            imageName,
		BootFile:         req.BootFile,
		BranchName:       req.BranchName,
		Description:      req.Description,
		CreatedUnix:      createTime,
		UpdatedUnix:      createTime,
		Spec:             req.Spec,
		ModelName:        req.ModelName,
		ModelVersion:     req.ModelVersion,
		LabelName:        req.LabelName,
		PreTrainModelUrl: req.PreTrainModelUrl,
		CkptName:         req.CkptName,
		ModelId:          req.ModelId,
	}

	err = models.CreateCloudbrain(task)
	if err != nil {
		return "", err
	}

	stringId := strconv.FormatInt(task.ID, 10)
	notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, stringId, req.DisplayJobName, models.ActionCreateDebugNPUTask)
	return jobResult.ID, nil
}

func GenerateTrainJob(ctx *context.Context, req *GenerateTrainJobReq) (jobId string, err error) {
	createTime := timeutil.TimeStampNow()
	var jobResult *models.CreateTrainJobResult
	var createErr error
	if req.EngineID < 0 {
		jobResult, createErr = CreateTrainJobUserImage(models.CreateUserImageTrainJobParams{
			JobName:     req.JobName,
			Description: req.Description,
			Config: models.UserImageConfig{
				WorkServerNum: req.WorkServerNumber,
				AppUrl:        req.CodeObsPath,
				BootFileUrl:   req.BootFileUrl,
				DataUrl:       req.DataUrl,
				TrainUrl:      req.TrainUrl,
				LogUrl:        req.LogUrl,
				PoolID:        req.PoolID,
				CreateVersion: true,
				Flavor: models.Flavor{
					Code: req.Spec.SourceSpecId,
				},
				Parameter:    req.Parameters,
				UserImageUrl: req.UserImageUrl,
				UserCommand:  req.UserCommand,
				ShareAddr:    setting.ModelArtsShareAddr,
				MountPath:    setting.ModelArtsMountPath,
				NasType:      setting.ModelArtsNasType,
			},
		})
	} else {
		jobResult, createErr = CreateTrainJob(models.CreateTrainJobParams{
			JobName:     req.JobName,
			Description: req.Description,
			Config: models.Config{
				WorkServerNum: req.WorkServerNumber,
				AppUrl:        req.CodeObsPath,
				BootFileUrl:   req.BootFileUrl,
				DataUrl:       req.DataUrl,
				EngineID:      req.EngineID,
				TrainUrl:      req.TrainUrl,
				LogUrl:        req.LogUrl,
				PoolID:        req.PoolID,
				CreateVersion: true,
				Flavor: models.Flavor{
					Code: req.Spec.SourceSpecId,
				},
				Parameter: req.Parameters,
				ShareAddr: setting.ModelArtsShareAddr,
				MountPath: setting.ModelArtsMountPath,
				NasType:   setting.ModelArtsNasType,
			},
		})
	}
	if createErr != nil {
		log.Error("createTrainJob failed: %v", createErr.Error())
		if strings.HasPrefix(createErr.Error(), UnknownErrorPrefix) {
			log.Info("(%s)unknown error, set temp status", req.DisplayJobName)
			errTemp := models.InsertCloudbrainTemp(&models.CloudbrainTemp{
				JobID:     models.TempJobId,
				VersionID: models.TempVersionId,
				Status:    models.TempJobStatus,
				Type:      models.TypeCloudBrainTwo,
				JobName:   req.JobName,
				JobType:   string(models.JobTypeTrain),
			})
			if errTemp != nil {
				log.Error("InsertCloudbrainTemp failed: %v", errTemp.Error())
				return "", errTemp
			}
		}
		return "", createErr
	}
	jobID := strconv.FormatInt(jobResult.JobID, 10)
	createErr = models.CreateCloudbrain(&models.Cloudbrain{
		Status:            TransTrainJobStatus(jobResult.Status),
		UserID:            ctx.User.ID,
		RepoID:            ctx.Repo.Repository.ID,
		JobID:             jobID,
		JobName:           req.JobName,
		DisplayJobName:    req.DisplayJobName,
		JobType:           string(models.JobTypeTrain),
		Type:              models.TypeCloudBrainTwo,
		VersionID:         jobResult.VersionID,
		VersionName:       jobResult.VersionName,
		Uuid:              req.Uuid,
		DatasetName:       req.DatasetName,
		CommitID:          req.CommitID,
		IsLatestVersion:   req.IsLatestVersion,
		ComputeResource:   models.NPUResource,
		EngineID:          req.EngineID,
		TrainUrl:          req.TrainUrl,
		BranchName:        req.BranchName,
		Parameters:        req.Params,
		BootFile:          req.BootFile,
		DataUrl:           req.DataUrl,
		LogUrl:            req.LogUrl,
		FlavorCode:        req.Spec.SourceSpecId,
		Description:       req.Description,
		WorkServerNumber:  req.WorkServerNumber,
		FlavorName:        req.FlavorName,
		EngineName:        req.EngineName,
		VersionCount:      req.VersionCount,
		TotalVersionCount: req.TotalVersionCount,
		CreatedUnix:       createTime,
		UpdatedUnix:       createTime,
		Spec:              req.Spec,
		ModelName:         req.ModelName,
		ModelVersion:      req.ModelVersion,
		LabelName:         req.LabelName,
		PreTrainModelUrl:  req.PreTrainModelUrl,
		CkptName:          req.CkptName,
		ModelId:           req.ModelId,
		FineTune:          req.FineTune,
		FineTuneModelType: req.FineTuneModelType,
		FineTuneCategory:  req.FineTuneCategory,
	})

	if createErr != nil {
		log.Error("CreateCloudbrain(%s) failed:%v", req.DisplayJobName, createErr.Error())
		return "", createErr
	}
	notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, jobID, req.DisplayJobName, models.ActionCreateTrainTask)
	return jobID, nil
}

func GenerateModelConvertTrainJob(req *GenerateTrainJobReq) (*models.CreateTrainJobResult, error) {

	return CreateTrainJobUserImage(models.CreateUserImageTrainJobParams{
		JobName:     req.JobName,
		Description: req.Description,
		Config: models.UserImageConfig{
			WorkServerNum: req.WorkServerNumber,
			AppUrl:        req.CodeObsPath,
			BootFileUrl:   req.BootFileUrl,
			DataUrl:       req.DataUrl,
			TrainUrl:      req.TrainUrl,
			LogUrl:        req.LogUrl,
			PoolID:        req.PoolID,
			CreateVersion: true,
			Flavor: models.Flavor{
				Code: req.FlavorCode,
			},
			Parameter:    req.Parameters,
			UserImageUrl: req.UserImageUrl,
			UserCommand:  req.UserCommand,
		},
	})
}

func GenerateTrainJobVersion(ctx *context.Context, req *GenerateTrainJobReq, jobId string) (err error) {
	createTime := timeutil.TimeStampNow()
	var jobResult *models.CreateTrainJobResult
	var createErr error

	if req.EngineID < 0 {
		jobResult, createErr = createTrainJobVersionUserImage(models.CreateTrainJobVersionUserImageParams{
			Description: req.Description,
			Config: models.TrainJobVersionUserImageConfig{
				WorkServerNum: req.WorkServerNumber,
				AppUrl:        req.CodeObsPath,
				BootFileUrl:   req.BootFileUrl,
				DataUrl:       req.DataUrl,
				TrainUrl:      req.TrainUrl,
				LogUrl:        req.LogUrl,
				PoolID:        req.PoolID,
				Flavor: models.Flavor{
					Code: req.Spec.SourceSpecId,
				},
				Parameter:    req.Parameters,
				PreVersionId: req.PreVersionId,
				UserImageUrl: req.UserImageUrl,
				UserCommand:  req.UserCommand,
				ShareAddr:    setting.ModelArtsShareAddr,
				MountPath:    setting.ModelArtsMountPath,
				NasType:      setting.ModelArtsNasType,
			},
		}, jobId)
	} else {
		jobResult, createErr = createTrainJobVersion(models.CreateTrainJobVersionParams{
			Description: req.Description,
			Config: models.TrainJobVersionConfig{
				WorkServerNum: req.WorkServerNumber,
				AppUrl:        req.CodeObsPath,
				BootFileUrl:   req.BootFileUrl,
				DataUrl:       req.DataUrl,
				EngineID:      req.EngineID,
				TrainUrl:      req.TrainUrl,
				LogUrl:        req.LogUrl,
				PoolID:        req.PoolID,
				Flavor: models.Flavor{
					Code: req.Spec.SourceSpecId,
				},
				Parameter:    req.Parameters,
				PreVersionId: req.PreVersionId,
				ShareAddr:    setting.ModelArtsShareAddr,
				MountPath:    setting.ModelArtsMountPath,
				NasType:      setting.ModelArtsNasType,
			},
		}, jobId)
	}
	if createErr != nil {
		log.Error("createTrainJobVersion failed: %v", createErr.Error())
		if strings.HasPrefix(createErr.Error(), UnknownErrorPrefix) {
			log.Info("(%s)unknown error, set temp status", req.DisplayJobName)
			errTemp := models.InsertCloudbrainTemp(&models.CloudbrainTemp{
				JobID:     jobId,
				VersionID: models.TempVersionId,
				Status:    models.TempJobStatus,
				Type:      models.TypeCloudBrainTwo,
				JobName:   req.JobName,
				JobType:   string(models.JobTypeTrain),
			})
			if errTemp != nil {
				log.Error("InsertCloudbrainTemp failed: %v", errTemp.Error())
				return errTemp
			}
		}
		return createErr
	}

	var jobTypes []string
	jobTypes = append(jobTypes, string(models.JobTypeTrain))
	repo := ctx.Repo.Repository
	VersionTaskList, VersionListCount, createErr := models.CloudbrainsVersionList(&models.CloudbrainsOptions{
		RepoID:   repo.ID,
		Type:     models.TypeCloudBrainTwo,
		JobTypes: jobTypes,
		JobID:    strconv.FormatInt(jobResult.JobID, 10),
	})
	if createErr != nil {
		ctx.ServerError("Cloudbrain", createErr)
		return createErr
	}
	//将当前版本的isLatestVersion设置为"1"和任务数量更新,任务数量包括当前版本数VersionCount和历史创建的总版本数TotalVersionCount

	createErr = models.CreateCloudbrain(&models.Cloudbrain{
		Status:            TransTrainJobStatus(jobResult.Status),
		UserID:            ctx.User.ID,
		RepoID:            ctx.Repo.Repository.ID,
		JobID:             strconv.FormatInt(jobResult.JobID, 10),
		JobName:           req.JobName,
		DisplayJobName:    req.DisplayJobName,
		JobType:           string(models.JobTypeTrain),
		Type:              models.TypeCloudBrainTwo,
		VersionID:         jobResult.VersionID,
		VersionName:       jobResult.VersionName,
		Uuid:              req.Uuid,
		DatasetName:       req.DatasetName,
		CommitID:          req.CommitID,
		IsLatestVersion:   req.IsLatestVersion,
		PreVersionName:    req.PreVersionName,
		ComputeResource:   models.NPUResource,
		EngineID:          req.EngineID,
		TrainUrl:          req.TrainUrl,
		BranchName:        req.BranchName,
		Parameters:        req.Params,
		BootFile:          req.BootFile,
		DataUrl:           req.DataUrl,
		LogUrl:            req.LogUrl,
		PreVersionId:      req.PreVersionId,
		FlavorCode:        req.Spec.SourceSpecId,
		Description:       req.Description,
		WorkServerNumber:  req.WorkServerNumber,
		FlavorName:        req.FlavorName,
		EngineName:        req.EngineName,
		TotalVersionCount: VersionTaskList[0].TotalVersionCount + 1,
		VersionCount:      VersionListCount + 1,
		CreatedUnix:       createTime,
		UpdatedUnix:       createTime,
		Spec:              req.Spec,
		ModelName:         req.ModelName,
		ModelVersion:      req.ModelVersion,
		LabelName:         req.LabelName,
		PreTrainModelUrl:  req.PreTrainModelUrl,
		CkptName:          req.CkptName,
		ModelId:           req.ModelId,
	})
	if createErr != nil {
		log.Error("CreateCloudbrain(%s) failed:%v", req.JobName, createErr.Error())
		return createErr
	}

	//将训练任务的上一版本的isLatestVersion设置为"0"
	createErr = models.SetVersionCountAndLatestVersion(strconv.FormatInt(jobResult.JobID, 10), VersionTaskList[0].VersionName, VersionCountOne, NotLatestVersion, TotalVersionCount)
	if createErr != nil {
		ctx.ServerError("Update IsLatestVersion failed", createErr)
		return createErr
	}

	return createErr
}

func TransTrainJobStatus(status int) string {
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

func GetOutputPathByCount(TotalVersionCount int) (VersionOutputPath string) {
	talVersionCountToString := fmt.Sprintf("%04d", TotalVersionCount)
	VersionOutputPath = "V" + talVersionCountToString
	return VersionOutputPath
}

func GenerateInferenceJob(ctx *context.Context, req *GenerateInferenceJobReq) (jobId string, err error) {
	createTime := timeutil.TimeStampNow()
	var jobResult *models.CreateTrainJobResult
	var createErr error
	if req.EngineID < 0 {
		jobResult, createErr = createInferenceJobUserImage(models.CreateInfUserImageParams{
			JobName:     req.JobName,
			Description: req.Description,
			Config: models.InfUserImageConfig{
				WorkServerNum: req.WorkServerNumber,
				AppUrl:        req.CodeObsPath,
				BootFileUrl:   req.BootFileUrl,
				DataUrl:       req.DataUrl,
				// TrainUrl:      req.TrainUrl,
				LogUrl:        req.LogUrl,
				PoolID:        req.PoolID,
				CreateVersion: true,
				Flavor: models.Flavor{
					Code: req.Spec.SourceSpecId,
				},
				Parameter:    req.Parameters,
				UserImageUrl: req.UserImageUrl,
				UserCommand:  req.UserCommand,
			},
		})
	} else {
		jobResult, createErr = createInferenceJob(models.CreateInferenceJobParams{
			JobName:     req.JobName,
			Description: req.Description,
			InfConfig: models.InfConfig{
				WorkServerNum: req.WorkServerNumber,
				AppUrl:        req.CodeObsPath,
				BootFileUrl:   req.BootFileUrl,
				DataUrl:       req.DataUrl,
				EngineID:      req.EngineID,
				// TrainUrl:      req.TrainUrl,
				LogUrl:        req.LogUrl,
				PoolID:        req.PoolID,
				CreateVersion: true,
				Flavor: models.Flavor{
					Code: req.Spec.SourceSpecId,
				},
				Parameter: req.Parameters,
			},
		})
	}
	if createErr != nil {
		log.Error("createInferenceJob failed: %v", err.Error())
		if strings.HasPrefix(err.Error(), UnknownErrorPrefix) {
			log.Info("(%s)unknown error, set temp status", req.DisplayJobName)
			err = models.InsertCloudbrainTemp(&models.CloudbrainTemp{
				JobID:     models.TempJobId,
				VersionID: models.TempVersionId,
				Status:    models.TempJobStatus,
				Type:      models.TypeCloudBrainTwo,
				JobName:   req.JobName,
				JobType:   req.JobType,
			})
			if err != nil {
				log.Error("InsertCloudbrainTemp failed: %v", err.Error())
				return "", err
			}
		}
		return "", err
	}

	// attach, err := models.GetAttachmentByUUID(req.Uuid)
	// if err != nil {
	// 	log.Error("GetAttachmentByUUID(%s) failed:%v", strconv.FormatInt(jobResult.JobID, 10), err.Error())
	// 	return err
	// }
	jobID := strconv.FormatInt(jobResult.JobID, 10)
	err = models.CreateCloudbrain(&models.Cloudbrain{
		Status:            TransTrainJobStatus(jobResult.Status),
		UserID:            ctx.User.ID,
		RepoID:            ctx.Repo.Repository.ID,
		JobID:             jobID,
		JobName:           req.JobName,
		DisplayJobName:    req.DisplayJobName,
		JobType:           req.JobType,
		Type:              models.TypeCloudBrainTwo,
		VersionID:         jobResult.VersionID,
		VersionName:       jobResult.VersionName,
		Uuid:              req.Uuid,
		DatasetName:       req.DatasetName,
		CommitID:          req.CommitID,
		EngineID:          req.EngineID,
		TrainUrl:          req.TrainUrl,
		BranchName:        req.BranchName,
		Parameters:        req.Params,
		BootFile:          req.BootFile,
		DataUrl:           req.DataUrl,
		LogUrl:            req.LogUrl,
		FlavorCode:        req.Spec.SourceSpecId,
		Description:       req.Description,
		WorkServerNumber:  req.WorkServerNumber,
		FlavorName:        req.FlavorName,
		EngineName:        req.EngineName,
		LabelName:         req.LabelName,
		IsLatestVersion:   req.IsLatestVersion,
		ComputeResource:   models.NPUResource,
		VersionCount:      req.VersionCount,
		TotalVersionCount: req.TotalVersionCount,
		ModelName:         req.ModelName,
		ModelVersion:      req.ModelVersion,
		CkptName:          req.CkptName,
		ModelId:           req.ModelId,
		ResultUrl:         req.ResultUrl,
		CreatedUnix:       createTime,
		UpdatedUnix:       createTime,
		Spec:              req.Spec,
	})

	if err != nil {
		log.Error("CreateCloudbrain(%s) failed:%v", req.JobName, err.Error())
		return "", err
	}
	if req.JobType == string(models.JobTypeModelSafety) {
		task, err := models.GetCloudbrainByJobID(jobID)
		if err == nil {
			notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, fmt.Sprint(task.ID), req.DisplayJobName, models.ActionCreateBenchMarkTask)
		}
	} else {
		notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, jobID, req.DisplayJobName, models.ActionCreateInferenceTask)
	}

	return jobID, nil
}

func GetNotebookImageName(imageId string) (string, error) {
	var validImage = false
	var imageName = ""

	for _, imageInfo := range setting.StImageInfos.ImageInfo {
		if imageInfo.Id == imageId {
			validImage = true
			imageName = imageInfo.Value
		}
	}

	if !validImage {
		log.Error("the image id(%s) is invalid", imageId)
		return imageName, errors.New("the image id is invalid")
	}

	return imageName, nil
}

func InitSpecialPool() {
	if SpecialPools == nil && setting.ModelArtsSpecialPools != "" {
		json.Unmarshal([]byte(setting.ModelArtsSpecialPools), &SpecialPools)
	}
}

func InitMultiNode() {
	if MultiNodeConfig == nil && setting.ModelArtsMultiNode != "" {
		json.Unmarshal([]byte(setting.ModelArtsMultiNode), &MultiNodeConfig)
	}

}

func HandleTrainJobInfo(task *models.Cloudbrain) error {

	result, err := GetTrainJob(task.JobID, strconv.FormatInt(task.VersionID, 10))
	if err != nil {
		log.Error("GetTrainJob(%s) failed:%v", task.DisplayJobName, err)
		return err
	}

	if result != nil {
		oldStatus := task.Status
		task.Status = TransTrainJobStatus(result.IntStatus)
		task.Duration = result.Duration / 1000
		task.TrainJobDuration = result.TrainJobDuration

		if task.StartTime == 0 && result.StartTime > 0 {
			task.StartTime = timeutil.TimeStamp(result.StartTime / 1000)
		}
		task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)
		if task.EndTime == 0 && models.IsTrainJobTerminal(task.Status) && task.StartTime > 0 {
			task.EndTime = task.StartTime.Add(task.Duration)
		}
		task.CorrectCreateUnix()
		if oldStatus != task.Status {
			notification.NotifyChangeCloudbrainStatus(task, oldStatus)
		}
		err = models.UpdateJob(task)
		if err != nil {
			log.Error("UpdateJob(%s) failed:%v", task.JobName, err)
			return err
		}
	}

	return nil
}

func HandleNotebookInfo(task *models.Cloudbrain) error {
	var result *models.GetNotebook2Result
	var err error
	if task.Type == models.TypeCloudBrainTwo {
		result, err = GetNotebook2(task.JobID)
	} else if task.Type == models.TypeCDCenter {
		result, err = modelarts_cd.GetNotebook(task.JobID)
	}
	if err != nil {
		log.Error("GetNotebook2(%s) failed:%v", task.DisplayJobName, err)
		return err
	}

	if result != nil {
		oldStatus := task.Status
		task.Status = result.Status
		if task.StartTime == 0 && result.Lease.UpdateTime > 0 {
			task.StartTime = timeutil.TimeStamp(result.Lease.UpdateTime / 1000)
		}
		if task.EndTime == 0 && models.IsModelArtsDebugJobTerminal(task.Status) {
			task.EndTime = timeutil.TimeStampNow()
		}
		task.CorrectCreateUnix()
		task.ComputeAndSetDuration()
		if oldStatus != task.Status {
			notification.NotifyChangeCloudbrainStatus(task, oldStatus)
		}
		if task.FlavorCode == "" {
			task.FlavorCode = result.Flavor
		}

		err = models.UpdateJob(task)
		if err != nil {
			log.Error("UpdateJob(%s) failed:%v", task.DisplayJobName, err)
			return err
		}
	}

	return nil
}

func SyncTempStatusJob() {
	jobs, err := models.GetCloudBrainTempJobs()
	if err != nil {
		log.Error("GetCloudBrainTempJobs failed:%v", err.Error())
		return
	}

	for _, temp := range jobs {
		log.Info("start to handle record: %s", temp.JobName)
		if temp.Type == models.TypeCloudBrainTwo {
			if temp.JobType == string(models.JobTypeDebug) {
				err = handleNotebook(temp)
				if err != nil {
					log.Error("handleNotebook falied:%v", err)
					break
				}
			} else if temp.JobType == string(models.JobTypeTrain) || temp.JobType == string(models.JobTypeInference) {
				_, err = models.GetCloudbrainByJobID(temp.JobID)
				if err != nil {
					//one version
					err = handleTrainJob(temp)
					if err != nil {
						log.Error("handleTrainJob falied:%v", err)
						break
					}
				} else {
					//multi version
					err = handleTrainJobMultiVersion(temp)
					if err != nil {
						log.Error("handleTrainJobMultiVersion falied:%v", err)
						break
					}
				}
			}
		}
	}

	return
}

func handleNotebook(temp *models.CloudbrainTemp) error {
	if temp.Status == models.TempJobStatus {
		err := handleTempNotebook(temp)
		if err != nil {
			log.Error("handleTempNotebook failed:%v", err)
			return err
		}
	} else if temp.Status == string(models.ModelArtsStopping) {
		res, err := GetNotebook2(temp.JobID)
		if err != nil {
			log.Error("GetNotebook2 failed:%v", err)
			return err
		}

		temp.Status = res.Status
		if temp.Status == string(models.ModelArtsStopped) {
			err = models.UpdateCloudbrainTemp(temp)
			if err != nil {
				log.Error("UpdateCloudbrainTemp failed:%v", err)
				return err
			}

			_, err := DelNotebook2(temp.JobID)
			if err != nil {
				log.Error("DelNotebook2 failed:%v", err)
				return err
			}

			temp.Status = string(models.ModelArtsDeleted)
			err = models.UpdateCloudbrainTemp(temp)
			if err != nil {
				log.Error("UpdateCloudbrainTemp failed:%v", err)
				return err
			}
		}
	}

	return nil
}

func handleTempNotebook(temp *models.CloudbrainTemp) error {
	var err error
	var isExist bool

	for {
		result, err := GetNotebookList(1000, 0, "createTime", "DESC", temp.JobName)
		if err != nil {
			log.Error("GetNotebookList failed:%v", err)
			break
		}

		temp.QueryTimes++
		err = models.UpdateCloudbrainTemp(temp)
		if err != nil {
			log.Error("UpdateCloudbrainTemp failed:%v", err)
		}

		if result != nil {
			for _, notebook := range result.NotebookList {
				if temp.JobID == models.TempJobId {
					//new notebook
					if notebook.JobName == temp.JobName {
						isExist = true
						temp.Status = notebook.Status
						temp.JobID = notebook.JobID
						break
					}
				} else {
					//restart: always can find one record
					if notebook.JobName == temp.JobName {
						if notebook.Status != string(models.ModelArtsStopped) {
							isExist = true
							temp.Status = notebook.Status
							temp.JobID = notebook.JobID
							break
						}
					}
				}
			}

			if isExist {
				log.Info("find the record(%s), status(%s)", temp.JobName, temp.Status)
				if temp.Status == string(models.ModelArtsCreateFailed) {
					err = models.UpdateCloudbrainTemp(temp)
					if err != nil {
						log.Error("UpdateCloudbrainTemp failed:%v", err)
						break
					}

					_, err := DelNotebook2(temp.JobID)
					if err != nil {
						log.Error("DelNotebook2(%s) failed:%v", temp.JobName, err)
						break
					}

					temp.Status = string(models.ModelArtsDeleted)
				} else {
					_, err := ManageNotebook2(temp.JobID, models.NotebookAction{Action: models.ActionStop})
					if err != nil {
						log.Error("ManageNotebook2(%s) failed:%v", temp.JobName, err)
						break
					}
					temp.Status = string(models.ModelArtsStopping)
				}

				models.UpdateCloudbrainTemp(temp)
			} else {
				log.Error("can not find the record(%s) till now", temp.JobName)
				err = errors.New("not found")
				break
			}
		} else {
			log.Error("can not find the record(%s) till now", temp.JobName)
			err = errors.New("not found")
			break
		}

		break
	}

	if temp.QueryTimes >= setting.MaxTempQueryTimes && !isExist {
		log.Info("reach MaxTempQueryTimes, set the job failed")

		temp.Status = string(models.ModelArtsTrainJobFailed)
		err = models.UpdateCloudbrainTemp(temp)
		if err != nil {
			log.Error("UpdateCloudbrainTemp(%s) failed:%v", temp.JobName, err)
			return err
		}
	}

	return err
}

func handleTrainJob(temp *models.CloudbrainTemp) error {
	if temp.Status == models.TempJobStatus {
		err := handleTempTrainJob(temp)
		if err != nil {
			log.Error("handleTempTrainJob failed:%v", err)
			return err
		}
	} else if temp.Status == string(models.ModelArtsTrainJobKilling) {
		res, err := GetTrainJob(temp.JobID, temp.VersionID)
		if err != nil {
			log.Error("GetTrainJob failed:%v", err)
			return err
		}

		temp.Status = TransTrainJobStatus(res.IntStatus)
		if temp.Status == string(models.ModelArtsTrainJobKilled) {
			err = models.UpdateCloudbrainTemp(temp)
			if err != nil {
				log.Error("UpdateCloudbrainTemp failed:%v", err)
				return err
			}

			_, err := DelTrainJob(temp.JobID)
			if err != nil {
				log.Error("DelTrainJob failed:%v", err)
				return err
			}

			temp.Status = string(models.ModelArtsDeleted)
			err = models.UpdateCloudbrainTemp(temp)
			if err != nil {
				log.Error("UpdateCloudbrainTemp failed:%v", err)
				return err
			}
		}
	}

	return nil
}

func handleTrainJobMultiVersion(temp *models.CloudbrainTemp) error {
	if temp.Status == models.TempJobStatus {
		err := handleTempTrainJobMultiVersion(temp)
		if err != nil {
			log.Error("handleTempTrainJobMultiVersion failed:%v", err)
			return err
		}
	} else if temp.Status == string(models.ModelArtsTrainJobKilling) {
		res, err := GetTrainJob(temp.JobID, temp.VersionID)
		if err != nil {
			log.Error("GetTrainJob failed:%v", err)
			return err
		}

		temp.Status = TransTrainJobStatus(res.IntStatus)
		if temp.Status == string(models.ModelArtsTrainJobKilled) {
			err = models.UpdateCloudbrainTemp(temp)
			if err != nil {
				log.Error("UpdateCloudbrainTemp failed:%v", err)
				return err
			}

			_, err := DelTrainJobVersion(temp.JobID, temp.VersionID)
			if err != nil {
				log.Error("DelTrainJob failed:%v", err)
				return err
			}

			temp.Status = string(models.ModelArtsDeleted)
			err = models.UpdateCloudbrainTemp(temp)
			if err != nil {
				log.Error("UpdateCloudbrainTemp failed:%v", err)
				return err
			}
		}

	}

	return nil
}

func handleTempTrainJobMultiVersion(temp *models.CloudbrainTemp) error {
	var err error
	var isExist bool

	for {
		result, err := GetTrainJobVersionList(1000, 1, temp.JobID)
		if err != nil {
			log.Error("GetTrainJobVersionList failed:%v", err)
			break
		}

		temp.QueryTimes++
		err = models.UpdateCloudbrainTemp(temp)
		if err != nil {
			log.Error("UpdateCloudbrainTemp failed:%v", err)
		}

		if result != nil {
			count, _ := models.GetCloudbrainCountByJobName(temp.JobName, temp.JobType, temp.Type)
			if result.VersionCount == int64(count+1) {
				isExist = true
				temp.Status = TransTrainJobStatus(result.JobVersionList[0].IntStatus)
				temp.VersionID = strconv.FormatInt(result.JobVersionList[0].VersionID, 10)

				log.Info("find the record(%s), status(%s)", temp.JobName, temp.Status)

				_, err := StopTrainJob(temp.JobID, temp.VersionID)
				if err != nil {
					log.Error("StopTrainJob failed:%v", err)
					break
				}
				temp.Status = string(models.ModelArtsTrainJobKilling)
				err = models.UpdateCloudbrainTemp(temp)
				if err != nil {
					log.Error("UpdateCloudbrainTemp(%s) failed:%v", temp.JobName, err)
					break
				}
			} else {
				log.Error("can not find the record(%s) till now", temp.JobName)
				err = errors.New("not found")
				break
			}
		}

		break
	}

	if temp.QueryTimes >= setting.MaxTempQueryTimes && !isExist {
		log.Info("reach MaxTempQueryTimes, set the job failed")

		temp.Status = string(models.ModelArtsTrainJobFailed)
		err = models.UpdateCloudbrainTemp(temp)
		if err != nil {
			log.Error("UpdateCloudbrainTemp(%s) failed:%v", temp.JobName, err)
			return err
		}
	}

	return err
}

func handleTempTrainJob(temp *models.CloudbrainTemp) error {
	var err error
	var isExist bool

	for {
		result, err := GetTrainJobList(1000, 1, "create_time", "desc", temp.JobName)
		if err != nil {
			log.Error("GetTrainJobList failed:%v", err)
			break
		}

		temp.QueryTimes++
		err = models.UpdateCloudbrainTemp(temp)
		if err != nil {
			log.Error("UpdateCloudbrainTemp failed:%v", err)
		}

		if result != nil {
			for _, job := range result.JobList {
				if temp.JobName == job.JobName && TransTrainJobStatus(job.IntStatus) != string(models.ModelArtsTrainJobFailed) {
					isExist = true
					temp.Status = TransTrainJobStatus(job.IntStatus)
					temp.JobID = strconv.FormatInt(job.JobID, 10)
					temp.VersionID = strconv.FormatInt(job.VersionID, 10)

					log.Info("find the record(%s), status(%s)", temp.JobName, temp.Status)

					_, err = StopTrainJob(temp.JobID, temp.VersionID)
					if err != nil {
						log.Error("StopTrainJob(%s) failed:%v", temp.JobName, err)
						break
					}

					temp.Status = string(models.ModelArtsTrainJobKilling)
					err = models.UpdateCloudbrainTemp(temp)
					if err != nil {
						log.Error("UpdateCloudbrainTemp(%s) failed:%v", temp.JobName, err)
						break
					}
				}
			}

			if !isExist {
				log.Error("can not find the record(%s) till now", temp.JobName)
				err = errors.New("not found")
				break
			}
		}

		break
	}

	if temp.QueryTimes >= setting.MaxTempQueryTimes && !isExist {
		log.Info("reach MaxTempQueryTimes, set the job failed")

		temp.Status = string(models.ModelArtsTrainJobFailed)
		err = models.UpdateCloudbrainTemp(temp)
		if err != nil {
			log.Error("UpdateCloudbrainTemp(%s) failed:%v", temp.JobName, err)
			return err
		}
	}

	return err
}

func GenerateDeployModel(ctx *context.Context, req *GenerateDeployModelReq) (string, error) {
	createTime := timeutil.TimeStampNow()
	deploy, _ := models.GetModelartsDeployByJobID(req.JobID)

	var deployModelResult *models.CreateDeployModelResult
	var createErr error
	deployModelResult, createErr = createDeployModel(models.CreateDeployModelParams{
		ModelName:      req.ModelName,
		ModelVersion:   req.ModelVersion,
		ModelType:      req.ModelType,
		SourceLocation: req.SourceLocation,
		Runtime:        req.Runtime,
		InstallType:    req.InstallType,
		Prebuild:       true,
	})
	if createErr != nil {
		deploy.UpdateUnix = createTime
		deploy.Status = "FAILED"
		err := models.UpdateDeploy(deploy)
		if err != nil {
			log.Error("CreateDeployModel DB(%s) failed:%v", deployModelResult.ModelID, createErr.Error())
			return "", err
		}
		log.Error("createDeployModelAPI failed: %v", createErr.Error())
		return "", createErr
	}

	deploy.ModelID = deployModelResult.ModelID
	deploy.ModelType = req.ModelType
	deploy.ModelName = req.ModelName
	deploy.UpdateUnix = createTime
	deploy.ModelStatus = "publishing"
	deploy.Status = "BUILDING"
	err := models.UpdateDeploy(deploy)
	if err != nil {
		log.Error("CreateDeployModel DB(%s) failed:%v", deployModelResult.ModelID, createErr.Error())
		return "", err
	}

	return deployModelResult.ModelID, nil
}

func GenerateDeployService(req *GenerateDeployServiceReq) (string, error) {
	createTime := timeutil.TimeStampNow()
	deploy, _ := models.GetModelartsDeployByJobID(req.JobID)

	var deployServiceResult *models.CreateDeployServiceResult
	var createErr error
	deployServiceResult, createErr = createDeployService(models.CreateDeployServiceParams{
		InferType:   req.InferType,
		ServiceName: req.ServiceName,
		ClusterID:   setting.FineTune.Pangu.Service.ClusterID,
		Config: []models.ServiceConfig{
			{
				Specification: req.Spec,
				ModelID:       req.ModelID,
				InstanceCount: req.InstanceCount,
				// CustomSpec: models.ServiceCustomSpec{
				// 	CPU:    setting.FineTune.Pangu.Service.CPU,
				// 	Memory: setting.FineTune.Pangu.Service.Memory,
				// 	Ascend: setting.FineTune.Pangu.Service.NPU,
				// 	CPUInfo: models.CPUInfo{
				// 		CPU:  setting.FineTune.Pangu.Service.CPU,
				// 		Arch: setting.FineTune.Pangu.Service.CPUArch,
				// 	},
				// 	MemoryInfo: models.MemoryInfo{
				// 		Memory: setting.FineTune.Pangu.Service.MemoryInfo,
				// 		Unit:   setting.FineTune.Pangu.Service.MemoryUnit,
				// 	},
				// 	NPUInfo: models.NPUInfo{
				// 		NPU:     setting.FineTune.Pangu.Service.NPU,
				// 		Brand:   setting.FineTune.Pangu.Service.NPUBrand,
				// 		Version: setting.FineTune.Pangu.Service.NPUVersion,
				// 		Unit:    setting.FineTune.Pangu.Service.NPUMemoryUnit,
				// 		Memory:  setting.FineTune.Pangu.Service.NPUMemory,
				// 	},
				// },
			},
		},
		Schedule: []models.DeploySchedule{
			{
				Duration: req.Duration,
				TimeUnit: req.TimeUnit,
				Type:     req.ScheduleType,
			},
		},
	})

	if createErr != nil {
		log.Error("createDeployServiceAPI failed: %v", createErr.Error())
		deploy.UpdateUnix = createTime
		deploy.Status = "FAILED"
		err := models.UpdateDeploy(deploy)
		if err != nil {
			log.Error("CreateDeployService DB(%s) failed:%v", deployServiceResult.ServiceID, createErr.Error())
			return "", err
		}
		return "", createErr
	}

	deploy.ServiceID = deployServiceResult.ServiceID
	deploy.ServiceName = req.ServiceName
	deploy.UpdateUnix = createTime
	deploy.ServiceStatus = "deploying"
	deploy.Status = "DEPLOYING"
	err := models.UpdateDeploy(deploy)
	if err != nil {
		log.Error("CreateDeployService DB(%s) failed:%v", deployServiceResult.ServiceID, createErr.Error())
		return "", err
	}

	return deployServiceResult.ServiceID, nil
}
