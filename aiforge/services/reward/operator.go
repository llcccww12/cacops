package reward

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/services/reward/point"
	"errors"
	"fmt"
	"time"
)

var RewardOperatorMap = map[string]RewardOperator{
	fmt.Sprint(models.RewardTypePoint): new(point.PointOperator),
}

type RewardOperator interface {
	IsLimited(ctx *models.RewardOperateContext) error
	Operate(ctx *models.RewardOperateContext) error
}

func Operate(ctx *models.RewardOperateContext) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	if !checkRewardOperationParam(ctx) {
		log.Error("send reward error,param incorrect")
		return errors.New("param incorrect")
	}
	//add lock
	var rewardLock = redis_lock.NewDistributeLock(redis_key.RewardOperateLock(ctx.RequestId, ctx.SourceType.Name(), ctx.OperateType.Name()))
	isOk, err := rewardLock.Lock(3 * time.Second)
	if err != nil {
		return err
	}
	if !isOk {
		log.Info("duplicated reward request,targetUserId=%d requestId=%s", ctx.TargetUserId, ctx.RequestId)
		return nil
	}
	defer rewardLock.UnLock()

	//is handled before?
	isHandled, err := isHandled(ctx.SourceType.Name(), ctx.RequestId, ctx.OperateType.Name())
	if err != nil {
		log.Error("reward is handled error,%v", err)
		return err
	}
	if isHandled {
		log.Info("reward has been handled,ctx=%+v", ctx)
		return nil
	}

	//get operator
	operator := GetOperator(ctx.Reward.Type)
	if operator == nil {
		log.Error("operator of reward type is not exist,ctx=%v", ctx)
		return errors.New("operator of reward type is not exist")
	}

	if ctx.OperateType == models.OperateTypeIncrease {
		//is limited?
		if err := operator.IsLimited(ctx); err != nil {
			log.Info("operator IsLimited, err=%v", err)
			return err
		}
	}

	//new reward operate record
	recordId, err := initRewardOperateRecord(ctx)
	if err != nil {
		log.Error("initRewardOperateRecord error,err=%v", err)
		return err
	}

	ctx.SourceId = recordId

	//operate
	if err := operator.Operate(ctx); err != nil {
		log.Error("operator Operate error,err=%v", err)
		UpdateRewardRecordToFinalStatus(ctx.SourceType.Name(), ctx.RequestId, models.OperateStatusFailed)
		return err
	}

	UpdateRewardRecordToFinalStatus(ctx.SourceType.Name(), ctx.RequestId, models.OperateStatusSucceeded)
	NotifyRewardOperation(ctx.TargetUserId, ctx.Reward.Amount, ctx.SourceType, ctx.Reward.Type, ctx.OperateType)
	return nil
}

func checkRewardOperationParam(ctx *models.RewardOperateContext) bool {
	if ctx.Reward.Type == "" {
		return false
	}
	return true
}

func GetOperator(rewardType models.RewardType) RewardOperator {
	return RewardOperatorMap[rewardType.Name()]
}

func isHandled(sourceType string, requestId string, operateType string) (bool, error) {
	_, err := models.GetPointOperateRecordBySourceTypeAndRequestId(sourceType, requestId, operateType)
	if err != nil {
		log.Error("operator isHandled error. %v", err)
		if models.IsErrRecordNotExist(err) {
			return false, nil
		}
		log.Error("GetPointOperateRecordBySourceTypeAndRequestId ZRangeByScore error. %v", err)
		return false, err
	}
	return true, nil

}

func initRewardOperateRecord(ctx *models.RewardOperateContext) (string, error) {
	sn, err := generateOperateSerialNo()
	if err != nil {
		log.Error("generateOperateSerialNo error. %v", err)
		return "", err
	}
	record := &models.RewardOperateRecord{
		UserId:           ctx.TargetUserId,
		Amount:           ctx.Reward.Amount,
		LossAmount:       ctx.LossAmount,
		RewardType:       ctx.Reward.Type.Name(),
		SourceType:       ctx.SourceType.Name(),
		SourceId:         ctx.SourceId,
		SourceTemplateId: ctx.SourceTemplateId,
		RequestId:        ctx.RequestId,
		OperateType:      ctx.OperateType.Name(),
		Status:           models.OperateStatusOperating,
		Remark:           ctx.Remark,
		Title:            ctx.Title,
		SerialNo:         sn,
		SourceContent:    ctx.SourceContent,
	}
	_, err = models.InsertRewardOperateRecord(record)
	if err != nil {
		log.Error("InsertRewardOperateRecord error. %v", err)
		return "", err
	}
	return record.SerialNo, nil
}

