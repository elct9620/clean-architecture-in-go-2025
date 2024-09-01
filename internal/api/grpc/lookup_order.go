package grpc

import (
	"context"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/elct9620/clean-architecture-in-go-2025/pkg/orderspb"
)

func (s *OrderServer) LookupOrder(ctx context.Context, req *orderspb.LookupOrderRequest) (*orderspb.LookupOrderResponse, error) {
	out, err := s.LookupOrderUsecase.Execute(ctx, &usecase.LookupOrderInput{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	res := buildLookupOrderResponse(out)

	return res, nil
}

func buildLookupOrderResponse(out *usecase.LookupOrderOutput) *orderspb.LookupOrderResponse {
	outItems := make([]*orderspb.OrderItem, 0, len(out.Items))
	for _, item := range out.Items {
		outItems = append(outItems, &orderspb.OrderItem{
			Name:      item.Name,
			Quantity:  int32(item.Quantity),
			UnitPrice: int32(item.UnitPrice),
		})
	}

	return &orderspb.LookupOrderResponse{
		Id:    out.Id,
		Name:  out.Name,
		Items: outItems,
	}
}
