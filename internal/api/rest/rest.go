//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml openapi.yaml
package rest

import (
	"net/http"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"

	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

var DefaultSet = wire.NewSet(
	chi.NewRouter,
	wire.Struct(new(Api), "*"),
	NewServer,
)

var _ ServerInterface = &Api{}

type Api struct {
	PlaceOrderUsecase *usecase.PlaceOrder
}

var _ http.Handler = &Server{}

type Server struct {
	router *chi.Mux
}

func NewServer(router *chi.Mux, api *Api) (*Server, error) {
	apiDoc, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	router.Use(nethttpmiddleware.OapiRequestValidator(apiDoc))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	HandlerFromMux(api, router)

	return &Server{router: router}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
