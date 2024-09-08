package validator

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	New,
	wire.Bind(new(usecase.Validator), new(*Validator)),
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(ctx context.Context, i any) error {
	return v.validate.StructCtx(ctx, i)
}
