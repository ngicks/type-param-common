package atomicparam

import (
	"sync/atomic"
)

// Value is type-param version of sync/atomic.Value.
// T is constrained to be complarable,
// because internal atomic.Value checks equality by comparison.
//
// Zero Value is invalid. Value must be initialized with NewValue.
type Value[T comparable] struct {
	inner         atomic.Value
	isInitialized bool
}

// NewValue returns newly created Value with inner value populated as zero of T.
func NewValue[T comparable]() Value[T] {
	var zero T
	val := atomic.Value{}
	val.Store(zero)

	return Value[T]{
		inner:         val,
		isInitialized: true,
	}
}

func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.inner.CompareAndSwap(old, new)
}

// Load loads inner value T atomically.
// If v is not initialized by NewValue, behavior is undefined.
func (v *Value[T]) Load() (val T) {
	loaded := v.inner.Load()
	if !v.isInitialized {
		// Checking if initialized, for consitent behavior.
		panic("value is not initialized: create it with NewValue")
	}
	return loaded.(T)
}

func (v *Value[T]) Store(val T) {
	v.inner.Store(val)
}

// Swap swaps old value with new atomically.
// If v is not initialized by NewValue, behavior is undefined.
func (v *Value[T]) Swap(new T) (old T) {
	loaded := v.inner.Swap(new)
	if !v.isInitialized {
		panic("value is not initialized: create it with NewValue")
	}
	return loaded.(T)
}
