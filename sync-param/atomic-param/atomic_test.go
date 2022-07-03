package atomicparam_test

import (
	"reflect"
	"testing"

	atomicparam "github.com/ngicks/type-param-common/sync-param/atomic-param"
)

func TestAtomic(t *testing.T) {
	t.Run("non pointer type", func(t *testing.T) {
		// string technically is a pointer type, being Vector of uint8.
		// Golang treats string as non-pointer type.
		val := atomicparam.NewValue[string]()

		var v string

		v = val.Load()
		if !reflect.ValueOf(v).IsZero() {
			t.Fatalf("must be zero value")
		}

		val.Store("foo")
		v = val.Load()
		if !reflect.DeepEqual(v, "foo") {
			t.Fatalf("unmatched result, expected=%v, actual=%v", "foo", v)
		}

		val.Swap("bar")

		v = val.Load()
		if !reflect.DeepEqual(v, "bar") {
			t.Fatalf("unmatched result, expected=%v, actual=%v", "bar", v)
		}

		var swapped bool
		swapped = val.CompareAndSwap("bar", "baz")
		if !swapped {
			t.Fatalf("must be swapped")
		}
		swapped = val.CompareAndSwap("bar", "baz")
		if swapped {
			t.Fatalf("must not be swapped")
		}
	})

	t.Run("slice", func(t *testing.T) {
		val := atomicparam.NewValue[*[]string]()

		var v *[]string

		v = val.Load()
		if !reflect.ValueOf(v).IsZero() {
			t.Fatalf("must be zero value")
		}

		val.Store(&[]string{"foo", "bar", "baz"})
		v = val.Load()
		if !reflect.DeepEqual(v, &[]string{"foo", "bar", "baz"}) {
			t.Fatalf("unmatched result, expected=%v, actual=%v", []string{"foo", "bar", "baz"}, *v)
		}

		val.Swap(&[]string{"qux", "quux", "corge"})

		v = val.Load()
		if !reflect.DeepEqual(v, &[]string{"qux", "quux", "corge"}) {
			t.Fatalf("unmatched result, expected=%v, actual=%v", []string{"qux", "quux", "corge"}, *v)
		}

		var swapped bool
		sl := []string{"grault", "garply", "waldo"}
		swapped = val.CompareAndSwap(&[]string{"qux", "quux", "corge"}, &sl)
		if swapped {
			// it is pointer value.
			t.Fatalf("must not be swapped")
		}
		val.Store(&sl)
		swapped = val.CompareAndSwap(&sl, &[]string{"qux", "quux", "corge"})
		if !swapped {
			t.Fatalf("must be swapped")
		}
		v = val.Load()
		if !reflect.DeepEqual(v, &[]string{"qux", "quux", "corge"}) {
			t.Fatalf("value mismatch, expeceted=%v, actual=%v", []string{"qux", "quux", "corge"}, *v)
		}
	})
}
