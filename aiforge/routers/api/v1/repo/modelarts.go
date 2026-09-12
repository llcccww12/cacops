// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/services/ai_task_service/schedule"
	"code.gitea.io/gitea/services/ai_task_service/task"

	"code.gitea.io/gitea/routers/response"

	"code.gitea.io/gitea/modules/cloudbrain"

	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
	routerRepo "code.gitea.io/gitea/routers/repo"
	ai_task "code.gitea.io/gitea/services/ai_task_service/task"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
)

func GetModelArtsNotebook2(ctx *context.APIContext) {
	var (
		err error
	)

	ID := ctx.Params(":id")

	job, err := cloudbrain.GetCloudBrainByIdOrJobId(ID, "id")

	if err != nil {
		ctx.NotFound(err)
		return
	}
	if !job.Cleared {
		if job.IsNewAITask() {
			job, _ = task.UpdateCloudbrain(job)
		} else {
			err = modelarts.HandleNotebookInfo(job)
			if err != nil {
				ctx.NotFound(err)
				return
			}
		}

	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"ID":             ID,
		"JobName":        job.JobName,
		"JobStatus":      job.Status,
		"JobDuration":    job.TrainJobDuration,
		"StartTime":      job.StartTime,
		"DetailedStatus": job.DetailedStatus,
	})

}

func GetModelArtsTrainJob(ctx *context.APIContext) {
	var (
		err error
	)

	jobID := ctx.Params(":jobid")
	repoID := ctx.Repo.Repository.ID
	job, err := models.GetRepoCloudBrainByJobID(repoID, jobID)
	if err != nil {
		ctx.NotFound(err)
		return
	}
	if !job.Cleared {
		result, err := modelarts.GetTrainJob(jobID, strconv.FormatInt(job.VersionID, 10))
		if err != nil {
			ctx.NotFound(err)
			return
		}
		oldStatus := job.Status
		job.Status = modelarts.TransTrainJobStatus(result.IntStatus)
		job.Duration = result.Duration
		job.TrainJobDuration = result.TrainJobDuration
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
		"JobDuration": job.Duration,
	})

}

