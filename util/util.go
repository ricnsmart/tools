package util

import "github.com/labstack/gommon/log"

func FatalOnError(err error, msg string, data ...interface{}) {
	if err != nil {
		log.Fatalf("%s: %s ; data: %v", msg, err, data)
	}
}

func LogOnError(code string, err error, data ...interface{}) {
	log.Errorf("%s: %s ;data: %v", code, err, data)
}
