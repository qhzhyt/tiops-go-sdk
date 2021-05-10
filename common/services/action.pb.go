// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: tiops/common/services/action.proto

//import "tiops/common/models/workflow.proto";

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DataType int32

const (
	DataType_List     DataType = 0
	DataType_Constant DataType = 1
	DataType_File     DataType = 2
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "List",
		1: "Constant",
		2: "File",
	}
	DataType_value = map[string]int32{
		"List":     0,
		"Constant": 1,
		"File":     2,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_tiops_common_services_action_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_tiops_common_services_action_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{0}
}

type ResponseCode int32

const (
	ResponseCode_UNKNOWN ResponseCode = 0
	ResponseCode_OK      ResponseCode = 1
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
	}
	ResponseCode_value = map[string]int32{
		"UNKNOWN": 0,
		"OK":      1,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_tiops_common_services_action_proto_enumTypes[1].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_tiops_common_services_action_proto_enumTypes[1]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{1}
}

type ActionArgument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" bson:"value"`
}

func (x *ActionArgument) Reset() {
	*x = ActionArgument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionArgument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionArgument) ProtoMessage() {}

func (x *ActionArgument) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionArgument.ProtoReflect.Descriptor instead.
func (*ActionArgument) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{0}
}

func (x *ActionArgument) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ActionArgument) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ActionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	Type      DataType `protobuf:"varint,2,opt,name=type,proto3,enum=services.DataType" json:"type,omitempty" bson:"type"`
	ValueType string   `protobuf:"bytes,3,opt,name=valueType,proto3" json:"valueType,omitempty" bson:"valueType"`
	Data      [][]byte `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" bson:"data"`
	Count     int32    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty" bson:"count"`
	Timestamp int32    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty" bson:"timestamp"`
}

func (x *ActionData) Reset() {
	*x = ActionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionData) ProtoMessage() {}

func (x *ActionData) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionData.ProtoReflect.Descriptor instead.
func (*ActionData) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{1}
}

func (x *ActionData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionData) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_List
}

func (x *ActionData) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *ActionData) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ActionData) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ActionData) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	NodeId     string `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty" bson:"nodeId"`
	ActionName string `protobuf:"bytes,3,opt,name=actionName,proto3" json:"actionName,omitempty" bson:"actionName"`
	//  ActionData inputs = 2;
	Inputs map[string]*ActionData `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"inputs"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{2}
}

func (x *ActionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ActionRequest) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *ActionRequest) GetInputs() map[string]*ActionData {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	Code    ResponseCode `protobuf:"varint,2,opt,name=code,proto3,enum=services.ResponseCode" json:"code,omitempty" bson:"code"`
	Message string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty" bson:"message"`
	//  ActionData result = 4;
	Outputs map[string]*ActionData `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"outputs"`
	Count   int32                  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty" bson:"count"`
	Done    bool                   `protobuf:"varint,6,opt,name=done,proto3" json:"done,omitempty" bson:"done"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{3}
}

func (x *ActionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionResponse) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_UNKNOWN
}

func (x *ActionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ActionResponse) GetOutputs() map[string]*ActionData {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *ActionResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ActionResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type RegisterActionNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionName    string            `protobuf:"bytes,1,opt,name=actionName,proto3" json:"actionName,omitempty" bson:"actionName"`
	NodeId        string            `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty" bson:"nodeId"`
	ActionOptions map[string]string `protobuf:"bytes,3,rep,name=actionOptions,proto3" json:"actionOptions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"actionOptions"`
}

func (x *RegisterActionNodeRequest) Reset() {
	*x = RegisterActionNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterActionNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterActionNodeRequest) ProtoMessage() {}

func (x *RegisterActionNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterActionNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterActionNodeRequest) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_action_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterActionNodeRequest) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *RegisterActionNodeRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RegisterActionNodeRequest) GetActionOptions() map[string]string {
	if x != nil {
		return x.ActionOptions
	}
	return nil
}

var File_tiops_common_services_action_proto protoreflect.FileDescriptor

var file_tiops_common_services_action_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x69, 0x6f, 0x70, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x1a,
	0x4f, 0x0a, 0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa3, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x1a, 0x50, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf3, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x0d,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2c, 0x0a, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x2a, 0x23, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x32,
	0xaa, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tiops_common_services_action_proto_rawDescOnce sync.Once
	file_tiops_common_services_action_proto_rawDescData = file_tiops_common_services_action_proto_rawDesc
)

func file_tiops_common_services_action_proto_rawDescGZIP() []byte {
	file_tiops_common_services_action_proto_rawDescOnce.Do(func() {
		file_tiops_common_services_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiops_common_services_action_proto_rawDescData)
	})
	return file_tiops_common_services_action_proto_rawDescData
}

