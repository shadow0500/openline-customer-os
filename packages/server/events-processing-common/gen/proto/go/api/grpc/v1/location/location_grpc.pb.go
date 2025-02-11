// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/location.proto

package location_grpc_service

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

// LocationGrpcServiceClient is the client API for LocationGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationGrpcServiceClient interface {
	UpsertLocation(ctx context.Context, in *UpsertLocationGrpcRequest, opts ...grpc.CallOption) (*LocationIdGrpcResponse, error)
}

type locationGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationGrpcServiceClient(cc grpc.ClientConnInterface) LocationGrpcServiceClient {
	return &locationGrpcServiceClient{cc}
}

func (c *locationGrpcServiceClient) UpsertLocation(ctx context.Context, in *UpsertLocationGrpcRequest, opts ...grpc.CallOption) (*LocationIdGrpcResponse, error) {
	out := new(LocationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/LocationGrpcService/UpsertLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationGrpcServiceServer is the server API for LocationGrpcService service.
// All implementations should embed UnimplementedLocationGrpcServiceServer
// for forward compatibility
type LocationGrpcServiceServer interface {
	UpsertLocation(context.Context, *UpsertLocationGrpcRequest) (*LocationIdGrpcResponse, error)
}

// UnimplementedLocationGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLocationGrpcServiceServer struct {
}

func (UnimplementedLocationGrpcServiceServer) UpsertLocation(context.Context, *UpsertLocationGrpcRequest) (*LocationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertLocation not implemented")
}

// UnsafeLocationGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationGrpcServiceServer will
// result in compilation errors.
type UnsafeLocationGrpcServiceServer interface {
	mustEmbedUnimplementedLocationGrpcServiceServer()
}

func RegisterLocationGrpcServiceServer(s grpc.ServiceRegistrar, srv LocationGrpcServiceServer) {
	s.RegisterService(&LocationGrpcService_ServiceDesc, srv)
}

func _LocationGrpcService_UpsertLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertLocationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationGrpcServiceServer).UpsertLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationGrpcService/UpsertLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationGrpcServiceServer).UpsertLocation(ctx, req.(*UpsertLocationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationGrpcService_ServiceDesc is the grpc.ServiceDesc for LocationGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LocationGrpcService",
	HandlerType: (*LocationGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertLocation",
			Handler:    _LocationGrpcService_UpsertLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/location.proto",
}
