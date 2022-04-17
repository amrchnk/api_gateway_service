// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.2.0
// source: account/account.proto

package account

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	//ACCOUNT
	CreateAccountByUserId(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	GetAccountByUserId(ctx context.Context, in *GetAccountByUserIdRequest, opts ...grpc.CallOption) (*GetAccountByUserIdResponse, error)
	DeleteAccountByUserId(ctx context.Context, in *DeleteAccountByUserIdRequest, opts ...grpc.CallOption) (*DeleteAccountByUserIdResponse, error)
	//POST
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	DeletePostById(ctx context.Context, in *DeletePostByIdRequest, opts ...grpc.CallOption) (*DeletePostByIdResponse, error)
	GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error)
	GetPostsByAccountId(ctx context.Context, in *GetUserPostsRequest, opts ...grpc.CallOption) (*GetUserPostsResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccountByUserId(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/CreateAccountByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByUserId(ctx context.Context, in *GetAccountByUserIdRequest, opts ...grpc.CallOption) (*GetAccountByUserIdResponse, error) {
	out := new(GetAccountByUserIdResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/GetAccountByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccountByUserId(ctx context.Context, in *DeleteAccountByUserIdRequest, opts ...grpc.CallOption) (*DeleteAccountByUserIdResponse, error) {
	out := new(DeleteAccountByUserIdResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/DeleteAccountByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeletePostById(ctx context.Context, in *DeletePostByIdRequest, opts ...grpc.CallOption) (*DeletePostByIdResponse, error) {
	out := new(DeletePostByIdResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/DeletePostById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error) {
	out := new(GetPostByIdResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/GetPostById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetPostsByAccountId(ctx context.Context, in *GetUserPostsRequest, opts ...grpc.CallOption) (*GetUserPostsResponse, error) {
	out := new(GetUserPostsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountService/GetPostsByAccountId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	//ACCOUNT
	CreateAccountByUserId(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	GetAccountByUserId(context.Context, *GetAccountByUserIdRequest) (*GetAccountByUserIdResponse, error)
	DeleteAccountByUserId(context.Context, *DeleteAccountByUserIdRequest) (*DeleteAccountByUserIdResponse, error)
	//POST
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	DeletePostById(context.Context, *DeletePostByIdRequest) (*DeletePostByIdResponse, error)
	GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error)
	GetPostsByAccountId(context.Context, *GetUserPostsRequest) (*GetUserPostsResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccountByUserId(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountByUserId not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByUserId(context.Context, *GetAccountByUserIdRequest) (*GetAccountByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByUserId not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccountByUserId(context.Context, *DeleteAccountByUserIdRequest) (*DeleteAccountByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccountByUserId not implemented")
}
func (UnimplementedAccountServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedAccountServiceServer) DeletePostById(context.Context, *DeletePostByIdRequest) (*DeletePostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostById not implemented")
}
func (UnimplementedAccountServiceServer) GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedAccountServiceServer) GetPostsByAccountId(context.Context, *GetUserPostsRequest) (*GetUserPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByAccountId not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccountByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccountByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/CreateAccountByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccountByUserId(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/GetAccountByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByUserId(ctx, req.(*GetAccountByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccountByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccountByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/DeleteAccountByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccountByUserId(ctx, req.(*DeleteAccountByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeletePostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeletePostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/DeletePostById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeletePostById(ctx, req.(*DeletePostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/GetPostById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetPostById(ctx, req.(*GetPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetPostsByAccountId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetPostsByAccountId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountService/GetPostsByAccountId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetPostsByAccountId(ctx, req.(*GetUserPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccountByUserId",
			Handler:    _AccountService_CreateAccountByUserId_Handler,
		},
		{
			MethodName: "GetAccountByUserId",
			Handler:    _AccountService_GetAccountByUserId_Handler,
		},
		{
			MethodName: "DeleteAccountByUserId",
			Handler:    _AccountService_DeleteAccountByUserId_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _AccountService_CreatePost_Handler,
		},
		{
			MethodName: "DeletePostById",
			Handler:    _AccountService_DeletePostById_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _AccountService_GetPostById_Handler,
		},
		{
			MethodName: "GetPostsByAccountId",
			Handler:    _AccountService_GetPostsByAccountId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account.proto",
}
