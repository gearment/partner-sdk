// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/support/v1/vendor_support.proto

package supportv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SupportRequest_Status int32

const (
	SupportRequest_SUPPORT_REQUEST_STATUS_UNSPECIFIED SupportRequest_Status = 0
	SupportRequest_SUPPORT_REQUEST_STATUS_PENDING     SupportRequest_Status = 1
	SupportRequest_SUPPORT_REQUEST_STATUS_PROCESSING  SupportRequest_Status = 2
	SupportRequest_SUPPORT_REQUEST_STATUS_COMPLETED   SupportRequest_Status = 3
	SupportRequest_SUPPORT_REQUEST_STATUS_REJECTED    SupportRequest_Status = 4
)

// Enum value maps for SupportRequest_Status.
var (
	SupportRequest_Status_name = map[int32]string{
		0: "SUPPORT_REQUEST_STATUS_UNSPECIFIED",
		1: "SUPPORT_REQUEST_STATUS_PENDING",
		2: "SUPPORT_REQUEST_STATUS_PROCESSING",
		3: "SUPPORT_REQUEST_STATUS_COMPLETED",
		4: "SUPPORT_REQUEST_STATUS_REJECTED",
	}
	SupportRequest_Status_value = map[string]int32{
		"SUPPORT_REQUEST_STATUS_UNSPECIFIED": 0,
		"SUPPORT_REQUEST_STATUS_PENDING":     1,
		"SUPPORT_REQUEST_STATUS_PROCESSING":  2,
		"SUPPORT_REQUEST_STATUS_COMPLETED":   3,
		"SUPPORT_REQUEST_STATUS_REJECTED":    4,
	}
)

func (x SupportRequest_Status) Enum() *SupportRequest_Status {
	p := new(SupportRequest_Status)
	*p = x
	return p
}

func (x SupportRequest_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupportRequest_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_support_v1_vendor_support_proto_enumTypes[0].Descriptor()
}

func (SupportRequest_Status) Type() protoreflect.EnumType {
	return &file_api_support_v1_vendor_support_proto_enumTypes[0]
}

