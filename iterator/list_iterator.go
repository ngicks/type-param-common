package iterator

import "github.com/ngicks/type-param-common/list"

type ListIter[T any] struct {
	listLen  int
	done     bool
	eleFront list.Element[T]
	eleBack  list.Element[T]
}

func FromList[T any](list list.List[T]) *ListIter[T] {
	return &ListIter[T]{
		listLen:  list.Len(),
		eleFront: list.Front(),
		eleBack:  list.Back(),
	}
}

func (li *ListIter[T]) Next() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront.Unwrap() == li.eleBack.Unwrap() {
		li.done = true
	}
	next = li.eleFront.Get()
	ok = true
	li.eleFront = li.eleFront.Next()
	return
}
func (li *ListIter[T]) NextBack() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront.Unwrap() == li.eleBack.Unwrap() {
		li.done = true
	}
	next = li.eleBack.Get()
	ok = true
	li.eleBack = li.eleBack.Prev()
	return
}
func (li *ListIter[T]) Len() int {
	return li.listLen
}
