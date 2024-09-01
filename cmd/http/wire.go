//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/rest"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/google/wire"
)

func initialize() (*rest.Server, error) {
	wire.Build(
		repository.DefaultSet,
		usecase.DefaultSet,
		rest.DefaultSet,
	)

	return nil, nil
}
