//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/grpc"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/validator"
	"github.com/google/wire"
)

func initialize() (*grpc.Server, error) {
	wire.Build(
		repository.DefaultSet,
		usecase.DefaultSet,
		validator.DefaultSet,
		grpc.DefaultSet,
	)

	return nil, nil
}
