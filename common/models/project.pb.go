// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tiops/common/models/project.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProjectType int32

const (
	ProjectType_BasicProject  ProjectType = 0
	ProjectType_EngineProject ProjectType = 1
	ProjectType_CustomProject ProjectType = 2
)

var ProjectType_name = map[int32]string{
	0: "BasicProject",
	1: "EngineProject",
	2: "CustomProject",
}

var ProjectType_value = map[string]int32{
	"BasicProject":  0,
	"EngineProject": 1,
	"CustomProject": 2,
}

func (x ProjectType) String() string {
	return proto.EnumName(ProjectType_name, int32(x))
}

func (ProjectType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ee799472a6c885c, []int{0}
}

type DatasetVolume struct {
	DatasetId            string   `protobuf:"bytes,1,opt,name=datasetId,proto3" json:"datasetId,omitempty" bson:"datasetId"`
	MountPath            string   `protobuf:"bytes,2,opt,name=mountPath,proto3" json:"mountPath,omitempty" bson:"mountPath"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DatasetVolume) Reset()         { *m = DatasetVolume{} }
func (m *DatasetVolume) String() string { return proto.CompactTextString(m) }
func (*DatasetVolume) ProtoMessage()    {}
func (*DatasetVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee799472a6c885c, []int{0}
}
func (m *DatasetVolume) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DatasetVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DatasetVolume.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DatasetVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetVolume.Merge(m, src)
}
func (m *DatasetVolume) XXX_Size() int {
	return m.Size()
}
func (m *DatasetVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetVolume.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetVolume proto.InternalMessageInfo

func (m *DatasetVolume) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *DatasetVolume) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

type ProjectConfig struct {
	UseVolume            bool             `protobuf:"varint,1,opt,name=useVolume,proto3" json:"useVolume,omitempty" bson:"useVolume"`
	UseGPU               bool             `protobuf:"varint,2,opt,name=useGPU,proto3" json:"useGPU,omitempty" bson:"useGPU"`
	DatasetVolumes       []*DatasetVolume `protobuf:"bytes,3,rep,name=datasetVolumes,proto3" json:"datasetVolumes,omitempty" bson:"datasetVolumes"`
	UseAllDatasetVolumes bool             `protobuf:"varint,4,opt,name=useAllDatasetVolumes,proto3" json:"useAllDatasetVolumes,omitempty" bson:"useAllDatasetVolumes"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" bson:"-"`
	XXX_unrecognized     []byte           `json:"-" bson:"-"`
	XXX_sizecache        int32            `json:"-" bson:"-"`
}

func (m *ProjectConfig) Reset()         { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()    {}
func (*ProjectConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee799472a6c885c, []int{1}
}
func (m *ProjectConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectConfig.Merge(m, src)
}
func (m *ProjectConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProjectConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectConfig proto.InternalMessageInfo

func (m *ProjectConfig) GetUseVolume() bool {
	if m != nil {
		return m.UseVolume
	}
	return false
}

func (m *ProjectConfig) GetUseGPU() bool {
	if m != nil {
		return m.UseGPU
	}
	return false
}

func (m *ProjectConfig) GetDatasetVolumes() []*DatasetVolume {
	if m != nil {
		return m.DatasetVolumes
	}
	return nil
}

func (m *ProjectConfig) GetUseAllDatasetVolumes() bool {
	if m != nil {
		return m.UseAllDatasetVolumes
	}
	return false
}

type ProjectInfo struct {
	XId                  string         `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Lang                 string         `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty" bson:"lang"`
	Title                string         `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty" bson:"title"`
	Description          string         `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	CreatedBy            string         `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	CreatedTime          int64          `protobuf:"varint,7,opt,name=createdTime,proto3" json:"createdTime,omitempty" bson:"createdTime"`
	UpdatedTime          int64          `protobuf:"varint,8,opt,name=updatedTime,proto3" json:"updatedTime,omitempty" bson:"updatedTime"`
	Readme               string         `protobuf:"bytes,9,opt,name=readme,proto3" json:"readme,omitempty" bson:"readme"`
	Actions              []*ActionInfo  `protobuf:"bytes,10,rep,name=actions,proto3" json:"actions,omitempty" bson:"actions"`
	Engines              []*ActionInfo  `protobuf:"bytes,11,rep,name=engines,proto3" json:"engines,omitempty" bson:"engines"`
	Config               *ProjectConfig `protobuf:"bytes,12,opt,name=config,proto3" json:"config,omitempty" bson:"config"`
	Type                 ProjectType    `protobuf:"varint,13,opt,name=type,proto3,enum=models.ProjectType" json:"type,omitempty" bson:"type"`
	IsPublic             bool           `protobuf:"varint,14,opt,name=isPublic,proto3" json:"isPublic,omitempty" bson:"isPublic"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-" bson:"-"`
	XXX_unrecognized     []byte         `json:"-" bson:"-"`
	XXX_sizecache        int32          `json:"-" bson:"-"`
}

func (m *ProjectInfo) Reset()         { *m = ProjectInfo{} }
func (m *ProjectInfo) String() string { return proto.CompactTextString(m) }
func (*ProjectInfo) ProtoMessage()    {}
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee799472a6c885c, []int{2}
}
func (m *ProjectInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectInfo.Merge(m, src)
}
func (m *ProjectInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProjectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectInfo proto.InternalMessageInfo

func (m *ProjectInfo) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *ProjectInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProjectInfo) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *ProjectInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProjectInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProjectInfo) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ProjectInfo) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *ProjectInfo) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *ProjectInfo) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *ProjectInfo) GetActions() []*ActionInfo {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *ProjectInfo) GetEngines() []*ActionInfo {
	if m != nil {
		return m.Engines
	}
	return nil
}

