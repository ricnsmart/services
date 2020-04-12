package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type mgo struct{}

var Mgo mgo

var mongodb *mongo.Database

func InitMongodb(address, dbName string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	mongodb = client.Database(dbName)
}

func (mgo) Collection(name string) *mongo.Collection {
	return mongodb.Collection(name)
}
