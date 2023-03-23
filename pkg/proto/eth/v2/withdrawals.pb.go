// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/proto/eth/v2/withdrawals.proto

// Note: largely inspired by
// https://github.com/prysmaticlabs/prysm/tree/develop/proto/eth/v2

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BLSToExecutionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorIndex     uint64 `protobuf:"varint,1,opt,name=validator_index,proto3" json:"validator_index,omitempty"`
	FromBlsPubkey      string `protobuf:"bytes,2,opt,name=from_bls_pubkey,proto3" json:"from_bls_pubkey,omitempty"`
	ToExecutionAddress string `protobuf:"bytes,3,opt,name=to_execution_address,proto3" json:"to_execution_address,omitempty"`
}

func (x *BLSToExecutionChange) Reset() {
	*x = BLSToExecutionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLSToExecutionChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLSToExecutionChange) ProtoMessage() {}

func (x *BLSToExecutionChange) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLSToExecutionChange.ProtoReflect.Descriptor instead.
func (*BLSToExecutionChange) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v2_withdrawals_proto_rawDescGZIP(), []int{0}
}

func (x *BLSToExecutionChange) GetValidatorIndex() uint64 {
	if x != nil {
		return x.ValidatorIndex
	}
	return 0
}

func (x *BLSToExecutionChange) GetFromBlsPubkey() string {
	if x != nil {
		return x.FromBlsPubkey
	}
	return ""
}

func (x *BLSToExecutionChange) GetToExecutionAddress() string {
	if x != nil {
		return x.ToExecutionAddress
	}
	return ""
}

type SignedBLSToExecutionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *BLSToExecutionChange `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature string                `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedBLSToExecutionChange) Reset() {
	*x = SignedBLSToExecutionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedBLSToExecutionChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedBLSToExecutionChange) ProtoMessage() {}

func (x *SignedBLSToExecutionChange) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedBLSToExecutionChange.ProtoReflect.Descriptor instead.
func (*SignedBLSToExecutionChange) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v2_withdrawals_proto_rawDescGZIP(), []int{1}
}

func (x *SignedBLSToExecutionChange) GetMessage() *BLSToExecutionChange {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedBLSToExecutionChange) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type SubmitBLSToExecutionChangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*SignedBLSToExecutionChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *SubmitBLSToExecutionChangesRequest) Reset() {
	*x = SubmitBLSToExecutionChangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitBLSToExecutionChangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitBLSToExecutionChangesRequest) ProtoMessage() {}

func (x *SubmitBLSToExecutionChangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitBLSToExecutionChangesRequest.ProtoReflect.Descriptor instead.
func (*SubmitBLSToExecutionChangesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v2_withdrawals_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitBLSToExecutionChangesRequest) GetChanges() []*SignedBLSToExecutionChange {
	if x != nil {
		return x.Changes
	}
	return nil
}

var File_pkg_proto_eth_v2_withdrawals_proto protoreflect.FileDescriptor

var file_pkg_proto_eth_v2_withdrawals_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x32, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x14, 0x42, 0x4c, 0x53, 0x54, 0x6f, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x62,
	0x6c, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x62, 0x6c, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x12, 0x32, 0x0a, 0x14, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x74, 0x6f, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x77, 0x0a, 0x1a, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x4c,
	0x53, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x32, 0x2e, 0x42, 0x4c, 0x53, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x67, 0x0a,
	0x22, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x4c, 0x53, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x4c, 0x53, 0x54, 0x6f, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x70, 0x73,
	0x2f, 0x78, 0x61, 0x74, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_eth_v2_withdrawals_proto_rawDescOnce sync.Once
	file_pkg_proto_eth_v2_withdrawals_proto_rawDescData = file_pkg_proto_eth_v2_withdrawals_proto_rawDesc
)

func file_pkg_proto_eth_v2_withdrawals_proto_rawDescGZIP() []byte {
	file_pkg_proto_eth_v2_withdrawals_proto_rawDescOnce.Do(func() {
		file_pkg_proto_eth_v2_withdrawals_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_eth_v2_withdrawals_proto_rawDescData)
	})
	return file_pkg_proto_eth_v2_withdrawals_proto_rawDescData
}

var file_pkg_proto_eth_v2_withdrawals_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_eth_v2_withdrawals_proto_goTypes = []interface{}{
	(*BLSToExecutionChange)(nil),               // 0: xatu.eth.v2.BLSToExecutionChange
	(*SignedBLSToExecutionChange)(nil),         // 1: xatu.eth.v2.SignedBLSToExecutionChange
	(*SubmitBLSToExecutionChangesRequest)(nil), // 2: xatu.eth.v2.SubmitBLSToExecutionChangesRequest
}
var file_pkg_proto_eth_v2_withdrawals_proto_depIdxs = []int32{
	0, // 0: xatu.eth.v2.SignedBLSToExecutionChange.message:type_name -> xatu.eth.v2.BLSToExecutionChange
	1, // 1: xatu.eth.v2.SubmitBLSToExecutionChangesRequest.changes:type_name -> xatu.eth.v2.SignedBLSToExecutionChange
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_eth_v2_withdrawals_proto_init() }
func file_pkg_proto_eth_v2_withdrawals_proto_init() {
	if File_pkg_proto_eth_v2_withdrawals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLSToExecutionChange); i {
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
		file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedBLSToExecutionChange); i {
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
		file_pkg_proto_eth_v2_withdrawals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitBLSToExecutionChangesRequest); i {
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
			RawDescriptor: file_pkg_proto_eth_v2_withdrawals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_eth_v2_withdrawals_proto_goTypes,
		DependencyIndexes: file_pkg_proto_eth_v2_withdrawals_proto_depIdxs,
		MessageInfos:      file_pkg_proto_eth_v2_withdrawals_proto_msgTypes,
	}.Build()
	File_pkg_proto_eth_v2_withdrawals_proto = out.File
	file_pkg_proto_eth_v2_withdrawals_proto_rawDesc = nil
	file_pkg_proto_eth_v2_withdrawals_proto_goTypes = nil
	file_pkg_proto_eth_v2_withdrawals_proto_depIdxs = nil
}
