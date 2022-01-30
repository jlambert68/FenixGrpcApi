// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_grpc_api

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

// FenixTestDataGrpcServicesClient is the client API for FenixTestDataGrpcServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FenixTestDataGrpcServicesClient interface {
	//Fenix client can check if Fenix Testdata sync server is alive with this service
	AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix client can register itself with the Fenix Testdata sync server
	RegisterTestDataClient(ctx context.Context, in *TestDataClientInformationMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix client can send TestData MerkleHash to Fenix Testdata sync server with this service
	SendMerkleHash(ctx context.Context, in *MerkleHashMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
	SendMerkleTree(ctx context.Context, in *MerkleTreeMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix client can send TestDataHeaderHash to Fenix Testdata sync server with this service
	SendTestDataHeaderHash(ctx context.Context, in *TestDataHeaderHashMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
	SendTestDataHeaders(ctx context.Context, in *TestDataHeadersMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix client can send TestData rows to Fenix Testdata sync server with this service
	SendTestDataRows(ctx context.Context, in *TestdataRowsMessages, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Retry to allow incoming gRPC calls and process outgoing calls
	AllowOrDisallowIncomingAndOutgoingMessages(ctx context.Context, in *AllowOrDisallowIncomingAndOutgoingMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Restart Fenix TestData Server processes
	RestartFenixServerProcesses(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
}

type fenixTestDataGrpcServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewFenixTestDataGrpcServicesClient(cc grpc.ClientConnInterface) FenixTestDataGrpcServicesClient {
	return &fenixTestDataGrpcServicesClient{cc}
}

func (c *fenixTestDataGrpcServicesClient) AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/AreYouAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) RegisterTestDataClient(ctx context.Context, in *TestDataClientInformationMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/RegisterTestDataClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) SendMerkleHash(ctx context.Context, in *MerkleHashMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendMerkleHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) SendMerkleTree(ctx context.Context, in *MerkleTreeMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendMerkleTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) SendTestDataHeaderHash(ctx context.Context, in *TestDataHeaderHashMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendTestDataHeaderHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) SendTestDataHeaders(ctx context.Context, in *TestDataHeadersMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendTestDataHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) SendTestDataRows(ctx context.Context, in *TestdataRowsMessages, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendTestDataRows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) AllowOrDisallowIncomingAndOutgoingMessages(ctx context.Context, in *AllowOrDisallowIncomingAndOutgoingMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/AllowOrDisallowIncomingAndOutgoingMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixTestDataGrpcServicesClient) RestartFenixServerProcesses(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/RestartFenixServerProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FenixTestDataGrpcServicesServer is the server API for FenixTestDataGrpcServices service.
// All implementations must embed UnimplementedFenixTestDataGrpcServicesServer
// for forward compatibility
type FenixTestDataGrpcServicesServer interface {
	//Fenix client can check if Fenix Testdata sync server is alive with this service
	AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// Fenix client can register itself with the Fenix Testdata sync server
	RegisterTestDataClient(context.Context, *TestDataClientInformationMessage) (*AckNackResponse, error)
	// Fenix client can send TestData MerkleHash to Fenix Testdata sync server with this service
	SendMerkleHash(context.Context, *MerkleHashMessage) (*AckNackResponse, error)
	// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
	SendMerkleTree(context.Context, *MerkleTreeMessage) (*AckNackResponse, error)
	// Fenix client can send TestDataHeaderHash to Fenix Testdata sync server with this service
	SendTestDataHeaderHash(context.Context, *TestDataHeaderHashMessage) (*AckNackResponse, error)
	// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
	SendTestDataHeaders(context.Context, *TestDataHeadersMessage) (*AckNackResponse, error)
	// Fenix client can send TestData rows to Fenix Testdata sync server with this service
	SendTestDataRows(context.Context, *TestdataRowsMessages) (*AckNackResponse, error)
	// Retry to allow incoming gRPC calls and process outgoing calls
	AllowOrDisallowIncomingAndOutgoingMessages(context.Context, *AllowOrDisallowIncomingAndOutgoingMessage) (*AckNackResponse, error)
	// Restart Fenix TestData Server processes
	RestartFenixServerProcesses(context.Context, *EmptyParameter) (*AckNackResponse, error)
	mustEmbedUnimplementedFenixTestDataGrpcServicesServer()
}

// UnimplementedFenixTestDataGrpcServicesServer must be embedded to have forward compatible implementations.
type UnimplementedFenixTestDataGrpcServicesServer struct {
}

func (UnimplementedFenixTestDataGrpcServicesServer) AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AreYouAlive not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) RegisterTestDataClient(context.Context, *TestDataClientInformationMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTestDataClient not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) SendMerkleHash(context.Context, *MerkleHashMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMerkleHash not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) SendMerkleTree(context.Context, *MerkleTreeMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMerkleTree not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) SendTestDataHeaderHash(context.Context, *TestDataHeaderHashMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTestDataHeaderHash not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) SendTestDataHeaders(context.Context, *TestDataHeadersMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTestDataHeaders not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) SendTestDataRows(context.Context, *TestdataRowsMessages) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTestDataRows not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) AllowOrDisallowIncomingAndOutgoingMessages(context.Context, *AllowOrDisallowIncomingAndOutgoingMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowOrDisallowIncomingAndOutgoingMessages not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) RestartFenixServerProcesses(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartFenixServerProcesses not implemented")
}
func (UnimplementedFenixTestDataGrpcServicesServer) mustEmbedUnimplementedFenixTestDataGrpcServicesServer() {
}

// UnsafeFenixTestDataGrpcServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FenixTestDataGrpcServicesServer will
// result in compilation errors.
type UnsafeFenixTestDataGrpcServicesServer interface {
	mustEmbedUnimplementedFenixTestDataGrpcServicesServer()
}

func RegisterFenixTestDataGrpcServicesServer(s grpc.ServiceRegistrar, srv FenixTestDataGrpcServicesServer) {
	s.RegisterService(&FenixTestDataGrpcServices_ServiceDesc, srv)
}

func _FenixTestDataGrpcServices_AreYouAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).AreYouAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/AreYouAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).AreYouAlive(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_RegisterTestDataClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestDataClientInformationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).RegisterTestDataClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/RegisterTestDataClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).RegisterTestDataClient(ctx, req.(*TestDataClientInformationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_SendMerkleHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerkleHashMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).SendMerkleHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendMerkleHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).SendMerkleHash(ctx, req.(*MerkleHashMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_SendMerkleTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerkleTreeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).SendMerkleTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendMerkleTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).SendMerkleTree(ctx, req.(*MerkleTreeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_SendTestDataHeaderHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestDataHeaderHashMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).SendTestDataHeaderHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendTestDataHeaderHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).SendTestDataHeaderHash(ctx, req.(*TestDataHeaderHashMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_SendTestDataHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestDataHeadersMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).SendTestDataHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendTestDataHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).SendTestDataHeaders(ctx, req.(*TestDataHeadersMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_SendTestDataRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestdataRowsMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).SendTestDataRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/SendTestDataRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).SendTestDataRows(ctx, req.(*TestdataRowsMessages))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_AllowOrDisallowIncomingAndOutgoingMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowOrDisallowIncomingAndOutgoingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).AllowOrDisallowIncomingAndOutgoingMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/AllowOrDisallowIncomingAndOutgoingMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).AllowOrDisallowIncomingAndOutgoingMessages(ctx, req.(*AllowOrDisallowIncomingAndOutgoingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixTestDataGrpcServices_RestartFenixServerProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixTestDataGrpcServicesServer).RestartFenixServerProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices/RestartFenixServerProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixTestDataGrpcServicesServer).RestartFenixServerProcesses(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

// FenixTestDataGrpcServices_ServiceDesc is the grpc.ServiceDesc for FenixTestDataGrpcServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FenixTestDataGrpcServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices",
	HandlerType: (*FenixTestDataGrpcServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AreYouAlive",
			Handler:    _FenixTestDataGrpcServices_AreYouAlive_Handler,
		},
		{
			MethodName: "RegisterTestDataClient",
			Handler:    _FenixTestDataGrpcServices_RegisterTestDataClient_Handler,
		},
		{
			MethodName: "SendMerkleHash",
			Handler:    _FenixTestDataGrpcServices_SendMerkleHash_Handler,
		},
		{
			MethodName: "SendMerkleTree",
			Handler:    _FenixTestDataGrpcServices_SendMerkleTree_Handler,
		},
		{
			MethodName: "SendTestDataHeaderHash",
			Handler:    _FenixTestDataGrpcServices_SendTestDataHeaderHash_Handler,
		},
		{
			MethodName: "SendTestDataHeaders",
			Handler:    _FenixTestDataGrpcServices_SendTestDataHeaders_Handler,
		},
		{
			MethodName: "SendTestDataRows",
			Handler:    _FenixTestDataGrpcServices_SendTestDataRows_Handler,
		},
		{
			MethodName: "AllowOrDisallowIncomingAndOutgoingMessages",
			Handler:    _FenixTestDataGrpcServices_AllowOrDisallowIncomingAndOutgoingMessages_Handler,
		},
		{
			MethodName: "RestartFenixServerProcesses",
			Handler:    _FenixTestDataGrpcServices_RestartFenixServerProcesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fenixTestDataSyncServerGrpcApi.proto",
}
