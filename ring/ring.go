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

// Unwrap returns internal *`container/rinng`.Ring.
// Setting non-T value may cause runtime panic in succeeding Get or Do call.
func (r Ring[T]) Unwrap() *ring.Ring {
	return r.inner
}

// Get returns internal Value. If internal Value is non-nil and then returns value and true.
// Otherwise returns zero of T and false.
func (r Ring[T]) Get() (v T, ok bool) {
	if r.inner.Value == nil {
		return
	}
	return r.inner.Value.(T), true
}

// Clear is equivalent to `ring.Value = v`
func (r Ring[T]) Set(v T) {
	r.inner.Value = v
}

// Clear is equivalent to `ring.Value = nil`
func (r Ring[T]) Clear() {
	r.inner.Value = nil
}

// Do is equivalent of `container/ring`.Do but added hasValue boolean.
// hasValue is false if internal Value is nil, indicating passed value is zero of T. hasValue is true othrewize.
func (r Ring[T]) Do(f func(v T, hasValue bool)) {
	r.inner.Do(func(a any) {
		if a == nil {
			var zero T
			f(zero, false)
		} else {
			f(a.(T), true)
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
