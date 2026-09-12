package repo

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	ai_task "code.gitea.io/gitea/services/ai_task_service/task"

	"code.gitea.io/gitea/services/cloudbrain/modelmanage"
	"code.gitea.io/gitea/services/lock"

	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"

	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"

	"code.gitea.io/gitea/modules/dataset"

	"code.gitea.io/gitea/modules/modelarts_cd"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward/point/account"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	tplDebugJobIndex base.TplName = "repo/debugjob/index"

	tplModelArtsNotebookIndex base.TplName = "repo/modelarts/notebook/index"
	tplModelArtsNotebookNew   base.TplName = "repo/modelarts/notebook/new"
	tplModelArtsNotebookShow  base.TplName = "repo/modelarts/notebook/show"

	tplModelArtsTrainJobIndex      base.TplName = "repo/modelarts/trainjob/index"
	tplModelArtsTrainJobNew        base.TplName = "repo/modelarts/trainjob/new"
	tplModelArtsTrainJobShow       base.TplName = "repo/modelarts/trainjob/show"
	tplModelArtsTrainJobVersionNew base.TplName = "repo/modelarts/trainjob/version_new"

	tplModelArtsInferenceJobIndex base.TplName = "repo/modelarts/inferencejob/index"
	tplModelArtsInferenceJobNew   base.TplName = "repo/modelarts/inferencejob/new"
	tplModelArtsInferenceJobShow  base.TplName = "repo/modelarts/inferencejob/show"
)

func DebugJobIndex(ctx *context.Context) {
	listType := ctx.Query("debugListType")
	if listType == "" {
		listType = models.AllResource
	}
	MustEnableCloudbrain(ctx)
	repo := ctx.Repo.Repository
	ctx.Data["ListType"] = listType
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["CanCreate"] = cloudbrain.CanCreateOrDebugJob(ctx)
	ctx.Data["RepoIsEmpty"] = repo.IsEmpty
	ctx.Data["debugListType"] = listType
	ctx.HTML(200, tplDebugJobIndex)

	// page := ctx.QueryInt("page")
	// if page <= 0 {
	// 	page = 1
	// }

	// jobTypeNot := false
	// var computeResource string
	// if listType != models.AllResource {
	// 	computeResource = listType
	// }

	// var jobTypes []string
	// jobTypes = append(jobTypes, string(models.JobTypeDebug))
	// ciTasks, count, err := models.Cloudbrains(&models.CloudbrainsOptions{
	// 	ListOptions: models.ListOptions{
	// 		Page:     page,
	// 		PageSize: setting.UI.IssuePagingNum,
	// 	},
	// 	RepoID:          repo.ID,
	// 	ComputeResource: computeResource,
	// 	Type:            models.TypeCloudBrainAll,
	// 	JobTypeNot:      jobTypeNot,
	// 	JobTypes:        jobTypes,
	// })
	// if err != nil {
	// 	ctx.ServerError("Get debugjob faild:", err)
	// 	return
	// }

	// for i, task := range ciTasks {
	// 	ciTasks[i].CanDebug = cloudbrain.CanModifyJob(ctx, &task.Cloudbrain)
	// 	ciTasks[i].CanDel = cloudbrain.CanDeleteJob(ctx, &task.Cloudbrain)
	// 	ciTasks[i].Cloudbrain.ComputeResource = task.ComputeResource
	// }

	// pager := context.NewPagination(int(count), setting.UI.IssuePagingNum, page, 5)
	// pager.AddParam(ctx, "debugListType", "ListType")
}

// MustEnableDataset check if repository enable internal cb
func MustEnableModelArts(ctx *context.Context) {
	if !ctx.Repo.CanRead(models.UnitTypeCloudBrain) {
		ctx.NotFound("MustEnableCloudbrain", nil)
		return
	}
}

func NotebookNew(ctx *context.Context) {
	// notebookNewDataPrepare(ctx)
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplModelArtsNotebookNew)
}

func notebookNewDataPrepare(ctx *context.Context) error {
	ctx.Data["PageIsCloudBrain"] = true
	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName

	attachs, err := models.GetModelArtsUserAttachments(ctx.User.ID)
	if err != nil {
		ctx.ServerError("GetAllUserAttachments failed:", err)
		return err
	}
	ctx.Data["attachments"] = attachs
	ctx.Data["images"] = setting.StImageInfos.ImageInfo

	prepareCloudbrainTwoDebugSpecs(ctx)

	ctx.Data["datasetType"] = models.TypeCloudBrainTwo

	waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeCloudBrainTwo, "")
	ctx.Data["WaitCount"] = waitCount
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeDebug))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount

	return nil
}

func prepareCloudbrainTwoDebugSpecs(ctx *context.Context) {
	aiCenterCode := models.AICenterOfCloudBrainTwo
	if setting.ModelartsCD.Enabled {
		aiCenterCode = models.AICenterOfChengdu
	}
	noteBookSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeDebug,
		ComputeResource: models.NPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    aiCenterCode,
	})
	ctx.Data["Specs"] = noteBookSpecs
}

func NotebookShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplModelArtsNotebookShow)
	return
	/*
		debugListType := ctx.Query("debugListType")
		if debugListType == "" {
			debugListType = "all"
		}
		var ID = ctx.Params(":id")
		task, err := models.GetCloudbrainByIDWithDeleted(ID)
		if err != nil {
			log.Error("GET job error", err.Error())
			ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
			return
		}

		if task.DeletedAt.IsZero() && !task.Cleared { //normal record
			err := modelarts.HandleNotebookInfo(task)
			if err != nil {
				ctx.Data["error"] = err.Error()
				ctx.RenderWithErr(err.Error(), tplModelArtsNotebookShow, nil)
				return
			}
		}
		datasetDownload := make([]*models.DatasetDownload, 0)
		var modelDownload models.ModelDownload
		if ctx.IsSigned {
			if task.Uuid != "" && task.UserID == ctx.User.ID {
				datasetDownload = GetCloudBrainDataSetInfo(task.Uuid, task.DatasetName, true)
			}
			if task.ModelName != "" && task.UserID == ctx.User.ID {
				modelDownload = GetModelDownload(task)

			}

		}
		user, err := models.GetUserByID(task.UserID)
		if err == nil {
			task.User = user
		}
		prepareSpec4Show(ctx, task)
		if task.TrainJobDuration == "" {
			if task.Duration == 0 {
				var duration int64
				if task.Status == string(models.JobRunning) {
					duration = time.Now().Unix() - int64(task.CreatedUnix)
				} else {
					duration = int64(task.UpdatedUnix) - int64(task.CreatedUnix)
				}
				task.Duration = duration
			}
			task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)
		}
		ctx.Data["duration"] = task.TrainJobDuration
		ctx.Data["datasetDownload"] = datasetDownload
		ctx.Data["modelDownload"] = modelDownload
		ctx.Data["task"] = task
		ctx.Data["ID"] = ID
		ctx.Data["jobName"] = task.JobName
		ctx.Data["debugListType"] = debugListType
		ctx.HTML(200, tplModelArtsNotebookShow)
	*/
}

func GetModelDownload(task *models.Cloudbrain) models.ModelDownload {
	index := strings.Index(task.PreTrainModelUrl, "/")
	key := task.PreTrainModelUrl[index+1:] + task.CkptName
	url, _ := storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, key)
	modelDownload := models.ModelDownload{
		Name:         task.CkptName,
		DownloadLink: url,
		IsDelete:     false,
	}

	if !modelmanage.HasModelFile(task) {
		log.Warn("Can not get model by path:" + task.PreTrainModelUrl)
		modelDownload.IsDelete = true
	}
	return modelDownload
}

