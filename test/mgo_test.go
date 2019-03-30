package test

import (
	"context"
	"fmt"
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

func TestOR(t *testing.T) {
	mgo.Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	a := bson.A{}

	a = append(a, bson.M{"deviceid": "1354b470-7733-4194-b406-f9a5999d8d57"}, bson.M{"deviceid": "aabb6480-6412-485a-9397-7277f110704d"})

	filter := bson.D{
		{"$or", a},
	}

	//filter := bson.M{"deviceid": "1354b470-7733-4194-b406-f9a5999d8d57"}

	//b, err := mgo.MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).DecodeBytes()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Print(b.String())

	//
	cur, err := mgo.MongoDB.Collection(rules.DevicesCollection).Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	var deviceMetrics []DeviceMetric
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var deviceMetric DeviceMetric
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
	mgo.Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	arr := bson.A{}

	for _, deviceID := range []string{"1354b470-7733-4194-b406-f9a5999d8d57", "aabb6480-6412-485a-9397-7277f110704d"} {
		arr = append(arr, bson.M{"deviceid": deviceID})
	}

	filter := bson.D{{"$or", arr}}

	_, err := mgo.MongoDB.Collection(rules.DevicesCollection).DeleteMany(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
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
