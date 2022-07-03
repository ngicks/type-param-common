package slice

// This file sticks its naming convention to it of Rust:
//   - https://doc.rust-lang.org/std/iter/trait.Iterator.html
//   - https://doc.rust-lang.org/std/vec/struct.Vec.html

func Has[T comparable](sl []T, target T) bool {
	return Position(sl, func(v T) bool { return v == target }) >= 0
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
func Position[T any](sl []T, predicate func(v T) bool) int {
	if sl == nil || len(sl) == 0 || predicate == nil {
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
	if sl == nil || len(sl) == 0 || predicate == nil {
		return -1
	}
	for i := len(sl) - 1; i >= 0; i-- {
		if predicate(sl[i]) {
			return i
		}
	}
	return -1
}
