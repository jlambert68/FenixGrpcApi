// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/fenixTestCaseBuilderServerGrpcApi.proto

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

var File_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto protoreflect.FileDescriptor

var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_rawDesc = []byte{
	0x0a, 0x64, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x1a, 0x76, 0x46, 0x65, 0x6e, 0x69, 0x78,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x7c, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x5f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f,
	0x42, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x75, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72,
	0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7c, 0x46, 0x65,
	0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x85, 0x01, 0x46, 0x65, 0x6e,
	0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x54, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x9f, 0x11, 0x0a, 0x26, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x76, 0x0a,
	0x0b, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x31, 0x2e, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a,
	0x32, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xdb, 0x01, 0x0a, 0x31, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x2e, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x66, 0x2e, 0x66, 0x65, 0x6e, 0x69,
	0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x12, 0xf2, 0x01, 0x0a, 0x37, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e,
	0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x3c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x77, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x6e, 0x65,
	0x64, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x90, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x3c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x37, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x6f, 0x6e,
	0x64, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0xd2, 0x01, 0x0a, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x6a,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64,
	0x50, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32, 0x2e, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41,
	0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x97, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x12, 0x3e, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3f, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x89, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x12, 0x3c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x32, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0xc2, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x2e, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x40, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0xdc, 0x01, 0x0a, 0x28,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x64, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x48,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x4d, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x0c, 0x53, 0x61,
	0x76, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x32, 0x2e, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0xb4, 0x01, 0x0a, 0x1f, 0x53, 0x61, 0x76, 0x65, 0x41, 0x6c, 0x6c,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xc6, 0x01, 0x0a, 0x28,
	0x53, 0x61, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x64, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_goTypes = []interface{}{
	(*EmptyParameter)(nil),            // 0: fenixTestCaseBuilderServerGrpcApi.EmptyParameter
	(*UserIdentificationMessage)(nil), // 1: fenixTestCaseBuilderServerGrpcApi.UserIdentificationMessage
	(*SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage)(nil),              // 2: fenixTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage
	(*ListTestCasesRequestMessage)(nil),                                                          // 3: fenixTestCaseBuilderServerGrpcApi.ListTestCasesRequestMessage
	(*GetTestCaseRequestMessage)(nil),                                                            // 4: fenixTestCaseBuilderServerGrpcApi.GetTestCaseRequestMessage
	(*ListAllTestInstructionsForSpecificTestCaseRequestMessage)(nil),                             // 5: fenixTestCaseBuilderServerGrpcApi.ListAllTestInstructionsForSpecificTestCaseRequestMessage
	(*ListAllTestInstructionContainersForSpecificTestCaseRequestMessage)(nil),                    // 6: fenixTestCaseBuilderServerGrpcApi.ListAllTestInstructionContainersForSpecificTestCaseRequestMessage
	(*TestCaseMessage)(nil),                                                                      // 7: fenixTestCaseBuilderServerGrpcApi.TestCaseMessage
	(*SaveAllTestInstructionsForSpecificTestCaseRequestMessage)(nil),                             // 8: fenixTestCaseBuilderServerGrpcApi.SaveAllTestInstructionsForSpecificTestCaseRequestMessage
	(*SaveAllTestInstructionContainersForSpecificTestCaseRequestMessage)(nil),                    // 9: fenixTestCaseBuilderServerGrpcApi.SaveAllTestInstructionContainersForSpecificTestCaseRequestMessage
	(*AckNackResponse)(nil),                                                                      // 10: fenixTestCaseBuilderServerGrpcApi.AckNackResponse
	(*AvailableTestInstructionsAndPreCreatedTestContainersResponseMessage)(nil),                  // 11: fenixTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestContainersResponseMessage
	(*AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage)(nil), // 12: fenixTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage
	(*ImmatureBondsMessage)(nil),                                                                 // 13: fenixTestCaseBuilderServerGrpcApi.ImmatureBondsMessage
	(*ListTestCasesResponseMessage)(nil),                                                         // 14: fenixTestCaseBuilderServerGrpcApi.ListTestCasesResponseMessage
	(*MatureTestInstructionsMessage)(nil),                                                        // 15: fenixTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage
	(*MatureTestInstructionContainerMessage)(nil),                                                // 16: fenixTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerMessage
}
var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_depIdxs = []int32{
	0,  // 0: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.AreYouAlive:input_type -> fenixTestCaseBuilderServerGrpcApi.EmptyParameter
	1,  // 1: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllAvailableTestInstructionsAndTestContainers:input_type -> fenixTestCaseBuilderServerGrpcApi.UserIdentificationMessage
	1,  // 2: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllAvailablePinnedTestInstructionsAndTestContainers:input_type -> fenixTestCaseBuilderServerGrpcApi.UserIdentificationMessage
	1,  // 3: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllAvailableBonds:input_type -> fenixTestCaseBuilderServerGrpcApi.UserIdentificationMessage
	2,  // 4: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveAllPinnedTestInstructionsAndTestContainers:input_type -> fenixTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage
	3,  // 5: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllTestCases:input_type -> fenixTestCaseBuilderServerGrpcApi.ListTestCasesRequestMessage
	4,  // 6: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.GetDetailedTestCase:input_type -> fenixTestCaseBuilderServerGrpcApi.GetTestCaseRequestMessage
	5,  // 7: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllTestCaseTestInstructions:input_type -> fenixTestCaseBuilderServerGrpcApi.ListAllTestInstructionsForSpecificTestCaseRequestMessage
	6,  // 8: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllTestCaseTestInstructionContainers:input_type -> fenixTestCaseBuilderServerGrpcApi.ListAllTestInstructionContainersForSpecificTestCaseRequestMessage
	7,  // 9: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveTestCase:input_type -> fenixTestCaseBuilderServerGrpcApi.TestCaseMessage
	8,  // 10: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveAllTestCaseTestInstructions:input_type -> fenixTestCaseBuilderServerGrpcApi.SaveAllTestInstructionsForSpecificTestCaseRequestMessage
	9,  // 11: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveAllTestCaseTestInstructionContainers:input_type -> fenixTestCaseBuilderServerGrpcApi.SaveAllTestInstructionContainersForSpecificTestCaseRequestMessage
	10, // 12: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.AreYouAlive:output_type -> fenixTestCaseBuilderServerGrpcApi.AckNackResponse
	11, // 13: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllAvailableTestInstructionsAndTestContainers:output_type -> fenixTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestContainersResponseMessage
	12, // 14: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllAvailablePinnedTestInstructionsAndTestContainers:output_type -> fenixTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage
	13, // 15: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllAvailableBonds:output_type -> fenixTestCaseBuilderServerGrpcApi.ImmatureBondsMessage
	10, // 16: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveAllPinnedTestInstructionsAndTestContainers:output_type -> fenixTestCaseBuilderServerGrpcApi.AckNackResponse
	14, // 17: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllTestCases:output_type -> fenixTestCaseBuilderServerGrpcApi.ListTestCasesResponseMessage
	7,  // 18: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.GetDetailedTestCase:output_type -> fenixTestCaseBuilderServerGrpcApi.TestCaseMessage
	15, // 19: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllTestCaseTestInstructions:output_type -> fenixTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage
	16, // 20: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.ListAllTestCaseTestInstructionContainers:output_type -> fenixTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerMessage
	10, // 21: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveTestCase:output_type -> fenixTestCaseBuilderServerGrpcApi.AckNackResponse
	10, // 22: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveAllTestCaseTestInstructions:output_type -> fenixTestCaseBuilderServerGrpcApi.AckNackResponse
	10, // 23: fenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcServices.SaveAllTestCaseTestInstructionContainers:output_type -> fenixTestCaseBuilderServerGrpcApi.AckNackResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_init()
}
func file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_init() {
	if File_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto != nil {
		return
	}
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_AvailableMessages_proto_init()
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_GeneralMessagesAndEnums_proto_init()
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_BondingMessages_proto_init()
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TestCaseMessages_proto_init()
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TestInstructionMessages_proto_init()
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TestInstructionContainerMessages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_goTypes,
		DependencyIndexes: file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_depIdxs,
	}.Build()
	File_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto = out.File
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_rawDesc = nil
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_goTypes = nil
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_proto_depIdxs = nil
}
