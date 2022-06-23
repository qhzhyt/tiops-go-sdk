// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: tiops/common/models/workflow-resource.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

type ServiceMode int32

const (
	ServiceMode_Nil     ServiceMode = 0
	ServiceMode_One     ServiceMode = 1
	ServiceMode_Replica ServiceMode = 2
)

// Enum value maps for ServiceMode.
var (
	ServiceMode_name = map[int32]string{
		0: "Nil",
		1: "One",
		2: "Replica",
	}
	ServiceMode_value = map[string]int32{
		"Nil":     0,
		"One":     1,
		"Replica": 2,
	}
)

func (x ServiceMode) Enum() *ServiceMode {
	p := new(ServiceMode)
	*p = x
	return p
}

func (x ServiceMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceMode) Descriptor() protoreflect.EnumDescriptor {
	return file_tiops_common_models_workflow_resource_proto_enumTypes[0].Descriptor()
}

func (ServiceMode) Type() protoreflect.EnumType {
	return &file_tiops_common_models_workflow_resource_proto_enumTypes[0]
}

func (x ServiceMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceMode.Descriptor instead.
func (ServiceMode) EnumDescriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{0}
}

type K8SAppType int32

const (
	K8SAppType_K8sDeployment K8SAppType = 0
	K8SAppType_K8sJob        K8SAppType = 1
)

// Enum value maps for K8SAppType.
var (
	K8SAppType_name = map[int32]string{
		0: "K8sDeployment",
		1: "K8sJob",
	}
	K8SAppType_value = map[string]int32{
		"K8sDeployment": 0,
		"K8sJob":        1,
	}
)

func (x K8SAppType) Enum() *K8SAppType {
	p := new(K8SAppType)
	*p = x
	return p
}

func (x K8SAppType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SAppType) Descriptor() protoreflect.EnumDescriptor {
	return file_tiops_common_models_workflow_resource_proto_enumTypes[1].Descriptor()
}

func (K8SAppType) Type() protoreflect.EnumType {
	return &file_tiops_common_models_workflow_resource_proto_enumTypes[1]
}

func (x K8SAppType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SAppType.Descriptor instead.
func (K8SAppType) EnumDescriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{1}
}

type ContainerResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu    string            `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty" bson:"cpu"`
	Memory string            `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty" bson:"memory"`
	Gpu    string            `protobuf:"bytes,3,opt,name=gpu,proto3" json:"gpu,omitempty" bson:"gpu"`
	Extra  map[string]string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
}

func (x *ContainerResources) Reset() {
	*x = ContainerResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerResources) ProtoMessage() {}

func (x *ContainerResources) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerResources.ProtoReflect.Descriptor instead.
func (*ContainerResources) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerResources) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *ContainerResources) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *ContainerResources) GetGpu() string {
	if x != nil {
		return x.Gpu
	}
	return ""
}

func (x *ContainerResources) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type K8SContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Image             string              `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty" bson:"image"`
	Cmd               []string            `protobuf:"bytes,3,rep,name=cmd,proto3" json:"cmd,omitempty" bson:"cmd"`
	Args              []string            `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty" bson:"args"`
	Env               map[string]string   `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"env"`
	ResourcesRequests *ContainerResources `protobuf:"bytes,6,opt,name=resourcesRequests,proto3" json:"resourcesRequests,omitempty" bson:"resourcesRequests"`
	ResourcesLimits   *ContainerResources `protobuf:"bytes,7,opt,name=resourcesLimits,proto3" json:"resourcesLimits,omitempty" bson:"resourcesLimits"`
}

func (x *K8SContainer) Reset() {
	*x = K8SContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SContainer) ProtoMessage() {}

func (x *K8SContainer) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SContainer.ProtoReflect.Descriptor instead.
func (*K8SContainer) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{1}
}

func (x *K8SContainer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SContainer) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *K8SContainer) GetCmd() []string {
	if x != nil {
		return x.Cmd
	}
	return nil
}

func (x *K8SContainer) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *K8SContainer) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *K8SContainer) GetResourcesRequests() *ContainerResources {
	if x != nil {
		return x.ResourcesRequests
	}
	return nil
}

func (x *K8SContainer) GetResourcesLimits() *ContainerResources {
	if x != nil {
		return x.ResourcesLimits
	}
	return nil
}

type K8SPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	ContainerPort int32  `protobuf:"varint,2,opt,name=containerPort,proto3" json:"containerPort,omitempty" bson:"containerPort"`
	ServicePort   int32  `protobuf:"varint,3,opt,name=servicePort,proto3" json:"servicePort,omitempty" bson:"servicePort"`
	NodePort      int32  `protobuf:"varint,4,opt,name=nodePort,proto3" json:"nodePort,omitempty" bson:"nodePort"`
	Proto         string `protobuf:"bytes,5,opt,name=proto,proto3" json:"proto,omitempty" bson:"proto"`
}

func (x *K8SPort) Reset() {
	*x = K8SPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SPort) ProtoMessage() {}

func (x *K8SPort) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SPort.ProtoReflect.Descriptor instead.
func (*K8SPort) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{2}
}

