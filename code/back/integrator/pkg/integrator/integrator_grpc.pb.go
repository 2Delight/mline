// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: integrator.proto

package integrator

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
	WorkerService_SaveSpecification_FullMethodName       = "/integrator.WorkerService/SaveSpecification"
	WorkerService_RunMLDev_FullMethodName                = "/integrator.WorkerService/RunMLDev"
	WorkerService_GetSpecificationFromGit_FullMethodName = "/integrator.WorkerService/GetSpecificationFromGit"
	WorkerService_GetHello_FullMethodName                = "/integrator.WorkerService/GetHello"
)

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	SaveSpecification(ctx context.Context, in *SaveSpecificationRequest, opts ...grpc.CallOption) (*CommitPushResult, error)
	RunMLDev(ctx context.Context, in *RunMLDevRequest, opts ...grpc.CallOption) (*MLDevResult, error)
	GetSpecificationFromGit(ctx context.Context, in *GetSpecificationRequest, opts ...grpc.CallOption) (*Specification, error)
	GetHello(ctx context.Context, in *GetHelloRequest, opts ...grpc.CallOption) (*GetHelloResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) SaveSpecification(ctx context.Context, in *SaveSpecificationRequest, opts ...grpc.CallOption) (*CommitPushResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommitPushResult)
	err := c.cc.Invoke(ctx, WorkerService_SaveSpecification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) RunMLDev(ctx context.Context, in *RunMLDevRequest, opts ...grpc.CallOption) (*MLDevResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MLDevResult)
	err := c.cc.Invoke(ctx, WorkerService_RunMLDev_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetSpecificationFromGit(ctx context.Context, in *GetSpecificationRequest, opts ...grpc.CallOption) (*Specification, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Specification)
	err := c.cc.Invoke(ctx, WorkerService_GetSpecificationFromGit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetHello(ctx context.Context, in *GetHelloRequest, opts ...grpc.CallOption) (*GetHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHelloResponse)
	err := c.cc.Invoke(ctx, WorkerService_GetHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations must embed UnimplementedWorkerServiceServer
// for forward compatibility.
type WorkerServiceServer interface {
	SaveSpecification(context.Context, *SaveSpecificationRequest) (*CommitPushResult, error)
	RunMLDev(context.Context, *RunMLDevRequest) (*MLDevResult, error)
	GetSpecificationFromGit(context.Context, *GetSpecificationRequest) (*Specification, error)
	GetHello(context.Context, *GetHelloRequest) (*GetHelloResponse, error)
	mustEmbedUnimplementedWorkerServiceServer()
}

// UnimplementedWorkerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkerServiceServer struct{}

func (UnimplementedWorkerServiceServer) SaveSpecification(context.Context, *SaveSpecificationRequest) (*CommitPushResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSpecification not implemented")
}
func (UnimplementedWorkerServiceServer) RunMLDev(context.Context, *RunMLDevRequest) (*MLDevResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunMLDev not implemented")
}
func (UnimplementedWorkerServiceServer) GetSpecificationFromGit(context.Context, *GetSpecificationRequest) (*Specification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecificationFromGit not implemented")
}
func (UnimplementedWorkerServiceServer) GetHello(context.Context, *GetHelloRequest) (*GetHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHello not implemented")
}
func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {}
func (UnimplementedWorkerServiceServer) testEmbeddedByValue()                       {}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	// If the following call pancis, it indicates UnimplementedWorkerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_SaveSpecification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSpecificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).SaveSpecification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_SaveSpecification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).SaveSpecification(ctx, req.(*SaveSpecificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_RunMLDev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunMLDevRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).RunMLDev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_RunMLDev_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).RunMLDev(ctx, req.(*RunMLDevRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetSpecificationFromGit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetSpecificationFromGit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_GetSpecificationFromGit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetSpecificationFromGit(ctx, req.(*GetSpecificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_GetHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetHello(ctx, req.(*GetHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "integrator.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveSpecification",
			Handler:    _WorkerService_SaveSpecification_Handler,
		},
		{
			MethodName: "RunMLDev",
			Handler:    _WorkerService_RunMLDev_Handler,
		},
		{
			MethodName: "GetSpecificationFromGit",
			Handler:    _WorkerService_GetSpecificationFromGit_Handler,
		},
		{
			MethodName: "GetHello",
			Handler:    _WorkerService_GetHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrator.proto",
}
