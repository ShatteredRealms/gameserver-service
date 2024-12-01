// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: sro/gameserver/dimension.proto

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
	DimensionService_GetDimension_FullMethodName       = "/sro.gameserver.DimensionService/GetDimension"
	DimensionService_GetDimensions_FullMethodName      = "/sro.gameserver.DimensionService/GetDimensions"
	DimensionService_CreateDimension_FullMethodName    = "/sro.gameserver.DimensionService/CreateDimension"
	DimensionService_DuplicateDimension_FullMethodName = "/sro.gameserver.DimensionService/DuplicateDimension"
	DimensionService_EditDimension_FullMethodName      = "/sro.gameserver.DimensionService/EditDimension"
	DimensionService_DeleteDimension_FullMethodName    = "/sro.gameserver.DimensionService/DeleteDimension"
)

// DimensionServiceClient is the client API for DimensionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DimensionServiceClient interface {
	GetDimension(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*Dimension, error)
	GetDimensions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Dimensions, error)
	CreateDimension(ctx context.Context, in *CreateDimensionRequest, opts ...grpc.CallOption) (*Dimension, error)
	DuplicateDimension(ctx context.Context, in *DuplicateDimensionRequest, opts ...grpc.CallOption) (*Dimension, error)
	EditDimension(ctx context.Context, in *EditDimensionRequest, opts ...grpc.CallOption) (*Dimension, error)
	DeleteDimension(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dimensionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDimensionServiceClient(cc grpc.ClientConnInterface) DimensionServiceClient {
	return &dimensionServiceClient{cc}
}

func (c *dimensionServiceClient) GetDimension(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*Dimension, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dimension)
	err := c.cc.Invoke(ctx, DimensionService_GetDimension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dimensionServiceClient) GetDimensions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Dimensions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dimensions)
	err := c.cc.Invoke(ctx, DimensionService_GetDimensions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dimensionServiceClient) CreateDimension(ctx context.Context, in *CreateDimensionRequest, opts ...grpc.CallOption) (*Dimension, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dimension)
	err := c.cc.Invoke(ctx, DimensionService_CreateDimension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dimensionServiceClient) DuplicateDimension(ctx context.Context, in *DuplicateDimensionRequest, opts ...grpc.CallOption) (*Dimension, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dimension)
	err := c.cc.Invoke(ctx, DimensionService_DuplicateDimension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dimensionServiceClient) EditDimension(ctx context.Context, in *EditDimensionRequest, opts ...grpc.CallOption) (*Dimension, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dimension)
	err := c.cc.Invoke(ctx, DimensionService_EditDimension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dimensionServiceClient) DeleteDimension(ctx context.Context, in *pb.TargetId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DimensionService_DeleteDimension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DimensionServiceServer is the server API for DimensionService service.
// All implementations must embed UnimplementedDimensionServiceServer
// for forward compatibility.
type DimensionServiceServer interface {
	GetDimension(context.Context, *pb.TargetId) (*Dimension, error)
	GetDimensions(context.Context, *emptypb.Empty) (*Dimensions, error)
	CreateDimension(context.Context, *CreateDimensionRequest) (*Dimension, error)
	DuplicateDimension(context.Context, *DuplicateDimensionRequest) (*Dimension, error)
	EditDimension(context.Context, *EditDimensionRequest) (*Dimension, error)
	DeleteDimension(context.Context, *pb.TargetId) (*emptypb.Empty, error)
	mustEmbedUnimplementedDimensionServiceServer()
}

// UnimplementedDimensionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDimensionServiceServer struct{}

func (UnimplementedDimensionServiceServer) GetDimension(context.Context, *pb.TargetId) (*Dimension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDimension not implemented")
}
func (UnimplementedDimensionServiceServer) GetDimensions(context.Context, *emptypb.Empty) (*Dimensions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDimensions not implemented")
}
func (UnimplementedDimensionServiceServer) CreateDimension(context.Context, *CreateDimensionRequest) (*Dimension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDimension not implemented")
}
func (UnimplementedDimensionServiceServer) DuplicateDimension(context.Context, *DuplicateDimensionRequest) (*Dimension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DuplicateDimension not implemented")
}
func (UnimplementedDimensionServiceServer) EditDimension(context.Context, *EditDimensionRequest) (*Dimension, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDimension not implemented")
}
func (UnimplementedDimensionServiceServer) DeleteDimension(context.Context, *pb.TargetId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDimension not implemented")
}
func (UnimplementedDimensionServiceServer) mustEmbedUnimplementedDimensionServiceServer() {}
func (UnimplementedDimensionServiceServer) testEmbeddedByValue()                          {}

// UnsafeDimensionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DimensionServiceServer will
// result in compilation errors.
type UnsafeDimensionServiceServer interface {
	mustEmbedUnimplementedDimensionServiceServer()
}

func RegisterDimensionServiceServer(s grpc.ServiceRegistrar, srv DimensionServiceServer) {
	// If the following call pancis, it indicates UnimplementedDimensionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DimensionService_ServiceDesc, srv)
}

func _DimensionService_GetDimension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DimensionServiceServer).GetDimension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DimensionService_GetDimension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DimensionServiceServer).GetDimension(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DimensionService_GetDimensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DimensionServiceServer).GetDimensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DimensionService_GetDimensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DimensionServiceServer).GetDimensions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DimensionService_CreateDimension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDimensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DimensionServiceServer).CreateDimension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DimensionService_CreateDimension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DimensionServiceServer).CreateDimension(ctx, req.(*CreateDimensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DimensionService_DuplicateDimension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DuplicateDimensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DimensionServiceServer).DuplicateDimension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DimensionService_DuplicateDimension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DimensionServiceServer).DuplicateDimension(ctx, req.(*DuplicateDimensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DimensionService_EditDimension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditDimensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DimensionServiceServer).EditDimension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DimensionService_EditDimension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DimensionServiceServer).EditDimension(ctx, req.(*EditDimensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DimensionService_DeleteDimension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TargetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DimensionServiceServer).DeleteDimension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DimensionService_DeleteDimension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DimensionServiceServer).DeleteDimension(ctx, req.(*pb.TargetId))
	}
	return interceptor(ctx, in, info, handler)
}

// DimensionService_ServiceDesc is the grpc.ServiceDesc for DimensionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DimensionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sro.gameserver.DimensionService",
	HandlerType: (*DimensionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDimension",
			Handler:    _DimensionService_GetDimension_Handler,
		},
		{
			MethodName: "GetDimensions",
			Handler:    _DimensionService_GetDimensions_Handler,
		},
		{
			MethodName: "CreateDimension",
			Handler:    _DimensionService_CreateDimension_Handler,
		},
		{
			MethodName: "DuplicateDimension",
			Handler:    _DimensionService_DuplicateDimension_Handler,
		},
		{
			MethodName: "EditDimension",
			Handler:    _DimensionService_EditDimension_Handler,
		},
		{
			MethodName: "DeleteDimension",
			Handler:    _DimensionService_DeleteDimension_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sro/gameserver/dimension.proto",
}
