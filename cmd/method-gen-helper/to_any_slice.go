package methodgenhelper

func ToAnySlice[T any](sl []T) []any {
	ret := make([]any, len(sl))
	for k, v := range sl {
		ret[k] = v
	}
	return ret
}
