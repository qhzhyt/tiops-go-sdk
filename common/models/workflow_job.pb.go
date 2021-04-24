// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tiops/common/models/workflow_job.proto

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

type WorkflowJob struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	WorkflowId           string   `protobuf:"bytes,2,opt,name=workflowId,proto3" json:"workflowId,omitempty" bson:"workflowId"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Namespace            string   `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty" bson:"namespace"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	LogId                string   `protobuf:"bytes,6,opt,name=logId,proto3" json:"logId,omitempty" bson:"logId"`
	CreatedBy            string   `protobuf:"bytes,7,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	CreatedTime          int64    `protobuf:"varint,8,opt,name=createdTime,proto3" json:"createdTime,omitempty" bson:"createdTime"`
	StartedTime          int64    `protobuf:"varint,9,opt,name=startedTime,proto3" json:"startedTime,omitempty" bson:"startedTime"`
	FinishedTime         int64    `protobuf:"varint,10,opt,name=finishedTime,proto3" json:"finishedTime,omitempty" bson:"finishedTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *WorkflowJob) Reset()         { *m = WorkflowJob{} }
func (m *WorkflowJob) String() string { return proto.CompactTextString(m) }
func (*WorkflowJob) ProtoMessage()    {}
func (*WorkflowJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_04c94e5c97aaaf4f, []int{0}
}
func (m *WorkflowJob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowJob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowJob.Merge(m, src)
}
func (m *WorkflowJob) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowJob) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowJob.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowJob proto.InternalMessageInfo

func (m *WorkflowJob) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *WorkflowJob) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowJob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowJob) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WorkflowJob) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *WorkflowJob) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func (m *WorkflowJob) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *WorkflowJob) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *WorkflowJob) GetStartedTime() int64 {
	if m != nil {
		return m.StartedTime
	}
	return 0
}

func (m *WorkflowJob) GetFinishedTime() int64 {
	if m != nil {
		return m.FinishedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*WorkflowJob)(nil), "models.WorkflowJob")
}

func init() {
	proto.RegisterFile("tiops/common/models/workflow_job.proto", fileDescriptor_04c94e5c97aaaf4f)
}

var fileDescriptor_04c94e5c97aaaf4f = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x71, 0xd3, 0x16, 0x7a, 0x45, 0x02, 0x59, 0x0c, 0x16, 0x42, 0x51, 0x94, 0x01, 0x75,
	0xa2, 0x03, 0x1b, 0x63, 0xb7, 0x30, 0x46, 0x48, 0x8c, 0x51, 0x12, 0xbb, 0xc1, 0x90, 0xe4, 0xa2,
	0xd8, 0x52, 0xd5, 0x9d, 0x87, 0x40, 0x0c, 0xbc, 0x03, 0x6f, 0xc1, 0xc8, 0x23, 0xa0, 0xf0, 0x22,
	0xc8, 0x76, 0x0a, 0x41, 0xea, 0xe4, 0xbb, 0xef, 0xff, 0xee, 0x6c, 0x19, 0x2e, 0xb5, 0xc4, 0x46,
	0x2d, 0x73, 0xac, 0x2a, 0xac, 0x97, 0x15, 0x72, 0x51, 0xaa, 0xe5, 0x06, 0xdb, 0xa7, 0x75, 0x89,
	0x9b, 0xe4, 0x11, 0xb3, 0xab, 0xa6, 0x45, 0x8d, 0x74, 0xea, 0xa2, 0xf3, 0x60, 0x9f, 0xef, 0x3a,
	0x67, 0xee, 0x37, 0x74, 0x5a, 0x14, 0xa2, 0x75, 0x46, 0xf8, 0x3e, 0x82, 0xf9, 0x7d, 0x7f, 0xc5,
	0x2d, 0x66, 0x34, 0x00, 0x2f, 0x91, 0x9c, 0x91, 0x80, 0x2c, 0x66, 0xab, 0x93, 0xd7, 0xe7, 0x37,
	0x0f, 0x32, 0x85, 0xf5, 0x4d, 0x98, 0x48, 0x1e, 0xc6, 0xa3, 0x88, 0x53, 0x1f, 0x60, 0xf7, 0xa6,
	0x88, 0xb3, 0x91, 0x11, 0xe3, 0x01, 0xa1, 0x14, 0xc6, 0x75, 0x5a, 0x09, 0xe6, 0xd9, 0xc4, 0xd6,
	0xf4, 0x02, 0x66, 0xe6, 0x54, 0x4d, 0x9a, 0x0b, 0x36, 0xb6, 0xc1, 0x1f, 0x30, 0x13, 0x7a, 0xdb,
	0x08, 0x36, 0x71, 0x13, 0xa6, 0xa6, 0x67, 0x30, 0x29, 0xb1, 0x88, 0x38, 0x9b, 0x5a, 0xe8, 0x1a,
	0xb3, 0x27, 0x6f, 0x45, 0xaa, 0x05, 0x5f, 0x6d, 0xd9, 0xa1, 0xdb, 0xf3, 0x0b, 0x68, 0x00, 0xf3,
	0xbe, 0xb9, 0x93, 0x95, 0x60, 0x47, 0x01, 0x59, 0x78, 0xf1, 0x10, 0x19, 0x43, 0xe9, 0xb4, 0xdd,
	0x19, 0x33, 0x67, 0x0c, 0x10, 0x0d, 0xe1, 0x78, 0x2d, 0x6b, 0xa9, 0x1e, 0x7a, 0x05, 0xac, 0xf2,
	0x8f, 0xad, 0x4e, 0x3f, 0x3a, 0x9f, 0x7c, 0x76, 0x3e, 0xf9, 0xea, 0x7c, 0xf2, 0xf2, 0xed, 0x1f,
	0x64, 0x53, 0xfb, 0x99, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xcb, 0x05, 0xc5, 0xc2,
	0x01, 0x00, 0x00,
}

