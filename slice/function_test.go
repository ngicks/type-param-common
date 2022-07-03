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
	if slice.Has(nil, 10) {
		t.Fatalf("returned value must be false if input slice is nil")
	}

	if idx, ok := slice.Find(sl, func(v int) bool { return v > 2 }); !ok || idx != 3 {
		t.Fatalf("index must be %d, but %d", 3, idx)
	}
	if idx, ok := slice.Find(sl, func(v int) bool { return v == 6 }); ok {
		t.Fatalf("returned ok must be false if no value matches to given predicate, but found %d", idx)
	}
	if _, ok := slice.Find(nil, func(v int) bool { return v == 1 }); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.Find(sl, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.Find[int](nil, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}

	if idx, ok := slice.FindLast(sl, func(v int) bool { return v > 2 }); !ok || idx != 5 {
		t.Fatalf("index must be %d, but %d", 5, idx)
	}
	if idx, ok := slice.FindLast(sl, func(v int) bool { return v == 6 }); ok {
		t.Fatalf("returned ok must be false if no value matches to given predicate, but found %d", idx)
	}
	if _, ok := slice.FindLast(nil, func(v int) bool { return v == 1 }); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.FindLast(sl, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.FindLast[int](nil, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}

	if p := slice.Position(sl, func(v int) bool { return v == 1 }); p != 0 {
		t.Fatalf("must be 0, but %d", p)
	}
	if p := slice.Position(sl, func(v int) bool { return v == 6 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.Position(nil, func(v int) bool { return v == 1 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.Position(sl, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.Position[int](nil, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}

	if p := slice.PositionLast(sl, func(v int) bool { return v == 1 }); p != 4 {
		t.Fatalf("must be 4, but %d", p)
	}
	if p := slice.PositionLast(sl, func(v int) bool { return v == 6 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.PositionLast(nil, func(v int) bool { return v == 1 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.PositionLast(sl, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.PositionLast[int](nil, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
}
