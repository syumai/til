package main

import (
	"fmt"

	pkg_errors "github.com/pkg/errors"
)

func A() error {
	return B()
}

func B() error {
	return pkg_errors.New("error happened")
}

func main() {
	err := A()
	if err != nil {
		// error happened
		fmt.Printf("%v\n", err)

		// error happened
		// main.B
		//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/basic/main.go:14
		// main.A
		//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/basic/main.go:10
		// main.main
		//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/basic/main.go:18
		// runtime.main
		//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/proc.go:250
		// runtime.goexit
		//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/asm_arm64.s:1172
		fmt.Printf("%+v\n", err)
	}
}
