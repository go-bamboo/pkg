package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/store/redis"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

// Option is redis registry option.
type Option func(*Registry)

var (
	registerScript = redis.NewScript(`
local ttl = tonumber(ARGV[2])
if ttl == nil or ttl <= 0 then
  return redis.error_reply('ttl required')
end
redis.call('HSET', KEYS[1], ARGV[3], ARGV[1])
redis.call('HEXPIRE', KEYS[1], ttl, 'FIELDS', 1, ARGV[3])
return 1
`)

	deregisterScript = redis.NewScript(`
redis.call('HDEL', KEYS[1], ARGV[1])
return 1
`)
)

func indexKey(namespace, name string) string {
	return fmt.Sprintf("%s:%s", namespace, name)
}

// WithTimeout with get services timeout option.
func WithTimeout(timeout time.Duration) Option {
	return func(o *Registry) {
		o.timeout = timeout
	}
}

// WithTTL sets how long each registered instance stays in the hash.
func WithTTL(ttl time.Duration) Option {
	return func(o *Registry) {
		o.ttl = ttl
	}
}

// WithNamespace with registry namespace.
func WithNamespace(ns string) Option {
	return func(o *Registry) {
		o.namespace = ns
	}
}

// Registry is redis registry.
type Registry struct {
	rdb       *redis.Client
	namespace string
	timeout   time.Duration
	ttl       time.Duration
	hbMu      sync.RWMutex
	sets      map[string]*registry.ServiceInstance
}

// New creates redis registry.
func New(rdb *redis.Client, opts ...Option) *Registry {
	r := &Registry{
		rdb:       rdb,
		namespace: "microservices",
		timeout:   30 * time.Second,
		ttl:       60 * time.Second,
		sets:      make(map[string]*registry.ServiceInstance),
	}
	for _, o := range opts {
		o(r)
	}
	go r.heartbeat()
	return r
}

// Register register service.
func (r *Registry) Register(ctx context.Context, svc *registry.ServiceInstance) error {
	r.hbMu.Lock()
	r.sets[svc.ID] = svc
	r.hbMu.Unlock()
	value, err := json.Marshal(svc)
	if err != nil {
		return err
	}
	if err = r.write(ctx, svc, value); err != nil {
		return err
	}
	return nil
}

// Update update service.
func (r *Registry) Update(ctx context.Context, svc *registry.ServiceInstance) error {
	r.hbMu.Lock()
	r.sets[svc.ID] = svc
	r.hbMu.Unlock()
	value, err := json.Marshal(svc)
	if err != nil {
		return err
	}
	if err = r.write(ctx, svc, value); err != nil {
		return err
	}
	return nil
}

// Deregister deregister service.
func (r *Registry) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	r.hbMu.Lock()
	delete(r.sets, svc.ID)
	r.hbMu.Unlock()
	_, err := deregisterScript.Run(ctx, &r.rdb.Client, []string{
		indexKey(r.namespace, svc.Name),
	}, svc.ID).Result()
	return err
}

// GetService return service by name.
func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	fields, err := r.rdb.HGetAll(ctx, indexKey(r.namespace, name)).Result()
	if err != nil {
		return nil, err
	}
	services := make([]*registry.ServiceInstance, 0, len(fields))
	for _, raw := range fields {
		var si registry.ServiceInstance
		if err := json.Unmarshal([]byte(raw), &si); err != nil {
			return nil, err
		}
		services = append(services, &si)
	}
	return services, nil
}

// Watch resolve service by name.
func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	return newWatcher(ctx, r.namespace, name, r.rdb)
}

func (r *Registry) write(ctx context.Context, svc *registry.ServiceInstance, value []byte) error {
	sec := int(r.ttl.Seconds())
	if sec < 1 {
		sec = 1
	}
	_, err := registerScript.Run(ctx, &r.rdb.Client, []string{
		indexKey(r.namespace, svc.Name),
	}, value, sec, svc.ID).Result()
	return err
}

func (r *Registry) heartbeat() {
	interval := r.ttl / 3
	if interval < time.Second {
		interval = time.Second
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			r.hbMu.RLock()
			svcs := make([]*registry.ServiceInstance, 0, len(r.sets))
			for _, svc := range r.sets {
				svcs = append(svcs, svc)
			}
			r.hbMu.RUnlock()
			for _, svc := range svcs {
				value, err := json.Marshal(svc)
				if err != nil {
					log.Errorf("redis heartbeat err: %v", err)
					continue
				}
				if err = r.write(context.TODO(), svc, value); err != nil {
					log.Errorf("redis heartbeat err: %v", err)
				}
			}
		}
	}
}
