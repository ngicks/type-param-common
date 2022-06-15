package ring

import "container/ring"

type Ring[T any] struct {
	inner *ring.Ring
}

func New[T any](n int) Ring[T] {
	return Ring[T]{
		inner: ring.New(n),
	}
}
func (r Ring[T]) Unwrap() *ring.Ring {
	return r.inner
}
func (r Ring[T]) Get() (v T) {
	if r.inner.Value == nil {
		return
	}
	return r.inner.Value.(T)
}
func (r Ring[T]) Set(v T) {
	r.inner.Value = v
}
func (r Ring[T]) Do(f func(T)) {
	r.inner.Do(func(a any) {
		var zero T
		if a == nil {
			f(zero)
		} else {
			f(a.(T))
		}
	})
}
func (r Ring[T]) Len() int {
	return r.inner.Len()
}
func (r Ring[T]) Link(s Ring[T]) Ring[T] {
	return Ring[T]{
		inner: r.inner.Link(s.inner),
	}
}
func (r Ring[T]) Move(n int) Ring[T] {
	return Ring[T]{
		inner: r.inner.Move(n),
	}
}
func (r Ring[T]) Next() Ring[T] {
	return Ring[T]{
		inner: r.inner.Next(),
	}
}
func (r Ring[T]) Prev() Ring[T] {
	return Ring[T]{
		inner: r.inner.Prev(),
	}
}
func (r Ring[T]) Unlink(n int) Ring[T] {
	return Ring[T]{
		inner: r.inner.Unlink(n),
	}
}
