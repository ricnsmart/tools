package cache

import (
	"github.com/go-redis/redis"
	"github.com/ricnsmart/tools/util"
)

var RedisDB *redis.Client

func Connect(address, password string) {
	/*Redis*/
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // use default DB
	})

	pong, err := RedisDB.Ping().Result()

	util.FatalOnError(err, "连接失败", pong)
}
