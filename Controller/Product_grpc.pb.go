// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: Product.proto

package Controller

import (
	Helper "GolangMongoDBGRPCSimpleCRUD/Controller/Helper"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductControllerClient is the client API for ProductController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductControllerClient interface {
	GetProducts(ctx context.Context, in *Helper.RequestOnAllData, opts ...grpc.CallOption) (*Helper.Response, error)
	GetProduct(ctx context.Context, in *Helper.ID, opts ...grpc.CallOption) (*Helper.Response, error)
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Helper.Response, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Helper.Response, error)
	DeleteProduct(ctx context.Context, in *Helper.ID, opts ...grpc.CallOption) (*Helper.Response, error)
	DeleteProductPermanently(ctx context.Context, in *Helper.ID, opts ...grpc.CallOption) (*Helper.Response, error)
}

type productControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewProductControllerClient(cc grpc.ClientConnInterface) ProductControllerClient {
	return &productControllerClient{cc}
}

func (c *productControllerClient) GetProducts(ctx context.Context, in *Helper.RequestOnAllData, opts ...grpc.CallOption) (*Helper.Response, error) {
	out := new(Helper.Response)
	err := c.cc.Invoke(ctx, "/GolangMongoDBGRPCSimpleCRUD.ProductController/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productControllerClient) GetProduct(ctx context.Context, in *Helper.ID, opts ...grpc.CallOption) (*Helper.Response, error) {
	out := new(Helper.Response)
	err := c.cc.Invoke(ctx, "/GolangMongoDBGRPCSimpleCRUD.ProductController/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productControllerClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Helper.Response, error) {
	out := new(Helper.Response)
	err := c.cc.Invoke(ctx, "/GolangMongoDBGRPCSimpleCRUD.ProductController/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productControllerClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Helper.Response, error) {
	out := new(Helper.Response)
	err := c.cc.Invoke(ctx, "/GolangMongoDBGRPCSimpleCRUD.ProductController/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productControllerClient) DeleteProduct(ctx context.Context, in *Helper.ID, opts ...grpc.CallOption) (*Helper.Response, error) {
	out := new(Helper.Response)
	err := c.cc.Invoke(ctx, "/GolangMongoDBGRPCSimpleCRUD.ProductController/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productControllerClient) DeleteProductPermanently(ctx context.Context, in *Helper.ID, opts ...grpc.CallOption) (*Helper.Response, error) {
	out := new(Helper.Response)
	err := c.cc.Invoke(ctx, "/GolangMongoDBGRPCSimpleCRUD.ProductController/DeleteProductPermanently", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductControllerServer is the server API for ProductController service.
// All implementations must embed UnimplementedProductControllerServer
// for forward compatibility
type ProductControllerServer interface {
	GetProducts(context.Context, *Helper.RequestOnAllData) (*Helper.Response, error)
	GetProduct(context.Context, *Helper.ID) (*Helper.Response, error)
	CreateProduct(context.Context, *Product) (*Helper.Response, error)
	UpdateProduct(context.Context, *Product) (*Helper.Response, error)
	DeleteProduct(context.Context, *Helper.ID) (*Helper.Response, error)
	DeleteProductPermanently(context.Context, *Helper.ID) (*Helper.Response, error)
	mustEmbedUnimplementedProductControllerServer()
}

// UnimplementedProductControllerServer must be embedded to have forward compatible implementations.
type UnimplementedProductControllerServer struct {
}

func (UnimplementedProductControllerServer) GetProducts(context.Context, *Helper.RequestOnAllData) (*Helper.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductControllerServer) GetProduct(context.Context, *Helper.ID) (*Helper.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductControllerServer) CreateProduct(context.Context, *Product) (*Helper.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductControllerServer) UpdateProduct(context.Context, *Product) (*Helper.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductControllerServer) DeleteProduct(context.Context, *Helper.ID) (*Helper.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductControllerServer) DeleteProductPermanently(context.Context, *Helper.ID) (*Helper.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductPermanently not implemented")
}
func (UnimplementedProductControllerServer) mustEmbedUnimplementedProductControllerServer() {}

// UnsafeProductControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductControllerServer will
// result in compilation errors.
type UnsafeProductControllerServer interface {
	mustEmbedUnimplementedProductControllerServer()
}

func RegisterProductControllerServer(s grpc.ServiceRegistrar, srv ProductControllerServer) {
	s.RegisterService(&ProductController_ServiceDesc, srv)
}

func _ProductController_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Helper.RequestOnAllData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductControllerServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GolangMongoDBGRPCSimpleCRUD.ProductController/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductControllerServer).GetProducts(ctx, req.(*Helper.RequestOnAllData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductController_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Helper.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductControllerServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GolangMongoDBGRPCSimpleCRUD.ProductController/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductControllerServer).GetProduct(ctx, req.(*Helper.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductController_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductControllerServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GolangMongoDBGRPCSimpleCRUD.ProductController/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductControllerServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductController_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductControllerServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GolangMongoDBGRPCSimpleCRUD.ProductController/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductControllerServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductController_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Helper.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductControllerServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GolangMongoDBGRPCSimpleCRUD.ProductController/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductControllerServer).DeleteProduct(ctx, req.(*Helper.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductController_DeleteProductPermanently_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Helper.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductControllerServer).DeleteProductPermanently(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GolangMongoDBGRPCSimpleCRUD.ProductController/DeleteProductPermanently",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductControllerServer).DeleteProductPermanently(ctx, req.(*Helper.ID))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductController_ServiceDesc is the grpc.ServiceDesc for ProductController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GolangMongoDBGRPCSimpleCRUD.ProductController",
	HandlerType: (*ProductControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductController_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductController_GetProduct_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ProductController_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductController_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductController_DeleteProduct_Handler,
		},
		{
			MethodName: "DeleteProductPermanently",
			Handler:    _ProductController_DeleteProductPermanently_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Product.proto",
}
