package redis_key

import "time"

const SERIAL_REDIS_PREFIX = "serial"

func RewardSerialCounter(now time.Time) string {
	h := now.Format("200601021504")
	return KeyJoin(SERIAL_REDIS_PREFIX, "reward_operate", h, "counter")
}
