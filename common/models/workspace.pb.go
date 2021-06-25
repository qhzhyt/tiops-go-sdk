// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tiops/common/models/workspace.proto

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

type Workspace struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	CreatedBy            string   `protobuf:"bytes,3,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	IsPublic             bool     `protobuf:"varint,4,opt,name=isPublic,proto3" json:"isPublic,omitempty" bson:"isPublic"`
	Members              []string `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty" bson:"members"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	CreatedTime          int64    `protobuf:"varint,9,opt,name=createdTime,proto3" json:"createdTime,omitempty" bson:"createdTime"`
	UpdatedTime          int64    `protobuf:"varint,10,opt,name=updatedTime,proto3" json:"updatedTime,omitempty" bson:"updatedTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_0507a3ab829754dc, []int{0}
}
func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(m, src)
}
func (m *Workspace) XXX_Size() int {
	return m.Size()
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Workspace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Workspace) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Workspace) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

func (m *Workspace) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Workspace) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Workspace) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Workspace) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Workspace)(nil), "models.Workspace")
}

func init() {
	proto.RegisterFile("tiops/common/models/workspace.proto", fileDescriptor_0507a3ab829754dc)
}

var fileDescriptor_0507a3ab829754dc = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0xcd, 0xf5, 0xac, 0xd7, 0xcf, 0x41, 0xc8, 0x20, 0xe1, 0x90, 0x12, 0xea, 0xd2, 0xa9,
	0x1d, 0xdc, 0x1c, 0xbb, 0xb9, 0x49, 0x11, 0x1c, 0x8f, 0xb6, 0x09, 0x47, 0xf0, 0xd2, 0xaf, 0x24,
	0x3d, 0xc4, 0xdd, 0x1f, 0x21, 0x0e, 0xfe, 0x1e, 0x47, 0x7f, 0x82, 0xd4, 0x9f, 0xe1, 0x22, 0x4d,
	0xee, 0xce, 0x0e, 0xb7, 0xe5, 0x7d, 0xf2, 0xf0, 0x7e, 0xf0, 0xc2, 0x75, 0xaf, 0xb0, 0xb3, 0x79,
	0x83, 0x5a, 0x63, 0x9b, 0x6b, 0x14, 0x72, 0x63, 0xf3, 0x67, 0x34, 0x4f, 0xb6, 0xab, 0x1a, 0x99,
	0x75, 0x06, 0x7b, 0xa4, 0xa1, 0xe7, 0x4b, 0x7e, 0x4c, 0xf6, 0xc9, 0x9b, 0xc7, 0x8d, 0xbe, 0x5a,
	0xaf, 0xa5, 0xf1, 0x46, 0xf2, 0x4b, 0x20, 0x7a, 0xdc, 0xf7, 0x53, 0x0e, 0xc1, 0x4a, 0x09, 0x46,
	0x38, 0x49, 0xa3, 0xe2, 0xe2, 0xfd, 0xf5, 0x23, 0x80, 0xda, 0x62, 0x7b, 0x9b, 0xac, 0x94, 0x48,
	0xca, 0xd9, 0x9d, 0xa0, 0x14, 0xe6, 0x6d, 0xa5, 0x25, 0x9b, 0x8d, 0x4a, 0xe9, 0xde, 0xf4, 0x0a,
	0xa2, 0xc6, 0xc8, 0xaa, 0x97, 0xa2, 0x78, 0x61, 0x81, 0xfb, 0xf8, 0x07, 0x74, 0x09, 0x0b, 0x65,
	0xef, 0xb7, 0xf5, 0x46, 0x35, 0x6c, 0xce, 0x49, 0xba, 0x28, 0x0f, 0x99, 0x32, 0x38, 0xd3, 0x52,
	0xd7, 0xd2, 0x58, 0x76, 0xca, 0x83, 0x34, 0x2a, 0xf7, 0x91, 0x72, 0x38, 0x17, 0xd2, 0x36, 0x46,
	0x75, 0xbd, 0xc2, 0x96, 0x85, 0xae, 0x75, 0x8a, 0x46, 0x63, 0x77, 0xe4, 0x41, 0x69, 0xc9, 0x22,
	0x4e, 0xd2, 0xa0, 0x9c, 0xa2, 0xd1, 0xd8, 0x76, 0xe2, 0x60, 0x80, 0x37, 0x26, 0xa8, 0x48, 0x3f,
	0x87, 0x98, 0x7c, 0x0d, 0x31, 0xf9, 0x1e, 0x62, 0xf2, 0xf6, 0x13, 0x9f, 0xc0, 0xa5, 0xc2, 0xcc,
	0x8d, 0x96, 0xed, 0x86, 0xf4, 0xa3, 0xd5, 0xa1, 0x9b, 0xeb, 0xe6, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0xe3, 0x0e, 0x84, 0xa1, 0x01, 0x00, 0x00,
}

func (m *Workspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Workspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Workspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdatedTime != 0 {
		i = encodeVarintWorkspace(dAtA, i, uint64(m.UpdatedTime))
		i--
		dAtA[i] = 0x50
	}
	if m.CreatedTime != 0 {
		i = encodeVarintWorkspace(dAtA, i, uint64(m.CreatedTime))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.IsPublic {
		i--
		if m.IsPublic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.CreatedBy) > 0 {
		i -= len(m.CreatedBy)
		copy(dAtA[i:], m.CreatedBy)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.CreatedBy)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XId) > 0 {
		i -= len(m.XId)
		copy(dAtA[i:], m.XId)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.XId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkspace(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkspace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Workspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XId)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	l = len(m.CreatedBy)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	if m.IsPublic {
		n += 2
	}
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovWorkspace(uint64(l))
		}
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	if m.CreatedTime != 0 {
		n += 1 + sovWorkspace(uint64(m.CreatedTime))
	}
	if m.UpdatedTime != 0 {
		n += 1 + sovWorkspace(uint64(m.UpdatedTime))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkspace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkspace(x uint64) (n int) {
	return sovWorkspace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Workspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkspace
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
			return fmt.Errorf("proto: Workspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Workspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
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
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPublic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedTime", wireType)
			}
			m.CreatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedTime", wireType)
			}
			m.UpdatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
		default:
			iNdEx = preIndex
			skippy, err := skipWorkspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorkspace
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
func skipWorkspace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
				return 0, ErrInvalidLengthWorkspace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorkspace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorkspace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorkspace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkspace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorkspace = fmt.Errorf("proto: unexpected end of group")
)
