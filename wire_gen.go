// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/rest"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/go-chi/chi/v5"
)

// Injectors from wire.go:

func initializeTest() (*rest.Server, error) {
	mux := chi.NewRouter()
	inMemoryOrderRepository := repository.NewInMemoryOrderRepository()
	placeOrder := usecase.NewPlaceOrder(inMemoryOrderRepository)
	api := &rest.Api{
		PlaceOrderUsecase: placeOrder,
	}
	server, err := rest.NewServer(mux, api)
	if err != nil {
		return nil, err
	}
	return server, nil
}
