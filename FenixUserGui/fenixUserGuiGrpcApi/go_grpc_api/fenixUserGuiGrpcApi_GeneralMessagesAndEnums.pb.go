// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: FenixUserGui/fenixUserGuiGrpcApi/fenixUserGuiGrpcApi_GeneralMessagesAndEnums.proto

package go_grpc_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Used to have client and server in sync with using the same proto file version
type CurrentFenixUserGuiProtoFileVersionEnum int32

const (
	// Deprecated: Do not use.
	CurrentFenixUserGuiProtoFileVersionEnum_CurrentFenixUserGuiProtoFileVersionEnum_DEFAULT_NOT_SET CurrentFenixUserGuiProtoFileVersionEnum = 0
	// Deprecated: Do not use.
	CurrentFenixUserGuiProtoFileVersionEnum_VERSION_0_0 CurrentFenixUserGuiProtoFileVersionEnum = 1
	// Deprecated: Do not use.
	CurrentFenixUserGuiProtoFileVersionEnum_VERSION_0_1 CurrentFenixUserGuiProtoFileVersionEnum = 2
	CurrentFenixUserGuiProtoFileVersionEnum_VERSION_0_2 CurrentFenixUserGuiProtoFileVersionEnum = 3 // Current version
)

// Enum value maps for CurrentFenixUserGuiProtoFileVersionEnum.
var (
	CurrentFenixUserGuiProtoFileVersionEnum_name = map[int32]string{
		0: "CurrentFenixUserGuiProtoFileVersionEnum_DEFAULT_NOT_SET",
		1: "VERSION_0_0",
		2: "VERSION_0_1",
		3: "VERSION_0_2",
	}
	CurrentFenixUserGuiProtoFileVersionEnum_value = map[string]int32{
		"CurrentFenixUserGuiProtoFileVersionEnum_DEFAULT_NOT_SET": 0,
		"VERSION_0_0": 1,
		"VERSION_0_1": 2,
		"VERSION_0_2": 3,
	}
)

func (x CurrentFenixUserGuiProtoFileVersionEnum) Enum() *CurrentFenixUserGuiProtoFileVersionEnum {
	p := new(CurrentFenixUserGuiProtoFileVersionEnum)
	*p = x
	return p
}

func (x CurrentFenixUserGuiProtoFileVersionEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CurrentFenixUserGuiProtoFileVersionEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[0].Descriptor()
}

func (CurrentFenixUserGuiProtoFileVersionEnum) Type() protoreflect.EnumType {
	return &file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[0]
}

