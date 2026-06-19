package routes

import (
	"net/http"
	"task-api/auth"
	"task-api/handlers"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetUpRoute(client *mongo.Client) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", handlers.HealthCheck())
	mux.HandleFunc("/api/auth/register", auth.Register(client))
	mux.HandleFunc("/api/auth/login", auth.Login(client))

	mux.HandleFunc("/api/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTask(client)(w, r)
		case http.MethodPost:
			handlers.PostTask(client)(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