func GetModelArtsTrainJobVersion(ctx *context.APIContext) {
	var (
		err          error
		aiCenterName string
	)

	jobID := ctx.Params(":jobid")
	versionName := ctx.Query("version_name")
	var job *models.Cloudbrain
	id := ctx.QueryInt64("id")
	if id > 0 {
		job, err = models.GetCloudbrainByCloudbrainID(id)
	} else {
		job, err = models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	}
	if err != nil {
		ctx.NotFound(err)
		return
	}
	if job.IsNewAITask() {
		var bizErr *response.BizError
		job, bizErr = ai_task.UpdateCloudbrain(job)
		if bizErr != nil {
			log.Error("UpdateCloudbrain err.job.DisplayJobName = %s  err=%v", job.DisplayJobName, err)
			ctx.NotFound(err)
			return
		}
		aiCenterName = cloudbrainService.GetAiCenterShow(job.AiCenter, ctx.Context)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":          jobID,
			"JobStatus":      job.Status,
			"DetailedStatus": job.DetailedStatus,
			"JobDuration":    job.TrainJobDuration,
			"AiCenter":       aiCenterName,
			"StartTime":      job.StartTime,
		})
		return
	}
	if job.Type == models.TypeCloudBrainOne {
		aiCenterName = routerRepo.GetAiCenterNameByCode(models.AICenterOfCloudBrainOne, ctx.Language())
		job, err = cloudbrainTask.SyncCloudBrainOneStatus(job)
		if err != nil {
			ctx.NotFound(err)
			return
		}
	} else if job.Type == models.TypeCloudBrainTwo {
		aiCenterName = routerRepo.GetAiCenterNameByCode(models.AICenterOfCloudBrainTwo, ctx.Language())
		if !job.Cleared {
			err := modelarts.HandleTrainJobInfo(job)
			if err != nil {
				ctx.NotFound(err)
				return
			}
		}
	} else if job.Type == models.TypeC2Net {
		result, err := grampus.GetJob(jobID)
		if err != nil {
			log.Error("GetJob(%s) failed:%v", job.JobName, err)
			ctx.NotFound(err)
			return
		}

		if job.StartTime == 0 && result.JobInfo.StartedAt > 0 {
			job.StartTime = timeutil.TimeStamp(result.JobInfo.StartedAt)
		}
		oldStatus := job.Status
		job.Status = grampus.TransTrainJobStatus(result.JobInfo.Status)
		job.Duration = result.JobInfo.RunSec
		job.TrainJobDuration = models.ConvertDurationToStr(job.Duration)

		if job.EndTime == 0 && models.IsTrainJobTerminal(job.Status) && job.StartTime > 0 {
			job.EndTime = job.StartTime.Add(job.Duration)
		}
		job.CorrectCreateUnix()

		if len(job.AiCenter) == 0 {
			if len(result.JobInfo.Tasks) > 0 {
				if len(result.JobInfo.Tasks[0].CenterID) > 0 && len(result.JobInfo.Tasks[0].CenterName) > 0 {
					job.AiCenter = result.JobInfo.Tasks[0].CenterID[0] + "+" + result.JobInfo.Tasks[0].CenterName[0]
					aiCenterName = cloudbrainService.GetAiCenterShow(job.AiCenter, ctx.Context)
				}
			}
		} else {
			aiCenterName = cloudbrainService.GetAiCenterShow(job.AiCenter, ctx.Context)
		}
		if oldStatus != job.Status {
			notification.NotifyChangeCloudbrainStatus(job, oldStatus)
		}
		err = models.UpdateTrainJobVersion(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":          jobID,
		"JobStatus":      job.Status,
		"JobDuration":    job.TrainJobDuration,
		"AiCenter":       aiCenterName,
		"StartTime":      job.StartTime,
		"DetailedStatus": job.DetailedStatus,
	})

}

func GetModelScheduleStatus(ctx *context.APIContext) {
	jobID := ctx.Params(":jobid")
	status, err := schedule.GetModelScheduleStatus(jobID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	m := map[string]interface{}{"status": status}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func RetryModelSchedule(ctx *context.APIContext) {
	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("RetryModelMigrate GetCloudbrainByJobID err.jobId=%s err=%v", jobID, err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(errors.New("jobId not correct")))
	}
	err = schedule.RetryModelMigrate(job)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func TrainJobForModelConvertGetLog(ctx *context.APIContext) {
	var (
		err error
	)

	var jobID = ctx.Params(":id")
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

	resultLogFile, result, err := trainJobForModelConvertGetLogContent(jobID, baseLine, order, lines_int)
	if err != nil {
		log.Error("trainJobGetLog(%s) failed:%v", jobID, err.Error())
		// ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobShow, nil)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":       jobID,
			"LogFileName": "",
			"StartLine":   "0",
			"EndLine":     "0",
			"Content":     "",
			"Lines":       0,
		})
		return
	}

	ctx.Data["log_file_name"] = resultLogFile.LogFileList[0]

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":       jobID,
		"LogFileName": resultLogFile.LogFileList[0],
		"StartLine":   result.StartLine,
		"EndLine":     result.EndLine,
		"Content":     result.Content,
		"Lines":       result.Lines,
	})
}

