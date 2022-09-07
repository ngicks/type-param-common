package iterator

// Code generated by github.com/ngicks/type-param-common/cmd/lenner. DO NOT EDIT.

func (iter Excluder[T]) SizeHint() int {
	if sizehinter, ok := iter.inner.(SizeHinter); ok {
		return sizehinter.SizeHint()
	}
	return -1
}

func (iter Selector[T]) SizeHint() int {
	if sizehinter, ok := iter.inner.(SizeHinter); ok {
		return sizehinter.SizeHint()
	}
	return -1
}

func (iter Iterator[T]) SizeHint() int {
	if sizehinter, ok := iter.SeIterator.(SizeHinter); ok {
		return sizehinter.SizeHint()
	}
	return -1
}

func (iter Mapper[T,U]) SizeHint() int {
	if sizehinter, ok := iter.inner.(SizeHinter); ok {
		return sizehinter.SizeHint()
	}
	return -1
}

func (iter ReversedDeIter[T]) SizeHint() int {
	if sizehinter, ok := iter.DeIterator.(SizeHinter); ok {
		return sizehinter.SizeHint()
	}
	return -1
}
