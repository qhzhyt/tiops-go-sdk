// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: tiops/common/services/api.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	models "tiops/common/models"
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

type ObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" bson:"code"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" bson:"message"`
	Data    *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty" bson:"data"`
}

func (x *ObjectResponse) Reset() {
	*x = ObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectResponse) ProtoMessage() {}

func (x *ObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectResponse.ProtoReflect.Descriptor instead.
func (*ObjectResponse) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_api_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ObjectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ObjectResponse) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" bson:"code"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" bson:"message"`
	Data    []*any.Any `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" bson:"data"` //  oneof msg { Child1 c1 = 2; Child2 c2 = 3; }
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListResponse) GetData() []*any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_api_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LogIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty" bson:"moduleName"`
}

func (x *LogIdRequest) Reset() {
	*x = LogIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogIdRequest) ProtoMessage() {}

func (x *LogIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogIdRequest.ProtoReflect.Descriptor instead.
func (*LogIdRequest) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_api_proto_rawDescGZIP(), []int{3}
}

func (x *LogIdRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

type LogIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogId int64 `protobuf:"varint,1,opt,name=logId,proto3" json:"logId,omitempty" bson:"logId"`
}

func (x *LogIdResponse) Reset() {
	*x = LogIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_services_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogIdResponse) ProtoMessage() {}

func (x *LogIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_services_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogIdResponse.ProtoReflect.Descriptor instead.
func (*LogIdResponse) Descriptor() ([]byte, []int) {
	return file_tiops_common_services_api_proto_rawDescGZIP(), []int{4}
}

func (x *LogIdResponse) GetLogId() int64 {
	if x != nil {
		return x.LogId
	}
	return 0
}

var File_tiops_common_services_api_proto protoreflect.FileDescriptor

var file_tiops_common_services_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1d, 0x74, 0x69, 0x6f,
	0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x74, 0x69, 0x6f, 0x70,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x69, 0x6f, 0x70, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x74,
	0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x66,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x32, 0xbd, 0x04,
	0x0a, 0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07,
	0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x0a,
	0x18, 0x69, 0x6f, 0x2e, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tiops_common_services_api_proto_rawDescOnce sync.Once
	file_tiops_common_services_api_proto_rawDescData = file_tiops_common_services_api_proto_rawDesc
)

func file_tiops_common_services_api_proto_rawDescGZIP() []byte {
	file_tiops_common_services_api_proto_rawDescOnce.Do(func() {
		file_tiops_common_services_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiops_common_services_api_proto_rawDescData)
	})
	return file_tiops_common_services_api_proto_rawDescData
}

