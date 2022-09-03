package iterator

//go:generate go run ../cmd/lenner/lenner.go -i . -ignore "lenner.go,chain.go"  -o lenner.go

// nextFunc should be type alias but Go does not provide aliasing for type-param func at the time.
type nextFunc[T any] func() (next T, ok bool)

type Nexter[T any] interface {
	Next() (next T, ok bool)
}
type NextBacker[T any] interface {
	NextBack() (next T, ok bool)
}

type Lenner interface {
	Len() int
}

type IterReverser[T any] interface {
	Reverse() Reverser[T]
}
type IterSelector[T any] interface {
	Select(selector func(T) bool) Selector[T]
}
type IterExcluder[T any] interface {
	Exclude(excluder func(T) bool) Excluder[T]
}
type IterMapper[T, U any] interface {
	Map(mapper func(T) U) Mapper[T, U]
}

type IterSkipper[T any] interface {
	SkipN() NSkipper[T]
	SkipWhile(skipIf func(T) bool) WhileSkipper[T]
}

type IterTaker[T any] interface {
	TakeN() NTaker[T]
	TakeWhile(skipIf func(T) bool) WhileTaker[T]
}

// Doubly ended iterator.
type DeIterator[T any] interface {
	Nexter[T]
	NextBacker[T]
}

type Iterator[T any] struct {
	DeIterator[T]
}

func (iter Iterator[T]) Reverse() Iterator[T] {
	return Iterator[T]{NewReverser[T](iter)}
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
func (iter Iterator[T]) ForEach(each func(T)) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		each(next)
	}
}
func (iter Iterator[T]) Collect() []T {
	collected := make([]T, 0)
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		collected = append(collected, next)
	}
	return collected
}
func (iter Iterator[T]) Chain(z Iterator[T]) Iterator[T] {
	return Iterator[T]{
		DeIterator: NewChainer[T](iter.DeIterator, z),
	}
}
