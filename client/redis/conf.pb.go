// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: client/redis/conf.proto

package redis

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

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout   *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout  *durationpb.Duration `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	DialTimeout   *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	ExpireTimeout *durationpb.Duration `protobuf:"bytes,6,opt,name=expire_timeout,json=expireTimeout,proto3" json:"expire_timeout,omitempty"`
	Username      string               `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	Password      string               `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Db            int32                `protobuf:"varint,9,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_redis_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_client_redis_conf_proto_msgTypes[0]
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
	return file_client_redis_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Conf) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Conf) GetAddr() string {
	if x != nil {
		return x.Addr
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

func (x *Conf) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Conf) GetExpireTimeout() *durationpb.Duration {
	if x != nil {
		return x.ExpireTimeout
	}
	return nil
}

func (x *Conf) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Conf) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Conf) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

var File_client_redis_conf_proto protoreflect.FileDescriptor

var file_client_redis_conf_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x3c,
	0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c,
	0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64,
	0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x64, 0x62, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x66, 0x61, 0x72, 0x6b, 0x61, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x3b,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_redis_conf_proto_rawDescOnce sync.Once
	file_client_redis_conf_proto_rawDescData = file_client_redis_conf_proto_rawDesc
)

func file_client_redis_conf_proto_rawDescGZIP() []byte {
	file_client_redis_conf_proto_rawDescOnce.Do(func() {
		file_client_redis_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_redis_conf_proto_rawDescData)
	})
	return file_client_redis_conf_proto_rawDescData
}

var file_client_redis_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_redis_conf_proto_goTypes = []interface{}{
	(*Conf)(nil),                // 0: client.redis.Conf
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_client_redis_conf_proto_depIdxs = []int32{
	1, // 0: client.redis.Conf.read_timeout:type_name -> google.protobuf.Duration
	1, // 1: client.redis.Conf.write_timeout:type_name -> google.protobuf.Duration
	1, // 2: client.redis.Conf.dial_timeout:type_name -> google.protobuf.Duration
	1, // 3: client.redis.Conf.expire_timeout:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_client_redis_conf_proto_init() }
func file_client_redis_conf_proto_init() {
	if File_client_redis_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_redis_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_client_redis_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_redis_conf_proto_goTypes,
		DependencyIndexes: file_client_redis_conf_proto_depIdxs,
		MessageInfos:      file_client_redis_conf_proto_msgTypes,
	}.Build()
	File_client_redis_conf_proto = out.File
	file_client_redis_conf_proto_rawDesc = nil
	file_client_redis_conf_proto_goTypes = nil
	file_client_redis_conf_proto_depIdxs = nil
}