// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: client/rocketmq/conf.proto

package rocketmq

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic      string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Expression string `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rocketmq_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_client_rocketmq_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_client_rocketmq_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Topic) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr      string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	AccessKey string   `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey string   `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Channel   string   `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	GroupId   string   `protobuf:"bytes,5,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Topics    []*Topic `protobuf:"bytes,6,rep,name=topics,proto3" json:"topics,omitempty"`
	Broadcast bool     `protobuf:"varint,7,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	Namespace string   `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rocketmq_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rocketmq_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf.ProtoReflect.Descriptor instead.
func (*Conf) Descriptor() ([]byte, []int) {
	return file_client_rocketmq_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Conf) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Conf) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Conf) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Conf) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Conf) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Conf) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Conf) GetBroadcast() bool {
	if x != nil {
		return x.Broadcast
	}
	return false
}

func (x *Conf) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_client_rocketmq_conf_proto protoreflect.FileDescriptor

var file_client_rocketmq_conf_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d,
	0x71, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x27,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x1b, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d,
	0x71, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61, 0x6d,
	0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x3b, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_rocketmq_conf_proto_rawDescOnce sync.Once
	file_client_rocketmq_conf_proto_rawDescData = file_client_rocketmq_conf_proto_rawDesc
)

func file_client_rocketmq_conf_proto_rawDescGZIP() []byte {
	file_client_rocketmq_conf_proto_rawDescOnce.Do(func() {
		file_client_rocketmq_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_rocketmq_conf_proto_rawDescData)
	})
	return file_client_rocketmq_conf_proto_rawDescData
}

var file_client_rocketmq_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_rocketmq_conf_proto_goTypes = []interface{}{
	(*Topic)(nil), // 0: client.rocketmq.Topic
	(*Conf)(nil),  // 1: client.rocketmq.Conf
}
var file_client_rocketmq_conf_proto_depIdxs = []int32{
	0, // 0: client.rocketmq.Conf.topics:type_name -> client.rocketmq.Topic
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_rocketmq_conf_proto_init() }
func file_client_rocketmq_conf_proto_init() {
	if File_client_rocketmq_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_rocketmq_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_client_rocketmq_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conf); i {
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
			RawDescriptor: file_client_rocketmq_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_rocketmq_conf_proto_goTypes,
		DependencyIndexes: file_client_rocketmq_conf_proto_depIdxs,
		MessageInfos:      file_client_rocketmq_conf_proto_msgTypes,
	}.Build()
	File_client_rocketmq_conf_proto = out.File
	file_client_rocketmq_conf_proto_rawDesc = nil
	file_client_rocketmq_conf_proto_goTypes = nil
	file_client_rocketmq_conf_proto_depIdxs = nil
}