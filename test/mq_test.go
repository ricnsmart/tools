package test

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/mq"
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

	mq.Connect("ricn", "ricnsmart2018", "localhost:5672")

	var alarm alarm

	result, err := mq.Subscribe("pmc350f/alarm")

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
	mq.Connect("ricn", "ricnsmart2018", "localhost:5672")

	err := mq.TopicEmit(`system_notification`, `1975b244-f1b5-4371-b03c-9bba6ee25b4b`, []byte{3})

	if err != nil {
		log.Error(err)
	}
}

func TestTopicReceive(t *testing.T) {
	mq.Connect("ricn", "ricnsmart2018", "localhost:5672")

	ch, err := mq.TopicReceive(`system_notification`, `*`, `ee7b760d-8068-4697-8b37-ddb625650b91`, `1975b244-f1b5-4371-b03c-9bba6ee25b4b`)

	if err != nil {
		log.Error(err)
		return
	}

	for msg := range ch {
		log.Info(msg.Body)
	}
}
