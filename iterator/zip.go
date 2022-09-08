package iterator

type TwoEleTuple[T any, U any] struct {
	Former T
	Latter U
}

type Zipper[T any, U any] struct {
	iterFormer SeIterator[T]
	iterLatter SeIterator[U]
}

func Zip[T any, U any](iterFormer SeIterator[T], iterLatter SeIterator[U]) Zipper[T, U] {
	return Zipper[T, U]{
		iterFormer: iterFormer,
		iterLatter: iterLatter,
	}
}

func (z Zipper[T, U]) Next() (next TwoEleTuple[T, U], ok bool) {
	nextFormer, okFormer := z.iterFormer.Next()
	nextLatter, okLatter := z.iterLatter.Next()
	if !(okFormer && okLatter) {
		return TwoEleTuple[T, U]{}, false
	}
	return TwoEleTuple[T, U]{Former: nextFormer, Latter: nextLatter}, true
}

func (z Zipper[T, U]) sizeHint() (former, latter int) {
	former, latter = -1, -1

	sizeFormer, okFormer := z.iterFormer.(SizeHinter)
	if okFormer {
		former = sizeFormer.SizeHint()
	}

	sizeLatter, okLatter := z.iterLatter.(SizeHinter)
	if okLatter {
		latter = sizeLatter.SizeHint()
	}

	return
}
func (z Zipper[T, U]) SizeHint() int {
	formerLen, latterLen := z.sizeHint()
	if formerLen < latterLen {
		return formerLen
	} else {
		return latterLen
	}
}

// Reverse implements Reverser.
// Reverse succeeds only when both are reversible and same size.
func (z Zipper[T, U]) Reverse() (rev SeIterator[TwoEleTuple[T, U]], ok bool) {
	formerLen, latterLen := z.sizeHint()
	if formerLen < 0 || latterLen < 0 || formerLen != latterLen {
		return nil, false
	}
	revFormer, okFormer := Reverse(z.iterFormer)
	revLatter, okLatter := Reverse(z.iterLatter)
	if !(okFormer && okLatter) {
		return nil, false
	}
	return Zipper[T, U]{
		iterFormer: revFormer,
		iterLatter: revLatter,
	}, true
}
