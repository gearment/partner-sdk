// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: common/platform/v1/platform.proto

package platformv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/gearment/gea-next/vendorsdk/common/option/v1"
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

type App int32

const (
	App_APP_UNSPECIFIED App = 0
	App_APP_POD         App = 2
	App_APP_FFM         App = 3
)

// Enum value maps for App.
var (
	App_name = map[int32]string{
		0: "APP_UNSPECIFIED",
		2: "APP_POD",
		3: "APP_FFM",
	}
	App_value = map[string]int32{
		"APP_UNSPECIFIED": 0,
		"APP_POD":         2,
		"APP_FFM":         3,
	}
)

func (x App) Enum() *App {
	p := new(App)
	*p = x
	return p
}

func (x App) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (App) Descriptor() protoreflect.EnumDescriptor {
	return file_common_platform_v1_platform_proto_enumTypes[0].Descriptor()
}

func (App) Type() protoreflect.EnumType {
	return &file_common_platform_v1_platform_proto_enumTypes[0]
}

func (x App) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use App.Descriptor instead.
func (App) EnumDescriptor() ([]byte, []int) {
	return file_common_platform_v1_platform_proto_rawDescGZIP(), []int{0}
}

type MarketplacePlatform int32

const (
	MarketplacePlatform_MARKETPLACE_PLATFORM_UNSPECIFIED MarketplacePlatform = 0
	MarketplacePlatform_MARKETPLACE_PLATFORM_ALL         MarketplacePlatform = 1
	MarketplacePlatform_MARKETPLACE_PLATFORM_EBAY        MarketplacePlatform = 2
	MarketplacePlatform_MARKETPLACE_PLATFORM_AMAZON      MarketplacePlatform = 3
	MarketplacePlatform_MARKETPLACE_PLATFORM_SHOPIFY     MarketplacePlatform = 4
	MarketplacePlatform_MARKETPLACE_PLATFORM_WOOCOMMERCE MarketplacePlatform = 5
	MarketplacePlatform_MARKETPLACE_PLATFORM_ETSY        MarketplacePlatform = 6
	MarketplacePlatform_MARKETPLACE_PLATFORM_SHOPBASE    MarketplacePlatform = 7
	MarketplacePlatform_MARKETPLACE_PLATFORM_GEARMENT    MarketplacePlatform = 8
	MarketplacePlatform_MARKETPLACE_PLATFORM_ORDERDESK   MarketplacePlatform = 9
	MarketplacePlatform_MARKETPLACE_PLATFORM_TIKTOKSHOP  MarketplacePlatform = 10
	MarketplacePlatform_MARKETPLACE_PLATFORM_POSHMARK    MarketplacePlatform = 11
	MarketplacePlatform_MARKETPLACE_PLATFORM_PRESTASHOP  MarketplacePlatform = 12
	MarketplacePlatform_MARKETPLACE_PLATFORM_INKGO       MarketplacePlatform = 13
	MarketplacePlatform_MARKETPLACE_PLATFORM_WISH        MarketplacePlatform = 14
	MarketplacePlatform_MARKETPLACE_PLATFORM_BIGCOMMERCE MarketplacePlatform = 15
)

// Enum value maps for MarketplacePlatform.
var (
	MarketplacePlatform_name = map[int32]string{
		0:  "MARKETPLACE_PLATFORM_UNSPECIFIED",
		1:  "MARKETPLACE_PLATFORM_ALL",
		2:  "MARKETPLACE_PLATFORM_EBAY",
		3:  "MARKETPLACE_PLATFORM_AMAZON",
		4:  "MARKETPLACE_PLATFORM_SHOPIFY",
		5:  "MARKETPLACE_PLATFORM_WOOCOMMERCE",
		6:  "MARKETPLACE_PLATFORM_ETSY",
		7:  "MARKETPLACE_PLATFORM_SHOPBASE",
		8:  "MARKETPLACE_PLATFORM_GEARMENT",
		9:  "MARKETPLACE_PLATFORM_ORDERDESK",
		10: "MARKETPLACE_PLATFORM_TIKTOKSHOP",
		11: "MARKETPLACE_PLATFORM_POSHMARK",
		12: "MARKETPLACE_PLATFORM_PRESTASHOP",
		13: "MARKETPLACE_PLATFORM_INKGO",
		14: "MARKETPLACE_PLATFORM_WISH",
		15: "MARKETPLACE_PLATFORM_BIGCOMMERCE",
	}
	MarketplacePlatform_value = map[string]int32{
		"MARKETPLACE_PLATFORM_UNSPECIFIED": 0,
		"MARKETPLACE_PLATFORM_ALL":         1,
		"MARKETPLACE_PLATFORM_EBAY":        2,
		"MARKETPLACE_PLATFORM_AMAZON":      3,
		"MARKETPLACE_PLATFORM_SHOPIFY":     4,
		"MARKETPLACE_PLATFORM_WOOCOMMERCE": 5,
		"MARKETPLACE_PLATFORM_ETSY":        6,
		"MARKETPLACE_PLATFORM_SHOPBASE":    7,
		"MARKETPLACE_PLATFORM_GEARMENT":    8,
		"MARKETPLACE_PLATFORM_ORDERDESK":   9,
		"MARKETPLACE_PLATFORM_TIKTOKSHOP":  10,
		"MARKETPLACE_PLATFORM_POSHMARK":    11,
		"MARKETPLACE_PLATFORM_PRESTASHOP":  12,
		"MARKETPLACE_PLATFORM_INKGO":       13,
		"MARKETPLACE_PLATFORM_WISH":        14,
		"MARKETPLACE_PLATFORM_BIGCOMMERCE": 15,
	}
)

