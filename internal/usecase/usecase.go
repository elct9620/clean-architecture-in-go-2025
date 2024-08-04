package usecase

import "github.com/google/wire"

var DefaultSet = wire.NewSet(
	NewPlaceOrder,
	NewLookupOrder,
)
