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
	VendorCatalogAPI_VendorListCatalog_FullMethodName        = "/api.catalog.v1.VendorCatalogAPI/VendorListCatalog"
	VendorCatalogAPI_VendorListCatalogVariant_FullMethodName = "/api.catalog.v1.VendorCatalogAPI/VendorListCatalogVariant"
)

// VendorCatalogAPIClient is the client API for VendorCatalogAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ProductListingVendorAPI provide an abstraction to all of read and write data access for order businesses
type VendorCatalogAPIClient interface {
	VendorListCatalog(ctx context.Context, in *VendorListCatalogRequest, opts ...grpc.CallOption) (*VendorListCatalogResponse, error)
	VendorListCatalogVariant(ctx context.Context, in *VendorListCatalogVariantRequest, opts ...grpc.CallOption) (*VendorListCatalogVariantResponse, error)
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

// VendorCatalogAPIServer is the server API for VendorCatalogAPI service.
// All implementations should embed UnimplementedVendorCatalogAPIServer
// for forward compatibility.
//
// ProductListingVendorAPI provide an abstraction to all of read and write data access for order businesses
type VendorCatalogAPIServer interface {
	VendorListCatalog(context.Context, *VendorListCatalogRequest) (*VendorListCatalogResponse, error)
	VendorListCatalogVariant(context.Context, *VendorListCatalogVariantRequest) (*VendorListCatalogVariantResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/catalog/v1/api_vendor_catalog.proto",
}
