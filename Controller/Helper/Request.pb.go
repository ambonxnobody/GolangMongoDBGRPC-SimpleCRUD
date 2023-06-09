// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: Request.proto

package Helper

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestOnAllData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit     *int64  `protobuf:"varint,1,opt,name=Limit,proto3,oneof" json:"Limit,omitempty"`
	Page      *int64  `protobuf:"varint,2,opt,name=Page,proto3,oneof" json:"Page,omitempty"`
	Filter    *string `protobuf:"bytes,3,opt,name=Filter,proto3,oneof" json:"Filter,omitempty"`
	UserID    *string `protobuf:"bytes,4,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	IsDeleted *bool   `protobuf:"varint,5,opt,name=IsDeleted,proto3,oneof" json:"IsDeleted,omitempty"`
}

func (x *RequestOnAllData) Reset() {
	*x = RequestOnAllData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestOnAllData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestOnAllData) ProtoMessage() {}

func (x *RequestOnAllData) ProtoReflect() protoreflect.Message {
	mi := &file_Request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestOnAllData.ProtoReflect.Descriptor instead.
func (*RequestOnAllData) Descriptor() ([]byte, []int) {
	return file_Request_proto_rawDescGZIP(), []int{0}
}

func (x *RequestOnAllData) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *RequestOnAllData) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *RequestOnAllData) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

func (x *RequestOnAllData) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *RequestOnAllData) GetIsDeleted() bool {
	if x != nil && x.IsDeleted != nil {
		return *x.IsDeleted
	}
	return false
}

type ResponseOnAllData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousPage int64        `protobuf:"varint,1,opt,name=PreviousPage,proto3" json:"PreviousPage,omitempty"`
	CurrentPage  int64        `protobuf:"varint,2,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	NextPage     int64        `protobuf:"varint,3,opt,name=NextPage,proto3" json:"NextPage,omitempty"`
	Limit        int64        `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset       int64        `protobuf:"varint,5,opt,name=Offset,proto3" json:"Offset,omitempty"`
	TotalPage    int64        `protobuf:"varint,6,opt,name=TotalPage,proto3" json:"TotalPage,omitempty"`
	TotalData    int64        `protobuf:"varint,7,opt,name=TotalData,proto3" json:"TotalData,omitempty"`
	Data         []*anypb.Any `protobuf:"bytes,8,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ResponseOnAllData) Reset() {
	*x = ResponseOnAllData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseOnAllData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseOnAllData) ProtoMessage() {}

func (x *ResponseOnAllData) ProtoReflect() protoreflect.Message {
	mi := &file_Request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseOnAllData.ProtoReflect.Descriptor instead.
func (*ResponseOnAllData) Descriptor() ([]byte, []int) {
	return file_Request_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseOnAllData) GetPreviousPage() int64 {
	if x != nil {
		return x.PreviousPage
	}
	return 0
}

func (x *ResponseOnAllData) GetCurrentPage() int64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *ResponseOnAllData) GetNextPage() int64 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

func (x *ResponseOnAllData) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ResponseOnAllData) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ResponseOnAllData) GetTotalPage() int64 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *ResponseOnAllData) GetTotalData() int64 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

func (x *ResponseOnAllData) GetData() []*anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Request_proto protoreflect.FileDescriptor

var file_Request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x47, 0x52,
	0x50, 0x43, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x52, 0x55, 0x44, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4f, 0x6e, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x49, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52,
	0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x50, 0x61, 0x67, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x49, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x22, 0x89, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4f, 0x6e, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x2f, 0x5a, 0x2d, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44,
	0x42, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x52, 0x55, 0x44, 0x2f,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Request_proto_rawDescOnce sync.Once
	file_Request_proto_rawDescData = file_Request_proto_rawDesc
)

func file_Request_proto_rawDescGZIP() []byte {
	file_Request_proto_rawDescOnce.Do(func() {
		file_Request_proto_rawDescData = protoimpl.X.CompressGZIP(file_Request_proto_rawDescData)
	})
	return file_Request_proto_rawDescData
}

var file_Request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Request_proto_goTypes = []interface{}{
	(*RequestOnAllData)(nil),  // 0: GolangMongoDBGRPCSimpleCRUD.RequestOnAllData
	(*ResponseOnAllData)(nil), // 1: GolangMongoDBGRPCSimpleCRUD.ResponseOnAllData
	(*anypb.Any)(nil),         // 2: google.protobuf.Any
}
var file_Request_proto_depIdxs = []int32{
	2, // 0: GolangMongoDBGRPCSimpleCRUD.ResponseOnAllData.Data:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Request_proto_init() }
func file_Request_proto_init() {
	if File_Request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestOnAllData); i {
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
		file_Request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseOnAllData); i {
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
	file_Request_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Request_proto_goTypes,
		DependencyIndexes: file_Request_proto_depIdxs,
		MessageInfos:      file_Request_proto_msgTypes,
	}.Build()
	File_Request_proto = out.File
	file_Request_proto_rawDesc = nil
	file_Request_proto_goTypes = nil
	file_Request_proto_depIdxs = nil
}
