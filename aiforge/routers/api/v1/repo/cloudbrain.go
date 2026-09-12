// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/task"
	"code.gitea.io/gitea/services/aim"
	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/services/reward/point/account"

	"code.gitea.io/gitea/modules/grampus"

	"code.gitea.io/gitea/modules/convert"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
	repo_service "code.gitea.io/gitea/services/repository"

	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/storage"
	routerRepo "code.gitea.io/gitea/routers/repo"
	ai_task "code.gitea.io/gitea/services/ai_task_service/task"
)

func CloudBrainShow(ctx *context.APIContext) {

	task, err := models.GetCloudbrainByJobID(ctx.Params(":jobid"))

	if err != nil {
		log.Info("error:" + err.Error())
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("repo.cloudbrain_query_fail"))
		return
	}
	cloudbrainTask.PrepareSpec4Show(task)
	task.ContainerIp = ""
	if cloudbrainTask.IsTaskNotStop(task) {
		cloudbrainTask.SyncTaskStatus(task)
	}

	if task.TrainJobDuration == "" {
		if task.Duration == 0 {
			var duration int64
			if task.Status == string(models.JobWaiting) {
				duration = 0
			} else if task.Status == string(models.JobRunning) {
				duration = time.Now().Unix() - int64(task.CreatedUnix)
			} else {
				duration = int64(task.UpdatedUnix) - int64(task.CreatedUnix)
			}
			task.Duration = duration
		}
		task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)
	}
	//to unify  image output
	if task.Type == models.TypeCloudBrainTwo || task.Type == models.TypeCDCenter {
		task.ImageID = strconv.FormatInt(task.EngineID, 10)
		task.Image = task.EngineName

	} else if task.Type == models.TypeC2Net {
		task.Image = task.EngineName
	}
	task.AiCenter = cloudbrainService.GetAiCenterShow(task.AiCenter, ctx.Context)

	ctx.JSON(http.StatusOK, models.BaseMessageWithDataApi{Code: 0, Message: "", Data: convert.ToCloudBrain(task)})

}
func ReposCanCreateCloudbrainJob(ctx *context.APIContext) {
	opts := &models.SearchRepoOptions{
		ListOptions: models.ListOptions{
			Page:     1,
			PageSize: 1000,
		},
		Actor:       ctx.User,
		OwnerID:     ctx.QueryInt64("uid"),
		Private:     ctx.IsSigned && (ctx.Query("private") == "" || ctx.QueryBool("private")),
		Collaborate: util.OptionalBoolNone,
		Template:    util.OptionalBoolNone,
		Archived:    util.OptionalBoolFalse,
	}

	var sortMode = ctx.Query("sort")
	if len(sortMode) > 0 {
		var sortOrder = ctx.Query("order")
		if len(sortOrder) == 0 {
			sortOrder = "asc"
		}
		if searchModeMap, ok := searchOrderByMap[sortOrder]; ok {
			if orderBy, ok := searchModeMap[sortMode]; ok {
				opts.OrderBy = orderBy
			} else {
				ctx.Error(http.StatusUnprocessableEntity, "", fmt.Errorf("Invalid sort mode: \"%s\"", sortMode))
				return
			}
		} else {
			ctx.Error(http.StatusUnprocessableEntity, "", fmt.Errorf("Invalid sort order: \"%s\"", sortOrder))
			return
		}
	}

	var err error
	repos, _, err := models.SearchRepository(opts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.SearchError{
			OK:    false,
			Error: err.Error(),
		})
		return
	}
	filterType := models.UnitTypeCloudBrain
	if ctx.Query("type") == "model" {
		filterType = models.UnitTypeModelManage
	} else if ctx.Query("type") == "dataset" {
		filterType = models.UnitTypeDatasets
	}

	results := make([]*api.Repository, 0)
	for _, repo := range repos {
		if err = repo.GetOwner(); err != nil {
			ctx.JSON(http.StatusInternalServerError, api.SearchError{
				OK:    false,
				Error: err.Error(),
			})
			return
		}

		if repo.UnitEnabled(filterType) {
			if filterType == models.UnitTypeDatasets && models.IsDatasetExistInRepo(repo.ID) {
				continue
			}
			accessMode, err := models.AccessLevel(ctx.User, repo)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, api.SearchError{
					OK:    false,
					Error: err.Error(),
				})
			}
			if accessMode >= models.AccessModeWrite {

				results = append(results, repo.APIFormat(accessMode))
			}
		}

	}

	ctx.JSON(http.StatusOK, api.SearchResults{
		OK:   true,
		Data: results,
	})

}

