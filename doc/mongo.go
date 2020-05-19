package doc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var db *mongo.Database

func Connect(address, dbName string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		return err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	db = client.Database(dbName)
	return nil
}

func Collection(name string) *mongo.Collection {
	return db.Collection(name)
}
