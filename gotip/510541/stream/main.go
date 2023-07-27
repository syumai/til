package main

import (
	"fmt"
	"strings"

	"github.com/syumai/til/gotip/510541/stream/maps"
	"github.com/syumai/til/gotip/510541/stream/stream"
)

func main() {
	m := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}
	upperValues := stream.Map(maps.Values(m), func(s string) string {
		return strings.ToUpper(s)
	})
	for v := range upperValues {
		fmt.Println(v)
	}
}
