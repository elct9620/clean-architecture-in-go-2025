package repository

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewInMemoryOrderRepository,
	wire.Bind(new(usecase.OrderRepository), new(*InMemoryOrderRepository)),
)
