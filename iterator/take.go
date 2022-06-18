package iterator

type NTaker[T any] struct {
	inner DeIterator[T]
	n     int
}

func NewNTaker[T any](iter DeIterator[T], n int) *NTaker[T] {
	return &NTaker[T]{
		inner: iter,
		n:     n,
	}
}

func (s *NTaker[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T
	v, ok = nextFn()
	if !ok {
		return
	}
	if s.n > 0 {
		s.n--
		return v, ok
	}
	return
}
func (s *NTaker[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}
func (s *NTaker[T]) NextBack() (next T, ok bool) {
	return s.next(s.inner.NextBack)
}

type WhileTaker[T any] struct {
	inner        DeIterator[T]
	isOutOfWhile bool
	takeIf       func(T) bool
}

func NewWhileTaker[T any](iter DeIterator[T], takeIf func(T) bool) *WhileTaker[T] {
	return &WhileTaker[T]{
		inner:  iter,
		takeIf: takeIf,
	}
}

func (s *WhileTaker[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T
	if s.isOutOfWhile {
		return
	}

	v, ok = nextFn()
	if !ok {
		s.isOutOfWhile = true
		return
	}
	if s.takeIf(v) {
		return v, ok
	}
	s.isOutOfWhile = true
	return next, false
}

func (s *WhileTaker[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}
func (s *WhileTaker[T]) NextBack() (next T, ok bool) {
	return s.next(s.inner.NextBack)
}
