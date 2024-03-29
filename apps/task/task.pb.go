// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/task/pb/task.proto

package task

import (
	resource "github.com/ahwhy/yCmdb/apps/resource"
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

type Status int32

const (
	Status_PENDDING Status = 0
	// 任务运行中
	Status_RUNNING Status = 1
	// 任务执行成功
	Status_SUCCESS Status = 2
	// 任务执行失败
	Status_FAILED Status = 3
	// 任务部分成功
	Status_WARNING Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDDING",
		1: "RUNNING",
		2: "SUCCESS",
		3: "FAILED",
		4: "WARNING",
	}
	Status_value = map[string]int32{
		"PENDDING": 0,
		"RUNNING":  1,
		"SUCCESS":  2,
		"FAILED":   3,
		"WARNING":  4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_app_task_pb_task_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_app_task_pb_task_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{0}
}

// Task 同个区域的同一种资源一次只能有1个task run
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 同步的区域
	// @gotags: json:"region"
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// 同步的资源
	// @gotags: json:"resource_type"
	ResourceType resource.Type `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=ahwhy.yCmdb.resource.Type" json:"resource_type,omitempty"`
	// 关联secret
	// @gotags: json:"secret_id"
	SecretId string `protobuf:"bytes,4,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	// secret描述
	// @gotags: json:"secret_description"
	SecretDescription string `protobuf:"bytes,5,opt,name=secret_description,json=secretDescription,proto3" json:"secret_description,omitempty"`
	// 任务超时时间
	// @gotags: json:"timeout"
	Timeout int32 `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// 任务状态
	// @gotags: json:"status"
	Status Status `protobuf:"varint,7,opt,name=status,proto3,enum=ahwhy.yCmdb.task.Status" json:"status,omitempty"`
	// 失败时的异常信息
	// @gotags: json:"message"
	Message string `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	// 开始同步的时间
	// @gotags: json:"start_at"
	StartAt int64 `protobuf:"varint,9,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// 同步结束时间
	// @gotags: json:"end_at"
	EndAt int64 `protobuf:"varint,10,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	// 成功的条数
	// @gotags: json:"total_succeed"
	TotalSucceed int64 `protobuf:"varint,11,opt,name=total_succeed,json=totalSucceed,proto3" json:"total_succeed,omitempty"`
	// 失败的条数
	// @gotags: json:"total_failed"
	TotalFailed int64 `protobuf:"varint,12,opt,name=total_failed,json=totalFailed,proto3" json:"total_failed,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Task) GetResourceType() resource.Type {
	if x != nil {
		return x.ResourceType
	}
	return resource.Type(0)
}

func (x *Task) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *Task) GetSecretDescription() string {
	if x != nil {
		return x.SecretDescription
	}
	return ""
}

func (x *Task) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Task) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDDING
}

func (x *Task) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Task) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *Task) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *Task) GetTotalSucceed() int64 {
	if x != nil {
		return x.TotalSucceed
	}
	return 0
}

func (x *Task) GetTotalFailed() int64 {
	if x != nil {
		return x.TotalFailed
	}
	return 0
}

type TaskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总条数
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 具体条目
	// @gotags: json:"items"
	Items []*Task `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TaskSet) Reset() {
	*x = TaskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSet) ProtoMessage() {}

func (x *TaskSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSet.ProtoReflect.Descriptor instead.
func (*TaskSet) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskSet) GetItems() []*Task {
	if x != nil {
		return x.Items
	}
	return nil
}

// 用于描述资源同步的详情信息
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务的Id
	// @gotags: json:"task_id"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// 记录创建时间
	// @gotags: json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	// 资源实例Id
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// 资源名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// 是否同步成功
	// @gotags: json:"is_success"
	IsSuccess bool `protobuf:"varint,5,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	// 同步失败原因
	// @gotags: json:"message"
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Record) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Record) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Record) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *Record) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RecordSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总条数
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 具体条目
	// @gotags: json:"items"
	Items []*Record `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RecordSet) Reset() {
	*x = RecordSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSet) ProtoMessage() {}

func (x *RecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSet.ProtoReflect.Descriptor instead.
func (*RecordSet) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{3}
}

func (x *RecordSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RecordSet) GetItems() []*Record {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateTaskRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"secret_id" validate:"required,lte=100"
	SecretId string `protobuf:"bytes,1,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	// @gotags: json:"region"
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// @gotags: json:"resource_type"
	ResourceType resource.Type `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=ahwhy.yCmdb.resource.Type" json:"resource_type,omitempty"`
	// @gotags: json:"timeout"
	Timeout int64 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *CreateTaskRequst) Reset() {
	*x = CreateTaskRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequst) ProtoMessage() {}

