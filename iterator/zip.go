package iterator

type TwoEleTuple[T any, U any] struct {
	Former T
	Latter U
}

type Zipper[T any, U any] struct {
	iterFormer SeIterator[T]
	iterLatter SeIterator[U]
}

func Zip[T any, U any](iterFormer SeIterator[T], iterLatter SeIterator[U]) SeIterator[TwoEleTuple[T, U]] {
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

func (z Zipper[T, U]) SizeHint() int {
	sizeFormer, okFormer := z.iterFormer.(SizeHinter)
	sizeLatter, okLatter := z.iterLatter.(SizeHinter)
	if !(okFormer && okLatter) {
		return -1
	}
	formerLen := sizeFormer.SizeHint()
	latterLen := sizeLatter.SizeHint()
	if formerLen < latterLen {
		return formerLen
	} else {
		return latterLen
	}
}

func (z Zipper[T, U]) Reverse() (rev SeIterator[TwoEleTuple[T, U]], ok bool) {
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
