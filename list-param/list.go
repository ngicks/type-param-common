package listparam

import "container/list"

type Element[T any] struct {
	inner *list.Element
}

func NewElement[T any]() Element[T] {
	return Element[T]{
		inner: new(list.Element),
	}
}

// Unwrap returns internal *`container/list`.Element.
// Setting non-T value may cause runtime panic in succeeding Get call.
func (e Element[T]) Unwrap() *list.Element {
	return e.inner
}

// Get returns internal Value. If internal Value is non-nil and then returns value and true.
// Otherwise returns zero of T and false.
func (e Element[T]) Get() (v T, ok bool) {
	if e.inner == nil || e.inner.Value == nil {
		return
	} else {
		return e.inner.Value.(T), true
	}
}

// Set is equivalent to `element.Value = v`
func (e Element[T]) Set(v T) {
	e.inner.Value = v
}

// Clear is equivalent to `element.Value = nil`
func (e Element[T]) Clear() {
	e.inner.Value = nil
}
func (e Element[T]) Next() Element[T] {
	if e.inner == nil {
		return Element[T]{
			inner: nil,
		}
	}
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

// Remove calls Remove method of internal `container/list`.List.
// If Remove returns non-nil value then removed is true, false otherwize.
// When removed is false returned v is zero of T.
func (l List[T]) Remove(e Element[T]) (v T, removed bool) {
	if ele := l.inner.Remove(e.inner); ele != nil {
		return ele.(T), true
	} else {
		return
	}
}
