package main

import (
	"fmt"

	"github.com/morikuni/failure"
	"github.com/syumai/til/go/errors/failure_example/codes"
)

func A() error {
	return B()
}

func B() error {
	return failure.New(codes.Internal, failure.Message("internal error happened"))
}

func main() {
	err := A()
	if err != nil {
		// main.B: internal error happened: code(Internal)
		fmt.Printf("%v\n", err)

		// [main.B] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/basic/main.go:15
		//    message("internal error happened")
		//    code(Internal)
		// [CallStack]
		//    [main.B] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/basic/main.go:15
		//    [main.A] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/basic/main.go:11
		//    [main.main] /Users/syumai/go/src/github.com/syumai/til/go/errors/failure_example/basic/main.go:19
		//    [runtime.main] /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/proc.go:250
		//    [runtime.goexit] /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/asm_arm64.s:1172
		fmt.Printf("%+v\n", err)
	}
}
