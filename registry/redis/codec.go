package redis

import (
	"encoding/json"

	"github.com/go-bamboo/pkg/registry"
)

func marshal(si *registry.ServiceInstance) (string, error) {
	data, err := json.Marshal(si)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unmarshal(data []byte) (*registry.ServiceInstance, error) {
	var si *registry.ServiceInstance
	err := json.Unmarshal(data, &si)
	return si, err
}
