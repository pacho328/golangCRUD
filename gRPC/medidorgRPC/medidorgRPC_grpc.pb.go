// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: medidorgRPC.proto

package medidorgRPC

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

// MedidorServiceClient is the client API for MedidorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedidorServiceClient interface {
	WriteMedidor(ctx context.Context, in *MedidorRequest, opts ...grpc.CallOption) (*MedidorCreateResponse, error)
	GetMedidor(ctx context.Context, in *MedidorUUID, opts ...grpc.CallOption) (*MedidorGet, error)
	UpdateMedidor(ctx context.Context, in *MedidorUpdate, opts ...grpc.CallOption) (*MedidorResponse, error)
	DeleteMedidor(ctx context.Context, in *MedidorUUID, opts ...grpc.CallOption) (*MedidorResponse, error)
	GetMedidorInstalled(ctx context.Context, in *MedidorIsActive, opts ...grpc.CallOption) (*MedidorIsActiveResponse, error)
	RecentInstallationMarca(ctx context.Context, in *MedidorMarca, opts ...grpc.CallOption) (*MedidorGet, error)
	RecentInstallationSerial(ctx context.Context, in *MedidorSerial, opts ...grpc.CallOption) (*MedidorGet, error)
}

type medidorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedidorServiceClient(cc grpc.ClientConnInterface) MedidorServiceClient {
	return &medidorServiceClient{cc}
}