func GetCloudBrainDataSetInfo(uuid string, datasetname string, isNeedDown bool) []*models.DatasetDownload {
	datasetDownload := make([]*models.DatasetDownload, 0)
	if len(uuid) == 0 {
		return datasetDownload
	}
	uuidList := strings.Split(uuid, ";")
	datasetnameList := strings.Split(datasetname, ";")
	for i, uuidStr := range uuidList {
		name := ""
		link := ""
		url := ""
		isDelete := false
		attachment, err := models.GetAttachmentByUUID(uuidStr)
		if err != nil {
			log.Error("GetAttachmentByUUID failed:%v", err.Error())
			if len(datasetnameList) <= i || len(datasetname) == 0 {
				continue
			}
			name = datasetnameList[i]
			isDelete = true
		} else {
			name = attachment.Name
			dataset, err := models.GetDatasetByID(attachment.DatasetID)
			if err != nil {
				log.Error("GetDatasetByID failed:%v", err.Error())
			} else {
				repo, err := models.GetRepositoryByID(dataset.RepoID)
				if err != nil {
					log.Error("GetRepositoryByID failed:%v", err.Error())
				} else {
					link = repo.Link() + "/datasets"
				}
			}
			if isNeedDown {
				url = attachment.S3DownloadURL()
			}
		}

		datasetDownload = append(datasetDownload, &models.DatasetDownload{
			DatasetName:         name,
			DatasetDownloadLink: url,
			RepositoryLink:      link,
			IsDelete:            isDelete,
		})
	}
	log.Info("dataset length=" + fmt.Sprint(len(datasetDownload)))
	return datasetDownload
}

func setShowSpecBySpecialPoolConfig(ctx *context.Context, findSpec bool, task *models.Cloudbrain) {
	modelarts.InitSpecialPool()
	if modelarts.SpecialPools != nil && !findSpec {
		for _, pool := range modelarts.SpecialPools.Pools {
			for _, flavor := range pool.Flavor {
				if flavor.Value == task.FlavorCode {
					ctx.Data["resource_spec"] = flavor.Desc
				}
			}
		}

	}
}

func NotebookDebug2(ctx *context.Context) {
	var err error
	var result *models.GetNotebook2Result
	task := ctx.Cloudbrain
	if task.Type == models.TypeCloudBrainTwo {
		result, err = modelarts.GetNotebook2(task.JobID)
	} else if task.Type == models.TypeCDCenter {
		result, err = modelarts_cd.GetNotebook(task.JobID)
	}

	if err != nil {
		ctx.RenderWithErr(err.Error(), tplModelArtsNotebookIndex, nil)
		return
	}

	if ctx.QueryTrim("file") != "" {
		ctx.Redirect(getFileUrl(result.Url, ctx.QueryTrim("file")) + "&token=" + result.Token)
	} else {
		if task.BootFile != "" {
			go cloudbrainTask.UploadNotebookFiles(task)
		}
		ctx.Redirect(result.Url + "?token=" + result.Token)
	}
}

func getFileUrl(url string, filename string) string {
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

	return url + middle + filename + "?reset"
}

func NotebookStop(ctx *context.Context) {
	var id = ctx.Params(":id")
	var resultCode = "0"
	var errorMsg = ""
	var status = ""

	task := ctx.Cloudbrain

	for {
		if task.Status != string(models.ModelArtsRunning) {
			log.Error("the job(%s) is not running", task.JobName, ctx.Data["MsgID"])
			resultCode = "-1"
			errorMsg = ctx.Tr("cloudbrain.Already_stopped")
			break
		}
		if res, isHandled, err := ai_task.HandleNewAITaskStop(task.ID); isHandled {
			if err != nil {
				log.Error("ManageNotebook2(%s) failed:%v", task.JobName, err.Error(), ctx.Data["MsgID"])
				resultCode = "-1"
				errorMsg = err.Error()
				if strings.Contains(err.Error(), modelarts.NotebookNotFound) {
					errorMsg = "the job's version is too old and can not be restarted"
				}
				break
			}
			status = res.Status
			break
		}

		err, res := StopModelArtsNotebook(task)

		if err != nil {
			log.Error("ManageNotebook2(%s) failed:%v", task.JobName, err.Error(), ctx.Data["MsgID"])
			resultCode = "-1"
			errorMsg = err.Error()
			if strings.Contains(err.Error(), modelarts.NotebookNotFound) {
				errorMsg = "the job's version is too old and can not be restarted"
			}
			break
		}

		status = res.Status
		oldStatus := task.Status
		task.Status = res.Status
		if task.EndTime == 0 && models.IsModelArtsDebugJobTerminal(task.Status) {
			task.EndTime = timeutil.TimeStampNow()
		}
		task.ComputeAndSetDuration()
		if oldStatus != task.Status {
			notification.NotifyChangeCloudbrainStatus(task, oldStatus)
		}
		err = models.UpdateJob(task)
		if err != nil {
			log.Error("UpdateJob(%s) failed:%v", task.JobName, err.Error(), ctx.Data["MsgID"])
			resultCode = "-1"
			errorMsg = "system error"
			break
		}

		break
	}

	ctx.JSON(200, map[string]string{
		"result_code": resultCode,
		"error_msg":   errorMsg,
		"status":      status,
		"id":          id,
	})
}

func StopModelArtsNotebook(task *models.Cloudbrain) (error, *models.NotebookActionResult) {
	param := models.NotebookAction{
		Action: models.ActionStop,
	}

	var err error
	var res *models.NotebookActionResult
	if task.Type == models.TypeCloudBrainTwo {
		res, err = modelarts.ManageNotebook2(task.JobID, param)
	} else if task.Type == models.TypeCDCenter {
		res, err = modelarts_cd.ManageNotebook(task.JobID, param)
	}
	return err, res
}

func NotebookDel(ctx *context.Context) {
	var listType = ctx.Query("debugListType")
	task := ctx.Cloudbrain

	if isHandled, err := ai_task.HandleNewAITaskDelete(task.ID); isHandled {
		if err != nil {
			log.Error("DeleteJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
			ctx.RenderWithErr("DeleteJob failed", tplDebugJobIndex, nil)
		}
		var isAdminPage = ctx.Query("isadminpage")
		var isHomePage = ctx.Query("ishomepage")
		if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
			ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
		} else if isHomePage == "true" {
			ctx.Redirect(setting.AppSubURL + "/cloudbrains")
		} else {
			ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/debugjob?debugListType=" + listType)
		}
	}

	if task.Status != string(models.ModelArtsCreateFailed) && task.Status != string(models.ModelArtsStartFailed) && task.Status != string(models.ModelArtsStopped) && task.Status != string(models.ModelArtsDeleted) {
		log.Error("the job(%s) has not been stopped", task.JobName)
		ctx.RenderWithErr("the job has not been stopped", tplDebugJobIndex, nil)
		return
	}

	var err error
	if task.Type == models.TypeCloudBrainTwo {
		_, err = modelarts.DelNotebook2(task.JobID)
	} else if task.Type == models.TypeCDCenter {
		_, err = modelarts_cd.DelNotebook(task.JobID)
	}

	if err != nil {
		log.Error("DelNotebook2(%s) failed:%v", task.JobName, err.Error())
		if strings.Contains(err.Error(), modelarts.NotebookNotFound) || strings.Contains(err.Error(), modelarts.NotebookNoPermission) || strings.Contains(err.Error(), modelarts.NotebookInvalid) {
			log.Info("old notebook version")
		} else {
			ctx.RenderWithErr(err.Error(), tplDebugJobIndex, nil)
			return
		}
	}

	err = models.DeleteJob(task)
	if err != nil {
		ctx.RenderWithErr(err.Error(), tplDebugJobIndex, nil)
		return
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/debugjob?debugListType=" + listType)
	}
}

func TrainJobIndex(ctx *context.Context) {
	MustEnableModelArts(ctx)
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplModelArtsTrainJobIndex)
	/*
		repo := ctx.Repo.Repository
		page := ctx.QueryInt("page")
		if page <= 0 {
			page = 1
		}

		listType := ctx.Query("listType")
		ctx.Data["ListType"] = listType

		if listType == models.AllResource {
			listType = ""
		}

		var jobTypes []string
		jobTypes = append(jobTypes, string(models.JobTypeTrain))
		tasks, count, err := models.Cloudbrains(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: setting.UI.IssuePagingNum,
			},
			RepoID:          repo.ID,
			JobTypeNot:      false,
			JobTypes:        jobTypes,
			IsLatestVersion: modelarts.IsLatestVersion,
			ComputeResource: listType,
			Type:            models.TypeCloudBrainAll,
		})
		if err != nil {
			ctx.ServerError("Cloudbrain", err)
			return
		}

		for i, task := range tasks {
			tasks[i].CanDel = cloudbrain.CanDeleteJob(ctx, &task.Cloudbrain)
			tasks[i].CanModify = cloudbrain.CanModifyJob(ctx, &task.Cloudbrain)
		}

		pager := context.NewPagination(int(count), setting.UI.IssuePagingNum, page, 5)
		pager.SetDefaultParams(ctx)
		pager.AddParam(ctx, "listType", "ListType")
		ctx.Data["Page"] = pager

		ctx.Data["PageIsCloudBrain"] = true
		ctx.Data["Tasks"] = tasks
		ctx.Data["CanCreate"] = cloudbrain.CanCreateOrDebugJob(ctx)
		ctx.Data["RepoIsEmpty"] = repo.IsEmpty
		ctx.HTML(200, tplModelArtsTrainJobIndex)
	*/
}

