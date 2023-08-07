package parallel

import "sync"

type Lazy[T any] struct {
	once    sync.Once
	value   T
	initErr error
	thunk   func() (T, error)
}

func NewLazy[T any](thunk func() (T, error)) *Lazy[T] {
	return &Lazy[T]{
		thunk: thunk,
	}
}

func (z *Lazy[T]) Val() (T, error) {
	z.once.Do(func() {
		z.value, z.initErr = z.thunk()
	})
	return z.value, z.initErr
}
