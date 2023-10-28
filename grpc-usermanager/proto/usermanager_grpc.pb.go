// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/usermanager.proto

package grpc_usermanager

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

// UsermanagerServiceClient is the client API for UsermanagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsermanagerServiceClient interface {
	GetUser(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type usermanagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsermanagerServiceClient(cc grpc.ClientConnInterface) UsermanagerServiceClient {
	return &usermanagerServiceClient{cc}
}

func (c *usermanagerServiceClient) GetUser(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/proto.UsermanagerService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsermanagerServiceServer is the server API for UsermanagerService service.
// All implementations must embed UnimplementedUsermanagerServiceServer
// for forward compatibility
type UsermanagerServiceServer interface {
	GetUser(context.Context, *Input) (*Output, error)
	mustEmbedUnimplementedUsermanagerServiceServer()
}

// UnimplementedUsermanagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsermanagerServiceServer struct {
}

func (UnimplementedUsermanagerServiceServer) GetUser(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsermanagerServiceServer) mustEmbedUnimplementedUsermanagerServiceServer() {}

// UnsafeUsermanagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsermanagerServiceServer will
// result in compilation errors.
type UnsafeUsermanagerServiceServer interface {
	mustEmbedUnimplementedUsermanagerServiceServer()
}

func RegisterUsermanagerServiceServer(s grpc.ServiceRegistrar, srv UsermanagerServiceServer) {
	s.RegisterService(&UsermanagerService_ServiceDesc, srv)
}

func _UsermanagerService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsermanagerServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UsermanagerService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsermanagerServiceServer).GetUser(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// UsermanagerService_ServiceDesc is the grpc.ServiceDesc for UsermanagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsermanagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UsermanagerService",
	HandlerType: (*UsermanagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UsermanagerService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/usermanager.proto",
}
