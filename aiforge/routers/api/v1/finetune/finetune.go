package finetune

import (
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/timeutil"

	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/services/cloudbrain/resource"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	repo_service "code.gitea.io/gitea/services/repository"
)

func CheckRepo(ctx *context.APIContext) {
	repo, _ := models.GetRepositoryByName(ctx.User.ID, setting.FileNoteBook.ProjectName)
	var dataset *models.Dataset
	var err error
	if repo == nil { //创建项目和数据集
		repo, err = repo_service.CreateRepository(ctx.User, ctx.User, models.CreateRepoOptions{
			Name:          setting.FileNoteBook.ProjectName,
			Alias:         "",
			Description:   "",
			IssueLabels:   "",
			Gitignores:    "",
			License:       "",
			Readme:        "Default",
			IsPrivate:     false,
			AutoInit:      true,
			DefaultBranch: "master",
		})
		if err != nil {
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.failed_to_create_notebook_repo", setting.FileNoteBook.ProjectName)))
			return
		}
		dataset, err = createDataset(ctx, repo)
		if err != nil {
			log.Error("fail to create dataset", err)
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("dataset.create_dataset_fail")))
			return
		}

	} else { //如果数据集不存在，创建数据集
		dataset, err = models.GetDatasetByRepo(repo)
		if err != nil && !models.IsErrNotExist(err) {
			log.Error("fail to get dataset", err)
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("dataset.query_dataset_fail")))
			return
		}
		if dataset == nil {
			dataset, err = createDataset(ctx, repo)
			if err != nil {
				log.Error("fail to create dataset", err)
				ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("dataset.create_dataset_fail")))
				return
			}
		}

	}

	ctx.JSON(http.StatusOK, models.BaseMessageApi{
		Code:    0,
		Message: strconv.FormatInt(dataset.ID, 10),
	})

}
func GetSpec(ctx *context.APIContext) {

	ctx.JSON(http.StatusOK, getFineTuneSpec(ctx.User.ID))
}

func getFineTuneSpec(uid int64) []*api.SpecificationShow {
	specs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(uid, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.NPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainTwo,
	})
	var filterNoteBookSpecs = make([]*api.SpecificationShow, 0)
	for _, spec := range specs {
		if spec.AccCardsNum == 2 {
			filterNoteBookSpecs = append(filterNoteBookSpecs, convert.ToSpecification(spec))
		}
	}

	return filterNoteBookSpecs

}

func GetFinetuneJobs(ctx *context.APIContext) {
	jobs, err := models.GetFinetuneCloudbrainsByUser(ctx.User.ID)
	if err != nil {
		log.Error("get finetune job error", err)
	}

	var fineTuneJobs = make([]*api.FinetuneJobShow, 0)

	for _, job := range jobs {
		fineTuneJobs = append(fineTuneJobs, convert.ToFineTuneJobShow(job))
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"maxJobNum":    setting.FineTune.MaxJobNum,
		"fineTuneJobs": fineTuneJobs,
	})

}
func CreateFineTune(ctx *context.APIContext, option api.CreateFineTuneJobOption) {
	fineTuneCategory := option.SampleDatasetType

	if (option.SampleDatasetType == 0 || option.SampleDatasetType > models.PanguOpenDialog) && option.Attachment == "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
		return
	}
	if !isSpecValid(ctx.User.ID, option.SpecId) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.wrong_specification")))
	}

	if option.Attachment != "" {
		fineTuneCategory = models.PanguCustom
	} else {
		if len(setting.FineTune.Pangu.Dataset.Attachments) < option.SampleDatasetType || len(setting.FineTune.Pangu.Dataset.DatasetNames) < option.SampleDatasetType {
			log.Error("fine tune pangu  attachment config is wrong.")
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
			return
		}
		option.Attachment = setting.FineTune.Pangu.Dataset.Attachments[option.SampleDatasetType-1]
		option.DatasetName = setting.FineTune.Pangu.Dataset.DatasetNames[option.SampleDatasetType-1]
	}
	for _, modelName := range setting.FineTune.Pangu.Model.ModelNames {
		if modelName == option.DatasetName {
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.dataset_same_fail")))
			return
		}
	}
	//supply models dataset
	option.Attachment = option.Attachment + ";" + strings.Join(setting.FineTune.Pangu.Model.ModelAttachments, ";")
	option.DatasetName = option.DatasetName + ";" + strings.Join(setting.FineTune.Pangu.Model.ModelNames, ";")

	fineTuneRepo, err := models.GetRepositoryByOwnerAndName(setting.FineTune.Pangu.Basic.OwnerName, setting.FineTune.Pangu.Basic.RepoName)
	if err != nil {
		log.Error("Can not get pangu repo", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
		return
	}

	jobCount, err := models.GetFinetuneCloudbrainsCountByUser(ctx.User.ID)
	if err != nil {
		log.Error("Can not get finetune job count", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.query_finetune_fail")))
		return
	}
	if jobCount >= int64(setting.FineTune.MaxJobNum) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.finetune_max")))
		return
	}

	trainJobOption := api.CreateTrainJobOption{
		Type:              cloudbrainTask.TaskTypeModelArts,
		DisplayJobName:    option.DisplayJobName,
		ImageID:           strconv.FormatInt(setting.FineTune.Pangu.Basic.ImageId, 10),
		Image:             setting.FineTune.Pangu.Basic.Image,
		Attachment:        option.Attachment,
		DatasetName:       option.DatasetName,
		Description:       option.Description,
		BootFile:          setting.FineTune.Pangu.Basic.BootFile,
		BranchName:        setting.FineTune.Pangu.Basic.BranchName,
		Params:            option.Params,
		WorkServerNumber:  1,
		SpecId:            option.SpecId,
		FineTune:          true,
		FineTuneModelType: models.PanguModelFineTune,
		FineTuneCategory:  fineTuneCategory,
		CodeRepo:          fineTuneRepo,
	}

	cloudbrainTask.ModelArtsTrainJobNpuCreate(ctx.Context, trainJobOption)

}

