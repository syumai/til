package main

import (
	"fmt"
	"strconv"
)

type SignedInt interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Stringer interface {
	String() string
}

type SignedIntStringer interface {
	SignedInt
	Stringer
}

type I int

func (i I) String() string {
	return strconv.Itoa(int(i))
}

type F float32

func (f F) String() string {
	return fmt.Sprintf("%f", f)
}

func print[T SignedIntStringer](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	print([]I{1, 2, 3})
	//Print([]int{1, 2, 3})
	//Print([]F{1, 2, 3})
	var sis SignedIntStringer
	sis = I(1)
	_ = sis
}

