package main

import (
	"errors"
	"fmt"

	"github.com/morikuni/failure"
)

func A() error {
	return B()
}

func B() error {
	err := C()
	if err != nil {
		return failure.Wrap(err)
	}
	return nil
}

func C() error {
	return errors.New("error happened")
}

func main() {
	err := A()
	if err != nil {
		// main.B: error happened
		fmt.Printf("%v\n", err)

		// [main.B] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/wrap/main.go:17
		//    *errors.errorString("error happened")
		// [CallStack]
		//    [main.B] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/wrap/main.go:17
		//    [main.A] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/wrap/main.go:11
		//    [main.main] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/wrap/main.go:27
		//    [runtime.main] /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/proc.go:250
		//    [runtime.goexit] /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/asm_arm64.s:1172
		fmt.Printf("%+v\n", err)
	}
}
