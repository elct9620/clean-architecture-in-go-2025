package repository

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewInMemoryOrderRepository,
	wire.Bind(new(usecase.OrderRepository), new(*InMemoryOrderRepository)),
	NewInMemoryTokenRepository,
	wire.Bind(new(usecase.TokenRepository), new(*InMemoryTokenRepository)),
)

var BoltSet = wire.NewSet(
	NewBoltOrderRepository,
	wire.Bind(new(usecase.OrderRepository), new(*BoltOrderRepository)),
	NewBoltTokenRepository,
	wire.Bind(new(usecase.TokenRepository), new(*BoltTokenRepository)),
)
