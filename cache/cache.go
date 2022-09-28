package cache

import (
	"context"
	"time"

	"github.com/emberfarkas/pkg/cache/memory"
)

var (
	// DefaultCache is the default cache.
	DefaultCache Cache = memory.NewCache()
)

// Cache is the interface that wraps the cache.
type Cache interface {
	// Get gets a cached value by key.
	Get(ctx context.Context, key string) (interface{}, time.Time, error)
	// Put stores a key-value pair into cache.
	Put(ctx context.Context, key string, val interface{}, d time.Duration) error
	// Delete removes a key from cache.
	Delete(ctx context.Context, key string) error
	// String returns the name of the implementation.
	String() string
}
