// Configuration proto for GCE targets.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: github.com/google/cloudprober/targets/gce/proto/config.proto

package proto

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

type Instances_NetworkInterface_IPType int32

const (
	// Private IP address.
	Instances_NetworkInterface_PRIVATE Instances_NetworkInterface_IPType = 0
	// IP address of the first access config.
	Instances_NetworkInterface_PUBLIC Instances_NetworkInterface_IPType = 1
	// First IP address from the first Alias IP range. For example, for
	// alias IP range "192.168.12.0/24", 192.168.12.0 will be returned.
	Instances_NetworkInterface_ALIAS Instances_NetworkInterface_IPType = 2
)

// Enum value maps for Instances_NetworkInterface_IPType.
var (
	Instances_NetworkInterface_IPType_name = map[int32]string{
		0: "PRIVATE",
		1: "PUBLIC",
		2: "ALIAS",
	}
	Instances_NetworkInterface_IPType_value = map[string]int32{
		"PRIVATE": 0,
		"PUBLIC":  1,
		"ALIAS":   2,
	}
)

func (x Instances_NetworkInterface_IPType) Enum() *Instances_NetworkInterface_IPType {
	p := new(Instances_NetworkInterface_IPType)
	*p = x
	return p
}

func (x Instances_NetworkInterface_IPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Instances_NetworkInterface_IPType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_enumTypes[0].Descriptor()
}

func (Instances_NetworkInterface_IPType) Type() protoreflect.EnumType {
	return &file_github_com_google_cloudprober_targets_gce_proto_config_proto_enumTypes[0]
}

func (x Instances_NetworkInterface_IPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Instances_NetworkInterface_IPType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Instances_NetworkInterface_IPType(num)
	return nil
}

// Deprecated: Use Instances_NetworkInterface_IPType.Descriptor instead.
func (Instances_NetworkInterface_IPType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP(), []int{1, 0, 0}
}

// TargetsConf represents GCE targets, e.g. instances, forwarding rules etc.
type TargetsConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If running on GCE, this defaults to the local project.
	// Note: Multiple projects support in targets is experimental and may go away
	// with future iterations.
	Project []string `protobuf:"bytes,1,rep,name=project" json:"project,omitempty"`
	// Types that are assignable to Type:
	//	*TargetsConf_Instances
	//	*TargetsConf_ForwardingRules
	Type isTargetsConf_Type `protobuf_oneof:"type"`
}

func (x *TargetsConf) Reset() {
	*x = TargetsConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetsConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetsConf) ProtoMessage() {}

func (x *TargetsConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetsConf.ProtoReflect.Descriptor instead.
func (*TargetsConf) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *TargetsConf) GetProject() []string {
	if x != nil {
		return x.Project
	}
	return nil
}

func (m *TargetsConf) GetType() isTargetsConf_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *TargetsConf) GetInstances() *Instances {
	if x, ok := x.GetType().(*TargetsConf_Instances); ok {
		return x.Instances
	}
	return nil
}

func (x *TargetsConf) GetForwardingRules() *ForwardingRules {
	if x, ok := x.GetType().(*TargetsConf_ForwardingRules); ok {
		return x.ForwardingRules
	}
	return nil
}

type isTargetsConf_Type interface {
	isTargetsConf_Type()
}

type TargetsConf_Instances struct {
	Instances *Instances `protobuf:"bytes,2,opt,name=instances,oneof"`
}

type TargetsConf_ForwardingRules struct {
	ForwardingRules *ForwardingRules `protobuf:"bytes,3,opt,name=forwarding_rules,json=forwardingRules,oneof"`
}

func (*TargetsConf_Instances) isTargetsConf_Type() {}

func (*TargetsConf_ForwardingRules) isTargetsConf_Type() {}

