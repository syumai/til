// WIP

package main

import "fmt"

type Signed interface {
	type int, int8, int16, int32, int64 // type list
}

func printSignedSlice[T Signed](i T) {
	fmt.Println(i)
}

func main() {
	printSigned(1) // Function type inference
}

