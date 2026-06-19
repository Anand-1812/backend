package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Task struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string        `bson:"name" json:"name"`
	Status string        `bson:"status" json:"status"`
}

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(client *mongo.Client) *MongoTaskRepository {
	return &MongoTaskRepository{
		collection: client.Database("db").Collection("tasks"),
	}
}

func HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "server is running")
	}
}

// implement get task
func GetTask(client *mongo.Client) http.HandlerFunc {
	repo := NewMongoTaskRepository(client)

	return func(w http.ResponseWriter, r *http.Request) {

		tasksFromDb, err := repo.collection.Find(context.TODO(), bson.M{})
		if err != nil {
			http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
			return
		}

		defer tasksFromDb.Close(context.TODO())

		var tasks []Task
		err = tasksFromDb.All(context.TODO(), &tasks)
		if err != nil {
			http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

func PostTask(client *mongo.Client) http.HandlerFunc {
	repo := NewMongoTaskRepository(client)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var task Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid json", http.StatusBadRequest)
			return
		}

		result, err := repo.collection.InsertOne(context.Background(), task)
		if err != nil {
			http.Error(w, "Failed to insert data", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Task created",
			"id":      result.InsertedID,
		})
	}
}
