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

type coreRegistry struct {
	cores map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &coreRegistry{
		cores: map[string]Factory{},
	}
}

func (d *coreRegistry) Register(name string, factory Factory) {
	d.cores[name] = factory
}

func (d *coreRegistry) Create(c *Conf) (core.Logger, error) {
	factory, ok := d.cores[c.Type.String()]
	if !ok {
		return nil, fmt.Errorf("cores %s has not been registered", c.Type.String())
	}
	impl, err := factory(c)
	if err != nil {
		return nil, fmt.Errorf("create cores error: %s", err)
	}
	return impl, nil
}

// Register registers one cores.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a cores based on `discoveryDSN`.
func Create(c *Conf) (core.Logger, error) {
	return globalRegistry.Create(c)
}

// MustCreate instantiates a cores based on `discoveryDSN`.
func MustCreate(c *Conf) core.Logger {
	co, err := globalRegistry.Create(c)
	if err != nil {
		panic(err)
	}
	return co
}
