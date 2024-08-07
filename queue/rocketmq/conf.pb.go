// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: queue/rocketmq/conf.proto

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

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addrs      []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
	AccessKey  string   `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey  string   `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Channel    string   `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	GroupId    string   `protobuf:"bytes,5,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Topic      string   `protobuf:"bytes,6,opt,name=topic,proto3" json:"topic,omitempty"`
	Expression string   `protobuf:"bytes,7,opt,name=expression,proto3" json:"expression,omitempty"`
	Broadcast  bool     `protobuf:"varint,8,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	Namespace  string   `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Conns      int32    `protobuf:"varint,10,opt,name=conns,proto3" json:"conns,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_rocketmq_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_queue_rocketmq_conf_proto_msgTypes[0]
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
	return file_queue_rocketmq_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Conf) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
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

func (x *Conf) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Conf) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
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

func (x *Conf) GetConns() int32 {
	if x != nil {
		return x.Conns
	}
	return 0
}

var File_queue_rocketmq_conf_proto protoreflect.FileDescriptor

var file_queue_rocketmq_conf_proto_rawDesc = []byte{
	0x0a, 0x19, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64,
	0x64, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6e,
	0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x3b, 0x72, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x6d, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queue_rocketmq_conf_proto_rawDescOnce sync.Once
	file_queue_rocketmq_conf_proto_rawDescData = file_queue_rocketmq_conf_proto_rawDesc
)

func file_queue_rocketmq_conf_proto_rawDescGZIP() []byte {
	file_queue_rocketmq_conf_proto_rawDescOnce.Do(func() {
		file_queue_rocketmq_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_rocketmq_conf_proto_rawDescData)
	})
	return file_queue_rocketmq_conf_proto_rawDescData
}

var file_queue_rocketmq_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_queue_rocketmq_conf_proto_goTypes = []interface{}{
	(*Conf)(nil), // 0: queue.rocketmq.Conf
}
var file_queue_rocketmq_conf_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_queue_rocketmq_conf_proto_init() }
func file_queue_rocketmq_conf_proto_init() {
	if File_queue_rocketmq_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queue_rocketmq_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_queue_rocketmq_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queue_rocketmq_conf_proto_goTypes,
		DependencyIndexes: file_queue_rocketmq_conf_proto_depIdxs,
		MessageInfos:      file_queue_rocketmq_conf_proto_msgTypes,
	}.Build()
	File_queue_rocketmq_conf_proto = out.File
	file_queue_rocketmq_conf_proto_rawDesc = nil
	file_queue_rocketmq_conf_proto_goTypes = nil
	file_queue_rocketmq_conf_proto_depIdxs = nil
}
