package limiter

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/services/task/period"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type limiterRunner struct {
	limiters      []models.LimitConfig
	index         int
	userId        int64
	amount        float64
	limitCode     string
	limitType     models.LimitType
	rejectPolicy  models.LimiterRejectPolicy
	resultMap     map[int]limitResult
	minRealAmount float64
}

type limitResult struct {
	isLoss     bool
	planAmount float64
	realAmount float64
}

func newLimitResult(isLoss bool, planAmount, realAmount float64) limitResult {
	return limitResult{
		isLoss:     isLoss,
		planAmount: planAmount,
		realAmount: realAmount,
	}
}

func newLimiterRunner(limitCode string, limitType models.LimitType, userId int64, amount float64, policy models.LimiterRejectPolicy) *limiterRunner {
	return &limiterRunner{
		userId:       userId,
		amount:       amount,
		limitCode:    limitCode,
		limitType:    limitType,
		index:        0,
		rejectPolicy: policy,
		resultMap:    make(map[int]limitResult, 0),
	}
}

//Run run all limiters
//return real used amount(when choose the FillUp reject policy, amount may only be partially used)
func (l *limiterRunner) Run() error {
	if err := l.LoadLimiters(); err != nil {
		return err
	}
	l.minRealAmount = l.amount
	for l.index < len(l.limiters) {
		err := l.limit(l.limiters[l.index])
		if err != nil {
			log.Info("limiter check failed,%v", err)
			l.Rollback()
			return err
		}
		result := l.resultMap[l.index]
		if result.isLoss {
			//find the minimum real amount
			if l.minRealAmount > result.realAmount {
				l.minRealAmount = result.realAmount
			}
		}
		l.index += 1
	}

	//post process
	l.PostProcess()
	return nil
}

//Rollback rollback the usedNum  from limiters[0] to limiters[index]
func (l *limiterRunner) Rollback() error {
	for i := l.index - 1; i >= 0; i-- {
		l.rollback(l.limiters[i], l.resultMap[i])
	}
	return nil
}

func (l *limiterRunner) rollback(r models.LimitConfig, result limitResult) error {
	p, err := period.GetPeriod(r.RefreshRate)
	if err != nil {
		return err
	}
	redisKey := redis_key.LimitCount(l.userId, r.LimitCode, r.LimitType, r.Scope, p)
	redis_client.IncrByFloat(redisKey, -1*result.realAmount)
	return nil
}

//PostProcess process loss,if realAmount < planAmount
func (l *limiterRunner) PostProcess() error {
	for i := l.index - 1; i >= 0; i-- {
		l.postProcess(l.limiters[i], l.resultMap[i])
	}
	return nil
}

func (l *limiterRunner) postProcess(r models.LimitConfig, result limitResult) error {
	if result.realAmount == l.minRealAmount {
		return nil
	}
	p, err := period.GetPeriod(r.RefreshRate)
	if err != nil {
		return err
	}
	diff := result.realAmount - l.minRealAmount
	redisKey := redis_key.LimitCount(l.userId, r.LimitCode, r.LimitType, r.Scope, p)
	redis_client.IncrByFloat(redisKey, -1*diff)
	return nil
}

