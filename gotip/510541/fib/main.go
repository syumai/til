package main

import "fmt"

func fibonacci(n int) func(func(int) bool) bool {
	a, b := 1, 0
	return func(yield func(int) bool) bool {
		for i := 0; i < n; i++ {
			a, b = b, a+b
			if !yield(a) {
				// 				return false
			}
		}
		return true
	}
}

func main() {
	f := fibonacci(10)
	for v := range f {
		fmt.Println(v)
		if v == 8 {
			break
		}
	}
}
