// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: vcs/v1/vcs.proto

package vcsv1

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

type GithubAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GithubAppRequest) Reset() {
	*x = GithubAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_v1_vcs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubAppRequest) ProtoMessage() {}

func (x *GithubAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_v1_vcs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubAppRequest.ProtoReflect.Descriptor instead.
func (*GithubAppRequest) Descriptor() ([]byte, []int) {
	return file_vcs_v1_vcs_proto_rawDescGZIP(), []int{0}
}

type GithubAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *GithubAppResponse) Reset() {
	*x = GithubAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_v1_vcs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubAppResponse) ProtoMessage() {}

func (x *GithubAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_v1_vcs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubAppResponse.ProtoReflect.Descriptor instead.
func (*GithubAppResponse) Descriptor() ([]byte, []int) {
	return file_vcs_v1_vcs_proto_rawDescGZIP(), []int{1}
}

func (x *GithubAppResponse) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type GithubLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizationCode string `protobuf:"bytes,1,opt,name=authorizationCode,proto3" json:"authorizationCode,omitempty"`
}

func (x *GithubLoginRequest) Reset() {
	*x = GithubLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_v1_vcs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubLoginRequest) ProtoMessage() {}

func (x *GithubLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_v1_vcs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubLoginRequest.ProtoReflect.Descriptor instead.
func (*GithubLoginRequest) Descriptor() ([]byte, []int) {
	return file_vcs_v1_vcs_proto_rawDescGZIP(), []int{2}
}

func (x *GithubLoginRequest) GetAuthorizationCode() string {
	if x != nil {
		return x.AuthorizationCode
	}
	return ""
}

type GithubLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GithubLoginResponse) Reset() {
	*x = GithubLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_v1_vcs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubLoginResponse) ProtoMessage() {}

func (x *GithubLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_v1_vcs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubLoginResponse.ProtoReflect.Descriptor instead.
func (*GithubLoginResponse) Descriptor() ([]byte, []int) {
	return file_vcs_v1_vcs_proto_rawDescGZIP(), []int{3}
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the full path to the repository
	RepositoryURL string `protobuf:"bytes,1,opt,name=repositoryURL,proto3" json:"repositoryURL,omitempty"`
	// the vcs ref to get the file from
	Ref string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// the path to the file as provided by the symbols
	LocalPath string `protobuf:"bytes,3,opt,name=localPath,proto3" json:"localPath,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_v1_vcs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_v1_vcs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_vcs_v1_vcs_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileRequest) GetRepositoryURL() string {
	if x != nil {
		return x.RepositoryURL
	}
	return ""
}

func (x *GetFileRequest) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *GetFileRequest) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base64 content of the file
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// the full URL to the file
	URL string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcs_v1_vcs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vcs_v1_vcs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_vcs_v1_vcs_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetFileResponse) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

var File_vcs_v1_vcs_proto protoreflect.FileDescriptor

var file_vcs_v1_vcs_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x76, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f,
	0x0a, 0x11, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22,
	0x42, 0x0a, 0x12, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55,
	0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52,
	0x4c, 0x32, 0xd8, 0x01, 0x0a, 0x0a, 0x56, 0x43, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x09, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x12, 0x18, 0x2e,
	0x76, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x76, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x76, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x76, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x76, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8b, 0x01, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x56, 0x63, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x63,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x56, 0x63, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x06, 0x56, 0x63, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x56, 0x63,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x07, 0x56, 0x63, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vcs_v1_vcs_proto_rawDescOnce sync.Once
	file_vcs_v1_vcs_proto_rawDescData = file_vcs_v1_vcs_proto_rawDesc
)

func file_vcs_v1_vcs_proto_rawDescGZIP() []byte {
	file_vcs_v1_vcs_proto_rawDescOnce.Do(func() {
		file_vcs_v1_vcs_proto_rawDescData = protoimpl.X.CompressGZIP(file_vcs_v1_vcs_proto_rawDescData)
	})
	return file_vcs_v1_vcs_proto_rawDescData
}

var file_vcs_v1_vcs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vcs_v1_vcs_proto_goTypes = []interface{}{
	(*GithubAppRequest)(nil),    // 0: vcs.v1.GithubAppRequest
	(*GithubAppResponse)(nil),   // 1: vcs.v1.GithubAppResponse
	(*GithubLoginRequest)(nil),  // 2: vcs.v1.GithubLoginRequest
	(*GithubLoginResponse)(nil), // 3: vcs.v1.GithubLoginResponse
	(*GetFileRequest)(nil),      // 4: vcs.v1.GetFileRequest
	(*GetFileResponse)(nil),     // 5: vcs.v1.GetFileResponse
}
var file_vcs_v1_vcs_proto_depIdxs = []int32{
	0, // 0: vcs.v1.VCSService.GithubApp:input_type -> vcs.v1.GithubAppRequest
	2, // 1: vcs.v1.VCSService.GithubLogin:input_type -> vcs.v1.GithubLoginRequest
	4, // 2: vcs.v1.VCSService.GetFile:input_type -> vcs.v1.GetFileRequest
	1, // 3: vcs.v1.VCSService.GithubApp:output_type -> vcs.v1.GithubAppResponse
	3, // 4: vcs.v1.VCSService.GithubLogin:output_type -> vcs.v1.GithubLoginResponse
	5, // 5: vcs.v1.VCSService.GetFile:output_type -> vcs.v1.GetFileResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vcs_v1_vcs_proto_init() }
func file_vcs_v1_vcs_proto_init() {
	if File_vcs_v1_vcs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vcs_v1_vcs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubAppRequest); i {
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
		file_vcs_v1_vcs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubAppResponse); i {
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
		file_vcs_v1_vcs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubLoginRequest); i {
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
		file_vcs_v1_vcs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubLoginResponse); i {
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
		file_vcs_v1_vcs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_vcs_v1_vcs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
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
			RawDescriptor: file_vcs_v1_vcs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vcs_v1_vcs_proto_goTypes,
		DependencyIndexes: file_vcs_v1_vcs_proto_depIdxs,
		MessageInfos:      file_vcs_v1_vcs_proto_msgTypes,
	}.Build()
	File_vcs_v1_vcs_proto = out.File
	file_vcs_v1_vcs_proto_rawDesc = nil
	file_vcs_v1_vcs_proto_goTypes = nil
	file_vcs_v1_vcs_proto_depIdxs = nil
}
