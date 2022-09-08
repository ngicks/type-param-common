package iterator_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/iterator"
	listparam "github.com/ngicks/type-param-common/list-param"
)

type skipNTestCase struct {
	input        []int
	skipN        int
	expectedSize int
	expected     []int
}

func TestSkipNSizeHint(t *testing.T) {
	testCases := []skipNTestCase{
		{
			input:        []int{1, 2, 3},
			skipN:        3,
			expectedSize: 0,
		},
		{
			input:        []int{3, 5, 6, 76, 1, 34, 8},
			skipN:        2,
			expectedSize: 5,
		},
		{
			input:        []int{3, 5, 6, 76, 1, 34, 8},
			skipN:        0,
			expectedSize: 7,
		},
		{
			input:        []int{3, 5, 6},
			skipN:        120,
			expectedSize: 0,
		},
	}

	for _, testCase := range testCases {
		iter := iterator.FromSlice(testCase.input).ToIterator().SkipN(testCase.skipN)
		if size := iter.SizeHint(); size != testCase.expectedSize {
			t.Fatalf("mismatched: expected = %d, actual = %d", testCase.expectedSize, size)
		}
	}

	iter := iterator.FromList(listparam.New[int]()).ToIterator().SkipN(5)
	if size := iter.SizeHint(); size != -1 {
		t.Fatalf("mismatched: expected = %d, actual = %d", -1, size)
	}
}

func TestSkipN(t *testing.T) {
	testCases := []skipNTestCase{
		{
			input:    []int{1, 2, 3},
			skipN:    3,
			expected: []int{},
		},
		{
			input:    []int{1, 2, 3},
			skipN:    0,
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{3, 5, 6, 76, 1, 34, 8},
			skipN:    5,
			expected: []int{34, 8},
		},
		{
			input:    []int{3, 5, 6},
			skipN:    120,
			expected: []int{},
		},
	}

	for _, testCase := range testCases {
		iter := iterator.FromSlice(testCase.input).ToIterator().SkipN(testCase.skipN)
		if collected := iter.Collect(); !reflect.DeepEqual(testCase.expected, collected) {
			t.Errorf("mismatched: expected = %+v, actual = %+v", testCase.expected, collected)
		}
	}
}

func TestSkipWhileSizeHint(t *testing.T) {
	for _, iter := range []iterator.Iterator[int]{
		iterator.FromList(listparam.New[int]()).ToIterator(),
		iterator.FromSlice([]int{1, 2, 3}).ToIterator(),
		iterator.NewRange(0, 25).ToIterator(),
	} {
		size := iter.SkipWhile(func(i int) bool { return true }).SizeHint()
		if size != -1 {
			t.Fatalf("mismatched: expected = %d, actual = %d", -1, size)
		}
	}
}

type skipWhileTestCase struct {
	input    []int
	skipIf   func(int) bool
	expected []int
}

func TestSkipWhile(t *testing.T) {
	testCases := []skipWhileTestCase{
		{
			input:    []int{1, 2, 3},
			skipIf:   func(i int) bool { return true },
			expected: []int{},
		},
		{
			input:    []int{1, 2, 3},
			skipIf:   func(i int) bool { return false },
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{3, 5, 6, 76, 1, 34, 8},
			skipIf:   func(i int) bool { return i < 15 },
			expected: []int{76, 1, 34, 8},
		},
		{
			input:    []int{3, 5, 6},
			skipIf:   func(i int) bool { return i%2 != 0 },
			expected: []int{6},
		},
	}

	for _, testCase := range testCases {
		{
			iter := iterator.FromSlice(testCase.input).ToIterator().SkipWhile(func(i int) bool { return false })
			if collected := iter.Collect(); !reflect.DeepEqual(testCase.input, collected) {
				t.Errorf("mismatched: expected = %+v, actual = %+v", testCase.input, collected)
			}
		}
		{
			iter := iterator.FromSlice(testCase.input).ToIterator().SkipWhile(testCase.skipIf)
			if collected := iter.Collect(); !reflect.DeepEqual(testCase.expected, collected) {
				t.Errorf("mismatched: expected = %+v, actual = %+v", testCase.expected, collected)
			}
		}
	}
}
