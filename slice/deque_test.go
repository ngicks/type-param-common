package slice_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/slice"
)

func TestDequeue(t *testing.T) {
	deque := slice.Deque[int]{}

	if deque.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, deque.Len())
	}

	if v, popped := deque.Pop(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}
	if v, popped := deque.PopBack(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}
	if v, popped := deque.PopFront(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}

	deque.PushFront(1)
	deque.PushBack(2)
	deque.PushFront(3)
	deque.PushBack(4)
	deque.Push(5)
	deque.PushFront(6)

	if deque.Len() != 6 {
		t.Fatalf("wrong len: expected to be %d but is %d", 6, deque.Len())
	}

	{
		expected := []int{6, 3, 1, 2, 4, 5}
		if !reflect.DeepEqual(([]int)(deque), expected) {
			t.Fatalf("incorrect push behavior: expected to be %v, but is %v", expected, deque)
		}
	}

	popped := []int{}
	var v int
	v, _ = deque.PopFront()
	popped = append(popped, v)
	v, _ = deque.Pop()
	popped = append(popped, v)
	v, _ = deque.PopBack()
	popped = append(popped, v)
	v, _ = deque.PopFront()
	popped = append(popped, v)
	v, _ = deque.PopBack()
	popped = append(popped, v)
	v, _ = deque.PopFront()
	popped = append(popped, v)

	if deque.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, deque.Len())
	}
	{
		expected := []int{6, 5, 4, 3, 2, 1}
		if !reflect.DeepEqual(popped, expected) {
			t.Fatalf("incorrect pop behavior: expected to be %v, but is %v", expected, popped)
		}
	}
}
