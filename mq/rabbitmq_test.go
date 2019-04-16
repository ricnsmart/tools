package mq

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"testing"
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

	Connect("ricn", "ricnsmart2018", "localhost:5672")

	var alarm alarm

	result, err := Subscribe("pmc350f/alarm")

	if err != nil {
		log.Print(err)
	}

	for ch := range result {
		err = json.Unmarshal(ch.Body, &alarm)

		if err != nil {
			log.Error(err, ch.Body)
			continue
		}

		log.Info(alarm)
	}
}

func TestTopicEmit(t *testing.T) {
	Connect("ricn", "ricnsmart2018", "dev.ricnsmart.com:5672")

	err := TopicEmit(`system_notification`, `1975b244-f1b5-4371-b03c-9bba6ee25b4b`, []byte{3})

	if err != nil {
		log.Error(err)
	}
}

func TestTopicReceive(t *testing.T) {
	Connect("ricn", "ricnsmart2018", "dev.ricnsmart.com:5672")

	ch, err := TopicReceive(`system_notification`, `*`, `ee7b760d-8068-4697-8b37-ddb625650b91`, `1975b244-f1b5-4371-b03c-9bba6ee25b4b`)

	if err != nil {
		log.Error(err)
		return
	}

	for msg := range ch {
		log.Info(msg.Body)
	}
}

func TestRoutePublish(t *testing.T) {

	Connect("ricn", "ricnsmart2018", "dev.ricnsmart.com:5672")

	err := RoutePublish(`test`, "1234", []byte{2})

	if err != nil {
		log.Fatal(err)
	}

}