func (l *limiterRunner) limit(r models.LimitConfig) error {
	p, err := period.GetPeriod(r.RefreshRate)
	if err != nil {
		return err
	}
	redisKey := redis_key.LimitCount(l.userId, r.LimitCode, r.LimitType, r.Scope, p)
	usedNum, err := redis_client.IncrByFloat(redisKey, l.amount)
	if err != nil {
		return err
	}
	//if usedNum equals  amount,it is the first operation in period or redis cache deleted
	//count in database to distinguish the two cases
	if usedNum == l.amount {
		n, err := l.countInPeriod(r, p)
		if err != nil {
			return err
		}
		if n > 0 {
			//means redis cache deleted,incr the cache with real value
			usedNum, err = redis_client.IncrByFloat(redisKey, n)
		}
		if p != nil {
			redis_client.Expire(redisKey, p.LeftTime)
		} else {
			//add default expire time if no period set
			redis_client.Expire(redisKey, 24*time.Hour)
		}
	}
	if usedNum > r.LimitNum {
		if usedNum-r.LimitNum >= l.amount {
			redis_client.IncrByFloat(redisKey, -1*l.amount)
			return errors.New(fmt.Sprintf("over limit,congfigId=%d", r.ID))
		}
		switch l.rejectPolicy {
		case models.FillUp:
			exceed := usedNum - r.LimitNum
			realAmount := l.amount - exceed
			redis_client.IncrByFloat(redisKey, -1*exceed)
			l.resultMap[l.index] = newLimitResult(true, l.amount, realAmount)
			return nil
		case models.JustReject:
			redis_client.IncrByFloat(redisKey, -1*l.amount)
			return errors.New(fmt.Sprintf("over limit,congfigId=%d", r.ID))
		case models.PermittedOnce:
			l.resultMap[l.index] = newLimitResult(false, l.amount, l.amount)
			return nil
		}

	}
	l.resultMap[l.index] = newLimitResult(false, l.amount, l.amount)
	return nil
}

func (l *limiterRunner) LoadLimiters() error {
	limiters, err := GetLimiters(l.limitCode, l.limitType)
	if err != nil {
		return err
	}
	if limiters != nil {
		l.limiters = limiters
	}
	return nil
}

func (l *limiterRunner) countInPeriod(r models.LimitConfig, p *models.PeriodResult) (float64, error) {
	switch r.LimitType {
	case models.LimitTypeTask.Name():
		return models.CountTaskAccomplishLogInTaskPeriod(r.LimitCode, l.userId, p)
	case models.LimitTypeRewardPoint.Name():
		return models.SumRewardAmountInTaskPeriod(models.RewardTypePoint.Name(), r.LimitCode, l.userId, p)
	default:
		return 0, nil

	}
}

func CheckLimit(limitCode string, limitType models.LimitType, userId int64, amount float64, rejectPolicy models.LimiterRejectPolicy) (float64, error) {
	if rejectPolicy == "" {
		rejectPolicy = models.JustReject
	}
	r := newLimiterRunner(limitCode, limitType, userId, amount, rejectPolicy)
	err := r.Run()
	if err != nil {
		return 0, err
	}
	return r.minRealAmount, nil
}

func GetLimiters(limitCode string, limitType models.LimitType) ([]models.LimitConfig, error) {
	limiters, err := GetLimitersByLimitType(limitType)
	if err != nil {
		return nil, err
	}
	result := make([]models.LimitConfig, 0)
	for i, v := range limiters {
		if v.LimitCode == "" || v.LimitCode == limitCode {
			result = append(result, limiters[i])
		}
	}
	return result, nil
}

func GetLimitersByLimitType(limitType models.LimitType) ([]models.LimitConfig, error) {
	redisKey := redis_key.LimitConfig(limitType.Name())
	val, _ := redis_client.Get(redisKey)
	if val != "" {
		if val == redis_key.EMPTY_REDIS_VAL {
			return nil, nil
		}
		limiters := make([]models.LimitConfig, 0)
		json.Unmarshal([]byte(val), &limiters)
		return limiters, nil
	}
	limiters, err := models.GetLimitConfigByLimitType(limitType)
	if err != nil {
		if models.IsErrRecordNotExist(err) {
			redis_client.Setex(redisKey, redis_key.EMPTY_REDIS_VAL, 5*time.Second)
			return nil, nil
		}
		return nil, err
	}
	jsonStr, _ := json.Marshal(limiters)
	redis_client.Setex(redisKey, string(jsonStr), 30*24*time.Hour)

	return limiters, nil
}

func GetLimitersByRelatedIdWithDeleted(limitType models.LimitType) ([]models.LimitConfig, error) {
	limiters, err := models.GetLimitersByRelatedIdWithDeleted(limitType)
	if err != nil {
		if models.IsErrRecordNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return limiters, nil
}
