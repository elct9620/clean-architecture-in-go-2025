package usecase

import (
	"context"

	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewPlaceOrder,
	NewLookupOrder,
)

type Validator interface {
	Validate(ctx context.Context, input any) error
}
