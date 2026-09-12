package lock

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
)

type CloudbrainUniquenessLock struct {
}

func (c CloudbrainUniquenessLock) IsMatch(ctx *LockContext) bool {
	return true
}

func (c CloudbrainUniquenessLock) Lock(ctx *LockContext) string {
	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainUniquenessKey(ctx.User.ID, ctx.Task.JobType))
	isOk, err := lock.Lock(setting.CloudbrainUniquenessLockTime)
	if !isOk {
		log.Error("CloudbrainDisplayJobNameLock lock failed:%v", err)
		return "repo.cloudbrain.morethanonejob"
	}
	return ""
}

func (c CloudbrainUniquenessLock) Unlock(ctx *LockContext) error {
	lock := redis_lock.NewDistributeLock(redis_key.CloudbrainUniquenessKey(ctx.User.ID, ctx.Task.JobType))
	return lock.UnLock()
}
