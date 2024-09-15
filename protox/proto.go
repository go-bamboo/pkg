package protox

import (
	"encoding/base64"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
)

// Codec defines the interface Transport uses to marshal and unmarshals messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v Message) ([]byte, error)
	// MarshalToString returns the wire format of v.
	MarshalToString(v Message) (string, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error
	// UnmarshalFromString parses the wire format into v.
	UnmarshalFromString(str string, v interface{}) error
	// UnmarshalFromReader parses the wire format into v.
	UnmarshalFromReader(reader io.Reader, v interface{}) error
	// Name returns the name of the Codec implementation. The returned string
	// will be used as part of content type in transmission.  The result must be
	// static; the result cannot change between calls.
	Name() string
}

type Message = protoreflect.ProtoMessage

func Marshal(m Message) ([]byte, error) {
	buf, err := proto.Marshal(m)
	if err != nil {
		return nil, WrapError(err)
	}
	return buf, nil
}

func MarshalToString(v Message) (string, error) {
	buf, err := proto.Marshal(v)
	if err != nil {
		return "", WrapError(err)
	}
	return base64.StdEncoding.EncodeToString(buf), nil
}

func Unmarshal(b []byte, m Message) error {
	if err := proto.Unmarshal(b, m); err != nil {
		return WrapError(err)
	}
	return nil
}

func UnmarshalFromString(str string, m Message) error {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return WrapError(err)
	}
	return Unmarshal(b, m)
}

func UnmarshalFromReader(r io.Reader, m Message) error {
	all, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return Unmarshal(all, m)
}
