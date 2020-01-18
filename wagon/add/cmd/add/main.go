package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/syumai/til/wagon/add"
)

func main() {
	os.Exit(realMain(os.Stdout, os.Args))
}

func realMain(w io.Writer, args []string) int {
	if len(args) < 3 {
		fmt.Printf("args must be given\n")
		return 1
	}

	var intArgs []int
	for _, s := range args[1:] {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("args must be integer\n")
			return 1
		}
		intArgs = append(intArgs, i)
	}
	result, err := add.Add(intArgs[0], intArgs[1])
	if err != nil {
		fmt.Printf("unexpected error: %v\n", err)
		return 1
	}
	fmt.Println(result)
	return 0
}
