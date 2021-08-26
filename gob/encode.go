package main

import (
	"encoding/gob"
	"os"
)

type State struct {
	A int
	B int
}

type StateInterface interface{}

func main() {
	var v StateInterface = State{
		A: 1,
		B: 2,
	}

	f, err := os.OpenFile("state.gob", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	gob.Register(State{})
	err = gob.NewEncoder(f).Encode(&v)
	if err != nil {
		panic(err)
	}
}
