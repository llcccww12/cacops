package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"code.gitea.io/gitea/services/cloudbrain/modelmanage"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"
	routerRepo "code.gitea.io/gitea/routers/repo"
	"code.gitea.io/gitea/services/ai_task_service/schedule"
)

type FileInfo struct {
	FileName string `json:"fileName"`
	ModTime  string `json:"modTime"`
	IsDir    bool   `json:"isDir"`
	Size     int64  `json:"size"`
	ParenDir string `json:"parenDir"`
	UUID     string `json:"uuid"`
}

func CreateNewModel(ctx *context.APIContext) {
	log.Info("CreateNewModel by api.")
	routerRepo.SaveModel(ctx.Context)
}

func ShowModelManageApi(ctx *context.APIContext) {
	log.Info("ShowModelManageApi by api.")
	routerRepo.ShowModelPageInfo(ctx.Context)
}

func DeleteModel(ctx *context.APIContext) {
	log.Info("DeleteModel by api.")
	routerRepo.DeleteModel(ctx.Context)
}

func DownloadModel(ctx *context.APIContext) {
	log.Info("DownloadModel by api.")
	routerRepo.DownloadMultiModelFile(ctx.Context)
}

func DownloadModelSingle(ctx *context.APIContext) {
	log.Info("DownloadModel by api.")
	routerRepo.DownloadSingleModelFile(ctx.Context)
}

func QueryModelById(ctx *context.APIContext) {
	log.Info("QueryModelById by api.")
	model := routerRepo.QueryModelObjById(ctx.Context)
	ctx.JSON(200, model)
}

func QueryModelByName(ctx *context.APIContext) {
	log.Info("QueryModelByName by api.")
	models := routerRepo.QueryModelObjByName(ctx.Context)

	ctx.JSON(200, models)
}

func QueryModelListForPredict(ctx *context.APIContext) {
	log.Info("QueryModelListForPredict by api.")
	ctx.Context.SetParams("isOnlyThisRepo", "true")
	routerRepo.QueryModelListForPredict(ctx.Context)
}

func QueryTrainJobList(ctx *context.APIContext) {
	result, err := routerRepo.QueryTrainJobListApi(ctx.Context)
	if err != nil {
		log.Info("query error." + err.Error())
		ctx.JSON(http.StatusOK, nil)
	} else {
		re := make([]*api.Cloudbrain, 0)
		for _, task := range result {
			conRe := convert.ToCloudBrain(task)
			re = append(re, conRe)
		}
		ctx.JSON(http.StatusOK, re)
	}
}

func QueryTrainModelList(ctx *context.APIContext) {
	result, err := routerRepo.QueryTrainModelFileById(ctx.Context)
	if err != nil {
		log.Info("query error." + err.Error())
	}
	re := convertFileFormat(result)
	ctx.JSON(http.StatusOK, re)
}

func QueryTrainJobVersionList(ctx *context.APIContext) {
	result, err := routerRepo.QueryTrainJobVersionListApi(ctx.Context)
	if err != nil {
		log.Info("query error." + err.Error())
		ctx.JSON(http.StatusOK, nil)
	} else {
		re := make([]*api.Cloudbrain, 0)
		for _, task := range result {
			conRe := convert.ToCloudBrain(task)
			re = append(re, conRe)
		}
		ctx.JSON(http.StatusOK, re)
	}
}

func convertFileFormat(result []storage.FileInfo) []FileInfo {
	re := make([]FileInfo, 0)
	if result != nil {
		for _, file := range result {
			tmpFile := FileInfo{
				FileName: file.FileName,
				ModTime:  file.ModTime,
				IsDir:    file.IsDir,
				Size:     file.Size,
				ParenDir: file.ParenDir,
				UUID:     file.UUID,
			}
			re = append(re, tmpFile)
		}
	}
	return re
}