func (x MarketplacePlatform) Enum() *MarketplacePlatform {
	p := new(MarketplacePlatform)
	*p = x
	return p
}

func (x MarketplacePlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MarketplacePlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_common_platform_v1_platform_proto_enumTypes[1].Descriptor()
}

func (MarketplacePlatform) Type() protoreflect.EnumType {
	return &file_common_platform_v1_platform_proto_enumTypes[1]
}

func (x MarketplacePlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MarketplacePlatform.Descriptor instead.
func (MarketplacePlatform) EnumDescriptor() ([]byte, []int) {
	return file_common_platform_v1_platform_proto_rawDescGZIP(), []int{1}
}

type ActivityCallback struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ActivityType  string                 `protobuf:"bytes,1,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	TaskQueue     string                 `protobuf:"bytes,2,opt,name=task_queue,json=taskQueue,proto3" json:"task_queue,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivityCallback) Reset() {
	*x = ActivityCallback{}
	mi := &file_common_platform_v1_platform_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivityCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityCallback) ProtoMessage() {}

func (x *ActivityCallback) ProtoReflect() protoreflect.Message {
	mi := &file_common_platform_v1_platform_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityCallback.ProtoReflect.Descriptor instead.
func (*ActivityCallback) Descriptor() ([]byte, []int) {
	return file_common_platform_v1_platform_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityCallback) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *ActivityCallback) GetTaskQueue() string {
	if x != nil {
		return x.TaskQueue
	}
	return ""
}

// utils
type MarketplacePlatformAndStringTuple struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Platform      MarketplacePlatform    `protobuf:"varint,1,opt,name=platform,proto3,enum=common.platform.v1.MarketplacePlatform" json:"platform,omitempty"`
	Data          string                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarketplacePlatformAndStringTuple) Reset() {
	*x = MarketplacePlatformAndStringTuple{}
	mi := &file_common_platform_v1_platform_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketplacePlatformAndStringTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketplacePlatformAndStringTuple) ProtoMessage() {}

