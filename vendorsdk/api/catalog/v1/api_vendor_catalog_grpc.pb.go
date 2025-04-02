// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/catalog/v1/api_vendor_catalog.proto

package catalogv1

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
	VendorCatalogAPI_VendorListCatalog_FullMethodName           = "/api.catalog.v1.VendorCatalogAPI/VendorListCatalog"
	VendorCatalogAPI_VendorListCatalogVariant_FullMethodName    = "/api.catalog.v1.VendorCatalogAPI/VendorListCatalogVariant"
	VendorCatalogAPI_VendorGetCatalogStockStatus_FullMethodName = "/api.catalog.v1.VendorCatalogAPI/VendorGetCatalogStockStatus"
)

// VendorCatalogAPIClient is the client API for VendorCatalogAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// VendorCatalogAPI provide an abstraction to all of read data access for catalog businesses
type VendorCatalogAPIClient interface {
	VendorListCatalog(ctx context.Context, in *VendorListCatalogRequest, opts ...grpc.CallOption) (*VendorListCatalogResponse, error)
	VendorListCatalogVariant(ctx context.Context, in *VendorListCatalogVariantRequest, opts ...grpc.CallOption) (*VendorListCatalogVariantResponse, error)
	VendorGetCatalogStockStatus(ctx context.Context, in *VendorListCatalogVariantRequest, opts ...grpc.CallOption) (*VendorListCatalogVariantResponse, error)
}

type vendorCatalogAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewVendorCatalogAPIClient(cc grpc.ClientConnInterface) VendorCatalogAPIClient {
	return &vendorCatalogAPIClient{cc}
}

