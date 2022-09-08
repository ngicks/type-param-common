package iterator_test

import (
	"reflect"
	"testing"

	"github.com/ngicks/type-param-common/iterator"
)

func TestChain(t *testing.T) {
	expectedFormer := []int{1, 2, 3, 4, 5}
	expectedLatter := []int{10, 11, 12, 13, 14}

	{
		expected := []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14}
		f := iterator.FromSlice(expectedFormer).ToIterator()
		l := iterator.FromSlice(expectedLatter).ToIterator()
		chained := f.Chain(l)

		if _, ok := chained.SeIterator.(*iterator.Chainer[int]); !ok {
			t.Fatalf("internal type must be *iterator.Chainer[int] but %T", chained.SeIterator)
		}

		collected := chained.Collect()
		if !reflect.DeepEqual(expected, collected) {
			t.Fatalf("must be deeply equal. expected = %+v, actual = %+v", expected, collected)
		}
	}
	{
		expected := []int{10, 11, 12, 13, 14, 1, 2, 3, 4, 5}
		f := iterator.FromSlice(expectedFormer).ToIterator()
		l := iterator.FromSlice(expectedLatter).ToIterator()
		chained := l.Chain(f)

		if _, ok := chained.SeIterator.(*iterator.Chainer[int]); !ok {
			t.Fatalf("internal type must be *iterator.Chainer[int] but %T", chained.SeIterator)
		}

		collected := chained.Collect()
		if !reflect.DeepEqual(expected, collected) {
			t.Fatalf("must be deeply equal. expected = %+v, actual = %+v", expected, collected)
		}

	}

}

func TestChainReversed(t *testing.T) {
	expectedFormer := []int{1, 2, 3, 4, 5}
	expectedLatter := []int{10, 11, 12, 13, 14}
	expected := []int{1, 2, 14, 13, 3, 4, 5, 12, 11, 10}

	f := iterator.FromSlice(expectedFormer).ToIterator()
	l := iterator.FromSlice(expectedLatter).ToIterator()
	chained := f.Chain(l)

	answer := []int{}

	answer = append(answer, chained.MustNext())
	answer = append(answer, chained.MustNext())
	chained = chained.MustReverse()
	answer = append(answer, chained.MustNext())
	answer = append(answer, chained.MustNext())
	chained = chained.MustReverse()
	answer = append(answer, chained.MustNext())
	answer = append(answer, chained.MustNext())
	answer = append(answer, chained.MustNext())
	chained = chained.MustReverse()
	answer = append(answer, chained.MustNext())
	answer = append(answer, chained.MustNext())
	answer = append(answer, chained.MustNext())

	if _, ok := chained.Next(); ok {
		t.Fatal("must be drained")
	}
	if _, ok := chained.MustReverse().Next(); ok {
		t.Fatal("must be drained")
	}

	if !reflect.DeepEqual(expected, answer) {
		t.Fatalf("must be deeply equal. expected = %+v, actual = %+v", expected, answer)
	}
}
