package atomicparam_test

import (
	"bytes"
	"reflect"
	"testing"

	atomicparam "github.com/ngicks/type-param-common/sync-param/atomic-param"
)

type dummyEnum uint8

const (
	invalid dummyEnum = iota
	foo
	bar
	baz
)

func TestAtomic(t *testing.T) {
	t.Run("dummyEnum", func(t *testing.T) {
		values := []dummyEnum{foo, bar, baz}
		testValue(t, atomicparam.NewValue[dummyEnum](), values, true)
		testValue(t, atomicparam.Value[dummyEnum]{}, values, false)
	})
	t.Run("string", func(t *testing.T) {
		values := []string{"foo", "bar", "baz"}
		testValue(t, atomicparam.NewValue[string](), values, true)
		testValue(t, atomicparam.Value[string]{}, values, false)
	})
	t.Run("*bytes.Buffer", func(t *testing.T) {
		values := []*bytes.Buffer{
			bytes.NewBuffer(make([]byte, 0)),
			nil,
			bytes.NewBuffer(make([]byte, 0)),
		}
		testValue(t, atomicparam.NewValue[*bytes.Buffer](), values, true)
		testValue(t, atomicparam.Value[*bytes.Buffer]{}, values, false)
	})
}

func testValue[T comparable](t *testing.T, val atomicparam.Value[T], values []T, shouldReturnNormally bool) {
	var v T

	var returnedNormally bool
	func() {
		defer func() {
			if recv := recover(); recv != nil {
				if shouldReturnNormally {
					t.Error(recv)
				}
			}
		}()
		v = val.Load()
		if !reflect.ValueOf(v).IsZero() {
			t.Fatalf("must be zero value")
		}
		returnedNormally = true
	}()

	if returnedNormally != shouldReturnNormally {
		t.Fatalf(
			"mismatched cond, shouldReturnNormally = %t, but returnedNormally = %t",
			shouldReturnNormally,
			returnedNormally,
		)
	}

	func() {
		defer func() {
			if recv := recover(); recv != nil {
				if shouldReturnNormally {
					t.Error(recv)
				}
			}
		}()
		v = val.Swap(values[0])
		if !reflect.ValueOf(v).IsZero() {
			t.Fatalf("must be zero value")
		}
		returnedNormally = true
	}()

	if returnedNormally != shouldReturnNormally {
		t.Fatalf(
			"mismatched cond, shouldReturnNormally = %t, but returnedNormally = %t",
			shouldReturnNormally,
			returnedNormally,
		)
	}

	val.Store(values[0])
	v = val.Load()
	if !reflect.DeepEqual(v, values[0]) {
		t.Fatalf("unmatched result, expected=%v, actual=%v", "foo", v)
	}

	val.Swap(values[1])

	v = val.Load()
	if !reflect.DeepEqual(v, values[1]) {
		t.Fatalf("unmatched result, expected=%v, actual=%v", bar, v)
	}

	var swapped bool
	swapped = val.CompareAndSwap(values[1], values[2])
	if !swapped {
		t.Fatalf("must be swapped")
	}
	swapped = val.CompareAndSwap(values[1], values[2])
	if swapped {
		t.Fatalf("must not be swapped")
	}
}
