package iterator

type ChanIter[T any] struct {
	channel <-chan T
}

// FromChannel makes ChanIter associated with given channel.
// To signal end of iterator, close passed channel.
//
// *ChanIter[T] only implements SeIterator[T].
func FromChannel[T any](channel <-chan T) *ChanIter[T] {
	return &ChanIter[T]{
		channel: channel,
	}
}

// Next earns next element from this iterator.
// Next blocks until internal channel receives.
func (ci *ChanIter[T]) Next() (next T, ok bool) {
	next, ok = <-ci.channel
	return
}

func (ci *ChanIter[T]) ToIterator() Iterator[T] {
	return Iterator[T]{
		SeIterator: ci,
	}
}