func QueryModelFileForPredict(ctx *context.APIContext) {
	log.Info("QueryModelFileForPredict by api.")
	id := ctx.Query("id")
	result := routerRepo.QueryModelFileByID(id)
	re := convertFileFormat(result)
	ctx.JSON(http.StatusOK, re)
}

func CreateModelConvert(ctx *context.APIContext) {
	log.Info("CreateModelConvert by api.")
	routerRepo.SaveModelConvert(ctx.Context)
}

func StopModelConvert(ctx *context.APIContext) {
	log.Info("StopModelConvert by api.")
	routerRepo.StopModelConvertApi(ctx.Context)
}

func ShowModelConvertPage(ctx *context.APIContext) {
	log.Info("ShowModelConvertPage by api.")
	modelResult, count, err := routerRepo.GetModelConvertPageData(ctx.Context)
	if err == nil {
		mapInterface := make(map[string]interface{})
		mapInterface["data"] = modelResult
		mapInterface["count"] = count
		ctx.JSON(http.StatusOK, mapInterface)
	} else {
		mapInterface := make(map[string]interface{})
		mapInterface["data"] = nil
		mapInterface["count"] = 0
		ctx.JSON(http.StatusOK, mapInterface)
	}

}

func QueryModelConvertById(ctx *context.APIContext) {
	modelResult, err := routerRepo.GetModelConvertById(ctx.Context)
	if err == nil {
		ctx.JSON(http.StatusOK, modelResult)
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}
func QueryOneLevelModelFile(ctx *context.APIContext) {
	routerRepo.QueryOneLevelModelFile(ctx.Context)
}

func QueryModelConvertByName(ctx *context.APIContext) {
	modelResult, err := routerRepo.GetModelConvertByName(ctx.Context)
	if err == nil {
		ctx.JSON(http.StatusOK, modelResult)
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}

func MultiModelDownload(ctx *context.APIContext) {
	log.Info("MultiModelDownload by api.")
	routerRepo.MultiModelDownload(ctx.Context)
}

func QueryModeConvertResultFile(ctx *context.APIContext) {
	log.Info("QueryModeConvertResultFile by api.")
	result, err := QueryModelConvertResultFileList(ctx, ctx.Query("id"))
	if err == nil {
		re := convertFileFormat(result)
		ctx.JSON(http.StatusOK, re)
	} else {
		ctx.JSON(http.StatusOK, nil)
	}
}

func DownloadModeConvertResultFile(ctx *context.APIContext) {
	log.Info("QueryModeConvertResultFile by api.")
	ctx.Context.SetParams("id", ctx.Query("id"))
	routerRepo.ModelConvertDownloadModel(ctx.Context)
}

func SaveLocalModel(ctx *context.APIContext) {
	log.Info("SaveLocalModel by api.")
	routerRepo.SaveLocalModel(ctx.Context)
}

func DeleteModelFile(ctx *context.APIContext) {
	log.Info("DeleteModelFile by api.")
	routerRepo.DeleteModelFile(ctx.Context)
}

func QueryAllModelFile(ctx *context.APIContext) {
	log.Info("QueryAllModelFile by api.")
	routerRepo.ModelSquareData(ctx.Context)
}

func GetCloudbrainModelConvertTask(ctx *context.APIContext) {
	var (
		err error
	)
	ID := ctx.Params(":id")
	job, err := models.QueryModelConvertById(ID)
	if err != nil {
		ctx.NotFound(err)
		log.Error("GetCloudbrainByID failed:", err)
		return
	}
	if job.IsGpuTrainTask() {
		jobResult, err := grampus.GetJob(job.CloudBrainTaskId)
		if err != nil {
			ctx.NotFound(err)
			log.Error("GetJob failed:", err)
			return
		}
		jobResultJson, _ := json.Marshal(jobResult)
		log.Info("grampus jobResultJson=" + string(jobResultJson))

		if jobResult.JobInfo.Status == models.GrampusStatusPending || jobResult.JobInfo.Status == "" {
			job.Status = models.GrampusStatusWaiting
		} else {
			job.Status = strings.ToUpper(jobResult.JobInfo.Status)
		}
		if jobResult.JobInfo.CompletedAt > 0 {
			job.EndTime = timeutil.TimeStamp(jobResult.JobInfo.CompletedAt)
		}
		if strings.ToUpper(jobResult.JobInfo.Status) != models.GrampusStatusWaiting && jobResult.JobInfo.Status != models.GrampusStatusPending {
			models.ModelConvertSetDuration(job)
			err = models.UpdateModelConvert(job)
			if err != nil {
				log.Error("UpdateJob failed:", err)
			}
		}
		doGrampusModelConvertMigrate(jobResult.JobInfo.Status, job)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ID":            ID,
			"JobName":       jobResult.JobInfo.Name,
			"JobStatus":     job.Status,
			"SubState":      "",
			"CreatedTime":   time.Unix(jobResult.JobInfo.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			"CompletedTime": time.Unix(jobResult.JobInfo.CompletedAt, 0).Format("2006-01-02 15:04:05"),
		})
	} else {

		result, err := modelarts.GetTrainJob(job.CloudBrainTaskId, job.ModelArtsVersionId)
		if err != nil {
			log.Error("get modelart job failed:", err)
			ctx.NotFound(err)
			return
		}

		job.Status = modelarts.TransTrainJobStatus(result.IntStatus)
		job.RunTime = result.Duration / 1000
		job.TrainJobDuration = models.ConvertDurationToStr(job.RunTime)
		err = models.UpdateModelConvert(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"ID":        ID,
			"JobStatus": job.Status,
		})

	}

}