func (x CurrentFenixUserGuiProtoFileVersionEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CurrentFenixUserGuiProtoFileVersionEnum.Descriptor instead.
func (CurrentFenixUserGuiProtoFileVersionEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{0}
}

// Error codes - for now a test
type ErrorCodesEnum int32

const (
	ErrorCodesEnum_ErrorCodesEnum_DEFAULT_NOT_SET ErrorCodesEnum = 0
	ErrorCodesEnum_OK                             ErrorCodesEnum = 1
	ErrorCodesEnum_ERROR_UNKNOWN_CALLER           ErrorCodesEnum = 2
	ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION ErrorCodesEnum = 3
	ErrorCodesEnum_ERROR_UNSPECIFIED              ErrorCodesEnum = 4
)

// Enum value maps for ErrorCodesEnum.
var (
	ErrorCodesEnum_name = map[int32]string{
		0: "ErrorCodesEnum_DEFAULT_NOT_SET",
		1: "OK",
		2: "ERROR_UNKNOWN_CALLER",
		3: "ERROR_WRONG_PROTO_FILE_VERSION",
		4: "ERROR_UNSPECIFIED",
	}
	ErrorCodesEnum_value = map[string]int32{
		"ErrorCodesEnum_DEFAULT_NOT_SET": 0,
		"OK":                             1,
		"ERROR_UNKNOWN_CALLER":           2,
		"ERROR_WRONG_PROTO_FILE_VERSION": 3,
		"ERROR_UNSPECIFIED":              4,
	}
)

func (x ErrorCodesEnum) Enum() *ErrorCodesEnum {
	p := new(ErrorCodesEnum)
	*p = x
	return p
}

func (x ErrorCodesEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCodesEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[1].Descriptor()
}

func (ErrorCodesEnum) Type() protoreflect.EnumType {
	return &file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[1]
}

func (x ErrorCodesEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCodesEnum.Descriptor instead.
func (ErrorCodesEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{1}
}

// Parameter used for Empty inputs, only containing current proto-file version
type EmptyParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoFileVersionUsedByCaller CurrentFenixUserGuiProtoFileVersionEnum `protobuf:"varint,1,opt,name=ProtoFileVersionUsedByCaller,proto3,enum=fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum" json:"ProtoFileVersionUsedByCaller,omitempty"` // The latest proto file version number to be used
}

func (x *EmptyParameter) Reset() {
	*x = EmptyParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyParameter) ProtoMessage() {}

func (x *EmptyParameter) ProtoReflect() protoreflect.Message {
	mi := &file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyParameter.ProtoReflect.Descriptor instead.
func (*EmptyParameter) Descriptor() ([]byte, []int) {
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{0}
}

func (x *EmptyParameter) GetProtoFileVersionUsedByCaller() CurrentFenixUserGuiProtoFileVersionEnum {
	if x != nil {
		return x.ProtoFileVersionUsedByCaller
	}
	return CurrentFenixUserGuiProtoFileVersionEnum_CurrentFenixUserGuiProtoFileVersionEnum_DEFAULT_NOT_SET
}

// Ack/Nack- Response message with comment
type AckNackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckNack                            bool                                    `protobuf:"varint,1,opt,name=AckNack,proto3" json:"AckNack,omitempty"`                                                                                                                        // True=Ack, False=Nack
	Comments                           string                                  `protobuf:"bytes,2,opt,name=Comments,proto3" json:"Comments,omitempty"`                                                                                                                       //Comments if needed
	ErrorCodes                         []ErrorCodesEnum                        `protobuf:"varint,3,rep,packed,name=ErrorCodes,proto3,enum=fenixUserGuiGrpcApi.ErrorCodesEnum" json:"ErrorCodes,omitempty"`                                                                   // List of Error codes
	ProtoFileVersionUsedByFenixUserGui CurrentFenixUserGuiProtoFileVersionEnum `protobuf:"varint,4,opt,name=ProtoFileVersionUsedByFenixUserGui,proto3,enum=fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum" json:"ProtoFileVersionUsedByFenixUserGui,omitempty"` // The latest proto file version number to be used
}

func (x *AckNackResponse) Reset() {
	*x = AckNackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckNackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckNackResponse) ProtoMessage() {}

func (x *AckNackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckNackResponse.ProtoReflect.Descriptor instead.
func (*AckNackResponse) Descriptor() ([]byte, []int) {
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{1}
}

func (x *AckNackResponse) GetAckNack() bool {
	if x != nil {
		return x.AckNack
	}
	return false
}

func (x *AckNackResponse) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

func (x *AckNackResponse) GetErrorCodes() []ErrorCodesEnum {
	if x != nil {
		return x.ErrorCodes
	}
	return nil
}

func (x *AckNackResponse) GetProtoFileVersionUsedByFenixUserGui() CurrentFenixUserGuiProtoFileVersionEnum {
	if x != nil {
		return x.ProtoFileVersionUsedByFenixUserGui
	}
	return CurrentFenixUserGuiProtoFileVersionEnum_CurrentFenixUserGuiProtoFileVersionEnum_DEFAULT_NOT_SET
}

var File_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto protoreflect.FileDescriptor

var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDesc = []byte{
	0x0a, 0x52, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x2f, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x22, 0x93, 0x01, 0x0a, 0x0e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a,
	0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x22,
	0x9b, 0x02, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x8c,
	0x01, 0x0a, 0x22, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x75, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x75, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x22, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42,
	0x79, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x2a, 0xa5, 0x01,
	0x0a, 0x27, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x75, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x3f, 0x0a, 0x37, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x13, 0x0a, 0x0b, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x30, 0x5f, 0x30, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x01, 0x12,
	0x13, 0x0a, 0x0b, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x30, 0x5f, 0x31, 0x10, 0x02,
	0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x30, 0x5f, 0x32, 0x10, 0x03, 0x2a, 0x91, 0x01, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x02, 0x12, 0x22,
	0x0a, 0x1e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x04, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67,
	0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescOnce sync.Once
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescData = file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDesc
)

func file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP() []byte {
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescOnce.Do(func() {
		file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescData = protoimpl.X.CompressGZIP(file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescData)
	})
	return file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDescData
}

var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_goTypes = []interface{}{
	(CurrentFenixUserGuiProtoFileVersionEnum)(0), // 0: fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum
	(ErrorCodesEnum)(0),                          // 1: fenixUserGuiGrpcApi.ErrorCodesEnum
	(*EmptyParameter)(nil),                       // 2: fenixUserGuiGrpcApi.EmptyParameter
	(*AckNackResponse)(nil),                      // 3: fenixUserGuiGrpcApi.AckNackResponse
}
var file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_depIdxs = []int32{
	0, // 0: fenixUserGuiGrpcApi.EmptyParameter.ProtoFileVersionUsedByCaller:type_name -> fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum
	1, // 1: fenixUserGuiGrpcApi.AckNackResponse.ErrorCodes:type_name -> fenixUserGuiGrpcApi.ErrorCodesEnum
	0, // 2: fenixUserGuiGrpcApi.AckNackResponse.ProtoFileVersionUsedByFenixUserGui:type_name -> fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_init()
}
func file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_init() {
	if File_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyParameter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckNackResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_goTypes,
		DependencyIndexes: file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_depIdxs,
		EnumInfos:         file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_enumTypes,
		MessageInfos:      file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_msgTypes,
	}.Build()
	File_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto = out.File
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_rawDesc = nil
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_goTypes = nil
	file_FenixUserGui_fenixUserGuiGrpcApi_fenixUserGuiGrpcApi_GeneralMessagesAndEnums_proto_depIdxs = nil
}
