package slices

func Collect[T any](iter func(func(T) bool) bool) []T {
	var s []T
	for v := range iter {
		s = append(s, v)
	}
	return s
}
