package main

import (
  "fmt"
)

type Chan[Elem any] interface { ~chan Elem }

func FlushSlice[C Chan[T], T any](ch C) []T {
  result := make([]T, 0, cap(ch))
Loop:
  for {
    select {
      case v, ok := <-ch:
        if !ok {
          break Loop
        }
		result = append(result, v)
      default:
        break Loop
    }
  }
  return result
}

func main() {
  intCh := make(chan int, 3)
  intCh <- 1; intCh <- 2; intCh <- 3

  var intRecvCh <-chan int = intCh
  fmt.Println(FlushSlice(intRecvCh))
}