func GeneralCloudBrainJobStop(ctx *context.APIContext) {
	task := ctx.Cloudbrain
	if task.IsNewAITask() {
		_, bizErr := ai_task.StopCloudbrain(task)
		if bizErr != nil {
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.Stopped_failed")))
			return
		}
		ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
	}

	if task.IsTerminal() {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("cloudbrain.Already_stopped"))
		return
	}
	var err error

	if ctx.Cloudbrain.Type == models.TypeCloudBrainOne {
		err = cloudbrain.StopJob(task.JobID)
	} else if ctx.Cloudbrain.Type == models.TypeCloudBrainTwo {
		_, err = modelarts.StopTrainJob(task.JobID, strconv.FormatInt(task.VersionID, 10))
	} else {
		_, err = grampus.StopJob(task.JobID)
	}

	if err != nil {
		log.Warn("cloud brain stopped failed.", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.Stopped_failed")))
		return
	}

	ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
}
func CreateFileNoteBook(ctx *context.APIContext, option api.CreateFileNotebookJobOption) {
	cloudbrainTask.FileNotebookCreate(ctx.Context, option)
	if ctx.Written() {
		return
	}
	CreateFileNotebookTask(ctx.Context, option)
}

func CreateModelNoteBook(ctx *context.APIContext, option api.CreateModelNoteBook) {
	log.Info("repoId 1=" + fmt.Sprint(option.RepoId) + " modelId=" + option.ModelId)
	re := make(map[string]interface{}, 0)
	oldRepoId := option.RepoId

	oldRepo, err := models.GetRepositoryByID(oldRepoId)
	if err != nil {
		log.Info("get repo err=" + err.Error())
		ctx.JSON(http.StatusOK, re)
		return
	}

	existRepo, _ := models.GetRepositoryByName(ctx.User.ID, oldRepo.Name)
	if existRepo == nil {
		_, err := repo_service.ForkRepository(ctx.User, ctx.User, oldRepo, oldRepo.Name, oldRepo.Description, oldRepo.Alias)
		if err != nil {
			log.Info("fork repository error.err=" + err.Error())
			ctx.JSON(http.StatusOK, re)
			return
		}
	}
	re["repoPath"] = ctx.User.Name + "/" + oldRepo.Name
	ctx.JSON(http.StatusOK, re)
}

func CreateFileNotebookTask(ctx *context.Context, option api.CreateFileNotebookJobOption) {
	displayJobName := cloudbrainService.GetDisplayJobName(ctx.User.Name)
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	jobType := models.JobTypeDebug
	specId := setting.FileNoteBook.SpecIdGPU
	ComputeSource := models.GPU
	imageUrl := setting.FileNoteBook.ImageGPU
	imageId := ""
	imageName := imageUrl
	cluster := entity.C2Net

	if option.Type == 0 {
		specId = setting.FileNoteBook.SpecIdCPU
		imageName = imageUrl
	}
	if option.Type > cloudbrainTask.GPUType {
		imageId = option.ImageId
		imageName = option.Image
		imageUrl = ""
		ComputeSource = models.NPU
		specId = setting.FileNoteBook.SpecIdNPU
	}

	sourceRepo, err := models.GetRepositoryByOwnerAndName(option.OwnerName, option.ProjectName)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_file_not_exist")))
		return
	}
	var traceErr error

	traceContext, _ := otel.StartTraceParent(&entity.TraceInfo{TaskName: displayJobName, JobType: string(jobType), SpanName: "CreateAITask"}, "POST")

	defer func() {

		otel.FinalizeSpan(traceContext, traceErr)
	}()

	res, bizErr := task.CreateAITask(entity.CreateReq{
		JobType:               jobType,
		DisplayJobName:        displayJobName,
		JobName:               jobName,
		SpecId:                specId,
		ComputeSourceStr:      ComputeSource,
		Cluster:               cluster,
		WorkServerNumber:      1,
		ImageUrl:              imageUrl,
		ImageName:             imageName,
		ImageID:               imageId,
		BootFile:              cloudbrainTask.GetBootFile(option.File, option.OwnerName, option.ProjectName, option.BranchName),
		FileRepository:        sourceRepo,
		FileBranchName:        option.BranchName,
		IsFileNoteBookRequest: true,
		Description:           getDescription(option),
	}, nil, sourceRepo, ctx.User, &traceContext)

	code := 0

	if bizErr != nil {
		traceErr = bizErr.ToError()
		switch bizErr.Code {
		case response.MULTI_TASK.Code:
			code = 2
		default:
			code = 1
		}
		msg := bizErr.DefaultMsg
		if bizErr.TrCode != "" {
			if bizErr.TrParams == nil || len(bizErr.TrParams) == 0 {
				msg = ctx.Tr(bizErr.TrCode)
			} else {
				msg = ctx.Tr(bizErr.TrCode, bizErr.TrParams...)
			}
		}
		ctx.JSON(http.StatusOK, models.BaseMessageApi{Code: code, Message: msg})
		return
	}
	ctx.JSON(http.StatusOK, models.BaseMessageApi{
		Code:    code,
		Message: fmt.Sprint(res.ID),
	})
}

