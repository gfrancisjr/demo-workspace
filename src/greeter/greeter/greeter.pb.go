// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: greeter.proto

package greeter

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

type SimpleGreetingRequest_Type int32

const (
	SimpleGreetingRequest_UNKNOWN        SimpleGreetingRequest_Type = 0
	SimpleGreetingRequest_HELLO          SimpleGreetingRequest_Type = 1
	SimpleGreetingRequest_GOOD_MORNING   SimpleGreetingRequest_Type = 2
	SimpleGreetingRequest_GOOD_NIGHT     SimpleGreetingRequest_Type = 3
	SimpleGreetingRequest_HAPPY_BIRTHDAY SimpleGreetingRequest_Type = 4
	SimpleGreetingRequest_MERRY_XMAS     SimpleGreetingRequest_Type = 5
)

// Enum value maps for SimpleGreetingRequest_Type.
var (
	SimpleGreetingRequest_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "HELLO",
		2: "GOOD_MORNING",
		3: "GOOD_NIGHT",
		4: "HAPPY_BIRTHDAY",
		5: "MERRY_XMAS",
	}
	SimpleGreetingRequest_Type_value = map[string]int32{
		"UNKNOWN":        0,
		"HELLO":          1,
		"GOOD_MORNING":   2,
		"GOOD_NIGHT":     3,
		"HAPPY_BIRTHDAY": 4,
		"MERRY_XMAS":     5,
	}
)

func (x SimpleGreetingRequest_Type) Enum() *SimpleGreetingRequest_Type {
	p := new(SimpleGreetingRequest_Type)
	*p = x
	return p
}

func (x SimpleGreetingRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SimpleGreetingRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_greeter_proto_enumTypes[0].Descriptor()
}

func (SimpleGreetingRequest_Type) Type() protoreflect.EnumType {
	return &file_greeter_proto_enumTypes[0]
}

func (x SimpleGreetingRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SimpleGreetingRequest_Type.Descriptor instead.
func (SimpleGreetingRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_greeter_proto_rawDescGZIP(), []int{0, 0}
}

// message containing the greeting request
type SimpleGreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type SimpleGreetingRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=greeter.SimpleGreetingRequest_Type" json:"type,omitempty"`
	Name string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimpleGreetingRequest) Reset() {
	*x = SimpleGreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleGreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleGreetingRequest) ProtoMessage() {}

func (x *SimpleGreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleGreetingRequest.ProtoReflect.Descriptor instead.
func (*SimpleGreetingRequest) Descriptor() ([]byte, []int) {
	return file_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleGreetingRequest) GetType() SimpleGreetingRequest_Type {
	if x != nil {
		return x.Type
	}
	return SimpleGreetingRequest_UNKNOWN
}

func (x *SimpleGreetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// message containing the greeting response
type SimpleGreetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *SimpleGreetingResponse) Reset() {
	*x = SimpleGreetingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleGreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleGreetingResponse) ProtoMessage() {}

func (x *SimpleGreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleGreetingResponse.ProtoReflect.Descriptor instead.
func (*SimpleGreetingResponse) Descriptor() ([]byte, []int) {
	return file_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleGreetingResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

var File_greeter_proto protoreflect.FileDescriptor

var file_greeter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x15, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x64, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x47, 0x4f, 0x4f, 0x44, 0x5f, 0x4d, 0x4f, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x4f, 0x4f, 0x44, 0x5f, 0x4e, 0x49, 0x47, 0x48, 0x54, 0x10,
	0x03, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x41, 0x50, 0x50, 0x59, 0x5f, 0x42, 0x49, 0x52, 0x54, 0x48,
	0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x52, 0x52, 0x59, 0x5f, 0x58,
	0x4d, 0x41, 0x53, 0x10, 0x05, 0x22, 0x34, 0x0a, 0x16, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x5c, 0x0a, 0x07, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x66, 0x72, 0x61, 0x6e, 0x63, 0x69, 0x73,
	0x6a, 0x72, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greeter_proto_rawDescOnce sync.Once
	file_greeter_proto_rawDescData = file_greeter_proto_rawDesc
)

func file_greeter_proto_rawDescGZIP() []byte {
	file_greeter_proto_rawDescOnce.Do(func() {
		file_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_greeter_proto_rawDescData)
	})
	return file_greeter_proto_rawDescData
}

var file_greeter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greeter_proto_goTypes = []interface{}{
	(SimpleGreetingRequest_Type)(0), // 0: greeter.SimpleGreetingRequest.Type
	(*SimpleGreetingRequest)(nil),   // 1: greeter.SimpleGreetingRequest
	(*SimpleGreetingResponse)(nil),  // 2: greeter.SimpleGreetingResponse
}
var file_greeter_proto_depIdxs = []int32{
	0, // 0: greeter.SimpleGreetingRequest.type:type_name -> greeter.SimpleGreetingRequest.Type
	1, // 1: greeter.Greeter.SimpleGreeting:input_type -> greeter.SimpleGreetingRequest
	2, // 2: greeter.Greeter.SimpleGreeting:output_type -> greeter.SimpleGreetingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_greeter_proto_init() }
func file_greeter_proto_init() {
	if File_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleGreetingRequest); i {
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
		file_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleGreetingResponse); i {
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
			RawDescriptor: file_greeter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greeter_proto_goTypes,
		DependencyIndexes: file_greeter_proto_depIdxs,
		EnumInfos:         file_greeter_proto_enumTypes,
		MessageInfos:      file_greeter_proto_msgTypes,
	}.Build()
	File_greeter_proto = out.File
	file_greeter_proto_rawDesc = nil
	file_greeter_proto_goTypes = nil
	file_greeter_proto_depIdxs = nil
}
