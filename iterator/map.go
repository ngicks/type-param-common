package iterator

func Map[T, U any](iter SeIterator[T], mapper func(T) U) Iterator[U] {
	return Iterator[U]{
		SeIterator: Mapper[T, U]{
			inner:  iter,
			mapper: mapper,
		},
	}
}

// Mapper applies mapper function.
type Mapper[T, U any] struct {
	inner  SeIterator[T]
	mapper func(T) U
}

func NewMapper[T, U any](iter SeIterator[T], mapper func(T) U) Mapper[T, U] {
	return Mapper[T, U]{
		inner:  iter,
		mapper: mapper,
	}
}

func (m Mapper[T, U]) next(nextFn nextFunc[T]) (next U, ok bool) {
	v, ok := nextFn()
	if ok {
		return m.mapper(v), ok
	}
	return
}
func (m Mapper[T, U]) Next() (next U, ok bool) {
	return m.next(m.inner.Next)
}

// SameTyMapper applies mapper function that returns value of same type to input.
type SameTyMapper[T any] struct {
	inner  SeIterator[T]
	mapper func(T) T
}

func NewSameTyMapper[T any](iter SeIterator[T], mapper func(T) T) SameTyMapper[T] {
	return SameTyMapper[T]{
		inner:  iter,
		mapper: mapper,
	}
}

func (m SameTyMapper[T]) Next() (next T, ok bool) {
	v, ok := m.inner.Next()
	if ok {
		return m.mapper(v), ok
	}
	return
}
