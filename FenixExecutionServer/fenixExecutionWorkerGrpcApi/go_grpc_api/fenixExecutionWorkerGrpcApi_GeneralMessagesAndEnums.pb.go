// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: FenixExecutionServer/fenixExecutionWorkerGrpcApi/fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums.proto

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
type CurrentFenixExecutionWorkerProtoFileVersionEnum int32

const (
	// Deprecated: Do not use.
	CurrentFenixExecutionWorkerProtoFileVersionEnum_VERSION_0_0 CurrentFenixExecutionWorkerProtoFileVersionEnum = 0
	CurrentFenixExecutionWorkerProtoFileVersionEnum_VERSION_0_1 CurrentFenixExecutionWorkerProtoFileVersionEnum = 1 // Current version
)

// Enum value maps for CurrentFenixExecutionWorkerProtoFileVersionEnum.
var (
	CurrentFenixExecutionWorkerProtoFileVersionEnum_name = map[int32]string{
		0: "VERSION_0_0",
		1: "VERSION_0_1",
	}
	CurrentFenixExecutionWorkerProtoFileVersionEnum_value = map[string]int32{
		"VERSION_0_0": 0,
		"VERSION_0_1": 1,
	}
)

func (x CurrentFenixExecutionWorkerProtoFileVersionEnum) Enum() *CurrentFenixExecutionWorkerProtoFileVersionEnum {
	p := new(CurrentFenixExecutionWorkerProtoFileVersionEnum)
	*p = x
	return p
}

func (x CurrentFenixExecutionWorkerProtoFileVersionEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CurrentFenixExecutionWorkerProtoFileVersionEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[0].Descriptor()
}

func (CurrentFenixExecutionWorkerProtoFileVersionEnum) Type() protoreflect.EnumType {
	return &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[0]
}

