package iterator_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/iterator"
	listparam "github.com/ngicks/type-param-common/list-param"
)

func testIteratorBasic[T any](t *testing.T, iter iterator.Nexter[T], expected any) {
	earned := make([]T, 0)
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		earned = append(earned, next)
	}

	if !reflect.DeepEqual(earned, expected) {
		t.Fatalf("interator incorrect behavior: expected = %v, actual = %v", expected, earned)
	}
}

func testIteratorBasicBack[T any](t *testing.T, iter iterator.NextBacker[T], expected any) {
	earned := make([]T, 0)
	for next, ok := iter.NextBack(); ok; next, ok = iter.NextBack() {
		earned = append(earned, next)
	}

	if !reflect.DeepEqual(earned, expected) {
		t.Fatalf("interator incorrect behavior: expected = %v, actual = %v", expected, earned)
	}
}

func TestIterator(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	iterSlice := iterator.FromSlice(expected)
	testIteratorBasic[int](t, iterSlice, expected)

	intLis := listparam.New[int]()
	for _, v := range expected {
		intLis.PushBack(v)
	}
	iterList := iterator.FromList(intLis)
	testIteratorBasic[int](t, iterList, expected)
}

// input must be iterator of []int{1,2,3,4,5}
func testIterBack(t *testing.T, iter iterator.DeIterator[int]) {
	assert := func(exexpected int, actual int) {
		if exexpected != actual {
			t.Fatalf("expected = %d, actual = %v", exexpected, actual)
		}
	}

	var val int
	var ok bool
	val, _ = iter.NextBack()
	assert(5, val)
	val, _ = iter.NextBack()
	assert(4, val)
	val, _ = iter.Next()
	assert(1, val)
	val, _ = iter.NextBack()
	assert(3, val)
	val, _ = iter.Next()
	assert(2, val)
	val, ok = iter.Next()
	// zero value.
	assert(0, val)
	if ok {
		t.Fatalf("iterator must be exhausted: %v", iter)
	}
}

func TestIterBack(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	iterSlice := iterator.FromSlice(expected)
	testIterBack(t, iterSlice)

	intLis := listparam.New[int]()
	for _, v := range expected {
		intLis.PushBack(v)
	}
	iterList := iterator.FromList(intLis)
	testIterBack(t, iterList)
}

func TestCombination(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	iter := iterator.FromSlice(expected)

	reduced := iterator.Reduce[int](
		// 1, 4, 9, 16, 25, 36, 49, 64
		iterator.Map[int](iter, func(next int) int { return next * next }).
			// 9, 16, 36, 49
			Select(
				func(i int) bool { return (i%10)%3 == 0 },
			).
			// 49, 36, 16, 9
			Reverse().
			// 16, 9
			SkipN(2),
		func(accumulator, next int) int {
			return accumulator + next
		},
		0,
	)

	if reduced != 25 {
		t.Fatalf("must be %d, but %d", 25, reduced)
	}
}
