// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: tiops/common/models/workflow_execution.proto

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

type WorkflowExecutionStatus int32

const (
	WorkflowExecutionStatus_Created   WorkflowExecutionStatus = 0
	WorkflowExecutionStatus_Queuing   WorkflowExecutionStatus = 1
	WorkflowExecutionStatus_Running   WorkflowExecutionStatus = 2
	WorkflowExecutionStatus_Succeeded WorkflowExecutionStatus = 3
	WorkflowExecutionStatus_Failed    WorkflowExecutionStatus = 4
	WorkflowExecutionStatus_Stopped   WorkflowExecutionStatus = 5
	WorkflowExecutionStatus_Abort     WorkflowExecutionStatus = 6
)

// Enum value maps for WorkflowExecutionStatus.
var (
	WorkflowExecutionStatus_name = map[int32]string{
		0: "Created",
		1: "Queuing",
		2: "Running",
		3: "Succeeded",
		4: "Failed",
		5: "Stopped",
		6: "Abort",
	}
	WorkflowExecutionStatus_value = map[string]int32{
		"Created":   0,
		"Queuing":   1,
		"Running":   2,
		"Succeeded": 3,
		"Failed":    4,
		"Stopped":   5,
		"Abort":     6,
	}
)

func (x WorkflowExecutionStatus) Enum() *WorkflowExecutionStatus {
	p := new(WorkflowExecutionStatus)
	*p = x
	return p
}

func (x WorkflowExecutionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowExecutionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tiops_common_models_workflow_execution_proto_enumTypes[0].Descriptor()
}

func (WorkflowExecutionStatus) Type() protoreflect.EnumType {
	return &file_tiops_common_models_workflow_execution_proto_enumTypes[0]
}

func (x WorkflowExecutionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowExecutionStatus.Descriptor instead.
func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_execution_proto_rawDescGZIP(), []int{0}
}

type WorkflowExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId              string                  `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty" bson:"_id"`
	JobId            string                  `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId,omitempty" bson:"jobId"`
	WorkflowId       string                  `protobuf:"bytes,3,opt,name=workflowId,proto3" json:"workflowId,omitempty" bson:"workflowId"`
	LogId            string                  `protobuf:"bytes,4,opt,name=logId,proto3" json:"logId,omitempty" bson:"logId"`
	CreatedBy        string                  `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty" bson:"createdBy"`
	CreatedTime      int64                   `protobuf:"varint,6,opt,name=createdTime,proto3" json:"createdTime,omitempty" bson:"createdTime"`
	StartTime        int64                   `protobuf:"varint,7,opt,name=startTime,proto3" json:"startTime,omitempty" bson:"startTime"`
	EndTime          int64                   `protobuf:"varint,8,opt,name=endTime,proto3" json:"endTime,omitempty" bson:"endTime"`
	Status           WorkflowExecutionStatus `protobuf:"varint,9,opt,name=status,proto3,enum=models.WorkflowExecutionStatus" json:"status,omitempty" bson:"status"`
	Workspace        string                  `protobuf:"bytes,10,opt,name=workspace,proto3" json:"workspace,omitempty" bson:"workspace"`
	WorkflowResource *WorkflowResources      `protobuf:"bytes,11,opt,name=workflowResource,proto3" json:"workflowResource,omitempty" bson:"workflowResource"`
}

func (x *WorkflowExecution) Reset() {
	*x = WorkflowExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecution) ProtoMessage() {}

func (x *WorkflowExecution) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecution.ProtoReflect.Descriptor instead.
func (*WorkflowExecution) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_execution_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowExecution) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *WorkflowExecution) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *WorkflowExecution) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowExecution) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *WorkflowExecution) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *WorkflowExecution) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *WorkflowExecution) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *WorkflowExecution) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *WorkflowExecution) GetStatus() WorkflowExecutionStatus {
	if x != nil {
		return x.Status
	}
	return WorkflowExecutionStatus_Created
}

func (x *WorkflowExecution) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *WorkflowExecution) GetWorkflowResource() *WorkflowResources {
	if x != nil {
		return x.WorkflowResource
	}
	return nil
}

var File_tiops_common_models_workflow_execution_proto protoreflect.FileDescriptor

var file_tiops_common_models_workflow_execution_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x20, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x69, 0x6f, 0x70,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03,
	0x0a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2a, 0x73, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x51, 0x75, 0x65,
	0x75, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x10, 0x06, 0x42, 0x2d, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x74, 0x69, 0x6f,
	0x70, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x5a, 0x13, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tiops_common_models_workflow_execution_proto_rawDescOnce sync.Once
	file_tiops_common_models_workflow_execution_proto_rawDescData = file_tiops_common_models_workflow_execution_proto_rawDesc
)

func file_tiops_common_models_workflow_execution_proto_rawDescGZIP() []byte {
	file_tiops_common_models_workflow_execution_proto_rawDescOnce.Do(func() {
		file_tiops_common_models_workflow_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiops_common_models_workflow_execution_proto_rawDescData)
	})
	return file_tiops_common_models_workflow_execution_proto_rawDescData
}

var file_tiops_common_models_workflow_execution_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tiops_common_models_workflow_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tiops_common_models_workflow_execution_proto_goTypes = []interface{}{
	(WorkflowExecutionStatus)(0), // 0: models.WorkflowExecutionStatus
	(*WorkflowExecution)(nil),    // 1: models.WorkflowExecution
	(*WorkflowResources)(nil),    // 2: models.WorkflowResources
}
var file_tiops_common_models_workflow_execution_proto_depIdxs = []int32{
	0, // 0: models.WorkflowExecution.status:type_name -> models.WorkflowExecutionStatus
	2, // 1: models.WorkflowExecution.workflowResource:type_name -> models.WorkflowResources
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tiops_common_models_workflow_execution_proto_init() }
func file_tiops_common_models_workflow_execution_proto_init() {
	if File_tiops_common_models_workflow_execution_proto != nil {
		return
	}
	file_tiops_common_models_common_proto_init()
	file_tiops_common_models_tagger_proto_init()
	file_tiops_common_models_workflow_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tiops_common_models_workflow_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowExecution); i {
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
			RawDescriptor: file_tiops_common_models_workflow_execution_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tiops_common_models_workflow_execution_proto_goTypes,
		DependencyIndexes: file_tiops_common_models_workflow_execution_proto_depIdxs,
		EnumInfos:         file_tiops_common_models_workflow_execution_proto_enumTypes,
		MessageInfos:      file_tiops_common_models_workflow_execution_proto_msgTypes,
	}.Build()
	File_tiops_common_models_workflow_execution_proto = out.File
	file_tiops_common_models_workflow_execution_proto_rawDesc = nil
	file_tiops_common_models_workflow_execution_proto_goTypes = nil
	file_tiops_common_models_workflow_execution_proto_depIdxs = nil
}
