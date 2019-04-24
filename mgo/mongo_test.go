package mgo

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/rules"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	//update := bson.D{{"$set",bson.A{}}}
	//log.Print(len(update))
	Connect("localhost:27017", "ricnsmart_dev")

	filter := bson.M{"deviceid": "6d8a69ce-8b4d-4caf-be55-eaa2c828bfaa"}
	//filter := bsonx.Doc{{"DeviceID", bsonx.String("6c76660a-d285-48b7-b4c8-ac7a70179153")}}
	//update := bsonx.Doc{{"$set", bsonx.Document(bsonx.Doc{{"Metric.Ia.Alert", bsonx.Int32(80)}})}}

	update := bson.D{{"$set", bson.M{"metrics.ia": bson.M{
		"smsswitch":   false,
		"warnswitch":  false,
		"alertswitch": false,
		"max":         300,
		"min":         0,
		"scale":       5,
		"warn":        62,
		"alert":       82,
		"spl":         12,
	}}}}

	_, err := MongoDB.Collection(rules.DevicesCollection).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

}

func TestDecode(t *testing.T) {
	var deviceMetric RCN350F

	Connect("localhost:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"deviceid": "6d8a69ce-8b4d-4caf-be55-eaa2c828bfaa"}

	err := MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).Decode(&deviceMetric)
	//b, err := MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).DecodeBytes()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(`%+v`, deviceMetric)
	//log.Print(b.String())

}

func TestOR(t *testing.T) {
	Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	a := bson.A{}

	a = append(a, bson.M{"deviceid": "1354b470-7733-4194-b406-f9a5999d8d57"}, bson.M{"deviceid": "aabb6480-6412-485a-9397-7277f110704d"})

	filter := bson.D{
		{"$or", a},
	}

	//filter := bson.M{"deviceid": "1354b470-7733-4194-b406-f9a5999d8d57"}

	//b, err := MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).DecodeBytes()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Print(b.String())

	//
	cur, err := MongoDB.Collection(rules.DevicesCollection).Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	var deviceMetrics []RCN350F
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var deviceMetric RCN350F
		err := cur.Decode(&deviceMetric)
		if err != nil {
			log.Fatal(err)
		}
		deviceMetrics = append(deviceMetrics, deviceMetric)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", deviceMetrics)

}

func TestDelete(t *testing.T) {
	Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	arr := bson.A{}

	for _, deviceID := range []string{"1354b470-7733-4194-b406-f9a5999d8d57", "aabb6480-6412-485a-9397-7277f110704d"} {
		arr = append(arr, bson.M{"deviceid": deviceID})
	}

	filter := bson.D{{"$or", arr}}

	_, err := MongoDB.Collection(rules.DevicesCollection).DeleteMany(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
}

type RCN350F struct {
	CreateAt     string
	UpdateAt     string
	DeviceID     string
	GPRSOperator int
	DomainRecord string
	CT           int
	Interval     int
	SMSLimit     int

	//  rcn350f特有
	Buzzer     byte // 蜂鸣器开关
	BreakShort byte
	ICCID      string

	Metrics struct {
		DO1 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		DO2 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		DI1 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		DI2 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		Ia struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        float32
			Alert       float32
			SPL         int
		}
		Ib struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        float32
			Alert       float32
			SPL         int
		}
		Ic struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        float32
			Alert       float32
			SPL         int
		}
		IR struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Warn        float32
			Alert       float32
			SPL         int
		}
		T1 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        float32
			Alert       float32
		}
		T2 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        float32
			Alert       float32
		}
		T3 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        float32
			Alert       float32
		}
		T4 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Warn        float32
			Alert       float32
		}
		Ua struct {
			WarnSwitch    bool
			AlertSwitch   bool
			SMSSwitch     bool
			WarnSMSSwitch bool
			Min           int
			Max           int
			Scale         int
			Warn          float32
			Alert         float32
		}
		Ub struct {
			WarnSwitch    bool
			AlertSwitch   bool
			SMSSwitch     bool
			WarnSMSSwitch bool
			Min           int
			Max           int
			Scale         int
			Warn          float32
			Alert         float32
		}
		Uc struct {
			WarnSwitch    bool
			AlertSwitch   bool
			SMSSwitch     bool
			WarnSMSSwitch bool
			Min           int
			Max           int
			Scale         int
			Warn          float32
			Alert         float32
		}
	}
}

func TestMigration(t *testing.T) {
	Connect("mongodb://39.104.158.136:27017", "ricnsmart")

	var docs []interface{}

	//var origin = bson.M{}

	arr := []string{
		"3697a8cd-6fc2-4416-ad0a-f45434756b75",
	}

	for _, deviceID := range arr[1:3] {
		var c = bson.M{}

		c["deviceid"] = deviceID

		docs = append(docs, c)
	}

	log.Print(docs)
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//
	//_, err := MongoDB.Collection(rules.DevicesCollection).InsertMany(ctx, docs)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

}

func TestFind(t *testing.T) {
	Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	//var result = struct {
	//	CT float64
	//}{}

	var d RCN350F

	err := MongoDB.Collection(rules.DevicesCollection).FindOne(context.Background(), bson.D{{"deviceid", "493be9e5-f429-40c0-85d8-5d38a299ff93"}}).Decode(&d)

	if err != nil {
		log.Fatal(err)
	}

	log.Infof("%+v", d)

	//log.Print(result)
}