func (x *CreateTaskRequst) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequst.ProtoReflect.Descriptor instead.
func (*CreateTaskRequst) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTaskRequst) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *CreateTaskRequst) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateTaskRequst) GetResourceType() resource.Type {
	if x != nil {
		return x.ResourceType
	}
	return resource.Type(0)
}

func (x *CreateTaskRequst) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type QueryTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page_size"
	PageSize uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// @gotags: json:"page_number"
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// @gotags: json:"resource_type"
	ResourceType resource.Type `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=ahwhy.yCmdb.resource.Type" json:"resource_type,omitempty"`
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *QueryTaskRequest) Reset() {
	*x = QueryTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTaskRequest) ProtoMessage() {}

func (x *QueryTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryTaskRequest) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{5}
}

func (x *QueryTaskRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryTaskRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *QueryTaskRequest) GetResourceType() resource.Type {
	if x != nil {
		return x.ResourceType
	}
	return resource.Type(0)
}

func (x *QueryTaskRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type DescribeTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task id
	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeTaskRequest) Reset() {
	*x = DescribeTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTaskRequest) ProtoMessage() {}

func (x *DescribeTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTaskRequest.ProtoReflect.Descriptor instead.
func (*DescribeTaskRequest) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QueryTaskRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task id
	// @gotags: json:"task_id" validate:"required"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *QueryTaskRecordRequest) Reset() {
	*x = QueryTaskRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_task_pb_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTaskRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTaskRecordRequest) ProtoMessage() {}

func (x *QueryTaskRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_task_pb_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTaskRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryTaskRecordRequest) Descriptor() ([]byte, []int) {
	return file_app_task_pb_task_proto_rawDescGZIP(), []int{7}
}

func (x *QueryTaskRecordRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_app_task_pb_task_proto protoreflect.FileDescriptor

var file_app_task_pb_task_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e,
	0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1e, 0x61, 0x70, 0x70, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x04, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x4d, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79,
	0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79,
	0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x68, 0x77,
	0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xad,
	0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79,
	0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x25,
	0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x2a, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x4e, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x32, 0xc7, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x47, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x2e, 0x61,
	0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x4a, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43,
	0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x68, 0x77, 0x68,
	0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x25, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d,
	0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x68,
	0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x58, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79,
	0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x61, 0x68, 0x77, 0x68, 0x79, 0x2e, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x42, 0x21, 0x5a,
	0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x68, 0x77, 0x68,
	0x79, 0x2f, 0x79, 0x43, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_task_pb_task_proto_rawDescOnce sync.Once
	file_app_task_pb_task_proto_rawDescData = file_app_task_pb_task_proto_rawDesc
)

func file_app_task_pb_task_proto_rawDescGZIP() []byte {
	file_app_task_pb_task_proto_rawDescOnce.Do(func() {
		file_app_task_pb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_task_pb_task_proto_rawDescData)
	})
	return file_app_task_pb_task_proto_rawDescData
}