var file_tiops_common_services_action_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tiops_common_services_action_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tiops_common_services_action_proto_goTypes = []interface{}{
	(DataType)(0),                     // 0: services.DataType
	(ResponseCode)(0),                 // 1: services.ResponseCode
	(*ActionArgument)(nil),            // 2: services.ActionArgument
	(*ActionData)(nil),                // 3: services.ActionData
	(*ActionRequest)(nil),             // 4: services.ActionRequest
	(*ActionResponse)(nil),            // 5: services.ActionResponse
	(*RegisterActionNodeRequest)(nil), // 6: services.RegisterActionNodeRequest
	nil,                               // 7: services.ActionRequest.InputsEntry
	nil,                               // 8: services.ActionResponse.OutputsEntry
	nil,                               // 9: services.RegisterActionNodeRequest.ActionOptionsEntry
	(*StatusResponse)(nil),            // 10: services.StatusResponse
}
var file_tiops_common_services_action_proto_depIdxs = []int32{
	0,  // 0: services.ActionData.type:type_name -> services.DataType
	7,  // 1: services.ActionRequest.inputs:type_name -> services.ActionRequest.InputsEntry
	1,  // 2: services.ActionResponse.code:type_name -> services.ResponseCode
	8,  // 3: services.ActionResponse.outputs:type_name -> services.ActionResponse.OutputsEntry
	9,  // 4: services.RegisterActionNodeRequest.actionOptions:type_name -> services.RegisterActionNodeRequest.ActionOptionsEntry
	3,  // 5: services.ActionRequest.InputsEntry.value:type_name -> services.ActionData
	3,  // 6: services.ActionResponse.OutputsEntry.value:type_name -> services.ActionData
	4,  // 7: services.ActionsService.CallAction:input_type -> services.ActionRequest
	6,  // 8: services.ActionsService.RegisterActionNode:input_type -> services.RegisterActionNodeRequest
	5,  // 9: services.ActionsService.CallAction:output_type -> services.ActionResponse
	10, // 10: services.ActionsService.RegisterActionNode:output_type -> services.StatusResponse
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_tiops_common_services_action_proto_init() }
func file_tiops_common_services_action_proto_init() {
	if File_tiops_common_services_action_proto != nil {
		return
	}
	file_tiops_common_services_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tiops_common_services_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionArgument); i {
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
		file_tiops_common_services_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionData); i {
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
		file_tiops_common_services_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_tiops_common_services_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
		file_tiops_common_services_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterActionNodeRequest); i {
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
			RawDescriptor: file_tiops_common_services_action_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tiops_common_services_action_proto_goTypes,
		DependencyIndexes: file_tiops_common_services_action_proto_depIdxs,
		EnumInfos:         file_tiops_common_services_action_proto_enumTypes,
		MessageInfos:      file_tiops_common_services_action_proto_msgTypes,
	}.Build()
	File_tiops_common_services_action_proto = out.File
	file_tiops_common_services_action_proto_rawDesc = nil
	file_tiops_common_services_action_proto_goTypes = nil
	file_tiops_common_services_action_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActionsServiceClient is the client API for ActionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActionsServiceClient interface {
	// define the interface and data type
	CallAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	RegisterActionNode(ctx context.Context, in *RegisterActionNodeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type actionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsServiceClient(cc grpc.ClientConnInterface) ActionsServiceClient {
	return &actionsServiceClient{cc}
}

func (c *actionsServiceClient) CallAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/services.ActionsService/CallAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) RegisterActionNode(ctx context.Context, in *RegisterActionNodeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/services.ActionsService/RegisterActionNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServiceServer is the server API for ActionsService service.
type ActionsServiceServer interface {
	// define the interface and data type
	CallAction(context.Context, *ActionRequest) (*ActionResponse, error)
	RegisterActionNode(context.Context, *RegisterActionNodeRequest) (*StatusResponse, error)
}

// UnimplementedActionsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActionsServiceServer struct {
}

func (*UnimplementedActionsServiceServer) CallAction(context.Context, *ActionRequest) (*ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallAction not implemented")
}
func (*UnimplementedActionsServiceServer) RegisterActionNode(context.Context, *RegisterActionNodeRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterActionNode not implemented")
}

func RegisterActionsServiceServer(s *grpc.Server, srv ActionsServiceServer) {
	s.RegisterService(&_ActionsService_serviceDesc, srv)
}

func _ActionsService_CallAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).CallAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ActionsService/CallAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).CallAction(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_RegisterActionNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterActionNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).RegisterActionNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ActionsService/RegisterActionNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).RegisterActionNode(ctx, req.(*RegisterActionNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ActionsService",
	HandlerType: (*ActionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallAction",
			Handler:    _ActionsService_CallAction_Handler,
		},
		{
			MethodName: "RegisterActionNode",
			Handler:    _ActionsService_RegisterActionNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tiops/common/services/action.proto",
}