func (x *MarketplacePlatformAndStringTuple) ProtoReflect() protoreflect.Message {
	mi := &file_common_platform_v1_platform_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketplacePlatformAndStringTuple.ProtoReflect.Descriptor instead.
func (*MarketplacePlatformAndStringTuple) Descriptor() ([]byte, []int) {
	return file_common_platform_v1_platform_proto_rawDescGZIP(), []int{1}
}

func (x *MarketplacePlatformAndStringTuple) GetPlatform() MarketplacePlatform {
	if x != nil {
		return x.Platform
	}
	return MarketplacePlatform_MARKETPLACE_PLATFORM_UNSPECIFIED
}

func (x *MarketplacePlatformAndStringTuple) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_common_platform_v1_platform_proto protoreflect.FileDescriptor

var file_common_platform_v1_platform_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x56, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x21, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12,
	0x4d, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x34, 0x0a, 0x03, 0x41,
	0x70, 0x70, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x50, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x50, 0x50, 0x5f, 0x50,
	0x4f, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x50, 0x50, 0x5f, 0x46, 0x46, 0x4d, 0x10,
	0x03, 0x2a, 0x88, 0x06, 0x0a, 0x13, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2a, 0x0a, 0x20, 0x4d, 0x41, 0x52,
	0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a,
	0x04, 0xa2, 0xc0, 0x02, 0x00, 0x12, 0x25, 0x0a, 0x18, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4c,
	0x4c, 0x10, 0x01, 0x1a, 0x07, 0xa2, 0xc0, 0x02, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x27, 0x0a, 0x19,
	0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x45, 0x42, 0x41, 0x59, 0x10, 0x02, 0x1a, 0x08, 0xa2, 0xc0, 0x02,
	0x04, 0x45, 0x62, 0x61, 0x79, 0x12, 0x2b, 0x0a, 0x1b, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4d,
	0x41, 0x5a, 0x4f, 0x4e, 0x10, 0x03, 0x1a, 0x0a, 0xa2, 0xc0, 0x02, 0x06, 0x41, 0x6d, 0x61, 0x7a,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x1c, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x49,
	0x46, 0x59, 0x10, 0x04, 0x1a, 0x0b, 0xa2, 0xc0, 0x02, 0x07, 0x53, 0x68, 0x6f, 0x70, 0x69, 0x66,
	0x79, 0x12, 0x35, 0x0a, 0x20, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x4f, 0x4f, 0x43, 0x4f, 0x4d,
	0x4d, 0x45, 0x52, 0x43, 0x45, 0x10, 0x05, 0x1a, 0x0f, 0xa2, 0xc0, 0x02, 0x0b, 0x57, 0x6f, 0x6f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x19, 0x4d, 0x41, 0x52, 0x4b,
	0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x45, 0x54, 0x53, 0x59, 0x10, 0x06, 0x1a, 0x08, 0xa2, 0xc0, 0x02, 0x04, 0x45, 0x74, 0x73,
	0x79, 0x12, 0x2f, 0x0a, 0x1d, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x42, 0x41,
	0x53, 0x45, 0x10, 0x07, 0x1a, 0x0c, 0xa2, 0xc0, 0x02, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x1d, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x47, 0x45, 0x41, 0x52, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x08, 0x1a, 0x0c, 0xa2, 0xc0, 0x02, 0x08, 0x47, 0x65, 0x61, 0x72, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x1e, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x44, 0x45, 0x53, 0x4b, 0x10, 0x09, 0x1a, 0x0d, 0xa2, 0xc0, 0x02, 0x09, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x6b, 0x12, 0x33, 0x0a, 0x1f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54,
	0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54,
	0x49, 0x4b, 0x54, 0x4f, 0x4b, 0x53, 0x48, 0x4f, 0x50, 0x10, 0x0a, 0x1a, 0x0e, 0xa2, 0xc0, 0x02,
	0x0a, 0x54, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x2f, 0x0a, 0x1d, 0x4d,
	0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x4f, 0x53, 0x48, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x0b, 0x1a, 0x0c,
	0xa2, 0xc0, 0x02, 0x08, 0x50, 0x6f, 0x73, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x33, 0x0a, 0x1f,
	0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x54, 0x41, 0x53, 0x48, 0x4f, 0x50, 0x10,
	0x0c, 0x1a, 0x0e, 0xa2, 0xc0, 0x02, 0x0a, 0x50, 0x72, 0x65, 0x73, 0x74, 0x61, 0x73, 0x68, 0x6f,
	0x70, 0x12, 0x29, 0x0a, 0x1a, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4e, 0x4b, 0x47, 0x4f, 0x10,
	0x0d, 0x1a, 0x09, 0xa2, 0xc0, 0x02, 0x05, 0x49, 0x6e, 0x6b, 0x67, 0x6f, 0x12, 0x27, 0x0a, 0x19,
	0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49, 0x53, 0x48, 0x10, 0x0e, 0x1a, 0x08, 0xa2, 0xc0, 0x02,
	0x04, 0x57, 0x69, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x20, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x42, 0x49,
	0x47, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x52, 0x43, 0x45, 0x10, 0x0f, 0x1a, 0x0f, 0xa2, 0xc0, 0x02,
	0x0b, 0x42, 0x69, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x42, 0xd7, 0x01, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65,
	0x61, 0x2d, 0x6e, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x64, 0x6b,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x50, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_platform_v1_platform_proto_rawDescOnce sync.Once
	file_common_platform_v1_platform_proto_rawDescData []byte
)

func file_common_platform_v1_platform_proto_rawDescGZIP() []byte {
	file_common_platform_v1_platform_proto_rawDescOnce.Do(func() {
		file_common_platform_v1_platform_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_platform_v1_platform_proto_rawDesc), len(file_common_platform_v1_platform_proto_rawDesc)))
	})
	return file_common_platform_v1_platform_proto_rawDescData
}

var file_common_platform_v1_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_platform_v1_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_platform_v1_platform_proto_goTypes = []any{
	(App)(0),                                  // 0: common.platform.v1.App
	(MarketplacePlatform)(0),                  // 1: common.platform.v1.MarketplacePlatform
	(*ActivityCallback)(nil),                  // 2: common.platform.v1.ActivityCallback
	(*MarketplacePlatformAndStringTuple)(nil), // 3: common.platform.v1.MarketplacePlatformAndStringTuple
}
var file_common_platform_v1_platform_proto_depIdxs = []int32{
	1, // 0: common.platform.v1.MarketplacePlatformAndStringTuple.platform:type_name -> common.platform.v1.MarketplacePlatform
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_platform_v1_platform_proto_init() }
func file_common_platform_v1_platform_proto_init() {
	if File_common_platform_v1_platform_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_platform_v1_platform_proto_rawDesc), len(file_common_platform_v1_platform_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_platform_v1_platform_proto_goTypes,
		DependencyIndexes: file_common_platform_v1_platform_proto_depIdxs,
		EnumInfos:         file_common_platform_v1_platform_proto_enumTypes,
		MessageInfos:      file_common_platform_v1_platform_proto_msgTypes,
	}.Build()
	File_common_platform_v1_platform_proto = out.File
	file_common_platform_v1_platform_proto_goTypes = nil
	file_common_platform_v1_platform_proto_depIdxs = nil
}
