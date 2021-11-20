package main

import (
	"fmt"
	"unsafe"
)

func GetInOut(i interface{}) (in, out uint16) {
	v := *(*[2]uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&i)) + 48))
	return v[0], v[1]
}

func main() {
	var (
		f1 = func(a int) bool { return false }
		f2 = func(a bool) int { return 1 }
		f3 = func(a bool) bool { return false }
		f4 = func(a int) int { return 1 }
		f5 = func(a, b int) int { return 1 }
		f6 = func(a int) (int, int) { return 1, 1 }
	)
	fmt.Println(GetInOut(f1))
	fmt.Println(GetInOut(f2))
	fmt.Println(GetInOut(f3))
	fmt.Println(GetInOut(f4))
	fmt.Println(GetInOut(f5))
	fmt.Println(GetInOut(f6))
}
