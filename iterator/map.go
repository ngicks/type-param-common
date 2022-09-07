package iterator

type Mapper[T, U any] struct {
	inner  SeIterator[T]
	mapper func(T) U
}

func Map[T, U any](iter SeIterator[T], mapper func(T) U) Mapper[T, U] {
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

func (m Mapper[T, U]) ToIterator() Iterator[U] {
	return Iterator[U]{
		SeIterator: m,
	}
}
