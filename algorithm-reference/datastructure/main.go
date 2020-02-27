package main

import (
	"fmt"

	"github.com/syumai/til/algorithm-reference/datastructure/list"
)

func main() {
	a := list.LinearListNode{
		Value: 1,
	}

	b := list.LinearListNode{
		Value: 2,
	}
	a.Next = &b

	l := list.LinearList{
		FirstNode: &a,
	}

	l.Add(1, 3)

	for _, n := range l.ToSlice() {
		fmt.Println(n.Value)
	}
}
