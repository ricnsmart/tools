package test

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/rules"
	"github.com/ricnsmart/tools/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	mgo.Connect("localhost:27017", "platform_dev")

	filter := bson.M{"DeviceID": "6c76660a-d285-48b7-b4c8-ac7a70179153"}
	//filter := bsonx.Doc{{"DeviceID", bsonx.String("6c76660a-d285-48b7-b4c8-ac7a70179153")}}
	update := bsonx.Doc{{"$set", bsonx.Document(bsonx.Doc{{"Metric.Ia.Alert", bsonx.Int32(80)}})}}

	_, err := mgo.MongoDB.Collection(rules.DevicesCollection).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

}

func TestDecode(t *testing.T) {
	var deviceMetric DeviceMetric

	mgo.Connect("localhost:27017", "platform_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"deviceid": "6c76660a-d285-48b7-b4c8-ac7a70179153"}

	err := mgo.MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).Decode(&deviceMetric)
	b, err := mgo.MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).DecodeBytes()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(deviceMetric)
	log.Print(b.String())

}

type DeviceMetric struct {
	GPRSOperator int
	DomainRecord string
	Interval     int
	SMSLimit     int
	DeviceID     string
	//Metric   struct {
	//	DO1 struct {
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//	}
	//	DO2 struct {
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//	}
	//	DI1 struct {
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//	}
	//	DI2 struct {
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//	}
	//	Ia struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//		SPL         int
	//	}
	//	Ib struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//		SPL         int
	//	}
	//	Ic struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//		SPL         int
	//	}
	//	IR struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Warn        int
	//		Alert       int
	//		SPL         int
	//	}
	//	T1 struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//	}
	//	T2 struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//	}
	//	T3 struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//	}
	//	T4 struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Warn        int
	//		Alert       int
	//	}
	//	Ua struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//	}
	//	Ub struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//	}
	//	Uc struct {
	//		WarnSwitch  bool
	//		AlertSwitch bool
	//		SMSSwitch   bool
	//		Min         int
	//		Max         int
	//		Scale       int
	//		Warn        int
	//		Alert       int
	//	}
	//}
}