const CharacterLength = 2550

func getDescription(option api.CreateFileNotebookJobOption) string {
	des := option.OwnerName + "/" + option.ProjectName + "/" + option.BranchName + "/" + option.File
	if len(des) <= CharacterLength {
		return des
	}
	return ""
}

func getNpuImageId(option api.CreateFileNotebookJobOption) (*setting.ImageInfoModelArts, error) {
	if option.Type != cloudbrainTask.NPUType {
		return nil, fmt.Errorf("type is not npu.")
	}
	if option.Image == "" {
		return nil, nil
	}
	for _, imageInfo := range setting.StImageInfos.ImageInfo {
		if imageInfo.Value == option.Image {
			return imageInfo, nil
		}
	}
	return nil, fmt.Errorf("invalid image parameter")
}

func FileNoteBookStatus(ctx *context.APIContext, option api.CreateFileNotebookJobOption) {
	cloudbrainTask.FileNotebookStatus(ctx.Context, option)
}

func GetFileNoteBookInfo(ctx *context.APIContext) {
	//image description  spec description  waiting count

	specs, err := models.GetResourceSpecificationByIds([]int64{setting.FileNoteBook.SpecIdCPU, setting.FileNoteBook.SpecIdGPU, setting.FileNoteBook.SpecIdNPU})
	if err != nil {
		log.Error("Fail to query specifications", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_query_fail")))
		return
	}

	var specCPU, specGpu, specNPU *api.SpecificationShow

	for _, spec := range specs {
		if spec.ID == setting.FileNoteBook.SpecIdCPU {
			specCPU = convert.ToSpecification(spec)
		} else if spec.ID == setting.FileNoteBook.SpecIdGPU {
			specGpu = convert.ToSpecification(spec)
		} else if spec.ID == setting.FileNoteBook.SpecIdNPU {
			specNPU = convert.ToSpecification(spec)
		}
	}

	waitCountNpu := cloudbrain.GetWaitingCloudbrainCount(models.TypeC2Net, models.NPUResource, models.JobTypeDebug)

	waitCountGPU := cloudbrain.GetWaitingCloudbrainCount(models.TypeC2Net, models.GPUResource, models.JobTypeDebug)

	var a *models.PointAccount
	if ctx.User != nil {
		a, _ = account.GetAccount(ctx.User.ID)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":                0,
		"projectName":         setting.FileNoteBook.ProjectName,
		"specCpu":             specCPU,
		"specGpu":             specGpu,
		"specNpu":             specNPU,
		"waitCountGpu":        waitCountGPU,
		"waitCountNpu":        waitCountNpu,
		"imageCpuDescription": setting.FileNoteBook.ImageCPUDescription,
		"imageGpuDescription": setting.FileNoteBook.ImageGPUDescription,
		"cloudBrainPaySwitch": setting.CloudBrainPaySwitch,
		"PointAccount":        a,
	})

}

func GrampusNoteBookDebug(ctx *context.APIContext) {

	err := cloudbrainTask.GrampusNotebookDebug(ctx.Context)
	if err != nil {
		ctx.NotFound(err.Error())
	}

}

func GrampusStopJob(ctx *context.APIContext) {
	cloudbrainTask.GrampusStopJob(ctx.Context)

}

func GrampusNotebookDel(ctx *context.APIContext) {
	if isHandled, err := task.HandleNewAITaskDelete(ctx.Cloudbrain.ID); isHandled {
		if err != nil {
			log.Error("DeleteJob(%s) failed:%v", ctx.Cloudbrain.JobName, err, ctx.Data["msgID"])
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
		return
	}
	err := cloudbrainTask.DeleteGrampusJob(ctx.Context)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessageApi)

}

