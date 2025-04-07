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

type CounterImpl struct {
	count int
}

func (c *CounterImpl) Increment()    { c.count++ }
func (c *CounterImpl) Decrement()    { c.count-- }
func (c *CounterImpl) GetCount() int { return c.count }

type Proxy[T any] struct {
	Impl T
}

func (p *Proxy[T]) Handle(m reflect.Method, values []reflect.Value) []reflect.Value {
	fmt.Println("Method called:", m.Name)
	if m.Name == "Increment" {
		fmt.Println("Increment is disabled!")
		return nil
	}
	// call original method implementation
	return reflect.ValueOf(p.Impl).MethodByName(m.Name).Call(values)
}

func main() {
	counter := &CounterImpl{count: 0}
	proxyHandler := &Proxy[Counter]{Impl: counter}
	proxy, _ := dyno.Dynamic[Counter](proxyHandler.Handle)
	proxy.Increment()
	proxy.Decrement()
	fmt.Println(proxy.GetCount())
}