func trainJobForModelConvertGetLogContent(jobID string, baseLine string, order string, lines int) (*models.GetTrainJobLogFileNamesResult, *models.GetTrainJobLogResult, error) {
	task, err := models.QueryModelConvertById(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", jobID, err.Error())
		return nil, nil, err
	}
	resultLogFile, err := modelarts.GetTrainJobLogFileNames(task.CloudBrainTaskId, task.ModelArtsVersionId)
	if err != nil {
		log.Error("GetTrainJobLogFileNames(%s) failed:%v", task.CloudBrainTaskId, err.Error())
		return nil, nil, err
	}
	result, err := modelarts.GetTrainJobLog(task.CloudBrainTaskId, task.ModelArtsVersionId, baseLine, resultLogFile.LogFileList[0], order, lines)
	if err != nil {
		log.Error("GetTrainJobLog(%s) failed:%v", task.CloudBrainTaskId, err.Error())
		return nil, nil, err
	}

	return resultLogFile, result, err
}

func TrainJobGetLog(ctx *context.APIContext) {
	var (
		err error
	)

	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
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

	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", jobID, err.Error())
		return
	}
	resultLogFile, result, err := trainJobGetLogContent(jobID, task.VersionID, baseLine, order, lines_int)
	if err != nil {
		log.Error("trainJobGetLog(%s) failed:%v", jobID, err.Error())
		// ctx.RenderWithErr(err.Error(), tplModelArtsTrainJobShow, nil)
		return
	}

	ctx.Data["log_file_name"] = resultLogFile.LogFileList[0]

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":          jobID,
		"LogFileName":    resultLogFile.LogFileList[0],
		"StartLine":      result.StartLine,
		"EndLine":        result.EndLine,
		"Content":        result.Content,
		"Lines":          result.Lines,
		"CanLogDownload": canLogDownload(ctx.User, task),
		"StartTime":      task.StartTime,
	})
}

func canLogDownload(user *models.User, task *models.Cloudbrain) bool {
	if task == nil || !task.IsUserHasRight(user) {
		return false
	}
	prefix := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, modelarts.LogPath, task.VersionName), "/") + "/job"
	_, err := storage.GetObsLogFileName(prefix)
	if err != nil {
		return false
	}
	return true
}

func trainJobGetLogContent(jobID string, versionID int64, baseLine string, order string, lines int) (*models.GetTrainJobLogFileNamesResult, *models.GetTrainJobLogResult, error) {

	resultLogFile, err := modelarts.GetTrainJobLogFileNames(jobID, strconv.FormatInt(versionID, 10))
	if err != nil {
		log.Error("GetTrainJobLogFileNames(%s) failed:%v", jobID, err.Error())
		return nil, nil, err
	}

	result, err := modelarts.GetTrainJobLog(jobID, strconv.FormatInt(versionID, 10), baseLine, resultLogFile.LogFileList[0], order, lines)
	if err != nil {
		log.Error("GetTrainJobLog(%s) failed:%v", jobID, err.Error())
		return nil, nil, err
	}

	return resultLogFile, result, err
}

