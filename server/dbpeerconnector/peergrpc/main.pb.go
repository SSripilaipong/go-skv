// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: server/dbpeerconnector/peergrpc/main.proto

package peergrpc

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

type SubscribeReplicaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdvertisedAddress string `protobuf:"bytes,1,opt,name=advertisedAddress,proto3" json:"advertisedAddress,omitempty"`
}

func (x *SubscribeReplicaRequest) Reset() {
	*x = SubscribeReplicaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_dbpeerconnector_peergrpc_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReplicaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReplicaRequest) ProtoMessage() {}

func (x *SubscribeReplicaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_dbpeerconnector_peergrpc_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReplicaRequest.ProtoReflect.Descriptor instead.
func (*SubscribeReplicaRequest) Descriptor() ([]byte, []int) {
	return file_server_dbpeerconnector_peergrpc_main_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeReplicaRequest) GetAdvertisedAddress() string {
	if x != nil {
		return x.AdvertisedAddress
	}
	return ""
}

type ReplicaUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ReplicaUpdate) Reset() {
	*x = ReplicaUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_dbpeerconnector_peergrpc_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicaUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicaUpdate) ProtoMessage() {}

func (x *ReplicaUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_server_dbpeerconnector_peergrpc_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicaUpdate.ProtoReflect.Descriptor instead.
func (*ReplicaUpdate) Descriptor() ([]byte, []int) {
	return file_server_dbpeerconnector_peergrpc_main_proto_rawDescGZIP(), []int{1}
}

func (x *ReplicaUpdate) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReplicaUpdate) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_server_dbpeerconnector_peergrpc_main_proto protoreflect.FileDescriptor

var file_server_dbpeerconnector_peergrpc_main_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x62, 0x70, 0x65, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x17,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x4f,
	0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x12, 0x18, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x21, 0x5a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x62, 0x70, 0x65, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_dbpeerconnector_peergrpc_main_proto_rawDescOnce sync.Once
	file_server_dbpeerconnector_peergrpc_main_proto_rawDescData = file_server_dbpeerconnector_peergrpc_main_proto_rawDesc
)

func file_server_dbpeerconnector_peergrpc_main_proto_rawDescGZIP() []byte {
	file_server_dbpeerconnector_peergrpc_main_proto_rawDescOnce.Do(func() {
		file_server_dbpeerconnector_peergrpc_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_dbpeerconnector_peergrpc_main_proto_rawDescData)
	})
	return file_server_dbpeerconnector_peergrpc_main_proto_rawDescData
}

var file_server_dbpeerconnector_peergrpc_main_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_dbpeerconnector_peergrpc_main_proto_goTypes = []interface{}{
	(*SubscribeReplicaRequest)(nil), // 0: SubscribeReplicaRequest
	(*ReplicaUpdate)(nil),           // 1: ReplicaUpdate
}
var file_server_dbpeerconnector_peergrpc_main_proto_depIdxs = []int32{
	0, // 0: PeerService.SubscribeReplica:input_type -> SubscribeReplicaRequest
	1, // 1: PeerService.SubscribeReplica:output_type -> ReplicaUpdate
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_dbpeerconnector_peergrpc_main_proto_init() }
func file_server_dbpeerconnector_peergrpc_main_proto_init() {
	if File_server_dbpeerconnector_peergrpc_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_dbpeerconnector_peergrpc_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReplicaRequest); i {
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
		file_server_dbpeerconnector_peergrpc_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicaUpdate); i {
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
			RawDescriptor: file_server_dbpeerconnector_peergrpc_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_dbpeerconnector_peergrpc_main_proto_goTypes,
		DependencyIndexes: file_server_dbpeerconnector_peergrpc_main_proto_depIdxs,
		MessageInfos:      file_server_dbpeerconnector_peergrpc_main_proto_msgTypes,
	}.Build()
	File_server_dbpeerconnector_peergrpc_main_proto = out.File
	file_server_dbpeerconnector_peergrpc_main_proto_rawDesc = nil
	file_server_dbpeerconnector_peergrpc_main_proto_goTypes = nil
	file_server_dbpeerconnector_peergrpc_main_proto_depIdxs = nil
}
