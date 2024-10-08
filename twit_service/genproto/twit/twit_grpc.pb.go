// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: twit.proto

package twit

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

const (
	TwitService_Create_FullMethodName  = "/twit.TwitService/Create"
	TwitService_GetById_FullMethodName = "/twit.TwitService/GetById"
	TwitService_GetAll_FullMethodName  = "/twit.TwitService/GetAll"
	TwitService_Update_FullMethodName  = "/twit.TwitService/Update"
	TwitService_Delete_FullMethodName  = "/twit.TwitService/Delete"
)

// TwitServiceClient is the client API for TwitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitServiceClient interface {
	Create(ctx context.Context, in *TwitCreateReq, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TwitRes, error)
	GetAll(ctx context.Context, in *TwitGetAllReq, opts ...grpc.CallOption) (*TwitGetAllRes, error)
	Update(ctx context.Context, in *TwitUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *TwitDeleteReq, opts ...grpc.CallOption) (*Void, error)
}

type twitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitServiceClient(cc grpc.ClientConnInterface) TwitServiceClient {
	return &twitServiceClient{cc}
}

func (c *twitServiceClient) Create(ctx context.Context, in *TwitCreateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, TwitService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TwitRes, error) {
	out := new(TwitRes)
	err := c.cc.Invoke(ctx, TwitService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClient) GetAll(ctx context.Context, in *TwitGetAllReq, opts ...grpc.CallOption) (*TwitGetAllRes, error) {
	out := new(TwitGetAllRes)
	err := c.cc.Invoke(ctx, TwitService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClient) Update(ctx context.Context, in *TwitUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, TwitService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClient) Delete(ctx context.Context, in *TwitDeleteReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, TwitService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitServiceServer is the server API for TwitService service.
// All implementations must embed UnimplementedTwitServiceServer
// for forward compatibility
type TwitServiceServer interface {
	Create(context.Context, *TwitCreateReq) (*Void, error)
	GetById(context.Context, *ById) (*TwitRes, error)
	GetAll(context.Context, *TwitGetAllReq) (*TwitGetAllRes, error)
	Update(context.Context, *TwitUpdateReq) (*Void, error)
	Delete(context.Context, *TwitDeleteReq) (*Void, error)
	mustEmbedUnimplementedTwitServiceServer()
}

// UnimplementedTwitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTwitServiceServer struct {
}

func (UnimplementedTwitServiceServer) Create(context.Context, *TwitCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTwitServiceServer) GetById(context.Context, *ById) (*TwitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTwitServiceServer) GetAll(context.Context, *TwitGetAllReq) (*TwitGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTwitServiceServer) Update(context.Context, *TwitUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTwitServiceServer) Delete(context.Context, *TwitDeleteReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTwitServiceServer) mustEmbedUnimplementedTwitServiceServer() {}

// UnsafeTwitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitServiceServer will
// result in compilation errors.
type UnsafeTwitServiceServer interface {
	mustEmbedUnimplementedTwitServiceServer()
}

func RegisterTwitServiceServer(s grpc.ServiceRegistrar, srv TwitServiceServer) {
	s.RegisterService(&TwitService_ServiceDesc, srv)
}

func _TwitService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).Create(ctx, req.(*TwitCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).GetAll(ctx, req.(*TwitGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).Update(ctx, req.(*TwitUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceServer).Delete(ctx, req.(*TwitDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TwitService_ServiceDesc is the grpc.ServiceDesc for TwitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "twit.TwitService",
	HandlerType: (*TwitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TwitService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _TwitService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TwitService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TwitService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TwitService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "twit.proto",
}
