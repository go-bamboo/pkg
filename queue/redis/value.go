package redis

import "fmt"

const (
	fieldKey   = "key"
	fieldValue = "value"
)

func payloadFromValues(values map[string]interface{}) (key, value []byte) {
	return toBytes(values[fieldKey]), toBytes(values[fieldValue])
}

func toBytes(v interface{}) []byte {
	switch t := v.(type) {
	case nil:
		return nil
	case []byte:
		return t
	case string:
		return []byte(t)
	default:
		return []byte(fmt.Sprint(t))
	}
}