func DelTrainJobVersion(ctx *context.APIContext) {
	var (
		err  error
		task *models.Cloudbrain
	)

	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
	var id = ctx.QueryInt64("id")
	if id > 0 {
		task, err = models.GetCloudbrainByCloudbrainID(id)
	} else {
		task, err = models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	}
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		ctx.NotFound(err)
		return
	}
	if task.IsNewAITask() {
		bizErr := ai_task.DelCloudbrain(task)
		if bizErr != nil {
			log.Error("DelCloudbrain(%s) failed:%v err=%v", task.JobName, bizErr)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"Message":  ctx.Tr(bizErr.TrCode),
				"StatusOK": 1,
			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":       task.JobID,
			"VersionName": task.VersionName,
			"StatusOK":    0,
		})
		return
	}

	if !task.IsTerminal() {
		log.Error("the job(%s) version has not been stopped", task.JobName)
		ctx.NotFound(err)
		return
	}

	//删除modelarts上的记录
	_, err = modelarts.DelTrainJobVersion(jobID, strconv.FormatInt(task.VersionID, 10))
	if err != nil {
		log.Error("DelTrainJobVersion(%s) failed:%v", task.JobName, err.Error())
		if err.Error() == "1" {
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"Message":  ctx.Tr("deployment.deletion_notice_trainjob"),
				"StatusOK": 1,
			})
		}
		return
	}

	//删除数据库记录
	err = models.DeleteJob(task)
	if err != nil {
		ctx.ServerError("DeleteJob failed", err)
		ctx.NotFound(err)
		return
	}

	//获取删除后的版本数量
	var jobTypes []string
	jobTypes = append(jobTypes, string(models.JobTypeTrain))
	repo := ctx.Repo.Repository
	VersionTaskList, VersionListCount, err := models.CloudbrainsVersionList(&models.CloudbrainsOptions{
		RepoID:   repo.ID,
		Type:     models.TypeCloudBrainTwo,
		JobTypes: jobTypes,
		JobID:    jobID,
	})
	if err != nil {
		ctx.ServerError("get VersionListCount failed", err)
		return
	}
	if VersionListCount > 0 {
		// 判断当前删掉的任务是否是最新版本，若是，将排序后的TotalVersionCount置为删掉的最新版本的TotalVersionCount，若不是，按时间排序后的版本列表的第一个版本设置为最新版本，TotalVersionCount不变
		if task.IsLatestVersion == modelarts.IsLatestVersion {
			err = models.SetVersionCountAndLatestVersion(jobID, VersionTaskList[0].Cloudbrain.VersionName, VersionListCount, modelarts.IsLatestVersion, task.TotalVersionCount)
			if err != nil {
				ctx.ServerError("UpdateJobVersionCount failed", err)
				return
			}
		} else {
			err = models.SetVersionCountAndLatestVersion(jobID, VersionTaskList[0].VersionName, VersionListCount, modelarts.IsLatestVersion, VersionTaskList[0].Cloudbrain.TotalVersionCount)
			if err != nil {
				ctx.ServerError("UpdateJobVersionCount failed", err)
				return
			}
		}
	} else { //已删除该任务下的所有版本
		routerRepo.DeleteJobStorage(task.JobName)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":            jobID,
		"VersionName":      versionName,
		"StatusOK":         0,
		"VersionListCount": VersionListCount,
	})
}

func StopTrainJobVersion(ctx *context.APIContext) {
	var (
		err  error
		task *models.Cloudbrain
	)
	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
	var id = ctx.QueryInt64("id")
	if id > 0 {
		task, err = models.GetCloudbrainByCloudbrainID(id)
	} else {
		task, err = models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	}
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		return
	}
	if task.IsNewAITask() {
		_, bizErr := ai_task.StopCloudbrain(task)
		if bizErr != nil {
			log.Error("StopCloudbrain(%s) failed:%v err=%v", task.JobName, bizErr)
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":       task.JobID,
			"VersionName": task.VersionName,
			"StatusOK":    0,
		})
	}
	_, err = modelarts.StopTrainJob(jobID, strconv.FormatInt(task.VersionID, 10))
	if err != nil {
		log.Error("StopTrainJob(%s) failed:%v", task.JobName, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":       jobID,
		"VersionName": versionName,
		"StatusOK":    0,
	})
}

