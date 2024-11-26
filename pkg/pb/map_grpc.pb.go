// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: sro/gameserver/map.proto

package pb

import (
	context "context"
	pb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MapService_GetMap_FullMethodName    = "/sro.gameserver.MapService/GetMap"
	MapService_GetMaps_FullMethodName   = "/sro.gameserver.MapService/GetMaps"
	MapService_CreateMap_FullMethodName = "/sro.gameserver.MapService/CreateMap"
	MapService_EditMap_FullMethodName   = "/sro.gameserver.MapService/EditMap"
	MapService_DeleteMap_FullMethodName = "/sro.gameserver.MapService/DeleteMap"
)

// MapServiceClient is the client API for MapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapServiceClient interface {
	GetMap(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*Map, error)
	GetMaps(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Maps, error)
	CreateMap(ctx context.Context, in *CreateMapRequest, opts ...grpc.CallOption) (*Map, error)
	EditMap(ctx context.Context, in *EditMapRequest, opts ...grpc.CallOption) (*Map, error)
	DeleteMap(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mapServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMapServiceClient(cc grpc.ClientConnInterface) MapServiceClient {
	return &mapServiceClient{cc}
}

func (c *mapServiceClient) GetMap(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*Map, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Map)
	err := c.cc.Invoke(ctx, MapService_GetMap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapServiceClient) GetMaps(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Maps, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Maps)
	err := c.cc.Invoke(ctx, MapService_GetMaps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapServiceClient) CreateMap(ctx context.Context, in *CreateMapRequest, opts ...grpc.CallOption) (*Map, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Map)
	err := c.cc.Invoke(ctx, MapService_CreateMap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapServiceClient) EditMap(ctx context.Context, in *EditMapRequest, opts ...grpc.CallOption) (*Map, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Map)
	err := c.cc.Invoke(ctx, MapService_EditMap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapServiceClient) DeleteMap(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MapService_DeleteMap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapServiceServer is the server API for MapService service.
// All implementations must embed UnimplementedMapServiceServer
// for forward compatibility.
type MapServiceServer interface {
	GetMap(context.Context, *pb.TargetId) (*Map, error)
	GetMaps(context.Context, *emptypb.Empty) (*Maps, error)
	CreateMap(context.Context, *CreateMapRequest) (*Map, error)
	EditMap(context.Context, *EditMapRequest) (*Map, error)
	DeleteMap(context.Context, *pb.TargetId) (*emptypb.Empty, error)
	mustEmbedUnimplementedMapServiceServer()
}

// UnimplementedMapServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMapServiceServer struct{}

func (UnimplementedMapServiceServer) GetMap(context.Context, *pb.TargetId) (*Map, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMap not implemented")
}
func (UnimplementedMapServiceServer) GetMaps(context.Context, *emptypb.Empty) (*Maps, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaps not implemented")
}
func (UnimplementedMapServiceServer) CreateMap(context.Context, *CreateMapRequest) (*Map, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMap not implemented")
}
func (UnimplementedMapServiceServer) EditMap(context.Context, *EditMapRequest) (*Map, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMap not implemented")
}
func (UnimplementedMapServiceServer) DeleteMap(context.Context, *pb.TargetId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMap not implemented")
}
func (UnimplementedMapServiceServer) mustEmbedUnimplementedMapServiceServer() {}
func (UnimplementedMapServiceServer) testEmbeddedByValue()                    {}

// UnsafeMapServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapServiceServer will
// result in compilation errors.
type UnsafeMapServiceServer interface {
	mustEmbedUnimplementedMapServiceServer()
}

func RegisterMapServiceServer(s grpc.ServiceRegistrar, srv MapServiceServer) {
	// If the following call pancis, it indicates UnimplementedMapServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MapService_ServiceDesc, srv)
}

func _MapService_GetMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServiceServer).GetMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapService_GetMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServiceServer).GetMap(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapService_GetMaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServiceServer).GetMaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapService_GetMaps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServiceServer).GetMaps(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapService_CreateMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServiceServer).CreateMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapService_CreateMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServiceServer).CreateMap(ctx, req.(*CreateMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapService_EditMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServiceServer).EditMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapService_EditMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServiceServer).EditMap(ctx, req.(*EditMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapService_DeleteMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServiceServer).DeleteMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapService_DeleteMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServiceServer).DeleteMap(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

// MapService_ServiceDesc is the grpc.ServiceDesc for MapService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sro.gameserver.MapService",
	HandlerType: (*MapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMap",
			Handler:    _MapService_GetMap_Handler,
		},
		{
			MethodName: "GetMaps",
			Handler:    _MapService_GetMaps_Handler,
		},
		{
			MethodName: "CreateMap",
			Handler:    _MapService_CreateMap_Handler,
		},
		{
			MethodName: "EditMap",
			Handler:    _MapService_EditMap_Handler,
		},
		{
			MethodName: "DeleteMap",
			Handler:    _MapService_DeleteMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sro/gameserver/map.proto",
}
