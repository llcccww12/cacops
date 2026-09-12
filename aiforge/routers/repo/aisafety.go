package repo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/services/ai_task_service/schedule"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/manager/client/grampus"
	grampus_m "code.gitea.io/gitea/modules/grampus"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/services/ai_task_service/task"

	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/aisafety"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

const (
	tplModelSafetyTestCreateGrampusGpu = "repo/modelsafety/newgrampusgpu"
	tplModelSafetyTestCreateGrampusNpu = "repo/modelsafety/newgrampusnpu"
	tplModelSafetyTestCreateGpu        = "repo/modelsafety/newgpu"
	tplModelSafetyTestCreateNpu        = "repo/modelsafety/newnpu"
	tplModelSafetyTestShow             = "repo/modelsafety/show"
)

func GetAiSafetyTaskByJob(job *models.Cloudbrain) {
	if job == nil {
		log.Error("GetCloudbrainByJobID failed")
		return
	}
	syncAiSafetyTaskStatus(job)
}

func GetAiSafetyTaskTmpl(ctx *context.Context) {
	ctx.Data["id"] = ctx.Params(":id")
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplModelSafetyTestShow)
}

func GetAiSafetyTask(ctx *context.Context) {
	var ID = ctx.Params(":id")
	job, err := models.GetCloudbrainByIDWithDeleted(ID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed:" + err.Error())
		return
	}
	syncAiSafetyTaskStatus(job)
	job, err = models.GetCloudbrainByIDWithDeleted(ID)
	job.BenchmarkType = "安全评测"
	job.BenchmarkTypeName = "Image Classification"
	job.CanModify = cloudbrain.CanModifyJob(ctx, job)
	job.CanDel = cloudbrain.CanDeleteJob(ctx, job)
	if job.Parameters == "{\"parameter\":[]}" {
		job.Parameters = ""
	}
	s, err := resource.GetCloudbrainSpec(job.ID)
	if err == nil {
		job.Spec = s
	}
	user, err := models.GetUserByID(job.UserID)
	if err == nil {
		tmpUser := &models.User{
			Name: user.Name,
		}
		job.User = tmpUser
	}

	ctx.JSON(200, job)
}

