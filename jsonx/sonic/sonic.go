package sonic

import (
	"github.com/bytedance/sonic"
	"github.com/go-bamboo/pkg/jsonx"
	"io"
)

type codec struct{}

func init() {
	jsonx.RegisterCodec(codec{})
}

// Marshal marshals v into json bytes.
func (codec) Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

// MarshalToString marshals v into a string.
func (codec) MarshalToString(v interface{}) (string, error) {
	return sonic.MarshalString(v)
}

// Unmarshal unmarshals data bytes into v.
func (codec) Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

// UnmarshalFromString unmarshals v from str.
func (codec) UnmarshalFromString(str string, v interface{}) error {
	return sonic.UnmarshalString(str, v)
}

// UnmarshalFromReader unmarshals v from reader.
func (codec) UnmarshalFromReader(reader io.Reader, v interface{}) error {
	return sonic.ConfigDefault.NewDecoder(reader).Decode(v)
}

func (codec) Get(str, path string) (string, error) {
	n, err := sonic.GetFromString(str, path)
	if err != nil {
		return "", err
	}
	return n.Raw()
}

func (codec) GetBytes(json []byte, path string) ([]byte, error) {
	n, err := sonic.Get(json, path)
	if err != nil {
		return nil, err
	}
	return n.MarshalJSON()
}

func (codec) Name() string {
	return "sonic"
}
