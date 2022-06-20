package slice

func Has[T comparable](sl []T, target T) bool {
	return Find(sl, target) >= 0
}

func Find[T comparable](sl []T, target T) int {
	if sl == nil || len(sl) == 0 {
		return -1
	}
	for idx, v := range sl {
		if target == v {
			return idx
		}
	}
	return -1
}
func FindLast[T comparable](sl []T, target T) int {
	if sl == nil || len(sl) == 0 {
		return -1
	}
	for i := len(sl) - 1; i >= 0; i-- {
		if target == sl[i] {
			return i
		}
	}
	return -1
}
