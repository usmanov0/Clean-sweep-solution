// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error)
	AddItems(ctx context.Context, in *AddItemsRequest, opts ...grpc.CallOption) (*AddItemsResponse, error)
	GetCartWithItems(ctx context.Context, in *GetCartWithItemsRequest, opts ...grpc.CallOption) (*GetCartWithItemsResponse, error)
	GetActiveCarts(ctx context.Context, in *GetActiveCartsRequest, opts ...grpc.CallOption) (*GetActiveCartsResponse, error)
	GetAllItems(ctx context.Context, in *GetAllItemsRequest, opts ...grpc.CallOption) (*GetAllItemsResponse, error)
	UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*UpdateCartResponse, error)
	MarkCartAsPurchased(ctx context.Context, in *MarkCartAsPurchasedRequest, opts ...grpc.CallOption) (*MarkCartAsPurchasedResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error) {
	out := new(CreateCartResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddItems(ctx context.Context, in *AddItemsRequest, opts ...grpc.CallOption) (*AddItemsResponse, error) {
	out := new(AddItemsResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/AddItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartWithItems(ctx context.Context, in *GetCartWithItemsRequest, opts ...grpc.CallOption) (*GetCartWithItemsResponse, error) {
	out := new(GetCartWithItemsResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/GetCartWithItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetActiveCarts(ctx context.Context, in *GetActiveCartsRequest, opts ...grpc.CallOption) (*GetActiveCartsResponse, error) {
	out := new(GetActiveCartsResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/GetActiveCarts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetAllItems(ctx context.Context, in *GetAllItemsRequest, opts ...grpc.CallOption) (*GetAllItemsResponse, error) {
	out := new(GetAllItemsResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*UpdateCartResponse, error) {
	out := new(UpdateCartResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/UpdateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) MarkCartAsPurchased(ctx context.Context, in *MarkCartAsPurchasedRequest, opts ...grpc.CallOption) (*MarkCartAsPurchasedResponse, error) {
	out := new(MarkCartAsPurchasedResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/MarkCartAsPurchased", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.CartService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error)
	AddItems(context.Context, *AddItemsRequest) (*AddItemsResponse, error)
	GetCartWithItems(context.Context, *GetCartWithItemsRequest) (*GetCartWithItemsResponse, error)
	GetActiveCarts(context.Context, *GetActiveCartsRequest) (*GetActiveCartsResponse, error)
	GetAllItems(context.Context, *GetAllItemsRequest) (*GetAllItemsResponse, error)
	UpdateCart(context.Context, *UpdateCartRequest) (*UpdateCartResponse, error)
	MarkCartAsPurchased(context.Context, *MarkCartAsPurchasedRequest) (*MarkCartAsPurchasedResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServiceServer) AddItems(context.Context, *AddItemsRequest) (*AddItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItems not implemented")
}
func (UnimplementedCartServiceServer) GetCartWithItems(context.Context, *GetCartWithItemsRequest) (*GetCartWithItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartWithItems not implemented")
}
func (UnimplementedCartServiceServer) GetActiveCarts(context.Context, *GetActiveCartsRequest) (*GetActiveCartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveCarts not implemented")
}
func (UnimplementedCartServiceServer) GetAllItems(context.Context, *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (UnimplementedCartServiceServer) UpdateCart(context.Context, *UpdateCartRequest) (*UpdateCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCartServiceServer) MarkCartAsPurchased(context.Context, *MarkCartAsPurchasedRequest) (*MarkCartAsPurchasedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkCartAsPurchased not implemented")
}
func (UnimplementedCartServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/AddItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItems(ctx, req.(*AddItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartWithItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartWithItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartWithItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/GetCartWithItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartWithItems(ctx, req.(*GetCartWithItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetActiveCarts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveCartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetActiveCarts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/GetActiveCarts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetActiveCarts(ctx, req.(*GetActiveCartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetAllItems(ctx, req.(*GetAllItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/UpdateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateCart(ctx, req.(*UpdateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_MarkCartAsPurchased_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkCartAsPurchasedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).MarkCartAsPurchased(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/MarkCartAsPurchased",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).MarkCartAsPurchased(ctx, req.(*MarkCartAsPurchasedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
		{
			MethodName: "AddItems",
			Handler:    _CartService_AddItems_Handler,
		},
		{
			MethodName: "GetCartWithItems",
			Handler:    _CartService_GetCartWithItems_Handler,
		},
		{
			MethodName: "GetActiveCarts",
			Handler:    _CartService_GetActiveCarts_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _CartService_GetAllItems_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _CartService_UpdateCart_Handler,
		},
		{
			MethodName: "MarkCartAsPurchased",
			Handler:    _CartService_MarkCartAsPurchased_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CartService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cart.proto",
}
