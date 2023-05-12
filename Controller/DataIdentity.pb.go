// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: DataIdentity.proto

package Controller

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

type DataIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedBy string `protobuf:"bytes,1,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	UpdatedBy string `protobuf:"bytes,2,opt,name=UpdatedBy,proto3" json:"UpdatedBy,omitempty"`
	DeletedBy string `protobuf:"bytes,3,opt,name=DeletedBy,proto3" json:"DeletedBy,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt int64  `protobuf:"varint,6,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *DataIdentity) Reset() {
	*x = DataIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataIdentity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataIdentity) ProtoMessage() {}

func (x *DataIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_DataIdentity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataIdentity.ProtoReflect.Descriptor instead.
func (*DataIdentity) Descriptor() ([]byte, []int) {
	return file_DataIdentity_proto_rawDescGZIP(), []int{0}
}

func (x *DataIdentity) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DataIdentity) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DataIdentity) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

func (x *DataIdentity) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DataIdentity) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DataIdentity) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_DataIdentity_proto protoreflect.FileDescriptor

var file_DataIdentity_proto_rawDesc = []byte{
	0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x44, 0x42, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x52, 0x55,
	0x44, 0x22, 0xc2, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x47, 0x52, 0x50, 0x43, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x43, 0x52, 0x55, 0x44, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DataIdentity_proto_rawDescOnce sync.Once
	file_DataIdentity_proto_rawDescData = file_DataIdentity_proto_rawDesc
)

func file_DataIdentity_proto_rawDescGZIP() []byte {
	file_DataIdentity_proto_rawDescOnce.Do(func() {
		file_DataIdentity_proto_rawDescData = protoimpl.X.CompressGZIP(file_DataIdentity_proto_rawDescData)
	})
	return file_DataIdentity_proto_rawDescData
}

var file_DataIdentity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DataIdentity_proto_goTypes = []interface{}{
	(*DataIdentity)(nil), // 0: GolangMongoDBGRPCSimpleCRUD.DataIdentity
}
var file_DataIdentity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DataIdentity_proto_init() }
func file_DataIdentity_proto_init() {
	if File_DataIdentity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DataIdentity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataIdentity); i {
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
			RawDescriptor: file_DataIdentity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DataIdentity_proto_goTypes,
		DependencyIndexes: file_DataIdentity_proto_depIdxs,
		MessageInfos:      file_DataIdentity_proto_msgTypes,
	}.Build()
	File_DataIdentity_proto = out.File
	file_DataIdentity_proto_rawDesc = nil
	file_DataIdentity_proto_goTypes = nil
	file_DataIdentity_proto_depIdxs = nil
}