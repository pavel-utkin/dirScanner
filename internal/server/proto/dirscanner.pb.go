// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: internal/server/proto/dirscanner.proto

package dirscanner

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

type DirScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DirScanRequest) Reset() {
	*x = DirScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_proto_dirscanner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirScanRequest) ProtoMessage() {}

func (x *DirScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_proto_dirscanner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirScanRequest.ProtoReflect.Descriptor instead.
func (*DirScanRequest) Descriptor() ([]byte, []int) {
	return file_internal_server_proto_dirscanner_proto_rawDescGZIP(), []int{0}
}

func (x *DirScanRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Files []*File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *DirScanResponse) Reset() {
	*x = DirScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_proto_dirscanner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirScanResponse) ProtoMessage() {}

func (x *DirScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_proto_dirscanner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirScanResponse.ProtoReflect.Descriptor instead.
func (*DirScanResponse) Descriptor() ([]byte, []int) {
	return file_internal_server_proto_dirscanner_proto_rawDescGZIP(), []int{1}
}

func (x *DirScanResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DirScanResponse) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_proto_dirscanner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_proto_dirscanner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_internal_server_proto_dirscanner_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_internal_server_proto_dirscanner_proto protoreflect.FileDescriptor

var file_internal_server_proto_dirscanner_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x24, 0x0a, 0x0e, 0x44, 0x69, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x49, 0x0a, 0x0f, 0x44, 0x69, 0x72, 0x53, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x22, 0x36, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0x47, 0x0a, 0x0a, 0x44, 0x69, 0x72,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x53, 0x63, 0x61, 0x6e, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x72, 0x53, 0x63, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x44, 0x69, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x64, 0x69, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_server_proto_dirscanner_proto_rawDescOnce sync.Once
	file_internal_server_proto_dirscanner_proto_rawDescData = file_internal_server_proto_dirscanner_proto_rawDesc
)

func file_internal_server_proto_dirscanner_proto_rawDescGZIP() []byte {
	file_internal_server_proto_dirscanner_proto_rawDescOnce.Do(func() {
		file_internal_server_proto_dirscanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_server_proto_dirscanner_proto_rawDescData)
	})
	return file_internal_server_proto_dirscanner_proto_rawDescData
}

var file_internal_server_proto_dirscanner_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_server_proto_dirscanner_proto_goTypes = []interface{}{
	(*DirScanRequest)(nil),  // 0: server.DirScanRequest
	(*DirScanResponse)(nil), // 1: server.DirScanResponse
	(*File)(nil),            // 2: server.File
}
var file_internal_server_proto_dirscanner_proto_depIdxs = []int32{
	2, // 0: server.DirScanResponse.files:type_name -> server.File
	0, // 1: server.Dirscanner.Scan:input_type -> server.DirScanRequest
	1, // 2: server.Dirscanner.Scan:output_type -> server.DirScanResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_server_proto_dirscanner_proto_init() }
func file_internal_server_proto_dirscanner_proto_init() {
	if File_internal_server_proto_dirscanner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_server_proto_dirscanner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirScanRequest); i {
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
		file_internal_server_proto_dirscanner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirScanResponse); i {
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
		file_internal_server_proto_dirscanner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_internal_server_proto_dirscanner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_server_proto_dirscanner_proto_goTypes,
		DependencyIndexes: file_internal_server_proto_dirscanner_proto_depIdxs,
		MessageInfos:      file_internal_server_proto_dirscanner_proto_msgTypes,
	}.Build()
	File_internal_server_proto_dirscanner_proto = out.File
	file_internal_server_proto_dirscanner_proto_rawDesc = nil
	file_internal_server_proto_dirscanner_proto_goTypes = nil
	file_internal_server_proto_dirscanner_proto_depIdxs = nil
}
