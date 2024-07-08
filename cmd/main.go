package main

import (
	"log"
	"net/http"
)

func main() {
	handler, err := initialize()
	if err != nil {
		log.Fatalf("Error initializing handler: %v", err)
	}

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
