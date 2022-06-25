// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: ping_pong.proto

package pingpong

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ball struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Ball) Reset() {
	*x = Ball{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_pong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ball) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ball) ProtoMessage() {}

func (x *Ball) ProtoReflect() protoreflect.Message {
	mi := &file_ping_pong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ball.ProtoReflect.Descriptor instead.
func (*Ball) Descriptor() ([]byte, []int) {
	return file_ping_pong_proto_rawDescGZIP(), []int{0}
}

func (x *Ball) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_ping_pong_proto protoreflect.FileDescriptor

var file_ping_pong_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x42, 0x61, 0x6c, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x32, 0x40, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67,
	0x12, 0x34, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x69, 0x74, 0x69, 0x6e, 0x74, 0x6c, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_pong_proto_rawDescOnce sync.Once
	file_ping_pong_proto_rawDescData = file_ping_pong_proto_rawDesc
)

func file_ping_pong_proto_rawDescGZIP() []byte {
	file_ping_pong_proto_rawDescOnce.Do(func() {
		file_ping_pong_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_pong_proto_rawDescData)
	})
	return file_ping_pong_proto_rawDescData
}

var file_ping_pong_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ping_pong_proto_goTypes = []interface{}{
	(*Ball)(nil),          // 0: PingPong.Ball
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_ping_pong_proto_depIdxs = []int32{
	0, // 0: PingPong.PingPong.SendBall:input_type -> PingPong.Ball
	1, // 1: PingPong.PingPong.SendBall:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_pong_proto_init() }
func file_ping_pong_proto_init() {
	if File_ping_pong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_pong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ball); i {
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
			RawDescriptor: file_ping_pong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_pong_proto_goTypes,
		DependencyIndexes: file_ping_pong_proto_depIdxs,
		MessageInfos:      file_ping_pong_proto_msgTypes,
	}.Build()
	File_ping_pong_proto = out.File
	file_ping_pong_proto_rawDesc = nil
	file_ping_pong_proto_goTypes = nil
	file_ping_pong_proto_depIdxs = nil
}
