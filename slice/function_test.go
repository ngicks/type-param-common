package slice_test

import (
	"testing"

	"github.com/ngicks/type-param-common/slice"
)

func TestFunc(t *testing.T) {
	sl := []int{1, 2, 3, 4, 1, 5}

	if !slice.Has(sl, 1) {
		t.Fatalf("must be true")
	}
	if slice.Has(sl, 10) {
		t.Fatalf("must be false")
	}
	if slice.FindIndex(sl, 1) != 0 {
		t.Fatalf("must be 0, but %d", slice.FindIndex(sl, 1))
	}
	if slice.FindIndexLast(sl, 1) != 4 {
		t.Fatalf("must be 4, but %d", slice.FindIndexLast(sl, 1))
	}
}
