package iterator

type Chainer[T any] struct {
	isFormerExhausted bool
	isLatterExhausted bool
	former            SeIterator[T]
	latter            SeIterator[T]
}

func NewChainer[T any](former SeIterator[T], latter SeIterator[T]) *Chainer[T] {
	return &Chainer[T]{
		former: former,
		latter: latter,
	}
}

// SizeHint implements SizeHinter.
func (c *Chainer[T]) SizeHint() int {
	lennerFormer, ok := c.former.(SizeHinter)
	if !ok {
		return -1
	}
	lennerLatter, ok := c.latter.(SizeHinter)
	if !ok {
		return -1
	}
	formerLen := lennerFormer.SizeHint()
	latterLen := lennerLatter.SizeHint()
	if formerLen < 0 || latterLen < 0 {
		return -1
	}
	return formerLen + latterLen
}

func (c *Chainer[T]) Next() (next T, ok bool) {
	if !c.isFormerExhausted {
		v, ok := c.former.Next()
		if ok {
			return v, ok
		}
		// former will not be Next-ed once it returns not-ok.
		c.isFormerExhausted = true
	}

	if !c.isLatterExhausted {
		v, ok := c.latter.Next()
		if ok {
			return v, ok
		}
		// latter is treated similary.
		c.isLatterExhausted = true
	}
	return
}

// Reverse implements Reverser.
func (c *Chainer[T]) Reverse() (rev SeIterator[T], ok bool) {
	revFormer, okFormer := Reverse(c.former)
	revLatter, okLatter := Reverse(c.latter)
	if !(okFormer && okLatter) {
		return nil, false
	}
	return &Chainer[T]{
		former:            revLatter,
		latter:            revFormer,
		isFormerExhausted: c.isLatterExhausted,
		isLatterExhausted: c.isFormerExhausted,
	}, true
}
