package main

import (
	"fmt"
	"reflect"

	"github.com/ovechkin-dm/go-dyno/pkg/dyno"
)

type Counter interface {
	Increment()
	Decrement()
	GetCount() int
}

var count int

func CounterHandler(m reflect.Method, values []reflect.Value) []reflect.Value {
	fmt.Println("Method called: ", m.Name)
	switch m.Name {
	case "Increment":
		count++
	case "Decrement":
		count--
	case "GetCount":
		return []reflect.Value{reflect.ValueOf(count)}
	}
	return nil
}

func main() {
	dynamicCounter, _ := dyno.Dynamic[Counter](CounterHandler)
	dynamicCounter.Increment()
	dynamicCounter.Increment()
	dynamicCounter.Decrement()
	fmt.Println(dynamicCounter.GetCount())
}
