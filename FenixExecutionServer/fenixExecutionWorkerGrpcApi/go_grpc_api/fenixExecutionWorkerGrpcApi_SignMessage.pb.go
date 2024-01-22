// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: FenixExecutionServer/fenixExecutionWorkerGrpcApi/fenixExecutionWorkerGrpcApi_SignMessage.proto

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

// The Message that BuilderServer send to Worker for Worker to sign
type SignMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoFileVersionUsedByClient CurrentFenixExecutionWorkerProtoFileVersionEnum `protobuf:"varint,1,opt,name=ProtoFileVersionUsedByClient,proto3,enum=fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum" json:"ProtoFileVersionUsedByClient,omitempty"` // The latest proto file version number to be used
	MessageToBeSigned            string                                          `protobuf:"bytes,2,opt,name=MessageToBeSigned,proto3" json:"MessageToBeSigned,omitempty"`                                                                                                         // The message to be Signe
}

func (x *SignMessageRequest) Reset() {
	*x = SignMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignMessageRequest) ProtoMessage() {}

func (x *SignMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignMessageRequest.ProtoReflect.Descriptor instead.
func (*SignMessageRequest) Descriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescGZIP(), []int{0}
}

func (x *SignMessageRequest) GetProtoFileVersionUsedByClient() CurrentFenixExecutionWorkerProtoFileVersionEnum {
	if x != nil {
		return x.ProtoFileVersionUsedByClient
	}
	return CurrentFenixExecutionWorkerProtoFileVersionEnum_CurrentFenixExecutionWorkerProtoFileVersionEnum_DEFAULT_NOT_SET
}

func (x *SignMessageRequest) GetMessageToBeSigned() string {
	if x != nil {
		return x.MessageToBeSigned
	}
	return ""
}

// The Message that holds the signature data
type SignMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckNackResponse                     *AckNackResponse                            `protobuf:"bytes,1,opt,name=AckNackResponse,proto3" json:"AckNackResponse,omitempty"`                                         // The AckNack message
	SignedMessageByWorkerServiceAccount *SignedMessageByWorkerServiceAccountMessage `protobuf:"bytes,2,opt,name=SignedMessageByWorkerServiceAccount,proto3" json:"SignedMessageByWorkerServiceAccount,omitempty"` // The message containing signature data
}

func (x *SignMessageResponse) Reset() {
	*x = SignMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignMessageResponse) ProtoMessage() {}

func (x *SignMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignMessageResponse.ProtoReflect.Descriptor instead.
func (*SignMessageResponse) Descriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescGZIP(), []int{1}
}

func (x *SignMessageResponse) GetAckNackResponse() *AckNackResponse {
	if x != nil {
		return x.AckNackResponse
	}
	return nil
}

func (x *SignMessageResponse) GetSignedMessageByWorkerServiceAccount() *SignedMessageByWorkerServiceAccountMessage {
	if x != nil {
		return x.SignedMessageByWorkerServiceAccount
	}
	return nil
}

// Holds information about signature signed by Workers Service Account
type SignedMessageByWorkerServiceAccountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageToBeSigned string `protobuf:"bytes,1,opt,name=MessageToBeSigned,proto3" json:"MessageToBeSigned,omitempty"` // The message that the service account signed
	HashOfSignature   string `protobuf:"bytes,2,opt,name=HashOfSignature,proto3" json:"HashOfSignature,omitempty"`     // The signature produced when signing the message
	HashedKeyId       string `protobuf:"bytes,3,opt,name=HashedKeyId,proto3" json:"HashedKeyId,omitempty"`             // THe Hash of the KeyId used when signing
}

func (x *SignedMessageByWorkerServiceAccountMessage) Reset() {
	*x = SignedMessageByWorkerServiceAccountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMessageByWorkerServiceAccountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMessageByWorkerServiceAccountMessage) ProtoMessage() {}

func (x *SignedMessageByWorkerServiceAccountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMessageByWorkerServiceAccountMessage.ProtoReflect.Descriptor instead.
func (*SignedMessageByWorkerServiceAccountMessage) Descriptor() ([]byte, []int) {
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescGZIP(), []int{2}
}

func (x *SignedMessageByWorkerServiceAccountMessage) GetMessageToBeSigned() string {
	if x != nil {
		return x.MessageToBeSigned
	}
	return ""
}

func (x *SignedMessageByWorkerServiceAccountMessage) GetHashOfSignature() string {
	if x != nil {
		return x.HashOfSignature
	}
	return ""
}

func (x *SignedMessageByWorkerServiceAccountMessage) GetHashedKeyId() string {
	if x != nil {
		return x.HashedKeyId
	}
	return ""
}

var File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto protoreflect.FileDescriptor

var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x53,
	0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x1a, 0x6a, 0x46,
	0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x5f, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x12, 0x53, 0x69,
	0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x90, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72,
	0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x6e,
	0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x42, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x42, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x22, 0x89, 0x02, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x41, 0x63, 0x6b,
	0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x2e, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0f, 0x41, 0x63, 0x6b, 0x4e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x99, 0x01, 0x0a, 0x23, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x47, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x23, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa6, 0x01,
	0x0a, 0x2a, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x42, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x42, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x48, 0x61,
	0x73, 0x68, 0x4f, 0x66, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x66, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x4b, 0x65,
	0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x48, 0x61, 0x73, 0x68, 0x65,
	0x64, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescOnce sync.Once
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescData = file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDesc
)

func file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescGZIP() []byte {
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescOnce.Do(func() {
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescData)
	})
	return file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDescData
}

var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_goTypes = []interface{}{
	(*SignMessageRequest)(nil),                           // 0: fenixExecutionWorkerGrpcApi.SignMessageRequest
	(*SignMessageResponse)(nil),                          // 1: fenixExecutionWorkerGrpcApi.SignMessageResponse
	(*SignedMessageByWorkerServiceAccountMessage)(nil),   // 2: fenixExecutionWorkerGrpcApi.SignedMessageByWorkerServiceAccountMessage
	(CurrentFenixExecutionWorkerProtoFileVersionEnum)(0), // 3: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum
	(*AckNackResponse)(nil),                              // 4: fenixExecutionWorkerGrpcApi.AckNackResponse
}
var file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_depIdxs = []int32{
	3, // 0: fenixExecutionWorkerGrpcApi.SignMessageRequest.ProtoFileVersionUsedByClient:type_name -> fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum
	4, // 1: fenixExecutionWorkerGrpcApi.SignMessageResponse.AckNackResponse:type_name -> fenixExecutionWorkerGrpcApi.AckNackResponse
	2, // 2: fenixExecutionWorkerGrpcApi.SignMessageResponse.SignedMessageByWorkerServiceAccount:type_name -> fenixExecutionWorkerGrpcApi.SignedMessageByWorkerServiceAccountMessage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_init()
}
func file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_init() {
	if File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto != nil {
		return
	}
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_GeneralMessagesAndEnums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignMessageRequest); i {
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
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignMessageResponse); i {
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
		file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMessageByWorkerServiceAccountMessage); i {
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
			RawDescriptor: file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_goTypes,
		DependencyIndexes: file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_depIdxs,
		MessageInfos:      file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_msgTypes,
	}.Build()
	File_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto = out.File
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_rawDesc = nil
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_goTypes = nil
	file_FenixExecutionServer_fenixExecutionWorkerGrpcApi_fenixExecutionWorkerGrpcApi_SignMessage_proto_depIdxs = nil
}