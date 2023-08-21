package registry

import (
	"fmt"
	"github.com/go-bamboo/pkg/registry/core"
)

var globalRegistry = NewRegistry()

type Factory func(c *Conf) (core.Registrar, core.Discovery, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(c *Conf) (core.Registrar, core.Discovery, error)
}

type discoveryRegistry struct {
	discovery map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &discoveryRegistry{
		discovery: map[string]Factory{},
	}
}

func (d *discoveryRegistry) Register(name string, factory Factory) {
	d.discovery[name] = factory
}

func (d *discoveryRegistry) Create(c *Conf) (core.Registrar, core.Discovery, error) {
	factory, ok := d.discovery[c.ProviderType.String()]
	if !ok {
		return nil, nil, fmt.Errorf("discovery %s has not been registered", c.ProviderType.String())
	}

	r, dd, err := factory(c)
	if err != nil {
		return nil, nil, fmt.Errorf("create discovery error: %s", err)
	}
	return r, dd, nil
}

// Register registers one discovery.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a discovery based on `discoveryDSN`.
func Create(c *Conf) (core.Registrar, core.Discovery, error) {
	return globalRegistry.Create(c)
}
