package slice

// Doubly ended queue.
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

func (d *Deque[T]) Get(index uint) (v T, ok bool) {
	return Get(*d, index)
}

// Clone copies inner slice.
func (d *Deque[T]) Clone() Deque[T] {
	return Clone(*d)
}

func (d *Deque[T]) Insert(index uint, ele T) {
	*d = Insert(*d, index, ele)
}

func (d *Deque[T]) Append(elements ...T) {
	if elements == nil {
		return
	}
	*d = Append(*d, elements...)
}

func (d *Deque[T]) Prepend(elements ...T) {
	if elements == nil {
		return
	}
	*d = Prepend(*d, elements...)
}

func pushBack[T any](sl *[]T, v T) {
	*sl = append(*sl, v)
}

func popBack[T any](sl *[]T) (v T, popped bool) {
	var zero T

	if len(*sl) == 0 {
		return zero, false
	}

	v = (*sl)[len(*sl)-1]

	// avoiding memory leak.
	(*sl)[len(*sl)-1] = zero

	*sl = (*sl)[:len(*sl)-1]

	return v, true
}

func pushFront[T any](sl *[]T, v T) {
	*sl = append([]T{v}, *sl...)
}

func popFront[T any](sl *[]T) (v T, popped bool) {
	var zero T

	if len(*sl) == 0 {
		return zero, false
	}

	v = (*sl)[0]

	// avoiding memory leak
	(*sl)[0] = zero

	*sl = (*sl)[1:]

	return v, true
}