func (c *vendorCatalogAPIClient) VendorListCatalog(ctx context.Context, in *VendorListCatalogRequest, opts ...grpc.CallOption) (*VendorListCatalogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VendorListCatalogResponse)
	err := c.cc.Invoke(ctx, VendorCatalogAPI_VendorListCatalog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorCatalogAPIClient) VendorListCatalogVariant(ctx context.Context, in *VendorListCatalogVariantRequest, opts ...grpc.CallOption) (*VendorListCatalogVariantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VendorListCatalogVariantResponse)
	err := c.cc.Invoke(ctx, VendorCatalogAPI_VendorListCatalogVariant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorCatalogAPIClient) VendorGetCatalogStockStatus(ctx context.Context, in *VendorListCatalogVariantRequest, opts ...grpc.CallOption) (*VendorListCatalogVariantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VendorListCatalogVariantResponse)
	err := c.cc.Invoke(ctx, VendorCatalogAPI_VendorGetCatalogStockStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendorCatalogAPIServer is the server API for VendorCatalogAPI service.
// All implementations should embed UnimplementedVendorCatalogAPIServer
// for forward compatibility.
//
// VendorCatalogAPI provide an abstraction to all of read data access for catalog businesses
type VendorCatalogAPIServer interface {
	VendorListCatalog(context.Context, *VendorListCatalogRequest) (*VendorListCatalogResponse, error)
	VendorListCatalogVariant(context.Context, *VendorListCatalogVariantRequest) (*VendorListCatalogVariantResponse, error)
	VendorGetCatalogStockStatus(context.Context, *VendorListCatalogVariantRequest) (*VendorListCatalogVariantResponse, error)
}

// UnimplementedVendorCatalogAPIServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVendorCatalogAPIServer struct{}

func (UnimplementedVendorCatalogAPIServer) VendorListCatalog(context.Context, *VendorListCatalogRequest) (*VendorListCatalogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorListCatalog not implemented")
}
func (UnimplementedVendorCatalogAPIServer) VendorListCatalogVariant(context.Context, *VendorListCatalogVariantRequest) (*VendorListCatalogVariantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorListCatalogVariant not implemented")
}
func (UnimplementedVendorCatalogAPIServer) VendorGetCatalogStockStatus(context.Context, *VendorListCatalogVariantRequest) (*VendorListCatalogVariantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorGetCatalogStockStatus not implemented")
}
func (UnimplementedVendorCatalogAPIServer) testEmbeddedByValue() {}

// UnsafeVendorCatalogAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendorCatalogAPIServer will
// result in compilation errors.
type UnsafeVendorCatalogAPIServer interface {
	mustEmbedUnimplementedVendorCatalogAPIServer()
}

func RegisterVendorCatalogAPIServer(s grpc.ServiceRegistrar, srv VendorCatalogAPIServer) {
	// If the following call pancis, it indicates UnimplementedVendorCatalogAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VendorCatalogAPI_ServiceDesc, srv)
}

func _VendorCatalogAPI_VendorListCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VendorListCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorCatalogAPIServer).VendorListCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorCatalogAPI_VendorListCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorCatalogAPIServer).VendorListCatalog(ctx, req.(*VendorListCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorCatalogAPI_VendorListCatalogVariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VendorListCatalogVariantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorCatalogAPIServer).VendorListCatalogVariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorCatalogAPI_VendorListCatalogVariant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorCatalogAPIServer).VendorListCatalogVariant(ctx, req.(*VendorListCatalogVariantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorCatalogAPI_VendorGetCatalogStockStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VendorListCatalogVariantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorCatalogAPIServer).VendorGetCatalogStockStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorCatalogAPI_VendorGetCatalogStockStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorCatalogAPIServer).VendorGetCatalogStockStatus(ctx, req.(*VendorListCatalogVariantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VendorCatalogAPI_ServiceDesc is the grpc.ServiceDesc for VendorCatalogAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendorCatalogAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.catalog.v1.VendorCatalogAPI",
	HandlerType: (*VendorCatalogAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VendorListCatalog",
			Handler:    _VendorCatalogAPI_VendorListCatalog_Handler,
		},
		{
			MethodName: "VendorListCatalogVariant",
			Handler:    _VendorCatalogAPI_VendorListCatalogVariant_Handler,
		},
		{
			MethodName: "VendorGetCatalogStockStatus",
			Handler:    _VendorCatalogAPI_VendorGetCatalogStockStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/catalog/v1/api_vendor_catalog.proto",
}

const (
	LegacyVendorCatalogAPI_LegacyVendorListCatalogVariant_FullMethodName     = "/api.catalog.v1.LegacyVendorCatalogAPI/LegacyVendorListCatalogVariant"
	LegacyVendorCatalogAPI_LegacyVendorListCatalogStockStatus_FullMethodName = "/api.catalog.v1.LegacyVendorCatalogAPI/LegacyVendorListCatalogStockStatus"
)

// LegacyVendorCatalogAPIClient is the client API for LegacyVendorCatalogAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// LegacyVendorCatalogAPI provide an abstraction to support the legacy api call format
type LegacyVendorCatalogAPIClient interface {
	LegacyVendorListCatalogVariant(ctx context.Context, in *LegacyVendorListCatalogVariantRequest, opts ...grpc.CallOption) (*LegacyVendorListCatalogVariantResponse, error)
	LegacyVendorListCatalogStockStatus(ctx context.Context, in *LegacyVendorListCatalogStockStatusRequest, opts ...grpc.CallOption) (*LegacyVendorListCatalogStockStatusResponse, error)
}

type legacyVendorCatalogAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewLegacyVendorCatalogAPIClient(cc grpc.ClientConnInterface) LegacyVendorCatalogAPIClient {
	return &legacyVendorCatalogAPIClient{cc}
}

func (c *legacyVendorCatalogAPIClient) LegacyVendorListCatalogVariant(ctx context.Context, in *LegacyVendorListCatalogVariantRequest, opts ...grpc.CallOption) (*LegacyVendorListCatalogVariantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LegacyVendorListCatalogVariantResponse)
	err := c.cc.Invoke(ctx, LegacyVendorCatalogAPI_LegacyVendorListCatalogVariant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *legacyVendorCatalogAPIClient) LegacyVendorListCatalogStockStatus(ctx context.Context, in *LegacyVendorListCatalogStockStatusRequest, opts ...grpc.CallOption) (*LegacyVendorListCatalogStockStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LegacyVendorListCatalogStockStatusResponse)
	err := c.cc.Invoke(ctx, LegacyVendorCatalogAPI_LegacyVendorListCatalogStockStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LegacyVendorCatalogAPIServer is the server API for LegacyVendorCatalogAPI service.
// All implementations should embed UnimplementedLegacyVendorCatalogAPIServer
// for forward compatibility.
//
// LegacyVendorCatalogAPI provide an abstraction to support the legacy api call format
type LegacyVendorCatalogAPIServer interface {
	LegacyVendorListCatalogVariant(context.Context, *LegacyVendorListCatalogVariantRequest) (*LegacyVendorListCatalogVariantResponse, error)
	LegacyVendorListCatalogStockStatus(context.Context, *LegacyVendorListCatalogStockStatusRequest) (*LegacyVendorListCatalogStockStatusResponse, error)
}

// UnimplementedLegacyVendorCatalogAPIServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLegacyVendorCatalogAPIServer struct{}

func (UnimplementedLegacyVendorCatalogAPIServer) LegacyVendorListCatalogVariant(context.Context, *LegacyVendorListCatalogVariantRequest) (*LegacyVendorListCatalogVariantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LegacyVendorListCatalogVariant not implemented")
}
func (UnimplementedLegacyVendorCatalogAPIServer) LegacyVendorListCatalogStockStatus(context.Context, *LegacyVendorListCatalogStockStatusRequest) (*LegacyVendorListCatalogStockStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LegacyVendorListCatalogStockStatus not implemented")
}
func (UnimplementedLegacyVendorCatalogAPIServer) testEmbeddedByValue() {}

// UnsafeLegacyVendorCatalogAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LegacyVendorCatalogAPIServer will
// result in compilation errors.
type UnsafeLegacyVendorCatalogAPIServer interface {
	mustEmbedUnimplementedLegacyVendorCatalogAPIServer()
}

func RegisterLegacyVendorCatalogAPIServer(s grpc.ServiceRegistrar, srv LegacyVendorCatalogAPIServer) {
	// If the following call pancis, it indicates UnimplementedLegacyVendorCatalogAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LegacyVendorCatalogAPI_ServiceDesc, srv)
}

func _LegacyVendorCatalogAPI_LegacyVendorListCatalogVariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LegacyVendorListCatalogVariantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LegacyVendorCatalogAPIServer).LegacyVendorListCatalogVariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LegacyVendorCatalogAPI_LegacyVendorListCatalogVariant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LegacyVendorCatalogAPIServer).LegacyVendorListCatalogVariant(ctx, req.(*LegacyVendorListCatalogVariantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LegacyVendorCatalogAPI_LegacyVendorListCatalogStockStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LegacyVendorListCatalogStockStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LegacyVendorCatalogAPIServer).LegacyVendorListCatalogStockStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LegacyVendorCatalogAPI_LegacyVendorListCatalogStockStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LegacyVendorCatalogAPIServer).LegacyVendorListCatalogStockStatus(ctx, req.(*LegacyVendorListCatalogStockStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LegacyVendorCatalogAPI_ServiceDesc is the grpc.ServiceDesc for LegacyVendorCatalogAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LegacyVendorCatalogAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.catalog.v1.LegacyVendorCatalogAPI",
	HandlerType: (*LegacyVendorCatalogAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LegacyVendorListCatalogVariant",
			Handler:    _LegacyVendorCatalogAPI_LegacyVendorListCatalogVariant_Handler,
		},
		{
			MethodName: "LegacyVendorListCatalogStockStatus",
			Handler:    _LegacyVendorCatalogAPI_LegacyVendorListCatalogStockStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/catalog/v1/api_vendor_catalog.proto",
}
