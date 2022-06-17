// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: tiops/common/models/notebook.proto

package models

import (
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

type NotebookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	CreatedBy      string            `protobuf:"bytes,2,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	Image          string            `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty" bson:"image"`
	RepoType       int32             `protobuf:"varint,4,opt,name=repoType,proto3" json:"repoType,omitempty" bson:"repoType"`
	RepoName       string            `protobuf:"bytes,5,opt,name=repoName,proto3" json:"repoName,omitempty" bson:"repoName"`
	Token          string            `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty" bson:"token"`
	Url            string            `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty" bson:"url"`
	UsePathRewrite bool              `protobuf:"varint,8,opt,name=usePathRewrite,proto3" json:"usePathRewrite,omitempty" bson:"usePathRewrite"`
	Description    string            `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	DatasetVolumes []*DatasetVolume  `protobuf:"bytes,10,rep,name=datasetVolumes,proto3" json:"datasetVolumes,omitempty" bson:"datasetVolumes"`
	WorkDir        string            `protobuf:"bytes,11,opt,name=workDir,proto3" json:"workDir,omitempty" bson:"workDir"`
	Env            map[string]string `protobuf:"bytes,12,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"env"`
	Cpu            string            `protobuf:"bytes,13,opt,name=cpu,proto3" json:"cpu,omitempty" bson:"cpu"`
	Memory         string            `protobuf:"bytes,14,opt,name=memory,proto3" json:"memory,omitempty" bson:"memory"`
	Gpu            string            `protobuf:"bytes,15,opt,name=gpu,proto3" json:"gpu,omitempty" bson:"gpu"`
	Ports          []string          `protobuf:"bytes,16,rep,name=ports,proto3" json:"ports,omitempty" bson:"ports"`
	Ready          bool              `protobuf:"varint,17,opt,name=ready,proto3" json:"ready,omitempty" bson:"ready"`
	Conditions     []string          `protobuf:"bytes,18,rep,name=conditions,proto3" json:"conditions,omitempty" bson:"conditions"`
	CreatedTime    int64             `protobuf:"varint,19,opt,name=createdTime,proto3" json:"createdTime,omitempty" bson:"createdTime"`
	ReadyTime      int64             `protobuf:"varint,20,opt,name=readyTime,proto3" json:"readyTime,omitempty" bson:"readyTime"`
}

func (x *NotebookInfo) Reset() {
	*x = NotebookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_notebook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotebookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotebookInfo) ProtoMessage() {}

func (x *NotebookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_notebook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotebookInfo.ProtoReflect.Descriptor instead.
func (*NotebookInfo) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_notebook_proto_rawDescGZIP(), []int{0}
}

func (x *NotebookInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotebookInfo) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *NotebookInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *NotebookInfo) GetRepoType() int32 {
	if x != nil {
		return x.RepoType
	}
	return 0
}

func (x *NotebookInfo) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *NotebookInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *NotebookInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NotebookInfo) GetUsePathRewrite() bool {
	if x != nil {
		return x.UsePathRewrite
	}
	return false
}

func (x *NotebookInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotebookInfo) GetDatasetVolumes() []*DatasetVolume {
	if x != nil {
		return x.DatasetVolumes
	}
	return nil
}

func (x *NotebookInfo) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *NotebookInfo) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *NotebookInfo) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *NotebookInfo) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *NotebookInfo) GetGpu() string {
	if x != nil {
		return x.Gpu
	}
	return ""
}

func (x *NotebookInfo) GetPorts() []string {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *NotebookInfo) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *NotebookInfo) GetConditions() []string {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *NotebookInfo) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *NotebookInfo) GetReadyTime() int64 {
	if x != nil {
		return x.ReadyTime
	}
	return 0
}

var File_tiops_common_models_notebook_proto protoreflect.FileDescriptor

var file_tiops_common_models_notebook_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x20, 0x74, 0x69,
	0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x05, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x75, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x12, 0x2f, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x6e,
	0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x70, 0x75, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2d, 0x0a, 0x16, 0x69,
	0x6f, 0x2e, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5a, 0x13, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tiops_common_models_notebook_proto_rawDescOnce sync.Once
	file_tiops_common_models_notebook_proto_rawDescData = file_tiops_common_models_notebook_proto_rawDesc
)

func file_tiops_common_models_notebook_proto_rawDescGZIP() []byte {
	file_tiops_common_models_notebook_proto_rawDescOnce.Do(func() {
		file_tiops_common_models_notebook_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiops_common_models_notebook_proto_rawDescData)
	})
	return file_tiops_common_models_notebook_proto_rawDescData
}

var file_tiops_common_models_notebook_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tiops_common_models_notebook_proto_goTypes = []interface{}{
	(*NotebookInfo)(nil),  // 0: models.NotebookInfo
	nil,                   // 1: models.NotebookInfo.EnvEntry
	(*DatasetVolume)(nil), // 2: models.DatasetVolume
}
var file_tiops_common_models_notebook_proto_depIdxs = []int32{
	2, // 0: models.NotebookInfo.datasetVolumes:type_name -> models.DatasetVolume
	1, // 1: models.NotebookInfo.env:type_name -> models.NotebookInfo.EnvEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tiops_common_models_notebook_proto_init() }
func file_tiops_common_models_notebook_proto_init() {
	if File_tiops_common_models_notebook_proto != nil {
		return
	}
	file_tiops_common_models_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tiops_common_models_notebook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotebookInfo); i {
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
			RawDescriptor: file_tiops_common_models_notebook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tiops_common_models_notebook_proto_goTypes,
		DependencyIndexes: file_tiops_common_models_notebook_proto_depIdxs,
		MessageInfos:      file_tiops_common_models_notebook_proto_msgTypes,
	}.Build()
	File_tiops_common_models_notebook_proto = out.File
	file_tiops_common_models_notebook_proto_rawDesc = nil
	file_tiops_common_models_notebook_proto_goTypes = nil
	file_tiops_common_models_notebook_proto_depIdxs = nil
}
