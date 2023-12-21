package apollo

import (
	"context"

	"github.com/go-bamboo/pkg/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/philchia/agollo/v4"
)

type Option func(*options)

type options struct {
	appID          string
	cluster        string
	namespaceNames []string
	metaAddr       string
	skipLocalCache bool
	logger         agollo.Logger
}

func AppID(appID string) Option {
	return func(c *options) {
		c.appID = appID
	}
}

func Cluster(cluster string) Option {
	return func(c *options) {
		c.cluster = cluster
	}
}

func Namespaces(ns ...string) Option {
	return func(c *options) {
		c.namespaceNames = append(c.namespaceNames, ns...)
	}
}

func MetaAddr(metaAddr string) Option {
	return func(c *options) {
		c.metaAddr = metaAddr
	}
}

func SkipLocalCache() Option {
	return func(c *options) {
		c.skipLocalCache = true
	}
}

// WithLogger with middleware logger.
func WithLogger(logger agollo.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

type Config struct {
	opts   options
	client agollo.Client
}

func NewConfigSource(opts ...Option) config.Source {
	_options := options{
		cluster: "default",
		logger:  log.GetLogger(),
	}
	for _, o := range opts {
		o(&_options)
	}
	c := &agollo.Conf{
		AppID:          _options.appID,
		Cluster:        _options.cluster,
		NameSpaceNames: _options.namespaceNames,
		MetaAddr:       _options.metaAddr,
	}
	if _options.skipLocalCache {
		client := agollo.NewClient(c, agollo.SkipLocalCache(), agollo.WithLogger(_options.logger))
		if err := client.Start(); err != nil {
			panic(err)
		}
		return &Config{client: client, opts: _options}
	} else {
		client := agollo.NewClient(c, agollo.WithLogger(_options.logger))
		if err := client.Start(); err != nil {
			panic(err)
		}
		return &Config{client: client, opts: _options}
	}
}

func (c *Config) Load() ([]*config.KeyValue, error) {
	content := c.client.GetContent(agollo.WithNamespace(c.opts.namespaceNames[0]))
	return []*config.KeyValue{
		{
			Key:    c.opts.namespaceNames[0],
			Value:  []byte(content),
			Format: "yaml",
		},
	}, nil
}

func (c *Config) Watch() (config.Watcher, error) {
	watcher := newWatcher(c.opts.logger, c.opts.namespaceNames[0])
	c.client.OnUpdate(func(ce *agollo.ChangeEvent) {
		watcher.onChange(ce)
	})
	return watcher, nil
}

type Watcher struct {
	namespaceID string
	content     chan string

	context.Context
	cancel context.CancelFunc
}

func newWatcher(logger agollo.Logger, namespaceID string) *Watcher {
	w := &Watcher{
		namespaceID: namespaceID,
		content:     make(chan string),
	}
	ctx, cancel := context.WithCancel(context.Background())
	w.Context = ctx
	w.cancel = cancel
	return w
}

// OnChange 增加变更监控
func (w *Watcher) onChange(ce *agollo.ChangeEvent) {
	if w.namespaceID == ce.Namespace {
		for _, change := range ce.Changes {
			if change.ChangeType == agollo.MODIFY || change.ChangeType == agollo.ADD {
				w.content <- change.NewValue
			}
		}
	}
}

func (w *Watcher) Next() ([]*config.KeyValue, error) {
	select {
	case <-w.Context.Done():
		return nil, nil
	case content := <-w.content:
		return []*config.KeyValue{
			{
				Key:    w.namespaceID,
				Value:  []byte(content),
				Format: "yaml",
			},
		}, nil
	}
}

func (w *Watcher) Stop() error {
	w.cancel()
	return nil
}
