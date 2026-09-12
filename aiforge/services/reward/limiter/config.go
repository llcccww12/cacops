package limiter

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
)

func GetSingleDailyPointLimitConfig() (*models.LimitConfigVO, error) {
	r, err := GetLimitConfigList(models.LimitConfigQueryOpts{
		RefreshRate: models.PeriodDaily,
		Scope:       models.LimitScopeSingleUser,
		LimitCode:   models.SourceTypeAccomplishTask.Name(),
		LimitType:   models.LimitTypeRewardPoint,
	})
	if err != nil {
		return nil, err
	}
	if r == nil || len(r) == 0 {
		return nil, nil
	}
	return r[0], nil
}

func SetSingleDailyPointLimitConfig(limitNum float64, doer *models.User) error {
	l := &models.LimitConfigVO{
		RefreshRate: models.PeriodDaily,
		Scope:       models.LimitScopeSingleUser.Name(),
		LimitCode:   models.SourceTypeAccomplishTask.Name(),
		LimitType:   models.LimitTypeRewardPoint.Name(),
		LimitNum:    limitNum,
	}
	return AddLimitConfig(l, doer)
}

func GetLimitConfigList(opts models.LimitConfigQueryOpts) ([]*models.LimitConfigVO, error) {
	r, err := GetLimitersByLimitType(opts.LimitType)
	if err != nil {
		log.Error("GetLimitConfigList error when getting limiters by limit type.err=%v", err)
		return nil, err
	}
	result := make([]*models.LimitConfigVO, 0)
	for _, v := range r {
		if opts.LimitCode != "" && opts.LimitCode != v.LimitCode {
			continue
		}
		if opts.Scope != "" && opts.Scope.Name() != v.Scope {
			continue
		}
		if opts.RefreshRate != "" && opts.RefreshRate != v.RefreshRate {
			continue
		}
		if opts.LimitType != "" && opts.LimitType.Name() != v.LimitType {
			continue
		}
		result = append(result, v.ToLimitConfigVO())
	}
	return result, nil
}
func GetLimitConfigById(id int64) (*models.LimitConfig, error) {
	return models.GetLimitConfigById(id)
}

func AddLimitConfig(config *models.LimitConfigVO, doer *models.User) error {
	r := &models.LimitConfig{
		Title:       config.Title,
		RefreshRate: config.RefreshRate,
		Scope:       config.Scope,
		LimitNum:    config.LimitNum,
		LimitCode:   config.LimitCode,
		LimitType:   config.LimitType,
		CreatorId:   doer.ID,
		CreatorName: doer.Name,
	}
	err := models.AddLimitConfig(r)

	if err != nil {
		log.Error("add limit config error,config:%v  err:%v", config, err)
		return err
	}
	redis_client.Del(redis_key.LimitConfig(config.LimitType))
	return nil
}

func DeleteLimitConfig(id int64, doer *models.User) error {
	config, err := GetLimitConfigById(id)
	if err != nil {
		log.Error("GetLimitConfigById err,e=%v", err)
		return err
	}
	err = models.DeleteLimitConfig(*config, doer.ID, doer.Name)

	if err != nil {
		log.Error("add limit config error,config:%v  err:%v", config, err)
		return err
	}
	redis_client.Del(redis_key.LimitConfig(config.LimitType))
	return nil
}
