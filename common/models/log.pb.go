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
type LogExtra struct {
	NodeId               string                  `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty" bson:"nodeId"`
	Count                int32                   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty" bson:"count"`
	Sum                  int64                   `protobuf:"varint,3,opt,name=sum,proto3" json:"sum,omitempty" bson:"sum"`
	Status               WorkflowExecutionStatus `protobuf:"varint,4,opt,name=status,proto3,enum=models.WorkflowExecutionStatus" json:"status,omitempty" bson:"status"`
	Others               map[string]string       `protobuf:"bytes,16,rep,name=others,proto3" json:"others,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"others"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-" bson:"-"`
	XXX_unrecognized     []byte                  `json:"-" bson:"-"`
	XXX_sizecache        int32                   `json:"-" bson:"-"`
}

func (m *LogExtra) Reset()         { *m = LogExtra{} }
func (m *LogExtra) String() string { return proto.CompactTextString(m) }
func (*LogExtra) ProtoMessage()    {}
func (*LogExtra) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed4936e38b207de, []int{0}
}
func (m *LogExtra) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogExtra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogExtra.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogExtra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogExtra.Merge(m, src)
}
func (m *LogExtra) XXX_Size() int {
	return m.Size()
}
func (m *LogExtra) XXX_DiscardUnknown() {
	xxx_messageInfo_LogExtra.DiscardUnknown(m)
}

var xxx_messageInfo_LogExtra proto.InternalMessageInfo

func (m *LogExtra) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *LogExtra) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LogExtra) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *LogExtra) GetStatus() WorkflowExecutionStatus {
	if m != nil {
		return m.Status
	}
	return WorkflowExecutionStatus_Created
}

func (m *LogExtra) GetOthers() map[string]string {
	if m != nil {
		return m.Others
	}
	return nil
}

