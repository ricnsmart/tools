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
	connectRedisSucceed = "Redis connected!"
)

func Connect(address, password string) {
	/*Redis*/
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := RedisDB.Ping().Result()

	util.FatalOnError(err, connectRedisFailed, " Address:", address)

	log.Info(connectRedisSucceed, " Address:", address)
}

// 广播结构体或者Map
func Publish(key string, i interface{}) error {

	bytes, err := json.Marshal(i)

	if err != nil {
		log.Error(util.MarshalFailed.String(), i, err)
		return err
	}

	err = RedisDB.Publish(key, bytes).Err()

	if err != nil {
		log.Error(util.RedisPublishFailed.String(), i, err)
	}

	return err
}

// 缓存结构体或者Map
func Set(key string, i interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(i)

	if err != nil {
		log.Error(util.MarshalFailed.String(), i, err)
		return err
	}

	err = RedisDB.Set(key, bytes, expiration).Err()

	if err != nil {
		log.Error(util.SetCacheFailed.String(), i, err)
	}

	return err
}

// 用hash表存储结构体或Map
func HSet(key, field string, i interface{}) error {
	bytes, err := json.Marshal(i)

	if err != nil {
		log.Error(util.MarshalFailed.String(), i, err)
		return err
	}

	err = RedisDB.HSet(key, field, string(bytes)).Err()

	if err != nil {
		log.Error(util.SetCacheFailed.String(), i, err)
	}

	return err
}
