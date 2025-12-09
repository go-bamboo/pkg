package config

import (
	"fmt"
	"net/url"

	"github.com/go-kratos/kratos/v2/config"
)

var globalRegistry = NewRegistry()

type Factory func(dsn *url.URL, v interface{}) (config.Config, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(dsn *url.URL, v interface{}) (config.Config, error)
}

type configRegistry struct {
	discovery map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &configRegistry{
		discovery: map[string]Factory{},
	}
}

func (d *configRegistry) Register(name string, factory Factory) {
	d.discovery[name] = factory
}

func (d *configRegistry) Create(dsn *url.URL, v interface{}) (config.Config, error) {

	factory, ok := d.discovery[dsn.Scheme]
	if !ok {
		return nil, fmt.Errorf("discovery %s has not been registered", dsn.Scheme)
	}

	impl, err := factory(dsn, v)
	if err != nil {
		return nil, fmt.Errorf("create discovery error: %s", err)
	}
	return impl, nil
}

// Register registers one discovery.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a discovery based on `discoveryDSN`.
func Create(dsn *url.URL, v interface{}) (config.Config, error) {
	return globalRegistry.Create(dsn, v)
}
