// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package identityv1

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

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/identity.v1.IdentityService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/identity.v1.IdentityService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility
type IdentityServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedIdentityServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.v1.IdentityService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.v1.IdentityService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.v1.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _IdentityService_CreateAccount_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _IdentityService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/v1/api.proto",
}
