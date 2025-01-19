// Copyright 2021 Woodpecker Authors
// Copyright 2011 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: crow.proto

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Crow_Version_FullMethodName         = "/proto.crow/Version"
	Crow_Next_FullMethodName            = "/proto.crow/Next"
	Crow_Init_FullMethodName            = "/proto.crow/Init"
	Crow_Wait_FullMethodName            = "/proto.crow/Wait"
	Crow_Done_FullMethodName            = "/proto.crow/Done"
	Crow_Extend_FullMethodName          = "/proto.crow/Extend"
	Crow_Update_FullMethodName          = "/proto.crow/Update"
	Crow_Log_FullMethodName             = "/proto.crow/Log"
	Crow_RegisterAgent_FullMethodName   = "/proto.crow/RegisterAgent"
	Crow_UnregisterAgent_FullMethodName = "/proto.crow/UnregisterAgent"
	Crow_ReportHealth_FullMethodName    = "/proto.crow/ReportHealth"
)

// CrowClient is the client API for Crow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Crow Server Service
type CrowClient interface {
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextResponse, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*Empty, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*Empty, error)
	Extend(ctx context.Context, in *ExtendRequest, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error)
	RegisterAgent(ctx context.Context, in *RegisterAgentRequest, opts ...grpc.CallOption) (*RegisterAgentResponse, error)
	UnregisterAgent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ReportHealth(ctx context.Context, in *ReportHealthRequest, opts ...grpc.CallOption) (*Empty, error)
}

type crowClient struct {
	cc grpc.ClientConnInterface
}

func NewCrowClient(cc grpc.ClientConnInterface) CrowClient {
	return &crowClient{cc}
}

func (c *crowClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, Crow_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NextResponse)
	err := c.cc.Invoke(ctx, Crow_Next_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_Wait_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_Done_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Extend(ctx context.Context, in *ExtendRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_Extend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_Log_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) RegisterAgent(ctx context.Context, in *RegisterAgentRequest, opts ...grpc.CallOption) (*RegisterAgentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterAgentResponse)
	err := c.cc.Invoke(ctx, Crow_RegisterAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) UnregisterAgent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_UnregisterAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowClient) ReportHealth(ctx context.Context, in *ReportHealthRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Crow_ReportHealth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrowServer is the server API for Crow service.
// All implementations must embed UnimplementedCrowServer
// for forward compatibility.
//
// Crow Server Service
type CrowServer interface {
	Version(context.Context, *Empty) (*VersionResponse, error)
	Next(context.Context, *NextRequest) (*NextResponse, error)
	Init(context.Context, *InitRequest) (*Empty, error)
	Wait(context.Context, *WaitRequest) (*Empty, error)
	Done(context.Context, *DoneRequest) (*Empty, error)
	Extend(context.Context, *ExtendRequest) (*Empty, error)
	Update(context.Context, *UpdateRequest) (*Empty, error)
	Log(context.Context, *LogRequest) (*Empty, error)
	RegisterAgent(context.Context, *RegisterAgentRequest) (*RegisterAgentResponse, error)
	UnregisterAgent(context.Context, *Empty) (*Empty, error)
	ReportHealth(context.Context, *ReportHealthRequest) (*Empty, error)
	mustEmbedUnimplementedCrowServer()
}

// UnimplementedCrowServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCrowServer struct{}

func (UnimplementedCrowServer) Version(context.Context, *Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedCrowServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedCrowServer) Init(context.Context, *InitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedCrowServer) Wait(context.Context, *WaitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wait not implemented")
}
func (UnimplementedCrowServer) Done(context.Context, *DoneRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Done not implemented")
}
func (UnimplementedCrowServer) Extend(context.Context, *ExtendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extend not implemented")
}
func (UnimplementedCrowServer) Update(context.Context, *UpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCrowServer) Log(context.Context, *LogRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedCrowServer) RegisterAgent(context.Context, *RegisterAgentRequest) (*RegisterAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgent not implemented")
}
func (UnimplementedCrowServer) UnregisterAgent(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterAgent not implemented")
}
func (UnimplementedCrowServer) ReportHealth(context.Context, *ReportHealthRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportHealth not implemented")
}
func (UnimplementedCrowServer) mustEmbedUnimplementedCrowServer() {}
func (UnimplementedCrowServer) testEmbeddedByValue()                    {}

// UnsafeCrowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrowServer will
// result in compilation errors.
type UnsafeCrowServer interface {
	mustEmbedUnimplementedCrowServer()
}

func RegisterCrowServer(s grpc.ServiceRegistrar, srv CrowServer) {
	// If the following call pancis, it indicates UnimplementedCrowServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Crow_ServiceDesc, srv)
}

func _Crow_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Next_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Next(ctx, req.(*NextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Wait_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Wait(ctx, req.(*WaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Done_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Done(ctx, req.(*DoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Extend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Extend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Extend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Extend(ctx, req.(*ExtendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_Log_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_RegisterAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).RegisterAgent(ctx, req.(*RegisterAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_UnregisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).UnregisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_UnregisterAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).UnregisterAgent(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crow_ReportHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowServer).ReportHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Crow_ReportHealth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowServer).ReportHealth(ctx, req.(*ReportHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Crow_ServiceDesc is the grpc.ServiceDesc for Crow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Crow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Crow",
	HandlerType: (*CrowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Crow_Version_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Crow_Next_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Crow_Init_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _Crow_Wait_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Crow_Done_Handler,
		},
		{
			MethodName: "Extend",
			Handler:    _Crow_Extend_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Crow_Update_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _Crow_Log_Handler,
		},
		{
			MethodName: "RegisterAgent",
			Handler:    _Crow_RegisterAgent_Handler,
		},
		{
			MethodName: "UnregisterAgent",
			Handler:    _Crow_UnregisterAgent_Handler,
		},
		{
			MethodName: "ReportHealth",
			Handler:    _Crow_ReportHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crow.proto",
}

const (
	CrowAuth_Auth_FullMethodName = "/proto.CrowAuth/Auth"
)

// crowAuthClient is the client API for CrowAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrowAuthClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type crowAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewCrowAuthClient(cc grpc.ClientConnInterface) CrowAuthClient {
	return &crowAuthClient{cc}
}

func (c *crowAuthClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, CrowAuth_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrowAuthServer is the server API for CrowAuth service.
// All implementations must embed UnimplementedCrowAuthServer
// for forward compatibility.
type CrowAuthServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedCrowAuthServer()
}

// UnimplementedCrowAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCrowAuthServer struct{}

func (UnimplementedCrowAuthServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedCrowAuthServer) mustEmbedUnimplementedCrowAuthServer() {}
func (UnimplementedCrowAuthServer) testEmbeddedByValue()                  {}

// UnsafeCrowAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrowAuthServer will
// result in compilation errors.
type UnsafeCrowAuthServer interface {
	mustEmbedUnimplementedCrowAuthServer()
}

func RegisterCrowAuthServer(s grpc.ServiceRegistrar, srv CrowAuthServer) {
	// If the following call pancis, it indicates UnimplementedCrowAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CrowAuth_ServiceDesc, srv)
}

func _CrowAuth_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowAuthServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrowAuth_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowAuthServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CrowAuth_ServiceDesc is the grpc.ServiceDesc for CrowAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrowAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CrowAuth",
	HandlerType: (*CrowAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _CrowAuth_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crow.proto",
}
