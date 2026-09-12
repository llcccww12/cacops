package redis_client

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/labelmsg"
	"github.com/gomodule/redigo/redis"
)

func Setex(key, value string, timeout time.Duration) (bool, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	seconds := int(math.Floor(timeout.Seconds()))
	reply, err := redisClient.Do("SETEX", key, seconds, value)
	if err != nil {
		return false, err
	}
	if reply != "OK" {
		return false, nil
	}
	return true, nil

}

func Setnx(key, value string, timeout time.Duration) (bool, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	seconds := int(math.Floor(timeout.Seconds()))
	reply, err := redisClient.Do("SET", key, value, "NX", "EX", seconds)
	if err != nil {
		return false, err
	}
	if reply != "OK" {
		return false, nil
	}
	return true, nil

}

func Set(key, value string) (bool, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("SET", key, value)
	if err != nil {
		return false, err
	}
	if reply != "OK" {
		return false, nil
	}
	return true, nil

}

func SETNX(conn redis.Conn, key, value string, seconds int) (bool, error) {
	reply, err := conn.Do("SET", key, value, "NX", "EX", seconds)
	return redis.Bool(reply, err)

}

func SET(conn redis.Conn, key, value string, seconds int) (bool, error) {
	reply, err := conn.Do("SETEX", key, seconds, value)
	return redis.Bool(reply, err)

}

func HSETNX(conn redis.Conn, key, subKey string, value interface{}) error {
	_, err := conn.Do("HSETNX", key, subKey, value)
	return err

}

func HGET(conn redis.Conn, key, subKey string) (interface{}, error) {
	return conn.Do("HGET", key, subKey)
}

func HDEL(conn redis.Conn, key string, subKey string) error {
	_, err := conn.Do("HDEL", key, subKey)
	return err

}
func EXISTS(conn redis.Conn, key string) (bool, error) {

	reply, err := conn.Do("EXISTS", key)
	return redis.Bool(reply, err)

}

func HEXISTS(conn redis.Conn, key string, subKey string) (bool, error) {

	reply, err := conn.Do("HEXISTS", key, subKey)
	return redis.Bool(reply, err)

}

func EXPIRE(conn redis.Conn, key string, seconds int) error {
	_, err := conn.Do("EXPIRE", key, seconds)
	return err

}

func HINCRBY(conn redis.Conn, key, subKey string, value int) error {
	_, err := conn.Do("HINCRBY", key, subKey, value)
	return err
}
func GET(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do("GET", key)
}

func Ttl(conn redis.Conn, key string) (int, error) {

	reply, err := conn.Do("TTL", key)
	if err != nil {
		return 0, err
	}
	n, _ := strconv.Atoi(fmt.Sprint(reply))
	return n, nil

}

func Get(key string) (string, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("GET", key)
	if err != nil {
		return "", err
	}
	if reply == nil {
		return "", err
	}
	s, _ := redis.String(reply, nil)
	return s, nil

}

func Del(key string) (int, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("DEL", key)
	if err != nil {
		return 0, err
	}
	if reply == nil {
		return 0, err
	}
	s, _ := redis.Int(reply, nil)
	return s, nil

}

func TTL(key string) (int, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("TTL", key)
	if err != nil {
		return 0, err
	}
	n, _ := strconv.Atoi(fmt.Sprint(reply))
	return n, nil

}

func IncrBy(key string, n int64) (int64, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("INCRBY", key, n)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(fmt.Sprint(reply), 10, 64)
	return i, nil

}

func IncrByFloat(key string, f float64) (float64, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("INCRBYFLOAT", key, f)
	if err != nil {
		return 0, err
	}
	replyStr, ok := reply.([]byte)
	if !ok {
		return 0, errors.New("INCRBYFLOAT response is not correct")
	}
	i, err := strconv.ParseFloat(string(replyStr), 64)
	return i, nil

}

func Expire(key string, expireTime time.Duration) error {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	_, err := redisClient.Do("EXPIRE", key, int64(expireTime.Seconds()))
	if err != nil {
		return err
	}
	return nil

}

