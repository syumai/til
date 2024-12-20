package main

import (
	"fmt"
	"unsafe"
)

const MaxCapacity = 1024

func main() {
	ints := MakeInts(3, 5)
	ints[0] = 1
	ints[1] = 2
	ints[2] = 3

	strings := MakeStrings(3, 5)
	strings[0] = "1"
	strings[1] = "2"
	strings[2] = "3"

	fmt.Printf("ints - values: %v, len: %d, cap: %d\n", ints, len(ints), cap(ints))
	fmt.Printf("strings - values: %v, len: %d, cap: %d\n", strings, len(strings), cap(strings))
}

type Slice struct {
	ArrayPtr uintptr
	Len      int
	Cap      int
}

func MakeInts(len_, cap_ int) []int {
	var underlyingArray [MaxCapacity]int
	s := Slice{
		ArrayPtr: uintptr(unsafe.Pointer(&underlyingArray)),
		Len:      len_,
		Cap:      cap_,
	}
	return *(*[]int)(unsafe.Pointer(&s))
}

func MakeStrings(len_, cap_ int) []string {
	var underlyingArray [MaxCapacity]string
	s := Slice{
		ArrayPtr: uintptr(unsafe.Pointer(&underlyingArray)),
		Len:      len_,
		Cap:      cap_,
	}
	return *(*[]string)(unsafe.Pointer(&s))
}