func (x CurrentFenixExecutionWorkerProtoFileVersionEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CurrentFenixExecutionWorkerProtoFileVersionEnum.Descriptor instead.
func (CurrentFenixExecutionWorkerProtoFileVersionEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{0}
}

// Error codes - for now a test
type ErrorCodesEnum int32

const (
	ErrorCodesEnum_OK                                ErrorCodesEnum = 0
	ErrorCodesEnum_ERROR_UNKNOWN_CALLER              ErrorCodesEnum = 1
	ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION    ErrorCodesEnum = 2
	ErrorCodesEnum_ERROR_UNSPECIFIED                 ErrorCodesEnum = 3
	ErrorCodesEnum_ERROR_DATABASE_PROBLEM            ErrorCodesEnum = 4
	ErrorCodesEnum_ERROR_EXECUTION_CONNECTOR_PROBLEM ErrorCodesEnum = 5
)

// Enum value maps for ErrorCodesEnum.
var (
	ErrorCodesEnum_name = map[int32]string{
		0: "OK",
		1: "ERROR_UNKNOWN_CALLER",
		2: "ERROR_WRONG_PROTO_FILE_VERSION",
		3: "ERROR_UNSPECIFIED",
		4: "ERROR_DATABASE_PROBLEM",
		5: "ERROR_EXECUTION_CONNECTOR_PROBLEM",
	}
	ErrorCodesEnum_value = map[string]int32{
		"OK":                                0,
		"ERROR_UNKNOWN_CALLER":              1,
		"ERROR_WRONG_PROTO_FILE_VERSION":    2,
		"ERROR_UNSPECIFIED":                 3,
		"ERROR_DATABASE_PROBLEM":            4,
		"ERROR_EXECUTION_CONNECTOR_PROBLEM": 5,
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
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[1].Descriptor()
}

func (ErrorCodesEnum) Type() protoreflect.EnumType {
	return &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[1]
}

func (x ErrorCodesEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCodesEnum.Descriptor instead.
func (ErrorCodesEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{1}
}

// Execution Enum for a TestInstruction Execution
type TestInstructionExecutionStatusEnum int32

const (
	TestInstructionExecutionStatusEnum_TIE_INITIATED               TestInstructionExecutionStatusEnum = 0 // All set up for execution, but has not been triggered to start execution
	TestInstructionExecutionStatusEnum_TIE_EXECUTING               TestInstructionExecutionStatusEnum = 1 // TestCase is execution
	TestInstructionExecutionStatusEnum_TIE_CONTROLLED_INTERRUPTION TestInstructionExecutionStatusEnum = 2 // Interrupted by in a controlled way
	TestInstructionExecutionStatusEnum_TIE_FINISHED_OK             TestInstructionExecutionStatusEnum = 3 // Finnish as expected to TestCase definition
	TestInstructionExecutionStatusEnum_TIE_FINISHED_NOT_OK         TestInstructionExecutionStatusEnum = 4 // Finnish with errors in validations
	TestInstructionExecutionStatusEnum_TIE_UNEXPECTED_INTERRUPTION TestInstructionExecutionStatusEnum = 5 // The TestCase stopped executed in an unexpected way
)

// Enum value maps for TestInstructionExecutionStatusEnum.
var (
	TestInstructionExecutionStatusEnum_name = map[int32]string{
		0: "TIE_INITIATED",
		1: "TIE_EXECUTING",
		2: "TIE_CONTROLLED_INTERRUPTION",
		3: "TIE_FINISHED_OK",
		4: "TIE_FINISHED_NOT_OK",
		5: "TIE_UNEXPECTED_INTERRUPTION",
	}
	TestInstructionExecutionStatusEnum_value = map[string]int32{
		"TIE_INITIATED":               0,
		"TIE_EXECUTING":               1,
		"TIE_CONTROLLED_INTERRUPTION": 2,
		"TIE_FINISHED_OK":             3,
		"TIE_FINISHED_NOT_OK":         4,
		"TIE_UNEXPECTED_INTERRUPTION": 5,
	}
)

func (x TestInstructionExecutionStatusEnum) Enum() *TestInstructionExecutionStatusEnum {
	p := new(TestInstructionExecutionStatusEnum)
	*p = x
	return p
}

func (x TestInstructionExecutionStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestInstructionExecutionStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[2].Descriptor()
}

func (TestInstructionExecutionStatusEnum) Type() protoreflect.EnumType {
	return &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[2]
}

func (x TestInstructionExecutionStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestInstructionExecutionStatusEnum.Descriptor instead.
func (TestInstructionExecutionStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{2}
}

// How a Client is able to process requests in serial or parallell, regarding TestInstructions, TestCases and/or TestSuites
type ProcessingCapabilityForClientSystemEnum int32

const (
	ProcessingCapabilityForClientSystemEnum_CAN_PROCESS_IN_PARALLELL      ProcessingCapabilityForClientSystemEnum = 0 // Can be processed in parallell, by the Client system
	ProcessingCapabilityForClientSystemEnum_CAN_ONLY_PROCESS_IN_SERIAL    ProcessingCapabilityForClientSystemEnum = 1 // Can only be processed in serial, by the Client system
	ProcessingCapabilityForClientSystemEnum_STOP_SEND_NEW_EXECUTION_TASKS ProcessingCapabilityForClientSystemEnum = 2 // Stopp sending new tasks for execution
)

// Enum value maps for ProcessingCapabilityForClientSystemEnum.
var (
	ProcessingCapabilityForClientSystemEnum_name = map[int32]string{
		0: "CAN_PROCESS_IN_PARALLELL",
		1: "CAN_ONLY_PROCESS_IN_SERIAL",
		2: "STOP_SEND_NEW_EXECUTION_TASKS",
	}
	ProcessingCapabilityForClientSystemEnum_value = map[string]int32{
		"CAN_PROCESS_IN_PARALLELL":      0,
		"CAN_ONLY_PROCESS_IN_SERIAL":    1,
		"STOP_SEND_NEW_EXECUTION_TASKS": 2,
	}
)

func (x ProcessingCapabilityForClientSystemEnum) Enum() *ProcessingCapabilityForClientSystemEnum {
	p := new(ProcessingCapabilityForClientSystemEnum)
	*p = x
	return p
}

func (x ProcessingCapabilityForClientSystemEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessingCapabilityForClientSystemEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[3].Descriptor()
}

func (ProcessingCapabilityForClientSystemEnum) Type() protoreflect.EnumType {
	return &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[3]
}

func (x ProcessingCapabilityForClientSystemEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessingCapabilityForClientSystemEnum.Descriptor instead.
func (ProcessingCapabilityForClientSystemEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{3}
}

// Execution priority for TestInstructions, TestCases and TestSuites
type ExecutionPriorityEnum int32

const (
	ExecutionPriorityEnum_HIGHEST_PROBES             ExecutionPriorityEnum = 0 // Only triggered by system that need to trigger probes ');
	ExecutionPriorityEnum_HIGH_SINGLE_TESTCASE       ExecutionPriorityEnum = 1 // Used for execution of single TestCases triggered by the user');
	ExecutionPriorityEnum_HIGH_SINGLE_TESTSUITE      ExecutionPriorityEnum = 2 // Used for execution of single TestSuite triggered by the user');
	ExecutionPriorityEnum_MEDIUM_MULTIPLE_TESTCASES  ExecutionPriorityEnum = 3 // Used for execution of multiple TestCases triggered by the user');
	ExecutionPriorityEnum_MEDIUM_MULTIPLE_TESTSUITES ExecutionPriorityEnum = 4 // Used for execution of Multiple TestSuites triggered by the user');
	ExecutionPriorityEnum_LOW_SCHEDULED_TESTSUITES   ExecutionPriorityEnum = 5 //Scheduled Suites use this priority');
)

// Enum value maps for ExecutionPriorityEnum.
var (
	ExecutionPriorityEnum_name = map[int32]string{
		0: "HIGHEST_PROBES",
		1: "HIGH_SINGLE_TESTCASE",
		2: "HIGH_SINGLE_TESTSUITE",
		3: "MEDIUM_MULTIPLE_TESTCASES",
		4: "MEDIUM_MULTIPLE_TESTSUITES",
		5: "LOW_SCHEDULED_TESTSUITES",
	}
	ExecutionPriorityEnum_value = map[string]int32{
		"HIGHEST_PROBES":             0,
		"HIGH_SINGLE_TESTCASE":       1,
		"HIGH_SINGLE_TESTSUITE":      2,
		"MEDIUM_MULTIPLE_TESTCASES":  3,
		"MEDIUM_MULTIPLE_TESTSUITES": 4,
		"LOW_SCHEDULED_TESTSUITES":   5,
	}
)

func (x ExecutionPriorityEnum) Enum() *ExecutionPriorityEnum {
	p := new(ExecutionPriorityEnum)
	*p = x
	return p
}

func (x ExecutionPriorityEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecutionPriorityEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[4].Descriptor()
}

func (ExecutionPriorityEnum) Type() protoreflect.EnumType {
	return &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[4]
}

func (x ExecutionPriorityEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecutionPriorityEnum.Descriptor instead.
func (ExecutionPriorityEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{4}
}

// The TestInstructionAttribute can have one of the following types: "TextBox", "ComboBox", "FileSelector", "FunctionSelector"
type TestInstructionAttributeTypeEnum int32

const (
	TestInstructionAttributeTypeEnum_TEXTBOX           TestInstructionAttributeTypeEnum = 0 // Defines that TestInstructionAttribute of type "TextBox"
	TestInstructionAttributeTypeEnum_COMBOBOX          TestInstructionAttributeTypeEnum = 1 // Defines that TestInstructionAttriebute of type ComboBox
	TestInstructionAttributeTypeEnum_FILE_SELECTOR     TestInstructionAttributeTypeEnum = 2 // Defines that TestInstructionAttribute of type FileSelector
	TestInstructionAttributeTypeEnum_FUNCTION_SELECTOR TestInstructionAttributeTypeEnum = 3 // Defines that TestInstructionAttribute of type FunctionSelector
)

// Enum value maps for TestInstructionAttributeTypeEnum.
var (
	TestInstructionAttributeTypeEnum_name = map[int32]string{
		0: "TEXTBOX",
		1: "COMBOBOX",
		2: "FILE_SELECTOR",
		3: "FUNCTION_SELECTOR",
	}
	TestInstructionAttributeTypeEnum_value = map[string]int32{
		"TEXTBOX":           0,
		"COMBOBOX":          1,
		"FILE_SELECTOR":     2,
		"FUNCTION_SELECTOR": 3,
	}
)

func (x TestInstructionAttributeTypeEnum) Enum() *TestInstructionAttributeTypeEnum {
	p := new(TestInstructionAttributeTypeEnum)
	*p = x
	return p
}

func (x TestInstructionAttributeTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestInstructionAttributeTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[5].Descriptor()
}

func (TestInstructionAttributeTypeEnum) Type() protoreflect.EnumType {
	return &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes[5]
}

func (x TestInstructionAttributeTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestInstructionAttributeTypeEnum.Descriptor instead.
func (TestInstructionAttributeTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{5}
}

// Parameter used for Empty inputs, only containing current proto-file version
type EmptyParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoFileVersionUsedByClient CurrentFenixExecutionWorkerProtoFileVersionEnum `protobuf:"varint,1,opt,name=ProtoFileVersionUsedByClient,proto3,enum=fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum" json:"ProtoFileVersionUsedByClient,omitempty"` // The latest proto file version number to be used
}

func (x *EmptyParameter) Reset() {
	*x = EmptyParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyParameter) ProtoMessage() {}

func (x *EmptyParameter) ProtoReflect() protoreflect.Message {
	mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[0]
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
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{0}
}

func (x *EmptyParameter) GetProtoFileVersionUsedByClient() CurrentFenixExecutionWorkerProtoFileVersionEnum {
	if x != nil {
		return x.ProtoFileVersionUsedByClient
	}
	return CurrentFenixExecutionWorkerProtoFileVersionEnum_VERSION_0_0
}

// Ack/Nack- Response message with comment
type AckNackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckNack                      bool                                            `protobuf:"varint,1,opt,name=AckNack,proto3" json:"AckNack,omitempty"`                                                                                                                            // True=Ack, False=Nack
	Comments                     string                                          `protobuf:"bytes,2,opt,name=Comments,proto3" json:"Comments,omitempty"`                                                                                                                           //Comments if needed
	ErrorCodes                   []ErrorCodesEnum                                `protobuf:"varint,3,rep,packed,name=ErrorCodes,proto3,enum=fenixExecutionWorkerGrpcApi.ErrorCodesEnum" json:"ErrorCodes,omitempty"`                                                               // List of Error codes
	ProtoFileVersionUsedByClient CurrentFenixExecutionWorkerProtoFileVersionEnum `protobuf:"varint,4,opt,name=ProtoFileVersionUsedByClient,proto3,enum=fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum" json:"ProtoFileVersionUsedByClient,omitempty"` // The latest proto file version number to be used
}

func (x *AckNackResponse) Reset() {
	*x = AckNackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckNackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckNackResponse) ProtoMessage() {}

func (x *AckNackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[1]
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
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{1}
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

func (x *AckNackResponse) GetProtoFileVersionUsedByClient() CurrentFenixExecutionWorkerProtoFileVersionEnum {
	if x != nil {
		return x.ProtoFileVersionUsedByClient
	}
	return CurrentFenixExecutionWorkerProtoFileVersionEnum_VERSION_0_0
}

// Message Holding information about who is the calling client system
type ClientSystemIdentificationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainUuid                   string                                          `protobuf:"bytes,1,opt,name=DomainUuid,proto3" json:"DomainUuid,omitempty"`                                                                                                                       // The Users Id, SEB S-id will be used for SEB
	ProtoFileVersionUsedByClient CurrentFenixExecutionWorkerProtoFileVersionEnum `protobuf:"varint,2,opt,name=ProtoFileVersionUsedByClient,proto3,enum=fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum" json:"ProtoFileVersionUsedByClient,omitempty"` // The latest proto file version number to be used
}

func (x *ClientSystemIdentificationMessage) Reset() {
	*x = ClientSystemIdentificationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSystemIdentificationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSystemIdentificationMessage) ProtoMessage() {}

func (x *ClientSystemIdentificationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSystemIdentificationMessage.ProtoReflect.Descriptor instead.
func (*ClientSystemIdentificationMessage) Descriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP(), []int{2}
}

func (x *ClientSystemIdentificationMessage) GetDomainUuid() string {
	if x != nil {
		return x.DomainUuid
	}
	return ""
}

func (x *ClientSystemIdentificationMessage) GetProtoFileVersionUsedByClient() CurrentFenixExecutionWorkerProtoFileVersionEnum {
	if x != nil {
		return x.ProtoFileVersionUsedByClient
	}
	return CurrentFenixExecutionWorkerProtoFileVersionEnum_VERSION_0_0
}

var File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto protoreflect.FileDescriptor

var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDesc = []byte{
	0x0a, 0x6a, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x6e,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x90, 0x01, 0x0a,
	0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0xa7, 0x02, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x0a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0a, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42,
	0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x1c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65,
	0x64, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xd6, 0x01, 0x0a, 0x21, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x90, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e, 0x69,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2a, 0x57, 0x0a, 0x2f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e,
	0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x13, 0x0a, 0x0b, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x30, 0x5f, 0x30, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x30, 0x5f, 0x31, 0x10, 0x01, 0x2a, 0xb0, 0x01, 0x0a, 0x0e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f,
	0x50, 0x52, 0x4f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x4f, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x05, 0x2a, 0xba,
	0x01, 0x0a, 0x22, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x49, 0x45, 0x5f, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x49, 0x45, 0x5f,
	0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x54,
	0x49, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f,
	0x54, 0x49, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x5f, 0x4f, 0x4b, 0x10,
	0x03, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x49, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4b, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x49,
	0x45, 0x5f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x2a, 0x8a, 0x01, 0x0a, 0x27,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4c, 0x4c,
	0x45, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x52,
	0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x45,
	0x4e, 0x44, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x53, 0x10, 0x02, 0x2a, 0xbd, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x49, 0x47, 0x48, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x52,
	0x4f, 0x42, 0x45, 0x53, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x53,
	0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x43, 0x41, 0x53, 0x45, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x53, 0x55, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4d,
	0x45, 0x44, 0x49, 0x55, 0x4d, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x54,
	0x45, 0x53, 0x54, 0x43, 0x41, 0x53, 0x45, 0x53, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45,
	0x44, 0x49, 0x55, 0x4d, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x45,
	0x53, 0x54, 0x53, 0x55, 0x49, 0x54, 0x45, 0x53, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x4f,
	0x57, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x45, 0x53, 0x54,
	0x53, 0x55, 0x49, 0x54, 0x45, 0x53, 0x10, 0x05, 0x2a, 0x67, 0x0a, 0x20, 0x54, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x45, 0x58, 0x54, 0x42, 0x4f, 0x58, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d,
	0x42, 0x4f, 0x42, 0x4f, 0x58, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x10,
	0x03, 0x42, 0x2d, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0xaa, 0x02, 0x1b, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescOnce sync.Once
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescData = file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDesc
)

func file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescGZIP() []byte {
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescOnce.Do(func() {
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescData = protoimpl.X.CompressGZIP(file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescData)
	})
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDescData
}

var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_goTypes = []interface{}{
	(CurrentFenixExecutionWorkerProtoFileVersionEnum)(0), // 0: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum
	(ErrorCodesEnum)(0),                          // 1: fenixExecutionWorkerGrpcApi.ErrorCodesEnum
	(TestInstructionExecutionStatusEnum)(0),      // 2: fenixExecutionWorkerGrpcApi.TestInstructionExecutionStatusEnum
	(ProcessingCapabilityForClientSystemEnum)(0), // 3: fenixExecutionWorkerGrpcApi.ProcessingCapabilityForClientSystemEnum
	(ExecutionPriorityEnum)(0),                   // 4: fenixExecutionWorkerGrpcApi.ExecutionPriorityEnum
	(TestInstructionAttributeTypeEnum)(0),        // 5: fenixExecutionWorkerGrpcApi.TestInstructionAttributeTypeEnum
	(*EmptyParameter)(nil),                       // 6: fenixExecutionWorkerGrpcApi.EmptyParameter
	(*AckNackResponse)(nil),                      // 7: fenixExecutionWorkerGrpcApi.AckNackResponse
	(*ClientSystemIdentificationMessage)(nil),    // 8: fenixExecutionWorkerGrpcApi.ClientSystemIdentificationMessage
}
var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_depIdxs = []int32{
	0, // 0: fenixExecutionWorkerGrpcApi.EmptyParameter.ProtoFileVersionUsedByClient:type_name -> fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum
	1, // 1: fenixExecutionWorkerGrpcApi.AckNackResponse.ErrorCodes:type_name -> fenixExecutionWorkerGrpcApi.ErrorCodesEnum
	0, // 2: fenixExecutionWorkerGrpcApi.AckNackResponse.ProtoFileVersionUsedByClient:type_name -> fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum
	0, // 3: fenixExecutionWorkerGrpcApi.ClientSystemIdentificationMessage.ProtoFileVersionUsedByClient:type_name -> fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_init()
}
func file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_init() {
	if File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSystemIdentificationMessage); i {
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
			RawDescriptor: file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_goTypes,
		DependencyIndexes: file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_depIdxs,
		EnumInfos:         file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_enumTypes,
		MessageInfos:      file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_msgTypes,
	}.Build()
	File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto = out.File
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_rawDesc = nil
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_goTypes = nil
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_depIdxs = nil
}
