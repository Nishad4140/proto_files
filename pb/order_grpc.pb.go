// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: order.proto

package pb

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	OrderAll(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*OrderId, error)
	CancelOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*OrderId, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) OrderAll(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*OrderId, error) {
	out := new(OrderId)
	err := c.cc.Invoke(ctx, "/cart.OrderService/OrderAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*OrderId, error) {
	out := new(OrderId)
	err := c.cc.Invoke(ctx, "/cart.OrderService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	OrderAll(context.Context, *OrderId) (*OrderId, error)
	CancelOrder(context.Context, *OrderId) (*OrderId, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) OrderAll(context.Context, *OrderId) (*OrderId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderAll not implemented")
}
func (UnimplementedOrderServiceServer) CancelOrder(context.Context, *OrderId) (*OrderId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_OrderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.OrderService/OrderAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderAll(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.OrderService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderAll",
			Handler:    _OrderService_OrderAll_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
