package iterator

var _ Iterator[any] = Mapper[any, any]{}

type Mapper[T, U any] struct {
	DeIterator[T]
	mapper func(T) U
}

func Map[T, U any](iter DeIterator[T], mapper func(T) U) Mapper[T, U] {
	return Mapper[T, U]{
		DeIterator: iter,
		mapper:     mapper,
	}
}

func (m Mapper[T, U]) Next() (next U, ok bool) {
	v, ok := m.DeIterator.Next()
	if ok {
		return m.mapper(v), ok
	}
	return
}
func (m Mapper[T, U]) NextBack() (next U, ok bool) {
	v, ok := m.DeIterator.NextBack()
	if ok {
		return m.mapper(v), ok
	}
	return
}

func (s Mapper[T, U]) Reverse() Reverser[U] {
	return Reverser[U]{s}
}
func (s Mapper[T, U]) Select(selector func(U) bool) Selector[U] {
	return Selector[U]{s, selector}
}
func (s Mapper[T, U]) Exclude(excluder func(U) bool) Excluder[U] {
	return Excluder[U]{s, excluder}
}
