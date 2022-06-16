package iterator_test

import (
	"testing"

	"github.com/ngicks/type-param-common/iterator"
	"github.com/ngicks/type-param-common/list"
)

func TestReverse(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	intLis := list.New[int]()
	for _, v := range expected {
		intLis.PushBack(v)
	}

	expectedRev := []int{5, 4, 3, 2, 1}
	iterSliceRev := &iterator.Reverser[int]{iterator.FromSlice[int](expected)}
	iterListRev := &iterator.Reverser[int]{iterator.FromList[int](intLis)}
	testIteratorBasic[int](t, iterSliceRev, expectedRev)
	testIteratorBasic[int](t, iterListRev, expectedRev)
}
