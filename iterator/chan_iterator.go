package iterator

type ChanIter[T any] struct {
	channel <-chan T
}

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

// NextBack is alias of Next. Reversing ChanIter is no-op.
func (ci *ChanIter[T]) NextBack() (next T, ok bool) {
	return ci.Next()
}
