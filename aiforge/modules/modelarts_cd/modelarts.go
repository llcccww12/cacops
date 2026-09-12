package modelarts_cd

import (
	"errors"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/cloudbrain"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	//notebook
	storageTypeOBS                 = "obs"
	autoStopDuration               = 4 * 60 * 60
	autoStopDurationMs             = 4 * 60 * 60 * 1000
	MORDELART_USER_IMAGE_ENGINE_ID = -1
	DataSetMountPath               = "/home/ma-user/work"
	NotebookEnv                    = "Python3"
	NotebookType                   = "Ascend"
	FlavorInfo                     = "Ascend: 1*Ascend 910 CPU: 24 核 96GiB （modelarts.kat1.xlarge）"

	//train-job
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

var ()

type VersionInfo struct {
	Version []struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
		Url   string `json:"url"`
	} `json:"version"`
}

type Flavor struct {
	Info []struct {
		Code  string `json:"code"`
		Value string `json:"value"`
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

type Parameters struct {
	Parameter []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"parameter"`
}

func GenerateNotebook(ctx *context.Context, req cloudbrain.GenerateModelArtsNotebookReq) (string, error) {
	imageName, err := GetNotebookImageName(req.ImageId)
	if err != nil {
		log.Error("GetNotebookImageName failed: %v", err.Error())
		return "", err
	}
	createTime := timeutil.TimeStampNow()
	jobResult, err := createNotebook(models.CreateNotebookWithoutPoolParams{
		JobName:     req.JobName,
		Description: req.Description,
		Flavor:      req.Spec.SourceSpecId,
		Duration:    req.AutoStopDurationMs,
		ImageID:     req.ImageId,
		Feature:     models.NotebookFeature,
		Volume: models.VolumeReq{
			Capacity:  setting.Capacity,
			Category:  models.EVSCategory,
			Ownership: models.ManagedOwnership,
		},
		WorkspaceID: "0",
	})
	if err != nil {
		log.Error("createNotebook failed: %v", err.Error())
		if strings.HasPrefix(err.Error(), UnknownErrorPrefix) {
			log.Info("(%s)unknown error, set temp status", req.DisplayJobName)
			errTemp := models.InsertCloudbrainTemp(&models.CloudbrainTemp{
				JobID:     models.TempJobId,
				VersionID: models.TempVersionId,
				Status:    models.TempJobStatus,
				Type:      models.TypeCDCenter,
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
		Type:             models.TypeCDCenter,
		Uuid:             req.Uuid,
		ComputeResource:  models.NPUResource,
		Image:            imageName,
		Description:      req.Description,
		CreatedUnix:      createTime,
		UpdatedUnix:      createTime,
		Spec:             req.Spec,
		BootFile:         req.BootFile,
		BranchName:       req.BranchName,
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
