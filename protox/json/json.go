package json

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
)

type Message = protoreflect.ProtoMessage

var p = protojson.MarshalOptions{
	EmitUnpopulated: true,
}

func Marshal(m Message) ([]byte, error) {
	buf, err := p.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func MarshalToString(v Message) (string, error) {
	buf, err := p.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// Unmarshal unmarshals data bytes into v.
func Unmarshal(data []byte, v proto.Message) error {
	return protojson.Unmarshal(data, v)
}

func UnmarshalFromString(str string, m Message) error {
	return Unmarshal([]byte(str), m)
}

func UnmarshalFromReader(r io.Reader, m Message) error {
	all, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return Unmarshal(all, m)
}
