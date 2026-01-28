package fs

import (
	"fmt"
)

var globalRegistry = NewRegistry()

type Factory func(c *Conf) (FileStorage, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(c *Conf, name string) (FileStorage, error)
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

func (d *discoveryRegistry) Create(c *Conf, name string) (FileStorage, error) {
	factory, ok := d.discovery[name]
	if !ok {
		return nil, fmt.Errorf("provider %s has not been registered", name)
	}

	impl, err := factory(c)
	if err != nil {
		return nil, fmt.Errorf("create provider error: %s", err)
	}
	return impl, nil
}

// Register registers one discovery.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a discovery based on `discoveryDSN`.
func Create(c *Conf, name string) (FileStorage, error) {
	return globalRegistry.Create(c, name)
}
