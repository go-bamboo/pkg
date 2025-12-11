package queue

import (
	"fmt"
)

var globalRegistry = NewRegistry()

type ConsumerFactory func(c *Conf) (MessageQueue, error)
type PusherFactory func(c *Conf) (Pusher, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	RegisterConsumer(name string, factory ConsumerFactory)
	CreateConsumer(name string, c *Conf) (MessageQueue, error)
	RegisterPusher(name string, factory PusherFactory)
	CreatePusher(name string, c *Conf) (Pusher, error)
}

type discoveryRegistry struct {
	c map[string]ConsumerFactory
	p map[string]PusherFactory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &discoveryRegistry{
		c: map[string]ConsumerFactory{},
		p: map[string]PusherFactory{},
	}
}

func (d *discoveryRegistry) RegisterConsumer(name string, factory ConsumerFactory) {
	d.c[name] = factory
}

func (d *discoveryRegistry) CreateConsumer(name string, c *Conf) (MessageQueue, error) {
	factory, ok := d.c[name]
	if !ok {
		return nil, fmt.Errorf("provider %s has not been registered", name)
	}

	consumer, err := factory(c)
	if err != nil {
		return nil, fmt.Errorf("create provider error: %s", err)
	}
	return consumer, nil
}

func (d *discoveryRegistry) RegisterPusher(name string, factory PusherFactory) {
	d.p[name] = factory
}

func (d *discoveryRegistry) CreatePusher(name string, c *Conf) (Pusher, error) {
	return nil, nil
}

// RegisterConsumer registers one discovery.
func RegisterConsumer(name string, factory ConsumerFactory) {
	globalRegistry.RegisterConsumer(name, factory)
}

// CreateConsumer instantiates a discovery based on `discoveryDSN`.
func CreateConsumer(name string, c *Conf) (MessageQueue, error) {
	return globalRegistry.CreateConsumer(name, c)
}

func RegisterPusher(name string, factory PusherFactory) {
	globalRegistry.RegisterPusher(name, factory)
}

func CreatePusher(name string, c *Conf) (Pusher, error) {
	return globalRegistry.CreatePusher(name, c)
}