func (m *ProjectInfo) GetConfig() *ProjectConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ProjectInfo) GetType() ProjectType {
	if m != nil {
		return m.Type
	}
	return ProjectType_BasicProject
}

func (m *ProjectInfo) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

func init() {
	proto.RegisterEnum("models.ProjectType", ProjectType_name, ProjectType_value)
	proto.RegisterType((*DatasetVolume)(nil), "models.DatasetVolume")
	proto.RegisterType((*ProjectConfig)(nil), "models.ProjectConfig")
	proto.RegisterType((*ProjectInfo)(nil), "models.ProjectInfo")
}

func init() { proto.RegisterFile("tiops/common/models/project.proto", fileDescriptor_7ee799472a6c885c) }

var fileDescriptor_7ee799472a6c885c = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0x87, 0xeb, 0x38, 0x4d, 0x93, 0x93, 0x26, 0xb7, 0x77, 0x28, 0xd5, 0xa8, 0x42, 0x91, 0xc9,
	0x86, 0x08, 0x41, 0x22, 0x85, 0x1d, 0x12, 0x8b, 0xa6, 0xad, 0x50, 0xc4, 0x26, 0xb2, 0x0a, 0xdb,
	0xca, 0xf1, 0x9c, 0x86, 0x41, 0xf6, 0x8c, 0xe5, 0x19, 0x2f, 0xb2, 0xe7, 0x21, 0x10, 0x0b, 0x9e,
	0x83, 0x47, 0x60, 0xc9, 0x96, 0x1d, 0x0a, 0x2f, 0x82, 0xe6, 0x4f, 0xd2, 0xa4, 0x0a, 0x3b, 0x9f,
	0x6f, 0x3e, 0x1f, 0x1f, 0xff, 0x66, 0x06, 0x9e, 0x6a, 0x2e, 0x0b, 0x35, 0x4a, 0x65, 0x9e, 0x4b,
	0x31, 0xca, 0x25, 0xc3, 0x4c, 0x8d, 0x8a, 0x52, 0x7e, 0xc2, 0x54, 0x0f, 0x8b, 0x52, 0x6a, 0x49,
	0x1a, 0x8e, 0x9e, 0x47, 0xfb, 0x54, 0x9d, 0x2c, 0x16, 0x58, 0x3a, 0x73, 0xbf, 0x91, 0xa4, 0x9a,
	0x4b, 0xe1, 0x8c, 0xfe, 0x3b, 0xe8, 0x5c, 0x25, 0x3a, 0x51, 0xa8, 0x3f, 0xc8, 0xac, 0xca, 0x91,
	0x3c, 0x81, 0x16, 0x73, 0x60, 0xca, 0x68, 0x10, 0x05, 0x83, 0x56, 0x7c, 0x0f, 0xcc, 0x6a, 0x2e,
	0x2b, 0xa1, 0x67, 0x89, 0xfe, 0x48, 0x6b, 0x6e, 0x75, 0x03, 0xfa, 0xdf, 0x03, 0xe8, 0xcc, 0xdc,
	0xa8, 0x97, 0x52, 0xdc, 0xf1, 0x85, 0xf1, 0x2b, 0x85, 0xae, 0xb5, 0xed, 0xd6, 0x8c, 0xef, 0x01,
	0x39, 0x83, 0x46, 0xa5, 0xf0, 0xed, 0xec, 0xbd, 0x6d, 0xd5, 0x8c, 0x7d, 0x45, 0xde, 0x40, 0x97,
	0x6d, 0x0f, 0xa5, 0x68, 0x18, 0x85, 0x83, 0xf6, 0xf8, 0xf1, 0xd0, 0xfd, 0xc2, 0x70, 0x67, 0xe4,
	0xf8, 0x81, 0x4c, 0xc6, 0x70, 0x5a, 0x29, 0xbc, 0xc8, 0xb2, 0xab, 0xdd, 0x26, 0x75, 0xfb, 0x91,
	0xbd, 0x6b, 0xfd, 0x5f, 0x21, 0xb4, 0xfd, 0xe8, 0x53, 0x71, 0x27, 0x49, 0x04, 0xe1, 0x2d, 0xf7,
	0x01, 0x4c, 0xfe, 0xfb, 0xfa, 0xf9, 0x5b, 0x08, 0x73, 0x25, 0xc5, 0xeb, 0xfe, 0x2d, 0x67, 0xfd,
	0xb8, 0x36, 0x65, 0x84, 0x40, 0x5d, 0x24, 0x39, 0xfa, 0x14, 0xec, 0xb3, 0x61, 0x59, 0x22, 0x16,
	0x34, 0x74, 0xcc, 0x3c, 0x93, 0x53, 0x38, 0xd4, 0x5c, 0x67, 0x68, 0x3f, 0xdf, 0x8a, 0x5d, 0x41,
	0x22, 0x68, 0x33, 0x54, 0x69, 0xc9, 0x0b, 0xb3, 0x19, 0xf4, 0xd0, 0xae, 0x6d, 0x23, 0x13, 0x5d,
	0x5a, 0x62, 0xa2, 0x91, 0x4d, 0x96, 0xb4, 0xe1, 0xa2, 0xde, 0x00, 0xf3, 0xbe, 0x2f, 0x6e, 0x78,
	0x8e, 0xf4, 0x28, 0x0a, 0x06, 0x61, 0xbc, 0x8d, 0x8c, 0x51, 0x15, 0x6c, 0x63, 0x34, 0x9d, 0xb1,
	0x85, 0x4c, 0xfc, 0x25, 0x26, 0x2c, 0x47, 0xda, 0xb2, 0xed, 0x7d, 0x45, 0x5e, 0xc0, 0x91, 0x3b,
	0x23, 0x8a, 0x82, 0xcd, 0x9d, 0xac, 0x73, 0xbf, 0xb0, 0xd8, 0x04, 0x14, 0xaf, 0x15, 0x63, 0xa3,
	0x58, 0x70, 0x81, 0x8a, 0xb6, 0xff, 0x6d, 0x7b, 0x85, 0xbc, 0x84, 0x46, 0x6a, 0x8f, 0x06, 0x3d,
	0x8e, 0x82, 0xed, 0x2d, 0xdd, 0x39, 0x37, 0xb1, 0x97, 0xc8, 0x33, 0xa8, 0xeb, 0x65, 0x81, 0xb4,
	0x13, 0x05, 0x83, 0xee, 0xf8, 0xd1, 0x03, 0xf9, 0x66, 0x59, 0x60, 0x6c, 0x05, 0x72, 0x0e, 0x4d,
	0xae, 0x66, 0xd5, 0x3c, 0xe3, 0x29, 0xed, 0xda, 0x7d, 0xde, 0xd4, 0xcf, 0xaf, 0x37, 0x5b, 0x6b,
	0x5e, 0x20, 0x27, 0x70, 0x3c, 0x49, 0x14, 0x4f, 0x3d, 0x3b, 0x39, 0x20, 0xff, 0x43, 0xe7, 0xda,
	0xce, 0xb7, 0x46, 0x81, 0x41, 0x97, 0x95, 0xd2, 0x32, 0x5f, 0xa3, 0xda, 0x64, 0xf0, 0x63, 0xd5,
	0x0b, 0x7e, 0xae, 0x7a, 0xc1, 0xef, 0x55, 0x2f, 0xf8, 0xf2, 0xa7, 0x77, 0x00, 0x67, 0x5c, 0x0e,
	0xed, 0x0d, 0x1b, 0xba, 0x1b, 0xe6, 0xc7, 0x9b, 0x37, 0xec, 0xdd, 0x7a, 0xf5, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xc0, 0x07, 0x53, 0x32, 0xcc, 0x03, 0x00, 0x00,
}

