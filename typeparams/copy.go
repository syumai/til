package main

import "fmt"

func CopyInts(a []int) []int {
  b := make([]int, len(a))
  copy(b, a)
  return b
}

func CopyStrs(a []string) []string {
  b := make([]string, len(a))
  copy(b, a)
  return b
}

func SlicesCopy[T any](a []T) []T {
  b := make([]T, len(a))
  copy(b, a)
  return b
}

func main() {
  is := []int{1, 2, 3}
  ss := []string{"a", "b", "c"}

  fmt.Println(CopyInts(is))
  fmt.Println(CopyStrs(ss))

  // with generics
  fmt.Println(SlicesCopy(is))
  fmt.Println(SlicesCopy(ss))
}
