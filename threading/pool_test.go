package threading

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	p, err := New()
	assert.Nil(t, err)
	defer p.Close()
	p.Submit(func(ctx context.Context) {
		log.Printf("xx")
	})
}
