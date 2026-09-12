package redis_lock

import (
	"time"

	"code.gitea.io/gitea/modules/redis/redis_client"
)

type DistributeLock struct {
	lockKey string
}

func NewDistributeLock(lockKey string) *DistributeLock {
	return &DistributeLock{lockKey: lockKey}
}

func (lock *DistributeLock) Lock(expireTime time.Duration) (bool, error) {
	isOk, err := redis_client.Setnx(lock.lockKey, "", expireTime)
	if err != nil {
		return false, err
	}
	return isOk, nil
}

func (lock *DistributeLock) LockWithWait(expireTime time.Duration, waitTime time.Duration) (bool, error) {
	start := time.Now().Unix() * 1000
	duration := waitTime.Milliseconds()
	for {
		isOk, err := redis_client.Setnx(lock.lockKey, "", expireTime)
		if err != nil {
			return false, err
		}
		if isOk {
			return true, nil
		}
		if time.Now().Unix()*1000-start > duration {
			return false, nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return false, nil
}

func (lock *DistributeLock) UnLock() error {
	_, err := redis_client.Del(lock.lockKey)
	return err
}
