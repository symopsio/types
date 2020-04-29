// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: sym/enums/service.proto

package enums

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

// Service represents a system Sym integrates with
type Service int32

const (
	Service_UNKNOWN Service = 0
	Service_SLACK   Service = 1
	Service_OKTA    Service = 2
	Service_AWS     Service = 3
	Service_CUSTOM  Service = 4
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0: "UNKNOWN",
		1: "SLACK",
		2: "OKTA",
		3: "AWS",
		4: "CUSTOM",
	}
	Service_value = map[string]int32{
		"UNKNOWN": 0,
		"SLACK":   1,
		"OKTA":    2,
		"AWS":     3,
		"CUSTOM":  4,
	}
)

func (x Service) Enum() *Service {
	p := new(Service)
	*p = x
	return p
}

func (x Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service) Descriptor() protoreflect.EnumDescriptor {
	return file_sym_enums_service_proto_enumTypes[0].Descriptor()
}

func (Service) Type() protoreflect.EnumType {
	return &file_sym_enums_service_proto_enumTypes[0]
}

func (x Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service.Descriptor instead.
func (Service) EnumDescriptor() ([]byte, []int) {
	return file_sym_enums_service_proto_rawDescGZIP(), []int{0}
}

var File_sym_enums_service_proto protoreflect.FileDescriptor

var file_sym_enums_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x79, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x79, 0x6d, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x40, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4b, 0x54, 0x41, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x57, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6d, 0x6f, 0x70, 0x73, 0x69, 0x6f, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sym_enums_service_proto_rawDescOnce sync.Once
	file_sym_enums_service_proto_rawDescData = file_sym_enums_service_proto_rawDesc
)

func file_sym_enums_service_proto_rawDescGZIP() []byte {
	file_sym_enums_service_proto_rawDescOnce.Do(func() {
		file_sym_enums_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sym_enums_service_proto_rawDescData)
	})
	return file_sym_enums_service_proto_rawDescData
}

var file_sym_enums_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sym_enums_service_proto_goTypes = []interface{}{
	(Service)(0), // 0: sym.enums.Service
}
var file_sym_enums_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sym_enums_service_proto_init() }
func file_sym_enums_service_proto_init() {
	if File_sym_enums_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sym_enums_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sym_enums_service_proto_goTypes,
		DependencyIndexes: file_sym_enums_service_proto_depIdxs,
		EnumInfos:         file_sym_enums_service_proto_enumTypes,
	}.Build()
	File_sym_enums_service_proto = out.File
	file_sym_enums_service_proto_rawDesc = nil
	file_sym_enums_service_proto_goTypes = nil
	file_sym_enums_service_proto_depIdxs = nil
}
