package redis

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/store/redis"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

// Option is redis registry option.
type Option func(*Registry)

// WithHealthCheck with registry health check option.
func WithHealthCheck(enable bool) Option {
	return func(o *Registry) {
		o.enableHealthCheck = enable
	}
}

// WithTimeout with get services timeout option.
func WithTimeout(timeout time.Duration) Option {
	return func(o *Registry) {
		o.timeout = timeout
	}
}

// WithHeartbeat enable or disable heartbeat.
func WithHeartbeat(enable bool) Option {
	return func(o *Registry) {
		if o.cli != nil {
			o.cli.heartbeat = enable
		}
	}
}

// WithHealthCheckInterval with healthcheck interval in seconds.
func WithHealthCheckInterval(interval int) Option {
	return func(o *Registry) {
		if o.cli != nil {
			o.cli.healthcheckInterval = interval
		}
	}
}

// WithNamespace with registry namespace.
func WithNamespace(ns string) Option {
	return func(o *Registry) {
		if o.cli != nil {
			o.cli.namespace = ns
		}
	}
}

// WithTTL with register ttl.
func WithTTL(ttl time.Duration) Option {
	return func(o *Registry) {
		if o.cli != nil {
			o.cli.ttl = ttl
		}
	}
}

// Registry is redis registry.
type Registry struct {
	cli               *Client
	enableHealthCheck bool
	registry          map[string]*serviceSet
	lock              sync.RWMutex
	timeout           time.Duration
}

// New creates redis registry.
func New(rdb *redis.Client, opts ...Option) *Registry {
	r := &Registry{
		registry:          make(map[string]*serviceSet),
		enableHealthCheck: true,
		timeout:           10 * time.Second,
		cli: &Client{
			rdb:                 rdb,
			namespace:           "/microservices",
			healthcheckInterval: 10,
			heartbeat:           true,
			instances:           make(map[string]*instRecord),
			heartbeats:          make(map[string]context.CancelFunc),
		},
	}
	for _, o := range opts {
		o(r)
	}
	return r
}

// Register register service.
func (r *Registry) Register(ctx context.Context, svc *registry.ServiceInstance) error {
	return r.cli.Register(ctx, svc, r.enableHealthCheck)
}

// Update update service.
func (r *Registry) Update(ctx context.Context, svc *registry.ServiceInstance) error {
	return r.cli.Update(ctx, svc, r.enableHealthCheck)
}

// Deregister deregister service.
func (r *Registry) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	return r.cli.Deregister(ctx, svc)
}

// GetService return service by name.
func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	set := r.registry[name]

	getRemote := func() []*registry.ServiceInstance {
		timeoutCtx, cancel := context.WithTimeout(ctx, r.timeout)
		defer cancel()
		services, _, err := r.cli.Service(timeoutCtx, name)
		if err == nil && len(services) > 0 {
			return services
		}
		return nil
	}

	if set == nil {
		if s := getRemote(); len(s) > 0 {
			return s, nil
		}
		return nil, fmt.Errorf("service %s not resolved in registry", name)
	}
	ss, _ := set.services.Load().([]*registry.ServiceInstance)
	if ss == nil {
		if s := getRemote(); len(s) > 0 {
			return s, nil
		}
		return nil, fmt.Errorf("service %s not found in registry", name)
	}
	return ss, nil
}

// Watch resolve service by name.
func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	set, ok := r.registry[name]
	if !ok {
		set = &serviceSet{
			watcher:     make(map[*watcher]struct{}),
			services:    &atomic.Value{},
			serviceName: name,
		}
		r.registry[name] = set
	}

	w := &watcher{
		event: make(chan struct{}, 1),
	}
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.set = set
	set.lock.Lock()
	set.watcher[w] = struct{}{}
	set.lock.Unlock()
	ss, _ := set.services.Load().([]*registry.ServiceInstance)
	if len(ss) > 0 {
		w.event <- struct{}{}
	}

	if !ok {
		err := r.resolve(ctx, set)
		if err != nil {
			set.lock.Lock()
			delete(set.watcher, w)
			set.lock.Unlock()
			w.cancel()
			return nil, err
		}
	}
	return w, nil
}

func (r *Registry) resolve(ctx context.Context, ss *serviceSet) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	services, fp, err := r.cli.Service(timeoutCtx, ss.serviceName)
	if err != nil {
		return err
	}
	ss.broadcast(services)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		idx := fp
		for {
			select {
			case <-ticker.C:
				timeoutCtx, cancel := context.WithTimeout(context.Background(), r.timeout)
				tmpService, tmpFp, err := r.cli.Service(timeoutCtx, ss.serviceName)
				cancel()
				if err != nil {
					time.Sleep(time.Second)
					continue
				}
				if tmpFp != idx {
					ss.broadcast(tmpService)
				}
				idx = tmpFp
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
