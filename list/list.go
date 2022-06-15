package list

import "container/list"

type Element[T any] struct {
	inner *list.Element
}

func NewElement[T any]() Element[T] {
	return Element[T]{
		inner: new(list.Element),
	}
}

func (e Element[T]) Unwrap() *list.Element {
	return e.inner
}
func (e Element[T]) Get() (v T) {
	if e.inner.Value == nil {
		return
	} else {
		return e.inner.Value.(T)
	}
}
func (e Element[T]) Set(v T) {
	e.inner.Value = v
}
func (e Element[T]) Next() Element[T] {
	return Element[T]{
		inner: e.inner.Next(),
	}
}
func (e Element[T]) Prev() Element[T] {
	return Element[T]{
		inner: e.inner.Prev(),
	}
}

type List[T any] struct {
	inner *list.List
}

func New[T any]() List[T] {
	return List[T]{
		inner: list.New(),
	}
}
func (l List[T]) Back() Element[T] {
	return Element[T]{
		inner: l.inner.Back(),
	}
}
func (l List[T]) Front() Element[T] {
	return Element[T]{
		inner: l.inner.Front(),
	}
}
func (l List[T]) Init() List[T] {
	return List[T]{
		inner: l.inner.Init(),
	}
}
func (l List[T]) InsertAfter(v T, mark Element[T]) Element[T] {
	return Element[T]{
		inner: l.inner.InsertAfter(v, mark.inner),
	}
}
func (l List[T]) InsertBefore(v T, mark Element[T]) Element[T] {
	return Element[T]{
		inner: l.inner.InsertBefore(v, mark.inner),
	}
}
func (l List[T]) Len() int {
	return l.inner.Len()
}
func (l List[T]) MoveAfter(e, mark Element[T]) {
	l.inner.MoveAfter(e.inner, mark.inner)
}
func (l List[T]) MoveBefore(e, mark Element[T]) {
	l.inner.MoveBefore(e.inner, mark.inner)
}
func (l List[T]) MoveToBack(e Element[T]) {
	l.inner.MoveToBack(e.inner)
}
func (l List[T]) MoveToFront(e Element[T]) {
	l.inner.MoveToFront(e.inner)
}
func (l List[T]) PushBack(v T) Element[T] {
	return Element[T]{
		inner: l.inner.PushBack(v),
	}
}
func (l List[T]) PushBackList(other List[T]) {
	l.inner.PushBackList(other.inner)
}
func (l List[T]) PushFront(v T) Element[T] {
	return Element[T]{
		inner: l.inner.PushFront(v),
	}
}
func (l List[T]) PushFrontList(other List[T]) {
	l.inner.PushFrontList(other.inner)
}
func (l List[T]) Remove(e Element[T]) (v T) {
	if ele := l.inner.Remove(e.inner); ele != nil {
		return ele.(T)
	} else {
		return
	}
}
