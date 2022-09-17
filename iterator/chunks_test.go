package iterator_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/iterator"
)

func TestChunker(t *testing.T) {
	c := iterator.NewChunker(iterator.FromRange(0, 12).Collect(), 5)

	var next []int
	var ok bool

	next, ok = c.Next()
	if len(next) == 0 || !ok {
		t.Fatalf("must be size of 5 but %d", len(next))
	}
	if !reflect.DeepEqual([]int{0, 1, 2, 3, 4}, next) {
		t.Fatalf("not Equal, expected = %+v, actual = %+v", []int{0, 1, 2, 3, 4}, next)
	}

	next, ok = c.Next()
	if len(next) == 0 || !ok {
		t.Fatalf("must be size of 5 but %d", len(next))
	}
	if !reflect.DeepEqual([]int{5, 6, 7, 8, 9}, next) {
		t.Fatalf("not Equal, expected = %+v, actual = %+v", []int{5, 6, 7, 8, 9}, next)
	}

	next, ok = c.Next()
	if len(next) != 2 || !ok {
		t.Fatalf("must be size of 2 but %d", len(next))
	}
	if !reflect.DeepEqual([]int{10, 11}, next) {
		t.Fatalf("not Equal, expected = %+v, actual = %+v", []int{10, 11}, next)
	}

	next, ok = c.Next()
	if len(next) != 0 || ok {
		t.Fatalf("must be ended but next returns ok = true")
	}
}
