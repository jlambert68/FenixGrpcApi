// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: FenixUserGui/fenixUserGuiGrpcApi/fenixUserGuiGrpcApi.proto

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

var File_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto protoreflect.FileDescriptor

var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x2f, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x1a, 0x52, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd8, 0x02, 0x0a, 0x18, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x5a, 0x0a, 0x0b, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x12, 0x23, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b,
	0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c,
	0x0a, 0x1d, 0x47, 0x75, 0x69, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12,
	0x23, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72,
	0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x23,
	0x47, 0x75, 0x69, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x12, 0x23, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x41,
	0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_goTypes = []interface{}{
	(*EmptyParameter)(nil),  // 0: fenixUserGuiGrpcApi.EmptyParameter
	(*AckNackResponse)(nil), // 1: fenixUserGuiGrpcApi.AckNackResponse
}
var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_depIdxs = []int32{
	0, // 0: fenixUserGuiGrpcApi.FenixUserGuiGrpcServices.AreYouAlive:input_type -> fenixUserGuiGrpcApi.EmptyParameter
	0, // 1: fenixUserGuiGrpcApi.FenixUserGuiGrpcServices.GuiExecutionServerAreYouAlive:input_type -> fenixUserGuiGrpcApi.EmptyParameter
	0, // 2: fenixUserGuiGrpcApi.FenixUserGuiGrpcServices.GuiTestCaseBuilderServerAreYouAlive:input_type -> fenixUserGuiGrpcApi.EmptyParameter
	1, // 3: fenixUserGuiGrpcApi.FenixUserGuiGrpcServices.AreYouAlive:output_type -> fenixUserGuiGrpcApi.AckNackResponse
	1, // 4: fenixUserGuiGrpcApi.FenixUserGuiGrpcServices.GuiExecutionServerAreYouAlive:output_type -> fenixUserGuiGrpcApi.AckNackResponse
	1, // 5: fenixUserGuiGrpcApi.FenixUserGuiGrpcServices.GuiTestCaseBuilderServerAreYouAlive:output_type -> fenixUserGuiGrpcApi.AckNackResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_init() }
func file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_init() {
	if File_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto != nil {
		return
	}
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_goTypes,
		DependencyIndexes: file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_depIdxs,
	}.Build()
	File_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto = out.File
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_rawDesc = nil
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_goTypes = nil
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_proto_depIdxs = nil
}
