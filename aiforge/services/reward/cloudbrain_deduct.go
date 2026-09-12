package reward

import (
	"encoding/json"
	"fmt"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

var (
	ResourceSpecs      *models.ResourceSpecs
	TrainResourceSpecs *models.ResourceSpecs
)

const RUN_CLOUDBRAIN_TASK_TITTLE = "运行云脑任务"

func AcceptStatusChangeAction() {
	for {
		select {
		case task := <-models.StatusChangeChan:
			log.Info("AcceptStatusChangeAction task.Id=%d  status=%s", task.ID, task.Status)
			DeductPoint4Cloudbrain(*task, time.Now())

			// 如果任务状态变为失败、主动停止或正在停止中，清理关联的 user_agent 记录（成功完成不清理）
			if task.IsFailedOrStopped() {
				log.Info("DeleteUserAgentByCloudBrainIDOnly task.UserID=%d task.Id=%d status=%s", task.UserID, task.ID, task.Status)
				// 判断删除原因
				deleteReason := "STOPPED"
				if task.Status == string(models.JobFailed) || task.Status == string(models.ModelArtsTrainJobFailed) ||
					task.Status == string(models.ModelArtsCreateFailed) || task.Status == string(models.ModelArtsStartFailed) ||
					task.Status == string(models.GrampusStatusFailed) || task.Status == string(models.LocalStatusFailed) {
					deleteReason = "FAILED"
				}
				models.DeleteUserAgentByCloudBrainIDOnly(task.ID, deleteReason, task.Status)
			}
		}
	}
}

func StartAndGetCloudBrainPointDeductTask(task models.Cloudbrain) (*models.RewardPeriodicTask, error) {
	sourceId := GetCloudBrainPointTaskSourceId(task)
	r, err := GetPeriodicTask(models.SourceTypeRunCloudbrainTask, sourceId, sourceId, models.OperateTypeDecrease)
	if err != nil {
		return nil, err
	}

	if r != nil {
		log.Debug("PeriodicTask is already exist.cloudbrain.ID = %d", task.ID)
		return r, nil
	}

	if !setting.CloudBrainPaySwitch {
		log.Debug("CloudBrainPaySwitch is off")
		return nil, nil
	}

	unitPrice, err := models.GetCloudbrainTaskUnitPrice(task)
	if err != nil {
		return nil, err
	}
	if unitPrice == 0 {
		log.Debug("Finish startAndGetCloudBrainPointDeductTask,  UnitPrice = 0 task.ID=%d", task.ID)
		return nil, nil
	}

	startTime := time.Unix(int64(task.StartTime), 0)
	if int64(task.StartTime) < setting.DeductTaskMinTimestamp {
		startTime = time.Unix(setting.DeductTaskMinTimestamp, 0)
	}

	unitAmount := unitPrice * float64(setting.CloudBrainPayInterval) / float64(time.Hour)

	var sourceContent string
	if contentByte, err := json.Marshal(task.ToShow()); err == nil {
		sourceContent = string(contentByte)
	}

	return StartPeriodicTask(&models.StartPeriodicTaskOpts{
		SourceType:    models.SourceTypeRunCloudbrainTask,
		SourceId:      GetCloudBrainPointTaskSourceId(task),
		TargetUserId:  task.UserID,
		RequestId:     GetCloudBrainPointTaskSourceId(task),
		OperateType:   models.OperateTypeDecrease,
		Delay:         setting.CloudBrainPayDelay,
		Interval:      setting.CloudBrainPayInterval,
		UnitAmount:    unitAmount,
		RewardType:    models.RewardTypePoint,
		StartTime:     startTime,
		Title:         RUN_CLOUDBRAIN_TASK_TITTLE,
		SourceContent: sourceContent,
	})
}

func StopCloudBrainPointDeductTask(task models.Cloudbrain) {
	var sourceContent string
	if contentByte, err := json.Marshal(task.ToShow()); err == nil {
		sourceContent = string(contentByte)
	}
	StopPeriodicTask(models.SourceTypeRunCloudbrainTask, GetCloudBrainPointTaskSourceId(task), models.OperateTypeDecrease, sourceContent)
}

func GetCloudBrainPointTaskSourceId(task models.Cloudbrain) string {
	return fmt.Sprint(task.ID)
}

var firstTimeFlag = true

func StartCloudbrainPointDeductTask() {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	log.Info("try to run CloudbrainPointDeductTask")
	end := time.Now()
	start := end.Add(-1 * setting.DeductTaskRange)
	if firstTimeFlag {
		//When it is executed for the first time, it needs to process the tasks of the last 3 hours.
		//This is done to prevent the application from hanging for a long time
		start = end.Add(-1 * setting.DeductTaskRangeForFirst)
		firstTimeFlag = false
	}
	taskList, err := models.GetStartedCloudbrainTaskByUpdatedUnix(start, end)
	if err != nil {
		log.Error("GetStartedCloudbrainTaskByUpdatedUnix error. %v", err)
		return
	}
	if taskList == nil || len(taskList) == 0 {
		log.Debug("No cloudbrain task need handled")
		return
	}
	for _, t := range taskList {
		log.Info("try to DeductPoint4Cloudbrain by CloudbrainPointDeductTask.task.id=%d", t.ID)
		DeductPoint4Cloudbrain(t, end)
	}
	log.Debug("CloudbrainPointDeductTask completed")
}

func DeductPoint4Cloudbrain(t models.Cloudbrain, now time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	log.Info("start to deduct point for cloudbrain[%d]", t.ID)
	if t.StartTime == 0 {
		log.Debug("cloudbrain[%d] task not start", t.ID)
		return nil
	}

	task, err := StartAndGetCloudBrainPointDeductTask(t)
	if err != nil {
		log.Error("run cloudbrain point deduct task error,err=%v", err)
		return err
	}
	if task == nil {
		log.Info("cloudbrain[%d] deduct task is nil", t.ID)
		return nil
	}
	if task.Status == models.PeriodicTaskStatusFinished {
		log.Info("Periodic task is finished")
		return nil
	}

	if t.EndTime > 0 {
		log.Info("cloudbrain[%d] task is finished,try to stop PeriodicTask.endTimeStamp = %d", t.ID, t.EndTime)
		endTime := time.Unix(int64(t.EndTime), 0)
		RunRewardTask(*task, endTime)
		var sourceContent string
		if contentByte, err := json.Marshal(t.ToShow()); err == nil {
			sourceContent = string(contentByte)
		}
		models.StopPeriodicTask(task.ID, task.OperateSerialNo, endTime, sourceContent)
	} else {
		log.Info("try to RunRewardTask.cloudbrain.id = %d", t.ID)
		RunRewardTask(*task, now)
	}
	log.Debug("finished deduct point for cloudbrain[%d]", t.ID)
	return nil
}
