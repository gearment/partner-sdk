// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/option/v1/enum.proto

package optionv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_common_option_v1_enum_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         5124,
		Name:          "common.option.v1.display_name",
		Tag:           "bytes,5124,opt,name=display_name",
		Filename:      "common/option/v1/enum.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string display_name = 5124;
	E_DisplayName = &file_common_option_v1_enum_proto_extTypes[0]
)

var File_common_option_v1_enum_proto protoreflect.FileDescriptor

var file_common_option_v1_enum_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3a, 0x48, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0xc5, 0x01, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65,
	0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x61, 0x2d, 0x6e, 0x65, 0x78, 0x74, 0x2f,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4f, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_common_option_v1_enum_proto_goTypes = []any{
	(*descriptorpb.EnumValueOptions)(nil), // 0: google.protobuf.EnumValueOptions
}
var file_common_option_v1_enum_proto_depIdxs = []int32{
	0, // 0: common.option.v1.display_name:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_option_v1_enum_proto_init() }
func file_common_option_v1_enum_proto_init() {
	if File_common_option_v1_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_option_v1_enum_proto_rawDesc), len(file_common_option_v1_enum_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_common_option_v1_enum_proto_goTypes,
		DependencyIndexes: file_common_option_v1_enum_proto_depIdxs,
		ExtensionInfos:    file_common_option_v1_enum_proto_extTypes,
	}.Build()
	File_common_option_v1_enum_proto = out.File
	file_common_option_v1_enum_proto_goTypes = nil
	file_common_option_v1_enum_proto_depIdxs = nil
}
