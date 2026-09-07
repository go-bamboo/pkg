package redis

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-bamboo/pkg/queue"
	"github.com/stretchr/testify/require"
)

func TestGroupStart(t *testing.T) {
	require.Equal(t, "0", groupStart("0"))
	require.Equal(t, "0", groupStart("earliest"))
	require.Equal(t, "$", groupStart(""))
	require.Equal(t, "$", groupStart("latest"))
	require.Equal(t, "$", groupStart("$"))
}

func TestPayloadFromValues(t *testing.T) {
	key, value := payloadFromValues(map[string]interface{}{
		fieldKey:   "k1",
		fieldValue: []byte("v1"),
	})
	require.Equal(t, []byte("k1"), key)
	require.Equal(t, []byte("v1"), value)
}

func TestNewProducerEmptyBrokers(t *testing.T) {
	_, err := NewProducer(&queue.Conf{})
	require.Error(t, err)
	_, err = NewConsumer(&queue.Conf{})
	require.Error(t, err)
}

func redisTestConf(t *testing.T) *queue.Conf {
	t.Helper()
	if testing.Short() {
		t.Skip("skip redis queue tests in short mode")
	}
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	c := &queue.Conf{
		Brokers: []string{addr},
		Group:   "test-" + strings.ReplaceAll(t.Name(), "/", "_"),
		Offset:  "0",
	}
	p, err := NewProducer(c)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := p.(*Producer).pub.Ping(ctx).Err(); err != nil {
		_ = p.Close()
		t.Skipf("redis not available: %v", err)
	}
	_ = p.Close()
	return c
}

func TestPushAndConsume(t *testing.T) {
	c := redisTestConf(t)
	topic := "q-" + strings.ReplaceAll(t.Name(), "/", "_")

	pusher, err := NewProducer(c)
	require.NoError(t, err)
	t.Cleanup(func() { _ = pusher.Close() })

	q, err := NewConsumer(c)
	require.NoError(t, err)
	t.Cleanup(func() { _ = q.Close() })

	got := make(chan struct {
		key, value []byte
	}, 1)
	sub, err := q.Subscribe(topic, func(ctx context.Context, topic string, key, message []byte) error {
		got <- struct {
			key, value []byte
		}{key, message}
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, topic, sub.Topic())
	t.Cleanup(func() { _ = sub.Unsubscribe(true) })

	time.Sleep(200 * time.Millisecond)
	require.NoError(t, pusher.Push(context.Background(), topic, []byte("k"), []byte("hello")))

	select {
	case msg := <-got:
		require.Equal(t, []byte("k"), msg.key)
		require.Equal(t, []byte("hello"), msg.value)
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for message")
	}
}

func TestCreatePusher(t *testing.T) {
	c := redisTestConf(t)
	p, err := queue.CreatePusher(Name, c)
	require.NoError(t, err)
	require.NotNil(t, p)
	require.Equal(t, Name, p.Name())
	t.Cleanup(func() { _ = p.Close() })
}
