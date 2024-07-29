package usecase

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
)

type OrderRepository interface {
	Save(ctx context.Context, order *orders.Order) error
}