// Represents GCE instances
type Instances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use DNS to resolve target names (instances). If set to false (default),
	// IP addresses specified in the compute.Instance resource is used. If set
	// to true all the other resolving options are ignored.
	UseDnsToResolve  *bool                       `protobuf:"varint,1,opt,name=use_dns_to_resolve,json=useDnsToResolve,def=0" json:"use_dns_to_resolve,omitempty"`
	NetworkInterface *Instances_NetworkInterface `protobuf:"bytes,2,opt,name=network_interface,json=networkInterface" json:"network_interface,omitempty"`
	// Labels to filter instances by ("key:value-regex" format).
	Label []string `protobuf:"bytes,3,rep,name=label" json:"label,omitempty"`
}

// Default values for Instances fields.
const (
	Default_Instances_UseDnsToResolve = bool(false)
)

func (x *Instances) Reset() {
	*x = Instances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instances) ProtoMessage() {}

func (x *Instances) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instances.ProtoReflect.Descriptor instead.
func (*Instances) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *Instances) GetUseDnsToResolve() bool {
	if x != nil && x.UseDnsToResolve != nil {
		return *x.UseDnsToResolve
	}
	return Default_Instances_UseDnsToResolve
}

func (x *Instances) GetNetworkInterface() *Instances_NetworkInterface {
	if x != nil {
		return x.NetworkInterface
	}
	return nil
}

func (x *Instances) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

// Represents GCE forwarding rules. Does not support multiple projects
type ForwardingRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Important: if multiple probes use forwarding_rules targets, only the
	// settings in the definition will take effect.
	// TODO(manugarg): Fix this behavior.
	//
	// For regional forwarding rules, regions to return forwarding rules for.
	// Default is to return forwarding rules from the region that the VM is
	// running in. To return forwarding rules from all regions, specify region as
	// "all".
	Region []string `protobuf:"bytes,1,rep,name=region" json:"region,omitempty"`
}

func (x *ForwardingRules) Reset() {
	*x = ForwardingRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardingRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardingRules) ProtoMessage() {}

func (x *ForwardingRules) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardingRules.ProtoReflect.Descriptor instead.
func (*ForwardingRules) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP(), []int{2}
}

func (x *ForwardingRules) GetRegion() []string {
	if x != nil {
		return x.Region
	}
	return nil
}

// Global GCE targets options. These options are independent of the per-probe
// targets which are defined by the "GCETargets" type above.
type GlobalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How often targets should be evaluated/expanded
	ReEvalSec *int32 `protobuf:"varint,1,opt,name=re_eval_sec,json=reEvalSec,def=900" json:"re_eval_sec,omitempty"` // default 15 min
	// Compute API version.
	ApiVersion *string `protobuf:"bytes,2,opt,name=api_version,json=apiVersion,def=v1" json:"api_version,omitempty"`
}

// Default values for GlobalOptions fields.
const (
	Default_GlobalOptions_ReEvalSec  = int32(900)
	Default_GlobalOptions_ApiVersion = string("v1")
)

func (x *GlobalOptions) Reset() {
	*x = GlobalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalOptions) ProtoMessage() {}

func (x *GlobalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalOptions.ProtoReflect.Descriptor instead.
func (*GlobalOptions) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP(), []int{3}
}

func (x *GlobalOptions) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_GlobalOptions_ReEvalSec
}

func (x *GlobalOptions) GetApiVersion() string {
	if x != nil && x.ApiVersion != nil {
		return *x.ApiVersion
	}
	return Default_GlobalOptions_ApiVersion
}

// Get the IP address from Network Interface
type Instances_NetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  *int32                             `protobuf:"varint,1,opt,name=index,def=0" json:"index,omitempty"`
	IpType *Instances_NetworkInterface_IPType `protobuf:"varint,2,opt,name=ip_type,json=ipType,enum=cloudprober.targets.gce.Instances_NetworkInterface_IPType,def=0" json:"ip_type,omitempty"`
}