var file_tiops_common_services_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tiops_common_services_api_proto_goTypes = []interface{}{
	(*ObjectResponse)(nil),           // 0: services.ObjectResponse
	(*ListResponse)(nil),             // 1: services.ListResponse
	(*QueryRequest)(nil),             // 2: services.QueryRequest
	(*LogIdRequest)(nil),             // 3: services.LogIdRequest
	(*LogIdResponse)(nil),            // 4: services.LogIdResponse
	(*any.Any)(nil),                  // 5: google.protobuf.Any
	(*models.Log)(nil),               // 6: models.Log
	(*models.WorkflowExecution)(nil), // 7: models.WorkflowExecution
	(*models.ExecutionRecord)(nil),   // 8: models.ExecutionRecord
	(*models.ProcessRecord)(nil),     // 9: models.ProcessRecord
	(*models.ProjectInfo)(nil),       // 10: models.ProjectInfo
	(*models.WorkflowInfo)(nil),      // 11: models.WorkflowInfo
	(*models.WorkflowJob)(nil),       // 12: models.WorkflowJob
	(*StatusResponse)(nil),           // 13: services.StatusResponse
}
var file_tiops_common_services_api_proto_depIdxs = []int32{
	5,  // 0: services.ObjectResponse.data:type_name -> google.protobuf.Any
	5,  // 1: services.ListResponse.data:type_name -> google.protobuf.Any
	6,  // 2: services.APIService.PostLog:input_type -> models.Log
	2,  // 3: services.APIService.GetProjectByID:input_type -> services.QueryRequest
	2,  // 4: services.APIService.GetWorkflowById:input_type -> services.QueryRequest
	3,  // 5: services.APIService.GetSystemLogId:input_type -> services.LogIdRequest
	2,  // 6: services.APIService.GetWorkflowJobById:input_type -> services.QueryRequest
	7,  // 7: services.APIService.CreateOrUpdateWorkflowExecution:input_type -> models.WorkflowExecution
	8,  // 8: services.APIService.CreateExecutionRecord:input_type -> models.ExecutionRecord
	9,  // 9: services.APIService.AddProcessRecord:input_type -> models.ProcessRecord
	0,  // 10: services.APIService.PostLog:output_type -> services.ObjectResponse
	10, // 11: services.APIService.GetProjectByID:output_type -> models.ProjectInfo
	11, // 12: services.APIService.GetWorkflowById:output_type -> models.WorkflowInfo
	4,  // 13: services.APIService.GetSystemLogId:output_type -> services.LogIdResponse
	12, // 14: services.APIService.GetWorkflowJobById:output_type -> models.WorkflowJob
	13, // 15: services.APIService.CreateOrUpdateWorkflowExecution:output_type -> services.StatusResponse
	13, // 16: services.APIService.CreateExecutionRecord:output_type -> services.StatusResponse
	13, // 17: services.APIService.AddProcessRecord:output_type -> services.StatusResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_tiops_common_services_api_proto_init() }
func file_tiops_common_services_api_proto_init() {
	if File_tiops_common_services_api_proto != nil {
		return
	}
	file_tiops_common_services_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tiops_common_services_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectResponse); i {
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
		file_tiops_common_services_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_tiops_common_services_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_tiops_common_services_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogIdRequest); i {
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
		file_tiops_common_services_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogIdResponse); i {
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
			RawDescriptor: file_tiops_common_services_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tiops_common_services_api_proto_goTypes,
		DependencyIndexes: file_tiops_common_services_api_proto_depIdxs,
		MessageInfos:      file_tiops_common_services_api_proto_msgTypes,
	}.Build()
	File_tiops_common_services_api_proto = out.File
	file_tiops_common_services_api_proto_rawDesc = nil
	file_tiops_common_services_api_proto_goTypes = nil
	file_tiops_common_services_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIServiceClient interface {
	// define the interface and data type
	PostLog(ctx context.Context, in *models.Log, opts ...grpc.CallOption) (*ObjectResponse, error)
	GetProjectByID(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*models.ProjectInfo, error)
	//  rpc GetProjects(QueryRequest) returns(ListResponse){}
	GetWorkflowById(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*models.WorkflowInfo, error)
	GetSystemLogId(ctx context.Context, in *LogIdRequest, opts ...grpc.CallOption) (*LogIdResponse, error)
	GetWorkflowJobById(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*models.WorkflowJob, error)
	CreateOrUpdateWorkflowExecution(ctx context.Context, in *models.WorkflowExecution, opts ...grpc.CallOption) (*StatusResponse, error)
	CreateExecutionRecord(ctx context.Context, in *models.ExecutionRecord, opts ...grpc.CallOption) (*StatusResponse, error)
	AddProcessRecord(ctx context.Context, in *models.ProcessRecord, opts ...grpc.CallOption) (*StatusResponse, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) PostLog(ctx context.Context, in *models.Log, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, "/services.APIService/PostLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetProjectByID(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*models.ProjectInfo, error) {
	out := new(models.ProjectInfo)
	err := c.cc.Invoke(ctx, "/services.APIService/GetProjectByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetWorkflowById(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*models.WorkflowInfo, error) {
	out := new(models.WorkflowInfo)
	err := c.cc.Invoke(ctx, "/services.APIService/GetWorkflowById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetSystemLogId(ctx context.Context, in *LogIdRequest, opts ...grpc.CallOption) (*LogIdResponse, error) {
	out := new(LogIdResponse)
	err := c.cc.Invoke(ctx, "/services.APIService/GetSystemLogId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetWorkflowJobById(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*models.WorkflowJob, error) {
	out := new(models.WorkflowJob)
	err := c.cc.Invoke(ctx, "/services.APIService/GetWorkflowJobById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) CreateOrUpdateWorkflowExecution(ctx context.Context, in *models.WorkflowExecution, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/services.APIService/CreateOrUpdateWorkflowExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) CreateExecutionRecord(ctx context.Context, in *models.ExecutionRecord, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/services.APIService/CreateExecutionRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) AddProcessRecord(ctx context.Context, in *models.ProcessRecord, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/services.APIService/AddProcessRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
type APIServiceServer interface {
	// define the interface and data type
	PostLog(context.Context, *models.Log) (*ObjectResponse, error)
	GetProjectByID(context.Context, *QueryRequest) (*models.ProjectInfo, error)
	//  rpc GetProjects(QueryRequest) returns(ListResponse){}
	GetWorkflowById(context.Context, *QueryRequest) (*models.WorkflowInfo, error)
	GetSystemLogId(context.Context, *LogIdRequest) (*LogIdResponse, error)
	GetWorkflowJobById(context.Context, *QueryRequest) (*models.WorkflowJob, error)
	CreateOrUpdateWorkflowExecution(context.Context, *models.WorkflowExecution) (*StatusResponse, error)
	CreateExecutionRecord(context.Context, *models.ExecutionRecord) (*StatusResponse, error)
	AddProcessRecord(context.Context, *models.ProcessRecord) (*StatusResponse, error)
}

// UnimplementedAPIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (*UnimplementedAPIServiceServer) PostLog(context.Context, *models.Log) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostLog not implemented")
}
func (*UnimplementedAPIServiceServer) GetProjectByID(context.Context, *QueryRequest) (*models.ProjectInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectByID not implemented")
}
func (*UnimplementedAPIServiceServer) GetWorkflowById(context.Context, *QueryRequest) (*models.WorkflowInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowById not implemented")
}
func (*UnimplementedAPIServiceServer) GetSystemLogId(context.Context, *LogIdRequest) (*LogIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemLogId not implemented")
}
func (*UnimplementedAPIServiceServer) GetWorkflowJobById(context.Context, *QueryRequest) (*models.WorkflowJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowJobById not implemented")
}
func (*UnimplementedAPIServiceServer) CreateOrUpdateWorkflowExecution(context.Context, *models.WorkflowExecution) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateWorkflowExecution not implemented")
}
func (*UnimplementedAPIServiceServer) CreateExecutionRecord(context.Context, *models.ExecutionRecord) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExecutionRecord not implemented")
}
func (*UnimplementedAPIServiceServer) AddProcessRecord(context.Context, *models.ProcessRecord) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProcessRecord not implemented")
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_PostLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).PostLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/PostLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).PostLog(ctx, req.(*models.Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetProjectByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetProjectByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/GetProjectByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetProjectByID(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetWorkflowById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetWorkflowById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/GetWorkflowById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetWorkflowById(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetSystemLogId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetSystemLogId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/GetSystemLogId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetSystemLogId(ctx, req.(*LogIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetWorkflowJobById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetWorkflowJobById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/GetWorkflowJobById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetWorkflowJobById(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_CreateOrUpdateWorkflowExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.WorkflowExecution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).CreateOrUpdateWorkflowExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/CreateOrUpdateWorkflowExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).CreateOrUpdateWorkflowExecution(ctx, req.(*models.WorkflowExecution))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_CreateExecutionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ExecutionRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).CreateExecutionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/CreateExecutionRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).CreateExecutionRecord(ctx, req.(*models.ExecutionRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_AddProcessRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ProcessRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).AddProcessRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.APIService/AddProcessRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).AddProcessRecord(ctx, req.(*models.ProcessRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostLog",
			Handler:    _APIService_PostLog_Handler,
		},
		{
			MethodName: "GetProjectByID",
			Handler:    _APIService_GetProjectByID_Handler,
		},
		{
			MethodName: "GetWorkflowById",
			Handler:    _APIService_GetWorkflowById_Handler,
		},
		{
			MethodName: "GetSystemLogId",
			Handler:    _APIService_GetSystemLogId_Handler,
		},
		{
			MethodName: "GetWorkflowJobById",
			Handler:    _APIService_GetWorkflowJobById_Handler,
		},
		{
			MethodName: "CreateOrUpdateWorkflowExecution",
			Handler:    _APIService_CreateOrUpdateWorkflowExecution_Handler,
		},
		{
			MethodName: "CreateExecutionRecord",
			Handler:    _APIService_CreateExecutionRecord_Handler,
		},
		{
			MethodName: "AddProcessRecord",
			Handler:    _APIService_AddProcessRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tiops/common/services/api.proto",
}
