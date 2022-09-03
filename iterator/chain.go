package iterator

type Chainer[T any] struct {
	isFormerExhausted bool
	isLatterExhausted bool
	former            DeIterator[T]
	latter            DeIterator[T]
}

func NewChainer[T any](former DeIterator[T], latter DeIterator[T]) *Chainer[T] {
	return &Chainer[T]{
		former: former,
		latter: latter,
	}
}

func (iter *Chainer[T]) Len() int {
	lennerFormer, ok := iter.former.(Lenner)
	if !ok {
		return -1
	}
	lennerLatter, ok := iter.latter.(Lenner)
	if !ok {
		return -1
	}
	formerLen := lennerFormer.Len()
	latterLen := lennerLatter.Len()
	if formerLen < 0 || latterLen < 0 {
		return -1
	}
	return formerLen + latterLen
}

func (z *Chainer[T]) next(
	nextFnFormer, nextFnLatter nextFunc[T],
	isFormerExhausted func() bool,
	setFormerExhausted func(),
) (next T, ok bool) {
	var v T
	if !isFormerExhausted() {
		v, ok = nextFnFormer()
		if ok {
			return v, ok
		}
		setFormerExhausted()
	}

	v, ok = nextFnLatter()
	if ok {
		return v, true
	}

	return
}
func (z *Chainer[T]) Next() (next T, ok bool) {
	return z.next(
		z.former.Next,
		z.latter.Next,
		func() bool { return z.isFormerExhausted },
		func() { z.isFormerExhausted = true },
	)
}
func (z *Chainer[T]) NextBack() (next T, ok bool) {
	return z.next(
		z.latter.NextBack,
		z.former.NextBack,
		func() bool { return z.isLatterExhausted },
		func() { z.isLatterExhausted = true },
	)
}
