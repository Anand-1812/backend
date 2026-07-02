package main

import (
	"log"
	"rest-net-http/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run();err != nil {
		log.Fatalf("Error while running server: %v", err)
	}
}
