package phone

import (
	"fmt"
	"time"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/modules/phone"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"

	"github.com/gomodule/redigo/redis"
)

//验证码存储前缀 使用时%s用手机号替代
const CODE_PREFIX = "P_C:%s"

//手机号发送验证码次数Hkey，%s对应日期， 存储在hset中，值是hashet,记录手机号和发送次数
const TIMES_PREFIX = "P_T:%s"

func GetPhoneNumberSendTimes(conn redis.Conn, phoneNumber string) (int, error) {
	i, err := redis_client.HGET(conn, GetPhoneTimesHKey(), phoneNumber)

	return redis.Int(i, err)
}

func GetPhoneCodeTTL(conn redis.Conn, phoneNumber string) (int, error) {
	return redis_client.Ttl(conn, GetPhoneCodeKey(phoneNumber))

}
func SendVerifyCode(conn redis.Conn, phoneNumber string) error {
	timesKey := GetPhoneTimesHKey()
	exists, err := redis_client.EXISTS(conn, timesKey)
	if err != nil {
		return err
	}
	code := phone.GenerateVerifyCode(setting.PhoneService.VerifyCodeLength)
	err = phone.SendVerifyCode(phoneNumber, code)
	if err != nil {
		return err
	}
	redis_client.SET(conn, GetPhoneCodeKey(phoneNumber), code, setting.PhoneService.CodeTimeout)
	if !exists {
		err = redis_client.HSETNX(conn, timesKey, phoneNumber, 1)
		if err != nil {
			return err
		}
		err = redis_client.EXPIRE(conn, timesKey, getRemainSecondOfDay(time.Now()))
		if err != nil {
			return err
		}

	} else {
		err = redis_client.HINCRBY(conn, timesKey, phoneNumber, 1)

		if err != nil {
			return err
		}

	}
	return nil

}

func IsVerifyCodeRight(phoneNumer string, verifyCode string) bool {
	if phoneNumer == "" {
		return false
	}
	value, err := redis_client.Get(GetPhoneCodeKey(phoneNumer))
	if err != nil {
		log.Warn("redis err", err)
		return false
	} else {
		if value == "" {
			return false
		}
		return value == verifyCode
	}
}

func GetPhoneCodeKey(phoneNumber string) string {
	return fmt.Sprintf(CODE_PREFIX, phoneNumber)
}

func GetPhoneTimesHKey() string {
	today := time.Now().Format("2006-01-02")
	return fmt.Sprintf(TIMES_PREFIX, today)
}

func getRemainSecondOfDay(t time.Time) int {
	return 86400 - 60*60*t.Hour() - 60*t.Minute() - t.Second()
}
