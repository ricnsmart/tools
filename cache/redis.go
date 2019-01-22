package cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"time"
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

func Publish(key string, i interface{}) error {

	bytes, err := json.Marshal(i)

	if err != nil {
		return err
	}

	err = RedisDB.Publish(key, bytes).Err()

	return err
}

func Set(key string, i interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(i)

	if err != nil {
		return err
	}

	err = RedisDB.Set(key, bytes, expiration).Err()

	return err
}
