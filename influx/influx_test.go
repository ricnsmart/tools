package influx

import (
	"encoding/json"
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	Connect("http://influx_dev.ricnsmart.com:8086", "ricn", "ricn@2018", "ricnsmart_dev")

	m, err := Query(`select count(ID) as total from "alarm"`)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(m[0]["total"].(json.Number))
}