// cloudbrain get job task by jobid
func GetCloudbrainTask(ctx *context.APIContext) {
	// swagger:operation GET /repos/{owner}/{repo}/cloudbrain/{jobid} cloudbrain jobTask
	// ---
	// summary: Get a single task
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: jobid
	//   in: path
	//   description: id of cloudbrain jobid
	//   type: string
	//   required: true
	// responses:
	//   "200":

	//     "$ref": "#/responses/Label"

	var (
		err error
	)

	ID := ctx.Params(":id")

	job, err := cloudbrain.GetCloudBrainByIdOrJobId(ID, "id")

	if err != nil {
		ctx.NotFound(err)
		return
	}

	if job.IsNewAITask() {
		jobAfter, _ := task.UpdateCloudbrain(job)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ID":             ID,
			"JobName":        jobAfter.JobName,
			"JobStatus":      jobAfter.Status,
			"DetailedStatus": jobAfter.DetailedStatus,
			"SubState":       "",
			"CreatedTime":    jobAfter.CreatedUnix.Format("2006-01-02 15:04:05"),
			"CompletedTime":  jobAfter.UpdatedUnix.Format("2006-01-02 15:04:05"),
			"JobDuration":    jobAfter.TrainJobDuration,
		})
		return
	}

	if job.JobType == string(models.JobTypeModelSafety) {
		routerRepo.GetAiSafetyTaskByJob(job)
		job, err = models.GetCloudbrainByID(ID)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ID":             ID,
			"JobName":        job.JobName,
			"JobStatus":      job.Status,
			"DetailedStatus": job.DetailedStatus,
			"SubState":       "",
			"CreatedTime":    job.CreatedUnix.Format("2006-01-02 15:04:05"),
			"CompletedTime":  job.UpdatedUnix.Format("2006-01-02 15:04:05"),
			"JobDuration":    job.TrainJobDuration,
		})
	} else {
		jobAfter, err := cloudbrainTask.SyncCloudBrainOneStatus(job)

		if err != nil {
			ctx.NotFound(err)
			log.Error("Sync cloud brain one status failed:", err)
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ID":             ID,
			"JobName":        jobAfter.JobName,
			"JobStatus":      jobAfter.Status,
			"DetailedStatus": jobAfter.DetailedStatus,
			"SubState":       "",
			"CreatedTime":    jobAfter.CreatedUnix.Format("2006-01-02 15:04:05"),
			"CompletedTime":  jobAfter.UpdatedUnix.Format("2006-01-02 15:04:05"),
			"JobDuration":    jobAfter.TrainJobDuration,
		})
	}
}

func GetCloudBrainInferenceJob(ctx *context.APIContext) {

	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		ctx.NotFound(err)
		return
	}
	jobResult, err := cloudbrain.GetJob(job.JobID)
	if err != nil {
		ctx.NotFound(err)
		log.Error("GetJob failed:", err)
		return
	}
	result, err := models.ConvertToJobResultPayload(jobResult.Payload)
	if err != nil {
		ctx.NotFound(err)
		log.Error("ConvertToJobResultPayload failed:", err)
		return
	}
	oldStatus := job.Status
	job.Status = result.JobStatus.State
	if result.JobStatus.State != string(models.JobWaiting) && result.JobStatus.State != string(models.JobFailed) {
		taskRoles := result.TaskRoles
		taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))

		job.ContainerIp = taskRes.TaskStatuses[0].ContainerIP
		job.ContainerID = taskRes.TaskStatuses[0].ContainerID
		job.Status = taskRes.TaskStatuses[0].State
	}

	if result.JobStatus.State != string(models.JobWaiting) {
		models.ParseAndSetDurationFromCloudBrainOne(result, job)
		if oldStatus != job.Status {
			notification.NotifyChangeCloudbrainStatus(job, oldStatus)
		}
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":       jobID,
		"JobStatus":   job.Status,
		"JobDuration": job.TrainJobDuration,
		"StartTime":   job.StartTime,
	})

}

func DelCloudBrainJob(ctx *context.APIContext) {
	jobID := ctx.Params(":jobid")

	errStr := cloudbrain.DelCloudBrainJob(jobID)

	if errStr != "" {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"Message":     ctx.Tr(errStr),
			"VersionName": "1",
			"Code":        1,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"Message":     "",
			"VersionName": "1",
			"Code":        0,
		})
	}

}

func InferencJobResultList(ctx *context.APIContext) {
	jobID := ctx.Params(":jobid")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")

	task, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("get cloud brain err:", err)
		ctx.ServerError("get cloud brain information failed:", err)

	}

	//get dirs
	dirs, err := routerRepo.GetResultDirs(task.JobName, parentDir)
	if err != nil {
		log.Error("GetModelDirs failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("GetModelDirs failed:", err)
		return
	}

	var fileInfos []storage.FileInfo
	err = json.Unmarshal([]byte(dirs), &fileInfos)
	if err != nil {
		log.Error("json.Unmarshal failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.ServerError("json.Unmarshal failed:", err)
		return
	}

	for i, fileInfo := range fileInfos {
		temp, _ := time.Parse("2006-01-02 15:04:05", fileInfo.ModTime)
		fileInfos[i].ModTime = temp.Local().Format("2006-01-02 15:04:05")
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime > fileInfos[j].ModTime
	})

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":            jobID,
		"StatusOK":         0,
		"Path":             dirArray,
		"Dirs":             fileInfos,
		"task":             task,
		"PageIsCloudBrain": true,
	})

}

