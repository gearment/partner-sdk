// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/order/v1/data_vendor_shipping.proto

package orderv1

import (
	_ "github.com/gearment/gea-next/vendorsdk/common/option/v1"
	v1 "github.com/gearment/gea-next/vendorsdk/common/type/v1"
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

type VendorShippingServiceType int32

const (
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_UNSPECIFIED VendorShippingServiceType = 0
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_ALL         VendorShippingServiceType = 1 // only used for filtering
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_EBAY        VendorShippingServiceType = 2
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_ETSY        VendorShippingServiceType = 3
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_POSHMARK    VendorShippingServiceType = 4
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_TIKTOK      VendorShippingServiceType = 5
	VendorShippingServiceType_VENDOR_SHIPPING_SERVICE_TYPE_FBA         VendorShippingServiceType = 6
)

// Enum value maps for VendorShippingServiceType.
var (
	VendorShippingServiceType_name = map[int32]string{
		0: "VENDOR_SHIPPING_SERVICE_TYPE_UNSPECIFIED",
		1: "VENDOR_SHIPPING_SERVICE_TYPE_ALL",
		2: "VENDOR_SHIPPING_SERVICE_TYPE_EBAY",
		3: "VENDOR_SHIPPING_SERVICE_TYPE_ETSY",
		4: "VENDOR_SHIPPING_SERVICE_TYPE_POSHMARK",
		5: "VENDOR_SHIPPING_SERVICE_TYPE_TIKTOK",
		6: "VENDOR_SHIPPING_SERVICE_TYPE_FBA",
	}
	VendorShippingServiceType_value = map[string]int32{
		"VENDOR_SHIPPING_SERVICE_TYPE_UNSPECIFIED": 0,
		"VENDOR_SHIPPING_SERVICE_TYPE_ALL":         1,
		"VENDOR_SHIPPING_SERVICE_TYPE_EBAY":        2,
		"VENDOR_SHIPPING_SERVICE_TYPE_ETSY":        3,
		"VENDOR_SHIPPING_SERVICE_TYPE_POSHMARK":    4,
		"VENDOR_SHIPPING_SERVICE_TYPE_TIKTOK":      5,
		"VENDOR_SHIPPING_SERVICE_TYPE_FBA":         6,
	}
)

func (x VendorShippingServiceType) Enum() *VendorShippingServiceType {
	p := new(VendorShippingServiceType)
	*p = x
	return p
}

func (x VendorShippingServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorShippingServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_v1_data_vendor_shipping_proto_enumTypes[0].Descriptor()
}

func (VendorShippingServiceType) Type() protoreflect.EnumType {
	return &file_api_order_v1_data_vendor_shipping_proto_enumTypes[0]
}

func (x VendorShippingServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorShippingServiceType.Descriptor instead.
func (VendorShippingServiceType) EnumDescriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_shipping_proto_rawDescGZIP(), []int{0}
}

type VendorShippingVerificationStatus int32

const (
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_UNKNOWN            VendorShippingVerificationStatus = 0
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_ALL                VendorShippingVerificationStatus = 1 // only used for filtering
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_PENDING            VendorShippingVerificationStatus = 2
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_CONSTRAINT_INVALID VendorShippingVerificationStatus = 3
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_CONSTRAINT_VALID   VendorShippingVerificationStatus = 4
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_ADDRESS_UNVERIFIED VendorShippingVerificationStatus = 5
	VendorShippingVerificationStatus_VENDOR_SHIPPING_VERIFICATION_STATUS_ADDRESS_VERIFIED   VendorShippingVerificationStatus = 6
)

// Enum value maps for VendorShippingVerificationStatus.
var (
	VendorShippingVerificationStatus_name = map[int32]string{
		0: "VENDOR_SHIPPING_VERIFICATION_STATUS_UNKNOWN",
		1: "VENDOR_SHIPPING_VERIFICATION_STATUS_ALL",
		2: "VENDOR_SHIPPING_VERIFICATION_STATUS_PENDING",
		3: "VENDOR_SHIPPING_VERIFICATION_STATUS_CONSTRAINT_INVALID",
		4: "VENDOR_SHIPPING_VERIFICATION_STATUS_CONSTRAINT_VALID",
		5: "VENDOR_SHIPPING_VERIFICATION_STATUS_ADDRESS_UNVERIFIED",
		6: "VENDOR_SHIPPING_VERIFICATION_STATUS_ADDRESS_VERIFIED",
	}
	VendorShippingVerificationStatus_value = map[string]int32{
		"VENDOR_SHIPPING_VERIFICATION_STATUS_UNKNOWN":            0,
		"VENDOR_SHIPPING_VERIFICATION_STATUS_ALL":                1,
		"VENDOR_SHIPPING_VERIFICATION_STATUS_PENDING":            2,
		"VENDOR_SHIPPING_VERIFICATION_STATUS_CONSTRAINT_INVALID": 3,
		"VENDOR_SHIPPING_VERIFICATION_STATUS_CONSTRAINT_VALID":   4,
		"VENDOR_SHIPPING_VERIFICATION_STATUS_ADDRESS_UNVERIFIED": 5,
		"VENDOR_SHIPPING_VERIFICATION_STATUS_ADDRESS_VERIFIED":   6,
	}
)

func (x VendorShippingVerificationStatus) Enum() *VendorShippingVerificationStatus {
	p := new(VendorShippingVerificationStatus)
	*p = x
	return p
}

func (x VendorShippingVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorShippingVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_v1_data_vendor_shipping_proto_enumTypes[1].Descriptor()
}

func (VendorShippingVerificationStatus) Type() protoreflect.EnumType {
	return &file_api_order_v1_data_vendor_shipping_proto_enumTypes[1]
}

func (x VendorShippingVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorShippingVerificationStatus.Descriptor instead.
func (VendorShippingVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_shipping_proto_rawDescGZIP(), []int{1}
}

type VendorShippingMethod int32

const (
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_UNKNOWN                VendorShippingMethod = 0
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_ALL                    VendorShippingMethod = 1 // only used for filtering
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_STANDARD               VendorShippingMethod = 2
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_GROUND                 VendorShippingMethod = 3
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_FAST_SHIP              VendorShippingMethod = 4
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_STAMP                  VendorShippingMethod = 5
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_INTERNATIONAL_STANDARD VendorShippingMethod = 6
	VendorShippingMethod_VENDOR_SHIPPING_METHOD_INTERNATIONAL_STAMP    VendorShippingMethod = 7
)

// Enum value maps for VendorShippingMethod.
var (
	VendorShippingMethod_name = map[int32]string{
		0: "VENDOR_SHIPPING_METHOD_UNKNOWN",
		1: "VENDOR_SHIPPING_METHOD_ALL",
		2: "VENDOR_SHIPPING_METHOD_STANDARD",
		3: "VENDOR_SHIPPING_METHOD_GROUND",
		4: "VENDOR_SHIPPING_METHOD_FAST_SHIP",
		5: "VENDOR_SHIPPING_METHOD_STAMP",
		6: "VENDOR_SHIPPING_METHOD_INTERNATIONAL_STANDARD",
		7: "VENDOR_SHIPPING_METHOD_INTERNATIONAL_STAMP",
	}
	VendorShippingMethod_value = map[string]int32{
		"VENDOR_SHIPPING_METHOD_UNKNOWN":                0,
		"VENDOR_SHIPPING_METHOD_ALL":                    1,
		"VENDOR_SHIPPING_METHOD_STANDARD":               2,
		"VENDOR_SHIPPING_METHOD_GROUND":                 3,
		"VENDOR_SHIPPING_METHOD_FAST_SHIP":              4,
		"VENDOR_SHIPPING_METHOD_STAMP":                  5,
		"VENDOR_SHIPPING_METHOD_INTERNATIONAL_STANDARD": 6,
		"VENDOR_SHIPPING_METHOD_INTERNATIONAL_STAMP":    7,
	}
)

func (x VendorShippingMethod) Enum() *VendorShippingMethod {
	p := new(VendorShippingMethod)
	*p = x
	return p
}

func (x VendorShippingMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorShippingMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_v1_data_vendor_shipping_proto_enumTypes[2].Descriptor()
}

func (VendorShippingMethod) Type() protoreflect.EnumType {
	return &file_api_order_v1_data_vendor_shipping_proto_enumTypes[2]
}

func (x VendorShippingMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorShippingMethod.Descriptor instead.
func (VendorShippingMethod) EnumDescriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_shipping_proto_rawDescGZIP(), []int{2}
}

type VendorShippingLabel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LabelFile     *v1.File               `protobuf:"bytes,1,opt,name=label_file,json=labelFile,proto3" json:"label_file,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VendorShippingLabel) Reset() {
	*x = VendorShippingLabel{}
	mi := &file_api_order_v1_data_vendor_shipping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VendorShippingLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorShippingLabel) ProtoMessage() {}

func (x *VendorShippingLabel) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_data_vendor_shipping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorShippingLabel.ProtoReflect.Descriptor instead.
func (*VendorShippingLabel) Descriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *VendorShippingLabel) GetLabelFile() *v1.File {
	if x != nil {
		return x.LabelFile
	}
	return nil
}

type VendorShippingOption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        VendorShippingMethod   `protobuf:"varint,1,opt,name=method,proto3,enum=api.order.v1.VendorShippingMethod" json:"method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VendorShippingOption) Reset() {
	*x = VendorShippingOption{}
	mi := &file_api_order_v1_data_vendor_shipping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VendorShippingOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorShippingOption) ProtoMessage() {}

func (x *VendorShippingOption) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_data_vendor_shipping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorShippingOption.ProtoReflect.Descriptor instead.
func (*VendorShippingOption) Descriptor() ([]byte, []int) {
	return file_api_order_v1_data_vendor_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *VendorShippingOption) GetMethod() VendorShippingMethod {
	if x != nil {
		return x.Method
	}
	return VendorShippingMethod_VENDOR_SHIPPING_METHOD_UNKNOWN
}

var File_api_order_v1_data_vendor_shipping_proto protoreflect.FileDescriptor

var file_api_order_v1_data_vendor_shipping_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x13, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2a,
	0xfd, 0x02, 0x0a, 0x19, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a,
	0x28, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa2, 0xc0, 0x02,
	0x00, 0x12, 0x2d, 0x0a, 0x20, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50,
	0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x1a, 0x07, 0xa2, 0xc0, 0x02, 0x03, 0x41, 0x6c, 0x6c,
	0x12, 0x2f, 0x0a, 0x21, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x42, 0x41, 0x59, 0x10, 0x02, 0x1a, 0x08, 0xa2, 0xc0, 0x02, 0x04, 0x45, 0x62, 0x61,
	0x79, 0x12, 0x2f, 0x0a, 0x21, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50,
	0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x54, 0x53, 0x59, 0x10, 0x03, 0x1a, 0x08, 0xa2, 0xc0, 0x02, 0x04, 0x45, 0x74,
	0x73, 0x79, 0x12, 0x37, 0x0a, 0x25, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49,
	0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x48, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x04, 0x1a, 0x0c, 0xa2,
	0xc0, 0x02, 0x08, 0x50, 0x6f, 0x73, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x33, 0x0a, 0x23, 0x56,
	0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4b, 0x54,
	0x4f, 0x4b, 0x10, 0x05, 0x1a, 0x0a, 0xa2, 0xc0, 0x02, 0x06, 0x54, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x12, 0x2d, 0x0a, 0x20, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x42, 0x41, 0x10, 0x06, 0x1a, 0x07, 0xa2, 0xc0, 0x02, 0x03, 0x46, 0x42, 0x41, 0x2a,
	0x95, 0x04, 0x0a, 0x20, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x2b, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53,
	0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa2, 0xc0, 0x02, 0x00, 0x12, 0x34, 0x0a, 0x27, 0x56,
	0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x1a, 0x07, 0xa2, 0xc0, 0x02, 0x03, 0x41, 0x6c,
	0x6c, 0x12, 0x3c, 0x0a, 0x2b, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50,
	0x50, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x1a, 0x0b, 0xa2, 0xc0, 0x02, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x52, 0x0a, 0x36, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49,
	0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e,
	0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x1a, 0x16, 0xa2, 0xc0, 0x02,
	0x12, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x34, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48,
	0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x1a, 0x14, 0xa2,
	0xc0, 0x02, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x52, 0x0a, 0x36, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48,
	0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45,
	0x53, 0x53, 0x5f, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x05, 0x1a,
	0x16, 0xa2, 0xc0, 0x02, 0x12, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x20, 0x55, 0x6e, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x34, 0x56, 0x45, 0x4e, 0x44, 0x4f,
	0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41,
	0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x06, 0x1a, 0x14, 0xa2, 0xc0, 0x02, 0x10, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x20, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2a, 0xc4, 0x03, 0x0a, 0x14, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x28, 0x0a, 0x1e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa2, 0xc0, 0x02, 0x00, 0x12, 0x27, 0x0a, 0x1a, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x1a, 0x07, 0xa2, 0xc0, 0x02, 0x03,
	0x41, 0x6c, 0x6c, 0x12, 0x31, 0x0a, 0x1f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48,
	0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x02, 0x1a, 0x0c, 0xa2, 0xc0, 0x02, 0x08, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x1d, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52,
	0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x1a, 0x0a, 0xa2, 0xc0, 0x02, 0x06, 0x47,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x20, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f,
	0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x46, 0x41, 0x53, 0x54, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x10, 0x04, 0x1a, 0x0c, 0xa2, 0xc0, 0x02,
	0x08, 0x46, 0x61, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x12, 0x2b, 0x0a, 0x1c, 0x56, 0x45, 0x4e,
	0x44, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x05, 0x1a, 0x09, 0xa2, 0xc0, 0x02,
	0x05, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4d, 0x0a, 0x2d, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52,
	0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x06, 0x1a, 0x1a, 0xa2, 0xc0, 0x02, 0x16, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x47, 0x0a, 0x2a, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f,
	0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x4d, 0x50, 0x10, 0x07, 0x1a, 0x17, 0xa2, 0xc0, 0x02, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0xba,
	0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x17, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x61, 0x2d, 0x6e, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x4f,
	0x58, 0xaa, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x18, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_api_order_v1_data_vendor_shipping_proto_rawDescOnce sync.Once
	file_api_order_v1_data_vendor_shipping_proto_rawDescData []byte
)

