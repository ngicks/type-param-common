// package listparam wraps doubly-linked list
// implemented in go programming langauge standard library,
// to accept type parameter.
package listparam

import "container/list"

type Element[T any] struct {
	inner  *list.Element
	entMap entMap[T]
}

func (e *Element[T]) ensureValid() {
	if e.entMap == nil {
		panic("not initialized: zero *Element[T] is invalid. e.entMap is nil")
	}
	if e.inner == nil {
		panic("not initialized: zero *Element[T] is invalid. e.inner is nil")
	}
}

// Unwrap returns internal *`container/list`.Element.
// Setting non-T value may cause runtime panic in succeeding Get call.
func (e *Element[T]) Unwrap() *list.Element {
	e.ensureValid()
	return e.inner
}

// Get returns internal Value. If internal Value is non-nil and then returns value and true.
// Otherwise returns zero of T and false.
func (e *Element[T]) Get() (v T, ok bool) {
	e.ensureValid()

	if e.inner.Value == nil {
		return
	} else {
		return e.inner.Value.(T), true
	}
}

// Set is equivalent to `element.Value = v`
func (e *Element[T]) Set(v T) {
	e.ensureValid()
	e.inner.Value = v
}

func (e *Element[T]) Next() *Element[T] {
	e.ensureValid()
	return e.entMap.ensureGet(e.inner.Next())
}
func (e *Element[T]) Prev() *Element[T] {
	e.ensureValid()
	return e.entMap.ensureGet(e.inner.Prev())
}

type List[T any] struct {
	inner  *list.List
	entMap entMap[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		inner:  list.New(),
		entMap: make(entMap[T]),
	}
}

func (l *List[T]) lazyInit() {
	if l.inner == nil {
		l.inner = list.New()
	}
	if l.entMap == nil {
		l.entMap = make(entMap[T])
	}
}

func (l *List[T]) Back() *Element[T] {
	l.lazyInit()
	return l.entMap.ensureGet(l.inner.Back())
}
func (l *List[T]) Front() *Element[T] {
	l.lazyInit()
	return l.entMap.ensureGet(l.inner.Front())
}
func (l *List[T]) Init() *List[T] {
	l.lazyInit()
	l.inner.Init()
	l.entMap = make(entMap[T])
	return l
}
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	l.lazyInit()
	return l.entMap.createInsertingElement(l.inner.InsertAfter(v, mark.inner))
}
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	l.lazyInit()
	return l.entMap.createInsertingElement(l.inner.InsertBefore(v, mark.inner))
}
func (l *List[T]) Len() int {
	l.lazyInit()
	return l.inner.Len()
}
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.lazyInit()
	l.inner.MoveAfter(e.inner, mark.inner)
}
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.lazyInit()
	l.inner.MoveBefore(e.inner, mark.inner)
}
func (l *List[T]) MoveToBack(e *Element[T]) {
	l.lazyInit()
	l.inner.MoveToBack(e.inner)
}
func (l *List[T]) MoveToFront(e *Element[T]) {
	l.lazyInit()
	l.inner.MoveToFront(e.inner)
}
func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()
	return l.entMap.createInsertingElement(l.inner.PushBack(v))
}
func (l *List[T]) PushBackList(other *List[T]) {
	l.lazyInit()
	other.lazyInit()
	l.inner.PushBackList(other.inner)
}
func (l *List[T]) PushFront(v T) *Element[T] {
	l.lazyInit()
	return l.entMap.createInsertingElement(l.inner.PushFront(v))
}
func (l *List[T]) PushFrontList(other *List[T]) {
	l.lazyInit()
	other.lazyInit()
	l.inner.PushFrontList(other.inner)
}

// Remove calls Remove method of internal `container/list`.List.
// If Remove returns non-nil value then removed is true, false otherwize.
// When removed is false returned v is zero of T.
func (l *List[T]) Remove(e *Element[T]) (v T, removed bool) {
	l.lazyInit()
	delete(l.entMap, e.inner)
	if ele := l.inner.Remove(e.inner); ele != nil {
		return ele.(T), true
	} else {
		return
	}
}
