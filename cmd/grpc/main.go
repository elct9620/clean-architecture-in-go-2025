package main

import "log"

func main() {
	server, err := initialize()
	if err != nil {
		log.Fatalf("Error initializing handler: %v", err)
	}

	if err := server.Serve(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
