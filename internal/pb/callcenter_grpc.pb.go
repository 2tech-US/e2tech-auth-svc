// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/pb/callcenter.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CallCenterServiceClient is the client API for CallCenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallCenterServiceClient interface {
	GetEmployee(ctx context.Context, in *GetEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error)
	GetListEmployee(ctx context.Context, in *GetListEmployeeRequest, opts ...grpc.CallOption) (*GetListEmployeeResponse, error)
	CreateEmployee(ctx context.Context, in *CreateCallCenterEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error)
	UpdateEmployee(ctx context.Context, in *UpdateCallCenterEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error)
}

type callCenterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCallCenterServiceClient(cc grpc.ClientConnInterface) CallCenterServiceClient {
	return &callCenterServiceClient{cc}
}

func (c *callCenterServiceClient) GetEmployee(ctx context.Context, in *GetEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error) {
	out := new(GetEmployeeResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.CallCenterService/getEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callCenterServiceClient) GetListEmployee(ctx context.Context, in *GetListEmployeeRequest, opts ...grpc.CallOption) (*GetListEmployeeResponse, error) {
	out := new(GetListEmployeeResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.CallCenterService/getListEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callCenterServiceClient) CreateEmployee(ctx context.Context, in *CreateCallCenterEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error) {
	out := new(GetEmployeeResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.CallCenterService/createEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callCenterServiceClient) UpdateEmployee(ctx context.Context, in *UpdateCallCenterEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error) {
	out := new(GetEmployeeResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.CallCenterService/updateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallCenterServiceServer is the server API for CallCenterService service.
// All implementations must embed UnimplementedCallCenterServiceServer
// for forward compatibility
type CallCenterServiceServer interface {
	GetEmployee(context.Context, *GetEmployeeRequest) (*GetEmployeeResponse, error)
	GetListEmployee(context.Context, *GetListEmployeeRequest) (*GetListEmployeeResponse, error)
	CreateEmployee(context.Context, *CreateCallCenterEmployeeRequest) (*GetEmployeeResponse, error)
	UpdateEmployee(context.Context, *UpdateCallCenterEmployeeRequest) (*GetEmployeeResponse, error)
	mustEmbedUnimplementedCallCenterServiceServer()
}

// UnimplementedCallCenterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCallCenterServiceServer struct {
}

func (UnimplementedCallCenterServiceServer) GetEmployee(context.Context, *GetEmployeeRequest) (*GetEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (UnimplementedCallCenterServiceServer) GetListEmployee(context.Context, *GetListEmployeeRequest) (*GetListEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListEmployee not implemented")
}
func (UnimplementedCallCenterServiceServer) CreateEmployee(context.Context, *CreateCallCenterEmployeeRequest) (*GetEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedCallCenterServiceServer) UpdateEmployee(context.Context, *UpdateCallCenterEmployeeRequest) (*GetEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedCallCenterServiceServer) mustEmbedUnimplementedCallCenterServiceServer() {}

// UnsafeCallCenterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallCenterServiceServer will
// result in compilation errors.
type UnsafeCallCenterServiceServer interface {
	mustEmbedUnimplementedCallCenterServiceServer()
}

func RegisterCallCenterServiceServer(s grpc.ServiceRegistrar, srv CallCenterServiceServer) {
	s.RegisterService(&CallCenterService_ServiceDesc, srv)
}

func _CallCenterService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallCenterServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.CallCenterService/getEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallCenterServiceServer).GetEmployee(ctx, req.(*GetEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallCenterService_GetListEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallCenterServiceServer).GetListEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.CallCenterService/getListEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallCenterServiceServer).GetListEmployee(ctx, req.(*GetListEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallCenterService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCallCenterEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallCenterServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.CallCenterService/createEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallCenterServiceServer).CreateEmployee(ctx, req.(*CreateCallCenterEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallCenterService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCallCenterEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallCenterServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.CallCenterService/updateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallCenterServiceServer).UpdateEmployee(ctx, req.(*UpdateCallCenterEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CallCenterService_ServiceDesc is the grpc.ServiceDesc for CallCenterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallCenterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tech2.microservice.CallCenterService",
	HandlerType: (*CallCenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getEmployee",
			Handler:    _CallCenterService_GetEmployee_Handler,
		},
		{
			MethodName: "getListEmployee",
			Handler:    _CallCenterService_GetListEmployee_Handler,
		},
		{
			MethodName: "createEmployee",
			Handler:    _CallCenterService_CreateEmployee_Handler,
		},
		{
			MethodName: "updateEmployee",
			Handler:    _CallCenterService_UpdateEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/callcenter.proto",
}
