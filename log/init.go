package log

import (
	"github.com/go-bamboo/pkg/log/core"
	"github.com/go-bamboo/pkg/log/multi"
	"github.com/go-bamboo/pkg/log/sugar"
)

func Init(c []*core.Conf, opts ...sugar.Option) core.Logger {
	hooks := make([]core.Logger, 0)
	for _, conf := range c {
		co, err := core.Create(conf)
		if err != nil {
			Fatal(err)
		}
		hooks = append(hooks, co)
	}
	co, err := multi.NewMultiCore(hooks...)
	if err != nil {
		panic(err)
	}

	// global
	logger := sugar.NewLogger(co, opts...)
	SetLogger(logger)
	return co
}
