package set

import "github.com/ngicks/type-param-common/iterator"

type SetLike[T comparable] interface {
	Len() int
	Add(v T)
	Clear()
	Delete(v T) (deleted bool)
	ForEach(f func(v T, idx int))
	Has(v T) (has bool)
	Values() iterator.Iterator[T]
}
