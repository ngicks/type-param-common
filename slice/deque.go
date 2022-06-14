package slice

// Doubly ended queueu.
type Deque[T any] []T

// Len returns length of d.
func (d *Deque[T]) Len() int {
	return len(*d)
}

func (d *Deque[T]) Push(v T) {
	d.PushBack(v)
}

func (d *Deque[T]) Pop() (v T, popped bool) {
	return d.PopBack()
}

func (d *Deque[T]) PushBack(v T) {
	pushBack((*[]T)(d), v)
}

func (d *Deque[T]) PopBack() (v T, popped bool) {
	return popBack((*[]T)(d))
}

func (d *Deque[T]) PushFront(v T) {
	pushFront((*[]T)(d), v)
}

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