func file_api_order_v1_data_vendor_shipping_proto_rawDescGZIP() []byte {
	file_api_order_v1_data_vendor_shipping_proto_rawDescOnce.Do(func() {
		file_api_order_v1_data_vendor_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_order_v1_data_vendor_shipping_proto_rawDesc), len(file_api_order_v1_data_vendor_shipping_proto_rawDesc)))
	})
	return file_api_order_v1_data_vendor_shipping_proto_rawDescData
}

var file_api_order_v1_data_vendor_shipping_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_order_v1_data_vendor_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_order_v1_data_vendor_shipping_proto_goTypes = []any{
	(VendorShippingServiceType)(0),        // 0: api.order.v1.VendorShippingServiceType
	(VendorShippingVerificationStatus)(0), // 1: api.order.v1.VendorShippingVerificationStatus
	(VendorShippingMethod)(0),             // 2: api.order.v1.VendorShippingMethod
	(*VendorShippingLabel)(nil),           // 3: api.order.v1.VendorShippingLabel
	(*VendorShippingOption)(nil),          // 4: api.order.v1.VendorShippingOption
	(*v1.File)(nil),                       // 5: common.type.v1.File
}
var file_api_order_v1_data_vendor_shipping_proto_depIdxs = []int32{
	5, // 0: api.order.v1.VendorShippingLabel.label_file:type_name -> common.type.v1.File
	2, // 1: api.order.v1.VendorShippingOption.method:type_name -> api.order.v1.VendorShippingMethod
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_order_v1_data_vendor_shipping_proto_init() }
func file_api_order_v1_data_vendor_shipping_proto_init() {
	if File_api_order_v1_data_vendor_shipping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_order_v1_data_vendor_shipping_proto_rawDesc), len(file_api_order_v1_data_vendor_shipping_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_order_v1_data_vendor_shipping_proto_goTypes,
		DependencyIndexes: file_api_order_v1_data_vendor_shipping_proto_depIdxs,
		EnumInfos:         file_api_order_v1_data_vendor_shipping_proto_enumTypes,
		MessageInfos:      file_api_order_v1_data_vendor_shipping_proto_msgTypes,
	}.Build()
	File_api_order_v1_data_vendor_shipping_proto = out.File
	file_api_order_v1_data_vendor_shipping_proto_goTypes = nil
	file_api_order_v1_data_vendor_shipping_proto_depIdxs = nil
}
