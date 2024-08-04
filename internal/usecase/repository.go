package usecase

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
)

type OrderRepository interface {
	Find(ctx context.Context, id string) (*orders.Order, error)
	Save(ctx context.Context, order *orders.Order) error
}