func GrampusTaskGetLogByJobId(jobId string, jobName string) map[string]interface{} {
	var content string
	result, err := grampus.GetTrainJobLog(jobId)
	if err != nil {
		log.Error("GetJobLog failed: %v", err)
		content = ""
	} else {
		content = result
	}
	return map[string]interface{}{
		"JobName": jobName,
		"Content": content,
	}
}

func GrampusForModelConvertGetLog(ctx *context.Context) {
	ID := ctx.Params(":id")
	job, err := models.QueryModelConvertById(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobName failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	result := GrampusTaskGetLogByJobId(job.CloudBrainTaskId, job.Name)
	if result == nil {
		log.Error("GetJobLog failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func CloudBrainModelConvertList(ctx *context.APIContext) {
	ID := ctx.Params(":id")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")
	var versionName = "V0001"

	job, err := models.QueryModelConvertById(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", job.Name, err.Error())
		ctx.ServerError("GetModelDirs failed:", err)
		return
	}
	result, err := QueryModelConvertResultFileList(ctx, ID)
	if err == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":            job.ID,
			"VersionName":      versionName,
			"StatusOK":         0,
			"Path":             dirArray,
			"Dirs":             result,
			"task":             job,
			"PageIsCloudBrain": true,
		})
	} else {
		log.Error("GetCloudbrainByJobID failed:%v", err.Error())
		ctx.ServerError("GetModelDirs failed:", err)
		return
	}
}

func QueryModelConvertResultFileList(ctx *context.APIContext, id string) ([]storage.FileInfo, error) {
	ID := id
	parentDir := ctx.Query("parentDir")
	job, err := models.QueryModelConvertById(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobID(%s) failed:%v", job.Name, err.Error())
		return nil, err
	}
	if job.IsGpuTrainTask() {

		path := setting.CBCodePathPrefix + job.ID + cloudbrain.ModelMountPath + "/"
		log.Info("get model convert result file path=" + path)
		fileInfos, err := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, path)
		if err == nil {
			for i, fileInfo := range fileInfos {
				temp, _ := time.Parse("2006-01-02 15:04:05", fileInfo.ModTime)
				fileInfos[i].ModTime = temp.Local().Format("2006-01-02 15:04:05")
			}

			sort.Slice(fileInfos, func(i, j int) bool {
				return fileInfos[i].ModTime > fileInfos[j].ModTime
			})

			return fileInfos, nil
		} else {
			log.Info("query models path error.")
			return nil, err
		}
	} else {
		var versionName = "V0001"
		models, err := storage.GetObsListObject(job.ID, "output/", parentDir, versionName)
		if err != nil {
			log.Info("get TrainJobListModel failed:", err)
			return nil, err
		}
		return models, nil
	}

}