func StopAiSafetyTask(ctx *context.Context) {
	log.Info("start to stop the task.")
	var ID = ctx.Params(":id")
	task, err := models.GetCloudbrainByIDWithDeleted(ID)
	result := make(map[string]interface{})
	result["result_code"] = "-1"
	if err != nil {
		log.Info("query task error.err=" + err.Error())
		log.Error("GetCloudbrainByJobID failed:" + err.Error())
		result["msg"] = "No such task."
		ctx.JSON(200, result)
		return
	}
	if isTaskNotFinished(task.Status) {
		if task.Type == models.TypeC2Net {
			log.Info("start to stop grampus task.")
			_, err := grampus.StopJob(&entity.JobIdAndVersionId{JobID: task.JobID, TaskID: task.ID})
			if err != nil {
				log.Info("stop failed.err=" + err.Error())
			}
			task.Status = string(models.JobStopped)
			if task.EndTime == 0 {
				task.EndTime = timeutil.TimeStampNow()
			}
			task.ComputeAndSetDuration()
			err = models.UpdateJob(task)
			if err != nil {
				log.Error("UpdateJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
				result["msg"] = "cloudbrain.Stopped_success_update_status_fail"
				ctx.JSON(200, result)
				return
			}
			//queryTaskStatusFromCloudbrainTwo(job)
		} else if task.Type == models.TypeCloudBrainOne {
			if task.Status == string(models.JobStopped) || task.Status == string(models.JobFailed) || task.Status == string(models.JobSucceeded) {
				log.Error("the job(%s) has been stopped", task.JobName, ctx.Data["msgID"])
				result["msg"] = "cloudbrain.Already_stopped"
				ctx.JSON(200, result)
				return
			}
			err := cloudbrain.StopJob(task.JobID)
			if err != nil {
				log.Error("StopJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
				result["msg"] = "cloudbrain.Stopped_failed"
				ctx.JSON(200, result)
				return
			}
			task.Status = string(models.JobStopped)
			if task.EndTime == 0 {
				task.EndTime = timeutil.TimeStampNow()
			}
			task.ComputeAndSetDuration()
			err = models.UpdateJob(task)
			if err != nil {
				log.Error("UpdateJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
				result["msg"] = "cloudbrain.Stopped_success_update_status_fail"
				ctx.JSON(200, result)
				return
			}
		}
	} else {
		if task.Status == string(models.ModelSafetyTesting) {
			//修改为Failed
			task.Status = string(models.JobStopped)
			if task.EndTime == 0 {
				task.EndTime = timeutil.TimeStampNow()
			}
			task.ComputeAndSetDuration()
			err = models.UpdateJob(task)
			if err != nil {
				log.Error("UpdateJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
				result["msg"] = "cloudbrain.Stopped_success_update_status_fail"
				ctx.JSON(200, result)
				return
			}
		} else {
			log.Info("The job is finished. status=" + task.Status)
		}
	}
	result["result_code"] = "0"
	result["msg"] = "succeed"
	ctx.JSON(200, result)

}

func DelAiSafetyTask(ctx *context.Context) {
	var ID = ctx.Params(":id")
	task, err := models.GetCloudbrainByIDWithDeleted(ID)

	if err != nil {
		log.Error("GetCloudbrainByJobID failed:" + err.Error())
		ctx.ServerError("No such task.", err)
		return
	}
	if task.Status != string(models.JobStopped) && task.Status != string(models.JobFailed) && task.Status != string(models.JobSucceeded) {
		log.Error("the job(%s) has not been stopped", task.JobName, ctx.Data["msgID"])
		ctx.ServerError("the job("+task.JobName+") has not been stopped", nil)
		return
	}
	if task.Type == models.TypeCloudBrainOne {
		cloudbrainTask.DeleteCloudbrainJobStorage(task.JobName, models.TypeCloudBrainOne)
	} else {
		grampus.DeleteJob(&entity.JobIdAndVersionId{JobID: task.JobID, JobType: task.JobType, TaskID: task.ID})

		cloudbrainTask.DeleteCloudbrainJobStorage(task.JobName, models.TypeCloudBrainTwo)
	}
	err = models.DeleteJob(task)
	if err != nil {
		ctx.ServerError(err.Error(), err)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/cloudbrain/benchmark")
}

func syncAiSafetyTaskStatus(job *models.Cloudbrain) {
	log.Info("start to query safety task status.")
	if isTaskNotFinished(job.Status) {
		if job.Type == models.TypeC2Net {
			queryTaskStatusFromGrampus(job)
		} else if job.Type == models.TypeCloudBrainOne {
			queryTaskStatusFromCloudbrain(job)
		}
	} else {
		if job.Status == string(models.ModelSafetyTesting) {
			queryTaskStatusFromModelSafetyTestServer(job)
		} else {
			log.Info("The job is finished. status=" + job.Status)
		}
	}
}

func TimerHandleModelSafetyTestTask() {
	log.Info("start to TimerHandleModelSafetyTestTask")
	tasks, err := models.GetModelSafetyTestTask()
	if err == nil {
		if tasks != nil && len(tasks) > 0 {
			for _, job := range tasks {
				syncAiSafetyTaskStatus(job)
			}
		} else {
			log.Info("query running model safety test task 0.")
		}
	} else {
		log.Info("query running model safety test task err." + err.Error())
	}
}

func queryTaskStatusFromGrampus(job *models.Cloudbrain) {
	jobResult, err := grampus.GetJob(&entity.JobIdAndVersionId{JobID: job.JobID})
	if err != nil {
		log.Info("query train job error." + err.Error())
		return
	}
	oldStatus := job.Status
	job.Status = grampus_m.TransTrainJobStatus(jobResult.JobInfo.Status)

	job.Duration = jobResult.JobInfo.RunSec
	if job.StartTime == 0 && jobResult.JobInfo.StartedAt > 0 && job.Status != models.GrampusStatusWaiting {
		job.StartTime = timeutil.TimeStamp(jobResult.JobInfo.StartedAt)
	}

	if job.EndTime == 0 && models.IsTrainJobTerminal(job.Status) && job.StartTime > 0 {
		job.EndTime = job.StartTime.Add(job.Duration)
	}
	job.CorrectCreateUnix()

	if len(job.AiCenter) == 0 {
		if len(jobResult.JobInfo.Tasks) > 0 {
			if len(jobResult.JobInfo.Tasks[0].CenterID) > 0 && len(jobResult.JobInfo.Tasks[0].CenterName) > 0 {
				job.AiCenter = jobResult.JobInfo.Tasks[0].CenterID[0] + "+" + jobResult.JobInfo.Tasks[0].CenterName[0]
			}
		}
	}

	if job.Status != oldStatus {
		notification.NotifyChangeCloudbrainStatus(job, oldStatus)
	}

	if job.Status != models.GrampusStatusSucceeded {
		log.Info("CloudbrainTwo task status=" + job.Status)
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
	} else {
		log.Info("start to deal ModelSafetyTesting, task status=" + job.Status)
		job.Status = string(models.ModelSafetyTesting)
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
		//send msg to beihang
		go sendGrampusInferenceResultToTest(job)
	}

}

func queryTaskStatusFromCloudbrainTwo(job *models.Cloudbrain) {
	log.Info("The task not finished,name=" + job.DisplayJobName)
	result, err := modelarts.GetTrainJob(job.JobID, strconv.FormatInt(job.VersionID, 10))
	if err != nil {
		log.Info("query train job error." + err.Error())
		return
	}

	job.Status = modelarts.TransTrainJobStatus(result.IntStatus)
	job.Duration = result.Duration / 1000
	job.TrainJobDuration = result.TrainJobDuration

	if job.StartTime == 0 && result.StartTime > 0 {
		job.StartTime = timeutil.TimeStamp(result.StartTime / 1000)
	}
	job.TrainJobDuration = models.ConvertDurationToStr(job.Duration)
	if job.EndTime == 0 && models.IsTrainJobTerminal(job.Status) && job.StartTime > 0 {
		job.EndTime = job.StartTime.Add(job.Duration)
	}
	job.CorrectCreateUnix()

	if job.Status != string(models.ModelArtsTrainJobCompleted) {
		log.Info("CloudbrainTwo task status=" + job.Status)
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
	} else {
		log.Info("start to deal ModelSafetyTesting, task status=" + job.Status)
		job.Status = string(models.ModelSafetyTesting)
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
		//send msg to beihang
		sendNPUInferenceResultToTest(job)
	}

}

func queryTaskStatusFromCloudbrain(job *models.Cloudbrain) {

	log.Info("The task not finished,name=" + job.DisplayJobName)
	jobResult, err := cloudbrain.GetJob(job.JobID)

	result, err := models.ConvertToJobResultPayload(jobResult.Payload)
	if err != nil {
		log.Error("ConvertToJobResultPayload failed:", err)
		return
	}
	job.Status = result.JobStatus.State
	if result.JobStatus.State != string(models.JobWaiting) && result.JobStatus.State != string(models.JobFailed) {
		taskRoles := result.TaskRoles
		taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))
		job.Status = taskRes.TaskStatuses[0].State
	}
	models.ParseAndSetDurationFromCloudBrainOne(result, job)
	//updateCloudBrainOneJobTime(job)
	log.Info("cloud brain one job status=" + job.Status)
	if result.JobStatus.State != string(models.JobSucceeded) {
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
	} else {
		//
		job.Status = string(models.ModelSafetyTesting)
		job.EndTime = 0
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
		//send msg to beihang
		sendGPUInferenceResultToTest(job)
	}
}

func queryTaskStatusFromModelSafetyTestServer(job *models.Cloudbrain) {
	if job.PreVersionName != "" {

		result, err := aisafety.GetTaskStatus(job.PreVersionName)
		if err == nil {
			if result.Code == "0" {
				if result.Data.Status == 1 {
					log.Info("The task is running....")
				} else {
					job.EndTime = timeutil.TimeStampNow()
					job.Duration = (job.EndTime.AsTime().Unix() - job.StartTime.AsTime().Unix())
					job.TrainJobDuration = models.ConvertDurationToStr(job.Duration)
					if result.Data.Code == 0 {
						job.ResultJson = result.Data.StandardJson
						job.Status = string(models.JobSucceeded)
						err = models.UpdateJob(job)
						if err != nil {
							log.Error("UpdateJob failed:", err)
						}
					} else {
						job.ResultJson = result.Data.Msg
						job.Status = string(models.JobFailed)
						err = models.UpdateJob(job)
						if err != nil {
							log.Error("UpdateJob failed:", err)
						}
					}
				}
			} else {
				log.Info("The task is failed.")
				job.Status = string(models.JobFailed)
				err = models.UpdateJob(job)
				if err != nil {
					log.Error("UpdateJob failed:", err)
				}
			}
		} else {
			log.Info("The task not found.....")
		}
	}
}

func getAisafetyTaskReq(job *models.Cloudbrain) aisafety.TaskReq {
	datasetname := job.DatasetName
	datasetnames := strings.Split(datasetname, ";")
	indicator := job.LabelName
	EvalContent := "test1"
	if job.Description != "" {
		EvalContent = job.Description
	}
	req := aisafety.TaskReq{
		UnionId:     job.JobID,
		EvalName:    job.DisplayJobName,
		EvalContent: EvalContent,
		TLPath:      "test1",
		Indicators:  strings.Split(indicator, ";"),
		CDName:      strings.Split(datasetnames[1], ".")[0],
		BDName:      strings.Split(datasetnames[0], ".")[0] + "基础数据集",
	}
	log.Info("CDName=" + req.CDName)
	log.Info("BDName=" + req.BDName)
	return req
}

func sendGPUInferenceResultToTest(job *models.Cloudbrain) {
	log.Info("send sendGPUInferenceResultToTest")
	req := getAisafetyTaskReq(job)
	resultDir := "/result"
	prefix := setting.CBCodePathPrefix + job.JobName + resultDir
	files, err := storage.GetOneLevelAllObjectUnderDirMinio(setting.Attachment.Minio.Bucket, prefix, "")
	if err != nil {
		log.Error("query cloudbrain one model failed: %v", err)
		return
	}
	jsonContent := ""
	for _, file := range files {
		if strings.HasSuffix(file.FileName, "result.json") {
			path := storage.GetMinioPath(job.JobName+resultDir+"/", file.FileName)
			log.Info("path=" + path)
			reader, err := os.Open(path)
			defer reader.Close()
			if err == nil {
				r := bufio.NewReader(reader)
				for {
					line, error := r.ReadString('\n')
					jsonContent += line
					if error == io.EOF {
						log.Info("read file completed.")
						break
					}
					if error != nil {
						log.Info("read file error." + error.Error())
						break
					}
				}
			}
			break
		}
	}
	if jsonContent != "" {
		sendHttpReqToBeihang(job, jsonContent, req)
	} else {
		updateJobFailed(job, "推理生成的Json数据为空，无法进行评测。")
	}
}

func sendGrampusInferenceResultToTest(job *models.Cloudbrain) {
	log.Info("start to sendNPUInferenceResultToTest")
	if isMigrateFinished(job) {

		req := getAisafetyTaskReq(job)
		jsonContent := ""
		resultPath := modelarts.JobPath + job.JobName + "/" + setting.OutPutPath + "result.json"
		resultPath = resultPath[1:]
		log.Info("bucket=" + setting.Bucket + "  resultPath=" + resultPath)
		body, err := storage.ObsDownloadAFile(setting.Bucket, resultPath)
		if err != nil {
			log.Info("ObsDownloadAFile  error." + err.Error() + " resultPath=" + resultPath)
		} else {
			defer body.Close()
			var data []byte
			p := make([]byte, 4096)
			var readErr error
			var readCount int
			for {
				readCount, readErr = body.Read(p)
				if readCount > 0 {
					data = append(data, p[:readCount]...)
				}
				if readErr != nil || readCount == 0 {
					break
				}
			}
			jsonContent = string(data)
		}

		if jsonContent != "" {
			sendHttpReqToBeihang(job, jsonContent, req)
		} else {
			updateJobFailed(job, "推理生成的Json数据为空，无法进行评测。")
		}
	} else {
		updateJobFailed(job, "推理生成的Json数据为空，无法进行评测。")
	}
}

func isMigrateFinished(job *models.Cloudbrain) bool {
	for i := 0; i < 20; i++ {
		migrate, err := models.GetModelMigrateRecordByCloudbrainId(job.ID)
		if err != nil {
			log.Error("GetModelMigrateRecordByCloudbrainId error.", err)
			return false
		}
		if migrate.Status == models.ModelMigrateSuccess {
			return true
		} else if migrate.Status == models.ModelMigrateFailed {
			schedule.RetryModelMigrate(job)
			time.Sleep(time.Second * 30)
		} else {
			time.Sleep(time.Second * 30)
		}

	}
	return false
}

func sendNPUInferenceResultToTest(job *models.Cloudbrain) {
	log.Info("start to sendNPUInferenceResultToTest")
	req := getAisafetyTaskReq(job)
	jsonContent := ""
	VersionOutputPath := modelarts.GetOutputPathByCount(modelarts.TotalVersionCount)
	resultPath := modelarts.JobPath + job.JobName + modelarts.ResultPath + VersionOutputPath + "/result.json"
	resultPath = resultPath[1:]
	log.Info("bucket=" + setting.Bucket + "  resultPath=" + resultPath)
	body, err := storage.ObsDownloadAFile(setting.Bucket, resultPath)
	if err != nil {
		log.Info("ObsDownloadAFile  error." + err.Error() + " resultPath=" + resultPath)
	} else {
		defer body.Close()
		var data []byte
		p := make([]byte, 4096)
		var readErr error
		var readCount int
		for {
			readCount, readErr = body.Read(p)
			if readCount > 0 {
				data = append(data, p[:readCount]...)
			}
			if readErr != nil || readCount == 0 {
				break
			}
		}
		jsonContent = string(data)
	}

	if jsonContent != "" {
		sendHttpReqToBeihang(job, jsonContent, req)
	} else {
		updateJobFailed(job, "推理生成的Json数据为空，无法进行评测。")
	}
}
func updateJobFailed(job *models.Cloudbrain, msg string) {
	log.Info("The json is null. so set it failed.")
	//update task failed.
	job.Status = string(models.ModelArtsTrainJobFailed)
	job.ResultJson = msg
	job.EndTime = timeutil.TimeStampNow()
	job.Duration = (job.EndTime.AsTime().Unix() - job.StartTime.AsTime().Unix()) / 1000
	job.TrainJobDuration = models.ConvertDurationToStr(job.Duration)
	err := models.UpdateJob(job)
	if err != nil {
		log.Error("UpdateJob failed:", err)
	}
}
func sendHttpReqToBeihang(job *models.Cloudbrain, jsonContent string, req aisafety.TaskReq) {
	log.Info("start to send beihang ...")
	serialNo, err := aisafety.CreateSafetyTask(req, jsonContent)
	if err == nil {
		//update serial no to db
		job.PreVersionName = serialNo
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
		}
	}
}

func isTaskNotFinished(status string) bool {
	if status == string(models.ModelArtsTrainJobRunning) || status == string(models.ModelArtsTrainJobWaiting) {
		return true
	}
	if status == string(models.JobWaiting) || status == string(models.JobRunning) {
		return true
	}

	if status == string(models.ModelArtsTrainJobUnknown) || status == string(models.ModelArtsTrainJobInit) {
		return true
	}
	if status == string(models.ModelArtsTrainJobImageCreating) || status == string(models.ModelArtsTrainJobSubmitTrying) {
		return true
	}
	return false
}

func AiSafetyCreateForGetGPU(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["IsCreate"] = true
	ctx.Data["type"] = models.TypeCloudBrainOne
	ctx.Data["compute_resource"] = models.GPUResource
	ctx.Data["datasetType"] = models.TypeCloudBrainOne
	ctx.Data["BaseDataSetName"] = setting.ModelSafetyTest.GPUBaseDataSetName
	ctx.Data["BaseDataSetUUID"] = setting.ModelSafetyTest.GPUBaseDataSetUUID
	ctx.Data["CombatDataSetName"] = setting.ModelSafetyTest.GPUCombatDataSetName
	ctx.Data["CombatDataSetUUID"] = setting.ModelSafetyTest.GPUCombatDataSetUUID
	log.Info("GPUBaseDataSetName=" + setting.ModelSafetyTest.GPUBaseDataSetName)
	log.Info("GPUBaseDataSetUUID=" + setting.ModelSafetyTest.GPUBaseDataSetUUID)
	log.Info("GPUCombatDataSetName=" + setting.ModelSafetyTest.GPUCombatDataSetName)
	log.Info("GPUCombatDataSetUUID=" + setting.ModelSafetyTest.GPUCombatDataSetUUID)
	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName
	prepareCloudbrainOneSpecs(ctx)
	queuesDetail, _ := cloudbrain.GetQueuesDetail()
	if queuesDetail != nil {
		ctx.Data["QueuesDetail"] = queuesDetail
		reqPara, _ := json.Marshal(queuesDetail)
		log.Warn("The GPU WaitCount json:", string(reqPara))
	} else {
		log.Info("The GPU WaitCount not get")
	}
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeBenchmark))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount
	ctx.HTML(200, tplModelSafetyTestCreateGpu)
}

