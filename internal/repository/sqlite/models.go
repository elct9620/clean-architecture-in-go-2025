// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlite

type Order struct {
	ID           string
	CustomerName string
}

type OrderItem struct {
	ID        string
	OrderID   string
	Name      string
	Quantity  int64
	UnitPrice int64
}

type Token struct {
	ID      string
	Data    []byte
	Version string
}
