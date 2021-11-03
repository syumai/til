package main

import (
	"constraints"
	"fmt"
)

func callFn[T constraints.Integer, F func()T](f F) T {
	return f()
}

func main() {
	result := callFn(func() int { return 123 })
	fmt.Println(result)
}
