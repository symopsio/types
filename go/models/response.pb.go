// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: models/response.proto

package models

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

type EscalationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EscalationResponse) Reset() {
	*x = EscalationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EscalationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EscalationResponse) ProtoMessage() {}

func (x *EscalationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EscalationResponse.ProtoReflect.Descriptor instead.
func (*EscalationResponse) Descriptor() ([]byte, []int) {
	return file_models_response_proto_rawDescGZIP(), []int{0}
}

func (x *EscalationResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *EscalationResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ExpirationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ExpirationResponse) Reset() {
	*x = ExpirationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpirationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirationResponse) ProtoMessage() {}

func (x *ExpirationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirationResponse.ProtoReflect.Descriptor instead.
func (*ExpirationResponse) Descriptor() ([]byte, []int) {
	return file_models_response_proto_rawDescGZIP(), []int{1}
}

func (x *ExpirationResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ExpirationResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_models_response_proto protoreflect.FileDescriptor

var file_models_response_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x79, 0x6d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x22, 0x3a, 0x0a, 0x12, 0x45, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x3a, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6d, 0x6f, 0x70, 0x73,
	0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_response_proto_rawDescOnce sync.Once
	file_models_response_proto_rawDescData = file_models_response_proto_rawDesc
)

func file_models_response_proto_rawDescGZIP() []byte {
	file_models_response_proto_rawDescOnce.Do(func() {
		file_models_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_response_proto_rawDescData)
	})
	return file_models_response_proto_rawDescData
}

var file_models_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_response_proto_goTypes = []interface{}{
	(*EscalationResponse)(nil), // 0: sym.models.EscalationResponse
	(*ExpirationResponse)(nil), // 1: sym.models.ExpirationResponse
}
var file_models_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_response_proto_init() }
func file_models_response_proto_init() {
	if File_models_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EscalationResponse); i {
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
		file_models_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpirationResponse); i {
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
			RawDescriptor: file_models_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_response_proto_goTypes,
		DependencyIndexes: file_models_response_proto_depIdxs,
		MessageInfos:      file_models_response_proto_msgTypes,
	}.Build()
	File_models_response_proto = out.File
	file_models_response_proto_rawDesc = nil
	file_models_response_proto_goTypes = nil
	file_models_response_proto_depIdxs = nil
}