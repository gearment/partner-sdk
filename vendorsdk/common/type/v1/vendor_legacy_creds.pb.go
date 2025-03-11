// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/type/v1/vendor_legacy_creds.proto

package typev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type LegacyCredentials struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiKey        string                 `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSignature  string                 `protobuf:"bytes,2,opt,name=api_signature,json=apiSignature,proto3" json:"api_signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LegacyCredentials) Reset() {
	*x = LegacyCredentials{}
	mi := &file_common_type_v1_vendor_legacy_creds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LegacyCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyCredentials) ProtoMessage() {}

func (x *LegacyCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_common_type_v1_vendor_legacy_creds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyCredentials.ProtoReflect.Descriptor instead.
func (*LegacyCredentials) Descriptor() ([]byte, []int) {
	return file_common_type_v1_vendor_legacy_creds_proto_rawDescGZIP(), []int{0}
}

func (x *LegacyCredentials) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *LegacyCredentials) GetApiSignature() string {
	if x != nil {
		return x.ApiSignature
	}
	return ""
}

var File_common_type_v1_vendor_legacy_creds_proto protoreflect.FileDescriptor

var file_common_type_v1_vendor_legacy_creds_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x11, 0x4c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x07,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba,
	0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x12, 0x2f, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10,
	0x01, 0x18, 0xff, 0x01, 0x52, 0x0c, 0x61, 0x70, 0x69, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0xc4, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x72, 0x65, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x61, 0x2d, 0x6e, 0x65, 0x78,
	0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a,
	0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_common_type_v1_vendor_legacy_creds_proto_rawDescOnce sync.Once
	file_common_type_v1_vendor_legacy_creds_proto_rawDescData []byte
)

func file_common_type_v1_vendor_legacy_creds_proto_rawDescGZIP() []byte {
	file_common_type_v1_vendor_legacy_creds_proto_rawDescOnce.Do(func() {
		file_common_type_v1_vendor_legacy_creds_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_type_v1_vendor_legacy_creds_proto_rawDesc), len(file_common_type_v1_vendor_legacy_creds_proto_rawDesc)))
	})
	return file_common_type_v1_vendor_legacy_creds_proto_rawDescData
}

var file_common_type_v1_vendor_legacy_creds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_type_v1_vendor_legacy_creds_proto_goTypes = []any{
	(*LegacyCredentials)(nil), // 0: common.type.v1.LegacyCredentials
}
var file_common_type_v1_vendor_legacy_creds_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_type_v1_vendor_legacy_creds_proto_init() }
func file_common_type_v1_vendor_legacy_creds_proto_init() {
	if File_common_type_v1_vendor_legacy_creds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_type_v1_vendor_legacy_creds_proto_rawDesc), len(file_common_type_v1_vendor_legacy_creds_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_type_v1_vendor_legacy_creds_proto_goTypes,
		DependencyIndexes: file_common_type_v1_vendor_legacy_creds_proto_depIdxs,
		MessageInfos:      file_common_type_v1_vendor_legacy_creds_proto_msgTypes,
	}.Build()
	File_common_type_v1_vendor_legacy_creds_proto = out.File
	file_common_type_v1_vendor_legacy_creds_proto_goTypes = nil
	file_common_type_v1_vendor_legacy_creds_proto_depIdxs = nil
}
