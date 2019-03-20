package mgo

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	connectMongoFailed  = "Failed to connect to Mongo"
	connectMongoSucceed = "Success to connect to Mongo"
)

var MongoDB *mongo.Database

func Connect(address, dbName string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))

	util.FatalOnError(err, connectMongoFailed)

	err = client.Ping(ctx, readpref.Primary())

	util.FatalOnError(err, connectMongoFailed)

	MongoDB = client.Database(dbName)

	log.Info(connectMongoSucceed)
}
