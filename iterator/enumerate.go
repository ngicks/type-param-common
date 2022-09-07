package iterator

type EnumerateEnt[T any] struct {
	Count int
	Next  T
}

func Enumerate[T any](iter SeIterator[T]) *Enumerator[T] {
	return &Enumerator[T]{
		inner: iter,
	}
}

type Enumerator[T any] struct {
	count int
	inner SeIterator[T]
}

func (e *Enumerator[T]) Next() (next *EnumerateEnt[T], ok bool) {
	nextInner, ok := e.inner.Next()
	if !ok {
		return &EnumerateEnt[T]{}, false
	}
	c := e.count
	e.count++
	return &EnumerateEnt[T]{Count: c, Next: nextInner}, true
}
