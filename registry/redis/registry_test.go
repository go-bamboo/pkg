package redis

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/store/redis"
	"github.com/stretchr/testify/require"
)

func newTestRegistry(t *testing.T, opts ...Option) *Registry {
	t.Helper()
	if testing.Short() {
		t.Skip("skip redis registry tests in short mode")
	}
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	rdb := redis.New(&redis.Conf{Addrs: []string{addr}})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		t.Skipf("redis not available: %v", err)
	}
	t.Cleanup(func() { _ = rdb.Close() })
	defaults := []Option{
		WithHealthCheckInterval(1),
		WithTimeout(2 * time.Second),
	}
	return New(rdb, append(defaults, opts...)...)
}

func testInstance(t *testing.T) *registry.ServiceInstance {
	t.Helper()
	name := strings.ReplaceAll(t.Name(), "/", "_")
	return &registry.ServiceInstance{
		ID:        "inst-" + name,
		Name:      "helloworld-" + name,
		Version:   "v1",
		Metadata:  map[string]string{"region": "local"},
		Endpoints: []string{"grpc://127.0.0.1:9000"},
	}
}

func TestRegisterGetService(t *testing.T) {
	r := newTestRegistry(t, WithHeartbeat(false))
	ctx := context.Background()
	svc := testInstance(t)
	t.Cleanup(func() { _ = r.Deregister(ctx, svc) })

	require.NoError(t, r.Register(ctx, svc))

	got, err := r.GetService(ctx, svc.Name)
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, svc.ID, got[0].ID)
	require.Equal(t, svc.Name, got[0].Name)
	require.Equal(t, svc.Version, got[0].Version)
	require.Equal(t, svc.Endpoints, got[0].Endpoints)
}

func TestUpdate(t *testing.T) {
	r := newTestRegistry(t, WithHeartbeat(false))
	ctx := context.Background()
	svc := testInstance(t)
	t.Cleanup(func() { _ = r.Deregister(ctx, svc) })
	require.NoError(t, r.Register(ctx, svc))

	svc.Version = "v2"
	svc.Metadata["offline"] = "true"
	require.NoError(t, r.Update(ctx, svc))

	got, err := r.GetService(ctx, svc.Name)
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, "v2", got[0].Version)
	require.Equal(t, "true", got[0].Metadata["offline"])
}

func TestDeregister(t *testing.T) {
	r := newTestRegistry(t, WithHeartbeat(false))
	ctx := context.Background()
	svc := testInstance(t)
	require.NoError(t, r.Register(ctx, svc))
	require.NoError(t, r.Deregister(ctx, svc))

	_, err := r.GetService(ctx, svc.Name)
	require.Error(t, err)
}

func TestTTLExpireWithoutHeartbeat(t *testing.T) {
	r := newTestRegistry(t, WithHeartbeat(false), WithHealthCheckInterval(1))
	ctx := context.Background()
	svc := testInstance(t)
	t.Cleanup(func() { _ = r.Deregister(ctx, svc) })
	require.NoError(t, r.Register(ctx, svc))

	got, err := r.GetService(ctx, svc.Name)
	require.NoError(t, err)
	require.Len(t, got, 1)

	time.Sleep(3 * time.Second)

	_, err = r.GetService(ctx, svc.Name)
	require.Error(t, err)
}

func TestHeartbeatRenewsTTL(t *testing.T) {
	r := newTestRegistry(t, WithHealthCheckInterval(1))
	ctx := context.Background()
	svc := testInstance(t)
	t.Cleanup(func() { _ = r.Deregister(ctx, svc) })
	require.NoError(t, r.Register(ctx, svc))

	// TTL is 2s; without heartbeat the key would be gone after 2.5s.
	time.Sleep(2500 * time.Millisecond)

	got, err := r.GetService(ctx, svc.Name)
	require.NoError(t, err)
	require.Len(t, got, 1)
}

func TestWatchRegisterAndDeregister(t *testing.T) {
	r := newTestRegistry(t, WithHeartbeat(false))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	svc := testInstance(t)
	t.Cleanup(func() { _ = r.Deregister(context.Background(), svc) })

	w, err := r.Watch(ctx, svc.Name)
	require.NoError(t, err)
	t.Cleanup(func() { _ = w.Stop() })

	type result struct {
		ss  []*registry.ServiceInstance
		err error
	}
	ch := make(chan result, 8)
	go func() {
		for {
			ss, err := w.Next()
			ch <- result{ss: ss, err: err}
			if err != nil {
				return
			}
		}
	}()

	require.NoError(t, r.Register(ctx, svc))

	found := false
	deadline := time.After(4 * time.Second)
	for !found {
		select {
		case <-deadline:
			t.Fatal("timeout waiting for watch register")
		case res := <-ch:
			require.NoError(t, res.err)
			for _, inst := range res.ss {
				if inst.ID == svc.ID {
					found = true
				}
			}
		}
	}

	require.NoError(t, r.Deregister(ctx, svc))
	empty := false
	deadline = time.After(4 * time.Second)
	for !empty {
		select {
		case <-deadline:
			t.Fatal("timeout waiting for watch deregister")
		case res := <-ch:
			require.NoError(t, res.err)
			if len(res.ss) == 0 {
				empty = true
			}
		}
	}
}

func TestFactoryCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip redis registry tests in short mode")
	}
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	reg, disc, err := Create(&registry.Conf{
		ProviderType:      registry.ProviderType_Redis,
		Endpoints:         []string{addr},
		EnableHealthCheck: true,
		Namespace:         "/testns",
	})
	require.NoError(t, err)
	require.NotNil(t, disc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := disc.(*Registry).cli.rdb.Ping(ctx).Err(); err != nil {
		t.Skipf("redis not available: %v", err)
	}

	svc := testInstance(t)
	bg := context.Background()
	require.NoError(t, reg.Register(bg, svc))
	t.Cleanup(func() { _ = reg.Deregister(bg, svc) })
	got, err := disc.GetService(bg, svc.Name)
	require.NoError(t, err)
	require.Len(t, got, 1)
}

func TestCreateEmptyEndpoints(t *testing.T) {
	_, _, err := Create(&registry.Conf{ProviderType: registry.ProviderType_Redis})
	require.Error(t, err)
}

func TestFingerprintStable(t *testing.T) {
	a := []*registry.ServiceInstance{
		{ID: "b", Version: "1", Endpoints: []string{"grpc://2", "http://1"}},
		{ID: "a", Version: "1", Endpoints: []string{"grpc://1"}},
	}
	b := []*registry.ServiceInstance{
		{ID: "a", Version: "1", Endpoints: []string{"grpc://1"}},
		{ID: "b", Version: "1", Endpoints: []string{"http://1", "grpc://2"}},
	}
	require.Equal(t, fingerprint(a), fingerprint(b))
	c := []*registry.ServiceInstance{
		{ID: "a", Version: "2", Endpoints: []string{"grpc://1"}},
		{ID: "b", Version: "1", Endpoints: []string{"http://1", "grpc://2"}},
	}
	require.NotEqual(t, fingerprint(a), fingerprint(c))
}
