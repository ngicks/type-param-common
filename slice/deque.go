package slice

// Doubly ended queueu.
type Deque[T any] []T

// Len returns length of underlying slice.
func (d *Deque[T]) Len() int {
	return len(*d)
}

// Push is an alias for PushBack.
func (d *Deque[T]) Push(v T) {
	d.PushBack(v)
}

// Pop is an alias for PopBack.
func (d *Deque[T]) Pop() (v T, popped bool) {
	return d.PopBack()
}

// PushBack adds an element to tail of underlying slice.
func (d *Deque[T]) PushBack(v T) {
	pushBack((*[]T)(d), v)
}

// PopBack removes an element from tail of underlying slice, and then returns removed value.
// If slice is empty, returns zero of T and false.
func (d *Deque[T]) PopBack() (v T, popped bool) {
	return popBack((*[]T)(d))
}

// PushFront adds an element to head of underlying slice.
func (d *Deque[T]) PushFront(v T) {
	pushFront((*[]T)(d), v)
}

// PopFront removes an element from head of underlying slice, and then returns removed value.
// If slice is empty, returns zero of T and false.
func (d *Deque[T]) PopFront() (v T, popped bool) {
	return popFront((*[]T)(d))
}

func pushBack[T any](sl *[]T, v T) {
	*sl = append(*sl, v)
}

func popBack[T any](sl *[]T) (v T, popped bool) {
	if len(*sl) == 0 {
		return
	}
	popped = true
	*sl, v = (*sl)[:len(*sl)-1], (*sl)[len(*sl)-1]
	return
}

func pushFront[T any](sl *[]T, v T) {
	*sl = append([]T{v}, *sl...)
}

func popFront[T any](sl *[]T) (v T, popped bool) {
	if len(*sl) == 0 {
		return
	}
	popped = true
	*sl, v = (*sl)[1:], (*sl)[0]
	return
}
