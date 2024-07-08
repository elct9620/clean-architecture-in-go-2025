//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/rest"
	"github.com/google/wire"
)

func initializeTest() (*rest.Server, error) {
	wire.Build(rest.DefaultSet)

	return nil, nil
}
