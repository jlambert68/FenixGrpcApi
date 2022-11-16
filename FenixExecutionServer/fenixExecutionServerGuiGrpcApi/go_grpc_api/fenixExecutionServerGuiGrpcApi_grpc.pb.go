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

// FenixExecutionServerGuiGrpcServicesForGuiClientClient is the client API for FenixExecutionServerGuiGrpcServicesForGuiClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FenixExecutionServerGuiGrpcServicesForGuiClientClient interface {
	//Anyone can check if Fenix Execution server is alive with this service
	AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// List TestCases that exists on Execution Queue, for specific Domains and DateTimes-span
	ListTestCasesOnExecutionQueue(ctx context.Context, in *ListTestCasesInExecutionQueueRequest, opts ...grpc.CallOption) (*ListTestCasesInExecutionQueueResponse, error)
	// List TestCases that is under execution , for specific Domains and DateTimes-span
	ListTestCasesUnderExecution(ctx context.Context, in *ListTestCasesUnderExecutionRequest, opts ...grpc.CallOption) (*ListTestCasesUnderExecutionResponse, error)
	// List TestCases that ire finished execution , for specific Domains and DateTimes-span
	ListTestCasesWithFinishedExecutions(ctx context.Context, in *ListTestCasesWithFinishedExecutionsRequest, opts ...grpc.CallOption) (*ListTestCasesWithFinishedExecutionsResponse, error)
	// Get a single TestCase Execution
	GetSingleTestCaseExecution(ctx context.Context, in *GetSingleTestCaseExecutionRequest, opts ...grpc.CallOption) (*GetSingleTestCaseExecutionResponse, error)
	// Initiate a single TestCase Execution with one specific TestDataSet
	InitiateTestCaseExecution(ctx context.Context, in *InitiateSingleTestCaseExecutionRequestMessage, opts ...grpc.CallOption) (*InitiateSingleTestCaseExecutionResponseMessage, error)
	// Execution TesterGui opens the gPRC-channel and messages are then streamed back to TestGui from GuiExecutionServer
	// Messages can be everything from execution status, information messages or Environment status
	SubscribeToMessageStream(ctx context.Context, in *UserAndApplicationRunTimeIdentificationMessage, opts ...grpc.CallOption) (FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStreamClient, error)
	// Call from TesterGui to GuiExecution regarding which messages to subscribe to
	SubscribeToMessages(ctx context.Context, in *SubscribeToMessagesRequest, opts ...grpc.CallOption) (*AckNackResponse, error)
}

type fenixExecutionServerGuiGrpcServicesForGuiClientClient struct {
	cc grpc.ClientConnInterface
}

