// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: FenixExecutionServer/fenixExecutionWorkerGrpcApi/fenixExecutionWorkerGrpcApi.proto

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

var File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto protoreflect.FileDescriptor

var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_rawDesc = []byte{
	0x0a, 0x52, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x1a, 0x6a, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41,
	0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5b, 0x46,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x71, 0x46, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x55, 0x73, 0x65, 0x64, 0x54, 0x6f, 0x67, 0x65, 0x74, 0x68, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc5, 0x08,
	0x0a, 0x20, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x6a, 0x0a, 0x0b, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x12, 0x2b, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x2c,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b,
	0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79,
	0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb5, 0x01, 0x0a, 0x1f, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xa2, 0x01, 0x0a, 0x25, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x12, 0x49, 0x2e, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa3, 0x01, 0x0a, 0x2c, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72,
	0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa2, 0x01, 0x0a,
	0x2b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e,
	0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x91, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x43, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x81, 0x08, 0x0a, 0x29, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x2b, 0x2e, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb0, 0x01, 0x0a, 0x35, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x47, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb1, 0x01, 0x0a, 0x32,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a,
	0x4a, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0xb8, 0x01, 0x0a, 0x38, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa8, 0x01, 0x0a, 0x30, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x49, 0x74, 0x49, 0x73, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x12, 0x34, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3c, 0x2e, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67,
	0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_goTypes = []interface{}{
	(*EmptyParameter)(nil),                                  // 0: fenixExecutionWorkerGrpcApi.EmptyParameter
	(*ProcessTestInstructionExecutionReveredRequest)(nil),   // 1: fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest
	(*ProcessTestInstructionExecutionPubSubRequest)(nil),    // 2: fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest
	(*TestInstructionExecutionRequestMessage)(nil),          // 3: fenixExecutionWorkerGrpcApi.TestInstructionExecutionRequestMessage
	(*FinalTestInstructionExecutionResultMessage)(nil),      // 4: fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
	(*ProcessTestInstructionExecutionReversedResponse)(nil), // 5: fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse
	(*ProcessTestInstructionExecutionResponse)(nil),         // 6: fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionResponse
	(*ConnectorIsReadyMessage)(nil),                         // 7: fenixExecutionWorkerGrpcApi.ConnectorIsReadyMessage
	(*AckNackResponse)(nil),                                 // 8: fenixExecutionWorkerGrpcApi.AckNackResponse
	(*ConnectorIsReadyResponseMessage)(nil),                 // 9: fenixExecutionWorkerGrpcApi.ConnectorIsReadyResponseMessage
}
var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_depIdxs = []int32{
	0,  // 0: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.AreYouAlive:input_type -> fenixExecutionWorkerGrpcApi.EmptyParameter
	0,  // 1: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ReportProcessingCapability:input_type -> fenixExecutionWorkerGrpcApi.EmptyParameter
	1,  // 2: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ProcessTestInstructionExecution:input_type -> fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest
	2,  // 3: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ProcessTestInstructionExecutionPubSub:input_type -> fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest
	3,  // 4: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ReportCompleteTestInstructionExecutionResult:input_type -> fenixExecutionWorkerGrpcApi.TestInstructionExecutionRequestMessage
	3,  // 5: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ReportCurrentTestInstructionExecutionResult:input_type -> fenixExecutionWorkerGrpcApi.TestInstructionExecutionRequestMessage
	3,  // 6: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.SendAllLogPostForExecution:input_type -> fenixExecutionWorkerGrpcApi.TestInstructionExecutionRequestMessage
	0,  // 7: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorAreYouAlive:input_type -> fenixExecutionWorkerGrpcApi.EmptyParameter
	4,  // 8: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorReportCompleteTestInstructionExecutionResult:input_type -> fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
	0,  // 9: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorRequestForProcessTestInstructionExecution:input_type -> fenixExecutionWorkerGrpcApi.EmptyParameter
	5,  // 10: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorProcessTestInstructionExecutionReversedResponse:input_type -> fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse
	6,  // 11: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorProcessTestInstructionExecutionResponse:input_type -> fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionResponse
	7,  // 12: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorInformsItIsAlive:input_type -> fenixExecutionWorkerGrpcApi.ConnectorIsReadyMessage
	8,  // 13: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.AreYouAlive:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 14: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ReportProcessingCapability:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	6,  // 15: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ProcessTestInstructionExecution:output_type -> fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionResponse
	8,  // 16: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ProcessTestInstructionExecutionPubSub:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 17: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ReportCompleteTestInstructionExecutionResult:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 18: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.ReportCurrentTestInstructionExecutionResult:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 19: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerGrpcServices.SendAllLogPostForExecution:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 20: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorAreYouAlive:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 21: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorReportCompleteTestInstructionExecutionResult:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	1,  // 22: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorRequestForProcessTestInstructionExecution:output_type -> fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReveredRequest
	8,  // 23: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorProcessTestInstructionExecutionReversedResponse:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	8,  // 24: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorProcessTestInstructionExecutionResponse:output_type -> fenixExecutionWorkerGrpcApi.AckNackResponse
	9,  // 25: fenixExecutionWorkerGrpcApi.FenixExecutionWorkerConnectorGrpcServices.ConnectorInformsItIsAlive:output_type -> fenixExecutionWorkerGrpcApi.ConnectorIsReadyResponseMessage
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_init()
}
func file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_init() {
	if File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto != nil {
		return
	}
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_init()
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_Messages_proto_init()
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_MessagesUsedTogetherWithPubSub_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_goTypes,
		DependencyIndexes: file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_depIdxs,
	}.Build()
	File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto = out.File
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_rawDesc = nil
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_goTypes = nil
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_proto_depIdxs = nil
}
