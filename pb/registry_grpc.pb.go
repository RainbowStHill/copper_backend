// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: registry.proto

package pb

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
	Registry_Register_FullMethodName   = "/pb.Registry/Register"
	Registry_Sync_FullMethodName       = "/pb.Registry/Sync"
	Registry_Deregister_FullMethodName = "/pb.Registry/Deregister"
	Registry_Discover_FullMethodName   = "/pb.Registry/Discover"
)

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*SyncResp, error)
	Deregister(ctx context.Context, in *DeregisterReq, opts ...grpc.CallOption) (*DeregisterResp, error)
	Discover(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DiscoverReq, DiscoverResp], error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, Registry_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*SyncResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, Registry_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Deregister(ctx context.Context, in *DeregisterReq, opts ...grpc.CallOption) (*DeregisterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeregisterResp)
	err := c.cc.Invoke(ctx, Registry_Deregister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Discover(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DiscoverReq, DiscoverResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[0], Registry_Discover_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DiscoverReq, DiscoverResp]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Registry_DiscoverClient = grpc.BidiStreamingClient[DiscoverReq, DiscoverResp]

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility.
type RegistryServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Sync(context.Context, *SyncReq) (*SyncResp, error)
	Deregister(context.Context, *DeregisterReq) (*DeregisterResp, error)
	Discover(grpc.BidiStreamingServer[DiscoverReq, DiscoverResp]) error
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRegistryServer struct{}

func (UnimplementedRegistryServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistryServer) Sync(context.Context, *SyncReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedRegistryServer) Deregister(context.Context, *DeregisterReq) (*DeregisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedRegistryServer) Discover(grpc.BidiStreamingServer[DiscoverReq, DiscoverResp]) error {
	return status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}
func (UnimplementedRegistryServer) testEmbeddedByValue()                  {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	// If the following call pancis, it indicates UnimplementedRegistryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Sync(ctx, req.(*SyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Deregister(ctx, req.(*DeregisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistryServer).Discover(&grpc.GenericServerStream[DiscoverReq, DiscoverResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Registry_DiscoverServer = grpc.BidiStreamingServer[DiscoverReq, DiscoverResp]

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registry_Register_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Registry_Sync_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _Registry_Deregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _Registry_Discover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "registry.proto",
}
