package util

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	provides requirements to interact with mongodb
*/
const URI string = "mongodb://localhost:27017"

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(URI)
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	return client
}

func GetDb(database string) *mongo.Database {
	client := GetClient()
	if client != nil {
		return client.Database(database)
	}
	return nil
}

func GetCollection(database string, collection string) *mongo.Collection {
	db := GetDb(database)
	if db != nil {
		return db.Collection(collection)
	}
	return nil
}
