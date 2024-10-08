// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.27.1
// source: queue/conf.proto

package queue

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Net struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sasl *Net_SASL `protobuf:"bytes,1,opt,name=sasl,proto3" json:"sasl,omitempty"`
	Tls  *Net_TLS  `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (x *Net) Reset() {
	*x = Net{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Net) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Net) ProtoMessage() {}

func (x *Net) ProtoReflect() protoreflect.Message {
	mi := &file_queue_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Net.ProtoReflect.Descriptor instead.
func (*Net) Descriptor() ([]byte, []int) {
	return file_queue_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Net) GetSasl() *Net_SASL {
	if x != nil {
		return x.Sasl
	}
	return nil
}

func (x *Net) GetTls() *Net_TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brokers      []string             `protobuf:"bytes,1,rep,name=brokers,proto3" json:"brokers,omitempty"`
	Net          *Net                 `protobuf:"bytes,2,opt,name=net,proto3" json:"net,omitempty"`
	Broadcast    bool                 `protobuf:"varint,3,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	Namespace    string               `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Topic        string               `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Expression   string               `protobuf:"bytes,6,opt,name=expression,proto3" json:"expression,omitempty"`
	Group        string               `protobuf:"bytes,7,opt,name=group,proto3" json:"group,omitempty"`
	Offset       string               `protobuf:"bytes,8,opt,name=offset,proto3" json:"offset,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,9,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,10,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	Acks         int32                `protobuf:"varint,11,opt,name=acks,proto3" json:"acks,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_queue_conf_proto_msgTypes[1]
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
	return file_queue_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Conf) GetBrokers() []string {
	if x != nil {
		return x.Brokers
	}
	return nil
}

func (x *Conf) GetNet() *Net {
	if x != nil {
		return x.Net
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

func (x *Conf) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Conf) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

func (x *Conf) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Conf) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Conf) GetAcks() int32 {
	if x != nil {
		return x.Acks
	}
	return 0
}

type Net_SASL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable   bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Net_SASL) Reset() {
	*x = Net_SASL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Net_SASL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Net_SASL) ProtoMessage() {}

func (x *Net_SASL) ProtoReflect() protoreflect.Message {
	mi := &file_queue_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Net_SASL.ProtoReflect.Descriptor instead.
func (*Net_SASL) Descriptor() ([]byte, []int) {
	return file_queue_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Net_SASL) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Net_SASL) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Net_SASL) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Net_TLS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable   bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	RootCa   string `protobuf:"bytes,2,opt,name=rootCa,proto3" json:"rootCa,omitempty"`
	Jks      string `protobuf:"bytes,3,opt,name=jks,proto3" json:"jks,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Cert     string `protobuf:"bytes,5,opt,name=cert,proto3" json:"cert,omitempty"`
	Key      string `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Net_TLS) Reset() {
	*x = Net_TLS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Net_TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Net_TLS) ProtoMessage() {}

func (x *Net_TLS) ProtoReflect() protoreflect.Message {
	mi := &file_queue_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Net_TLS.ProtoReflect.Descriptor instead.
func (*Net_TLS) Descriptor() ([]byte, []int) {
	return file_queue_conf_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Net_TLS) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Net_TLS) GetRootCa() string {
	if x != nil {
		return x.RootCa
	}
	return ""
}

func (x *Net_TLS) GetJks() string {
	if x != nil {
		return x.Jks
	}
	return ""
}

func (x *Net_TLS) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Net_TLS) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *Net_TLS) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_queue_conf_proto protoreflect.FileDescriptor

var file_queue_conf_proto_rawDesc = []byte{
	0x0a, 0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x03, 0x4e, 0x65,
	0x74, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x61, 0x73, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x2e, 0x53, 0x41, 0x53, 0x4c,
	0x52, 0x04, 0x73, 0x61, 0x73, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x2e,
	0x54, 0x4c, 0x53, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x1a, 0x4e, 0x0a, 0x04, 0x53, 0x41, 0x53, 0x4c,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x89, 0x01, 0x0a, 0x03, 0x54, 0x4c, 0x53,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x74,
	0x43, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x61,
	0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a,
	0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0xf0, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x03, 0x6e, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x4e, 0x65, 0x74,
	0x52, 0x03, 0x6e, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x3b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queue_conf_proto_rawDescOnce sync.Once
	file_queue_conf_proto_rawDescData = file_queue_conf_proto_rawDesc
)

func file_queue_conf_proto_rawDescGZIP() []byte {
	file_queue_conf_proto_rawDescOnce.Do(func() {
		file_queue_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_conf_proto_rawDescData)
	})
	return file_queue_conf_proto_rawDescData
}

var file_queue_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_queue_conf_proto_goTypes = []interface{}{
	(*Net)(nil),                 // 0: queue.Net
	(*Conf)(nil),                // 1: queue.Conf
	(*Net_SASL)(nil),            // 2: queue.Net.SASL
	(*Net_TLS)(nil),             // 3: queue.Net.TLS
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_queue_conf_proto_depIdxs = []int32{
	2, // 0: queue.Net.sasl:type_name -> queue.Net.SASL
	3, // 1: queue.Net.tls:type_name -> queue.Net.TLS
	0, // 2: queue.Conf.net:type_name -> queue.Net
	4, // 3: queue.Conf.read_timeout:type_name -> google.protobuf.Duration
	4, // 4: queue.Conf.write_timeout:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_queue_conf_proto_init() }
func file_queue_conf_proto_init() {
	if File_queue_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queue_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Net); i {
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
		file_queue_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_queue_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Net_SASL); i {
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
		file_queue_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Net_TLS); i {
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
			RawDescriptor: file_queue_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queue_conf_proto_goTypes,
		DependencyIndexes: file_queue_conf_proto_depIdxs,
		MessageInfos:      file_queue_conf_proto_msgTypes,
	}.Build()
	File_queue_conf_proto = out.File
	file_queue_conf_proto_rawDesc = nil
	file_queue_conf_proto_goTypes = nil
	file_queue_conf_proto_depIdxs = nil
}
