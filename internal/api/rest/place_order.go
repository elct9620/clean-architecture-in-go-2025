package rest

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
)

func (api *Api) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input := buildPlaceOrderInput(req)

	out, err := api.PlaceOrderUsecase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := buildPlaceOrderResponse(out)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func buildPlaceOrderInput(req PlaceOrderRequest) *usecase.PlaceOrderInput {
	inItems := make([]usecase.PlaceOrderItem, 0, len(req.Items))
	for _, item := range req.Items {
		inItems = append(inItems, usecase.PlaceOrderItem{
			Name:      item.Name,
			Quantity:  int(item.Quantity),
			UnitPrice: int(item.UnitPrice),
		})
	}

	return &usecase.PlaceOrderInput{
		Name:  req.Name,
		Items: inItems,
	}
}

func buildPlaceOrderResponse(out *usecase.PlaceOrderOutput) *PlaceOrderResponse {
	outItems := make([]OrderItem, 0, len(out.Items))
	for _, item := range out.Items {
		outItems = append(outItems, OrderItem{
			Name:      item.Name,
			Quantity:  float32(item.Quantity),
			UnitPrice: float32(item.UnitPrice),
		})
	}

	return &PlaceOrderResponse{
		Id:    out.Id,
		Name:  out.Name,
		Items: outItems,
	}
}
