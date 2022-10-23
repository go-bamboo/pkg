package threading

import (
	"context"
	"sync"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/panjf2000/ants/v2"
)

type Pool struct {
	pool *ants.Pool

	ctx context.Context
	cf  context.CancelFunc
	wg  sync.WaitGroup
}

func NewPool(size int) (*Pool, error) {
	p, err := ants.NewPool(size, ants.WithLogger(log.GetLogger()))
	if err != nil {
		return nil, err
	}
	ctx, cf := context.WithCancel(context.Background())
	return &Pool{
		pool: p,
		ctx:  ctx,
		cf:   cf,
	}, nil
}

func (p *Pool) Submit(task func(ctx context.Context)) {
	p.wg.Add(1)
	p.pool.Submit(func() {
		defer rescue.Recover(func() { p.wg.Done() })
		task(p.ctx)
	})
}

func (p *Pool) Close() {
	p.cf()
	p.wg.Wait()
	p.pool.Release()
}

type PoolWithFunc struct {
	pool *ants.PoolWithFunc

	ctx context.Context
	cf  context.CancelFunc
	wg  sync.WaitGroup
}

func NewPoolWithFunc(size int, pf func(context.Context, interface{})) (pwf *PoolWithFunc, err error) {
	ctx, cf := context.WithCancel(context.Background())
	pwf = &PoolWithFunc{
		ctx: ctx,
		cf:  cf,
	}
	p, err := ants.NewPoolWithFunc(size, func(i interface{}) {
		defer rescue.Recover(func() { pwf.wg.Done() })
		pf(ctx, i)
	}, ants.WithLogger(log.GetLogger()))
	if err != nil {
		return nil, err
	}
	pwf.pool = p
	return
}

func (p *PoolWithFunc) Submit(i interface{}) error {
	p.wg.Add(1)
	if err := p.pool.Invoke(i); err != nil {
		return err
	}
	return nil
}

func (p *PoolWithFunc) Close() {
	p.cf()
	p.wg.Wait()
	p.pool.Release()
}
