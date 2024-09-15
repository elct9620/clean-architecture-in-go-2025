package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"log"

	appDb "github.com/elct9620/clean-architecture-in-go-2025/db"

	bolt "go.etcd.io/bbolt"
)

func provideBoltDb() (*bolt.DB, func(), error) {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
}

func provideSQLiteDb() (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		return nil, nil, err
	}

	if _, err := db.Exec(appDb.Schema); err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
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
