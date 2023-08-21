package log

import (
	"fmt"
	"github.com/go-bamboo/pkg/log/core"
)

var globalRegistry = NewRegistry()

type Factory func(c *Conf) (core.Logger, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(c *Conf) (core.Logger, error)
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

func (d *discoveryRegistry) Create(c *Conf) (core.Logger, error) {
	factory, ok := d.discovery[c.Type.String()]
	if !ok {
		return nil, fmt.Errorf("discovery %s has not been registered", c.Type.String())
	}

	impl, err := factory(c)
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
func Create(c *Conf) (core.Logger, error) {
	return globalRegistry.Create(c)
}

// MustCreate instantiates a discovery based on `discoveryDSN`.
func MustCreate(c *Conf) core.Logger {
	co, err := globalRegistry.Create(c)
	if err != nil {
		panic(err)
	}
	return co
}
