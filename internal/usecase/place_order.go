package usecase

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
	"github.com/google/uuid"
)

type PlaceOrderItem struct {
	Name      string `validate:"required"`
	Quantity  int    `validate:"required,gte=1"`
	UnitPrice int    `validate:"required,gte=1"`
}

type PlaceOrderInput struct {
	Name  string           `validate:"required"`
	Items []PlaceOrderItem `validate:"required,gt=0,dive,required"`
}

type PlaceOrderOutput struct {
	Id    string
	Name  string
	Items []PlaceOrderItem
}

type PlaceOrder struct {
	orders    OrderRepository
	tokens    TokenRepository
	validator Validator
}

func NewPlaceOrder(orders OrderRepository, tokens TokenRepository, validator Validator) *PlaceOrder {
	return &PlaceOrder{
		orders:    orders,
		tokens:    tokens,
		validator: validator,
	}
}

func (u *PlaceOrder) Execute(ctx context.Context, input *PlaceOrderInput) (*PlaceOrderOutput, error) {
	if err := u.validator.Validate(ctx, input); err != nil {
		return nil, err
	}

	nameToken := tokens.New(uuid.NewString())
	nameToken.SetData([]byte(input.Name))

	if err := u.tokens.Save(ctx, nameToken); err != nil {
		return nil, err
	}

	order := orders.New(
		uuid.NewString(),
		nameToken.String(),
	)

	for _, item := range input.Items {
		err := order.AddItem(item.Name, item.Quantity, item.UnitPrice)
		if err != nil {
			return nil, err
		}
	}

	if err := u.orders.Save(ctx, order); err != nil {
		return nil, err
	}

	out := &PlaceOrderOutput{
		Id:    order.Id(),
		Name:  string(nameToken.Data()),
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
