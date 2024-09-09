package main

import (
	"flag"
	"log"

	bolt "go.etcd.io/bbolt"
)

func provideBoltDb() (*bolt.DB, error) {
	return bolt.Open("bolt.db", 0600, nil)
}

func main() {
	databaseType := flag.String("database", "in-memory", "Database type to use")
	flag.Parse()

	server, cleanup, err := initialize(*databaseType)
	if err != nil {
		log.Fatalf("Error initializing handler: %v", err)
	}
	defer cleanup()

	if err := server.Serve(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