func doGrampusModelConvertMigrate(status string, job *models.AiModelConvert) error {
	if strings.ToUpper(status) == models.GrampusStatusSucceeded {
		if job.StatusResult != "MIGRATE_SUCCEED" {
			lock := redis_lock.NewDistributeLock(redis_key.RecordHandleLock(job.CloudBrainTaskId))
			success, err := lock.Lock(60 * time.Second)
			if err != nil {
				log.Error("HandleUnfinishedMigrateRecord lock err.ID=%d %v", job.CloudBrainTaskId, err)
				return err
			}
			if !success {
				log.Error("HandleUnfinishedMigrateRecord lock failed.ID=%d ", job.CloudBrainTaskId)
				return nil
			}
			//todo migrate
			grampus.PostModelMigrate(job.CloudBrainTaskId)
			go dealModelConvertModelMigrate(job)
		}
	}

	return nil
}

func dealModelConvertModelMigrate(job *models.AiModelConvert) {
	count := 0
	for {
		if count > 20 {
			break
		}
		log.Info("deal count= " + fmt.Sprint(count))
		if updateModelMigrateStatus(job) {
			count++
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
}

func updateModelMigrateStatus(job *models.AiModelConvert) bool {
	res, err := grampus.ModelMigrateInfo(job.CloudBrainTaskId)
	if err != nil {
		log.Error("ModelMigrateInfo err. r.ID=%d  %v", job.CloudBrainTaskId, err)
		return true
	}
	log.Info("model migrate status=" + fmt.Sprint(res.Status))
	status := models.GrampusMigrateResponse(res.Status).ConvertToModelMigrateStep()
	if status == models.GrampusMigrateSuccess {
		log.Info("start to move grampus bucket to minio")
		//to move bucket
		if err := schedule.MoveBucketInOpenIMinio(res.DestObjectKey, grampus.GetGPUModelObjectKey(job.ID), res.DestBucket, setting.Attachment.Minio.Bucket); err != nil {
			log.Error("MoveBucketInOpenIMinio err.%v", err)
		} else {
			models.UpdateResultMigrateFlag(job.ID, "MIGRATE_SUCCEED")
		}
		return false
	}
	if status == models.GrampusMigrateFailed || status == models.GrampusMigrateNoNeed {
		return false
	}
	return true
}

func getModelMigrateStatusFromGrampus(jobId string) models.ModelMigrateStep {
	res, err := grampus.ModelMigrateInfo(jobId)
	if err != nil {
		log.Error("ModelMigrateInfo err. r.ID=%d  %v", jobId, err)
		return -1
	}
	log.Info("model convert ModelMigrateInfo r.ID=%d res=%+v", jobId, res)
	return models.GrampusMigrateResponse(res.Status).ConvertToModelMigrateStep()
}

func ListModelFilesForSDK(ctx *context.APIContext) {
	log.Info("ListModelFilesForSDK by api.")

	aiModelList := routerRepo.QueryModelObjByName(ctx.Context)
	if len(aiModelList) == 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  "model not exists",
			"data": nil,
		})
		return
	}

	aiModel := aiModelList[0]
	modelFileList := modelmanage.QueryModelFileByModel(aiModel)
	aiModel.ModelFileList = modelFileList

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  "success",
		"data": aiModel,
	})
}
