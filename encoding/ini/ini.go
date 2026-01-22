package yaml

import (
	"bytes"

	"github.com/go-kratos/kratos/v2/encoding"
	"gopkg.in/ini.v1"
)

// Name is the name registered for the base64 codec.
const Name = "ini"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with base64.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	// 2. 创建一个空的空配置对象
	cfg := ini.Empty()

	// 3. 将结构体反射(Reflect)到配置对象中
	err := ini.ReflectFrom(cfg, v)
	if err != nil {
		return nil, err
	}

	// 4. 将配置对象写入字节缓冲区
	var buf bytes.Buffer
	_, err = cfg.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	// 得到字节流
	dst := buf.Bytes()

	return dst, nil
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	cfg, err := ini.Load(data)
	if err != nil {
		return err
	}

	// 3. 映射到结构体
	err = cfg.MapTo(v)
	if err != nil {
		return err
	}
	return nil
}

func (codec) Name() string {
	return Name
}
