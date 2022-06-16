package iterator

import (
	"github.com/ngicks/type-param-common/list"
)

type Nexter[T any] interface {
	Next() (next T, ok bool)
}
type NextBacker[T any] interface {
	NextBack() (next T, ok bool)
}

type Lenner interface {
	Len() int
}

type IterReverser[T any] interface {
	Reverse() Reverser[T]
}
type IterSelector[T any] interface {
	Select(selector func(T) bool) Selector[T]
}
type IterExcluder[T any] interface {
	Exclude(excluder func(T) bool) Excluder[T]
}
type IterMapper[T, U any] interface {
	Map(mapper func(T) U) Mapper[T, U]
}

// Reverse() Reverser[T]
// Select(selector func(T) bool) Selector[T]
// Exclude(excluder func(T) bool) Excluder[T]
// Map(mapper func(T) U) Mapper[T, U]

type Iterator[T any] interface {
	Nexter[T]
	NextBacker[T]
	IterReverser[T]
	IterSelector[T]
	IterExcluder[T]
}

// Doubly ended iterator.
type DeIterator[T any] interface {
	Nexter[T]
	NextBacker[T]
}

var _ Iterator[any] = &SliceIter[any]{}

type SliceIter[T any] struct {
	innerSlice []T
	idxFront   int
	idxBack    int
}

func FromSlice[T any](sl []T) *SliceIter[T] {
	return &SliceIter[T]{
		innerSlice: sl,
		idxFront:   0,
		idxBack:    len(sl) - 1,
	}
}

func (si *SliceIter[T]) Next() (next T, ok bool) {
	if si.idxFront > si.idxBack {
		return
	}
	next, ok = si.innerSlice[si.idxFront], true
	si.idxFront++
	return
}
func (si *SliceIter[T]) NextBack() (next T, ok bool) {
	if si.idxFront > si.idxBack {
		return
	}
	next, ok = si.innerSlice[si.idxBack], true
	si.idxBack--
	return
}
func (si *SliceIter[T]) Len() int {
	return len(si.innerSlice)
}
func (si *SliceIter[T]) Reverse() Reverser[T] {
	return Reverser[T]{si}
}
func (si *SliceIter[T]) Select(selector func(T) bool) Selector[T] {
	return Selector[T]{si, selector}
}
func (si *SliceIter[T]) Exclude(excluder func(T) bool) Excluder[T] {
	return Excluder[T]{si, excluder}
}

var _ Iterator[any] = &ListIter[any]{}

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
func (li *ListIter[T]) Reverse() Reverser[T] {
	return Reverser[T]{li}
}
func (li *ListIter[T]) Select(selector func(T) bool) Selector[T] {
	return Selector[T]{li, selector}
}
func (li *ListIter[T]) Exclude(excluder func(T) bool) Excluder[T] {
	return Excluder[T]{li, excluder}
}
