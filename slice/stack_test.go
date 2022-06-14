package slice_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/slice"
)

func TestStack(t *testing.T) {
	stack := slice.Stack[int]{}

	if stack.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, stack.Len())
	}
	if v, popped := stack.Pop(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}

	stack.Push(6)
	stack.Push(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(5)

	if stack.Len() != 6 {
		t.Fatalf("wrong len: expected to be %d but is %d", 6, stack.Len())
	}
	{
		expected := []int{6, 3, 1, 2, 4, 5}
		if !reflect.DeepEqual(([]int)(stack), expected) {
			t.Fatalf("incorrect push behavior: expected to be %v, but is %v", expected, stack)
		}
	}

	popped := []int{}
	var v int
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)

	if stack.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, stack.Len())
	}
	{
		reversed := []int{6, 3, 1, 2, 4, 5}
		expected := make([]int, len(reversed))
		for idx, v := range reversed {
			expected[len(reversed)-(idx+1)] = v
		}
		if !reflect.DeepEqual(popped, expected) {
			t.Fatalf("incorrect pop behavior: expected to be %v, but is %v", expected, popped)
		}
	}
}
