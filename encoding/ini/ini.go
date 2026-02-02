package ini

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

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

	// 3. 将结构体或 map 反射到配置对象中
	if v == nil {
		return nil, fmt.Errorf("v is nil")
	}
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return nil, fmt.Errorf("v is nil")
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Struct:
		if err := ini.ReflectFrom(cfg, v); err != nil {
			return nil, err
		}
	case reflect.Map:
		if err := reflectMapToIni(cfg, rv); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported kind: %v", rv.Kind())
	}

	// 4. 将配置对象写入字节缓冲区
	var buf bytes.Buffer
	_, err := cfg.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	// 得到字节流
	dst := buf.Bytes()

	return dst, nil
}

func reflectMapToIni(cfg *ini.File, rv reflect.Value) error {
	if rv.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("map key must be string")
	}
	defaultSection := cfg.Section("DEFAULT")
	for _, key := range rv.MapKeys() {
		sectionName := key.String()
		value := unwrapValue(rv.MapIndex(key))
		if !value.IsValid() {
			continue
		}
		if value.Kind() == reflect.Map {
			section := cfg.Section(sectionName)
			if err := reflectSectionFromMap(section, value); err != nil {
				return err
			}
			continue
		}
		defaultSection.Key(sectionName).SetValue(fmt.Sprint(value.Interface()))
	}
	return nil
}

func reflectSectionFromMap(section *ini.Section, rv reflect.Value) error {
	if rv.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("section map key must be string")
	}
	for _, key := range rv.MapKeys() {
		value := unwrapValue(rv.MapIndex(key))
		if !value.IsValid() {
			continue
		}
		section.Key(key.String()).SetValue(fmt.Sprint(value.Interface()))
	}
	return nil
}

func unwrapValue(v reflect.Value) reflect.Value {
	for v.IsValid() && (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) {
		if v.IsNil() {
			return reflect.Value{}
		}
		v = v.Elem()
	}
	return v
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		PreserveSurroundedQuote:  true, // 设置为 false，保留引号和转义符
		SpaceBeforeInlineComment: true,
	}, data)
	if err != nil {
		fmt.Println("load err:", err)
		return err
	}
	rv := reflect.ValueOf(v)
	// 基础检查：必须是指针且不为空
	if rv.Kind() == reflect.Ptr {
		return cfg.MapTo(v)
	} else if rv.Kind() == reflect.Map {
		// 遍历所有 Section
		for _, section := range cfg.Sections() {
			sectionMap := make(map[string]any)

			for _, key := range section.Keys() {
				val := key.Value()
				keyName := key.Name()
				// 2. 判断是否有双引号包裹
				if strings.HasPrefix(val, `"`) && strings.HasSuffix(val, `"`) {
					// 发现引号：说明用户想强行指定为字符串
					// 我们去掉引号后，直接存入 Map，跳过后续的 bool/int 转换
					sectionMap[keyName] = strings.Trim(val, `"`)
				} else {
					// 无引号：进入自动类型匹配逻辑
					if boolVal, err := strconv.ParseBool(val); err == nil && val != "0" {
						sectionMap[keyName] = boolVal
					} else if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
						sectionMap[keyName] = intVal
					} else {
						// 兜底：普通的字符串
						sectionMap[keyName] = val
					}
				}
			}
			// 注意：INI 默认有一个 "DEFAULT" section
			if section.Name() == "DEFAULT" {
				for key, value := range sectionMap {
					rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
				}
			} else {
				// 将 Section 存入主 Map
				rv.SetMapIndex(reflect.ValueOf(section.Name()), reflect.ValueOf(sectionMap))
			}
		}
		return nil
	} else {
		return fmt.Errorf("v(%v) must be a non-nil pointer", rv.Kind())
	}
}

func (codec) Name() string {
	return Name
}
