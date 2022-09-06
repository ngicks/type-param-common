package set

import (
	"github.com/ngicks/type-param-common/iterator"
	listparam "github.com/ngicks/type-param-common/list-param"
)

// OrderedSet is same as Set but remembers insertion order.
type OrderedSet[T comparable] struct {
	insertionOrder    listparam.List[T]
	insertionOrderMap map[T]listparam.Element[T]
}

func NewOrdered[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		insertionOrder:    listparam.New[T](),
		insertionOrderMap: make(map[T]listparam.Element[T]),
	}
}

func (s *OrderedSet[T]) Len() int {
	return len(s.insertionOrderMap)
}

func (s *OrderedSet[T]) Add(v T) {
	_, has := s.insertionOrderMap[v]
	if !has {
		ele := s.insertionOrder.PushBack(v)
		s.insertionOrderMap[v] = ele
	}
}
func (s *OrderedSet[T]) Clear() {
	s.insertionOrder = s.insertionOrder.Init()
	s.insertionOrderMap = make(map[T]listparam.Element[T])
}
func (s *OrderedSet[T]) Delete(v T) (deleted bool) {
	ele, deleted := s.insertionOrderMap[v]
	if deleted {
		s.insertionOrder.Remove(ele)
		delete(s.insertionOrderMap, v)
	}
	return
}

// ForEach iterates over set and invoke f with each elements.
// Order is FIFO. Call of Add with existing v does not change order.
func (s *OrderedSet[T]) ForEach(f func(v T, idx int)) {
	var idx int
	for next := s.insertionOrder.Front(); next.Unwrap() != nil; next = next.Next() {
		v, _ := next.Get()
		f(v, idx)
		idx++
	}
}
func (s *OrderedSet[T]) Has(v T) (has bool) {
	_, has = s.insertionOrderMap[v]
	return
}
func (s *OrderedSet[T]) Values() iterator.Iterator[T] {
	sl := make([]T, 0)
	s.ForEach(func(v T, _ int) {
		sl = append(sl, v)
	})
	return iterator.Iterator[T]{
		DeIterator: iterator.FromSlice(sl),
	}
}
