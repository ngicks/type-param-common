package slice

import (
	"reflect"
)

// This file sticks its naming convention to it of Rust:
//   - https://doc.rust-lang.org/std/iter/trait.Iterator.html
//   - https://doc.rust-lang.org/std/vec/struct.Vec.html

func Append[T any](sl []T, elements ...T) []T {
	return append(sl, elements...)
}

func Clone[T any](sl []T) []T {
	cloned := make([]T, len(sl))
	copy(cloned, sl)
	return cloned
}

func Eq[T comparable](left, right []T) bool {
	if len(left) != len(right) {
		return false
	}

	// Comparing types which is constrained by comparable
	// may not be actually comparable in GO 1.20 or later version.
	// It might cause a runtime panic.
	var zero T
	isComparable := reflect.TypeOf(zero).Comparable()
	if !isComparable {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func Find[T any](sl []T, predicate func(v T) bool) (v T, found bool) {
	p := Position(sl, predicate)
	if p < 0 {
		return
	}
	return sl[p], true
}

func FindLast[T any](sl []T, predicate func(v T) bool) (v T, found bool) {
	p := PositionLast(sl, predicate)
	if p < 0 {
		return
	}
	return sl[p], true
}

func Get[T any](sl []T, index uint) (v T, ok bool) {
	if index >= uint(len(sl)) {
		return
	}
	return sl[index], true
}

func Has[T comparable](sl []T, target T) bool {
	return Position(sl, func(v T) bool { return v == target }) >= 0
}

func Insert[T any](sl []T, index uint, ele T) []T {
	prefix, suffix := sl[:index], sl[index:]
	// Avoiding mutation of input slice.
	// This effectively clones input.
	prefix = append([]T{}, prefix...)
	return append(prefix, append([]T{ele}, suffix...)...)
}

func Position[T any](sl []T, predicate func(v T) bool) int {
	if len(sl) == 0 || predicate == nil {
		return -1
	}
	for idx, v := range sl {
		if predicate(v) {
			return idx
		}
	}
	return -1
}

func PositionLast[T any](sl []T, predicate func(v T) bool) int {
	if len(sl) == 0 || predicate == nil {
		return -1
	}
	for i := len(sl) - 1; i >= 0; i-- {
		if predicate(sl[i]) {
			return i
		}
	}
	return -1
}

func Prepend[T any](sl []T, elements ...T) []T {
	for i := 0; i < len(elements); i++ {
		sl = append([]T{elements[i]}, sl...)
	}
	return sl
}
