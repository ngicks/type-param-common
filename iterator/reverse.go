package iterator

import "fmt"

func Reverse[T any](iter SeIterator[T]) (rev SeIterator[T], ok bool) {
	switch x := iter.(type) {
	case Reverser[T]:
		return x.Reverse()
	case DeIterator[T]:
		return ReversedDeIter[T]{DeIterator: x}, true
	case Unwrapper[T]:
		return Reverse(x.Unwrap())
	}
	return nil, false
}

func MustReverse[T any](iter SeIterator[T]) (rev SeIterator[T]) {
	rev, ok := Reverse(iter)
	if !ok {
		panic(fmt.Sprintf("MustReverse: failed: %+v", iter))
	}
	return
}

type ReversedDeIter[T any] struct {
	DeIterator[T]
}

func (rev ReversedDeIter[T]) Next() (next T, ok bool) {
	return rev.DeIterator.NextBack()
}
func (rev ReversedDeIter[T]) NextBack() (next T, ok bool) {
	return rev.DeIterator.Next()
}

// Reverse implements Reverser[T].
// This simply unwrap iterator.
func (iter ReversedDeIter[T]) Reverse() (rev SeIterator[T], ok bool) {
	return iter.DeIterator, true
}
