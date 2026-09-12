package redis

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
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

var (
	registerScript = redis.NewScript(`
local ttl = tonumber(ARGV[2])
if ttl ~= nil and ttl > 0 then
  redis.call('SET', KEYS[1], ARGV[1], 'EX', ttl)
else
  redis.call('SET', KEYS[1], ARGV[1])
end
redis.call('SADD', KEYS[2], ARGV[3])
return 1
`)

	deregisterScript = redis.NewScript(`
redis.call('DEL', KEYS[1])
redis.call('SREM', KEYS[2], ARGV[1])
return 1
`)
)

// WithTimeout with get services timeout option.
func WithTimeout(timeout time.Duration) Option {
	return func(o *Registry) {
		o.timeout = timeout
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
	registry  map[string]*serviceSet
	lock      sync.RWMutex
	timeout   time.Duration
}

// New creates redis registry.
func New(rdb *redis.Client, opts ...Option) *Registry {
	r := &Registry{
		rdb:       rdb,
		namespace: "microservices",
		registry:  make(map[string]*serviceSet),
		timeout:   10 * time.Second,
	}
	for _, o := range opts {
		o(r)
	}
	return r
}

// Register register service.
func (r *Registry) Register(ctx context.Context, svc *registry.ServiceInstance) error {
	value, err := json.Marshal(svc)
	if err != nil {
		return err
	}
	_, err = registerScript.Run(ctx, &r.rdb.Client, []string{
		r.instanceKey(svc.Name, svc.ID),
		r.indexKey(svc.Name),
	}, value, 0, svc.ID).Result()
	return err
}

// Update update service.
func (r *Registry) Update(ctx context.Context, svc *registry.ServiceInstance) error {
	return r.Register(ctx, svc)
}

// Deregister deregister service.
func (r *Registry) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	_, err := deregisterScript.Run(ctx, &r.rdb.Client, []string{
		r.instanceKey(svc.Name, svc.ID),
		r.indexKey(svc.Name),
	}, svc.ID).Result()
	return err
}

// GetService return service by name.
func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	set := r.registry[name]

	getRemote := func() []*registry.ServiceInstance {
		timeoutCtx, cancel := context.WithTimeout(ctx, r.timeout)
		defer cancel()
		services, _, err := r.Service(timeoutCtx, name)
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

	services, fp, err := r.Service(timeoutCtx, ss.serviceName)
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
				tmpService, tmpFp, err := r.Service(timeoutCtx, ss.serviceName)
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

func (r *Registry) instanceKey(name, id string) string {
	return fmt.Sprintf("%s:%s:%s", r.namespace, name, id)
}

func (r *Registry) indexKey(name string) string {
	return fmt.Sprintf("%s:%s", r.namespace, name)
}

func (r *Registry) Service(ctx context.Context, name string) ([]*registry.ServiceInstance, string, error) {
	indexKey := r.indexKey(name)
	ids, err := r.rdb.SMembers(ctx, indexKey).Result()
	if err != nil {
		return nil, "", err
	}
	if len(ids) == 0 {
		return []*registry.ServiceInstance{}, "", nil
	}
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = r.instanceKey(name, id)
	}
	vals, err := r.rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, "", err
	}
	services := make([]*registry.ServiceInstance, 0, len(vals))
	stale := make([]interface{}, 0)
	for i, v := range vals {
		if v == nil {
			stale = append(stale, ids[i])
			continue
		}
		raw, ok := v.(string)
		if !ok {
			stale = append(stale, ids[i])
			continue
		}
		var si registry.ServiceInstance
		if err := json.Unmarshal([]byte(raw), &si); err != nil {
			return nil, "", err
		}
		services = append(services, &si)
	}
	if len(stale) > 0 {
		_ = r.rdb.SRem(ctx, indexKey, stale...).Err()
	}
	return services, fingerprint(services), nil
}

func fingerprint(ss []*registry.ServiceInstance) string {
	if len(ss) == 0 {
		return ""
	}
	parts := make([]string, 0, len(ss))
	for _, s := range ss {
		eps := append([]string(nil), s.Endpoints...)
		sort.Strings(eps)
		parts = append(parts, strings.Join([]string{s.ID, s.Version, strings.Join(eps, ",")}, "|"))
	}
	sort.Strings(parts)
	sum := sha256.Sum256([]byte(strings.Join(parts, ";")))
	return hex.EncodeToString(sum[:])
}
