// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.15.8
// source: sro/gameserver/dimension.proto

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

type CreateDimensionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	MapIds        []string               `protobuf:"bytes,3,rep,name=map_ids,json=mapIds,proto3" json:"map_ids,omitempty"`
	Location      string                 `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDimensionRequest) Reset() {
	*x = CreateDimensionRequest{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDimensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDimensionRequest) ProtoMessage() {}

func (x *CreateDimensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDimensionRequest.ProtoReflect.Descriptor instead.
func (*CreateDimensionRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDimensionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDimensionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateDimensionRequest) GetMapIds() []string {
	if x != nil {
		return x.MapIds
	}
	return nil
}

func (x *CreateDimensionRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type DuplicateDimensionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetId      string                 `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DuplicateDimensionRequest) Reset() {
	*x = DuplicateDimensionRequest{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DuplicateDimensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuplicateDimensionRequest) ProtoMessage() {}

func (x *DuplicateDimensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuplicateDimensionRequest.ProtoReflect.Descriptor instead.
func (*DuplicateDimensionRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{1}
}

func (x *DuplicateDimensionRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *DuplicateDimensionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EditDimensionRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	TargetId string                 `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Types that are valid to be assigned to OptionalName:
	//
	//	*EditDimensionRequest_Name
	OptionalName isEditDimensionRequest_OptionalName `protobuf_oneof:"optional_name"`
	// Types that are valid to be assigned to OptionalVersion:
	//
	//	*EditDimensionRequest_Version
	OptionalVersion isEditDimensionRequest_OptionalVersion `protobuf_oneof:"optional_version"`
	EditMaps        bool                                   `protobuf:"varint,4,opt,name=edit_maps,json=editMaps,proto3" json:"edit_maps,omitempty"`
	MapIds          []string                               `protobuf:"bytes,5,rep,name=map_ids,json=mapIds,proto3" json:"map_ids,omitempty"`
	// Types that are valid to be assigned to OptionalLocation:
	//
	//	*EditDimensionRequest_Location
	OptionalLocation isEditDimensionRequest_OptionalLocation `protobuf_oneof:"optional_location"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *EditDimensionRequest) Reset() {
	*x = EditDimensionRequest{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditDimensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDimensionRequest) ProtoMessage() {}

func (x *EditDimensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDimensionRequest.ProtoReflect.Descriptor instead.
func (*EditDimensionRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{2}
}

func (x *EditDimensionRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *EditDimensionRequest) GetOptionalName() isEditDimensionRequest_OptionalName {
	if x != nil {
		return x.OptionalName
	}
	return nil
}

func (x *EditDimensionRequest) GetName() string {
	if x != nil {
		if x, ok := x.OptionalName.(*EditDimensionRequest_Name); ok {
			return x.Name
		}
	}
	return ""
}

func (x *EditDimensionRequest) GetOptionalVersion() isEditDimensionRequest_OptionalVersion {
	if x != nil {
		return x.OptionalVersion
	}
	return nil
}

func (x *EditDimensionRequest) GetVersion() string {
	if x != nil {
		if x, ok := x.OptionalVersion.(*EditDimensionRequest_Version); ok {
			return x.Version
		}
	}
	return ""
}

func (x *EditDimensionRequest) GetEditMaps() bool {
	if x != nil {
		return x.EditMaps
	}
	return false
}

func (x *EditDimensionRequest) GetMapIds() []string {
	if x != nil {
		return x.MapIds
	}
	return nil
}

func (x *EditDimensionRequest) GetOptionalLocation() isEditDimensionRequest_OptionalLocation {
	if x != nil {
		return x.OptionalLocation
	}
	return nil
}

func (x *EditDimensionRequest) GetLocation() string {
	if x != nil {
		if x, ok := x.OptionalLocation.(*EditDimensionRequest_Location); ok {
			return x.Location
		}
	}
	return ""
}

type isEditDimensionRequest_OptionalName interface {
	isEditDimensionRequest_OptionalName()
}

type EditDimensionRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*EditDimensionRequest_Name) isEditDimensionRequest_OptionalName() {}

type isEditDimensionRequest_OptionalVersion interface {
	isEditDimensionRequest_OptionalVersion()
}

type EditDimensionRequest_Version struct {
	Version string `protobuf:"bytes,3,opt,name=version,proto3,oneof"`
}

func (*EditDimensionRequest_Version) isEditDimensionRequest_OptionalVersion() {}

type isEditDimensionRequest_OptionalLocation interface {
	isEditDimensionRequest_OptionalLocation()
}

type EditDimensionRequest_Location struct {
	Location string `protobuf:"bytes,8,opt,name=location,proto3,oneof"`
}

func (*EditDimensionRequest_Location) isEditDimensionRequest_OptionalLocation() {}

type Dimension struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique id for the dimension
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique name of the dimension
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Server version used for each server instance
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// All maps the realm should have available
	MapIds []string `protobuf:"bytes,4,rep,name=map_ids,json=mapIds,proto3" json:"map_ids,omitempty"`
	// Physical server location
	Location      string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{3}
}

func (x *Dimension) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dimension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dimension) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Dimension) GetMapIds() []string {
	if x != nil {
		return x.MapIds
	}
	return nil
}

