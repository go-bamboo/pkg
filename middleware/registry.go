package middleware

import (
	"errors"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
)

func init() {

}

type Middleware = middleware.Middleware

type Handler = middleware.Handler

var globalRegistry = NewRegistry()

// ErrNotFound is middleware not found.
var ErrNotFound = errors.New("Middleware has not been registered")

// Factory is a middleware factory.
type Factory func(conf *Conf) (Middleware, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(cfg *Conf) (Middleware, error)
}

type middlewareRegistry struct {
	middleware map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &middlewareRegistry{
		middleware: map[string]Factory{},
	}
}

// Register registers one middleware.
func (p *middlewareRegistry) Register(name string, factory Factory) {
	p.middleware[createFullName(name)] = factory
}

// Create instantiates a middleware based on `cfg`.
func (p *middlewareRegistry) Create(cfg *Conf) (Middleware, error) {
	if method, ok := p.getMiddleware(createFullName(cfg.Name)); ok {
		if cfg.Required {
			// If the middleware is required, it must be created successfully.
			instance, err := method(cfg)
			if err != nil {
				return nil, err
			}
			return instance, nil
		}
		instance, err := method(cfg)
		if err != nil {
			return nil, err
		}
		return instance, nil
	}
	return nil, ErrNotFound
}

func (p *middlewareRegistry) getMiddleware(name string) (Factory, bool) {
	nameLower := strings.ToLower(name)
	middlewareFn, ok := p.middleware[nameLower]
	if ok {
		return middlewareFn, true
	}
	return nil, false
}

func createFullName(name string) string {
	return strings.ToLower("gateway.middleware." + name)
}

// Register registers one middleware.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a middleware based on `cfg`.
func Create(cfg *Conf) (Middleware, error) {
	return globalRegistry.Create(cfg)
}