func (m *DatasetVolume) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DatasetVolume) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DatasetVolume) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MountPath) > 0 {
		i -= len(m.MountPath)
		copy(dAtA[i:], m.MountPath)
		i = encodeVarintProject(dAtA, i, uint64(len(m.MountPath)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DatasetId) > 0 {
		i -= len(m.DatasetId)
		copy(dAtA[i:], m.DatasetId)
		i = encodeVarintProject(dAtA, i, uint64(len(m.DatasetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProjectConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UseAllDatasetVolumes {
		i--
		if m.UseAllDatasetVolumes {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.DatasetVolumes) > 0 {
		for iNdEx := len(m.DatasetVolumes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DatasetVolumes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.UseGPU {
		i--
		if m.UseGPU {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.UseVolume {
		i--
		if m.UseVolume {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProjectInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsPublic {
		i--
		if m.IsPublic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	if m.Type != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x68
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.Engines) > 0 {
		for iNdEx := len(m.Engines) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Engines[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Readme) > 0 {
		i -= len(m.Readme)
		copy(dAtA[i:], m.Readme)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Readme)))
		i--
		dAtA[i] = 0x4a
	}
	if m.UpdatedTime != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.UpdatedTime))
		i--
		dAtA[i] = 0x40
	}
	if m.CreatedTime != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.CreatedTime))
		i--
		dAtA[i] = 0x38
	}
	if len(m.CreatedBy) > 0 {
		i -= len(m.CreatedBy)
		copy(dAtA[i:], m.CreatedBy)
		i = encodeVarintProject(dAtA, i, uint64(len(m.CreatedBy)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Lang) > 0 {
		i -= len(m.Lang)
		copy(dAtA[i:], m.Lang)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Lang)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XId) > 0 {
		i -= len(m.XId)
		copy(dAtA[i:], m.XId)
		i = encodeVarintProject(dAtA, i, uint64(len(m.XId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProject(dAtA []byte, offset int, v uint64) int {
	offset -= sovProject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DatasetVolume) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DatasetId)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.MountPath)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProjectConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UseVolume {
		n += 2
	}
	if m.UseGPU {
		n += 2
	}
	if len(m.DatasetVolumes) > 0 {
		for _, e := range m.DatasetVolumes {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if m.UseAllDatasetVolumes {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProjectInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XId)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Lang)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.CreatedBy)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.CreatedTime != 0 {
		n += 1 + sovProject(uint64(m.CreatedTime))
	}
	if m.UpdatedTime != 0 {
		n += 1 + sovProject(uint64(m.UpdatedTime))
	}
	l = len(m.Readme)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if len(m.Engines) > 0 {
		for _, e := range m.Engines {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovProject(uint64(m.Type))
	}
	if m.IsPublic {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProject(x uint64) (n int) {
	return sovProject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DatasetVolume) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DatasetVolume: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DatasetVolume: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatasetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatasetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MountPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MountPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseVolume", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseVolume = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseGPU", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseGPU = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatasetVolumes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatasetVolumes = append(m.DatasetVolumes, &DatasetVolume{})
			if err := m.DatasetVolumes[len(m.DatasetVolumes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseAllDatasetVolumes", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseAllDatasetVolumes = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lang = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedTime", wireType)
			}
			m.CreatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedTime", wireType)
			}
			m.UpdatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Readme", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Readme = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &ActionInfo{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Engines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Engines = append(m.Engines, &ActionInfo{})
			if err := m.Engines[len(m.Engines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &ProjectConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ProjectType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPublic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPublic = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProject
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProject = fmt.Errorf("proto: unexpected end of group")
)
