// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Start(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Stop(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Start(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/v1.Service/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Stop(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/v1.Service/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Start(context.Context, *ServiceRequest) (*StatusResponse, error)
	Stop(context.Context, *ServiceRequest) (*StatusResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Start(context.Context, *ServiceRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedServiceServer) Stop(context.Context, *ServiceRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Service/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Start(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Service/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Stop(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Service_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Service_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

// DockerClient is the client API for Docker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerClient interface {
	Start(ctx context.Context, in *DockerRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Stop(ctx context.Context, in *DockerRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type dockerClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerClient(cc grpc.ClientConnInterface) DockerClient {
	return &dockerClient{cc}
}

func (c *dockerClient) Start(ctx context.Context, in *DockerRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/v1.Docker/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerClient) Stop(ctx context.Context, in *DockerRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/v1.Docker/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerServer is the server API for Docker service.
// All implementations must embed UnimplementedDockerServer
// for forward compatibility
type DockerServer interface {
	Start(context.Context, *DockerRequest) (*StatusResponse, error)
	Stop(context.Context, *DockerRequest) (*StatusResponse, error)
	mustEmbedUnimplementedDockerServer()
}

// UnimplementedDockerServer must be embedded to have forward compatible implementations.
type UnimplementedDockerServer struct {
}

func (UnimplementedDockerServer) Start(context.Context, *DockerRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedDockerServer) Stop(context.Context, *DockerRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedDockerServer) mustEmbedUnimplementedDockerServer() {}

// UnsafeDockerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerServer will
// result in compilation errors.
type UnsafeDockerServer interface {
	mustEmbedUnimplementedDockerServer()
}

func RegisterDockerServer(s grpc.ServiceRegistrar, srv DockerServer) {
	s.RegisterService(&Docker_ServiceDesc, srv)
}

func _Docker_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Docker/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServer).Start(ctx, req.(*DockerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Docker_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Docker/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServer).Stop(ctx, req.(*DockerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Docker_ServiceDesc is the grpc.ServiceDesc for Docker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Docker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Docker",
	HandlerType: (*DockerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Docker_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Docker_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

// CPUClient is the client API for CPU service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CPUClient interface {
	Start(ctx context.Context, in *CPURequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Stop(ctx context.Context, in *CPURequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type cPUClient struct {
	cc grpc.ClientConnInterface
}

func NewCPUClient(cc grpc.ClientConnInterface) CPUClient {
	return &cPUClient{cc}
}

func (c *cPUClient) Start(ctx context.Context, in *CPURequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/v1.CPU/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPUClient) Stop(ctx context.Context, in *CPURequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/v1.CPU/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CPUServer is the server API for CPU service.
// All implementations must embed UnimplementedCPUServer
// for forward compatibility
type CPUServer interface {
	Start(context.Context, *CPURequest) (*StatusResponse, error)
	Stop(context.Context, *CPURequest) (*StatusResponse, error)
	mustEmbedUnimplementedCPUServer()
}

// UnimplementedCPUServer must be embedded to have forward compatible implementations.
type UnimplementedCPUServer struct {
}

func (UnimplementedCPUServer) Start(context.Context, *CPURequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedCPUServer) Stop(context.Context, *CPURequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedCPUServer) mustEmbedUnimplementedCPUServer() {}

// UnsafeCPUServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CPUServer will
// result in compilation errors.
type UnsafeCPUServer interface {
	mustEmbedUnimplementedCPUServer()
}

func RegisterCPUServer(s grpc.ServiceRegistrar, srv CPUServer) {
	s.RegisterService(&CPU_ServiceDesc, srv)
}

func _CPU_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CPURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPUServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CPU/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPUServer).Start(ctx, req.(*CPURequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPU_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CPURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPUServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CPU/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPUServer).Stop(ctx, req.(*CPURequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CPU_ServiceDesc is the grpc.ServiceDesc for CPU service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CPU_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CPU",
	HandlerType: (*CPUServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _CPU_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _CPU_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

// StrategyClient is the client API for Strategy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyClient interface {
	Recover(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type strategyClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyClient(cc grpc.ClientConnInterface) StrategyClient {
	return &strategyClient{cc}
}

func (c *strategyClient) Recover(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/v1.Strategy/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategyServer is the server API for Strategy service.
// All implementations must embed UnimplementedStrategyServer
// for forward compatibility
type StrategyServer interface {
	Recover(context.Context, *RecoverRequest) (*ResolveResponse, error)
	mustEmbedUnimplementedStrategyServer()
}

// UnimplementedStrategyServer must be embedded to have forward compatible implementations.
type UnimplementedStrategyServer struct {
}

func (UnimplementedStrategyServer) Recover(context.Context, *RecoverRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (UnimplementedStrategyServer) mustEmbedUnimplementedStrategyServer() {}

// UnsafeStrategyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategyServer will
// result in compilation errors.
type UnsafeStrategyServer interface {
	mustEmbedUnimplementedStrategyServer()
}

func RegisterStrategyServer(s grpc.ServiceRegistrar, srv StrategyServer) {
	s.RegisterService(&Strategy_ServiceDesc, srv)
}

func _Strategy_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Strategy/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServer).Recover(ctx, req.(*RecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Strategy_ServiceDesc is the grpc.ServiceDesc for Strategy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Strategy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Strategy",
	HandlerType: (*StrategyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recover",
			Handler:    _Strategy_Recover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}