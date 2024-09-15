package jsonx

import (
	"io"
)

// Codec defines the interface Transport uses to marshal and unmarshals messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v interface{}) ([]byte, error)
	// MarshalToString returns the wire format of v.
	MarshalToString(v interface{}) (string, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error
	// UnmarshalFromString parses the wire format into v.
	UnmarshalFromString(str string, v interface{}) error
	// UnmarshalFromReader parses the wire format into v.
	UnmarshalFromReader(reader io.Reader, v interface{}) error
	// Get parses the wire format into v.
	Get(str, path string) (string, error)
	// GetBytes parses the wire format into v.
	GetBytes(json []byte, path string) ([]byte, error)
	// Name returns the name of the Codec implementation. The returned string
	// will be used as part of content type in transmission.  The result must be
	// static; the result cannot change between calls.
	Name() string
}

var registeredCodec Codec

// RegisterCodec registers the provided Codec for use with all Transport clients and
// servers.
func RegisterCodec(codec Codec) {
	if codec == nil {
		panic("cannot register a nil Codec")
	}
	registeredCodec = codec
}

// GetCodec gets a registered Codec by content-subtype, or nil if no Codec is
// registered for the content-subtype.
//
// The content-subtype is expected to be lowercase.
func GetCodec() Codec {
	return registeredCodec
}

// Marshal marshals v into json bytes.
func Marshal(v interface{}) ([]byte, error) {
	return registeredCodec.Marshal(v)
}

// MarshalToString marshals v into a string.
func MarshalToString(v interface{}) (string, error) {
	return registeredCodec.MarshalToString(v)
}

// Unmarshal unmarshals data bytes into v.
func Unmarshal(data []byte, v interface{}) error {
	return registeredCodec.Unmarshal(data, v)
}

// UnmarshalFromString unmarshals v from str.
func UnmarshalFromString(str string, v interface{}) error {
	return registeredCodec.UnmarshalFromString(str, v)
}

// UnmarshalFromReader unmarshals v from reader.
func UnmarshalFromReader(reader io.Reader, v interface{}) error {
	return registeredCodec.UnmarshalFromReader(reader, v)
}

func Get(json, path string) (string, error) {
	return registeredCodec.Get(json, path)
}

func GetBytes(json []byte, path string) ([]byte, error) {
	return registeredCodec.GetBytes(json, path)
}