func (x SupportRequest_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupportRequest_Status.Descriptor instead.
func (SupportRequest_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{1, 0}
}

type SupportRequest_Category int32

const (
	// img_quality, item_damaged, item_missing, item_not_expected, no_update_from_carrier, order_late_production, shipping_problem, wrong_item_size, wrong_print, shipping_to_wrong_address, size_not_expected, print_missing, another, cancel_by_buyer
	SupportRequest_SUPPORT_REQUEST_CATEGORY_UNSPECIFIED               SupportRequest_Category = 0
	SupportRequest_SUPPORT_REQUEST_CATEGORY_IMG_QUALITY               SupportRequest_Category = 1
	SupportRequest_SUPPORT_REQUEST_CATEGORY_ITEM_DAMAGED              SupportRequest_Category = 2
	SupportRequest_SUPPORT_REQUEST_CATEGORY_ITEM_MISSING              SupportRequest_Category = 3
	SupportRequest_SUPPORT_REQUEST_CATEGORY_ITEM_NOT_EXPECTED         SupportRequest_Category = 4
	SupportRequest_SUPPORT_REQUEST_CATEGORY_NO_UPDATE_FROM_CARRIER    SupportRequest_Category = 5
	SupportRequest_SUPPORT_REQUEST_CATEGORY_ORDER_LATE_PRODUCTION     SupportRequest_Category = 6
	SupportRequest_SUPPORT_REQUEST_CATEGORY_SHIPPING_PROBLEM          SupportRequest_Category = 7
	SupportRequest_SUPPORT_REQUEST_CATEGORY_WRONG_ITEM_SIZE           SupportRequest_Category = 8
	SupportRequest_SUPPORT_REQUEST_CATEGORY_WRONG_PRINT               SupportRequest_Category = 9
	SupportRequest_SUPPORT_REQUEST_CATEGORY_SHIPPING_TO_WRONG_ADDRESS SupportRequest_Category = 10
	SupportRequest_SUPPORT_REQUEST_CATEGORY_SIZE_NOT_EXPECTED         SupportRequest_Category = 11
	SupportRequest_SUPPORT_REQUEST_CATEGORY_PRINT_MISSING             SupportRequest_Category = 12
	SupportRequest_SUPPORT_REQUEST_CATEGORY_ANOTHER                   SupportRequest_Category = 13
	SupportRequest_SUPPORT_REQUEST_CATEGORY_CANCEL_BY_BUYER           SupportRequest_Category = 14
)

// Enum value maps for SupportRequest_Category.
var (
	SupportRequest_Category_name = map[int32]string{
		0:  "SUPPORT_REQUEST_CATEGORY_UNSPECIFIED",
		1:  "SUPPORT_REQUEST_CATEGORY_IMG_QUALITY",
		2:  "SUPPORT_REQUEST_CATEGORY_ITEM_DAMAGED",
		3:  "SUPPORT_REQUEST_CATEGORY_ITEM_MISSING",
		4:  "SUPPORT_REQUEST_CATEGORY_ITEM_NOT_EXPECTED",
		5:  "SUPPORT_REQUEST_CATEGORY_NO_UPDATE_FROM_CARRIER",
		6:  "SUPPORT_REQUEST_CATEGORY_ORDER_LATE_PRODUCTION",
		7:  "SUPPORT_REQUEST_CATEGORY_SHIPPING_PROBLEM",
		8:  "SUPPORT_REQUEST_CATEGORY_WRONG_ITEM_SIZE",
		9:  "SUPPORT_REQUEST_CATEGORY_WRONG_PRINT",
		10: "SUPPORT_REQUEST_CATEGORY_SHIPPING_TO_WRONG_ADDRESS",
		11: "SUPPORT_REQUEST_CATEGORY_SIZE_NOT_EXPECTED",
		12: "SUPPORT_REQUEST_CATEGORY_PRINT_MISSING",
		13: "SUPPORT_REQUEST_CATEGORY_ANOTHER",
		14: "SUPPORT_REQUEST_CATEGORY_CANCEL_BY_BUYER",
	}
	SupportRequest_Category_value = map[string]int32{
		"SUPPORT_REQUEST_CATEGORY_UNSPECIFIED":               0,
		"SUPPORT_REQUEST_CATEGORY_IMG_QUALITY":               1,
		"SUPPORT_REQUEST_CATEGORY_ITEM_DAMAGED":              2,
		"SUPPORT_REQUEST_CATEGORY_ITEM_MISSING":              3,
		"SUPPORT_REQUEST_CATEGORY_ITEM_NOT_EXPECTED":         4,
		"SUPPORT_REQUEST_CATEGORY_NO_UPDATE_FROM_CARRIER":    5,
		"SUPPORT_REQUEST_CATEGORY_ORDER_LATE_PRODUCTION":     6,
		"SUPPORT_REQUEST_CATEGORY_SHIPPING_PROBLEM":          7,
		"SUPPORT_REQUEST_CATEGORY_WRONG_ITEM_SIZE":           8,
		"SUPPORT_REQUEST_CATEGORY_WRONG_PRINT":               9,
		"SUPPORT_REQUEST_CATEGORY_SHIPPING_TO_WRONG_ADDRESS": 10,
		"SUPPORT_REQUEST_CATEGORY_SIZE_NOT_EXPECTED":         11,
		"SUPPORT_REQUEST_CATEGORY_PRINT_MISSING":             12,
		"SUPPORT_REQUEST_CATEGORY_ANOTHER":                   13,
		"SUPPORT_REQUEST_CATEGORY_CANCEL_BY_BUYER":           14,
	}
)

func (x SupportRequest_Category) Enum() *SupportRequest_Category {
	p := new(SupportRequest_Category)
	*p = x
	return p
}

func (x SupportRequest_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupportRequest_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_api_support_v1_vendor_support_proto_enumTypes[1].Descriptor()
}

func (SupportRequest_Category) Type() protoreflect.EnumType {
	return &file_api_support_v1_vendor_support_proto_enumTypes[1]
}

func (x SupportRequest_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupportRequest_Category.Descriptor instead.
func (SupportRequest_Category) EnumDescriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{1, 1}
}

type ListSupportRequestRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestStatus SupportRequest_Status  `protobuf:"varint,1,opt,name=request_status,json=requestStatus,proto3,enum=api.support.v1.SupportRequest_Status" json:"request_status,omitempty"`
	CreatedFrom   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_from,json=createdFrom,proto3" json:"created_from,omitempty"`
	CreatedTo     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_to,json=createdTo,proto3" json:"created_to,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSupportRequestRequest) Reset() {
	*x = ListSupportRequestRequest{}
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSupportRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupportRequestRequest) ProtoMessage() {}

