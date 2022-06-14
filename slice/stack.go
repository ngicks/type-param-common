package slice

// LIFO stack
type Stack[T any] []T

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(v T) {
	pushBack((*[]T)(s), v)
}

func (s *Stack[T]) Pop() (v T, popped bool) {
	return popBack((*[]T)(s))
}
