package base64

import (
	"github.com/go-kratos/kratos/v2/encoding"
)

// Name is the name registered for the base64 codec.
const Name = "base64"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with yaml.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return []byte(""), nil
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (codec) Name() string {
	return Name
}
