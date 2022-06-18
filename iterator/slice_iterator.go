package iterator

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
