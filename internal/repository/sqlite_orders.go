package repository

import (
	"context"
	"database/sql"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository/sqlite"
	"github.com/google/uuid"
)

type SQLiteOrderRepository struct {
	db      *sql.DB
	queries *sqlite.Queries
}

func NewSQLiteOrderRepository(db *sql.DB, queries *sqlite.Queries) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{
		db:      db,
		queries: queries,
	}
}

func (r *SQLiteOrderRepository) Find(ctx context.Context, id string) (*orders.Order, error) {
	order, err := r.queries.FindOrder(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, orders.ErrOrderNotFound
		}

		return nil, err
	}

	orderEntity := orders.New(order.ID, order.CustomerName)

	orderItems, err := r.queries.FindOrderItems(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range orderItems {
		err := orderEntity.AddItem(item.Name, int(item.Quantity), int(item.UnitPrice))
		if err != nil {
			return nil, err
		}
	}

	return orderEntity, nil
}

func (r *SQLiteOrderRepository) Save(ctx context.Context, order *orders.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)

	createdOrder, err := qtx.CreateOrder(ctx, sqlite.CreateOrderParams{
		ID:           order.Id(),
		CustomerName: order.CustomerName(),
	})
	if err != nil {
		return err
	}

	for _, item := range order.Items() {
		_, err := qtx.CreateOrderItem(ctx, sqlite.CreateOrderItemParams{
			ID:        uuid.NewString(),
			OrderID:   createdOrder.ID,
			Name:      item.Name(),
			Quantity:  int64(item.Quantity()),
			UnitPrice: int64(item.UnitPrice()),
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
