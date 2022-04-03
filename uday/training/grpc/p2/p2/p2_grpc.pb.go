// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: pb/p2.proto

package p2

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
	NameService(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error)
	NumberService(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) NameService(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/p2.Service/NameService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NumberService(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error) {
	out := new(NumberResponse)
	err := c.cc.Invoke(ctx, "/p2.Service/NumberService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	NameService(context.Context, *NameRequest) (*NameResponse, error)
	NumberService(context.Context, *NumberRequest) (*NumberResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) NameService(context.Context, *NameRequest) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NameService not implemented")
}
func (UnimplementedServiceServer) NumberService(context.Context, *NumberRequest) (*NumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NumberService not implemented")
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

func _Service_NameService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).NameService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2.Service/NameService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).NameService(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_NumberService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).NumberService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2.Service/NumberService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).NumberService(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p2.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NameService",
			Handler:    _Service_NameService_Handler,
		},
		{
			MethodName: "NumberService",
			Handler:    _Service_NumberService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/p2.proto",
}