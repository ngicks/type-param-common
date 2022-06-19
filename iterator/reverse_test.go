package iterator_test

import (
	"testing"

	"github.com/ngicks/type-param-common/iterator"
	list "github.com/ngicks/type-param-common/list-param"
)

func TestReverse(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	intLis := list.New[int]()
	for _, v := range expected {
		intLis.PushBack(v)
	}

	expectedRev := []int{5, 4, 3, 2, 1}
	iterSliceRev := iterator.NewReverser[int](iterator.FromSlice(expected))
	iterListRev := iterator.NewReverser[int](iterator.FromList(intLis))
	testIteratorBasic[int](t, iterSliceRev, expectedRev)
	testIteratorBasic[int](t, iterListRev, expectedRev)
}
