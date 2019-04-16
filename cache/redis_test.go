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
