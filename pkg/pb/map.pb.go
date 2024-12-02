// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sro/gameserver/map.proto

package pb

import (
	pb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MapPath string `protobuf:"bytes,2,opt,name=map_path,json=mapPath,proto3" json:"map_path,omitempty"`
}

func (x *CreateMapRequest) Reset() {
	*x = CreateMapRequest{}
	mi := &file_sro_gameserver_map_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMapRequest) ProtoMessage() {}

func (x *CreateMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_map_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMapRequest.ProtoReflect.Descriptor instead.
func (*CreateMapRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_map_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMapRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMapRequest) GetMapPath() string {
	if x != nil {
		return x.MapPath
	}
	return ""
}

type EditMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Types that are assignable to OptionalName:
	//
	//	*EditMapRequest_Name
	OptionalName isEditMapRequest_OptionalName `protobuf_oneof:"optional_name"`
	// Types that are assignable to OptionalMapPath:
	//
	//	*EditMapRequest_MapPath
	OptionalMapPath isEditMapRequest_OptionalMapPath `protobuf_oneof:"optional_map_path"`
}

func (x *EditMapRequest) Reset() {
	*x = EditMapRequest{}
	mi := &file_sro_gameserver_map_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMapRequest) ProtoMessage() {}

func (x *EditMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_map_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMapRequest.ProtoReflect.Descriptor instead.
func (*EditMapRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_map_proto_rawDescGZIP(), []int{1}
}

func (x *EditMapRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (m *EditMapRequest) GetOptionalName() isEditMapRequest_OptionalName {
	if m != nil {
		return m.OptionalName
	}
	return nil
}

func (x *EditMapRequest) GetName() string {
	if x, ok := x.GetOptionalName().(*EditMapRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *EditMapRequest) GetOptionalMapPath() isEditMapRequest_OptionalMapPath {
	if m != nil {
		return m.OptionalMapPath
	}
	return nil
}

func (x *EditMapRequest) GetMapPath() string {
	if x, ok := x.GetOptionalMapPath().(*EditMapRequest_MapPath); ok {
		return x.MapPath
	}
	return ""
}

type isEditMapRequest_OptionalName interface {
	isEditMapRequest_OptionalName()
}

type EditMapRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*EditMapRequest_Name) isEditMapRequest_OptionalName() {}

type isEditMapRequest_OptionalMapPath interface {
	isEditMapRequest_OptionalMapPath()
}

type EditMapRequest_MapPath struct {
	MapPath string `protobuf:"bytes,3,opt,name=map_path,json=mapPath,proto3,oneof"`
}

func (*EditMapRequest_MapPath) isEditMapRequest_OptionalMapPath() {}

type Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MapPath string `protobuf:"bytes,3,opt,name=map_path,json=mapPath,proto3" json:"map_path,omitempty"`
}

func (x *Map) Reset() {
	*x = Map{}
	mi := &file_sro_gameserver_map_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Map) ProtoMessage() {}

func (x *Map) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_map_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Map.ProtoReflect.Descriptor instead.
func (*Map) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_map_proto_rawDescGZIP(), []int{2}
}

func (x *Map) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Map) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Map) GetMapPath() string {
	if x != nil {
		return x.MapPath
	}
	return ""
}

type Maps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maps []*Map `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
}

func (x *Maps) Reset() {
	*x = Maps{}
	mi := &file_sro_gameserver_map_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Maps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maps) ProtoMessage() {}

func (x *Maps) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_map_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maps.ProtoReflect.Descriptor instead.
func (*Maps) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_map_proto_rawDescGZIP(), []int{3}
}

func (x *Maps) GetMaps() []*Map {
	if x != nil {
		return x.Maps
	}
	return nil
}

var File_sro_gameserver_map_proto protoreflect.FileDescriptor

var file_sro_gameserver_map_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x72, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x72, 0x6f, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x72, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x50, 0x61, 0x74, 0x68, 0x22, 0x86, 0x01, 0x0a, 0x0e,
	0x45, 0x64, 0x69, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x50, 0x61, 0x74, 0x68, 0x42, 0x0f,
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x13, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x44, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2f, 0x0a, 0x04, 0x4d, 0x61,
	0x70, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x32, 0xaa, 0x03, 0x0a, 0x0a,
	0x4d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x12, 0x0d, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x12, 0x57, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x20, 0x2e, 0x73, 0x72, 0x6f,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73,
	0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x61,
	0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x12, 0x62, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x61,
	0x70, 0x12, 0x1e, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x4d, 0x61, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x4c, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x0d, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sro_gameserver_map_proto_rawDescOnce sync.Once
	file_sro_gameserver_map_proto_rawDescData = file_sro_gameserver_map_proto_rawDesc
)

func file_sro_gameserver_map_proto_rawDescGZIP() []byte {
	file_sro_gameserver_map_proto_rawDescOnce.Do(func() {
		file_sro_gameserver_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_sro_gameserver_map_proto_rawDescData)
	})
	return file_sro_gameserver_map_proto_rawDescData
}

var file_sro_gameserver_map_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sro_gameserver_map_proto_goTypes = []any{
	(*CreateMapRequest)(nil), // 0: sro.gameserver.CreateMapRequest
	(*EditMapRequest)(nil),   // 1: sro.gameserver.EditMapRequest
	(*Map)(nil),              // 2: sro.gameserver.Map
	(*Maps)(nil),             // 3: sro.gameserver.Maps
	(*pb.TargetId)(nil),      // 4: sro.TargetId
	(*emptypb.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_sro_gameserver_map_proto_depIdxs = []int32{
	2, // 0: sro.gameserver.Maps.maps:type_name -> sro.gameserver.Map
	4, // 1: sro.gameserver.MapService.GetMap:input_type -> sro.TargetId
	5, // 2: sro.gameserver.MapService.GetMaps:input_type -> google.protobuf.Empty
	0, // 3: sro.gameserver.MapService.CreateMap:input_type -> sro.gameserver.CreateMapRequest
	1, // 4: sro.gameserver.MapService.EditMap:input_type -> sro.gameserver.EditMapRequest
	4, // 5: sro.gameserver.MapService.DeleteMap:input_type -> sro.TargetId
	2, // 6: sro.gameserver.MapService.GetMap:output_type -> sro.gameserver.Map
	3, // 7: sro.gameserver.MapService.GetMaps:output_type -> sro.gameserver.Maps
	2, // 8: sro.gameserver.MapService.CreateMap:output_type -> sro.gameserver.Map
	2, // 9: sro.gameserver.MapService.EditMap:output_type -> sro.gameserver.Map
	5, // 10: sro.gameserver.MapService.DeleteMap:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sro_gameserver_map_proto_init() }
func file_sro_gameserver_map_proto_init() {
	if File_sro_gameserver_map_proto != nil {
		return
	}
	file_sro_gameserver_map_proto_msgTypes[1].OneofWrappers = []any{
		(*EditMapRequest_Name)(nil),
		(*EditMapRequest_MapPath)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sro_gameserver_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sro_gameserver_map_proto_goTypes,
		DependencyIndexes: file_sro_gameserver_map_proto_depIdxs,
		MessageInfos:      file_sro_gameserver_map_proto_msgTypes,
	}.Build()
	File_sro_gameserver_map_proto = out.File
	file_sro_gameserver_map_proto_rawDesc = nil
	file_sro_gameserver_map_proto_goTypes = nil
	file_sro_gameserver_map_proto_depIdxs = nil
}
