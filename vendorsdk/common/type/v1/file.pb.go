// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/type/v1/file.proto

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

type File struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath      string                 `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileUrl       string                 `protobuf:"bytes,3,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	FileId        string                 `protobuf:"bytes,4,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_common_type_v1_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_common_type_v1_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_common_type_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *File) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *File) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *File) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

var File_common_type_v1_file_proto protoreflect.FileDescriptor

var file_common_type_v1_file_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0xe8, 0x07, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x72, 0x03, 0x18, 0xe8, 0x07, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x23, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0xe8, 0x07, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0x80, 0x10, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x42, 0xb7, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x67, 0x65, 0x61, 0x2d, 0x6e, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa,
	0x02, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_type_v1_file_proto_rawDescOnce sync.Once
	file_common_type_v1_file_proto_rawDescData []byte
)

func file_common_type_v1_file_proto_rawDescGZIP() []byte {
	file_common_type_v1_file_proto_rawDescOnce.Do(func() {
		file_common_type_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_type_v1_file_proto_rawDesc), len(file_common_type_v1_file_proto_rawDesc)))
	})
	return file_common_type_v1_file_proto_rawDescData
}

var file_common_type_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_type_v1_file_proto_goTypes = []any{
	(*File)(nil), // 0: common.type.v1.File
}
var file_common_type_v1_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_type_v1_file_proto_init() }
func file_common_type_v1_file_proto_init() {
	if File_common_type_v1_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_type_v1_file_proto_rawDesc), len(file_common_type_v1_file_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_type_v1_file_proto_goTypes,
		DependencyIndexes: file_common_type_v1_file_proto_depIdxs,
		MessageInfos:      file_common_type_v1_file_proto_msgTypes,
	}.Build()
	File_common_type_v1_file_proto = out.File
	file_common_type_v1_file_proto_goTypes = nil
	file_common_type_v1_file_proto_depIdxs = nil
}
