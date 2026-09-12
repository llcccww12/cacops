package wechat

import (
	"time"

	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
)

var accessTokenLock = redis_lock.NewDistributeLock(redis_key.AccessTokenLockKey())

func GetWechatAccessToken() string {
	token, _ := redis_client.Get(redis_key.WechatAccessTokenKey())
	if token != "" {
		if token == redis_key.EMPTY_REDIS_VAL {
			return ""
		}
		live, _ := redis_client.TTL(redis_key.WechatAccessTokenKey())
		//refresh wechat access token when expire time less than 5 minutes
		if live > 0 && live < 300 {
			refreshAccessToken()
		}
		return token
	}
	return refreshAndGetAccessToken()
}

func refreshAccessToken() {
	if ok, _ := accessTokenLock.Lock(3 * time.Second); ok {
		defer accessTokenLock.UnLock()
		callAccessTokenAndUpdateCache()
	}
}

func refreshAndGetAccessToken() string {
	isOk, err := accessTokenLock.LockWithWait(3*time.Second, 3*time.Second)
	if err != nil {
		return ""
	}
	if isOk {
		defer accessTokenLock.UnLock()
		token, _ := redis_client.Get(redis_key.WechatAccessTokenKey())
		if token != "" {
			if token == redis_key.EMPTY_REDIS_VAL {
				return ""
			}
			return token
		}
		return callAccessTokenAndUpdateCache()
	}
	return ""

}

func callAccessTokenAndUpdateCache() string {
	r := callAccessToken()

	var token string
	if r != nil {
		token = r.Access_token
	}

	if token == "" {
		redis_client.Setex(redis_key.WechatAccessTokenKey(), redis_key.EMPTY_REDIS_VAL, 10*time.Second)
		return ""
	}
	redis_client.Setex(redis_key.WechatAccessTokenKey(), token, time.Duration(r.Expires_in)*time.Second)
	return token
}
