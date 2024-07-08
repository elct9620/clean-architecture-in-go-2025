package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	chi.NewRouter,
	NewServer,
)

var _ http.Handler = &Server{}

type Server struct {
	router *chi.Mux
}

func NewServer(router *chi.Mux) *Server {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	return &Server{router: router}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
