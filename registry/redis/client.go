package redis

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/store/redis"
)

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

	expireScript = redis.NewScript(`
return redis.call('EXPIRE', KEYS[1], ARGV[1])
`)

	deregisterScript = redis.NewScript(`
redis.call('DEL', KEYS[1])
redis.call('SREM', KEYS[2], ARGV[1])
return 1
`)
)

type instRecord struct {
	name  string
	id    string
	value string
}

// Client is redis registry client.
type Client struct {
	rdb                 *redis.Client
	namespace           string
	healthcheckInterval int
	ttl                 time.Duration
	heartbeat           bool

	mu         sync.Mutex
	instances  map[string]*instRecord
	heartbeats map[string]context.CancelFunc
}

func (c *Client) instanceKey(name, id string) string {
	return fmt.Sprintf("%s/%s/%s", c.namespace, name, id)
}

func (c *Client) indexKey(name string) string {
	return fmt.Sprintf("%s/%s", c.namespace, name)
}

func (c *Client) ttlSeconds() int {
	if c.ttl > 0 {
		sec := int(c.ttl.Seconds())
		if sec > 0 {
			return sec
		}
	}
	interval := c.healthcheckInterval
	if interval <= 0 {
		interval = 10
	}
	return interval * 2
}

func (c *Client) put(ctx context.Context, svc *registry.ServiceInstance, value string, ttl int) error {
	_, err := registerScript.Run(ctx, &c.rdb.Client, []string{
		c.instanceKey(svc.Name, svc.ID),
		c.indexKey(svc.Name),
	}, value, ttl, svc.ID).Result()
	return err
}

// Register register service instance to redis.
func (c *Client) Register(ctx context.Context, svc *registry.ServiceInstance, enableHealthCheck bool) error {
	value, err := marshal(svc)
	if err != nil {
		return err
	}
	ttl := 0
	if enableHealthCheck {
		ttl = c.ttlSeconds()
	}
	if err := c.put(ctx, svc, value, ttl); err != nil {
		return err
	}
	c.storeRecord(svc.Name, svc.ID, value)
	if enableHealthCheck && c.heartbeat {
		c.startHeartbeat(svc.ID)
	}
	return nil
}

// Update overwrite a registered instance without restarting heartbeat.
func (c *Client) Update(ctx context.Context, svc *registry.ServiceInstance, enableHealthCheck bool) error {
	value, err := marshal(svc)
	if err != nil {
		return err
	}
	ttl := 0
	if enableHealthCheck {
		ttl = c.ttlSeconds()
	}
	if err := c.put(ctx, svc, value, ttl); err != nil {
		return err
	}
	c.storeRecord(svc.Name, svc.ID, value)
	return nil
}

// Deregister service by service ID.
func (c *Client) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	c.stopHeartbeat(svc.ID)
	_, err := deregisterScript.Run(ctx, &c.rdb.Client, []string{
		c.instanceKey(svc.Name, svc.ID),
		c.indexKey(svc.Name),
	}, svc.ID).Result()
	c.mu.Lock()
	delete(c.instances, svc.ID)
	c.mu.Unlock()
	return err
}

// Service lists healthy (unexpired) instances and a content fingerprint.
func (c *Client) Service(ctx context.Context, name string) ([]*registry.ServiceInstance, string, error) {
	indexKey := c.indexKey(name)
	ids, err := c.rdb.SMembers(ctx, indexKey).Result()
	if err != nil {
		return nil, "", err
	}
	if len(ids) == 0 {
		return []*registry.ServiceInstance{}, "", nil
	}
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = c.instanceKey(name, id)
	}
	vals, err := c.rdb.MGet(ctx, keys...).Result()
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
		si, err := unmarshal([]byte(raw))
		if err != nil {
			return nil, "", err
		}
		services = append(services, si)
	}
	if len(stale) > 0 {
		if err := c.rdb.SRem(ctx, indexKey, stale...).Err(); err != nil {
			log.Errorf("[Redis] cleanup stale index members failed! err=%v", err)
		}
	}
	return services, fingerprint(services), nil
}

func (c *Client) storeRecord(name, id, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.instances[id] = &instRecord{name: name, id: id, value: value}
}

func (c *Client) startHeartbeat(id string) {
	c.mu.Lock()
	if cancel, ok := c.heartbeats[id]; ok {
		cancel()
	}
	ctx, cancel := context.WithCancel(context.Background())
	c.heartbeats[id] = cancel
	c.mu.Unlock()
	go c.heartbeatLoop(ctx, id)
}

func (c *Client) stopHeartbeat(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cancel, ok := c.heartbeats[id]; ok {
		cancel()
		delete(c.heartbeats, id)
	}
}

func (c *Client) heartbeatLoop(ctx context.Context, id string) {
	select {
	case <-time.After(time.Second):
		c.renewOrReRegister(ctx, id)
	case <-ctx.Done():
		return
	}
	interval := c.healthcheckInterval
	if interval <= 0 {
		interval = 10
	}
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.renewOrReRegister(ctx, id)
		case <-ctx.Done():
			return
		}
	}
}

func (c *Client) renewOrReRegister(ctx context.Context, id string) {
	c.mu.Lock()
	rec, ok := c.instances[id]
	ttl := c.ttlSeconds()
	c.mu.Unlock()
	if !ok {
		return
	}
	n, err := expireScript.Run(ctx, &c.rdb.Client, []string{c.instanceKey(rec.name, rec.id)}, ttl).Int()
	if err != nil {
		log.Errorf("[Redis] update ttl heartbeat failed! err=%v", err)
		c.scheduleReRegister(rec, ttl)
		return
	}
	if n == 0 {
		c.scheduleReRegister(rec, ttl)
	}
}

func (c *Client) scheduleReRegister(rec *instRecord, ttl int) {
	delay := time.Duration(rand.Intn(5)) * time.Second
	time.AfterFunc(delay, func() {
		c.mu.Lock()
		cur, ok := c.instances[rec.id]
		c.mu.Unlock()
		if !ok || cur != rec {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := c.put(ctx, &registry.ServiceInstance{Name: rec.name, ID: rec.id}, rec.value, ttl)
		if err != nil {
			log.Errorf("[Redis] re registry service failed! err=%v", err)
			return
		}
		log.Warn("[Redis] re registry of service occurred success")
	})
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
