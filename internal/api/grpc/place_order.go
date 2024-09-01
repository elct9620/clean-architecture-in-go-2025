package grpc

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/elct9620/clean-architecture-in-go-2025/pkg/orderspb"
)

func (s *OrderServer) PlaceOrder(ctx context.Context, req *orderspb.PlaceOrderRequest) (*orderspb.PlaceOrderResponse, error) {
	input := buildPlaceOrderInput(req)

	out, err := s.PlaceOrderUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	res := buildPlaceOrderResponse(out)

	return res, nil
}

func buildPlaceOrderInput(req *orderspb.PlaceOrderRequest) *usecase.PlaceOrderInput {
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

func buildPlaceOrderResponse(out *usecase.PlaceOrderOutput) *orderspb.PlaceOrderResponse {
	outItems := make([]*orderspb.OrderItem, 0, len(out.Items))
	for _, item := range out.Items {
		outItems = append(outItems, &orderspb.OrderItem{
			Name:      item.Name,
			Quantity:  int32(item.Quantity),
			UnitPrice: int32(item.UnitPrice),
		})
	}

	return &orderspb.PlaceOrderResponse{
		Id:    out.Id,
		Name:  out.Name,
		Items: outItems,
	}
}
