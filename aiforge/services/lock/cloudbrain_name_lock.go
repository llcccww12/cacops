package lock

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
)

type CloudbrainDisplayJobNameLock struct {
}

func (c CloudbrainDisplayJobNameLock) IsMatch(ctx *LockContext) bool {
	return true
}

func (c CloudbrainDisplayJobNameLock) Lock(ctx *LockContext) string {
	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainBindingJobNameKey(ctx.Task.JobType, ctx.Task.DisplayJobName))
	isOk, err := lock.Lock(models.CloudbrainKeyDuration)
	if !isOk {
		log.Error("CloudbrainDisplayJobNameLock lock failed:%v", err)
		return "repo.cloudbrain_samejob_err"
	}
	return ""
}

func (c CloudbrainDisplayJobNameLock) Unlock(ctx *LockContext) error {
	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainBindingJobNameKey(ctx.Task.JobType, ctx.Task.DisplayJobName))
	return lock.UnLock()
}
