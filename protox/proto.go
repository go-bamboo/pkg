package protox

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Message = protoreflect.ProtoMessage

func Marshal(m Message) ([]byte, error) {
	buf, err := proto.Marshal(m)
	if err != nil {
		return nil, WrapError(err)
	}
	return buf, nil
}

func Unmarshal(b []byte, m Message) error {
	if err := proto.Unmarshal(b, m); err != nil {
		return WrapError(err)
	}
	return nil
}
