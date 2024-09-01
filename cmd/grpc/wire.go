//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/grpc"
	"github.com/google/wire"
)

func initialize() (*grpc.Server, error) {
	wire.Build(
		grpc.DefaultSet,
	)

	return nil, nil
}