func TrainJobNew(ctx *context.Context) {
	err := trainJobNewDataPrepare(ctx)
	if err != nil {
		ctx.ServerError("get new train-job info failed", err)
		return
	}
	ctx.HTML(200, tplModelArtsTrainJobNew)
}

func trainJobNewDataPrepare(ctx *context.Context) error {
	ctx.Data["PageIsCloudBrain"] = true

	//can, err := canUserCreateTrainJob(ctx.User.ID)
	//if err != nil {
	//	ctx.ServerError("canUserCreateTrainJob", err)
	//	return
	//}
	//
	//if !can {
	//	log.Error("the user can not create train-job")
	//	ctx.ServerError("the user can not create train-job", fmt.Errorf("the user can not create train-job"))
	//	return
	//}

	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName

	attachs, err := models.GetModelArtsTrainAttachments(ctx.User.ID)
	if err != nil {
		ctx.ServerError("GetAllUserAttachments failed:", err)
		return err
	}
	ctx.Data["attachments"] = attachs

	var resourcePools modelarts.ResourcePool
	if err = json.Unmarshal([]byte(setting.ResourcePools), &resourcePools); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["resource_pools"] = resourcePools.Info

	var engines modelarts.Engine
	if err = json.Unmarshal([]byte(setting.Engines), &engines); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engines"] = engines.Info

	var versionInfos modelarts.VersionInfo
	if err = json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engine_versions"] = versionInfos.Version

	prepareCloudbrainTwoTrainSpecs(ctx)

	ctx.Data["params"] = ""
	ctx.Data["branchName"] = ctx.Repo.BranchName

	configList, err := getConfigList(modelarts.PerPage, 1, modelarts.SortByCreateTime, "desc", "", modelarts.ConfigTypeCustom)
	if err != nil {
		ctx.ServerError("getConfigList failed:", err)
		return err
	}
	ctx.Data["config_list"] = configList.ParaConfigs
	ctx.Data["datasetType"] = models.TypeCloudBrainTwo
	waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeCloudBrainTwo, "")
	ctx.Data["WaitCount"] = waitCount
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount

	setMultiNodeIfConfigureMatch(ctx)

	return nil
}

func prepareCloudbrainTwoTrainSpecs(ctx *context.Context) {
	noteBookSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.NPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainTwo,
	})
	ctx.Data["Specs"] = noteBookSpecs
}

func setMultiNodeIfConfigureMatch(ctx *context.Context) {
	modelarts.InitMultiNode()
	if modelarts.MultiNodeConfig != nil {
		for _, info := range modelarts.MultiNodeConfig.Info {
			if isInOrg, _ := models.IsOrganizationMemberByOrgName(info.Org, ctx.User.ID); isInOrg {
				ctx.Data["WorkNode"] = info.Node
				break
			}
		}
	}
}

func setSpecBySpecialPoolConfig(ctx *context.Context, jobType string) {
	modelarts.InitSpecialPool()

	if modelarts.SpecialPools != nil {
		for _, specialPool := range modelarts.SpecialPools.Pools {
			if cloudbrain.IsElementExist(specialPool.JobType, jobType) {

				if isInOrg, _ := models.IsOrganizationMemberByOrgName(specialPool.Org, ctx.User.ID); isInOrg {
					var specialFlavor []struct {
						Code  string
						Value string
					}

					if jobType == string(models.JobTypeDebug) {
						ctx.Data["flavors"] = specialPool.Flavor
					} else {

						for _, tempFlavor := range specialPool.Flavor {
							specialFlavor = append(specialFlavor, struct {
								Code  string
								Value string
							}{Code: tempFlavor.Value, Value: tempFlavor.Desc})
						}

						ctx.Data["flavor_infos"] = specialFlavor
					}

				}

			}
		}

	}
}

func TrainJobNewVersion(ctx *context.Context) {

	err := trainJobNewVersionDataPrepare(ctx)
	if err != nil {
		ctx.ServerError("get new train-job info failed", err)
		return
	}
	ctx.HTML(200, tplModelArtsTrainJobVersionNew)
}

func trainJobNewVersionDataPrepare(ctx *context.Context) error {
	ctx.Data["PageIsCloudBrain"] = true
	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")

	// canNewJob, err := canUserCreateTrainJobVersion(ctx, jobID, versionName)
	// if err != nil {
	// 	ctx.ServerError("canNewJob can info failed", err)
	// 	return err
	// }

	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobIDAndVersionName(%s) failed:%v", jobID, err.Error())
		return err
	}

	ctx.Data["display_job_name"] = task.DisplayJobName
	ctx.Data["job_name"] = task.JobName

	// attachs, err := models.GetModelArtsTrainAttachments(ctx.User.ID)
	// if err != nil {
	// 	ctx.ServerError("GetAllUserAttachments failed:", err)
	// 	return err
	// }
	// ctx.Data["attachments"] = attachs

	var resourcePools modelarts.ResourcePool
	if err = json.Unmarshal([]byte(setting.ResourcePools), &resourcePools); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["resource_pools"] = resourcePools.Info

	var engines modelarts.Engine
	if err = json.Unmarshal([]byte(setting.Engines), &engines); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engines"] = engines.Info

	var versionInfos modelarts.VersionInfo
	if err = json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engine_versions"] = versionInfos.Version

	prepareCloudbrainTwoTrainSpecs(ctx)
	spec, _ := resource.GetCloudbrainSpec(task.ID)
	if spec != nil {
		log.Info("spec_id = %d", spec.ID)
		ctx.Data["spec_id"] = spec.ID
	}

	ctx.Data["run_para_list"] = task.Parameters

	branches, _, err := ctx.Repo.GitRepo.GetBranches(0, 0)
	if err != nil {
		ctx.ServerError("GetBranches error:", err)
		return err
	}

	uuids, datasetNames := dataset.GetFilterDeletedAttachments(task.Uuid)

	ctx.Data["dataset_name"] = datasetNames

	ctx.Data["branches"] = branches
	ctx.Data["branch_name"] = task.BranchName
	ctx.Data["description"] = task.Description
	ctx.Data["boot_file"] = task.BootFile

	ctx.Data["work_server_number"] = task.WorkServerNumber
	ctx.Data["flavor_name"] = task.FlavorName
	ctx.Data["engine_name"] = task.EngineName
	ctx.Data["attachment"] = uuids
	ctx.Data["flavor_code"] = task.FlavorCode
	ctx.Data["engine_id"] = task.EngineID
	ctx.Data["datasetType"] = models.TypeCloudBrainTwo

	//pretrain model
	ctx.Data["model_name"] = task.ModelName
	ctx.Data["model_version"] = task.ModelVersion
	ctx.Data["ckpt_name"] = task.CkptName
	ctx.Data["model_id"] = ctx.Cloudbrain.ModelId
	ctx.Data["label_names"] = task.LabelName
	ctx.Data["pre_train_model_url"] = task.PreTrainModelUrl

	configList, err := getConfigList(modelarts.PerPage, 1, modelarts.SortByCreateTime, "desc", "", modelarts.ConfigTypeCustom)
	if err != nil {
		ctx.ServerError("getConfigList failed:", err)
		return err
	}
	ctx.Data["config_list"] = configList.ParaConfigs
	waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeCloudBrainTwo, "")
	ctx.Data["WaitCount"] = waitCount
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount

	return nil
}

