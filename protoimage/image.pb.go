// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: image.proto

package protoimage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImageMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*ImageMessage_Metadata
	//	*ImageMessage_ImageChunk
	Data          isImageMessage_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageMessage) Reset() {
	*x = ImageMessage{}
	mi := &file_image_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMessage) ProtoMessage() {}

func (x *ImageMessage) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMessage.ProtoReflect.Descriptor instead.
func (*ImageMessage) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *ImageMessage) GetData() isImageMessage_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ImageMessage) GetMetadata() *Metadata {
	if x != nil {
		if x, ok := x.Data.(*ImageMessage_Metadata); ok {
			return x.Metadata
		}
	}
	return nil
}

func (x *ImageMessage) GetImageChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*ImageMessage_ImageChunk); ok {
			return x.ImageChunk
		}
	}
	return nil
}

type isImageMessage_Data interface {
	isImageMessage_Data()
}

type ImageMessage_Metadata struct {
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type ImageMessage_ImageChunk struct {
	ImageChunk []byte `protobuf:"bytes,2,opt,name=image_chunk,json=imageChunk,proto3,oneof"`
}

func (*ImageMessage_Metadata) isImageMessage_Data() {}

func (*ImageMessage_ImageChunk) isImageMessage_Data() {}

type UploadResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageId       string                 `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResult) Reset() {
	*x = UploadResult{}
	mi := &file_image_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResult) ProtoMessage() {}

func (x *UploadResult) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResult.ProtoReflect.Descriptor instead.
func (*UploadResult) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResult) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

type Metadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	EntityId      string                 `protobuf:"bytes,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	IsCover       bool                   `protobuf:"varint,3,opt,name=is_cover,json=isCover,proto3" json:"is_cover,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_image_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Metadata) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *Metadata) GetIsCover() bool {
	if x != nil {
		return x.IsCover
	}
	return false
}

var File_image_proto protoreflect.FileDescriptor

const file_image_proto_rawDesc = "" +
	"\n" +
	"\vimage.proto\"b\n" +
	"\fImageMessage\x12'\n" +
	"\bmetadata\x18\x01 \x01(\v2\t.MetadataH\x00R\bmetadata\x12!\n" +
	"\vimage_chunk\x18\x02 \x01(\fH\x00R\n" +
	"imageChunkB\x06\n" +
	"\x04data\")\n" +
	"\fUploadResult\x12\x19\n" +
	"\bimage_id\x18\x01 \x01(\tR\aimageId\"\\\n" +
	"\bMetadata\x12\x18\n" +
	"\aservice\x18\x01 \x01(\tR\aservice\x12\x1b\n" +
	"\tentity_id\x18\x02 \x01(\tR\bentityId\x12\x19\n" +
	"\bis_cover\x18\x03 \x01(\bR\aisCover26\n" +
	"\x05Image\x12-\n" +
	"\vUploadImage\x12\r.ImageMessage\x1a\r.UploadResult(\x01B1Z/github.com/Gleb988/online-shop_proto/protoimageb\x06proto3"

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData []byte
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_image_proto_rawDesc), len(file_image_proto_rawDesc)))
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_image_proto_goTypes = []any{
	(*ImageMessage)(nil), // 0: ImageMessage
	(*UploadResult)(nil), // 1: UploadResult
	(*Metadata)(nil),     // 2: Metadata
}
var file_image_proto_depIdxs = []int32{
	2, // 0: ImageMessage.metadata:type_name -> Metadata
	0, // 1: Image.UploadImage:input_type -> ImageMessage
	1, // 2: Image.UploadImage:output_type -> UploadResult
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	file_image_proto_msgTypes[0].OneofWrappers = []any{
		(*ImageMessage_Metadata)(nil),
		(*ImageMessage_ImageChunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_image_proto_rawDesc), len(file_image_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}
