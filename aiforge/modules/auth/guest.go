package auth

import (
	"time"

	"code.gitea.io/gitea/modules/labelmsg"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"
	"github.com/gomodule/redigo/redis"
)

type Guest struct {
	Enabled         bool
	WhiteExpiration int // 白名单有效期
	BlackExpiration int // 黑名单有效期
	TempExpiration  int // 临时名单有效期
	WhiteKeyPrefix  string
	BlackKeyPrefix  string
	TempKeyPrefix   string
	MaxFailedKey    string
	MaxFailedCount  int
}

var guest *Guest

func (g *Guest) AddWhiteList(ip string) {

	redis_client.Setex(g.WhiteKeyPrefix+ip, "", time.Second*time.Duration(g.WhiteExpiration))

}

func (g *Guest) AddBlackList(ip string) {

	redis_client.Setex(g.BlackKeyPrefix+ip, "", time.Second*time.Duration(g.BlackExpiration))

}

func (g *Guest) AddTempList(ip string) {

	redis_client.Setex(g.TempKeyPrefix+ip, "", time.Second*time.Duration(g.TempExpiration))

}
func (g *Guest) AddMaxFailedCount(ip string) {
	conn := labelmsg.Get()
	defer conn.Close()
	exists, err := redis_client.EXISTS(conn, g.MaxFailedKey)
	if err != nil {
		return
	}
	if !exists {
		redis_client.HSETNX(conn, g.MaxFailedKey, ip, 1)

	} else {
		redis_client.HINCRBY(conn, g.MaxFailedKey, ip, 1)
	}

}

func (g *Guest) ResetMaxFailedCount(ip string) {
	conn := labelmsg.Get()
	defer conn.Close()
	redis_client.HDEL(conn, g.MaxFailedKey, ip)
}

func (g *Guest) GetMaxFailedCount(ip string) int {
	conn := labelmsg.Get()
	defer conn.Close()
	i, err := redis_client.HGET(conn, g.MaxFailedKey, ip)
	count, _ := redis.Int(i, err)
	return count
}

func (g *Guest) IsInWhiteList(ip string) bool {
	redisConn := labelmsg.Get()
	defer redisConn.Close()
	has, _ := redis_client.EXISTS(redisConn, g.WhiteKeyPrefix+ip)
	return has
}

func (g *Guest) IsInBlackList(ip string) bool {
	redisConn := labelmsg.Get()
	defer redisConn.Close()
	has, _ := redis_client.EXISTS(redisConn, g.BlackKeyPrefix+ip)
	return has
}

func (g *Guest) IsInTempList(ip string) bool {
	redisConn := labelmsg.Get()
	defer redisConn.Close()
	has, _ := redis_client.EXISTS(redisConn, g.TempKeyPrefix+ip)
	return has
}

func GetGuest() *Guest {

	if guest == nil {
		guest = &Guest{
			Enabled:         setting.GuestInfo.Enabled,
			WhiteExpiration: setting.GuestInfo.WhiteExpiration,
			BlackExpiration: setting.GuestInfo.BlackExpiration,
			TempExpiration:  setting.GuestInfo.TempExpiration,
			WhiteKeyPrefix:  "guest:white:",
			BlackKeyPrefix:  "guest:black:",
			TempKeyPrefix:   "guest:temp:",
			MaxFailedKey:    "guest:maxfailed:",
			MaxFailedCount:  setting.GuestInfo.MaxFailedCount,
		}
	}
	return guest

}
