package atomicparam

import (
	"sync/atomic"
)

// Value is type-param version of sync/atomic.Value.
// T is constrained to be complarable,
// because internal atomic.Value checks equality by comparison.
//
// Zero Value is invalid, without proceeding Store call.
type Value[T comparable] struct {
	inner atomic.Value
}

// NewValue returns newly created Value with inner value populated as zero of T.
func NewValue[T comparable]() Value[T] {
	var zero T
	val := atomic.Value{}
	val.Store(zero)

	return Value[T]{
		inner: val,
	}
}

// CompareAndSwap does compare-and-swap operation for T.
//
// `swapped` is always false if value is not stored.
// To avoid this, call Store beforehand or use v initialized by NewValue.
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.inner.CompareAndSwap(old, new)
}

// Load loads inner value T atomically.
//
// Load panics if value is not stored. Call Store beforehand or use v initialized by NewValue.
func (v *Value[T]) Load() (val T) {
	return v.inner.Load().(T)
}

func (v *Value[T]) Store(val T) {
	v.inner.Store(val)
}

// Swap swaps old value with new atomically.
//
// Swap panics if value is not stored. Call Store beforehand or use v initialized by NewValue.
func (v *Value[T]) Swap(new T) (old T) {
	return v.inner.Swap(new).(T)
}
