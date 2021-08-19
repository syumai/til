package main

import "fmt"

type Integer interface {
	int
}

type Map[K comparable, V any] interface {
	~map[K]V
}

func Clone[M Map[K, V], K comparable, V any](a M) M {
	b := make(map[K]V)
	for key, value := range a {
		b[key] = value
	}
	return b
}

func Filter[M Map[K, V], K comparable, V any](m M, keep func(K, V) bool) {
	for k, v := range m {
		if !keep(k, v) {
			delete(m, k)
		}
	}
}

func SeparateEvenOddMaps[M Map[K, V], K comparable, V Integer](m M) (even, odd M) {
	even, odd = Clone[M, K, V](m), Clone[M, K, V](m)
	Filter(even, func (k K, v V) bool { return v % 2 == 0 })
	Filter(odd, func (k K, v V) bool { return v % 2 == 1 })
	return
}

func main() {
	strIntMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}
	even, odd := SeparateEvenOddMaps(strIntMap)
	fmt.Println(even)
	fmt.Println(odd)
}
