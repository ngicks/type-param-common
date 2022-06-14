package slice

// FIFO queue.
type Queue[T any] []T

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Push(v T) {
	pushBack((*[]T)(q), v)
}

func (q *Queue[T]) Pop() (v T, popped bool) {
	return popFront((*[]T)(q))
}
