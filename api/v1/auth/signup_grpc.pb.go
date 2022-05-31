// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: v1/auth/signup.proto

package auth

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

// SignupClient is the client API for Signup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignupClient interface {
	UsingEmail(ctx context.Context, in *UsingEmailRequest, opts ...grpc.CallOption) (*UsingEmailReply, error)
	UsingPhone(ctx context.Context, in *UsingPhoneRequest, opts ...grpc.CallOption) (*UsingPhoneReply, error)
}

type signupClient struct {
	cc grpc.ClientConnInterface
}

func NewSignupClient(cc grpc.ClientConnInterface) SignupClient {
	return &signupClient{cc}
}

func (c *signupClient) UsingEmail(ctx context.Context, in *UsingEmailRequest, opts ...grpc.CallOption) (*UsingEmailReply, error) {
	out := new(UsingEmailReply)
	err := c.cc.Invoke(ctx, "/api.v1.auth.Signup/UsingEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupClient) UsingPhone(ctx context.Context, in *UsingPhoneRequest, opts ...grpc.CallOption) (*UsingPhoneReply, error) {
	out := new(UsingPhoneReply)
	err := c.cc.Invoke(ctx, "/api.v1.auth.Signup/UsingPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServer is the server API for Signup service.
// All implementations must embed UnimplementedSignupServer
// for forward compatibility
type SignupServer interface {
	UsingEmail(context.Context, *UsingEmailRequest) (*UsingEmailReply, error)
	UsingPhone(context.Context, *UsingPhoneRequest) (*UsingPhoneReply, error)
	mustEmbedUnimplementedSignupServer()
}

// UnimplementedSignupServer must be embedded to have forward compatible implementations.
type UnimplementedSignupServer struct {
}

func (UnimplementedSignupServer) UsingEmail(context.Context, *UsingEmailRequest) (*UsingEmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsingEmail not implemented")
}
func (UnimplementedSignupServer) UsingPhone(context.Context, *UsingPhoneRequest) (*UsingPhoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsingPhone not implemented")
}
func (UnimplementedSignupServer) mustEmbedUnimplementedSignupServer() {}

// UnsafeSignupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignupServer will
// result in compilation errors.
type UnsafeSignupServer interface {
	mustEmbedUnimplementedSignupServer()
}

func RegisterSignupServer(s grpc.ServiceRegistrar, srv SignupServer) {
	s.RegisterService(&Signup_ServiceDesc, srv)
}

func _Signup_UsingEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsingEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).UsingEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.auth.Signup/UsingEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).UsingEmail(ctx, req.(*UsingEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signup_UsingPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsingPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).UsingPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.auth.Signup/UsingPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).UsingPhone(ctx, req.(*UsingPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Signup_ServiceDesc is the grpc.ServiceDesc for Signup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.auth.Signup",
	HandlerType: (*SignupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UsingEmail",
			Handler:    _Signup_UsingEmail_Handler,
		},
		{
			MethodName: "UsingPhone",
			Handler:    _Signup_UsingPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/auth/signup.proto",
}
