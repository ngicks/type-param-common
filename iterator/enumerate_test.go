package iterator_test

import (
	"testing"

	"github.com/ngicks/type-param-common/iterator"
)

func TestEnumerate(t *testing.T) {
	input := []string{"foo", "bar", "baz", "qux", "quux"}
	iter := iterator.FromSlice(input).ToIterator()
	enum := iterator.Iterator[iterator.EnumerateEnt[string]]{
		SeIterator: iterator.Enumerate[string](iter),
	}

	for idx, v := range input {
		next := enum.NextMust()
		if next.Count != idx {
			t.Fatalf("%d, %d\n", idx, next.Count)
		}
		if next.Next != v {
			t.Fatalf("%s, %s\n", v, next.Next)
		}
	}

	if _, ok := enum.Next(); ok {
		t.Fatal("must be exhausted")
	}
}
