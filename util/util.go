package util

import "github.com/labstack/gommon/log"

func FatalOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Reply struct {
	Code    int
	Data    interface{}
	Message string
	Name    string
}