func (x *Dimension) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type Dimensions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dimensions    []*Dimension           `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Dimensions) Reset() {
	*x = Dimensions{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dimensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimensions) ProtoMessage() {}

func (x *Dimensions) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimensions.ProtoReflect.Descriptor instead.
func (*Dimensions) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{4}
}

func (x *Dimensions) GetDimensions() []*Dimension {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

type GameServerUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MapIdsCreated []string               `protobuf:"bytes,1,rep,name=map_ids_created,json=mapIdsCreated,proto3" json:"map_ids_created,omitempty"`
	MapIdsDeleted []string               `protobuf:"bytes,2,rep,name=map_ids_deleted,json=mapIdsDeleted,proto3" json:"map_ids_deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameServerUpdateResponse) Reset() {
	*x = GameServerUpdateResponse{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameServerUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameServerUpdateResponse) ProtoMessage() {}

func (x *GameServerUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameServerUpdateResponse.ProtoReflect.Descriptor instead.
func (*GameServerUpdateResponse) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{5}
}

func (x *GameServerUpdateResponse) GetMapIdsCreated() []string {
	if x != nil {
		return x.MapIdsCreated
	}
	return nil
}

func (x *GameServerUpdateResponse) GetMapIdsDeleted() []string {
	if x != nil {
		return x.MapIdsDeleted
	}
	return nil
}

var File_sro_gameserver_dimension_proto protoreflect.FileDescriptor

const file_sro_gameserver_dimension_proto_rawDesc = "" +
	"\n" +
	"\x1esro/gameserver/dimension.proto\x12\x0esro.gameserver\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x11sro/globals.proto\"{\n" +
	"\x16CreateDimensionRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion\x12\x17\n" +
	"\amap_ids\x18\x03 \x03(\tR\x06mapIds\x12\x1a\n" +
	"\blocation\x18\x04 \x01(\tR\blocation\"L\n" +
	"\x19DuplicateDimensionRequest\x12\x1b\n" +
	"\ttarget_id\x18\x01 \x01(\tR\btargetId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"\xf3\x01\n" +
	"\x14EditDimensionRequest\x12\x1b\n" +
	"\ttarget_id\x18\x01 \x01(\tR\btargetId\x12\x14\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x12\x1a\n" +
	"\aversion\x18\x03 \x01(\tH\x01R\aversion\x12\x1b\n" +
	"\tedit_maps\x18\x04 \x01(\bR\beditMaps\x12\x17\n" +
	"\amap_ids\x18\x05 \x03(\tR\x06mapIds\x12\x1c\n" +
	"\blocation\x18\b \x01(\tH\x02R\blocationB\x0f\n" +
	"\roptional_nameB\x12\n" +
	"\x10optional_versionB\x13\n" +
	"\x11optional_location\"~\n" +
	"\tDimension\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n" +
	"\aversion\x18\x03 \x01(\tR\aversion\x12\x17\n" +
	"\amap_ids\x18\x04 \x03(\tR\x06mapIds\x12\x1a\n" +
	"\blocation\x18\x06 \x01(\tR\blocation\"G\n" +
	"\n" +
	"Dimensions\x129\n" +
	"\n" +
	"dimensions\x18\x01 \x03(\v2\x19.sro.gameserver.DimensionR\n" +
	"dimensions\"j\n" +
	"\x18GameServerUpdateResponse\x12&\n" +
	"\x0fmap_ids_created\x18\x01 \x03(\tR\rmapIdsCreated\x12&\n" +
	"\x0fmap_ids_deleted\x18\x02 \x03(\tR\rmapIdsDeleted2\x95\x06\n" +
	"\x10DimensionService\x12X\n" +
	"\fGetDimension\x12\r.sro.TargetId\x1a\x19.sro.gameserver.Dimension\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/dimensions/id/{id}\x12[\n" +
	"\rGetDimensions\x12\x16.google.protobuf.Empty\x1a\x1a.sro.gameserver.Dimensions\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/dimensions\x12o\n" +
	"\x0fCreateDimension\x12&.sro.gameserver.CreateDimensionRequest\x1a\x19.sro.gameserver.Dimension\"\x19\x82\xd3\xe4\x93\x02\x13:\x01*\"\x0e/v1/dimensions\x12\x8e\x01\n" +
	"\x12DuplicateDimension\x12).sro.gameserver.DuplicateDimensionRequest\x1a\x19.sro.gameserver.Dimension\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/v1/dimensions/id/{target_id}/duplicate\x12z\n" +
	"\rEditDimension\x12$.sro.gameserver.EditDimensionRequest\x1a\x19.sro.gameserver.Dimension\"(\x82\xd3\xe4\x93\x02\":\x01*\x1a\x1d/v1/dimensions/id/{target_id}\x12X\n" +
	"\x0fDeleteDimension\x12\r.sro.TargetId\x1a\x16.google.protobuf.Empty\"\x1e\x82\xd3\xe4\x93\x02\x18*\x16/v1/dimensions/id/{id}\x12r\n" +
	"\x10GameServerUpdate\x12\r.sro.TargetId\x1a(.sro.gameserver.GameServerUpdateResponse\"%\x82\xd3\xe4\x93\x02\x1f\"\x1d/v1/dimensions/id/{id}/updateB9Z7github.com/ShatteredRealms/gameserver-service/pkg/pb;pbb\x06proto3"

