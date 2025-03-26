// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/order/v1/vendor_order.proto

package orderv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrderAPI_CreateDraftOrder_FullMethodName = "/api.order.v1.OrderAPI/CreateDraftOrder"
	OrderAPI_GetDraftOrder_FullMethodName    = "/api.order.v1.OrderAPI/GetDraftOrder"
	OrderAPI_PlaceOrder_FullMethodName       = "/api.order.v1.OrderAPI/PlaceOrder"
	OrderAPI_ListDraftOrders_FullMethodName  = "/api.order.v1.OrderAPI/ListDraftOrders"
)

// OrderAPIClient is the client API for OrderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderAPIClient interface {
	CreateDraftOrder(ctx context.Context, in *CreateDraftOrderRequest, opts ...grpc.CallOption) (*CreateDraftOrderResponse, error)
	GetDraftOrder(ctx context.Context, in *GetDraftOrderRequest, opts ...grpc.CallOption) (*GetDraftOrderResponse, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	ListDraftOrders(ctx context.Context, in *ListDraftOrdersRequest, opts ...grpc.CallOption) (*ListDraftOrdersResponse, error)
}

type orderAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderAPIClient(cc grpc.ClientConnInterface) OrderAPIClient {
	return &orderAPIClient{cc}
}

func (c *orderAPIClient) CreateDraftOrder(ctx context.Context, in *CreateDraftOrderRequest, opts ...grpc.CallOption) (*CreateDraftOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDraftOrderResponse)
	err := c.cc.Invoke(ctx, OrderAPI_CreateDraftOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAPIClient) GetDraftOrder(ctx context.Context, in *GetDraftOrderRequest, opts ...grpc.CallOption) (*GetDraftOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDraftOrderResponse)
	err := c.cc.Invoke(ctx, OrderAPI_GetDraftOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAPIClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, OrderAPI_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAPIClient) ListDraftOrders(ctx context.Context, in *ListDraftOrdersRequest, opts ...grpc.CallOption) (*ListDraftOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDraftOrdersResponse)
	err := c.cc.Invoke(ctx, OrderAPI_ListDraftOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderAPIServer is the server API for OrderAPI service.
// All implementations should embed UnimplementedOrderAPIServer
// for forward compatibility.
type OrderAPIServer interface {
	CreateDraftOrder(context.Context, *CreateDraftOrderRequest) (*CreateDraftOrderResponse, error)
	GetDraftOrder(context.Context, *GetDraftOrderRequest) (*GetDraftOrderResponse, error)
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	ListDraftOrders(context.Context, *ListDraftOrdersRequest) (*ListDraftOrdersResponse, error)
}

// UnimplementedOrderAPIServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderAPIServer struct{}

func (UnimplementedOrderAPIServer) CreateDraftOrder(context.Context, *CreateDraftOrderRequest) (*CreateDraftOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDraftOrder not implemented")
}
func (UnimplementedOrderAPIServer) GetDraftOrder(context.Context, *GetDraftOrderRequest) (*GetDraftOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDraftOrder not implemented")
}
func (UnimplementedOrderAPIServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderAPIServer) ListDraftOrders(context.Context, *ListDraftOrdersRequest) (*ListDraftOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDraftOrders not implemented")
}
func (UnimplementedOrderAPIServer) testEmbeddedByValue() {}

// UnsafeOrderAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderAPIServer will
// result in compilation errors.
type UnsafeOrderAPIServer interface {
	mustEmbedUnimplementedOrderAPIServer()
}

func RegisterOrderAPIServer(s grpc.ServiceRegistrar, srv OrderAPIServer) {
	// If the following call pancis, it indicates UnimplementedOrderAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderAPI_ServiceDesc, srv)
}

func _OrderAPI_CreateDraftOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDraftOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).CreateDraftOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAPI_CreateDraftOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).CreateDraftOrder(ctx, req.(*CreateDraftOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAPI_GetDraftOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDraftOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).GetDraftOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAPI_GetDraftOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).GetDraftOrder(ctx, req.(*GetDraftOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAPI_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAPI_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAPI_ListDraftOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDraftOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).ListDraftOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAPI_ListDraftOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).ListDraftOrders(ctx, req.(*ListDraftOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderAPI_ServiceDesc is the grpc.ServiceDesc for OrderAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order.v1.OrderAPI",
	HandlerType: (*OrderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDraftOrder",
			Handler:    _OrderAPI_CreateDraftOrder_Handler,
		},
		{
			MethodName: "GetDraftOrder",
			Handler:    _OrderAPI_GetDraftOrder_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _OrderAPI_PlaceOrder_Handler,
		},
		{
			MethodName: "ListDraftOrders",
			Handler:    _OrderAPI_ListDraftOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/order/v1/vendor_order.proto",
}
