//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml openapi.yaml
package testability

import (
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	wire.Struct(new(Api), "*"),
)

var _ ServerInterface = &Api{}

type Api struct {
	OrderRepository usecase.OrderRepository
	TokenRepository usecase.TokenRepository
}
