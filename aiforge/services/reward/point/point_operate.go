package point

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/services/reward/limiter"
	"code.gitea.io/gitea/services/reward/point/account"
	"errors"
	"time"
)

type PointOperator struct {
}

func (operator *PointOperator) IsLimited(ctx *models.RewardOperateContext) error {
	realAmount, err := limiter.CheckLimit(ctx.SourceType.Name(), models.LimitTypeRewardPoint, ctx.TargetUserId, ctx.Reward.Amount, ctx.RejectPolicy)
	if err != nil {
		log.Error("PointOperator IsLimited error,err=%v", err)
		return err
	}
	if realAmount < ctx.Reward.Amount {
		ctx.LossAmount = ctx.Reward.Amount - realAmount
		ctx.Reward.Amount = realAmount
	}
	return nil
}

func (operator *PointOperator) Operate(ctx *models.RewardOperateContext) error {
	lock := redis_lock.NewDistributeLock(redis_key.PointAccountOperateLock(ctx.TargetUserId))
	isOk, err := lock.LockWithWait(3*time.Second, 3*time.Second)
	if err != nil {
		log.Error("Get PointAccountOperateLock error,err=%v", err)
		return err
	}
	if isOk {
		defer lock.UnLock()
		na, err := account.GetAccount(ctx.TargetUserId)
		if err != nil || na == nil {
			log.Error("operator get account error error,err=%v", err)
			return errors.New("get account error")
		}
		if ctx.OperateType == models.OperateTypeIncrease {
			err = na.Increase(ctx.Reward.Amount, ctx.SourceId)
		} else if ctx.OperateType == models.OperateTypeDecrease {
			if !ctx.PermittedNegative && na.Balance < ctx.Reward.Amount {
				log.Info("account balance is not enough,ctx=%v", ctx)
				return models.ErrInsufficientPointsBalance{}
			}
			err = na.Decrease(ctx.Reward.Amount, ctx.SourceId)
		}
		if err != nil {
			log.Error("operate account balance error,err=%v", err)
			return err
		}
		redis_client.Del(redis_key.PointAccountInfo(ctx.TargetUserId))

	} else {
		log.Error("Get account operate lock failed,ctx=%v", ctx)
		return errors.New("Get account operate lock failed")
	}
	return nil
}
