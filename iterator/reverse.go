package iterator

type Reverser[T any] struct {
	inner DeIterator[T]
}

func NewReverser[T any](iter DeIterator[T]) Reverser[T] {
	return Reverser[T]{iter}
}

func (rev Reverser[T]) Next() (next T, ok bool) {
	return rev.inner.NextBack()
}
func (rev Reverser[T]) NextBack() (next T, ok bool) {
	return rev.inner.Next()
}
