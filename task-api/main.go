package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"task-api/handlers"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const PORT = ":8080"

type Task struct {
	Name string
	Status string
}

func main() {
	uri := os.Getenv("MONGO_URI")
	
	if uri == "" {
		log.Fatal("Set your 'MONGO_URI' environment variable.")
	}

	// mongodb connection
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO());err != nil {
			panic(err)
		}
	}()
	fmt.Println("Connection successfull")


	// creating a task
	coll := client.Database("db").Collection("tasks")
	doc := Task{Name: "study", Status: "pending"}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Insert task with _id: %v\n", result.InsertedID)
	var resultTask bson.M

	err = coll.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&resultTask)
	if err != nil {
		fmt.Printf("Failed to find task with err = %v\n", err)
	}

	fmt.Printf("Found task: %v\n", resultTask)


	// routes
	http.HandleFunc("/api", handlers.HealthCheck())

	fmt.Printf("server is running on port %s", PORT)
	serverErr := http.ListenAndServe(PORT, nil)
	if serverErr != nil {
		fmt.Printf("server error: %v", serverErr)
	}
}
