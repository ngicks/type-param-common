package set

import (
	"github.com/ngicks/type-param-common/iterator"
	listparam "github.com/ngicks/type-param-common/list-param"
)

// OrderedSet is same as Set but remembers insertion order.
type OrderedSet[T comparable] struct {
	order  *listparam.List[T]
	eleMap map[T]*listparam.Element[T]
}

func NewOrdered[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		order:  listparam.New[T](),
		eleMap: make(map[T]*listparam.Element[T]),
	}
}

func (s *OrderedSet[T]) Len() int {
	return len(s.eleMap)
}

func (s *OrderedSet[T]) Add(v T) {
	_, has := s.eleMap[v]
	if !has {
		ele := s.order.PushBack(v)
		s.eleMap[v] = ele
	}
}
func (s *OrderedSet[T]) Clear() {
	s.order = s.order.Init()
	s.eleMap = make(map[T]*listparam.Element[T])
}
func (s *OrderedSet[T]) Delete(v T) (deleted bool) {
	ele, deleted := s.eleMap[v]
	if deleted {
		s.order.Remove(ele)
		delete(s.eleMap, v)
	}
	return
}

// ForEach iterates over set and invoke f with each elements.
// Order is FIFO. Call of Add with existing v does not change order.
func (s *OrderedSet[T]) ForEach(f func(v T, idx int)) {
	var idx int
	for next := s.order.Front(); next != nil; next = next.Next() {
		v, _ := next.Get()
		f(v, idx)
		idx++
	}
}
func (s *OrderedSet[T]) Has(v T) (has bool) {
	_, has = s.eleMap[v]
	return
}
func (s *OrderedSet[T]) Values() iterator.Iterator[T] {
	sl := make([]T, 0)
	s.ForEach(func(v T, _ int) {
		sl = append(sl, v)
	})
	return iterator.FromSlice(sl)
}