func ModelSafetyGetLog(ctx *context.APIContext) {
	ID := ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobName failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	if job.JobType == string(models.JobTypeModelSafety) {
		if job.Type == models.TypeCloudBrainTwo {
			//TrainJobForModelConvertGetLog(ctx)
			var baseLine = ctx.Query("base_line")
			var order = ctx.Query("order")
			var lines = ctx.Query("lines")
			lines_int, err := strconv.Atoi(lines)
			if err != nil {
				log.Error("change lines(%d) string to int failed", lines_int)
			}

			if order != modelarts.OrderDesc && order != modelarts.OrderAsc {
				log.Error("order(%s) check failed", order)
				ctx.JSON(http.StatusBadRequest, map[string]interface{}{
					"err_msg": "order check failed",
				})
				return
			}
			resultLogFile, err := modelarts.GetTrainJobLogFileNames(job.JobID, strconv.FormatInt(job.VersionID, 10))
			if err != nil {
				log.Error("GetTrainJobLogFileNames(%s) failed:%v", job.JobID, err.Error())
			}
			result, err := modelarts.GetTrainJobLog(job.JobID, strconv.FormatInt(job.VersionID, 10), baseLine, resultLogFile.LogFileList[0], order, lines_int)
			if err != nil {
				log.Error("GetTrainJobLog(%s) failed:%v", job.JobID, err.Error())
			}
			if err != nil {
				log.Error("trainJobGetLog(%s) failed:%v", job.JobID, err.Error())
				// ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobShow, nil)
				ctx.JSON(http.StatusOK, map[string]interface{}{
					"JobID":          job.JobID,
					"LogFileName":    "",
					"StartLine":      "0",
					"EndLine":        "0",
					"Content":        "",
					"Lines":          0,
					"CanLogDownload": false,
				})
				return
			}
			prefix := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, job.JobName, modelarts.LogPath, job.VersionName), "/") + "/job"
			_, err = storage.GetObsLogFileName(prefix)
			canLogDownload := isCanDownloadLog(ctx, job)
			if err != nil {
				canLogDownload = false
			}
			ctx.Data["log_file_name"] = resultLogFile.LogFileList[0]
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"JobID":          job.JobID,
				"LogFileName":    resultLogFile.LogFileList[0],
				"StartLine":      result.StartLine,
				"EndLine":        result.EndLine,
				"Content":        result.Content,
				"Lines":          result.Lines,
				"CanLogDownload": canLogDownload,
				"StartTime":      job.StartTime,
			})
		}
	}
	//result := ""
	//ctx.JSON(http.StatusOK, result)
}

func isCanDownloadLog(ctx *context.APIContext, job *models.Cloudbrain) bool {
	if !ctx.IsSigned {
		return false
	}
	return ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
}

func ModelSafetyDownloadLogFile(ctx *context.Context) {
	ID := ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobName failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	if job.JobType == string(models.JobTypeModelSafety) {
		if job.Type == models.TypeCloudBrainOne {
			CloudbrainDownloadLogFile(ctx)
		} else if job.Type == models.TypeCloudBrainTwo {
			ctx.SetParams("jobid", job.JobID)
			ctx.Req.Form.Set("version_name", job.VersionName)
			routerRepo.TrainJobDownloadLogFile(ctx)
		}
	}
}

func CloudbrainDownloadLogFile(ctx *context.Context) {
	ID := ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobName failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	if job.JobType == string(models.JobTypeModelSafety) {
		if job.Type == models.TypeCloudBrainTwo {
			ModelSafetyDownloadLogFile(ctx)
			return
		}
	}

	existStr := ""
	if job.JobType == string(models.JobTypeTrain) || job.JobType == string(models.JobTypeInference) {
		if job.Type == models.TypeCloudBrainOne {
			result, err := cloudbrain.GetJob(job.JobID)
			if err == nil && result != nil {
				jobRes, _ := models.ConvertToJobResultPayload(result.Payload)
				taskRoles := jobRes.TaskRoles
				taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
				existStr = taskRes.TaskStatuses[0].ExitDiagnostics
			}
		}
	}

	logDir := "/model"
	if job.JobType == string(models.JobTypeInference) || job.JobType == string(models.JobTypeModelSafety) {
		logDir = cloudbrain.ResultPath
	}
	files, err := storage.GetOneLevelAllObjectUnderDirMinio(setting.Attachment.Minio.Bucket, setting.CBCodePathPrefix+job.JobName+logDir, "")
	if err != nil {
		log.Error("query cloudbrain model failed: %v", err)
		return
	}
	fileName := ""
	for _, file := range files {
		if strings.HasSuffix(file.FileName, "log.txt") {
			fileName = file.FileName
			break
		}
	}
	if fileName != "" {
		prefix := "/" + setting.CBCodePathPrefix + job.JobName + logDir
		filePath := setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + prefix + "/" + fileName
		// Read the file contents into a byte slice
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			ctx.ServerError("ReadFile", err)
			return
		}

		// Set the appropriate response headers
		ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
		ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+fileName)

		// Write the file contents to the response
		if _, err := ctx.Resp.Write(data); err != nil {
			ctx.ServerError("Write", err)
			return
		}
		if _, err := ctx.Resp.Write([]byte(existStr)); err != nil {
			log.Error("Write failed: %v", err.Error(), ctx.Data["msgID"])
			return
		}
	} else {
		log.Info("fileName is null.")
	}
}

