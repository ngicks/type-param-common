package iterator

func Fold[T, U any](iter SeIterator[T], reducer func(accumulator U, next T) U, inital U) U {
	var accum U = inital
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		accum = reducer(accum, next)
	}
	return accum
}
