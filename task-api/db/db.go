package db

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDb() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	
	if uri == "" {
		log.Fatal("Set your 'MONGO_URI' environment variable.")
	}

	// mongodb connection
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client

}
