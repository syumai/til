package main

import "fmt"

type iter[T any] func(func(T) bool) bool

func MapKeys[M ~map[K]V, K comparable, V any](m M) iter[K] {
	return func(yield func(K) bool) bool {
		for k := range m {
			if !yield(k) {
				return false
			}
		}
		return true
	}
}

func MapValues[M ~map[K]V, K comparable, V any](m M) iter[V] {
	return func(yield func(V) bool) bool {
		for _, v := range m {
			if !yield(v) {
				return false
			}
		}
		return true
	}
}

func main() {
	m := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}
	for k := range MapKeys(m) {
		fmt.Println("key:", k)
	}
	for v := range MapValues(m) {
		fmt.Println("value:", v)
	}
}
