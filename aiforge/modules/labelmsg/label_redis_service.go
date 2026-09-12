package labelmsg

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	redigo "github.com/gomodule/redigo/redis"
)

var pool *redigo.Pool

func Init() {
	redisBroker := setting.Broker
	pool_size := 20
	log.Info("start to connect redis.")
	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.DialURL(redisBroker)
		if err != nil {
			return nil, err
		}
		return c, nil
	}, pool_size)

}

func Get() redigo.Conn {
	return pool.Get()
}
