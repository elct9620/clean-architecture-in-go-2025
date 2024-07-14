//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml openapi.yaml
package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"

	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

var DefaultSet = wire.NewSet(
	chi.NewRouter,
	NewServer,
)

var _ http.Handler = &Server{}
var _ ServerInterface = &Server{}

type Server struct {
	router *chi.Mux
}

func NewServer(router *chi.Mux) (*Server, error) {
	apiDoc, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	router.Use(nethttpmiddleware.OapiRequestValidator(apiDoc))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	server := &Server{router: router}
	HandlerFromMux(server, router)

	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
