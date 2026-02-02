package ini_test

import (
	"github.com/go-bamboo/pkg/encoding/ini"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestIni(t *testing.T) {
	codec := encoding.GetCodec(ini.Name)
	val := map[string]any{
		"key1": "value1",
		"key2": 3,
	}
	data, err := codec.Marshal(val)
	assert.Nil(t, err)
	t.Log(string(data))
	val2 := map[string]any{}
	err = codec.Unmarshal(data, val2)
	assert.Nil(t, err)
	assert.Equal(t, val["key1"], val2["key1"])
}
