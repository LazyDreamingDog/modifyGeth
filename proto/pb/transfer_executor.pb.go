// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.0
// source: transfer_executor.proto

package pb

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

type CommitWithdrawTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxData     []byte `protobuf:"bytes,1,opt,name=TxData,proto3" json:"TxData,omitempty"`
	To         []byte `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	Value      uint64 `protobuf:"varint,3,opt,name=Value,proto3" json:"Value,omitempty"`
	VerifyHash []byte `protobuf:"bytes,4,opt,name=VerifyHash,proto3" json:"VerifyHash,omitempty"`
	Height     uint64 `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
}

func (x *CommitWithdrawTxRequest) Reset() {
	*x = CommitWithdrawTxRequest{}
	mi := &file_transfer_executor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommitWithdrawTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitWithdrawTxRequest) ProtoMessage() {}

func (x *CommitWithdrawTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_executor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitWithdrawTxRequest.ProtoReflect.Descriptor instead.
func (*CommitWithdrawTxRequest) Descriptor() ([]byte, []int) {
	return file_transfer_executor_proto_rawDescGZIP(), []int{0}
}

func (x *CommitWithdrawTxRequest) GetTxData() []byte {
	if x != nil {
		return x.TxData
	}
	return nil
}

func (x *CommitWithdrawTxRequest) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *CommitWithdrawTxRequest) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CommitWithdrawTxRequest) GetVerifyHash() []byte {
	if x != nil {
		return x.VerifyHash
	}
	return nil
}

func (x *CommitWithdrawTxRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type VerifyChangeTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      uint64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Height     uint64 `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	VerifyHash []byte `protobuf:"bytes,3,opt,name=VerifyHash,proto3" json:"VerifyHash,omitempty"`
}

func (x *VerifyChangeTxRequest) Reset() {
	*x = VerifyChangeTxRequest{}
	mi := &file_transfer_executor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyChangeTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyChangeTxRequest) ProtoMessage() {}

func (x *VerifyChangeTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_executor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyChangeTxRequest.ProtoReflect.Descriptor instead.
func (*VerifyChangeTxRequest) Descriptor() ([]byte, []int) {
	return file_transfer_executor_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyChangeTxRequest) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *VerifyChangeTxRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VerifyChangeTxRequest) GetVerifyHash() []byte {
	if x != nil {
		return x.VerifyHash
	}
	return nil
}

type VerifyChangeTxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyRes bool `protobuf:"varint,1,opt,name=VerifyRes,proto3" json:"VerifyRes,omitempty"`
}

func (x *VerifyChangeTxReply) Reset() {
	*x = VerifyChangeTxReply{}
	mi := &file_transfer_executor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyChangeTxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyChangeTxReply) ProtoMessage() {}

func (x *VerifyChangeTxReply) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_executor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyChangeTxReply.ProtoReflect.Descriptor instead.
func (*VerifyChangeTxReply) Descriptor() ([]byte, []int) {
	return file_transfer_executor_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyChangeTxReply) GetVerifyRes() bool {
	if x != nil {
		return x.VerifyRes
	}
	return false
}

var File_transfer_executor_proto protoreflect.FileDescriptor

var file_transfer_executor_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x17,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x78, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x54, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x65, 0x0a,
	0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x48, 0x61, 0x73, 0x68, 0x22, 0x33, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x32, 0x9c, 0x01, 0x0a, 0x14, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x3c, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x54, 0x78, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x78, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_executor_proto_rawDescOnce sync.Once
	file_transfer_executor_proto_rawDescData = file_transfer_executor_proto_rawDesc
)

func file_transfer_executor_proto_rawDescGZIP() []byte {
	file_transfer_executor_proto_rawDescOnce.Do(func() {
		file_transfer_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_executor_proto_rawDescData)
	})
	return file_transfer_executor_proto_rawDescData
}

var file_transfer_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transfer_executor_proto_goTypes = []any{
	(*CommitWithdrawTxRequest)(nil), // 0: pb.CommitWithdrawTxRequest
	(*VerifyChangeTxRequest)(nil),   // 1: pb.VerifyChangeTxRequest
	(*VerifyChangeTxReply)(nil),     // 2: pb.VerifyChangeTxReply
	(*Empty)(nil),                   // 3: pb.Empty
}
var file_transfer_executor_proto_depIdxs = []int32{
	0, // 0: pb.TransferExecutorGRPC.CommitWithdrawTx:input_type -> pb.CommitWithdrawTxRequest
	1, // 1: pb.TransferExecutorGRPC.VerifyChangeTx:input_type -> pb.VerifyChangeTxRequest
	3, // 2: pb.TransferExecutorGRPC.CommitWithdrawTx:output_type -> pb.Empty
	2, // 3: pb.TransferExecutorGRPC.VerifyChangeTx:output_type -> pb.VerifyChangeTxReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_executor_proto_init() }
func file_transfer_executor_proto_init() {
	if File_transfer_executor_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_executor_proto_goTypes,
		DependencyIndexes: file_transfer_executor_proto_depIdxs,
		MessageInfos:      file_transfer_executor_proto_msgTypes,
	}.Build()
	File_transfer_executor_proto = out.File
	file_transfer_executor_proto_rawDesc = nil
	file_transfer_executor_proto_goTypes = nil
	file_transfer_executor_proto_depIdxs = nil
}