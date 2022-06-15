package heap

import (
	"container/heap"
	"sort"
)

type Interface[T any] interface {
	sort.Interface
	Push(x T)
	Pop() T
}

type interfaceReverter[T any] struct {
	Interface[T]
}

func (ir interfaceReverter[T]) Push(x any) {
	ir.Interface.Push(x.(T))
}

func (ir interfaceReverter[T]) Pop() any {
	return ir.Interface.Pop()
}

func Init[T any](h Interface[T]) {
	heap.Init(interfaceReverter[T]{h})
}

func Push[T any](h Interface[T], x T) {
	heap.Push(interfaceReverter[T]{h}, x)
}

func Pop[T any](h Interface[T]) (v T) {
	if popped := heap.Pop(interfaceReverter[T]{h}); popped != nil {
		// Pop'ping empty Interface may or may not panic.
		return popped.(T)
	} else {
		return
	}
}

func Remove[T any](h Interface[T], i int) (v T) {
	if removed := heap.Remove(interfaceReverter[T]{h}, i); removed != nil {
		// Remove'ing empty Interface may or may not panic.
		return removed.(T)
	} else {
		return
	}
}

func Fix[T any](h Interface[T], i int) {
	heap.Fix(interfaceReverter[T]{h}, i)
}
