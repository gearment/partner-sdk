// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/order/v1/data_vendor_fulfillment.proto

package orderv1

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

type VendorFulfillmentVendor int32

const (
	VendorFulfillmentVendor_VENDOR_FULFILLMENT_VENDOR_UNKNOWN  VendorFulfillmentVendor = 0
	VendorFulfillmentVendor_VENDOR_FULFILLMENT_VENDOR_ALL      VendorFulfillmentVendor = 1
	VendorFulfillmentVendor_VENDOR_FULFILLMENT_VENDOR_GEARMENT VendorFulfillmentVendor = 2
	VendorFulfillmentVendor_VENDOR_FULFILLMENT_VENDOR_AMAZON   VendorFulfillmentVendor = 3
)

// Enum value maps for VendorFulfillmentVendor.
var (
	VendorFulfillmentVendor_name = map[int32]string{
		0: "VENDOR_FULFILLMENT_VENDOR_UNKNOWN",
		1: "VENDOR_FULFILLMENT_VENDOR_ALL",
		2: "VENDOR_FULFILLMENT_VENDOR_GEARMENT",
		3: "VENDOR_FULFILLMENT_VENDOR_AMAZON",
	}
	VendorFulfillmentVendor_value = map[string]int32{
		"VENDOR_FULFILLMENT_VENDOR_UNKNOWN":  0,
		"VENDOR_FULFILLMENT_VENDOR_ALL":      1,
		"VENDOR_FULFILLMENT_VENDOR_GEARMENT": 2,
		"VENDOR_FULFILLMENT_VENDOR_AMAZON":   3,
	}
)

func (x VendorFulfillmentVendor) Enum() *VendorFulfillmentVendor {
	p := new(VendorFulfillmentVendor)
	*p = x
	return p
}

func (x VendorFulfillmentVendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorFulfillmentVendor) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_v1_data_vendor_fulfillment_proto_enumTypes[0].Descriptor()
}

func (VendorFulfillmentVendor) Type() protoreflect.EnumType {
	return &file_api_order_v1_data_vendor_fulfillment_proto_enumTypes[0]
}

func (x VendorFulfillmentVendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorFulfillmentVendor.Descriptor instead.
func (VendorFulfillmentVendor) EnumDescriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_fulfillment_proto_rawDescGZIP(), []int{0}
}

type VendorFulfillmentPriority int32

const (
	VendorFulfillmentPriority_VENDOR_FULFILLMENT_PRIORITY_UNKNOWN VendorFulfillmentPriority = 0
	VendorFulfillmentPriority_VENDOR_FULFILLMENT_PRIORITY_ALL     VendorFulfillmentPriority = 1 // only used for filtering
	VendorFulfillmentPriority_VENDOR_FULFILLMENT_PRIORITY_NORMAL  VendorFulfillmentPriority = 2
	VendorFulfillmentPriority_VENDOR_FULFILLMENT_PRIORITY_RUSH    VendorFulfillmentPriority = 3
)

// Enum value maps for VendorFulfillmentPriority.
var (
	VendorFulfillmentPriority_name = map[int32]string{
		0: "VENDOR_FULFILLMENT_PRIORITY_UNKNOWN",
		1: "VENDOR_FULFILLMENT_PRIORITY_ALL",
		2: "VENDOR_FULFILLMENT_PRIORITY_NORMAL",
		3: "VENDOR_FULFILLMENT_PRIORITY_RUSH",
	}
	VendorFulfillmentPriority_value = map[string]int32{
		"VENDOR_FULFILLMENT_PRIORITY_UNKNOWN": 0,
		"VENDOR_FULFILLMENT_PRIORITY_ALL":     1,
		"VENDOR_FULFILLMENT_PRIORITY_NORMAL":  2,
		"VENDOR_FULFILLMENT_PRIORITY_RUSH":    3,
	}
)

func (x VendorFulfillmentPriority) Enum() *VendorFulfillmentPriority {
	p := new(VendorFulfillmentPriority)
	*p = x
	return p
}

func (x VendorFulfillmentPriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorFulfillmentPriority) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_v1_data_vendor_fulfillment_proto_enumTypes[1].Descriptor()
}

