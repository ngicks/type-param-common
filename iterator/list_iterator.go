package iterator

import (
	list "github.com/ngicks/type-param-common/list-param"
)

// Doubly ended iterator made from List.
type ListIterDe[T any] struct {
	listLen      int
	done         bool
	advanceFront int
	advanceBack  int
	eleFront     list.Element[T]
	eleBack      list.Element[T]
}

// FromFixedList makes *ListIterDe[T] from list.List[T].
// Range is fixed at the time FromFixedList returns.
// Mutating passed list outside this iterator may cause undefined behavior.
func FromFixedList[T any](list list.List[T]) *ListIterDe[T] {
	return &ListIterDe[T]{
		listLen:  list.Len(),
		eleFront: list.Front(),
		eleBack:  list.Back(),
	}
}

func (li *ListIterDe[T]) Next() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront.Unwrap() == li.eleBack.Unwrap() {
		li.done = true
	}
	next, _ = li.eleFront.Get()
	ok = true
	li.eleFront = li.eleFront.Next()
	li.advanceFront++
	return
}
func (li *ListIterDe[T]) NextBack() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront.Unwrap() == li.eleBack.Unwrap() {
		li.done = true
	}
	next, _ = li.eleBack.Get()
	ok = true
	li.eleBack = li.eleBack.Prev()
	li.advanceBack++
	return
}

// SizeHint hints size of remaining elements.
// Size would be incorrect if and only if new element is inserted
// into between head and tail of the iterator.
func (li *ListIterDe[T]) SizeHint() int {
	return li.listLen - li.advanceFront - li.advanceBack
}

// ListIterSe is monotonic list iterator. It only advances to tail.
// ListIterSe is not fused, its Next might return ok=true after it returns ok=false.
type ListIterSe[T any] struct {
	ele list.Element[T]
}

func FromList[T any](list list.List[T]) *ListIterSe[T] {
	return &ListIterSe[T]{
		ele: list.Front(),
	}
}

func (li *ListIterSe[T]) Next() (next T, ok bool) {
	n := li.ele.Next()
	if n.Unwrap() == nil {
		return
	}
	li.ele = n
	next, _ = n.Get()
	return next, true
}
