package typeparamcommon

type Lessable[T any] interface {
	Inner() T
	// LessThan determines self is less than input Lessable[T].
	LessThan(Lessable[T]) bool
}

type FilterableHeap[U any, T Lessable[U]] struct {
	*HeapWrapper[T]
	internal *SliceInterface[T]
}

func genericLess[U any, T Lessable[U]](i T, j T) bool {
	return i.LessThan(j)
}

func NewFilterableHeap[U any, T Lessable[U]]() *FilterableHeap[U, T] {
	heapInternal, interfaceInternal := MakeHeap(genericLess[U, T])
	return &FilterableHeap[U, T]{
		HeapWrapper: heapInternal,
		internal:    interfaceInternal,
	}
}

// Peek returns most prioritized value in heap without removing it.
// If this heap contains 0 element, returned p is zero value for type T.
//
// The complexity is O(1).
func (h *FilterableHeap[U, T]) Peek() (p T) {
	if len(h.internal.Inner) == 0 {
		return
	}
	return h.internal.Inner[0]
}

func (h *FilterableHeap[U, T]) Len() int {
	return h.internal.Len()
}

// Clone clones internal slice and creates new FilterableHeap with it.
// This is done by simple slice copy, without succeeding Init call,
// which means it also clones broken invariants if any.
//
// If type T or one of its internal value is pointer type,
// mutation of T porpagates cloned to original, and vice versa.
func (h *FilterableHeap[U, T]) Clone() *FilterableHeap[U, T] {
	cloned := make([]T, len(h.internal.Inner))
	copy(cloned, h.internal.Inner)

	n := NewFilterableHeap[U, T]()
	n.internal.Inner = cloned
	return n
}

// Filter calls filterFuncs one by one, in given order, with following heap.Init().
// Each filterFunc receives innerSlice so that it can be mutated in that func.
//
// Filter calls heap.Init() at the end of the method.
// So the complexity is at least O(n) where n is h.Len().
func (h *FilterableHeap[U, T]) Filter(filterFuncs ...func(innerSlice *[]T)) {
	for _, v := range filterFuncs {
		if v != nil {
			v(&h.internal.Inner)
		}
	}

	h.Init()
}

func BuildExcludeFilter[T any](start, end int, filterPredicate func(T) bool) func(sl *[]T) {
	return func(sl *[]T) {
		if start < 0 {
			start = 0
		} else if start >= len(*sl) {
			return
		}
		if end > len(*sl) {
			end = len(*sl)
		}

		if start > end {
			return
		}

		for i := start; i < end; i++ {
			if filterPredicate((*sl)[i]) {
				*sl = append((*sl)[:i], (*sl)[i+1:]...)
				end--
				i--
			}
		}
	}
}
