package main

import "fmt"

//go:wasmimport math add
func add(a, b int64) int64

func main() {
	fmt.Println(add(1, 2))
}
