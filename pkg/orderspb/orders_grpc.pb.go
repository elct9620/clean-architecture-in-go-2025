// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.4
// source: pkg/orderspb/orders.proto

package orderspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	LookupOrder(ctx context.Context, in *LookupOrderRequest, opts ...grpc.CallOption) (*LookupOrderResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, "/orders.Order/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) LookupOrder(ctx context.Context, in *LookupOrderRequest, opts ...grpc.CallOption) (*LookupOrderResponse, error) {
	out := new(LookupOrderResponse)
	err := c.cc.Invoke(ctx, "/orders.Order/LookupOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	LookupOrder(context.Context, *LookupOrderRequest) (*LookupOrderResponse, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderServer) LookupOrder(context.Context, *LookupOrderRequest) (*LookupOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.Order/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_LookupOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).LookupOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.Order/LookupOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).LookupOrder(ctx, req.(*LookupOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _Order_PlaceOrder_Handler,
		},
		{
			MethodName: "LookupOrder",
			Handler:    _Order_LookupOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/orderspb/orders.proto",
}
