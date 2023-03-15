package json

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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

func MarshalAny(m interface{}) ([]byte, error) {
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// MarshalAnyToString marshals v into a string.
func MarshalAnyToString(v interface{}) (string, error) {
	data, err := MarshalAny(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Unmarshal unmarshals data bytes into v.
func Unmarshal(data []byte, v proto.Message) error {
	return protojson.Unmarshal(data, v)
}

// UnmarshalAny unmarshals data bytes into v.
func UnmarshalAny(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
