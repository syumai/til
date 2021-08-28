package main

import (
	"fmt"
)

// The playground now uses square brackets for type parameters. Otherwise,
// the syntax of type parameter lists matches the one of regular parameter
// lists except that all type parameters must have a name, and the type
// parameter list cannot be empty. The predeclared identifier "any" may be
// used in the position of a type parameter constraint (and only there);
// it indicates that there are no constraints.

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print([]string{"Hello, ", "playground\n"})
}

