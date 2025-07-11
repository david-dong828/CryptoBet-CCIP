// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: write_error.proto

package writetarget

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WT errored while processing write request
type WriteError struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Code      uint32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Summary   string                 `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Cause     string                 `protobuf:"bytes,3,opt,name=cause,proto3" json:"cause,omitempty"`
	Node      string                 `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	Forwarder string                 `protobuf:"bytes,5,opt,name=forwarder,proto3" json:"forwarder,omitempty"`
	Receiver  string                 `protobuf:"bytes,6,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Report Info
	ReportId uint32 `protobuf:"varint,7,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	// [Execution Context]
	// TODO: replace with a proto reference once supported
	// Execution Context - Source
	MetaSourceId string `protobuf:"bytes,20,opt,name=meta_source_id,json=metaSourceId,proto3" json:"meta_source_id,omitempty"`
	// Execution Context - Chain
	MetaChainFamilyName string `protobuf:"bytes,21,opt,name=meta_chain_family_name,json=metaChainFamilyName,proto3" json:"meta_chain_family_name,omitempty"`
	MetaChainId         string `protobuf:"bytes,22,opt,name=meta_chain_id,json=metaChainId,proto3" json:"meta_chain_id,omitempty"`
	MetaNetworkName     string `protobuf:"bytes,23,opt,name=meta_network_name,json=metaNetworkName,proto3" json:"meta_network_name,omitempty"`
	MetaNetworkNameFull string `protobuf:"bytes,24,opt,name=meta_network_name_full,json=metaNetworkNameFull,proto3" json:"meta_network_name_full,omitempty"`
	// Execution Context - Workflow (capabilities.RequestMetadata)
	MetaWorkflowId               string `protobuf:"bytes,25,opt,name=meta_workflow_id,json=metaWorkflowId,proto3" json:"meta_workflow_id,omitempty"`
	MetaWorkflowOwner            string `protobuf:"bytes,26,opt,name=meta_workflow_owner,json=metaWorkflowOwner,proto3" json:"meta_workflow_owner,omitempty"`
	MetaWorkflowExecutionId      string `protobuf:"bytes,27,opt,name=meta_workflow_execution_id,json=metaWorkflowExecutionId,proto3" json:"meta_workflow_execution_id,omitempty"`
	MetaWorkflowName             string `protobuf:"bytes,28,opt,name=meta_workflow_name,json=metaWorkflowName,proto3" json:"meta_workflow_name,omitempty"`
	MetaWorkflowDonId            uint32 `protobuf:"varint,29,opt,name=meta_workflow_don_id,json=metaWorkflowDonId,proto3" json:"meta_workflow_don_id,omitempty"`
	MetaWorkflowDonConfigVersion uint32 `protobuf:"varint,30,opt,name=meta_workflow_don_config_version,json=metaWorkflowDonConfigVersion,proto3" json:"meta_workflow_don_config_version,omitempty"`
	MetaReferenceId              string `protobuf:"bytes,31,opt,name=meta_reference_id,json=metaReferenceId,proto3" json:"meta_reference_id,omitempty"`
	// Execution Context - Capability
	MetaCapabilityType           string `protobuf:"bytes,32,opt,name=meta_capability_type,json=metaCapabilityType,proto3" json:"meta_capability_type,omitempty"`
	MetaCapabilityId             string `protobuf:"bytes,33,opt,name=meta_capability_id,json=metaCapabilityId,proto3" json:"meta_capability_id,omitempty"`
	MetaCapabilityTimestampStart uint64 `protobuf:"varint,34,opt,name=meta_capability_timestamp_start,json=metaCapabilityTimestampStart,proto3" json:"meta_capability_timestamp_start,omitempty"`
	MetaCapabilityTimestampEmit  uint64 `protobuf:"varint,35,opt,name=meta_capability_timestamp_emit,json=metaCapabilityTimestampEmit,proto3" json:"meta_capability_timestamp_emit,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *WriteError) Reset() {
	*x = WriteError{}
	mi := &file_write_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteError) ProtoMessage() {}

func (x *WriteError) ProtoReflect() protoreflect.Message {
	mi := &file_write_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteError.ProtoReflect.Descriptor instead.
func (*WriteError) Descriptor() ([]byte, []int) {
	return file_write_error_proto_rawDescGZIP(), []int{0}
}

func (x *WriteError) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *WriteError) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *WriteError) GetCause() string {
	if x != nil {
		return x.Cause
	}
	return ""
}

func (x *WriteError) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *WriteError) GetForwarder() string {
	if x != nil {
		return x.Forwarder
	}
	return ""
}

func (x *WriteError) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *WriteError) GetReportId() uint32 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

func (x *WriteError) GetMetaSourceId() string {
	if x != nil {
		return x.MetaSourceId
	}
	return ""
}

func (x *WriteError) GetMetaChainFamilyName() string {
	if x != nil {
		return x.MetaChainFamilyName
	}
	return ""
}

func (x *WriteError) GetMetaChainId() string {
	if x != nil {
		return x.MetaChainId
	}
	return ""
}

func (x *WriteError) GetMetaNetworkName() string {
	if x != nil {
		return x.MetaNetworkName
	}
	return ""
}

func (x *WriteError) GetMetaNetworkNameFull() string {
	if x != nil {
		return x.MetaNetworkNameFull
	}
	return ""
}

func (x *WriteError) GetMetaWorkflowId() string {
	if x != nil {
		return x.MetaWorkflowId
	}
	return ""
}

func (x *WriteError) GetMetaWorkflowOwner() string {
	if x != nil {
		return x.MetaWorkflowOwner
	}
	return ""
}

func (x *WriteError) GetMetaWorkflowExecutionId() string {
	if x != nil {
		return x.MetaWorkflowExecutionId
	}
	return ""
}

func (x *WriteError) GetMetaWorkflowName() string {
	if x != nil {
		return x.MetaWorkflowName
	}
	return ""
}

func (x *WriteError) GetMetaWorkflowDonId() uint32 {
	if x != nil {
		return x.MetaWorkflowDonId
	}
	return 0
}

func (x *WriteError) GetMetaWorkflowDonConfigVersion() uint32 {
	if x != nil {
		return x.MetaWorkflowDonConfigVersion
	}
	return 0
}

func (x *WriteError) GetMetaReferenceId() string {
	if x != nil {
		return x.MetaReferenceId
	}
	return ""
}

func (x *WriteError) GetMetaCapabilityType() string {
	if x != nil {
		return x.MetaCapabilityType
	}
	return ""
}

func (x *WriteError) GetMetaCapabilityId() string {
	if x != nil {
		return x.MetaCapabilityId
	}
	return ""
}

func (x *WriteError) GetMetaCapabilityTimestampStart() uint64 {
	if x != nil {
		return x.MetaCapabilityTimestampStart
	}
	return 0
}

func (x *WriteError) GetMetaCapabilityTimestampEmit() uint64 {
	if x != nil {
		return x.MetaCapabilityTimestampEmit
	}
	return 0
}

var File_write_error_proto protoreflect.FileDescriptor

const file_write_error_proto_rawDesc = "" +
	"\n" +
	"\x11write_error.proto\x12\x15platform.write_target\"\xf1\a\n" +
	"\n" +
	"WriteError\x12\x12\n" +
	"\x04code\x18\x01 \x01(\rR\x04code\x12\x18\n" +
	"\asummary\x18\x02 \x01(\tR\asummary\x12\x14\n" +
	"\x05cause\x18\x03 \x01(\tR\x05cause\x12\x12\n" +
	"\x04node\x18\x04 \x01(\tR\x04node\x12\x1c\n" +
	"\tforwarder\x18\x05 \x01(\tR\tforwarder\x12\x1a\n" +
	"\breceiver\x18\x06 \x01(\tR\breceiver\x12\x1b\n" +
	"\treport_id\x18\a \x01(\rR\breportId\x12$\n" +
	"\x0emeta_source_id\x18\x14 \x01(\tR\fmetaSourceId\x123\n" +
	"\x16meta_chain_family_name\x18\x15 \x01(\tR\x13metaChainFamilyName\x12\"\n" +
	"\rmeta_chain_id\x18\x16 \x01(\tR\vmetaChainId\x12*\n" +
	"\x11meta_network_name\x18\x17 \x01(\tR\x0fmetaNetworkName\x123\n" +
	"\x16meta_network_name_full\x18\x18 \x01(\tR\x13metaNetworkNameFull\x12(\n" +
	"\x10meta_workflow_id\x18\x19 \x01(\tR\x0emetaWorkflowId\x12.\n" +
	"\x13meta_workflow_owner\x18\x1a \x01(\tR\x11metaWorkflowOwner\x12;\n" +
	"\x1ameta_workflow_execution_id\x18\x1b \x01(\tR\x17metaWorkflowExecutionId\x12,\n" +
	"\x12meta_workflow_name\x18\x1c \x01(\tR\x10metaWorkflowName\x12/\n" +
	"\x14meta_workflow_don_id\x18\x1d \x01(\rR\x11metaWorkflowDonId\x12F\n" +
	" meta_workflow_don_config_version\x18\x1e \x01(\rR\x1cmetaWorkflowDonConfigVersion\x12*\n" +
	"\x11meta_reference_id\x18\x1f \x01(\tR\x0fmetaReferenceId\x120\n" +
	"\x14meta_capability_type\x18  \x01(\tR\x12metaCapabilityType\x12,\n" +
	"\x12meta_capability_id\x18! \x01(\tR\x10metaCapabilityId\x12E\n" +
	"\x1fmeta_capability_timestamp_start\x18\" \x01(\x04R\x1cmetaCapabilityTimestampStart\x12C\n" +
	"\x1emeta_capability_timestamp_emit\x18# \x01(\x04R\x1bmetaCapabilityTimestampEmitB\x0fZ\r.;writetargetb\x06proto3"

var (
	file_write_error_proto_rawDescOnce sync.Once
	file_write_error_proto_rawDescData []byte
)

func file_write_error_proto_rawDescGZIP() []byte {
	file_write_error_proto_rawDescOnce.Do(func() {
		file_write_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_write_error_proto_rawDesc), len(file_write_error_proto_rawDesc)))
	})
	return file_write_error_proto_rawDescData
}

var file_write_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_write_error_proto_goTypes = []any{
	(*WriteError)(nil), // 0: platform.write_target.WriteError
}
var file_write_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_write_error_proto_init() }
func file_write_error_proto_init() {
	if File_write_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_write_error_proto_rawDesc), len(file_write_error_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_write_error_proto_goTypes,
		DependencyIndexes: file_write_error_proto_depIdxs,
		MessageInfos:      file_write_error_proto_msgTypes,
	}.Build()
	File_write_error_proto = out.File
	file_write_error_proto_goTypes = nil
	file_write_error_proto_depIdxs = nil
}
