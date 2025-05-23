// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.15.8
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMapRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MapPath       string                 `protobuf:"bytes,2,opt,name=map_path,json=mapPath,proto3" json:"map_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state    protoimpl.MessageState `protogen:"open.v1"`
	TargetId string                 `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Types that are valid to be assigned to OptionalName:
	//
	//	*EditMapRequest_Name
	OptionalName isEditMapRequest_OptionalName `protobuf_oneof:"optional_name"`
	// Types that are valid to be assigned to OptionalMapPath:
	//
	//	*EditMapRequest_MapPath
	OptionalMapPath isEditMapRequest_OptionalMapPath `protobuf_oneof:"optional_map_path"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
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

func (x *EditMapRequest) GetOptionalName() isEditMapRequest_OptionalName {
	if x != nil {
		return x.OptionalName
	}
	return nil
}

func (x *EditMapRequest) GetName() string {
	if x != nil {
		if x, ok := x.OptionalName.(*EditMapRequest_Name); ok {
			return x.Name
		}
	}
	return ""
}

func (x *EditMapRequest) GetOptionalMapPath() isEditMapRequest_OptionalMapPath {
	if x != nil {
		return x.OptionalMapPath
	}
	return nil
}

func (x *EditMapRequest) GetMapPath() string {
	if x != nil {
		if x, ok := x.OptionalMapPath.(*EditMapRequest_MapPath); ok {
			return x.MapPath
		}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MapPath       string                 `protobuf:"bytes,3,opt,name=map_path,json=mapPath,proto3" json:"map_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Maps          []*Map                 `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_sro_gameserver_map_proto_rawDesc = "" +
	"\n" +
	"\x18sro/gameserver/map.proto\x12\x0esro.gameserver\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x11sro/globals.proto\"A\n" +
	"\x10CreateMapRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x19\n" +
	"\bmap_path\x18\x02 \x01(\tR\amapPath\"\x86\x01\n" +
	"\x0eEditMapRequest\x12\x1b\n" +
	"\ttarget_id\x18\x01 \x01(\tR\btargetId\x12\x14\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x12\x1b\n" +
	"\bmap_path\x18\x03 \x01(\tH\x01R\amapPathB\x0f\n" +
	"\roptional_nameB\x13\n" +
	"\x11optional_map_path\"D\n" +
	"\x03Map\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x19\n" +
	"\bmap_path\x18\x03 \x01(\tR\amapPath\"/\n" +
	"\x04Maps\x12'\n" +
	"\x04maps\x18\x01 \x03(\v2\x13.sro.gameserver.MapR\x04maps2\xaa\x03\n" +
	"\n" +
	"MapService\x12F\n" +
	"\x06GetMap\x12\r.sro.TargetId\x1a\x13.sro.gameserver.Map\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/v1/maps/id/{id}\x12I\n" +
	"\aGetMaps\x12\x16.google.protobuf.Empty\x1a\x14.sro.gameserver.Maps\"\x10\x82\xd3\xe4\x93\x02\n" +
	"\x12\b/v1/maps\x12W\n" +
	"\tCreateMap\x12 .sro.gameserver.CreateMapRequest\x1a\x13.sro.gameserver.Map\"\x13\x82\xd3\xe4\x93\x02\r:\x01*\"\b/v1/maps\x12b\n" +
	"\aEditMap\x12\x1e.sro.gameserver.EditMapRequest\x1a\x13.sro.gameserver.Map\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\x1a\x17/v1/maps/id/{target_id}\x12L\n" +
	"\tDeleteMap\x12\r.sro.TargetId\x1a\x16.google.protobuf.Empty\"\x18\x82\xd3\xe4\x93\x02\x12*\x10/v1/maps/id/{id}B9Z7github.com/ShatteredRealms/gameserver-service/pkg/pb;pbb\x06proto3"

var (
	file_sro_gameserver_map_proto_rawDescOnce sync.Once
	file_sro_gameserver_map_proto_rawDescData []byte
)

func file_sro_gameserver_map_proto_rawDescGZIP() []byte {
	file_sro_gameserver_map_proto_rawDescOnce.Do(func() {
		file_sro_gameserver_map_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sro_gameserver_map_proto_rawDesc), len(file_sro_gameserver_map_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sro_gameserver_map_proto_rawDesc), len(file_sro_gameserver_map_proto_rawDesc)),
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
	file_sro_gameserver_map_proto_goTypes = nil
	file_sro_gameserver_map_proto_depIdxs = nil
}