func checkMultiNode(userId int64, serverNum int) string {
	if serverNum == 1 {
		return ""
	}
	modelarts.InitMultiNode()
	var isServerNumValid = false
	if modelarts.MultiNodeConfig != nil {
		for _, info := range modelarts.MultiNodeConfig.Info {
			if isInOrg, _ := models.IsOrganizationMemberByOrgName(info.Org, userId); isInOrg {
				if isInNodes(info.Node, serverNum) {
					isServerNumValid = true
					break
				}

			}
		}
	}
	if isServerNumValid {
		return ""
	} else {
		return "repo.modelarts.no_node_right"
	}
}
func addModelUrlParam(param []models.Parameter, preTrainModelUrl string, ckptName string) []models.Parameter {
	ckptNames := strings.Split(ckptName, ";")
	ckptUrl := "/" + preTrainModelUrl + ckptNames[0]
	param = append(param, models.Parameter{
		Label: modelarts.CkptUrl,
		Value: "s3:/" + ckptUrl,
	})

	var modelUrlList []models.ModelUrls
	for _, ckptName := range ckptNames {
		modelUrlList = append(modelUrlList, models.ModelUrls{
			ModelUrl:  "s3://" + preTrainModelUrl + ckptName,
			ModelName: ckptName,
		})
	}
	modelUrlsJson, err := json.Marshal(modelUrlList)
	if err != nil {
		log.Error("Failed to Marshal: %v", err)
		return param
	}
	param = append(param, models.Parameter{
		Label: modelarts.PretrainUrl,
		Value: string(modelUrlsJson),
	})
	return param
}

func checkInferenceJobMultiNode(userId int64, serverNum int) string {
	if serverNum == 1 {
		return ""
	}

	return "repo.modelarts.no_node_right"

}

func isInNodes(nodes []int, num int) bool {
	for _, node := range nodes {
		if node == num {
			return true
		}
	}
	return false

}

func getUserCommand(engineId int, req *modelarts.GenerateTrainJobReq) (string, string) {
	userImageUrl := ""
	userCommand := ""
	if engineId < 0 {
		tmpCodeObsPath := strings.Trim(req.CodeObsPath, "/")
		tmpCodeObsPaths := strings.Split(tmpCodeObsPath, "/")
		lastCodeDir := "code"
		if len(tmpCodeObsPaths) > 0 {
			lastCodeDir = tmpCodeObsPaths[len(tmpCodeObsPaths)-1]
		}
		userCommand = "/bin/bash /home/work/run_train.sh 's3://" + req.CodeObsPath + "' '" + lastCodeDir + "/" + req.BootFile + "' '/tmp/log/train.log' --'data_url'='s3://" + req.DataUrl + "' --'train_url'='s3://" + req.TrainUrl + "'"
		var versionInfos modelarts.VersionInfo
		if err := json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
			log.Info("json parse err." + err.Error())
		} else {
			for _, engine := range versionInfos.Version {
				if engine.ID == engineId {
					userImageUrl = engine.Url
					break
				}
			}
		}
		for _, param := range req.Parameters {
			userCommand += " --'" + param.Label + "'='" + param.Value + "'"
		}
		return userCommand, userImageUrl
	}
	return userCommand, userImageUrl
}

func getInfJobUserCommand(engineId int, req *modelarts.GenerateInferenceJobReq) (string, string) {
	userImageUrl := ""
	userCommand := ""
	if engineId < 0 {
		tmpCodeObsPath := strings.Trim(req.CodeObsPath, "/")
		tmpCodeObsPaths := strings.Split(tmpCodeObsPath, "/")
		lastCodeDir := "code"
		if len(tmpCodeObsPaths) > 0 {
			lastCodeDir = tmpCodeObsPaths[len(tmpCodeObsPaths)-1]
		}
		userCommand = "/bin/bash /home/work/run_train.sh 's3://" + req.CodeObsPath + "' '" + lastCodeDir + "/" + req.BootFile + "' '/tmp/log/train.log' --'data_url'='s3://" + req.DataUrl + "' --'train_url'='s3://" + req.TrainUrl + "'"
		var versionInfos modelarts.VersionInfo
		if err := json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
			log.Info("json parse err." + err.Error())
		} else {
			for _, engine := range versionInfos.Version {
				if engine.ID == engineId {
					userImageUrl = engine.Url
					break
				}
			}
		}
		for _, param := range req.Parameters {
			userCommand += " --'" + param.Label + "'='" + param.Value + "'"
		}
		return userCommand, userImageUrl
	}
	return userCommand, userImageUrl
}

