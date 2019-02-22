package influx

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	c  client.Client
	db string
)

const (
	connectInfluxDBFailed  = "Failed to connect to InfluxDB"
	testInfluxDBFailed     = "Failed to test connect InfluxDB"
	connectInfluxDBSucceed = "Success to connect to InfluxDB"
)

type Or struct {
	Body string
	Keys []string
}

// 除了需要指定连接的用户名、密码、地址，还需要指定db
func Connect(address, userName, password, dbName string) {

	var err error

	c, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     address,
		Username: userName,
		Password: password,
	})

	util.FatalOnError(err, connectInfluxDBFailed, address, userName, password)

	duration, version, err := c.Ping(5 * time.Minute)

	util.FatalOnError(err, testInfluxDBFailed, duration, version)

	db = dbName

	log.Info(connectInfluxDBSucceed)
}

func Write(measurement string, tags map[string]string, fields map[string]interface{}) (err error) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	if err != nil {
		log.Error(db, err)
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

func Query(cmd string) ([]map[string]interface{}, error) {

	q := client.Query{
		Command:  cmd,
		Database: db,
	}

	response, err := c.Query(q)

	if err != nil {
		return nil, err
	}

	if err == nil && response.Error() != nil {
		return nil, response.Error()
	}

	var slc []map[string]interface{}

	if len(response.Results[0].Series) > 0 {

		result := response.Results[0].Series[0]

		columns := result.Columns

		values := result.Values

		for _, value := range values {

			m := make(map[string]interface{})

			for index, column := range columns {
				m[column] = value[index]
			}
			slc = append(slc, m)
		}
	}

	return slc, nil
}

// 解决influx对float和uint64支持不良的问题
// @param limit 误差范围
func FixInfluxType(fields map[string]interface{}, limit float32) {
	for key, value := range fields {
		switch value.(type) {
		case float64:
			i := value.(float64)
			// 将正整数形式的float值全部+0.00001
			// 因为influxDB认为字面量4就是整数，而不是浮点数；而go中不是这样，字面量4可能是浮点数
			if math.Ceil(i) == i {
				fields[key] = i + float64(limit)
			}
		case float32:
			i := value.(float32)
			if math.Ceil(float64(i)) == float64(i) {
				fields[key] = i + limit
			}
			// influx不支持uint64
		case uint64:
			fields[key] = strconv.FormatUint(value.(uint64), 10)
		}
	}
}

func Joint(body, table, option string, or *Or, andQuery []string) (cmd string) {

	cmd = fmt.Sprintf(`%v from "%v"`, body, table)

	if or != nil && len(or.Keys) > 0 {
		var tmp []string

		for _, key := range or.Keys {
			tmp = append(tmp, fmt.Sprintf(or.Body, key))
		}

		andQuery = append(andQuery, fmt.Sprintf(`(%v)`, strings.Join(tmp, " OR ")))
	}

	if len(andQuery) > 0 {
		cmd = fmt.Sprintf(`%v where %v`, cmd, strings.Join(andQuery, " AND "))
	}

	if option != "" {
		cmd = fmt.Sprintf(`%v %v`, cmd, option)
	}

	return
}
