// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/type/v1/slice.proto

package typev1

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

type SliceString struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []string               `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SliceString) Reset() {
	*x = SliceString{}
	mi := &file_common_type_v1_slice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SliceString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceString) ProtoMessage() {}

func (x *SliceString) ProtoReflect() protoreflect.Message {
	mi := &file_common_type_v1_slice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceString.ProtoReflect.Descriptor instead.
func (*SliceString) Descriptor() ([]byte, []int) {
	return file_common_type_v1_slice_proto_rawDescGZIP(), []int{0}
}

func (x *SliceString) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

type SliceInt32 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []int32                `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SliceInt32) Reset() {
	*x = SliceInt32{}
	mi := &file_common_type_v1_slice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SliceInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceInt32) ProtoMessage() {}

func (x *SliceInt32) ProtoReflect() protoreflect.Message {
	mi := &file_common_type_v1_slice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceInt32.ProtoReflect.Descriptor instead.
func (*SliceInt32) Descriptor() ([]byte, []int) {
	return file_common_type_v1_slice_proto_rawDescGZIP(), []int{1}
}

func (x *SliceInt32) GetData() []int32 {
	if x != nil {
		return x.Data
	}
	return nil
}

type SliceInt64 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []int64                `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SliceInt64) Reset() {
	*x = SliceInt64{}
	mi := &file_common_type_v1_slice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SliceInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceInt64) ProtoMessage() {}

func (x *SliceInt64) ProtoReflect() protoreflect.Message {
	mi := &file_common_type_v1_slice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceInt64.ProtoReflect.Descriptor instead.
func (*SliceInt64) Descriptor() ([]byte, []int) {
	return file_common_type_v1_slice_proto_rawDescGZIP(), []int{2}
}

func (x *SliceInt64) GetData() []int64 {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_type_v1_slice_proto protoreflect.FileDescriptor

var file_common_type_v1_slice_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x21, 0x0a, 0x0b,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x20, 0x0a, 0x0a, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x20, 0x0a, 0x0a, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0xb8, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65,
	0x61, 0x2d, 0x6e, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x64, 0x6b,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x0e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_type_v1_slice_proto_rawDescOnce sync.Once
	file_common_type_v1_slice_proto_rawDescData []byte
)

func file_common_type_v1_slice_proto_rawDescGZIP() []byte {
	file_common_type_v1_slice_proto_rawDescOnce.Do(func() {
		file_common_type_v1_slice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_type_v1_slice_proto_rawDesc), len(file_common_type_v1_slice_proto_rawDesc)))
	})
	return file_common_type_v1_slice_proto_rawDescData
}

var file_common_type_v1_slice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_type_v1_slice_proto_goTypes = []any{
	(*SliceString)(nil), // 0: common.type.v1.SliceString
	(*SliceInt32)(nil),  // 1: common.type.v1.SliceInt32
	(*SliceInt64)(nil),  // 2: common.type.v1.SliceInt64
}
var file_common_type_v1_slice_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_type_v1_slice_proto_init() }
func file_common_type_v1_slice_proto_init() {
	if File_common_type_v1_slice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_type_v1_slice_proto_rawDesc), len(file_common_type_v1_slice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_type_v1_slice_proto_goTypes,
		DependencyIndexes: file_common_type_v1_slice_proto_depIdxs,
		MessageInfos:      file_common_type_v1_slice_proto_msgTypes,
	}.Build()
	File_common_type_v1_slice_proto = out.File
	file_common_type_v1_slice_proto_goTypes = nil
	file_common_type_v1_slice_proto_depIdxs = nil
}
