package testability

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
)

func (api *Api) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := orders.New(req.Id, req.Name)
	for _, item := range req.Items {
		order.AddItem(item.Name, int(item.Quantity), int(item.UnitPrice))
	}

	err := api.OrderRepository.Save(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &CreatedResponse{
		Id: order.Id(),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
