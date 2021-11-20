package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	ArrayPtr uintptr
	Len      int
	Cap      int
}

func main() {
	var s []int
	for i := 0; i < 20; i++ {
		s = AppendInts(s, i, i, i)
		fmt.Printf("len: %d, cap: %d, v: %v\n", IntsLen(s), IntsCap(s), s)
	}
}

func AppendInts(v []int, elems ...int) []int {
	s := *(*Slice)(unsafe.Pointer((&v)))
	if s.Cap-s.Len >= IntsLen(elems) {
		s1 := Slice{
			ArrayPtr: s.ArrayPtr,
			Len:      s.Len + IntsLen(elems),
			Cap:      s.Cap,
		}
		result := *(*[]int)(unsafe.Pointer(&s1))
		copy(result[s.Len:], elems)
		return result
	}

	var (
		arrayPtr uintptr
		cap_     int
		len_     = s.Len + IntsLen(elems)
	)
	switch {
	case len_ < 1<<1:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 1]int{}))
		cap_ = 1 << 1
	case len_ < 1<<2:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 2]int{}))
		cap_ = 1 << 2
	case len_ < 1<<3:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 3]int{}))
		cap_ = 1 << 3
	case len_ < 1<<4:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 4]int{}))
		cap_ = 1 << 4
	case len_ < 1<<5:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 5]int{}))
		cap_ = 1 << 5
	case len_ < 1<<6:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 6]int{}))
		cap_ = 1 << 6
	case len_ < 1<<7:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 7]int{}))
		cap_ = 1 << 7
	case len_ < 1<<8:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 8]int{}))
		cap_ = 1 << 8
	case len_ < 1<<9:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 9]int{}))
		cap_ = 1 << 9
	case len_ < 1<<10:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 10]int{}))
		cap_ = 1 << 10
	}
	s2 := Slice{
		ArrayPtr: arrayPtr,
		Len:      len_,
		Cap:      cap_,
	}
	result := *(*[]int)(unsafe.Pointer(&s2))
	copy(result, v)
	copy(result[IntsLen(v):], elems)
	return result
}

func IntsLen(v []int) int {
	s := *(*Slice)(unsafe.Pointer(&v))
	return s.Len
}

func IntsCap(v []int) int {
	s := *(*Slice)(unsafe.Pointer(&v))
	return s.Cap
}
