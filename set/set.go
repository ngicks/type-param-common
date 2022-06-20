package set

import "github.com/ngicks/type-param-common/iterator"

type Set[T comparable] struct {
	inner map[T]struct{}
}

func (s *Set[T]) lazyInit() {
	if s.inner == nil {
		s.inner = make(map[T]struct{})
	}
}

func (s *Set[T]) Len() int {
	s.lazyInit()
	return len(s.inner)
}
func (s *Set[T]) Add(v T) {
	s.lazyInit()
	s.inner[v] = struct{}{}
}
func (s *Set[T]) Clear() {
	s.inner = make(map[T]struct{})
}
func (s *Set[T]) Delete(v T) (deleted bool) {
	s.lazyInit()
	_, deleted = s.inner[v]
	if deleted {
		delete(s.inner, v)
	}
	return
}

// Order is undefined.
func (s *Set[T]) ForEach(f func(v T)) {
	s.lazyInit()
	for v := range s.inner {
		f(v)
	}
}
func (s *Set[T]) Has(v T) (has bool) {
	s.lazyInit()
	_, has = s.inner[v]
	return
}
func (s *Set[T]) Values() iterator.Iterator[T] {
	s.lazyInit()
	sl := make([]T, 0)
	s.ForEach(func(v T) {
		sl = append(sl, v)
	})
	return iterator.Iterator[T]{
		DeIterator: iterator.FromSlice(sl),
	}
}
