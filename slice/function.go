package slice

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
	// avoiding mutation of input slice.
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
