package listparam

import (
	"container/list"
)

type entMap[T any] map[*list.Element]*Element[T]

// ensureGet gets mm[l] and ensures it to have been existing before.
//
//   - returns nil if l is nil.
//   - creates new *Element[T] if mm does not have l.
//     this happens if after calling PushBackList/ PushFrontList.
func (e entMap[T]) ensureGet(l *list.Element) *Element[T] {
	if l == nil {
		return nil
	}
	ele, ok := e[l]
	if !ok {
		ent := &Element[T]{
			inner:  l,
			entMap: e,
		}
		e[l] = ent
		return ent
	}
	return ele
}

// createInsertingElement creates *Element[T] and insert into e,
// returning that *Element[T].
//
//   - returns nil if l is nil.
func (e entMap[T]) createInsertingElement(l *list.Element) *Element[T] {
	if l == nil {
		return nil
	}
	ent := &Element[T]{
		inner:  l,
		entMap: e,
	}
	e[l] = ent
	return ent
}
