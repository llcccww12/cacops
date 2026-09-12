package reward

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"fmt"
	"math/rand"
	"time"
)

func GetSerialNoByRedis() (string, error) {
	now := time.Now()
	r := int64(rand.Intn(3)) + 1
	n, err := redis_client.IncrBy(redis_key.RewardSerialCounter(now), r)
	if err != nil {
		log.Error("GetSerialNoByRedis RewardSerialCounter error. %v", err)
		return "", err
	}
	if n == r {
		redis_client.Expire(redis_key.RewardSerialCounter(now), 2*time.Minute)
	}
	//when the counter n exceeds 1000, the length of the serial number will become longer
	if n >= 1000 {
		return now.Format("200601021504") + fmt.Sprintf("%d", n) + fmt.Sprint(rand.Intn(10)), nil
	}
	return now.Format("200601021504") + fmt.Sprintf("%03d", n) + fmt.Sprint(rand.Intn(10)), nil
}
