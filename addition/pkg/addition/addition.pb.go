// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: pkg/addition/addition.proto

package addition

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the user's name.
type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args string `protobuf:"bytes,1,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_addition_addition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_addition_addition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_pkg_addition_addition_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

// The response message containing the greetings
type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum string `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_addition_addition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_addition_addition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_pkg_addition_addition_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() string {
	if x != nil {
		return x.Sum
	}
	return ""
}

var File_pkg_addition_addition_proto protoreflect.FileDescriptor

var file_pkg_addition_addition_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x41, 0x0a, 0x05, 0x41, 0x64,
	0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x63, 0x41, 0x64, 0x64, 0x12, 0x14,
	0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_addition_addition_proto_rawDescOnce sync.Once
	file_pkg_addition_addition_proto_rawDescData = file_pkg_addition_addition_proto_rawDesc
)

func file_pkg_addition_addition_proto_rawDescGZIP() []byte {
	file_pkg_addition_addition_proto_rawDescOnce.Do(func() {
		file_pkg_addition_addition_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_addition_addition_proto_rawDescData)
	})
	return file_pkg_addition_addition_proto_rawDescData
}

var file_pkg_addition_addition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_addition_addition_proto_goTypes = []interface{}{
	(*AddRequest)(nil),  // 0: addition.AddRequest
	(*AddResponse)(nil), // 1: addition.AddResponse
}
var file_pkg_addition_addition_proto_depIdxs = []int32{
	0, // 0: addition.Adder.CalcAdd:input_type -> addition.AddRequest
	1, // 1: addition.Adder.CalcAdd:output_type -> addition.AddResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_addition_addition_proto_init() }
func file_pkg_addition_addition_proto_init() {
	if File_pkg_addition_addition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_addition_addition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_pkg_addition_addition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
			RawDescriptor: file_pkg_addition_addition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_addition_addition_proto_goTypes,
		DependencyIndexes: file_pkg_addition_addition_proto_depIdxs,
		MessageInfos:      file_pkg_addition_addition_proto_msgTypes,
	}.Build()
	File_pkg_addition_addition_proto = out.File
	file_pkg_addition_addition_proto_rawDesc = nil
	file_pkg_addition_addition_proto_goTypes = nil
	file_pkg_addition_addition_proto_depIdxs = nil
}
