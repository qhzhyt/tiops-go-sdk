// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tiops/common/models/data_source.proto

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

type DataSourceInfo struct {
	XId                  string            `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Type                 string            `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	SupportOutput        bool              `protobuf:"varint,4,opt,name=supportOutput,proto3" json:"supportOutput,omitempty" bson:"supportOutput"`
	Configs              map[string]string `protobuf:"bytes,5,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"configs"`
	ServiceName          string            `protobuf:"bytes,6,opt,name=serviceName,proto3" json:"serviceName,omitempty" bson:"serviceName"`
	Tags                 []string          `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" bson:"tags"`
	Description          string            `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	CreatedBy            string            `protobuf:"bytes,9,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	CreatedTime          int64             `protobuf:"varint,10,opt,name=createdTime,proto3" json:"createdTime,omitempty" bson:"createdTime"`
	UpdatedTime          int64             `protobuf:"varint,11,opt,name=updatedTime,proto3" json:"updatedTime,omitempty" bson:"updatedTime"`
	Readme               string            `protobuf:"bytes,12,opt,name=readme,proto3" json:"readme,omitempty" bson:"readme"`
	Url                  string            `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty" bson:"url"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *DataSourceInfo) Reset()         { *m = DataSourceInfo{} }
func (m *DataSourceInfo) String() string { return proto.CompactTextString(m) }
func (*DataSourceInfo) ProtoMessage()    {}
func (*DataSourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5867aca64158c3d, []int{0}
}
func (m *DataSourceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSourceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSourceInfo.Merge(m, src)
}
func (m *DataSourceInfo) XXX_Size() int {
	return m.Size()
}
func (m *DataSourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DataSourceInfo proto.InternalMessageInfo

func (m *DataSourceInfo) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *DataSourceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataSourceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DataSourceInfo) GetSupportOutput() bool {
	if m != nil {
		return m.SupportOutput
	}
	return false
}

func (m *DataSourceInfo) GetConfigs() map[string]string {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *DataSourceInfo) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *DataSourceInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DataSourceInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataSourceInfo) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *DataSourceInfo) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *DataSourceInfo) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *DataSourceInfo) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *DataSourceInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*DataSourceInfo)(nil), "models.DataSourceInfo")
	proto.RegisterMapType((map[string]string)(nil), "models.DataSourceInfo.ConfigsEntry")
}

func init() {
	proto.RegisterFile("tiops/common/models/data_source.proto", fileDescriptor_a5867aca64158c3d)
}

var fileDescriptor_a5867aca64158c3d = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0xcb, 0xd3, 0x30,
	0x1c, 0xc6, 0xcd, 0xf2, 0xbe, 0x7d, 0xb7, 0x6c, 0x53, 0x09, 0x32, 0xc2, 0x90, 0x12, 0xa6, 0x42,
	0x4f, 0x1d, 0xe8, 0x45, 0x06, 0x5e, 0xa6, 0x1e, 0x76, 0x51, 0xa8, 0xde, 0x47, 0xd6, 0x64, 0x25,
	0xb8, 0x36, 0x21, 0x49, 0x07, 0xbd, 0xfb, 0x21, 0xc4, 0x83, 0x9f, 0xc7, 0xa3, 0xdf, 0x40, 0x99,
	0x5f, 0x44, 0x92, 0x6c, 0xda, 0x81, 0xb7, 0xe7, 0xff, 0xcb, 0xf3, 0xb4, 0x4f, 0x9b, 0x3f, 0x7a,
	0xe6, 0xa4, 0xd2, 0x76, 0x59, 0xaa, 0xba, 0x56, 0xcd, 0xb2, 0x56, 0x5c, 0x1c, 0xec, 0x92, 0x33,
	0xc7, 0xb6, 0x56, 0xb5, 0xa6, 0x14, 0xb9, 0x36, 0xca, 0x29, 0x9c, 0xc4, 0x93, 0x39, 0xfd, 0x9f,
	0xdd, 0xb1, 0xaa, 0x12, 0x26, 0x3a, 0x17, 0x3f, 0x21, 0xba, 0xff, 0x86, 0x39, 0xf6, 0x21, 0xc4,
	0x37, 0xcd, 0x5e, 0x61, 0x8a, 0xe0, 0x56, 0x72, 0x02, 0x28, 0xc8, 0x46, 0xeb, 0x07, 0x5f, 0x3f,
	0x7f, 0x83, 0x68, 0x67, 0x55, 0xb3, 0x5a, 0x6c, 0x25, 0x5f, 0x14, 0x83, 0x0d, 0xc7, 0x18, 0xdd,
	0x34, 0xac, 0x16, 0x64, 0xe0, 0x2d, 0x45, 0xd0, 0x9e, 0xb9, 0x4e, 0x0b, 0x02, 0x23, 0xf3, 0x1a,
	0x3f, 0x45, 0x53, 0xdb, 0x6a, 0xad, 0x8c, 0x7b, 0xdf, 0x3a, 0xdd, 0x3a, 0x72, 0x43, 0x41, 0x36,
	0x2c, 0xae, 0x21, 0x7e, 0x85, 0xee, 0x4a, 0xd5, 0xec, 0x65, 0x65, 0xc9, 0x2d, 0x85, 0xd9, 0xf8,
	0xf9, 0x93, 0x3c, 0x36, 0xcd, 0xaf, 0x8b, 0xe5, 0xaf, 0xa3, 0xeb, 0x6d, 0xe3, 0x4c, 0x57, 0x5c,
	0x32, 0x98, 0xa2, 0xb1, 0x15, 0xe6, 0x28, 0x4b, 0xf1, 0xce, 0x77, 0x4a, 0xc2, 0xfb, 0xfb, 0x28,
	0x54, 0x63, 0x95, 0x25, 0x77, 0x14, 0x86, 0x6a, 0x2c, 0xa6, 0xb8, 0xb0, 0xa5, 0x91, 0xda, 0x49,
	0xd5, 0x90, 0x61, 0x4c, 0xf5, 0x10, 0x7e, 0x8c, 0x46, 0xa5, 0x11, 0xcc, 0x09, 0xbe, 0xee, 0xc8,
	0x28, 0x9c, 0xff, 0x03, 0x3e, 0x7f, 0x1e, 0x3e, 0xca, 0x5a, 0x10, 0x44, 0x41, 0x06, 0x8b, 0x3e,
	0xf2, 0x8e, 0x56, 0xf3, 0xbf, 0x8e, 0x71, 0x74, 0xf4, 0x10, 0x9e, 0xa1, 0xc4, 0x08, 0xc6, 0x6b,
	0x41, 0x26, 0xe1, 0xf1, 0xe7, 0x09, 0x3f, 0x44, 0xb0, 0x35, 0x07, 0x32, 0x0d, 0xd0, 0xcb, 0xf9,
	0x0a, 0x4d, 0xfa, 0x1f, 0xef, 0x1d, 0x9f, 0x44, 0x17, 0xaf, 0xa8, 0xf0, 0x12, 0x3f, 0x42, 0xb7,
	0x47, 0x76, 0x68, 0x2f, 0x77, 0x12, 0x87, 0xd5, 0xe0, 0x25, 0x58, 0x67, 0xdf, 0x4f, 0x29, 0xf8,
	0x71, 0x4a, 0xc1, 0xaf, 0x53, 0x0a, 0xbe, 0xfc, 0x4e, 0xef, 0xa1, 0x99, 0x54, 0x79, 0x58, 0x8c,
	0x3c, 0x2e, 0xc6, 0xf9, 0x77, 0xef, 0x92, 0xb0, 0x12, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x01, 0x8a, 0x9f, 0x92, 0x65, 0x02, 0x00, 0x00,
}

func (m *DataSourceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSourceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSourceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Readme) > 0 {
		i -= len(m.Readme)
		copy(dAtA[i:], m.Readme)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.Readme)))
		i--
		dAtA[i] = 0x62
	}
	if m.UpdatedTime != 0 {
		i = encodeVarintDataSource(dAtA, i, uint64(m.UpdatedTime))
		i--
		dAtA[i] = 0x58
	}
	if m.CreatedTime != 0 {
		i = encodeVarintDataSource(dAtA, i, uint64(m.CreatedTime))
		i--
		dAtA[i] = 0x50
	}
	if len(m.CreatedBy) > 0 {
		i -= len(m.CreatedBy)
		copy(dAtA[i:], m.CreatedBy)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.CreatedBy)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintDataSource(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Configs) > 0 {
		for k := range m.Configs {
			v := m.Configs[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintDataSource(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintDataSource(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintDataSource(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.SupportOutput {
		i--
		if m.SupportOutput {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XId) > 0 {
		i -= len(m.XId)
		copy(dAtA[i:], m.XId)
		i = encodeVarintDataSource(dAtA, i, uint64(len(m.XId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataSource(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataSource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataSourceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XId)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	if m.SupportOutput {
		n += 2
	}
	if len(m.Configs) > 0 {
		for k, v := range m.Configs {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDataSource(uint64(len(k))) + 1 + len(v) + sovDataSource(uint64(len(v)))
			n += mapEntrySize + 1 + sovDataSource(uint64(mapEntrySize))
		}
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovDataSource(uint64(l))
		}
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	l = len(m.CreatedBy)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	if m.CreatedTime != 0 {
		n += 1 + sovDataSource(uint64(m.CreatedTime))
	}
	if m.UpdatedTime != 0 {
		n += 1 + sovDataSource(uint64(m.UpdatedTime))
	}
	l = len(m.Readme)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovDataSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDataSource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataSource(x uint64) (n int) {
	return sovDataSource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataSourceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataSource
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
			return fmt.Errorf("proto: DataSourceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSourceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
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
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportOutput", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
			m.SupportOutput = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Configs == nil {
				m.Configs = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDataSource
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDataSource
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDataSource
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthDataSource
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDataSource
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDataSource
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthDataSource
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDataSource(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthDataSource
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Configs[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedTime", wireType)
			}
			m.CreatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedTime", wireType)
			}
			m.UpdatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Readme", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Readme = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataSource
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
				return ErrInvalidLengthDataSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataSource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataSource
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
func skipDataSource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataSource
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
					return 0, ErrIntOverflowDataSource
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
					return 0, ErrIntOverflowDataSource
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
				return 0, ErrInvalidLengthDataSource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataSource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataSource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataSource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataSource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataSource = fmt.Errorf("proto: unexpected end of group")
)
