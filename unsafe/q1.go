package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// research()
	fmt.Println(IntsLen([]int{1, 2, 3}))
	fmt.Println(IntsLen([]int{1, 2, 3, 4}))
	fmt.Println(StringsLen([]string{"1", "2", "3"}))
	fmt.Println(StringsLen([]string{"1", "2", "3", "4"}))
}

type SliceLen struct {
	_   [8]byte
	Len uint64
}

func IntsLen(v []int) uint64 {
	s := *(*SliceLen)(unsafe.Pointer(&v))
	return s.Len
}

func StringsLen(v []string) uint64 {
	s := *(*SliceLen)(unsafe.Pointer(&v))
	return s.Len
}

func research() {
	s1 := make([]int, 0, 0)
	s2 := make([]int, 0, 1)
	s3 := make([]int, 1, 1)
	s4 := make([]int, 0x1111, 0x2222)

	/* step 1 */
	fmt.Printf("s1: %v\n", *(*[32]byte)(unsafe.Pointer(&s1)))
	fmt.Printf("s2: %v\n", *(*[32]byte)(unsafe.Pointer(&s2)))
	fmt.Printf("s3: %v\n", *(*[32]byte)(unsafe.Pointer(&s3)))
	fmt.Printf("s4: %v\n", *(*[32]byte)(unsafe.Pointer(&s4)))

	/*
		s1: [232 174 7 0 192 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 208 175 7 0 192 0 0 0]
		s2: [240 174 7 0 192 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 232 174 7 0 192 0 0 0]
		s3: [232 174 7 0 192 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 240 174 7 0 192 0 0 0]
		s4: [0 192 16 0 192 0 0 0 17 17 0 0 0 0 0 0 34 34 0 0 0 0 0 0 192 189 5 0 192 0 0 0]
		8 byte - 8 byte (len) - 8 byte (cap)
		先頭 8 byte を読み飛ばして 4 byte 取得すれば OK
	*/
}
