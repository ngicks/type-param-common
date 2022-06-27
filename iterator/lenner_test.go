package iterator_test

import (
	"testing"

	"github.com/ngicks/type-param-common/iterator"
)

func TestLenner(t *testing.T) {
	const leng = 5
	innerSl := []int{1, 2, 3, 4, 5}
	sliceIter := iterator.FromSlice(innerSl)
	iter := iterator.Iterator[int]{DeIterator: sliceIter}

	assertEq := func() {
		if iter.Len() != leng {
			t.Fatalf("Len incorrect: expected = %d, actual  %d", leng, iter.Len())
		}
	}

	assertEq()

	iter = iterator.Iterator[int]{DeIterator: iterator.NewExcluder[int](sliceIter, func(i int) bool { return false })}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.NewSelector[int](sliceIter, func(i int) bool { return false })}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.Map[int](sliceIter, func(i int) int { return i })}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.NewReverser[int](sliceIter)}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.NewNSkipper[int](sliceIter, 1)}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.NewWhileSkipper[int](sliceIter, func(i int) bool { return false })}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.NewNTaker[int](sliceIter, 1)}
	assertEq()
	iter = iterator.Iterator[int]{DeIterator: iterator.NewWhileTaker[int](sliceIter, func(i int) bool { return false })}
	assertEq()
}
