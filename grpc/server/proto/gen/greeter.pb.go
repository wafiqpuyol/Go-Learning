// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: proto/greeter.proto

package mainpb

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

type GreetingMsg struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greet         string                 `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GreetingMsg) Reset() {
	*x = GreetingMsg{}
	mi := &file_proto_greeter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetingMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingMsg) ProtoMessage() {}

func (x *GreetingMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greeter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingMsg.ProtoReflect.Descriptor instead.
func (*GreetingMsg) Descriptor() ([]byte, []int) {
	return file_proto_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingMsg) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_proto_greeter_proto protoreflect.FileDescriptor

var file_proto_greeter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a,
	0x0b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x32, 0x37, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_greeter_proto_rawDescOnce sync.Once
	file_proto_greeter_proto_rawDescData = file_proto_greeter_proto_rawDesc
)

func file_proto_greeter_proto_rawDescGZIP() []byte {
	file_proto_greeter_proto_rawDescOnce.Do(func() {
		file_proto_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_greeter_proto_rawDescData)
	})
	return file_proto_greeter_proto_rawDescData
}

var file_proto_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_greeter_proto_goTypes = []any{
	(*GreetingMsg)(nil), // 0: main.GreetingMsg
	(*Empty)(nil),       // 1: main.Empty
}
var file_proto_greeter_proto_depIdxs = []int32{
	1, // 0: main.Greeter.SayHello:input_type -> main.Empty
	0, // 1: main.Greeter.SayHello:output_type -> main.GreetingMsg
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_greeter_proto_init() }
func file_proto_greeter_proto_init() {
	if File_proto_greeter_proto != nil {
		return
	}
	file_proto_main_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_greeter_proto_goTypes,
		DependencyIndexes: file_proto_greeter_proto_depIdxs,
		MessageInfos:      file_proto_greeter_proto_msgTypes,
	}.Build()
	File_proto_greeter_proto = out.File
	file_proto_greeter_proto_rawDesc = nil
	file_proto_greeter_proto_goTypes = nil
	file_proto_greeter_proto_depIdxs = nil
}
