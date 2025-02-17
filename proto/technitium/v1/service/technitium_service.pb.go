// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: technitium_service.proto

package service

import (
	message "github.com/chris-birch/docker-dns-sync/proto/technitium/v1/message"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/empty.proto.

type Empty = emptypb.Empty

var File_technitium_service_proto protoreflect.FileDescriptor

var file_technitium_service_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x74, 0x69, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x74, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x58, 0x0a, 0x11, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x74, 0x69,
	0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x69, 0x74, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x72,
	0x69, 0x73, 0x2d, 0x62, 0x69, 0x72, 0x63, 0x68, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2d,
	0x64, 0x6e, 0x73, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x63, 0x68, 0x6e, 0x69, 0x74, 0x69, 0x75, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_technitium_service_proto_goTypes = []any{
	(*message.DnsRecord)(nil), // 0: technitium.v1.DnsRecord
	(*emptypb.Empty)(nil),     // 1: google.protobuf.Empty
}
var file_technitium_service_proto_depIdxs = []int32{
	0, // 0: technitium.v1.TechnitiumService.ProcessRecord:input_type -> technitium.v1.DnsRecord
	1, // 1: technitium.v1.TechnitiumService.ProcessRecord:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_technitium_service_proto_init() }
func file_technitium_service_proto_init() {
	if File_technitium_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_technitium_service_proto_rawDesc), len(file_technitium_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_technitium_service_proto_goTypes,
		DependencyIndexes: file_technitium_service_proto_depIdxs,
	}.Build()
	File_technitium_service_proto = out.File
	file_technitium_service_proto_goTypes = nil
	file_technitium_service_proto_depIdxs = nil
}
