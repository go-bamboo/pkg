package memory

import (
	"context"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/cache"
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// DefaultExpiration is the default duration for items stored in
	// the cache to expire.
	DefaultExpiration time.Duration = 0
	// ErrItemExpired is returned in Cache.Get when the item found in the cache
	// has expired.
	ErrItemExpired error = errors.InternalServer("ErrItemExpired", "item has expired")
	// ErrKeyNotFound is returned in Cache.Get and Cache.Delete when the
	// provided key could not be found in cache.
	ErrKeyNotFound error = errors.InternalServer("ErrKeyNotFound", "key not found in cache")
)

var (
	// DefaultCache is the default cache.
	DefaultCache cache.Cache = NewCache()
)

type memCache struct {
	opts Options
	sync.RWMutex

	items map[string]Item
}

// NewCache returns a new cache.
func NewCache(opts ...Option) *memCache {
	options := NewOptions(opts...)
	items := make(map[string]Item)

	if len(options.Items) > 0 {
		items = options.Items
	}
	return &memCache{
		opts:  options,
		items: items,
	}
}

func (c *memCache) Get(ctx context.Context, key string) (interface{}, time.Time, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, time.Time{}, ErrKeyNotFound
	}
	if item.Expired() {
		return nil, time.Time{}, ErrItemExpired
	}

	return item.Value, time.Unix(0, item.Expiration), nil
}

func (c *memCache) Put(ctx context.Context, key string, val interface{}, d time.Duration) error {
	var e int64
	if d == DefaultExpiration {
		d = c.opts.Expiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}

	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.items[key] = Item{
		Value:      val,
		Expiration: e,
	}

	return nil
}

func (c *memCache) Delete(ctx context.Context, key string) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	_, found := c.items[key]
	if !found {
		return ErrKeyNotFound
	}

	delete(c.items, key)
	return nil
}

func (m *memCache) String() string {
	return "memory"
}
