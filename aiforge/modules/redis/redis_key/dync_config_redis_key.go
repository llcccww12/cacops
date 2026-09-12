package redis_key

const DYNC_CONFIG_REDIS_PREFIX = "dync_config"

func DyncConfigLockKey(key string) string {
	return KeyJoin(DYNC_CONFIG_REDIS_PREFIX, key, "update", "lock")
}
