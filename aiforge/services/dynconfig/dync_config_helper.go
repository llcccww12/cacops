package dynconfig

import (
	"errors"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/dynconfig/config_cache"
	"code.gitea.io/gitea/services/dynconfig/dync_parser"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
)

const (
	CacheNil = "nil"
)

type DyncConfigHelper struct {
	localCache config_cache.Cache
	fetcher    fetcher.Fetcher
}

func NewDyncConfigHelper(local config_cache.Cache, fetcher fetcher.Fetcher) *DyncConfigHelper {
	return &DyncConfigHelper{local, fetcher}
}

func (s *DyncConfigHelper) GetConfig(key string, noCache bool) (interface{}, error) {
	if noCache {
		return s.getConfigWithNoCache(key)
	}
	return s.getConfigWithCache(key)
}

func (s *DyncConfigHelper) getConfigWithCache(key string) (interface{}, error) {
	data, err := s.localCache.Get(key)
	if err == nil {
		if data == CacheNil {
			log.Info("GetConfig from local cache nil, key=%s", key)
			return nil, nil
		}
		return data, nil
	}
	lockKey := redis_key.DyncConfigLockKey(key)
	lock := redis_lock.NewDistributeLock(lockKey)
	success, err := lock.LockWithWait(5*time.Second, 5*time.Second)
	if err != nil {
		log.Error("GetConfig lock failed, error=%v", err)
		return nil, err
	}
	if !success {
		log.Error("GetConfig lock failed, key=%s", lockKey)
		return nil, errors.New("GetConfig lock failed")
	}
	defer lock.UnLock()

	data, err = s.localCache.Get(key)
	if err == nil {
		log.Info("GetConfig from local cache success, key=%s", key)
		return data, nil
	}

	fetechd, err := s.fetcher.Fetch(key)
	if err != nil || fetechd == nil {
		log.Error("FetchConfig failed, error=%v", err)
		if models.IsErrNotExist(err) {
			//如果没有找到配置文件，设置空标记
			s.localCache.Set(key, CacheNil, setting.DyncConfigDefaultTTL)
			return nil, nil
		}
		return nil, err
	}
	parsed, err := dync_parser.ParseConfig(*fetechd)
	if err != nil {
		log.Error("ParseConfig failed, error=%v", err)
		return nil, err
	}
	s.localCache.Set(key, parsed, setting.DyncConfigDefaultTTL)

	return parsed, nil

}

func (s *DyncConfigHelper) getConfigWithNoCache(key string) (interface{}, error) {

	fetechd, err := s.fetcher.Fetch(key)
	if err != nil || fetechd == nil {
		log.Error("FetchConfig failed, error=%v", err)
		if models.IsErrNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	parsed, err := dync_parser.ParseConfig(*fetechd)
	if err != nil {
		log.Error("ParseConfig failed, error=%v", err)
		return nil, err
	}

	return parsed, nil

}