func (x *K8SPort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SPort) GetContainerPort() int32 {
	if x != nil {
		return x.ContainerPort
	}
	return 0
}

func (x *K8SPort) GetServicePort() int32 {
	if x != nil {
		return x.ServicePort
	}
	return 0
}

func (x *K8SPort) GetNodePort() int32 {
	if x != nil {
		return x.NodePort
	}
	return 0
}

func (x *K8SPort) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

type K8SService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Ports []*K8SPort `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty" bson:"ports"`
}

func (x *K8SService) Reset() {
	*x = K8SService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SService) ProtoMessage() {}

func (x *K8SService) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SService.ProtoReflect.Descriptor instead.
func (*K8SService) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{3}
}

func (x *K8SService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SService) GetPorts() []*K8SPort {
	if x != nil {
		return x.Ports
	}
	return nil
}

type K8SApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	ActionId string `protobuf:"bytes,2,opt,name=actionId,proto3" json:"actionId,omitempty" bson:"actionId"`
	NodeId   string `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty" bson:"nodeId"`
	Replica  int32  `protobuf:"varint,4,opt,name=replica,proto3" json:"replica,omitempty" bson:"replica"`
	//  K8sContainer mainContainer = 4;
	WorkContainers  []*K8SContainer   `protobuf:"bytes,5,rep,name=workContainers,proto3" json:"workContainers,omitempty" bson:"workContainers"`
	InitContainers  []*K8SContainer   `protobuf:"bytes,6,rep,name=initContainers,proto3" json:"initContainers,omitempty" bson:"initContainers"`
	ServiceMode     ServiceMode       `protobuf:"varint,7,opt,name=serviceMode,proto3,enum=models.ServiceMode" json:"serviceMode,omitempty" bson:"serviceMode"`
	Volumes         map[string]string `protobuf:"bytes,8,rep,name=volumes,proto3" json:"volumes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"volumes"`
	ServiceTCPPorts map[string]int32  `protobuf:"bytes,9,rep,name=serviceTCPPorts,proto3" json:"serviceTCPPorts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" bson:"serviceTCPPorts"`
	ServiceUDPPorts map[string]int32  `protobuf:"bytes,10,rep,name=serviceUDPPorts,proto3" json:"serviceUDPPorts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" bson:"serviceUDPPorts"`
	Type            K8SAppType        `protobuf:"varint,11,opt,name=type,proto3,enum=models.K8SAppType" json:"type,omitempty" bson:"type"`
}

func (x *K8SApp) Reset() {
	*x = K8SApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SApp) ProtoMessage() {}

func (x *K8SApp) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SApp.ProtoReflect.Descriptor instead.
func (*K8SApp) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{4}
}

func (x *K8SApp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SApp) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *K8SApp) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *K8SApp) GetReplica() int32 {
	if x != nil {
		return x.Replica
	}
	return 0
}

func (x *K8SApp) GetWorkContainers() []*K8SContainer {
	if x != nil {
		return x.WorkContainers
	}
	return nil
}

func (x *K8SApp) GetInitContainers() []*K8SContainer {
	if x != nil {
		return x.InitContainers
	}
	return nil
}

func (x *K8SApp) GetServiceMode() ServiceMode {
	if x != nil {
		return x.ServiceMode
	}
	return ServiceMode_Nil
}

func (x *K8SApp) GetVolumes() map[string]string {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (x *K8SApp) GetServiceTCPPorts() map[string]int32 {
	if x != nil {
		return x.ServiceTCPPorts
	}
	return nil
}

func (x *K8SApp) GetServiceUDPPorts() map[string]int32 {
	if x != nil {
		return x.ServiceUDPPorts
	}
	return nil
}

func (x *K8SApp) GetType() K8SAppType {
	if x != nil {
		return x.Type
	}
	return K8SAppType_K8sDeployment
}

type WorkflowResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*K8SApp `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty" bson:"apps"`
}

func (x *WorkflowResources) Reset() {
	*x = WorkflowResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowResources) ProtoMessage() {}

func (x *WorkflowResources) ProtoReflect() protoreflect.Message {
	mi := &file_tiops_common_models_workflow_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowResources.ProtoReflect.Descriptor instead.
func (*WorkflowResources) Descriptor() ([]byte, []int) {
	return file_tiops_common_models_workflow_resource_proto_rawDescGZIP(), []int{5}
}

func (x *WorkflowResources) GetApps() []*K8SApp {
	if x != nil {
		return x.Apps
	}
	return nil
}

var File_tiops_common_models_workflow_resource_proto protoreflect.FileDescriptor

var file_tiops_common_models_workflow_resource_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x20, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x70,
	0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x3b, 0x0a, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xd7, 0x02, 0x0a, 0x0c, 0x4b, 0x38, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x12, 0x2f, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x48, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x11, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x44, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x97, 0x01,
	0x0a, 0x07, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0a, 0x4b, 0x38, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x4b, 0x38, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x22, 0xde, 0x05, 0x0a, 0x06, 0x4b, 0x38, 0x73, 0x41, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x3c, 0x0a,
	0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b,
	0x38, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0e, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x0e, 0x69,
	0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x35, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x41, 0x70,
	0x70, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x43, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x41, 0x70, 0x70,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x43, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x43,
	0x50, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x55, 0x44, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x41, 0x70, 0x70, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x44, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x44, 0x50,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38, 0x73,
	0x41, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x3a, 0x0a,
	0x0c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x43, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a,
	0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x44, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x37, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4b, 0x38,
	0x73, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x2a, 0x2c, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x69, 0x6c,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x0a, 0x4b, 0x38, 0x73, 0x41,
	0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4b, 0x38, 0x73, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x38, 0x73,
	0x4a, 0x6f, 0x62, 0x10, 0x01, 0x42, 0x2d, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x74, 0x69, 0x6f, 0x70,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5a,
	0x13, 0x74, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tiops_common_models_workflow_resource_proto_rawDescOnce sync.Once
	file_tiops_common_models_workflow_resource_proto_rawDescData = file_tiops_common_models_workflow_resource_proto_rawDesc
)

