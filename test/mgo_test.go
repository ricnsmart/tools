package test

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"testing"
)

func TestUpdate(t *testing.T) {
	mgo.Connect("localhost:27017", "platform_dev")

	filter := bson.M{"DeviceID": "6c76660a-d285-48b7-b4c8-ac7a70179153"}
	//filter := bsonx.Doc{{"DeviceID", bsonx.String("6c76660a-d285-48b7-b4c8-ac7a70179153")}}
	update := bsonx.Doc{{"$set", bsonx.Document(bsonx.Doc{{"Interval", bsonx.Int32(40)}})}}

	_, err := mgo.MongoDB.Collection(`devices`).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

}
