package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	_   uintptr
	Len int
}

type Strings []string

func main() {
	fmt.Println(Len([]int{1, 2, 3}))
	fmt.Println(Len([]int{1, 2, 3, 4}))
	fmt.Println(Len(Strings{"1", "2", "3"}))
	fmt.Println(Len(Strings{"1", "2", "3", "4"}))
}

func Len[S ~[]T, T any](v S) int {
	s := *(*Slice)(unsafe.Pointer(&v))
	return s.Len
}


