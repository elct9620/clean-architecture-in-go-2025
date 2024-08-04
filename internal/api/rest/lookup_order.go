package rest

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
)

func (api *Api) LookupOrder(w http.ResponseWriter, r *http.Request, orderId string) {
	out, err := api.LookupOrderUsecase.Execute(r.Context(), &usecase.LookupOrderInput{
		Id: orderId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := buildLookupOrderResponse(out)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func buildLookupOrderResponse(out *usecase.LookupOrderOutput) *LookupOrderResponse {
	outItems := make([]OrderItem, 0, len(out.Items))
	for _, item := range out.Items {
		outItems = append(outItems, OrderItem{
			Name:      item.Name,
			Quantity:  float32(item.Quantity),
			UnitPrice: float32(item.UnitPrice),
		})
	}

	return &LookupOrderResponse{
		Id:    out.Id,
		Name:  out.Name,
		Items: outItems,
	}
}
