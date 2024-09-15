//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"errors"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/grpc"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository/sqlite"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/validator"
	"github.com/google/wire"
)

func initialize(databaseType string) (*grpc.Server, func(), error) {
	switch databaseType {
	case "in-memory":
		return initializeInMemory()
	case "bolt":
		return initializeBolt()
	case "sqlite":
		return initializeSQLite()
	default:
		return nil, func() {}, errors.New("unsupported database type")
	}
}

func initializeInMemory() (*grpc.Server, func(), error) {
	wire.Build(
		repository.DefaultSet,
		usecase.DefaultSet,
		validator.DefaultSet,
		grpc.DefaultSet,
	)

	return nil, nil, nil
}

func initializeBolt() (*grpc.Server, func(), error) {
	wire.Build(
		provideBoltDb,
		repository.BoltSet,
		usecase.DefaultSet,
		validator.DefaultSet,
		grpc.DefaultSet,
	)

	return nil, nil, nil
}

func initializeSQLite() (*grpc.Server, func(), error) {
	wire.Build(
		provideSQLiteDb,
		wire.Bind(new(sqlite.DBTX), new(*sql.DB)),
		repository.SQLiteSet,
		usecase.DefaultSet,
		validator.DefaultSet,
		grpc.DefaultSet,
	)

	return nil, nil, nil
}