func (x *ListSupportRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupportRequestRequest.ProtoReflect.Descriptor instead.
func (*ListSupportRequestRequest) Descriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{0}
}

func (x *ListSupportRequestRequest) GetRequestStatus() SupportRequest_Status {
	if x != nil {
		return x.RequestStatus
	}
	return SupportRequest_SUPPORT_REQUEST_STATUS_UNSPECIFIED
}

func (x *ListSupportRequestRequest) GetCreatedFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedFrom
	}
	return nil
}

func (x *ListSupportRequestRequest) GetCreatedTo() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTo
	}
	return nil
}

func (x *ListSupportRequestRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListSupportRequestRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type SupportRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId        string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status         SupportRequest_Status  `protobuf:"varint,4,opt,name=status,proto3,enum=api.support.v1.SupportRequest_Status" json:"status,omitempty"`
	IssueCategory  string                 `protobuf:"bytes,5,opt,name=issue_category,json=issueCategory,proto3" json:"issue_category,omitempty"`
	IssueDetail    string                 `protobuf:"bytes,6,opt,name=issue_detail,json=issueDetail,proto3" json:"issue_detail,omitempty"`
	IssuePriority  string                 `protobuf:"bytes,7,opt,name=issue_priority,json=issuePriority,proto3" json:"issue_priority,omitempty"`
	IssueSolution  string                 `protobuf:"bytes,8,opt,name=issue_solution,json=issueSolution,proto3" json:"issue_solution,omitempty"`
	IssueImageUrls []string               `protobuf:"bytes,9,rep,name=issue_image_urls,json=issueImageUrls,proto3" json:"issue_image_urls,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SupportRequest) Reset() {
	*x = SupportRequest{}
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportRequest) ProtoMessage() {}

func (x *SupportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportRequest.ProtoReflect.Descriptor instead.
func (*SupportRequest) Descriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{1}
}

func (x *SupportRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SupportRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *SupportRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SupportRequest) GetStatus() SupportRequest_Status {
	if x != nil {
		return x.Status
	}
	return SupportRequest_SUPPORT_REQUEST_STATUS_UNSPECIFIED
}

func (x *SupportRequest) GetIssueCategory() string {
	if x != nil {
		return x.IssueCategory
	}
	return ""
}

func (x *SupportRequest) GetIssueDetail() string {
	if x != nil {
		return x.IssueDetail
	}
	return ""
}

func (x *SupportRequest) GetIssuePriority() string {
	if x != nil {
		return x.IssuePriority
	}
	return ""
}

func (x *SupportRequest) GetIssueSolution() string {
	if x != nil {
		return x.IssueSolution
	}
	return ""
}

func (x *SupportRequest) GetIssueImageUrls() []string {
	if x != nil {
		return x.IssueImageUrls
	}
	return nil
}

type ListSupportRequestResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SupportRequests []*SupportRequest      `protobuf:"bytes,1,rep,name=support_requests,json=supportRequests,proto3" json:"support_requests,omitempty"`
	Total           int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListSupportRequestResponse) Reset() {
	*x = ListSupportRequestResponse{}
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSupportRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupportRequestResponse) ProtoMessage() {}

func (x *ListSupportRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupportRequestResponse.ProtoReflect.Descriptor instead.
func (*ListSupportRequestResponse) Descriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{2}
}

func (x *ListSupportRequestResponse) GetSupportRequests() []*SupportRequest {
	if x != nil {
		return x.SupportRequests
	}
	return nil
}

func (x *ListSupportRequestResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateSupportRequestRequest struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	OrderId        string                  `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	IssueCategory  SupportRequest_Category `protobuf:"varint,2,opt,name=issue_category,json=issueCategory,proto3,enum=api.support.v1.SupportRequest_Category" json:"issue_category,omitempty"`
	IssueDetail    string                  `protobuf:"bytes,3,opt,name=issue_detail,json=issueDetail,proto3" json:"issue_detail,omitempty"`
	IssuePriority  string                  `protobuf:"bytes,4,opt,name=issue_priority,json=issuePriority,proto3" json:"issue_priority,omitempty"`
	IssueSolution  string                  `protobuf:"bytes,5,opt,name=issue_solution,json=issueSolution,proto3" json:"issue_solution,omitempty"`
	IssueImageUrls []string                `protobuf:"bytes,6,rep,name=issue_image_urls,json=issueImageUrls,proto3" json:"issue_image_urls,omitempty"`
	ItemAffected   int32                   `protobuf:"varint,7,opt,name=item_affected,json=itemAffected,proto3" json:"item_affected,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateSupportRequestRequest) Reset() {
	*x = CreateSupportRequestRequest{}
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSupportRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupportRequestRequest) ProtoMessage() {}

func (x *CreateSupportRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupportRequestRequest.ProtoReflect.Descriptor instead.
func (*CreateSupportRequestRequest) Descriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSupportRequestRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateSupportRequestRequest) GetIssueCategory() SupportRequest_Category {
	if x != nil {
		return x.IssueCategory
	}
	return SupportRequest_SUPPORT_REQUEST_CATEGORY_UNSPECIFIED
}

func (x *CreateSupportRequestRequest) GetIssueDetail() string {
	if x != nil {
		return x.IssueDetail
	}
	return ""
}

func (x *CreateSupportRequestRequest) GetIssuePriority() string {
	if x != nil {
		return x.IssuePriority
	}
	return ""
}

func (x *CreateSupportRequestRequest) GetIssueSolution() string {
	if x != nil {
		return x.IssueSolution
	}
	return ""
}

func (x *CreateSupportRequestRequest) GetIssueImageUrls() []string {
	if x != nil {
		return x.IssueImageUrls
	}
	return nil
}

func (x *CreateSupportRequestRequest) GetItemAffected() int32 {
	if x != nil {
		return x.ItemAffected
	}
	return 0
}

type GetSupportRequestResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	SupportRequest *SupportRequest        `protobuf:"bytes,1,opt,name=support_request,json=supportRequest,proto3" json:"support_request,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetSupportRequestResponse) Reset() {
	*x = GetSupportRequestResponse{}
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSupportRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportRequestResponse) ProtoMessage() {}

