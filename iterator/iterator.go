package iterator

import (
	"fmt"

	listparam "github.com/ngicks/type-param-common/list-param"
)

//go:generate go run ../cmd/sizehinter/sizehinter.go -i . -ignore "sizehinter.go,chain.go,reverser.go,skip.go,take.go,zip.go"  -o sizehinter.go
//go:generate go run ../cmd/reverser/reverser.go -i . -ignore "reverse.go,iterator.go,sizehinter.go,chain.go,reverser.go,enumerate.go,zip.go"  -o reverser.go

// nextFunc should be type alias but Go does not provide aliasing for type-param func at the time.
type nextFunc[T any] func() (next T, ok bool)

type Nexter[T any] interface {
	Next() (next T, ok bool)
}
type NextBacker[T any] interface {
	NextBack() (next T, ok bool)
}

type SizeHinter interface {
	SizeHint() int
}

type Reverser[T any] interface {
	Reverse() (rev SeIterator[T], ok bool)
}

type Unwrapper[T any] interface {
	Unwrap() SeIterator[T]
}

// Singly ended iterator.
type SeIterator[T any] interface {
	Nexter[T]
}

// Doubly ended iterator.
type DeIterator[T any] interface {
	Nexter[T]
	NextBacker[T]
}

type Iterator[T any] struct {
	SeIterator[T]
}

func FromSlice[T any](sl []T) Iterator[T] {
	return Iterator[T]{
		SeIterator: NewSliceIterDe(sl),
	}
}

func FromFixedList[T any](list *listparam.List[T]) Iterator[T] {
	return Iterator[T]{
		SeIterator: NewListIterDe(list),
	}
}

func FromList[T any](list *listparam.List[T]) Iterator[T] {
	return Iterator[T]{
		SeIterator: NewListIterSe(list),
	}
}

func FromChannel[T any](channel <-chan T) Iterator[T] {
	return Iterator[T]{
		SeIterator: NewChanIter(channel),
	}
}

func FromRange(start, end int) Iterator[int] {
	return Iterator[int]{
		SeIterator: NewRange(start, end),
	}
}

func (iter Iterator[T]) Select(selector func(T) bool) Iterator[T] {
	return Iterator[T]{NewSelector[T](iter, selector)}
}
func (iter Iterator[T]) Exclude(excluder func(T) bool) Iterator[T] {
	return Iterator[T]{NewExcluder[T](iter, excluder)}
}
func (iter Iterator[T]) SkipN(n int) Iterator[T] {
	return Iterator[T]{NewNSkipper[T](iter, n)}
}
func (iter Iterator[T]) SkipWhile(skipIf func(T) bool) Iterator[T] {
	return Iterator[T]{NewWhileSkipper[T](iter, skipIf)}
}
func (iter Iterator[T]) TakeN(n int) Iterator[T] {
	return Iterator[T]{NewNTaker[T](iter, n)}
}
func (iter Iterator[T]) TakeWhile(takeIf func(T) bool) Iterator[T] {
	return Iterator[T]{NewWhileTaker[T](iter, takeIf)}
}
func (iter Iterator[T]) Chain(z SeIterator[T]) Iterator[T] {
	return Iterator[T]{NewChainer(iter.SeIterator, z)}
}
func (iter Iterator[T]) Map(mapper func(T) T) Iterator[T] {
	return Iterator[T]{NewSameTyMapper(iter.SeIterator, mapper)}
}

func (iter Iterator[T]) Unwrap() SeIterator[T] {
	return iter.SeIterator
}
func (iter Iterator[T]) MustNext() T {
	v, ok := iter.SeIterator.Next()
	if !ok {
		panic("NextMust: failed")
	}
	return v
}
func (iter Iterator[T]) Reverse() (rev Iterator[T], ok bool) {
	reversed, ok := Reverse(iter.SeIterator)
	if !ok {
		return
	}
	return Iterator[T]{reversed}, true
}
func (iter Iterator[T]) MustReverse() (rev Iterator[T]) {
	rev, ok := iter.Reverse()
	if !ok {
		panic(fmt.Sprintf("MustReverse: failed: %+v", iter))
	}
	return
}
func (iter Iterator[T]) iterateUntil(predicate func(T) (continueIteration bool)) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !predicate(next) {
			break
		}
	}
}
func (iter Iterator[T]) ForEach(each func(T)) {
	iter.iterateUntil(func(t T) bool {
		each(t)
		return true
	})
}
func (iter Iterator[T]) Collect() []T {
	collected := make([]T, 0)
	iter.iterateUntil(func(t T) bool {
		collected = append(collected, t)
		return true
	})
	return collected
}
func (iter Iterator[T]) Find(predicate func(T) bool) (v T, found bool) {
	var lastElement T
	iter.iterateUntil(func(t T) bool {
		b := predicate(t)
		if b {
			lastElement = t
			found = true
		}
		return !b
	})
	return lastElement, found
}

func (iter Iterator[T]) Reduce(reducer func(accumulator T, next T) T) T {
	var accum T
	iter.iterateUntil(func(t T) bool {
		accum = reducer(accum, t)
		return true
	})
	return accum
}
