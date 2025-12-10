package base64

import (
	"encoding/base64"

	"github.com/go-kratos/kratos/v2/encoding"
)

// Name is the name registered for the base64 codec.
const Name = "base64"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with base64.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	var dst []byte
	base64.StdEncoding.Encode(dst, v.([]byte))
	return dst, nil
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (codec) Name() string {
	return Name
}
