package typeparamcommon

// Must unwraps error probablity from return value of functions
// so that you can use it in const expr.
//
// It fails if err is non nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func Must3[T any, U any](val1 T, val2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return val1, val2
}

// Escape escapes v to a pointer of v.
//
// It is useful when setting built-in type T to struct fields of *T.
func Escape[T any](v T) *T {
	return &v
}