func ModelList(ctx *context.APIContext) {
	var (
		err error
	)

	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")
	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		return
	}

	status := models.ModelMigrateSuccess

	if task.Type == models.TypeC2Net {
		if !task.IsTerminal() {
			log.Info("GetModelScheduleStatus job is not terminal.jobId=%s", jobID)
			status = models.JobNoTeminal
		} else {
			status, err = schedule.GetModelScheduleStatus(task.JobID)
			if err != nil {
				log.Error("GetModelScheduleStatus(%s) failed:%v", task.JobName, err.Error())
				return
			}
		}
	}

	if status != models.ModelMigrateSuccess {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":            jobID,
			"VersionName":      versionName,
			"StatusOK":         status,
			"Path":             dirArray,
			"Dirs":             []storage.FileInfo{},
			"task":             task,
			"PageIsCloudBrain": true,
		})
		return
	}

	var fileInfos []storage.FileInfo
	if task.ComputeResource == models.NPUResource {
		prefix := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, task.JobName, setting.OutPutPath, versionName), "/")
		if !strings.HasSuffix(prefix, "/") {
			prefix += "/"
		}
		fileInfos, err = storage.GetOneLevelObjectsUnderDir(setting.Bucket, prefix, parentDir)
		if err != nil {
			log.Info("get TrainJobListModel failed:", err)
			ctx.ServerError("GetObsListObject:", err)
			return
		}

	} else if task.ComputeResource == models.GPUResource || task.ComputeResource == models.GCUResource {
		files, err := routerRepo.GetModelDirs(task.JobName, parentDir)
		if err != nil {
			log.Info("GetModelDirs failed:", err)
			ctx.ServerError("GetModelDirs:", err)
			return
		}

		err = json.Unmarshal([]byte(files), &fileInfos)
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
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":            jobID,
		"VersionName":      versionName,
		"StatusOK":         status,
		"Path":             dirArray,
		"Dirs":             fileInfos,
		"task":             task,
		"PageIsCloudBrain": true,
	})
}

func GetModelArtsInferenceJob(ctx *context.APIContext) {
	var (
		err error
	)

	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		ctx.NotFound(err)
		return
	}
	err = modelarts.HandleTrainJobInfo(job)
	if err != nil {
		ctx.NotFound(err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":       jobID,
		"JobStatus":   job.Status,
		"JobDuration": job.TrainJobDuration,
		"StartTime":   job.StartTime,
	})

}

func ResultList(ctx *context.APIContext) {
	var (
		err error
	)

	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")
	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", task.JobName, err.Error())
		return
	}
	models, err := storage.GetObsListObject(task.JobName, "result/", parentDir, versionName)
	if err != nil {
		log.Info("get TrainJobListModel failed:", err)
		ctx.ServerError("GetObsListObject:", err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":            jobID,
		"VersionName":      versionName,
		"StatusOK":         0,
		"Path":             dirArray,
		"Dirs":             models,
		"task":             task,
		"PageIsCloudBrain": true,
	})
}

func TrainJobGetMetricStatistic(ctx *context.APIContext) {
	var (
		err error
	)

	var jobID = ctx.Params(":jobid")
	var versionName = ctx.Query("version_name")

	result, err := trainJobGetMetricStatistic(jobID, versionName)
	if err != nil {
		log.Error("trainJobGetMetricStatistic(%s) failed:%v", jobID, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":       jobID,
		"Interval":    result.Interval,
		"MetricsInfo": result.MetricsInfo,
	})
}

func trainJobGetMetricStatistic(jobID string, versionName string) (*models.GetTrainJobMetricStatisticResult, error) {
	task, err := models.GetCloudbrainByJobIDAndVersionName(jobID, versionName)
	if err != nil {
		log.Error("GetCloudbrainByJobIDAndVersionName(%s) failed:%v", jobID, err.Error())
		return nil, err
	}

	resultLogFile, err := modelarts.GetTrainJobLogFileNames(jobID, strconv.FormatInt(task.VersionID, 10))
	if err != nil {
		log.Error("GetTrainJobLogFileNames(%s) failed:%v", jobID, err.Error())
		return nil, err
	}

	result, err := modelarts.GetTrainJobMetricStatistic(jobID, strconv.FormatInt(task.VersionID, 10), resultLogFile.LogFileList[0])
	if err != nil {
		log.Error("GetTrainJobMetricStatistic(%s) failed:%v", jobID, err.Error())
		return nil, err
	}

	return result, err
}

func DownloadMultiResultFile(ctx *context.APIContext) {
	log.Info("DownloadMultiResultFile by api")
	routerRepo.DownloadMultiResultFile(ctx.Context)
}
