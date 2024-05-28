package threading

import "sync"

// A RoutineGroup is used to group goroutines together and all wait all goroutines to be done.
type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

// NewRoutineGroup returns a RoutineGroup.
func NewRoutineGroup() *RoutineGroup {
	return new(RoutineGroup)
}

// Run runs the given fn in RoutineGroup.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

// RunSafe runs the given fn in RoutineGroup, and avoid panics.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup) RunSafe(fn func()) {
	g.waitGroup.Add(1)

	GoSafe(func() {
		defer g.waitGroup.Done()
		fn()
	})
}

// Wait waits all running functions to be done.
func (g *RoutineGroup) Wait() {
	g.waitGroup.Wait()
}

// A RoutineGroup0 is used to group goroutines together and all wait all goroutines to be done.
type RoutineGroup0[Arg0 any] struct {
	waitGroup sync.WaitGroup
}

// NewRoutineGroup0 returns a RoutineGroup.
func NewRoutineGroup0[Arg0 any]() *RoutineGroup0[Arg0] {
	return new(RoutineGroup0[Arg0])
}

// Run runs the given fn in RoutineGroup.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup0[Arg0]) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

// RunSafe runs the given fn in RoutineGroup, and avoid panics.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup0[Arg0]) RunSafe(fn func()) {
	g.waitGroup.Add(1)

	GoSafe(func() {
		defer g.waitGroup.Done()
		fn()
	})
}

// Run0 runs the given fn in RoutineGroup.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup0[Arg0]) Run0(fn func(arg0 Arg0), arg0 Arg0) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn(arg0)
	}()
}

// RunSafe0 runs the given fn in RoutineGroup, and avoid panics.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup0[Arg0]) RunSafe0(fn func(arg0 Arg0), arg0 Arg0) {
	g.waitGroup.Add(1)

	GoSafe0[Arg0](func(arg0 Arg0) {
		defer g.waitGroup.Done()
		fn(arg0)
	}, arg0)
}

// Wait waits all running functions to be done.
func (g *RoutineGroup0[Arg0]) Wait() {
	g.waitGroup.Wait()
}

// A RoutineGroup1 is used to group goroutines together and all wait all goroutines to be done.
type RoutineGroup1[Arg0 any, Arg1 any] struct {
	waitGroup sync.WaitGroup
}

// NewRoutineGroup1 returns a RoutineGroup.
func NewRoutineGroup1[Arg0 any, Arg1 any]() *RoutineGroup1[Arg0, Arg1] {
	return new(RoutineGroup1[Arg0, Arg1])
}

// Run runs the given fn in RoutineGroup.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup1[Arg0, Arg1]) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

// RunSafe runs the given fn in RoutineGroup, and avoid panics.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup1[Arg0, Arg1]) RunSafe(fn func()) {
	g.waitGroup.Add(1)

	GoSafe(func() {
		defer g.waitGroup.Done()
		fn()
	})
}

// Run1 runs the given fn in RoutineGroup.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup1[Arg0, Arg1]) Run1(fn func(arg0 Arg0, arg1 Arg1), arg0 Arg0, arg1 Arg1) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn(arg0, arg1)
	}()
}

// RunSafe1 runs the given fn in RoutineGroup, and avoid panics.
// Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup1[Arg0, Arg1]) RunSafe1(fn func(arg0 Arg0, arg1 Arg1), arg0 Arg0, arg1 Arg1) {
	g.waitGroup.Add(1)

	GoSafe0[Arg0](func(arg0 Arg0) {
		defer g.waitGroup.Done()
		fn(arg0, arg1)
	}, arg0)
}

// Wait waits all running functions to be done.
func (g *RoutineGroup1[Arg0, Arg1]) Wait() {
	g.waitGroup.Wait()
}
