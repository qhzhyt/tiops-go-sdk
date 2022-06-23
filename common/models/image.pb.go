// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: tiops/common/models/image.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ImageType int32

const (
	ImageType_ActionServer   ImageType = 0
	ImageType_EngineServer   ImageType = 1
	ImageType_ServiceServer  ImageType = 2
	ImageType_NotebookServer ImageType = 3
	ImageType_CustomServer   ImageType = 6
)

// Enum value maps for ImageType.
var (
	ImageType_name = map[int32]string{
		0: "ActionServer",
		1: "EngineServer",
		2: "ServiceServer",
		3: "NotebookServer",
		6: "CustomServer",
	}
	ImageType_value = map[string]int32{
		"ActionServer":   0,
		"EngineServer":   1,
		"ServiceServer":  2,
		"NotebookServer": 3,
		"CustomServer":   6,
	}
)

func (x ImageType) Enum() *ImageType {
	p := new(ImageType)
	*p = x
	return p
}

func (x ImageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageType) Descriptor() protoreflect.EnumDescriptor {
	return file_tiops_common_models_image_proto_enumTypes[0].Descriptor()
}

func (ImageType) Type() protoreflect.EnumType {
	return &file_tiops_common_models_image_proto_enumTypes[0]
}

func (x ImageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageType.Descriptor instead.
func (ImageType) EnumDescriptor() ([]byte, []int) {
	return file_tiops_common_models_image_proto_rawDescGZIP(), []int{0}
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId         string            `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Registry    string            `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty" bson:"registry"`
	Repo        string            `protobuf:"bytes,4,opt,name=repo,proto3" json:"repo,omitempty" bson:"repo"`
	Tag         string            `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty" bson:"tag"`
	Size        int64             `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty" bson:"size"`
	Type        ImageType         `protobuf:"varint,7,opt,name=type,proto3,enum=models.ImageType" json:"type,omitempty" bson:"type"`
	Project     string            `protobuf:"bytes,8,opt,name=project,proto3" json:"project,omitempty" bson:"project"`
	CreatedBy   string            `protobuf:"bytes,9,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	Description string            `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	Public      bool              `protobuf:"varint,11,opt,name=public,proto3" json:"public,omitempty" bson:"public"`
	BuildTime   int64             `protobuf:"varint,12,opt,name=buildTime,proto3" json:"buildTime,omitempty" bson:"buildTime"`
	Extra       map[string]string `protobuf:"bytes,13,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *Image) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Image) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Image) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Image) GetType() ImageType {
	if x != nil {
		return x.Type
	}
	return ImageType_ActionServer
}

func (x *Image) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Image) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Image) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Image) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *Image) GetBuildTime() int64 {
	if x != nil {
		return x.BuildTime
	}
	return 0
}

func (x *Image) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_tiops_common_models_image_proto protoreflect.FileDescriptor

var file_tiops_common_models_image_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x20, 0x74, 0x69, 0x6f, 0x70, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x03, 0x0a, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03, 0x0a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f,
	0x69, 0x64, 0x22, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x68, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x06, 0x42, 0x2d, 0x0a, 0x16,
	0x69, 0x6f, 0x2e, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5a, 0x13, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tiops_common_models_image_proto_rawDescOnce sync.Once
	file_tiops_common_models_image_proto_rawDescData = file_tiops_common_models_image_proto_rawDesc
)

func file_tiops_common_models_image_proto_rawDescGZIP() []byte {
	file_tiops_common_models_image_proto_rawDescOnce.Do(func() {
		file_tiops_common_models_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiops_common_models_image_proto_rawDescData)
	})
	return file_tiops_common_models_image_proto_rawDescData
}

var file_tiops_common_models_image_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tiops_common_models_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tiops_common_models_image_proto_goTypes = []interface{}{
	(ImageType)(0), // 0: models.ImageType
	(*Image)(nil),  // 1: models.Image
	nil,            // 2: models.Image.ExtraEntry
}
var file_tiops_common_models_image_proto_depIdxs = []int32{
	0, // 0: models.Image.type:type_name -> models.ImageType
	2, // 1: models.Image.extra:type_name -> models.Image.ExtraEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tiops_common_models_image_proto_init() }
func file_tiops_common_models_image_proto_init() {
	if File_tiops_common_models_image_proto != nil {
		return
	}
	file_tiops_common_models_tagger_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tiops_common_models_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tiops_common_models_image_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tiops_common_models_image_proto_goTypes,
		DependencyIndexes: file_tiops_common_models_image_proto_depIdxs,
		EnumInfos:         file_tiops_common_models_image_proto_enumTypes,
		MessageInfos:      file_tiops_common_models_image_proto_msgTypes,
	}.Build()
	File_tiops_common_models_image_proto = out.File
	file_tiops_common_models_image_proto_rawDesc = nil
	file_tiops_common_models_image_proto_goTypes = nil
	file_tiops_common_models_image_proto_depIdxs = nil
}
