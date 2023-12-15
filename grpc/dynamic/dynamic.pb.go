// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: dynamic.proto

package dynamic

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

type UnaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Meta    map[string]string `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Action  string            `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *UnaryRequest) Reset() {
	*x = UnaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRequest) ProtoMessage() {}

func (x *UnaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRequest.ProtoReflect.Descriptor instead.
func (*UnaryRequest) Descriptor() ([]byte, []int) {
	return file_dynamic_proto_rawDescGZIP(), []int{0}
}

func (x *UnaryRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *UnaryRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *UnaryRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type UnaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Meta    map[string]string `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Action  string            `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *UnaryResponse) Reset() {
	*x = UnaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryResponse) ProtoMessage() {}

func (x *UnaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynamic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryResponse.ProtoReflect.Descriptor instead.
func (*UnaryResponse) Descriptor() ([]byte, []int) {
	return file_dynamic_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *UnaryResponse) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *UnaryResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_dynamic_proto protoreflect.FileDescriptor

var file_dynamic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa6, 0x01, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x55, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x31, 0x0a, 0x07, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x12, 0x26,
	0x0a, 0x05, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x0d, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x48, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x66, 0x75, 0x6e,
	0x2f, 0x78, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x3b,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dynamic_proto_rawDescOnce sync.Once
	file_dynamic_proto_rawDescData = file_dynamic_proto_rawDesc
)

func file_dynamic_proto_rawDescGZIP() []byte {
	file_dynamic_proto_rawDescOnce.Do(func() {
		file_dynamic_proto_rawDescData = protoimpl.X.CompressGZIP(file_dynamic_proto_rawDescData)
	})
	return file_dynamic_proto_rawDescData
}

var file_dynamic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dynamic_proto_goTypes = []interface{}{
	(*UnaryRequest)(nil),  // 0: UnaryRequest
	(*UnaryResponse)(nil), // 1: UnaryResponse
	nil,                   // 2: UnaryRequest.MetaEntry
	nil,                   // 3: UnaryResponse.MetaEntry
}
var file_dynamic_proto_depIdxs = []int32{
	2, // 0: UnaryRequest.meta:type_name -> UnaryRequest.MetaEntry
	3, // 1: UnaryResponse.meta:type_name -> UnaryResponse.MetaEntry
	0, // 2: Dynamic.Unary:input_type -> UnaryRequest
	1, // 3: Dynamic.Unary:output_type -> UnaryResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dynamic_proto_init() }
func file_dynamic_proto_init() {
	if File_dynamic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dynamic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryRequest); i {
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
		file_dynamic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryResponse); i {
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
			RawDescriptor: file_dynamic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dynamic_proto_goTypes,
		DependencyIndexes: file_dynamic_proto_depIdxs,
		MessageInfos:      file_dynamic_proto_msgTypes,
	}.Build()
	File_dynamic_proto = out.File
	file_dynamic_proto_rawDesc = nil
	file_dynamic_proto_goTypes = nil
	file_dynamic_proto_depIdxs = nil
}