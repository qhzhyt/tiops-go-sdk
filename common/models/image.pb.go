// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tiops/common/models/image.proto

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

type ImageType int32

const (
	ImageType_ActionServer   ImageType = 0
	ImageType_EngineServer   ImageType = 1
	ImageType_ServiceServer  ImageType = 2
	ImageType_NotebookServer ImageType = 3
	ImageType_CustomServer   ImageType = 6
)

var ImageType_name = map[int32]string{
	0: "ActionServer",
	1: "EngineServer",
	2: "ServiceServer",
	3: "NotebookServer",
	6: "CustomServer",
}

var ImageType_value = map[string]int32{
	"ActionServer":   0,
	"EngineServer":   1,
	"ServiceServer":  2,
	"NotebookServer": 3,
	"CustomServer":   6,
}

func (x ImageType) String() string {
	return proto.EnumName(ImageType_name, int32(x))
}

func (ImageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e751725b8a351bd, []int{0}
}

type Image struct {
	XId                  string            `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Registry             string            `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty" bson:"registry"`
	Repo                 string            `protobuf:"bytes,4,opt,name=repo,proto3" json:"repo,omitempty" bson:"repo"`
	Tag                  string            `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty" bson:"tag"`
	Size_                int64             `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Type                 ImageType         `protobuf:"varint,7,opt,name=type,proto3,enum=models.ImageType" json:"type,omitempty" bson:"type"`
	Project              string            `protobuf:"bytes,8,opt,name=project,proto3" json:"project,omitempty" bson:"project"`
	CreatedBy            string            `protobuf:"bytes,9,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	Description          string            `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	Public               bool              `protobuf:"varint,11,opt,name=public,proto3" json:"public,omitempty" bson:"public"`
	BuildTime            int64             `protobuf:"varint,12,opt,name=buildTime,proto3" json:"buildTime,omitempty" bson:"buildTime"`
	Extra                map[string]string `protobuf:"bytes,13,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e751725b8a351bd, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Image.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return m.Size()
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *Image) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Image) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Image) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Image) GetType() ImageType {
	if m != nil {
		return m.Type
	}
	return ImageType_ActionServer
}

func (m *Image) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Image) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Image) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Image) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *Image) GetBuildTime() int64 {
	if m != nil {
		return m.BuildTime
	}
	return 0
}

func (m *Image) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterEnum("models.ImageType", ImageType_name, ImageType_value)
	proto.RegisterType((*Image)(nil), "models.Image")
	proto.RegisterMapType((map[string]string)(nil), "models.Image.ExtraEntry")
}

func init() { proto.RegisterFile("tiops/common/models/image.proto", fileDescriptor_7e751725b8a351bd) }

var fileDescriptor_7e751725b8a351bd = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x71, 0xe2, 0x26, 0x93, 0xb6, 0xa4, 0x23, 0x84, 0x56, 0x11, 0x0a, 0xab, 0x48,
	0x48, 0x16, 0x07, 0x47, 0x2a, 0x97, 0xaa, 0x37, 0x82, 0x72, 0xe8, 0x85, 0x83, 0xe9, 0xbd, 0xf2,
	0x9f, 0x91, 0x59, 0x1a, 0x7b, 0x57, 0xeb, 0x4d, 0x85, 0x39, 0xf3, 0x10, 0x88, 0x03, 0x4f, 0xc1,
	0x43, 0x70, 0xe4, 0x11, 0x50, 0x78, 0x11, 0xb4, 0xeb, 0xa4, 0x01, 0xa9, 0xb7, 0x6f, 0x7e, 0xdf,
	0xb7, 0xa3, 0x4f, 0x63, 0xc3, 0x0b, 0x2b, 0x95, 0x6e, 0x16, 0xb9, 0xaa, 0x2a, 0x55, 0x2f, 0x2a,
	0x55, 0xd0, 0xba, 0x59, 0xc8, 0x2a, 0x2d, 0x29, 0xd6, 0x46, 0x59, 0x85, 0x61, 0xc7, 0xa6, 0xe2,
	0xb1, 0xa0, 0x4d, 0xcb, 0x92, 0x4c, 0x97, 0x9c, 0xff, 0x08, 0x60, 0x70, 0xed, 0x5e, 0xa2, 0x80,
	0xe0, 0x56, 0x16, 0x9c, 0x09, 0x16, 0x8d, 0x96, 0x4f, 0xbe, 0x7d, 0xf9, 0x1e, 0x40, 0xd6, 0xa8,
	0xfa, 0x6a, 0x7e, 0x2b, 0x8b, 0x79, 0xd2, 0xbb, 0x2e, 0x10, 0xa1, 0x5f, 0xa7, 0x15, 0xf1, 0x9e,
	0x8b, 0x24, 0x5e, 0xe3, 0x14, 0x86, 0x86, 0x4a, 0xd9, 0x58, 0xd3, 0xf2, 0xc0, 0xf3, 0x87, 0xd9,
	0xe5, 0x0d, 0x69, 0xc5, 0xfb, 0x5d, 0xde, 0x69, 0x9c, 0x40, 0x60, 0xd3, 0x92, 0x0f, 0x3c, 0x72,
	0xd2, 0xa5, 0x1a, 0xf9, 0x99, 0x78, 0x28, 0x58, 0x14, 0x24, 0x5e, 0xe3, 0x4b, 0xe8, 0xdb, 0x56,
	0x13, 0x3f, 0x16, 0x2c, 0x3a, 0xbb, 0x38, 0x8f, 0xbb, 0xe6, 0xb1, 0x2f, 0x7a, 0xd3, 0x6a, 0x4a,
	0xbc, 0x8d, 0x1c, 0x8e, 0xb5, 0x51, 0x1f, 0x29, 0xb7, 0x7c, 0xe8, 0x17, 0xee, 0x47, 0x7c, 0x0e,
	0xa3, 0xdc, 0x50, 0x6a, 0xa9, 0x58, 0xb6, 0x7c, 0xe4, 0xbd, 0x03, 0x40, 0x01, 0xe3, 0x82, 0x9a,
	0xdc, 0x48, 0x6d, 0xa5, 0xaa, 0x39, 0x78, 0xff, 0x5f, 0x84, 0xcf, 0x20, 0xd4, 0x9b, 0x6c, 0x2d,
	0x73, 0x3e, 0x16, 0x2c, 0x1a, 0x26, 0xbb, 0xc9, 0xed, 0xcd, 0x36, 0x72, 0x5d, 0xdc, 0xc8, 0x8a,
	0xf8, 0x89, 0x6f, 0x7c, 0x00, 0x18, 0xc3, 0x80, 0x3e, 0x59, 0x93, 0xf2, 0x53, 0x11, 0x44, 0xe3,
	0x0b, 0xfe, 0x5f, 0xef, 0x78, 0xe5, 0xac, 0x55, 0x6d, 0x4d, 0x9b, 0x74, 0xb1, 0xe9, 0x25, 0xc0,
	0x01, 0xba, 0xd3, 0xdc, 0x51, 0xdb, 0x7d, 0x80, 0xc4, 0x49, 0x7c, 0x0a, 0x83, 0xfb, 0x74, 0xbd,
	0xd9, 0x5f, 0xbc, 0x1b, 0xae, 0x7a, 0x97, 0xec, 0xd5, 0x07, 0x18, 0x3d, 0x1c, 0x03, 0x27, 0x70,
	0xf2, 0x26, 0x77, 0xb5, 0xdf, 0x93, 0xb9, 0x27, 0x33, 0x39, 0x72, 0x64, 0x55, 0x97, 0xb2, 0xa6,
	0x1d, 0x61, 0x78, 0x0e, 0xa7, 0x4e, 0xcb, 0x7c, 0x8f, 0x7a, 0x88, 0x70, 0xf6, 0x4e, 0x59, 0xca,
	0x94, 0xba, 0xdb, 0xb1, 0xc0, 0x3d, 0x7c, 0xbb, 0x69, 0xac, 0xaa, 0x76, 0x24, 0x5c, 0x4e, 0x7e,
	0x6e, 0x67, 0xec, 0xd7, 0x76, 0xc6, 0x7e, 0x6f, 0x67, 0xec, 0xeb, 0x9f, 0xd9, 0x51, 0x16, 0xfa,
	0x3f, 0xe7, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x3b, 0x06, 0xb4, 0x86, 0x02, 0x00,
	0x00,
}

func (m *Image) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Image) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Image) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Extra) > 0 {
		for k := range m.Extra {
			v := m.Extra[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintImage(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintImage(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintImage(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.BuildTime != 0 {
		i = encodeVarintImage(dAtA, i, uint64(m.BuildTime))
		i--
		dAtA[i] = 0x60
	}
	if m.Public {
		i--
		if m.Public {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.CreatedBy) > 0 {
		i -= len(m.CreatedBy)
		copy(dAtA[i:], m.CreatedBy)
		i = encodeVarintImage(dAtA, i, uint64(len(m.CreatedBy)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Project) > 0 {
		i -= len(m.Project)
		copy(dAtA[i:], m.Project)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Project)))
		i--
		dAtA[i] = 0x42
	}
	if m.Type != 0 {
		i = encodeVarintImage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x38
	}
	if m.Size_ != 0 {
		i = encodeVarintImage(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Repo) > 0 {
		i -= len(m.Repo)
		copy(dAtA[i:], m.Repo)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Repo)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Registry) > 0 {
		i -= len(m.Registry)
		copy(dAtA[i:], m.Registry)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Registry)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XId) > 0 {
		i -= len(m.XId)
		copy(dAtA[i:], m.XId)
		i = encodeVarintImage(dAtA, i, uint64(len(m.XId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintImage(dAtA []byte, offset int, v uint64) int {
	offset -= sovImage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Image) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XId)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Registry)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Repo)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovImage(uint64(m.Size_))
	}
	if m.Type != 0 {
		n += 1 + sovImage(uint64(m.Type))
	}
	l = len(m.Project)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.CreatedBy)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if m.Public {
		n += 2
	}
	if m.BuildTime != 0 {
		n += 1 + sovImage(uint64(m.BuildTime))
	}
	if len(m.Extra) > 0 {
		for k, v := range m.Extra {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovImage(uint64(len(k))) + 1 + len(v) + sovImage(uint64(len(v)))
			n += mapEntrySize + 1 + sovImage(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovImage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozImage(x uint64) (n int) {
	return sovImage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Image) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImage
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
			return fmt.Errorf("proto: Image: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Image: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
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
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ImageType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Project", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Project = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Public", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
			m.Public = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildTime", wireType)
			}
			m.BuildTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BuildTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Extra == nil {
				m.Extra = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowImage
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
							return ErrIntOverflowImage
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
						return ErrInvalidLengthImage
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthImage
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
							return ErrIntOverflowImage
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
						return ErrInvalidLengthImage
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthImage
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipImage(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthImage
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Extra[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImage
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
func skipImage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImage
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
					return 0, ErrIntOverflowImage
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
					return 0, ErrIntOverflowImage
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
				return 0, ErrInvalidLengthImage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupImage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthImage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthImage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupImage = fmt.Errorf("proto: unexpected end of group")
)