func TrainJobCreateVersion(ctx *context.Context, form auth.CreateModelArtsTrainJobForm) {
	ctx.Data["PageIsTrainJob"] = true
	var jobID = ctx.Params(":jobid")

	spec, err := resource.GetAndCheckSpec(ctx.User.ID, form.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.NPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainTwo})
	if err != nil || spec == nil {
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("Resource specification not available", tplModelArtsTrainJobVersionNew, &form)
		return
	}
	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice, WorkServerNumber: form.WorkServerNumber}) {
		log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr("points.insufficient_points_balance"), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: form.DisplayJobName, JobType: string(models.JobTypeTrain)}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr(errMsg), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	errStr := checkMultiNode(ctx.User.ID, form.WorkServerNumber)
	if errStr != "" {
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr(errStr), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	if err != nil {
		log.Error("GetCloudbrainTrainJobCountByUserID failed:%v", err, ctx.Data["MsgID"])
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("system error", tplModelArtsTrainJobVersionNew, &form)
		return
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
			trainJobNewVersionDataPrepare(ctx)
			ctx.RenderWithErr("you have already a running or waiting task, can not create more", tplModelArtsTrainJobVersionNew, &form)
			return
		}
	}

	latestTask, err := models.GetCloudbrainByJobIDAndIsLatestVersion(jobID, modelarts.IsLatestVersion)
	if err != nil {
		ctx.ServerError("GetCloudbrainByJobIDAndIsLatestVersion faild:", err)
		return
	}

	if latestTask.Cleared {
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.train_result_cleared"), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	VersionOutputPath := modelarts.GetOutputPathByCount(latestTask.TotalVersionCount + 1)

	displayJobName := form.DisplayJobName
	jobName := form.JobName
	uuid := form.Attachment
	description := form.Description
	workServerNumber := form.WorkServerNumber
	engineID := form.EngineID
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params
	poolID := form.PoolID
	//isSaveParam := form.IsSaveParam
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + modelarts.CodePath
	codeObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.CodePath + VersionOutputPath + "/"
	outputObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.OutputPath + VersionOutputPath + "/"
	logObsPath := "/" + setting.Bucket + modelarts.JobPath + jobName + modelarts.LogPath + VersionOutputPath + "/"

	branchName := form.BranchName
	PreVersionName := form.VersionName
	FlavorName := form.FlavorName
	EngineName := form.EngineName
	isLatestVersion := modelarts.IsLatestVersion

	canNewJob, _ := canUserCreateTrainJobVersion(ctx, latestTask.UserID)
	if !canNewJob {
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("user cann't new trainjob", tplModelArtsTrainJobVersionNew, &form)
		return
	}

	if err := paramCheckCreateTrainJob(form); err != nil {
		log.Error("paramCheckCreateTrainJob failed:(%v)", err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	bootFileExist, err := ctx.Repo.FileExists(bootFile, branchName)
	if err != nil || !bootFileExist {
		log.Error("Get bootfile error:", err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_bootfile_err"), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	//todo: del the codeLocalPath
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	gitRepo, _ := git.OpenRepository(repo.RepoPath())
	commitID, _ := gitRepo.GetBranchCommitID(branchName)
	if err := downloadCode(repo, codeLocalPath, branchName); err != nil {
		log.Error("Failed git clone repo to local(！: %s (%v)", repo.FullName(), err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	//todo: upload code (send to file_server todo this work?)
	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.OutputPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_output: %s (%v)", repo.FullName(), err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("Failed to obsMkdir_output", tplModelArtsTrainJobVersionNew, &form)
		return
	}

	if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.LogPath + VersionOutputPath + "/"); err != nil {
		log.Error("Failed to obsMkdir_log: %s (%v)", repo.FullName(), err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("Failed to obsMkdir_log", tplModelArtsTrainJobVersionNew, &form)
		return
	}

	parentDir := VersionOutputPath + "/"
	// parentDir := ""
	// if err := uploadCodeToObs(codeLocalPath, jobName, ""); err != nil {
	if err := uploadCodeToObs(codeLocalPath, jobName, parentDir); err != nil {
		log.Error("Failed to uploadCodeToObs: %s (%v)", repo.FullName(), err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tplModelArtsTrainJobVersionNew, &form)
		return
	}

	//todo: del local code?

	var parameters models.Parameters
	param := make([]models.Parameter, 0)
	existDeviceTarget := false
	if len(params) != 0 {
		err := json.Unmarshal([]byte(params), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", params, err)
			trainJobNewVersionDataPrepare(ctx)
			ctx.RenderWithErr("运行参数错误", tplModelArtsTrainJobVersionNew, &form)
			return
		}
		for _, parameter := range parameters.Parameter {
			if parameter.Label == modelarts.DeviceTarget {
				existDeviceTarget = true
			}
			if parameter.Label != modelarts.TrainUrl && parameter.Label != modelarts.DataUrl {
				param = append(param, models.Parameter{
					Label: parameter.Label,
					Value: parameter.Value,
				})
			}
		}
	}
	if !existDeviceTarget {
		param = append(param, models.Parameter{
			Label: modelarts.DeviceTarget,
			Value: modelarts.Ascend,
		})
	}

	datasUrlList, dataUrl, datasetNames, _, err := getDatasUrlListByUUIDS(uuid)
	if err != nil {
		log.Error("Failed to getDatasUrlListByUUIDS: %v", err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("Failed to getDatasUrlListByUUIDS:"+err.Error(), tplModelArtsTrainJobVersionNew, &form)
		return
	}
	dataPath := dataUrl
	jsondatas, err := json.Marshal(datasUrlList)
	if err != nil {
		log.Error("Failed to Marshal: %v", err)
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr("json error:"+err.Error(), tplModelArtsTrainJobVersionNew, &form)
		return
	}
	// if isMultiDataset {
	param = append(param, models.Parameter{
		Label: modelarts.MultiDataUrl,
		Value: string(jsondatas),
	})
	// }

	if form.ModelName != "" { //使用预训练模型训练
		param = addModelUrlParam(param, form.PreTrainModelUrl, form.CkptName)
	}

	if form.IsContinue { // qizhi NPU 继续训练，将旧任务输出文件拷贝至新任务输出路径
		srcPath := path.Join(setting.TrainJobModelPath, jobName, modelarts.OutputPath, form.VersionName, "") + "/"
		destPath := path.Join(setting.TrainJobModelPath, jobName, modelarts.OutputPath, VersionOutputPath, "") + "/"
		err := ObsCopyResults(srcPath, destPath)
		if err != nil {
			log.Error("Copy Prev Task Result Files failed:", err)
			trainJobNewVersionDataPrepare(ctx)
			ctx.RenderWithErr("Failed to copy output files from previous train job", tplModelArtsTrainJobVersionNew, &form)
			return
		}
	}

	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, PreVersionName)
	if err != nil {
		log.Error("GetCloudbrainByJobIDAndVersionName(%s) failed:%v", jobID, err.Error())
		ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobVersionNew, &form)
		return
	}
	req := &modelarts.GenerateTrainJobReq{
		JobName:           jobName,
		DisplayJobName:    displayJobName,
		DataUrl:           dataPath,
		Description:       description,
		CodeObsPath:       codeObsPath,
		BootFileUrl:       codeObsPath + bootFile,
		BootFile:          bootFile,
		TrainUrl:          outputObsPath,
		WorkServerNumber:  workServerNumber,
		IsLatestVersion:   isLatestVersion,
		EngineID:          int64(engineID),
		LogUrl:            logObsPath,
		PoolID:            poolID,
		Uuid:              uuid,
		Params:            form.Params,
		Parameters:        param,
		PreVersionId:      task.VersionID,
		CommitID:          commitID,
		BranchName:        branchName,
		FlavorName:        FlavorName,
		EngineName:        EngineName,
		PreVersionName:    PreVersionName,
		TotalVersionCount: latestTask.TotalVersionCount + 1,
		DatasetName:       datasetNames,
		Spec:              spec,
	}

	if form.ModelName != "" { //使用预训练模型训练
		req.ModelName = form.ModelName
		req.LabelName = form.LabelName
		req.CkptName = form.CkptName
		req.ModelId = form.ModelId
		req.ModelVersion = form.ModelVersion
		req.PreTrainModelUrl = form.PreTrainModelUrl

	}
	userCommand, userImageUrl := getUserCommand(engineID, req)
	req.UserCommand = userCommand
	req.UserImageUrl = userImageUrl

	err = modelarts.GenerateTrainJobVersion(ctx, req, jobID)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error())
		trainJobNewVersionDataPrepare(ctx)
		ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobVersionNew, &form)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job/" + jobID)
	// ctx.HTML(http.StatusOK, tplModelArtsTrainJobShow)
}

func ObsCopyResults(srcPath string, destPath string) error {
	log.Info("prev task obs path:", setting.Bucket+srcPath)
	log.Info("current task obs path:", setting.Bucket+destPath)
	allfile, _ := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, srcPath)
	var fileNames []string
	for _, file := range allfile {
		if strings.Contains(file.FileName, "README") || strings.Contains(file.FileName, ".txt") {
			continue
		}
		fileNames = append(fileNames, file.FileName)
	}
	log.Info("Previous task all files", fileNames)
	fileSizeAll, err := storage.ObsCopyManyFile(setting.Bucket, srcPath, setting.Bucket, destPath, fileNames)
	log.Info("%v output files copied from previous task", fileSizeAll)

	return err

}

// readDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	list, err := f.Readdir(0)
	f.Close()
	if err != nil {
		//todo: can not upload empty folder
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	//sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

func uploadCodeToObs(codePath, jobName, parentDir string) error {
	files, err := readDir(codePath)
	if err != nil {
		log.Error("readDir(%s) failed: %s", codePath, err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			input := &obs.PutObjectInput{}
			input.Bucket = setting.Bucket
			input.Key = parentDir + file.Name() + "/"
			_, err = storage.ObsCli.PutObject(input)
			if err != nil {
				log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
				return err
			}

			if err = uploadCodeToObs(codePath+file.Name()+"/", jobName, parentDir+file.Name()+"/"); err != nil {
				log.Error("uploadCodeToObs(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		} else {
			input := &obs.PutFileInput{}
			input.Bucket = setting.Bucket
			input.Key = setting.CodePathPrefix + jobName + "/code/" + parentDir + file.Name()
			input.SourceFile = codePath + file.Name()
			_, err = storage.ObsCli.PutFile(input)
			if err != nil {
				log.Error("PutFile(%s) failed: %s", input.SourceFile, err.Error())
				return err
			}
		}
	}

	return nil
}

func obsMkdir(dir string) error {
	input := &obs.PutObjectInput{}
	input.Bucket = setting.Bucket
	input.Key = dir
	_, err := storage.ObsCli.PutObject(input)
	if err != nil {
		log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
		return err
	}

	return nil
}

func paramCheckCreateTrainJob(form auth.CreateModelArtsTrainJobForm) error {
	if !strings.HasSuffix(strings.TrimSpace(form.BootFile), ".py") {
		log.Error("the boot file(%s) must be a python file", strings.TrimSpace(form.BootFile))
		return errors.New("启动文件必须是python文件")
	}

	if form.BranchName == "" {
		log.Error("the branch must not be null!", form.BranchName)
		return errors.New("代码分支不能为空！")
	}

	return nil
}

func paramCheckCreateInferenceJob(form auth.CreateModelArtsInferenceJobForm) error {
	if !strings.HasSuffix(strings.TrimSpace(form.BootFile), ".py") {
		log.Error("the boot file(%s) must be a python file", strings.TrimSpace(form.BootFile))
		return errors.New("启动文件必须是python文件")
	}

	if form.WorkServerNumber > 2 || form.WorkServerNumber < 1 {
		log.Error("the WorkServerNumber(%d) must be in (1,2)", form.WorkServerNumber)
		return errors.New("计算节点数必须在1-2之间")
	}

	if form.ModelName == "" {
		log.Error("the ModelName(%d) must not be nil", form.ModelName)
		return errors.New("模型名称不能为空")
	}
	if form.ModelVersion == "" {
		log.Error("the ModelVersion(%d) must not be nil", form.ModelVersion)
		return errors.New("模型版本不能为空")
	}
	if form.CkptName == "" {
		log.Error("the CkptName(%d) must not be nil", form.CkptName)
		return errors.New("权重文件不能为空")
	}
	if form.BranchName == "" {
		log.Error("the Branch(%d) must not be nil", form.BranchName)
		return errors.New("分支名不能为空")
	}

	if utf8.RuneCountInString(form.Description) > 255 {
		log.Error("the Description length(%d) must not more than 255", form.Description)
		return errors.New("描述字符不能超过255个字符")
	}

	return nil
}

func TrainJobShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplModelArtsTrainJobShow)
	/*
		var jobID = ctx.Params(":jobid")

		repo := ctx.Repo.Repository
		page := ctx.QueryInt("page")
		if page <= 0 {
			page = 1
		}

		var jobTypes []string
		jobTypes = append(jobTypes, string(models.JobTypeTrain))
		VersionListTasks, VersionListCount, err := models.CloudbrainsVersionList(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: setting.UI.IssuePagingNum,
			},
			RepoID:   repo.ID,
			Type:     models.TypeCloudBrainTwo,
			JobTypes: jobTypes,
			JobID:    jobID,
		})

		if err != nil {
			log.Error("GetVersionListTasks(%s) failed:%v", jobID, err.Error())
			ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
			return
		}
		if len(VersionListTasks) == 0 {
			ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
			return
		}
		//设置权限
		canNewJob, err := canUserCreateTrainJobVersion(ctx, VersionListTasks[0].UserID)
		if err != nil {
			ctx.ServerError("canNewJob failed", err)
			return
		}
		ctx.Data["canNewJob"] = canNewJob
		datasetList := make([][]*models.DatasetDownload, 0)
		//将运行参数转化为epoch_size = 3, device_target = Ascend的格式
		for i, task := range VersionListTasks {

			var parameters models.Parameters
			if VersionListTasks[i].Parameters != "" {
				err := json.Unmarshal([]byte(VersionListTasks[i].Parameters), &parameters)
				if err != nil {
					log.Error("Failed to Unmarshal Parameters: %s (%v)", VersionListTasks[i].Parameters, err)
				}
			}

			if len(parameters.Parameter) > 0 {
				paramTemp := ""
				for _, Parameter := range parameters.Parameter {
					param := Parameter.Label + " = " + Parameter.Value + "; "
					paramTemp = paramTemp + param
				}
				VersionListTasks[i].Parameters = paramTemp[:len(paramTemp)-2]
			} else {
				VersionListTasks[i].Parameters = ""
			}
			datasetList = append(datasetList, GetCloudBrainDataSetInfo(task.Uuid, task.DatasetName, false))
			VersionListTasks[i].CanDel = cloudbrain.CanDeleteJob(ctx, &task.Cloudbrain)
			VersionListTasks[i].CanModify = cloudbrain.CanModifyJob(ctx, &task.Cloudbrain)
			VersionListTasks[i].ContainerIp = ""
			//add spec
			s, err := resource.GetCloudbrainSpec(task.Cloudbrain.ID)
			if err != nil {
				log.Error("TrainJobShow GetCloudbrainSpec error:" + err.Error())
				continue
			}
			VersionListTasks[i].Cloudbrain.Spec = s
		}

		pager := context.NewPagination(VersionListCount, setting.UI.IssuePagingNum, page, 5)
		pager.SetDefaultParams(ctx)
		ctx.Data["Page"] = pager
		ctx.Data["jobID"] = jobID
		ctx.Data["displayJobName"] = VersionListTasks[0].DisplayJobName
		ctx.Data["version_list_task"] = VersionListTasks
		ctx.Data["version_list_count"] = VersionListCount
		ctx.Data["datasetList"] = datasetList
		ctx.Data["canDownload"] = cloudbrain.CanDownloadJob(ctx, &VersionListTasks[0].Cloudbrain)
		ctx.HTML(http.StatusOK, tplModelArtsTrainJobShow)
	*/
}

func TrainJobDel(ctx *context.Context) {
	var jobID = ctx.Params(":jobid")
	var listType = ctx.Query("listType")
	var id = ctx.QueryInt64("id")
	if id > 0 {
		task, _ := models.GetCloudbrainByCloudbrainID(id)
		if task != nil && task.IsNewAITask() {
			bizErr := ai_task.DelCloudbrain(task)
			if bizErr != nil {
				log.Error("DelCloudbrain(%s) failed:%v err=%v", task.JobName, bizErr)
				ctx.ServerError("DelCloudbrain failed", bizErr.ToError())
				return
			}
			var isAdminPage = ctx.Query("isadminpage")
			var isHomePage = ctx.Query("ishomepage")
			if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
				ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
			} else if isHomePage == "true" {
				ctx.Redirect(setting.AppSubURL + "/cloudbrains")
			} else {
				ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=" + listType)
			}
			return
		}

	}
	repo := ctx.Repo.Repository

	var jobTypes []string
	jobTypes = append(jobTypes, string(models.JobTypeTrain))
	VersionListTasks, _, err := models.CloudbrainsVersionList(&models.CloudbrainsOptions{
		RepoID:   repo.ID,
		Type:     models.TypeCloudBrainTwo,
		JobTypes: jobTypes,
		JobID:    jobID,
	})
	if err != nil {
		ctx.ServerError("get VersionListTasks failed", err)
		return
	}

	for _, task := range VersionListTasks {
		if !task.IsTerminal() {
			log.Error("the job(%s) version has not been stopped", task.JobName)
			ctx.RenderWithErr("the job version has not been stopped", tplModelArtsTrainJobIndex, nil)
			return
		}
	}

	//删除modelarts上的任务记录
	_, err = modelarts.DelTrainJob(jobID)
	if err != nil {
		log.Error("DelTrainJob(%s) failed:%v", jobID, err.Error())
		if err.Error() == "1" {
			ctx.Flash.Error(ctx.Tr("deployment.deletion_notice_trainjob"))
			ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=" + listType)
		}
		return
	}

	//删除数据库Cloudbrain表的记录
	for _, task := range VersionListTasks {
		err = models.DeleteJob(&task.Cloudbrain)
		if err != nil {
			ctx.ServerError("DeleteJob failed", err)
			return
		}
	}

	//删除存储
	if len(VersionListTasks) > 0 {
		DeleteJobStorage(VersionListTasks[0].JobName)
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=" + listType)
	}
}

func TrainJobStop(ctx *context.Context) {
	var jobID = ctx.Params(":jobid")
	var listType = ctx.Query("listType")
	task := ctx.Cloudbrain

	_, err := modelarts.StopTrainJob(jobID, strconv.FormatInt(task.VersionID, 10))
	if err != nil {
		log.Error("StopTrainJob(%s) failed:%v", task.JobName, err.Error())
		ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobIndex, nil)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=" + listType)
}

func canUserCreateTrainJobVersion(ctx *context.Context, userID int64) (bool, error) {
	if ctx == nil || ctx.User == nil {
		log.Error("user unlogin!")
		return false, nil
	}
	if userID == ctx.User.ID || ctx.User.IsAdmin {
		return true, nil
	} else {
		log.Error("Only user itself and admin can new trainjob!")
		return false, nil
	}
}

func TrainJobGetConfigList(ctx *context.Context) {
	ctx.Data["PageIsTrainJob"] = true

	var jobID = ctx.Params(":jobid")
	var logFileName = ctx.Query("file_name")
	var baseLine = ctx.Query("base_line")
	var order = ctx.Query("order")

	if order != modelarts.OrderDesc && order != modelarts.OrderAsc {
		log.Error("order(%s) check failed", order)
		ctx.HTML(http.StatusBadRequest, tplModelArtsTrainJobShow)
		return
	}

	task, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", jobID, err.Error())
		ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobShow, nil)
		return
	}

	result, err := modelarts.GetTrainJobLog(jobID, strconv.FormatInt(task.VersionID, 10), baseLine, logFileName, order, modelarts.Lines)
	if err != nil {
		log.Error("GetTrainJobLog(%s) failed:%v", jobID, err.Error())
		ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobShow, nil)
		return
	}

	ctx.Data["log"] = result
	//ctx.HTML(http.StatusOK, tplModelArtsTrainJobShow)
}

