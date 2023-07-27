package stream

func Map[T, U any](it func(func(T) bool) bool, f func(T) U) func(func(U) bool) bool {
	return func(yield func(U) bool) bool {
		for v := range it {
			if !yield(f(v)) {
				return false
			}
		}
		return true
	}
}