func AiSafetyCreateForGetNPU(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["IsCreate"] = true
	ctx.Data["type"] = models.TypeC2Net
	ctx.Data["compute_resource"] = models.NPUResource
	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName
	ctx.Data["datasetType"] = models.TypeCloudBrainTwo
	ctx.Data["BaseDataSetName"] = setting.ModelSafetyTest.NPUBaseDataSetName
	ctx.Data["BaseDataSetUUID"] = setting.ModelSafetyTest.NPUBaseDataSetUUID
	ctx.Data["CombatDataSetName"] = setting.ModelSafetyTest.NPUCombatDataSetName
	ctx.Data["CombatDataSetUUID"] = setting.ModelSafetyTest.NPUCombatDataSetUUID

	log.Info("NPUBaseDataSetName=" + setting.ModelSafetyTest.NPUBaseDataSetName)
	log.Info("NPUBaseDataSetUUID=" + setting.ModelSafetyTest.NPUBaseDataSetUUID)
	log.Info("NPUCombatDataSetName=" + setting.ModelSafetyTest.NPUCombatDataSetName)
	log.Info("NPUCombatDataSetUUID=" + setting.ModelSafetyTest.NPUCombatDataSetUUID)

	noteBookSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.NPU,
		Cluster:         models.C2NetCluster,
	})
	if len(noteBookSpecs) > 0 {
		ctx.Data["Specs"] = []*models.Specification{noteBookSpecs[0]}

		computeSource := models.GetComputeSourceInstance(models.NPU)
		result, bizerr := task.GetAvailableImageInfoBySpec(entity.GetAITaskCreationImageInfoReq{
			ClusterType:   entity.C2Net,
			ComputeSource: computeSource,
			Spec:          noteBookSpecs[0],
			JobType:       models.JobTypeTrain,
			UserID:        ctx.User.ID,
		})
		if bizerr != nil {
			log.Warn("get image err", bizerr.ToError())
		}

		ctx.Data["imageInfo"] = result

	} else {
		ctx.Data["Specs"] = []*models.Specification{}
	}

	waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeC2Net, models.NPUResource, models.JobTypeModelSafety)
	ctx.Data["WaitCount"] = waitCount
	log.Info("The NPU WaitCount is " + fmt.Sprint(waitCount))
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeModelSafety))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount
	ctx.HTML(200, tplModelSafetyTestCreateNpu)
}
