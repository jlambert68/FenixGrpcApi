// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters.proto

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

// AllTemplateRepositoryConnectionParameters
// Message holding parameters to be used to get all Templates from all Repositories
type AllTemplateRepositoryConnectionParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientSystemIdentification          *ClientSystemIdentificationMessage          `protobuf:"bytes,1,opt,name=ClientSystemIdentification,proto3" json:"ClientSystemIdentification,omitempty"`                   // Identifies Client System and Proto-file version used
	AllTemplateRepositories             []*TemplateRepositoryConnectionParameters   `protobuf:"bytes,2,rep,name=AllTemplateRepositories,proto3" json:"AllTemplateRepositories,omitempty"`                         // A list with all parameters to all different template repositories
	SignedMessageByWorkerServiceAccount *SignedMessageByWorkerServiceAccountMessage `protobuf:"bytes,3,opt,name=SignedMessageByWorkerServiceAccount,proto3" json:"SignedMessageByWorkerServiceAccount,omitempty"` // Holds information about signature signed by Workers Service Account
}

func (x *AllTemplateRepositoryConnectionParameters) Reset() {
	*x = AllTemplateRepositoryConnectionParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllTemplateRepositoryConnectionParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllTemplateRepositoryConnectionParameters) ProtoMessage() {}

func (x *AllTemplateRepositoryConnectionParameters) ProtoReflect() protoreflect.Message {
	mi := &file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllTemplateRepositoryConnectionParameters.ProtoReflect.Descriptor instead.
func (*AllTemplateRepositoryConnectionParameters) Descriptor() ([]byte, []int) {
	return file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescGZIP(), []int{0}
}

func (x *AllTemplateRepositoryConnectionParameters) GetClientSystemIdentification() *ClientSystemIdentificationMessage {
	if x != nil {
		return x.ClientSystemIdentification
	}
	return nil
}

func (x *AllTemplateRepositoryConnectionParameters) GetAllTemplateRepositories() []*TemplateRepositoryConnectionParameters {
	if x != nil {
		return x.AllTemplateRepositories
	}
	return nil
}

func (x *AllTemplateRepositoryConnectionParameters) GetSignedMessageByWorkerServiceAccount() *SignedMessageByWorkerServiceAccountMessage {
	if x != nil {
		return x.SignedMessageByWorkerServiceAccount
	}
	return nil
}

// TemplateRepositoryConnectionParameters
// Message holding parameters to be used to get a Template or group of Templates from one Repository
// TotalApiUrl := "repositoryApiUrl" + "/" +  repositoryOwner + "/" + repositoryName + "/contents" + repositoryPath
// TotalApiUrl := "https://api.github.com/repos/jlambert68/FenixTesterGui/contents"
type TemplateRepositoryConnectionParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryApiUrl string `protobuf:"bytes,1,opt,name=RepositoryApiUrl,proto3" json:"RepositoryApiUrl,omitempty"` // the base URL to the API e.g. "https://api.github.com/repos"
	RepositoryOwner  string `protobuf:"bytes,2,opt,name=RepositoryOwner,proto3" json:"RepositoryOwner,omitempty"`   // The owner of the repository e.g. "jlambert68"
	RepositoryName   string `protobuf:"bytes,3,opt,name=RepositoryName,proto3" json:"RepositoryName,omitempty"`     // The name of the repository e.g. "FenixTesterGui"
	RepositoryPath   string `protobuf:"bytes,4,opt,name=RepositoryPath,proto3" json:"RepositoryPath,omitempty"`     // The path within the repository e.g "" for storing templates in base of repository
	GitHubApiKey     string `protobuf:"bytes,5,opt,name=GitHubApiKey,proto3" json:"GitHubApiKey,omitempty"`         // The key used to access the files in this repository
}

func (x *TemplateRepositoryConnectionParameters) Reset() {
	*x = TemplateRepositoryConnectionParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRepositoryConnectionParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRepositoryConnectionParameters) ProtoMessage() {}