// Default values for Instances_NetworkInterface fields.
const (
	Default_Instances_NetworkInterface_Index  = int32(0)
	Default_Instances_NetworkInterface_IpType = Instances_NetworkInterface_PRIVATE
)

func (x *Instances_NetworkInterface) Reset() {
	*x = Instances_NetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instances_NetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instances_NetworkInterface) ProtoMessage() {}

func (x *Instances_NetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instances_NetworkInterface.ProtoReflect.Descriptor instead.
func (*Instances_NetworkInterface) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Instances_NetworkInterface) GetIndex() int32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return Default_Instances_NetworkInterface_Index
}

func (x *Instances_NetworkInterface) GetIpType() Instances_NetworkInterface_IPType {
	if x != nil && x.IpType != nil {
		return *x.IpType
	}
	return Default_Instances_NetworkInterface_IpType
}

var File_github_com_google_cloudprober_targets_gce_proto_config_proto protoreflect.FileDescriptor

var file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x2e, 0x67, 0x63, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x42, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xf1, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x6f,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x44, 0x6e, 0x73, 0x54, 0x6f, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0xb7,
	0x01, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x5c, 0x0a, 0x07,
	0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x2e, 0x67, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x49, 0x50, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x52, 0x06, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x49, 0x50,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x4c, 0x49, 0x41, 0x53, 0x10, 0x02, 0x22, 0x29, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x0d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x39, 0x30, 0x30, 0x52, 0x09,
	0x72, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x02,
	0x76, 0x31, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f,
}

var (
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescOnce sync.Once
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescData = file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDesc
)

func file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescData)
	})
	return file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDescData
}

var file_github_com_google_cloudprober_targets_gce_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_google_cloudprober_targets_gce_proto_config_proto_goTypes = []interface{}{
	(Instances_NetworkInterface_IPType)(0), // 0: cloudprober.targets.gce.Instances.NetworkInterface.IPType
	(*TargetsConf)(nil),                    // 1: cloudprober.targets.gce.TargetsConf
	(*Instances)(nil),                      // 2: cloudprober.targets.gce.Instances
	(*ForwardingRules)(nil),                // 3: cloudprober.targets.gce.ForwardingRules
	(*GlobalOptions)(nil),                  // 4: cloudprober.targets.gce.GlobalOptions
	(*Instances_NetworkInterface)(nil),     // 5: cloudprober.targets.gce.Instances.NetworkInterface
}
var file_github_com_google_cloudprober_targets_gce_proto_config_proto_depIdxs = []int32{
	2, // 0: cloudprober.targets.gce.TargetsConf.instances:type_name -> cloudprober.targets.gce.Instances
	3, // 1: cloudprober.targets.gce.TargetsConf.forwarding_rules:type_name -> cloudprober.targets.gce.ForwardingRules
	5, // 2: cloudprober.targets.gce.Instances.network_interface:type_name -> cloudprober.targets.gce.Instances.NetworkInterface
	0, // 3: cloudprober.targets.gce.Instances.NetworkInterface.ip_type:type_name -> cloudprober.targets.gce.Instances.NetworkInterface.IPType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_google_cloudprober_targets_gce_proto_config_proto_init() }
func file_github_com_google_cloudprober_targets_gce_proto_config_proto_init() {
	if File_github_com_google_cloudprober_targets_gce_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetsConf); i {
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
		file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instances); i {
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
		file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardingRules); i {
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
		file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalOptions); i {
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
		file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instances_NetworkInterface); i {
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
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TargetsConf_Instances)(nil),
		(*TargetsConf_ForwardingRules)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_google_cloudprober_targets_gce_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_google_cloudprober_targets_gce_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_google_cloudprober_targets_gce_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_google_cloudprober_targets_gce_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_google_cloudprober_targets_gce_proto_config_proto = out.File
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_rawDesc = nil
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_goTypes = nil
	file_github_com_google_cloudprober_targets_gce_proto_config_proto_depIdxs = nil
}