func isSpecValid(uid int64, specId int64) bool {
	availableSpecs := getFineTuneSpec(uid)
	for _, availableSpec := range availableSpecs {
		if availableSpec.ID == specId {
			return true
		}
	}
	return false
}

func createDataset(ctx *context.APIContext, repoNew *models.Repository) (*models.Dataset, error) {
	dataset := &models.Dataset{}

	dataset.RepoID = repoNew.ID
	dataset.UserID = ctx.User.ID
	dataset.Category = "natural_language_processing"
	dataset.Task = "machine_translation"
	dataset.Title = repoNew.Name

	dataset.Description = "Fine tuning"
	dataset.DownloadTimes = 0
	if repoNew.IsPrivate {
		dataset.Status = 0
	} else {
		dataset.Status = 1
	}
	err := models.CreateDataset(dataset)
	return dataset, err

}

func FineTuneDeployCreate(ctx *context.APIContext, option api.CreateFineTuneDeployOption) {
	// get running service count
	deployments, _ := models.GetRunningServiceByUser(ctx.User.ID)
	log.Info("调试，%s, %v", ctx.User.ID, len(deployments))
	if len(deployments) >= setting.FineTune.Pangu.Deploy.MaxDeployPerUser {
		log.Error("盘古微调部署: 每个用户最多只能同时部署%v个模型", setting.FineTune.Pangu.Deploy.MaxDeployPerUser)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("deployment.deploy_max", setting.FineTune.Pangu.Deploy.MaxDeployPerUser)))
		return
	}
	cloudbrainJob, _ := models.GetCloudbrainByJobID(option.JobID)
	models.CreateModelartsDeploy(&models.ModelartsDeploy{
		JobID:             option.JobID,
		UserID:            cloudbrainJob.UserID,
		JobName:           cloudbrainJob.JobName,
		Status:            "BUILDING",
		DisplayJobName:    cloudbrainJob.DisplayJobName,
		CreateUnix:        timeutil.TimeStampNow(),
		UpdateUnix:        timeutil.TimeStampNow(),
		Finetune:          true,
		FinetuneModelType: cloudbrainJob.FineTuneModelType,
		FinetuneCategory:  cloudbrainJob.FineTuneCategory,
	})

	err := CreateAIModel(ctx, option.JobID)
	if err != nil {
		log.Error("盘古微调部署: 微调任务 %s 创建AI应用失败 %v", option.JobID, err.Error())
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{
		Code:    0,
		Message: option.JobID,
	})
}