func (x *TemplateRepositoryConnectionParameters) ProtoReflect() protoreflect.Message {
	mi := &file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRepositoryConnectionParameters.ProtoReflect.Descriptor instead.
func (*TemplateRepositoryConnectionParameters) Descriptor() ([]byte, []int) {
	return file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateRepositoryConnectionParameters) GetRepositoryApiUrl() string {
	if x != nil {
		return x.RepositoryApiUrl
	}
	return ""
}

func (x *TemplateRepositoryConnectionParameters) GetRepositoryOwner() string {
	if x != nil {
		return x.RepositoryOwner
	}
	return ""
}

func (x *TemplateRepositoryConnectionParameters) GetRepositoryName() string {
	if x != nil {
		return x.RepositoryName
	}
	return ""
}

func (x *TemplateRepositoryConnectionParameters) GetRepositoryPath() string {
	if x != nil {
		return x.RepositoryPath
	}
	return ""
}

func (x *TemplateRepositoryConnectionParameters) GetGitHubApiKey() string {
	if x != nil {
		return x.GitHubApiKey
	}
	return ""
}

var File_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto protoreflect.FileDescriptor

var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDesc = []byte{
	0x0a, 0x8b, 0x01, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x5f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x1a, 0x7c, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65,
	0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x5f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0xa9, 0x01, 0x46, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x6e,
	0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x66,
	0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69,
	0x5f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x03, 0x0a, 0x29,
	0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x1a, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44,
	0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x1a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x83, 0x01, 0x0a, 0x17, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x49, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x17, 0x41,
	0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x23, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x66, 0x65, 0x6e, 0x69, 0x78, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x23, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf2, 0x01, 0x0a, 0x26, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x41, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x12,
	0x28, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x69, 0x74,
	0x48, 0x75, 0x62, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescOnce sync.Once
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescData = file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDesc
)

func file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescGZIP() []byte {
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescOnce.Do(func() {
		file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescData = protoimpl.X.CompressGZIP(file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescData)
	})
	return file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDescData
}

var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_goTypes = []interface{}{
	(*AllTemplateRepositoryConnectionParameters)(nil),  // 0: fenixTestCaseBuilderServerGrpcApi.AllTemplateRepositoryConnectionParameters
	(*TemplateRepositoryConnectionParameters)(nil),     // 1: fenixTestCaseBuilderServerGrpcApi.TemplateRepositoryConnectionParameters
	(*ClientSystemIdentificationMessage)(nil),          // 2: fenixTestCaseBuilderServerGrpcApi.ClientSystemIdentificationMessage
	(*SignedMessageByWorkerServiceAccountMessage)(nil), // 3: fenixTestCaseBuilderServerGrpcApi.SignedMessageByWorkerServiceAccountMessage
}
var file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_depIdxs = []int32{
	2, // 0: fenixTestCaseBuilderServerGrpcApi.AllTemplateRepositoryConnectionParameters.ClientSystemIdentification:type_name -> fenixTestCaseBuilderServerGrpcApi.ClientSystemIdentificationMessage
	1, // 1: fenixTestCaseBuilderServerGrpcApi.AllTemplateRepositoryConnectionParameters.AllTemplateRepositories:type_name -> fenixTestCaseBuilderServerGrpcApi.TemplateRepositoryConnectionParameters
	3, // 2: fenixTestCaseBuilderServerGrpcApi.AllTemplateRepositoryConnectionParameters.SignedMessageByWorkerServiceAccount:type_name -> fenixTestCaseBuilderServerGrpcApi.SignedMessageByWorkerServiceAccountMessage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_init()
}
func file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_init() {
	if File_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto != nil {
		return
	}
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_GeneralMessagesAndEnums_proto_init()
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllTemplateRepositoryConnectionParameters); i {
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
		file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRepositoryConnectionParameters); i {
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
			RawDescriptor: file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_goTypes,
		DependencyIndexes: file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_depIdxs,
		MessageInfos:      file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_msgTypes,
	}.Build()
	File_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto = out.File
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_rawDesc = nil
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_goTypes = nil
	file_FenixTestCaseBuilderServer_fenixTestCaseBuilderServerGrpcApi_fenixTestCaseBuilderServerGrpcApi_TemplateRepositoryConnectionParameters_proto_depIdxs = nil
}
