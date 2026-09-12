package reward

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/repo"
)

func NewRewardPeriodicTask(operateRecordId string, opts *models.StartPeriodicTaskOpts) error {
	task := &models.RewardPeriodicTask{}
	task.DelaySeconds = int64(opts.Delay.Seconds())
	task.IntervalSeconds = int64(opts.Interval.Seconds())
	task.Amount = opts.UnitAmount
	task.OperateSerialNo = operateRecordId
	task.Status = models.PeriodicTaskStatusRunning
	task.NextExecuteTime = timeutil.TimeStamp(opts.StartTime.Add(opts.Delay).Unix())

	_, err := models.InsertPeriodicTask(task)
	return err
}

func StartRewardTask() {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	log.Debug("try to run reward tasks")
	now := time.Now()
	taskList, err := models.GetRunningRewardTask(now)
	if err != nil {
		log.Error("GetRunningRewardTask error. %v", err)
		return
	}
	if taskList == nil || len(taskList) == 0 {
		log.Debug("No GetRunningRewardTask need handled")
		return
	}
	for _, t := range taskList {
		RunRewardTask(t, now)
	}
}

func RunRewardTask(t models.RewardPeriodicTask, now time.Time) error {
	lock := redis_lock.NewDistributeLock(redis_key.RewardTaskRunningLock(t.ID))
	isOk, _ := lock.LockWithWait(5*time.Second, 5*time.Second)
	if !isOk {
		log.Error("get RewardTaskRunningLock failed,t=%+v", t)
		return errors.New("get RewardTaskRunningLock failed")
	}
	defer lock.UnLock()
	record, err := models.GetPointOperateRecordBySerialNo(t.OperateSerialNo)
	if err != nil {
		log.Error("RunRewardTask. GetPointOperateRecordBySerialNo error. %v", err)
		return errors.New("GetPointOperateRecordBySerialNo error")
	}
	if record.Status != models.OperateStatusOperating {
		log.Info("RunRewardTask. operate record is finished,record=%+v", record)
		return nil
	}
	task, err := models.GetCloudbrainByID(record.SourceId)
	if task == nil {
		log.Info("RunRewardTask task is not exists.taskID = %d", record.SourceId)
		return nil
	}
	log.Info("RunRewardTask. operate record=%+v", record)
	n, _ := countExecuteTimes(t, now)
	if n == 0 {
		log.Info("countExecuteTimes result is 0")
		TryToNotifyInsufficientBalance(task, t.Amount, t.NextExecuteTime)
		return nil
	}
	log.Info("RunRewardTask. n=%d  now=%v", n, now)
	//get operator
	operator := GetOperator(models.GetRewardTypeInstance(record.RewardType))
	if operator == nil {
		log.Error("RunRewardTask. operator of reward type is not exist")
		return errors.New("operator of reward type is not exist")
	}
	nextTime := timeutil.TimeStamp(int64(t.NextExecuteTime) + t.IntervalSeconds)

	if err != nil {
		log.Error("RunRewardTask GetCloudbrainByID error. %v", err)
		return err
	}

	log.Info("RunRewardTask n=%d", n)
	for i := 1; int64(i) <= n; i++ {
		log.Debug("operator.Operate i=%d n=%d", i, n)
		err = operator.Operate(&models.RewardOperateContext{
			SourceType: models.SourceTypeRunCloudbrainTask,
			SourceId:   t.OperateSerialNo,
			Reward: models.Reward{
				Amount: t.Amount,
				Type:   models.GetRewardTypeInstance(record.RewardType),
			},
			TargetUserId: record.UserId,
			OperateType:  models.GetRewardOperateTypeInstance(record.OperateType),
		})
		if err != nil {
			log.Error("RunRewardTask.operator operate error.%v", err)
			if models.IsErrInsufficientPointsBalance(err) {
				tryToStopCloudbrainTask(task.ID, t.ID, t.OperateSerialNo)
				return nil
			}
			return nil
		}

		models.IncrRewardTaskSuccessCount(t, 1, nextTime)
		TryToNotifyInsufficientBalance(task, t.Amount, nextTime)
		nextTime = timeutil.TimeStamp(int64(nextTime) + t.IntervalSeconds)
	}
	return nil

}

func tryToStopCloudbrainTask(cloudbrainId int64, periodTaskId int64, operateSerialNo string) {

	t, err := models.GetCloudbrainByID(fmt.Sprint(cloudbrainId))
	if err != nil {
		log.Error("tryToStopCloudbrainTask GetCloudbrainByID error,task.ID=%d  %v", cloudbrainId, err)
		if !models.IsErrJobNotExist(err) {
			var sourceContent string
			if contentByte, err := json.Marshal(t.ToShow()); err == nil {
				sourceContent = string(contentByte)
			}
			models.StopPeriodicTask(periodTaskId, operateSerialNo, time.Now(), sourceContent)
		}
		return
	}

	isTerminal := false
	//可能已经被其他线程更新为终态
	if !t.IsTerminalOrStopping() {
		repo.StopJobs([]*models.Cloudbrain{t})
		t, _ = models.GetCloudbrainByID(fmt.Sprint(cloudbrainId))
		isTerminal = t.IsTerminalOrStopping()
	} else {
		isTerminal = true
	}

	if !isTerminal {
		log.Error("tryToStopCloudbrainTask stop cloudbrain task failed.task.ID=%d ", cloudbrainId)
		return
	}
	var sourceContent string
	if contentByte, err := json.Marshal(t.ToShow()); err == nil {
		sourceContent = string(contentByte)
	}
	err = models.StopPeriodicTask(periodTaskId, operateSerialNo, time.Now(), sourceContent)
	if err != nil {
		log.Error("tryToStopCloudbrainTask StopPeriodicTask err.%v", err)
	}
}

func countExecuteTimes(t models.RewardPeriodicTask, now time.Time) (int64, timeutil.TimeStamp) {
	interval := t.IntervalSeconds
	nextTime := int64(t.NextExecuteTime)
	if nextTime > now.Unix() {
		return 0, 0
	}
	diff := now.Unix() - nextTime
	var n int64
	if diff%interval == 0 {
		n = diff / interval
	} else {
		n = diff/interval + 1
	}

	newNextTime := timeutil.TimeStamp(nextTime + n*interval)
	return n, newNextTime
}