func CloudbrainGetLog(ctx *context.APIContext) {
	ID := ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobName failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	if job.JobType == string(models.JobTypeModelSafety) {
		if job.Type == models.TypeCloudBrainOne {
			result, err := cloudbrain.GetJob(job.JobID)
			existStr := ""
			if err == nil && result != nil {
				jobRes, _ := models.ConvertToJobResultPayload(result.Payload)
				taskRoles := jobRes.TaskRoles
				taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
				existStr = taskRes.TaskStatuses[0].ExitDiagnostics
			}
			ctx.Data["existStr"] = existStr
		} else {
			ModelSafetyGetLog(ctx)
			return
		}
	}

	if job.JobType == string(models.JobTypeTrain) || job.JobType == string(models.JobTypeInference) {
		if job.Type == models.TypeCloudBrainOne {
			result, err := cloudbrain.GetJob(job.JobID)
			existStr := ""
			if err == nil && result != nil {
				jobRes, _ := models.ConvertToJobResultPayload(result.Payload)
				taskRoles := jobRes.TaskRoles
				taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
				existStr = taskRes.TaskStatuses[0].ExitDiagnostics
			}
			ctx.Data["existStr"] = existStr
		}
	}

	lines := ctx.QueryInt("lines")
	baseLine := ctx.Query("base_line")
	order := ctx.Query("order")
	var result map[string]interface{}
	resultPath := "/model"
	if job.JobType == string(models.JobTypeInference) || job.JobType == string(models.JobTypeModelSafety) {
		resultPath = "/result"
	}
	if baseLine == "" && order == "desc" {
		result = getLastLogFromModelDir(job.JobName, lines, resultPath)
	} else {
		startLine := ctx.QueryInt("base_line")
		endLine := startLine + lines
		if order == "asc" {
			if baseLine == "" {
				startLine = 0
				endLine = lines
			} else {
				endLine = startLine
				startLine = endLine - lines
				if startLine < 0 {
					startLine = 0
				}
			}
		} else {
			if startLine > 0 {
				startLine += 1
				endLine += 1
			}
		}

		result = getLogFromModelDir(job.JobName, startLine, endLine, resultPath)
		if result == nil {
			log.Error("GetJobLog failed: %v", err, ctx.Data["MsgID"])
			//ctx.ServerError(err.Error(), err)
			return
		}
	}
	content := ""
	if result["Content"] != nil {
		content = result["Content"].(string)
	}

	if (job.JobType == string(models.JobTypeTrain) || job.JobType == string(models.JobTypeInference)) && job.Type == models.TypeCloudBrainOne && job.Status == string(models.JobFailed) {
		if ctx.Data["existStr"] != nil {
			if baseLine == "" && order == "desc" && result["Lines"].(int) == 0 {
				result["Lines"] = 1
				result["EndLine"] = 1
				content = content + ctx.Data["existStr"].(string)
			}

			if result["Lines"].(int) == 0 && result["StartLine"] == result["EndLine"] && result["StartLine"].(int) != 0 {
				content = content + ctx.Data["existStr"].(string)
				result["Lines"] = 1
				result["StartLine"] = result["StartLine"].(int) - 1
			}
			if result["Lines"].(int) == 1 && result["StartLine"] == result["EndLine"] {
				result["Lines"] = 0
				result["StartLine"] = result["StartLine"].(int) + 1
			}
		}
	} else {
		if ctx.Data["existStr"] != nil && result["Lines"].(int) < 50 {
			content = content + ctx.Data["existStr"].(string)
		}
	}

	logFileName := result["FileName"]

	//Logs can only be downloaded if the file exists
	//and the current user is an administrator or the creator of the task
	canLogDownload := logFileName != nil && logFileName != "" && job.IsUserHasRight(ctx.User)

	re := map[string]interface{}{
		"JobID":          ID,
		"LogFileName":    logFileName,
		"StartLine":      result["StartLine"],
		"EndLine":        result["EndLine"],
		"Content":        content,
		"Lines":          result["Lines"],
		"CanLogDownload": canLogDownload,
		"StartTime":      job.StartTime,
	}
	//result := CloudbrainGetLogByJobId(job.JobID, job.JobName)
	ctx.JSON(http.StatusOK, re)
}

