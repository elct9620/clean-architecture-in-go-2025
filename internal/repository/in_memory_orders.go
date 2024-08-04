package repository

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
)

type InMemoryOrderItemSchema struct {
	Name      string
	Quantity  int
	UnitPrice int
}

type InMemoryOrderSchema struct {
	Id           string
	CustomerName string
	Items        []InMemoryOrderItemSchema
}

type InMemoryOrderRepository struct {
	orders map[string]InMemoryOrderSchema
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: map[string]InMemoryOrderSchema{},
	}
}

func (r *InMemoryOrderRepository) Find(ctx context.Context, id string) (*orders.Order, error) {
	orderSchema, ok := r.orders[id]
	if !ok {
		return nil, orders.ErrOrderNotFound
	}

	order := orders.New(id, orderSchema.CustomerName)
	for _, itemSchema := range orderSchema.Items {
		err := order.AddItem(itemSchema.Name, itemSchema.Quantity, itemSchema.UnitPrice)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (r *InMemoryOrderRepository) Save(ctx context.Context, order *orders.Order) error {
	items := []InMemoryOrderItemSchema{}

	for _, item := range order.Items() {
		items = append(items, InMemoryOrderItemSchema{
			Name:      item.Name(),
			Quantity:  item.Quantity(),
			UnitPrice: item.UnitPrice(),
		})
	}

	r.orders[order.Id()] = InMemoryOrderSchema{
		Id:           order.Id(),
		CustomerName: order.CustomerName(),
		Items:        items,
	}

	return nil
}