// GetInt64 get redis value by Get(key)
// and then parse the value to int64
// return {isExist(bool)} {value(int64)} {error(error)}
func GetInt64(key string) (bool, int64, error) {
	str, err := Get(key)
	if err != nil {
		return false, 0, err
	}
	if str == "" {
		return false, 0, nil
	}

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false, 0, err
	}
	return true, i, nil

}

func ZAdd(key, value string, score float64) error {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	_, err := redisClient.Do("ZADD", key, score, value)
	if err != nil {
		return err
	}
	return nil
}

func ZRem(key, value string) error {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	_, err := redisClient.Do("ZREM", key, value)
	if err != nil {
		return err
	}
	return nil
}

func ZScore(key, value string) (float64, bool, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("ZSCORE", key, value)
	if err != nil {
		return 0, false, err
	}
	if reply == nil {
		return 0, false, nil
	}
	score, err := redis.Float64(reply, nil)
	if err != nil {
		return 0, false, err
	}
	return score, true, nil
}

func batchZRem(key string, members ...string) (int, error) {
	if len(members) == 0 {
		return 0, nil
	}

	redisClient := labelmsg.Get()
	defer redisClient.Close()

	args := make([]interface{}, 0, len(members)+2)
	args = append(args, key)
	for _, member := range members {
		args = append(args, member)
	}

	reply, err := redisClient.Do("ZREM", args...)
	if err != nil {
		return 0, err
	}

	deletedCount, ok := reply.(int64)
	if !ok {
		return 0, fmt.Errorf("unexpected reply type: %T", reply)
	}

	return int(deletedCount), nil
}

func BatchZRemSafe(key string, members []string) (totalDeleted int, err error) {
	batchSize := 200
	for i := 0; i < len(members); i += batchSize {
		end := i + batchSize
		if end > len(members) {
			end = len(members)
		}

		batch := members[i:end]
		deleted, err := batchZRem(key, batch...)
		if err != nil {
			return totalDeleted, err
		}
		totalDeleted += deleted
	}
	return totalDeleted, nil
}

func ZRangeByScore(key string, min, max float64) ([]string, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("ZRANGEBYSCORE", key, min, max)
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return nil, err
	}
	s, _ := redis.Strings(reply, nil)
	return s, nil
}

func ZRangeByScoreLimit(key string, min, max float64, offset, count int) ([]string, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count)
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return nil, err
	}
	s, _ := redis.Strings(reply, nil)
	return s, nil
}

type MemberScore struct {
	Member string
	Score  float64
}

func ZRangeWithScoresStruct(key string, start, stop int64) ([]MemberScore, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("ZRANGE", key, start, stop, "WITHSCORES")
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return nil, nil
	}

	stringResults, err := redis.Strings(reply, nil)
	if err != nil {
		return nil, err
	}

	var results []MemberScore
	for i := 0; i < len(stringResults); i += 2 {
		score, err := strconv.ParseFloat(stringResults[i+1], 64)
		if err != nil {
			return nil, err
		}
		results = append(results, MemberScore{
			Member: stringResults[i],
			Score:  score,
		})
	}

	return results, nil
}

func ZRemRangeByScore(key string, min, max float64) error {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	_, err := redisClient.Do("ZREMRANGEBYSCORE", key, min, max)
	if err != nil {
		return err
	}
	return nil
}

func ZCard(key string) (int64, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("ZCARD", key)
	if err != nil {
		return 0, err
	}
	return redis.Int64(reply, nil)
}

func SAdd(key string, vals ...string) (int64, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	args := make([]interface{}, 0, len(vals)+1)
	args = append(args, key)
	for _, val := range vals {
		args = append(args, val)
	}

	reply, err := redisClient.Do("SADD", args...)
	if err != nil {
		return 0, err
	}
	return redis.Int64(reply, nil)
}

func SRem(key string, vals ...string) (int64, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	args := make([]interface{}, 0, len(vals)+1)
	args = append(args, key)
	for _, val := range vals {
		args = append(args, val)
	}

	reply, err := redisClient.Do("SREM", args...)
	if err != nil {
		return 0, err
	}
	return redis.Int64(reply, nil)
}

func SISMember(key string, val string) (int, error) {
	redisClient := labelmsg.Get()
	defer redisClient.Close()

	reply, err := redisClient.Do("SISMEMBER", key, val)
	if err != nil {
		return 0, err
	}
	return redis.Int(reply, nil)
}
