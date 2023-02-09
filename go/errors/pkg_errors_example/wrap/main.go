package main

import (
	"errors"
	"fmt"

	pkg_errors "github.com/pkg/errors"
)

func A() error {
	return B()
}

func B() error {
	err := C()
	if err != nil {
		return pkg_errors.Wrap(err, "got error")
	}
	return nil
}

func C() error {
	return errors.New("error happened")
}

func main() {
	err := A()
	if err != nil {
		// got error: error happened
		fmt.Printf("%v\n", err)

		// error happened
		// got error
		// main.B
		//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/wrap/main.go:17
		// main.A
		//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/wrap/main.go:11
		// main.main
		//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/wrap/main.go:27
		// runtime.main
		//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/proc.go:250
		// runtime.goexit
		//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/asm_arm64.s:1172
		fmt.Printf("%+v\n", err)
	}
}
