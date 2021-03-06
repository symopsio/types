// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: sym/messages/authz.proto

package messages

import (
	proto "github.com/golang/protobuf/proto"
	models "github.com/symopsio/types/go/models"
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

// Authz messages test if a user is authorized for a role
type Authz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema *models.Schema `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	User   string         `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Role   string         `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Authz) Reset() {
	*x = Authz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sym_messages_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authz) ProtoMessage() {}

func (x *Authz) ProtoReflect() protoreflect.Message {
	mi := &file_sym_messages_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authz.ProtoReflect.Descriptor instead.
func (*Authz) Descriptor() ([]byte, []int) {
	return file_sym_messages_authz_proto_rawDescGZIP(), []int{0}
}

func (x *Authz) GetSchema() *models.Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *Authz) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Authz) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// AuthzResponse gets sent back by authz requests
type AuthzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok           bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"` // error message when not ok
}

func (x *AuthzResponse) Reset() {
	*x = AuthzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sym_messages_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzResponse) ProtoMessage() {}

func (x *AuthzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sym_messages_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzResponse.ProtoReflect.Descriptor instead.
func (*AuthzResponse) Descriptor() ([]byte, []int) {
	return file_sym_messages_authz_proto_rawDescGZIP(), []int{1}
}

func (x *AuthzResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *AuthzResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_sym_messages_authz_proto protoreflect.FileDescriptor

var file_sym_messages_authz_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x79, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x79, 0x6d, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x17, 0x73, 0x79, 0x6d, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5b, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x79, 0x6d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x43,
	0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x79, 0x6d, 0x6f, 0x70, 0x73, 0x69, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sym_messages_authz_proto_rawDescOnce sync.Once
	file_sym_messages_authz_proto_rawDescData = file_sym_messages_authz_proto_rawDesc
)

func file_sym_messages_authz_proto_rawDescGZIP() []byte {
	file_sym_messages_authz_proto_rawDescOnce.Do(func() {
		file_sym_messages_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_sym_messages_authz_proto_rawDescData)
	})
	return file_sym_messages_authz_proto_rawDescData
}

var file_sym_messages_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sym_messages_authz_proto_goTypes = []interface{}{
	(*Authz)(nil),         // 0: sym.messages.Authz
	(*AuthzResponse)(nil), // 1: sym.messages.AuthzResponse
	(*models.Schema)(nil), // 2: sym.models.Schema
}
var file_sym_messages_authz_proto_depIdxs = []int32{
	2, // 0: sym.messages.Authz.schema:type_name -> sym.models.Schema
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sym_messages_authz_proto_init() }
func file_sym_messages_authz_proto_init() {
	if File_sym_messages_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sym_messages_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authz); i {
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
		file_sym_messages_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzResponse); i {
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
			RawDescriptor: file_sym_messages_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sym_messages_authz_proto_goTypes,
		DependencyIndexes: file_sym_messages_authz_proto_depIdxs,
		MessageInfos:      file_sym_messages_authz_proto_msgTypes,
	}.Build()
	File_sym_messages_authz_proto = out.File
	file_sym_messages_authz_proto_rawDesc = nil
	file_sym_messages_authz_proto_goTypes = nil
	file_sym_messages_authz_proto_depIdxs = nil
}
