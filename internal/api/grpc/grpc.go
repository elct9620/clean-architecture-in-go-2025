package grpc

import (
	"net"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/elct9620/clean-architecture-in-go-2025/pkg/orderspb"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var DefaultSet = wire.NewSet(
	wire.Struct(new(OrderServer), "*"),
	NewServer,
)

var _ orderspb.OrderServer = &OrderServer{}

type OrderServer struct {
	orderspb.OrderServer `wire:"-"`
	PlaceOrderUsecase    *usecase.PlaceOrder
	LookupOrderUsecase   *usecase.LookupOrder
}

type Server struct {
	grpc *grpc.Server
}

func NewServer(
	orderServer *OrderServer,
) *Server {
	server := grpc.NewServer()

	orderspb.RegisterOrderServer(server, orderServer)
	reflection.Register(server)

	return &Server{
		grpc: server,
	}
}

func (s *Server) Serve() error {
	socket, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	return s.grpc.Serve(socket)
}
