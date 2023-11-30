// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/patient.proto

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

// PatientServiceClient is the client API for PatientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientServiceClient interface {
	Create(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*Patient, error)
	FindById(ctx context.Context, in *FindPatientByIDRequest, opts ...grpc.CallOption) (*Patient, error)
	Update(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*Blank, error)
	Delete(ctx context.Context, in *DeletePatientRequest, opts ...grpc.CallOption) (*Blank, error)
	NewSession(ctx context.Context, opts ...grpc.CallOption) (PatientService_NewSessionClient, error)
	Logout(ctx context.Context, in *LogoutPatientRequest, opts ...grpc.CallOption) (*Blank, error)
	GetHelp(ctx context.Context, in *GetHelpRequest, opts ...grpc.CallOption) (*Blank, error)
}

type patientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientServiceClient(cc grpc.ClientConnInterface) PatientServiceClient {
	return &patientServiceClient{cc}
}

func (c *patientServiceClient) Create(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*Patient, error) {
	out := new(Patient)
	err := c.cc.Invoke(ctx, "/pb.PatientService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) FindById(ctx context.Context, in *FindPatientByIDRequest, opts ...grpc.CallOption) (*Patient, error) {
	out := new(Patient)
	err := c.cc.Invoke(ctx, "/pb.PatientService/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) Update(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/pb.PatientService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) Delete(ctx context.Context, in *DeletePatientRequest, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/pb.PatientService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) NewSession(ctx context.Context, opts ...grpc.CallOption) (PatientService_NewSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &PatientService_ServiceDesc.Streams[0], "/pb.PatientService/NewSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &patientServiceNewSessionClient{stream}
	return x, nil
}

type PatientService_NewSessionClient interface {
	Send(*NewPatientSessionRequest) error
	Recv() (*Session, error)
	grpc.ClientStream
}

type patientServiceNewSessionClient struct {
	grpc.ClientStream
}

func (x *patientServiceNewSessionClient) Send(m *NewPatientSessionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *patientServiceNewSessionClient) Recv() (*Session, error) {
	m := new(Session)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *patientServiceClient) Logout(ctx context.Context, in *LogoutPatientRequest, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/pb.PatientService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) GetHelp(ctx context.Context, in *GetHelpRequest, opts ...grpc.CallOption) (*Blank, error) {
	out := new(Blank)
	err := c.cc.Invoke(ctx, "/pb.PatientService/GetHelp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServiceServer is the server API for PatientService service.
// All implementations must embed UnimplementedPatientServiceServer
// for forward compatibility
type PatientServiceServer interface {
	Create(context.Context, *CreatePatientRequest) (*Patient, error)
	FindById(context.Context, *FindPatientByIDRequest) (*Patient, error)
	Update(context.Context, *UpdatePatientRequest) (*Blank, error)
	Delete(context.Context, *DeletePatientRequest) (*Blank, error)
	NewSession(PatientService_NewSessionServer) error
	Logout(context.Context, *LogoutPatientRequest) (*Blank, error)
	GetHelp(context.Context, *GetHelpRequest) (*Blank, error)
	mustEmbedUnimplementedPatientServiceServer()
}

// UnimplementedPatientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServiceServer struct {
}

func (UnimplementedPatientServiceServer) Create(context.Context, *CreatePatientRequest) (*Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPatientServiceServer) FindById(context.Context, *FindPatientByIDRequest) (*Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedPatientServiceServer) Update(context.Context, *UpdatePatientRequest) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPatientServiceServer) Delete(context.Context, *DeletePatientRequest) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPatientServiceServer) NewSession(PatientService_NewSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method NewSession not implemented")
}
func (UnimplementedPatientServiceServer) Logout(context.Context, *LogoutPatientRequest) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedPatientServiceServer) GetHelp(context.Context, *GetHelpRequest) (*Blank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelp not implemented")
}
func (UnimplementedPatientServiceServer) mustEmbedUnimplementedPatientServiceServer() {}

// UnsafePatientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientServiceServer will
// result in compilation errors.
type UnsafePatientServiceServer interface {
	mustEmbedUnimplementedPatientServiceServer()
}

func RegisterPatientServiceServer(s grpc.ServiceRegistrar, srv PatientServiceServer) {
	s.RegisterService(&PatientService_ServiceDesc, srv)
}

func _PatientService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PatientService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).Create(ctx, req.(*CreatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPatientByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PatientService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).FindById(ctx, req.(*FindPatientByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PatientService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).Update(ctx, req.(*UpdatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PatientService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).Delete(ctx, req.(*DeletePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_NewSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PatientServiceServer).NewSession(&patientServiceNewSessionServer{stream})
}

type PatientService_NewSessionServer interface {
	Send(*Session) error
	Recv() (*NewPatientSessionRequest, error)
	grpc.ServerStream
}

type patientServiceNewSessionServer struct {
	grpc.ServerStream
}

func (x *patientServiceNewSessionServer) Send(m *Session) error {
	return x.ServerStream.SendMsg(m)
}

func (x *patientServiceNewSessionServer) Recv() (*NewPatientSessionRequest, error) {
	m := new(NewPatientSessionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PatientService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PatientService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).Logout(ctx, req.(*LogoutPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_GetHelp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHelpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).GetHelp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PatientService/GetHelp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).GetHelp(ctx, req.(*GetHelpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientService_ServiceDesc is the grpc.ServiceDesc for PatientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PatientService",
	HandlerType: (*PatientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PatientService_Create_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _PatientService_FindById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PatientService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PatientService_Delete_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _PatientService_Logout_Handler,
		},
		{
			MethodName: "GetHelp",
			Handler:    _PatientService_GetHelp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewSession",
			Handler:       _PatientService_NewSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/patient.proto",
}