func getConfigList(perPage, page int, sortBy, order, searchContent, configType string) (*models.GetConfigListResult, error) {
	var result models.GetConfigListResult

	list, err := modelarts.GetConfigList(perPage, page, sortBy, order, searchContent, configType)
	if err != nil {
		log.Error("GetConfigList failed:", err)
		return &result, err
	}

	for _, config := range list.ParaConfigs {
		paraConfig, err := modelarts.GetParaConfig(config.ConfigName, configType)
		if err != nil {
			log.Error("GetParaConfig failed:", err)
			return &result, err
		}

		config.Result = paraConfig
	}

	return list, nil
}

func checkModelArtsSpecialPool(ctx *context.Context, flavorCode string, jobType string) string {
	if modelarts.SpecialPools != nil {
		isMatchPool := false

		for _, specialPool := range modelarts.SpecialPools.Pools {
			if cloudbrain.IsElementExist(specialPool.JobType, jobType) {
				if isInOrg, _ := models.IsOrganizationMemberByOrgName(specialPool.Org, ctx.User.ID); isInOrg {
					isMatchPool = true
					isMatchSpec := false
					for _, flavor := range specialPool.Flavor {
						if flavor.Value == flavorCode {
							isMatchSpec = true
							break
						}
					}

					if !isMatchSpec {
						return "cloudbrain.wrong_specification"

					}

				}

			}
		}

		if !isMatchPool {
			isMatchSpec := false
			if jobType == string(models.JobTypeDebug) {
				for _, flavor := range setting.StFlavorInfo.FlavorInfo {
					if flavor.Value == flavorCode {
						isMatchSpec = true
						break
					}
				}
			} else {

				var flavorInfos modelarts.Flavor
				json.Unmarshal([]byte(setting.TrainJobFLAVORINFOS), &flavorInfos)

				for _, flavor := range flavorInfos.Info {
					if flavor.Code == flavorCode {
						isMatchSpec = true
						break
					}
				}
			}

			if !isMatchSpec {

				return "cloudbrain.wrong_specification"
			}

		}

	}
	return ""
}
func InferenceJobIndex(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplModelArtsInferenceJobIndex)
}
func InferenceJobNew(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplModelArtsInferenceJobNew)
}
func inferenceJobNewDataPrepare(ctx *context.Context) error {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["newInference"] = true

	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName

	attachs, err := models.GetModelArtsTrainAttachments(ctx.User.ID)
	if err != nil {
		ctx.ServerError("GetAllUserAttachments failed:", err)
		return err
	}
	ctx.Data["attachments"] = attachs

	var resourcePools modelarts.ResourcePool
	if err = json.Unmarshal([]byte(setting.ResourcePools), &resourcePools); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["resource_pools"] = resourcePools.Info

	var engines modelarts.Engine
	if err = json.Unmarshal([]byte(setting.Engines), &engines); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engines"] = engines.Info

	var versionInfos modelarts.VersionInfo
	if err = json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engine_versions"] = versionInfos.Version

	prepareCloudbrainTwoInferenceSpecs(ctx)

	ctx.Data["params"] = ""
	ctx.Data["branchName"] = ctx.Repo.BranchName

	configList, err := getConfigList(modelarts.PerPage, 1, modelarts.SortByCreateTime, "desc", "", modelarts.ConfigTypeCustom)
	if err != nil {
		ctx.ServerError("getConfigList failed:", err)
		return err
	}
	ctx.Data["config_list"] = configList.ParaConfigs
	isQueryPrivate := isQueryPrivateModel(ctx)
	repoId := ctx.Repo.Repository.ID
	Type := -1
	_, model_count, _ := models.QueryModel(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     1,
			PageSize: 2,
		},
		RepoID:         repoId,
		Type:           Type,
		New:            MODEL_LATEST,
		IsOnlyThisRepo: true,
		Status:         0,
		IsQueryPrivate: isQueryPrivate,
	})
	ctx.Data["MODEL_COUNT"] = model_count
	ctx.Data["datasetType"] = models.TypeCloudBrainTwo
	waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeCloudBrainTwo, "")
	ctx.Data["WaitCount"] = waitCount
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeInference))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount

	return nil
}

