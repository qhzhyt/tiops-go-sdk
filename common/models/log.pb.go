// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tiops/common/models/log.proto

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

type LogLevel int32

const (
	LogLevel_DEBUG    LogLevel = 0
	LogLevel_INFO     LogLevel = 1
	LogLevel_WARNING  LogLevel = 2
	LogLevel_ERROR    LogLevel = 3
	LogLevel_CRITICAL LogLevel = 4
)

var LogLevel_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "CRITICAL",
}

var LogLevel_value = map[string]int32{
	"DEBUG":    0,
	"INFO":     1,
	"WARNING":  2,
	"ERROR":    3,
	"CRITICAL": 4,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}

func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ed4936e38b207de, []int{0}
}

//enum LogType {System=0; WorkFlow=6;};
type Log struct {
	XId                  int64    `protobuf:"varint,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	LogID                int64    `protobuf:"varint,2,opt,name=logID,proto3" json:"logID,omitempty" bson:"logID"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty" bson:"source"`
	Module               string   `protobuf:"bytes,5,opt,name=module,proto3" json:"module,omitempty" bson:"module"`
	Position             string   `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty" bson:"position"`
	Sn                   int32    `protobuf:"varint,7,opt,name=sn,proto3" json:"sn,omitempty" bson:"sn"`
	Level                LogLevel `protobuf:"varint,8,opt,name=level,proto3,enum=models.LogLevel" json:"level,omitempty" bson:"level"`
	Time                 int64    `protobuf:"varint,9,opt,name=time,proto3" json:"time,omitempty" bson:"time"`
	Content              string   `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty" bson:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed4936e38b207de, []int{0}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetXId() int64 {
	if m != nil {
		return m.XId
	}
	return 0
}

func (m *Log) GetLogID() int64 {
	if m != nil {
		return m.LogID
	}
	return 0
}

func (m *Log) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Log) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Log) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *Log) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Log) GetSn() int32 {
	if m != nil {
		return m.Sn
	}
	return 0
}

func (m *Log) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_DEBUG
}

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterEnum("models.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterType((*Log)(nil), "models.Log")
}

func init() { proto.RegisterFile("tiops/common/models/log.proto", fileDescriptor_7ed4936e38b207de) }

var fileDescriptor_7ed4936e38b207de = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x3b, 0xf9, 0x6b, 0x7a, 0xbe, 0x8f, 0x1a, 0x0e, 0x22, 0x43, 0xc1, 0x10, 0xba, 0x90,
	0xe0, 0xa2, 0x05, 0xdd, 0xb9, 0xeb, 0x9f, 0x25, 0x10, 0x5a, 0x18, 0x14, 0x97, 0xc5, 0x36, 0x43,
	0x08, 0x24, 0x39, 0xa1, 0x49, 0x05, 0xd7, 0x7a, 0x11, 0xe2, 0xc2, 0xeb, 0x71, 0xe9, 0x25, 0x48,
	0xbd, 0x11, 0xc9, 0xc4, 0xba, 0x72, 0x77, 0x9e, 0xe7, 0xbc, 0xcc, 0xbc, 0x33, 0x70, 0x5a, 0x25,
	0x54, 0x94, 0xc3, 0x0d, 0x65, 0x19, 0xe5, 0xc3, 0x8c, 0x22, 0x99, 0x96, 0xc3, 0x94, 0xe2, 0x41,
	0xb1, 0xa5, 0x8a, 0xd0, 0x6a, 0x4c, 0xcf, 0xfb, 0x2b, 0x56, 0xdd, 0xc7, 0xb1, 0xdc, 0x36, 0xc9,
	0xfe, 0x93, 0x06, 0x7a, 0x48, 0x31, 0x7a, 0xa0, 0xaf, 0x92, 0x88, 0x33, 0x8f, 0xf9, 0xfa, 0xf8,
	0xe8, 0xf5, 0xf9, 0x4d, 0x87, 0x75, 0x49, 0xf9, 0x55, 0x7f, 0x95, 0x44, 0x7d, 0xa1, 0x05, 0x11,
	0x1e, 0x83, 0x99, 0x52, 0x1c, 0x4c, 0xb9, 0x56, 0x67, 0x44, 0x03, 0x88, 0x60, 0x54, 0x8f, 0x85,
	0xe4, 0xba, 0xc7, 0xfc, 0x8e, 0x50, 0x33, 0x9e, 0x80, 0x55, 0xd2, 0x6e, 0xbb, 0x91, 0xdc, 0x50,
	0xf6, 0x87, 0x6a, 0x9f, 0x51, 0xb4, 0x4b, 0x25, 0x37, 0x1b, 0xdf, 0x10, 0xf6, 0xc0, 0x2e, 0xa8,
	0x4c, 0xaa, 0x84, 0x72, 0x6e, 0xa9, 0xcd, 0x2f, 0x63, 0x17, 0xb4, 0x32, 0xe7, 0x6d, 0x8f, 0xf9,
	0xa6, 0xd0, 0xca, 0x1c, 0xcf, 0xc0, 0x4c, 0xe5, 0x83, 0x4c, 0xb9, 0xed, 0x31, 0xbf, 0x7b, 0xe1,
	0x0c, 0x9a, 0x47, 0x0d, 0x42, 0x8a, 0xc3, 0xda, 0x8b, 0x66, 0xad, 0x7a, 0x25, 0x99, 0xe4, 0x1d,
	0x55, 0x56, 0xcd, 0xc8, 0xa1, 0xbd, 0xa1, 0xbc, 0x92, 0x79, 0xc5, 0x41, 0x5d, 0x73, 0xc0, 0xf3,
	0x19, 0xd8, 0x87, 0x03, 0xb0, 0x03, 0xe6, 0x74, 0x36, 0xbe, 0x9d, 0x3b, 0x2d, 0xb4, 0xc1, 0x08,
	0x16, 0xd7, 0x4b, 0x87, 0xe1, 0x3f, 0x68, 0xdf, 0x8d, 0xc4, 0x22, 0x58, 0xcc, 0x1d, 0xad, 0x4e,
	0xcc, 0x84, 0x58, 0x0a, 0x47, 0xc7, 0xff, 0x60, 0x4f, 0x44, 0x70, 0x13, 0x4c, 0x46, 0xa1, 0x63,
	0x8c, 0x9d, 0xf7, 0xbd, 0xcb, 0x3e, 0xf6, 0x2e, 0xfb, 0xdc, 0xbb, 0xec, 0xe5, 0xcb, 0x6d, 0xad,
	0x2d, 0xf5, 0xcb, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0xc9, 0x14, 0xd9, 0xb0, 0x01,
	0x00, 0x00,
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x52
	}
	if m.Time != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x48
	}
	if m.Level != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x40
	}
	if m.Sn != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Sn))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Position) > 0 {
		i -= len(m.Position)
		copy(dAtA[i:], m.Position)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Position)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintLog(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LogID != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.LogID))
		i--
		dAtA[i] = 0x10
	}
	if m.XId != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.XId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	offset -= sovLog(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XId != 0 {
		n += 1 + sovLog(uint64(m.XId))
	}
	if m.LogID != 0 {
		n += 1 + sovLog(uint64(m.LogID))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Sn != 0 {
		n += 1 + sovLog(uint64(m.Sn))
	}
	if m.Level != 0 {
		n += 1 + sovLog(uint64(m.Level))
	}
	if m.Time != 0 {
		n += 1 + sovLog(uint64(m.Time))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLog(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
			}
			m.XId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.XId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogID", wireType)
			}
			m.LogID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sn", wireType)
			}
			m.Sn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sn |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= LogLevel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLog
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
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
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
				return 0, ErrInvalidLengthLog
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLog
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLog
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLog        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLog = fmt.Errorf("proto: unexpected end of group")
)
