package iterator_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/iterator"
)

func testIterator[T any](t *testing.T, input []T) {
	t.Run("Next", func(t *testing.T) {
		iter := iterator.FromSlice(input).ToIterator()
		nexted := []T{}
		for {
			next, ok := iter.Next()
			if !ok {
				break
			}
			nexted = append(nexted, next)
		}
		if !reflect.DeepEqual(input, nexted) {
			t.Fatalf("must equal. actual = %+v, expected = %+v", nexted, input)
		}
	})

	t.Run("Collect", func(t *testing.T) {
		iter := iterator.FromSlice(input).ToIterator()
		collected := iter.Collect()
		if !reflect.DeepEqual(input, collected) {
			t.Fatalf("must equal. actual = %+v, expected = %+v", collected, input)
		}
	})

	t.Run("ForEach", func(t *testing.T) {
		iter := iterator.FromSlice(input).ToIterator()
		eached := []T{}
		iter.ForEach(func(i T) {
			eached = append(eached, i)
		})
		if !reflect.DeepEqual(input, eached) {
			t.Fatalf("must equal. actual = %+v, expected = %+v", eached, input)
		}
	})
}
