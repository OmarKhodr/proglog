// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/v1/log.proto

package log_v1

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

type Record struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []byte                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset        uint64                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Term          uint64                 `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Type          uint32                 `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_api_v1_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Record) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Record) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Record) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type ProduceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Record        *Record                `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProduceRequest) Reset() {
	*x = ProduceRequest{}
	mi := &file_api_v1_log_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProduceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceRequest) ProtoMessage() {}

func (x *ProduceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceRequest.ProtoReflect.Descriptor instead.
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{1}
}

func (x *ProduceRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type ProduceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        uint64                 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProduceResponse) Reset() {
	*x = ProduceResponse{}
	mi := &file_api_v1_log_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProduceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceResponse) ProtoMessage() {}

func (x *ProduceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceResponse.ProtoReflect.Descriptor instead.
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{2}
}

func (x *ProduceResponse) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ConsumeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        uint64                 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	mi := &file_api_v1_log_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRequest.ProtoReflect.Descriptor instead.
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{3}
}

func (x *ConsumeRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ConsumeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Record        *Record                `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsumeResponse) Reset() {
	*x = ConsumeResponse{}
	mi := &file_api_v1_log_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeResponse) ProtoMessage() {}

func (x *ConsumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeResponse.ProtoReflect.Descriptor instead.
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{4}
}

func (x *ConsumeResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

var File_api_v1_log_proto protoreflect.FileDescriptor

const file_api_v1_log_proto_rawDesc = "" +
	"\n" +
	"\x10api/v1/log.proto\x12\x06log.v1\"^\n" +
	"\x06Record\x12\x14\n" +
	"\x05value\x18\x01 \x01(\fR\x05value\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x04R\x06offset\x12\x12\n" +
	"\x04term\x18\x03 \x01(\x04R\x04term\x12\x12\n" +
	"\x04type\x18\x04 \x01(\rR\x04type\"8\n" +
	"\x0eProduceRequest\x12&\n" +
	"\x06record\x18\x01 \x01(\v2\x0e.log.v1.RecordR\x06record\")\n" +
	"\x0fProduceResponse\x12\x16\n" +
	"\x06offset\x18\x01 \x01(\x04R\x06offset\"(\n" +
	"\x0eConsumeRequest\x12\x16\n" +
	"\x06offset\x18\x01 \x01(\x04R\x06offset\"9\n" +
	"\x0fConsumeResponse\x12&\n" +
	"\x06record\x18\x01 \x01(\v2\x0e.log.v1.RecordR\x06record2\x8f\x02\n" +
	"\x03Log\x12<\n" +
	"\aProduce\x12\x16.log.v1.ProduceRequest\x1a\x17.log.v1.ProduceResponse\"\x00\x12<\n" +
	"\aConsume\x12\x16.log.v1.ConsumeRequest\x1a\x17.log.v1.ConsumeResponse\"\x00\x12D\n" +
	"\rConsumeStream\x12\x16.log.v1.ConsumeRequest\x1a\x17.log.v1.ConsumeResponse\"\x000\x01\x12F\n" +
	"\rProduceStream\x12\x16.log.v1.ProduceRequest\x1a\x17.log.v1.ProduceResponse\"\x00(\x010\x01B!Z\x1fgithub.com/omarkhodr/api/log_v1b\x06proto3"

var (
	file_api_v1_log_proto_rawDescOnce sync.Once
	file_api_v1_log_proto_rawDescData []byte
)

func file_api_v1_log_proto_rawDescGZIP() []byte {
	file_api_v1_log_proto_rawDescOnce.Do(func() {
		file_api_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_log_proto_rawDesc), len(file_api_v1_log_proto_rawDesc)))
	})
	return file_api_v1_log_proto_rawDescData
}

var file_api_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_log_proto_goTypes = []any{
	(*Record)(nil),          // 0: log.v1.Record
	(*ProduceRequest)(nil),  // 1: log.v1.ProduceRequest
	(*ProduceResponse)(nil), // 2: log.v1.ProduceResponse
	(*ConsumeRequest)(nil),  // 3: log.v1.ConsumeRequest
	(*ConsumeResponse)(nil), // 4: log.v1.ConsumeResponse
}
var file_api_v1_log_proto_depIdxs = []int32{
	0, // 0: log.v1.ProduceRequest.record:type_name -> log.v1.Record
	0, // 1: log.v1.ConsumeResponse.record:type_name -> log.v1.Record
	1, // 2: log.v1.Log.Produce:input_type -> log.v1.ProduceRequest
	3, // 3: log.v1.Log.Consume:input_type -> log.v1.ConsumeRequest
	3, // 4: log.v1.Log.ConsumeStream:input_type -> log.v1.ConsumeRequest
	1, // 5: log.v1.Log.ProduceStream:input_type -> log.v1.ProduceRequest
	2, // 6: log.v1.Log.Produce:output_type -> log.v1.ProduceResponse
	4, // 7: log.v1.Log.Consume:output_type -> log.v1.ConsumeResponse
	4, // 8: log.v1.Log.ConsumeStream:output_type -> log.v1.ConsumeResponse
	2, // 9: log.v1.Log.ProduceStream:output_type -> log.v1.ProduceResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_log_proto_init() }
func file_api_v1_log_proto_init() {
	if File_api_v1_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_log_proto_rawDesc), len(file_api_v1_log_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_log_proto_goTypes,
		DependencyIndexes: file_api_v1_log_proto_depIdxs,
		MessageInfos:      file_api_v1_log_proto_msgTypes,
	}.Build()
	File_api_v1_log_proto = out.File
	file_api_v1_log_proto_goTypes = nil
	file_api_v1_log_proto_depIdxs = nil
}
