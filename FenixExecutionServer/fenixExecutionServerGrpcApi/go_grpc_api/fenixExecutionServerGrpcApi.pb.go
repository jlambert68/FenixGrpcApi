// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: FenixExecutionServer/fenixExecutionServerGrpcApi/fenixExecutionServerGrpcApi.proto

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

var File_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto protoreflect.FileDescriptor

var file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_rawDesc = []byte{
	0x0a, 0x52, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x1a, 0x6a, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41,
	0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5b, 0x46,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb3, 0x0a, 0x0a, 0x20, 0x46,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x6a, 0x0a, 0x0b, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x2b,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x2c, 0x2e, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x0d, 0x49,
	0x73, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x31, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x49, 0x73, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63,
	0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x9a, 0x01, 0x0a, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x68, 0x61, 0x74, 0x54, 0x68,
	0x65, 0x72, 0x65, 0x41, 0x72, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x73, 0x4f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x38, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x2c, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa1, 0x01, 0x0a,
	0x35, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x68, 0x61, 0x74, 0x54, 0x68, 0x65, 0x72, 0x65,
	0x41, 0x72, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x38, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41,
	0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x86, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x38, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x92, 0x01, 0x0a, 0x33, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x68, 0x61, 0x74,
	0x49, 0x73, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x42, 0x65, 0x53, 0x65, 0x6e,
	0x74, 0x12, 0x2b, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x2c,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b,
	0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa7,
	0x01, 0x0a, 0x2c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x47, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xaa, 0x01, 0x0a, 0x2b, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x49, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x79, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c,
	0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b,
	0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x42, 0x25, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0xaa, 0x02, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_goTypes = []interface{}{
	(*EmptyParameter)(nil),                               // 0: fenixExecutionServerGrpcApi.EmptyParameter
	(*IsWorkerAliveRequest)(nil),                         // 1: fenixExecutionServerGrpcApi.IsWorkerAliveRequest
	(*TestCaseExecutionsToProcess)(nil),                  // 2: fenixExecutionServerGrpcApi.TestCaseExecutionsToProcess
	(*ProcessingCapabilityMessage)(nil),                  // 3: fenixExecutionServerGrpcApi.ProcessingCapabilityMessage
	(*FinalTestInstructionExecutionResultMessage)(nil),   // 4: fenixExecutionServerGrpcApi.FinalTestInstructionExecutionResultMessage
	(*CurrentTestInstructionExecutionResultMessage)(nil), // 5: fenixExecutionServerGrpcApi.CurrentTestInstructionExecutionResultMessage
	(*LogPostsMessage)(nil),                              // 6: fenixExecutionServerGrpcApi.LogPostsMessage
	(*AckNackResponse)(nil),                              // 7: fenixExecutionServerGrpcApi.AckNackResponse
}
var file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_depIdxs = []int32{
	0, // 0: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.AreYouAlive:input_type -> fenixExecutionServerGrpcApi.EmptyParameter
	1, // 1: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.IsWorkerAlive:input_type -> fenixExecutionServerGrpcApi.IsWorkerAliveRequest
	2, // 2: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.InformThatThereAreNewTestCasesOnExecutionQueue:input_type -> fenixExecutionServerGrpcApi.TestCaseExecutionsToProcess
	2, // 3: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.InformThatThereAreNewTestInstructionsOnExecutionQueue:input_type -> fenixExecutionServerGrpcApi.TestCaseExecutionsToProcess
	3, // 4: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.ReportProcessingCapability:input_type -> fenixExecutionServerGrpcApi.ProcessingCapabilityMessage
	0, // 5: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.TriggerSendNewTestInstructionsThatIsWaitingToBeSent:input_type -> fenixExecutionServerGrpcApi.EmptyParameter
	4, // 6: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.ReportCompleteTestInstructionExecutionResult:input_type -> fenixExecutionServerGrpcApi.FinalTestInstructionExecutionResultMessage
	5, // 7: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.ReportCurrentTestInstructionExecutionResult:input_type -> fenixExecutionServerGrpcApi.CurrentTestInstructionExecutionResultMessage
	6, // 8: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.SendLogPostForExecution:input_type -> fenixExecutionServerGrpcApi.LogPostsMessage
	7, // 9: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.AreYouAlive:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 10: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.IsWorkerAlive:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 11: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.InformThatThereAreNewTestCasesOnExecutionQueue:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 12: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.InformThatThereAreNewTestInstructionsOnExecutionQueue:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 13: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.ReportProcessingCapability:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 14: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.TriggerSendNewTestInstructionsThatIsWaitingToBeSent:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 15: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.ReportCompleteTestInstructionExecutionResult:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 16: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.ReportCurrentTestInstructionExecutionResult:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	7, // 17: fenixExecutionServerGrpcApi.FenixExecutionServerGrpcServices.SendLogPostForExecution:output_type -> fenixExecutionServerGrpcApi.AckNackResponse
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_init()
}
func file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_init() {
	if File_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto != nil {
		return
	}
	file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_GeneralMessagesAndEnums_proto_init()
	file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_Messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_goTypes,
		DependencyIndexes: file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_depIdxs,
	}.Build()
	File_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto = out.File
	file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_rawDesc = nil
	file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_goTypes = nil
	file_FenixExecutionServer_fenixExecutionServerGrpcApi_fenixExecutionServerGrpcApi_proto_depIdxs = nil
}
