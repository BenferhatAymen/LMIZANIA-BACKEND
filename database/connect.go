package database

import (
	"context"

	"lmizania/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBInstance() *mongo.Client {

	log.Println("Connecting to database")
	log.Println(config.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_URI))

	if err != nil {
		log.Fatal("connection error", err)
	}
	client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	return client
}

var Client *mongo.Client
