package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
)

// Marshal marshals v into json bytes.
func Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, WrapJsonError(err)
	}
	return buf, nil
}

// MarshalToString marshals v into a string.
func MarshalToString(v interface{}) (string, error) {
	data, err := Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Unmarshal unmarshals data bytes into v.
func Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(string(data), err)
	}

	return nil
}

// UnmarshalFromString unmarshals v from str.
func UnmarshalFromString(str string, v interface{}) error {
	decoder := json.NewDecoder(strings.NewReader(str))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(str, err)
	}

	return nil
}

// UnmarshalFromReader unmarshals v from reader.
func UnmarshalFromReader(reader io.Reader, v interface{}) error {
	var buf strings.Builder
	teeReader := io.TeeReader(reader, &buf)
	decoder := json.NewDecoder(teeReader)
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(buf.String(), err)
	}

	return nil
}

func unmarshalUseNumber(decoder *json.Decoder, v interface{}) error {
	decoder.UseNumber()
	return decoder.Decode(v)
}

func formatError(v string, err error) error {
	return errors.InternalServer("FormatErr", fmt.Sprintf("string: `%s`, error: `%v`", v, err))
}
