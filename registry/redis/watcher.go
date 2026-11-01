package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/store/redis"
)

const watchInterval = time.Second

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	ctx         context.Context
	cf          context.CancelFunc
	rdb         *redis.Client
	namespace   string
	serviceName string
	first       bool
	sets        map[string]*registry.ServiceInstance
}

func newWatcher(ctx context.Context, namespace, name string, client *redis.Client) (*watcher, error) {
	ctx, cf := context.WithCancel(ctx)
	w := &watcher{
		ctx:         ctx,
		cf:          cf,
		rdb:         client,
		namespace:   namespace,
		serviceName: name,
		first:       true,
	}
	return w, nil
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	if w.first {
		fields, err := w.rdb.HGetAll(w.ctx, indexKey(w.namespace, w.serviceName)).Result()
		if err != nil {
			return nil, err
		}
		var sets map[string]*registry.ServiceInstance = make(map[string]*registry.ServiceInstance)
		var instances []*registry.ServiceInstance
		for id, value := range fields {
			var si registry.ServiceInstance
			if err := json.Unmarshal([]byte(value), &si); err != nil {
				log.Errorf("watcher Next %s: %v", w.serviceName, err)
				continue
			}
			sets[id] = &si
			instances = append(instances, &si)
		}
		w.sets = sets
		w.first = false
		return instances, nil
	}
	for {
		fields, err := w.rdb.HGetAll(w.ctx, indexKey(w.namespace, w.serviceName)).Result()
		if err != nil {
			return nil, err
		}
		var sets map[string]*registry.ServiceInstance = make(map[string]*registry.ServiceInstance)
		var instances []*registry.ServiceInstance
		for id, value := range fields {
			var si registry.ServiceInstance
			if err := json.Unmarshal([]byte(value), &si); err != nil {
				log.Errorf("watcher Next %s: %v", w.serviceName, err)
				continue
			}
			sets[id] = &si
			instances = append(instances, &si)
		}
		ss, _ := json.Marshal(sets)
		oldSs, _ := json.Marshal(w.sets)
		if string(ss) != string(oldSs) {
			w.sets = sets
			return instances, nil
		}
		time.Sleep(watchInterval)
	}
}

func (w *watcher) Stop() error {
	w.cf()
	return nil
}