func (x *GetSupportRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_support_v1_vendor_support_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportRequestResponse.ProtoReflect.Descriptor instead.
func (*GetSupportRequestResponse) Descriptor() ([]byte, []int) {
	return file_api_support_v1_vendor_support_proto_rawDescGZIP(), []int{4}
}

func (x *GetSupportRequestResponse) GetSupportRequest() *SupportRequest {
	if x != nil {
		return x.SupportRequest
	}
	return nil
}

var File_api_support_v1_vendor_support_proto protoreflect.FileDescriptor

var file_api_support_v1_vendor_support_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x56, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x6f, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x18, 0x64, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xff, 0x09, 0x0a, 0x0e, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x26, 0x0a, 0x22, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x55, 0x50, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x22, 0xbc, 0x05,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x49, 0x4d, 0x47, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x29,
	0x0a, 0x25, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f,
	0x44, 0x41, 0x4d, 0x41, 0x47, 0x45, 0x44, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x25, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x2e, 0x0a, 0x2a, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x33, 0x0a, 0x2f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x4e, 0x4f, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f,
	0x43, 0x41, 0x52, 0x52, 0x49, 0x45, 0x52, 0x10, 0x05, 0x12, 0x32, 0x0a, 0x2e, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x4c, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x2d, 0x0a,
	0x29, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x07, 0x12, 0x2c, 0x0a, 0x28,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x49,
	0x54, 0x45, 0x4d, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x08, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x49,
	0x4e, 0x54, 0x10, 0x09, 0x12, 0x36, 0x0a, 0x32, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x57, 0x52, 0x4f,
	0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x2e, 0x0a, 0x2a,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x2a, 0x0a, 0x26,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x0c, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x4e, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x2c,
	0x0a, 0x28, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x5f, 0x42, 0x59, 0x5f, 0x42, 0x55, 0x59, 0x45, 0x52, 0x10, 0x0e, 0x22, 0x7d, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xe3, 0x02, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x10, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x10, 0xba, 0x48, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x72, 0x06, 0x18, 0x80, 0x10,
	0x88, 0x01, 0x01, 0x52, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a,
	0x02, 0x28, 0x01, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x22, 0x64, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xf3, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x41, 0x50, 0x49, 0x12, 0x70, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x73, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x02, 0x42, 0xc3, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x67, 0x65, 0x61, 0x2d, 0x6e, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53,
	0x58, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_support_v1_vendor_support_proto_rawDescOnce sync.Once
	file_api_support_v1_vendor_support_proto_rawDescData []byte
)

