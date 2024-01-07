// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: services.proto

package pb

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
	PortfolioService_LoginUser_FullMethodName      = "/pb.PortfolioService/LoginUser"
	PortfolioService_LogoutUser_FullMethodName     = "/pb.PortfolioService/LogoutUser"
	PortfolioService_RefreshToken_FullMethodName   = "/pb.PortfolioService/RefreshToken"
	PortfolioService_CreateUser_FullMethodName     = "/pb.PortfolioService/CreateUser"
	PortfolioService_UpdateUser_FullMethodName     = "/pb.PortfolioService/UpdateUser"
	PortfolioService_GetUser_FullMethodName        = "/pb.PortfolioService/GetUser"
	PortfolioService_ResetPassword_FullMethodName  = "/pb.PortfolioService/ResetPassword"
	PortfolioService_ChangePassword_FullMethodName = "/pb.PortfolioService/ChangePassword"
	PortfolioService_CreateArt_FullMethodName      = "/pb.PortfolioService/CreateArt"
	PortfolioService_UpdateArt_FullMethodName      = "/pb.PortfolioService/UpdateArt"
	PortfolioService_GetArt_FullMethodName         = "/pb.PortfolioService/GetArt"
	PortfolioService_ListArts_FullMethodName       = "/pb.PortfolioService/ListArts"
	PortfolioService_DeleteArt_FullMethodName      = "/pb.PortfolioService/DeleteArt"
)

// PortfolioServiceClient is the client API for PortfolioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortfolioServiceClient interface {
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LogoutUser(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	CreateArt(ctx context.Context, in *CreateArtRequest, opts ...grpc.CallOption) (*CreateArtResponse, error)
	UpdateArt(ctx context.Context, in *UpdateArtRequest, opts ...grpc.CallOption) (*UpdateArtResponse, error)
	GetArt(ctx context.Context, in *GetArtRequest, opts ...grpc.CallOption) (*GetArtResponse, error)
	ListArts(ctx context.Context, in *ListArtRequest, opts ...grpc.CallOption) (*ListArtResponse, error)
	DeleteArt(ctx context.Context, in *DeleteArtRequest, opts ...grpc.CallOption) (*DeleteArtResponse, error)
}

type portfolioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortfolioServiceClient(cc grpc.ClientConnInterface) PortfolioServiceClient {
	return &portfolioServiceClient{cc}
}

func (c *portfolioServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PortfolioService_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) LogoutUser(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, PortfolioService_LogoutUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, PortfolioService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, PortfolioService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) CreateArt(ctx context.Context, in *CreateArtRequest, opts ...grpc.CallOption) (*CreateArtResponse, error) {
	out := new(CreateArtResponse)
	err := c.cc.Invoke(ctx, PortfolioService_CreateArt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) UpdateArt(ctx context.Context, in *UpdateArtRequest, opts ...grpc.CallOption) (*UpdateArtResponse, error) {
	out := new(UpdateArtResponse)
	err := c.cc.Invoke(ctx, PortfolioService_UpdateArt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) GetArt(ctx context.Context, in *GetArtRequest, opts ...grpc.CallOption) (*GetArtResponse, error) {
	out := new(GetArtResponse)
	err := c.cc.Invoke(ctx, PortfolioService_GetArt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ListArts(ctx context.Context, in *ListArtRequest, opts ...grpc.CallOption) (*ListArtResponse, error) {
	out := new(ListArtResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ListArts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) DeleteArt(ctx context.Context, in *DeleteArtRequest, opts ...grpc.CallOption) (*DeleteArtResponse, error) {
	out := new(DeleteArtResponse)
	err := c.cc.Invoke(ctx, PortfolioService_DeleteArt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortfolioServiceServer is the server API for PortfolioService service.
// All implementations must embed UnimplementedPortfolioServiceServer
// for forward compatibility
type PortfolioServiceServer interface {
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	CreateArt(context.Context, *CreateArtRequest) (*CreateArtResponse, error)
	UpdateArt(context.Context, *UpdateArtRequest) (*UpdateArtResponse, error)
	GetArt(context.Context, *GetArtRequest) (*GetArtResponse, error)
	ListArts(context.Context, *ListArtRequest) (*ListArtResponse, error)
	DeleteArt(context.Context, *DeleteArtRequest) (*DeleteArtResponse, error)
	mustEmbedUnimplementedPortfolioServiceServer()
}

// UnimplementedPortfolioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortfolioServiceServer struct {
}

func (UnimplementedPortfolioServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedPortfolioServiceServer) LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (UnimplementedPortfolioServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedPortfolioServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedPortfolioServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedPortfolioServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedPortfolioServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedPortfolioServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedPortfolioServiceServer) CreateArt(context.Context, *CreateArtRequest) (*CreateArtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArt not implemented")
}
func (UnimplementedPortfolioServiceServer) UpdateArt(context.Context, *UpdateArtRequest) (*UpdateArtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArt not implemented")
}
func (UnimplementedPortfolioServiceServer) GetArt(context.Context, *GetArtRequest) (*GetArtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArt not implemented")
}
func (UnimplementedPortfolioServiceServer) ListArts(context.Context, *ListArtRequest) (*ListArtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArts not implemented")
}
func (UnimplementedPortfolioServiceServer) DeleteArt(context.Context, *DeleteArtRequest) (*DeleteArtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArt not implemented")
}
func (UnimplementedPortfolioServiceServer) mustEmbedUnimplementedPortfolioServiceServer() {}

// UnsafePortfolioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortfolioServiceServer will
// result in compilation errors.
type UnsafePortfolioServiceServer interface {
	mustEmbedUnimplementedPortfolioServiceServer()
}

func RegisterPortfolioServiceServer(s grpc.ServiceRegistrar, srv PortfolioServiceServer) {
	s.RegisterService(&PortfolioService_ServiceDesc, srv)
}

func _PortfolioService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_LogoutUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).LogoutUser(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_CreateArt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).CreateArt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_CreateArt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).CreateArt(ctx, req.(*CreateArtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_UpdateArt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).UpdateArt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_UpdateArt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).UpdateArt(ctx, req.(*UpdateArtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_GetArt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).GetArt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_GetArt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).GetArt(ctx, req.(*GetArtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ListArts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ListArts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ListArts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ListArts(ctx, req.(*ListArtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_DeleteArt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).DeleteArt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_DeleteArt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).DeleteArt(ctx, req.(*DeleteArtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortfolioService_ServiceDesc is the grpc.ServiceDesc for PortfolioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortfolioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PortfolioService",
	HandlerType: (*PortfolioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _PortfolioService_LoginUser_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _PortfolioService_LogoutUser_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _PortfolioService_RefreshToken_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _PortfolioService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _PortfolioService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _PortfolioService_GetUser_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _PortfolioService_ResetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _PortfolioService_ChangePassword_Handler,
		},
		{
			MethodName: "CreateArt",
			Handler:    _PortfolioService_CreateArt_Handler,
		},
		{
			MethodName: "UpdateArt",
			Handler:    _PortfolioService_UpdateArt_Handler,
		},
		{
			MethodName: "GetArt",
			Handler:    _PortfolioService_GetArt_Handler,
		},
		{
			MethodName: "ListArts",
			Handler:    _PortfolioService_ListArts_Handler,
		},
		{
			MethodName: "DeleteArt",
			Handler:    _PortfolioService_DeleteArt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}