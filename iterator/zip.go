package iterator

type Zipper[T any] struct {
	isFormerExhausted bool
	isLatterExhausted bool
	former            DeIterator[T]
	latter            DeIterator[T]
}

func NewZipper[T any](former DeIterator[T], latter DeIterator[T]) *Zipper[T] {
	return &Zipper[T]{
		former: former,
		latter: latter,
	}
}

func (iter *Zipper[T]) Len() int {
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

func (z *Zipper[T]) next(
	nextFnFoermer, nextFnLatter nextFunc[T],
	isExhausted func() bool,
	setExhausted func(),
) (next T, ok bool) {
	var v T
	if !isExhausted() {
		v, ok = nextFnFoermer()
		if ok {
			return v, ok
		}
		setExhausted()
	}

	v, ok = nextFnLatter()
	if ok {
		return v, true
	}

	return
}
func (z *Zipper[T]) Next() (next T, ok bool) {
	return z.next(
		z.former.Next,
		z.latter.Next,
		func() bool { return z.isFormerExhausted },
		func() { z.isFormerExhausted = true },
	)
}
func (z *Zipper[T]) NextBack() (next T, ok bool) {
	return z.next(
		z.latter.NextBack,
		z.former.NextBack,
		func() bool { return z.isLatterExhausted },
		func() { z.isLatterExhausted = true },
	)
}
