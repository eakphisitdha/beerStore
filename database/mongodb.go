package database

import (
	"fmt"

	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collection *mongo.Collection
var ConnectionURI = "mongodb://localhost:27017/"

func MongoDB() *mongo.Database {

	var Client, Err = mongo.NewClient(options.Client().ApplyURI(ConnectionURI))
	if Err != nil {
		log.Fatal(Err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Err = Client.Connect(ctx)
	if Err != nil {
		log.Fatal(Err)
	}

	Err = Client.Ping(ctx, readpref.Primary())
	if Err != nil {
		log.Fatal(Err)
	}
	conn := Client.Database("beerStore")

	fmt.Println("MongoDB Connected")
	return conn
}
