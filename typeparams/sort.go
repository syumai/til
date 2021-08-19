package main

import (
	"fmt"
	"sort"
)

type Slice[Elem any] interface { ~[]Elem }

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func SortSlice[S Slice[T], T Ordered](s S) {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

func main() {
  ints := []int{3,1,4,2,5}
  SortSlice(ints)
  fmt.Println(ints) // []int{1,2,3,4,5}
  strs := []string{"c", "b", "a"}
  SortSlice(strs)
  fmt.Println(strs) // []string{"a", "b", "c"}
}

