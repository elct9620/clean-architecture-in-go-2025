package usecase

import "context"

type LookupOrderItem struct {
	Name      string
	Quantity  int
	UnitPrice int
}

type LookupOrderInput struct {
	Id string
}

type LookupOrderOutput struct {
	Id    string
	Name  string
	Items []LookupOrderItem
}

type LookupOrder struct {
	orders OrderRepository
	tokens TokenRepository
}

func NewLookupOrder(orders OrderRepository, tokens TokenRepository) *LookupOrder {
	return &LookupOrder{
		orders: orders,
		tokens: tokens,
	}
}

func (u *LookupOrder) Execute(ctx context.Context, input *LookupOrderInput) (*LookupOrderOutput, error) {
	order, err := u.orders.Find(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	nameToken, err := u.tokens.Find(ctx, order.CustomerName())
	if err != nil {
		return nil, err
	}

	out := &LookupOrderOutput{
		Id:    order.Id(),
		Name:  string(nameToken.Raw()),
		Items: []LookupOrderItem{},
	}

	for _, item := range order.Items() {
		out.Items = append(out.Items, LookupOrderItem{
			Name:      item.Name(),
			Quantity:  item.Quantity(),
			UnitPrice: item.UnitPrice(),
		})
	}

	return out, nil
}
