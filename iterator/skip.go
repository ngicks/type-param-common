package iterator

type NSkipper[T any] struct {
	inner SeIterator[T]
	n     int
}

func NewNSkipper[T any](iter SeIterator[T], n int) *NSkipper[T] {
	return &NSkipper[T]{
		inner: iter,
		n:     n,
	}
}

func (iter *NSkipper[T]) SizeHint() int {
	if hint, ok := iter.inner.(SizeHinter); ok {
		l := hint.SizeHint()
		if l < 0 {
			return l
		}
		if l > iter.n {
			return l - iter.n
		} else {
			return 0
		}
	}
	return -1
}

func (s *NSkipper[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T
	for {
		v, ok = nextFn()
		if !ok {
			return
		}
		if s.n <= 0 {
			return v, ok
		}
		s.n--
	}
}
func (s *NSkipper[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}

type WhileSkipper[T any] struct {
	inner        SeIterator[T]
	isOutOfWhile bool
	skipIf       func(T) bool
}

func NewWhileSkipper[T any](iter SeIterator[T], skipIf func(T) bool) *WhileSkipper[T] {
	return &WhileSkipper[T]{
		inner:  iter,
		skipIf: skipIf,
	}
}

// SizeHint implements SizeHinter.
func (s WhileSkipper[T]) SizeHint() int {
	return -1
}

func (s *WhileSkipper[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T

	if s.isOutOfWhile {
		return nextFn()
	}

	for {
		v, ok = nextFn()
		if !ok {
			s.isOutOfWhile = true
			return
		}
		if !s.skipIf(v) {
			s.isOutOfWhile = true
			return v, ok
		}
	}
}

func (s *WhileSkipper[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}
