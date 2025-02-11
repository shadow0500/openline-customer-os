// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: MessageStore.proto

package generated

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

// MessageStoreServiceClient is the client API for MessageStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageStoreServiceClient interface {
	GetFeeds(ctx context.Context, in *GetFeedsPagedRequest, opts ...grpc.CallOption) (*FeedItemPagedResponse, error)
	GetFeed(ctx context.Context, in *FeedId, opts ...grpc.CallOption) (*FeedItem, error)
	GetMessagesForFeed(ctx context.Context, in *PagedMessages, opts ...grpc.CallOption) (*MessageListResponse, error)
	GetMessage(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*Message, error)
	SaveMessage(ctx context.Context, in *InputMessage, opts ...grpc.CallOption) (*MessageId, error)
	GetParticipants(ctx context.Context, in *FeedId, opts ...grpc.CallOption) (*ParticipantsListResponse, error)
	GetParticipantIds(ctx context.Context, in *FeedId, opts ...grpc.CallOption) (*ParticipantObjectListResponse, error)
}

type messageStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageStoreServiceClient(cc grpc.ClientConnInterface) MessageStoreServiceClient {
	return &messageStoreServiceClient{cc}
}

func (c *messageStoreServiceClient) GetFeeds(ctx context.Context, in *GetFeedsPagedRequest, opts ...grpc.CallOption) (*FeedItemPagedResponse, error) {
	out := new(FeedItemPagedResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/getFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetFeed(ctx context.Context, in *FeedId, opts ...grpc.CallOption) (*FeedItem, error) {
	out := new(FeedItem)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/getFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetMessagesForFeed(ctx context.Context, in *PagedMessages, opts ...grpc.CallOption) (*MessageListResponse, error) {
	out := new(MessageListResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/getMessagesForFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetMessage(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/getMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) SaveMessage(ctx context.Context, in *InputMessage, opts ...grpc.CallOption) (*MessageId, error) {
	out := new(MessageId)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/saveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetParticipants(ctx context.Context, in *FeedId, opts ...grpc.CallOption) (*ParticipantsListResponse, error) {
	out := new(ParticipantsListResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/getParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageStoreServiceClient) GetParticipantIds(ctx context.Context, in *FeedId, opts ...grpc.CallOption) (*ParticipantObjectListResponse, error) {
	out := new(ParticipantObjectListResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageStoreService/getParticipantIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageStoreServiceServer is the server API for MessageStoreService service.
// All implementations must embed UnimplementedMessageStoreServiceServer
// for forward compatibility
type MessageStoreServiceServer interface {
	GetFeeds(context.Context, *GetFeedsPagedRequest) (*FeedItemPagedResponse, error)
	GetFeed(context.Context, *FeedId) (*FeedItem, error)
	GetMessagesForFeed(context.Context, *PagedMessages) (*MessageListResponse, error)
	GetMessage(context.Context, *MessageId) (*Message, error)
	SaveMessage(context.Context, *InputMessage) (*MessageId, error)
	GetParticipants(context.Context, *FeedId) (*ParticipantsListResponse, error)
	GetParticipantIds(context.Context, *FeedId) (*ParticipantObjectListResponse, error)
	mustEmbedUnimplementedMessageStoreServiceServer()
}

// UnimplementedMessageStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageStoreServiceServer struct {
}

func (UnimplementedMessageStoreServiceServer) GetFeeds(context.Context, *GetFeedsPagedRequest) (*FeedItemPagedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeds not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetFeed(context.Context, *FeedId) (*FeedItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetMessagesForFeed(context.Context, *PagedMessages) (*MessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesForFeed not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetMessage(context.Context, *MessageId) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessageStoreServiceServer) SaveMessage(context.Context, *InputMessage) (*MessageId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetParticipants(context.Context, *FeedId) (*ParticipantsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipants not implemented")
}
func (UnimplementedMessageStoreServiceServer) GetParticipantIds(context.Context, *FeedId) (*ParticipantObjectListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipantIds not implemented")
}
func (UnimplementedMessageStoreServiceServer) mustEmbedUnimplementedMessageStoreServiceServer() {}

// UnsafeMessageStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageStoreServiceServer will
// result in compilation errors.
type UnsafeMessageStoreServiceServer interface {
	mustEmbedUnimplementedMessageStoreServiceServer()
}

func RegisterMessageStoreServiceServer(s grpc.ServiceRegistrar, srv MessageStoreServiceServer) {
	s.RegisterService(&MessageStoreService_ServiceDesc, srv)
}

func _MessageStoreService_GetFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedsPagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/getFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetFeeds(ctx, req.(*GetFeedsPagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/getFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetFeed(ctx, req.(*FeedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetMessagesForFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagedMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetMessagesForFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/getMessagesForFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetMessagesForFeed(ctx, req.(*PagedMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/getMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetMessage(ctx, req.(*MessageId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_SaveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).SaveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/saveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).SaveMessage(ctx, req.(*InputMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/getParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetParticipants(ctx, req.(*FeedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageStoreService_GetParticipantIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServiceServer).GetParticipantIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageStoreService/getParticipantIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServiceServer).GetParticipantIds(ctx, req.(*FeedId))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageStoreService_ServiceDesc is the grpc.ServiceDesc for MessageStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageStoreService",
	HandlerType: (*MessageStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getFeeds",
			Handler:    _MessageStoreService_GetFeeds_Handler,
		},
		{
			MethodName: "getFeed",
			Handler:    _MessageStoreService_GetFeed_Handler,
		},
		{
			MethodName: "getMessagesForFeed",
			Handler:    _MessageStoreService_GetMessagesForFeed_Handler,
		},
		{
			MethodName: "getMessage",
			Handler:    _MessageStoreService_GetMessage_Handler,
		},
		{
			MethodName: "saveMessage",
			Handler:    _MessageStoreService_SaveMessage_Handler,
		},
		{
			MethodName: "getParticipants",
			Handler:    _MessageStoreService_GetParticipants_Handler,
		},
		{
			MethodName: "getParticipantIds",
			Handler:    _MessageStoreService_GetParticipantIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MessageStore.proto",
}