func (c *medidorServiceClient) WriteMedidor(ctx context.Context, in *MedidorRequest, opts ...grpc.CallOption) (*MedidorCreateResponse, error) {
	out := new(MedidorCreateResponse)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/WriteMedidor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medidorServiceClient) GetMedidor(ctx context.Context, in *MedidorUUID, opts ...grpc.CallOption) (*MedidorGet, error) {
	out := new(MedidorGet)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/GetMedidor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medidorServiceClient) UpdateMedidor(ctx context.Context, in *MedidorUpdate, opts ...grpc.CallOption) (*MedidorResponse, error) {
	out := new(MedidorResponse)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/UpdateMedidor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medidorServiceClient) DeleteMedidor(ctx context.Context, in *MedidorUUID, opts ...grpc.CallOption) (*MedidorResponse, error) {
	out := new(MedidorResponse)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/DeleteMedidor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medidorServiceClient) GetMedidorInstalled(ctx context.Context, in *MedidorIsActive, opts ...grpc.CallOption) (*MedidorIsActiveResponse, error) {
	out := new(MedidorIsActiveResponse)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/GetMedidorInstalled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medidorServiceClient) RecentInstallationMarca(ctx context.Context, in *MedidorMarca, opts ...grpc.CallOption) (*MedidorGet, error) {
	out := new(MedidorGet)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/RecentInstallationMarca", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medidorServiceClient) RecentInstallationSerial(ctx context.Context, in *MedidorSerial, opts ...grpc.CallOption) (*MedidorGet, error) {
	out := new(MedidorGet)
	err := c.cc.Invoke(ctx, "/medidorgRPC.MedidorService/RecentInstallationSerial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedidorServiceServer is the server API for MedidorService service.
// All implementations must embed UnimplementedMedidorServiceServer
// for forward compatibility
type MedidorServiceServer interface {
	WriteMedidor(context.Context, *MedidorRequest) (*MedidorCreateResponse, error)
	GetMedidor(context.Context, *MedidorUUID) (*MedidorGet, error)
	UpdateMedidor(context.Context, *MedidorUpdate) (*MedidorResponse, error)
	DeleteMedidor(context.Context, *MedidorUUID) (*MedidorResponse, error)
	GetMedidorInstalled(context.Context, *MedidorIsActive) (*MedidorIsActiveResponse, error)
	RecentInstallationMarca(context.Context, *MedidorMarca) (*MedidorGet, error)
	RecentInstallationSerial(context.Context, *MedidorSerial) (*MedidorGet, error)
	mustEmbedUnimplementedMedidorServiceServer()
}

// UnimplementedMedidorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMedidorServiceServer struct {
}

func (UnimplementedMedidorServiceServer) WriteMedidor(context.Context, *MedidorRequest) (*MedidorCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMedidor not implemented")
}
func (UnimplementedMedidorServiceServer) GetMedidor(context.Context, *MedidorUUID) (*MedidorGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedidor not implemented")
}
func (UnimplementedMedidorServiceServer) UpdateMedidor(context.Context, *MedidorUpdate) (*MedidorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMedidor not implemented")
}
func (UnimplementedMedidorServiceServer) DeleteMedidor(context.Context, *MedidorUUID) (*MedidorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedidor not implemented")
}
func (UnimplementedMedidorServiceServer) GetMedidorInstalled(context.Context, *MedidorIsActive) (*MedidorIsActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedidorInstalled not implemented")
}
func (UnimplementedMedidorServiceServer) RecentInstallationMarca(context.Context, *MedidorMarca) (*MedidorGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecentInstallationMarca not implemented")
}
func (UnimplementedMedidorServiceServer) RecentInstallationSerial(context.Context, *MedidorSerial) (*MedidorGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecentInstallationSerial not implemented")
}
func (UnimplementedMedidorServiceServer) mustEmbedUnimplementedMedidorServiceServer() {}

// UnsafeMedidorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedidorServiceServer will
// result in compilation errors.
type UnsafeMedidorServiceServer interface {
	mustEmbedUnimplementedMedidorServiceServer()
}

func RegisterMedidorServiceServer(s grpc.ServiceRegistrar, srv MedidorServiceServer) {
	s.RegisterService(&MedidorService_ServiceDesc, srv)
}

func _MedidorService_WriteMedidor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).WriteMedidor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/WriteMedidor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).WriteMedidor(ctx, req.(*MedidorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedidorService_GetMedidor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).GetMedidor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/GetMedidor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).GetMedidor(ctx, req.(*MedidorUUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedidorService_UpdateMedidor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).UpdateMedidor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/UpdateMedidor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).UpdateMedidor(ctx, req.(*MedidorUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedidorService_DeleteMedidor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).DeleteMedidor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/DeleteMedidor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).DeleteMedidor(ctx, req.(*MedidorUUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedidorService_GetMedidorInstalled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorIsActive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).GetMedidorInstalled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/GetMedidorInstalled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).GetMedidorInstalled(ctx, req.(*MedidorIsActive))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedidorService_RecentInstallationMarca_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorMarca)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).RecentInstallationMarca(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/RecentInstallationMarca",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).RecentInstallationMarca(ctx, req.(*MedidorMarca))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedidorService_RecentInstallationSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedidorSerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedidorServiceServer).RecentInstallationSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medidorgRPC.MedidorService/RecentInstallationSerial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedidorServiceServer).RecentInstallationSerial(ctx, req.(*MedidorSerial))
	}
	return interceptor(ctx, in, info, handler)
}

// MedidorService_ServiceDesc is the grpc.ServiceDesc for MedidorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedidorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medidorgRPC.MedidorService",
	HandlerType: (*MedidorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteMedidor",
			Handler:    _MedidorService_WriteMedidor_Handler,
		},
		{
			MethodName: "GetMedidor",
			Handler:    _MedidorService_GetMedidor_Handler,
		},
		{
			MethodName: "UpdateMedidor",
			Handler:    _MedidorService_UpdateMedidor_Handler,
		},
		{
			MethodName: "DeleteMedidor",
			Handler:    _MedidorService_DeleteMedidor_Handler,
		},
		{
			MethodName: "GetMedidorInstalled",
			Handler:    _MedidorService_GetMedidorInstalled_Handler,
		},
		{
			MethodName: "RecentInstallationMarca",
			Handler:    _MedidorService_RecentInstallationMarca_Handler,
		},
		{
			MethodName: "RecentInstallationSerial",
			Handler:    _MedidorService_RecentInstallationSerial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medidorgRPC.proto",
}