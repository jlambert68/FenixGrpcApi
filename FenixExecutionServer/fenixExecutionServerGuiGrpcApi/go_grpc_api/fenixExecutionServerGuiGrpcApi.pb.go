// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: FenixExecutionServer/fenixExecutionServerGuiGrpcApi/fenixExecutionServerGuiGrpcApi.proto

package go_grpc_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto protoreflect.FileDescriptor

var file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_rawDesc = []byte{
	0x0a, 0x58, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72,
	0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x1a, 0x70, 0x46, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x6e,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x78, 0x46, 0x65,
	0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x5f, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9e, 0x08, 0x0a, 0x23, 0x46, 0x65, 0x6e, 0x69, 0x78,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x70,
	0x0a, 0x0b, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x2e, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x2f, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41,
	0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xae, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x73, 0x4f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x44, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x49, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x49, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xa9, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x43, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73,
	0x55, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x43, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xc0, 0x01,
	0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4b, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xa5, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x41, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x42, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xbc, 0x01, 0x0a, 0x19, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x4e, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_goTypes = []interface{}{
	(*EmptyParameter)(nil),                                 // 0: fenixExecutionServerGuiGrpcApi.EmptyParameter
	(*ListTestCasesInExecutionQueueRequest)(nil),           // 1: fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest
	(*ListTestCasesUnderExecutionResponse)(nil),            // 2: fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse
	(*ListTestCasesWithFinishedExecutionsRequest)(nil),     // 3: fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest
	(*GetSingleTestCaseExecutionRequest)(nil),              // 4: fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest
	(*InitiateSingleTestCaseExecutionRequestMessage)(nil),  // 5: fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionRequestMessage
	(*AckNackResponse)(nil),                                // 6: fenixExecutionServerGuiGrpcApi.AckNackResponse
	(*ListTestCasesInExecutionQueueResponse)(nil),          // 7: fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse
	(*ListTestCasesWithFinishedExecutionsResponse)(nil),    // 8: fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse
	(*GetSingleTestCaseExecutionResponse)(nil),             // 9: fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse
	(*InitiateSingleTestCaseExecutionResponseMessage)(nil), // 10: fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage
}
var file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_depIdxs = []int32{
	0,  // 0: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.AreYouAlive:input_type -> fenixExecutionServerGuiGrpcApi.EmptyParameter
	1,  // 1: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.ListTestCasesOnExecutionQueue:input_type -> fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueRequest
	2,  // 2: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.ListTestCasesUnderExecution:input_type -> fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse
	3,  // 3: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.ListTestCasesWithFinishedExecutions:input_type -> fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsRequest
	4,  // 4: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.GetSingleTestCaseExecution:input_type -> fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest
	5,  // 5: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.InitiateTestCaseExecution:input_type -> fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionRequestMessage
	6,  // 6: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.AreYouAlive:output_type -> fenixExecutionServerGuiGrpcApi.AckNackResponse
	7,  // 7: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.ListTestCasesOnExecutionQueue:output_type -> fenixExecutionServerGuiGrpcApi.ListTestCasesInExecutionQueueResponse
	2,  // 8: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.ListTestCasesUnderExecution:output_type -> fenixExecutionServerGuiGrpcApi.ListTestCasesUnderExecutionResponse
	8,  // 9: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.ListTestCasesWithFinishedExecutions:output_type -> fenixExecutionServerGuiGrpcApi.ListTestCasesWithFinishedExecutionsResponse
	9,  // 10: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.GetSingleTestCaseExecution:output_type -> fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse
	10, // 11: fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServices.InitiateTestCaseExecution:output_type -> fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_init()
}
func file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_init() {
	if File_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto != nil {
		return
	}
	file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_GeneralMessagesAndEnums_proto_init()
	file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_SingleTestCaseExecutionMessages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_goTypes,
		DependencyIndexes: file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_depIdxs,
	}.Build()
	File_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto = out.File
	file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_rawDesc = nil
	file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_goTypes = nil
	file_FenixExecutionServer_fenixExecutionServerGuiGrpcApi_fenixExecutionServerGuiGrpcApi_proto_depIdxs = nil
}
