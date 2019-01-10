package cache

import (
	"github.com/go-redis/redis"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
)

var RedisDB *redis.Client

const (
	connectRedisFailed  = "Failed to connect to Redis"
	connectRedisSucceed = "Success to connect to Redis"
)

func Connect(address, password string) {
	/*Redis*/
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // use default DB
	})

	pong, err := RedisDB.Ping().Result()

	util.FatalOnError(err, connectRedisFailed, pong)

	log.Info(connectRedisSucceed)
}
