// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: gateway.proto

package gateway

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
	GatewayService_GetSpecification_FullMethodName      = "/gateway.GatewayService/GetSpecification"
	GatewayService_UpdateSpecification_FullMethodName   = "/gateway.GatewayService/UpdateSpecification"
	GatewayService_GetStatus_FullMethodName             = "/gateway.GatewayService/GetStatus"
	GatewayService_UpdateStatus_FullMethodName          = "/gateway.GatewayService/UpdateStatus"
	GatewayService_ValidateSpecification_FullMethodName = "/gateway.GatewayService/ValidateSpecification"
	GatewayService_GetHello_FullMethodName              = "/gateway.GatewayService/GetHello"
)

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	GetSpecification(ctx context.Context, in *GetSpecificationRequest, opts ...grpc.CallOption) (*Specification, error)
	UpdateSpecification(ctx context.Context, in *UpdateSpecificationRequest, opts ...grpc.CallOption) (*UpdateSpecificationResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*Status, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*StatusUpdateResponse, error)
	ValidateSpecification(ctx context.Context, in *ValidateSpecificationRequest, opts ...grpc.CallOption) (*ValidationResult, error)
	GetHello(ctx context.Context, in *GetHelloRequest, opts ...grpc.CallOption) (*GetHelloResponse, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) GetSpecification(ctx context.Context, in *GetSpecificationRequest, opts ...grpc.CallOption) (*Specification, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Specification)
	err := c.cc.Invoke(ctx, GatewayService_GetSpecification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) UpdateSpecification(ctx context.Context, in *UpdateSpecificationRequest, opts ...grpc.CallOption) (*UpdateSpecificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSpecificationResponse)
	err := c.cc.Invoke(ctx, GatewayService_UpdateSpecification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, GatewayService_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*StatusUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusUpdateResponse)
	err := c.cc.Invoke(ctx, GatewayService_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) ValidateSpecification(ctx context.Context, in *ValidateSpecificationRequest, opts ...grpc.CallOption) (*ValidationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidationResult)
	err := c.cc.Invoke(ctx, GatewayService_ValidateSpecification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetHello(ctx context.Context, in *GetHelloRequest, opts ...grpc.CallOption) (*GetHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHelloResponse)
	err := c.cc.Invoke(ctx, GatewayService_GetHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations must embed UnimplementedGatewayServiceServer
// for forward compatibility.
type GatewayServiceServer interface {
	GetSpecification(context.Context, *GetSpecificationRequest) (*Specification, error)
	UpdateSpecification(context.Context, *UpdateSpecificationRequest) (*UpdateSpecificationResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*Status, error)
	UpdateStatus(context.Context, *UpdateStatusRequest) (*StatusUpdateResponse, error)
	ValidateSpecification(context.Context, *ValidateSpecificationRequest) (*ValidationResult, error)
	GetHello(context.Context, *GetHelloRequest) (*GetHelloResponse, error)
	mustEmbedUnimplementedGatewayServiceServer()
}

// UnimplementedGatewayServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGatewayServiceServer struct{}

func (UnimplementedGatewayServiceServer) GetSpecification(context.Context, *GetSpecificationRequest) (*Specification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecification not implemented")
}
func (UnimplementedGatewayServiceServer) UpdateSpecification(context.Context, *UpdateSpecificationRequest) (*UpdateSpecificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpecification not implemented")
}
func (UnimplementedGatewayServiceServer) GetStatus(context.Context, *GetStatusRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedGatewayServiceServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*StatusUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedGatewayServiceServer) ValidateSpecification(context.Context, *ValidateSpecificationRequest) (*ValidationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSpecification not implemented")
}
func (UnimplementedGatewayServiceServer) GetHello(context.Context, *GetHelloRequest) (*GetHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHello not implemented")
}
func (UnimplementedGatewayServiceServer) mustEmbedUnimplementedGatewayServiceServer() {}
func (UnimplementedGatewayServiceServer) testEmbeddedByValue()                        {}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	// If the following call pancis, it indicates UnimplementedGatewayServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_GetSpecification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetSpecification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_GetSpecification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetSpecification(ctx, req.(*GetSpecificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_UpdateSpecification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpecificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).UpdateSpecification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_UpdateSpecification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).UpdateSpecification(ctx, req.(*UpdateSpecificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_ValidateSpecification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSpecificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ValidateSpecification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_ValidateSpecification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ValidateSpecification(ctx, req.(*ValidateSpecificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_GetHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetHello(ctx, req.(*GetHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSpecification",
			Handler:    _GatewayService_GetSpecification_Handler,
		},
		{
			MethodName: "UpdateSpecification",
			Handler:    _GatewayService_UpdateSpecification_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _GatewayService_GetStatus_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _GatewayService_UpdateStatus_Handler,
		},
		{
			MethodName: "ValidateSpecification",
			Handler:    _GatewayService_ValidateSpecification_Handler,
		},
		{
			MethodName: "GetHello",
			Handler:    _GatewayService_GetHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
