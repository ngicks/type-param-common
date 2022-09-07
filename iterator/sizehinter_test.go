package iterator

import (
	"testing"
)

func TestSizeHinter(t *testing.T) {
	innerSl := []int{1, 2, 3, 4, 5}
	sliceIter := FromSlice(innerSl)
	iter := Iterator[int]{SeIterator: sliceIter}

	assertEq := func(leng int) {
		if iter.SizeHint() != leng {
			t.Fatalf("Len incorrect: expected = %d, actual  %d", leng, iter.SizeHint())
		}
	}

	assertEq(5)

	iter = Iterator[int]{
		SeIterator: Excluder[int]{
			inner:    sliceIter,
			excluder: func(i int) bool { return i%3 == 0 },
		},
	}
	assertEq(5)
	iter = Iterator[int]{
		SeIterator: Selector[int]{
			inner:    sliceIter,
			selector: func(i int) bool { return i%3 == 0 },
		},
	}
	assertEq(5)
	iter = Iterator[int]{
		SeIterator: Iterator[int]{SeIterator: sliceIter},
	}
	assertEq(5)
	iter = Iterator[int]{
		SeIterator: Mapper[int, int]{
			inner:  sliceIter,
			mapper: func(i int) int { return i },
		},
	}
	assertEq(5)
	iter = Iterator[int]{SeIterator: ReversedDeIter[int]{DeIterator: sliceIter}}
	assertEq(5)
}
