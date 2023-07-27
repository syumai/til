package maps

type iter[T any] func(func(T) bool) bool

func Keys[M ~map[K]V, K comparable, V any](m M) iter[K] {
	return func(yield func(K) bool) bool {
		for k := range m {
			if !yield(k) {
				return false
			}
		}
		return true
	}
}

func Values[M ~map[K]V, K comparable, V any](m M) iter[V] {
	return func(yield func(V) bool) bool {
		for _, v := range m {
			if !yield(v) {
				return false
			}
		}
		return true
	}
}
