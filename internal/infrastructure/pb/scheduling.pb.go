// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/scheduling.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DoneSchedulingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DoneSchedulingRequest) Reset() {
	*x = DoneSchedulingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneSchedulingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneSchedulingRequest) ProtoMessage() {}

func (x *DoneSchedulingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneSchedulingRequest.ProtoReflect.Descriptor instead.
func (*DoneSchedulingRequest) Descriptor() ([]byte, []int) {
	return file_proto_scheduling_proto_rawDescGZIP(), []int{0}
}

func (x *DoneSchedulingRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetNextSchedulingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNextSchedulingRequest) Reset() {
	*x = GetNextSchedulingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextSchedulingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextSchedulingRequest) ProtoMessage() {}

func (x *GetNextSchedulingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextSchedulingRequest.ProtoReflect.Descriptor instead.
func (*GetNextSchedulingRequest) Descriptor() ([]byte, []int) {
	return file_proto_scheduling_proto_rawDescGZIP(), []int{1}
}

type NextScheduling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextScheduling) Reset() {
	*x = NextScheduling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextScheduling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextScheduling) ProtoMessage() {}

func (x *NextScheduling) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextScheduling.ProtoReflect.Descriptor instead.
func (*NextScheduling) Descriptor() ([]byte, []int) {
	return file_proto_scheduling_proto_rawDescGZIP(), []int{2}
}

var File_proto_scheduling_proto protoreflect.FileDescriptor

var file_proto_scheduling_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x27, 0x0a, 0x15,
	0x44, 0x6f, 0x6e, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x32, 0xa2, 0x01, 0x0a, 0x11, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x44, 0x6f, 0x6e,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x6e, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x78, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x1c, 0x5a, 0x1a, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scheduling_proto_rawDescOnce sync.Once
	file_proto_scheduling_proto_rawDescData = file_proto_scheduling_proto_rawDesc
)

func file_proto_scheduling_proto_rawDescGZIP() []byte {
	file_proto_scheduling_proto_rawDescOnce.Do(func() {
		file_proto_scheduling_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scheduling_proto_rawDescData)
	})
	return file_proto_scheduling_proto_rawDescData
}

var file_proto_scheduling_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_scheduling_proto_goTypes = []interface{}{
	(*DoneSchedulingRequest)(nil),    // 0: pb.DoneSchedulingRequest
	(*GetNextSchedulingRequest)(nil), // 1: pb.GetNextSchedulingRequest
	(*NextScheduling)(nil),           // 2: pb.NextScheduling
}
var file_proto_scheduling_proto_depIdxs = []int32{
	0, // 0: pb.SchedulingService.DoneScheduling:input_type -> pb.DoneSchedulingRequest
	1, // 1: pb.SchedulingService.GetNextScheduling:input_type -> pb.GetNextSchedulingRequest
	0, // 2: pb.SchedulingService.DoneScheduling:output_type -> pb.DoneSchedulingRequest
	2, // 3: pb.SchedulingService.GetNextScheduling:output_type -> pb.NextScheduling
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scheduling_proto_init() }
func file_proto_scheduling_proto_init() {
	if File_proto_scheduling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scheduling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneSchedulingRequest); i {
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
		file_proto_scheduling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextSchedulingRequest); i {
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
		file_proto_scheduling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextScheduling); i {
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
			RawDescriptor: file_proto_scheduling_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scheduling_proto_goTypes,
		DependencyIndexes: file_proto_scheduling_proto_depIdxs,
		MessageInfos:      file_proto_scheduling_proto_msgTypes,
	}.Build()
	File_proto_scheduling_proto = out.File
	file_proto_scheduling_proto_rawDesc = nil
	file_proto_scheduling_proto_goTypes = nil
	file_proto_scheduling_proto_depIdxs = nil
}