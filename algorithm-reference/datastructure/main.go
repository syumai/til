package main

import (
	"fmt"

	"github.com/syumai/til/algorithm-reference/datastructure/list"
)

func main() {
	l := list.LinearListNode{
		Value: 1,
	}

	a := list.LinearListNode{
		Value: 2,
	}
	l.Next = &a

	l.Add(1, 3)

	for _, n := range l.ToSlice() {
		fmt.Println(n.Value)
	}
}