func NewFenixExecutionServerGuiGrpcServicesForGuiClientClient(cc grpc.ClientConnInterface) FenixExecutionServerGuiGrpcServicesForGuiClientClient {
	return &fenixExecutionServerGuiGrpcServicesForGuiClientClient{cc}
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/AreYouAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) ListTestCasesOnExecutionQueue(ctx context.Context, in *ListTestCasesInExecutionQueueRequest, opts ...grpc.CallOption) (*ListTestCasesInExecutionQueueResponse, error) {
	out := new(ListTestCasesInExecutionQueueResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/ListTestCasesOnExecutionQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) ListTestCasesUnderExecution(ctx context.Context, in *ListTestCasesUnderExecutionRequest, opts ...grpc.CallOption) (*ListTestCasesUnderExecutionResponse, error) {
	out := new(ListTestCasesUnderExecutionResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/ListTestCasesUnderExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) ListTestCasesWithFinishedExecutions(ctx context.Context, in *ListTestCasesWithFinishedExecutionsRequest, opts ...grpc.CallOption) (*ListTestCasesWithFinishedExecutionsResponse, error) {
	out := new(ListTestCasesWithFinishedExecutionsResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/ListTestCasesWithFinishedExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) GetSingleTestCaseExecution(ctx context.Context, in *GetSingleTestCaseExecutionRequest, opts ...grpc.CallOption) (*GetSingleTestCaseExecutionResponse, error) {
	out := new(GetSingleTestCaseExecutionResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/GetSingleTestCaseExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) InitiateTestCaseExecution(ctx context.Context, in *InitiateSingleTestCaseExecutionRequestMessage, opts ...grpc.CallOption) (*InitiateSingleTestCaseExecutionResponseMessage, error) {
	out := new(InitiateSingleTestCaseExecutionResponseMessage)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/InitiateTestCaseExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) SubscribeToMessageStream(ctx context.Context, in *UserAndApplicationRunTimeIdentificationMessage, opts ...grpc.CallOption) (FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FenixExecutionServerGuiGrpcServicesForGuiClient_ServiceDesc.Streams[0], "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/SubscribeToMessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fenixExecutionServerGuiGrpcServicesForGuiClientSubscribeToMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStreamClient interface {
	Recv() (*SubscribeToMessagesStreamResponse, error)
	grpc.ClientStream
}

type fenixExecutionServerGuiGrpcServicesForGuiClientSubscribeToMessageStreamClient struct {
	grpc.ClientStream
}

func (x *fenixExecutionServerGuiGrpcServicesForGuiClientSubscribeToMessageStreamClient) Recv() (*SubscribeToMessagesStreamResponse, error) {
	m := new(SubscribeToMessagesStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fenixExecutionServerGuiGrpcServicesForGuiClientClient) SubscribeToMessages(ctx context.Context, in *SubscribeToMessagesRequest, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/SubscribeToMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FenixExecutionServerGuiGrpcServicesForGuiClientServer is the server API for FenixExecutionServerGuiGrpcServicesForGuiClient service.
// All implementations must embed UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer
// for forward compatibility
type FenixExecutionServerGuiGrpcServicesForGuiClientServer interface {
	//Anyone can check if Fenix Execution server is alive with this service
	AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// List TestCases that exists on Execution Queue, for specific Domains and DateTimes-span
	ListTestCasesOnExecutionQueue(context.Context, *ListTestCasesInExecutionQueueRequest) (*ListTestCasesInExecutionQueueResponse, error)
	// List TestCases that is under execution , for specific Domains and DateTimes-span
	ListTestCasesUnderExecution(context.Context, *ListTestCasesUnderExecutionRequest) (*ListTestCasesUnderExecutionResponse, error)
	// List TestCases that ire finished execution , for specific Domains and DateTimes-span
	ListTestCasesWithFinishedExecutions(context.Context, *ListTestCasesWithFinishedExecutionsRequest) (*ListTestCasesWithFinishedExecutionsResponse, error)
	// Get a single TestCase Execution
	GetSingleTestCaseExecution(context.Context, *GetSingleTestCaseExecutionRequest) (*GetSingleTestCaseExecutionResponse, error)
	// Initiate a single TestCase Execution with one specific TestDataSet
	InitiateTestCaseExecution(context.Context, *InitiateSingleTestCaseExecutionRequestMessage) (*InitiateSingleTestCaseExecutionResponseMessage, error)
	// Execution TesterGui opens the gPRC-channel and messages are then streamed back to TestGui from GuiExecutionServer
	// Messages can be everything from execution status, information messages or Environment status
	SubscribeToMessageStream(*UserAndApplicationRunTimeIdentificationMessage, FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStreamServer) error
	// Call from TesterGui to GuiExecution regarding which messages to subscribe to
	SubscribeToMessages(context.Context, *SubscribeToMessagesRequest) (*AckNackResponse, error)
	mustEmbedUnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer()
}

// UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer must be embedded to have forward compatible implementations.
type UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer struct {
}

func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AreYouAlive not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) ListTestCasesOnExecutionQueue(context.Context, *ListTestCasesInExecutionQueueRequest) (*ListTestCasesInExecutionQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestCasesOnExecutionQueue not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) ListTestCasesUnderExecution(context.Context, *ListTestCasesUnderExecutionRequest) (*ListTestCasesUnderExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestCasesUnderExecution not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) ListTestCasesWithFinishedExecutions(context.Context, *ListTestCasesWithFinishedExecutionsRequest) (*ListTestCasesWithFinishedExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestCasesWithFinishedExecutions not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) GetSingleTestCaseExecution(context.Context, *GetSingleTestCaseExecutionRequest) (*GetSingleTestCaseExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleTestCaseExecution not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) InitiateTestCaseExecution(context.Context, *InitiateSingleTestCaseExecutionRequestMessage) (*InitiateSingleTestCaseExecutionResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateTestCaseExecution not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) SubscribeToMessageStream(*UserAndApplicationRunTimeIdentificationMessage, FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToMessageStream not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) SubscribeToMessages(context.Context, *SubscribeToMessagesRequest) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToMessages not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer) mustEmbedUnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer() {
}

// UnsafeFenixExecutionServerGuiGrpcServicesForGuiClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FenixExecutionServerGuiGrpcServicesForGuiClientServer will
// result in compilation errors.
type UnsafeFenixExecutionServerGuiGrpcServicesForGuiClientServer interface {
	mustEmbedUnimplementedFenixExecutionServerGuiGrpcServicesForGuiClientServer()
}

func RegisterFenixExecutionServerGuiGrpcServicesForGuiClientServer(s grpc.ServiceRegistrar, srv FenixExecutionServerGuiGrpcServicesForGuiClientServer) {
	s.RegisterService(&FenixExecutionServerGuiGrpcServicesForGuiClient_ServiceDesc, srv)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_AreYouAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).AreYouAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/AreYouAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).AreYouAlive(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_ListTestCasesOnExecutionQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestCasesInExecutionQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).ListTestCasesOnExecutionQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/ListTestCasesOnExecutionQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).ListTestCasesOnExecutionQueue(ctx, req.(*ListTestCasesInExecutionQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_ListTestCasesUnderExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestCasesUnderExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).ListTestCasesUnderExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/ListTestCasesUnderExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).ListTestCasesUnderExecution(ctx, req.(*ListTestCasesUnderExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_ListTestCasesWithFinishedExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestCasesWithFinishedExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).ListTestCasesWithFinishedExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/ListTestCasesWithFinishedExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).ListTestCasesWithFinishedExecutions(ctx, req.(*ListTestCasesWithFinishedExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_GetSingleTestCaseExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleTestCaseExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).GetSingleTestCaseExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/GetSingleTestCaseExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).GetSingleTestCaseExecution(ctx, req.(*GetSingleTestCaseExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_InitiateTestCaseExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitiateSingleTestCaseExecutionRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).InitiateTestCaseExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/InitiateTestCaseExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).InitiateTestCaseExecution(ctx, req.(*InitiateSingleTestCaseExecutionRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserAndApplicationRunTimeIdentificationMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).SubscribeToMessageStream(m, &fenixExecutionServerGuiGrpcServicesForGuiClientSubscribeToMessageStreamServer{stream})
}

type FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStreamServer interface {
	Send(*SubscribeToMessagesStreamResponse) error
	grpc.ServerStream
}

type fenixExecutionServerGuiGrpcServicesForGuiClientSubscribeToMessageStreamServer struct {
	grpc.ServerStream
}

func (x *fenixExecutionServerGuiGrpcServicesForGuiClientSubscribeToMessageStreamServer) Send(m *SubscribeToMessagesStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeToMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).SubscribeToMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient/SubscribeToMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForGuiClientServer).SubscribeToMessages(ctx, req.(*SubscribeToMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FenixExecutionServerGuiGrpcServicesForGuiClient_ServiceDesc is the grpc.ServiceDesc for FenixExecutionServerGuiGrpcServicesForGuiClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FenixExecutionServerGuiGrpcServicesForGuiClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClient",
	HandlerType: (*FenixExecutionServerGuiGrpcServicesForGuiClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AreYouAlive",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_AreYouAlive_Handler,
		},
		{
			MethodName: "ListTestCasesOnExecutionQueue",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_ListTestCasesOnExecutionQueue_Handler,
		},
		{
			MethodName: "ListTestCasesUnderExecution",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_ListTestCasesUnderExecution_Handler,
		},
		{
			MethodName: "ListTestCasesWithFinishedExecutions",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_ListTestCasesWithFinishedExecutions_Handler,
		},
		{
			MethodName: "GetSingleTestCaseExecution",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_GetSingleTestCaseExecution_Handler,
		},
		{
			MethodName: "InitiateTestCaseExecution",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_InitiateTestCaseExecution_Handler,
		},
		{
			MethodName: "SubscribeToMessages",
			Handler:    _FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToMessageStream",
			Handler:       _FenixExecutionServerGuiGrpcServicesForGuiClient_SubscribeToMessageStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "FenixExecutionServer/fenixExecutionServerGuiGrpcApi/fenixExecutionServerGuiGrpcApi.proto",
}

// FenixExecutionServerGuiGrpcServicesForExecutionServerClient is the client API for FenixExecutionServerGuiGrpcServicesForExecutionServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FenixExecutionServerGuiGrpcServicesForExecutionServerClient interface {
	// ExecutionServer send over status for TestCaseExecutions and TestInstructionExecutions
	SendExecutionStatusTowardsGuiClient(ctx context.Context, in *TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
}

type fenixExecutionServerGuiGrpcServicesForExecutionServerClient struct {
	cc grpc.ClientConnInterface
}

func NewFenixExecutionServerGuiGrpcServicesForExecutionServerClient(cc grpc.ClientConnInterface) FenixExecutionServerGuiGrpcServicesForExecutionServerClient {
	return &fenixExecutionServerGuiGrpcServicesForExecutionServerClient{cc}
}

func (c *fenixExecutionServerGuiGrpcServicesForExecutionServerClient) SendExecutionStatusTowardsGuiClient(ctx context.Context, in *TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForExecutionServer/SendExecutionStatusTowardsGuiClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FenixExecutionServerGuiGrpcServicesForExecutionServerServer is the server API for FenixExecutionServerGuiGrpcServicesForExecutionServer service.
// All implementations must embed UnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer
// for forward compatibility
type FenixExecutionServerGuiGrpcServicesForExecutionServerServer interface {
	// ExecutionServer send over status for TestCaseExecutions and TestInstructionExecutions
	SendExecutionStatusTowardsGuiClient(context.Context, *TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage) (*AckNackResponse, error)
	mustEmbedUnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer()
}

// UnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer must be embedded to have forward compatible implementations.
type UnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer struct {
}

func (UnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer) SendExecutionStatusTowardsGuiClient(context.Context, *TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendExecutionStatusTowardsGuiClient not implemented")
}
func (UnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer) mustEmbedUnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer() {
}

// UnsafeFenixExecutionServerGuiGrpcServicesForExecutionServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FenixExecutionServerGuiGrpcServicesForExecutionServerServer will
// result in compilation errors.
type UnsafeFenixExecutionServerGuiGrpcServicesForExecutionServerServer interface {
	mustEmbedUnimplementedFenixExecutionServerGuiGrpcServicesForExecutionServerServer()
}

func RegisterFenixExecutionServerGuiGrpcServicesForExecutionServerServer(s grpc.ServiceRegistrar, srv FenixExecutionServerGuiGrpcServicesForExecutionServerServer) {
	s.RegisterService(&FenixExecutionServerGuiGrpcServicesForExecutionServer_ServiceDesc, srv)
}

func _FenixExecutionServerGuiGrpcServicesForExecutionServer_SendExecutionStatusTowardsGuiClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixExecutionServerGuiGrpcServicesForExecutionServerServer).SendExecutionStatusTowardsGuiClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForExecutionServer/SendExecutionStatusTowardsGuiClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixExecutionServerGuiGrpcServicesForExecutionServerServer).SendExecutionStatusTowardsGuiClient(ctx, req.(*TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FenixExecutionServerGuiGrpcServicesForExecutionServer_ServiceDesc is the grpc.ServiceDesc for FenixExecutionServerGuiGrpcServicesForExecutionServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FenixExecutionServerGuiGrpcServicesForExecutionServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForExecutionServer",
	HandlerType: (*FenixExecutionServerGuiGrpcServicesForExecutionServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendExecutionStatusTowardsGuiClient",
			Handler:    _FenixExecutionServerGuiGrpcServicesForExecutionServer_SendExecutionStatusTowardsGuiClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "FenixExecutionServer/fenixExecutionServerGuiGrpcApi/fenixExecutionServerGuiGrpcApi.proto",
}
