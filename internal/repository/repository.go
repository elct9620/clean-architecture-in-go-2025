package repository

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository/sqlite"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
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

type SQLiteDDL string

var SQLiteSet = wire.NewSet(
	sqlite.New,
	NewSQLiteOrderRepository,
	wire.Bind(new(usecase.OrderRepository), new(*SQLiteOrderRepository)),
	NewSQLiteTokenRepository,
	wire.Bind(new(usecase.TokenRepository), new(*SQLiteTokenRepository)),
)
