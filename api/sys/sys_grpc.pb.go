// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api/sys/sys.proto

package sys

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

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResp, error)
	CheckResource(ctx context.Context, in *CheckResourceRequest, opts ...grpc.CallOption) (*CheckResourceReply, error)
	CheckRoleLevel(ctx context.Context, in *CheckRoleLevelRequest, opts ...grpc.CallOption) (*CheckRoleLevelReply, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/api.sys.Sys/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) CheckResource(ctx context.Context, in *CheckResourceRequest, opts ...grpc.CallOption) (*CheckResourceReply, error) {
	out := new(CheckResourceReply)
	err := c.cc.Invoke(ctx, "/api.sys.Sys/CheckResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) CheckRoleLevel(ctx context.Context, in *CheckRoleLevelRequest, opts ...grpc.CallOption) (*CheckRoleLevelReply, error) {
	out := new(CheckRoleLevelReply)
	err := c.cc.Invoke(ctx, "/api.sys.Sys/CheckRoleLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResp, error)
	CheckResource(context.Context, *CheckResourceRequest) (*CheckResourceReply, error)
	CheckRoleLevel(context.Context, *CheckRoleLevelRequest) (*CheckRoleLevelReply, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) Auth(context.Context, *AuthRequest) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedSysServer) CheckResource(context.Context, *CheckResourceRequest) (*CheckResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckResource not implemented")
}
func (UnimplementedSysServer) CheckRoleLevel(context.Context, *CheckRoleLevelRequest) (*CheckRoleLevelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRoleLevel not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sys.Sys/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_CheckResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).CheckResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sys.Sys/CheckResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).CheckResource(ctx, req.(*CheckResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_CheckRoleLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoleLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).CheckRoleLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sys.Sys/CheckRoleLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).CheckRoleLevel(ctx, req.(*CheckRoleLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sys.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Sys_Auth_Handler,
		},
		{
			MethodName: "CheckResource",
			Handler:    _Sys_CheckResource_Handler,
		},
		{
			MethodName: "CheckRoleLevel",
			Handler:    _Sys_CheckRoleLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sys/sys.proto",
}