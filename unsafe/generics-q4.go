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
		s = Append(s, i, i, i)
		fmt.Printf("len: %d, cap: %d, v: %v\n", Len(s), Cap(s), s)
	}
}

func Append[S ~[]T, T any](v S, elems ...T) S {
	s := *(*Slice)(unsafe.Pointer((&v)))
	if s.Cap-s.Len >= Len(elems) {
		s1 := Slice{
			ArrayPtr: s.ArrayPtr,
			Len:      s.Len + Len(elems),
			Cap:      s.Cap,
		}
		result := *(*S)(unsafe.Pointer(&s1))
		copy(result[s.Len:], elems)
		return result
	}

	var (
		arrayPtr uintptr
		cap_     int
		len_     = s.Len + Len(elems)
	)
	switch {
	case len_ < 1<<1:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 1]T{}))
		cap_ = 1 << 1
	case len_ < 1<<2:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 2]T{}))
		cap_ = 1 << 2
	case len_ < 1<<3:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 3]T{}))
		cap_ = 1 << 3
	case len_ < 1<<4:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 4]T{}))
		cap_ = 1 << 4
	case len_ < 1<<5:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 5]T{}))
		cap_ = 1 << 5
	case len_ < 1<<6:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 6]T{}))
		cap_ = 1 << 6
	case len_ < 1<<7:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 7]T{}))
		cap_ = 1 << 7
	case len_ < 1<<8:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 8]T{}))
		cap_ = 1 << 8
	case len_ < 1<<9:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 9]T{}))
		cap_ = 1 << 9
	case len_ < 1<<10:
		arrayPtr = uintptr(unsafe.Pointer(&[1 << 10]T{}))
		cap_ = 1 << 10
	}
	s2 := Slice{
		ArrayPtr: arrayPtr,
		Len:      len_,
		Cap:      cap_,
	}
	result := *(*S)(unsafe.Pointer(&s2))
	copy(result, v)
	copy(result[Len[S, T](v):], elems)
	return result
}

func Len[S ~[]T, T any](v S) int {
	s := *(*Slice)(unsafe.Pointer(&v))
	return s.Len
}

func Cap[S ~[]T, T any](v S) int {
	s := *(*Slice)(unsafe.Pointer(&v))
	return s.Cap
}

