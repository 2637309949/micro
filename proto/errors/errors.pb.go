// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/errors/errors.proto

package errors

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code      int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Detail    string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_errors_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_errors_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_errors_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_errors_errors_proto protoreflect.FileDescriptor

var file_proto_errors_errors_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0x6a, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x32, 0x36,
	0x33, 0x37, 0x33, 0x30, 0x39, 0x39, 0x34, 0x39, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x76,
	0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_errors_errors_proto_rawDescOnce sync.Once
	file_proto_errors_errors_proto_rawDescData = file_proto_errors_errors_proto_rawDesc
)

func file_proto_errors_errors_proto_rawDescGZIP() []byte {
	file_proto_errors_errors_proto_rawDescOnce.Do(func() {
		file_proto_errors_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_errors_errors_proto_rawDescData)
	})
	return file_proto_errors_errors_proto_rawDescData
}

var file_proto_errors_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_errors_errors_proto_goTypes = []interface{}{
	(*Error)(nil), // 0: errors.Error
}
var file_proto_errors_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_errors_errors_proto_init() }
func file_proto_errors_errors_proto_init() {
	if File_proto_errors_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_errors_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_proto_errors_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_errors_errors_proto_goTypes,
		DependencyIndexes: file_proto_errors_errors_proto_depIdxs,
		MessageInfos:      file_proto_errors_errors_proto_msgTypes,
	}.Build()
	File_proto_errors_errors_proto = out.File
	file_proto_errors_errors_proto_rawDesc = nil
	file_proto_errors_errors_proto_goTypes = nil
	file_proto_errors_errors_proto_depIdxs = nil
}