func file_tiops_common_models_workflow_resource_proto_rawDescGZIP() []byte {
	file_tiops_common_models_workflow_resource_proto_rawDescOnce.Do(func() {
		file_tiops_common_models_workflow_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiops_common_models_workflow_resource_proto_rawDescData)
	})
	return file_tiops_common_models_workflow_resource_proto_rawDescData
}

var file_tiops_common_models_workflow_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tiops_common_models_workflow_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tiops_common_models_workflow_resource_proto_goTypes = []interface{}{
	(ServiceMode)(0),           // 0: models.ServiceMode
	(K8SAppType)(0),            // 1: models.K8sAppType
	(*ContainerResources)(nil), // 2: models.ContainerResources
	(*K8SContainer)(nil),       // 3: models.K8sContainer
	(*K8SPort)(nil),            // 4: models.K8sPort
	(*K8SService)(nil),         // 5: models.K8sService
	(*K8SApp)(nil),             // 6: models.K8sApp
	(*WorkflowResources)(nil),  // 7: models.WorkflowResources
	nil,                        // 8: models.ContainerResources.ExtraEntry
	nil,                        // 9: models.K8sContainer.EnvEntry
	nil,                        // 10: models.K8sApp.VolumesEntry
	nil,                        // 11: models.K8sApp.ServiceTCPPortsEntry
	nil,                        // 12: models.K8sApp.ServiceUDPPortsEntry
}
var file_tiops_common_models_workflow_resource_proto_depIdxs = []int32{
	8,  // 0: models.ContainerResources.extra:type_name -> models.ContainerResources.ExtraEntry
	9,  // 1: models.K8sContainer.env:type_name -> models.K8sContainer.EnvEntry
	2,  // 2: models.K8sContainer.resourcesRequests:type_name -> models.ContainerResources
	2,  // 3: models.K8sContainer.resourcesLimits:type_name -> models.ContainerResources
	4,  // 4: models.K8sService.ports:type_name -> models.K8sPort
	3,  // 5: models.K8sApp.workContainers:type_name -> models.K8sContainer
	3,  // 6: models.K8sApp.initContainers:type_name -> models.K8sContainer
	0,  // 7: models.K8sApp.serviceMode:type_name -> models.ServiceMode
	10, // 8: models.K8sApp.volumes:type_name -> models.K8sApp.VolumesEntry
	11, // 9: models.K8sApp.serviceTCPPorts:type_name -> models.K8sApp.ServiceTCPPortsEntry
	12, // 10: models.K8sApp.serviceUDPPorts:type_name -> models.K8sApp.ServiceUDPPortsEntry
	1,  // 11: models.K8sApp.type:type_name -> models.K8sAppType
	6,  // 12: models.WorkflowResources.apps:type_name -> models.K8sApp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_tiops_common_models_workflow_resource_proto_init() }
func file_tiops_common_models_workflow_resource_proto_init() {
	if File_tiops_common_models_workflow_resource_proto != nil {
		return
	}
	file_tiops_common_models_common_proto_init()
	file_tiops_common_models_tagger_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tiops_common_models_workflow_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerResources); i {
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
		file_tiops_common_models_workflow_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SContainer); i {
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
		file_tiops_common_models_workflow_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SPort); i {
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
		file_tiops_common_models_workflow_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SService); i {
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
		file_tiops_common_models_workflow_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SApp); i {
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
		file_tiops_common_models_workflow_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowResources); i {
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
			RawDescriptor: file_tiops_common_models_workflow_resource_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tiops_common_models_workflow_resource_proto_goTypes,
		DependencyIndexes: file_tiops_common_models_workflow_resource_proto_depIdxs,
		EnumInfos:         file_tiops_common_models_workflow_resource_proto_enumTypes,
		MessageInfos:      file_tiops_common_models_workflow_resource_proto_msgTypes,
	}.Build()
	File_tiops_common_models_workflow_resource_proto = out.File
	file_tiops_common_models_workflow_resource_proto_rawDesc = nil
	file_tiops_common_models_workflow_resource_proto_goTypes = nil
	file_tiops_common_models_workflow_resource_proto_depIdxs = nil
}
