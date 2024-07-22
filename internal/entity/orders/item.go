package orders

type Item struct {
	name      string
	quantity  int
	unitPrice int
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Quantity() int {
	return i.quantity
}

func (i *Item) UnitPrice() int {
	return i.unitPrice
}