var (
	file_sro_gameserver_dimension_proto_rawDescOnce sync.Once
	file_sro_gameserver_dimension_proto_rawDescData []byte
)

func file_sro_gameserver_dimension_proto_rawDescGZIP() []byte {
	file_sro_gameserver_dimension_proto_rawDescOnce.Do(func() {
		file_sro_gameserver_dimension_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sro_gameserver_dimension_proto_rawDesc), len(file_sro_gameserver_dimension_proto_rawDesc)))
	})
	return file_sro_gameserver_dimension_proto_rawDescData
}

var file_sro_gameserver_dimension_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sro_gameserver_dimension_proto_goTypes = []any{
	(*CreateDimensionRequest)(nil),    // 0: sro.gameserver.CreateDimensionRequest
	(*DuplicateDimensionRequest)(nil), // 1: sro.gameserver.DuplicateDimensionRequest
	(*EditDimensionRequest)(nil),      // 2: sro.gameserver.EditDimensionRequest
	(*Dimension)(nil),                 // 3: sro.gameserver.Dimension
	(*Dimensions)(nil),                // 4: sro.gameserver.Dimensions
	(*GameServerUpdateResponse)(nil),  // 5: sro.gameserver.GameServerUpdateResponse
	(*pb.TargetId)(nil),               // 6: sro.TargetId
	(*emptypb.Empty)(nil),             // 7: google.protobuf.Empty
}
var file_sro_gameserver_dimension_proto_depIdxs = []int32{
	3, // 0: sro.gameserver.Dimensions.dimensions:type_name -> sro.gameserver.Dimension
	6, // 1: sro.gameserver.DimensionService.GetDimension:input_type -> sro.TargetId
	7, // 2: sro.gameserver.DimensionService.GetDimensions:input_type -> google.protobuf.Empty
	0, // 3: sro.gameserver.DimensionService.CreateDimension:input_type -> sro.gameserver.CreateDimensionRequest
	1, // 4: sro.gameserver.DimensionService.DuplicateDimension:input_type -> sro.gameserver.DuplicateDimensionRequest
	2, // 5: sro.gameserver.DimensionService.EditDimension:input_type -> sro.gameserver.EditDimensionRequest
	6, // 6: sro.gameserver.DimensionService.DeleteDimension:input_type -> sro.TargetId
	6, // 7: sro.gameserver.DimensionService.GameServerUpdate:input_type -> sro.TargetId
	3, // 8: sro.gameserver.DimensionService.GetDimension:output_type -> sro.gameserver.Dimension
	4, // 9: sro.gameserver.DimensionService.GetDimensions:output_type -> sro.gameserver.Dimensions
	3, // 10: sro.gameserver.DimensionService.CreateDimension:output_type -> sro.gameserver.Dimension
	3, // 11: sro.gameserver.DimensionService.DuplicateDimension:output_type -> sro.gameserver.Dimension
	3, // 12: sro.gameserver.DimensionService.EditDimension:output_type -> sro.gameserver.Dimension
	7, // 13: sro.gameserver.DimensionService.DeleteDimension:output_type -> google.protobuf.Empty
	5, // 14: sro.gameserver.DimensionService.GameServerUpdate:output_type -> sro.gameserver.GameServerUpdateResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sro_gameserver_dimension_proto_init() }
func file_sro_gameserver_dimension_proto_init() {
	if File_sro_gameserver_dimension_proto != nil {
		return
	}
	file_sro_gameserver_dimension_proto_msgTypes[2].OneofWrappers = []any{
		(*EditDimensionRequest_Name)(nil),
		(*EditDimensionRequest_Version)(nil),
		(*EditDimensionRequest_Location)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sro_gameserver_dimension_proto_rawDesc), len(file_sro_gameserver_dimension_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sro_gameserver_dimension_proto_goTypes,
		DependencyIndexes: file_sro_gameserver_dimension_proto_depIdxs,
		MessageInfos:      file_sro_gameserver_dimension_proto_msgTypes,
	}.Build()
	File_sro_gameserver_dimension_proto = out.File
	file_sro_gameserver_dimension_proto_goTypes = nil
	file_sro_gameserver_dimension_proto_depIdxs = nil
}
