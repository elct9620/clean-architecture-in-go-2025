package usecase

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
	"github.com/google/uuid"
)

type PlaceOrderItem struct {
	Name      string
	Quantity  int
	UnitPrice int
}

type PlaceOrderInput struct {
	Name  string
	Items []PlaceOrderItem
}

type PlaceOrderOutput struct {
	Id    string
	Name  string
	Items []PlaceOrderItem
}

type PlaceOrder struct {
}

func NewPlaceOrder() *PlaceOrder {
	return &PlaceOrder{}
}

func (u *PlaceOrder) Execute(ctx context.Context, input *PlaceOrderInput) (*PlaceOrderOutput, error) {
	order := orders.New(
		uuid.NewString(),
		input.Name,
	)

	for _, item := range input.Items {
		err := order.AddItem(item.Name, item.Quantity, item.UnitPrice)
		if err != nil {
			return nil, err
		}
	}

	out := &PlaceOrderOutput{
		Id:    order.Id(),
		Name:  order.CustomerName(),
		Items: []PlaceOrderItem{},
	}

	for _, item := range order.Items() {
		out.Items = append(out.Items, PlaceOrderItem{
			Name:      item.Name(),
			Quantity:  item.Quantity(),
			UnitPrice: item.UnitPrice(),
		})
	}

	return out, nil
}
