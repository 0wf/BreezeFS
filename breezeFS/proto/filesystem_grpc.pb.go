// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: proto/filesystem.proto

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
	ManagerService_RegisterNode_FullMethodName      = "/filesystem.ManagerService/RegisterNode"
	ManagerService_GetNodesForChunks_FullMethodName = "/filesystem.ManagerService/GetNodesForChunks"
	ManagerService_GetChunkLocations_FullMethodName = "/filesystem.ManagerService/GetChunkLocations"
)

// ManagerServiceClient is the client API for ManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerServiceClient interface {
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	GetNodesForChunks(ctx context.Context, in *GetNodesForChunksRequest, opts ...grpc.CallOption) (*GetNodesForChunksResponse, error)
	GetChunkLocations(ctx context.Context, in *GetChunkLocationsRequest, opts ...grpc.CallOption) (*GetChunkLocationsResponse, error)
}

type managerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerServiceClient(cc grpc.ClientConnInterface) ManagerServiceClient {
	return &managerServiceClient{cc}
}

func (c *managerServiceClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, ManagerService_RegisterNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) GetNodesForChunks(ctx context.Context, in *GetNodesForChunksRequest, opts ...grpc.CallOption) (*GetNodesForChunksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodesForChunksResponse)
	err := c.cc.Invoke(ctx, ManagerService_GetNodesForChunks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) GetChunkLocations(ctx context.Context, in *GetChunkLocationsRequest, opts ...grpc.CallOption) (*GetChunkLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChunkLocationsResponse)
	err := c.cc.Invoke(ctx, ManagerService_GetChunkLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServiceServer is the server API for ManagerService service.
// All implementations must embed UnimplementedManagerServiceServer
// for forward compatibility.
type ManagerServiceServer interface {
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	GetNodesForChunks(context.Context, *GetNodesForChunksRequest) (*GetNodesForChunksResponse, error)
	GetChunkLocations(context.Context, *GetChunkLocationsRequest) (*GetChunkLocationsResponse, error)
	mustEmbedUnimplementedManagerServiceServer()
}

// UnimplementedManagerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedManagerServiceServer struct{}

func (UnimplementedManagerServiceServer) RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (UnimplementedManagerServiceServer) GetNodesForChunks(context.Context, *GetNodesForChunksRequest) (*GetNodesForChunksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodesForChunks not implemented")
}
func (UnimplementedManagerServiceServer) GetChunkLocations(context.Context, *GetChunkLocationsRequest) (*GetChunkLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunkLocations not implemented")
}
func (UnimplementedManagerServiceServer) mustEmbedUnimplementedManagerServiceServer() {}
func (UnimplementedManagerServiceServer) testEmbeddedByValue()                        {}

// UnsafeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServiceServer will
// result in compilation errors.
type UnsafeManagerServiceServer interface {
	mustEmbedUnimplementedManagerServiceServer()
}

func RegisterManagerServiceServer(s grpc.ServiceRegistrar, srv ManagerServiceServer) {
	// If the following call pancis, it indicates UnimplementedManagerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ManagerService_ServiceDesc, srv)
}

func _ManagerService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_RegisterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_GetNodesForChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesForChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).GetNodesForChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_GetNodesForChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).GetNodesForChunks(ctx, req.(*GetNodesForChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_GetChunkLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunkLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).GetChunkLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_GetChunkLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).GetChunkLocations(ctx, req.(*GetChunkLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerService_ServiceDesc is the grpc.ServiceDesc for ManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filesystem.ManagerService",
	HandlerType: (*ManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _ManagerService_RegisterNode_Handler,
		},
		{
			MethodName: "GetNodesForChunks",
			Handler:    _ManagerService_GetNodesForChunks_Handler,
		},
		{
			MethodName: "GetChunkLocations",
			Handler:    _ManagerService_GetChunkLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/filesystem.proto",
}
