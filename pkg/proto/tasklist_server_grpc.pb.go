// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskListClient is the client API for TaskList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskListClient interface {
	Create(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Add(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Update(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Delete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Get(ctx context.Context, in *TaskRequestID, opts ...grpc.CallOption) (*TaskResponse, error)
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TaskResponseSlice, error)
}

type taskListClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskListClient(cc grpc.ClientConnInterface) TaskListClient {
	return &taskListClient{cc}
}

func (c *taskListClient) Create(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/api.taskList/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskListClient) Add(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/api.taskList/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskListClient) Update(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/api.taskList/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskListClient) Delete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/api.taskList/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskListClient) Get(ctx context.Context, in *TaskRequestID, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/api.taskList/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskListClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TaskResponseSlice, error) {
	out := new(TaskResponseSlice)
	err := c.cc.Invoke(ctx, "/api.taskList/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskListServer is the server API for TaskList service.
// All implementations must embed UnimplementedTaskListServer
// for forward compatibility
type TaskListServer interface {
	Create(context.Context, *TaskRequest) (*TaskResponse, error)
	Add(context.Context, *TaskRequest) (*TaskResponse, error)
	Update(context.Context, *TaskRequest) (*TaskResponse, error)
	Delete(context.Context, *TaskRequest) (*TaskResponse, error)
	Get(context.Context, *TaskRequestID) (*TaskResponse, error)
	GetAll(context.Context, *emptypb.Empty) (*TaskResponseSlice, error)
	mustEmbedUnimplementedTaskListServer()
}

// UnimplementedTaskListServer must be embedded to have forward compatible implementations.
type UnimplementedTaskListServer struct {
}

func (UnimplementedTaskListServer) Create(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaskListServer) Add(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTaskListServer) Update(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTaskListServer) Delete(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTaskListServer) Get(context.Context, *TaskRequestID) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTaskListServer) GetAll(context.Context, *emptypb.Empty) (*TaskResponseSlice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTaskListServer) mustEmbedUnimplementedTaskListServer() {}

// UnsafeTaskListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskListServer will
// result in compilation errors.
type UnsafeTaskListServer interface {
	mustEmbedUnimplementedTaskListServer()
}

func RegisterTaskListServer(s grpc.ServiceRegistrar, srv TaskListServer) {
	s.RegisterService(&TaskList_ServiceDesc, srv)
}

func _TaskList_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.taskList/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServer).Create(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskList_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.taskList/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServer).Add(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskList_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.taskList/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServer).Update(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskList_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.taskList/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServer).Delete(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskList_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.taskList/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServer).Get(ctx, req.(*TaskRequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskList_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.taskList/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskList_ServiceDesc is the grpc.ServiceDesc for TaskList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.taskList",
	HandlerType: (*TaskListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TaskList_Create_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _TaskList_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TaskList_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskList_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TaskList_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TaskList_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/tasklist_server.proto",
}
