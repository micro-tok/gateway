// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/video/pb/video.proto

package videopb

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

type UploadVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video       []byte   `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	OwnerId     string   `protobuf:"bytes,4,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Tags        []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"` // Optional, if you want to include tags
}

func (x *UploadVideoRequest) Reset() {
	*x = UploadVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_video_pb_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoRequest) ProtoMessage() {}

func (x *UploadVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_video_pb_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoRequest) Descriptor() ([]byte, []int) {
	return file_pkg_video_pb_video_proto_rawDescGZIP(), []int{0}
}

func (x *UploadVideoRequest) GetVideo() []byte {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *UploadVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UploadVideoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UploadVideoRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *UploadVideoRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UploadVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadVideoResponse) Reset() {
	*x = UploadVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_video_pb_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoResponse) ProtoMessage() {}

func (x *UploadVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_video_pb_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoResponse.ProtoReflect.Descriptor instead.
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_video_pb_video_proto_rawDescGZIP(), []int{1}
}

func (x *UploadVideoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadVideoResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetVideoMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVideoMetadataRequest) Reset() {
	*x = GetVideoMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_video_pb_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoMetadataRequest) ProtoMessage() {}

func (x *GetVideoMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_video_pb_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetVideoMetadataRequest) Descriptor() ([]byte, []int) {
	return file_pkg_video_pb_video_proto_rawDescGZIP(), []int{2}
}

func (x *GetVideoMetadataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetVideoMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId     string   `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Url         string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Tags        []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"` // Optional
}

func (x *GetVideoMetadataResponse) Reset() {
	*x = GetVideoMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_video_pb_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoMetadataResponse) ProtoMessage() {}

func (x *GetVideoMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_video_pb_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetVideoMetadataResponse) Descriptor() ([]byte, []int) {
	return file_pkg_video_pb_video_proto_rawDescGZIP(), []int{3}
}

func (x *GetVideoMetadataResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetVideoMetadataResponse) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *GetVideoMetadataResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetVideoMetadataResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetVideoMetadataResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetVideoMetadataResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_pkg_video_pb_video_proto protoreflect.FileDescriptor

var file_pkg_video_pb_video_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0x37, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0xad, 0x01,
	0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x19, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x74, 0x6f, 0x6b, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_video_pb_video_proto_rawDescOnce sync.Once
	file_pkg_video_pb_video_proto_rawDescData = file_pkg_video_pb_video_proto_rawDesc
)

func file_pkg_video_pb_video_proto_rawDescGZIP() []byte {
	file_pkg_video_pb_video_proto_rawDescOnce.Do(func() {
		file_pkg_video_pb_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_video_pb_video_proto_rawDescData)
	})
	return file_pkg_video_pb_video_proto_rawDescData
}

var file_pkg_video_pb_video_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_video_pb_video_proto_goTypes = []interface{}{
	(*UploadVideoRequest)(nil),       // 0: video.UploadVideoRequest
	(*UploadVideoResponse)(nil),      // 1: video.UploadVideoResponse
	(*GetVideoMetadataRequest)(nil),  // 2: video.GetVideoMetadataRequest
	(*GetVideoMetadataResponse)(nil), // 3: video.GetVideoMetadataResponse
}
var file_pkg_video_pb_video_proto_depIdxs = []int32{
	0, // 0: video.VideoService.UploadVideo:input_type -> video.UploadVideoRequest
	2, // 1: video.VideoService.GetVideoMetadata:input_type -> video.GetVideoMetadataRequest
	1, // 2: video.VideoService.UploadVideo:output_type -> video.UploadVideoResponse
	3, // 3: video.VideoService.GetVideoMetadata:output_type -> video.GetVideoMetadataResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_video_pb_video_proto_init() }
func file_pkg_video_pb_video_proto_init() {
	if File_pkg_video_pb_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_video_pb_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoRequest); i {
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
		file_pkg_video_pb_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoResponse); i {
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
		file_pkg_video_pb_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoMetadataRequest); i {
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
		file_pkg_video_pb_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoMetadataResponse); i {
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
			RawDescriptor: file_pkg_video_pb_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_video_pb_video_proto_goTypes,
		DependencyIndexes: file_pkg_video_pb_video_proto_depIdxs,
		MessageInfos:      file_pkg_video_pb_video_proto_msgTypes,
	}.Build()
	File_pkg_video_pb_video_proto = out.File
	file_pkg_video_pb_video_proto_rawDesc = nil
	file_pkg_video_pb_video_proto_goTypes = nil
	file_pkg_video_pb_video_proto_depIdxs = nil
}
