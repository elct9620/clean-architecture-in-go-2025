package orders

import "errors"

var (
	ErrItemNameMustBeUnique = errors.New("item name must be unique")
	ErrOrderNotFound        = errors.New("order not found")
)

type Order struct {
	id           string
	customerName string
	items        []*Item
}

func New(id string, customerName string) *Order {
	return &Order{
		id:           id,
		customerName: customerName,
		items:        []*Item{},
	}
}

func (o *Order) Id() string {
	return o.id
}

func (o *Order) CustomerName() string {
	return o.customerName
}

func (o *Order) Items() []*Item {
	return o.items
}

func (o *Order) HasItem(name string) bool {
	for _, item := range o.items {
		if item.name == name {
			return true
		}
	}

	return false
}

func (o *Order) AddItem(name string, quantity int, unitPrice int) error {
	if o.HasItem(name) {
		return ErrItemNameMustBeUnique
	}

	o.items = append(o.items, &Item{
		name:      name,
		quantity:  quantity,
		unitPrice: unitPrice,
	})

	return nil
}
