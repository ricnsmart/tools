package cache

import (
	"log"
	"reflect"
	"testing"
)

func TestRedis(t *testing.T) {
	Connect("localhost:10032", "")

	strNum, err := RedisDB.HGet("11111", "222").Result()

	if err != nil {
		log.Print(strNum, reflect.TypeOf(strNum))
		log.Fatal(err)
	}

	log.Print(strNum, reflect.TypeOf(strNum))
}

func TestPipeline(t *testing.T) {
	Connect("39.104.176.28:10032", "ricnsmart2018")

	pipe := RedisDB.Pipeline()

	for _, v := range []string{"1", "2"} {
		pipe.Del(v)
	}

	_, err := pipe.Exec()

	if err != nil {
		log.Fatal("批量删除失败", err)
	}
}
