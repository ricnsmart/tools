package influx

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"strconv"
	"time"
)

var (
	c      client.Client
	DbName string
)

const (
	connectInfluxDBFailed = "Failed to connect to InfluxDB"
	testInfluxDBFailed    = "Failed to test connect InfluxDB"
)

func Connect(address, userName, password string) {

	var err error

	c, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     address,
		Username: userName,
		Password: password,
	})

	util.FatalOnError(err, connectInfluxDBFailed, address, userName, password)

	duration, version, err := c.Ping(5 * time.Minute)

	util.FatalOnError(err, testInfluxDBFailed, duration, version)
}

func Write(measurement string, tags map[string]string, fields map[string]interface{}) (err error) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DbName,
		Precision: "s",
	})
	if err != nil {
		log.Error(DbName, err)
		return
	}

	// Create a point and add to batch
	pt, err := client.NewPoint(measurement, tags, fields, time.Now())

	if err != nil {
		log.Error(tags, fields, err)
		return
	}

	bp.AddPoint(pt)

	if err = c.Write(bp); err != nil {
		log.Error(err)
		return
	}

	return
}

func Query(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: DbName,
	}
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// 防止float类型 0，1等存成int类型
func SolveFloatInt(fields map[string]interface{}) {
	for key, value := range fields {
		switch value.(type) {
		case float64:
			// 将可能为正整数的float值全部+0.00001
			// 因为influxDB认为字面量4就是整数，而不是浮点数；而go中不是这样，字面量4可能是浮点数
			fields[key] = value.(float64) + 0.00001
		case float32:
			fields[key] = value.(float32) + 0.00001
			// influx不支持uint64
		case uint64:
			fields[key] = strconv.FormatUint(value.(uint64), 10)
		}
	}
}
