// util implements trivial util functions.
package util

// Must unwraps error probability from return value of functions
// so that you can use it in outside of function easily.
//
// It panic if err is non nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Must3 is same as Must but takes 3 args.
//
// It panics if err is non nil.
func Must3[T any, U any](val1 T, val2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return val1, val2
}

// Escape escapes v to a pointer of v.
//
// It is useful when setting built-in type T (e.g. string, int) to struct fields of *T.
func Escape[T any](v T) *T {
	return &v
}

// Assert does type assertion on v.
// v is asserted against T and *T.
// Assert returns true ok when v is either of T or *T.
// If v is nil *T, then asserted is zero value of T.
func Assert[T any](v any) (asserted T, ok bool) {
	asserted, ok = v.(T)
	if !ok {
		p, ok := v.(*T)
		if !ok {
			return asserted, false
		}
		if p != nil {
			asserted = *p
		}
	}

	return asserted, true
}
