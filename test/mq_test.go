package test

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/mq"
	"testing"
	"time"
)

type alarm struct {
	ID        string
	DeviceID  string
	Metrics   string
	AlarmType uint8
	Current   float32
	SetValue  float32
}

func TestAlarm(t *testing.T) {

	mq.Connect("ricn", "ricnsmart2018", "localhost:5672")

	var alarm alarm

	result, ch, err := mq.Subscribe("pmc350f/alarm")

	if err != nil {
		log.Print(err)
	}

	go func() {
		time.Sleep(5 * time.Second)
		ch.Close()
	}()

	for ch := range result {
		err = json.Unmarshal(ch.Body, &alarm)

		if err != nil {
			log.Error(err, ch.Body)
			continue
		}

		log.Info(alarm)
	}
}
