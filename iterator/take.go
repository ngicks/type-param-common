package iterator

type NTaker[T any] struct {
	inner SeIterator[T]
	n     int
}

func NewNTaker[T any](iter SeIterator[T], n int) *NTaker[T] {
	return &NTaker[T]{
		inner: iter,
		n:     n,
	}
}

func (t *NTaker[T]) SizeHint() int {
	if lenner, ok := t.inner.(SizeHinter); ok {
		return lenner.SizeHint()
	}
	return -1
}

func (t *NTaker[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T
	v, ok = nextFn()
	if !ok {
		return
	}
	if t.n > 0 {
		t.n--
		return v, ok
	}
	return
}
func (s *NTaker[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}

type WhileTaker[T any] struct {
	inner        SeIterator[T]
	isOutOfWhile bool
	takeIf       func(T) bool
}

func NewWhileTaker[T any](iter SeIterator[T], takeIf func(T) bool) *WhileTaker[T] {
	return &WhileTaker[T]{
		inner:  iter,
		takeIf: takeIf,
	}
}

func (t WhileTaker[T]) SizeHint() int {
	if lenner, ok := t.inner.(SizeHinter); ok {
		return lenner.SizeHint()
	}
	return -1
}

func (t *WhileTaker[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T
	if t.isOutOfWhile {
		return
	}

	v, ok = nextFn()
	if !ok {
		t.isOutOfWhile = true
		return
	}
	if t.takeIf(v) {
		return v, ok
	}
	t.isOutOfWhile = true
	return next, false
}

func (s *WhileTaker[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}
