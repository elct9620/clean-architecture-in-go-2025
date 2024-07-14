package rest

import "net/http"

func (s *Server) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
