package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type State struct {
	A int
	B int
}

type StateInterface interface{}

func main() {
	f, err := os.Open("state.gob")
	if err != nil {
		panic(err)
	}

	var v StateInterface
	gob.Register(State{})
	err = gob.NewDecoder(f).Decode(&v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", v)
}
