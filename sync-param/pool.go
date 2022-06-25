package syncparam

import (
	"sync"
)

type Pool[T any] struct {
	inner sync.Pool
}

func NewPool[T any](new func() T) Pool[T] {
	if new == nil {
		return Pool[T]{}
	}
	return Pool[T]{
		inner: sync.Pool{
			New: func() any {
				return new()
			},
		},
	}
}

func (p *Pool[T]) Get() (content T) {
	got := p.inner.Get()
	if got == nil {
		return
	}
	return got.(T)
}
func (p *Pool[T]) Put(x T) {
	p.inner.Put(x)
}
