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
}

func NewLookupOrder(orders OrderRepository) *LookupOrder {
	return &LookupOrder{
		orders: orders,
	}
}

func (u *LookupOrder) Execute(ctx context.Context, input *LookupOrderInput) (*LookupOrderOutput, error) {
	order, err := u.orders.Find(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	out := &LookupOrderOutput{
		Id:    order.Id(),
		Name:  order.CustomerName(),
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
