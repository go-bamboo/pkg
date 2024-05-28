package threading

import (
	"bytes"
	"context"
	"runtime"
	"strconv"

	"github.com/go-bamboo/pkg/rescue"
)

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()) {
	go RunSafe(fn)
}

func GoSafe0[Arg0 any](fn func(arg0 Arg0), arg0 Arg0) {
	go RunSafe0(fn, arg0)
}

func GoSafe1[Arg0 any, Arg1 any](fn func(arg0 Arg0, arg1 Arg1), arg0 Arg0, arg1 Arg1) {
	go RunSafe1(fn, arg0, arg1)
}

// GoSafeCtx runs the given fn using another goroutine, recovers if fn panics with ctx.
func GoSafeCtx(ctx context.Context, fn func()) {
	go RunSafeCtx(ctx, fn)
}

// RoutineId is only for debug, never use it in production.
func RoutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	// if error, just return 0
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}

// RunSafe runs the given fn, recovers if fn panics.
func RunSafe(fn func()) {
	defer rescue.Recover()

	fn()
}

func RunSafe0[Arg0 any](fn func(arg0 Arg0), arg0 Arg0) {
	defer rescue.Recover()

	fn(arg0)
}

func RunSafe1[Arg0 any, Arg1 any](fn func(arg0 Arg0, arg1 Arg1), arg0 Arg0, arg1 Arg1) {
	defer rescue.Recover()

	fn(arg0, arg1)
}

// RunSafeCtx runs the given fn, recovers if fn panics with ctx.
func RunSafeCtx(ctx context.Context, fn func()) {
	defer rescue.RecoverCtx(ctx)

	fn()
}
