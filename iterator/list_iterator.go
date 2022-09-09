package iterator

import (
	listparam "github.com/ngicks/type-param-common/list-param"
)

// Doubly ended iterator made from List.
type ListIterDe[T any] struct {
	listLen      int
	done         bool
	advanceFront int
	advanceBack  int
	eleFront     *listparam.Element[T]
	eleBack      *listparam.Element[T]
}

// FromFixedList makes *ListIterDe[T] from list.List[T].
// Range is fixed at the time FromFixedList returns.
// Mutating passed list outside this iterator may cause undefined behavior.
func FromFixedList[T any](list *listparam.List[T]) *ListIterDe[T] {
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
	if li.eleFront == li.eleBack {
		li.done = true
	}
	if li.eleFront == nil {
		return
	}
	next, ok = li.eleFront.Get()
	if !ok {
		return
	}
	li.eleFront = li.eleFront.Next()
	li.advanceFront++
	return
}
func (li *ListIterDe[T]) NextBack() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront == li.eleBack {
		li.done = true
	}
	if li.eleBack == nil {
		return
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

func (li *ListIterDe[T]) ToIterator() Iterator[T] {
	return Iterator[T]{li}
}

// ListIterSe is monotonic list iterator. It only advances to tail.
// ListIterSe is not fused, its Next might return ok=true after it returns ok=false.
// This happens when passed list grows its tail afterwards.
type ListIterSe[T any] struct {
	root     *listparam.List[T]
	ele      *listparam.Element[T]
	advanced bool
}

func FromList[T any](list *listparam.List[T]) *ListIterSe[T] {
	return &ListIterSe[T]{
		root:     list,
		ele:      list.Front(),
		advanced: true,
	}
}

func (li *ListIterSe[T]) Next() (next T, ok bool) {
	if li.ele == nil {
		if li.root.Front() == nil {
			return
		}
		li.ele = li.root.Front()
	}

	if !li.advanced {
		nextEle := li.ele.Next()
		if nextEle == nil {
			return next, false
		} else {
			li.ele = nextEle
		}
	}

	ele, ok := li.ele.Get()
	nextEle := li.ele.Next()
	if nextEle == nil {
		li.advanced = false
	} else {
		li.ele = nextEle
		li.advanced = true
	}
	return ele, ok
}

func (li *ListIterSe[T]) ToIterator() Iterator[T] {
	return Iterator[T]{li}
}
