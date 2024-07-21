package usecase

import (
	"context"

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
	return &PlaceOrderOutput{
		Id:    uuid.NewString(),
		Name:  input.Name,
		Items: input.Items,
	}, nil
}
