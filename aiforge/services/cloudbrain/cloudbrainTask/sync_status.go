package cloudbrainTask

import (
	"net/http"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/modelarts_cd"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

var noteBookOKMap = make(map[int64]int, 20)
var noteBookFailMap = make(map[int64]int, 20)

// if a task notebook url can get successfulCount times,  the notebook can browser.
const successfulCount = 3
const maxSuccessfulCount = 10

func SyncCloudBrainOneStatus(task *models.Cloudbrain) (*models.Cloudbrain, error) {
	jobResult, err := cloudbrain.GetJob(task.JobID)
	if err != nil {

		log.Error("GetJob failed:", err)

		return task, err
	}
	result, err := models.ConvertToJobResultPayload(jobResult.Payload)
	if err != nil {
		log.Error("ConvertToJobResultPayload failed:", err)
		return task, err
	}
	oldStatus := task.Status

	if result.JobStatus.State != string(models.JobWaiting) && result.JobStatus.State != string(models.JobFailed) {
		taskRoles := result.TaskRoles
		taskRes, _ := models.ConvertToTaskPod(taskRoles[cloudbrain.SubTaskName].(map[string]interface{}))

		task.ContainerIp = taskRes.TaskStatuses[0].ContainerIP
		task.ContainerID = taskRes.TaskStatuses[0].ContainerID
	}

	if (result.JobStatus.State != string(models.JobWaiting) && result.JobStatus.State != string(models.JobRunning)) ||
		task.Status == string(models.JobRunning) || (result.JobStatus.State == string(models.JobRunning) && isNoteBookReady(task)) {

		models.ParseAndSetDurationFromCloudBrainOne(result, task)
		task.Status = result.JobStatus.State
		if oldStatus != task.Status {
			notification.NotifyChangeCloudbrainStatus(task, oldStatus)
		}
		err = models.UpdateJob(task)
		if err != nil {
			log.Error("UpdateJob failed:", err)
			return task, err
		}
	}
	return task, nil

}

func SyncGrampusNotebookStatus(job *models.Cloudbrain) (*models.Cloudbrain, error) {
	result, err := grampus.GetNotebookJob(job.JobID)
	if err != nil {

		log.Error("GetJob(%s) failed:%v", job.JobName, err)

		return job, err
	}

	oldStatus := job.Status
	job.Status = grampus.TransTrainJobStatus(result.JobInfo.Status)
	job.Duration = result.JobInfo.RunSec
	job.TrainJobDuration = models.ConvertDurationToStr(job.Duration)

	if job.StartTime == 0 && result.JobInfo.StartedAt > 0 && job.Status != models.GrampusStatusWaiting {
		job.StartTime = timeutil.TimeStamp(result.JobInfo.StartedAt)
	}

	if job.EndTime == 0 && models.IsTrainJobTerminal(job.Status) && job.StartTime > 0 {
		job.EndTime = job.StartTime.Add(job.Duration)
	}
	job.CorrectCreateUnix()

	if len(job.AiCenter) == 0 {
		if len(result.JobInfo.Tasks) > 0 {
			if len(result.JobInfo.Tasks[0].CenterID) > 0 && len(result.JobInfo.Tasks[0].CenterName) > 0 {
				job.AiCenter = result.JobInfo.Tasks[0].CenterID[0] + "+" + result.JobInfo.Tasks[0].CenterName[0]
			}
		}
	}

	if job.Status != models.GrampusStatusWaiting {
		if oldStatus != job.Status {
			notification.NotifyChangeCloudbrainStatus(job, oldStatus)
		}
		if job.ComputeResource == models.NPUResource {
			if len(result.JobInfo.Tasks) > 0 {
				job.TrainUrl = result.JobInfo.Tasks[0].CodeUrl
				job.DataUrl = result.JobInfo.Tasks[0].DataUrl
			}
		}
		err = models.UpdateJob(job)
		if err != nil {
			log.Error("UpdateJob failed:", err)
			return nil, err
		}
	}

	return job, nil

}

func isNoteBookReady(task *models.Cloudbrain) bool {
	if task.JobType != string(models.JobTypeDebug) {
		return true
	}
	noteBookUrl := setting.DebugServerHost + "jpylab_" + task.JobID + "_" + task.SubTaskName
	res, err := http.Get(noteBookUrl)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	log.Info("notebook success count:" + strconv.Itoa(noteBookOKMap[task.ID]) + ",fail count:" + strconv.Itoa(noteBookFailMap[task.ID]))
	if res.StatusCode == http.StatusOK {
		count := noteBookOKMap[task.ID]
		if count == 0 { //如果是第一次成功，把失败数重置为0
			noteBookFailMap[task.ID] = 0
		}

		if count < successfulCount-1 || (noteBookFailMap[task.ID] == 0 && count < maxSuccessfulCount-1) {
			noteBookOKMap[task.ID] = count + 1
			return false
		} else {
			log.Info("notebook success count:" + strconv.Itoa(count) + ",fail count:" + strconv.Itoa(noteBookFailMap[task.ID]))
			delete(noteBookOKMap, task.ID)
			delete(noteBookFailMap, task.ID)
			return true
		}

	} else {
		noteBookFailMap[task.ID] += 1
	}
	return false

}

func StopDebugJob(task *models.Cloudbrain) error {
	param := models.NotebookAction{
		Action: models.ActionStop,
	}
	var err error = nil

	if task.JobType == string(models.JobTypeDebug) {
		if task.Type == models.TypeCloudBrainOne {
			return cloudbrain.StopJob(task.JobID)
		} else if task.Type == models.TypeCloudBrainTwo {
			_, err = modelarts.ManageNotebook2(task.JobID, param)

		} else if task.Type == models.TypeCDCenter {
			_, err = modelarts_cd.ManageNotebook(task.JobID, param)

		} else if task.Type == models.TypeC2Net {
			_, err = grampus.StopJob(task.JobID, task.JobType)

		}

	}

	if task.JobType == string(models.JobTypeSuperCompute) {
		_, err = grampus.StopJob(task.JobID, task.JobType)
	}

	return err

}
