package main

import (
	"context"
	"fmt"
	"net/http"

	"task-api/db"
	"task-api/handlers"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const PORT = ":8080"

type Task struct {
	Name string
	Status string
}

func main() {
	client := db.ConnectDb()

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

	// findind the task
	err = coll.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&resultTask)
	if err != nil {
		fmt.Printf("Failed to find task with err = %v\n", err)
	} else {
		fmt.Printf("Found task: %v\n", resultTask)
	}

	// deleting the task
	delResult, delErr := coll.DeleteMany(context.TODO(), bson.M{"_id": result.InsertedID})

	if delErr != nil {
		panic(delErr)
	}

	fmt.Printf("Document delete: %d\n", delResult.DeletedCount)

	// routes
	http.HandleFunc("/api", handlers.HealthCheck())

	fmt.Printf("server is running on port %s", PORT)
	serverErr := http.ListenAndServe(PORT, nil)
	if serverErr != nil {
		fmt.Printf("server error: %v", serverErr)
	}
}
