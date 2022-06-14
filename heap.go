package typeparamcommon

import "github.com/ngicks/type-param-common/heap"

// MakeHeap makes a heap for the type T using a less[T] function.
//
// 1st returned value is struct with basic set of heap methods.
// 2nd is a point to struct that implements heap.Interface[T] which is used in *HeapWrapper[T].
// To add your own heap methods, embed *HeapWrapper[T] to your own struct type
// and manipulate SliceInterface[T].Inner slice in that struct methods with succeeding *HeapWrapper.Init call.
func MakeHeap[T any](less func(i, j T) bool) (*HeapWrapper[T], *SliceInterface[T]) {
	internal := NewSliceInterface(nil, less)
	return NewHeapWrapper[T](internal), internal
}

type Lessable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 | ~string | ~uint |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func less[T Lessable](i, j T) bool {
	return i < j
}

func more[T Lessable](i, j T) bool {
	return i > j
}

// MakeMinHeap makes a minheap for the type T.
//
// MakeMinHeap does what MakeHeap does but with predeclared less function.
// T is constrained to predeclared primitive types which are compatible with less and greater comparison operations.
func MakeMinHeap[T Lessable]() (*HeapWrapper[T], *SliceInterface[T]) {
	internal := NewSliceInterface(nil, less[T])
	return NewHeapWrapper[T](internal), internal
}

// MakeMaxHeap makes a maxheap for the type T.
// This is same as MakeMinHeap but uses more[T] instead.
func MakeMaxHeap[T Lessable]() (*HeapWrapper[T], *SliceInterface[T]) {
	internal := NewSliceInterface(nil, more[T])
	return NewHeapWrapper[T](internal), internal
}

type HeapWrapper[T any] struct {
	inter heap.Interface[T]
}

func NewHeapWrapper[T any](inter heap.Interface[T]) *HeapWrapper[T] {
	return &HeapWrapper[T]{
		inter: inter,
	}
}

func (hw *HeapWrapper[T]) Fix(i int) {
	heap.Fix(hw.inter, i)
}
func (hw *HeapWrapper[T]) Init() {
	heap.Init(hw.inter)
}
func (hw *HeapWrapper[T]) Pop() T {
	return heap.Pop(hw.inter)
}
func (hw *HeapWrapper[T]) Push(x T) {
	heap.Push(hw.inter, x)
}
func (hw *HeapWrapper[T]) Remove(i int) T {
	return heap.Remove(hw.inter, i)
}

var _ heap.Interface[int] = NewSliceInterface[int](nil, nil)

type SliceInterface[T any] struct {
	Inner []T
	less  func(i, j T) bool
}

func NewSliceInterface[T any](init []T, less func(i, j T) bool) *SliceInterface[T] {
	if init == nil {
		init = make([]T, 0)
	}
	return &SliceInterface[T]{
		Inner: init,
		less:  less,
	}
}

func (sl *SliceInterface[T]) Len() int {
	return len(sl.Inner)
}
func (sl *SliceInterface[T]) Less(i, j int) bool {
	return sl.less(sl.Inner[i], sl.Inner[j])
}
func (sl *SliceInterface[T]) Swap(i, j int) {
	sl.Inner[i], sl.Inner[j] = sl.Inner[j], sl.Inner[i]
}
func (sl *SliceInterface[T]) Push(x T) {
	sl.Inner = append(sl.Inner, x)
}
func (sl *SliceInterface[T]) Pop() (p T) {
	p, sl.Inner = sl.Inner[len(sl.Inner)-1], sl.Inner[:len(sl.Inner)-1]
	return p
}