var file_app_task_pb_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_task_pb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_task_pb_task_proto_goTypes = []interface{}{
	(Status)(0),                    // 0: ahwhy.yCmdb.task.Status
	(*Task)(nil),                   // 1: ahwhy.yCmdb.task.Task
	(*TaskSet)(nil),                // 2: ahwhy.yCmdb.task.TaskSet
	(*Record)(nil),                 // 3: ahwhy.yCmdb.task.Record
	(*RecordSet)(nil),              // 4: ahwhy.yCmdb.task.RecordSet
	(*CreateTaskRequst)(nil),       // 5: ahwhy.yCmdb.task.CreateTaskRequst
	(*QueryTaskRequest)(nil),       // 6: ahwhy.yCmdb.task.QueryTaskRequest
	(*DescribeTaskRequest)(nil),    // 7: ahwhy.yCmdb.task.DescribeTaskRequest
	(*QueryTaskRecordRequest)(nil), // 8: ahwhy.yCmdb.task.QueryTaskRecordRequest
	(resource.Type)(0),             // 9: ahwhy.yCmdb.resource.Type
}
var file_app_task_pb_task_proto_depIdxs = []int32{
	9,  // 0: ahwhy.yCmdb.task.Task.resource_type:type_name -> ahwhy.yCmdb.resource.Type
	0,  // 1: ahwhy.yCmdb.task.Task.status:type_name -> ahwhy.yCmdb.task.Status
	1,  // 2: ahwhy.yCmdb.task.TaskSet.items:type_name -> ahwhy.yCmdb.task.Task
	3,  // 3: ahwhy.yCmdb.task.RecordSet.items:type_name -> ahwhy.yCmdb.task.Record
	9,  // 4: ahwhy.yCmdb.task.CreateTaskRequst.resource_type:type_name -> ahwhy.yCmdb.resource.Type
	9,  // 5: ahwhy.yCmdb.task.QueryTaskRequest.resource_type:type_name -> ahwhy.yCmdb.resource.Type
	5,  // 6: ahwhy.yCmdb.task.Service.CreatTask:input_type -> ahwhy.yCmdb.task.CreateTaskRequst
	6,  // 7: ahwhy.yCmdb.task.Service.QueryTask:input_type -> ahwhy.yCmdb.task.QueryTaskRequest
	7,  // 8: ahwhy.yCmdb.task.Service.DescribeTask:input_type -> ahwhy.yCmdb.task.DescribeTaskRequest
	8,  // 9: ahwhy.yCmdb.task.Service.QueryTaskRecord:input_type -> ahwhy.yCmdb.task.QueryTaskRecordRequest
	1,  // 10: ahwhy.yCmdb.task.Service.CreatTask:output_type -> ahwhy.yCmdb.task.Task
	2,  // 11: ahwhy.yCmdb.task.Service.QueryTask:output_type -> ahwhy.yCmdb.task.TaskSet
	1,  // 12: ahwhy.yCmdb.task.Service.DescribeTask:output_type -> ahwhy.yCmdb.task.Task
	4,  // 13: ahwhy.yCmdb.task.Service.QueryTaskRecord:output_type -> ahwhy.yCmdb.task.RecordSet
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_app_task_pb_task_proto_init() }
func file_app_task_pb_task_proto_init() {
	if File_app_task_pb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_task_pb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_app_task_pb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSet); i {
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
		file_app_task_pb_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_app_task_pb_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSet); i {
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
		file_app_task_pb_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequst); i {
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
		file_app_task_pb_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTaskRequest); i {
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
		file_app_task_pb_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTaskRequest); i {
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
		file_app_task_pb_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTaskRecordRequest); i {
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
			RawDescriptor: file_app_task_pb_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_task_pb_task_proto_goTypes,
		DependencyIndexes: file_app_task_pb_task_proto_depIdxs,
		EnumInfos:         file_app_task_pb_task_proto_enumTypes,
		MessageInfos:      file_app_task_pb_task_proto_msgTypes,
	}.Build()
	File_app_task_pb_task_proto = out.File
	file_app_task_pb_task_proto_rawDesc = nil
	file_app_task_pb_task_proto_goTypes = nil
	file_app_task_pb_task_proto_depIdxs = nil
}
