package iterator

var _ Iterator[any] = Reverser[any]{}

type Reverser[T any] struct {
	DeIterator[T]
}

func (rev Reverser[T]) Next() (next T, ok bool) {
	return rev.DeIterator.NextBack()
}
func (rev Reverser[T]) NextBack() (next T, ok bool) {
	return rev.DeIterator.Next()
}

func (e Reverser[T]) Reverse() Reverser[T] {
	return Reverser[T]{e}
}
func (rev Reverser[T]) Select(selector func(T) bool) Selector[T] {
	return Selector[T]{
		DeIterator: rev,
		selector:   selector,
	}
}
func (rev Reverser[T]) Exclude(excluder func(T) bool) Excluder[T] {
	return Excluder[T]{
		DeIterator: rev,
		excluder:   excluder,
	}
}
