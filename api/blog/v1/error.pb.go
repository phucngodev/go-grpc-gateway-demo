// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: blog/v1/error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_BLOG_INVALID_ID ErrorReason = 0
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "BLOG_INVALID_ID",
	}
	ErrorReason_value = map[string]int32{
		"BLOG_INVALID_ID": 0,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_blog_v1_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_blog_v1_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_blog_v1_error_proto_rawDescGZIP(), []int{0}
}

var File_blog_v1_error_proto protoreflect.FileDescriptor

var file_blog_v1_error_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x28, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x4c, 0x4f, 0x47, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa0, 0x45, 0xf4,
	0x03, 0x42, 0x2b, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x1a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_v1_error_proto_rawDescOnce sync.Once
	file_blog_v1_error_proto_rawDescData = file_blog_v1_error_proto_rawDesc
)

func file_blog_v1_error_proto_rawDescGZIP() []byte {
	file_blog_v1_error_proto_rawDescOnce.Do(func() {
		file_blog_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_v1_error_proto_rawDescData)
	})
	return file_blog_v1_error_proto_rawDescData
}

var file_blog_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blog_v1_error_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.blog.v1.ErrorReason
}
var file_blog_v1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blog_v1_error_proto_init() }
func file_blog_v1_error_proto_init() {
	if File_blog_v1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blog_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blog_v1_error_proto_goTypes,
		DependencyIndexes: file_blog_v1_error_proto_depIdxs,
		EnumInfos:         file_blog_v1_error_proto_enumTypes,
	}.Build()
	File_blog_v1_error_proto = out.File
	file_blog_v1_error_proto_rawDesc = nil
	file_blog_v1_error_proto_goTypes = nil
	file_blog_v1_error_proto_depIdxs = nil
}
