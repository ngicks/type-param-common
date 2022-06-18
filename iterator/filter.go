package iterator

type Excluder[T any] struct {
	inner    DeIterator[T]
	excluder func(T) bool
}

func NewExcluder[T any](iter DeIterator[T], excluder func(T) bool) Excluder[T] {
	return Excluder[T]{
		inner:    iter,
		excluder: excluder,
	}
}

func (e Excluder[T]) next(nextFn nextFunc[T]) (next T, ok bool) {
	var v T
	for {
		v, ok = nextFn()
		if !ok {
			return
		}
		if e.excluder(v) {
			continue
		}
		return v, ok
	}
}
func (e Excluder[T]) Next() (next T, ok bool) {
	return e.next(e.inner.Next)
}
func (e Excluder[T]) NextBack() (next T, ok bool) {
	return e.next(e.inner.NextBack)
}

type Selector[T any] struct {
	inner    DeIterator[T]
	selector func(T) bool
}

func NewSelector[T any](iter DeIterator[T], selector func(T) bool) Selector[T] {
	return Selector[T]{
		inner:    iter,
		selector: selector,
	}
}

func (s Selector[T]) next(nexter nextFunc[T]) (next T, ok bool) {
	var v T
	for {
		v, ok = nexter()
		if !ok {
			return
		}
		if s.selector(v) {
			return v, ok
		}
	}
}
func (s Selector[T]) Next() (next T, ok bool) {
	return s.next(s.inner.Next)
}
func (s Selector[T]) NextBack() (next T, ok bool) {
	return s.next(s.inner.NextBack)
}