func createPeriodicRewardOperateRecord(ctx *models.StartPeriodicTaskOpts) (string, error) {
	sn, err := generateOperateSerialNo()
	if err != nil {
		log.Error("createPeriodic generateOperateSerialNo error. %v", err)
		return "", err
	}
	record := &models.RewardOperateRecord{
		UserId:        ctx.TargetUserId,
		Amount:        0,
		RewardType:    ctx.RewardType.Name(),
		SourceType:    ctx.SourceType.Name(),
		SourceId:      ctx.SourceId,
		RequestId:     ctx.RequestId,
		OperateType:   ctx.OperateType.Name(),
		Status:        models.OperateStatusOperating,
		Remark:        ctx.Remark,
		Title:         ctx.Title,
		SerialNo:      sn,
		SourceContent: ctx.SourceContent,
	}
	_, err = models.InsertRewardOperateRecord(record)
	if err != nil {
		log.Error("createPeriodic InsertRewardOperateRecord error. %v", err)
		return "", err
	}
	return record.SerialNo, nil
}

func UpdateRewardRecordToFinalStatus(sourceType, requestId, newStatus string) error {
	_, err := models.UpdateRewardRecordToFinalStatus(sourceType, requestId, newStatus)
	if err != nil {
		log.Error("UpdateRewardRecord UpdateRewardRecordToFinalStatus error. %v", err)
		return err
	}
	return nil
}

func GetPeriodicTask(sourceType models.SourceType, sourceId, requestId string, operateType models.RewardOperateType) (*models.RewardPeriodicTask, error) {
	_, err := models.GetPointOperateRecordBySourceTypeAndRequestId(sourceType.Name(), requestId, operateType.Name())
	if err == nil {
		task, err := models.GetPeriodicTaskBySourceIdAndType(sourceType, sourceId, operateType)
		if err != nil {
			log.Error("GetPeriodicTaskBySourceIdAndType error,%v", err)
			return nil, err
		}
		return task, nil
	}

	if err != nil && !models.IsErrRecordNotExist(err) {
		log.Error("GetPointOperateRecordBySourceTypeAndRequestId error,%v", err)
		return nil, err
	}
	return nil, nil
}

func StartPeriodicTask(opts *models.StartPeriodicTaskOpts) (*models.RewardPeriodicTask, error) {
	//add lock
	var rewardLock = redis_lock.NewDistributeLock(redis_key.RewardOperateLock(opts.RequestId, opts.SourceType.Name(), opts.OperateType.Name()))
	isOk, err := rewardLock.Lock(3 * time.Second)
	if !isOk {
		log.Info("duplicated operate request,targetUserId=%d requestId=%s", opts.TargetUserId, opts.RequestId)
		return nil, nil
	}
	defer rewardLock.UnLock()

	r, err := GetPeriodicTask(opts.SourceType, opts.SourceId, opts.RequestId, opts.OperateType)
	if err != nil {
		return nil, err
	}

	if r != nil {
		return r, nil
	}

	//new reward operate record
	recordId, err := createPeriodicRewardOperateRecord(opts)
	if err != nil {
		log.Error("StartAndGetPeriodicTask createPeriodicRewardOperateRecord error. %v", err)
		return nil, err
	}

	if err = NewRewardPeriodicTask(recordId, opts); err != nil {
		log.Error("StartAndGetPeriodicTask NewRewardPeriodicTask error. %v", err)
		UpdateRewardRecordToFinalStatus(opts.SourceType.Name(), opts.RequestId, models.OperateStatusFailed)
		return nil, err
	}

	task, err := models.GetPeriodicTaskBySourceIdAndType(opts.SourceType, opts.SourceId, opts.OperateType)
	if err != nil {
		log.Error("GetPeriodicTaskBySourceIdAndType error,%v", err)
		return nil, err
	}
	return task, nil
}

func StopPeriodicTaskAsyn(sourceType models.SourceType, sourceId string, operateType models.RewardOperateType, sourcContent string) {
	go StopPeriodicTask(sourceType, sourceId, operateType, sourcContent)
}

func StopPeriodicTask(sourceType models.SourceType, sourceId string, operateType models.RewardOperateType, sourceContent string) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	task, err := models.GetPeriodicTaskBySourceIdAndType(sourceType, sourceId, operateType)
	if err != nil {
		log.Error("StopPeriodicTask. GetPeriodicTaskBySourceIdAndType error. %v", err)
		return err
	}
	if task == nil {
		log.Info("Periodic task is not exist")
		return nil
	}
	if task.Status == models.PeriodicTaskStatusFinished {
		log.Info("Periodic task is finished")
		return nil
	}
	now := time.Now()
	RunRewardTask(*task, now)
	return models.StopPeriodicTask(task.ID, task.OperateSerialNo, now, sourceContent)
}

func generateOperateSerialNo() (string, error) {
	s, err := GetSerialNoByRedis()
	if err != nil {
		log.Error("generateOperateSerialNo error. %v", err)

		return "", err
	}
	return s, nil
}
