package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type State struct {
	A int
	B int
	C int
}

type OtherStateInterface interface{}

func main() {
	f, err := os.Open("state.gob")
	if err != nil {
		panic(err)
	}

	var v OtherStateInterface
	gob.Register(State{})
	err = gob.NewDecoder(f).Decode(&v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", v)
}
