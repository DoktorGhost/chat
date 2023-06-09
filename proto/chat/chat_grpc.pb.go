// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: chat.proto

package chat

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
	ChatService_AddMessage_FullMethodName  = "/chat.ChatService/AddMessage"
	ChatService_GetMessages_FullMethodName = "/chat.ChatService/GetMessages"
	ChatService_GetChats_FullMethodName    = "/chat.ChatService/GetChats"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	AddMessage(ctx context.Context, in *AddMessageRequest, opts ...grpc.CallOption) (*AddMessageResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	GetChats(ctx context.Context, in *GetChatsRequest, opts ...grpc.CallOption) (*GetChatsResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) AddMessage(ctx context.Context, in *AddMessageRequest, opts ...grpc.CallOption) (*AddMessageResponse, error) {
	out := new(AddMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_AddMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, ChatService_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChats(ctx context.Context, in *GetChatsRequest, opts ...grpc.CallOption) (*GetChatsResponse, error) {
	out := new(GetChatsResponse)
	err := c.cc.Invoke(ctx, ChatService_GetChats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	AddMessage(context.Context, *AddMessageRequest) (*AddMessageResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	GetChats(context.Context, *GetChatsRequest) (*GetChatsResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) AddMessage(context.Context, *AddMessageRequest) (*AddMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessage not implemented")
}
func (UnimplementedChatServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatServiceServer) GetChats(context.Context, *GetChatsRequest) (*GetChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_AddMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMessage(ctx, req.(*AddMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChats(ctx, req.(*GetChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMessage",
			Handler:    _ChatService_AddMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _ChatService_GetMessages_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _ChatService_GetChats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