func prepareCloudbrainTwoInferenceSpecs(ctx *context.Context) {
	noteBookSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeInference,
		ComputeResource: models.NPU,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainTwo,
	})
	ctx.Data["Specs"] = noteBookSpecs
}

func inferenceJobErrorNewDataPrepare(ctx *context.Context, form auth.CreateModelArtsInferenceJobForm) error {
	ctx.Data["PageIsCloudBrain"] = true

	t := time.Now()
	var jobName = "inference" + t.Format("2006010215") + strconv.Itoa(int(t.Unix()))[5:]
	ctx.Data["job_name"] = jobName

	attachs, err := models.GetModelArtsTrainAttachments(ctx.User.ID)
	if err != nil {
		ctx.ServerError("GetAllUserAttachments failed:", err)
		return err
	}
	ctx.Data["attachments"] = attachs

	var resourcePools modelarts.ResourcePool
	if err = json.Unmarshal([]byte(setting.ResourcePools), &resourcePools); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["resource_pools"] = resourcePools.Info

	var engines modelarts.Engine
	if err = json.Unmarshal([]byte(setting.Engines), &engines); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engines"] = engines.Info

	var versionInfos modelarts.VersionInfo
	if err = json.Unmarshal([]byte(setting.EngineVersions), &versionInfos); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["engine_versions"] = versionInfos.Version
	prepareCloudbrainTwoInferenceSpecs(ctx)

	configList, err := getConfigList(modelarts.PerPage, 1, modelarts.SortByCreateTime, "desc", "", modelarts.ConfigTypeCustom)
	if err != nil {
		ctx.ServerError("getConfigList failed:", err)
		return err
	}
	var Parameters modelarts.Parameters
	if err = json.Unmarshal([]byte(form.Params), &Parameters); err != nil {
		ctx.ServerError("json.Unmarshal failed:", err)
		return err
	}
	ctx.Data["params"] = Parameters.Parameter
	ctx.Data["config_list"] = configList.ParaConfigs
	ctx.Data["bootFile"] = form.BootFile
	ctx.Data["uuid"] = form.Attachment
	_, datasetNames, err := models.GetDatasetInfo(form.Attachment)
	if err != nil {
		log.Error("GetDatasetInfo failed: %v", err, ctx.Data["MsgID"])
		return nil
	}
	ctx.Data["dataset_name"] = datasetNames
	ctx.Data["branch_name"] = form.BranchName
	ctx.Data["model_name"] = form.ModelName
	ctx.Data["model_version"] = form.ModelVersion
	ctx.Data["ckpt_name"] = form.CkptName
	ctx.Data["model_id"] = form.ModelId
	ctx.Data["pre_train_model_url"] = form.PreTrainModelUrl
	ctx.Data["datasetType"] = models.TypeCloudBrainTwo
	waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeCloudBrainTwo, "")
	ctx.Data["WaitCount"] = waitCount
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeInference))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount
	return nil
}
func InferenceJobShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplModelArtsInferenceJobShow)
}

func MultiModelDownload(ctx *context.Context) {
	var (
		err error
	)
	jobID := ctx.Params(":jobid")
	versionName := ctx.Query("version_name")
	parentDir := ctx.Query("parent_dir")

	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobIDAndVersionName(%s) failed:%v", task.JobName, err.Error())
		return
	}

	if task.ComputeResource == models.NPUResource {
		path := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, setting.OutPutPath, versionName, parentDir), "/")
		path = strings.TrimSuffix(path, "/")
		path += "/"
		allFile, err := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, path)
		if err == nil {
			returnFileName := task.DisplayJobName + ".zip"
			ObsDownloadManyFile(path, ctx, returnFileName, allFile)
		} else {
			log.Info("error,msg=" + err.Error())
			ctx.ServerError("no file to download.", err)
		}
	} else if task.ComputeResource == models.GPUResource || task.ComputeResource == models.GCUResource {
		filePath := setting.CBCodePathPrefix + task.JobName + cloudbrain.ModelMountPath + "/" + parentDir
		allFile, err := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, filePath)
		if err == nil {
			returnFileName := task.DisplayJobName + ".zip"
			MinioDownloadManyFile(filePath, ctx, returnFileName, allFile)
		} else {
			log.Info("error,msg=" + err.Error())
			ctx.ServerError("no file to download.", err)
		}
	}

}

