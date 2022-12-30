package heap_test

import (
	"testing"

	"github.com/ngicks/type-param-common/heap"
	"github.com/ngicks/type-param-common/slice"
)

type Int int

func (i Int) Unwrap() Int {
	return i
}
func (i Int) Less(j heap.Lessable[Int]) bool {
	return i < j.Unwrap()
}

func TestHeapWithAdditionalProps(t *testing.T) {
	t.Run("Exclude", func(t *testing.T) {
		h := heap.NewFilterableHeap[Int, heap.Lessable[Int]]()

		h.Push(Int(7))
		h.Push(Int(4))
		h.Push(Int(1))
		h.Push(Int(6))
		h.Push(Int(5))
		h.Push(Int(3))
		h.Push(Int(2))

		exclude := heap.BuildExcludeFilter(
			-1,
			100,
			func(ent heap.Lessable[Int]) bool { return ent.Unwrap()%2 == 0 },
		)

		lenBefore := h.Len()
		h.Filter(exclude)
		removed := lenBefore - h.Len()

		if removed != 3 {
			t.Fatalf("removed len must be %d, but is %d", 3, removed)
		}

		for i := 1; i <= 7; i = i + 2 {
			popped := h.Pop()
			if int(popped.Unwrap()) != i {
				t.Errorf("pop = %v expected %v", popped, i)
			}
		}

		if h.Len() != 0 {
			t.Errorf("expect empty but size = %v", h.Len())
		}

		h.Push(Int(7))
		h.Push(Int(4))
		h.Push(Int(1))
		h.Push(Int(6))
		h.Push(Int(5))
		h.Push(Int(3))
		h.Push(Int(2))

		exclude = heap.BuildExcludeFilter(
			0,
			3,
			func(ent heap.Lessable[Int]) bool { return ent.Unwrap()%2 == 0 },
		)

		lenBefore = h.Len()
		h.Filter(exclude)
		removed = lenBefore - h.Len()

		if removed != 1 {
			t.Fatalf("removed len must be %d, but is %d", 3, removed)
		}

		for h.Len() != 0 {
			h.Pop()
		}

		h.Push(Int(7))
		h.Push(Int(4))
		h.Push(Int(1))
		h.Push(Int(6))
		h.Push(Int(5))
		h.Push(Int(3))
		h.Push(Int(2))

		exclude = heap.BuildExcludeFilter(
			3,
			6,
			func(ent heap.Lessable[Int]) bool { return ent.Unwrap()%2 == 0 },
		)
		lenBefore = h.Len()
		h.Filter(exclude)
		removed = lenBefore - h.Len()
		if removed != 2 {
			t.Fatalf("removed len must be %d, but is %d", 3, removed)
		}
	})

	t.Run("Clone", func(t *testing.T) {
		h := heap.NewFilterableHeap[Int, heap.Lessable[Int]]()

		h.Push(Int(7))
		h.Push(Int(4))
		h.Push(Int(1))
		h.Push(Int(6))
		h.Push(Int(5))
		h.Push(Int(3))
		h.Push(Int(2))

		cloned := h.Clone()

		var out slice.Deque[int]
		for h.Len() > 0 {
			out.PushBack(int(h.Pop().Unwrap()))
		}

		var outCloned slice.Deque[int]
		for cloned.Len() > 0 {
			outCloned.PushBack(int(cloned.Pop().Unwrap()))
		}

		for i := 0; i < len(out); i++ {
			if out[i] != outCloned[i] {
				t.Fatalf("not equal. expected = %d, actual = %d", out[i], outCloned[i])
			}
		}
	})
}
