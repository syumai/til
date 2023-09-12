package main

import (
	"fmt"
	"os"
)

func main() {
	sp := string(rune(0xfeff))
	code := fmt.Sprint("const" + sp + "a" + sp + "=" + sp + "1;\n" + "console.log(a);\n")
	f, _ := os.Create("whitespace.js")
	defer f.Close()
	f.Write([]byte(code))
}
