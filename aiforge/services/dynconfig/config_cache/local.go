package config_cache

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

type LocalCache struct {
	cache *cache.Cache
}

func NewLocalCache(defaultExpiration, cleanupInterval time.Duration) *LocalCache {
	return &LocalCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (l *LocalCache) Get(key string) (interface{}, error) {
	v, success := l.cache.Get(key)
	if success {
		if v == nil {
			return nil, nil
		}

		return v, nil
	}
	return nil, errors.New("key not found")
}

func (l *LocalCache) Set(key string, value interface{}, ttl time.Duration) error {
	l.cache.Set(key, value, ttl)
	return nil
}