func (VendorFulfillmentPriority) Type() protoreflect.EnumType {
	return &file_api_order_v1_data_vendor_fulfillment_proto_enumTypes[1]
}

func (x VendorFulfillmentPriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorFulfillmentPriority.Descriptor instead.
func (VendorFulfillmentPriority) EnumDescriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_fulfillment_proto_rawDescGZIP(), []int{1}
}

type VendorFulfillmentOption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VendorFulfillmentOption) Reset() {
	*x = VendorFulfillmentOption{}
	mi := &file_api_order_v1_data_vendor_fulfillment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VendorFulfillmentOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorFulfillmentOption) ProtoMessage() {}

func (x *VendorFulfillmentOption) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_data_vendor_fulfillment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorFulfillmentOption.ProtoReflect.Descriptor instead.
func (*VendorFulfillmentOption) Descriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_fulfillment_proto_rawDescGZIP(), []int{0}
}

var File_api_order_v1_data_vendor_fulfillment_proto protoreflect.FileDescriptor

var file_api_order_v1_data_vendor_fulfillment_proto_rawDesc = string([]byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x19, 0x0a, 0x17, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xb1, 0x01, 0x0a, 0x17, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x21, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x55, 0x4c, 0x46,
	0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x56, 0x45, 0x4e, 0x44,
	0x4f, 0x52, 0x5f, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x56,
	0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x56,
	0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x47, 0x45, 0x41, 0x52, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x55,
	0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52,
	0x5f, 0x41, 0x4d, 0x41, 0x5a, 0x4f, 0x4e, 0x10, 0x03, 0x2a, 0xb7, 0x01, 0x0a, 0x19, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x23, 0x56, 0x45, 0x4e, 0x44, 0x4f,
	0x52, 0x5f, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52,
	0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x23, 0x0a, 0x1f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x55, 0x4c, 0x46, 0x49,
	0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f,
	0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x4f,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x24, 0x0a,
	0x20, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x55, 0x53,
	0x48, 0x10, 0x03, 0x42, 0xbd, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x61, 0x2d,
	0x6e, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x64, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x4f, 0x58, 0xaa, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x5c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_order_v1_data_vendor_fulfillment_proto_rawDescOnce sync.Once
	file_api_order_v1_data_vendor_fulfillment_proto_rawDescData []byte
)

func file_api_order_v1_data_vendor_fulfillment_proto_rawDescGZIP() []byte {
	file_api_order_v1_data_vendor_fulfillment_proto_rawDescOnce.Do(func() {
		file_api_order_v1_data_vendor_fulfillment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_order_v1_data_vendor_fulfillment_proto_rawDesc), len(file_api_order_v1_data_vendor_fulfillment_proto_rawDesc)))
	})
	return file_api_order_v1_data_vendor_fulfillment_proto_rawDescData
}

var file_api_order_v1_data_vendor_fulfillment_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_order_v1_data_vendor_fulfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_order_v1_data_vendor_fulfillment_proto_goTypes = []any{
	(VendorFulfillmentVendor)(0),    // 0: api.order.v1.VendorFulfillmentVendor
	(VendorFulfillmentPriority)(0),  // 1: api.order.v1.VendorFulfillmentPriority
	(*VendorFulfillmentOption)(nil), // 2: api.order.v1.VendorFulfillmentOption
}
var file_api_order_v1_data_vendor_fulfillment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_order_v1_data_vendor_fulfillment_proto_init() }
func file_api_order_v1_data_vendor_fulfillment_proto_init() {
	if File_api_order_v1_data_vendor_fulfillment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_order_v1_data_vendor_fulfillment_proto_rawDesc), len(file_api_order_v1_data_vendor_fulfillment_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_order_v1_data_vendor_fulfillment_proto_goTypes,
		DependencyIndexes: file_api_order_v1_data_vendor_fulfillment_proto_depIdxs,
		EnumInfos:         file_api_order_v1_data_vendor_fulfillment_proto_enumTypes,
		MessageInfos:      file_api_order_v1_data_vendor_fulfillment_proto_msgTypes,
	}.Build()
	File_api_order_v1_data_vendor_fulfillment_proto = out.File
	file_api_order_v1_data_vendor_fulfillment_proto_goTypes = nil
	file_api_order_v1_data_vendor_fulfillment_proto_depIdxs = nil
}
