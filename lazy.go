package lazy

import "sync"

type lazyFunc[T any] func() T

type lazy[T any] struct {
	f lazyFunc[T]
	o sync.Once
	v T
}

// Represents Lazy evaluated value
type Lazy[T any] interface {
	// Returns Lazy value
	Value() T
}

// Creates new Lazy instance
//
// Example:
//
// l := NewLazy(func() int { return 42; })
// val := l.Value() // returns 42
func NewLazy[T any](f lazyFunc[T]) Lazy[T] {
	return &lazy[T]{
		f: f,
		o: sync.Once{},
	}
}

func (l *lazy[T]) Value() T {
	l.o.Do(func() {
		l.v = l.f()
	})

	return l.v
}
