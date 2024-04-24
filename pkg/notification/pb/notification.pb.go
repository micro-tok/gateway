// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/notification/pb/notification.proto

package notificationpb

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

type RegisterTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterTokenRequest) Reset() {
	*x = RegisterTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterTokenRequest) ProtoMessage() {}

func (x *RegisterTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterTokenRequest.ProtoReflect.Descriptor instead.
func (*RegisterTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterTokenRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RegisterTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegisterTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Token  string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *RegisterTokenResponse) Reset() {
	*x = RegisterTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterTokenResponse) ProtoMessage() {}

func (x *RegisterTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterTokenResponse.ProtoReflect.Descriptor instead.
func (*RegisterTokenResponse) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterTokenResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RegisterTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterTokenResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type RemoveTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RemoveTokenRequest) Reset() {
	*x = RemoveTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTokenRequest) ProtoMessage() {}

func (x *RemoveTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTokenRequest.ProtoReflect.Descriptor instead.
func (*RemoveTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveTokenRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RemoveTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RemoveTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RemoveTokenResponse) Reset() {
	*x = RemoveTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTokenResponse) ProtoMessage() {}

func (x *RemoveTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTokenResponse.ProtoReflect.Descriptor instead.
func (*RemoveTokenResponse) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveTokenResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RemoveTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RemoveUserTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *RemoveUserTokensRequest) Reset() {
	*x = RemoveUserTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserTokensRequest) ProtoMessage() {}

func (x *RemoveUserTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserTokensRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserTokensRequest) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveUserTokensRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RemoveUserTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *RemoveUserTokensResponse) Reset() {
	*x = RemoveUserTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserTokensResponse) ProtoMessage() {}

func (x *RemoveUserTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserTokensResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserTokensResponse) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveUserTokensResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetUserTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetUserTokensRequest) Reset() {
	*x = GetUserTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokensRequest) ProtoMessage() {}

func (x *GetUserTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokensRequest.ProtoReflect.Descriptor instead.
func (*GetUserTokensRequest) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTokensRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetUserTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *GetUserTokensResponse) Reset() {
	*x = GetUserTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokensResponse) ProtoMessage() {}

func (x *GetUserTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokensResponse.ProtoReflect.Descriptor instead.
func (*GetUserTokensResponse) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokensResponse) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type GetAllTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllTokensRequest) Reset() {
	*x = GetAllTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTokensRequest) ProtoMessage() {}

func (x *GetAllTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTokensRequest.ProtoReflect.Descriptor instead.
func (*GetAllTokensRequest) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{8}
}

type UserTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Tokens []string `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *UserTokens) Reset() {
	*x = UserTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokens) ProtoMessage() {}

func (x *UserTokens) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokens.ProtoReflect.Descriptor instead.
func (*UserTokens) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{9}
}

func (x *UserTokens) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserTokens) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type GetAllTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTokens []*UserTokens `protobuf:"bytes,1,rep,name=userTokens,proto3" json:"userTokens,omitempty"`
}

func (x *GetAllTokensResponse) Reset() {
	*x = GetAllTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_pb_notification_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTokensResponse) ProtoMessage() {}

func (x *GetAllTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_pb_notification_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTokensResponse.ProtoReflect.Descriptor instead.
func (*GetAllTokensResponse) Descriptor() ([]byte, []int) {
	return file_pkg_notification_pb_notification_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllTokensResponse) GetUserTokens() []*UserTokens {
	if x != nil {
		return x.UserTokens
	}
	return nil
}

var File_pkg_notification_pb_notification_proto protoreflect.FileDescriptor

var file_pkg_notification_pb_notification_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x59, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0x3e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x22, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x32, 0xd7, 0x03, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x25, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x22, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x74, 0x6f, 0x6b, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_notification_pb_notification_proto_rawDescOnce sync.Once
	file_pkg_notification_pb_notification_proto_rawDescData = file_pkg_notification_pb_notification_proto_rawDesc
)

func file_pkg_notification_pb_notification_proto_rawDescGZIP() []byte {
	file_pkg_notification_pb_notification_proto_rawDescOnce.Do(func() {
		file_pkg_notification_pb_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_notification_pb_notification_proto_rawDescData)
	})
	return file_pkg_notification_pb_notification_proto_rawDescData
}

var file_pkg_notification_pb_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_notification_pb_notification_proto_goTypes = []interface{}{
	(*RegisterTokenRequest)(nil),     // 0: notification.RegisterTokenRequest
	(*RegisterTokenResponse)(nil),    // 1: notification.RegisterTokenResponse
	(*RemoveTokenRequest)(nil),       // 2: notification.RemoveTokenRequest
	(*RemoveTokenResponse)(nil),      // 3: notification.RemoveTokenResponse
	(*RemoveUserTokensRequest)(nil),  // 4: notification.RemoveUserTokensRequest
	(*RemoveUserTokensResponse)(nil), // 5: notification.RemoveUserTokensResponse
	(*GetUserTokensRequest)(nil),     // 6: notification.GetUserTokensRequest
	(*GetUserTokensResponse)(nil),    // 7: notification.GetUserTokensResponse
	(*GetAllTokensRequest)(nil),      // 8: notification.GetAllTokensRequest
	(*UserTokens)(nil),               // 9: notification.UserTokens
	(*GetAllTokensResponse)(nil),     // 10: notification.GetAllTokensResponse
}
var file_pkg_notification_pb_notification_proto_depIdxs = []int32{
	9,  // 0: notification.GetAllTokensResponse.userTokens:type_name -> notification.UserTokens
	0,  // 1: notification.NotificationService.RegisterToken:input_type -> notification.RegisterTokenRequest
	2,  // 2: notification.NotificationService.RemoveToken:input_type -> notification.RemoveTokenRequest
	4,  // 3: notification.NotificationService.RemoveUserTokens:input_type -> notification.RemoveUserTokensRequest
	6,  // 4: notification.NotificationService.GetUserTokens:input_type -> notification.GetUserTokensRequest
	8,  // 5: notification.NotificationService.GetAllTokens:input_type -> notification.GetAllTokensRequest
	1,  // 6: notification.NotificationService.RegisterToken:output_type -> notification.RegisterTokenResponse
	3,  // 7: notification.NotificationService.RemoveToken:output_type -> notification.RemoveTokenResponse
	5,  // 8: notification.NotificationService.RemoveUserTokens:output_type -> notification.RemoveUserTokensResponse
	7,  // 9: notification.NotificationService.GetUserTokens:output_type -> notification.GetUserTokensResponse
	10, // 10: notification.NotificationService.GetAllTokens:output_type -> notification.GetAllTokensResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_notification_pb_notification_proto_init() }
func file_pkg_notification_pb_notification_proto_init() {
	if File_pkg_notification_pb_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_notification_pb_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterTokenRequest); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterTokenResponse); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTokenRequest); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTokenResponse); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserTokensRequest); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserTokensResponse); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokensRequest); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokensResponse); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTokensRequest); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokens); i {
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
		file_pkg_notification_pb_notification_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTokensResponse); i {
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
			RawDescriptor: file_pkg_notification_pb_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_notification_pb_notification_proto_goTypes,
		DependencyIndexes: file_pkg_notification_pb_notification_proto_depIdxs,
		MessageInfos:      file_pkg_notification_pb_notification_proto_msgTypes,
	}.Build()
	File_pkg_notification_pb_notification_proto = out.File
	file_pkg_notification_pb_notification_proto_rawDesc = nil
	file_pkg_notification_pb_notification_proto_goTypes = nil
	file_pkg_notification_pb_notification_proto_depIdxs = nil
}