package iterator

type NSkipper[T any] struct {
	inner DeIterator[T]
	n     int
}

func NewNSkipper[T any](iter DeIterator[T], n int) *NSkipper[T] {
	return &NSkipper[T]{
		inner: iter,
		n:     n,
	}
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
func (s *NSkipper[T]) NextBack() (next T, ok bool) {
	return s.next(s.inner.NextBack)
}

type WhileSkipper[T any] struct {
	inner        DeIterator[T]
	isOutOfWhile bool
	skipIf       func(T) bool
}

func NewWhileSkipper[T any](iter DeIterator[T], skipIf func(T) bool) *WhileSkipper[T] {
	return &WhileSkipper[T]{
		inner:  iter,
		skipIf: skipIf,
	}
}

func (s *WhileSkipper[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T

	if s.isOutOfWhile {
		return
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
func (s *WhileSkipper[T]) NextBack() (next T, ok bool) {
	return s.next(s.inner.NextBack)
}
