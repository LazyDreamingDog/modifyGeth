// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.0
// source: pow.proto

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

type PoWBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash   []byte   `protobuf:"bytes,1,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	Height       uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Txs          [][]byte `protobuf:"bytes,3,rep,name=txs,proto3" json:"txs,omitempty"`
	Nonce        uint64   `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	BlockHash    []byte   `protobuf:"bytes,5,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Commited     bool     `protobuf:"varint,6,opt,name=commited,proto3" json:"commited,omitempty"`
	ShardingName []byte   `protobuf:"bytes,7,opt,name=shardingName,proto3" json:"shardingName,omitempty"`
	Incentive    []byte   `protobuf:"bytes,8,opt,name=incentive,proto3" json:"incentive,omitempty"`
}

func (x *PoWBlock) Reset() {
	*x = PoWBlock{}
	mi := &file_pow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoWBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoWBlock) ProtoMessage() {}

func (x *PoWBlock) ProtoReflect() protoreflect.Message {
	mi := &file_pow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoWBlock.ProtoReflect.Descriptor instead.
func (*PoWBlock) Descriptor() ([]byte, []int) {
	return file_pow_proto_rawDescGZIP(), []int{0}
}

func (x *PoWBlock) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *PoWBlock) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *PoWBlock) GetTxs() [][]byte {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *PoWBlock) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *PoWBlock) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *PoWBlock) GetCommited() bool {
	if x != nil {
		return x.Commited
	}
	return false
}

func (x *PoWBlock) GetShardingName() []byte {
	if x != nil {
		return x.ShardingName
	}
	return nil
}

func (x *PoWBlock) GetIncentive() []byte {
	if x != nil {
		return x.Incentive
	}
	return nil
}

type PoWMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*PoWMessage_Block
	//	*PoWMessage_Request
	Msg isPoWMessage_Msg `protobuf_oneof:"msg"`
}

func (x *PoWMessage) Reset() {
	*x = PoWMessage{}
	mi := &file_pow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoWMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoWMessage) ProtoMessage() {}

func (x *PoWMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoWMessage.ProtoReflect.Descriptor instead.
func (*PoWMessage) Descriptor() ([]byte, []int) {
	return file_pow_proto_rawDescGZIP(), []int{1}
}

func (m *PoWMessage) GetMsg() isPoWMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *PoWMessage) GetBlock() *PoWBlock {
	if x, ok := x.GetMsg().(*PoWMessage_Block); ok {
		return x.Block
	}
	return nil
}

func (x *PoWMessage) GetRequest() *Request {
	if x, ok := x.GetMsg().(*PoWMessage_Request); ok {
		return x.Request
	}
	return nil
}

type isPoWMessage_Msg interface {
	isPoWMessage_Msg()
}

type PoWMessage_Block struct {
	Block *PoWBlock `protobuf:"bytes,1,opt,name=block,proto3,oneof"`
}

type PoWMessage_Request struct {
	Request *Request `protobuf:"bytes,2,opt,name=request,proto3,oneof"`
}

func (*PoWMessage_Block) isPoWMessage_Msg() {}

func (*PoWMessage_Request) isPoWMessage_Msg() {}

var File_pow_proto protoreflect.FileDescriptor

var file_pow_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01,
	0x0a, 0x08, 0x50, 0x6f, 0x57, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x03, 0x74, 0x78, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x63,
	0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x22, 0x62, 0x0a, 0x0a, 0x50, 0x6f, 0x57, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x57, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pow_proto_rawDescOnce sync.Once
	file_pow_proto_rawDescData = file_pow_proto_rawDesc
)

func file_pow_proto_rawDescGZIP() []byte {
	file_pow_proto_rawDescOnce.Do(func() {
		file_pow_proto_rawDescData = protoimpl.X.CompressGZIP(file_pow_proto_rawDescData)
	})
	return file_pow_proto_rawDescData
}

var file_pow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pow_proto_goTypes = []any{
	(*PoWBlock)(nil),   // 0: pb.PoWBlock
	(*PoWMessage)(nil), // 1: pb.PoWMessage
	(*Request)(nil),    // 2: pb.Request
}
var file_pow_proto_depIdxs = []int32{
	0, // 0: pb.PoWMessage.block:type_name -> pb.PoWBlock
	2, // 1: pb.PoWMessage.request:type_name -> pb.Request
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pow_proto_init() }
func file_pow_proto_init() {
	if File_pow_proto != nil {
		return
	}
	file_common_proto_init()
	file_pow_proto_msgTypes[1].OneofWrappers = []any{
		(*PoWMessage_Block)(nil),
		(*PoWMessage_Request)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pow_proto_goTypes,
		DependencyIndexes: file_pow_proto_depIdxs,
		MessageInfos:      file_pow_proto_msgTypes,
	}.Build()
	File_pow_proto = out.File
	file_pow_proto_rawDesc = nil
	file_pow_proto_goTypes = nil
	file_pow_proto_depIdxs = nil
}
