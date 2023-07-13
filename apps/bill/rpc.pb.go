// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/bill/pb/rpc.proto

package bill

import (
	resource "github.com/ahwhy/yCmdb/apps/resource"
	page "github.com/infraboard/mcube/pb/page"
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

type QueryBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页
	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// 实例所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// 实例所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 资源所属环境
	// @gotags: json:"env"
	Env string `protobuf:"bytes,4,opt,name=env,proto3" json:"env,omitempty"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.VENDOR `protobuf:"varint,5,opt,name=vendor,proto3,enum=ahwhy.yCmdb.resource.VENDOR" json:"vendor,omitempty"`
	// 账单的月份
	// @gotags: json:"month"
	Month string `protobuf:"bytes,6,opt,name=month,proto3" json:"month,omitempty"`
	// 账号Id
	// @gotags: json:"account_id"
	AccountId string `protobuf:"bytes,7,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// 地域Id
	// @gotags: json:"region_code"
	RegionCode string `protobuf:"bytes,8,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// 产品编码
	// @gotags: json:"product_code"
	ProductCode string `protobuf:"bytes,9,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	// 账单月当时标签
	// @gotags: json:"tags"
	Tags []*resource.Tag `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	// 实例ID
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,11,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *QueryBillRequest) Reset() {
	*x = QueryBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBillRequest) ProtoMessage() {}

func (x *QueryBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBillRequest.ProtoReflect.Descriptor instead.
func (*QueryBillRequest) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryBillRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryBillRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryBillRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryBillRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *QueryBillRequest) GetVendor() resource.VENDOR {
	if x != nil {
		return x.Vendor
	}
	return resource.VENDOR(0)
}

func (x *QueryBillRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *QueryBillRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *QueryBillRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *QueryBillRequest) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *QueryBillRequest) GetTags() []*resource.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *QueryBillRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type ConfirmBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账单的月份
	// @gotags: json:"task_id" validate:"required"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *ConfirmBillRequest) Reset() {
	*x = ConfirmBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmBillRequest) ProtoMessage() {}

func (x *ConfirmBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmBillRequest.ProtoReflect.Descriptor instead.
func (*ConfirmBillRequest) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmBillRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type DeleteBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账单的月份
	// @gotags: json:"task_id" validate:"required"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeleteBillRequest) Reset() {
	*x = DeleteBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBillRequest) ProtoMessage() {}

func (x *DeleteBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBillRequest.ProtoReflect.Descriptor instead.
func (*DeleteBillRequest) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBillRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_apps_bill_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_bill_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e,
	0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x68, 0x77, 0x68,
	0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79,
	0x43, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x32, 0xaf, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x16, 0x2e,
	0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x1a, 0x16, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x4a, 0x0a,
	0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x61, 0x68, 0x77,
	0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x24, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79,
	0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x23, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e,
	0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e,
	0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2f, 0x79, 0x43, 0x6d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_bill_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_bill_pb_rpc_proto_rawDescData = file_apps_bill_pb_rpc_proto_rawDesc
)

func file_apps_bill_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_bill_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_bill_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_bill_pb_rpc_proto_rawDescData)
	})
	return file_apps_bill_pb_rpc_proto_rawDescData
}

var file_apps_bill_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_bill_pb_rpc_proto_goTypes = []interface{}{
	(*QueryBillRequest)(nil),   // 0: ahwhy.yCmdb.bill.QueryBillRequest
	(*ConfirmBillRequest)(nil), // 1: ahwhy.yCmdb.bill.ConfirmBillRequest
	(*DeleteBillRequest)(nil),  // 2: ahwhy.yCmdb.bill.DeleteBillRequest
	(*page.PageRequest)(nil),   // 3: infraboard.mcube.page.PageRequest
	(resource.VENDOR)(0),       // 4: ahwhy.yCmdb.resource.VENDOR
	(*resource.Tag)(nil),       // 5: ahwhy.yCmdb.resource.Tag
	(*Bill)(nil),               // 6: ahwhy.yCmdb.bill.Bill
	(*BillSet)(nil),            // 7: ahwhy.yCmdb.bill.BillSet
}
var file_apps_bill_pb_rpc_proto_depIdxs = []int32{
	3, // 0: ahwhy.yCmdb.bill.QueryBillRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4, // 1: ahwhy.yCmdb.bill.QueryBillRequest.vendor:type_name -> ahwhy.yCmdb.resource.VENDOR
	5, // 2: ahwhy.yCmdb.bill.QueryBillRequest.tags:type_name -> ahwhy.yCmdb.resource.Tag
	6, // 3: ahwhy.yCmdb.bill.Service.SyncBill:input_type -> ahwhy.yCmdb.bill.Bill
	0, // 4: ahwhy.yCmdb.bill.Service.QueryBill:input_type -> ahwhy.yCmdb.bill.QueryBillRequest
	1, // 5: ahwhy.yCmdb.bill.Service.ConfirmBill:input_type -> ahwhy.yCmdb.bill.ConfirmBillRequest
	2, // 6: ahwhy.yCmdb.bill.Service.DeleteBill:input_type -> ahwhy.yCmdb.bill.DeleteBillRequest
	6, // 7: ahwhy.yCmdb.bill.Service.SyncBill:output_type -> ahwhy.yCmdb.bill.Bill
	7, // 8: ahwhy.yCmdb.bill.Service.QueryBill:output_type -> ahwhy.yCmdb.bill.BillSet
	7, // 9: ahwhy.yCmdb.bill.Service.ConfirmBill:output_type -> ahwhy.yCmdb.bill.BillSet
	7, // 10: ahwhy.yCmdb.bill.Service.DeleteBill:output_type -> ahwhy.yCmdb.bill.BillSet
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_bill_pb_rpc_proto_init() }
func file_apps_bill_pb_rpc_proto_init() {
	if File_apps_bill_pb_rpc_proto != nil {
		return
	}
	file_apps_bill_pb_bill_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_bill_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBillRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_bill_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmBillRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_bill_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBillRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_bill_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_bill_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_bill_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_bill_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_bill_pb_rpc_proto = out.File
	file_apps_bill_pb_rpc_proto_rawDesc = nil
	file_apps_bill_pb_rpc_proto_goTypes = nil
	file_apps_bill_pb_rpc_proto_depIdxs = nil
}
