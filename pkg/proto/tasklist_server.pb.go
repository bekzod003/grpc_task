// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/proto/tasklist_server.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Creating struct Task
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Priority  string `protobuf:"bytes,4,opt,name=Priority,proto3" json:"Priority,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	CreatedBy string `protobuf:"bytes,6,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	DueDate   string `protobuf:"bytes,7,opt,name=DueDate,proto3" json:"DueDate,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_tasklist_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_tasklist_server_proto_msgTypes[0]
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
	return file_api_proto_tasklist_server_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Task) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *Task) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Task) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Task) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

// Request structs
type TaskRequestID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *TaskRequestID) Reset() {
	*x = TaskRequestID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_tasklist_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequestID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequestID) ProtoMessage() {}

func (x *TaskRequestID) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_tasklist_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequestID.ProtoReflect.Descriptor instead.
func (*TaskRequestID) Descriptor() ([]byte, []int) {
	return file_api_proto_tasklist_server_proto_rawDescGZIP(), []int{1}
}

func (x *TaskRequestID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T *Task `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_tasklist_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_tasklist_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_tasklist_server_proto_rawDescGZIP(), []int{2}
}

func (x *TaskRequest) GetT() *Task {
	if x != nil {
		return x.T
	}
	return nil
}

// Response structs
type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T *Task `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_tasklist_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_tasklist_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_tasklist_server_proto_rawDescGZIP(), []int{3}
}

func (x *TaskResponse) GetT() *Task {
	if x != nil {
		return x.T
	}
	return nil
}

type TaskResponseSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T []*Task `protobuf:"bytes,1,rep,name=t,proto3" json:"t,omitempty"`
}

func (x *TaskResponseSlice) Reset() {
	*x = TaskResponseSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_tasklist_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponseSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponseSlice) ProtoMessage() {}

func (x *TaskResponseSlice) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_tasklist_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponseSlice.ProtoReflect.Descriptor instead.
func (*TaskResponseSlice) Descriptor() ([]byte, []int) {
	return file_api_proto_tasklist_server_proto_rawDescGZIP(), []int{4}
}

func (x *TaskResponseSlice) GetT() []*Task {
	if x != nil {
		return x.T
	}
	return nil
}

var File_api_proto_tasklist_server_proto protoreflect.FileDescriptor

var file_api_proto_tasklist_server_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0b, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x01, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x01, 0x74, 0x22, 0x27, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x01, 0x74, 0x22, 0x2c, 0x0a, 0x11,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x12, 0x17, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x01, 0x74, 0x32, 0xb7, 0x02, 0x0a, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_tasklist_server_proto_rawDescOnce sync.Once
	file_api_proto_tasklist_server_proto_rawDescData = file_api_proto_tasklist_server_proto_rawDesc
)

func file_api_proto_tasklist_server_proto_rawDescGZIP() []byte {
	file_api_proto_tasklist_server_proto_rawDescOnce.Do(func() {
		file_api_proto_tasklist_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_tasklist_server_proto_rawDescData)
	})
	return file_api_proto_tasklist_server_proto_rawDescData
}

var file_api_proto_tasklist_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_tasklist_server_proto_goTypes = []interface{}{
	(*Task)(nil),              // 0: api.Task
	(*TaskRequestID)(nil),     // 1: api.TaskRequestID
	(*TaskRequest)(nil),       // 2: api.TaskRequest
	(*TaskResponse)(nil),      // 3: api.TaskResponse
	(*TaskResponseSlice)(nil), // 4: api.TaskResponseSlice
	(*emptypb.Empty)(nil),     // 5: google.protobuf.Empty
}
var file_api_proto_tasklist_server_proto_depIdxs = []int32{
	0, // 0: api.TaskRequest.t:type_name -> api.Task
	0, // 1: api.TaskResponse.t:type_name -> api.Task
	0, // 2: api.TaskResponseSlice.t:type_name -> api.Task
	2, // 3: api.taskList.Create:input_type -> api.TaskRequest
	2, // 4: api.taskList.Add:input_type -> api.TaskRequest
	2, // 5: api.taskList.Update:input_type -> api.TaskRequest
	2, // 6: api.taskList.Delete:input_type -> api.TaskRequest
	1, // 7: api.taskList.Get:input_type -> api.TaskRequestID
	5, // 8: api.taskList.GetAll:input_type -> google.protobuf.Empty
	3, // 9: api.taskList.Create:output_type -> api.TaskResponse
	3, // 10: api.taskList.Add:output_type -> api.TaskResponse
	3, // 11: api.taskList.Update:output_type -> api.TaskResponse
	3, // 12: api.taskList.Delete:output_type -> api.TaskResponse
	3, // 13: api.taskList.Get:output_type -> api.TaskResponse
	4, // 14: api.taskList.GetAll:output_type -> api.TaskResponseSlice
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_tasklist_server_proto_init() }
func file_api_proto_tasklist_server_proto_init() {
	if File_api_proto_tasklist_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_tasklist_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_tasklist_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequestID); i {
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
		file_api_proto_tasklist_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_api_proto_tasklist_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
		file_api_proto_tasklist_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponseSlice); i {
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
			RawDescriptor: file_api_proto_tasklist_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_tasklist_server_proto_goTypes,
		DependencyIndexes: file_api_proto_tasklist_server_proto_depIdxs,
		MessageInfos:      file_api_proto_tasklist_server_proto_msgTypes,
	}.Build()
	File_api_proto_tasklist_server_proto = out.File
	file_api_proto_tasklist_server_proto_rawDesc = nil
	file_api_proto_tasklist_server_proto_goTypes = nil
	file_api_proto_tasklist_server_proto_depIdxs = nil
}