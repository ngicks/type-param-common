package iterator_test

import (
	"testing"

	"github.com/ngicks/type-param-common/iterator"
	listparam "github.com/ngicks/type-param-common/list-param"
)

func TestZip(t *testing.T) {
	expectedFormer := []int{1, 2, 3, 4, 5}
	expectedLatter := []int{10, 11, 12, 13, 14}

	intLisFormer := listparam.New[int]()
	intLisLatter := listparam.New[int]()
	for _, v := range expectedFormer {
		intLisFormer.PushBack(v)
	}
	for _, v := range expectedLatter {
		intLisLatter.PushBack(v)
	}

	expected := []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14}
	expectedRev := iterator.Iterator[int]{iterator.FromSlice(expected)}.Reverse().Collect()
	{
		iterSliceZipped := iterator.NewChainer[int](iterator.FromSlice(expectedFormer), iterator.FromSlice(expectedLatter))
		iterListZipped := iterator.NewChainer[int](iterator.FromList(intLisFormer), iterator.FromList(intLisLatter))

		testIteratorBasic[int](t, iterSliceZipped, expected)
		testIteratorBasic[int](t, iterListZipped, expected)
	}
	{
		iterSliceZipped := iterator.NewChainer[int](iterator.FromSlice(expectedFormer), iterator.FromSlice(expectedLatter))
		iterListZipped := iterator.NewChainer[int](iterator.FromList(intLisFormer), iterator.FromList(intLisLatter))
		testIteratorBasicBack[int](t, iterSliceZipped, expectedRev)
		testIteratorBasicBack[int](t, iterListZipped, expectedRev)
	}

}
