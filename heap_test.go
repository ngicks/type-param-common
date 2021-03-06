package typeparamcommon_test

import (
	"testing"
	"time"

	typeparamcommon "github.com/ngicks/type-param-common"
)

func TestSimpleHeap(t *testing.T) {
	// Seeing basic delegation.
	t.Run("int heap", func(t *testing.T) {
		h, inter := typeparamcommon.MakeMinHeap[int]()
		ans := []int{3, 4, 4, 5, 6}
		h.Push(5)
		h.Push(4)
		h.Push(6)
		h.Push(3)
		h.Push(4)

		for _, i := range ans {
			popped := h.Pop()
			if popped != i {
				t.Errorf("pop = %v expected %v", popped, i)
			}
		}
		if inter.Len() != 0 {
			t.Errorf("expect empty but size = %v", inter.Len())
		}
	})

	t.Run("struct heap", func(t *testing.T) {
		type testStruct struct {
			t time.Time
		}
		less := func(i, j *testStruct) bool {
			return i.t.Before(j.t)
		}

		h, inter := typeparamcommon.MakeHeap(less)
		ans := []*testStruct{
			{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			{t: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)},
			{t: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)},
			{t: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)},
			{t: time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC)},
		}
		h.Push(ans[2])
		h.Push(ans[1])
		h.Push(ans[3])
		h.Push(ans[0])
		h.Push(ans[4])

		for _, i := range ans {
			popped := h.Pop()
			if popped.t != i.t {
				t.Errorf("pop = %v expected %v", popped.t, i.t)
			}
		}
		if inter.Len() != 0 {
			t.Errorf("expect empty but size = %v", inter.Len())
		}
	})
}

type AdditionalPropHeap[T any] struct {
	*typeparamcommon.HeapWrapper[T]
	internal *typeparamcommon.SliceInterface[T]
}

func (aph *AdditionalPropHeap[T]) Len() int {
	return aph.internal.Len()
}
func (aph *AdditionalPropHeap[T]) Exclude(filter func(ent T) bool, start, end int) (removed []T) {
	if filter == nil {
		return
	}

	if start < 0 {
		start = 0
	} else if start >= len(aph.internal.Inner) {
		return
	}
	if end > len(aph.internal.Inner) {
		end = len(aph.internal.Inner)
	}

	if start > end {
		return
	}

	for i := start; i < end; i++ {
		if filter(aph.internal.Inner[i]) {
			removed = append(removed, aph.internal.Inner[i])
			aph.internal.Inner = append(aph.internal.Inner[:i], aph.internal.Inner[i+1:]...)
			end--
			i--
		}
	}

	aph.Init()
	return removed
}

func TestHeapWithAdditionalProps(t *testing.T) {
	t.Run("Exclude", func(t *testing.T) {
		h_, inter := typeparamcommon.MakeMinHeap[int]()
		h := &AdditionalPropHeap[int]{
			HeapWrapper: h_,
			internal:    inter,
		}

		h.Push(7)
		h.Push(4)
		h.Push(1)
		h.Push(6)
		h.Push(5)
		h.Push(3)
		h.Push(2)

		removed := h.Exclude(func(ent int) bool { return ent%2 == 0 }, -1, 100)

		for i := 1; i <= 7; i = i + 2 {
			popped := h.Pop()
			if popped != i {
				t.Errorf("pop = %v expected %v", popped, i)
			}
		}
		if h.Len() != 0 {
			t.Errorf("expect empty but size = %v", h.Len())
		}

		h.Push(1)
		h.Push(3)
		h.Push(5)
		h.Push(7)
		h.Push(9)
		h.Push(11)
		h.Push(13)

		removed = h.Exclude(func(ent int) bool { return ent%2 != 0 }, 0, 3)

		if len(removed) != 3 {
			t.Fatalf("removed len must be %d, but is %d: %v", 3, len(removed), removed)
		}
		for h.Len() != 0 {
			h.Pop()
		}

		h.Push(1)
		h.Push(3)
		h.Push(5)
		h.Push(7)
		h.Push(9)
		h.Push(11)
		h.Push(13)

		removed = h.Exclude(func(ent int) bool { return ent%2 != 0 }, 3, 6)

		if len(removed) != 3 {
			t.Fatalf("removed len must be %d, but is %d: %v", 3, len(removed), removed)
		}
	})
}