func (m *WorkflowJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowJob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowJob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.FinishedTime != 0 {
		i = encodeVarintWorkflowJob(dAtA, i, uint64(m.FinishedTime))
		i--
		dAtA[i] = 0x50
	}
	if m.StartedTime != 0 {
		i = encodeVarintWorkflowJob(dAtA, i, uint64(m.StartedTime))
		i--
		dAtA[i] = 0x48
	}
	if m.CreatedTime != 0 {
		i = encodeVarintWorkflowJob(dAtA, i, uint64(m.CreatedTime))
		i--
		dAtA[i] = 0x40
	}
	if len(m.CreatedBy) > 0 {
		i -= len(m.CreatedBy)
		copy(dAtA[i:], m.CreatedBy)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.CreatedBy)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LogId) > 0 {
		i -= len(m.LogId)
		copy(dAtA[i:], m.LogId)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.LogId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.WorkflowId) > 0 {
		i -= len(m.WorkflowId)
		copy(dAtA[i:], m.WorkflowId)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.WorkflowId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XId) > 0 {
		i -= len(m.XId)
		copy(dAtA[i:], m.XId)
		i = encodeVarintWorkflowJob(dAtA, i, uint64(len(m.XId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkflowJob(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkflowJob(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkflowJob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XId)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	l = len(m.WorkflowId)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	l = len(m.LogId)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	l = len(m.CreatedBy)
	if l > 0 {
		n += 1 + l + sovWorkflowJob(uint64(l))
	}
	if m.CreatedTime != 0 {
		n += 1 + sovWorkflowJob(uint64(m.CreatedTime))
	}
	if m.StartedTime != 0 {
		n += 1 + sovWorkflowJob(uint64(m.StartedTime))
	}
	if m.FinishedTime != 0 {
		n += 1 + sovWorkflowJob(uint64(m.FinishedTime))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkflowJob(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkflowJob(x uint64) (n int) {
	return sovWorkflowJob(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkflowJob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowJob
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
			return fmt.Errorf("proto: WorkflowJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowJob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkflowId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
				return ErrInvalidLengthWorkflowJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedTime", wireType)
			}
			m.CreatedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedTime", wireType)
			}
			m.StartedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishedTime", wireType)
			}
			m.FinishedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinishedTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowJob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorkflowJob
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
func skipWorkflowJob(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkflowJob
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
					return 0, ErrIntOverflowWorkflowJob
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
					return 0, ErrIntOverflowWorkflowJob
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
				return 0, ErrInvalidLengthWorkflowJob
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorkflowJob
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorkflowJob
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorkflowJob        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkflowJob          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorkflowJob = fmt.Errorf("proto: unexpected end of group")
)
