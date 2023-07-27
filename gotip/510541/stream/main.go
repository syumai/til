package main

import (
	"fmt"

	"github.com/syumai/til/gotip/510541/stream/maps"
	"github.com/syumai/til/gotip/510541/stream/slices"
)

func main() {
	m := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}
	keys := slices.Collect(maps.Keys(m))
	values := slices.Collect(maps.Values(m))
	fmt.Println("keys:", keys)
	fmt.Println("values:", values)
}
