package main

import (
	"context"
	"fmt"
	"net/http"

	"task-api/db"
	"task-api/routes"
)

const PORT = ":8080"

func main() {
	client := db.ConnectDb()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Connection successfull")

	router := routes.SetUpRoute(client)

	fmt.Printf("server is running on port %s", PORT)
	serverErr := http.ListenAndServe(PORT, router)
	if serverErr != nil {
		fmt.Printf("server error: %v", serverErr)
	}
}