func file_api_support_v1_vendor_support_proto_rawDescGZIP() []byte {
	file_api_support_v1_vendor_support_proto_rawDescOnce.Do(func() {
		file_api_support_v1_vendor_support_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_support_v1_vendor_support_proto_rawDesc), len(file_api_support_v1_vendor_support_proto_rawDesc)))
	})
	return file_api_support_v1_vendor_support_proto_rawDescData
}

var file_api_support_v1_vendor_support_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_support_v1_vendor_support_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_support_v1_vendor_support_proto_goTypes = []any{
	(SupportRequest_Status)(0),          // 0: api.support.v1.SupportRequest.Status
	(SupportRequest_Category)(0),        // 1: api.support.v1.SupportRequest.Category
	(*ListSupportRequestRequest)(nil),   // 2: api.support.v1.ListSupportRequestRequest
	(*SupportRequest)(nil),              // 3: api.support.v1.SupportRequest
	(*ListSupportRequestResponse)(nil),  // 4: api.support.v1.ListSupportRequestResponse
	(*CreateSupportRequestRequest)(nil), // 5: api.support.v1.CreateSupportRequestRequest
	(*GetSupportRequestResponse)(nil),   // 6: api.support.v1.GetSupportRequestResponse
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_api_support_v1_vendor_support_proto_depIdxs = []int32{
	0,  // 0: api.support.v1.ListSupportRequestRequest.request_status:type_name -> api.support.v1.SupportRequest.Status
	7,  // 1: api.support.v1.ListSupportRequestRequest.created_from:type_name -> google.protobuf.Timestamp
	7,  // 2: api.support.v1.ListSupportRequestRequest.created_to:type_name -> google.protobuf.Timestamp
	7,  // 3: api.support.v1.SupportRequest.created_at:type_name -> google.protobuf.Timestamp
	0,  // 4: api.support.v1.SupportRequest.status:type_name -> api.support.v1.SupportRequest.Status
	3,  // 5: api.support.v1.ListSupportRequestResponse.support_requests:type_name -> api.support.v1.SupportRequest
	1,  // 6: api.support.v1.CreateSupportRequestRequest.issue_category:type_name -> api.support.v1.SupportRequest.Category
	3,  // 7: api.support.v1.GetSupportRequestResponse.support_request:type_name -> api.support.v1.SupportRequest
	2,  // 8: api.support.v1.SupportAPI.ListSupportRequest:input_type -> api.support.v1.ListSupportRequestRequest
	5,  // 9: api.support.v1.SupportAPI.CreateSupportRequest:input_type -> api.support.v1.CreateSupportRequestRequest
	4,  // 10: api.support.v1.SupportAPI.ListSupportRequest:output_type -> api.support.v1.ListSupportRequestResponse
	6,  // 11: api.support.v1.SupportAPI.CreateSupportRequest:output_type -> api.support.v1.GetSupportRequestResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_support_v1_vendor_support_proto_init() }
func file_api_support_v1_vendor_support_proto_init() {
	if File_api_support_v1_vendor_support_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_support_v1_vendor_support_proto_rawDesc), len(file_api_support_v1_vendor_support_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_support_v1_vendor_support_proto_goTypes,
		DependencyIndexes: file_api_support_v1_vendor_support_proto_depIdxs,
		EnumInfos:         file_api_support_v1_vendor_support_proto_enumTypes,
		MessageInfos:      file_api_support_v1_vendor_support_proto_msgTypes,
	}.Build()
	File_api_support_v1_vendor_support_proto = out.File
	file_api_support_v1_vendor_support_proto_goTypes = nil
	file_api_support_v1_vendor_support_proto_depIdxs = nil
}