func ModelDownload(ctx *context.Context) {
	var (
		err error
	)

	jobID := ctx.Params(":jobid")
	versionName := ctx.Query("version_name")
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobIDAndVersionName(%s) failed:%v", task.JobName, err.Error())
		return
	}

	var url string
	if task.ComputeResource == models.NPUResource {
		path := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, setting.OutPutPath, versionName, parentDir, fileName), "/")
		url, err = storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, path)
		if err != nil {
			log.Error("GetObsCreateSignedUrl failed: %v", err.Error(), ctx.Data["msgID"])
			ctx.ServerError("GetObsCreateSignedUrl", err)
			return
		}
	} else if task.ComputeResource == models.GPUResource || task.ComputeResource == models.GCUResource {
		filePath := setting.CBCodePathPrefix + task.JobName + cloudbrain.ModelMountPath + "/" + parentDir
		url, err = storage.Attachments.PresignedGetURL(filePath, fileName)
		if err != nil {
			log.Error("PresignedGetURL failed: %v", err.Error(), ctx.Data["msgID"])
			ctx.ServerError("PresignedGetURL", err)
			return
		}
	}

	ctx.Resp.Header().Set("Cache-Control", "max-age=0")
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusTemporaryRedirect)
}

func ResultDownload(ctx *context.Context) {
	var (
		err error
	)

	versionName := ctx.Query("version_name")
	parentDir := ctx.Query("parent_dir")
	fileName := ctx.Query("file_name")
	task := ctx.Cloudbrain
	if err != nil {
		ctx.Data["error"] = err.Error()
	}
	path := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, "result/", versionName, parentDir, fileName), "/")

	url, err := storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, path)
	if err != nil {
		log.Error("GetObsCreateSignedUrl failed: %v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("GetObsCreateSignedUrl", err)
		return
	}
	ctx.Resp.Header().Set("Cache-Control", "max-age=0")
	http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
}
func DeleteJobStorage(jobName string) error {
	//delete local
	localJobPath := setting.JobPath + jobName
	err := os.RemoveAll(localJobPath)
	if err != nil {
		log.Error("RemoveAll(%s) failed:%v", localJobPath, err)
	}

	//delete oss
	dirPath := setting.CodePathPrefix + jobName + "/"
	err = storage.ObsRemoveObject(setting.Bucket, dirPath)
	if err != nil {
		log.Error("ObsRemoveObject(%s) failed:%v", localJobPath, err)
	}

	return nil
}

func DownloadMultiResultFile(ctx *context.Context) {
	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		return
	}
	// if !isCanDeleteOrDownload(ctx, task) {
	// 	ctx.ServerError("no right.", errors.New(ctx.Tr("repo.model_noright")))
	// 	return
	// }

	// path := Model_prefix + models.AttachmentRelativePath(id) + "/"
	path := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, "result/", versionName), "/") + "/"

	allFile, err := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, path)
	if err == nil {
		//count++
		// models.ModifyModelDownloadCount(id)

		returnFileName := task.DisplayJobName + ".zip"
		ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+returnFileName)
		ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
		w := zip.NewWriter(ctx.Resp)
		defer w.Close()
		for _, oneFile := range allFile {
			if oneFile.IsDir {
				log.Info("zip dir name:" + oneFile.FileName)
			} else {
				log.Info("zip file name:" + oneFile.FileName)
				fDest, err := w.Create(oneFile.FileName)
				if err != nil {
					log.Info("create zip entry error, download file failed: %s\n", err.Error())
					ctx.ServerError("download file failed:", err)
					return
				}
				body, err := storage.ObsDownloadAFile(setting.Bucket, path+oneFile.FileName)
				if err != nil {
					log.Info("download file failed: %s\n", err.Error())
					ctx.ServerError("download file failed:", err)
					return
				} else {
					defer body.Close()
					p := make([]byte, 1024)
					var readErr error
					var readCount int
					// 读取对象内容
					for {
						readCount, readErr = body.Read(p)
						if readCount > 0 {
							fDest.Write(p[:readCount])
						}
						if readErr != nil {
							break
						}
					}
				}
			}
		}
	} else {
		log.Info("error,msg=" + err.Error())
		ctx.ServerError("no file to download.", err)
	}
}

func SetJobCount(ctx *context.Context) {
	repoId := ctx.Repo.Repository.ID
	_, jobCount, err := models.Cloudbrains(&models.CloudbrainsOptions{
		RepoID: repoId,
		Type:   models.TypeCloudBrainAll,
	})
	if err != nil {
		ctx.ServerError("Get job faild:", err)
		return
	}
	ctx.Data["jobCount"] = jobCount
}

func TrainJobDownloadLogFile(ctx *context.Context) {
	var (
		err error
	)

	var jobID = ctx.Params(":jobid")
	versionName := ctx.Query("version_name")
	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobIDAndVersionName(%s) failed:%v", task.JobName, err.Error(), ctx.Data["msgID"])
		ctx.ServerError("GetCloudbrainByJobIDAndVersionName", err)
		return
	}

	prefix := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, modelarts.LogPath, versionName), "/") + "/job"
	key, err := storage.GetObsLogFileName(prefix)
	if err != nil {
		log.Error("GetObsLogFileName(%s) failed:%v", jobID, err.Error(), ctx.Data["msgID"])
		ctx.ServerError("GetObsLogFileName", err)
		return
	}
	if len(key) > 1 {
		ObsDownloadManyFile(prefix[0:len(prefix)-3], ctx, task.DisplayJobName+".zip", key)
	} else {
		url, err := storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, key[0].ParenDir+key[0].FileName)
		if err != nil {
			log.Error("GetObsCreateSignedUrlByBucketAndKey failed: %v", err.Error(), ctx.Data["msgID"])
			ctx.ServerError("GetObsCreateSignedUrlByBucketAndKey", err)
			return
		}
		ctx.Resp.Header().Set("Cache-Control", "max-age=0")
		http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusTemporaryRedirect)
	}

}
func getDatasUrlListByUUIDS(uuidStr string) ([]models.Datasurl, string, string, bool, error) {
	var isMultiDataset bool
	var dataUrl string
	var datasetNames string
	var datasUrlList []models.Datasurl
	uuids := strings.Split(uuidStr, ";")
	if len(uuids) > setting.MaxDatasetNum {
		log.Error("the dataset count(%d) exceed the limit", len(uuids))
		return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("the dataset count exceed the limit")
	}

	datasetInfos := make(map[string]models.DatasetInfo)
	attachs, err := models.GetAttachmentsByUUIDs(uuids)
	if err != nil || len(attachs) != len(uuids) {
		log.Error("GetAttachmentsByUUIDs failed: %v", err)
		return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("GetAttachmentsByUUIDs failed")
	}

	for i, tmpUuid := range uuids {
		var attach *models.Attachment
		for _, tmpAttach := range attachs {
			if tmpAttach.UUID == tmpUuid {
				attach = tmpAttach
				break
			}
		}
		if attach == nil {
			log.Error("GetAttachmentsByUUIDs failed: %v", err)
			return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("GetAttachmentsByUUIDs failed")
		}
		fileName := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(attach.Name, ".zip"), ".tar.gz"), ".tgz")
		for _, datasetInfo := range datasetInfos {
			if fileName == datasetInfo.Name {
				log.Error("the dataset name is same: %v", attach.Name)
				return datasUrlList, dataUrl, datasetNames, isMultiDataset, errors.New("the dataset name is same")
			}
		}

		//dataUrl = "/" + setting.Bucket + "/" + setting.BasePath + path.Join(attachs[0].UUID[0:1], attachs[0].UUID[1:2]) + "/" + attachs[0].UUID + attachs[0].UUID + "/"
		dataUrl = "/" + setting.Bucket + "/" + attachs[0].UnzipPath
		//datasetUrl := "s3://" + setting.Bucket + "/" + setting.BasePath + path.Join(attach.UUID[0:1], attach.UUID[1:2]) + "/" + attach.UUID + attach.UUID + "/"
		datasetUrl := "s3://" + setting.Bucket + "/" + attach.UnzipPath
		datasUrlList = append(datasUrlList, models.Datasurl{
			DatasetUrl:  datasetUrl,
			DatasetName: fileName,
		})

		if i == 0 {
			datasetNames = attach.Name
		} else {
			datasetNames += ";" + attach.Name
		}
	}

	return datasUrlList, dataUrl, datasetNames, isMultiDataset, nil
}
