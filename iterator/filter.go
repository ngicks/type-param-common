package iterator

var _ Iterator[any] = Excluder[any]{}

type Excluder[T any] struct {
	DeIterator[T]
	excluder func(T) bool
}

func Exclude[T any](iter DeIterator[T], excluder func(T) bool) Excluder[T] {
	return Excluder[T]{
		DeIterator: iter,
		excluder:   excluder,
	}
}

func (e Excluder[T]) Next() (next T, ok bool) {
	var v T
	for {
		v, ok = e.DeIterator.Next()
		if !ok {
			return
		}
		if e.excluder(v) {
			continue
		}
		return v, ok
	}
}
func (e Excluder[T]) NextBack() (next T, ok bool) {
	var v T
	for {
		v, ok = e.DeIterator.NextBack()
		if !ok {
			return
		}
		if e.excluder(v) {
			continue
		}
		return v, ok
	}
}
func (e Excluder[T]) Reverse() Reverser[T] {
	return Reverser[T]{e}
}
func (e Excluder[T]) Select(selector func(T) bool) Selector[T] {
	return Selector[T]{e, selector}
}
func (e Excluder[T]) Exclude(excluder func(T) bool) Excluder[T] {
	return Excluder[T]{e, excluder}
}

var _ Iterator[any] = Selector[any]{}

type Selector[T any] struct {
	DeIterator[T]
	selector func(T) bool
}

func Select[T any](iter DeIterator[T], selector func(T) bool) Selector[T] {
	return Selector[T]{
		DeIterator: iter,
		selector:   selector,
	}
}

func (e Selector[T]) Next() (next T, ok bool) {
	var v T
	for {
		v, ok = e.DeIterator.Next()
		if !ok {
			return
		}
		if e.selector(v) {
			return v, ok
		}
	}
}
func (e Selector[T]) NextBack() (next T, ok bool) {
	var v T
	for {
		v, ok = e.DeIterator.NextBack()
		if !ok {
			return
		}
		if e.selector(v) {
			return v, ok
		}
	}
}

func (s Selector[T]) Reverse() Reverser[T] {
	return Reverser[T]{s}
}
func (s Selector[T]) Select(selector func(T) bool) Selector[T] {
	return Selector[T]{s, selector}
}
func (s Selector[T]) Exclude(excluder func(T) bool) Excluder[T] {
	return Excluder[T]{s, excluder}
}
