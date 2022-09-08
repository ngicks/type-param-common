package iterator_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ngicks/type-param-common/iterator"
)

func TestChanIterator(t *testing.T) {
	intChan := make(chan int)

	expected := []int{5, 2, 8, 1, 3, 4, 8, 9}
	go func() {
		for _, v := range expected {
			intChan <- v
		}
	}()
	iter := iterator.FromChannel(intChan).ToIterator()

	var collected []int
	for i := 0; i < len(expected); i++ {
		collected = append(collected, iter.MustNext())
	}
	if !reflect.DeepEqual(expected, collected) {
		t.Fatalf("must match. expected = %+v, actual = %+v", expected, collected)
	}

	received := make(chan struct{})

	go func() {
		iter.Next()
		close(received)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	select {
	case <-ctx.Done():
	case <-received:
		t.Fatalf("Next must not return")
	}
	cancel()

	close(intChan)

	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond)
	select {
	case <-ctx.Done():
		t.Fatalf("Next must return at this point")
	case <-received:
	}
	cancel()
}
