// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: registry/conf.proto

package registry

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

type Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints   []string             `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	DialTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`
}

func (x *Etcd) Reset() {
	*x = Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etcd) ProtoMessage() {}

func (x *Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_registry_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etcd.ProtoReflect.Descriptor instead.
func (*Etcd) Descriptor() ([]byte, []int) {
	return file_registry_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Etcd) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

type Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DialTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`
}

func (x *Consul) Reset() {
	*x = Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consul) ProtoMessage() {}

func (x *Consul) ProtoReflect() protoreflect.Message {
	mi := &file_registry_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consul.ProtoReflect.Descriptor instead.
func (*Consul) Descriptor() ([]byte, []int) {
	return file_registry_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Consul) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etcd   *Etcd   `protobuf:"bytes,1,opt,name=etcd,proto3" json:"etcd,omitempty"`
	Consul *Consul `protobuf:"bytes,2,opt,name=consul,proto3" json:"consul,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_registry_conf_proto_msgTypes[2]
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
	return file_registry_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Conf) GetEtcd() *Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Conf) GetConsul() *Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

var File_registry_conf_proto protoreflect.FileDescriptor

var file_registry_conf_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x61, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x22, 0x5f, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x22, 0x54, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x22, 0x0a, 0x04, 0x65,
	0x74, 0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x12,
	0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x66, 0x61, 0x72,
	0x6b, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_registry_conf_proto_rawDescOnce sync.Once
	file_registry_conf_proto_rawDescData = file_registry_conf_proto_rawDesc
)

func file_registry_conf_proto_rawDescGZIP() []byte {
	file_registry_conf_proto_rawDescOnce.Do(func() {
		file_registry_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_conf_proto_rawDescData)
	})
	return file_registry_conf_proto_rawDescData
}

var file_registry_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_registry_conf_proto_goTypes = []interface{}{
	(*Etcd)(nil),                // 0: registry.Etcd
	(*Consul)(nil),              // 1: registry.Consul
	(*Conf)(nil),                // 2: registry.Conf
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_registry_conf_proto_depIdxs = []int32{
	3, // 0: registry.Etcd.dialTimeout:type_name -> google.protobuf.Duration
	3, // 1: registry.Consul.dialTimeout:type_name -> google.protobuf.Duration
	0, // 2: registry.Conf.etcd:type_name -> registry.Etcd
	1, // 3: registry.Conf.consul:type_name -> registry.Consul
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_registry_conf_proto_init() }
func file_registry_conf_proto_init() {
	if File_registry_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Etcd); i {
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
		file_registry_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consul); i {
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
		file_registry_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_registry_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registry_conf_proto_goTypes,
		DependencyIndexes: file_registry_conf_proto_depIdxs,
		MessageInfos:      file_registry_conf_proto_msgTypes,
	}.Build()
	File_registry_conf_proto = out.File
	file_registry_conf_proto_rawDesc = nil
	file_registry_conf_proto_goTypes = nil
	file_registry_conf_proto_depIdxs = nil
}
