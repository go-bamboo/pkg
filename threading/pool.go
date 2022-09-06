package threading

import (
	"bls/pkg/log"
	"bls/pkg/rescue"
	"context"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type Pool struct {
	pool *ants.Pool

	ctx context.Context
	cf  context.CancelFunc
	wg  sync.WaitGroup
}

func New() (*Pool, error) {
	p, err := ants.NewPool(1000, ants.WithLogger(log.GetLogger()))
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
