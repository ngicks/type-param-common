package iterator

func Enumerate[T any](iter SeIterator[T]) Iterator[EnumerateEnt[T]] {
	return Iterator[EnumerateEnt[T]]{
		SeIterator: NewEnumerator(iter),
	}
}

type EnumerateEnt[T any] struct {
	Count int
	Next  T
}

type Enumerator[T any] struct {
	count int
	inner SeIterator[T]
}

func NewEnumerator[T any](iter SeIterator[T]) *Enumerator[T] {
	return &Enumerator[T]{
		inner: iter,
	}
}

func (e *Enumerator[T]) Next() (next EnumerateEnt[T], ok bool) {
	nextInner, ok := e.inner.Next()
	if !ok {
		return EnumerateEnt[T]{}, false
	}
	c := e.count
	e.count++
	return EnumerateEnt[T]{Count: c, Next: nextInner}, true
}
