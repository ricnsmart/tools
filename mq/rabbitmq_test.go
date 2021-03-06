package mq

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
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

func Co() {
	Connect("ricn", "ricnsmart2018", "dev.ricnsmart.com:5672")
}

func TestReceive(t *testing.T) {
	Co()

	ch, err := Conn.Channel()

	util.FatalOnError(err, "Failed to open a channel")

	// 经过测试，channel上可以申明多个queue，因为channel数量是有限的，这点很重要
	go func() {
		msgs2, err := Receive(ch, "test2")

		util.FatalOnError(err, "Failed to receive2")

		for msg := range msgs2 {
			log.Print(string(msg.Body))
		}
	}()

	msgs, err := Receive(ch, "test")

	util.FatalOnError(err, "Failed to receive")

	for msg := range msgs {
		log.Print(string(msg.Body))
	}

}

func TestSend(t *testing.T) {
	Co()

	//err := Send("test", []byte("hello"))
	//
	//util.FatalOnError(err, "Failed to send")

	d := struct {
		ID    string
		Param string
		Value interface{}
	}{ID: "91b19ec8-52f5-4b3d-8385-8c4ba9ebac6f", Param: "Ia", Value: 88}

	b, _ := json.Marshal(d)

	err := Send("param", b)

	util.FatalOnError(err, "Failed to send")
}

func TestPublish(t *testing.T) {
	Co()

	d := struct {
		ID    string
		Param string
		Value interface{}
	}{ID: "91b19ec8-52f5-4b3d-8385-8c4ba9ebac6f", Param: "Ia", Value: 88}

	b, _ := json.Marshal(d)

	err := Publish("param", b)

	if err != nil {
		log.Fatal(err)
	}

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

func TestRoutePublish2(t *testing.T) {

	Connect("ricn", "ricnsmart2018", "dev.ricnsmart.com:5672")

	c, err := Conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	ch, err := Receive(c, "test2")

	go func() {
		time.Sleep(5000)

		err = c.Cancel("test2", true)

		if err != nil {
			log.Fatal(err)
		}

	}()

	for s := range ch {
		log.Print(s.Body)
	}

}
