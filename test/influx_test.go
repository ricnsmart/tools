package test

import (
	"encoding/json"
	"fmt"
	"github.com/ricnsmart/tools/influx"
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	influx.Connect("http://influx_dev.ricnsmart.com:8086", "ricn", "ricn@2018", "ricnsmart_dev")

	m, err := influx.Query(`select count(ID) as total from "alarm"`)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(m[0]["total"].(json.Number))
}

func TestJoint(t *testing.T) {

	keys := []string{"12345", "23423ja"}

	cmd := influx.Joint("select * ", "alarm",
		fmt.Sprintf(`ORDER BY time DESC  LIMIT %v OFFSET %v tz('Asia/Shanghai')`, 10,
			10*(1-1)),
		&influx.Or{Body: `"DeviceID"='%v'`, Keys: keys},
		[]string{fmt.Sprintf(`(time >='%v' and time <= '%v')`, 1, 2)},
	)
	log.Print(cmd)
}
