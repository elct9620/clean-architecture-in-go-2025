package rest

import "net/http"

func (s *Api) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
