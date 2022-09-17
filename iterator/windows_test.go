package iterator_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/iterator"
)

func TestWindower(t *testing.T) {
	c := iterator.NewWindower(iterator.FromRange(0, 12).Collect(), 5)

	for i := 0; i <= 12-5; i++ {
		next, ok := c.Next()
		if len(next) == 0 || !ok {
			t.Fatalf("must be size of 5 but %d", len(next))
		}
		expected := []int{i, i + 1, i + 2, i + 3, i + 4}
		if !reflect.DeepEqual(expected, next) {
			t.Fatalf("not Equal, expected = %+v, actual = %+v", expected, next)
		}
	}

	next, ok := c.Next()
	if len(next) != 0 || ok {
		t.Fatalf("must be ended but next returns ok = true, %+v", next)
	}
}
