// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: heartbeat_pb/heartbeat.proto

package heartbeat_pb

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

type HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bpm      int32  `protobuf:"varint,1,opt,name=bpm,proto3" json:"bpm,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *HeartBeat) Reset() {
	*x = HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_pb_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeat) ProtoMessage() {}

func (x *HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_pb_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeat.ProtoReflect.Descriptor instead.
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return file_heartbeat_pb_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *HeartBeat) GetBpm() int32 {
	if x != nil {
		return x.Bpm
	}
	return 0
}

func (x *HeartBeat) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type HeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heartbeat *HeartBeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *HeartBeatRequest) Reset() {
	*x = HeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_pb_heartbeat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRequest) ProtoMessage() {}

func (x *HeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_pb_heartbeat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_heartbeat_pb_heartbeat_proto_rawDescGZIP(), []int{1}
}

func (x *HeartBeatRequest) GetHeartbeat() *HeartBeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type HeartBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HeartBeatResponse) Reset() {
	*x = HeartBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_pb_heartbeat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatResponse) ProtoMessage() {}

func (x *HeartBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_pb_heartbeat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatResponse.ProtoReflect.Descriptor instead.
func (*HeartBeatResponse) Descriptor() ([]byte, []int) {
	return file_heartbeat_pb_heartbeat_proto_rawDescGZIP(), []int{2}
}

func (x *HeartBeatResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_heartbeat_pb_heartbeat_proto protoreflect.FileDescriptor

var file_heartbeat_pb_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x2f, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x22, 0x39, 0x0a, 0x09,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x70, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x66, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x5f, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x5f, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heartbeat_pb_heartbeat_proto_rawDescOnce sync.Once
	file_heartbeat_pb_heartbeat_proto_rawDescData = file_heartbeat_pb_heartbeat_proto_rawDesc
)

func file_heartbeat_pb_heartbeat_proto_rawDescGZIP() []byte {
	file_heartbeat_pb_heartbeat_proto_rawDescOnce.Do(func() {
		file_heartbeat_pb_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_heartbeat_pb_heartbeat_proto_rawDescData)
	})
	return file_heartbeat_pb_heartbeat_proto_rawDescData
}

var file_heartbeat_pb_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_heartbeat_pb_heartbeat_proto_goTypes = []interface{}{
	(*HeartBeat)(nil),         // 0: heartbeat_pb.HeartBeat
	(*HeartBeatRequest)(nil),  // 1: heartbeat_pb.HeartBeatRequest
	(*HeartBeatResponse)(nil), // 2: heartbeat_pb.HeartBeatResponse
}
var file_heartbeat_pb_heartbeat_proto_depIdxs = []int32{
	0, // 0: heartbeat_pb.HeartBeatRequest.heartbeat:type_name -> heartbeat_pb.HeartBeat
	1, // 1: heartbeat_pb.HeartBeatService.UserHeartBeat:input_type -> heartbeat_pb.HeartBeatRequest
	2, // 2: heartbeat_pb.HeartBeatService.UserHeartBeat:output_type -> heartbeat_pb.HeartBeatResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_heartbeat_pb_heartbeat_proto_init() }
func file_heartbeat_pb_heartbeat_proto_init() {
	if File_heartbeat_pb_heartbeat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heartbeat_pb_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeat); i {
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
		file_heartbeat_pb_heartbeat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatRequest); i {
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
		file_heartbeat_pb_heartbeat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatResponse); i {
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
			RawDescriptor: file_heartbeat_pb_heartbeat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heartbeat_pb_heartbeat_proto_goTypes,
		DependencyIndexes: file_heartbeat_pb_heartbeat_proto_depIdxs,
		MessageInfos:      file_heartbeat_pb_heartbeat_proto_msgTypes,
	}.Build()
	File_heartbeat_pb_heartbeat_proto = out.File
	file_heartbeat_pb_heartbeat_proto_rawDesc = nil
	file_heartbeat_pb_heartbeat_proto_goTypes = nil
	file_heartbeat_pb_heartbeat_proto_depIdxs = nil
}
