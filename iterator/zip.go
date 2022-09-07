package iterator

type TwoEleTuple[T any, U any] struct {
	Former T
	Latter U
}

type Zipper[T any, U any] struct {
	iterFormer SeIterator[T]
	iterLatter SeIterator[U]
}

func (z Zipper[T, U]) Next() (next TwoEleTuple[T, U], ok bool) {
	nextFormer, okFormer := z.iterFormer.Next()
	nextLatter, okLatter := z.iterLatter.Next()
	if !(okFormer && okLatter) {
		return TwoEleTuple[T, U]{}, false
	}
	return TwoEleTuple[T, U]{Former: nextFormer, Latter: nextLatter}, true
}

func Zip[T any, U any](iterFormer SeIterator[T], iterLatter SeIterator[U]) SeIterator[TwoEleTuple[T, U]] {
	return Zipper[T, U]{
		iterFormer: iterFormer,
		iterLatter: iterLatter,
	}
}
