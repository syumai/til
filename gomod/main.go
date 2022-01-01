package main

import (
	"fmt"
	"strings"

	"github.com/syumai/emojidata"
	"github.com/syumai/til/gomod116"
	"github.com/syumai/til/gomod117"
)

func main() {
	fmt.Println(emojidata.Get("star").String())
	gomod116.ReadAll(strings.NewReader("abcde"))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(gomod117.ConvertSliceToArrayOf5Ints(s))
}
