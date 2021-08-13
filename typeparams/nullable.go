package main

import "fmt"

type Nilable[T any] interface {
	T
}

type NilableInt Nilable[int]
type NilableFloat64 Nilable[float64]

func PrintNilableInt[T NilableInt](v T) {
	if v == nil {
		fmt.Println("value is nil")
	}
	fmt.Println(v)
}

func main() {
	/*
	var ni NullableInt
	ni = 1
	ni = 100
	ni = nil
	*/
	PrintNilableInt(100)
	PrintNilableInt(nil)
}