type Log struct {
	XId                  string    `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	LogID                string    `protobuf:"bytes,2,opt,name=logID,proto3" json:"logID,omitempty" bson:"logID"`
	Type                 string    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	Source               string    `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty" bson:"source"`
	Module               string    `protobuf:"bytes,5,opt,name=module,proto3" json:"module,omitempty" bson:"module"`
	Position             string    `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty" bson:"position"`
	Sn                   int32     `protobuf:"varint,7,opt,name=sn,proto3" json:"sn,omitempty" bson:"sn"`
	Level                LogLevel  `protobuf:"varint,8,opt,name=level,proto3,enum=models.LogLevel" json:"level,omitempty" bson:"level"`
	Time                 int64     `protobuf:"varint,9,opt,name=time,proto3" json:"time,omitempty" bson:"time"`
	Content              string    `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty" bson:"content"`
	Extra                *LogExtra `protobuf:"bytes,11,opt,name=extra,proto3" json:"extra,omitempty" bson:"extra"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed4936e38b207de, []int{1}
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

func (m *Log) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Log) GetLogID() string {
	if m != nil {
		return m.LogID
	}
	return ""
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

func (m *Log) GetExtra() *LogExtra {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterEnum("models.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterType((*LogExtra)(nil), "models.LogExtra")
	proto.RegisterMapType((map[string]string)(nil), "models.LogExtra.OthersEntry")
	proto.RegisterType((*Log)(nil), "models.Log")
}

func init() { proto.RegisterFile("tiops/common/models/log.proto", fileDescriptor_7ed4936e38b207de) }

var fileDescriptor_7ed4936e38b207de = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xdd, 0x49, 0xda, 0xb4, 0xb9, 0x95, 0x35, 0x0c, 0xb2, 0x0c, 0x45, 0x6b, 0xe8, 0x83, 0x04,
	0x91, 0x14, 0x56, 0x41, 0xdd, 0xb7, 0xed, 0x6e, 0x5c, 0x02, 0xa5, 0x85, 0x51, 0xd9, 0xc7, 0xd2,
	0x6d, 0xc7, 0x18, 0x36, 0xc9, 0x2d, 0xc9, 0x64, 0x77, 0xfb, 0xee, 0x47, 0x88, 0x0f, 0xfe, 0x80,
	0x3f, 0xe2, 0xa3, 0x9f, 0x20, 0xf5, 0x17, 0xfc, 0x00, 0x99, 0x99, 0x54, 0x76, 0x61, 0xdf, 0xee,
	0xb9, 0xf7, 0x64, 0xce, 0xb9, 0x27, 0x17, 0x9e, 0xc8, 0x14, 0xd7, 0xd5, 0x68, 0x89, 0x79, 0x8e,
	0xc5, 0x28, 0xc7, 0x95, 0xc8, 0xaa, 0x51, 0x86, 0x49, 0xb8, 0x2e, 0x51, 0x22, 0x75, 0x4c, 0xa7,
	0xef, 0xdf, 0x47, 0x93, 0x8b, 0x24, 0x11, 0xa5, 0x61, 0xf6, 0x5f, 0xdc, 0xc7, 0xb8, 0xc6, 0xf2,
	0xf2, 0x53, 0x86, 0xd7, 0x73, 0x71, 0x23, 0x96, 0xb5, 0x4c, 0xb1, 0x30, 0xec, 0xe1, 0x5f, 0x02,
	0xdd, 0x09, 0x26, 0xd1, 0x8d, 0x2c, 0x17, 0xf4, 0x00, 0x9c, 0x02, 0x57, 0x22, 0x5e, 0x31, 0xe2,
	0x93, 0xc0, 0xe5, 0x0d, 0xa2, 0x8f, 0xa0, 0xbd, 0xc4, 0xba, 0x90, 0xcc, 0xf2, 0x49, 0xd0, 0xe6,
	0x06, 0x50, 0x0f, 0xec, 0xaa, 0xce, 0x99, 0xed, 0x93, 0xc0, 0xe6, 0xaa, 0xa4, 0xaf, 0xc1, 0xa9,
	0xe4, 0x42, 0xd6, 0x15, 0x6b, 0xf9, 0x24, 0xd8, 0x3f, 0x7c, 0x1a, 0x1a, 0xf9, 0xf0, 0xbc, 0x91,
	0x8f, 0x76, 0xea, 0xef, 0x35, 0x8d, 0x37, 0x74, 0xfa, 0x0a, 0x1c, 0x94, 0x9f, 0x45, 0x59, 0x31,
	0xcf, 0xb7, 0x83, 0xde, 0xe1, 0xe3, 0xdd, 0x87, 0x3b, 0x6b, 0xe1, 0x4c, 0x8f, 0xa3, 0x42, 0x96,
	0x1b, 0xde, 0x70, 0xfb, 0x6f, 0xa1, 0x77, 0xab, 0xad, 0xfc, 0x5c, 0x8a, 0x4d, 0x63, 0x5d, 0x95,
	0xca, 0xf7, 0xd5, 0x22, 0xab, 0x85, 0xf6, 0xed, 0x72, 0x03, 0x8e, 0xac, 0x37, 0x64, 0xf8, 0xc3,
	0x02, 0x7b, 0x82, 0x09, 0xf5, 0xc1, 0x9e, 0xa7, 0xcd, 0xba, 0xe3, 0x87, 0xdf, 0xbe, 0x7c, 0xb7,
	0xe1, 0xa2, 0xc2, 0xe2, 0x68, 0x38, 0x4f, 0x57, 0x43, 0x6e, 0x99, 0xdd, 0x33, 0x4c, 0xe2, 0xd3,
	0xdd, 0x1b, 0x1a, 0x50, 0x0a, 0x2d, 0xb9, 0x59, 0x0b, 0xbd, 0xbc, 0xcb, 0x75, 0xad, 0xd2, 0xab,
	0xb0, 0x2e, 0x97, 0x42, 0x6f, 0xef, 0xf2, 0x06, 0xa9, 0x7e, 0x8e, 0xab, 0x3a, 0x13, 0xac, 0x6d,
	0xfa, 0x06, 0xd1, 0x3e, 0x74, 0xd7, 0x58, 0xa5, 0x2a, 0x0e, 0xe6, 0xe8, 0xc9, 0x7f, 0x4c, 0xf7,
	0xc1, 0xaa, 0x0a, 0xd6, 0xd1, 0x71, 0x5b, 0x55, 0x41, 0x9f, 0x41, 0x3b, 0x13, 0x57, 0x22, 0x63,
	0x5d, 0x1d, 0xac, 0x77, 0x2b, 0x9f, 0x89, 0xea, 0x73, 0x33, 0xd6, 0xbe, 0xd2, 0x5c, 0x30, 0x57,
	0xff, 0x14, 0x5d, 0x53, 0x06, 0x9d, 0x25, 0x16, 0x52, 0x14, 0x92, 0x81, 0x96, 0xd9, 0x41, 0xf5,
	0xaa, 0x50, 0xe9, 0xb2, 0x9e, 0x4f, 0x82, 0xde, 0x9d, 0x57, 0x75, 0xea, 0xdc, 0x8c, 0x9f, 0x47,
	0xfa, 0x46, 0xb4, 0x10, 0x75, 0xa1, 0x7d, 0x1a, 0x8d, 0x3f, 0x9e, 0x79, 0x7b, 0xb4, 0x0b, 0xad,
	0x78, 0xfa, 0x6e, 0xe6, 0x11, 0xda, 0x83, 0xce, 0xf9, 0x31, 0x9f, 0xc6, 0xd3, 0x33, 0xcf, 0x52,
	0x8c, 0x88, 0xf3, 0x19, 0xf7, 0x6c, 0xfa, 0x00, 0xba, 0x27, 0x3c, 0xfe, 0x10, 0x9f, 0x1c, 0x4f,
	0xbc, 0xd6, 0x38, 0xf8, 0xb9, 0x1d, 0x90, 0x5f, 0xdb, 0x01, 0xf9, 0xbd, 0x1d, 0x90, 0xaf, 0x7f,
	0x06, 0x7b, 0x70, 0x90, 0x62, 0xa8, 0xcf, 0x35, 0x34, 0xe7, 0xda, 0x18, 0xb8, 0x70, 0xf4, 0x71,
	0xbe, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xb9, 0xfd, 0xe9, 0x15, 0x03, 0x00, 0x00,
}

func (m *LogExtra) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogExtra) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogExtra) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Others) > 0 {
		for k := range m.Others {
			v := m.Others[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintLog(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintLog(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintLog(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if m.Status != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Sum != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Sum))
		i--
		dAtA[i] = 0x18
	}
	if m.Count != 0 {
		i = encodeVarintLog(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintLog(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if m.Extra != nil {
		{
			size, err := m.Extra.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLog(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
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
	if len(m.LogID) > 0 {
		i -= len(m.LogID)
		copy(dAtA[i:], m.LogID)
		i = encodeVarintLog(dAtA, i, uint64(len(m.LogID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XId) > 0 {
		i -= len(m.XId)
		copy(dAtA[i:], m.XId)
		i = encodeVarintLog(dAtA, i, uint64(len(m.XId)))
		i--
		dAtA[i] = 0xa
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
func (m *LogExtra) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovLog(uint64(m.Count))
	}
	if m.Sum != 0 {
		n += 1 + sovLog(uint64(m.Sum))
	}
	if m.Status != 0 {
		n += 1 + sovLog(uint64(m.Status))
	}
	if len(m.Others) > 0 {
		for k, v := range m.Others {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLog(uint64(len(k))) + 1 + len(v) + sovLog(uint64(len(v)))
			n += mapEntrySize + 2 + sovLog(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XId)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.LogID)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
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
	if m.Extra != nil {
		l = m.Extra.Size()
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
func (m *LogExtra) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LogExtra: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogExtra: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
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
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			m.Sum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sum |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= WorkflowExecutionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Others", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Others == nil {
				m.Others = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLog
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
						return ErrInvalidLengthLog
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthLog
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
							return ErrIntOverflowLog
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
						return ErrInvalidLengthLog
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthLog
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLog(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthLog
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Others[mapkey] = mapvalue
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XId", wireType)
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
			m.XId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogID", wireType)
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
			m.LogID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Extra == nil {
				m.Extra = &LogExtra{}
			}
			if err := m.Extra.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
