package task

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/services/reward/limiter"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//GetTaskConfig get task config from redis cache first
// if not exist in redis, find in db and refresh the redis key
func GetTaskConfig(taskType string) (*models.TaskConfig, error) {
	list, err := GetTaskConfigList()
	if err != nil {
		log.Error(" GetTaskConfigList error. %v", err)
		return nil, err
	}
	for _, v := range list {
		if v.TaskCode == taskType {
			return v, nil
		}
	}
	return nil, nil
}

func GetTaskConfigList() ([]*models.TaskConfig, error) {
	redisKey := redis_key.TaskConfigList()
	configStr, _ := redis_client.Get(redisKey)
	if configStr != "" {
		if configStr == redis_key.EMPTY_REDIS_VAL {
			return nil, nil
		}
		config := make([]*models.TaskConfig, 0)
		json.Unmarshal([]byte(configStr), &config)
		return config, nil
	}
	config, err := models.GetTaskConfigList()
	if err != nil {
		log.Error(" GetTaskConfigList from model error. %v", err)
		if models.IsErrRecordNotExist(err) {
			redis_client.Setex(redisKey, redis_key.EMPTY_REDIS_VAL, 5*time.Second)
			return nil, nil
		}
		return nil, err
	}
	jsonStr, _ := json.Marshal(config)
	redis_client.Setex(redisKey, string(jsonStr), 30*24*time.Hour)
	return config, nil
}

func GetTaskConfigPageWithDeleted(opt models.GetTaskConfigOpts) ([]*models.TaskAndLimiterConfig, int64, error) {
	config, count, err := models.GetTaskConfigPageWithDeleted(opt)
	if err != nil {
		log.Error(" GetTaskConfigPageWithDeleted from model error. %v", err)
		if models.IsErrRecordNotExist(err) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return config, count, nil
}

func GetTaskConfigWithLimitList(opt models.GetTaskConfigOpts) (*models.TaskConfigWithLimitResponse, error) {
	list, n, err := GetTaskConfigPageWithDeleted(opt)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	r := make([]*models.TaskConfigWithSingleLimit, 0)
	for i := 0; i < len(list); i++ {
		li := list[i]
		t := &models.TaskConfigWithSingleLimit{
			ID:          li.TaskConfig.ID,
			TaskCode:    li.TaskConfig.TaskCode,
			AwardType:   li.TaskConfig.AwardType,
			AwardAmount: li.TaskConfig.AwardAmount,
			Creator:     li.TaskConfig.CreatorName,
			CreatedUnix: li.TaskConfig.CreatedUnix,
			IsDeleted:   li.TaskConfig.DeletedAt > 0,
			DeleteAt:    li.TaskConfig.DeletedAt,
			LimitNum:    li.LimitConfig.LimitNum,
			RefreshRate: li.LimitConfig.RefreshRate,
		}
		r = append(r, t)
	}

	return &models.TaskConfigWithLimitResponse{
		Records:  r,
		Page:     opt.Page,
		PageSize: opt.PageSize,
		Total:    n,
	}, nil
}

func AddTaskConfig(config models.TaskConfigWithLimit, doer *models.User) error {
	if config.TaskCode == "" || config.AwardType == "" {
		log.Error(" EditTaskConfig param error")
		return errors.New("param error")
	}
	var lock = redis_lock.NewDistributeLock(redis_key.TaskConfigOperateLock(config.TaskCode, config.AwardType))
	isOk, _ := lock.LockWithWait(3*time.Second, 3*time.Second)
	if !isOk {
		return errors.New("Get lock failed")
	}
	defer lock.UnLock()

	t, err := models.GetTaskConfigByTaskCode(config.TaskCode)
	if err != nil && !models.IsErrRecordNotExist(err) {
		return err
	}
	if t != nil {
		return errors.New("task config is exist")
	}

	for i, l := range config.Limiters {
		if l.Scope == "" {
			config.Limiters[i].Scope = models.LimitScopeSingleUser.Name()
		}
	}
	err = models.NewTaskConfig(config, doer)
	if err != nil {
		log.Error("add task config error,config:%v  err:%v", config, err)
		return err
	}
	redis_client.Del(redis_key.LimitConfig(models.LimitTypeTask.Name()))
	redis_client.Del(redis_key.TaskConfigList())
	return nil
}

func EditTaskConfig(config models.TaskConfigWithLimit, doer *models.User) error {
	if config.TaskCode == "" || config.AwardType == "" || config.ID <= 0 {
		log.Error(" EditTaskConfig param error")
		return errors.New("param error")
	}
	var lock = redis_lock.NewDistributeLock(redis_key.TaskConfigOperateLock(config.TaskCode, config.AwardType))
	isOk, _ := lock.LockWithWait(3*time.Second, 3*time.Second)
	if !isOk {
		return errors.New("Get lock failed")
	}
	defer lock.UnLock()
	t, err := models.GetTaskConfigByID(config.ID)
	if err != nil {
		return err
	}
	if t == nil {
		return errors.New("task config is not exist")
	}

	for i, l := range config.Limiters {
		if l.Scope == "" {
			config.Limiters[i].Scope = models.LimitScopeSingleUser.Name()
		}
	}
	err = models.EditTaskConfig(config, doer)
	if err != nil {
		log.Error("add task config error,config:%v  err:%v", config, err)
		return err
	}
	redis_client.Del(redis_key.LimitConfig(models.LimitTypeTask.Name()))
	redis_client.Del(redis_key.TaskConfigList())
	return nil
}

func DelTaskConfig(id int64, doer *models.User) error {
	if id == 0 {
		log.Error(" EditTaskConfig param error")
		return errors.New("param error")
	}
	err := models.DelTaskConfig(id, doer)
	if err != nil {
		log.Error("del task config error,err:%v", err)
		return err
	}
	redis_client.Del(redis_key.LimitConfig(models.LimitTypeTask.Name()))
	redis_client.Del(redis_key.TaskConfigList())
	return nil
}

func GetPointRule() (*models.PointRule, error) {
	r, err := limiter.GetSingleDailyPointLimitConfig()
	if err != nil {
		return nil, err
	}
	limiters, err := limiter.GetLimitersByLimitType(models.LimitTypeTask)
	if err != nil {
		return nil, err
	}
	limiterMap := make(map[string]*models.LimitConfig, 0)
	for i := 0; i < len(limiters); i++ {
		limiterMap[limiters[i].LimitCode] = &limiters[i]
	}

	taskConfigs, err := GetTaskConfigList()
	if err != nil {
		return nil, err
	}
	taskRules := make([]models.TaskRule, len(taskConfigs))

	for i, taskConfig := range taskConfigs {
		rule := models.TaskRule{
			TaskCode:    taskConfig.TaskCode,
			AwardType:   taskConfig.AwardType,
			AwardAmount: taskConfig.AwardAmount,
		}
		limiter := limiterMap[fmt.Sprint(taskConfig.TaskCode)]
		if limiter != nil {
			rule.RefreshRate = limiter.RefreshRate
			rule.LimitNum = limiter.LimitNum
		}
		taskRules[i] = rule
	}

	pointRule := &models.PointRule{
		TaskRules: taskRules,
	}
	if r != nil {
		pointRule.UserDailyLimit = r.LimitNum
	}
	return pointRule, nil
}