func getAllLineFromFile(path string) int {
	count := 0
	reader, err := os.Open(path)
	defer reader.Close()
	if err == nil {
		r := bufio.NewReader(reader)
		for {
			_, error := r.ReadString('\n')
			if error == io.EOF {
				log.Info("read file completed.")
				break
			}
			if error != nil {
				log.Info("read file error." + error.Error())
				break
			}
			count = count + 1
		}
	} else {
		log.Info("error:" + err.Error())
	}
	return count
}

func getLastLogFromModelDir(jobName string, lines int, resultPath string) map[string]interface{} {
	prefix := setting.CBCodePathPrefix + jobName + resultPath
	files, err := storage.GetOneLevelAllObjectUnderDirMinio(setting.Attachment.Minio.Bucket, prefix, "")
	if err != nil {
		log.Error("query cloudbrain model failed: %v", err)
		return nil
	}

	re := ""
	fileName := ""
	count := 0
	allLines := 0
	startLine := 0
	for _, file := range files {
		if strings.HasSuffix(file.FileName, "log.txt") {
			fileName = file.FileName
			path := storage.GetMinioPath(jobName+resultPath+"/", file.FileName)
			allLines = getAllLineFromFile(path)
			startLine = allLines - lines
			if startLine < 0 {
				startLine = 0
			}
			count = allLines - startLine
			log.Info("path=" + path)
			reader, err := os.Open(path)
			defer reader.Close()
			if err == nil {
				r := bufio.NewReader(reader)
				for i := 0; i < allLines; i++ {
					line, error := r.ReadString('\n')
					if error == io.EOF {
						log.Info("read file completed.")
						break
					}
					if error != nil {
						log.Info("read file error." + error.Error())
						break
					}
					if error == nil {
						if i >= startLine {
							re = re + line
						}
					}
				}
			} else {
				log.Info("error:" + err.Error())
			}
			break
		}
	}

	return map[string]interface{}{
		"JobName":   jobName,
		"Content":   re,
		"FileName":  fileName,
		"Lines":     count,
		"EndLine":   allLines,
		"StartLine": startLine,
	}
}

func getLogFromModelDir(jobName string, startLine int, endLine int, resultPath string) map[string]interface{} {
	prefix := setting.CBCodePathPrefix + jobName + resultPath
	files, err := storage.GetOneLevelAllObjectUnderDirMinio(setting.Attachment.Minio.Bucket, prefix, "")
	if err != nil {
		log.Error("query cloudbrain model failed: %v", err)
		return nil
	}
	if startLine == endLine {
		return map[string]interface{}{
			"JobName":   jobName,
			"Content":   "",
			"FileName":  "",
			"Lines":     0,
			"EndLine":   startLine,
			"StartLine": startLine,
		}
	}
	re := ""
	fileName := ""
	count := 0
	fileEndLine := endLine
	for _, file := range files {
		if strings.HasSuffix(file.FileName, "log.txt") {
			fileName = file.FileName
			path := storage.GetMinioPath(jobName+resultPath+"/", file.FileName)
			log.Info("path=" + path)
			reader, err := os.Open(path)
			defer reader.Close()
			if err == nil {
				r := bufio.NewReader(reader)
				for i := 0; i < endLine; i++ {
					line, error := r.ReadString('\n')
					if error == io.EOF {
						if i >= startLine {
							re = re + line
							count++
						}
						fileEndLine = i + 1
						log.Info("read file completed.")
						break
					}
					if error != nil {
						log.Info("read file error." + error.Error())
						break
					}
					if error == nil {
						if i >= startLine {
							fileEndLine = i + 1
							re = re + line
							count++
						}
					}
				}
			} else {
				log.Info("error:" + err.Error())
			}
			break
		}
	}

	return map[string]interface{}{
		"JobName":   jobName,
		"Content":   re,
		"FileName":  fileName,
		"Lines":     count,
		"EndLine":   fileEndLine,
		"StartLine": startLine,
	}
}

func CloudBrainModelList(ctx *context.APIContext) {
	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")

	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		return
	}

	//get dirs
	dirs, err := routerRepo.GetModelDirs(task.JobName, parentDir)
	if err != nil {
		log.Error("GetModelDirs failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":            jobID,
			"VersionName":      versionName,
			"StatusOK":         0,
			"Path":             dirArray,
			"Dirs":             "",
			"task":             task,
			"CanDownload":      false,
			"PageIsCloudBrain": true,
		})
		return
	}

	var fileInfos []storage.FileInfo
	err = json.Unmarshal([]byte(dirs), &fileInfos)
	if err != nil {
		log.Error("json.Unmarshal failed:%v", err.Error(), ctx.Data["msgID"])
		//ctx.ServerError("json.Unmarshal failed:", err)
		return
	}

	for i, fileInfo := range fileInfos {
		temp, _ := time.Parse("2006-01-02 15:04:05", fileInfo.ModTime)
		fileInfos[i].ModTime = temp.Local().Format("2006-01-02 15:04:05")
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime > fileInfos[j].ModTime
	})

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":            jobID,
		"VersionName":      versionName,
		"StatusOK":         0,
		"Path":             dirArray,
		"Dirs":             fileInfos,
		"task":             task,
		"CanDownload":      cloudbrain.CanDeleteJob(ctx.Context, task),
		"PageIsCloudBrain": true,
	})
}

