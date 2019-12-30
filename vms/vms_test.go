package vms

import (
	"log"
	"testing"
)

func TestCall(t *testing.T) {
	Connect("ricnsmart.com", "localhost:7781")
	reply, err := Call("13205173164", "TTS_180347394", `{"Name":"测试","Time":"2019-12-30"}`)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(reply)
}