type JobInfo struct {
	JobName    string `json:"job_name"`
	AiCenterId int    `json:"ai_center_id"`
}

func GetNewestJobs(ctx *context.APIContext) {
	idsC2Net, err := models.GetNewestJobsByAiCenter()
	if err != nil {
		log.Error("GetNewestJobsByAiCenter(%s) failed:%v", err.Error())
		return
	}

	idsCloudbrain, err := models.GetNewestJobsByType()
	if err != nil {
		log.Error("GetNewestJobsByType(%s) failed:%v", err.Error())
		return
	}

	ids := make([]int64, len(idsC2Net), cap(idsC2Net)*2)
	copy(ids, idsC2Net)
	for _, id := range idsCloudbrain {
		ids = append(ids, id)
	}

	jobs, err := models.GetCloudbrainByIDs(ids)
	if err != nil {
		log.Error("GetCloudbrainByIDs(%s) failed:%v", err.Error())
		return
	}

	jobInfos := make([]JobInfo, 0)
	for _, job := range jobs {
		var id int
		var content string
		switch job.Type {
		case models.TypeCloudBrainOne:
			id, content = getAICenterID("cloudbrain_one")
			if content == "" {
				log.Error("job(%s) has no match config info", job.DisplayJobName)
				continue
			}
		case models.TypeCloudBrainTwo:
			id, content = getAICenterID("cloudbrain_two")
			if content == "" {
				log.Error("job(%s) has no match config info", job.DisplayJobName)
				continue
			}
		case models.TypeC2Net:
			centerInfo := strings.Split(job.AiCenter, "+")
			if len(centerInfo) != 2 {
				log.Error("job(%s):ai_center(%s) is wrong", job.DisplayJobName, job.AiCenter)
				continue
			}
			id, content = getAICenterID(centerInfo[0])
			if content == "" {
				log.Error("job(%s) has no match config info", job.DisplayJobName)
				continue
			}
		default:
			log.Error("no match info")
			continue
		}

		jobInfos = append(jobInfos, JobInfo{
			JobName:    job.DisplayJobName,
			AiCenterId: id,
		})
	}

	ctx.JSON(http.StatusOK, jobInfos)
}

func GetAICenterInfo(ctx *context.APIContext) {
	if setting.C2NetInfos == nil {
		log.Error("C2NET_SEQUENCE is incorrect")
		return
	}

	ctx.JSON(http.StatusOK, setting.C2NetInfos.C2NetSqInfo)
}

type AimJobName struct {
	JobName []string `json:"job_name"`
}

func GetAimUrl(ctx *context.APIContext, aimJobName AimJobName) {

	job_names := aimJobName.JobName

	if len(job_names) == 0 {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx))
		return
	}

	if len(job_names) == 1 {
		job_name := job_names[0]
		url, err := aim.GetDetailUrl(job_name)
		if err != nil {
			log.Error("GetDetailUrl failed:%v", err)
			ctx.JSON(http.StatusOK, response.OuterTrBizError(response.SYSTEM_ERROR, ctx))
			return
		} else {
			m := map[string]interface{}{"url": url}
			ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
		}

	} else {

		url, err := aim.GetComparelUrl(job_names)
		if err != nil {
			log.Error("GetCompareUrl failed:%v", err)
			ctx.JSON(http.StatusOK, response.OuterTrBizError(response.SYSTEM_ERROR, ctx))
			return
		} else {
			m := map[string]interface{}{"url": url}
			ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
		}

	}

}

func HasMLopsRight(ctx *context.APIContext) {
	if !ctx.IsSigned {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"mlops_right": false,
		})
		return
	}
	if !setting.MLOPS || !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"mlops_right": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"mlops_right": true,
	})
}

func HasAIMRight(ctx *context.APIContext) {
	if !ctx.IsSigned {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"aim_right": false,
		})
		return
	}
	if !setting.AimConfig.Enabled || !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"aim_right": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"aim_right": true,
	})
}

func getAICenterID(name string) (int, string) {
	for _, info := range setting.C2NetInfos.C2NetSqInfo {
		if name == info.Name {
			return info.ID, info.Content
		}
	}

	return 0, ""
}
